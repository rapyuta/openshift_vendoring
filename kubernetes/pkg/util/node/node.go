/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package node

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/openshift/kubernetes/pkg/api"
	"github.com/openshift/kubernetes/pkg/api/unversioned"
	clientset "github.com/openshift/kubernetes/pkg/client/clientset_generated/internalclientset"
	"github.com/openshift/kubernetes/pkg/types"
)

const (
	// The reason and message set on a pod when its state cannot be confirmed as kubelet is unresponsive
	// on the node it is (was) running.
	NodeUnreachablePodReason  = "NodeLost"
	NodeUnreachablePodMessage = "Node %v which was running pod %v is unresponsive"
)

func GetHostname(hostnameOverride string) string {
	var hostname string = hostnameOverride
	if hostname == "" {
		nodename, err := os.Hostname()
		if err != nil {
			glog.Fatalf("Couldn't determine hostname: %v", err)
		}
		hostname = nodename
	}
	return strings.ToLower(strings.TrimSpace(hostname))
}

// GetPreferredNodeAddress returns the address of the provided node, using the provided preference order.
// If none of the preferred address types are found, an error is returned.
func GetPreferredNodeAddress(node *api.Node, preferredAddressTypes []api.NodeAddressType) (string, error) {
	for _, addressType := range preferredAddressTypes {
		for _, address := range node.Status.Addresses {
			if address.Type == addressType {
				return address.Address, nil
			}
		}
		// If hostname was requested and no Hostname address was registered...
		if addressType == api.NodeHostName {
			// ...fall back to the kubernetes.io/hostname label for compatibility with kubelets before 1.5
			if hostname, ok := node.Labels[unversioned.LabelHostname]; ok && len(hostname) > 0 {
				return hostname, nil
			}
		}
	}
	return "", fmt.Errorf("no preferred addresses found; known addresses: %v", node.Status.Addresses)
}

// GetNodeHostIP returns the provided node's IP, based on the priority:
// 1. NodeInternalIP
// 2. NodeExternalIP
// 3. NodeLegacyHostIP
func GetNodeHostIP(node *api.Node) (net.IP, error) {
	addresses := node.Status.Addresses
	addressMap := make(map[api.NodeAddressType][]api.NodeAddress)
	for i := range addresses {
		addressMap[addresses[i].Type] = append(addressMap[addresses[i].Type], addresses[i])
	}
	if addresses, ok := addressMap[api.NodeInternalIP]; ok {
		return net.ParseIP(addresses[0].Address), nil
	}
	if addresses, ok := addressMap[api.NodeExternalIP]; ok {
		return net.ParseIP(addresses[0].Address), nil
	}
	if addresses, ok := addressMap[api.NodeLegacyHostIP]; ok {
		return net.ParseIP(addresses[0].Address), nil
	}
	return nil, fmt.Errorf("host IP unknown; known addresses: %v", addresses)
}

// Helper function that builds a string identifier that is unique per failure-zone
// Returns empty-string for no zone
func GetZoneKey(node *api.Node) string {
	labels := node.Labels
	if labels == nil {
		return ""
	}

	region, _ := labels[unversioned.LabelZoneRegion]
	failureDomain, _ := labels[unversioned.LabelZoneFailureDomain]

	if region == "" && failureDomain == "" {
		return ""
	}

	// We include the null character just in case region or failureDomain has a colon
	// (We do assume there's no null characters in a region or failureDomain)
	// As a nice side-benefit, the null character is not printed by fmt.Print or glog
	return region + ":\x00:" + failureDomain
}

// SetNodeCondition updates specific node condition with patch operation.
func SetNodeCondition(c clientset.Interface, node types.NodeName, condition api.NodeCondition) error {
	generatePatch := func(condition api.NodeCondition) ([]byte, error) {
		raw, err := json.Marshal(&[]api.NodeCondition{condition})
		if err != nil {
			return nil, err
		}
		return []byte(fmt.Sprintf(`{"status":{"conditions":%s}}`, raw)), nil
	}
	condition.LastHeartbeatTime = unversioned.NewTime(time.Now())
	patch, err := generatePatch(condition)
	if err != nil {
		return nil
	}
	_, err = c.Core().Nodes().PatchStatus(string(node), patch)
	return err
}

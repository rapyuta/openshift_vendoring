package tenants

import "github.com/openshift/github.com/rackspace/gophercloud"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("tenants")
}

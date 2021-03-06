/*
Copyright 2016 The Kubernetes Authors.

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

package e2e

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/openshift/kubernetes/pkg/api"
	"github.com/openshift/kubernetes/pkg/labels"
	"github.com/openshift/kubernetes/pkg/util/uuid"
	"github.com/openshift/kubernetes/pkg/util/wait"
	"github.com/openshift/kubernetes/pkg/watch"
	"github.com/openshift/kubernetes/test/e2e/framework"

	. "github.com/openshift/github.com/onsi/ginkgo"
	. "github.com/openshift/github.com/onsi/gomega"
)

var _ = framework.KubeDescribe("Pods Delete Grace Period", func() {
	f := framework.NewDefaultFramework("pods")
	var podClient *framework.PodClient
	BeforeEach(func() {
		podClient = f.PodClient()
	})
	It("should be submitted and removed [Conformance]", func() {
		By("creating the pod")
		name := "pod-submit-remove-" + string(uuid.NewUUID())
		value := strconv.Itoa(time.Now().Nanosecond())
		pod := &api.Pod{
			ObjectMeta: api.ObjectMeta{
				Name: name,
				Labels: map[string]string{
					"name": "foo",
					"time": value,
				},
			},
			Spec: api.PodSpec{
				Containers: []api.Container{
					{
						Name:  "nginx",
						Image: "gcr.io/google_containers/nginx-slim:0.7",
					},
				},
			},
		}

		By("setting up watch")
		selector := labels.SelectorFromSet(labels.Set(map[string]string{"time": value}))
		options := api.ListOptions{LabelSelector: selector}
		pods, err := podClient.List(options)
		Expect(err).NotTo(HaveOccurred(), "failed to query for pod")
		Expect(len(pods.Items)).To(Equal(0))
		options = api.ListOptions{
			LabelSelector:   selector,
			ResourceVersion: pods.ListMeta.ResourceVersion,
		}
		w, err := podClient.Watch(options)
		Expect(err).NotTo(HaveOccurred(), "failed to set up watch")

		By("submitting the pod to kubernetes")
		podClient.Create(pod)

		By("verifying the pod is in kubernetes")
		selector = labels.SelectorFromSet(labels.Set(map[string]string{"time": value}))
		options = api.ListOptions{LabelSelector: selector}
		pods, err = podClient.List(options)
		Expect(err).NotTo(HaveOccurred(), "failed to query for pod")
		Expect(len(pods.Items)).To(Equal(1))

		By("verifying pod creation was observed")
		select {
		case event, _ := <-w.ResultChan():
			if event.Type != watch.Added {
				framework.Failf("Failed to observe pod creation: %v", event)
			}
		case <-time.After(framework.PodStartTimeout):
			Fail("Timeout while waiting for pod creation")
		}

		// We need to wait for the pod to be running, otherwise the deletion
		// may be carried out immediately rather than gracefully.
		framework.ExpectNoError(f.WaitForPodRunning(pod.Name))
		// save the running pod
		pod, err = podClient.Get(pod.Name)
		Expect(err).NotTo(HaveOccurred(), "failed to GET scheduled pod")

		// start local proxy, so we can send graceful deletion over query string, rather than body parameter
		cmd := framework.KubectlCmd("proxy", "-p", "0")
		stdout, stderr, err := framework.StartCmdAndStreamOutput(cmd)
		Expect(err).NotTo(HaveOccurred(), "failed to start up proxy")
		defer stdout.Close()
		defer stderr.Close()
		defer framework.TryKill(cmd)
		buf := make([]byte, 128)
		var n int
		n, err = stdout.Read(buf)
		Expect(err).NotTo(HaveOccurred(), "failed to read from kubectl proxy stdout")
		output := string(buf[:n])
		proxyRegexp := regexp.MustCompile("Starting to serve on 127.0.0.1:([0-9]+)")
		match := proxyRegexp.FindStringSubmatch(output)
		Expect(len(match)).To(Equal(2))
		port, err := strconv.Atoi(match[1])
		Expect(err).NotTo(HaveOccurred(), "failed to convert port into string")

		endpoint := fmt.Sprintf("http://localhost:%d/api/v1/namespaces/%s/pods/%s?gracePeriodSeconds=30", port, pod.Namespace, pod.Name)
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		req, err := http.NewRequest("DELETE", endpoint, nil)
		Expect(err).NotTo(HaveOccurred(), "failed to create http request")

		By("deleting the pod gracefully")
		rsp, err := client.Do(req)
		Expect(err).NotTo(HaveOccurred(), "failed to use http client to send delete")

		defer rsp.Body.Close()

		By("verifying the kubelet observed the termination notice")
		Expect(wait.Poll(time.Second*5, time.Second*30, func() (bool, error) {
			podList, err := framework.GetKubeletPods(f.ClientSet, pod.Spec.NodeName)
			if err != nil {
				framework.Logf("Unable to retrieve kubelet pods for node %v: %v", pod.Spec.NodeName, err)
				return false, nil
			}
			for _, kubeletPod := range podList.Items {
				if pod.Name != kubeletPod.Name {
					continue
				}
				if kubeletPod.ObjectMeta.DeletionTimestamp == nil {
					framework.Logf("deletion has not yet been observed")
					return false, nil
				}
				return true, nil
			}
			framework.Logf("no pod exists with the name we were looking for, assuming the termination request was observed and completed")
			return true, nil
		})).NotTo(HaveOccurred(), "kubelet never observed the termination notice")

		By("verifying pod deletion was observed")
		deleted := false
		timeout := false
		var lastPod *api.Pod
		timer := time.After(2 * time.Minute)
		for !deleted && !timeout {
			select {
			case event, _ := <-w.ResultChan():
				if event.Type == watch.Deleted {
					lastPod = event.Object.(*api.Pod)
					deleted = true
				}
			case <-timer:
				timeout = true
			}
		}
		if !deleted {
			Fail("Failed to observe pod deletion")
		}

		Expect(lastPod.DeletionTimestamp).ToNot(BeNil())
		Expect(lastPod.Spec.TerminationGracePeriodSeconds).ToNot(BeZero())

		selector = labels.SelectorFromSet(labels.Set(map[string]string{"time": value}))
		options = api.ListOptions{LabelSelector: selector}
		pods, err = podClient.List(options)
		Expect(err).NotTo(HaveOccurred(), "failed to query for pods")
		Expect(len(pods.Items)).To(Equal(0))

	})
})

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

package e2e

import (
	"github.com/openshift/kubernetes/pkg/api"
	"github.com/openshift/kubernetes/pkg/api/resource"
	"github.com/openshift/kubernetes/pkg/util/intstr"
	"github.com/openshift/kubernetes/pkg/util/uuid"
	"github.com/openshift/kubernetes/test/e2e/framework"

	"fmt"
	"strconv"

	. "github.com/openshift/github.com/onsi/ginkgo"
	. "github.com/openshift/github.com/onsi/gomega"
)

const (
	// These numbers are obtained empirically.
	// If you make them too low, you'll get flaky
	// tests instead of failing ones if the race bug reappears.
	// If you make volume counts or pod counts too high,
	// the tests may fail because mounting configmap/git_repo
	// volumes is not very fast and the tests may time out
	// waiting for pods to become Running.
	// And of course the higher are the numbers, the
	// slower are the tests.
	wrappedVolumeRaceConfigMapVolumeCount    = 50
	wrappedVolumeRaceConfigMapPodCount       = 5
	wrappedVolumeRaceConfigMapIterationCount = 3
	wrappedVolumeRaceGitRepoVolumeCount      = 50
	wrappedVolumeRaceGitRepoPodCount         = 5
	wrappedVolumeRaceGitRepoIterationCount   = 3
	wrappedVolumeRaceRCNamePrefix            = "wrapped-volume-race-"
)

var _ = framework.KubeDescribe("EmptyDir wrapper volumes", func() {
	f := framework.NewDefaultFramework("emptydir-wrapper")

	It("should not conflict", func() {
		name := "emptydir-wrapper-test-" + string(uuid.NewUUID())
		volumeName := "secret-volume"
		volumeMountPath := "/etc/secret-volume"

		secret := &api.Secret{
			ObjectMeta: api.ObjectMeta{
				Namespace: f.Namespace.Name,
				Name:      name,
			},
			Data: map[string][]byte{
				"data-1": []byte("value-1\n"),
			},
		}

		var err error
		if secret, err = f.ClientSet.Core().Secrets(f.Namespace.Name).Create(secret); err != nil {
			framework.Failf("unable to create test secret %s: %v", secret.Name, err)
		}

		gitVolumeName := "git-volume"
		gitVolumeMountPath := "/etc/git-volume"
		gitURL, gitRepo, gitCleanup := createGitServer(f)
		defer gitCleanup()

		pod := &api.Pod{
			ObjectMeta: api.ObjectMeta{
				Name: "pod-secrets-" + string(uuid.NewUUID()),
			},
			Spec: api.PodSpec{
				Volumes: []api.Volume{
					{
						Name: volumeName,
						VolumeSource: api.VolumeSource{
							Secret: &api.SecretVolumeSource{
								SecretName: name,
							},
						},
					},
					{
						Name: gitVolumeName,
						VolumeSource: api.VolumeSource{
							GitRepo: &api.GitRepoVolumeSource{
								Repository: gitURL,
								Directory:  gitRepo,
							},
						},
					},
				},
				Containers: []api.Container{
					{
						Name:  "secret-test",
						Image: "gcr.io/google_containers/test-webserver:e2e",
						VolumeMounts: []api.VolumeMount{
							{
								Name:      volumeName,
								MountPath: volumeMountPath,
								ReadOnly:  true,
							},
							{
								Name:      gitVolumeName,
								MountPath: gitVolumeMountPath,
							},
						},
					},
				},
			},
		}
		pod = f.PodClient().CreateSync(pod)

		defer func() {
			By("Cleaning up the secret")
			if err := f.ClientSet.Core().Secrets(f.Namespace.Name).Delete(secret.Name, nil); err != nil {
				framework.Failf("unable to delete secret %v: %v", secret.Name, err)
			}
			By("Cleaning up the git vol pod")
			if err = f.ClientSet.Core().Pods(f.Namespace.Name).Delete(pod.Name, api.NewDeleteOptions(0)); err != nil {
				framework.Failf("unable to delete git vol pod %v: %v", pod.Name, err)
			}
		}()
	})

	// The following two tests check for the problem fixed in #29641.
	// In order to reproduce it you need to revert the fix, e.g. via
	// git revert -n df1e925143daf34199b55ffb91d0598244888cce
	// or
	// curl -sL https://github.com/kubernetes/kubernetes/pull/29641.patch | patch -p1 -R
	//
	// After that these tests will fail because some of the pods
	// they create never enter Running state.
	//
	// They need to be [Serial] and [Slow] because they try to induce
	// the race by creating pods with many volumes and container volume mounts,
	// which takes considerable time and may interfere with other tests.
	//
	// Probably should also try making tests for secrets and downwardapi,
	// but these cases are harder because tmpfs-based emptyDir
	// appears to be less prone to the race problem.

	It("should not cause race condition when used for configmaps [Serial] [Slow]", func() {
		configMapNames := createConfigmapsForRace(f)
		defer deleteConfigMaps(f, configMapNames)
		volumes, volumeMounts := makeConfigMapVolumes(configMapNames)
		for i := 0; i < wrappedVolumeRaceConfigMapIterationCount; i++ {
			testNoWrappedVolumeRace(f, volumes, volumeMounts, wrappedVolumeRaceConfigMapPodCount)
		}
	})

	It("should not cause race condition when used for git_repo [Serial] [Slow]", func() {
		gitURL, gitRepo, cleanup := createGitServer(f)
		defer cleanup()
		volumes, volumeMounts := makeGitRepoVolumes(gitURL, gitRepo)
		for i := 0; i < wrappedVolumeRaceGitRepoIterationCount; i++ {
			testNoWrappedVolumeRace(f, volumes, volumeMounts, wrappedVolumeRaceGitRepoPodCount)
		}
	})
})

func createGitServer(f *framework.Framework) (gitURL string, gitRepo string, cleanup func()) {
	var err error
	gitServerPodName := "git-server-" + string(uuid.NewUUID())
	containerPort := 8000

	labels := map[string]string{"name": gitServerPodName}

	gitServerPod := &api.Pod{
		ObjectMeta: api.ObjectMeta{
			Name:   gitServerPodName,
			Labels: labels,
		},
		Spec: api.PodSpec{
			Containers: []api.Container{
				{
					Name:            "git-repo",
					Image:           "gcr.io/google_containers/fakegitserver:0.1",
					ImagePullPolicy: "IfNotPresent",
					Ports: []api.ContainerPort{
						{ContainerPort: int32(containerPort)},
					},
				},
			},
		},
	}
	f.PodClient().CreateSync(gitServerPod)

	// Portal IP and port
	httpPort := 2345

	gitServerSvc := &api.Service{
		ObjectMeta: api.ObjectMeta{
			Name: "git-server-svc",
		},
		Spec: api.ServiceSpec{
			Selector: labels,
			Ports: []api.ServicePort{
				{
					Name:       "http-portal",
					Port:       int32(httpPort),
					TargetPort: intstr.FromInt(containerPort),
				},
			},
		},
	}

	if gitServerSvc, err = f.ClientSet.Core().Services(f.Namespace.Name).Create(gitServerSvc); err != nil {
		framework.Failf("unable to create test git server service %s: %v", gitServerSvc.Name, err)
	}

	return "http://" + gitServerSvc.Spec.ClusterIP + ":" + strconv.Itoa(httpPort), "test", func() {
		By("Cleaning up the git server pod")
		if err := f.ClientSet.Core().Pods(f.Namespace.Name).Delete(gitServerPod.Name, api.NewDeleteOptions(0)); err != nil {
			framework.Failf("unable to delete git server pod %v: %v", gitServerPod.Name, err)
		}
		By("Cleaning up the git server svc")
		if err := f.ClientSet.Core().Services(f.Namespace.Name).Delete(gitServerSvc.Name, nil); err != nil {
			framework.Failf("unable to delete git server svc %v: %v", gitServerSvc.Name, err)
		}
	}
}

func makeGitRepoVolumes(gitURL, gitRepo string) (volumes []api.Volume, volumeMounts []api.VolumeMount) {
	for i := 0; i < wrappedVolumeRaceGitRepoVolumeCount; i++ {
		volumeName := fmt.Sprintf("racey-git-repo-%d", i)
		volumes = append(volumes, api.Volume{
			Name: volumeName,
			VolumeSource: api.VolumeSource{
				GitRepo: &api.GitRepoVolumeSource{
					Repository: gitURL,
					Directory:  gitRepo,
				},
			},
		})
		volumeMounts = append(volumeMounts, api.VolumeMount{
			Name:      volumeName,
			MountPath: fmt.Sprintf("/etc/git-volume-%d", i),
		})
	}
	return
}

func createConfigmapsForRace(f *framework.Framework) (configMapNames []string) {
	By(fmt.Sprintf("Creating %d configmaps", wrappedVolumeRaceConfigMapVolumeCount))
	for i := 0; i < wrappedVolumeRaceConfigMapVolumeCount; i++ {
		configMapName := fmt.Sprintf("racey-configmap-%d", i)
		configMapNames = append(configMapNames, configMapName)
		configMap := &api.ConfigMap{
			ObjectMeta: api.ObjectMeta{
				Namespace: f.Namespace.Name,
				Name:      configMapName,
			},
			Data: map[string]string{
				"data-1": "value-1",
			},
		}
		_, err := f.ClientSet.Core().ConfigMaps(f.Namespace.Name).Create(configMap)
		framework.ExpectNoError(err)
	}
	return
}

func deleteConfigMaps(f *framework.Framework, configMapNames []string) {
	By("Cleaning up the configMaps")
	for _, configMapName := range configMapNames {
		err := f.ClientSet.Core().ConfigMaps(f.Namespace.Name).Delete(configMapName, nil)
		Expect(err).NotTo(HaveOccurred(), "unable to delete configMap %v", configMapName)
	}
}

func makeConfigMapVolumes(configMapNames []string) (volumes []api.Volume, volumeMounts []api.VolumeMount) {
	for i, configMapName := range configMapNames {
		volumeName := fmt.Sprintf("racey-configmap-%d", i)
		volumes = append(volumes, api.Volume{
			Name: volumeName,
			VolumeSource: api.VolumeSource{
				ConfigMap: &api.ConfigMapVolumeSource{
					LocalObjectReference: api.LocalObjectReference{
						Name: configMapName,
					},
					Items: []api.KeyToPath{
						{
							Key:  "data-1",
							Path: "data-1",
						},
					},
				},
			},
		})
		volumeMounts = append(volumeMounts, api.VolumeMount{
			Name:      volumeName,
			MountPath: fmt.Sprintf("/etc/config-%d", i),
		})
	}
	return
}

func testNoWrappedVolumeRace(f *framework.Framework, volumes []api.Volume, volumeMounts []api.VolumeMount, podCount int32) {
	rcName := wrappedVolumeRaceRCNamePrefix + string(uuid.NewUUID())
	nodeList := framework.GetReadySchedulableNodesOrDie(f.ClientSet)
	Expect(len(nodeList.Items)).To(BeNumerically(">", 0))
	targetNode := nodeList.Items[0]

	By("Creating RC which spawns configmap-volume pods")
	affinity := map[string]string{
		api.AffinityAnnotationKey: fmt.Sprintf(`
				{"nodeAffinity": { "requiredDuringSchedulingIgnoredDuringExecution": {
					"nodeSelectorTerms": [{
						"matchExpressions": [{
							"key": "kubernetes.io/hostname",
							"operator": "In",
							"values": ["%s"]
					}]
				}]
			}}}`, targetNode.Name),
	}

	rc := &api.ReplicationController{
		ObjectMeta: api.ObjectMeta{
			Name: rcName,
		},
		Spec: api.ReplicationControllerSpec{
			Replicas: podCount,
			Selector: map[string]string{
				"name": rcName,
			},
			Template: &api.PodTemplateSpec{
				ObjectMeta: api.ObjectMeta{
					Annotations: affinity,
					Labels:      map[string]string{"name": rcName},
				},
				Spec: api.PodSpec{
					Containers: []api.Container{
						{
							Name:    "test-container",
							Image:   "gcr.io/google_containers/busybox:1.24",
							Command: []string{"sleep", "10000"},
							Resources: api.ResourceRequirements{
								Requests: api.ResourceList{
									api.ResourceCPU: resource.MustParse("10m"),
								},
							},
							VolumeMounts: volumeMounts,
						},
					},
					DNSPolicy: api.DNSDefault,
					Volumes:   volumes,
				},
			},
		},
	}
	_, err := f.ClientSet.Core().ReplicationControllers(f.Namespace.Name).Create(rc)
	Expect(err).NotTo(HaveOccurred(), "error creating replication controller")

	defer func() {
		err := framework.DeleteRCAndPods(f.ClientSet, f.Namespace.Name, rcName)
		framework.ExpectNoError(err)
	}()

	pods, err := framework.PodsCreated(f.ClientSet, f.Namespace.Name, rcName, podCount)

	By("Ensuring each pod is running")

	// Wait for the pods to enter the running state. Waiting loops until the pods
	// are running so non-running pods cause a timeout for this test.
	for _, pod := range pods.Items {
		if pod.DeletionTimestamp != nil {
			continue
		}
		err = f.WaitForPodRunning(pod.Name)
		Expect(err).NotTo(HaveOccurred(), "Failed waiting for pod %s to enter running state", pod.Name)
	}
}

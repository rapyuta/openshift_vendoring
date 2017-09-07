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

package metrics

import (
	"encoding/json"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/openshift/kubernetes/pkg/api"
	"github.com/openshift/kubernetes/pkg/api/resource"
	"github.com/openshift/kubernetes/pkg/api/unversioned"
	"github.com/openshift/kubernetes/pkg/api/v1"
	_ "github.com/openshift/kubernetes/pkg/apimachinery/registered"
	"github.com/openshift/kubernetes/pkg/client/clientset_generated/internalclientset/fake"
	"github.com/openshift/kubernetes/pkg/client/restclient"
	"github.com/openshift/kubernetes/pkg/client/testing/core"
	"github.com/openshift/kubernetes/pkg/labels"
	"github.com/openshift/kubernetes/pkg/runtime"

	heapster "github.com/openshift/k8s.io/heapster/metrics/api/v1/types"
	metrics_api "github.com/openshift/k8s.io/heapster/metrics/apis/metrics/v1alpha1"

	"github.com/openshift/github.com/stretchr/testify/assert"
)

var fixedTimestamp = time.Date(2015, time.November, 10, 12, 30, 0, 0, time.UTC)

func (w fakeResponseWrapper) DoRaw() ([]byte, error) {
	return w.raw, nil
}

func (w fakeResponseWrapper) Stream() (io.ReadCloser, error) {
	return nil, nil
}

func newFakeResponseWrapper(raw []byte) fakeResponseWrapper {
	return fakeResponseWrapper{raw: raw}
}

type fakeResponseWrapper struct {
	raw []byte
}

// timestamp is used for establishing order on metricPoints
type metricPoint struct {
	level     uint64
	timestamp int
}

type testCase struct {
	desiredResourceValues PodResourceInfo
	desiredMetricValues   PodMetricsInfo
	desiredError          error

	replicas              int
	targetTimestamp       int
	reportedMetricsPoints [][]metricPoint
	reportedPodMetrics    [][]int64

	namespace    string
	selector     labels.Selector
	resourceName api.ResourceName
	metricName   string
}

func (tc *testCase) prepareTestClient(t *testing.T) *fake.Clientset {
	namespace := "test-namespace"
	tc.namespace = namespace
	podNamePrefix := "test-pod"
	podLabels := map[string]string{"name": podNamePrefix}
	tc.selector = labels.SelectorFromSet(podLabels)

	// it's a resource test if we have a resource name
	isResource := len(tc.resourceName) > 0

	fakeClient := &fake.Clientset{}

	fakeClient.AddReactor("list", "pods", func(action core.Action) (handled bool, ret runtime.Object, err error) {
		obj := &api.PodList{}
		for i := 0; i < tc.replicas; i++ {
			podName := fmt.Sprintf("%s-%d", podNamePrefix, i)
			pod := buildPod(namespace, podName, podLabels, api.PodRunning, "1024")
			obj.Items = append(obj.Items, pod)
		}
		return true, obj, nil
	})

	if isResource {
		fakeClient.AddProxyReactor("services", func(action core.Action) (handled bool, ret restclient.ResponseWrapper, err error) {
			metrics := metrics_api.PodMetricsList{}
			for i, containers := range tc.reportedPodMetrics {
				metric := metrics_api.PodMetrics{
					ObjectMeta: v1.ObjectMeta{
						Name:      fmt.Sprintf("%s-%d", podNamePrefix, i),
						Namespace: namespace,
					},
					Timestamp:  unversioned.Time{Time: fixedTimestamp.Add(time.Duration(tc.targetTimestamp) * time.Minute)},
					Containers: []metrics_api.ContainerMetrics{},
				}
				for j, cpu := range containers {
					cm := metrics_api.ContainerMetrics{
						Name: fmt.Sprintf("%s-%d-container-%d", podNamePrefix, i, j),
						Usage: v1.ResourceList{
							v1.ResourceCPU: *resource.NewMilliQuantity(
								cpu,
								resource.DecimalSI),
							v1.ResourceMemory: *resource.NewQuantity(
								int64(1024*1024),
								resource.BinarySI),
						},
					}
					metric.Containers = append(metric.Containers, cm)
				}
				metrics.Items = append(metrics.Items, metric)
			}
			heapsterRawMemResponse, _ := json.Marshal(&metrics)
			return true, newFakeResponseWrapper(heapsterRawMemResponse), nil
		})
	} else {
		fakeClient.AddProxyReactor("services", func(action core.Action) (handled bool, ret restclient.ResponseWrapper, err error) {
			metrics := heapster.MetricResultList{}
			var latestTimestamp time.Time
			for _, reportedMetricPoints := range tc.reportedMetricsPoints {
				var heapsterMetricPoints []heapster.MetricPoint
				for _, reportedMetricPoint := range reportedMetricPoints {
					timestamp := fixedTimestamp.Add(time.Duration(reportedMetricPoint.timestamp) * time.Minute)
					if latestTimestamp.Before(timestamp) {
						latestTimestamp = timestamp
					}
					heapsterMetricPoint := heapster.MetricPoint{Timestamp: timestamp, Value: reportedMetricPoint.level, FloatValue: nil}
					heapsterMetricPoints = append(heapsterMetricPoints, heapsterMetricPoint)
				}
				metric := heapster.MetricResult{
					Metrics:         heapsterMetricPoints,
					LatestTimestamp: latestTimestamp,
				}
				metrics.Items = append(metrics.Items, metric)
			}
			heapsterRawMemResponse, _ := json.Marshal(&metrics)
			return true, newFakeResponseWrapper(heapsterRawMemResponse), nil
		})
	}

	return fakeClient
}

func buildPod(namespace, podName string, podLabels map[string]string, phase api.PodPhase, request string) api.Pod {
	return api.Pod{
		ObjectMeta: api.ObjectMeta{
			Name:      podName,
			Namespace: namespace,
			Labels:    podLabels,
		},
		Spec: api.PodSpec{
			Containers: []api.Container{
				{
					Resources: api.ResourceRequirements{
						Requests: api.ResourceList{
							api.ResourceCPU: resource.MustParse(request),
						},
					},
				},
			},
		},
		Status: api.PodStatus{
			Phase: phase,
			Conditions: []api.PodCondition{
				{
					Type:   api.PodReady,
					Status: api.ConditionTrue,
				},
			},
		},
	}
}

func (tc *testCase) verifyResults(t *testing.T, metrics interface{}, timestamp time.Time, err error) {
	if tc.desiredError != nil {
		assert.Error(t, err, "there should be an error retrieving the metrics")
		assert.Contains(t, fmt.Sprintf("%v", err), fmt.Sprintf("%v", tc.desiredError), "the error message should be eas expected")
		return
	}
	assert.NoError(t, err, "there should be no error retrieving the metrics")
	assert.NotNil(t, metrics, "there should be metrics returned")

	if metricsInfo, wasRaw := metrics.(PodMetricsInfo); wasRaw {
		assert.Equal(t, tc.desiredMetricValues, metricsInfo, "the raw metrics values should be as expected")
	} else if resourceInfo, wasResource := metrics.(PodResourceInfo); wasResource {
		assert.Equal(t, tc.desiredResourceValues, resourceInfo, "the resource metrics values be been as expected")
	} else {
		assert.False(t, true, "should return either resource metrics info or raw metrics info")
	}

	targetTimestamp := fixedTimestamp.Add(time.Duration(tc.targetTimestamp) * time.Minute)
	assert.True(t, targetTimestamp.Equal(timestamp), fmt.Sprintf("the timestamp should be as expected (%s) but was %s", targetTimestamp, timestamp))
}

func (tc *testCase) runTest(t *testing.T) {
	testClient := tc.prepareTestClient(t)
	metricsClient := NewHeapsterMetricsClient(testClient, DefaultHeapsterNamespace, DefaultHeapsterScheme, DefaultHeapsterService, DefaultHeapsterPort)
	isResource := len(tc.resourceName) > 0
	if isResource {
		info, timestamp, err := metricsClient.GetResourceMetric(tc.resourceName, tc.namespace, tc.selector)
		tc.verifyResults(t, info, timestamp, err)
	} else {
		info, timestamp, err := metricsClient.GetRawMetric(tc.metricName, tc.namespace, tc.selector)
		tc.verifyResults(t, info, timestamp, err)
	}
}

func TestCPU(t *testing.T) {
	tc := testCase{
		replicas: 3,
		desiredResourceValues: PodResourceInfo{
			"test-pod-0": 5000, "test-pod-1": 5000, "test-pod-2": 5000,
		},
		resourceName:       api.ResourceCPU,
		targetTimestamp:    1,
		reportedPodMetrics: [][]int64{{5000}, {5000}, {5000}},
	}
	tc.runTest(t)
}

func TestQPS(t *testing.T) {
	tc := testCase{
		replicas: 3,
		desiredMetricValues: PodMetricsInfo{
			"test-pod-0": 10, "test-pod-1": 20, "test-pod-2": 10,
		},
		metricName:            "qps",
		targetTimestamp:       1,
		reportedMetricsPoints: [][]metricPoint{{{10, 1}}, {{20, 1}}, {{10, 1}}},
	}
	tc.runTest(t)
}

func TestQpsSumEqualZero(t *testing.T) {
	tc := testCase{
		replicas: 3,
		desiredMetricValues: PodMetricsInfo{
			"test-pod-0": 0, "test-pod-1": 0, "test-pod-2": 0,
		},
		metricName:            "qps",
		targetTimestamp:       0,
		reportedMetricsPoints: [][]metricPoint{{{0, 0}}, {{0, 0}}, {{0, 0}}},
	}
	tc.runTest(t)
}

func TestCPUMoreMetrics(t *testing.T) {
	tc := testCase{
		replicas: 5,
		desiredResourceValues: PodResourceInfo{
			"test-pod-0": 5000, "test-pod-1": 5000, "test-pod-2": 5000,
			"test-pod-3": 5000, "test-pod-4": 5000,
		},
		resourceName:       api.ResourceCPU,
		targetTimestamp:    10,
		reportedPodMetrics: [][]int64{{1000, 2000, 2000}, {5000}, {1000, 1000, 1000, 2000}, {4000, 1000}, {5000}},
	}
	tc.runTest(t)
}

func TestCPUMissingMetrics(t *testing.T) {
	tc := testCase{
		replicas: 3,
		desiredResourceValues: PodResourceInfo{
			"test-pod-0": 4000,
		},
		resourceName:       api.ResourceCPU,
		reportedPodMetrics: [][]int64{{4000}},
	}
	tc.runTest(t)
}

func TestQpsMissingMetrics(t *testing.T) {
	tc := testCase{
		replicas:              3,
		desiredError:          fmt.Errorf("requested metrics for 3 pods, got metrics for 1"),
		metricName:            "qps",
		targetTimestamp:       1,
		reportedMetricsPoints: [][]metricPoint{{{4000, 4}}},
	}
	tc.runTest(t)
}

func TestQpsSuperfluousMetrics(t *testing.T) {
	tc := testCase{
		replicas:              3,
		desiredError:          fmt.Errorf("requested metrics for 3 pods, got metrics for 6"),
		metricName:            "qps",
		reportedMetricsPoints: [][]metricPoint{{{1000, 1}}, {{2000, 4}}, {{2000, 1}}, {{4000, 5}}, {{2000, 1}}, {{4000, 4}}},
	}
	tc.runTest(t)
}

func TestCPUEmptyMetrics(t *testing.T) {
	tc := testCase{
		replicas:              3,
		resourceName:          api.ResourceCPU,
		desiredError:          fmt.Errorf("no metrics returned from heapster"),
		reportedMetricsPoints: [][]metricPoint{},
		reportedPodMetrics:    [][]int64{},
	}
	tc.runTest(t)
}

func TestQpsEmptyEntries(t *testing.T) {
	tc := testCase{
		replicas:   3,
		metricName: "qps",
		desiredMetricValues: PodMetricsInfo{
			"test-pod-0": 4000, "test-pod-2": 2000,
		},
		targetTimestamp:       4,
		reportedMetricsPoints: [][]metricPoint{{{4000, 4}}, {}, {{2000, 4}}},
	}
	tc.runTest(t)
}

func TestCPUZeroReplicas(t *testing.T) {
	tc := testCase{
		replicas:           0,
		resourceName:       api.ResourceCPU,
		desiredError:       fmt.Errorf("no metrics returned from heapster"),
		reportedPodMetrics: [][]int64{},
	}
	tc.runTest(t)
}

func TestCPUEmptyMetricsForOnePod(t *testing.T) {
	tc := testCase{
		replicas:     3,
		resourceName: api.ResourceCPU,
		desiredResourceValues: PodResourceInfo{
			"test-pod-0": 100, "test-pod-1": 700,
		},
		reportedPodMetrics: [][]int64{{100}, {300, 400}, {}},
	}
	tc.runTest(t)
}

func testCollapseTimeSamples(t *testing.T) {
	now := time.Now()
	metrics := heapster.MetricResult{
		Metrics: []heapster.MetricPoint{
			{Timestamp: now, Value: 50, FloatValue: nil},
			{Timestamp: now.Add(-15 * time.Second), Value: 100, FloatValue: nil},
			{Timestamp: now.Add(-60 * time.Second), Value: 100000, FloatValue: nil}},
		LatestTimestamp: now,
	}

	val, timestamp, hadMetrics := collapseTimeSamples(metrics, time.Minute)
	assert.True(t, hadMetrics, "should report that it received a populated list of metrics")
	assert.InEpsilon(t, float64(75), val, 0.1, "collapsed sample value should be as expected")
	assert.True(t, timestamp.Equal(now), "timestamp should be the current time (the newest)")
}

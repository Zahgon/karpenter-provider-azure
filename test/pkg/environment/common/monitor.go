/*
Portions Copyright (c) Microsoft Corporation.

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

package common

import (
	"context"
	"sync"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Monitor is used to monitor the cluster state during a running test
type Monitor struct {
	ctx        context.Context
	kubeClient client.Client

	mu sync.RWMutex

	nodesAtReset map[string]*corev1.Node
}

type state struct {
	pods         corev1.PodList
	nodes        map[string]*corev1.Node        // node name -> node
	nodePods     map[string][]*corev1.Pod       // node name -> pods bound to the node
	nodeRequests map[string]corev1.ResourceList // node name -> sum of pod resource requests
}

func NewMonitor(ctx context.Context, kubeClient client.Client) *Monitor {
	_ = "STUB: not implemented"
	return nil
}

// Reset resets the cluster monitor prior to running a test.
func (m *Monitor) Reset() { _ = "STUB: not implemented"; return }

// RestartCount returns the containers and number of restarts for that container for all containers in the pods in the
// given namespace
func (m *Monitor) RestartCount(namespace string) map[string]int {
	_ = "STUB: not implemented"
	return nil
}

// NodeCount returns the current number of nodes
func (m *Monitor) NodeCount() int { _ = "STUB: not implemented"; return 0 }

// NodeCountAtReset returns the number of nodes that were running when the monitor was last reset, typically at the
// beginning of a test
func (m *Monitor) NodeCountAtReset() int { _ = "STUB: not implemented"; return 0 }

// CreatedNodeCount returns the number of nodes created since the last reset
func (m *Monitor) CreatedNodeCount() int { _ = "STUB: not implemented"; return 0 }

// NodesAtReset returns a slice of nodes that the monitor saw at the last reset
func (m *Monitor) NodesAtReset() []*corev1.Node { _ = "STUB: not implemented"; return nil }

// Nodes returns all the nodes on the cluster
func (m *Monitor) Nodes() []*corev1.Node { _ = "STUB: not implemented"; return nil }

// CreatedNodes returns the nodes that have been created since the last reset (essentially Nodes - NodesAtReset)
func (m *Monitor) CreatedNodes() []*corev1.Node { _ = "STUB: not implemented"; return nil }

// DeletedNodes returns the nodes that have been deleted since the last reset (essentially NodesAtReset - Nodes)
func (m *Monitor) DeletedNodes() []*corev1.Node { _ = "STUB: not implemented"; return nil }

// PendingPods returns the number of pending pods matching the given selector
func (m *Monitor) PendingPods(selector labels.Selector) []*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (m *Monitor) PendingPodsCount(selector labels.Selector) int {
	_ = "STUB: not implemented"
	return 0
}

// RunningPods returns the number of running pods matching the given selector
func (m *Monitor) RunningPods(selector labels.Selector) []*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (m *Monitor) RunningPodsCount(selector labels.Selector) int {
	_ = "STUB: not implemented"
	return 0
}

func (m *Monitor) poll() state { _ = "STUB: not implemented"; return *new(state) }

// collect pods per node

func (m *Monitor) AvgUtilization(resource corev1.ResourceName) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (m *Monitor) MinUtilization(resource corev1.ResourceName) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (m *Monitor) nodeUtilization(resource corev1.ResourceName) []float64 {
	_ = "STUB: not implemented"
	return nil
}

// skip any nodes we didn't launch

type copyable[T any] interface {
	DeepCopy() T
}

func deepCopyMap[K comparable, V copyable[V]](m map[K]V) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

func deepCopySlice[T copyable[T]](s []T) []T { _ = "STUB: not implemented"; return nil }

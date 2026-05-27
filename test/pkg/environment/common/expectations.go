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
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	karpv1 "sigs.k8s.io/karpenter/pkg/apis/v1"
)

const (
	karpenterControllerNamespace = "kube-system"
	karpenterDeploymentName      = "karpenter"
)

func (env *Environment) getInClusterController() bool { _ = "STUB: not implemented"; return false }

func (env *Environment) ExpectCreated(objects ...client.Object) { _ = "STUB: not implemented"; return }

func (env *Environment) ExpectDeleted(objects ...client.Object) { _ = "STUB: not implemented"; return }

// ExpectUpdated will update objects in the cluster to match the inputs.
// WARNING: This ignores the resource version check, which can result in
// overwriting changes made by other controllers in the cluster.
// This is useful in ensuring that we can clean up resources by patching
// out finalizers.
// Grab the object before making the updates to reduce the chance of this race.
func (env *Environment) ExpectUpdated(objects ...client.Object) { _ = "STUB: not implemented"; return }

// ExpectStatusUpdated will update objects in the cluster to match the inputs.
// WARNING: This ignores the resource version check, which can result in
// overwriting changes made by other controllers in the cluster.
// This is useful in ensuring that we can clean up resources by patching
// out finalizers.
// Grab the object before making the updates to reduce the chance of this race.
func (env *Environment) ExpectStatusUpdated(objects ...client.Object) {
	_ = "STUB: not implemented"
	return
}

func ReplaceNodeConditions(node *corev1.Node, conds ...corev1.NodeCondition) *corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

// ExpectCreatedOrUpdated can update objects in the cluster to match the inputs.
// WARNING: ExpectUpdated ignores the resource version check, which can result in
// overwriting changes made by other controllers in the cluster.
// This is useful in ensuring that we can clean up resources by patching
// out finalizers.
// Grab the object before making the updates to reduce the chance of this race.
func (env *Environment) ExpectCreatedOrUpdated(objects ...client.Object) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ExpectSettings() (res []corev1.EnvVar) {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) ExpectSettingsReplaced(vars ...corev1.EnvVar) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ExpectSettingsOverridden(vars ...corev1.EnvVar) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ExpectSettingsRemoved(vars ...corev1.EnvVar) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ExpectConfigMapExists(key types.NamespacedName) *corev1.ConfigMap {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) ExpectConfigMapDataReplaced(key types.NamespacedName, data ...map[string]string) (changed bool) {
	_ = "STUB: not implemented"
	return false
}

// Completely replace the data

// If the data hasn't changed, we can just return and not update anything

// Update the configMap to update the settings

func (env *Environment) ExpectConfigMapDataOverridden(key types.NamespacedName, data ...map[string]string) (changed bool) {
	_ = "STUB: not implemented"
	return false
}

// Update the configMap to update the settings

func (env *Environment) ExpectExists(obj client.Object) client.Object {
	_ = "STUB: not implemented"
	return *new(client.Object)
}

func (env *Environment) ExpectAllExist(objs ...client.Object) { _ = "STUB: not implemented"; return }

func (env *Environment) EventuallyExpectBound(pods ...*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) EventuallyExpectHealthy(pods ...*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) EventuallyExpectTerminating(pods ...*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) EventuallyExpectTerminatingWithTimeout(timeout time.Duration, pods ...*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) EventuallyExpectNoLeakedKubeNodeLease() {
	_ = "STUB: not implemented"

	// expect no kube node lease to be leaked
	return
}

func (env *Environment) EventuallyExpectHealthyWithTimeout(timeout time.Duration, pods ...*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ConsistentlyExpectTerminatingPods(duration time.Duration, pods ...*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ConsistentlyExpectActivePods(duration time.Duration, pods ...*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ConsistentlyExpectHealthyPods(duration time.Duration, pods ...*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) EventuallyExpectKarpenterRestarted() { _ = "STUB: not implemented"; return }

func (env *Environment) ExpectKarpenterLeaseOwnerChanged() { _ = "STUB: not implemented"; return }

func (env *Environment) EventuallyExpectRollout(name, namespace string) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ExpectKarpenterPods() []*corev1.Pod { _ = "STUB: not implemented"; return nil }

func (env *Environment) ExpectActiveKarpenterPodName() string { _ = "STUB: not implemented"; return "" }

// Holder identity for lease is always in the format "<pod-name>_<pseudo-random-value>

func (env *Environment) ExpectActiveKarpenterPod() *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectPendingPodCount(selector labels.Selector, numPods int) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) EventuallyExpectBoundPodCount(selector labels.Selector, numPods int) []*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectHealthyPodCount(selector labels.Selector, numPods int) []*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectHealthyDeployment(deployment *appsv1.Deployment) []*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectHealthyDeploymentWithTimeout(timeout time.Duration, deployment *appsv1.Deployment) []*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectHealthyPodCountWithTimeout(timeout time.Duration, selector labels.Selector, numPods int) []*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

// EventuallyExpectPVCBound waits for the PVC to reach Bound phase and returns its PV.
func (env *Environment) EventuallyExpectPVCBound(pvc *corev1.PersistentVolumeClaim) *corev1.PersistentVolume {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) ExpectHealthyPodCount(selector labels.Selector, numPods int) []*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) ExpectPodsMatchingSelector(selector labels.Selector) []*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectUniqueNodeNames(selector labels.Selector, uniqueNames int) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) eventuallyExpectScaleDown() { _ = "STUB: not implemented"; return }

// expect the current node count to be what it was when the test started

func (env *Environment) EventuallyExpectNotFound(objects ...client.Object) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ExpectCreatedNodeCount(comparator string, count int) []*corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) ExpectNodeCount(comparator string, count int) []*corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) ExpectNodeClaimCount(comparator string, count int) []*karpv1.NodeClaim {
	_ = "STUB: not implemented"
	return nil
}

func NodeClaimNames(nodeClaims []*karpv1.NodeClaim) []string { _ = "STUB: not implemented"; return nil }

func NodeNames(nodes []*corev1.Node) []string { _ = "STUB: not implemented"; return nil }

func (env *Environment) ConsistentlyExpectNodeCount(comparator string, count int, duration time.Duration) []*corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) ConsistentlyExpectCreatedNodeCount(comparator string, count int, duration time.Duration) []*corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

// ConsistentlyExpectNoDisruptions asserts that the number of tainted nodes remains the same.
// And that the number of nodeclaims remains the same.
func (env *Environment) ConsistentlyExpectNoDisruptions(nodeCount int, duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

// ConsistentlyExpectNodesNotDisrupted asserts that the captured nodes never receive the disrupted taint.
// Use this instead of ConsistentlyExpectNoDisruptions when the test only needs to prove that a specific
// node set was not disrupted, since ConsistentlyExpectNoDisruptions also asserts exact Node/NodeClaim counts.
func (env *Environment) ConsistentlyExpectNodesNotDisrupted(nodes []*corev1.Node, duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

// ConsistentlyExpectDisruptionsUntilNoneLeft consistently ensures a max on number of concurrently disrupting and non-terminating nodes.
// This actually uses an Eventually() under the hood so that when we reach 0 tainted nodes we exit early.
// We use the StopTrying() so that we can exit the Eventually() if we've breached an assertion on total concurrency of disruptions.
// For example: if we have 5 nodes, with a budget of 2 nodes, we ensure that `disruptingNodes <= maxNodesDisrupting=2`
// We use nodesAtStart+maxNodesDisrupting to assert that we're not creating too many instances in replacement.
func (env *Environment) ConsistentlyExpectDisruptionsUntilNoneLeft(nodesAtStart, maxNodesDisrupting int, timeout time.Duration) {
	_ = "STUB: not implemented"
	return
}

// We use an eventually to exit when we detect the number of tainted/disrupted nodes matches our target.

// Grab Nodes and NodeClaims

// Don't include NodeClaims with the `Terminating` status condition, as they're not included in budgets

// Don't include Nodes whose NodeClaims have been ignored

// Filter further by the number of tainted nodes to get the number of nodes that are disrupting

func (env *Environment) EventuallyExpectTaintedNodeCount(comparator string, count int) []*corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectNodesUntaintedWithTimeout(timeout time.Duration, nodes ...*corev1.Node) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) EventuallyExpectRegisteredNodeClaimCount(comparator string, count int) []*karpv1.NodeClaim {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectRegisteredNodeClaimCountWithSelector(comparator string, count int, selector labels.Selector) []*karpv1.NodeClaim {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectLaunchedNodeClaimCount(comparator string, count int) []*karpv1.NodeClaim {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectNodeCount(comparator string, count int) []*corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectNodeCountWithSelector(comparator string, count int, selector labels.Selector) []*corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectCreatedNodeCount(comparator string, count int) []*corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectCreatedNodeCountWithSelector(comparator string, count int, selector labels.Selector) []*corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectDeletedNodeCount(comparator string, count int) []*corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectDeletedNodeCountWithSelector(comparator string, count int, selector labels.Selector) []*corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectInitializedNodeCount(comparator string, count int) []*corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectCreatedNodeClaimCount(comparator string, count int) []*karpv1.NodeClaim {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) EventuallyExpectNodeClaimsReady(nodeClaims ...*karpv1.NodeClaim) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) EventuallyExpectDrifted(nodeClaims ...*karpv1.NodeClaim) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) EventuallyExpectDriftedWithTimeout(timeout time.Duration, nodeClaims ...*karpv1.NodeClaim) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ConsistentlyExpectNodeClaimsNotDrifted(duration time.Duration, nodeClaims ...*karpv1.NodeClaim) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) EventuallyExpectConsolidatable(nodeClaims ...*karpv1.NodeClaim) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) GetNode(nodeName string) *corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) ExpectNoCrashes() { _ = "STUB: not implemented"; return }

var (
	lastLogged = metav1.Now()
)

func (env *Environment) printControllerLogs(options *corev1.PodLogOptions) {
	_ = "STUB: not implemented"
	return
}

// local version of the log options

func (env *Environment) EventuallyExpectMinUtilization(resource corev1.ResourceName, comparator string, value float64) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) EventuallyExpectAvgUtilization(resource corev1.ResourceName, comparator string, value float64) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ExpectDaemonSetEnvironmentVariableUpdated(obj client.ObjectKey, name, value string, containers ...string) {
	_ = "STUB: not implemented"
	return
}

// If the env var already exists, update its value. Otherwise, create a new var.

// ForcePodsToSpread ensures that currently scheduled pods get spread evenly across all passed nodes by deleting pods off of existing
// nodes and waiting them to reschedule. This is useful for scenarios where you want to force the nodes be underutilized
// but you want to keep a consistent count of nodes rather than leaving around empty ones.
func (env *Environment) ForcePodsToSpread(nodes ...*corev1.Node) {
	_ = "STUB: not implemented"

	// Get the total count of pods across
	return
}

// Set the nodes to unschedulable so that the pods won't reschedule.

// TODO: Consider moving this time check to an Eventually poll. This gets a little tricker with helper functions
// since you need to make sure that your Expectation helper functions are scoped to to your "g Gomega" scope
// so that you don't fail the first time you get a failure on your expectation

func (env *Environment) ExpectActivePodsForNode(nodeName string) []*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) ExpectCABundle() string {
	_ = "STUB: not implemented"
	// Discover CA Bundle from the REST client. We could alternatively
	// have used the simpler client-go InClusterConfig() method.
	// However, that only works when Karpenter is running as a Pod
	// within the same cluster it's managing.
	return ""
}

// fills in CAData!

func (env *Environment) GetDaemonSetCount(np *karpv1.NodePool) int {
	_ = "STUB: not implemented"

	// Performs the same logic as the scheduler to get the number of daemonset
	// pods that we estimate we will need to schedule as overhead to each node
	return 0
}

func (env *Environment) GetDaemonSetOverhead(np *karpv1.NodePool) corev1.ResourceList {
	_ = "STUB: not implemented"

	// Performs the same logic as the scheduler to get the number of daemonset
	// pods that we estimate we will need to schedule as overhead to each node
	return *new(corev1.ResourceList)
}

func (env *Environment) IsCilium() bool { _ = "STUB: not implemented"; return false }

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

package expectations

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
	"github.com/Azure/karpenter-provider-azure/pkg/test"
	"github.com/Azure/skewer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	karpv1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	corecloudprovider "sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/controllers/provisioning"
	"sigs.k8s.io/karpenter/pkg/controllers/provisioning/scheduling"
	"sigs.k8s.io/karpenter/pkg/controllers/state"
)

func ExpectUnavailable(env *test.Environment, sku *skewer.SKU, zone string, capacityType string) {
	_ = "STUB: not implemented"
	return
}

func ExpectKubeletFlags(_ *test.Environment, customData string, expectedFlags map[string]string) {
	_ = "STUB: not implemented"
	return
}

func ExpectDecodedCustomData(env *test.Environment) string { _ = "STUB: not implemented"; return "" }

func ExpectCSEProvisioned(env *test.Environment) armcompute.VirtualMachineExtension {
	_ = "STUB: not implemented"
	return *new(armcompute.VirtualMachineExtension)
}

// CSE provisioning is asynchronous, starting after VM creation LRO completes

func ExpectCSENotProvisioned(env *test.Environment) { _ = "STUB: not implemented"; return }

// ExpectCleanUp handled the cleanup of all Objects we need within testing that core does not
//
// Core's ExpectCleanedUp function does not currently cleanup ConfigMaps:
// https://github.com/kubernetes-sigs/karpenter/blob/db8df23ffb0b689b116d99597316612c98d382ab/pkg/test/expectations/expectations.go#L244
// TODO: surface this within core and remove this function
func ExpectCleanUp(ctx context.Context, c client.Client) { _ = "STUB: not implemented"; return }

func ExpectInstanceResourcesHaveTags(ctx context.Context, name string, azureEnv *test.Environment, tags map[string]*string) *armcompute.VirtualMachine {
	_ = "STUB: not implemented"

	// The VM should be updated
	return nil
}

// Expect the identities to remain unchanged

// The NIC should be updated

// The extensions should be updated -- Note that we expect only 1 Extension update here because we're simulating scriptless
// mode which doesn't have a CSE extension.

// TODO: Upstream this?
func ExpectLaunched(ctx context.Context, c client.Client, cloudProvider corecloudprovider.CloudProvider, provisioner *provisioning.Provisioner, pods ...*corev1.Pod) {
	_ = "STUB: not implemented"

	// Persist objects
	return
}

func ExpectNodeClassHashUpdated(ctx context.Context, c client.Client, nodeClass *v1beta1.AKSNodeClass) {
	_ = "STUB: not implemented"
	return
}

// instancePromiseWaiter breaks the import cycle between pkg/cloudprovider and
// this package: cloudprovider tests import this package, so this package
// cannot import cloudprovider back. *cloudprovider.CloudProvider satisfies it.
type instancePromiseWaiter interface {
	WaitForInstancePromises()
}

// ExpectProvisionedAndWaitForPromises provisions pods and waits for async polling goroutines to complete.
// This ensures that any background Create operations (including GET poller) finish before
// the test continues, preventing goroutines from interfering with subsequent assertions.
//
// Use this instead of upstream ExpectProvisioned to ensure proper async cleanup.
func ExpectProvisionedAndWaitForPromises(
	ctx context.Context,
	c client.Client,
	cluster *state.Cluster,
	cp corecloudprovider.CloudProvider,
	provisioner *provisioning.Provisioner,
	azureEnv *test.Environment,
	pods ...*corev1.Pod,
) {
	_ = "STUB: not implemented"
	return
}

func ExpectScheduledNodeClaimsCreated(
	ctx context.Context,
	client client.Client,
	coreProvisioner *provisioning.Provisioner,
	claims ...*scheduling.NodeClaim,
) []*karpv1.NodeClaim {
	_ = "STUB: not implemented"
	return nil
}

func BindPodsToNode(ctx context.Context, client client.Client, cluster *state.Cluster, scheduledClaim *scheduling.NodeClaim, node *corev1.Node) {
	_ = "STUB: not implemented"
	return
}

// We have to manually bind the pod to the node when using a fakeClient by setting the value for pod.Spec.NodeName
// Note: This is a bit hacky but is what upstream does, see https://github.com/kubernetes-sigs/karpenter/blob/defdfae64097b8e58a211c429fa955896e515400/pkg/test/expectations/expectations.go#L307

// track pod bindings

// CreateAndWaitForPromises calls cloudProvider.Create and waits for async polling goroutines to complete.
// It sets the Launched condition on the NodeClaim (mirroring what the core lifecycle controller
// does in production) so that the async goroutine's waitUntilLaunched unblocks.
// Returns the created NodeClaim and any error from the Create operation.
//
// Use this instead of direct cloudProvider.Create() calls in tests.
func CreateAndWaitForPromises(
	ctx context.Context,
	cp corecloudprovider.CloudProvider,
	azureEnv *test.Environment,
	nodeClaim *karpv1.NodeClaim,
) (*karpv1.NodeClaim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Simulate what the core lifecycle Launch controller does after Create():
// set Launched=True so the async goroutine's waitUntilLaunched unblocks.
// We fetch a fresh copy from the API server and do a status-only update to
// avoid "spec is immutable" errors when the test has modified the spec
// (e.g., conflicted NodeClaim tests).

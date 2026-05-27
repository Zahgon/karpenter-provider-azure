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

package cloudprovider

import (
	"context"
	"net/http"
	"sync"

	"github.com/awslabs/operatorpkg/status"
	"sigs.k8s.io/karpenter/pkg/controllers/nodeoverlay"

	"sigs.k8s.io/controller-runtime/pkg/client"

	//nolint:SA1019 // deprecated package
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"

	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"

	"github.com/Azure/karpenter-provider-azure/pkg/providers/imagefamily"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/instance"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/instancetype"

	karpv1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/events"

	"sigs.k8s.io/karpenter/pkg/cloudprovider"
)

const (
	NodeClassReadinessUnknownReason    = "NodeClassReadinessUnknown"
	InstanceTypeResolutionFailedReason = "InstanceTypeResolutionFailed"
	CreateInstanceFailedReason         = "CreateInstanceFailed"
)

var _ cloudprovider.CloudProvider = (*CloudProvider)(nil)

type CloudProvider struct {
	instanceTypeProvider       instancetype.Provider
	vmInstanceProvider         instance.VMProvider // Note that even when provision mode does not create with VM instance provider, it is still being used to handle existing VM instances.
	aksMachineInstanceProvider instance.AKSMachineProvider
	kubeClient                 client.Client
	imageProvider              imagefamily.NodeImageProvider
	recorder                   events.Recorder
	instanceTypeStore          *nodeoverlay.InstanceTypeStore
	instancePromiseWg          sync.WaitGroup
}

func New(
	instanceTypeProvider instancetype.Provider,
	vmInstanceProvider instance.VMProvider,
	aksMachineInstanceProvider instance.AKSMachineProvider,
	recorder events.Recorder,
	kubeClient client.Client,
	imageProvider imagefamily.NodeImageProvider,
	store *nodeoverlay.InstanceTypeStore,
) *CloudProvider {
	_ = "STUB: not implemented"
	return nil
}

// WaitForInstancePromises blocks until all in-flight async Create goroutines have completed.
func (c *CloudProvider) WaitForInstancePromises() { _ = "STUB: not implemented"; return }

func (c *CloudProvider) validateNodeClass(nodeClass *v1beta1.AKSNodeClass) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CloudProvider) Create(ctx context.Context, nodeClaim *karpv1.NodeClaim) (*karpv1.NodeClaim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We treat a failure to resolve the NodeClass as an ICE since this means there is no capacity possibilities for this NodeClaim

/*
	// TODO: Remove this after v1
	nodePool, err := utils.ResolveNodePoolFromNodeClaim(ctx, c.kubeClient, nodeClaim)
	if err != nil {
		return nil, err
	}
	kubeletHash, err := utils.GetHashKubelet(nodePool, nodeClass)
	if err != nil {
		return nil, err
	}
*/

// Note: This filters out any instance types which we're out of capacity for

// Choose provider based on provision mode

func (c *CloudProvider) createVMInstance(ctx context.Context, nodeClass *v1beta1.AKSNodeClass, nodeClaim *karpv1.NodeClaim, instanceTypes []*cloudprovider.InstanceType) (*karpv1.NodeClaim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This is best-effort populated by Karpenter to be used to create the VM server-side. Not all fields are guaranteed to be populated, especially status fields.
// Double-check the code before making assumptions on their presence.

// Propagate single-value wellKnownLabels from the nodeClaim requirements to the labels.
// This is required for scheduling in core to work correctly. If this is not done, on subsequent scheduling passes before the Node is
// registered, the NodeClaim will not have the labels required to match the Pod and so a new NodeClaim will be created each time.
// Note that AWS does this by explicitly setting the labels in their CloudProvider (see https://github.com/aws/karpenter-provider-aws/blob/main/pkg/cloudprovider/cloudprovider.go#L456)
// rather than doing it in bulk here.
// TODO: should we do like AWS and smuggle all of these labels through VM tags rather than just setting them here?

func (c *CloudProvider) createAKSMachineInstance(ctx context.Context, nodeClass *v1beta1.AKSNodeClass, nodeClaim *karpv1.NodeClaim, instanceTypes []*cloudprovider.InstanceType) (*karpv1.NodeClaim, error) {
	_ = "STUB: not implemented"
	// Begin the creation of the instance
	return nil, nil
}

// Handle the promise

// Convert the AKS machine to a NodeClaim

// Propagate single-value wellKnownLabels from the nodeClaim requirements to the labels.
// This is required for scheduling in core to work correctly. If this is not done, on subsequent scheduling passes before the Node is
// registered, the NodeClaim will not have the labels required to match the Pod and so a new NodeClaim will be created each time.
// Note that AWS does this by explicitly setting the labels in their CloudProvider (see https://github.com/aws/karpenter-provider-aws/blob/main/pkg/cloudprovider/cloudprovider.go#L456)
// rather than doing it in bulk here.
// TODO: should we do like AWS and smuggle all of these labels through VM tags rather than just setting them here?

// handleInstancePromise handles the instance promise, primarily deciding on sync/async provisioning.
func (c *CloudProvider) handleInstancePromise(ctx context.Context, instancePromise instance.Promise, nodeClaim *karpv1.NodeClaim) error {
	_ = "STUB: not implemented"
	return nil
}

// Standalone NodeClaims aren't re-queued for reconciliation in the provision_trigger controller,
// so we delete them synchronously. After marking Launched=true,
// their status can't be reverted to false once the delete completes due to how core caches nodeclaims in
// the launch controller. This ensures we retry continuously until we hit the registration TTL

// For NodePool-managed nodeclaims, launch a single goroutine to poll the returned promise.
// Note that we could store the LRO details on the NodeClaim, but we don't bother today because Karpenter
// crashes should be rare, and even in the case of a crash, as long as the node comes up successfully there's
// no issue. If the node doesn't come up successfully in that case, the node and the linked claim will
// be garbage collected after the TTL, but the cause of the nodes issue will be lost, as the LRO URL was
// only held in memory.

// Only log if context is still active to avoid logging after test completes

// Wait until the claim is Launched, to avoid racing with creation.
// This isn't strictly required, but without this, failure test scenarios are harder
// to write because the nodeClaim gets deleted by error handling below before
// the EnsureApplied call finishes, so EnsureApplied creates it again (which is wrong/isn't how
// it would actually happen in production).

// For async provisioning, also delete the NodeClaim

// Only log if context is still active to avoid logging after test completes

func (c *CloudProvider) handleInstancePromiseWaitError(ctx context.Context, instancePromise instance.Promise, nodeClaim *karpv1.NodeClaim, waitErr error) {
	_ = "STUB: not implemented"
	return
}

// Only log if context is still active to avoid logging after test completes

// Fallback to garbage collection to clean up the instance, if it survived.

// Only log if context is still active to avoid logging after test completes

func (c *CloudProvider) waitUntilLaunched(ctx context.Context, nodeClaim *karpv1.NodeClaim) {
	_ = "STUB: not implemented"
	return
}

// Only log if context is still active to avoid logging after test completes

// context was canceled

func (c *CloudProvider) List(ctx context.Context) ([]*karpv1.NodeClaim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List AKS machine-based nodes

// List VM-based nodes

func (c *CloudProvider) Get(ctx context.Context, providerID string) (*karpv1.NodeClaim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AKS machine-based node

func (c *CloudProvider) LivenessProbe(req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// GetInstanceTypes returns all available InstanceTypes
// May return apimachinery.NotFoundError if NodeClass is not found.
func (c *CloudProvider) GetInstanceTypes(ctx context.Context, nodePool *karpv1.NodePool) ([]*cloudprovider.InstanceType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We must return an error here in the event of the node class not being found. Otherwise users just get
// no instance types and a failure to schedule with no indicator pointing to a bad configuration
// as the cause.

// Delete deletes the underlying node
// Note: Delete may be called many times while delete is ongoing (blocking) as the core Karpenter termination controller
// watches and reconciles on all node updates (including node status updates, which happen during deletion), so while
// one Delete call is blocking more will come in every ~5s due to excess Node events + requeues.
func (c *CloudProvider) Delete(ctx context.Context, nodeClaim *karpv1.NodeClaim) error {
	_ = "STUB: not implemented"
	return nil
}

// AKS machine-based node

// VM-based node

// IsDrifted checks if the NodeClaim has drifted from our goal state.
// Note: During the initial launch and registration of a NodeClaim,
// core calls IsDrifted quite frequently as it waits for the Node to register and become ready. This is
// because the core pkg/controllers/nodeclaim/disruption/controller.go watches NodeClaims without a
// generation filter, so any update to the NodeClaim (including updates to status such as when updating conditions during launch)
// will trigger a call to IsDrifted.
// The following things produce a large number of IsDrifted calls:
//   - The initialization controller pkg/controllers/nodeclaim/lifecycle/initialization.go changes the ConditionTypeInitialized condition a number of times during
//     this process, which triggers disruption/controller.go to call IsDrifted each time.
//   - Any pod scheduling that happens during this time will trigger the core disruption/controller.go, because it watches pod updates and
//     maps each pod update to a NodeClaim event. This means every time a pod (including a DaemonSet pod) is scheduled to the node
//     we'll get called.
func (c *CloudProvider) IsDrifted(ctx context.Context, nodeClaim *karpv1.NodeClaim) (cloudprovider.DriftReason, error) {
	_ = "STUB: not implemented"
	// Not needed when GetInstanceTypes removes nodepool dependency
	return *new(cloudprovider.DriftReason), nil
}

// Name returns the CloudProvider implementation name.
func (c *CloudProvider) Name() string { _ = "STUB: not implemented"; return "" }

func (c *CloudProvider) GetSupportedNodeClasses() []status.Object {
	_ = "STUB: not implemented"
	return nil
}

// TODO: review repair policies
func (c *CloudProvider) RepairPolicies() []cloudprovider.RepairPolicy {
	_ = "STUB: not implemented"
	return nil
}

// Supported Kubelet fields

// May return apimachinery.NotFoundError if NodePool is not found.
func (c *CloudProvider) resolveNodeClassFromNodePool(ctx context.Context, nodePool *karpv1.NodePool) (*v1beta1.AKSNodeClass, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// For the purposes of NodeClass CloudProvider resolution, we treat deleting NodeClasses as NotFound

// For the purposes of NodeClass CloudProvider resolution, we treat deleting NodeClasses as NotFound,
// but we return a different error message to be clearer to users

func (c *CloudProvider) resolveInstanceTypes(ctx context.Context, nodeClaim *karpv1.NodeClaim, nodeClass *v1beta1.AKSNodeClass) ([]*cloudprovider.InstanceType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Old logic
// return lo.Filter(instanceTypes, func(i *cloudprovider.InstanceType, _ int) bool {
//	return reqs.Get(v1.LabelInstanceTypeStable).Has(i.Name) &&
//		len(i.Offerings.Requirements(reqs).Available()) > 0
// }), nil

func (c *CloudProvider) resolveInstanceTypeFromVMInstance(ctx context.Context, vm *armcompute.VirtualMachine) (*cloudprovider.InstanceType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If we can't resolve the provisioner, we fallback to not getting instance type info

// If we can't resolve the nodepool, we fallback to not getting instance type info

func (c *CloudProvider) resolveNodePoolFromVMInstance(ctx context.Context, vm *armcompute.VirtualMachine) (*karpv1.NodePool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CloudProvider) vmInstanceToNodeClaim(ctx context.Context, vm *armcompute.VirtualMachine, instanceType *cloudprovider.InstanceType) (*karpv1.NodeClaim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Fallback to current time to ensure garbage collection grace period is enforced
// when TimeCreated is unavailable. Without this, CreationTimestamp would be epoch (zero value)
// and the instance could be immediately garbage collected, bypassing the 5-minute grace period.
// TODO: Investigate a more fail-safe approach. If vm.Properties.TimeCreated is NEVER populated,
// this fallback means the VM will never be garbage collected since we call this helper every time
// we create an in-memory NodeClaim. We currently assume this shouldn't happen because VMs that fail
// to come up should eventually stop appearing in Azure API responses.

// Set the deletionTimestamp to be the current time if the instance is currently terminating

func GetNodeClaimNameFromVMName(vmName string) string { _ = "STUB: not implemented"; return "" }

const truncateAt = 1200

func isNodeClaimStandalone(nodeClaim *karpv1.NodeClaim) bool {
	_ = "STUB: not implemented"
	// NodeClaims without the nodepool label are considered standalone
	return false
}

func truncateMessage(msg string) string { _ = "STUB: not implemented"; return "" }

func setAdditionalAnnotationsForNewNodeClaim(ctx context.Context, nodeClaim *karpv1.NodeClaim, nodeClass *v1beta1.AKSNodeClass) error {
	_ = "STUB: not implemented"
	// Additional annotations
	// ASSUMPTION: this is not needed in other places that the core also wants NodeClaim (e.g., Get, List).
	// As of the time of writing, AWS is doing something similar.
	// Suggestion: could have added this in instance.BuildNodeClaimFromAKSMachine, but might sacrifice some performance (little?), and need to consider that the calculated hash may change.
	return nil
}

func (c *CloudProvider) resolveNodeClaimFromAKSMachine(ctx context.Context, aksMachine *armcontainerservice.Machine) (*karpv1.NodeClaim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unknown error

// If GetInstanceTypes returns not found, we tolerate. But, possible instance types will be empty.

// Unknown error

// If FindNodePoolFromAKSMachine returns not found, we tolerate. But, possible instance types will be empty.

// ASSUMPTION: all machines are in the same location, and in the current pool.

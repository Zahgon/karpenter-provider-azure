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

package inplaceupdate

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	karpv1 "sigs.k8s.io/karpenter/pkg/apis/v1"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
	"github.com/Azure/karpenter-provider-azure/pkg/operator/options"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/instance"
)

type Controller struct {
	kubeClient                 client.Client
	vmInstanceProvider         instance.VMProvider
	aksMachineInstanceProvider instance.AKSMachineProvider
}

func NewController(
	kubeClient client.Client,
	vmInstanceProvider instance.VMProvider,
	aksMachineInstanceProvider instance.AKSMachineProvider,
) *Controller {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) Reconcile(ctx context.Context, nodeClaim *karpv1.NodeClaim) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// No need to add nodeClaim name to the context as it's already there

// Get the NodeClass

// TODO: When we have sources of truth coming from NodePool we can do:
// nodePool, err := nodeclaimutil.Owner(ctx, c.kubeClient, nodeClaim)
// TODO: To look it up and use that as input to calculate the goal state as well

// Compare the expected hash with the actual hash

// If there's no difference from goal state, no need to do anything else

// AKS machine-based nodeClaim

// VM-based nodeClaim

// Regardless of whether we actually changed anything in Azure, we have confirmed that
// the goal shape is in alignment with our expected shape, so update the annotation to reflect that

func (c *Controller) shouldProcess(ctx context.Context, nodeClaim *karpv1.NodeClaim) (bool, reconcile.Result) {
	_ = "STUB: not implemented"
	return false, *new(reconcile.Result)
}

// If the node isn't registered yet, we need to wait until it is as otherwise all the resources we need to update may not exist yet

func (c *Controller) processVMInstance(
	ctx context.Context,
	options *options.Options,
	nodeClaim *karpv1.NodeClaim,
	nodeClass *v1beta1.AKSNodeClass,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) processAKSMachineInstance(
	ctx context.Context,
	options *options.Options,
	nodeClaim *karpv1.NodeClaim,
	nodeClass *v1beta1.AKSNodeClass,
	aksMachineName string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) applyVMPatch(
	ctx context.Context,
	options *options.Options,
	nodeClaim *karpv1.NodeClaim,
	nodeClass *v1beta1.AKSNodeClass,
	vm *armcompute.VirtualMachine,
) error {
	_ = "STUB: not implemented"
	return nil
}

// This is safe only as long as we're not updating fields which we consider secret.
// If we do/are, we need to redact them.

// Apply the update, if one is needed

func (c *Controller) applyAKSMachinePatch(
	ctx context.Context,
	options *options.Options,
	nodeClaim *karpv1.NodeClaim,
	nodeClass *v1beta1.AKSNodeClass,
	aksMachineName string,
	aksMachine *armcontainerservice.Machine,
) error {
	_ = "STUB: not implemented"
	// Create a deep copy of the original for diff comparison
	return nil
}

// This is safe only as long as we're not updating fields which we consider secret.
// If we do/are, we need to redact them.

// Apply the update, if one is needed

// Extract ETag for optimistic concurrency control

// Given AKS machine support PUT, but not PATCH, the AKS machine object will be updated directly w/ etag check

// ASSUMPTION: if it is etag mismatch, the next try would work (if without another mismatch)

func (c *Controller) Register(_ context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// Note that this will trigger on pod restart for all Machines.

// TODO: Can add .Watches(&karpv1.NodePool{}, nodeclaimutil.NodePoolEventHandler(c.kubeClient))
// TODO: similar to https://github.com/kubernetes-sigs/karpenter/blob/main/pkg/controllers/nodeclaim/disruption/controller.go#L214C3-L217C5
// TODO: if/when we need to monitor provisioner changes and flow updates on the NodePool down to the underlying VMs.

// TODO: Document why this magic number used. If we want to consistently use it accoss reconcilers, refactor to a reused const.
// Comments thread discussing this: https://github.com/Azure/karpenter-provider-azure/pull/729#discussion_r2006629809

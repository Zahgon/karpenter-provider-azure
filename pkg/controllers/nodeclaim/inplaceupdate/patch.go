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

	karpv1 "sigs.k8s.io/karpenter/pkg/apis/v1"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
	"github.com/Azure/karpenter-provider-azure/pkg/operator/options"
)

func logVMPatch(ctx context.Context, update *armcompute.VirtualMachineUpdate) {
	_ = "STUB: not implemented"
	return
}

func logAKSMachinePatch(ctx context.Context, before, after *armcontainerservice.Machine) {
	_ = "STUB: not implemented"
	return
}

type patchParameters struct {
	opts      *options.Options
	nodeClaim *karpv1.NodeClaim
	nodeClass *v1beta1.AKSNodeClass
}

var vmPatchers = []func(*armcompute.VirtualMachineUpdate, *patchParameters, *armcompute.VirtualMachine) bool{
	patchVMIdentities,
	patchVMTags,
}

var aksMachinePatchers = []func(*patchParameters, *armcontainerservice.Machine) bool{
	// VM identities are handled server-side for AKS machines. No need here.
	patchAKSMachineTags,
}

func stringPtrEqual(v1, v2 *string) bool { _ = "STUB: not implemented"; return false }

func tagsEqual(expected, current map[string]*string) bool { _ = "STUB: not implemented"; return false }

func CalculateVMPatch(
	options *options.Options,
	nodeClaim *karpv1.NodeClaim,
	nodeClass *v1beta1.AKSNodeClass,
	currentVM *armcompute.VirtualMachine,
) *armcompute.VirtualMachineUpdate {
	_ = "STUB: not implemented"
	return nil
}

// No update to perform

// Note: AKS machine patching flow is different from VM patching, given AKS machine API supports PUT but not PATCH (i.e., send only diff to the API rather than the whole object).
// Thus, the patch will be applied locally on the AKS machine object, before the object is sent to the API.
func CalculateAKSMachinePatch(
	options *options.Options,
	nodeClaim *karpv1.NodeClaim,
	nodeClass *v1beta1.AKSNodeClass,
	patchingAKSMachine *armcontainerservice.Machine,
) bool {
	_ = "STUB: not implemented"
	return false
}

func patchVMIdentities(
	update *armcompute.VirtualMachineUpdate,
	params *patchParameters,
	currentVM *armcompute.VirtualMachine,
) bool {
	_ = "STUB: not implemented"
	return false
}

// It's not possible to PATCH identities away, so for now we never remove them even if they've been removed from
// the configmap. This matches the RPs behavior and also ensures that we don't remove identities which users have
// manually added.

// No update to perform

func patchVMTags(
	update *armcompute.VirtualMachineUpdate,
	params *patchParameters,
	currentVM *armcompute.VirtualMachine,
) bool {
	_ = "STUB: not implemented"
	return false
}

// No update to perform

func patchAKSMachineTags(
	params *patchParameters,
	patchingAKSMachine *armcontainerservice.Machine,
) bool {
	_ = "STUB: not implemented"
	// For NodeClaim name tag, given this controller is based on actual NodeClaim like during Create(), the patch will repair the tag if needed.
	return false
}

// Should not be possible, but handle it gracefully

// No update to perform

// No update to perform

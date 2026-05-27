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

package instance

import (
	"context"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
	"sigs.k8s.io/controller-runtime/pkg/client"
	karpv1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	corecloudprovider "sigs.k8s.io/karpenter/pkg/cloudprovider"

	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
)

var (
	KarpCapacityTypeToAKSScaleSetPriority = map[string]armcontainerservice.ScaleSetPriority{
		karpv1.CapacityTypeSpot:     armcontainerservice.ScaleSetPrioritySpot,
		karpv1.CapacityTypeOnDemand: armcontainerservice.ScaleSetPriorityRegular,
	}
	AKSScaleSetPriorityToKarpCapacityType = map[armcontainerservice.ScaleSetPriority]string{
		armcontainerservice.ScaleSetPrioritySpot:    karpv1.CapacityTypeSpot,
		armcontainerservice.ScaleSetPriorityRegular: karpv1.CapacityTypeOnDemand,
	}
	KarpCapacityTypeToScaleSetPriorityLabel = map[string]string{
		karpv1.CapacityTypeSpot:     v1beta1.ScaleSetPrioritySpot,
		karpv1.CapacityTypeOnDemand: v1beta1.ScaleSetPriorityRegular,
	}
	KarpCapacityTypeToPriorityLabel = map[string]string{
		karpv1.CapacityTypeSpot:     v1beta1.PrioritySpot,
		karpv1.CapacityTypeOnDemand: v1beta1.PriorityRegular,
	}
)

// Convention(?) These can change if needed. The purpose is to make assumptions more visible.
// Find:
//   Input: structs or values
//   Output: a struct, representing a resource to look for
// Build:
//   Input: structs or values
//   Output: a new struct
// Get:
//   Input: structs or values
//   Output: a struct or value

// Note that the template is not guaranteed to have status fields, thus they are made explicit here.
// Other Karpenter-level fields are also included as they may be easily retrieved during templating phase.
// Not assuming that NodeClaim exists.
func BuildNodeClaimFromAKSMachineTemplate(
	ctx context.Context,
	aksMachineTemplate *armcontainerservice.Machine,
	instanceType *corecloudprovider.InstanceType, // optional; won't be populated for standalone nodeclaims
	capacityType string,
	zone *string, // <region>-<zone-id>, optional
	aksMachineResourceID string,
	vmResourceID string,
	isDeleting bool,
	aksMachineNodeImageVersion string,
	creationTimestamp time.Time,
) (*karpv1.NodeClaim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Missing tag (by design, only possible if user intervenes) will eventually be repaired by in-place update controller.
// By the time of writing, this is being used for logging purposes within provider only.
// That is unlikely to change for core. But be mindful of provider is to rely on this in that situation. Still, rare.
// This was less of a concern for VM instance as NodeClaim name is always inferable from instance name.

// Determine instance creation timestamp for garbage collection purposes.
// Note: this assignment to NodeClaim is not effective to the actual object in the cluster, which still represents NodeClaim's (not instance's) creation time.
// This "borrowed struct field" is used by provider for instance garbage collection. AWS does the same.
// Note that it is incorrect to use actual NodeClaim's creation time, as retries can occur on the same NodeClaim, hurting grace period with each.

// Set the deletionTimestamp to be the current time if the instance is currently terminating

// ASSUMPTION: this doesn't need to be full image ID (should be fine on core, as the definition of ID is provider agnostic)

// Expect AKS machine struct to be fully populated as if it comes from GET.
// Not assuming that NodeClaim exists.
func BuildNodeClaimFromAKSMachine(ctx context.Context, aksMachine *armcontainerservice.Machine, possibleInstanceTypes []*corecloudprovider.InstanceType, aksMachineLocation string) (*karpv1.NodeClaim, error) {
	_ = "STUB: not implemented"
	// ASSUMPTION: unless said otherwise, the fields below must exist on the AKS machine instance. Either set by Karpenter (see Create()) or visibly defaulted by the API.
	return nil, nil
}

// Empty: not fatal, no need to check

// May return apimachinery.NotFoundError if NodePool is not found.
func FindNodePoolFromAKSMachine(ctx context.Context, aksMachine *armcontainerservice.Machine, kubeClient client.Client) (*karpv1.NodePool, error) {
	_ = "STUB: not implemented"
	// ASSUMPTION: NodePool name is stored in the all AKS machine tags.
	return nil, nil
}

// ASSUMPTION: NodeClaim name is in the format of <NodePool name>-<hash suffix>
// If total length exceeds AKS machine name limit, the exceeded part will be replaced with another deterministic hash.
// E.g., "thisisalongnodepoolname-a1b2c" --> "thisisalongnoz9y8x7-a1b2c"
func GetAKSMachineNameFromNodeClaimName(nodeClaimName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Defined by AKS machine API.
// The length of the hashed part replacing the exceeded part of the prefix.
// If 6, given alphanumeric hash, there will be a total of 36^6 = 2,176,782,336 combinations.

// Safe to use the whole name

// Keep the legit part of the prefix intact, but hash the rest
// ASSUMPTION: prefix length is at least 6 characters at this point (which means suffix length is not too large)
// At the time of writing, suffix length is 6 (e.g., "-a1b2c"). This is unlikely to change.

// GetAKSMachineNameFromNodeClaim extracts the AKS machine name from the NodeClaim annotations
// Returns false if the annotation is not present or the value is empty, which can indicate that the NodeClaim is not associated with an AKS machine.
func GetAKSMachineNameFromNodeClaim(nodeClaim *karpv1.NodeClaim) (string, bool) {
	_ = "STUB: not implemented"
	// ASSUMPTION: A NodeClaim is associated with an AKS machine iff the annotation is present (e.g., annotated when creating NodeClaim from a AKS machine instance).
	// ASSUMPTION: If exists, the annotation is always set to the correct AKS machine resource ID.
	return "", false
}

// The last part of the resource ID is the AKS machine name

func getCapacityTypeFromAKSScaleSetPriority(scaleSetPriority armcontainerservice.ScaleSetPriority) string {
	_ = "STUB: not implemented"
	return ""
}

// vmName = aks-<machinesPoolName>-<aksMachineName>-########-vm
// This is distinguishable from VM instance name as its suffix will always be 5 alphanumerics rather than "vm"
func GetAKSMachineNameFromVMName(aksMachinesPoolName, vmName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Check it ends with -vm

// Remove -vm, then find the last dash to remove ########

func isAKSMachineDeleting(aksMachine *armcontainerservice.Machine) bool {
	_ = "STUB: not implemented"
	return false
}

// GetAKSLabelZoneFromAKSMachine returns the zone for the given AKS machine, or RegionalZone ("0") if there is no zone specified.
// This function is analogous to zones.MakeAKSLabelZoneFromVM but for AKS machines.
func GetAKSLabelZoneFromAKSMachine(aksMachine *armcontainerservice.Machine, location string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func validateRetrievedAKSMachineBasicProperties(aksMachine *armcontainerservice.Machine) error {
	_ = "STUB: not implemented"
	// Assumptions may be made after this function returns no error.
	// Thus, check every usage before removing each validation.
	return nil
}

// This not being guaranteed is per the current behavior of both AKS machine API and AKS AgentPool API: priority will shows up only for spot.
// Although, it is expected that it gets rehydrated client-side before this validation function is called.
// Suggestion: rework/research more on this pattern RP-side?

func shouldAKSMachinesBeVisible(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

// BuildJSONFromAKSMachine returns a JSON string representation of an AKS Machine for logging/debugging purposes.
// Returns an error string if marshaling fails.
func BuildJSONFromAKSMachine(aksMachine *armcontainerservice.Machine) string {
	_ = "STUB: not implemented"
	return ""
}

// Safely marshal an error object to ensure any quotes in the error message are properly escaped.

// Fallback to a simple, static JSON error string if marshaling the error object also fails.

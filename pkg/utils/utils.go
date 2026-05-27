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

package utils

import (
	"context"
	"regexp"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
	"github.com/Azure/skewer"
	"github.com/mitchellh/hashstructure/v2"

	v1 "k8s.io/api/core/v1"
)

// extractVersionFromVMSize extracts and normalizes the version from VMSizeType, dropping "v" prefix and backfilling "1"
func ExtractVersionFromVMSize(vmsize *skewer.VMSizeType) string {
	_ = "STUB: not implemented"
	// safety-check to avoid panics, shouldn't happen in practice
	return ""
}

// should never happen; don't capture in label (won't be available for selection by version)

// azureResourceGroupNameRE is used to extract the resource group name from an Azure resource ID.
var azureResourceGroupNameRE = regexp.MustCompile(`.*/subscriptions/(?:.*)/resourceGroups/(.+)/providers/(?:.*)`)

// convertResourceGroupNameToLower converts the resource group name in the resource ID to be lowered.
// Inlined from sigs.k8s.io/cloud-provider-azure/pkg/provider to avoid pulling in a dependency
// that has incompatible armcompute version requirements.
func convertResourceGroupNameToLower(resourceID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func VMResourceIDToProviderID(ctx context.Context, id string) string {
	_ = "STUB: not implemented"
	return ""
}

// for historical reasons Azure providerID has the resource group name in lower case

// fallback to original providerID

// WithDefaultFloat64 returns the float64 value of the supplied environment variable or, if not present,
// the supplied default value. If the float64 conversion fails, returns the default
func WithDefaultFloat64(key string, def float64) float64 { _ = "STUB: not implemented"; return 0 }

func ImageReferenceToString(imageRef *armcompute.ImageReference) string {
	_ = "STUB: not implemented"
	// Check for Custom Image
	return ""
}

// Check for Community Image

// Check for Shared Gallery Image

// Check for Platform Image and use standard string representation

// Use the standard format: Publisher:Offer:Sku:Version

func IsVMDeleting(vm armcompute.VirtualMachine) bool { _ = "STUB: not implemented"; return false }

// StringMap returns the string map representation of the resource list
func StringMap(list v1.ResourceList) map[string]string { _ = "STUB: not implemented"; return nil }

// PrettySlice truncates a slice after a certain number of max items to ensure
// that the Slice isn't too long
func PrettySlice[T any](s []T, maxItems int) string { _ = "STUB: not implemented"; return "" }

// GetMaxPods resolves what we should set max pods to for a given nodeclass.
// If not specified, defaults based on network-plugin. 30 for "azure", 110 for "kubenet",
// or 250 for "none" and network plugin mode overlay.
func GetMaxPods(nodeClass *v1beta1.AKSNodeClass, networkPlugin, networkPluginMode string) int32 {
	_ = "STUB: not implemented"
	return 0
}

var managedVNETPattern = regexp.MustCompile(`(?i)^aks-vnet-\d{8}$`)

const managedSubnetName = "aks-subnet"

// IsAKSManagedVNET determines if the vnet managed or not.
// Note: You can "trick" this function if you really try by (for example) createding a VNET that looks like
// an AKS managed VNET, with the same resource group as the MC RG, in a different subscription, or by creating
// your own VNET in the MC RG whose name matches the AKS pattern but the VNET is actually yours rather than ours.
func IsAKSManagedVNET(nodeResourceGroup string, subnetID string) (bool, error) {
	_ = "STUB: not implemented"
	// TODO: I kinda think we should be using arm.ParseResourceID rather than rolling our own
	return false, nil
}

// HasChanged returns if the given value has changed, given the existing and new instance
//
// This option is accessible in place of using a ChangeMonitor, when there's access to both
// the existing and new data.
func HasChanged(existing, new any, options *hashstructure.HashOptions) bool {
	_ = "STUB: not implemented"
	// In the case of errors, the zero value from hashing will be compared, similar to ChangeMonitor
	return false
}

// GetAlphanumericHash generates a base36 alphanumeric hash of the input string with the specified length.
// Be mindful of collision risks with short lengths. Also note that length > 13 provides no additional
// collision resistance because the underlying hashstructure library returns a 64-bit hash, which only
// fills ~13 base36 characters; extra characters are just leading zeros.
// At the time of writing, this is being used in AKS machine instance provider/GetAKSMachineNameFromNodeClaimName(). See that for context.
func GetAlphanumericHash(input string, length int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Reverse order to have the same sense of significance as normal text

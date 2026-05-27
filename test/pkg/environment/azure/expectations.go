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

package azure

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
	containerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

func (env *Environment) EventuallyExpectKarpenterNicsToBeDeleted() {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ExpectCreatedInterface(networkInterface armnetwork.Interface) {
	_ = "STUB: not implemented"
	return
}

func (env *Environment) ExpectGetManagedCluster() *containerservice.ManagedCluster {
	_ = "STUB: not implemented"
	return nil
}

// ExpectClusterProvisioningState checks that the cluster's provisioning state matches the expected state,
// and fails the test if it does not.
func (env *Environment) ExpectClusterProvisioningState(expectedProvisioningState string) *containerservice.ManagedCluster {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) ExpectSuccessfulGetOfAvailableKubernetesVersionUpgradesForManagedCluster() []*containerservice.ManagedClusterPoolUpgradeProfileUpgradesItem {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) ExpectSuccessfulUpgradeOfManagedCluster(kubernetesUpgradeVersion string) *containerservice.ManagedCluster {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) ExpectUpgradeOfManagedCluster(kubernetesUpgradeVersion string) *runtime.Poller[containerservice.ManagedClustersClientCreateOrUpdateResponse] {
	_ = "STUB: not implemented"
	return nil
}

// See documentation for KubernetesVersion (client specified) and CurrentKubernetesVersion (version under use):
// https://learn.microsoft.com/en-us/rest/api/aks/managed-clusters/get?view=rest-aks-2025-01-01&tabs=HTTP

// Note that this is an update not a create so we don't need to add it to the tracker

func (env *Environment) ExpectParsedProviderID(providerID string) string {
	_ = "STUB: not implemented"
	return ""
}

func (env *Environment) ExpectCreatedSubnet(vnetName string, subnet *armnetwork.Subnet) {
	_ = "STUB: not implemented"
	return
}

// EventuallyExpectTags checks that all of the resources in the resource group have the expected tags.
func (env *Environment) EventuallyExpectTags(expectedTags map[string]string) {
	_ = "STUB: not implemented"

	// Convert the expectedTags to ptrs
	return
}

// Check extension tags

// EventuallyExpectMissingTags checks that all of the resources in the resource group are missing the expected tags.
func (env *Environment) EventuallyExpectMissingTags(expectedMissingTags map[string]string) {
	_ = "STUB: not implemented"
	return
}

// Check extension tags

func (env *Environment) EventuallyExpectAzureResources(
	verifyNIC func(nic *armnetwork.Interface) error,
	verifyVM func(vm *armcompute.VirtualMachine) error,
	verifyExt func(ext *armcompute.VirtualMachineExtension) error,
) {
	_ = "STUB: not implemented"
	return
}

// NICs

// Ignore nodes that don't have the expected Karpenter tag

// Note that disks also exist, but are automatically created and managed by Azure so we don't check them here.

// VMs

// Ignore nodes that don't have the expected Karpenter tag

// Extensions

// Only check extensions are that managed by Karpenter

func isMapSubset[K comparable, V comparable](m map[K]V, subset map[K]V, eq func(a, b V) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func eqPtr(v1, v2 *string) bool { _ = "STUB: not implemented"; return false }

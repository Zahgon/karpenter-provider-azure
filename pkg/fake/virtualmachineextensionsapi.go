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

package fake

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
	fakesync "github.com/Azure/karpenter-provider-azure/pkg/fake/sync"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/azclient/azapi"
)

type VirtualMachineExtensionCreateOrUpdateInput struct {
	ResourceGroupName           string
	VirtualMachineName          string
	VirtualMachineExtensionName string
	VirtualMachineExtension     armcompute.VirtualMachineExtension
	Options                     *armcompute.VirtualMachineExtensionsClientBeginCreateOrUpdateOptions
}

type VirtualMachineExtensionUpdateInput struct {
	ResourceGroupName             string
	VirtualMachineName            string
	VirtualMachineExtensionName   string
	VirtualMachineExtensionUpdate armcompute.VirtualMachineExtensionUpdate
	Options                       *armcompute.VirtualMachineExtensionsClientBeginUpdateOptions
}

type VirtualMachineExtensionsBehavior struct {
	VirtualMachineExtensionsCreateOrUpdateBehavior MockedLRO[VirtualMachineExtensionCreateOrUpdateInput, armcompute.VirtualMachineExtensionsClientCreateOrUpdateResponse]
	VirtualMachineExtensionsUpdateBehavior         MockedLRO[VirtualMachineExtensionUpdateInput, armcompute.VirtualMachineExtensionsClientUpdateResponse]
	Extensions                                     fakesync.Map[string, armcompute.VirtualMachineExtension]
}

// assert that ComputeAPI implements ARMComputeAPI
var _ azapi.VirtualMachineExtensionsAPI = &VirtualMachineExtensionsAPI{}

type VirtualMachineExtensionsAPI struct {
	// azapi.VirtualMachineExtensionsAPI
	VirtualMachineExtensionsBehavior
}

// Reset must be called between tests otherwise tests will pollute each other.
func (c *VirtualMachineExtensionsAPI) Reset() { _ = "STUB: not implemented"; return }

func (c *VirtualMachineExtensionsAPI) BeginCreateOrUpdate(
	_ context.Context,
	resourceGroupName,
	vmName,
	extensionName string,
	extension armcompute.VirtualMachineExtension,
	options *armcompute.VirtualMachineExtensionsClientBeginCreateOrUpdateOptions,
) (*runtime.Poller[armcompute.VirtualMachineExtensionsClientCreateOrUpdateResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// only store latest, but could be improved

func (c *VirtualMachineExtensionsAPI) BeginUpdate(
	_ context.Context,
	resourceGroupName string,
	vmName string,
	extensionName string,
	updates armcompute.VirtualMachineExtensionUpdate,
	options *armcompute.VirtualMachineExtensionsClientBeginUpdateOptions,
) (*runtime.Poller[armcompute.VirtualMachineExtensionsClientUpdateResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// VM tags are full-replace if they're specified

func MakeVMExtensionID(resourceGroupName, vmName, extensionName string) string {
	_ = "STUB: not implemented"
	return ""
}

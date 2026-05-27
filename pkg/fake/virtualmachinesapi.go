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
	"io"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
	"github.com/Azure/karpenter-provider-azure/pkg/auth"
	fakesync "github.com/Azure/karpenter-provider-azure/pkg/fake/sync"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/azclient/azapi"
)

type VirtualMachineCreateOrUpdateInput struct {
	ResourceGroupName string
	VMName            string
	VM                armcompute.VirtualMachine
	Options           *armcompute.VirtualMachinesClientBeginCreateOrUpdateOptions
}

type VirtualMachineUpdateInput struct {
	ResourceGroupName string
	VMName            string
	Updates           armcompute.VirtualMachineUpdate
	Options           *armcompute.VirtualMachinesClientBeginUpdateOptions
}

type VirtualMachineDeleteInput struct {
	ResourceGroupName string
	VMName            string
	Options           *armcompute.VirtualMachinesClientBeginDeleteOptions
}

type VirtualMachineGetInput struct {
	ResourceGroupName string
	VMName            string
	Options           *armcompute.VirtualMachinesClientGetOptions
}

type VirtualMachinesBehavior struct {
	VirtualMachineCreateOrUpdateBehavior MockedLRO[VirtualMachineCreateOrUpdateInput, armcompute.VirtualMachinesClientCreateOrUpdateResponse]
	VirtualMachineUpdateBehavior         MockedLRO[VirtualMachineUpdateInput, armcompute.VirtualMachinesClientUpdateResponse]
	VirtualMachineDeleteBehavior         MockedLRO[VirtualMachineDeleteInput, armcompute.VirtualMachinesClientDeleteResponse]
	VirtualMachineGetBehavior            MockedFunction[VirtualMachineGetInput, armcompute.VirtualMachinesClientGetResponse]
	Instances                            fakesync.Map[string, armcompute.VirtualMachine]
}

// assert that the fake implements the interface
var _ azapi.VirtualMachinesAPI = &VirtualMachinesAPI{}

type VirtualMachinesAPI struct {
	// TODO: document the implications of embedding vs. not embedding the interface here
	// azapi.VirtualMachinesAPI // - this is the interface we are mocking.
	VirtualMachinesBehavior
	AuxiliaryTokenPolicy *auth.AuxiliaryTokenPolicy
}

// Reset must be called between tests otherwise tests will pollute each other.
func (c *VirtualMachinesAPI) Reset() { _ = "STUB: not implemented"; return }

// UseAuxiliaryTokenPolicy simulates AuxiliaryTokenPolicy.Do() method being called at the beginning of each API call
// This is useful for testing scenarios where the auxiliary token is required for the API call to succeed.
// If the AuxiliaryTokenPolicy is not set (USE_SIG: false), this method does nothing and returns nil.
func (c *VirtualMachinesAPI) UseAuxiliaryTokenPolicy() error { _ = "STUB: not implemented"; return nil }

// req.Next() returns this if there are no more policies.

func (c *VirtualMachinesAPI) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vmName string, parameters armcompute.VirtualMachine, options *armcompute.VirtualMachinesClientBeginCreateOrUpdateOptions) (*runtime.Poller[armcompute.VirtualMachinesClientCreateOrUpdateResponse], error) {
	_ = "STUB: not implemented"
	// gather input parameters (may get rid of this with multiple mocked function signatures to reflect common patterns)
	return nil, nil
}

// BeginCreateOrUpdate should fail, if the vm exists in the cache, and we are attempting to change properties for zone

//if input.ResourceGroupName == "" {
//	return nil, errors.New("ResourceGroupName is required")
//}
// TODO: may have to clone ...
// TODO: subscription ID?

// Check store for existing vm by name

// Note: this assumes at least 1 zone and only one zone is put on our vm

// Currently only returning for zones, but osProfile.customData will also return this error

// Use existing vm rather than restoring

// TODO: use simulated time?

func (c *VirtualMachinesAPI) BeginUpdate(_ context.Context, resourceGroupName string, vmName string, updates armcompute.VirtualMachineUpdate, options *armcompute.VirtualMachinesClientBeginUpdateOptions) (*runtime.Poller[armcompute.VirtualMachinesClientUpdateResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If other fields need to be updated in the future, you can similarly
// update the VM object by merging with updates.<New Field>.

// VM tags are full-replace if they're specified

// Update the stored shape

func (c *VirtualMachinesAPI) Get(_ context.Context, resourceGroupName string, vmName string, options *armcompute.VirtualMachinesClientGetOptions) (armcompute.VirtualMachinesClientGetResponse, error) {
	_ = "STUB: not implemented"
	return *new(armcompute.VirtualMachinesClientGetResponse), nil
}

func (c *VirtualMachinesAPI) BeginDelete(_ context.Context, resourceGroupName string, vmName string, options *armcompute.VirtualMachinesClientBeginDeleteOptions) (*runtime.Poller[armcompute.VirtualMachinesClientDeleteResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CreateSDKErrorBody(code, message string) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

func MkVMID(resourceGroupName string, vmName string) string { _ = "STUB: not implemented"; return "" }

func getAuthTokenError(err error) *azcore.ResponseError { _ = "STUB: not implemented"; return nil }

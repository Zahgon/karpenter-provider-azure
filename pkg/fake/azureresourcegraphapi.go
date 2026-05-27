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

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/azclient/azapi"
)

type AzureResourceGraphResourcesInput struct {
	Query   armresourcegraph.QueryRequest
	Options *armresourcegraph.ClientResourcesOptions
}

type AzureResourceGraphBehavior struct {
	AzureResourceGraphResourcesBehavior MockedFunction[AzureResourceGraphResourcesInput, armresourcegraph.ClientResourcesResponse]
	VirtualMachinesAPI                  *VirtualMachinesAPI
	NetworkInterfacesAPI                *NetworkInterfacesAPI
	ResourceGroup                       string
}

// assert that the fake implements the interface
var _ azapi.AzureResourceGraphAPI = &AzureResourceGraphAPI{}

type AzureResourceGraphAPI struct {
	vmListQuery  string
	nicListQuery string
	AzureResourceGraphBehavior
}

func NewAzureResourceGraphAPI(resourceGroup string, virtualMachinesAPI *VirtualMachinesAPI, networkInterfacesAPI *NetworkInterfacesAPI) *AzureResourceGraphAPI {
	_ = "STUB: not implemented"
	return nil
}

// Reset must be called between tests otherwise tests will pollute each other.
func (c *AzureResourceGraphAPI) Reset() { _ = "STUB: not implemented"; return }

func (c *AzureResourceGraphAPI) Resources(_ context.Context, query armresourcegraph.QueryRequest, options *armresourcegraph.ClientResourcesOptions) (armresourcegraph.ClientResourcesResponse, error) {
	_ = "STUB: not implemented"
	return *new(armresourcegraph.ClientResourcesResponse), nil
}

func (c *AzureResourceGraphAPI) getResourceList(query string) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (c *AzureResourceGraphAPI) loadVMObjects() (vmList []armcompute.VirtualMachine) {
	_ = "STUB: not implemented"
	return nil
}

func (c *AzureResourceGraphAPI) loadNicObjects() (nicList []armnetwork.Interface) {
	_ = "STUB: not implemented"
	return nil
}

func convertBytesToInterface(b []byte) interface{} { _ = "STUB: not implemented"; return nil }

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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	fakesync "github.com/Azure/karpenter-provider-azure/pkg/fake/sync"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/azclient/azapi"
)

type NetworkInterfaceCreateOrUpdateInput struct {
	ResourceGroupName string
	InterfaceName     string
	Interface         armnetwork.Interface
	Options           *armnetwork.InterfacesClientBeginCreateOrUpdateOptions
}

type NetworkInterfaceDeleteInput struct {
	ResourceGroupName, InterfaceName string
}

type NetworkInterfaceUpdateTagsInput struct {
	ResourceGroupName string
	InterfaceName     string
	Tags              armnetwork.TagsObject
	Options           *armnetwork.InterfacesClientUpdateTagsOptions
}

type NetworkInterfacesBehavior struct {
	NetworkInterfacesCreateOrUpdateBehavior MockedLRO[NetworkInterfaceCreateOrUpdateInput, armnetwork.InterfacesClientCreateOrUpdateResponse]
	NetworkInterfacesDeleteBehavior         MockedLRO[NetworkInterfaceDeleteInput, armnetwork.InterfacesClientDeleteResponse]
	NetworkInterfacesUpdateTagsBehavior     MockedFunction[NetworkInterfaceUpdateTagsInput, armnetwork.InterfacesClientUpdateTagsResponse]
	NetworkInterfaces                       fakesync.Map[string, armnetwork.Interface]
}

// assert that the fake implements the interface
var _ azapi.NetworkInterfacesAPI = &NetworkInterfacesAPI{}

type NetworkInterfacesAPI struct {
	// azapi.NetworkInterfacesAPI
	NetworkInterfacesBehavior
}

// Reset must be called between tests otherwise tests will pollute each other.
func (c *NetworkInterfacesAPI) Reset() { _ = "STUB: not implemented"; return }

func (c *NetworkInterfacesAPI) BeginCreateOrUpdate(_ context.Context, resourceGroupName string, interfaceName string, iface armnetwork.Interface, options *armnetwork.InterfacesClientBeginCreateOrUpdateOptions) (*runtime.Poller[armnetwork.InterfacesClientCreateOrUpdateResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *NetworkInterfacesAPI) Get(_ context.Context, resourceGroupName string, interfaceName string, _ *armnetwork.InterfacesClientGetOptions) (armnetwork.InterfacesClientGetResponse, error) {
	_ = "STUB: not implemented"
	return *new(armnetwork.InterfacesClientGetResponse), nil
}

func (c *NetworkInterfacesAPI) BeginDelete(_ context.Context, resourceGroupName string, interfaceName string, _ *armnetwork.InterfacesClientBeginDeleteOptions) (*runtime.Poller[armnetwork.InterfacesClientDeleteResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *NetworkInterfacesAPI) UpdateTags(
	ctx context.Context,
	resourceGroupName string,
	interfaceName string,
	tags armnetwork.TagsObject,
	options *armnetwork.InterfacesClientUpdateTagsOptions,
) (armnetwork.InterfacesClientUpdateTagsResponse, error) {
	_ = "STUB: not implemented"
	return *new(armnetwork.InterfacesClientUpdateTagsResponse), nil
}

// Tags are full-replace if they're specified

func MakeNetworkInterfaceID(resourceGroupName, interfaceName string) string {
	_ = "STUB: not implemented"
	return ""
}

// not important for fake

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
	"github.com/Azure/karpenter-provider-azure/pkg/providers/networksecuritygroup"
)

type NetworkSecurityGroupBevhavior struct {
	NSGs fakesync.Map[string, armnetwork.SecurityGroup]
}

// assert that the fake implements the interface
var _ networksecuritygroup.API = &NetworkSecurityGroupAPI{}

type NetworkSecurityGroupAPI struct {
	NetworkSecurityGroupBevhavior
}

// Reset must be called between tests otherwise tests will pollute each other.
func (api *NetworkSecurityGroupAPI) Reset() {
	_ = "STUB: not implemented"

	// Get implements networksecuritygroup.API.
	return
}

func (api *NetworkSecurityGroupAPI) Get(
	ctx context.Context,
	resourceGroupName string,
	securityGroupName string,
	options *armnetwork.SecurityGroupsClientGetOptions,
) (armnetwork.SecurityGroupsClientGetResponse, error) {
	_ = "STUB: not implemented"
	return *new(armnetwork.SecurityGroupsClientGetResponse), nil
}

// NewListPager implements networksecuritygroup.API.
func (api *NetworkSecurityGroupAPI) NewListPager(resourceGroupName string, options *armnetwork.SecurityGroupsClientListOptions) *runtime.Pager[armnetwork.SecurityGroupsClientListResponse] {
	_ = "STUB: not implemented"
	return nil
}

// TODO: It might be ideal if we had a MockPager which sometimes simulated multiple pages of results to ensure we handle that correctly

// Sort the result according to ID so that we have a stable base to write asserts upon

func MakeNetworkSecurityGroupID(resourceGroupName, networkSecurityGroupName string) string {
	_ = "STUB: not implemented"
	return ""
}

// TODO: This is duplicated from other places, we should consider putting it in a common place

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
	"github.com/Azure/karpenter-provider-azure/pkg/providers/loadbalancer"
)

type LoadBalancersBehavior struct {
	LoadBalancers fakesync.Map[string, armnetwork.LoadBalancer]
}

// assert that the fake implements the interface
var _ loadbalancer.LoadBalancersAPI = &LoadBalancersAPI{}

type LoadBalancersAPI struct {
	LoadBalancersBehavior
}

// Reset must be called between tests otherwise tests will pollute each other.
func (api *LoadBalancersAPI) Reset() { _ = "STUB: not implemented"; return }

func (api *LoadBalancersAPI) Get(_ context.Context, resourceGroupName string, loadBalancerName string, _ *armnetwork.LoadBalancersClientGetOptions) (armnetwork.LoadBalancersClientGetResponse, error) {
	_ = "STUB: not implemented"
	return *new(armnetwork.LoadBalancersClientGetResponse), nil
}

func (api *LoadBalancersAPI) NewListPager(_ string, _ *armnetwork.LoadBalancersClientListOptions) *runtime.Pager[armnetwork.LoadBalancersClientListResponse] {
	_ = "STUB: not implemented"
	return nil
}

// TODO: It might be ideal if we had a MockPager which sometimes simulated multiple pages of results to ensure we handle that correctly

// Sort the result according to ID so that we have a stable base to write asserts upon

func MakeLoadBalancerID(resourceGroupName, loadBalancerName string) string {
	_ = "STUB: not implemented"
	return ""
}

// not important for fake

func MakeBackendAddressPoolID(resourceGroupName, loadBalancerName string, backendAddressPoolName string) string {
	_ = "STUB: not implemented"
	return ""
}

// not important for fake

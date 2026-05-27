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
	_ "embed"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	fakesync "github.com/Azure/karpenter-provider-azure/pkg/fake/sync"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/zone"
)

//go:embed locations.json
var fakeLocationsJSON string

type ListLocationsInput struct {
	SubscriptionID string
	Options        *armsubscriptions.ClientListLocationsOptions
}

type SubscriptionsAPIBehavior struct {
	NewListLocationsPagerBehavior MockedFunction[ListLocationsInput, armsubscriptions.ClientListLocationsResponse]
	Locations                     fakesync.Map[string, armsubscriptions.Location]
	UseFakeData                   bool
}

var _ zone.SubscriptionsAPI = &SubscriptionsAPI{}

type SubscriptionsAPI struct {
	SubscriptionsAPIBehavior
}

func NewSubscriptionsAPI() (*SubscriptionsAPI, error) { _ = "STUB: not implemented"; return nil, nil }

func (api *SubscriptionsAPI) Reset() { _ = "STUB: not implemented"; return }

// Not ideal, but shouldn't happen

func (api *SubscriptionsAPI) NewListLocationsPager(
	subscriptionID string,
	options *armsubscriptions.ClientListLocationsOptions,
) *runtime.Pager[armsubscriptions.ClientListLocationsResponse] {
	_ = "STUB: not implemented"
	return nil
}

// TODO: It might be ideal if we had a MockPager which sometimes simulated multiple pages of results to ensure we handle that correctly

// Sort the result according to Name so that we have a stable base to write asserts upon

func loadLocationsFromFile(api *SubscriptionsAPI) error { _ = "STUB: not implemented"; return nil }

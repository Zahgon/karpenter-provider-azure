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

package pricing

import (
	"context"
	"net/http"
	"sync"
	"time"

	"sigs.k8s.io/karpenter/pkg/utils/pretty"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/karpenter-provider-azure/pkg/auth"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/pricing/client"
)

// pricingUpdatePeriod is how often we try to update our pricing information after the initial update on startup
const pricingUpdatePeriod = 12 * time.Hour

const defaultRegion = "eastus"

// Provider provides actual pricing data to the Azure cloud provider to allow it to make more informed decisions
// regarding which instances to launch.  This is initialized at startup with a periodically updated static price list to
// support running in locations where pricing data is unavailable.  In those cases the static pricing data provides a
// relative ordering that is still more accurate than our previous pricing model.  In the event that a pricing update
// fails, the previous pricing information is retained and used which may be the static initial pricing data if pricing
// updates never succeed.
type Provider struct {
	pricing client.PricingAPI
	region  string
	cm      *pretty.ChangeMonitor

	mu                 sync.RWMutex
	onDemandUpdateTime time.Time
	onDemandPrices     map[string]float64
	spotUpdateTime     time.Time
	spotPrices         map[string]float64
	done               chan struct{}
}

// NewPricingAPI returns a pricing API
func NewAPI(cloud cloud.Configuration) client.PricingAPI {
	_ = "STUB: not implemented"
	return *new(client.PricingAPI)
}

func NewProvider(
	ctx context.Context,
	env *auth.Environment,
	pricing client.PricingAPI,
	region string,
	startAsync <-chan struct{},
) *Provider {
	_ = "STUB: not implemented"
	// see if we've got region specific pricing data
	return nil
}

// and if not, fall back to the always available eastus

// default our spot pricing to the same as the on-demand pricing until a price update

// Only poll in public cloud. Other clouds aren't supported currently

// perform an initial price update at startup

// wait for leader election or to be signaled to exit

// if it took many hours to be elected leader, we want to re-fetch pricing before we start our periodic
// polling

// done immediately

// InstanceTypes returns the list of all instance types for which either a price is known.
func (p *Provider) InstanceTypes() []string { _ = "STUB: not implemented"; return nil }

// OnDemandLastUpdated returns the time that the on-demand pricing was last updated
func (p *Provider) OnDemandLastUpdated() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// SpotLastUpdated returns the time that the spot pricing was last updated
func (p *Provider) SpotLastUpdated() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// OnDemandPrice returns the last known on-demand price for a given instance type, returning false if there is no
// known on-demand pricing for the instance type.
func (p *Provider) OnDemandPrice(instanceType string) (float64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// if we don't have a price, check if it's a known SKU with missing price

// SpotPrice returns the last known spot price for a given instance type, returning false
// if there is no known spot pricing for that instance type
func (p *Provider) SpotPrice(instanceType string) (float64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// if we don't have a price, check if it's a known SKU with missing price

func (p *Provider) updatePricing(ctx context.Context) { _ = "STUB: not implemented"; return }

// FetchPricing fetches VM pricing from the Azure retail pricing API for the given region,
// returning on-demand and spot prices keyed by ARM SKU name.
func FetchPricing(ctx context.Context, pricingAPI client.PricingAPI, region string) (onDemandPrices, spotPrices map[string]float64, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func processPage(prices map[client.Item]bool) func(page *client.ProductsPricePage) {
	_ = "STUB: not implemented"
	return nil
}

// https://learn.microsoft.com/en-us/azure/batch/batch-spot-vms#differences-between-spot-and-low-priority-vms

func categorizePrices(prices map[client.Item]bool) (map[string]float64, map[string]float64) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Provider) LivenessProbe(_ *http.Request) error {
	_ = "STUB: not implemented"
	// ensure we don't deadlock and nolint for the empty critical section
	return nil
}

//nolint: staticcheck

func (p *Provider) Reset() {
	_ = "STUB: not implemented"
	// see if we've got region specific pricing data
	return
}

// and if not, fall back to the always available eastus

// WaitUntilDone should be called after canceling the context passed to NewProvider to wait until all goroutines have exited
func (p *Provider) WaitUntilDone() error { _ = "STUB: not implemented"; return nil }

func Regions() []string { _ = "STUB: not implemented"; return nil }

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

package zone

import (
	"context"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
)

// We don't want to retry too aggressively here because this API is somewhat slow,
// but at the same time we want to wake back up eventually and try again in the case of an outage.
// These values were picked somewhat arbitrarily to achieve that.
const (
	maxFailuresPerWindow = 10
	windowBackoff        = 60 * time.Minute
)

type Clock interface {
	Now() time.Time
}

// SubscriptionsAPI defines the interface for Azure Subscriptions client operations
type SubscriptionsAPI interface {
	NewListLocationsPager(
		subscriptionID string,
		options *armsubscriptions.ClientListLocationsOptions,
	) *runtime.Pager[armsubscriptions.ClientListLocationsResponse]
}

// Provider handles zone support detection for Azure regions
// TODO: This provider is currently unused. Keeping it around for now though as we will likely want to adapt it
// to provide physical to logical zone mappings.
type Provider struct {
	subscriptionsAPI SubscriptionsAPI
	subscriptionID   string
	clock            Clock

	// Cached zone list data - maps region name to list of available zones
	zoneList  map[string][]string
	hasLoaded bool
	// failures is the number of times loading zone support from the Azure API has failed
	failures    int
	lastAttempt time.Time
	mu          sync.Mutex
}

// NewProvider creates a new zone provider
func NewProvider(
	subscriptionsAPI SubscriptionsAPI,
	clock Clock,
	subscriptionID string,
) *Provider {
	_ = "STUB: not implemented"
	return nil
}

// SupportsZones returns true if the given region supports availability zones
func (p *Provider) SupportsZones(ctx context.Context, region string) bool {
	_ = "STUB: not implemented"
	return false
}

// GetAvailableZones returns the list of available zones for a given region.
// Returns nil if the region doesn't support zones or zone data hasn't been loaded.
func (p *Provider) GetAvailableZones(ctx context.Context, region string) []string {
	_ = "STUB: not implemented"
	return nil
}

// ensureLoadedLocked attempts to load zone data from Azure API if not already loaded.
// Must be called with p.mu held.
func (p *Provider) ensureLoadedLocked(ctx context.Context) {
	_ = "STUB: not implemented"
	// NOTE: We considered doing this in a separate goroutine or inline on provider construction but
	// we want:
	// 1. To block provisioning until we've at least attempted to load zone support data from the API once.
	// 2. To avoid blocking provisioning forever if the API is unavailable.
	// It seems like this is the simplest way to accomplish that.
	return
}

// loadFromAzure discovers zone support by calling Azure Subscriptions API
func (p *Provider) loadFromAzure(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// shouldTryAgain determines if the provider should attempt to load zone support data again
// after failures have happened.
func (p *Provider) shouldTryAgain() bool { _ = "STUB: not implemented"; return false }

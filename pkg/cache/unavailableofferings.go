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

package cache

import (
	"context"
	"time"

	"github.com/Azure/skewer"
	"github.com/patrickmn/go-cache"
	karpv1 "sigs.k8s.io/karpenter/pkg/apis/v1"
)

const (
	// wholeVMFamilyBlockedSentinel means that entire SKU family is blocked, not just certain instance types with a CPU count above a threshold
	wholeVMFamilyBlockedSentinel = -1
)

var (
	spotKey = singleInstanceKey("", "", karpv1.CapacityTypeSpot)
)

// UnavailableOfferings stores any offerings that return ICE (insufficient capacity errors) when
// attempting to launch the capacity. These offerings are ignored as long as they are in the cache on
// GetInstanceTypes responses
// Information available from skewer.SKU is used to determine details about the VM SKU for which we encountered allocation errors.
type UnavailableOfferings struct {
	// TODO: I think this singleOfferingCache could basically be removed in favor of the family cache at this point, as we are now marking family unavailable at CPU count for all error cases.
	// I didn't do that purely because of the defensive fallback we have in markFamilyUnavailableAtCPUCountImpl, but in practice I am not sure if we'll ever have a nil family (I don't
	// see any evidence it can happen in logs)
	// key: <capacityType>:<instanceType>:<zone>, value: struct{}{}
	singleOfferingCache *cache.Cache
	// key: <skuFamilyName>:<zone>:<capacityType> (lowercase), value: int64 (CPU count at or above which we block, or wholeVMFamilyBlockedSentinel if entire family is blocked)
	vmFamilyCache *cache.Cache
	SeqNum        uint64
}

func NewUnavailableOfferingsWithCache(singleOfferingCache, vmFamilyCache *cache.Cache) *UnavailableOfferings {
	_ = "STUB: not implemented"
	return nil
}

func NewUnavailableOfferings() *UnavailableOfferings { _ = "STUB: not implemented"; return nil }

// IsUnavailable returns true if the offering appears in the cache
func (u *UnavailableOfferings) IsUnavailable(sku *skewer.SKU, zone, capacityType string) bool {
	_ = "STUB: not implemented"
	return false
}

// check if the offering is marked as unavailable at vm family level

// lastly check if the offering is marked as unavailable for the specific instance type, zone and capacity type

func (u *UnavailableOfferings) isFamilyUnavailable(sku *skewer.SKU, zone, capacityType string) bool {
	_ = "STUB: not implemented"
	return false
}

// default to 0 if we can't determine VCPU count, this shouldn't happen as long as data in skewer.SKU is correct

// Check if VM family is blocked in the specific zone

// Entire VM family is blocked in this zone

// VM sizes from this family are blocked for CPU counts >= blockedCPUCount in this zone

// markFamilyUnavailableAtCPUCount marks a VM family with custom TTL in a specific zone for all instance types that have CPU count at or above the SKU's vCPU count.
// Information is derived from the provided skewer.SKU: family name via GetFamilyName() and CPU count via VCPU().
func (u *UnavailableOfferings) markFamilyUnavailableAtCPUCount(ctx context.Context, sku *skewer.SKU, zone, capacityType string, ttl time.Duration) {
	_ = "STUB: not implemented"
	return
}

// default to 0 if we can't determine VCPU count, this shouldn't happen as long as data in skewer.SKU is correct

// MarkFamilyUnavailable marks the entire VM family as unavailable in a specific zone for a specific capacity type with custom TTL.
// Family name is derived from the provided skewer.SKU.
func (u *UnavailableOfferings) MarkFamilyUnavailable(ctx context.Context, sku *skewer.SKU, zone, capacityType string, ttl time.Duration) {
	_ = "STUB: not implemented"
	return
}

// markFamilyUnavailableAtCPUCountImpl is the internal implementation that marks a VM family unavailable at a given CPU count threshold.
// Value of -1 is used as a "wholeVMFamilyBlockedSentinel" to indicate that the entire VM family is blocked in this zone for the specified capacity type.
func (u *UnavailableOfferings) markFamilyUnavailableAtCPUCountImpl(ctx context.Context, sku *skewer.SKU, zone, capacityType string, cpuCount int64, ttl time.Duration) {
	_ = "STUB: not implemented"
	return
}

// This is a hedge against skewer having bad data where family name is missing,
// If family name is missing, we won't do any family level blocking, but we'll still mark the specific offering as unavailable.

// Keep the more restrictive limit for CPU count(lower value, with -1 being most restrictive - wholeVMFamilyBlockedSentinel)

// call Set to update the cache entry, even if it already exists, to extend its TTL

// MarkSpotUnavailable communicates recently observed temporary capacity shortages for spot
func (u *UnavailableOfferings) MarkSpotUnavailableWithTTL(ctx context.Context, ttl time.Duration) {
	_ = "STUB: not implemented"
	return
}

// even if the key is already in the cache, we still need to call Set to extend the cached entry's TTL

// MarkUnavailableWithTTL allows us to mark an offering unavailable with a custom TTL.
// In addition to marking the specific instance type unavailable, it also marks the VM family
// unavailable at the SKU's vCPU count, so that larger sizes of the same family are also blocked.
func (u *UnavailableOfferings) MarkUnavailableWithTTL(ctx context.Context, unavailableReason string, sku *skewer.SKU, zone, capacityType string, ttl time.Duration) {
	_ = "STUB: not implemented"
	return
}

// even if the key is already in the cache, we still need to call Set to extend the cached entry's TTL

// Also mark the VM family unavailable at this SKU's vCPU count, so larger sizes of the same family are blocked too

// MarkUnavailable communicates recently observed temporary capacity shortages in the provided offerings
func (u *UnavailableOfferings) MarkUnavailable(ctx context.Context, unavailableReason string, sku *skewer.SKU, zone, capacityType string) {
	_ = "STUB: not implemented"
	return
}

func (u *UnavailableOfferings) Flush() { _ = "STUB: not implemented"; return }

// singleInstanceKey returns the cache singleInstanceKey for all offerings in the cache
func singleInstanceKey(instanceType string, zone string, capacityType string) string {
	_ = "STUB: not implemented"
	return ""
}

// vmFamilyKey returns the cache key for VM family blocks in a specific zone
func vmFamilyKey(skuFamilyName, zone, capacityType string) string {
	_ = "STUB: not implemented"
	return ""
}

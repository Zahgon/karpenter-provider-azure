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

package instancetype

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
	"github.com/patrickmn/go-cache"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
	kcache "github.com/Azure/karpenter-provider-azure/pkg/cache"

	"github.com/Azure/karpenter-provider-azure/pkg/providers/pricing"

	"github.com/Azure/skewer"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/utils/pretty"
)

const (
	InstanceTypesCacheTTL = 23 * time.Hour
)

type Provider interface {
	LivenessProbe(*http.Request) error
	List(context.Context, *v1beta1.AKSNodeClass) ([]*cloudprovider.InstanceType, error)

	// Return Azure Skewer Representation of the instance type
	Get(context.Context, string) (*skewer.SKU, error)

	// UpdateInstanceTypes fetches instance types from Azure and updates the cache
	UpdateInstanceTypes(ctx context.Context) error

	// UpdateInstanceTypeOfferings(ctx context.Context) error
}

// assert that DefaultProvider implements Provider interface
var _ Provider = (*DefaultProvider)(nil)

type DefaultProvider struct {
	region               string
	skuClient            skewer.ResourceClient
	pricingProvider      *pricing.Provider
	unavailableOfferings *kcache.UnavailableOfferings

	// Values cached *before* considering insufficient capacity errors from the unavailableOfferings cache.
	// Fully initialized Instance Types are also cached based on the set of all instance types,
	// unavailableOfferings cache, AWSNodeClass, and kubelet configuration from the NodePool
	instanceTypesCache *cache.Cache

	cm *pretty.ChangeMonitor

	// instanceTypesSeqNum is a monotonically increasing change counter used to avoid the expensive hashing operation on instance types
	instanceTypesSeqNum uint64
	muInstanceTypesInfo sync.RWMutex
	instanceTypesInfo   map[string]*skewer.SKU
}

func NewDefaultProvider(
	region string,
	cache *cache.Cache,
	skuClient skewer.ResourceClient,
	pricingProvider *pricing.Provider,
	offeringsCache *kcache.UnavailableOfferings,
) *DefaultProvider {
	_ = "STUB: not implemented"
	return nil

	// TODO: skewer api, subnetprovider, pricing provider, unavailable offerings, ...
}

// Get all instance type options
func (p *DefaultProvider) List(
	ctx context.Context,
	nodeClass *v1beta1.AKSNodeClass,
) ([]*cloudprovider.InstanceType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Compute fully initialized instance types hash key

// Ensure what's returned from this function is a shallow-copy of the slice (not a deep-copy of the data itself)
// so that modifications to the ordering of the data don't affect the original

// Get Viable offerings
// Azure has zones availability directly from SKU info

// !!! Important !!!
// Any changes to the values passed into the NewInstanceType method will require making updates to the cache key
// so that Karpenter is able to cache the set of InstanceTypes based on values that alter the set of instance types
// !!! Important !!!

func (p *DefaultProvider) LivenessProbe(req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *DefaultProvider) Get(ctx context.Context, instanceType string) (*skewer.SKU, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// instanceTypeZones generates the set of all supported zones for a given SKU
// The strings have to match Zone labels that will be placed on Node
func (p *DefaultProvider) instanceTypeZones(sku *skewer.SKU) sets.Set[string] {
	_ = "STUB: not implemented"
	// skewer returns numerical zones, like "1" (as keys in the map);
	// prefix each zone with "<region>-", to have them match the labels placed on Node (e.g. "westus2-1")
	// Note this data comes from LocationInfo, then skewer is used to get the SKU info
	// If an offering is regional (non-zonal), the availability zones will be empty.
	return nil
}

// Regional (non-zonal) SKUs use zone "0" to match the label AKS places on regional nodes
// (topology.kubernetes.io/zone=0).

// TODO: review; switch to controller-driven updates
// createOfferings creates a set of mutually exclusive offerings for a given instance type. This provider maintains an
// invariant that each offering is mutually exclusive. Specifically, there is an offering for each permutation of zone
// and capacity type. ZoneID is also injected into the offering requirements, when available, but there is a 1-1
// mapping between zone and zoneID so this does not change the number of offerings.
//
// Each requirement on the offering is guaranteed to have a single value. To get the value for a requirement on an
// offering, you can do the following thanks to this invariant:
//
//	offering.Requirements.Get(v1.TopologyLabelZone).Any()
func (p *DefaultProvider) createOfferings(sku *skewer.SKU, offeringZones sets.Set[string]) cloudprovider.Offerings {
	_ = "STUB: not implemented"
	return *new(cloudprovider.Offerings)
}

/*
	instanceTypeOfferingAvailable.With(prometheus.Labels{
		instanceTypeLabel: *instanceType.InstanceType,
		capacityTypeLabel: capacityType,
		zoneLabel:         zone,
	}).Set(float64(lo.Ternary(available, 1, 0)))
	instanceTypeOfferingPriceEstimate.With(prometheus.Labels{
		instanceTypeLabel: *instanceType.InstanceType,
		capacityTypeLabel: capacityType,
		zoneLabel:         zone,
	}).Set(price)
*/

// isInstanceTypeSupportedByFilters consolidates all per-NodeClass instance type
// filters into a single call to keep the List() method's cyclomatic complexity low.
func (p *DefaultProvider) isInstanceTypeSupportedByFilters(sku *skewer.SKU, architecture string, nodeClass *v1beta1.AKSNodeClass) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *DefaultProvider) isInstanceTypeSupportedByImageFamily(skuName, imageFamily string) bool {
	_ = "STUB: not implemented"
	// Non-GPU SKUs are supported by all image families
	return false
}

func (p *DefaultProvider) isInstanceTypeSupportedByEncryptionAtHost(sku *skewer.SKU, nodeClass *v1beta1.AKSNodeClass) bool {
	_ = "STUB: not implemented"
	// If EncryptionAtHost is not enabled in the nodeclass, all instance types are supported
	return false
}

// If EncryptionAtHost is enabled, only include instance types that support it

// supportsEncryptionAtHost checks if the SKU supports encryption at host
func (p *DefaultProvider) supportsEncryptionAtHost(sku *skewer.SKU) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *DefaultProvider) isInstanceTypeSupportedByLocalDNS(sku *skewer.SKU, nodeClass *v1beta1.AKSNodeClass) bool {
	_ = "STUB: not implemented"
	// Read the resolved state from Status.LocalDNSState. The
	// nodeclass.localdns sub-reconciler is the sole writer.
	// If LocalDNS won't be enabled, all instance types are supported
	return false
}

// LocalDNS requires at least 4 vCPUs and 256 MB (244.140625 MiB) of memory

// 256 MB = 244.140625 MiB

func (p *DefaultProvider) isInstanceTypeSupportedByGPUDriverMode(sku *skewer.SKU, nodeClass *v1beta1.AKSNodeClass) bool {
	_ = "STUB: not implemented"
	// Only "Driver" mode filters out GPU SKUs without driver installation support.
	// "None" mode allows all GPU SKUs.
	return false
}

// Non-GPU SKUs are always allowed

// In "Driver" mode, only allow GPU SKUs with driver installation support

// isInstanceTypeSupportedByArtifactStreaming filters out ARM64 instance types when artifact streaming
// is explicitly enabled, since ARM64 does not support artifact streaming.
// When artifact streaming is not set (nil/default) or explicitly disabled, all architectures are allowed.
func (p *DefaultProvider) isInstanceTypeSupportedByArtifactStreaming(architecture string, nodeClass *v1beta1.AKSNodeClass) bool {
	_ = "STUB: not implemented"
	// Only filter when the user explicitly requested artifact streaming enabled
	return false
}

// Artifact streaming is explicitly enabled; exclude ARM64 since it doesn't support it

// UpdateInstanceTypes fetches all instance types from Azure (using skewer) and updates the cache.
// This is called periodically by the instance type controller.
func (p *DefaultProvider) UpdateInstanceTypes(ctx context.Context) error {
	_ = "STUB: not implemented"
	// DO NOT REMOVE THIS LOCK ----------------------------------------------------------------------------
	// We lock here so that multiple callers to UpdateInstanceTypes do not result in multiple
	// calls to Resource API when we could have just made one call. This lock is here because multiple callers result
	// in A LOT of extra memory generated from the response for simultaneous callers.
	return nil
}

// Only update instanceTypesSeqNum with the instance types have been changed
// This is to not create new keys with duplicate instance types option

// isSupported indicates SKU is supported by AKS, based on SKU properties
func (p *DefaultProvider) isSupported(sku *skewer.SKU, vmsize *skewer.VMSizeType) bool {
	_ = "STUB: not implemented"
	return false
}

// at least 2 cpus
func (p *DefaultProvider) hasMinimumCPU(sku *skewer.SKU) bool {
	_ = "STUB: not implemented"
	return false
}

// at least 3.5 GiB of memory
func (p *DefaultProvider) hasMinimumMemory(sku *skewer.SKU) bool {
	_ = "STUB: not implemented"
	return false
}

// instances AKS does not support
func (p *DefaultProvider) isUnsupportedByAKS(sku *skewer.SKU) bool {
	_ = "STUB: not implemented"
	return false
}

// GPU SKUs not in the supported GPU registry
func (p *DefaultProvider) isUnsupportedGPU(sku *skewer.SKU) bool {
	_ = "STUB: not implemented"
	return false
}

// SKU with constrained CPUs
func (p *DefaultProvider) hasConstrainedCPUs(vmsize *skewer.VMSizeType) bool {
	_ = "STUB: not implemented"
	return false
}

// confidential VMs (DC, EC) are not yet supported by this Karpenter provider
func (p *DefaultProvider) isConfidential(sku *skewer.SKU) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *DefaultProvider) Reset() { _ = "STUB: not implemented"; return }

func FindMaxEphemeralSizeGBAndPlacement(sku *skewer.SKU) (sizeGB int64, placement *armcompute.DiffDiskPlacement) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ephemeral OS disk is not supported by this SKU

// Check NVMe disk first (highest priority)

// NOTE: MaxResourceVolumeMB is actually in MiBs

func supportsNVMeEphemeralOSDisk(sku *skewer.SKU) bool { _ = "STUB: not implemented"; return false }

func UseEphemeralDisk(sku *skewer.SKU, nodeClass *v1beta1.AKSNodeClass) bool {
	_ = "STUB: not implemented"
	return false
}

// use ephemeral disk if it is large enough

func nvmeDiskSizeInMiB(s *skewer.SKU) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

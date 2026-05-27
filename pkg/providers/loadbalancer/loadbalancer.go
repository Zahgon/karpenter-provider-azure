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

package loadbalancer

import (
	"context"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/patrickmn/go-cache"
)

const (
	SLBName     = "kubernetes"
	SLBNameIPv6 = "kubernetes-ipv6"

	// InternalSLBName is the name of the internal SLB created by cloudprovider when Services with type LoadBalancer are created. This may not exist
	// when the pod launches/Karpenter is created. We query it as an optimization to save cloudprovider work, as otherwise cloudprovider must edit
	// ever VM we deploy to include this LB.
	InternalSLBName = "kubernetes-internal"

	// SLBOutboundBackendPoolName is the AKS SLB outbound backend pool name
	SLBOutboundBackendPoolName = "aksOutboundBackendPool"
	// SLBOutboundBackendPoolNameIPv6 is the AKS SLB outbound backend pool name for IPv6 traffic
	SLBOutboundBackendPoolNameIPv6 = "aksOutboundBackendPool-ipv6"
	// SLBInboundBackendPoolName is the AKS SLB inbound backend pool name
	SLBInboundBackendPoolName = "kubernetes"
	// SLBInboundBackendPoolNameIPv6 is the AKS SLB inbound backend pool name for IPv6 traffic
	SLBInboundBackendPoolNameIPv6 = "kubernetes-ipv6"

	loadBalancersCacheKey = "LoadBalancers"

	// LoadBalancersCacheTTL configures how freuqently we check for updates to the LBs.
	// Currently the choice of this value is entirely "how much work do we want to save cloudprovider".
	// The faster we do this, the faster we notice the creation of a kubernetes-internal LB and start
	// including it on new VMs, which saves CloudProvider needing to do that.
	LoadBalancersCacheTTL = 2 * time.Hour
)

type Provider struct {
	loadBalancersAPI LoadBalancersAPI
	resourceGroup    string
	cache            *cache.Cache
	mu               sync.Mutex
}

type BackendAddressPools struct {
	IPv4PoolIDs []string
	IPv6PoolIDs []string // TODO: This is always empty currently
}

// NewProvider creates a new LoadBalancer provider
func NewProvider(loadBalancersAPI LoadBalancersAPI, cache *cache.Cache, resourceGroup string) *Provider {
	_ = "STUB: not implemented"
	return nil
}

// LoadBalancerBackendPools returns a collection of IPv4 and IPv6 LoadBalancer backend pools.
// This collection is collected from Azure periodically but usually served from a cache to reduce
// Azure request load.
func (p *Provider) LoadBalancerBackendPools(ctx context.Context) (*BackendAddressPools, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RP only actually assigns the LB backend pools to VMs if OutboundType is LoadBalancer,
// but that's also the only OutboundType which creates the LoadBalancer, so as long as we're not allowing
// OutboundType changes, we can just infer that if the LBs exist we should assign them.

// TODO: IPv6 deferred for now. When they're used they must be put onto a non-primary NIC.

func (p *Provider) getLoadBalancers(ctx context.Context) ([]*armnetwork.LoadBalancer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If we wanted to hyper-optimize, we could set a much longer timeout once we find the -internal LB, as at that point we're "done" and
// aren't particularly interested in LB changes anymore.

func (p *Provider) loadFromAzure(ctx context.Context) ([]*armnetwork.LoadBalancer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Only return the LBs we actually care about

func isClusterLoadBalancer(lb *armnetwork.LoadBalancer, _ int) bool {
	_ = "STUB: not implemented"
	return false
}

// TODO: Not currently supporting IPv6

func extractBackendAddressPools(lb *armnetwork.LoadBalancer, _ int) []*armnetwork.BackendAddressPool {
	_ = "STUB: not implemented"
	return nil
}

func isBackendAddressPoolApplicable(backendPool *armnetwork.BackendAddressPool, _ int) bool {
	_ = "STUB: not implemented"
	return false
}

// shouldn't ever happen

// Ignore well-known named ipv6 pools for now

// Ignore IP-based pools, which are a thing in NodeIP mode. We don't need to assign these pools.
// See isIPBasedBackendPool in RP.

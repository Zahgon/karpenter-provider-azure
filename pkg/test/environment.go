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

package test

import (
	"context"

	gomegaformat "github.com/onsi/gomega/format"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"

	karpv1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/controllers/nodeoverlay"

	"github.com/patrickmn/go-cache"
	"k8s.io/client-go/dynamic"
	"sigs.k8s.io/controller-runtime/pkg/client"
	coretest "sigs.k8s.io/karpenter/pkg/test"

	azurecache "github.com/Azure/karpenter-provider-azure/pkg/cache"
	"github.com/Azure/karpenter-provider-azure/pkg/fake"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/allocationstrategy"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/imagefamily"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/instance"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/instance/machinecache"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/instancetype"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/kubernetesversion"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/launchtemplate"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/loadbalancer"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/networksecuritygroup"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/pricing"
)

func init() {
	karpv1.NormalizedLabels = lo.Assign(karpv1.NormalizedLabels, map[string]string{"topology.disk.csi.azure.com/zone": corev1.LabelTopologyZone})

	// Configuring this here because it's commonly imported and has an init already
	gomegaformat.CharactersAroundMismatchToInclude = 40
}

const (
	clusterName   = "test-cluster"
	resourceGroup = "test-resourceGroup"
	subscription  = "12345678-1234-1234-1234-123456789012"
)

type Environment struct {
	// API
	VirtualMachinesAPI          *fake.VirtualMachinesAPI
	AzureResourceGraphAPI       *fake.AzureResourceGraphAPI
	VirtualMachineExtensionsAPI *fake.VirtualMachineExtensionsAPI
	NetworkInterfacesAPI        *fake.NetworkInterfacesAPI
	CommunityImageVersionsAPI   *fake.CommunityGalleryImageVersionsAPI
	NodeImageVersionsAPI        *fake.NodeImageVersionsAPI
	SKUsAPI                     *fake.ResourceSKUsAPI
	PricingAPI                  *fake.PricingAPI
	LoadBalancersAPI            *fake.LoadBalancersAPI
	NetworkSecurityGroupAPI     *fake.NetworkSecurityGroupAPI
	SubnetsAPI                  *fake.SubnetsAPI
	DiskEncryptionSetsAPI       *fake.DiskEncryptionSetsAPI
	AuxiliaryTokenServer        *fake.AuxiliaryTokenServer
	SubscriptionAPI             *fake.SubscriptionsAPI
	NodeBootstrappingAPI        *fake.NodeBootstrappingAPI
	AKSMachinesAPI              *fake.AKSMachinesAPI
	AKSAgentPoolsAPI            *fake.AKSAgentPoolsAPI
	DynamicInterface            dynamic.Interface

	// Fake data stores for the APIs
	AKSDataStorage *fake.AKSDataStorage

	// Cache
	AKSMachineCache           *machinecache.MachineCache
	KubernetesVersionCache    *cache.Cache
	NodeImagesCache           *cache.Cache
	InstanceTypeCache         *cache.Cache
	LoadBalancerCache         *cache.Cache
	UnavailableOfferingsCache *azurecache.UnavailableOfferings

	// Providers
	InstanceTypesProvider        *instancetype.DefaultProvider
	VMInstanceProvider           instance.VMProvider
	AKSMachineProvider           instance.AKSMachineProvider
	PricingProvider              *pricing.Provider
	KubernetesVersionProvider    kubernetesversion.KubernetesVersionProvider
	ImageProvider                imagefamily.NodeImageProvider
	ImageResolver                imagefamily.Resolver
	LaunchTemplateProvider       *launchtemplate.Provider
	LoadBalancerProvider         *loadbalancer.Provider
	NetworkSecurityGroupProvider *networksecuritygroup.Provider
	AllocationStrategyProvider   allocationstrategy.Provider

	InstanceTypeStore *nodeoverlay.InstanceTypeStore

	// Settings
	nonZonal       bool
	SubscriptionID string
	coreEnv        *coretest.Environment
	region         string
}

func NewEnvironment(ctx context.Context, env *coretest.Environment) *Environment {
	_ = "STUB: not implemented"
	return nil
}

func NewEnvironmentNonZonal(ctx context.Context, env *coretest.Environment) *Environment {
	_ = "STUB: not implemented"
	return nil
}

func NewRegionalEnvironment(ctx context.Context, env *coretest.Environment, region string, nonZonal bool) *Environment {
	_ = "STUB: not implemented"
	return nil
}

// API

// Cache

// Providers

// Set up batching if provision mode is header batch

// For this configuration, we assume the AKS machines pool already exists

// Populate the instance type cache before returning the environment, as many tests assume it's populated and it simplifies test setup.
// We can update it in individual tests as needed.

// Seed the managed NSG

func (env *Environment) Reset(ctx context.Context) { _ = "STUB: not implemented"; return }

// Re-seed the managed NSG so launchtemplate provider can resolve it

func (env *Environment) Zones() []string { _ = "STUB: not implemented"; return nil }

// Client returns the controller-runtime client from the underlying core test environment.
func (env *Environment) Client() client.Client {
	_ = "STUB: not implemented"
	return *new(client.Client)
}

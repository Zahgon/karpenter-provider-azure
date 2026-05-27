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

package azclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/karpenter-provider-azure/pkg/auth"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/azclient/aksmachinesheaderbatch"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/azclient/azapi"
	imagefamilytypes "github.com/Azure/karpenter-provider-azure/pkg/providers/imagefamily/types"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/loadbalancer"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/networksecuritygroup"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/zone"
	"github.com/Azure/skewer"
)

type AZClient struct {
	azureResourceGraphClient       azapi.AzureResourceGraphAPI
	virtualMachinesClient          azapi.VirtualMachinesAPI
	aksMachinesClient              azapi.AKSMachinesAPI
	aksMachinesBatchClient         aksmachinesheaderbatch.AKSMachinesHeaderBatchAPI
	agentPoolsClient               azapi.AKSAgentPoolsAPI
	virtualMachinesExtensionClient azapi.VirtualMachineExtensionsAPI
	networkInterfacesClient        azapi.NetworkInterfacesAPI
	subnetsClient                  azapi.SubnetsAPI
	diskEncryptionSetsClient       azapi.DiskEncryptionSetsAPI

	NodeImageVersionsClient imagefamilytypes.NodeImageVersionsAPI
	ImageVersionsClient     imagefamilytypes.CommunityGalleryImageVersionsAPI
	NodeBootstrappingClient imagefamilytypes.NodeBootstrappingAPI
	// SKU CLIENT is still using track 1 because skewer does not support the track 2 path. We need to refactor this once skewer supports track 2
	SKUClient                   skewer.ResourceClient
	LoadBalancersClient         loadbalancer.LoadBalancersAPI
	NetworkSecurityGroupsClient networksecuritygroup.API
	SubscriptionsClient         zone.SubscriptionsAPI
}

func (c *AZClient) SubnetsClient() azapi.SubnetsAPI {
	_ = "STUB: not implemented"
	return *new(azapi.SubnetsAPI)
}

func (c *AZClient) DiskEncryptionSetsClient() azapi.DiskEncryptionSetsAPI {
	_ = "STUB: not implemented"
	return *new(azapi.DiskEncryptionSetsAPI)
}

func (c *AZClient) AKSMachinesClient() azapi.AKSMachinesAPI {
	_ = "STUB: not implemented"
	return *new(azapi.AKSMachinesAPI)
}

func (c *AZClient) AKSMachinesBatchClient() aksmachinesheaderbatch.AKSMachinesHeaderBatchAPI {
	_ = "STUB: not implemented"
	return *new(aksmachinesheaderbatch.AKSMachinesHeaderBatchAPI)
}

func (c *AZClient) AgentPoolsClient() azapi.AKSAgentPoolsAPI {
	_ = "STUB: not implemented"
	return *new(azapi.AKSAgentPoolsAPI)
}

func (c *AZClient) VirtualMachinesClient() azapi.VirtualMachinesAPI {
	_ = "STUB: not implemented"
	return *new(azapi.VirtualMachinesAPI)
}

func (c *AZClient) VirtualMachineExtensionsClient() azapi.VirtualMachineExtensionsAPI {
	_ = "STUB: not implemented"
	return *new(azapi.VirtualMachineExtensionsAPI)
}

func (c *AZClient) NetworkInterfacesClient() azapi.NetworkInterfacesAPI {
	_ = "STUB: not implemented"
	return *new(azapi.NetworkInterfacesAPI)
}

func (c *AZClient) AzureResourceGraphClient() azapi.AzureResourceGraphAPI {
	_ = "STUB: not implemented"
	return *new(azapi.AzureResourceGraphAPI)
}

func NewAZClientFromAPI(
	virtualMachinesClient azapi.VirtualMachinesAPI,
	azureResourceGraphClient azapi.AzureResourceGraphAPI,
	aksMachinesClient azapi.AKSMachinesAPI,
	aksMachinesBatchClient aksmachinesheaderbatch.AKSMachinesHeaderBatchAPI,
	agentPoolsClient azapi.AKSAgentPoolsAPI,
	virtualMachinesExtensionClient azapi.VirtualMachineExtensionsAPI,
	interfacesClient azapi.NetworkInterfacesAPI,
	subnetsClient azapi.SubnetsAPI,
	diskEncryptionSetsClient azapi.DiskEncryptionSetsAPI,
	loadBalancersClient loadbalancer.LoadBalancersAPI,
	networkSecurityGroupsClient networksecuritygroup.API,
	imageVersionsClient imagefamilytypes.CommunityGalleryImageVersionsAPI,
	nodeImageVersionsClient imagefamilytypes.NodeImageVersionsAPI,
	nodeBootstrappingClient imagefamilytypes.NodeBootstrappingAPI,
	skuClient skewer.ResourceClient,
	subscriptionsClient zone.SubscriptionsAPI,
) *AZClient {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gocyclo
func NewAZClient(ctx context.Context, cfg *auth.Config, env *auth.Environment, cred azcore.TokenCredential) (*AZClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// copy the options to avoid modifying the original

// Not doing this if PROVISION_MODE is an AKS machine API mode as Create will never use VM client, but want to allow other VM client operations

// TODO: this one is not enabled for rate limiting / throttling ...
// TODO Move this over to track 2 when skewer is migrated

// These clients are used for Azure instance management.

// Only create the bootstrapping client if we need to use it.

// Only create AKS machine clients if we need to use them.
// Otherwise, use the no-op dry clients, which will act like there are no AKS machines present.

// copy the options to avoid modifying the original

// Try create true clients. This is just for diagnostic purposes and serves no real functionality.
// This portion of code can be removed once we are confident that this works reliably.

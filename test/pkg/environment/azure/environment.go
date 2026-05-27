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

package azure

import (
	"testing"
	"time"

	"github.com/samber/lo"
	v1 "k8s.io/api/core/v1"
	karpv1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	coretest "sigs.k8s.io/karpenter/pkg/test"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
	containerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
	"github.com/Azure/karpenter-provider-azure/pkg/auth"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/zone"
	"github.com/Azure/karpenter-provider-azure/pkg/test/azure"
	"github.com/Azure/karpenter-provider-azure/test/pkg/environment/common"
)

func init() {
	karpv1.NormalizedLabels = lo.Assign(karpv1.NormalizedLabels, map[string]string{"topology.disk.csi.azure.com/zone": v1.LabelTopologyZone})
	coretest.DefaultImage = "mcr.microsoft.com/oss/kubernetes/pause:3.6"
}

const (
	CiliumAgentNotReadyTaint    = "node.cilium.io/agent-not-ready"
	EphemeralInitContainerImage = "alpine"

	ciliumStartupTaintTolerationSeconds int64 = 120
)

type Environment struct {
	*common.Environment

	NodeResourceGroup    string
	Region               string
	SubscriptionID       string
	VNETResourceGroup    string
	ACRName              string
	ClusterName          string
	MachineAgentPoolName string
	ClusterResourceGroup string
	CloudConfig          cloud.Configuration
	ProvisionMode        string

	tracker *azure.Tracker

	// These should be unexported and access should be through the Environment methods
	// Any create calls should make sure they also register the created resources with the Environment's tracker
	// to ensure they are cleaned up after the test.
	vmClient             *armcompute.VirtualMachinesClient
	vnetClient           *armnetwork.VirtualNetworksClient
	subnetClient         *armnetwork.SubnetsClient
	interfacesClient     *armnetwork.InterfacesClient
	managedClusterClient *containerservice.ManagedClustersClient
	agentPoolClient      *containerservice.AgentPoolsClient
	machinesClient       *containerservice.MachinesClient
	zoneProvider         *zone.Provider

	// Public Clients
	KeyVaultClient          *armkeyvault.VaultsClient
	DiskEncryptionSetClient *armcompute.DiskEncryptionSetsClient

	defaultCredential azcore.TokenCredential

	RBACManager *RBACManager
}

func readEnvRequired(name string) string { _ = "STUB: not implemented"; return "" }

func readEnvOptional(name string) string { _ = "STUB: not implemented"; return "" }

func getCloudEnvironment() *auth.Environment { _ = "STUB: not implemented"; return nil }

// This is a hack so we can re-use the same validate, even though in this test context we don't need a real subscription ID

func NewEnvironment(t *testing.T) *Environment { _ = "STUB: not implemented"; return nil }

// If ProvisionMode wasn't set, default to scriptless, though note that this is
// actually defaulted dynamically based on the value of a toggle in AKS which means
// assuming we're always in ProvisionMode Scriptless here is incorrect at times, though OK
// for our current usage.

// Default to reserved managed machine agentpool name for NAP

// Confirm we have a machine pool

type realClock struct{}

func (realClock) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (env *Environment) GetDefaultCredential() azcore.TokenCredential {
	_ = "STUB: not implemented"
	return *new(azcore.TokenCredential)
}

// SupportsZones returns true if the region supports availability zones
func (env *Environment) SupportsZones() bool { _ = "STUB: not implemented"; return false }

// GetAvailableZones returns the list of available zones for the current region.
// Returns nil if the region doesn't support zones.
func (env *Environment) GetAvailableZones() []string { _ = "STUB: not implemented"; return nil }

// Retry options for BYOK-related clients that may encounter RBAC propagation delays
// RBAC assignments can take time to propagate, resulting in 403 Forbidden errors
// With 15 retries at 5 second intervals = 75 seconds total retry time
func (env *Environment) ClientOptionsForRBACPropagation() *arm.ClientOptions {
	_ = "STUB: not implemented"
	return nil
}

// RBAC assignments haven't propagated yet

func (env *Environment) IsAKSMachineAPIMode() bool { _ = "STUB: not implemented"; return false }

func (env *Environment) IsMachineModeOrNPS() bool {
	_ = "STUB: not implemented"
	// Assumption is if we're not in the cluster, we're in NPS mode. Ideally we would just check this via ProvisionMode, but
	// we can't do that right now as depending on context we may not set provision mode for the tests
	return false
}

func (env *Environment) UsesSharedImageGallery() bool { _ = "STUB: not implemented"; return false }

func (env *Environment) DefaultAKSNodeClass() *v1beta1.AKSNodeClass {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) AZLinuxNodeClass() *v1beta1.AKSNodeClass {
	_ = "STUB: not implemented"
	return nil
}

// Pod wraps coretest.Pod for Azure E2E tests; use it instead of coretest.Pod when the test should apply Azure environment defaults.
// Currently this is any time one has to work around taint race described in https://github.com/Azure/karpenter-provider-azure/issues/1625
// and cannot use Deployment instead.
func (env *Environment) Pod(options coretest.PodOptions) *v1.Pod {
	_ = "STUB: not implemented"
	// Keep pod-based tests resilient to the Cilium startup-taint race while bounding how long the pod can tolerate it.
	return nil
}

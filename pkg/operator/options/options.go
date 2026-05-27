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

package options

import (
	"context"

	coreoptions "sigs.k8s.io/karpenter/pkg/operator/options"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

func init() {
	coreoptions.Injectables = append(coreoptions.Injectables, &Options{})
}

type nodeIdentitiesValue []string

func newNodeIdentitiesValue(val string, p *[]string) *nodeIdentitiesValue {
	_ = "STUB: not implemented"
	return nil
}

func (s *nodeIdentitiesValue) Set(val string) error { _ = "STUB: not implemented"; return nil }

func (s *nodeIdentitiesValue) Get() any { _ = "STUB: not implemented"; return *new(any) }

func (s *nodeIdentitiesValue) String() string { _ = "STUB: not implemented"; return "" }

type optionsKey struct{}

type Options struct {
	ClusterName                    string  `json:"clusterName,omitempty"`
	ClusterEndpoint                string  `json:"clusterEndpoint,omitempty"` // => APIServerName in bootstrap, except needs to be w/o https/port
	VMMemoryOverheadPercent        float64 `json:"vmMemoryOverheadPercent,omitempty"`
	KubeletClientTLSBootstrapToken string  `json:"-"` // => TLSBootstrapToken in bootstrap (may need to be per node/nodepool)
	LinuxAdminUsername             string  `json:"-"`
	SSHPublicKey                   string  `json:"-"` // ssh.publicKeys.keyData => VM SSH public key // TODO: move to v1beta1.AKSNodeClass?

	NetworkPlugin     string `json:"networkPlugin,omitempty"`     // => NetworkPlugin in bootstrap
	NetworkPolicy     string `json:"networkPolicy,omitempty"`     // => NetworkPolicy in bootstrap
	NetworkPluginMode string `json:"networkPluginMode,omitempty"` // => Network Plugin Mode is used to control the mode the network plugin should operate in. For example, "overlay" used with --network-plugin=azure will use an overlay network (non-VNET IPs) for pods in the cluster. Learn more about overlay networking here: https://learn.microsoft.com/en-us/azure/aks/azure-cni-overlay?tabs=kubectl#overview-of-overlay-networking
	NetworkDataplane  string `json:"networkDataplane,omitempty"`
	DNSServiceIP      string `json:"dnsServiceIP,omitempty"`

	NodeIdentities          []string `json:"nodeIdentities,omitempty"`          // => Applied onto each VM
	KubeletIdentityClientID string   `json:"kubeletIdentityClientID,omitempty"` // => Flows to bootstrap and used in drift
	VnetGUID                string   `json:"vnetGuid,omitempty"`                // resource guid used by azure cni for identifying the right vnet
	SubnetID                string   `json:"subnetId,omitempty"`                // => VnetSubnetID to use (for nodes in Azure CNI Overlay and Azure CNI + pod subnet; for for nodes and pods in Azure CNI), unless overridden via AKSNodeClass
	setFlags                map[string]bool

	ProvisionMode              string            `json:"provisionMode,omitempty"`
	NodeBootstrappingServerURL string            `json:"-"`
	UseSIG                     bool              `json:"useSIG,omitempty"` // => UseSIG is true if Karpenter is managed by AKS, false if it is a self-hosted karpenter installation
	SIGAccessTokenServerURL    string            `json:"-"`                // => SIGAccessTokenServerURL used to access SIG, not set if it is a self-hosted karpenter installation
	SIGSubscriptionID          string            `json:"sigSubscriptionId,omitempty"`
	NodeResourceGroup          string            `json:"nodeResourceGroup,omitempty"`
	AdditionalTags             map[string]string `json:"additionalTags,omitempty"`
	EnableAzureSDKLogging      bool              `json:"enableAzureSDKLogging,omitempty"` // Controls whether Azure SDK middleware logging is enabled
	DiskEncryptionSetID        string            `json:"diskEncryptionSetId,omitempty"`

	// If set to true, existing AKS machines created with an AKS Machine API provision mode will be managed even with other provision modes. This option does not have any effect if PROVISION_MODE is already an AKS Machine API mode, as it will behave as if this option is set to true.
	ManageExistingAKSMachines bool `json:"manageExistingAKSMachines,omitempty"`

	AKSMachinesPoolName string `json:"aksMachinesPoolName,omitempty"` // The name of the agent pool for the AKS machine API, assuming that all machines belong to the same agent pool. Only used on AKS machine API provision modes.
	BatchIdleTimeoutMS  int    `json:"batchIdleTimeoutMS,omitempty"`  // Idle timeout in milliseconds for batch accumulation (default 1000ms). Only used on provision mode aksmachineapiheaderbatch.
	BatchMaxTimeoutMS   int    `json:"batchMaxTimeoutMS,omitempty"`   // Maximum timeout in milliseconds for batch accumulation (default 5000ms). Only used on provision mode aksmachineapiheaderbatch.
	MaxBatchSize        int    `json:"maxBatchSize,omitempty"`        // Maximum number of machines per batch (default 50, AKS API limit). Only used on provision mode aksmachineapiheaderbatch.

	// computed options; do not set.
	ParsedDiskEncryptionSetID *arm.ResourceID `json:"-"`
}

func (o *Options) AddFlags(fs *coreoptions.FlagSet) { _ = "STUB: not implemented"; return }

// See https://github.com/Azure/karpenter-provider-azure/issues/1042 for issue discussing improvements around this

// IsAKSMachineAPIMode returns true if the current provision mode creates instances via the AKS Machine API.
func (o *Options) IsAKSMachineAPIMode() bool { _ = "STUB: not implemented"; return false }

func (o *Options) GetAPIServerName() string { _ = "STUB: not implemented"; return "" }

// assume to already validated

func (o *Options) Parse(fs *coreoptions.FlagSet, args ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if each option has been set. This is a little brute force and better options might exist,
// but this only needs to be here for one version

func (o *Options) String() string { _ = "STUB: not implemented"; return "" }

func (o *Options) ToContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func ToContext(ctx context.Context, opts *Options) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func FromContext(ctx context.Context) *Options { _ = "STUB: not implemented"; return nil }

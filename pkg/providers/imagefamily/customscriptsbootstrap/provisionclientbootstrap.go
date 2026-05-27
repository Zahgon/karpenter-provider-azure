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

package customscriptsbootstrap

import (
	"context"

	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/imagefamily/bootstrap"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/imagefamily/types"
	"github.com/Azure/karpenter-provider-azure/pkg/provisionclients/models"

	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
)

const (
	ImageFamilyOSSKUUbuntu2004  = "Ubuntu2004"
	ImageFamilyOSSKUUbuntu2204  = "Ubuntu2204"
	ImageFamilyOSSKUUbuntu2404  = "Ubuntu2404"
	ImageFamilyOSSKUAzureLinux2 = "AzureLinux2"
	ImageFamilyOSSKUAzureLinux3 = "AzureLinux3"
)

type ProvisionClientBootstrap struct {
	ClusterName                    string
	KubeletConfig                  *bootstrap.KubeletConfiguration
	Taints                         []v1.Taint        `hash:"set"`
	StartupTaints                  []v1.Taint        `hash:"set"`
	Labels                         map[string]string `hash:"set"`
	SubnetID                       string
	Arch                           string
	SubscriptionID                 string
	ClusterResourceGroup           string
	ResourceGroup                  string
	KubeletClientTLSBootstrapToken string
	KubernetesVersion              string
	ImageDistro                    string
	IsWindows                      bool
	InstanceType                   *cloudprovider.InstanceType
	StorageProfile                 string
	OSSKU                          string
	NodeBootstrappingProvider      types.NodeBootstrappingAPI
	GPUDriverInstallationEnabled   bool
	FIPSMode                       *v1beta1.FIPSMode
	LocalDNSProfile                *v1beta1.LocalDNS
	ArtifactStreaming              *v1beta1.ArtifactStreaming
	LinuxOSConfig                  *v1beta1.LinuxOSConfiguration
}

var _ Bootstrapper = (*ProvisionClientBootstrap)(nil) // assert ProvisionClientBootstrap implements customscriptsbootstrapper

func (p ProvisionClientBootstrap) GetCustomDataAndCSE(ctx context.Context) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// As of now we just fail the provisioning given the unlikely scenario of retriable error, but could be revisited along with retriable status on the server side.

// ATTENTION!!!: changes here may NOT be effective on AKS machine nodes (ProvisionModeAKSMachineAPI); See aksmachineinstance.go/aksmachineinstancehelpers.go.
// Refactoring for code unification is not being invested immediately.
//
//nolint:gocyclo
func (p *ProvisionClientBootstrap) ConstructProvisionValues(ctx context.Context) (*models.ProvisionValues, error) {
	_ = "STUB: not implemented"

	// TODO(Windows)
	return nil, nil
}

// unspecified FIPSMode is effectively no FIPS for now

// EnableVTPM:       lo.ToPtr(false), // Unsupported as of now (Trusted launch)
// EnableSecureBoot: lo.ToPtr(false), // Unsupported as of now (Trusted launch)

// Unsupported as of now; TODO(Windows)
// MessageOfTheDay:         lo.ToPtr(""),                                    // Unsupported as of now
// AgentPoolWindowsProfile: &models.AgentPoolWindowsProfile{},               // Unsupported as of now; TODO(Windows)
// KubeletDiskType:         lo.ToPtr(models.KubeletDiskTypeUnspecified),    // Unsupported as of now
// CustomLinuxOSConfig:     &models.CustomLinuxOSConfig{},                   // Unsupported as of now (sysctl)

// GpuInstanceProfile:      lo.ToPtr(models.GPUInstanceProfileUnspecified), // Unsupported as of now (MIG)
// WorkloadRuntime:         lo.ToPtr(models.WorkloadRuntimeUnspecified),    // Unsupported as of now (Kata)

// Map OS SKU to AKS provision client's expectation
// Note that the direction forward is to be more specific with OS versions. Be careful when supporting new ones.

// https://go.dev/wiki/Switch#multiple-cases

// NodeClaim defaults don't work somehow and keep giving invalid values. Can be improved later.

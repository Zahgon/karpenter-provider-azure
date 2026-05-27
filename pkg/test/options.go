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
	azoptions "github.com/Azure/karpenter-provider-azure/pkg/operator/options"
)

type OptionsFields struct {
	ClusterName                    *string
	ClusterEndpoint                *string
	KubeletClientTLSBootstrapToken *string
	LinuxAdminUsername             *string
	SSHPublicKey                   *string
	NetworkPlugin                  *string
	NetworkPluginMode              *string
	NetworkPolicy                  *string
	NetworkDataplane               *string
	VMMemoryOverheadPercent        *float64
	NodeIdentities                 []string
	SubnetID                       *string
	NodeResourceGroup              *string
	ProvisionMode                  *string
	NodeBootstrappingServerURL     *string
	VnetGUID                       *string
	KubeletIdentityClientID        *string
	AdditionalTags                 map[string]string
	EnableAzureSDKLogging          *bool
	DiskEncryptionSetID            *string
	ClusterDNSServiceIP            *string
	ManageExistingAKSMachines      *bool
	AKSMachinesPoolName            *string
	BatchIdleTimeoutMS             *int
	BatchMaxTimeoutMS              *int
	MaxBatchSize                   *int

	// SIG Flags not required by the self hosted offering
	UseSIG                  *bool
	SIGAccessTokenServerURL *string
	SIGSubscriptionID       *string
}

func Options(overrides ...OptionsFields) *azoptions.Options { _ = "STUB: not implemented"; return nil }

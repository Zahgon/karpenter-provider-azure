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

package imagefamily

import (
	v1 "k8s.io/api/core/v1"

	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/imagefamily/bootstrap"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/imagefamily/customscriptsbootstrap"
	types "github.com/Azure/karpenter-provider-azure/pkg/providers/imagefamily/types"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/launchtemplate/parameters"

	"sigs.k8s.io/karpenter/pkg/cloudprovider"
)

const (
	AzureLinux3Gen2ImageDefinition          = "V3gen2"
	AzureLinux3Gen1ImageDefinition          = "V3"
	AzureLinux3Gen2ArmImageDefinition       = "V3gen2arm64"
	AzureLinux3Gen2FIPSImageDefinition      = "V3gen2fips"
	AzureLinux3Gen1FIPSImageDefinition      = "V3fips"
	AzureLinux3Gen2Arm64FIPSImageDefinition = "V3gen2arm64fips"
)

type AzureLinux3 struct {
	Options *parameters.StaticParameters
}

func (u AzureLinux3) Name() string { _ = "STUB: not implemented"; return "" }

func (u AzureLinux3) DefaultImages(useSIG bool, fipsMode *v1beta1.FIPSMode) []types.DefaultImageOutput {
	_ = "STUB: not implemented"
	return nil
}

// Note: FIPS images aren't supported in public galleries, only shared image galleries
// image provider will select these images in order, first match wins

// image provider will select these images in order, first match wins

// UserData returns the default userdata script for the image Family
func (u AzureLinux3) ScriptlessCustomData(
	kubeletConfig *bootstrap.KubeletConfiguration,
	taints []v1.Taint,
	labels map[string]string,
	caBundle *string,
	_ *cloudprovider.InstanceType,
) bootstrap.Bootstrapper {
	_ = "STUB: not implemented"
	return *new(bootstrap.Bootstrapper)
}

// UserData returns the default userdata script for the image Family
func (u AzureLinux3) CustomScriptsNodeBootstrapping(
	kubeletConfig *bootstrap.KubeletConfiguration,
	taints []v1.Taint,
	startupTaints []v1.Taint,
	labels map[string]string,
	instanceType *cloudprovider.InstanceType,
	imageDistro string,
	storageProfile string,
	nodeBootstrappingClient types.NodeBootstrappingAPI,
	fipsMode *v1beta1.FIPSMode,
	localDNS *v1beta1.LocalDNS,
	artifactStreaming *v1beta1.ArtifactStreaming,
	linuxOSConfig *v1beta1.LinuxOSConfiguration,
) customscriptsbootstrap.Bootstrapper {
	_ = "STUB: not implemented"
	return *new(customscriptsbootstrap.Bootstrapper)
}

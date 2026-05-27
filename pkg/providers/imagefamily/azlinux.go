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
	AzureLinuxGen2ImageDefinition      = "V2gen2"
	AzureLinuxGen1ImageDefinition      = "V2"
	AzureLinuxGen2ArmImageDefinition   = "V2gen2arm64"
	AzureLinux2Gen2FIPSImageDefinition = "V2gen2fips"
	AzureLinux2Gen1FIPSImageDefinition = "V2fips"
)

type AzureLinux struct {
	Options *parameters.StaticParameters
}

func (u AzureLinux) Name() string { _ = "STUB: not implemented"; return "" }

func (u AzureLinux) DefaultImages(useSIG bool, fipsMode *v1beta1.FIPSMode) []types.DefaultImageOutput {
	_ = "STUB: not implemented"
	return nil
}

// Note: FIPS images aren't supported in public galleries, only shared image galleries
// image provider will select these images in order, first match wins

// image provider will select these images in order, first match wins. This is why we chose to put Gen2 first in the defaultImages, as we prefer gen2 over gen1

// UserData returns the default userdata script for the image Family
func (u AzureLinux) ScriptlessCustomData(
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
func (u AzureLinux) CustomScriptsNodeBootstrapping(
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

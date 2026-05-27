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
	Ubuntu2404Gen2ImageDefinition    = "2404gen2containerd"
	Ubuntu2404Gen1ImageDefinition    = "2404containerd"
	Ubuntu2404Gen2ArmImageDefinition = "2404gen2arm64containerd"
)

type Ubuntu2404 struct {
	Options *parameters.StaticParameters
}

func (u Ubuntu2404) Name() string { _ = "STUB: not implemented"; return "" }

func (u Ubuntu2404) DefaultImages(useSIG bool, fipsMode *v1beta1.FIPSMode) []types.DefaultImageOutput {
	_ = "STUB: not implemented"
	return nil
}

// Note: FIPS images aren't supported in public galleries, only shared image galleries

//TODO: Fill out when Ubuntu 24.04 with FIPS becomes available

// image provider will select these images in order, first match wins. This is why we chose to put Ubuntu2404Gen2containerd first in the defaultImages

// UserData returns the default userdata script for the image Family
func (u Ubuntu2404) ScriptlessCustomData(
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
func (u Ubuntu2404) CustomScriptsNodeBootstrapping(
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

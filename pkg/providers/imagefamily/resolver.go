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
	"context"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/imagefamily/bootstrap"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/imagefamily/customscriptsbootstrap"
	types "github.com/Azure/karpenter-provider-azure/pkg/providers/imagefamily/types"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/instancetype"
	template "github.com/Azure/karpenter-provider-azure/pkg/providers/launchtemplate/parameters"
	karpv1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
)

type Resolver interface {
	Resolve(
		ctx context.Context,
		nodeClass *v1beta1.AKSNodeClass,
		nodeClaim *karpv1.NodeClaim,
		instanceType *cloudprovider.InstanceType,
		staticParameters *template.StaticParameters) (*template.Parameters, error)
	ResolveNodeImageFromNodeClass(nodeClass *v1beta1.AKSNodeClass, instanceType *cloudprovider.InstanceType) (string, error)
}

// assert that defaultResolver implements Resolver interface
var _ Resolver = &defaultResolver{}

// defaultResolver is able to fill-in dynamic launch template parameters
// ATTENTION!!!: changes here may NOT be effective on AKS machine nodes (ProvisionModeAKSMachineAPI); See aksmachineinstance.go/aksmachineinstancehelpers.go.
// Refactoring for code unification is not being invested immediately.
type defaultResolver struct {
	nodeBootstrappingProvider types.NodeBootstrappingAPI
	imageProvider             *provider
	instanceTypeProvider      instancetype.Provider
}

// ImageFamily can be implemented to override the default logic for generating dynamic launch template parameters
// ATTENTION!!!: changes here may NOT be effective on AKS machine nodes (ProvisionModeAKSMachineAPI); See aksmachineinstance.go/aksmachineinstancehelpers.go.
// Refactoring for code unification is not being invested immediately.
type ImageFamily interface {
	ScriptlessCustomData(
		kubeletConfig *bootstrap.KubeletConfiguration,
		taints []corev1.Taint,
		labels map[string]string,
		caBundle *string,
		instanceType *cloudprovider.InstanceType,
	) bootstrap.Bootstrapper
	CustomScriptsNodeBootstrapping(
		kubeletConfig *bootstrap.KubeletConfiguration,
		taints []corev1.Taint,
		startupTaints []corev1.Taint,
		labels map[string]string,
		instanceType *cloudprovider.InstanceType,
		imageDistro string,
		storageProfile string,
		nodeBootstrappingClient types.NodeBootstrappingAPI,
		fipsMode *v1beta1.FIPSMode,
		localDNS *v1beta1.LocalDNS,
		artifactStreaming *v1beta1.ArtifactStreaming,
		linuxOSConfig *v1beta1.LinuxOSConfiguration,
	) customscriptsbootstrap.Bootstrapper
	Name() string
	// DefaultImages returns a list of default CommunityImage definitions for this ImageFamily.
	// Our Image Selection logic relies on the ordering of the default images to be ordered from most preferred to least, then we will select the latest image version available for that CommunityImage definition.
	// Our Release pipeline ensures all images are released together within 24 hours of each other for community image gallery, so selecting based on image feature priorities, then by date, and not vice-versa is acceptable.
	// If fipsMode is FIPSModeFIPS, only FIPS-enabled images will be returned
	DefaultImages(useSIG bool, fipsMode *v1beta1.FIPSMode) []types.DefaultImageOutput
}

// NewDefaultResolver constructs a new launch template Resolver
func NewDefaultResolver(_ client.Client, imageProvider *provider, instanceTypeProvider instancetype.Provider, nodeBootstrappingClient types.NodeBootstrappingAPI) *defaultResolver {
	_ = "STUB: not implemented"
	return nil
}

// Resolve fills in dynamic launch template parameters.
// The name "imageFamilyResolver.Resolve()" is potentially misleading here.
// Suggestion: refactor would help, but this won't be used by PROVISION_MODE=aksmachineapi anyway. May not be worth it.
// ATTENTION!!!: changes here may NOT be effective on AKS machine nodes (ProvisionModeAKSMachineAPI); See aksmachineinstance.go/aksmachineinstancehelpers.go.
// Refactoring for code unification is not being invested immediately.
func (r *defaultResolver) Resolve(
	ctx context.Context,
	nodeClass *v1beta1.AKSNodeClass,
	nodeClaim *karpv1.NodeClaim,
	instanceType *cloudprovider.InstanceType,
	staticParameters *template.StaticParameters,
) (*template.Parameters, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: as ProvisionModeBootstrappingClient path develops, we will eventually be able to drop the retrieval of imageDistro here.

// ATTENTION!!!: changes here will NOT be effective on AKS machine nodes (ProvisionModeAKSMachineAPI); See aksmachineinstance.go/aksmachineinstancehelpers.go.
// Refactoring for code unification is not being invested immediately.

// TODO: We could potentially use the instance type to do defaulting like
// traditional AKS, so putting this here along with the other settings

// TODO(Windows)

func (r *defaultResolver) getStorageProfile(ctx context.Context, instanceType *cloudprovider.InstanceType, nodeClass *v1beta1.AKSNodeClass) (diskType string, placement *armcompute.DiffDiskPlacement, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func mapToImageDistro(imageID string, fipsMode *v1beta1.FIPSMode, imageFamily ImageFamily, useSIG bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ATTENTION!!!: changes here may NOT be effective on AKS machine nodes (ProvisionModeAKSMachineAPI); See aksmachineinstance.go/aksmachineinstancehelpers.go.
// Refactoring for code unification is not being invested immediately.
func prepareKubeletConfiguration(ctx context.Context, instanceType *cloudprovider.InstanceType, nodeClass *v1beta1.AKSNodeClass) *bootstrap.KubeletConfiguration {
	_ = "STUB: not implemented"
	return nil
}

// TODO: revisit computeResources implementation

func getSupportedImages(familyName *string, fipsMode *v1beta1.FIPSMode, kubernetesVersion string, useSIG bool) []types.DefaultImageOutput {
	_ = "STUB: not implemented"
	// TODO: Options aren't used within DefaultImages, so safe to be using nil here. Refactor so we don't actually need to pass in Options for getting DefaultImage.
	return nil
}

func GetImageFamily(familyName *string, fipsMode *v1beta1.FIPSMode, kubernetesVersion string, parameters *template.StaticParameters) ImageFamily {
	_ = "STUB: not implemented"
	return *new(ImageFamily)
}

func defaultUbuntu(fipsMode *v1beta1.FIPSMode, kubernetesVersion string, parameters *template.StaticParameters) ImageFamily {
	_ = "STUB: not implemented"
	return *new(ImageFamily)
}

// ResolveNodeImageFromNodeClass resolves Distro and image ID for the given node class and instance type. Images may vary due to architecture, accelerator, etc
func (r *defaultResolver) ResolveNodeImageFromNodeClass(nodeClass *v1beta1.AKSNodeClass, instanceType *cloudprovider.InstanceType) (string, error) {
	_ = "STUB: not implemented"
	// ASSUMPTION: nodeImages in a NodeClass are always sorted by priority order.
	return "", nil
}

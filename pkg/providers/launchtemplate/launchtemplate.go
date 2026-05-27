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

package launchtemplate

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/imagefamily"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/launchtemplate/parameters"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/networksecuritygroup"

	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
	karpv1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
)

// ATTENTION!!!: changes here may NOT be effective on AKS machine nodes (ProvisionModeAKSMachineAPI); See aksmachineinstance.go/aksmachineinstancehelpers.go.
// Refactoring for code unification is not being invested immediately.
type Template struct {
	ScriptlessCustomData      string
	ImageID                   string
	SubnetID                  string
	Tags                      map[string]*string
	CustomScriptsCustomData   string
	CustomScriptsCSE          string
	IsWindows                 bool
	StorageProfileDiskType    string
	StorageProfileIsEphemeral bool
	StorageProfilePlacement   armcompute.DiffDiskPlacement
	StorageProfileSizeGB      int32
}

type Provider struct {
	imageFamily             imagefamily.Resolver
	imageProvider           imagefamily.NodeImageProvider
	nsgProvider             *networksecuritygroup.Provider
	caBundle                *string
	clusterEndpoint         string
	tenantID                string
	subscriptionID          string
	kubeletIdentityClientID string
	resourceGroup           string
	clusterResourceGroup    string
	location                string
	provisionMode           string
}

// TODO: add caching of launch templates

// ATTENTION!!!: changes here may NOT be effective on AKS machine nodes (ProvisionModeAKSMachineAPI); See aksmachineinstance.go/aksmachineinstancehelpers.go.
// Refactoring for code unification is not being invested immediately.
func NewProvider(
	_ context.Context,
	imageFamily imagefamily.Resolver,
	imageProvider imagefamily.NodeImageProvider,
	nsgProvider *networksecuritygroup.Provider,
	caBundle *string,
	clusterEndpoint string,
	tenantID,
	subscriptionID,
	clusterResourceGroup string,
	kubeletIdentityClientID,
	resourceGroup,
	location,
	provisionMode string,
) *Provider {
	_ = "STUB: not implemented"
	return nil
}

// ATTENTION!!!: changes here may NOT be effective on AKS machine nodes (ProvisionModeAKSMachineAPI); See aksmachineinstance.go/aksmachineinstancehelpers.go.
// Refactoring for code unification is not being invested immediately.
func (p *Provider) GetTemplate(
	ctx context.Context,
	nodeClass *v1beta1.AKSNodeClass,
	nodeClaim *karpv1.NodeClaim,
	instanceType *cloudprovider.InstanceType,
	additionalLabels map[string]string,
) (*Template, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note: we check GetKubernetesVersion for errors at the start of the Create call, so this case should not happen.

// ATTENTION!!!: changes here may NOT be effective on AKS machine nodes (ProvisionModeAKSMachineAPI); See aksmachineinstance.go/aksmachineinstancehelpers.go.
// Refactoring for code unification is not being invested immediately.
func (p *Provider) getStaticParameters(
	ctx context.Context,
	instanceType *cloudprovider.InstanceType,
	nodeClass *v1beta1.AKSNodeClass,
	labels map[string]string,
) (*parameters.StaticParameters, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Remove labels kubelet can't set (e.g. kubernetes.io/*, k8s.io/* outside allowed namespaces)

// ATTENTION!!!: changes here will NOT be effective on AKS machine nodes (ProvisionModeAKSMachineAPI); See aksmachineinstance.go/aksmachineinstancehelpers.go.
// Refactoring for code unification is not being invested immediately.

func getAgentbakerNetworkPlugin(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

// ATTENTION!!!: changes here may NOT be effective on AKS machine nodes (ProvisionModeAKSMachineAPI); See aksmachineinstance.go/aksmachineinstancehelpers.go.
// Refactoring for code unification is not being invested immediately.
func (p *Provider) createLaunchTemplate(ctx context.Context, params *parameters.Parameters) (*Template, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// render user data

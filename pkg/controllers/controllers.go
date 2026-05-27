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

package controllers

import (
	"context"

	"github.com/awslabs/operatorpkg/controller"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"

	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/events"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"

	"github.com/Azure/karpenter-provider-azure/pkg/providers/azclient/azapi"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/imagefamily"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/instance"
	instancetypeprovider "github.com/Azure/karpenter-provider-azure/pkg/providers/instancetype"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/kubernetesversion"
)

func NewControllers(
	ctx context.Context,
	mgr manager.Manager,
	kubeClient client.Client,
	recorder events.Recorder,
	cloudProvider cloudprovider.CloudProvider,
	vmInstanceProvider instance.VMProvider,
	aksMachineInstanceProvider instance.AKSMachineProvider,
	kubernetesVersionProvider kubernetesversion.KubernetesVersionProvider,
	nodeImageProvider imagefamily.NodeImageProvider,
	instanceTypesProvider instancetypeprovider.Provider,
	inClusterKubernetesInterface kubernetes.Interface,
	managedKubernetesInterface kubernetes.Interface,
	managedDynamicInterface dynamic.Interface,
	subnetsClient azapi.SubnetsAPI,
	diskEncryptionSetsClient azapi.DiskEncryptionSetsAPI,
	parsedDiskEncryptionSetID *arm.ResourceID,
	networkPolicy string,
	networkPlugin string,
) []controller.Controller {
	_ = "STUB: not implemented"
	return nil
}

// TODO: nodeclaim tagging

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

package status

import (
	"context"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/azclient/azapi"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/imagefamily"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/kubernetesversion"
)

type reconciler interface {
	Reconcile(context.Context, *v1beta1.AKSNodeClass) (reconcile.Result, error)
}

type Controller struct {
	kubeClient client.Client

	kubernetesVersion *KubernetesVersionReconciler
	nodeImage         *NodeImageReconciler
	subnet            *SubnetReconciler
	validation        *ValidationReconciler
	localDNS          *LocalDNSReconciler
}

// TODO: Consider splitting this (and other similar constructors)
// into some kind of builder struct to make the calling code easier to read.
func NewController(
	kubeClient client.Client,
	kubernetesVersionProvider kubernetesversion.KubernetesVersionProvider,
	nodeImageProvider imagefamily.NodeImageProvider,
	inClusterKubernetesInterface kubernetes.Interface,
	managedKubernetesInterface kubernetes.Interface,
	managedDynamicInterface dynamic.Interface,
	subnetClient azapi.SubnetsAPI,
	diskEncryptionSetsClient azapi.DiskEncryptionSetsAPI,
	parsedDiskEncryptionSetID *arm.ResourceID,
	networkPolicy string,
	networkPlugin string,
) *Controller {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) Reconcile(ctx context.Context, nodeClass *v1beta1.AKSNodeClass) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// We use client.MergeFromWithOptimisticLock because patching a list with a JSON merge patch
// can cause races due to the fact that it fully replaces the list on a change
// Here, we are updating the status condition list

func (c *Controller) Register(_ context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Document why this magic number used. If we want to consistently use it accoss reconcilers, refactor to a reused const.
// Comments thread discussing this: https://github.com/Azure/karpenter-provider-azure/pull/729#discussion_r2006629809

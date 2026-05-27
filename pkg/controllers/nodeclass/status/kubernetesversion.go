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

	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/kubernetesversion"

	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/karpenter/pkg/utils/pretty"
)

const (
	kubernetesVersionReconcilerName = "nodeclass.kubernetesversion"
)

type KubernetesVersionReconciler struct {
	kubernetesVersionProvider kubernetesversion.KubernetesVersionProvider
	cm                        *pretty.ChangeMonitor
}

func NewKubernetesVersionReconciler(provider kubernetesversion.KubernetesVersionProvider) *KubernetesVersionReconciler {
	_ = "STUB: not implemented"
	return nil
}

func (r *KubernetesVersionReconciler) Register(_ context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Document why this magic number used. If we want to consistently use it accoss reconcilers, refactor to a reused const.
// Comments thread discussing this: https://github.com/Azure/karpenter-provider-azure/pull/729#discussion_r2006629809

// The kubernetes version reconciler will detect reasons to bump the kubernetes version:
//  1. Newly created AKSNodeClass, will select the version discovered from the API server
//  2. If a later kubernetes version is discovered from the API server, we will upgrade to it. [don't currently support rollback]
//     - Note: We will indirectly trigger an upgrade to latest image version as well, by resetting the Images readiness.
func (r *KubernetesVersionReconciler) Reconcile(ctx context.Context, nodeClass *v1beta1.AKSNodeClass) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Handles case 1: init, update kubernetes status to API server version found

// Check if there is an upgrade

// Handles case 2: Upgrade kubernetes version [Note: we set node image to not ready, since we upgrade node image when there is a kubernetes upgrade]

// We do not currently support downgrading, so keep the kubernetes version the same

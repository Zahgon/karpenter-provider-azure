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

package hash

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
)

type Controller struct {
	kubeClient client.Client
}

func NewController(kubeClient client.Client) *Controller { _ = "STUB: not implemented"; return nil }

func (c *Controller) Reconcile(ctx context.Context, nodeClass *v1beta1.AKSNodeClass) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

func (c *Controller) Register(_ context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Document why this magic number used. If we want to consistently use it accoss reconcilers, refactor to a reused const.
// Comments thread discussing this: https://github.com/Azure/karpenter-provider-azure/pull/729#discussion_r2006629809

// Updating `aksnodeclass-hash-version` annotation inside the karpenter controller means a breaking change has been made to the hash calculation.
// `aksnodeclass-hash` annotation on the AKSNodeClass will be updated, due to the breaking change, making the `aksnodeclass-hash` on the NodeClaim different from
// AKSNodeClass. Since, we cannot rely on the `aksnodeclass-hash` on the NodeClaims, due to the breaking change, we will need to re-calculate the hash and update the annotation.
// For more information on the Drift Hash Versioning: https://github.com/kubernetes-sigs/karpenter/blob/main/designs/drift-hash-versioning.md
func (c *Controller) updateNodeClaimHash(ctx context.Context, nodeClass *v1beta1.AKSNodeClass) error {
	_ = "STUB: not implemented"
	return nil
}

// Any NodeClaim that is already drifted will remain drifted if the karpenter.azure.com/nodepool-hash-version doesn't match
// Since the hashing mechanism has changed we will not be able to determine if the drifted status of the NodeClaim has changed

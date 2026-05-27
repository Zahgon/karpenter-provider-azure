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

package debug

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	karpv1 "sigs.k8s.io/karpenter/pkg/apis/v1"
)

type NodeClaimController struct {
	kubeClient client.Client
}

func NewNodeClaimController(kubeClient client.Client) *NodeClaimController {
	_ = "STUB: not implemented"
	return nil
}

func (c *NodeClaimController) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

func (c *NodeClaimController) GetInfo(nc *karpv1.NodeClaim) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *NodeClaimController) Register(_ context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

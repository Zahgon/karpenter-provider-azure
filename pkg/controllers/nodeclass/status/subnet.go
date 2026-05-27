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
	"time"

	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/azclient/azapi"
)

type SubnetReconciler struct {
	subnetClient azapi.SubnetsAPI
}

func NewSubnetReconciler(subnetClient azapi.SubnetsAPI) *SubnetReconciler {
	_ = "STUB: not implemented"
	return nil
}

const (
	SubnetUnreadyReasonNotFound     = "SubnetNotFound"
	SubnetUnreadyReasonIDInvalid    = "SubnetIDInvalid"
	SubnetUnreadyReasonUnknownError = "SubnetUnknownError"
)

const (
	subnetReconcilerName = "nodeclass.subnet"
	// we set 3 minutes for a healthy requeue interval because NRP reserves NICs for 180 seconds.
	// which means that we will not be able to free a given NIC for up to 3 minutes, for now setting it as
	// the default requeue interval at that timestamp, we may choose to redesign as we implement subnet fullness
	healthyRequeueInterval = time.Minute * 3
)

func (r *SubnetReconciler) Reconcile(ctx context.Context, nodeClass *v1beta1.AKSNodeClass) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	// TODO: Handle podSubnetID readiness here as well
	return *new(reconcile.Result), nil
}

func (r *SubnetReconciler) validateVNETSubnetID(ctx context.Context, nodeClass *v1beta1.AKSNodeClass) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Assume valid cluster subnet id
// Highly unlikely case but putting it in nonetheless

// Periodically requeue just in case subnet has been removed or later revalidating things like fullness etc

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

package garbagecollection

import (
	"context"

	"github.com/awslabs/operatorpkg/reconciler"

	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	karpv1 "sigs.k8s.io/karpenter/pkg/apis/v1"

	corecloudprovider "sigs.k8s.io/karpenter/pkg/cloudprovider"
)

type Instance struct {
	kubeClient    client.Client
	cloudProvider corecloudprovider.CloudProvider
}

func NewInstance(kubeClient client.Client, cloudProvider corecloudprovider.CloudProvider) *Instance {
	_ = "STUB: not implemented"
	return nil
}

func (c *Instance) Reconcile(ctx context.Context) (reconciler.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconciler.Result), nil
}

// We LIST instances on the CloudProvider BEFORE we grab NodeClaims/Nodes on the cluster so that we make sure that, if
// LISTing instances takes a long time, our information is more updated by the time we get to nodeclaim and Node LIST
// This works since our CloudProvider instances are deleted based on whether the NodeClaim exists or not, not vice-versa

// Garbage collect if the cloud instance has been around for more than 5 minutes, yet still no matching (per ProviderID) cluster NodeClaim.
// Note that the "match" occurs after cloudprovider.Create() returns and cluster NodeClaim ProviderID is populated as a result.
// Although, the intention of garbage collection is to clear instances with missing/deleted NodeClaim.
// This 5m is more of a grace period for newly-created instances that have yet to populate NodeClaim after.

// In the case that CreationTimestamp is irretrievable (technically, when CreationTimestamp = 0 = epoch), grace period will effectively be disabled.
// Which could be dangerous if the instance is legitimately awaiting NodeClaim population.

func (c *Instance) garbageCollect(ctx context.Context, nodeClaim *karpv1.NodeClaim, nodeList *v1.NodeList) error {
	_ = "STUB: not implemented"
	return nil
}

// Go ahead and cleanup the node if we know that it exists to make scheduling go quicker

func (c *Instance) Register(_ context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

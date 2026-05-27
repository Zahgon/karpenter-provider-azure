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

package cloudprovider

import (
	"context"

	v1 "k8s.io/api/core/v1"

	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"

	karpv1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
)

const (
	NodeClassDrift       cloudprovider.DriftReason = "NodeClassDrift"
	K8sVersionDrift      cloudprovider.DriftReason = "K8sVersionDrift"
	ImageDrift           cloudprovider.DriftReason = "ImageDrift"
	KubeletIdentityDrift cloudprovider.DriftReason = "KubeletIdentityDrift"
	ClusterConfigDrift   cloudprovider.DriftReason = "ClusterConfigDrift" // This is a catch-all for cluster-level config changes (e.g., from PUT ManagedCluster), where Karpenter does not directly "own" them.

	// TODO (charliedmcb): Use this const across code and test locations which are signaling/checking for "no drift"
	NoDrift cloudprovider.DriftReason = ""
)

func (c *CloudProvider) isNodeClassDrifted(ctx context.Context, nodeClaim *karpv1.NodeClaim, nodeClass *v1beta1.AKSNodeClass) (cloudprovider.DriftReason, error) {
	_ = "STUB: not implemented"
	// TODO: if we find more expensive checks, such as reading VMs or NICs from Azure, are being duplicated between checks, we should
	//       produce a lazy at-most-once that allows a check to cache a value for later checks to read.
	return *new(cloudprovider.DriftReason), nil
}

// This is technically not possible (as of the time of writing). IsDrifted() won't be called by core until NodeClaim is launched.

// For legacy nodes

func (c *CloudProvider) areStaticFieldsDrifted(ctx context.Context, nodeClaim *karpv1.NodeClaim, nodeClass *v1beta1.AKSNodeClass) (cloudprovider.DriftReason, error) {
	_ = "STUB: not implemented"
	return *new(cloudprovider.DriftReason), nil
}

// validate that the hash version for the AKSNodeClass is the same as the NodeClaim before evaluating for static drift

func (c *CloudProvider) isK8sVersionDrifted(ctx context.Context, nodeClaim *karpv1.NodeClaim, nodeClass *v1beta1.AKSNodeClass) (cloudprovider.DriftReason, error) {
	_ = "STUB: not implemented"
	return *new(cloudprovider.DriftReason), nil
}

// Note: this differs from AWS, as they don't check for status readiness during Drift.

// Note: we don't consider this a hard failure for drift if the KubernetesVersion is invalid/not ready to use, so we ignore returning the error here.
// We simply ensure the stored version is valid and ready to use, if we are to calculate potential Drift based on it.
// TODO (charliedmcb): I'm wondering if we actually want to have these soft-error cases switch to return an error if no-drift condition was found across all of IsDrifted.

func (c *CloudProvider) isImageVersionDrifted(
	ctx context.Context,
	nodeClaim *karpv1.NodeClaim,
	nodeClass *v1beta1.AKSNodeClass,
) (cloudprovider.DriftReason, error) {
	_ = "STUB: not implemented"
	return *new(cloudprovider.DriftReason), nil
}

// Note: this differs from AWS, as they don't check for status readiness during Drift.

// Note: we don't consider this a hard failure for drift if the Images are not ready to use, so we ignore returning the error here.
// The stored Images must be ready to use if we are to calculate potential Drift based on them.
// TODO (charliedmcb): I'm wondering if we actually want to have these soft-error cases switch to return an error if no-drift condition was found across all of IsDrifted.

// Note: this case shouldn't happen, since if there are no nodeImages, the ConditionTypeImagesReady should be false.
//     However, if it would happen, we want this to error, as it means the NodeClass is in a state it can't provision nodes.

// bail out early if core called this before the node is done creating.

// Note: not supporting drift across galleries yet, as AKS machine does not hold gallery info, as of now.
// Alternatively, could call GET VM, if not propose API changes.
// WARNING: verify whether this function support the desired gallery

// isKubeletIdentityDrifted returns drift if the kubelet identity has drifted
func (c *CloudProvider) isKubeletIdentityDrifted(ctx context.Context, nodeClaim *karpv1.NodeClaim, _ *v1beta1.AKSNodeClass) (cloudprovider.DriftReason, error) {
	_ = "STUB: not implemented"
	return *new(cloudprovider.DriftReason), nil
}

// The kubelet identity label is supposed to be set on every node, but prior to
// 1.4.0 it was not set by Karpenter. In order to avoid rolling all existing nodes,
// we don't count a missing kubelet identity as drift. This situation should resolve itself as
// image version and Kubernetes version drift is performed.
// TODO: This short-circuit should be removed post 1.4.0 (~2025-07-01)

func (c *CloudProvider) getNodeForDrift(ctx context.Context, nodeClaim *karpv1.NodeClaim) (*v1.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We do not return an error here as its expected within the lifecycle of the nodeclaims registration.
// Core's checks only for Launched status which means we've started the create, but the node doesn't nessicarially exist yet
// https://github.com/kubernetes-sigs/karpenter/blob/9877cf639e665eadcae9e46e5a702a1b30ced1d3/pkg/controllers/nodeclaim/disruption/drift.go#L51

// We do not need to check for drift if the node is being deleted.

// isMachineDrifted checks the DriftAction field of the AKS machine to determine if drift exists
func (c *CloudProvider) isMachineDrifted(ctx context.Context, nodeClaim *karpv1.NodeClaim, _ *v1beta1.AKSNodeClass) (cloudprovider.DriftReason, error) {
	_ = "STUB: not implemented"
	return *new(cloudprovider.DriftReason), nil
}

// Not an AKS machine node, no drift action to check

// Note: this is not being incorporated yet, and we currently return ClusterConfigDrift for all reasons. // Suggestion: could be extended.

// AKS machine API may add additional drift actions in the future (e.g., restart, reimage). Karpenter (core) need to support them explicitly.
// Meanwhile, re-create covers all cases.

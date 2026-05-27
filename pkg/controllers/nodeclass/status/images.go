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

	"k8s.io/client-go/kubernetes"

	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/imagefamily"

	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/karpenter/pkg/utils/pretty"
)

const (
	nodeImageReconcilerName = "nodeclass.images"

	// ConfigMap consts
	maintenanceWindowConfigMapName = "upcoming-maintenance-window"
	nodeOSMaintenanceWindowChannel = "aksManagedNodeOSUpgradeSchedule"
	configMapStartTimeFormat       = "%s-start"
	configMapEndTimeFormat         = "%s-end"
)

type NodeImageReconciler struct {
	nodeImageProvider            imagefamily.NodeImageProvider
	inClusterKubernetesInterface kubernetes.Interface
	systemNamespace              string
	cm                           *pretty.ChangeMonitor
}

func NewNodeImageReconciler(
	provider imagefamily.NodeImageProvider,
	inClusterKubernetesInterface kubernetes.Interface,
) *NodeImageReconciler {
	_ = "STUB: not implemented"
	return nil
}

func (r *NodeImageReconciler) Register(_ context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Document why this magic number used. If we want to consistently use it accoss reconcilers, refactor to a reused const.
// Comments thread discussing this: https://github.com/Azure/karpenter-provider-azure/pull/729#discussion_r2006629809

// The image version reconciler will detect reasons to bump the node image version as follows in order:
//
// Scenario A: Update all image versions to latest
//   - 1. Initializes the images versions for a newly created AKSNodeClass, based on customer configuration.
//   - 2. Indirectly handle image bump for k8s upgrade
//   - 3. Can indirectly handle bumps for any images unsupported by node features, if required to in the future
//     Note: Currently there are no node features to be handled in this way.
//   - 4. TODO: Update Images to latest if in an open maintenance window [retrieved from ConfigMap]
//
// Scenario B: Calculate images to be updated based on delta of available images
//   - 5. Handles update cases when customer changes image family, SIG usage, or other means of image selectors
//   - 6. Handles softly adding newest image version of any newly supported SKUs by Karpenter
//
// Note: While we'd currently only need to store a SKU -> version mapping in the status for avilaible Images
// we decided to store the full image ID, plus Requirements associated with it. Storing the complete ID is a simple
// and clean approach while allowing us to extend future capabilities off of it. Additionally, while the decision to
// store Requirements adds minor bloat, it also provides extra visibility into the avilaible images and how their
// selection will work, which is seen as worth the tradeoff.
func (r *NodeImageReconciler) Reconcile(ctx context.Context, nodeClass *v1beta1.AKSNodeClass) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// validate FIPS + useSIG

// sorted for consistency

// Scenario A: Check if we should do a full update to latest before processing any partial update
//
// Note: We want to handle cases 1-3 regardless of maintenance window state, since they are either
// for initialization, based off an underlying customer operation, or a different update we're
// dependant upon which would have already been preformed within its required maintenance Window.

// Case 4: Check if the maintenance window is open

// Scenario B: Calculate any partial update based on image selectors, or newly supports SKUs

// We care about the ordering of the slices here, as it translates to priority during selection, so not treating them as sets

// Handles case 1: This is a new AKSNodeClass, where images haven't been populated yet
// Handles case 2: This is indirectly handling k8s version image bump, since k8s version sets this status to false
// Handles case 3: Note: like k8s we would also indirectly handle node features that required an image version bump, but none required atm.
func imageVersionsUnready(nodeClass *v1beta1.AKSNodeClass) bool {
	_ = "STUB: not implemented"
	return false
}

// Handles case 4: check if the maintenance window is open
// TODO (charliedmcb): remove nolint on gocyclo. Added for now in order to pass "make verify"
// I think the best way to get rid of gocyclo is to break the section retrieving the maintenance window
// range from the ConfigMap into its own helper function using channel as a parameter.
//
//nolint:gocyclo
func (r *NodeImageReconciler) isMaintenanceWindowOpen(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// We fail open here, since the default case should be to upgrade

// We fail open here, since the default case should be to upgrade

// Monitoring the entire ConfigMap's data might catch more data changes than we care about. However, I think it makes sense to monitor
//     here as it does catch the entire spread of cases we care about, and will give us direct insight on the raw data.
// Note: we don't need to add the nodeclass name into the monitoring here, as we actually want the entries to collide, since
//     maintenance windows are a cluster level concept, rather that a nodeclass level type, meaning we'd have repeat redundant info
//     if scoping to the nodeclass.
// TODO: In the longer run, the maintenance window handling should be factored out into a sharable provider, rather than being contained
//     within the image controller itself.

// An empty configmap means there's no maintenance windows defined, and its up to us when to preform maintenance

// Treat empty string values as missing, since the ConfigMap may have keys present with empty values

// No maintenance window defined for aksManagedNodeOSUpgradeSchedule, so its up to us when to preform maintenance

// overrideAnyGoalStateVersionsWithExisting: will look over all the discovered images, and choose to either keep the existing version if already found in the status
// or merge the new version in. This will discard any images that are no longer selected for as well. Results in picking up new images, while also not bumping
// image versions outside of a maintenance window for existing ones.
//
// Handles case 5: users updating image selectors.
//   - Currently, this is just image family, and/or usage of SIG, which means that we should just be looking at the baseID of the images
//
// Handles case 6: We will softly add newly supported SKUs by Karpenter on their latest version
//   - Note: I think this should be re-assessed if this is the exact behavior we want to give users before any actual new SKU support is released.
//
// TODO: Need longer term design for handling newly supported versions, and other image selectors.
func overrideAnyGoalStateVersionsWithExisting(nodeClass *v1beta1.AKSNodeClass, discoveredImages []v1beta1.NodeImage) []v1beta1.NodeImage {
	_ = "STUB: not implemented"
	return nil
}

// Note: we have to range over the discovered images here, instead of converting to a baseIDMapping, to keep the ordering consistent

func mapImageBasesToImages(images []v1beta1.NodeImage) map[string]*v1beta1.NodeImage {
	_ = "STUB: not implemented"
	return nil
}

// Trims off the version suffix, and leaves just the image base id
// Examples:
//
// - CIG:
//   - Input: /CommunityGalleries/AKSUbuntu-38d80f77-467a-481f-a8d4-09b6d4220bd2/images/2204gen2containerd/versions/2022.10.03
//   - Output: /CommunityGalleries/AKSUbuntu-38d80f77-467a-481f-a8d4-09b6d4220bd2/images/2204gen2containerd
//
// - SIG:
//   - Input: /subscriptions/10945678-1234-1234-1234-123456789012/resourceGroups/AKS-Ubuntu/providers/Microsoft.Compute/galleries/AKSUbuntu/images/2204gen2containerd/versions/2022.10.03
//   - Output: /subscriptions/10945678-1234-1234-1234-123456789012/resourceGroups/AKS-Ubuntu/providers/Microsoft.Compute/galleries/AKSUbuntu/images/2204gen2containerd
func trimVersionSuffix(imageID string) string { _ = "STUB: not implemented"; return "" }

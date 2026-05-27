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

package imagefamily

import (
	"context"
	"time"

	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
	types "github.com/Azure/karpenter-provider-azure/pkg/providers/imagefamily/types"
	"github.com/patrickmn/go-cache"
	"sigs.k8s.io/karpenter/pkg/scheduling"
	"sigs.k8s.io/karpenter/pkg/utils/pretty"
)

const (
	ImageExpirationInterval    = time.Hour * 24 * 3
	ImageCacheCleaningInterval = time.Hour * 1

	sharedImageGalleryImageIDFormat = "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/galleries/%s/images/%s/versions/%s"
	communityImageIDFormat          = "/CommunityGalleries/%s/images/%s/versions/%s"
)

type NodeImage struct {
	ID           string
	Requirements scheduling.Requirements
}

type NodeImageProvider interface {
	List(ctx context.Context, nodeClass *v1beta1.AKSNodeClass) ([]NodeImage, error)
}

type provider struct {
	subscription string
	location     string

	imageVersionsClient types.CommunityGalleryImageVersionsAPI
	nodeImageVersions   types.NodeImageVersionsAPI

	nodeImagesCache *cache.Cache
	cm              *pretty.ChangeMonitor
}

func NewProvider(versionsClient types.CommunityGalleryImageVersionsAPI, location, subscription string, nodeImageVersionsClient types.NodeImageVersionsAPI, nodeImagesCache *cache.Cache) *provider {
	_ = "STUB: not implemented"
	return nil
}

// Returns the list of available NodeImages for the given AKSNodeClass sorted in priority ordering
func (p *provider) List(ctx context.Context, nodeClass *v1beta1.AKSNodeClass) ([]NodeImage, error) {
	_ = "STUB: not implemented"
	// TODO: refactor to be part of construction, since this is a karpenter setting and won't change across the process.
	return nil, nil
}

// CIG has no FIPS images; FIPS images can only be accessed through SIG
// (this won't be an error since there just aren't any FIPS images for CIG)

func (p *provider) listSIG(ctx context.Context, supportedImages []types.DefaultImageOutput) ([]NodeImage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unable to find given image version

func (p *provider) listCIG(_ context.Context, supportedImages []types.DefaultImageOutput) ([]NodeImage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *provider) cacheKey(supportedImages []types.DefaultImageOutput, k8sVersion string) (string, error) {
	_ = "STUB: not implemented"
	// Note: the kubernetes version is part of the cache key here, because we bump images on kubernetes upgrade meaning
	// we want to ensure if there is a kubernetes change we'll get fresh images if there are any.
	return "", nil
}

func (p *provider) getCIGImageID(publicGalleryURL, communityImageName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *provider) latestNodeImageVersionCommunity(publicGalleryURL, communityImageName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// BuildImageIDCIG builds a Community Image Gallery image ID
func BuildImageIDCIG(publicGalleryURL, communityImageName, imageVersion string) string {
	_ = "STUB: not implemented"
	return ""
}

// BuildImageIDSIG builds a Shared Image Gallery image ID
func BuildImageIDSIG(subscriptionID, resourceGroup, galleryName, imageDefinition, imageVersion string) string {
	_ = "STUB: not implemented"
	return ""
}

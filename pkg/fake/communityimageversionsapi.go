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

package fake

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"

	imagefamilytypes "github.com/Azure/karpenter-provider-azure/pkg/providers/imagefamily/types"
)

// TODO: no ability to simulate errors on the Pager.
type CommunityGalleryImageVersionsAPI struct {
	ImageVersions AtomicPtrSlice[armcompute.CommunityGalleryImageVersion]
}

// assert that the fake implements the interface
var _ imagefamilytypes.CommunityGalleryImageVersionsAPI = &CommunityGalleryImageVersionsAPI{}

// NewListPager returns a new pager to return the next page of CommunityGalleryImageVersionsClientListResponse
func (c *CommunityGalleryImageVersionsAPI) NewListPager(_ string, _ string, _ string, _ *armcompute.CommunityGalleryImageVersionsClientListOptions) *runtime.Pager[armcompute.CommunityGalleryImageVersionsClientListResponse] {
	_ = "STUB: not implemented"
	return nil
}

func (c *CommunityGalleryImageVersionsAPI) Reset() { _ = "STUB: not implemented"; return }

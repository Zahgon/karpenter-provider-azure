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

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
)

type NodeImageVersionsClient struct {
	client *armcontainerservice.Client
}

func NewNodeImageVersionsClient(subscriptionID string, cred azcore.TokenCredential, opts *arm.ClientOptions) (*NodeImageVersionsClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *NodeImageVersionsClient) List(ctx context.Context, location string) ([]*armcontainerservice.NodeImageVersion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FilteredNodeImages filters on two conditions
// 1. The image is the latest version for the given OS and SKU
// 2. the image belongs to a supported gallery(AKS Ubuntu or Azure Linux)
func FilteredNodeImages(nodeImageVersions []*armcontainerservice.NodeImageVersion) []*armcontainerservice.NodeImageVersion {
	_ = "STUB: not implemented"
	return nil
}

// Skip the galleries that Karpenter does not support

// isNewerVersion will return if version1 is greater than version2, note the new versioning scheme is yearmm.dd.build, previously it was yy.mm.dd without the build id.
func isNewerVersion(version1, version2 string) bool {
	_ = "STUB: not implemented"
	// Split by dots and compare each segment as an integer getting the largest vhd version
	return false
}

// If all segments are equal up to the length of the shorter version,
// the longer version is considered newer if it has additional segments
// the legacy linux versions use "yy.mm.dd" whereas new linux versions use "yymm.dd.build"

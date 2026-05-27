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

package test

import (
	"context"

	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/imagefamily"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	coretest "sigs.k8s.io/karpenter/pkg/test"
)

const (
	DefaultCIGImageVersion = "202501.02.0"
	DefaultSIGImageVersion = "202410.09.0"
)

func AKSNodeClass(overrides ...v1beta1.AKSNodeClass) *v1beta1.AKSNodeClass {
	_ = "STUB: not implemented"
	return nil
}

// In reality, these default values will be set via the defaulting done by the API server. The reason we provide them here is
// we sometimes reference a test.AKSNodeClass without applying it, and in that case we need to set the default values ourselves

// TODO: Pass in test.Options if we want to use more options within this func
func ApplyDefaultStatus(nodeClass *v1beta1.AKSNodeClass, env *coretest.Environment, useSIG bool) {
	_ = "STUB: not implemented"
	return
}

// Using the magic number 1, as it appears the Generation is always equal to 1 on the NodeClass in testing. If that appears to not be the case,
// than we should add some function for allows bumps as needed to match.

func ApplyCIGImages(nodeClass *v1beta1.AKSNodeClass) { _ = "STUB: not implemented"; return }

func ApplyCIGImagesWithVersion(nodeClass *v1beta1.AKSNodeClass, cigImageVersion string) {
	_ = "STUB: not implemented"
	return
}

func ApplySIGImages(nodeClass *v1beta1.AKSNodeClass) { _ = "STUB: not implemented"; return }

func ApplySIGImagesWithVersion(nodeClass *v1beta1.AKSNodeClass, sigImageVersion string) {
	_ = "STUB: not implemented"
	return
}

func getExpectedTestSIGImages(imageFamily string, fipsMode *v1beta1.FIPSMode, version string, kubernetesVersion string) []imagefamily.NodeImage {
	_ = "STUB: not implemented"
	return nil
}

func translateToStatusNodeImages(imageFamilyNodeImages []imagefamily.NodeImage) []v1beta1.NodeImage {
	_ = "STUB: not implemented"
	return nil
}

// sorted for consistency

func AKSNodeClassFieldIndexer(ctx context.Context) func(cache.Cache) error {
	_ = "STUB: not implemented"
	return nil
}

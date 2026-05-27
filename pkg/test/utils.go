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

const (
	// Note (charliedmcb): this already exists in test/pkg/environment/common
	// https://github.com/Azure/karpenter-provider-azure/blob/84e449787ec72268efb0c7af81ec87a6b3ee95fa/test/pkg/environment/common/setup.go#L47
	// However, I'd prefer to keep our unit test dependants self-contained instead of depending upon the e2e testing package.
	TestingFinalizer = "testing/finalizer"
)

// RandomName returns a pseudo-random resource name with a given prefix.
func RandomName(prefix string) string {
	_ = "STUB: not implemented"
	// You could make this more robust by including additional random characters.
	return ""
}

func ManagedTags(nodepoolName string) map[string]*string { _ = "STUB: not implemented"; return nil }

func ManagedTagsAKSMachine(nodepoolName string, nodeClaimName string) map[string]*string {
	_ = "STUB: not implemented"
	return nil
}

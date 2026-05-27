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

package offerings

import (
	corecloudprovider "sigs.k8s.io/karpenter/pkg/cloudprovider"
)

// Suggestion: consider merging this package with instancetype package, as both of their responsibilities deal with instance types management

func getOfferingCapacityType(offering *corecloudprovider.Offering) string {
	_ = "STUB: not implemented"
	return ""
}

func getOfferingZone(offering *corecloudprovider.Offering) string {
	_ = "STUB: not implemented"
	return ""
}

// May return nil if there is no match
func GetInstanceTypeFromVMSize(vmSize string, possibleInstanceTypes []*corecloudprovider.InstanceType) *corecloudprovider.InstanceType {
	_ = "STUB: not implemented"
	return nil
}

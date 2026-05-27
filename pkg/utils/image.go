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

package utils

import (
	"regexp"
)

var (
	sigImageIDRegex = regexp.MustCompile(`(?i)/subscriptions/(\S+)/resourceGroups/(\S+)/providers/Microsoft.Compute/galleries/(\S+)/images/(\S+)/versions/(\S+)`)
)

// WARNING: not supporting CIG images yet.
func GetAKSMachineNodeImageVersionFromImageID(imageID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Requires AKS machine API support

// Convert from "/subscriptions/10945678-1234-1234-1234-123456789012/resourceGroups/AKS-Ubuntu/providers/Microsoft.Compute/galleries/AKSUbuntu/images/2204gen2containerd/versions/2022.10.03"
// to "AKSUbuntu-2204gen2containerd-2022.10.03".
func GetAKSMachineNodeImageVersionFromSIGImageID(imageID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// subscriptionID := matches[1]
// resourceGroup := matches[2]

// if strings.Contains(prefix, windowsPrefix) {		// TODO(Windows)
// 	osVersion = extractOsVersionForWindows(definition)
// }

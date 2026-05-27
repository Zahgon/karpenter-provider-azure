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

package options

func (o *Options) Validate() error { _ = "STUB: not implemented"; return nil }

func (o *Options) validateClusterDNSIP() error { _ = "STUB: not implemented"; return nil }

func (o *Options) validateVNETGUID() error { _ = "STUB: not implemented"; return nil }

func (o *Options) validateKubeletIdentityClientID() error { _ = "STUB: not implemented"; return nil }

func (o *Options) validateNetworkingOptions() error { _ = "STUB: not implemented"; return nil }

func (o *Options) validateVnetSubnetID() error { _ = "STUB: not implemented"; return nil }

func (o *Options) validateEndpoint() error { _ = "STUB: not implemented"; return nil }

func (o *Options) validateVMMemoryOverheadPercent() error { _ = "STUB: not implemented"; return nil }

func (o *Options) validateProvisionMode() error { _ = "STUB: not implemented"; return nil }

func (o *Options) validateBatchOptions() error { _ = "STUB: not implemented"; return nil }

func (o *Options) validateRequiredFields() error { _ = "STUB: not implemented"; return nil }

func (o *Options) validateUseSIG() error { _ = "STUB: not implemented"; return nil }

// For AKS Machine API modes, we don't need SIGAccessTokenServerURL etc. given AKS Machine API would handle it.

func (o *Options) validateAdminUsername() error { _ = "STUB: not implemented"; return nil }

// Must start with a letter and only contain letters, numbers, hyphens, and underscores

// validateAdditionalTags checks that additional tags are valid according to Azure's tag rules.
// - Keys must be unique (case-insensitive)
// - Keys must not exceed 512 characters
// - Values must not exceed 256 characters
// - Keys must not contain invalid characters: <, >, %, &, \, ?, /
func (o *Options) validateAdditionalTags() error { _ = "STUB: not implemented"; return nil }

func isValidURL(u string) bool { _ = "STUB: not implemented"; return false }

// url.Parse() will accept a lot of input without error; make
// sure it's a real URL

func (o *Options) validateDiskEncryptionSetID() error { _ = "STUB: not implemented"; return nil }

// Parse with Azure SDK for validation
// arm.ParseResourceID will validate the format of the resource ID and extract its components

// Validate resource type is Microsoft.Compute/diskEncryptionSets

// Validate required fields are not empty

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

package auth

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient"
)

type Environment struct {
	// Environment is the track 1 representation of the Azure environment
	Environment *azclient.Environment
	// Cloud is the track 2 representation of the Azure environment
	Cloud cloud.Configuration
}

// readEnvironmentFromFile reads environment configuration from a JSON file.
// The expected file format is the same one that CloudProvider expects:
// https://github.com/kubernetes-sigs/cloud-provider-azure/blob/master/pkg/azclient/cloud.go#L153.
// This also happens to be the original Track1 SDK format.
func readEnvironmentFromFile(path string) (*azclient.Environment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validateEnvironment validates that required fields are present in the environment
func validateEnvironment(env *azclient.Environment) error { _ = "STUB: not implemented"; return nil }

// mapTrack1ToTrack2Environment converts azclient.Environment (Track1 and CloudProvider format, written to all nodes automatically by AKS)
// to cloud.Configuration (Track2 format).
// This is similar to what CloudProvider does here: https://github.com/kubernetes-sigs/cloud-provider-azure/blob/master/pkg/azclient/cloud.go#L121
// but we don't use that as it couples loading the file and mapping track1 to track2, in addition to allowing partial overrides
// which is less than ideal.
// TODO: We could move this upstream to azure-sdk-for-go-extensions or refactor how CloudProvider parses and share that.
func mapTrack1ToTrack2Environment(env *azclient.Environment) (cloud.Configuration, error) {
	_ = "STUB: not implemented"
	return *new(cloud.Configuration), nil
}

// environmentFromName returns a Track1-style environment from a cloud name.
// This is very similar to https://github.com/kubernetes-sigs/cloud-provider-azure/blob/master/pkg/azclient/cloud.go#L361
// but returns an error rather than defaulting to PublicCloud if the user provides an unknown cloud name.
func environmentFromName(cloudName string) (*azclient.Environment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EnvironmentFromName(cloudName string) (*Environment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IsPublic returns if the specified configuration is public.
// This takes the track2 format rather than being a method on Environment because
// usage in api/sdk contexts use the track2 format and may not have access to the
// auth.Environment struct.
func IsPublic(env cloud.Configuration) bool { _ = "STUB: not implemented"; return false }

// Shouldn't differ by case but let's be safe

// ResolveCloudEnvironment resolves the cloud environment using the following precedence:
// 1. File-based environment (AZURE_ENVIRONMENT_FILEPATH)
// 2. Known cloud names (ARM_CLOUD)
// 3. Default (Azure Public Cloud)
func ResolveCloudEnvironment(cfg *Config) (*Environment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 1. Try file-based environment first (highest precedence)

// 2. Try known cloud names (ARM_CLOUD)

// 3. Default to Azure Public Cloud -- this code shouldn't be hit regularly
// as we already default cfg.Cloud to AzurePublicCloud in the Config.Default method.

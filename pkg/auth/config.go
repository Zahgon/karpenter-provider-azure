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
	"github.com/Azure/go-autorest/autorest"
)

type cfgField struct {
	val  string
	name string
}

// ClientConfig contains all essential information to create an Azure client.
type ClientConfig struct {
	CloudName               string
	Location                string
	SubscriptionID          string
	ResourceManagerEndpoint string
	Authorizer              autorest.Authorizer
	UserAgent               string
}

// Config holds the configuration parsed from the --cloud-config flag
type Config struct {
	Cloud                    string `json:"cloud" yaml:"cloud"`
	Location                 string `json:"location" yaml:"location"`
	TenantID                 string `json:"tenantId" yaml:"tenantId"`
	SubscriptionID           string `json:"subscriptionId" yaml:"subscriptionId"`
	ResourceGroup            string `json:"resourceGroup" yaml:"resourceGroup"`
	AzureEnvironmentFilepath string `json:"azureEnvironmentFilepath" yaml:"azureEnvironmentFilepath"`
}

// BuildAzureConfig returns a Config object for the Azure clients
func BuildAzureConfig() (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

func (cfg *Config) Build() error {
	_ = "STUB: not implemented"
	// May require more than this behind the scenes: https://github.com/Azure/azure-sdk-for-go/blob/main/sdk/azidentity/README.md#defaultazurecredential
	return nil
}

// TODO: We probably can remove both of these "ARM_" fallbacks in mid 2026. We also don't currently
// use the TENANT_ID anyway so possibly could remove that too.
// Read both ARM_TENANT_ID and AZURE_TENANT_ID env vars, preferring AZURE_ when both are set

// Read both ARM_SUBSCRIPTION_ID and AZURE_SUBSCRIPTION_ID env vars, preferring AZURE_ when both are set

func (cfg *Config) Default() error {
	_ = "STUB: not implemented"
	// Default is AzurePublicCloud if not set
	return nil
}

func (cfg *Config) Validate() error {
	_ = "STUB: not implemented"
	// Validate that ARM_CLOUD and AZURE_ENVIRONMENT_FILEPATH are not both set
	return nil
}

// Setup fields and validate all of them are not empty

// Even though the config doesnt use some of these,
// its good to validate they were set in the environment

func (cfg *Config) String() string { _ = "STUB: not implemented"; return "" }

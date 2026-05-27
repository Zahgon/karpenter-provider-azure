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
)

func GetUserAgentExtension() string {
	_ = "STUB: not implemented"
	// Note: do not change "karpenter-aks/" prefix, some infra depends on it
	return ""
}

// TokenScope returns the token scope for the Azure environment, such as "https://management.azure.com/.default" (for public cloud)
func TokenScope(cloudCfg cloud.Configuration) string { _ = "STUB: not implemented"; return "" }

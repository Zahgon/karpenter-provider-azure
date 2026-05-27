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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
)

// AKSAgentPoolOptions customizes an AKS Agent Pool for testing.
type AKSAgentPoolOptions struct {
	Name          string
	ResourceGroup string
	// SubscriptionID      string
	ClusterName         string
	Count               int32
	VMSize              string
	OrchestratorVersion string
	Tags                map[string]*string
}

// AKSAgentPool creates a test AKS Agent Pool with defaults that can be overridden by AKSAgentPoolOptions.
// Overrides are applied in order, with last-write-wins semantics.
func AKSAgentPool(overrides ...AKSAgentPoolOptions) *armcontainerservice.AgentPool {
	_ = "STUB: not implemented"
	return nil
}

// Provide default values if none are set

// if options.SubscriptionID == "" {
// 	options.SubscriptionID = "test-subscription"
// }

// Create the agent pool ID using the fake helper

// Construct the AKS Agent Pool

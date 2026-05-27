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

package fake

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/azclient/azapi"
)

type AgentPoolDeleteMachinesInput struct {
	ResourceGroupName string
	ResourceName      string
	AgentPoolName     string
	AKSMachines       armcontainerservice.AgentPoolDeleteMachinesParameter
	Options           *armcontainerservice.AgentPoolsClientBeginDeleteMachinesOptions
}

type AgentPoolGetInput struct {
	ResourceGroupName string
	ResourceName      string
	AgentPoolName     string
	Options           *armcontainerservice.AgentPoolsClientGetOptions
}

type AgentPoolCreateOrUpdateInput struct {
	ResourceGroupName string
	ResourceName      string
	AgentPoolName     string
	Parameters        armcontainerservice.AgentPool
	Options           *armcontainerservice.AgentPoolsClientBeginCreateOrUpdateOptions
}

type AgentPoolsBehavior struct {
	AgentPoolDeleteMachinesBehavior MockedLRO[AgentPoolDeleteMachinesInput, armcontainerservice.AgentPoolsClientDeleteMachinesResponse]
	AgentPoolGetBehavior            MockedFunction[AgentPoolGetInput, armcontainerservice.AgentPoolsClientGetResponse]
}

var AKSAgentPoolsAPIErrorFromAKSAgentPoolNotFound = &azcore.ResponseError{
	ErrorCode:  "NotFound",
	StatusCode: http.StatusNotFound,
}

// AKSAgentPoolsAPIErrorFromAKSMachineNotFound creates the specific error for when machines cannot be found during delete
func AKSAgentPoolsAPIErrorFromAKSMachineNotFound(agentPoolName string, validMachines []string) error {
	_ = "STUB: not implemented"
	return nil
}

// assert that the fake implements the interface
var _ azapi.AKSAgentPoolsAPI = &AKSAgentPoolsAPI{}

type AKSAgentPoolsAPI struct {
	AgentPoolsBehavior
	aksDataStorage *AKSDataStorage
}

func NewAKSAgentPoolsAPI(aksDataStorage *AKSDataStorage) *AKSAgentPoolsAPI {
	_ = "STUB: not implemented"
	return nil
}

// Reset must be called between tests otherwise tests will pollute each other.
func (c *AKSAgentPoolsAPI) Reset() { _ = "STUB: not implemented"; return }

func (c *AKSAgentPoolsAPI) Get(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *armcontainerservice.AgentPoolsClientGetOptions) (armcontainerservice.AgentPoolsClientGetResponse, error) {
	_ = "STUB: not implemented"
	return *new(armcontainerservice.AgentPoolsClientGetResponse), nil
}

// Already procedural, and is a fake
//
//nolint:gocyclo
func (c *AKSAgentPoolsAPI) BeginDeleteMachines(
	ctx context.Context,
	resourceGroupName string,
	resourceName string,
	agentPoolName string,
	aksMachines armcontainerservice.AgentPoolDeleteMachinesParameter,
	options *armcontainerservice.AgentPoolsClientBeginDeleteMachinesOptions,
) (*runtime.Poller[armcontainerservice.AgentPoolsClientDeleteMachinesResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if agent pool exists before deleting machines

// First, validate that all machines exist and collect valid/invalid machines

// Collect all existing machines in the agent pool for error message

// Check if this machine belongs to the same agent pool

// Check if requested machines exist

// If any machines are invalid, return the InvalidParameter error

// Delete only the valid machines

func MkAgentPoolID(resourceGroupName string, clusterName string, agentPoolName string) string {
	_ = "STUB: not implemented"
	return ""
}

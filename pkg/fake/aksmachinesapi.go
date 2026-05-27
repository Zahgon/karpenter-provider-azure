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
	fakesync "github.com/Azure/karpenter-provider-azure/pkg/fake/sync"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/azclient/aksmachinesheaderbatch"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/azclient/azapi"
)

// AKSDataStorage contains the shared data stores for both AKS agent pools and machines
type AKSDataStorage struct {
	AgentPools  *fakesync.Map[string, armcontainerservice.AgentPool]
	AKSMachines *fakesync.Map[string, armcontainerservice.Machine]
}

// NewAKSDataStorage creates a new instance of shared data stores
func NewAKSDataStorage() *AKSDataStorage { _ = "STUB: not implemented"; return nil }

type AKSMachineCreateOrUpdateInput struct {
	ResourceGroupName string
	ResourceName      string
	AgentPoolName     string
	AKSMachineName    string
	AKSMachine        armcontainerservice.Machine
	Options           *armcontainerservice.MachinesClientBeginCreateOrUpdateOptions
}

type AKSMachineGetInput struct {
	ResourceGroupName string
	ResourceName      string
	AgentPoolName     string
	AKSMachineName    string
	Options           *armcontainerservice.MachinesClientGetOptions
}

type AKSMachineListInput struct {
	ResourceGroupName string
	ResourceName      string
	AgentPoolName     string
	Options           *armcontainerservice.MachinesClientListOptions
}

type AKSMachinesBehavior struct {
	AKSMachineCreateOrUpdateBehavior   MockedLRO[AKSMachineCreateOrUpdateInput, armcontainerservice.MachinesClientCreateOrUpdateResponse]
	AKSMachineGetBehavior              MockedFunction[AKSMachineGetInput, armcontainerservice.MachinesClientGetResponse]
	AKSMachineNewListPagerBehavior     MockedFunction[AKSMachineListInput, *runtime.Pager[armcontainerservice.MachinesClientListResponse]]
	AfterPollProvisioningErrorOverride *armcontainerservice.ErrorDetail

	// BatchMachineErrorFunc, if set, is called during batch creation to determine per-machine
	// errors. It receives a machine name and returns an error code and message; if the error code
	// is non-empty, the machine is treated as failed. Machines with empty error code succeed.
	// When any per-machine errors are returned, the fake produces a batch error response
	// (BatchMachineClientError or BatchMachineInternalServerError) matching the real Azure API format.
	BatchMachineErrorFunc func(machineName string) (errorCode string, errorMessage string)
}

var AKSMachineAPIErrorFromAKSMachineNotFound = &azcore.ResponseError{
	ErrorCode:  "NotFound",
	StatusCode: http.StatusNotFound,
}
var AKSMachineAPIErrorFromAKSMachinesPoolNotFound = &azcore.ResponseError{
	ErrorCode:  "NotFound",
	StatusCode: http.StatusNotFound,
}
var AKSMachineAPIErrorFromAKSMachineImmutablePropertyChangeAttempted = &azcore.ResponseError{
	ErrorCode:  "OperationNotAllowed",
	StatusCode: http.StatusBadRequest,
}
var AKSMachineAPIErrorAny = &azcore.ResponseError{
	ErrorCode: "SomeRandomError",
}

func AKSMachineAPIErrorVMSizeNotSupported(vmSize, subscription, location string) *azcore.ResponseError {
	_ = "STUB: not implemented"
	return nil
}

func AKSMachineAPIErrorVMSizeNotSupportedBadRequest(vmSize, subscription, location string) *azcore.ResponseError {
	_ = "STUB: not implemented"
	return nil
}

// statusCode is always BadRequest today but kept as a parameter for generality
func newResponseError(errorCode string, statusCode int, message string) *azcore.ResponseError {
	_ = "STUB: not implemented"
	return nil
}

func AKSMachineAPIProvisioningErrorSkuNotAvailable(sku string, location string) *armcontainerservice.ErrorDetail {
	_ = "STUB: not implemented"
	return nil
}

func AKSMachineAPIProvisioningErrorLowPriorityCoresQuota(location string) *armcontainerservice.ErrorDetail {
	_ = "STUB: not implemented"
	return nil
}

func AKSMachineAPIProvisioningErrorOverconstrainedZonalAllocation() *armcontainerservice.ErrorDetail {
	_ = "STUB: not implemented"
	return nil
}

func AKSMachineAPIProvisioningErrorOverconstrainedAllocation() *armcontainerservice.ErrorDetail {
	_ = "STUB: not implemented"
	return nil
}

func AKSMachineAPIProvisioningErrorAllocationFailed() *armcontainerservice.ErrorDetail {
	_ = "STUB: not implemented"
	return nil
}

func AKSMachineAPIProvisioningErrorVMFamilyQuotaExceeded(location string, familyName string, currentLimit int32, currentUsage int32, additionalRequired int32, newLimitRequired int32) *armcontainerservice.ErrorDetail {
	_ = "STUB: not implemented"
	return nil
}

func AKSMachineAPIProvisioningErrorTotalRegionalCoresQuota(location string) *armcontainerservice.ErrorDetail {
	_ = "STUB: not implemented"
	return nil
}

func AKSMachineAPIProvisioningErrorZoneAllocationFailed(sku string, zone string) *armcontainerservice.ErrorDetail {
	_ = "STUB: not implemented"
	return nil
}

func AKSMachineAPIProvisioningErrorAny() *armcontainerservice.ErrorDetail {
	_ = "STUB: not implemented"
	return nil
}

// assert that the fake implements the interface
var _ azapi.AKSMachinesAPI = &AKSMachinesAPI{}

type AKSMachinesAPI struct {
	AKSMachinesBehavior
	aksDataStorage *AKSDataStorage
}

// NewAKSMachinesAPI creates a new AKSMachinesAPI instance with shared data stores
func NewAKSMachinesAPI(aksDataStorage *AKSDataStorage) *AKSMachinesAPI {
	_ = "STUB: not implemented"
	return nil
}

// Reset must be called between tests otherwise tests will pollute each other.
func (c *AKSMachinesAPI) Reset() { _ = "STUB: not implemented"; return }

func (c *AKSMachinesAPI) BeginCreateOrUpdate(
	ctx context.Context,
	resourceGroupName string,
	resourceName string,
	agentPoolName string,
	aksMachineName string,
	parameters armcontainerservice.Machine,
	options *armcontainerservice.MachinesClientBeginCreateOrUpdateOptions,
) (*runtime.Poller[armcontainerservice.MachinesClientCreateOrUpdateResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate parent AgentPool

// If batch entries are present, create a machine for each entry.
// This simulates what the real Azure API does when it reads the BatchPutMachine header.

// Non-batch path: single machine creation (original behavior)

// createSingleMachine handles non-batch (single machine) creation.
func (c *AKSMachinesAPI) createSingleMachine(input *AKSMachineCreateOrUpdateInput, parameters armcontainerservice.Machine) (*runtime.Poller[armcontainerservice.MachinesClientCreateOrUpdateResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if AKS machine already exists, if so, consider this an update than a create

// Default values + update status, for sync phase

// Update status, for async phase

// createBatchMachines handles batch creation: creates one machine per entry
// using the shared template body + per-machine zones/tags from the batch entries.
// This simulates what the real Azure API does when reading the BatchPutMachine header.
//
// If BatchMachineErrorFunc is set, it is called for each machine to determine per-machine
// errors. Failed machines are NOT created; successful machines are stored normally.
// The error response matches the real Azure batch API format:
//   - If any client error (4xx-style code) is present: returns 400 BatchMachineClientError
//   - If only internal errors (5xx-style): returns 500 BatchMachineInternalServerError
//   - If all succeed: returns success as before
func (c *AKSMachinesAPI) createBatchMachines(input *AKSMachineCreateOrUpdateInput, template armcontainerservice.Machine, entries []aksmachinesheaderbatch.MachineEntry) (*runtime.Poller[armcontainerservice.MachinesClientCreateOrUpdateResponse], error) {
	_ = "STUB: not implemented"
	// Collect per-machine errors if the error function is set
	return nil, nil
}

// If there are per-machine errors, build and return a batch error response

// Enrich input.AKSMachine with the primary entry's zones/tags so that
// CalledWithInput captures meaningful per-machine data (not the cleared template).

// Return the poller for the primary (first) machine, matching coordinator behavior

// createOneBatchMachine builds and stores a single machine from a batch entry.
func (c *AKSMachinesAPI) createOneBatchMachine(input *AKSMachineCreateOrUpdateInput, template armcontainerservice.Machine, entry aksmachinesheaderbatch.MachineEntry) (armcontainerservice.Machine, error) {
	_ = "STUB: not implemented"

	// Shallow-copy Properties to avoid mutating the shared template across loop iterations.
	return *new(armcontainerservice.Machine), nil
}

// Apply per-machine zones from the batch entry

// Apply per-machine tags from the batch entry

// Check if AKS machine already exists — if so, check for immutable property conflicts

// fakeBatchMachineError represents a per-machine error for building fake batch error responses.
type fakeBatchMachineError struct {
	code    string
	message string
	target  string
}

// batchErrorDetailJSON is the JSON shape for a per-machine error detail in batch API responses.
type batchErrorDetailJSON struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Target  string `json:"target"`
}

// buildFakeBatchError constructs an *azcore.ResponseError matching the real Azure batch error format.
// Mimics the JoinBatchPutMachineErrors logic tested in the wiki:
//   - If any error code looks like a client error → 400 BatchMachineClientError with details[] at top level
//   - If only internal errors → 500 BatchMachineInternalServerError with details[] JSON-encoded in message
func buildFakeBatchError(errors []fakeBatchMachineError) *azcore.ResponseError {
	_ = "STUB: not implemented"
	return nil
}

// Client errors are non-Internal* codes (e.g., InvalidParameter, SkuNotAvailable, etc.)

func (c *AKSMachinesAPI) updateExistingAKSMachine(input *AKSMachineCreateOrUpdateInput, existing armcontainerservice.Machine, aksMachine armcontainerservice.Machine) (*runtime.Poller[armcontainerservice.MachinesClientCreateOrUpdateResponse], error) {
	_ = "STUB: not implemented"
	// Check ETag for optimistic concurrency control
	return nil, nil
}

// Validate immutable properties not violated

// Patch with new values

// Update ETag after successful update

// Write the updated machine

func (c *AKSMachinesAPI) simulateCreateStatusAtAsync(aksMachine armcontainerservice.Machine) (armcontainerservice.Machine, error) {
	_ = "STUB: not implemented"
	return *new(armcontainerservice.Machine), nil
}

func (c *AKSMachinesAPI) Get(
	ctx context.Context,
	resourceGroupName string,
	resourceName string,
	agentPoolName string,
	aksMachineName string,
	options *armcontainerservice.MachinesClientGetOptions,
) (armcontainerservice.MachinesClientGetResponse, error) {
	_ = "STUB: not implemented"
	return *new(armcontainerservice.MachinesClientGetResponse), nil
}

// Validate that the agent pool exists before attempting to get machines

// First try direct lookup using the standard machine ID format

func (c *AKSMachinesAPI) NewListPager(
	resourceGroupName string,
	resourceName string,
	agentPoolName string,
	options *armcontainerservice.MachinesClientListOptions,
) *runtime.Pager[armcontainerservice.MachinesClientListResponse] {
	_ = "STUB: not implemented"
	return nil
}

// For this fake implementation, return a simple pager that lists all AKS machines

// Single page for fake implementation

// Check if the agent pool exists when fetching the page

// AKS machines pool not found. Return ARM not found error to match real API behavior.

// doesAgentPoolExists checks if the agent pool exists
func (c *AKSMachinesAPI) doesAgentPoolExists(resourceGroupName, resourceName, agentPoolName string) bool {
	_ = "STUB: not implemented"
	return false
}

// No store means agent pool doesn't exist

// Return true ONLY if agent pool is actually found

// validateMachinePropertyChanges checks if the immutable properties of an AKS machine are being changed
//
//nolint:gocyclo
func (c *AKSMachinesAPI) doImmutablePropertiesChanged(existing, incoming *armcontainerservice.Machine) bool {
	_ = "STUB: not implemented"
	return false
}

// Skip validation if properties are missing

// Check VM size changes (not allowed)

// Check priority changes (not allowed)

// Check zone changes (not allowed)

func MkMachineID(resourceGroupName string, clusterName string, agentPoolName string, aksMachineName string) string {
	_ = "STUB: not implemented"
	return ""
}

// setDefaultMachineValues sets comprehensive default values for AKS machine creation
// Note: this may not be accurate. But likely sufficient for testing.
func (c *AKSMachinesAPI) setDefaultMachineValues(machine *armcontainerservice.Machine, resourceGroupName string, agentPoolName string) {
	_ = "STUB: not implemented"
	return
}

// Set Status with creation timestamp

// Set ProvisioningState

// Set Priority - default to Regular if not set

// Set ResourceID - simulates VM resource ID
// vmName = aks-<machinesPoolName>-<aksMachineName>-########-vm

// NodeImageVersion is now set directly on the machine template by the caller.
// Only apply default if not provided.

// Default node image version if none provided

// Set ETag for optimistic concurrency control

// deepCopyMachine returns a fully independent copy of an AKS Machine via JSON
// round-trip, simulating the serialization boundary of a real HTTP call.
func deepCopyMachine(src armcontainerservice.Machine) armcontainerservice.Machine {
	_ = "STUB: not implemented"
	return *new(armcontainerservice.Machine)
}

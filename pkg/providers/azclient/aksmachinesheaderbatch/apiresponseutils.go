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

package aksmachinesheaderbatch

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/instance/offerings"
)

const (
	// BatchMachineClientError and BatchMachineInternalServerError are error codes returned by AKS RP for batch PUT operations.
	// They indicate that the batch request was processed, but some or all machines in the batch failed with per-machine errors in the details.
	BatchMachineClientError         = "BatchMachineClientError"
	BatchMachineInternalServerError = "BatchMachineInternalServerError"
)

// extractPerMachineErrors takes an input map that should be pre-populated with all machine names in the batch,
// then fills in the corresponding HandlableError for each machine based on the API error.
// If the API error code is not a recognized batch error code, the whole error (top-level code/message) will be applied to all machines.
func extractPerMachineErrors(apiError error, perMachineErrors map[string]*offerings.HandlableError) error {
	_ = "STUB: not implemented"
	// Design notes:
	// - The logic/API assumptions are based on the contract noted in the design doc for batch (0010-aks-machines-batch-creation.md).
	// - This function assumes that the contract is upheld strictly. Any deviation will result in an error, aborting the operation and failing the whole batch.
	// - For the case where the whole error will be applied to all machines, we made an ASSUMPTION that the top-level code/message is enough to be used by handle logic.
	// - perMachineErrors being pre-populated will help validate that no unexpected machine names are referenced by the API error.
	// - The contract and parsing being fragile is a known issue. It is due to server-side and Azure SDK's technical limitations.
	// 	 Given the impact, the accepted trade-off is that it is not worth solving now. Upcoming ARM batch integration is expected to be a natural resolution.
	//   See design doc for details.
	return nil
}

// Recognized batch error codes — parse per-machine details from the response body.

// Not a recognized batch error code — parse top-level code + message and apply to all machines.

// parsePerMachineDetails extracts per-machine error details from a batch API error response body.
func parsePerMachineDetails(respErr *azcore.ResponseError) ([]*armcontainerservice.ErrorDetail, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note: both batch error codes are AKS RP responses (always flat format), so the response body
// unmarshals directly into armcontainerservice.ErrorDetail.

// Use different parsing logic based on the error code, since the location of details[] is
// different for BatchMachineInternalServerError vs BatchMachineClientError. This is per the contract.

// Details are JSON-encoded inside the message field.

// Details are at the top level.

// parseTopLevelError extracts the top-level code + message from an API error response body.
func parseTopLevelError(respErr *azcore.ResponseError) (*offerings.HandlableError, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The top-level code/message may be in different formats based on whether the error is from
// ARM infrastructure (wrapped) vs directly from AKS RP (flat). We will try known formats in
// order, but ultimately if we cannot parse it, we will return an error instead of silently
// losing the information.

// Try wrapped first (ARM infrastructure errors).
// E.g., {"error": {"code": "ResourceGroupNotFound", "message": "Resource group 'rg' could not be found."}}

// Try flat (AKS RP errors).
// E.g., {"code": "BadRequest", "message": "Agent pool 'np' is not in 'Machines' mode.", "details": [...]}

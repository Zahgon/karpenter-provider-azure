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
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/instance/offerings"
	"github.com/Azure/karpenter-provider-azure/pkg/utils/batcher"
)

// aksMachineCreatePayload is the per-request payload to be batched on.
type aksMachineCreatePayload struct {
	resourceGroupName string
	resourceName      string
	agentPoolName     string
	machineName       string
	machineBody       *armcontainerservice.Machine
}

// executor sends batches to Azure using the BatchPutMachine HTTP header.
// It transforms a pending batch into a single API call, then distributes
// per-machine results back to each request's channel.
type executor struct {
	realClient AKSMachinesCreateAPI
}

func newExecutor(realClient AKSMachinesCreateAPI) *executor { _ = "STUB: not implemented"; return nil }

// executeBatch is the batcher.ExecuteBatch — it sends a batch to Azure as one
// API call, then distributes results back to each request's channel.
func (e *executor) executeBatch(ctx context.Context, batch *batcher.Batch[aksMachineCreatePayload, *offerings.HandlableError]) {
	_ = "STUB: not implemented"
	return
}

// Attach batch header for the real Azure API.

// Also mirror entries into context for fakes/testing.
// See WithFakeBatchEntries for why this duplication is necessary.

// Use resource params from the first request (all requests in a batch
// share the same resource path due to the key function).

// Build the template body to be used as a base for the batch.

// Note: We discard the SDK poller - callers should use the GET-based poller instead

// Extract per-machine errors from the parsed API error.

// Default to no error for each machine.
// Also, this is to catch the case where API error is erroneously referencing a non-existent machine.

// All machines in the batch are successfully created (in sync phase)

// distributePerMachine sends individual API errors back to each request based on the map of machineName → HandlableError.
// Machines with nil HandlableError are treated as successes. Returns the count of successes and failures.
func distributePerMachine(batch *batcher.Batch[aksMachineCreatePayload, *offerings.HandlableError], perMachineErrors map[string]*offerings.HandlableError) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// distributeSuccess sends a nil-nil response to all requests.
// Returns the count of requests notified.
func distributeSuccess(batch *batcher.Batch[aksMachineCreatePayload, *offerings.HandlableError]) int {
	_ = "STUB: not implemented"
	return 0
}

// distributeOperationalError sends the same operational error (via Err) to all requests.
// Use this only for errors that are not API responses (e.g., header build failure, parse failure).
func distributeOperationalError(batch *batcher.Batch[aksMachineCreatePayload, *offerings.HandlableError], err error) {
	_ = "STUB: not implemented"
	return
}

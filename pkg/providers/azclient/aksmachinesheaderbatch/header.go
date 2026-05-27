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
	"github.com/Azure/karpenter-provider-azure/pkg/providers/instance/offerings"
	"github.com/Azure/karpenter-provider-azure/pkg/utils/batcher"
)

// Header-specific types for the BatchPutMachine HTTP header.

// batchPutMachineHeader is the JSON structure sent via HTTP header to Azure.
type batchPutMachineHeader struct {
	BatchMachines []MachineEntry `json:"batchMachines"`
}

type MachineEntry struct {
	MachineName string            `json:"machineName"`
	Zones       []string          `json:"zones"`
	Tags        map[string]string `json:"tags"`
}

// buildBatchHeader creates the JSON for the BatchPutMachine HTTP header
func buildBatchHeader(batch *batcher.Batch[aksMachineCreatePayload, *offerings.HandlableError]) (string, []MachineEntry, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Helpers to convert Azure SDK pointer types to concrete values.

func extractZones(zones []*string) []string { _ = "STUB: not implemented"; return nil }

func extractTags(tags map[string]*string) map[string]string { _ = "STUB: not implemented"; return nil }

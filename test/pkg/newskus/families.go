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

package newskus

import (
	"github.com/Azure/karpenter-provider-azure/pkg/providers/instancetype"
)

// keptFeatures are additive features that represent fundamentally different hardware
// and should NOT be stripped when computing canonical families.
var keptFeatures = map[rune]bool{
	'a': true, // AMD
	'p': true, // ARM
	'm': true, // memory-intensive
}

// CanonicalFamily computes a canonical family key from a VM size name.
// It strips the "Standard_" prefix, parses with skewer, and rebuilds
// a key keeping only hardware-significant additive features (a, p, m).
// If the name can't be parsed, falls back to the SKU entry's Azure family.
func CanonicalFamily(entry instancetype.SKUEntry) string { _ = "STUB: not implemented"; return "" }

// VcpuCount extracts the vCPU count from a VM size name.
func VcpuCount(name string) int { _ = "STUB: not implemented"; return 0 }

// FeatureCount returns the number of additive features in a VM size name.
func FeatureCount(name string) int { _ = "STUB: not implemented"; return 0 }

// IsConstrainedCPU returns true if a VM size name represents a constrained-CPU SKU
// (e.g. Standard_E4-2as_v7 has 4 vCPUs but only 2 active).
func IsConstrainedCPU(name string) bool { _ = "STUB: not implemented"; return false }

// PickRepresentativeSize picks the best VM size from a group sharing a canonical family:
// 1. Fewest additive features (simplest variant)
// 2. Smallest vCPU count
// 3. Alphabetical name (deterministic tiebreaker)
// Then returns the first entry with >= 4 vCPUs, or falls back to the first overall.
// Constrained-CPU SKUs are excluded because we don't want to pick a representative that appears to have high vCPU count but is actually constrained.
func PickRepresentativeSize(skus []instancetype.SKUEntry) instancetype.SKUEntry {
	_ = "STUB: not implemented"
	return *new(instancetype.SKUEntry)
}

// Filter out constrained-CPU SKUs (e.g. Standard_E4-2as_v7)

// All SKUs are constrained — fall back to original list

// Prefer >= 4 vCPUs over < 4

// Prefer smaller vCPU as long as it's over 4

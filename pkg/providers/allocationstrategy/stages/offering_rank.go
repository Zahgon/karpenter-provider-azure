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

package stages

import (
	"context"

	corecloudprovider "sigs.k8s.io/karpenter/pkg/cloudprovider"
)

type defaultOfferingRankStage struct{}

func NewDefaultOfferingRankStage() Stage { _ = "STUB: not implemented"; return *new(Stage) }

func (s *defaultOfferingRankStage) Process(_ context.Context, instanceOfferings []InstanceOffering) []InstanceOffering {
	_ = "STUB: not implemented"
	return nil
}

func rankOfferings(offerings corecloudprovider.Offerings) {
	_ = "STUB: not implemented"
	// Shuffle before the stable sort so that offerings tied on every comparison
	// dimension (price, capacity type, placement scope) end up in random order.
	// This avoids concentrating launches in the lexically first zone when zonal
	// offerings are otherwise equivalent. Non-cryptographic randomness is
	// intentional here.
	return
}

func firstOffering(instanceOffering InstanceOffering) *corecloudprovider.Offering {
	_ = "STUB: not implemented"
	return nil
}

// compareOfferings returns a negative value when i should sort before j. The
// default precedence order is: non-nil offerings, lowest price, capacity type
// with spot preferred over on-demand, then placement scope with zonal preferred
// over regional. Since capacity type is evaluated before placement scope,
// regional spot is preferred over zonal on-demand when their prices are equal.
// Zones are intentionally not compared here; rankOfferings shuffles before
// stable sorting so otherwise equivalent offerings are spread across zones over
// time.
func compareOfferings(i, j *corecloudprovider.Offering) int { _ = "STUB: not implemented"; return 0 }

// Prefer the lower priced offering.

// Preserve Karpenter's spot-before-on-demand tie-break.

// Prefer zonal over regional within the same price and capacity type.

func placementScopeRank(offering *corecloudprovider.Offering) int {
	_ = "STUB: not implemented"
	return 0
}

// This should be unreachable for provider-generated offerings. Rank
// malformed offerings last so they never outrank known placement scopes.

func capacityTypeRank(offering *corecloudprovider.Offering) int {
	_ = "STUB: not implemented"
	return 0
}

func instanceOfferingName(instanceOffering InstanceOffering) string {
	_ = "STUB: not implemented"
	return ""
}

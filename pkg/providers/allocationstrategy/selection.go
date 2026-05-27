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

package allocationstrategy

import (
	corecloudprovider "sigs.k8s.io/karpenter/pkg/cloudprovider"
)

// Selection is the outcome of a client-side allocation decision: a single
// concrete instance type and offering chosen to satisfy a NodeClaim. Selection
// hides how the choice was made (e.g. local ranking + head-pick today, or a
// future client-side strategy that consults external Azure APIs to inform the
// decision).
//
// Both InstanceType and Offering are non-nil for any Selection returned by a
// Provider; Provider.Allocate returns nil when no compatible offering is
// available rather than a Selection with nil fields. The accessor methods
// below rely on this invariant.
type Selection struct {
	InstanceType *corecloudprovider.InstanceType
	Offering     *corecloudprovider.Offering
}

// CapacityType returns the karpenter.sh/capacity-type value of the chosen offering.
func (s *Selection) CapacityType() string { _ = "STUB: not implemented"; return "" }

// Zone returns the topology.kubernetes.io/zone value of the chosen offering.
func (s *Selection) Zone() string { _ = "STUB: not implemented"; return "" }

// PlacementScope returns the karpenter.azure.com/placement-scope value of the
// chosen offering, falling back to inferring it from the offering's zone when
// the label is absent.
func (s *Selection) PlacementScope() string { _ = "STUB: not implemented"; return "" }

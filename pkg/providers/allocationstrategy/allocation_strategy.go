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
	"context"

	corecloudprovider "sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/scheduling"
)

type Provider interface {
	// Allocate selects a single instance type and offering for a NodeClaim.
	// Returns nil when no compatible offering is available.
	//
	// This interface models client-side allocation: the provider chooses one
	// concrete offering and the caller is responsible for provisioning. It
	// accommodates future client-side strategies that consult external Azure
	// APIs (e.g. placement-score, capacity advice) to inform the decision,
	// since the contract here is just "given candidates, return a choice."
	//
	// It does NOT accommodate future server-side APIs that combine decision
	// and provisioning into a single call (Fleet-like APIs): such APIs do
	// not return a separable Selection that the caller then provisions, so
	// they would replace this provider rather than implement it. Those paths
	// can still reuse the default filtering/ranking stages when they need an
	// ordered candidate set to pass to the server-side API.
	Allocate(ctx context.Context, instanceTypes []*corecloudprovider.InstanceType, requirements scheduling.Requirements) *Selection
}

var _ Provider = &DefaultProvider{}

type DefaultProvider struct{}

func NewProvider() *DefaultProvider { _ = "STUB: not implemented"; return nil }

func (p *DefaultProvider) Allocate(ctx context.Context, instanceTypes []*corecloudprovider.InstanceType, requirements scheduling.Requirements) *Selection {
	_ = "STUB: not implemented"
	return nil
}

func (p *DefaultProvider) FilterInstanceOfferings(ctx context.Context, instanceOfferings []InstanceOffering, requirements scheduling.Requirements) []InstanceOffering {
	_ = "STUB: not implemented"
	return nil
}

// Keep offering ranking in a single stage so future customizable allocation strategy work can swap or parameterize the ranker
// without introducing multiple reorder stages where the last reorder wins.

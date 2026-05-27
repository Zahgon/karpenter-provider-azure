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

	"github.com/Azure/karpenter-provider-azure/pkg/providers/imagefamily/types"
	"github.com/Azure/karpenter-provider-azure/pkg/provisionclients/models"
)

type NodeBootstrappingGetInput struct {
	Params *models.ProvisionValues
}

type NodeBootstrappingBehavior struct {
	NodeBootstrappingGetBehavior MockedFunction[NodeBootstrappingGetInput, types.NodeBootstrapping]
}

// NodeBootstrappingAPI implements a fake version of the imagefamily.types.NodeBootstrappingAPI
// for testing purposes.
type NodeBootstrappingAPI struct {
	NodeBootstrappingBehavior
	SimulateDown bool
}

// Ensure NodeBootstrappingAPI implements the types.NodeBootstrappingAPI interface
var _ types.NodeBootstrappingAPI = &NodeBootstrappingAPI{}

// Reset must be called between tests otherwise tests will pollute each other.
func (n *NodeBootstrappingAPI) Reset() { _ = "STUB: not implemented"; return }

// Get implements the NodeBootstrappingAPI interface for testing
func (n *NodeBootstrappingAPI) Get(ctx context.Context, params *models.ProvisionValues) (types.NodeBootstrapping, error) {
	_ = "STUB: not implemented"
	return *new(types.NodeBootstrapping), nil
}

//nolint:gocyclo
func validateProvisionProfile(p *models.ProvisionProfile) error {
	_ = "STUB: not implemented"
	return nil
}

func validateProvisionHelperValues(p *models.ProvisionHelperValues) error {
	_ = "STUB: not implemented"
	return nil
}

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

package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
)

type RBACManager struct {
	subscriptionID string
	client         *armauthorization.RoleAssignmentsClient
}

// NewRBACManager builds a client with the provided TokenCredential.
func NewRBACManager(subscriptionID string, cred azcore.TokenCredential) (*RBACManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EnsureRole assigns roleDefinitionID to principalID at scope if not already present.
// It lists for the scope and returns nil if a matching assignment exists.
func (r *RBACManager) EnsureRole(ctx context.Context, scope, roleDefinitionID, principalID string) error {
	_ = "STUB: not implemented"
	return nil
}

// EnsureRoleWithPrincipalType assigns roleDefinitionID to principalID at scope with optional principalType.
// Setting principalType helps handle replication delays when creating principals and immediately assigning roles.
// See https://aka.ms/docs-principaltype for more information.
func (r *RBACManager) EnsureRoleWithPrincipalType(ctx context.Context, scope, roleDefinitionID, principalID, principalType string) error {
	_ = "STUB: not implemented"
	// Quick scan to avoid duplicates
	return nil
}

// Already assigned

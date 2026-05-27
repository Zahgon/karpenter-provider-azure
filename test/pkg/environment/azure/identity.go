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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	containerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func (env *Environment) GetClusterIdentity(ctx context.Context) *containerservice.ManagedClusterIdentity {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) GetKarpenterWorkloadIdentity(ctx context.Context) string {
	_ = "STUB: not implemented"
	return ""
}

// matches AZURE_KARPENTER_USER_ASSIGNED_IDENTITY_NAME

// getCurrentUserPrincipalID gets the principal ID of the current authenticated identity
func (env *Environment) GetCurrentUserPrincipalID(ctx context.Context, cred azcore.TokenCredential) string {
	_ = "STUB: not implemented"
	return ""
}

// ExpectCreatedManagedIdentity creates a new user-assigned managed identity
func (env *Environment) ExpectCreatedManagedIdentity(ctx context.Context, identityName string) *armmsi.Identity {
	_ = "STUB: not implemented"
	return nil
}

// Note: we don't register for cleanup in the env.tracker, in case there are more tests to run. We don't want to break the cluster by deleting the kubelet identity.
// It will get deleted when the node resource group is cleaned up.

// ExpectUpdatedManagedClusterKubeletIdentityAsync updates the kubelet identity of a managed cluster asynchronously
func (env *Environment) ExpectUpdatedManagedClusterKubeletIdentityAsync(ctx context.Context, newIdentity *armmsi.Identity) *runtime.Poller[containerservice.ManagedClustersClientCreateOrUpdateResponse] {
	_ = "STUB: not implemented"
	return nil
}

// Update the kubelet identity in the identity profile

// Update the cluster and wait for the operation to complete

// ExpectGrantedACRAccess grants the specified identity access to pull from the ACR
func (env *Environment) ExpectGrantedACRAccess(ctx context.Context, identity *armmsi.Identity) {
	_ = "STUB: not implemented"
	return
}

// Get the ACR resource ID

// AcrPull role definition ID: 7f951dda-4ed3-4680-a7ca-43fe172d538d

// CheckClusterIdentityType returns the type of managed identity used by the cluster
func (env *Environment) CheckClusterIdentityType(ctx context.Context) string {
	_ = "STUB: not implemented"
	return ""
}

// IsClusterUserAssignedIdentity checks if the cluster uses user-assigned managed identity
func (env *Environment) IsClusterUserAssignedIdentity(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// GetKubeletIdentity returns the current kubelet identity
func (env *Environment) GetKubeletIdentity(ctx context.Context) *containerservice.UserAssignedIdentity {
	_ = "STUB: not implemented"
	return nil
}

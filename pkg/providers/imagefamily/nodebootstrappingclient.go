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

package imagefamily

import (
	"context"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	types "github.com/Azure/karpenter-provider-azure/pkg/providers/imagefamily/types"
	"github.com/Azure/karpenter-provider-azure/pkg/provisionclients/models"
)

// As of this time, for managed identity, token caching exists in the implementation beneath MSAL GetToken(...) as below. This means our caching is not needed if MSAL GetToken(...) caching is valid.
// https://github.com/AzureAD/microsoft-authentication-library-for-go/blob/b4b8bfc9569042572ccb82b648ea509075fadb74/apps/managedidentity/managedidentity.go#L318
// However, this is never made clear in the interface of this layer nor its documentation, thus relying on that assumption may not be perfect. Still, the likelihood of change is low given history.
// Alternatively, we could try to pair our cache implementation with what's in Azure clients as below. However, the cost makes it arguably not worth it, given the current circumstances.
// https://github.com/Azure/azure-sdk-for-go/blob/f72e2ad4f23b02eba6387dc31580c0e66333f2ae/sdk/internal/temporal/resource.go#L78-L140
type tokenProvider struct {
	// credential.GetToken(...) doesn't have locks, thus could have resulted in multiple token requests at the same time.
	// This could produce unnecessarily more traffic with the token provider service, and more tokens to be issued if there is no server-side caching.
	// MSAL-side issue: https://github.com/AzureAD/microsoft-authentication-library-for-go/issues/569
	mu sync.Mutex

	cloud cloud.Configuration
}

func (t *tokenProvider) getToken(ctx context.Context, credential azcore.TokenCredential) (azcore.AccessToken, error) {
	_ = "STUB: not implemented"
	return *new(azcore.AccessToken), nil
}

// NodeBootstrappingClient implements the NodeBootstrappingAPI interface using the swagger-generated client.
type NodeBootstrappingClient struct {
	serverURL         string
	subscriptionID    string
	resourceGroupName string
	resourceName      string
	credential        azcore.TokenCredential
	tokenProvider     *tokenProvider
	enableLogging     bool
}

// NewNodeBootstrappingClient creates a new NodeBootstrappingClient with token caching enabled.
func NewNodeBootstrappingClient(
	ctx context.Context,
	cloud cloud.Configuration,
	subscriptionID string,
	resourceGroupName string,
	resourceName string,
	credential azcore.TokenCredential,
	serverURL string,
	enableLogging bool,
) (*NodeBootstrappingClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get implements the NodeBootstrappingAPI interface.
// It retrieves node bootstrapping data (CSE and base64-encoded CustomData), but may omit the TLS bootstrap token.
func (c *NodeBootstrappingClient) Get(
	ctx context.Context,
	parameters *models.ProvisionValues,
) (types.NodeBootstrapping, error) {
	_ = "STUB: not implemented"
	return *new(types.NodeBootstrapping), nil
}

// Add Authorization Bearer token header

// Middleware logging only if ENABLE_AZURE_SDK_LOGGING flag is enabled

// Create the client

// Prepare the parameters for the request

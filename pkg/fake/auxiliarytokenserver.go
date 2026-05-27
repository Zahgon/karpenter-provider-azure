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
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/karpenter-provider-azure/pkg/auth"
)

type AuxiliaryTokenDoInput struct {
	request *http.Request
}

type AuxiliaryTokenBehavior struct {
	AuxiliaryTokenDoBehavior MockedFunction[AuxiliaryTokenDoInput, *http.Response]
}

// assert that the fake implements the interface
var _ auth.AuxiliaryTokenServer = &AuxiliaryTokenServer{}

type AuxiliaryTokenServer struct {
	AuxiliaryTokenBehavior
	Token azcore.AccessToken
}

// NewAuxiliaryTokenServer creates a new AuxiliaryTokenServer with the given token.
func NewAuxiliaryTokenServer(token string, expiresOn time.Time, refreshOn time.Time) *AuxiliaryTokenServer {
	_ = "STUB: not implemented"
	return nil
}

func (c *AuxiliaryTokenServer) SetToken(token string, expiresOn time.Time, refreshOn time.Time) {
	_ = "STUB: not implemented"
	return
}

// Reset must be called between tests otherwise tests will pollute each other.
func (c *AuxiliaryTokenServer) Reset() { _ = "STUB: not implemented"; return }

func (c *AuxiliaryTokenServer) Do(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// init response writer

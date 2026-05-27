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

package auth

import (
	"net/http"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
)

type AuxiliaryTokenServer interface {
	Do(req *http.Request) (*http.Response, error)
}

var _ policy.Policy = &AuxiliaryTokenPolicy{}

// AuxiliaryTokenPolicy provides a custom policy used to authenticate
// with shared node image galleries.
type AuxiliaryTokenPolicy struct {
	Token  azcore.AccessToken
	url    string
	scope  string
	client AuxiliaryTokenServer
	lock   sync.Mutex
}

func (p *AuxiliaryTokenPolicy) GetAuxiliaryToken() error { _ = "STUB: not implemented"; return nil }

// If the token is uninitialized or close to expiration, fetch a new one

func (p *AuxiliaryTokenPolicy) Do(req *policy.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewAuxiliaryTokenPolicy(client AuxiliaryTokenServer, url string, scope string) *AuxiliaryTokenPolicy {
	_ = "STUB: not implemented"
	return nil
}

func getAuxiliaryToken(client AuxiliaryTokenServer, url string, scope string) (azcore.AccessToken, error) {
	_ = "STUB: not implemented"
	return *new(azcore.AccessToken), nil
}

// Construct the request

// Send the request

// Decode the response body into the AccessToken struct

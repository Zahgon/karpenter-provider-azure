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

package client

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
)

const (
	apiVersion = "2023-01-01-preview"
	pricingURL = "https://prices.azure.com/api/retail/prices?api-version=" + apiVersion // TODO: This API is not available in national clouds.
)

type PricingAPI interface {
	GetProductsPricePages(context.Context, []*Filter, func(output *ProductsPricePage)) error
}

type pricingAPI struct {
	cloud cloud.Configuration
}

func New(cloud cloud.Configuration) PricingAPI { _ = "STUB: not implemented"; return *new(PricingAPI) }

func (papi *pricingAPI) GetProductsPricePages(_ context.Context, filters []*Filter, pageHandler func(output *ProductsPricePage)) error {
	_ = "STUB: not implemented"
	return nil
}

// If the cloud is not Azure Public, the pricing API isn't supported and we return an error

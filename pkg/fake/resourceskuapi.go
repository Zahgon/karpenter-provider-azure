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

	//nolint:staticcheck // deprecated package
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2022-08-01/compute"
	"github.com/Azure/skewer"
)

// TODO: consider using fakes from skewer itself

// ResourceSkus is a map of location to resource skus
var ResourceSkus = make(map[string][]compute.ResourceSku)

// assert that the fake implements the interface
var _ skewer.ResourceClient = &ResourceSKUsAPI{}

type ResourceSKUsAPI struct {
	Location string
	// skewer.ResourceClient
	Error error
}

// Reset must be called between tests otherwise tests will pollute each other.
func (s *ResourceSKUsAPI) Reset() {
	_ = "STUB: not implemented"
	// c.ResourceSKUsBehavior.Reset()
	return
}

func (s *ResourceSKUsAPI) ListComplete(_ context.Context, _, _ string) (compute.ResourceSkusResultIterator, error) {
	_ = "STUB: not implemented"
	return *new(compute.ResourceSkusResultIterator), nil
}

// cur

// fn

// end of iterator

// MakeSKU looks up a full *skewer.SKU from the fake ResourceSkus data for the default Region.
// This includes Name, Family, Capabilities (vCPU count, etc.), and other SKU metadata.
// Panics if the SKU is not found in the fake data.
func MakeSKU(skuName string) *skewer.SKU { _ = "STUB: not implemented"; return nil }

// MakeSKUForRegion looks up a full *skewer.SKU from the fake ResourceSkus data for the given region.
// This includes Name, Family, Capabilities (vCPU count, etc.), and other SKU metadata.
// Panics if the SKU is not found in the fake data.
func MakeSKUForRegion(skuName, region string) *skewer.SKU { _ = "STUB: not implemented"; return nil }

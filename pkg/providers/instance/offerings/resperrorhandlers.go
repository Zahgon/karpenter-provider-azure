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

package offerings

import (
	"context"

	"github.com/Azure/karpenter-provider-azure/pkg/cache"
	"github.com/Azure/skewer"
	corecloudprovider "sigs.k8s.io/karpenter/pkg/cloudprovider"
)

type responseErrorHandlerEntry struct {
	match  func(error) bool
	handle errorHandle
}

type ResponseErrorHandler struct {
	UnavailableOfferings *cache.UnavailableOfferings
	HandlerEntries       []responseErrorHandlerEntry
}

// Comparing to ErrorDetailHandler, this is handling same errors, but for a different error data model.
// HandlerEntries should generally be kept in sync.
// See ErrorDetailHandler for more details.
func NewResponseErrorHandler(unavailableOfferings *cache.UnavailableOfferings) *ResponseErrorHandler {
	_ = "STUB: not implemented"
	return nil
}

func (h *ResponseErrorHandler) extractErrorCodeAndMessage(err error) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

func (h *ResponseErrorHandler) Handle(ctx context.Context, sku *skewer.SKU, instanceType *corecloudprovider.InstanceType, zone, capacityType string, responseError error) error {
	_ = "STUB: not implemented"
	return nil
}

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

// HandlableError is a code + message pair extracted from an API error response.
// It is intentionally minimal and API-agnostic. This module decides how to interpret it.
type HandlableError struct {
	Code    string
	Message string
}

func (e *HandlableError) Error() string { _ = "STUB: not implemented"; return "" }

func ErrorToHandlableError(err error) *HandlableError { _ = "STUB: not implemented"; return nil }

// Note: this is not truly extracting the message, but rather dump the whole error in.
// This is okay for now, as all handling logic just does substring match on the message, or care only about ErrorCode (ideal in long run).
// TODO: rework this? Once we revisit this whole module to perhaps share the handling logic to azure-sdk-for-go-extensions.

type aksMachineBeginCreateErrorHandlerEntry struct {
	match  func(*HandlableError) bool
	handle errorHandle
}

// AKSMachineBeginCreateErrorHandler classifies HandlableErrors and takes appropriate offerings cache actions.
type AKSMachineBeginCreateErrorHandler struct {
	unavailableOfferings *cache.UnavailableOfferings
	handlerEntries       []aksMachineBeginCreateErrorHandlerEntry
}

// NewAKSMachineBeginCreateErrorHandler creates a handler for AKS Machine API sync-phase errors.
// TODO: consider sharing this on azure-sdk-for-go-extensions like other error handlers.
func NewAKSMachineBeginCreateErrorHandler(unavailableOfferings *cache.UnavailableOfferings) *AKSMachineBeginCreateErrorHandler {
	_ = "STUB: not implemented"
	return nil
}

func (h *AKSMachineBeginCreateErrorHandler) Handle(ctx context.Context, sku *skewer.SKU, instanceType *corecloudprovider.InstanceType, zone, capacityType string, he *HandlableError) error {
	_ = "STUB: not implemented"
	return nil
}

// handleSKUNotAvailableForSubscriptionError marks both capacity types unavailable because
// subscription-level VM size restrictions apply equally to spot and on-demand.
// This differs from CRP's SkuNotAvailable (handleSKUNotAvailableError) which may be
// spot-only due to capacity — here the error is a hard subscription restriction.
func handleSKUNotAvailableForSubscriptionError(
	ctx context.Context,
	unavailableOfferings *cache.UnavailableOfferings,
	sku *skewer.SKU,
	instanceType *corecloudprovider.InstanceType,
	zone,
	capacityType,
	errorCode,
	errorMessage string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// For "Virtual Machine size: '%s' is not supported for subscription %s in location '%[3]s'. %s. Please refer to aka.ms/aks/vm-size-selector to find supported VM sizes in location '%[3]s'."
// ASSUMPTION: this error occurring means the whole VM family is not available.
func isSKUNotAvailableForSubscription(he *HandlableError) bool {
	_ = "STUB: not implemented"
	return false
}

// For "Virtual Machine size: '%s' is not supported for subscription %s in location '%[3]s'. %s. Please refer to aka.ms/aks/vm-size-selector to find supported VM sizes in location '%[3]s'."
// Similar to IsSKUNotAvailableForSubscription, but this different error code is another possible variant.
// ASSUMPTION: this error occurring means the whole VM family is not available.
func isSKUNotAvailableForSubscriptionBadRequest(he *HandlableError) bool {
	_ = "STUB: not implemented"
	return false
}

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

package status

import (
	"context"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/azclient/azapi"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	DiskEncryptionSetRBACMissing = "DiskEncryptionSetRBACMissing"
	// TODO: May want to rethink how we handle successful validation + potential for RBAC removal.
	// See this PR comment for considerations:
	// https://github.com/Azure/karpenter-provider-azure/pull/1372#discussion_r2795367386
	// ValidationSuccessRequeueInterval defines how often to re-validate DES RBAC after success
	// Set to 1 hour since RBAC changes are infrequent in production
	ValidationSuccessRequeueInterval = 1 * time.Hour
	// ValidationFailureRequeueInterval defines how often to retry DES RBAC validation after auth failure
	// Set to 1 minute to detect when permissions are granted without creating a high system load
	ValidationFailureRequeueInterval = 1 * time.Minute
	// DiskEncryptionSetRBACErrorMessage is the error message shown when the controlling identity lacks Reader permissions
	DiskEncryptionSetRBACErrorMessage = "controlling identity does not have Reader role on Disk Encryption Set"
)

type ValidationReconciler struct {
	diskEncryptionSetsAPI     azapi.DiskEncryptionSetsAPI
	parsedDiskEncryptionSetID *arm.ResourceID // parsed by options.Validate(), will be nil if DiskEncryptionSetID is not set
}

func NewValidationReconciler(
	diskEncryptionSetsAPI azapi.DiskEncryptionSetsAPI,
	parsedDiskEncryptionSetID *arm.ResourceID,
) *ValidationReconciler {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidationReconciler) Reconcile(ctx context.Context, nodeClass *v1beta1.AKSNodeClass) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Check BYOK RBAC if DES ID is configured

// Auth failure (403/401) - set condition to False, requeue soon to detect permission grants

// Unexpected error (network, parsing, etc.) - don't change condition, return error for retry

// All validations passed - requeue to detect permission revocations

func (r *ValidationReconciler) validateDiskEncryptionSetRBAC(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Attempt to read the DiskEncryptionSet
	// This uses the controller's current credentials (DefaultAzureCredential)
	return nil
}

// Wrap the original error to preserve the error chain for isAuthorizationErr checks

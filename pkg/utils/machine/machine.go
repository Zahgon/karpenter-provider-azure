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

package machine

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
)

// HandleProvisioningState inspects the provisioning state of an AKS machine.
// It returns the provisioning error if the machine failed, a polling error if the provisioning
// state indicates a canceled or otherwise fatal state, and a boolean indicating whether the
// provisioning is complete (either succeeded, deleting or failed).
func HandleProvisioningState(ctx context.Context, aksMachine *armcontainerservice.Machine) (*armcontainerservice.ErrorDetail, error, bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

// IsAKSMachineOrMachinesPoolNotFound returns true if the error
// returned from the Azure API indicates that the AKS machine or the AKS machines pool
// could not be found.
func IsAKSMachineOrMachinesPoolNotFound(err error) bool { _ = "STUB: not implemented"; return false }

// Covers AKS machines pool not found on PUT machine, GET machine, GET (list) machines, POST agent pool (DELETE machines), and AKS machine not found on GET machine
// Covers AKS machine not found on POST agent pool (DELETE machines)

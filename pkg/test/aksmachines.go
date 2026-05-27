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

package test

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
)

// AKSMachineOptions customizes an AKS Machine for testing.
type AKSMachineOptions struct {
	Name                 string
	MachinesPoolName     string
	ClusterResourceGroup string
	ClusterName          string
	Location             string
	VMSize               string
	Priority             *armcontainerservice.ScaleSetPriority
	Zones                []*string
	Properties           *armcontainerservice.MachineProperties
	NodepoolName         string
}

// AKSMachine creates a test AKS Machine with defaults that can be overridden by AKSMachineOptions.
// This implementation matches the setDefaultMachineValues pattern from the fake API.
// Overrides are applied in order, with last-write-wins semantics.
//
//nolint:gocyclo
func AKSMachine(overrides ...AKSMachineOptions) *armcontainerservice.Machine {
	_ = "STUB: not implemented"
	return nil
}

// Provide default values if none are set

// Set default properties if not provided - matching setDefaultMachineValues pattern

// Set Priority field (required field that was missing) - must be set AFTER default Priority is established

// Set Status with creation timestamp (required field) - matching setDefaultMachineValues

// Set ResourceID (required field) - simulates VM resource ID following AKS naming convention
// vmName = aks-<machinesPoolName>-<aksMachineName>-########-vm

// Generate a VM name following AKS convention: aks-{agentPoolName}-{machineName}-{randomId}-vm{id}

// Set NodeImageVersion - matching setDefaultMachineValues default

// Default node image version if none provided

// Construct the AKS Machine

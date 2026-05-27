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

package instance

import (
	"github.com/Azure/azure-kusto-go/kusto/kql"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

const (
	vmResourceType  = "microsoft.compute/virtualmachines"
	nicResourceType = "microsoft.network/networkinterfaces"
)

// getResourceListQueryBuilder returns a KQL query builder for listing resources with nodepool tags
// but excluding AKS machine-created resources
func getResourceListQueryBuilder(rg string, resourceType string) *kql.Builder {
	_ = "STUB: not implemented"
	return nil
}

// ARG resources appear to have lowercase RG

// GetVMListQueryBuilder returns a KQL query builder for listing VMs with nodepool tags
func GetVMListQueryBuilder(rg string) *kql.Builder { _ = "STUB: not implemented"; return nil }

// GetNICListQueryBuilder returns a KQL query builder for listing NICs with nodepool tags
func GetNICListQueryBuilder(rg string) *kql.Builder { _ = "STUB: not implemented"; return nil }

// createVMFromQueryResponseData converts ARG query response data into a VirtualMachine object
func createVMFromQueryResponseData(data map[string]interface{}) (*armcompute.VirtualMachine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We see inconsistent casing being returned by ARG for the last segment
// of the vm.ID string. This forces it to be lowercase.

// createNICFromQueryResponseData converts ARG query response data into a Network Interface object
func createNICFromQueryResponseData(data map[string]interface{}) (*armnetwork.Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We see inconsistent casing being returned by ARG for the last segment
// of the nic.ID string. This forces it to be lowercase.

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

package utils

// this parsing function replaces three different functions in different packages that all had bugs. Please don't use a regex to parse these
type VnetSubnetResource struct {
	SubscriptionID    string
	ResourceGroupName string
	VNetName          string
	SubnetName        string
}

func (v VnetSubnetResource) IsSameVNET(cmp VnetSubnetResource) bool {
	_ = "STUB: not implemented"
	return false
}

// GetSubnetResourceID constructs the subnet resource id
func GetSubnetResourceID(subscriptionID, resourceGroupName, virtualNetworkName, subnetName string) string {
	_ = "STUB: not implemented"
	// an example subnet resource: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
	return ""
}

// GetVnetSubnetIDComponents parses an Azure subnet resource ID into its component parts.
// Input: A fully qualified Azure subnet resource ID in the format:
//
//	/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
//
// The input is case-insensitive and must contain exactly 11 slash-separated segments.
// Output: A vnetSubnetResource struct containing:
//   - SubscriptionID: The Azure subscription ID
//   - ResourceGroupName: The resource group name
//   - VNetName: The virtual network name
//   - SubnetName: The subnet name
//
// Returns an error if the input format is invalid or doesn't match the expected structure.
func GetVnetSubnetIDComponents(vnetSubnetID string) (VnetSubnetResource, error) {
	_ = "STUB: not implemented"
	return *new(VnetSubnetResource), nil
}

//this is a cheap way of ensure all the names match

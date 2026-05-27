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

package azure

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

func (env *Environment) GetVM(nodeName string) armcompute.VirtualMachine {
	_ = "STUB: not implemented"
	return *new(armcompute.VirtualMachine)
}

func (env *Environment) GetVMSKU(nodeName string) string { _ = "STUB: not implemented"; return "" }

func (env *Environment) GetVMByName(vmName string) armcompute.VirtualMachine {
	_ = "STUB: not implemented"
	return *new(armcompute.VirtualMachine)
}

func (env *Environment) SimulateVMEviction(nodeName string) { _ = "STUB: not implemented"; return }

func (env *Environment) GetNetworkInterface(nicName string) armnetwork.Interface {
	_ = "STUB: not implemented"
	return *new(armnetwork.Interface)
}

func (env *Environment) GetClusterVNET() *armnetwork.VirtualNetwork {
	_ = "STUB: not implemented"
	return nil
}

func (env *Environment) GetClusterSubnet() *armnetwork.Subnet {
	_ = "STUB: not implemented"
	return nil
}

// This returns the first vnet we find in the resource group, works for managed vnet, it hasn't been tested on custom vnet.
func firstVNETInRG(ctx context.Context, client *armnetwork.VirtualNetworksClient, vnetRG string) (*armnetwork.VirtualNetwork, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (env *Environment) K8sVersion() string { _ = "STUB: not implemented"; return "" }

func (env *Environment) K8sVersionWithOffset(offset int) string {
	_ = "STUB: not implemented"
	return ""
}

// Choose a minor version one lesser than the server's minor version. This ensures that we choose an AMI for
// this test that wouldn't be selected as Karpenter's SSM default (therefore avoiding false positives), and also
// ensures that we aren't violating version skew.

func (env *Environment) K8sMinorVersion() int { _ = "STUB: not implemented"; return 0 }

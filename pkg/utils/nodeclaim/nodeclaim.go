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

package nodeclaim

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	karpv1 "sigs.k8s.io/karpenter/pkg/apis/v1"

	armcompute "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/instance"
)

// UsingAKSNodeClassPredicate creates a predicate to filter node claim using AKS node class.
func UsingAKSNodeClassPredicate() predicate.Funcs {
	_ = "STUB: not implemented"
	return *new(predicate.Funcs)
}

// GetAKSNodeClass resolves the AKSNodeClass from the NodeClaim's NodeClassRef.
// If the NodeClass for the nodeClaim has DeletionTimestamp set, an error is returned.
func GetAKSNodeClass(ctx context.Context, kubeClient client.Client, nodeClaim *karpv1.NodeClaim) (*v1beta1.AKSNodeClass, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: Could go onto vmInstanceProvider?
// GetVM gets the Azure VM associated with the NodeClaim
func GetVM(ctx context.Context, vmInstanceProvider instance.VMProvider, nodeClaim *karpv1.NodeClaim) (*armcompute.VirtualMachine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetVMName parses the provider ID stored on the node to get the vmName
// associated with a node
func GetVMName(providerID string) (string, error) {
	_ = "STUB: not implemented"
	// standalone VMs have providerID in the format: azure:///subscriptions/<subscriptionID>/resourceGroups/<resourceGroup>/providers/Microsoft.Compute/virtualMachines/<instanceID>
	return "", nil
}

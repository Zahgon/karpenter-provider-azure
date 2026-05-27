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

package garbagecollection

import (
	"context"
	"time"

	"github.com/awslabs/operatorpkg/reconciler"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/Azure/karpenter-provider-azure/pkg/providers/instance"
)

const (
	NicReservationDuration = time.Second * 180
	// We set this interval at 5 minutes, as thats how often our NRP limits are reset.
	// See: https://learn.microsoft.com/en-us/azure/azure-resource-manager/management/request-limits-and-throttling#network-throttling
	NicGarbageCollectionInterval = time.Minute * 5
)

type NetworkInterface struct {
	kubeClient         client.Client
	vmInstanceProvider instance.VMProvider
}

func NewNetworkInterface(kubeClient client.Client, vmInstanceProvider instance.VMProvider) *NetworkInterface {
	_ = "STUB: not implemented"
	return nil
}

func (c *NetworkInterface) populateUnremovableInterfaces(ctx context.Context) (sets.Set[string], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *NetworkInterface) Reconcile(ctx context.Context) (reconciler.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconciler.Result), nil
}

func (c *NetworkInterface) Register(_ context.Context, m manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

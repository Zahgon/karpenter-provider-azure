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

package common

import (
	"context"
	"testing"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"

	. "sigs.k8s.io/karpenter/pkg/utils/testing" //nolint:staticcheck

	"sigs.k8s.io/controller-runtime/pkg/client"

	karpv1 "sigs.k8s.io/karpenter/pkg/apis/v1"
)

type ContextKey string

const (
	GitRefContextKey = ContextKey("gitRef")

	NetworkDataplaneCilium = "cilium"
	NetworkDataplaneAzure  = "azure"
)

type Environment struct {
	context.Context
	cancel context.CancelFunc

	Client     client.Client
	Config     *rest.Config
	KubeClient kubernetes.Interface
	Monitor    *Monitor

	// Resolved from cluster
	InClusterController bool
	NetworkDataplane    string

	StartingNodeCount int
}

func NewEnvironment(t *testing.T) *Environment { _ = "STUB: not implemented"; return nil }

func (env *Environment) Stop() { _ = "STUB: not implemented"; return }

func NewConfig() *rest.Config { _ = "STUB: not implemented"; return nil }

func NewClient(ctx context.Context, config *rest.Config) client.Client {
	_ = "STUB: not implemented"
	return *new(client.Client)
}

func (env *Environment) DefaultNodePool(nodeClass *v1beta1.AKSNodeClass) *karpv1.NodePool {
	_ = "STUB: not implemented"
	return nil
}

// TODO: do we need that much?

// AdaptToClusterConfig modifies NodePool to match the cluster configuration.
// It has to be applied to any custom node pools constructed by tests;
// is already applied by default test NodePool constructors.
func (env *Environment) AdaptToClusterConfig(nodePool *karpv1.NodePool) *karpv1.NodePool {
	_ = "STUB: not implemented"
	return nil
}

// https://karpenter.sh/docs/concepts/nodepools/#cilium-startup-taint

// required for Karpenter to predict overhead from cilium DaemonSet

func (env *Environment) ArmNodepool(nodeClass *v1beta1.AKSNodeClass) *karpv1.NodePool {
	_ = "STUB: not implemented"
	return nil
}

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
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1"
	schedulingv1 "k8s.io/api/scheduling/v1"
	storagev1 "k8s.io/api/storage/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	karpv1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/apis/v1alpha1"

	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const TestingFinalizer = "testing/finalizer"

var (
	CleanableObjects = []client.Object{
		&corev1.Pod{},
		&appsv1.Deployment{},
		&appsv1.StatefulSet{},
		&appsv1.DaemonSet{},
		&policyv1.PodDisruptionBudget{},
		&corev1.PersistentVolumeClaim{},
		&corev1.PersistentVolume{},
		&storagev1.StorageClass{},
		&karpv1.NodePool{},
		&corev1.LimitRange{},
		&schedulingv1.PriorityClass{},
		&corev1.Node{},
		&karpv1.NodeClaim{},
		&v1beta1.AKSNodeClass{},
		&v1alpha1.NodeOverlay{},
	}
)

func (env *Environment) BeforeEach() { _ = "STUB: not implemented"; return }

// Expect this cluster to be clean for test runs to execute successfully

func (env *Environment) ExpectCleanCluster() { _ = "STUB: not implemented"; return }

// This assumes the cluster is created without managed pools

func (env *Environment) ExpectSystemNodesTainted() { _ = "STUB: not implemented"; return }

func (env *Environment) ExpectNoPodsInDefaultNamespace() { _ = "STUB: not implemented"; return }

func (env *Environment) EventuallyExpectNoProvisionablePods() { _ = "STUB: not implemented"; return }

func (env *Environment) ExpectNoNodePoolOrAKSNodeClass() { _ = "STUB: not implemented"; return }

func (env *Environment) Cleanup() { _ = "STUB: not implemented"; return }

func (env *Environment) AfterEach() { _ = "STUB: not implemented"; return }

func (env *Environment) CleanupObjects(cleanableObjects ...client.Object) {
	_ = "STUB: not implemented"
	return
	// wait one second to let the caches get up-to-date for deletion
}

// This only gets the metadata for the objects since we don't need all the details of the objects

// Limit the concurrency of these calls to 50 workers per object so that we try to limit how aggressively we
// are deleting so that we avoid getting client-side throttled

// If the deletes eventually succeed, we should have no elements here at the end of the test

func (env *Environment) ExpectTestingFinalizerRemoved(obj client.Object) error {
	_ = "STUB: not implemented"
	return nil
}

// If the Group is the "core" APIs, then we can strategic merge patch
// CRDs do not currently have support for strategic merge patching, so we can't blindly do it
// https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/#advanced-features-and-flexibility:~:text=Yes-,strategic%2Dmerge%2Dpatch,-The%20new%20endpoints

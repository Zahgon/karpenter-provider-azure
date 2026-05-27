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

package labels

import (
	"context"

	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/karpenter/pkg/scheduling"
)

// These labels are defined here rather than v1beta1 because we do not support scheduling simulation
// on these labels
var (
	AKSLabelEBPFDataplane               = v1beta1.AKSLabelDomain + "/ebpf-dataplane"
	AKSLabelAzureCNIOverlay             = v1beta1.AKSLabelDomain + "/azure-cni-overlay"
	AKSLabelSubnetName                  = v1beta1.AKSLabelDomain + "/network-subnet"
	AKSLabelVNetGUID                    = v1beta1.AKSLabelDomain + "/nodenetwork-vnetguid"
	AKSLabelPodNetworkType              = v1beta1.AKSLabelDomain + "/podnetwork-type"
	AKSLabelNetworkStatelessCNI         = v1beta1.AKSLabelDomain + "/network-stateless-cni"
	AKSLocalDNSStateLabelKey            = v1beta1.AKSLabelDomain + "/localdns-state"
	AKSArtifactStreamingEnabledLabelKey = v1beta1.AKSLabelDomain + "/artifactstreaming-enabled"

	AKSLabelRole = v1beta1.AKSLabelDomain + "/role"

	kubeletLabelNamespaces = sets.NewString(
		v1.LabelNamespaceSuffixKubelet,
		v1.LabelNamespaceSuffixNode,
	)

	kubeletLabels = sets.NewString(
		v1.LabelHostname,
		v1.LabelTopologyZone,
		v1.LabelTopologyRegion,
		v1.LabelFailureDomainBetaZone,
		v1.LabelFailureDomainBetaRegion,
		v1.LabelInstanceType,
		v1.LabelInstanceTypeStable,
		v1.LabelOSStable,
		v1.LabelArchStable,

		LabelOS,
		LabelArch,
	)

	K8sLabelDomains = sets.New(
		"kubernetes.io",
		"k8s.io",
	)
)

// These label definitions taken from here: https://github.com/kubernetes/kubernetes/blob/e319c541f144e9bee6160f1dd8671638a9029f4c/staging/src/k8s.io/kubelet/pkg/apis/well_known_labels.go#L67
const (
	// LabelOS is a label to indicate the operating system of the node.
	// The OS labels are promoted to GA in 1.14. kubelet applies GA labels and stop applying the beta OS labels in Kubernetes 1.19.
	LabelOS = "beta.kubernetes.io/os"
	// LabelArch is a label to indicate the architecture of the node.
	// The Arch labels are promoted to GA in 1.14. kubelet applies GA labels and stop applying the beta Arch labels in Kubernetes 1.19.
	LabelArch = "beta.kubernetes.io/arch"
)

func Get(
	ctx context.Context,
	nodeClass *v1beta1.AKSNodeClass,
	arch string,
) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add labels that are always there

// Note that while we set the Kubelet identity label here, in bootstrap API mode, the actual kubelet identity that is set in the bootstrapping
// script is configured by the NPS service. That means the label can be set to the older client ID if the client ID
// changed recently. This is OK because drift will correct it.

// Prevent race conditions with startup taints by telling Karpenter not to sync taints from the NodeClaim to the Node
// See https://github.com/kubernetes-sigs/karpenter/issues/1772

// Add os-sku label based on imageFamily

// Add os-sku-requested label that exactly matches the imageFamily specified on the NodeClass

// nil static parameters here is safe only because we're not using the resulting imageFamily for anything except to get its name

// TODO: make conditional on pod subnet
// good

// Sanity Check: in production we should always have a k8s version set

// This label is required for the cilium agent daemonset because
// we select the nodes for the daemonset based on this label
//              - key: kubernetes.azure.com/ebpf-dataplane
//            operator: In
//            values:
//              - cilium

// Only set the artifact streaming label when it's enabled (matching AKS RP behavior)
// ARM64 nodes do not support artifact streaming

// CanKubeletSetLabel returns true if the given label is a label kubelet is allowed to set.
// This is similar to the method one used by the node restriction admission
// https://github.com/kubernetes/kubernetes/blob/e319c541f144e9bee6160f1dd8671638a9029f4c/staging/src/k8s.io/kubelet/pkg/apis/well_known_labels.go#L67
// with the isKubernetesLabel check from https://github.com/kubernetes/kubernetes/blob/4bed36e03e7bd699b089d33da6f7d7c9ef9eb661/cmd/kubelet/app/options/options.go#L176C6-L176C23.
func CanKubeletSetLabel(key string) bool { _ = "STUB: not implemented"; return false }

func IsLabelKubeletManaged(key string) bool { _ = "STUB: not implemented"; return false }

// GetWellKnownSingleValuedRequirementLabels converts well-known Azure single-value instanceType.Requirements to labels
// This is useful for projecting requirements from the NodeClaim to labels, which is required for scheduling simulation to work
// correctly.
func GetWellKnownSingleValuedRequirementLabels(requirements scheduling.Requirements) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// GetAllSingleValuedRequirementLabels converts single-value instanceType.Requirements to labels
// Like   instanceType.Requirements.Labels() it uses single-valued requirements
// Unlike instanceType.Requirements.Labels() it does not filter out restricted Node labels
func GetAllSingleValuedRequirementLabels(requirements scheduling.Requirements) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func GetFilteredSingleValuedRequirementLabels(requirements scheduling.Requirements, predicate func(k string, r *scheduling.Requirement) bool) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// isKubernetesNamespace checks if the given namespace belongs to kubernetes.io or k8s.io domains
// Similar to https://github.com/kubernetes/kubernetes/blob/4bed36e03e7bd699b089d33da6f7d7c9ef9eb661/cmd/kubelet/app/options/options.go#L176C6-L176C23
func isKubernetesNamespace(namespace string) bool { _ = "STUB: not implemented"; return false }

func getLabelNamespace(key string) string { _ = "STUB: not implemented"; return "" }

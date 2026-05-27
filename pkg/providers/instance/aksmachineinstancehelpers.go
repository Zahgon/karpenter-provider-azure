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
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
	karpv1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	corecloudprovider "sigs.k8s.io/karpenter/pkg/cloudprovider"

	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
	"github.com/Azure/karpenter-provider-azure/pkg/operator/options"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/instancetype"
)

// buildAKSMachineTemplate creates an in-memory AKS machine template from the provided specs.
// May return error whenever required fields are not set (check carefully).
//
// BATCH AWARENESS: This template is used both for single-VM creates and as the shared
// template body in batch creates. When adding new MachineProperties fields:
//   - If the field varies per NodeClaim (like Tags), also update clearPerMachineFields in
//     pkg/providers/azclient/aksmachinesheaderbatch/batch_field_registry.go so the batch
//     system knows to move it to the per-machine header.
//   - If the field is the same for all NodeClaims in a NodePool+NodeClass (like VMSize),
//     no action needed — it's automatically part of the shared template and batch grouping hash.
func (p *DefaultAKSMachineProvider) buildAKSMachineTemplate(ctx context.Context, instanceType *corecloudprovider.InstanceType, capacityType string, placementScope string, zone string, nodeClass *v1beta1.AKSNodeClass, nodeClaim *karpv1.NodeClaim) (*armcontainerservice.Machine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NodeImageVersion
// E.g., "AKSUbuntu-2204gen2containerd-2023.11.15"

// GPUProfile

// OrchestratorVersion (i.e., Kubernetes version)

// OSSKU, EnableFIPS

// OSDiskType

// NodeTaints, NodeInitializationTaints

// NodeLabels, Mode

// Priority (e.g., regular, spot)

// Tags (to be put on AKS machine and all affiliated resources)
// Note: as of the time of writing, AKS machine API does not support tags on NICs. This could be fixed server-side.
//
// BATCH: Tags is a per-machine field (varies per NodeClaim). If you change this,
// see batch_field_registry.go — ClearPerMachineFields must cover any new per-machine
// MachineProperties field so batch grouping and header extraction stay correct.

// BATCH: Zones is a per-machine field (selected from instance type offerings).
// It lives on Machine (not MachineProperties) so it's excluded from template
// hashing by design. See batch_field_registry.go for the full field classification.

// AKS machine API take control, if nil
// As of the time of writing, the current version of AKS machine API support just that with nil. That is unlikely to change.
// PodSubnetID:          "",
// EnableNodePublicIP:   nil,
// NodePublicIPPrefixID: "",
// IPTags:               nil,

// GPUInstanceProfile: nil,

// AKS machine API defaults it if nil

// WindowsProfile: nil,

// KubeletDiskType:          "",

// AKS machine API defaults it per network plugins if nil.
// WorkloadRuntime:          nil,

// EnableVTPM:             nil,
// EnableSecureBoot:       nil,

func configureGPUProfile(instanceType *corecloudprovider.InstanceType, nodeClass *v1beta1.AKSNodeClass) *armcontainerservice.GPUProfile {
	_ = "STUB: not implemented"
	// Non-GPU SKUs don't need a GPU profile.
	return nil
}

// GPU SKUs: pass through the driver setting from nodeClass.
// "Driver" mode -> Install, "None" mode -> None (treat as non-GPU).
// Upstream instance type filtering already ensures invalid SKU+mode combinations
// (e.g., AMD GPU with Driver mode) are excluded before reaching here.

func configureArtifactStreamingProfile(nodeClass *v1beta1.AKSNodeClass, instanceType *corecloudprovider.InstanceType) *armcontainerservice.AgentPoolArtifactStreamingProfile {
	_ = "STUB: not implemented"
	return nil
}

func configureLocalDNSProfile(nodeClass *v1beta1.AKSNodeClass) *armcontainerservice.LocalDNSProfile {
	_ = "STUB: not implemented"
	// Use the wire-resolved spec so Karpenter's Preferred-mode decision
	// (recorded on the AKSNodeClass via Status.LocalDNSState) is honored
	// end-to-end downstream. This guarantees a deterministic and consistent
	// LocalDNS decision across Machines spawned from the same NodeClass:
	// Preferred is never sent downstream, so it can never be re-interpreted.
	// See AKSNodeClass.ResolvedLocalDNSForWire for the full rationale.
	return nil
}

func convertLocalDNSOverrides(overrides []v1beta1.LocalDNSZoneOverride) map[string]*armcontainerservice.LocalDNSOverride {
	_ = "STUB: not implemented"
	return nil
}

func configureOSDiskType(ctx context.Context, instanceTypeProvider instancetype.Provider, nodeClass *v1beta1.AKSNodeClass, instanceType *corecloudprovider.InstanceType) (*armcontainerservice.OSDiskType, error) {
	_ = "STUB: not implemented"
	// Karpenter defaults to Managed, but decides whether to use Ephemeral
	return nil, nil
}

func configurePriority(capacityType string) *armcontainerservice.ScaleSetPriority {
	_ = "STUB: not implemented"
	return nil
}

// Karpenter defaults to Regular

func configureOSSKUAndFIPs(nodeClass *v1beta1.AKSNodeClass, orchestratorVersion string) (*armcontainerservice.OSSKU, *bool, error) {
	_ = "STUB: not implemented"
	// Counterpart for ProvisionModeBootstrappingClient is in customscriptsbootstrap/provisionclientbootstrap.go
	return nil, nil, nil
}

func configureTaints(nodeClaim *karpv1.NodeClaim) ([]*string, []*string) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deduplicate (original behavior used sets.NewString for deduplication)

// Currently, we will use "nodeInitializationTaints" field for all taints, as "taints" field are subjected to server-side reconciliation and extra validation
// Server-side reconciliation is not necessarily a bad thing, but needs to resolve validation conflicts at least. E.g., system node cannot have hard taints other than CriticalAddonsOnly, per AKS Machine API.
// If changing, don't forget to update unit + acceptance tests accordingly.

func configureLabelsAndMode(nodeClaim *karpv1.NodeClaim, instanceType *corecloudprovider.InstanceType, capacityType string, placementScope string) (map[string]*string, *armcontainerservice.AgentPoolMode) {
	_ = "STUB: not implemented"
	// Counterpart for ProvisionModeBootstrappingClient is in customscriptsbootstrap/provisionclientbootstrap.go and instance/vminstance.go
	return nil, nil
}

// We need to get all single-valued requirement labels from the instance type and the nodeClaim to pass down to kubelet.
// We don't just include single-value labels from the instance type because in the case where the label is NOT single-value on the instance
// (i.e. there are options), the nodeClaim may have selected one of those options via its requirements which we want to include.
// These may contain restricted labels from the pod that we need to filter out; that's done by the OmitBy below.

// TODO: also do the same for taints (which don't have sanitization logic like this yet)
// Remove labels that shouldn't be sent to the API:
// - AKS-managed labels (kubernetes.azure.com/*) and legacy AKS labels
// - Kubelet-managed labels (set automatically by kubelet, e.g. hostname, zone)
// - Labels kubelet can't set (e.g. kubernetes.io/*, k8s.io/* outside allowed namespaces)

// ConfigureAKSMachineTags returns the tags to be applied to AKS machine instances and their affiliated resources.
// This includes all standard tags plus the AKS machine distinguishing tag.
func ConfigureAKSMachineTags(opts *options.Options, nodeClass *v1beta1.AKSNodeClass, nodeClaim *karpv1.NodeClaim) map[string]*string {
	_ = "STUB: not implemented"
	// TODO: move that code here instead, as AKS machine instances will be the main path forward
	// Can move when other provision modes are removed too.
	// Right now we are willing to call this just to avoid unnecessary code duplication.
	return nil
}

// Add AKS machine distinguishing tags

//nolint:gocyclo // borderline complexity violation, code is not hard to read
func configureKubeletConfig(nodeClass *v1beta1.AKSNodeClass) *armcontainerservice.KubeletConfig {
	_ = "STUB: not implemented"
	// Counterpart for ProvisionModeBootstrappingClient is in customscriptsbootstrap/provisionclientbootstrap.go and imagefamily/resolver.go
	return nil
}

// Map from v1beta1.KubeletConfiguration to AKS machine KubeletConfig

// Convert container log max size to MB

// Convert PodPidsLimit to PodMaxPids

// convertContainerLogMaxSizeToMB converts string size to MB integer
// TODO: refactor this to generic "convertSizeToMB" than just "convertContainerLogMaxSizeToMB"
func convertContainerLogMaxSizeToMB(containerLogMaxSize string) *int32 {
	_ = "STUB: not implemented"
	// TODO: move that code here instead, as AKS machine instances will be the main path forward
	// Can move when other provision modes are removed too.
	// Right now we are willing to call this just to avoid unnecessary code duplication.
	return nil
}

func convertSwapFileSizeToMB(swapFileSize string) *int32 {
	_ = "STUB: not implemented"
	// TODO: rename the utils below.
	return nil
}

func convertPodMaxPids(podPidsLimit int64) *int32 {
	_ = "STUB: not implemented"
	// TODO: move that code here instead, as AKS machine instances will be the main path forward
	// Can move when other provision modes are removed too.
	// Right now we are willing to call this just to avoid unnecessary code duplication.
	return nil
}

func configureLinuxOSConfig(nodeClass *v1beta1.AKSNodeClass) *armcontainerservice.LinuxOSConfig {
	_ = "STUB: not implemented"
	return nil
}

func configureSysctlConfig(sysctls *v1beta1.SysctlConfiguration) *armcontainerservice.SysctlConfig {
	_ = "STUB: not implemented"
	return nil
}

// parseVMImageID parses a VM image ID and extracts the required components for custom OS image headers.
// Expected format: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{imageName}/versions/{version}
func parseVMImageID(vmImageID string) (subscriptionID, resourceGroup, gallery, imageName, version string, err error) {
	_ = "STUB: not implemented"
	return "", "", "", "", "", nil
}

// Validate it's a gallery image version resource type (case-insensitive, as Azure resource IDs are case-insensitive)

// Validate we have the required parent chain (gallery -> image -> version)

// Validate none of the extracted values are empty

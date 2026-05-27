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

package instancetype

import (
	"context"
	"math"

	"github.com/Azure/skewer"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/scheduling"

	"github.com/Azure/karpenter-provider-azure/pkg/operator/options"
)

const (
	MemoryAvailable        = "memory.available"
	DefaultMemoryAvailable = "750Mi"
)

var (
	// reservedMemoryTaxGi denotes the tax brackets for memory in Gi.
	reservedMemoryTaxGi = TaxBrackets{
		{
			UpperBound: 4,
			Rate:       .25,
		},
		{
			UpperBound: 8,
			Rate:       .20,
		},
		{
			UpperBound: 16,
			Rate:       .10,
		},
		{
			UpperBound: 128,
			Rate:       .06,
		},
		{
			UpperBound: math.MaxFloat64,
			Rate:       .02,
		},
	}

	//reservedCPUTaxVCPU denotes the tax brackets for Virtual CPU cores.
	reservedCPUTaxVCPU = TaxBrackets{
		{
			UpperBound: 1,
			Rate:       .06,
		},
		{
			UpperBound: 2,
			Rate:       .04,
		},
		{
			UpperBound: 4,
			Rate:       .02,
		},
		{
			UpperBound: math.MaxFloat64,
			Rate:       .01,
		},
	}
)

// TaxBrackets implements a simple bracketed tax structure.
type TaxBrackets []struct {
	// UpperBound is the largest value this bracket is applied to.
	// The first bracket's lower bound is always 0.
	UpperBound float64

	// Rate is the percent rate of tax expressed as a float i.e. .5 for 50%.
	Rate float64
}

// Calculate expects Memory in Gi and CPU in cores.
func (t TaxBrackets) Calculate(amount float64) float64 { _ = "STUB: not implemented"; return 0 }

func NewInstanceType(
	ctx context.Context,
	sku *skewer.SKU,
	vmsize *skewer.VMSizeType,
	kc *v1beta1.KubeletConfiguration,
	region string,
	offerings cloudprovider.Offerings,
	nodeClass *v1beta1.AKSNodeClass,
	architecture string,
) *cloudprovider.InstanceType {
	_ = "STUB: not implemented"
	return nil
}

func computeRequirements(
	opts *options.Options,
	sku *skewer.SKU,
	vmsize *skewer.VMSizeType,
	architecture string,
	offerings cloudprovider.Offerings,
	region string,
	nodeClass *v1beta1.AKSNodeClass,
) scheduling.Requirements {
	_ = "STUB: not implemented"
	return *new(scheduling.Requirements)
}

// Well Known Upstream

// Well Known to Karpenter

// Well Known to Azure

// in MiB
// AKS domain.
// AKS domain.

// AKS only sets this label if FIPS is enabled, otherwise it's expected to be empty

// composites

// size parts

// SKU capabilities

// all additive feature initialized elsewhere

// composites

// size parts

func setRequirementsEphemeralOSDiskSupported(requirements scheduling.Requirements, sku *skewer.SKU) {
	_ = "STUB: not implemented"
	return
}

func setRequirementsHyperVGeneration(requirements scheduling.Requirements, sku *skewer.SKU) {
	_ = "STUB: not implemented"
	return
}

func setRequirementsGPU(requirements scheduling.Requirements, sku *skewer.SKU, vmsize *skewer.VMSizeType) {
	_ = "STUB: not implemented"
	return
}

// setRequirementsVersion sets the SKU version label, dropping "v" prefix and backfilling "1"
func setRequirementsVersion(requirements scheduling.Requirements, vmsize *skewer.VMSizeType) {
	_ = "STUB: not implemented"
	return
}

func getArchitecture(architecture string) string { _ = "STUB: not implemented"; return "" }

// unrecognized

func computeCapacity(ctx context.Context, sku *skewer.SKU, nodeClass *v1beta1.AKSNodeClass) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

// gpuNvidiaCount returns the number of Nvidia GPUs in the SKU.
func gpuNvidiaCount(sku *skewer.SKU) *resource.Quantity { _ = "STUB: not implemented"; return nil }

// gpuAMDCount returns the number of AMD GPUs in the SKU.
func gpuAMDCount(sku *skewer.SKU) *resource.Quantity { _ = "STUB: not implemented"; return nil }

// gpuTotalCount returns the total number of GPUs in the SKU for any supported vendor.
func gpuTotalCount(sku *skewer.SKU) *resource.Quantity { _ = "STUB: not implemented"; return nil }

func vcpuCount(sku *skewer.SKU) int64 { _ = "STUB: not implemented"; return 0 }

func cpu(sku *skewer.SKU) *resource.Quantity { _ = "STUB: not implemented"; return nil }

func memoryGiB(sku *skewer.SKU) float64 { _ = "STUB: not implemented"; return 0 }

// contrary to "MemoryGB" capability name, it is in GiB (!)

func memoryMiB(sku *skewer.SKU) int64 { _ = "STUB: not implemented"; return 0 }

func memoryWithoutOverhead(ctx context.Context, sku *skewer.SKU) *resource.Quantity {
	_ = "STUB: not implemented"
	return nil
}

func CalculateMemoryWithoutOverhead(vmMemoryOverheadPercent float64, skuMemoryGiB float64) *resource.Quantity {
	_ = "STUB: not implemented"
	// Consistency in abstractions could be improved here (e.g., units, returning types)
	return nil
}

func ephemeralStorage(nodeClass *v1beta1.AKSNodeClass) *resource.Quantity {
	_ = "STUB: not implemented"
	return nil
}

func pods(ctx context.Context, nc *v1beta1.AKSNodeClass) *resource.Quantity {
	_ = "STUB: not implemented"
	return nil
}

func SystemReservedResources() corev1.ResourceList {
	_ = "STUB: not implemented"
	// AKS does not set system-reserved values and only CPU and memory are considered
	// https://learn.microsoft.com/en-us/azure/aks/concepts-clusters-workloads#resource-reservations
	return *new(corev1.ResourceList)
}

func KubeReservedResources(vcpus int64, memoryGib float64) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

func EvictionThreshold() corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

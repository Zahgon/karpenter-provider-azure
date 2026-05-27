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

package customscriptsbootstrap

import (
	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
	"github.com/Azure/karpenter-provider-azure/pkg/provisionclients/models"
)

func hydrateBootstrapTokenIfNeeded(customDataDehydratable string, cseDehydratable string, bootstrapToken string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func reverseVMMemoryOverhead(vmMemoryOverheadPercent float64, adjustedMemory float64) float64 {
	_ = "STUB: not implemented"
	// This is not the best way to do it... But will be refactored later, given that retrieving the original memory properly might involves some restructure.
	// Due to the fact that it is abstracted behind the cloudprovider interface.
	return 0
}

func ConvertContainerLogMaxSizeToMB(containerLogMaxSize string) *int32 {
	_ = "STUB: not implemented"
	return nil
}

// This could be improved later

func ConvertPodMaxPids(podPidsLimit *int64) *int32 { _ = "STUB: not implemented"; return nil }

// This could be improved later

// This as well

// golint:ignore G115 already check overflow

// convertLocalDNSToModel converts v1beta1.LocalDNS to models.LocalDNSProfile
func convertLocalDNSToModel(localDNS *v1beta1.LocalDNS) *models.LocalDNSProfile {
	_ = "STUB: not implemented"
	return nil
}

// Convert VnetDNSOverrides

// Convert KubeDNSOverrides

// convertLocalDNSZoneOverrideToModel converts v1beta1.LocalDNSZoneOverride to models.LocalDNSOverride
func convertLocalDNSZoneOverrideToModel(override *v1beta1.LocalDNSZoneOverride) *models.LocalDNSOverride {
	_ = "STUB: not implemented"
	return nil
}

// convertLinuxOSConfigToModel converts v1beta1.LinuxOSConfiguration to models.CustomLinuxOSConfig
func convertLinuxOSConfigToModel(linuxOSConfig *v1beta1.LinuxOSConfiguration) *models.CustomLinuxOSConfig {
	_ = "STUB: not implemented"
	return nil
}

// convertSysctlConfigToModel converts v1beta1.SysctlConfiguration to models.SysctlConfig
func convertSysctlConfigToModel(sysctls *v1beta1.SysctlConfiguration) *models.SysctlConfig {
	_ = "STUB: not implemented"
	return nil
}

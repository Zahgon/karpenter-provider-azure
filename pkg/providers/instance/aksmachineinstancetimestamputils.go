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
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func standardizeAKSMachineTimestamp(t time.Time) time.Time {
	_ = "STUB: not implemented"
	// Truncate to centisecond precision (10ms) to ensure consistent 2-digit format
	return *new(time.Time)
}

// NewAKSMachineTimestamp returns the current time truncated to centisecond precision for AKS machine creation timestamps
func NewAKSMachineTimestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func ZeroAKSMachineTimestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// AKSMachineTimestampToMeta converts a time.Time to metav1.Time for AKS machine creation timestamps
func AKSMachineTimestampToMeta(t time.Time) metav1.Time {
	_ = "STUB: not implemented"
	return *new(metav1.Time)
}

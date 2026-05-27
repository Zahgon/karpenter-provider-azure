// Portions Copyright (c) Microsoft Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metrics

import (
	dto "github.com/prometheus/client_model/go"
)

// FindMetricWithLabelValues locates a metric by name and label values within the controller-runtime registry.
// It returns the metric (or nil if not found) and any error produced while gathering metrics.
func FindMetricWithLabelValues(metricName string, labels map[string]string) (*dto.Metric, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FailureMetricLabels produces a failure metric label set by merging the provided base labels with
// the phase label and any additional label maps. Later maps override earlier values when conflicts occur.
func FailureMetricLabels(base map[string]string, phase string, extra ...map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func metricLabelsEqual(metric *dto.Metric, expected map[string]string) bool {
	_ = "STUB: not implemented"
	return false
}

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

package debug

import (
	"context"
	"time"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type EventClient struct {
	start      time.Time
	kubeClient client.Client
}

func NewEventClient(kubeClient client.Client) *EventClient { _ = "STUB: not implemented"; return nil }

func (c *EventClient) DumpEvents(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *EventClient) dumpPodEvents(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *EventClient) dumpNodeEvents(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func filterTestEvents(events []corev1.Event, startTime time.Time) []corev1.Event {
	_ = "STUB: not implemented"
	return nil
}

func collateEvents(events []corev1.Event) map[corev1.ObjectReference]*corev1.EventList {
	_ = "STUB: not implemented"
	return nil
}

// Sort the events in ascending order by event time

// Return the time that should be used for sorting, which can come from
// various places in corev1.Event.
func eventTime(event corev1.Event) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Partially copied from
// https://github.com/kubernetes/kubernetes/blob/04ee339c7a4d36b4037ce3635993e2a9e395ebf3/staging/src/k8s.io/kubectl/pkg/describe/describe.go#L4232
func getEventInformation(o corev1.ObjectReference, el *corev1.EventList) string {
	_ = "STUB: not implemented"
	return ""
}

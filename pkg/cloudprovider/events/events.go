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

package events

import (
	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/events"
)

const (
	AsyncProvisioningReason   = "AsyncProvisioningError"
	NodeClassResolutionReason = "NodeClassResolutionError"
)

func NodePoolFailedToResolveNodeClass(nodePool *v1.NodePool) events.Event {
	_ = "STUB: not implemented"
	return *new(events.Event)
}

func NodeClaimFailedToResolveNodeClass(nodeClaim *v1.NodeClaim) events.Event {
	_ = "STUB: not implemented"
	return *new(events.Event)
}

func NodeClaimFailedToRegister(nodeClaim *v1.NodeClaim, err error) events.Event {
	_ = "STUB: not implemented"
	return *new(events.Event)
}

const truncateAt = 500

func truncateMessage(msg string) string { _ = "STUB: not implemented"; return "" }

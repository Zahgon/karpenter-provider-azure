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

package inplaceupdate

import (
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

type tagsChangedPredicate struct {
	predicate.Funcs
}

var _ predicate.Predicate = tagsChangedPredicate{}

func (p tagsChangedPredicate) Delete(e event.DeleteEvent) bool {
	_ = "STUB: not implemented"
	// We never want updates on delete
	return false
}

func (p tagsChangedPredicate) Update(e event.UpdateEvent) bool {
	_ = "STUB: not implemented"
	return false
}

// This isn't expected, so propagate the event so we don't miss anything

// This isn't expected, so propagate the event so we don't miss anything

// If we don't know the type, we assume it has changed

// If we don't know the type, we assume it has changed

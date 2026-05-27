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

package sync

import "sync"

// Map is a typesafe wrapper around a map with sync.RWMutex.
// It exposes the same methods as sync.Map but with generic type parameters.
// NOTE: This lives here for now because most of the time you may want to enforce some other invariants while under the RWMutex (not just the map invaraints),
// so this helper exists primarily for the fake pkg which is only used for tests. We can think about promoting this to a more general utility package if we
// find more use cases for it.
type Map[K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
}

func (m *Map[K, V]) init() {
	if m.m == nil {
		m.m = make(map[K]V)
	}
}

func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (m *Map[K, V]) Store(key K, value V) { _ = "STUB: not implemented"; return }

func (m *Map[K, V]) Delete(key K) { _ = "STUB: not implemented"; return }

func (m *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (m *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (m *Map[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (m *Map[K, V]) CompareAndSwap(key K, old, new V) (swapped bool) {
	_ = "STUB: not implemented"
	return false
}

func (m *Map[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	_ = "STUB: not implemented"
	return false
}

func (m *Map[K, V]) Range(f func(key K, value V) bool) { _ = "STUB: not implemented"; return }

func (m *Map[K, V]) Clear() { _ = "STUB: not implemented"; return }

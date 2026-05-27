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

package azure

import (
	"sync"
)

type Cleanup func() error

type Tracker struct {
	// Can't easily use the armresources.NewClient here for generic deletion because we would need to know the api-version for each resource. It's easier
	// to just have cleanup passed in as part of Add. If we really wanted to, we could probably hook the clients, detect PUTs in a PerRequestPolicy, extract
	// the API version and save it here, but that seems like more trouble than it's worth.
	ids map[string]Cleanup
	mu  sync.Mutex
}

func NewTracker() *Tracker { _ = "STUB: not implemented"; return nil }

func (t *Tracker) Add(id string, cleanup Cleanup) { _ = "STUB: not implemented"; return }

// TODO: This isn't case-insensitive, but it probably should be

func (t *Tracker) Cleanup() error {
	_ = "STUB: not implemented"
	// We could avoid holding the lock across cleanup - we don't expect adds to happen during cleanup so for now not worrying about it
	return nil
}

// TODO: Should be using the test logger

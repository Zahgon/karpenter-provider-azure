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

package fake

import (
	"iter"
	"sync"
)

// AtomicPtr is intended for use in mocks to easily expose variables for use in testing.  It makes setting and retrieving
// the values race free by wrapping the pointer itself in a mutex.  There is no Get() method, but instead a Clone() method
// deep copies the object being stored by serializing/de-serializing it from JSON.  This pattern shouldn't be followed
// anywhere else but is an easy way to eliminate races in our tests.
type AtomicPtr[T any] struct {
	mu    sync.Mutex
	value *T
}

func (a *AtomicPtr[T]) Set(v *T) { _ = "STUB: not implemented"; return }

func (a *AtomicPtr[T]) IsNil() bool { _ = "STUB: not implemented"; return false }

func (a *AtomicPtr[T]) Clone() *T { _ = "STUB: not implemented"; return nil }

func clone[T any](v *T) *T { _ = "STUB: not implemented"; return nil }

func (a *AtomicPtr[T]) Reset() { _ = "STUB: not implemented"; return }

type AtomicError struct {
	mu  sync.Mutex
	err error

	calls    int
	maxCalls int
}

func (e *AtomicError) Reset() { _ = "STUB: not implemented"; return }

func (e *AtomicError) IsNil() bool { _ = "STUB: not implemented"; return false }

// Get is equivalent to the error being called, so we increase
// number of calls in this function
func (e *AtomicError) Get() error { _ = "STUB: not implemented"; return nil }

func (e *AtomicError) Set(err error, opts ...AtomicErrorOption) { _ = "STUB: not implemented"; return }

type AtomicErrorOption func(atomicError *AtomicError)

func MaxCalls(maxCalls int) AtomicErrorOption {
	_ = "STUB: not implemented"
	// Setting to 0 is equivalent to allowing infinite errors to API
	return *new(AtomicErrorOption)
}

// AtomicPtrStack exposes a slice of a pointer type in a race-free manner. The interface is just enough to replace the
// set.Set usage in our previous tests.
type AtomicPtrStack[T any] struct {
	mu     sync.Mutex
	values []*T
}

func (a *AtomicPtrStack[T]) Reset() { _ = "STUB: not implemented"; return }

func (a *AtomicPtrStack[T]) Add(input *T) { _ = "STUB: not implemented"; return }

func (a *AtomicPtrStack[T]) Len() int { _ = "STUB: not implemented"; return 0 }

func (a *AtomicPtrStack[T]) Pop() *T { _ = "STUB: not implemented"; return nil }

// All returns an iterator over all values in the stack
// NOTE: This pops the values from the stack
func (a *AtomicPtrStack[T]) All() iter.Seq[*T] { _ = "STUB: not implemented"; return nil }

// AtomicPtrSlice exposes a slice of a pointer type in a race-free manner.
type AtomicPtrSlice[T any] struct {
	mu     sync.Mutex
	values []*T
}

func (a *AtomicPtrSlice[T]) Reset() { _ = "STUB: not implemented"; return }

func (a *AtomicPtrSlice[T]) Append(input ...*T) { _ = "STUB: not implemented"; return }

func (a *AtomicPtrSlice[T]) Len() int { _ = "STUB: not implemented"; return 0 }

func (a *AtomicPtrSlice[T]) Get(index int) *T { _ = "STUB: not implemented"; return nil }

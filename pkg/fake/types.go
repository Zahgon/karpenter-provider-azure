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
	"context"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

type MockedFunction[I any, O any] struct {
	Output          AtomicPtr[O]      // Output to return on call to this function
	CalledWithInput AtomicPtrStack[I] // Stack used to keep track of passed input to this function
	Error           AtomicError       // Error to return a certain number of times defined by custom error options

	successfulCalls     atomic.Int32 // Internal construct to keep track of the number of times this function has successfully been called
	failedCalls         atomic.Int32 // Internal construct to keep track of the number of times this function has failed (with error)
	customTransformerMu sync.RWMutex
	customTransformer   func(*I) error // Optional hook called before the default transformer; can modify input or return an error to short-circuit
}

// SetCustomTransformer sets an optional hook that is called before the default
// transformer in Invoke. The hook may modify the input in place or return an
// error to short-circuit the call.
func (m *MockedFunction[I, O]) SetCustomTransformer(fn func(*I) error) {
	_ = "STUB: not implemented"
	return
}

func (m *MockedFunction[I, O]) getCustomTransformer() func(*I) error {
	_ = "STUB: not implemented"
	return nil
}

// Reset must be called between tests otherwise tests will pollute
// each other.
func (m *MockedFunction[I, O]) Reset() { _ = "STUB: not implemented"; return }

func (m *MockedFunction[I, O]) Invoke(input *I, defaultTransformer func(*I) (O, error)) (O, error) {
	_ = "STUB: not implemented"
	return *new(O), nil
}

func (m *MockedFunction[I, O]) Calls() int { _ = "STUB: not implemented"; return 0 }

func (m *MockedFunction[I, O]) SuccessfulCalls() int { _ = "STUB: not implemented"; return 0 }

func (m *MockedFunction[I, O]) FailedCalls() int { _ = "STUB: not implemented"; return 0 }

type MockedLRO[I any, O any] struct {
	MockedFunction[I, O]
	BeginError AtomicError // Error to return a certain number of times defined by custom error options (for Begin)
}

// Reset must be called between tests otherwise tests will pollute each other.
func (m *MockedLRO[I, O]) Reset() { _ = "STUB: not implemented"; return }

func (m *MockedLRO[I, O]) Invoke(input *I, defaultTransformer func(*I) (*O, error)) (*runtime.Poller[O], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MockedLRO[I, O]) Calls() int { _ = "STUB: not implemented"; return 0 }

func (m *MockedLRO[I, O]) SuccessfulCalls() int { _ = "STUB: not implemented"; return 0 }

func (m *MockedLRO[I, O]) FailedCalls() int { _ = "STUB: not implemented"; return 0 }

// MockHandler returns a pre-defined result or error.
type MockHandler[T any] struct {
	result *T
	err    error
}

// Done returns true if the LRO has reached a terminal state. TrivialHandler is always done.
func (h MockHandler[T]) Done() bool {
	_ = "STUB: not implemented"

	// Poll fetches the latest state of the LRO.
	return false
}

func (h MockHandler[T]) Poll(context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Result is called once the LRO has reached a terminal state. It populates the out parameter
// with the result of the operation.
func (h MockHandler[T]) Result(_ context.Context, result *T) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: may need to deep copy

// newMockPoller returns a poller with a mock handler that returns the given result and error.
func newMockPoller[T any](result *T, err error) (*runtime.Poller[T], error) {
	_ = "STUB: not implemented"
	// http.Response and Pipeline are not used
	return nil, nil
}

// Response: &result at the poller level is not needed, result from handler is always used

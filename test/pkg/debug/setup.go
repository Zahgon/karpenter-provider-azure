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

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

	. "github.com/onsi/gomega"
)

const (
	NoWatch  = "NoWatch"
	NoEvents = "NoEvents"
)

var m *Monitor
var e *EventClient

func BeforeEach(ctx context.Context, config *rest.Config, kubeClient client.Client) {
	_ = "STUB: not implemented"
	// If the test is labeled as NoWatch, then the node/pod monitor will just list at the beginning
	// of the test rather than perform a watch during it
	return
}

func AfterEach(ctx context.Context) { _ = "STUB: not implemented"; return }

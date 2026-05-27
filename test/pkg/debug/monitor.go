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
	"sync"

	"github.com/awslabs/operatorpkg/controller"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

type Monitor struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
	mgr    manager.Manager
}

func New(ctx context.Context, config *rest.Config, kubeClient client.Client) *Monitor {
	_ = "STUB: not implemented"
	return nil
}

// this context is only meant for monitor start/stop

// MustStart starts the debug monitor
func (m *Monitor) MustStart() { _ = "STUB: not implemented"; return }

// Stop stops the monitor
func (m *Monitor) Stop() { _ = "STUB: not implemented"; return }

func newControllers(kubeClient client.Client) []controller.Controller {
	_ = "STUB: not implemented"
	return nil
}

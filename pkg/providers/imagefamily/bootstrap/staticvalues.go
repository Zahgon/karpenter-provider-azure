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

package bootstrap

import (
	_ "embed"
	"text/template"
)

const (
	globalAKSMirror = "https://acs-mirror.azureedge.net"
)

// NOTE: embed only works on vars not defined in a function, so without putting this into an internal package for encapulation, we are stuck with these remaining vars.
var (
	//go:embed cse_cmd.sh.gtpl
	customDataTemplateText string

	//go:embed  containerd.toml.gtpl
	containerdConfigTemplateText string

	//go:embed sysctl.conf
	sysctlContent []byte
)

func getCustomDataTemplate() *template.Template { _ = "STUB: not implemented"; return nil }

func getContainerdConfigTemplate() *template.Template { _ = "STUB: not implemented"; return nil }

func getBaseKubeletFlags() map[string]string {
	_ = "STUB: not implemented"
	// source note: unique per nodepool. partially user-specified, static, and RP-generated
	// removed --image-pull-progress-deadline=30m  (not in 1.24?)
	// removed --network-plugin=cni (not in 1.24?)
	// removed --azure-container-registry-config (not in 1.30)
	// removed --keep-terminated-pod-volumes (not in 1.31)
	return nil
}

func getStaticNodeBootstrapVars() *NodeBootstrapVariables { _ = "STUB: not implemented"; return nil }

// baseline, covering unused (-), static (s), and unsupported (n) fields,
// as well as defaults, cluster/node level (cd/td/xd)

// n
// n
// n
// td
// -
// -
// cd
// -
// -
// xd
// xd
// xd
// -
// -
// ad
// - [currently required, installCNI in provisioning scripts depends on CNI_PLUGINS_URL]
// - [currently required, same]
// s
// s
// s
// s
// s
// s
// s
// s
// s
// s
// s
// xd
// s
// s
// xd
// s
// xd
// s
// s
// -
// cd
// s
// -
// td
// s
// -
// -
// -
// -
// -
// -
// td
// n
// s
// td
// cd
// cd
// cd
// cd
// cd

// s
// cd
// td
// td
// td
// n
// n
// n
// n
// s
// s
// s
// s
// s
// cd
// cd
// s
// cd
// cd
// cd
// s
// s
// s
// cd
// cd
// s
// -
// -
// s
// s
// td
// td
// cd
// cd
// cd
// td
// td
// td
// psX
// s
// kd
// n
// s only static for karpenter
// s karpenter does not support kubenet

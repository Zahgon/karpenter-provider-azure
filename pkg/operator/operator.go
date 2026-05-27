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

package operator

import (
	"context"
	"net"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/go-logr/logr"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	karpv1 "sigs.k8s.io/karpenter/pkg/apis/v1"

	"sigs.k8s.io/karpenter/pkg/operator"

	"github.com/Azure/karpenter-provider-azure/pkg/auth"
	azurecache "github.com/Azure/karpenter-provider-azure/pkg/cache"

	"github.com/Azure/karpenter-provider-azure/pkg/providers/azclient"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/imagefamily"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/instance"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/instancetype"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/kubernetesversion"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/launchtemplate"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/loadbalancer"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/pricing"
)

func init() {
	karpv1.NormalizedLabels = lo.Assign(karpv1.NormalizedLabels, map[string]string{"topology.disk.csi.azure.com/zone": corev1.LabelTopologyZone})
}

type Operator struct {
	*operator.Operator

	// InClusterKubernetesInterface is a Kubernetes client that can be used to talk to the APIServer
	// of the cluster where the Karpenter pod is running. This is usually the same as operator.KubernetesInterface,
	// but may be different if Karpenter is running in a different cluster than the one it manages.
	InClusterKubernetesInterface kubernetes.Interface

	// ManagedDynamicInterface is a dynamic client over the managed (workload) cluster, used
	// by callers that need to inspect out-of-tree CRDs (e.g. Cilium / Calico NetworkPolicy)
	// alongside the typed client on the embedded operator (operator.KubernetesInterface).
	ManagedDynamicInterface dynamic.Interface

	UnavailableOfferingsCache *azurecache.UnavailableOfferings

	KubernetesVersionProvider kubernetesversion.KubernetesVersionProvider
	ImageProvider             imagefamily.NodeImageProvider
	ImageResolver             imagefamily.Resolver
	LaunchTemplateProvider    *launchtemplate.Provider
	PricingProvider           *pricing.Provider
	InstanceTypesProvider     instancetype.Provider
	VMInstanceProvider        *instance.DefaultVMProvider
	AKSMachineProvider        *instance.DefaultAKSMachineProvider
	LoadBalancerProvider      *loadbalancer.Provider
	AZClient                  *azclient.AZClient
}

func kubeDNSIP(ctx context.Context, kubernetesInterface kubernetes.Interface) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

func NewOperator(ctx context.Context, operator *operator.Operator) (context.Context, *Operator) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// NOTE: we prefer this over the cleaner azConfig := lo.Must(GetAzConfig()), as when initializing the client there are helpful error messages in initializing clients and the azure config

// Get a token to ensure we can

// These options are set similarly to those used by operator.KubernetesInterface

// Build a dynamic client over the managed (workload) cluster config. operator.GetConfig()
// is the rest.Config that controller-runtime uses for the manager, which targets the
// workload cluster (via mounted kubeconfig in CCP, or same as in-cluster in non-CCP).
// LocalDNS gate evaluation reads out-of-tree CRDs (Cilium / Calico NetworkPolicy) through
// this; the typed client lives on the embedded operator (operator.KubernetesInterface).

// fall back to default

// Ensure we're able to hydrate instance types before starting any controllers
// that depend on them. The instance type controller will refresh this list
// perioidcally once all controllers are running.

func GetAZConfig() (*auth.Config, error) { _ = "STUB: not implemented"; return nil, nil }

func getCABundle(restConfig *rest.Config) (*string, error) {
	_ = "STUB: not implemented"
	// Discover CA Bundle from the REST client. We could alternatively
	// have used the simpler client-go InClusterConfig() method.
	// However, that only works when Karpenter is running as a Pod
	// within the same cluster it's managing.
	return nil, nil
}

// fills in CAData!

func getVnetGUID(ctx context.Context, creds azcore.TokenCredential, cfg *auth.Config, subnetID string) (string, error) {
	_ = "STUB: not implemented"
	// TODO: Current the VNET client isn't used anywhere but this method. As such, it is not
	// held on azclient like the other clients.
	// We should possibly just put the vnet client on azclient, and then pass azclient in here, rather than
	// constructing the VNET client here separate from all the other Azure clients.
	return "", nil
}

// WaitForCRDs waits for the required CRDs to be available with a timeout
func WaitForCRDs(ctx context.Context, timeout time.Duration, config *rest.Config, log logr.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// ensureToken ensures we can get a token for the Azure environment. Note that this doesn't actually
// use the token for anything, it just checks that we can get one.
func ensureToken(cred azcore.TokenCredential, env *auth.Environment) error {
	_ = "STUB: not implemented"
	// Short timeout to avoid hanging forever if something bad happens
	return nil
}

func getCredential(env *auth.Environment) (azcore.TokenCredential, error) {
	_ = "STUB: not implemented"
	// TODO: Don't use NewDefaultAzureCredential
	return *new(azcore.TokenCredential), nil
}

func getRequiredGVKs() []schema.GroupVersionKind {
	_ = "STUB: not implemented"
	// controller-runtime internal, ignore them as we don't watch them
	return nil
}

// Ignore lists as well, we don't watch these

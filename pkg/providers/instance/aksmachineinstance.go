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

package instance

import (
	"context"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
	"k8s.io/apimachinery/pkg/util/sets"
	karpv1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	corecloudprovider "sigs.k8s.io/karpenter/pkg/cloudprovider"

	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
	"github.com/Azure/karpenter-provider-azure/pkg/cache"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/allocationstrategy"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/azclient"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/imagefamily"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/instance/machinecache"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/instance/offerings"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/instancetype"
)

// Notes on terminology:
// An "instance" is a remote object, created by the API based on the template.
// A "template" is a local struct, populated from Karpenter-provided parameters with the logic further below.
// A "template" shares the struct with an "instance" representation. But read-only fields may not be populated. Ideally, the types should have been separated to avoid making cross-module assumption of the existence of certain fields.
//
// TODO: Consider extracting the template-related fields (AKSMachineTemplate, AKSMachineName, InstanceType, CapacityType, Zone, AKSMachineID, AKSMachineNodeImageVersion, VMResourceID)
// into a dedicated struct (e.g., AKSMachineDetails or AKSMachineTemplateInfo). This would clarify the relationship between
// the promise fields and functions like BuildNodeClaimFromAKSMachineTemplate, as well as reduce the number of loose arguments passed around.
// More discussion: https://github.com/Azure/karpenter-provider-azure/pull/1197#discussion_r2482957255
type AKSMachinePromise struct {
	waitFunc    func() error
	providerRef AKSMachineProvider

	AKSMachineTemplate *armcontainerservice.Machine
	AKSMachineName     string
	InstanceType       *corecloudprovider.InstanceType // Despite the reference nature, this is guaranteed to exist
	CapacityType       string
	Zone               string

	AKSMachineID               string
	AKSMachineNodeImageVersion string
	VMResourceID               string
	CreationTimestamp          time.Time
}

func NewAKSMachinePromise(
	providerRef AKSMachineProvider,
	aksMachineTemplate *armcontainerservice.Machine,
	waitFunc func() error,
	aksMachineName string,
	instanceType *corecloudprovider.InstanceType,
	capacityType string,
	zone string,
	aksMachineID string,
	aksMachineNodeImageVersion string,
	vmResourceID string,
	creationTimestamp time.Time,
) *AKSMachinePromise {
	_ = "STUB: not implemented"
	return nil
}

func (p *AKSMachinePromise) Cleanup(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *AKSMachinePromise) Wait() error { _ = "STUB: not implemented"; return nil }

func (p *AKSMachinePromise) GetInstanceName() string { _ = "STUB: not implemented"; return "" }

type opt struct {
	useCache bool
}

// Option defines a functional option for configuring AKSMachineProvider methods that accept optional behavior.
type Option func(*opt)

// WithCache is an Option that tells the AKSMachineProvider to first use the machine cache
// when attempting to retrieve an AKS machine before falling back to calling the Azure API.
func WithCache() Option { _ = "STUB: not implemented"; return *new(Option) }

type AKSMachineProvider interface {
	// BeginCreate starts the creation of an AKS machine instance.
	// Returns a promise that must be waited on to complete the creation.
	BeginCreate(ctx context.Context, nodeClass *v1beta1.AKSNodeClass, nodeClaim *karpv1.NodeClaim, instanceTypes []*corecloudprovider.InstanceType) (*AKSMachinePromise, error)
	// Update updates the AKS machine instance with the specified name. Uses ETag for optimistic concurrency control.
	// Return NodeClaimNotFoundError if not found.
	Update(ctx context.Context, aksMachineName string, aksMachine armcontainerservice.Machine, etag *string) error
	// Get retrieves the AKS machine instance with the specified AKS machine name. Return NodeClaimNotFoundError if not found.
	Get(ctx context.Context, aksMachineName string, opts ...Option) (*armcontainerservice.Machine, error)
	// List lists all AKS machine instances in the cluster.
	List(ctx context.Context, opts ...Option) ([]*armcontainerservice.Machine, error)
	// Delete deletes the AKS machine instance with the specified name. Return NodeClaimNotFoundError if not found.
	Delete(ctx context.Context, aksMachineName string) error
	// GetMachinesPoolLocation returns the location of the AKS machines pool. The only reason this need to be exported is because armcontainerservice.Machine does not have the location field.
	GetMachinesPoolLocation() string
}

// assert that DefaultAKSMachineProvider implements Provider interface
var _ AKSMachineProvider = (*DefaultAKSMachineProvider)(nil)

type DefaultAKSMachineProvider struct {
	azClient                   *azclient.AZClient
	instanceTypeProvider       instancetype.Provider
	allocationStrategyProvider allocationstrategy.Provider
	imageResolver              imagefamily.Resolver
	subscriptionID             string
	clusterResourceGroup       string
	clusterName                string
	aksMachinesPoolName        string // Only support one AKS machine pool at a time, for now.
	aksMachinesPoolLocation    string
	batchCreationEnabled       bool
	provisioningErrorHandling  *offerings.ErrorDetailHandler
	beginCreateErrorHandling   *offerings.AKSMachineBeginCreateErrorHandler
	deletingMachines           sets.Set[string] // tracks in-flight delete operations by machine name
	deletingMachinesMu         sync.RWMutex
	machineCache               *machinecache.MachineCache
}

func NewAKSMachineProvider(
	azClient *azclient.AZClient,
	instanceTypeProvider instancetype.Provider,
	allocationStrategyProvider allocationstrategy.Provider,
	imageResolver imagefamily.Resolver,
	offeringsCache *cache.UnavailableOfferings,
	subscriptionID string,
	clusterResourceGroup string,
	clusterName string,
	aksMachinesPoolName string,
	aksMachinesPoolLocation string,
	batchCreationEnabled bool,
	machineCache *machinecache.MachineCache,
) *DefaultAKSMachineProvider {
	_ = "STUB: not implemented"
	return nil
}

// BeginCreate creates an instance given the constraints.
// Note that the returned instance may not be finished provisioning yet.
// Errors that occur on the "sync side" of the VM create, such as BadRequest due to invalid user input, and similar, will have the error returned here.
// Errors that occur on the "async side" of the VM create (after the request is accepted) will be returned from AKSMachinePromise.Wait().
func (p *DefaultAKSMachineProvider) BeginCreate(
	ctx context.Context,
	nodeClass *v1beta1.AKSNodeClass,
	nodeClaim *karpv1.NodeClaim,
	instanceTypes []*corecloudprovider.InstanceType,
) (*AKSMachinePromise, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Clean up if creation fails.

// We don't return the cleanup error here, as we want to return the original error from beginCreateMachine

func (p *DefaultAKSMachineProvider) Update(ctx context.Context, aksMachineName string, aksMachine armcontainerservice.Machine, etag *string) error {
	_ = "STUB: not implemented"
	return nil
}

// Can only be AKS machines pool not found.
// Suggestion: separate the util function to not cover more than needed?

// ASSUMPTION: the AKS machine will be in the current p.aksMachinesPoolName. Otherwise need rework to pass the pool name in.
func (p *DefaultAKSMachineProvider) Get(ctx context.Context, aksMachineName string, opts ...Option) (*armcontainerservice.Machine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Possible when this option field is not populated, which is not required when PROVISION_MODE is not aksmachineapi.
// But an AKS machine instance exists, whether added manually or from before switching PROVISION_MODE.
// So, we respond similarly to if AKS machines pool is not found.

func (p *DefaultAKSMachineProvider) List(ctx context.Context, opts ...Option) ([]*armcontainerservice.Machine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Possible when this option field is not populated, which is not required when PROVISION_MODE is not aksmachineapi.
// So, we respond similarly to if AKS machines pool is not found.

func (p *DefaultAKSMachineProvider) Delete(ctx context.Context, aksMachineName string) error {
	_ = "STUB: not implemented"
	return nil
}

// If there's already an in-flight delete for this machine, return immediately.

// Note that 'Get' also satisfies cloudprovider.Delete contract expectation (from v1.3.0)
// of returning cloudprovider.NewNodeClaimNotFoundError if the instance is already deleted
// This get exists to deal with the case where the operator restarted during the course of a deletion.
// With it, we may do an extra unneeded get before delete, but without it we may erroneously issue
// 2 deletes if the instance was being deleted and the operator restarted.
// Since get quota is generally higher, we prefer to check w/ get rather than issue 2 deletes.

func (p *DefaultAKSMachineProvider) GetMachinesPoolLocation() string {
	_ = "STUB: not implemented"
	return ""
}

func (p *DefaultAKSMachineProvider) deleteMachine(ctx context.Context, aksMachineName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Suggestion: we could utilize this batch capability to optimize performance

// beginCreateMachine starts the creation of an AKS machine instance.
// The returned AKSMachinePromise must be called to gather any errors
// that are retrieved during async provisioning, as well as to complete the provisioning process.
//

func (p *DefaultAKSMachineProvider) beginCreateMachine(
	ctx context.Context,
	nodeClass *v1beta1.AKSNodeClass,
	nodeClaim *karpv1.NodeClaim,
	instanceTypes []*corecloudprovider.InstanceType,
	aksMachineName string,
) (*AKSMachinePromise, error) {
	_ = "STUB: not implemented"
	// Reuse the existing AKS machine if it exists, and skip the creation.
	//
	// This repeated creation is possible in a race condition where we successfully launched the AKS machine
	// but crashed/restarted before we could set the NodeClaim's Launched status. In that case, the NodeClaim
	// will be re-queued for creation, but the AKS machine already exists.
	//
	// We assume that if an AKS machine exists, we successfully created it with the right parameters from the
	// NodeClaim during a previous run.
	// However, offerings properties (e.g., instanceType, capacityType, zone) are decided below and not deterministic,
	// thus, they may differ in the new attempts.
	//
	// If we attempted to recreate with different properties, the API would reject the request due to property
	// conflicts, blocking the NodeClaim until liveness TTL is hit. This guard will just reuse the existing AKS machine,
	// potentially with original offerings properties, which is acceptable, as it just complete the original intention.
	return nil, nil
}

// Existing AKS machine found, reuse it.

// Not fatal. Will fall back to normal creation.

// Decide on offerings

// Build the AKS machine template

// Call the AKS machine API with the template to create the AKS machine instance

// Branch between batch and non-batch creation paths.

// beginCreateMachineBatch handles the batch creation path using the AKS machines header batch API and GET-based poller.
func (p *DefaultAKSMachineProvider) beginCreateMachineBatch(
	ctx context.Context,
	aksMachineTemplate *armcontainerservice.Machine,
	aksMachineName string,
	instanceType *corecloudprovider.InstanceType,
	capacityType string,
	zone string,
) (*AKSMachinePromise, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get once after begin create to retrieve VMResourceID.
// In fact, the AKS machine object we want here is already returned with the PUT request above. However, the SDK have prevented us from accessing it easily.
// TODO: find a way to access that instead of making another GET call like this.

// Return LRO

// beginCreateMachineNonBatch handles the non-batch creation path using the standard AKS machines API and SDK poller.
func (p *DefaultAKSMachineProvider) beginCreateMachineNonBatch(
	ctx context.Context,
	aksMachineTemplate *armcontainerservice.Machine,
	aksMachineName string,
	instanceType *corecloudprovider.InstanceType,
	capacityType string,
	zone string,
) (*AKSMachinePromise, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get once after begin create to retrieve VMResourceID.
// In fact, the AKS machine object we want here is already returned with the PUT request above. However, the SDK have prevented us from accessing it easily.
// TODO: find a way to access that instead of making another GET call like this.

// Return LRO

// Use SDK poller (non-batch case)
// This may panic if it is deleted mid-way.

// Could be quota error; will be handled with custom logic below

// Get once after begin create to retrieve error details. This is because if the poller returns error, the sdk doesn't let us look at the real results.

// This should not be expected.

// For use in beginCreateMachine only. Otherwise need to rework parameters, do nil check better, and generalize error messaging.
func (p *DefaultAKSMachineProvider) handleMachineProvisioningError(ctx context.Context, phase string, aksMachineName string, instanceType *corecloudprovider.InstanceType, zone string, capacityType string, provisioningError *armcontainerservice.ErrorDetail) error {
	_ = "STUB: not implemented"
	return nil
}

// This should be VM creation error.
// ASSUMPTION: the length of details is always <= 1. And VM creation error Karpenter may expect is always at Details[0].
// Suggestion: suggest API change to have an explicit VM create error, if not changing Karpenter to rely on AKS machine ProvisioningError instead?

// Fallback to AKS machine API-level error. Though, this is unlikely to be handled by Karpenter.

// If error is handled, return it (wrapped)

func (p *DefaultAKSMachineProvider) handleMachineBeginCreateError(ctx context.Context, aksMachineName string, instanceType *corecloudprovider.InstanceType, zone string, capacityType string, he *offerings.HandlableError) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *DefaultAKSMachineProvider) reuseExistingMachine(ctx context.Context, aksMachineName string, nodeClaim *karpv1.NodeClaim, instanceTypes []*corecloudprovider.InstanceType, existingAKSMachine *armcontainerservice.Machine) (*AKSMachinePromise, error) {
	_ = "STUB: not implemented"
	// Reconstruct properties from existing AKS machine instance.
	return nil, nil
}

// This is not included in validateRetrievedAKSMachineBasicProperties as inplaceupdate can repair it.
// Although, we don't want to reuse a machine until that happens.

// Might be possible from NodePool name hash collision within AKS machine name
// See how AKS machine name is generated for more details.
// ASSUMPTION: repeated failure will eventually result in NodeClaim reaching registration TTL, then gets re-created with the new hash, recovering from the collision.

// Unfortunately, that was more like a remain than a usable aksMachine.
// ASSUMPTION: this is irrecoverable (i.e., polling would have failed).

// We hope the AKS machine completed provisioning at this point. Otherwise, if fails, it would not be handled until registration TTL.
// Suggestion: create a new poller just to handle it the same way as new machines. That will improve performance for such cases.

func (p *DefaultAKSMachineProvider) getCreatedMachineAndHandleEarlyProvisioningError(ctx context.Context, aksMachineName string, instanceType *corecloudprovider.InstanceType, zone string, capacityType string) (*armcontainerservice.Machine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We luckily catch failed state early (compared to during polling).
// ASSUMPTION: this is irrecoverable (i.e., polling would have failed).

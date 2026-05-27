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

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"k8s.io/apimachinery/pkg/util/sets"
	azureclouds "sigs.k8s.io/cloud-provider-azure/pkg/azclient"
	karpv1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	corecloudprovider "sigs.k8s.io/karpenter/pkg/cloudprovider"

	"github.com/Azure/karpenter-provider-azure/pkg/apis/v1beta1"
	"github.com/Azure/karpenter-provider-azure/pkg/auth"
	"github.com/Azure/karpenter-provider-azure/pkg/cache"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/allocationstrategy"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/azclient"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/instance/offerings"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/instancetype"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/launchtemplate"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/loadbalancer"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/networksecuritygroup"
)

var (
	KarpCapacityTypeToVMPriority = map[string]armcompute.VirtualMachinePriorityTypes{
		karpv1.CapacityTypeSpot:     armcompute.VirtualMachinePriorityTypesSpot,
		karpv1.CapacityTypeOnDemand: armcompute.VirtualMachinePriorityTypesRegular,
	}
	VMPriorityToKarpCapacityType = map[armcompute.VirtualMachinePriorityTypes]string{
		armcompute.VirtualMachinePriorityTypesSpot:    karpv1.CapacityTypeSpot,
		armcompute.VirtualMachinePriorityTypesRegular: karpv1.CapacityTypeOnDemand,
	}
	// Note that there is no ScaleSetPriorityToKarpCapacityType because the karpenter.sh/capacity-type
	// label is the "official" label that we actually key priority off of. Selection still works though
	// because when we list instance types on-demand offerings always have v1beta1.ScaleSetPriorityRegular
	// and spot instances always have v1beta1.ScaleSetPrioritySpot, so the correct karpenter.sh/capacity-type
	// label is still selected even if the user is using kubernetes.azure.com/scalesetpriority only on the NodePool.
	VMPriorityToScaleSetPriority = map[armcompute.VirtualMachinePriorityTypes]string{
		armcompute.VirtualMachinePriorityTypesSpot:    v1beta1.ScaleSetPrioritySpot,
		armcompute.VirtualMachinePriorityTypesRegular: v1beta1.ScaleSetPriorityRegular,
	}
	VMPriorityToPriority = map[armcompute.VirtualMachinePriorityTypes]string{
		armcompute.VirtualMachinePriorityTypesSpot:    v1beta1.PrioritySpot,
		armcompute.VirtualMachinePriorityTypesRegular: v1beta1.PriorityRegular,
	}

	aksIdentifyingExtensionEnvs = sets.New(
		azureclouds.PublicCloud.Name,
		azureclouds.ChinaCloud.Name,
		azureclouds.USGovernmentCloud.Name,
	)
)

const (
	aksIdentifyingExtensionName = "computeAksLinuxBilling"
	// TODO: Why bother with a different CSE name for Windows?
	cseNameWindows = "windows-cse-agent-karpenter"
	cseNameLinux   = "cse-agent-karpenter"
)

// ErrorCodeForMetrics extracts a stable Azure error code for metric labeling when possible.
func ErrorCodeForMetrics(err error) string { _ = "STUB: not implemented"; return "" }

// GetManagedExtensionNames gets the names of the VM extensions managed by Karpenter.
// This is a set of 1 or 2 extensions (depending on provisionMode): aksIdentifyingExtension and (sometimes) cse.
func GetManagedExtensionNames(provisionMode string, env *auth.Environment) []string {
	_ = "STUB: not implemented"

	// Only including AKS identifying extension in the clouds it is supported in
	return nil
}

// TODO: Windows

func isAKSIdentifyingExtensionEnabled(env *auth.Environment) bool {
	_ = "STUB: not implemented"
	return false
}

type Resource = map[string]interface{}

type VirtualMachinePromise struct {
	VM       *armcompute.VirtualMachine
	WaitFunc func() error

	providerRef VMProvider
}

func (p *VirtualMachinePromise) Cleanup(ctx context.Context) error {
	_ = "STUB: not implemented"
	// This won't clean up leaked NICs if the VM doesn't exist... intentional?
	// From Delete(): "Leftover network interfaces (if any) will be cleaned by by GC controller"
	// Still, we could try to DeleteNic()?
	return nil
}

func (p *VirtualMachinePromise) Wait() error { _ = "STUB: not implemented"; return nil }

func (p *VirtualMachinePromise) GetInstanceName() string { _ = "STUB: not implemented"; return "" }

type VMProvider interface {
	BeginCreate(context.Context, *v1beta1.AKSNodeClass, *karpv1.NodeClaim, []*corecloudprovider.InstanceType) (*VirtualMachinePromise, error)
	Get(context.Context, string) (*armcompute.VirtualMachine, error)
	List(context.Context) ([]*armcompute.VirtualMachine, error)
	Delete(context.Context, string) error
	Update(context.Context, string, armcompute.VirtualMachineUpdate) error
	GetNic(context.Context, string, string) (*armnetwork.Interface, error)
	DeleteNic(context.Context, string) error
	ListNics(context.Context) ([]*armnetwork.Interface, error)
}

// assert that DefaultProvider implements Provider interface
var _ VMProvider = (*DefaultVMProvider)(nil)

type DefaultVMProvider struct {
	location                     string
	azClient                     *azclient.AZClient
	instanceTypeProvider         instancetype.Provider
	allocationStrategyProvider   allocationstrategy.Provider
	launchTemplateProvider       *launchtemplate.Provider
	loadBalancerProvider         *loadbalancer.Provider
	networkSecurityGroupProvider *networksecuritygroup.Provider
	resourceGroup                string
	subscriptionID               string
	provisionMode                string
	diskEncryptionSetID          string
	errorHandling                *offerings.ResponseErrorHandler
	env                          *auth.Environment

	vmListQuery, nicListQuery string
	deletingVMs               sets.Set[string] // tracks in-flight delete operations by VM name
	deletingVMsMu             sync.RWMutex
}

func NewDefaultVMProvider(
	azClient *azclient.AZClient,
	instanceTypeProvider instancetype.Provider,
	allocationStrategyProvider allocationstrategy.Provider,
	launchTemplateProvider *launchtemplate.Provider,
	loadBalancerProvider *loadbalancer.Provider,
	networkSecurityGroupProvider *networksecuritygroup.Provider,
	offeringsCache *cache.UnavailableOfferings,
	location string,
	resourceGroup string,
	subscriptionID string,
	provisionMode string,
	diskEncryptionSetID string,
	env *auth.Environment,
) *DefaultVMProvider {
	_ = "STUB: not implemented"
	return nil
}

// BeginCreate creates an instance given the constraints.
// instanceTypes should be sorted by priority for spot capacity type.
// Note that the returned instance may not be finished provisioning yet.
// Errors that occur on the "sync side" of the VM create, such as quota/capacity, BadRequest due to invalid user input, and similar, will have the error returned here.
// Errors that occur on the "async side" of the VM create (after the request is accepted) will be returned from VirtualMachinePromise.Wait().
func (p *DefaultVMProvider) BeginCreate(
	ctx context.Context,
	nodeClass *v1beta1.AKSNodeClass,
	nodeClaim *karpv1.NodeClaim,
	instanceTypes []*corecloudprovider.InstanceType,
) (*VirtualMachinePromise, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// There may be orphan NICs (created before promise started)
// This err block is hit only for sync failures. Async (VM provisioning) failures will be returned by the vmPromise.Wait() function

// Update updates the VM with the given updates. If Tags are specified, the tags are also updated on the associated network interface and VM extensions.
// Note that this means that this method can fail if the extensions have not been created yet. It is expected that the caller handles this and retries the update
// to propagate the tags to the extensions once they're created.
func (p *DefaultVMProvider) Update(ctx context.Context, vmName string, update armcompute.VirtualMachineUpdate) error {
	_ = "STUB: not implemented"
	return nil

	// If there are tags for other resources, do those first. This is a hedge to avoid updating the VM first which may cause us to think subsequent updates aren't needed
	// because the VM already has the updates
}

// Update NIC tags

// NIC is named the same as the VM

// Update tags on VM extensions

// TODO: This is a bit of a hack based on how this Update function is currently used.
// Currently this function will not be called by any callers until a claim has been Registered, which means that the CSE had to have succeeded.
// The aksIdentifyingExtensionName is not currently guaranteed to be on the VM though, as Karpenter could have failed over during the initial VM create
// after CSE but before aksIdentifyingExtensionName. So, for now, we just ignore NotFound errors for the aksIdentifyingExtensionName.

func (p *DefaultVMProvider) Get(ctx context.Context, vmName string) (*armcompute.VirtualMachine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *DefaultVMProvider) List(ctx context.Context) ([]*armcompute.VirtualMachine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *DefaultVMProvider) Delete(ctx context.Context, resourceName string) error {
	_ = "STUB: not implemented"
	// If there's already an in-flight delete for this VM, return immediately.
	return nil
}

// Note that 'Get' also satisfies cloudprovider.Delete contract expectation (from v1.3.0)
// of returning cloudprovider.NewNodeClaimNotFoundError if the instance is already deleted.
// This get exists to deal with the case where the operator restarted during the course of a deletion.
// With it, we may do an extra unneeded get before delete, but without it we may erroneously issue
// 2 deletes if the instance was being deleted and the operator restarted.
// Since get quota is generally higher, we prefer to check w/ get rather than issue 2 deletes.

// Check if the instance is already shutting down to reduce the number of API calls.
// Leftover network interfaces (if any) will be cleaned by by GC controller.

func (p *DefaultVMProvider) GetNic(ctx context.Context, rg, nicName string) (*armnetwork.Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListNics returns all network interfaces in the resource group that have the nodepool tag
func (p *DefaultVMProvider) ListNics(ctx context.Context) ([]*armnetwork.Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *DefaultVMProvider) DeleteNic(ctx context.Context, nicName string) error {
	_ = "STUB: not implemented"
	return nil
}

// createAKSIdentifyingExtension attaches a VM extension to identify that this VM participates in an AKS cluster
func (p *DefaultVMProvider) createAKSIdentifyingExtension(ctx context.Context, vmName string, tags map[string]*string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (p *DefaultVMProvider) createCSExtension(ctx context.Context, vmName string, cse string, isWindows bool, tags map[string]*string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *DefaultVMProvider) newNetworkInterfaceForVM(opts *createNICOptions) armnetwork.Interface {
	_ = "STUB: not implemented"
	return *new(armnetwork.Interface)
}

// AzureCNI without overlay requires secondary IPs, for pods. (These IPs are not included in backend address pools.)
// NOTE: Unlike AKS RP, this logic does not reduce secondary IP count by the number of expected hostNetwork pods, favoring simplicity instead

// E.g., aks-default-2jf98
func GenerateResourceName(nodeClaimName string) string { _ = "STUB: not implemented"; return "" }

type createNICOptions struct {
	NICName                string
	BackendPools           *loadbalancer.BackendAddressPools
	InstanceType           *corecloudprovider.InstanceType
	LaunchTemplate         *launchtemplate.Template
	NetworkPlugin          string
	NetworkPluginMode      string
	MaxPods                int32
	NetworkSecurityGroupID string
}

func (p *DefaultVMProvider) createNetworkInterface(ctx context.Context, opts *createNICOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// createVMOptions contains all the parameters needed to create a VM
type createVMOptions struct {
	VMName              string
	NicReference        string
	Zone                string
	CapacityType        string
	Location            string
	SSHPublicKey        string
	LinuxAdminUsername  string
	NodeIdentities      []string
	NodeClass           *v1beta1.AKSNodeClass
	LaunchTemplate      *launchtemplate.Template
	InstanceType        *corecloudprovider.InstanceType
	ProvisionMode       string
	UseSIG              bool
	DiskEncryptionSetID string
	NodePoolName        string
}

// newVMObject creates a new armcompute.VirtualMachine from the provided options
func newVMObject(opts *createVMOptions) *armcompute.VirtualMachine {
	_ = "STUB: not implemented"
	return nil
}

// TODO(Windows)

// TODO: I think it's safe to set this, even though it's read only

func setVMPropertiesOSDiskType(vmProperties *armcompute.VirtualMachineProperties, launchTemplate *launchtemplate.Template) {
	_ = "STUB: not implemented"
	return
}

func setVMPropertiesOSDiskEncryption(vmProperties *armcompute.VirtualMachineProperties, diskEncryptionSetID string) {
	_ = "STUB: not implemented"
	return
}

// setImageReference sets the image reference for the VM based on if we are using self hosted karpenter or the node auto provisioning addon
func setImageReference(vmProperties *armcompute.VirtualMachineProperties, imageID string, useSIG bool) {
	_ = "STUB: not implemented"
	return
}

// setVMPropertiesBillingProfile sets a default MaxPrice of -1 for Spot
func setVMPropertiesBillingProfile(vmProperties *armcompute.VirtualMachineProperties, capacityType string) {
	_ = "STUB: not implemented"
	return
}

func setVMPropertiesSecurityProfile(vmProperties *armcompute.VirtualMachineProperties, nodeClass *v1beta1.AKSNodeClass) {
	_ = "STUB: not implemented"
	return
}

type createResult struct {
	Poller *runtime.Poller[armcompute.VirtualMachinesClientCreateOrUpdateResponse]
	VM     *armcompute.VirtualMachine
}

// createVirtualMachine creates a new VM using the provided options or skips the creation of a vm if it already exists, which means opts is not guaranteed except VMName
func (p *DefaultVMProvider) createVirtualMachine(ctx context.Context, opts *createVMOptions) (*createResult, error) {
	_ = "STUB: not implemented"
	// We assume that if a vm exists, we successfully created it with the right parameters from the nodeclaims during another run before a restart.
	// there are some non-deterministic properties that may change.
	// Zones: zones are non-detrminsitic as we do a random pick out of zones on the nodeclaim that satisfy the workload requirements.
	// 	      Nodeclaim can have Requirements: Zone-1, Zone-2, Zone-3
	//        Then we pick a random zone from that list in each create call that satisfies the workload
	// UnavailableOfferingsCache: The unavailable offerings cache is used to determine if we should pick the sku, zone, or even priority.
	//        Errors for things like subscription level spot quota, SKU Quota, etc are stored in the unavailable offerings cache.
	//        So values like the SKU, Priority(Spot/On-Demand), may be different, which results in a different image, different
	//        os.CustomData.
	// If any of these properties are modified, the existing vm will return a 409 status code "PropertyChangeNotAllowed".
	// this results in create being blocked on the nodeclaim until liveness TTL is hit.
	return nil, nil
}

// If status == ok, we want to return the existing vmm

// if status != ok, and for a reason other than we did not find the vm

// beginLaunchInstance starts the launch of a VM instance.
// The returned VirtualMachinePromise must be called to gather any errors
// that are retrieved during async provisioning, as well as to complete the provisioning process.
//
//nolint:gocyclo
func (p *DefaultVMProvider) beginLaunchInstance(
	ctx context.Context,
	nodeClass *v1beta1.AKSNodeClass,
	nodeClaim *karpv1.NodeClaim,
	instanceTypes []*corecloudprovider.InstanceType,
) (*VirtualMachinePromise, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// resourceName for the NIC, VM, and Disk

// TODO: Not returning after launching this LRO because
// TODO: doing so would bypass the capacity and other errors that are currently handled by
// TODO: core pkg/controllers/nodeclaim/lifecycle/controller.go - in particular, there are metrics/events
// TODO: emitted in capacity failure cases that we probably want.

// At this point, the error is handled in provider layer (e.g., unavailable offerings cache), but not yet Karpenter core.
// Thus the error needs to be returned.
// Assuming that `HandleResponseError` already format/convert the error for such (e.g., `InsufficientCapacityError`).

// Patch the VM object to fill out a few fields that are needed later.
// This is a bit of a hack that saves us doing a GET now.
// The reason to avoid a GET is that it can fail, and if it does the future above will be lost,
// which we don't want.

// Poller is nil means the VM existed already and we're done.
// TODO: if the VM doesn't have extensions this will still happen and we will have to
// TODO: wait for the TTL for the claim to be deleted and recreated. This will most likely
// TODO: happen during Karpenter pod restart.

// At this point, the error is handled in provider layer (e.g., unavailable offerings cache), but not yet Karpenter core.
// Thus the error needs to be returned.
// Assuming that `HandleResponseError` already format/convert the error for such (e.g., `InsufficientCapacityError`).

// An error here is handled by CloudProvider create and calls vmInstanceProvider.Delete (which cleans up the azure resources)

func (p *DefaultVMProvider) applyTemplateToNic(nic *armnetwork.Interface, template *launchtemplate.Template) {
	_ = "STUB: not implemented"
	// set tags
	return
}

func (p *DefaultVMProvider) getLaunchTemplate(
	ctx context.Context,
	nodeClass *v1beta1.AKSNodeClass,
	nodeClaim *karpv1.NodeClaim,
	instanceType *corecloudprovider.InstanceType,
	capacityType string,
	placementScope string,
) (*launchtemplate.Template, error) {
	_ = "STUB: not implemented"
	// We need to get all single-valued requirement labels from the instance type and the nodeClaim to pass down to kubelet.
	// We don't just include single-value labels from the instance type because in the case where the label is NOT single-value on the instance
	// (i.e. there are options), the nodeClaim may have selected one of those options via its requirements which we want to include.
	return nil, nil
}

// These may contain restricted labels from the pod that we need to filter out; that's done in getStaticParameters.

// mustDeleteNic parameter is used to determine whether NIC deletion failure is considered an error.
// We may not want to return error of NIC cannot be deleted, as it is "by design" that NIC deletion may not be successful when VM deletion is not completed.
// NIC garbage collector is expected to handle such cases.
func (p *DefaultVMProvider) cleanupAzureResources(ctx context.Context, resourceName string, mustDeleteNic bool) error {
	_ = "STUB: not implemented"
	return nil
}

// The order here is intentional, if the VM was created successfully, then we attempt to delete the vm, the
// nic, disk and all associated resources will be removed. If the VM was not created successfully and a nic was found,
// then we attempt to delete the nic.

// Don't log NIC error here since mustDeleteNic is true (critical cleanup scenario).
// Both VM and NIC errors are returned to the caller for proper handling and logging.
// Logging here would create duplicate logs when the caller processes the joined error.

// Log NIC error here since mustDeleteNic is false (best-effort cleanup scenario).
// Because we're not returning nicErr to the caller we need to log here.
// Without this log, NIC deletion failures would be silently ignored.

// deleteVirtualMachineIfExists checks if a virtual machine exists, and if it does, we delete it with a cascading delete
func (p *DefaultVMProvider) deleteVirtualMachineIfExists(ctx context.Context, vmName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *DefaultVMProvider) deleteVirtualMachine(ctx context.Context, vmName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *DefaultVMProvider) getAKSIdentifyingExtension(tags map[string]*string) *armcompute.VirtualMachineExtension {
	_ = "STUB: not implemented"
	return nil
}

func (p *DefaultVMProvider) getCSExtension(cse string, isWindows bool, tags map[string]*string) *armcompute.VirtualMachineExtension {
	_ = "STUB: not implemented"
	return nil
}

func ConvertToVirtualMachineIdentity(nodeIdentities []string) *armcompute.VirtualMachineIdentity {
	_ = "STUB: not implemented"
	return nil
}

func GetCapacityTypeFromVM(vm *armcompute.VirtualMachine) string {
	_ = "STUB: not implemented"
	return ""
}

func GetScaleSetPriorityLabelFromVM(vm *armcompute.VirtualMachine) string {
	_ = "STUB: not implemented"
	return ""
}

func GetPriorityLabelFromVM(vm *armcompute.VirtualMachine) string {
	_ = "STUB: not implemented"
	return ""
}

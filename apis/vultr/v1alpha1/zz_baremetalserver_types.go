/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BareMetalServerObservation struct {

	// Whether an activation email will be sent when the server is ready.
	ActivationEmail *bool `json:"activationEmail,omitempty" tf:"activation_email,omitempty"`

	// The ID of the Vultr application to be installed on the server. See List Applications
	AppID *float64 `json:"appId,omitempty" tf:"app_id,omitempty"`

	// The number of CPUs available on the server.
	CPUCount *float64 `json:"cpuCount,omitempty" tf:"cpu_count,omitempty"`

	// The date the server was added to your Vultr account.
	DateCreated *string `json:"dateCreated,omitempty" tf:"date_created,omitempty"`

	// The description of the disk(s) on the server.
	Disk *string `json:"disk,omitempty" tf:"disk,omitempty"`

	// Whether the server has IPv6 networking activated.
	EnableIPv6 *bool `json:"enableIpv6,omitempty" tf:"enable_ipv6,omitempty"`

	// The server's IPv4 gateway.
	GatewayV4 *string `json:"gatewayV4,omitempty" tf:"gateway_v4,omitempty"`

	// The hostname to assign to the server.
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// ID of the server.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Vultr marketplace application to be installed on the server. See List Applications Note marketplace applications are denoted by type: marketplace and you must use the image_id not the id.
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// A label for the server.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The MAC address associated with the server.
	MacAddress *float64 `json:"macAddress,omitempty" tf:"mac_address,omitempty"`

	// The server's main IP address.
	MainIP *string `json:"mainIp,omitempty" tf:"main_ip,omitempty"`

	// The server's IPv4 netmask.
	NetmaskV4 *string `json:"netmaskV4,omitempty" tf:"netmask_v4,omitempty"`

	// The string description of the operating system installed on the server.
	Os *string `json:"os,omitempty" tf:"os,omitempty"`

	// The ID of the operating system to be installed on the server. See List OS
	OsID *float64 `json:"osId,omitempty" tf:"os_id,omitempty"`

	// The ID of the plan that you want the server to subscribe to. See List Plans
	Plan *string `json:"plan,omitempty" tf:"plan,omitempty"`

	// The amount of memory available on the server in MB.
	RAM *string `json:"ram,omitempty" tf:"ram,omitempty"`

	// The ID of the region that the server is to be created in. See List Regions
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The ID of the floating IP to use as the main IP of this server. See Reserved IPs
	ReservedIPv4 *string `json:"reservedIpv4,omitempty" tf:"reserved_ipv4,omitempty"`

	// A list of SSH key IDs to apply to the server on install (only valid for Linux/FreeBSD).
	SSHKeyIds []*string `json:"sshKeyIds,omitempty" tf:"ssh_key_ids,omitempty"`

	// The ID of the startup script you want added to the server.
	ScriptID *string `json:"scriptId,omitempty" tf:"script_id,omitempty"`

	// The ID of the Vultr snapshot that the server will restore for the initial installation. See List Snapshots
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// The status of the server's subscription.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A list of tags to apply to the servier.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// The main IPv6 network address.
	V6MainIP *string `json:"v6MainIp,omitempty" tf:"v6_main_ip,omitempty"`

	// The IPv6 subnet.
	V6Network *string `json:"v6Network,omitempty" tf:"v6_network,omitempty"`

	// The IPv6 network size in bits.
	V6NetworkSize *float64 `json:"v6NetworkSize,omitempty" tf:"v6_network_size,omitempty"`
}

type BareMetalServerParameters struct {

	// Whether an activation email will be sent when the server is ready.
	// +kubebuilder:validation:Optional
	ActivationEmail *bool `json:"activationEmail,omitempty" tf:"activation_email,omitempty"`

	// The ID of the Vultr application to be installed on the server. See List Applications
	// +kubebuilder:validation:Optional
	AppID *float64 `json:"appId,omitempty" tf:"app_id,omitempty"`

	// Whether the server has IPv6 networking activated.
	// +kubebuilder:validation:Optional
	EnableIPv6 *bool `json:"enableIpv6,omitempty" tf:"enable_ipv6,omitempty"`

	// The hostname to assign to the server.
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// The ID of the Vultr marketplace application to be installed on the server. See List Applications Note marketplace applications are denoted by type: marketplace and you must use the image_id not the id.
	// +kubebuilder:validation:Optional
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// A label for the server.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The ID of the operating system to be installed on the server. See List OS
	// +kubebuilder:validation:Optional
	OsID *float64 `json:"osId,omitempty" tf:"os_id,omitempty"`

	// The ID of the plan that you want the server to subscribe to. See List Plans
	// +kubebuilder:validation:Optional
	Plan *string `json:"plan,omitempty" tf:"plan,omitempty"`

	// The ID of the region that the server is to be created in. See List Regions
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The ID of the floating IP to use as the main IP of this server. See Reserved IPs
	// +kubebuilder:validation:Optional
	ReservedIPv4 *string `json:"reservedIpv4,omitempty" tf:"reserved_ipv4,omitempty"`

	// A list of SSH key IDs to apply to the server on install (only valid for Linux/FreeBSD).
	// +kubebuilder:validation:Optional
	SSHKeyIds []*string `json:"sshKeyIds,omitempty" tf:"ssh_key_ids,omitempty"`

	// The ID of the startup script you want added to the server.
	// +kubebuilder:validation:Optional
	ScriptID *string `json:"scriptId,omitempty" tf:"script_id,omitempty"`

	// The ID of the Vultr snapshot that the server will restore for the initial installation. See List Snapshots
	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// A list of tags to apply to the servier.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.
	// +kubebuilder:validation:Optional
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`
}

// BareMetalServerSpec defines the desired state of BareMetalServer
type BareMetalServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BareMetalServerParameters `json:"forProvider"`
}

// BareMetalServerStatus defines the observed state of BareMetalServer.
type BareMetalServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BareMetalServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BareMetalServer is the Schema for the BareMetalServers API. Provides a Vultr bare metal server resource. This can be used to create, read, modify, and delete bare metal servers on your Vultr account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vultr}
type BareMetalServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.plan)",message="plan is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.region)",message="region is a required parameter"
	Spec   BareMetalServerSpec   `json:"spec"`
	Status BareMetalServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BareMetalServerList contains a list of BareMetalServers
type BareMetalServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BareMetalServer `json:"items"`
}

// Repository type metadata.
var (
	BareMetalServer_Kind             = "BareMetalServer"
	BareMetalServer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BareMetalServer_Kind}.String()
	BareMetalServer_KindAPIVersion   = BareMetalServer_Kind + "." + CRDGroupVersion.String()
	BareMetalServer_GroupVersionKind = CRDGroupVersion.WithKind(BareMetalServer_Kind)
)

func init() {
	SchemeBuilder.Register(&BareMetalServer{}, &BareMetalServerList{})
}

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

type ConnectionPoolObservation struct {

	// The logical database to use for the new managed database connection pool.
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// The managed database ID you want to attach this connection pool to.
	DatabaseID *string `json:"databaseId,omitempty" tf:"database_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The mode to configure for the new managed database connection pool (session, transaction, statement).
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The size of the new managed database connection pool.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The database user to use for the new managed database connection pool.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type ConnectionPoolParameters struct {

	// The logical database to use for the new managed database connection pool.
	// +kubebuilder:validation:Optional
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// The managed database ID you want to attach this connection pool to.
	// +kubebuilder:validation:Optional
	DatabaseID *string `json:"databaseId,omitempty" tf:"database_id,omitempty"`

	// The mode to configure for the new managed database connection pool (session, transaction, statement).
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The size of the new managed database connection pool.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The database user to use for the new managed database connection pool.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// ConnectionPoolSpec defines the desired state of ConnectionPool
type ConnectionPoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectionPoolParameters `json:"forProvider"`
}

// ConnectionPoolStatus defines the observed state of ConnectionPool.
type ConnectionPoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectionPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionPool is the Schema for the ConnectionPools API. Provides a Vultr database connection pool resource. This can be used to create, read, modify, and delete connection pools for a PostgreSQL managed database on your Vultr account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vultr}
type ConnectionPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.database)",message="database is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.databaseId)",message="databaseId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.mode)",message="mode is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.size)",message="size is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.username)",message="username is a required parameter"
	Spec   ConnectionPoolSpec   `json:"spec"`
	Status ConnectionPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionPoolList contains a list of ConnectionPools
type ConnectionPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConnectionPool `json:"items"`
}

// Repository type metadata.
var (
	ConnectionPool_Kind             = "ConnectionPool"
	ConnectionPool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConnectionPool_Kind}.String()
	ConnectionPool_KindAPIVersion   = ConnectionPool_Kind + "." + CRDGroupVersion.String()
	ConnectionPool_GroupVersionKind = CRDGroupVersion.WithKind(ConnectionPool_Kind)
)

func init() {
	SchemeBuilder.Register(&ConnectionPool{}, &ConnectionPoolList{})
}

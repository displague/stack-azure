/*
Copyright 2019 The Crossplane Authors.

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

package v1alpha2

import (
	runtimev1alpha1 "github.com/crossplaneio/crossplane-runtime/apis/core/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/crossplaneio/crossplane-runtime/pkg/resource"
)

const (
	// OperationCreateServer is the operation type for creating a new mysql server
	OperationCreateServer = "createServer"
	// OperationCreateFirewallRules is the operation type for creating a firewall rule
	OperationCreateFirewallRules = "createFirewallRules"
)

// +kubebuilder:object:generate=false

// SQLServer represents a generic Azure SQL server.
type SQLServer interface {
	resource.Managed

	GetSpec() *SQLServerSpec
	GetStatus() *SQLServerStatus
	SetStatus(*SQLServerStatus)
}

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// +kubebuilder:object:root=true

// MysqlServer is the Schema for the instances API
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.bindingPhase"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="CLASS",type="string",JSONPath=".spec.classRef.name"
// +kubebuilder:printcolumn:name="VERSION",type="string",JSONPath=".spec.version"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
type MysqlServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SQLServerSpec   `json:"spec,omitempty"`
	Status SQLServerStatus `json:"status,omitempty"`
}

// GetSpec returns the MySQL server's spec.
func (s *MysqlServer) GetSpec() *SQLServerSpec {
	return &s.Spec
}

// GetStatus returns the MySQL server's status.
func (s *MysqlServer) GetStatus() *SQLServerStatus {
	return &s.Status
}

// SetStatus sets the MySQL server's status.
func (s *MysqlServer) SetStatus(status *SQLServerStatus) {
	s.Status = *status
}

// SetBindingPhase of this MysqlServer.
func (s *MysqlServer) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	s.Status.SetBindingPhase(p)
}

// GetBindingPhase of this MysqlServer.
func (s *MysqlServer) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return s.Status.GetBindingPhase()
}

// SetConditions of this MysqlServer.
func (s *MysqlServer) SetConditions(c ...runtimev1alpha1.Condition) {
	s.Status.SetConditions(c...)
}

// SetClaimReference of this MysqlServer.
func (s *MysqlServer) SetClaimReference(r *corev1.ObjectReference) {
	s.Spec.ClaimReference = r
}

// GetClaimReference of this MysqlServer.
func (s *MysqlServer) GetClaimReference() *corev1.ObjectReference {
	return s.Spec.ClaimReference
}

// SetNonPortableClassReference of this MysqlServer.
func (s *MysqlServer) SetNonPortableClassReference(r *corev1.ObjectReference) {
	s.Spec.NonPortableClassReference = r
}

// GetNonPortableClassReference of this MysqlServer.
func (s *MysqlServer) GetNonPortableClassReference() *corev1.ObjectReference {
	return s.Spec.NonPortableClassReference
}

// SetWriteConnectionSecretToReference of this MysqlServer.
func (s *MysqlServer) SetWriteConnectionSecretToReference(r corev1.LocalObjectReference) {
	s.Spec.WriteConnectionSecretToReference = r
}

// GetWriteConnectionSecretToReference of this MysqlServer.
func (s *MysqlServer) GetWriteConnectionSecretToReference() corev1.LocalObjectReference {
	return s.Spec.WriteConnectionSecretToReference
}

// GetReclaimPolicy of this MysqlServer.
func (s *MysqlServer) GetReclaimPolicy() runtimev1alpha1.ReclaimPolicy {
	return s.Spec.ReclaimPolicy
}

// SetReclaimPolicy of this MysqlServer.
func (s *MysqlServer) SetReclaimPolicy(p runtimev1alpha1.ReclaimPolicy) {
	s.Spec.ReclaimPolicy = p
}

// +kubebuilder:object:root=true

// MysqlServerList contains a list of MysqlServer
type MysqlServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MysqlServer `json:"items"`
}

// +kubebuilder:object:root=true

// PostgresqlServer is the Schema for the instances API
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.bindingPhase"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="CLASS",type="string",JSONPath=".spec.classRef.name"
// +kubebuilder:printcolumn:name="VERSION",type="string",JSONPath=".spec.version"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
type PostgresqlServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SQLServerSpec   `json:"spec,omitempty"`
	Status SQLServerStatus `json:"status,omitempty"`
}

// GetSpec gets the PostgreSQL server's spec.
func (s *PostgresqlServer) GetSpec() *SQLServerSpec {
	return &s.Spec
}

// GetStatus gets the PostgreSQL server's status.
func (s *PostgresqlServer) GetStatus() *SQLServerStatus {
	return &s.Status
}

// SetStatus sets the PostgreSQL server's status.
func (s *PostgresqlServer) SetStatus(status *SQLServerStatus) {
	s.Status = *status
}

// SetBindingPhase of this PostgresqlServer.
func (s *PostgresqlServer) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	s.Status.SetBindingPhase(p)
}

// GetBindingPhase of this PostgresqlServer.
func (s *PostgresqlServer) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return s.Status.GetBindingPhase()
}

// SetConditions of this PostgresqlServer.
func (s *PostgresqlServer) SetConditions(c ...runtimev1alpha1.Condition) {
	s.Status.SetConditions(c...)
}

// SetClaimReference of this PostgresqlServer.
func (s *PostgresqlServer) SetClaimReference(r *corev1.ObjectReference) {
	s.Spec.ClaimReference = r
}

// GetClaimReference of this PostgresqlServer.
func (s *PostgresqlServer) GetClaimReference() *corev1.ObjectReference {
	return s.Spec.ClaimReference
}

// SetNonPortableClassReference of this PostgresqlServer.
func (s *PostgresqlServer) SetNonPortableClassReference(r *corev1.ObjectReference) {
	s.Spec.NonPortableClassReference = r
}

// GetNonPortableClassReference of this PostgresqlServer.
func (s *PostgresqlServer) GetNonPortableClassReference() *corev1.ObjectReference {
	return s.Spec.NonPortableClassReference
}

// SetWriteConnectionSecretToReference of this PostgresqlServer.
func (s *PostgresqlServer) SetWriteConnectionSecretToReference(r corev1.LocalObjectReference) {
	s.Spec.WriteConnectionSecretToReference = r
}

// GetWriteConnectionSecretToReference of this PostgresqlServer.
func (s *PostgresqlServer) GetWriteConnectionSecretToReference() corev1.LocalObjectReference {
	return s.Spec.WriteConnectionSecretToReference
}

// GetReclaimPolicy of this PostgresqlServer.
func (s *PostgresqlServer) GetReclaimPolicy() runtimev1alpha1.ReclaimPolicy {
	return s.Spec.ReclaimPolicy
}

// SetReclaimPolicy of this PostgresqlServer.
func (s *PostgresqlServer) SetReclaimPolicy(p runtimev1alpha1.ReclaimPolicy) {
	s.Spec.ReclaimPolicy = p
}

// +kubebuilder:object:root=true

// PostgresqlServerList contains a list of PostgresqlServer
type PostgresqlServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgresqlServer `json:"items"`
}

// SQLServerClassSpecTemplate is the Schema for the resource class
type SQLServerClassSpecTemplate struct {
	runtimev1alpha1.NonPortableClassSpecTemplate `json:",inline"`
	SQLServerParameters                          `json:",inline"`
}

// All non-portable classes must implement the NonPortableClass interface.
var _ resource.NonPortableClass = &SQLServerClass{}

// +kubebuilder:object:root=true

// SQLServerClass is the Schema for the resource class
// +kubebuilder:printcolumn:name="PROVIDER-REF",type="string",JSONPath=".specTemplate.providerRef.name"
// +kubebuilder:printcolumn:name="RECLAIM-POLICY",type="string",JSONPath=".specTemplate.reclaimPolicy"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
type SQLServerClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	SpecTemplate SQLServerClassSpecTemplate `json:"specTemplate,omitempty"`
}

// GetReclaimPolicy of this PostgresqlServerClass.
func (i *SQLServerClass) GetReclaimPolicy() runtimev1alpha1.ReclaimPolicy {
	return i.SpecTemplate.ReclaimPolicy
}

// SetReclaimPolicy of this PostgresqlServerClass.
func (i *SQLServerClass) SetReclaimPolicy(p runtimev1alpha1.ReclaimPolicy) {
	i.SpecTemplate.ReclaimPolicy = p
}

// +kubebuilder:object:root=true

// SQLServerClassList contains a list of cloud memorystore resource classes.
type SQLServerClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLServerClass `json:"items"`
}

// SQLServerParameters defines the desired state of SQLServer
type SQLServerParameters struct {
	ResourceGroupName string             `json:"resourceGroupName"`
	Location          string             `json:"location"`
	PricingTier       PricingTierSpec    `json:"pricingTier"`
	StorageProfile    StorageProfileSpec `json:"storageProfile"`
	AdminLoginName    string             `json:"adminLoginName"`
	Version           string             `json:"version"`
	SSLEnforced       bool               `json:"sslEnforced,omitempty"`
}

// SQLServerSpec defines the desired state of SQLServer
type SQLServerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	SQLServerParameters          `json:",inline"`
}

// SQLServerStatus defines the observed state of SQLServer
type SQLServerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`

	State   string `json:"state,omitempty"`
	Message string `json:"message,omitempty"`

	// the external ID to identify this resource in the cloud provider
	ProviderID string `json:"providerID,omitempty"`

	// Endpoint of the MySQL Server instance used in connection strings
	Endpoint string `json:"endpoint,omitempty"`

	// RunningOperation stores any current long running operation for this instance across
	// reconciliation attempts.  This will be a serialized Azure MySQL Server API object that will
	// be used to check the status and completion of the operation during each reconciliation.
	// Once the operation has completed, this field will be cleared out.
	RunningOperation string `json:"runningOperation,omitempty"`

	// RunningOperationType is the type of the currently running operation
	RunningOperationType string `json:"runningOperationType,omitempty"`
}

// PricingTierSpec represents the performance and cost oriented properties of the server
type PricingTierSpec struct {
	Tier   string `json:"tier"`
	VCores int    `json:"vcores"`
	Family string `json:"family"`
}

// StorageProfileSpec represents storage related properties of the server
type StorageProfileSpec struct {
	StorageGB           int  `json:"storageGB"`
	BackupRetentionDays int  `json:"backupRetentionDays,omitempty"`
	GeoRedundantBackup  bool `json:"geoRedundantBackup,omitempty"`
}

// ValidMySQLVersionValues returns the valid set of engine version values.
func ValidMySQLVersionValues() []string {
	return []string{"5.6", "5.7"}
}

// ValidPostgreSQLVersionValues returns the valid set of engine version values.
func ValidPostgreSQLVersionValues() []string {
	return []string{"9.5", "9.6", "10", "10.0", "10.2"}
}

// VirtualNetworkRuleProperties defines the properties of the VirtualNetworkRule
type VirtualNetworkRuleProperties struct {
	// VirtualNetworkSubnetID - The ARM resource id of the virtual network subnet.
	VirtualNetworkSubnetID string `json:"virtualNetworkSubnetId"`
	// IgnoreMissingVnetServiceEndpoint - Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVnetServiceEndpoint bool `json:"ignoreMissingVnetServiceEndpoint,omitempty"`
}

// VirtualNetworkRuleSpec defines the desired state of VirtualNetworkRule
type VirtualNetworkRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`

	// Name - Name of the Virtual Network Rule.
	Name string `json:"name"`

	// ServerName - Name of the Virtual Network Rule's server.
	ServerName string `json:"serverName"`

	// ResourceGroupName - Name of the Virtual Network Rule's resource group.
	ResourceGroupName string `json:"resourceGroupName"`

	// VirtualNetworkRuleProperties - Resource properties.
	VirtualNetworkRuleProperties `json:"properties"`
}

// VirtualNetworkRuleStatus defines the current state of VirtualNetworkRule
type VirtualNetworkRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`

	State   string `json:"state,omitempty"`
	Message string `json:"message,omitempty"`

	// ID - Resource ID
	ID string `json:"id,omitempty"`
	// Type - Resource type.
	Type string `json:"type,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlServerVirtualNetworkRule is the Schema for the instances API
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.bindingPhase"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
type PostgresqlServerVirtualNetworkRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualNetworkRuleSpec   `json:"spec,omitempty"`
	Status VirtualNetworkRuleStatus `json:"status,omitempty"`
}

// SetBindingPhase of this PostgresqlServerVirtualNetworkRule.
func (s *PostgresqlServerVirtualNetworkRule) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	s.Status.SetBindingPhase(p)
}

// GetBindingPhase of this PostgresqlServerVirtualNetworkRule.
func (s *PostgresqlServerVirtualNetworkRule) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return s.Status.GetBindingPhase()
}

// SetConditions of this PostgresqlServerVirtualNetworkRule.
func (s *PostgresqlServerVirtualNetworkRule) SetConditions(c ...runtimev1alpha1.Condition) {
	s.Status.SetConditions(c...)
}

// SetClaimReference of this PostgresqlServerVirtualNetworkRule.
func (s *PostgresqlServerVirtualNetworkRule) SetClaimReference(r *corev1.ObjectReference) {
	s.Spec.ClaimReference = r
}

// GetClaimReference of this PostgresqlServerVirtualNetworkRule.
func (s *PostgresqlServerVirtualNetworkRule) GetClaimReference() *corev1.ObjectReference {
	return s.Spec.ClaimReference
}

// SetNonPortableClassReference of this PostgresqlServerVirtualNetworkRule.
func (s *PostgresqlServerVirtualNetworkRule) SetNonPortableClassReference(r *corev1.ObjectReference) {
	s.Spec.NonPortableClassReference = r
}

// GetNonPortableClassReference of this PostgresqlServerVirtualNetworkRule.
func (s *PostgresqlServerVirtualNetworkRule) GetNonPortableClassReference() *corev1.ObjectReference {
	return s.Spec.NonPortableClassReference
}

// SetWriteConnectionSecretToReference of this PostgresqlServerVirtualNetworkRule.
func (s *PostgresqlServerVirtualNetworkRule) SetWriteConnectionSecretToReference(r corev1.LocalObjectReference) {
	s.Spec.WriteConnectionSecretToReference = r
}

// GetWriteConnectionSecretToReference of this PostgresqlServerVirtualNetworkRule.
func (s *PostgresqlServerVirtualNetworkRule) GetWriteConnectionSecretToReference() corev1.LocalObjectReference {
	return s.Spec.WriteConnectionSecretToReference
}

// GetReclaimPolicy of this PostgresqlServerVirtualNetworkRule.
func (s *PostgresqlServerVirtualNetworkRule) GetReclaimPolicy() runtimev1alpha1.ReclaimPolicy {
	return s.Spec.ReclaimPolicy
}

// SetReclaimPolicy of this PostgresqlServerVirtualNetworkRule.
func (s *PostgresqlServerVirtualNetworkRule) SetReclaimPolicy(p runtimev1alpha1.ReclaimPolicy) {
	s.Spec.ReclaimPolicy = p
}

// +kubebuilder:object:root=true

// PostgresqlServerVirtualNetworkRuleList contains a list of PostgresqlServerVirtualNetworkRule
type PostgresqlServerVirtualNetworkRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgresqlServerVirtualNetworkRule `json:"items"`
}

// +kubebuilder:object:root=true

// MysqlServerVirtualNetworkRule is the Schema for the instances API
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.bindingPhase"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
type MysqlServerVirtualNetworkRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualNetworkRuleSpec   `json:"spec,omitempty"`
	Status VirtualNetworkRuleStatus `json:"status,omitempty"`
}

// SetBindingPhase of this MysqlServerVirtualNetworkRule.
func (s *MysqlServerVirtualNetworkRule) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	s.Status.SetBindingPhase(p)
}

// GetBindingPhase of this MysqlServerVirtualNetworkRule.
func (s *MysqlServerVirtualNetworkRule) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return s.Status.GetBindingPhase()
}

// SetConditions of this MysqlServerVirtualNetworkRule.
func (s *MysqlServerVirtualNetworkRule) SetConditions(c ...runtimev1alpha1.Condition) {
	s.Status.SetConditions(c...)
}

// SetClaimReference of this MysqlServerVirtualNetworkRule.
func (s *MysqlServerVirtualNetworkRule) SetClaimReference(r *corev1.ObjectReference) {
	s.Spec.ClaimReference = r
}

// GetClaimReference of this MysqlServerVirtualNetworkRule.
func (s *MysqlServerVirtualNetworkRule) GetClaimReference() *corev1.ObjectReference {
	return s.Spec.ClaimReference
}

// SetNonPortableClassReference of this MysqlServerVirtualNetworkRule.
func (s *MysqlServerVirtualNetworkRule) SetNonPortableClassReference(r *corev1.ObjectReference) {
	s.Spec.NonPortableClassReference = r
}

// GetNonPortableClassReference of this MysqlServerVirtualNetworkRule.
func (s *MysqlServerVirtualNetworkRule) GetNonPortableClassReference() *corev1.ObjectReference {
	return s.Spec.NonPortableClassReference
}

// SetWriteConnectionSecretToReference of this MysqlServerVirtualNetworkRule.
func (s *MysqlServerVirtualNetworkRule) SetWriteConnectionSecretToReference(r corev1.LocalObjectReference) {
	s.Spec.WriteConnectionSecretToReference = r
}

// GetWriteConnectionSecretToReference of this MysqlServerVirtualNetworkRule.
func (s *MysqlServerVirtualNetworkRule) GetWriteConnectionSecretToReference() corev1.LocalObjectReference {
	return s.Spec.WriteConnectionSecretToReference
}

// GetReclaimPolicy of this MysqlServerVirtualNetworkRule.
func (s *MysqlServerVirtualNetworkRule) GetReclaimPolicy() runtimev1alpha1.ReclaimPolicy {
	return s.Spec.ReclaimPolicy
}

// SetReclaimPolicy of this MysqlServerVirtualNetworkRule.
func (s *MysqlServerVirtualNetworkRule) SetReclaimPolicy(p runtimev1alpha1.ReclaimPolicy) {
	s.Spec.ReclaimPolicy = p
}

// +kubebuilder:object:root=true

// MysqlServerVirtualNetworkRuleList contains a list of MysqlServerVirtualNetworkRule
type MysqlServerVirtualNetworkRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MysqlServerVirtualNetworkRule `json:"items"`
}

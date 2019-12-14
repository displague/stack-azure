// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha3

import (
	runtime "k8s.io/apimachinery/pkg/runtime"

	"github.com/crossplaneio/stack-azure/apis/database/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLServerVirtualNetworkRule) DeepCopyInto(out *MySQLServerVirtualNetworkRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLServerVirtualNetworkRule.
func (in *MySQLServerVirtualNetworkRule) DeepCopy() *MySQLServerVirtualNetworkRule {
	if in == nil {
		return nil
	}
	out := new(MySQLServerVirtualNetworkRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLServerVirtualNetworkRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLServerVirtualNetworkRuleList) DeepCopyInto(out *MySQLServerVirtualNetworkRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MySQLServerVirtualNetworkRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLServerVirtualNetworkRuleList.
func (in *MySQLServerVirtualNetworkRuleList) DeepCopy() *MySQLServerVirtualNetworkRuleList {
	if in == nil {
		return nil
	}
	out := new(MySQLServerVirtualNetworkRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLServerVirtualNetworkRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLVirtualNetworkRuleSpec) DeepCopyInto(out *MySQLVirtualNetworkRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	if in.ServerNameRef != nil {
		in, out := &in.ServerNameRef, &out.ServerNameRef
		*out = new(v1beta1.MySQLServerNameReferencer)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(ResourceGroupNameReferencerForVirtualNetworkRule)
		**out = **in
	}
	in.VirtualNetworkRuleProperties.DeepCopyInto(&out.VirtualNetworkRuleProperties)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLVirtualNetworkRuleSpec.
func (in *MySQLVirtualNetworkRuleSpec) DeepCopy() *MySQLVirtualNetworkRuleSpec {
	if in == nil {
		return nil
	}
	out := new(MySQLVirtualNetworkRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLServerVirtualNetworkRule) DeepCopyInto(out *PostgreSQLServerVirtualNetworkRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLServerVirtualNetworkRule.
func (in *PostgreSQLServerVirtualNetworkRule) DeepCopy() *PostgreSQLServerVirtualNetworkRule {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLServerVirtualNetworkRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgreSQLServerVirtualNetworkRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLServerVirtualNetworkRuleList) DeepCopyInto(out *PostgreSQLServerVirtualNetworkRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgreSQLServerVirtualNetworkRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLServerVirtualNetworkRuleList.
func (in *PostgreSQLServerVirtualNetworkRuleList) DeepCopy() *PostgreSQLServerVirtualNetworkRuleList {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLServerVirtualNetworkRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgreSQLServerVirtualNetworkRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLVirtualNetworkRuleSpec) DeepCopyInto(out *PostgreSQLVirtualNetworkRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	if in.ServerNameRef != nil {
		in, out := &in.ServerNameRef, &out.ServerNameRef
		*out = new(v1beta1.PostgreSQLServerNameReferencer)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(ResourceGroupNameReferencerForVirtualNetworkRule)
		**out = **in
	}
	in.VirtualNetworkRuleProperties.DeepCopyInto(&out.VirtualNetworkRuleProperties)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLVirtualNetworkRuleSpec.
func (in *PostgreSQLVirtualNetworkRuleSpec) DeepCopy() *PostgreSQLVirtualNetworkRuleSpec {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLVirtualNetworkRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroupNameReferencerForVirtualNetworkRule) DeepCopyInto(out *ResourceGroupNameReferencerForVirtualNetworkRule) {
	*out = *in
	out.ResourceGroupNameReferencer = in.ResourceGroupNameReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroupNameReferencerForVirtualNetworkRule.
func (in *ResourceGroupNameReferencerForVirtualNetworkRule) DeepCopy() *ResourceGroupNameReferencerForVirtualNetworkRule {
	if in == nil {
		return nil
	}
	out := new(ResourceGroupNameReferencerForVirtualNetworkRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerNameReferencerForMySQLServerVirtualNetworkRule) DeepCopyInto(out *ServerNameReferencerForMySQLServerVirtualNetworkRule) {
	*out = *in
	out.MySQLServerNameReferencer = in.MySQLServerNameReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerNameReferencerForMySQLServerVirtualNetworkRule.
func (in *ServerNameReferencerForMySQLServerVirtualNetworkRule) DeepCopy() *ServerNameReferencerForMySQLServerVirtualNetworkRule {
	if in == nil {
		return nil
	}
	out := new(ServerNameReferencerForMySQLServerVirtualNetworkRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerNameReferencerForPostgreSQLServerVirtualNetworkRule) DeepCopyInto(out *ServerNameReferencerForPostgreSQLServerVirtualNetworkRule) {
	*out = *in
	out.PostgreSQLServerNameReferencer = in.PostgreSQLServerNameReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerNameReferencerForPostgreSQLServerVirtualNetworkRule.
func (in *ServerNameReferencerForPostgreSQLServerVirtualNetworkRule) DeepCopy() *ServerNameReferencerForPostgreSQLServerVirtualNetworkRule {
	if in == nil {
		return nil
	}
	out := new(ServerNameReferencerForPostgreSQLServerVirtualNetworkRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetIDReferencerForVirtualNetworkRule) DeepCopyInto(out *SubnetIDReferencerForVirtualNetworkRule) {
	*out = *in
	out.SubnetIDReferencer = in.SubnetIDReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetIDReferencerForVirtualNetworkRule.
func (in *SubnetIDReferencerForVirtualNetworkRule) DeepCopy() *SubnetIDReferencerForVirtualNetworkRule {
	if in == nil {
		return nil
	}
	out := new(SubnetIDReferencerForVirtualNetworkRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkRuleProperties) DeepCopyInto(out *VirtualNetworkRuleProperties) {
	*out = *in
	if in.VirtualNetworkSubnetIDRef != nil {
		in, out := &in.VirtualNetworkSubnetIDRef, &out.VirtualNetworkSubnetIDRef
		*out = new(SubnetIDReferencerForVirtualNetworkRule)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkRuleProperties.
func (in *VirtualNetworkRuleProperties) DeepCopy() *VirtualNetworkRuleProperties {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkRuleProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkRuleStatus) DeepCopyInto(out *VirtualNetworkRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkRuleStatus.
func (in *VirtualNetworkRuleStatus) DeepCopy() *VirtualNetworkRuleStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationExecution) DeepCopyInto(out *AuthenticationExecution) {
	*out = *in
	if in.AuthenticatorConfig != nil {
		in, out := &in.AuthenticatorConfig, &out.AuthenticatorConfig
		*out = new(AuthenticatorConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationExecution.
func (in *AuthenticationExecution) DeepCopy() *AuthenticationExecution {
	if in == nil {
		return nil
	}
	out := new(AuthenticationExecution)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticatorConfig) DeepCopyInto(out *AuthenticatorConfig) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticatorConfig.
func (in *AuthenticatorConfig) DeepCopy() *AuthenticatorConfig {
	if in == nil {
		return nil
	}
	out := new(AuthenticatorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchRole) DeepCopyInto(out *BatchRole) {
	*out = *in
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.Composites != nil {
		in, out := &in.Composites, &out.Composites
		*out = make([]Composite, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchRole.
func (in *BatchRole) DeepCopy() *BatchRole {
	if in == nil {
		return nil
	}
	out := new(BatchRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientRole) DeepCopyInto(out *ClientRole) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientRole.
func (in *ClientRole) DeepCopy() *ClientRole {
	if in == nil {
		return nil
	}
	out := new(ClientRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Composite) DeepCopyInto(out *Composite) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Composite.
func (in *Composite) DeepCopy() *Composite {
	if in == nil {
		return nil
	}
	out := new(Composite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Keycloak) DeepCopyInto(out *Keycloak) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Keycloak.
func (in *Keycloak) DeepCopy() *Keycloak {
	if in == nil {
		return nil
	}
	out := new(Keycloak)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Keycloak) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakAuthFlow) DeepCopyInto(out *KeycloakAuthFlow) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakAuthFlow.
func (in *KeycloakAuthFlow) DeepCopy() *KeycloakAuthFlow {
	if in == nil {
		return nil
	}
	out := new(KeycloakAuthFlow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakAuthFlow) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakAuthFlowList) DeepCopyInto(out *KeycloakAuthFlowList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeycloakAuthFlow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakAuthFlowList.
func (in *KeycloakAuthFlowList) DeepCopy() *KeycloakAuthFlowList {
	if in == nil {
		return nil
	}
	out := new(KeycloakAuthFlowList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakAuthFlowList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakAuthFlowSpec) DeepCopyInto(out *KeycloakAuthFlowSpec) {
	*out = *in
	if in.AuthenticationExecutions != nil {
		in, out := &in.AuthenticationExecutions, &out.AuthenticationExecutions
		*out = make([]AuthenticationExecution, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakAuthFlowSpec.
func (in *KeycloakAuthFlowSpec) DeepCopy() *KeycloakAuthFlowSpec {
	if in == nil {
		return nil
	}
	out := new(KeycloakAuthFlowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakAuthFlowStatus) DeepCopyInto(out *KeycloakAuthFlowStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakAuthFlowStatus.
func (in *KeycloakAuthFlowStatus) DeepCopy() *KeycloakAuthFlowStatus {
	if in == nil {
		return nil
	}
	out := new(KeycloakAuthFlowStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakClient) DeepCopyInto(out *KeycloakClient) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakClient.
func (in *KeycloakClient) DeepCopy() *KeycloakClient {
	if in == nil {
		return nil
	}
	out := new(KeycloakClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakClient) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakClientList) DeepCopyInto(out *KeycloakClientList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeycloakClient, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakClientList.
func (in *KeycloakClientList) DeepCopy() *KeycloakClientList {
	if in == nil {
		return nil
	}
	out := new(KeycloakClientList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakClientList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakClientSpec) DeepCopyInto(out *KeycloakClientSpec) {
	*out = *in
	if in.RealmRoles != nil {
		in, out := &in.RealmRoles, &out.RealmRoles
		*out = new([]RealmRole)
		if **in != nil {
			in, out := *in, *out
			*out = make([]RealmRole, len(*in))
			copy(*out, *in)
		}
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ClientRoles != nil {
		in, out := &in.ClientRoles, &out.ClientRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ProtocolMappers != nil {
		in, out := &in.ProtocolMappers, &out.ProtocolMappers
		*out = new([]ProtocolMapper)
		if **in != nil {
			in, out := *in, *out
			*out = make([]ProtocolMapper, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(ServiceAccount)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakClientSpec.
func (in *KeycloakClientSpec) DeepCopy() *KeycloakClientSpec {
	if in == nil {
		return nil
	}
	out := new(KeycloakClientSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakClientStatus) DeepCopyInto(out *KeycloakClientStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakClientStatus.
func (in *KeycloakClientStatus) DeepCopy() *KeycloakClientStatus {
	if in == nil {
		return nil
	}
	out := new(KeycloakClientStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakList) DeepCopyInto(out *KeycloakList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Keycloak, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakList.
func (in *KeycloakList) DeepCopy() *KeycloakList {
	if in == nil {
		return nil
	}
	out := new(KeycloakList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealm) DeepCopyInto(out *KeycloakRealm) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealm.
func (in *KeycloakRealm) DeepCopy() *KeycloakRealm {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakRealm) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealmGroup) DeepCopyInto(out *KeycloakRealmGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealmGroup.
func (in *KeycloakRealmGroup) DeepCopy() *KeycloakRealmGroup {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealmGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakRealmGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealmGroupList) DeepCopyInto(out *KeycloakRealmGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeycloakRealmGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealmGroupList.
func (in *KeycloakRealmGroupList) DeepCopy() *KeycloakRealmGroupList {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealmGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakRealmGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealmGroupSpec) DeepCopyInto(out *KeycloakRealmGroupSpec) {
	*out = *in
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.Access != nil {
		in, out := &in.Access, &out.Access
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RealmRoles != nil {
		in, out := &in.RealmRoles, &out.RealmRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubGroups != nil {
		in, out := &in.SubGroups, &out.SubGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ClientRoles != nil {
		in, out := &in.ClientRoles, &out.ClientRoles
		*out = make([]ClientRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealmGroupSpec.
func (in *KeycloakRealmGroupSpec) DeepCopy() *KeycloakRealmGroupSpec {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealmGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealmGroupStatus) DeepCopyInto(out *KeycloakRealmGroupStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealmGroupStatus.
func (in *KeycloakRealmGroupStatus) DeepCopy() *KeycloakRealmGroupStatus {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealmGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealmList) DeepCopyInto(out *KeycloakRealmList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeycloakRealm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealmList.
func (in *KeycloakRealmList) DeepCopy() *KeycloakRealmList {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealmList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakRealmList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealmRole) DeepCopyInto(out *KeycloakRealmRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealmRole.
func (in *KeycloakRealmRole) DeepCopy() *KeycloakRealmRole {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealmRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakRealmRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealmRoleBatch) DeepCopyInto(out *KeycloakRealmRoleBatch) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealmRoleBatch.
func (in *KeycloakRealmRoleBatch) DeepCopy() *KeycloakRealmRoleBatch {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealmRoleBatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakRealmRoleBatch) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealmRoleBatchList) DeepCopyInto(out *KeycloakRealmRoleBatchList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeycloakRealmRoleBatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealmRoleBatchList.
func (in *KeycloakRealmRoleBatchList) DeepCopy() *KeycloakRealmRoleBatchList {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealmRoleBatchList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakRealmRoleBatchList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealmRoleBatchSpec) DeepCopyInto(out *KeycloakRealmRoleBatchSpec) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]BatchRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealmRoleBatchSpec.
func (in *KeycloakRealmRoleBatchSpec) DeepCopy() *KeycloakRealmRoleBatchSpec {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealmRoleBatchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealmRoleBatchStatus) DeepCopyInto(out *KeycloakRealmRoleBatchStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealmRoleBatchStatus.
func (in *KeycloakRealmRoleBatchStatus) DeepCopy() *KeycloakRealmRoleBatchStatus {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealmRoleBatchStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealmRoleList) DeepCopyInto(out *KeycloakRealmRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeycloakRealmRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealmRoleList.
func (in *KeycloakRealmRoleList) DeepCopy() *KeycloakRealmRoleList {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealmRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakRealmRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealmRoleSpec) DeepCopyInto(out *KeycloakRealmRoleSpec) {
	*out = *in
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.Composites != nil {
		in, out := &in.Composites, &out.Composites
		*out = make([]Composite, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealmRoleSpec.
func (in *KeycloakRealmRoleSpec) DeepCopy() *KeycloakRealmRoleSpec {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealmRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealmRoleStatus) DeepCopyInto(out *KeycloakRealmRoleStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealmRoleStatus.
func (in *KeycloakRealmRoleStatus) DeepCopy() *KeycloakRealmRoleStatus {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealmRoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealmSpec) DeepCopyInto(out *KeycloakRealmSpec) {
	*out = *in
	if in.SsoRealmEnabled != nil {
		in, out := &in.SsoRealmEnabled, &out.SsoRealmEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SsoAutoRedirectEnabled != nil {
		in, out := &in.SsoAutoRedirectEnabled, &out.SsoAutoRedirectEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]User, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SSORealmMappers != nil {
		in, out := &in.SSORealmMappers, &out.SSORealmMappers
		*out = new([]SSORealmMapper)
		if **in != nil {
			in, out := *in, *out
			*out = make([]SSORealmMapper, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.BrowserFlow != nil {
		in, out := &in.BrowserFlow, &out.BrowserFlow
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealmSpec.
func (in *KeycloakRealmSpec) DeepCopy() *KeycloakRealmSpec {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealmSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealmStatus) DeepCopyInto(out *KeycloakRealmStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealmStatus.
func (in *KeycloakRealmStatus) DeepCopy() *KeycloakRealmStatus {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealmStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakSpec) DeepCopyInto(out *KeycloakSpec) {
	*out = *in
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]User, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InstallMainRealm != nil {
		in, out := &in.InstallMainRealm, &out.InstallMainRealm
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakSpec.
func (in *KeycloakSpec) DeepCopy() *KeycloakSpec {
	if in == nil {
		return nil
	}
	out := new(KeycloakSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakStatus) DeepCopyInto(out *KeycloakStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakStatus.
func (in *KeycloakStatus) DeepCopy() *KeycloakStatus {
	if in == nil {
		return nil
	}
	out := new(KeycloakStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtocolMapper) DeepCopyInto(out *ProtocolMapper) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtocolMapper.
func (in *ProtocolMapper) DeepCopy() *ProtocolMapper {
	if in == nil {
		return nil
	}
	out := new(ProtocolMapper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RealmRole) DeepCopyInto(out *RealmRole) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RealmRole.
func (in *RealmRole) DeepCopy() *RealmRole {
	if in == nil {
		return nil
	}
	out := new(RealmRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSORealmMapper) DeepCopyInto(out *SSORealmMapper) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSORealmMapper.
func (in *SSORealmMapper) DeepCopy() *SSORealmMapper {
	if in == nil {
		return nil
	}
	out := new(SSORealmMapper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccount) DeepCopyInto(out *ServiceAccount) {
	*out = *in
	if in.RealmRoles != nil {
		in, out := &in.RealmRoles, &out.RealmRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ClientRoles != nil {
		in, out := &in.ClientRoles, &out.ClientRoles
		*out = make([]ClientRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccount.
func (in *ServiceAccount) DeepCopy() *ServiceAccount {
	if in == nil {
		return nil
	}
	out := new(ServiceAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User) DeepCopyInto(out *User) {
	*out = *in
	if in.RealmRoles != nil {
		in, out := &in.RealmRoles, &out.RealmRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User.
func (in *User) DeepCopy() *User {
	if in == nil {
		return nil
	}
	out := new(User)
	in.DeepCopyInto(out)
	return out
}

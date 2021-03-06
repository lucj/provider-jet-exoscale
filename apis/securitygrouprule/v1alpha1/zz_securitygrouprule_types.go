/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SecurityGroupRuleObservation struct {
}

type SecurityGroupRuleParameters struct {

	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	EndPort *int64 `json:"endPort,omitempty" tf:"end_port,omitempty"`

	// +kubebuilder:validation:Optional
	IcmpCode *int64 `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

	// +kubebuilder:validation:Optional
	IcmpType *int64 `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroup *string `json:"securityGroup,omitempty" tf:"security_group,omitempty"`

	// +crossplane:generate:reference:type=github.com/lucj/provider-jet-exoscale/apis/securitygroup/v1alpha1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupIdRef
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupIdSelector
	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupIdRef *v1.Reference `json:"securityGroupIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityGroupIdSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	StartPort *int64 `json:"startPort,omitempty" tf:"start_port,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	UserSecurityGroup *string `json:"userSecurityGroup,omitempty" tf:"user_security_group,omitempty"`

	// +crossplane:generate:reference:type=github.com/lucj/provider-jet-exoscale/apis/securitygroup/v1alpha1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=UserSecurityGroupIdRef
	// +crossplane:generate:reference:selectorFieldName=UserSecurityGroupIdSelector
	// +kubebuilder:validation:Optional
	UserSecurityGroupID *string `json:"userSecurityGroupId,omitempty" tf:"user_security_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	UserSecurityGroupIdRef *v1.Reference `json:"userSecurityGroupIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	UserSecurityGroupIdSelector *v1.Selector `json:"userSecurityGroupIdSelector,omitempty" tf:"-"`
}

// SecurityGroupRuleSpec defines the desired state of SecurityGroupRule
type SecurityGroupRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityGroupRuleParameters `json:"forProvider"`
}

// SecurityGroupRuleStatus defines the observed state of SecurityGroupRule.
type SecurityGroupRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityGroupRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupRule is the Schema for the SecurityGroupRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,exoscalejet}
type SecurityGroupRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityGroupRuleSpec   `json:"spec"`
	Status            SecurityGroupRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupRuleList contains a list of SecurityGroupRules
type SecurityGroupRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroupRule `json:"items"`
}

// Repository type metadata.
var (
	SecurityGroupRule_Kind             = "SecurityGroupRule"
	SecurityGroupRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityGroupRule_Kind}.String()
	SecurityGroupRule_KindAPIVersion   = SecurityGroupRule_Kind + "." + CRDGroupVersion.String()
	SecurityGroupRule_GroupVersionKind = CRDGroupVersion.WithKind(SecurityGroupRule_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityGroupRule{}, &SecurityGroupRuleList{})
}

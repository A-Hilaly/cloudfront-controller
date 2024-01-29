// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ResponseHeadersPolicySpec defines the desired state of ResponseHeadersPolicy.
//
// A response headers policy.
//
// A response headers policy contains information about a set of HTTP response
// headers.
//
// After you create a response headers policy, you can use its ID to attach
// it to one or more cache behaviors in a CloudFront distribution. When it's
// attached to a cache behavior, the response headers policy affects the HTTP
// headers that CloudFront includes in HTTP responses to requests that match
// the cache behavior. CloudFront adds or removes response headers according
// to the configuration of the response headers policy.
//
// For more information, see Adding or removing HTTP headers in CloudFront responses
// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/modifying-response-headers.html)
// in the Amazon CloudFront Developer Guide.
type ResponseHeadersPolicySpec struct {

	// Contains metadata about the response headers policy, and a set of configurations
	// that specify the HTTP headers.
	// +kubebuilder:validation:Required
	ResponseHeadersPolicyConfig *ResponseHeadersPolicyConfig `json:"responseHeadersPolicyConfig"`
}

// ResponseHeadersPolicyStatus defines the observed state of ResponseHeadersPolicy
type ResponseHeadersPolicyStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// +kubebuilder:validation:Optional
	ETag *string `json:"eTag,omitempty"`
	// The identifier for the response headers policy.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty"`
	// The date and time when the response headers policy was last modified.
	// +kubebuilder:validation:Optional
	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`
}

// ResponseHeadersPolicy is the Schema for the ResponseHeadersPolicies API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type ResponseHeadersPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResponseHeadersPolicySpec   `json:"spec,omitempty"`
	Status            ResponseHeadersPolicyStatus `json:"status,omitempty"`
}

// ResponseHeadersPolicyList contains a list of ResponseHeadersPolicy
// +kubebuilder:object:root=true
type ResponseHeadersPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResponseHeadersPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ResponseHeadersPolicy{}, &ResponseHeadersPolicyList{})
}
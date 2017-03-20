/*
Copyright 2016 The Kubernetes Authors.

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

package v1alpha1

import (
	"github.com/OpenPlatformSDN/client-go/pkg/api/v1"
	metav1 "github.com/OpenPlatformSDN/client-go/pkg/apis/meta/v1"
)

// +genclient=true
// +nonNamespaced=true

// Describes a certificate signing request
type CertificateSigningRequest struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// The certificate request itself and any additional information.
	// +optional
	Spec CertificateSigningRequestSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Derived information about the request.
	// +optional
	Status CertificateSigningRequestStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// This information is immutable after the request is created. Only the Request
// and ExtraInfo fields can be set on creation, other fields are derived by
// Kubernetes and cannot be modified by users.
type CertificateSigningRequestSpec struct {
	// Base64-encoded PKCS#10 CSR data
	Request []byte `json:"request" protobuf:"bytes,1,opt,name=request"`

	// Information about the requesting user (if relevant)
	// See user.Info interface for details
	// +optional
	Username string `json:"username,omitempty" protobuf:"bytes,2,opt,name=username"`
	// +optional
	UID string `json:"uid,omitempty" protobuf:"bytes,3,opt,name=uid"`
	// +optional
	Groups []string `json:"groups,omitempty" protobuf:"bytes,4,rep,name=groups"`
}

type CertificateSigningRequestStatus struct {
	// Conditions applied to the request, such as approval or denial.
	// +optional
	Conditions []CertificateSigningRequestCondition `json:"conditions,omitempty" protobuf:"bytes,1,rep,name=conditions"`

	// If request was approved, the controller will place the issued certificate here.
	// +optional
	Certificate []byte `json:"certificate,omitempty" protobuf:"bytes,2,opt,name=certificate"`
}

type RequestConditionType string

// These are the possible conditions for a certificate request.
const (
	CertificateApproved RequestConditionType = "Approved"
	CertificateDenied   RequestConditionType = "Denied"
)

type CertificateSigningRequestCondition struct {
	// request approval state, currently Approved or Denied.
	Type RequestConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=RequestConditionType"`
	// brief reason for the request state
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,2,opt,name=reason"`
	// human readable message with details about the request state
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,3,opt,name=message"`
	// timestamp for the last update to this condition
	// +optional
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty" protobuf:"bytes,4,opt,name=lastUpdateTime"`
}

type CertificateSigningRequestList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []CertificateSigningRequest `json:"items" protobuf:"bytes,2,rep,name=items"`
}

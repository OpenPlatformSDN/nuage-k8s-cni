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

package storage

import (
	"github.com/OpenPlatformSDN/client-go/pkg/api"
	metav1 "github.com/OpenPlatformSDN/client-go/pkg/apis/meta/v1"
)

// +genclient=true
// +nonNamespaced=true

// StorageClass describes a named "class" of storage offered in a cluster.
// Different classes might map to quality-of-service levels, or to backup policies,
// or to arbitrary policies determined by the cluster administrators.  Kubernetes
// itself is unopinionated about what classes represent.  This concept is sometimes
// called "profiles" in other storage systems.
// The name of a StorageClass object is significant, and is how users can request a particular class.
type StorageClass struct {
	metav1.TypeMeta
	// +optional
	api.ObjectMeta

	// provisioner is the driver expected to handle this StorageClass.
	// This is an optionally-prefixed name, like a label key.
	// For example: "kubernetes.io/gce-pd" or "kubernetes.io/aws-ebs".
	// This value may not be empty.
	Provisioner string

	// parameters holds parameters for the provisioner.
	// These values are opaque to the  system and are passed directly
	// to the provisioner.  The only validation done on keys is that they are
	// not empty.  The maximum number of parameters is
	// 512, with a cumulative max size of 256K
	// +optional
	Parameters map[string]string
}

// StorageClassList is a collection of storage classes.
type StorageClassList struct {
	metav1.TypeMeta
	// Standard list metadata
	// More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
	// +optional
	metav1.ListMeta

	// Items is the list of StorageClasses
	Items []StorageClass
}

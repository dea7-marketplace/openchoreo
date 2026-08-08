// Copyright 2026 The OpenChoreo Authors
// SPDX-License-Identifier: Apache-2.0

package renderedrelease

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	openchoreov1alpha1 "github.com/openchoreo/openchoreo/api/v1alpha1"
)

// resourceIdentity uniquely identifies one rendered Kubernetes object. Resource IDs
// alone are insufficient because controllers commonly propagate a parent's labels to
// generated children and a rendered entry may replace an immutable object by name.
type resourceIdentity struct {
	Group     string
	Version   string
	Kind      string
	Namespace string
	Name      string
}

func identityForObject(obj *unstructured.Unstructured) resourceIdentity {
	gvk := obj.GroupVersionKind()
	return resourceIdentity{
		Group:     gvk.Group,
		Version:   gvk.Version,
		Kind:      gvk.Kind,
		Namespace: obj.GetNamespace(),
		Name:      obj.GetName(),
	}
}

func identityForStatus(status openchoreov1alpha1.RenderedManifestStatus) resourceIdentity {
	return resourceIdentity{
		Group:     status.Group,
		Version:   status.Version,
		Kind:      status.Kind,
		Namespace: status.Namespace,
		Name:      status.Name,
	}
}

func hasControllerOwner(obj metav1.Object) bool {
	return metav1.GetControllerOf(obj) != nil
}

/*
Copyright The OpenShift authors.

SPDX-License-Identifier: Apache-2.0
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/openshift/csi-driver-shared-resource/pkg/api/sharedresource/v1alpha1"
	scheme "github.com/openshift/csi-driver-shared-resource/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SharesGetter has a method to return a ShareInterface.
// A group's client should implement this interface.
type SharesGetter interface {
	Shares() ShareInterface
}

// ShareInterface has methods to work with Share resources.
type ShareInterface interface {
	Create(ctx context.Context, share *v1alpha1.Share, opts v1.CreateOptions) (*v1alpha1.Share, error)
	Update(ctx context.Context, share *v1alpha1.Share, opts v1.UpdateOptions) (*v1alpha1.Share, error)
	UpdateStatus(ctx context.Context, share *v1alpha1.Share, opts v1.UpdateOptions) (*v1alpha1.Share, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Share, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ShareList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Share, err error)
	ShareExpansion
}

// shares implements ShareInterface
type shares struct {
	client rest.Interface
}

// newShares returns a Shares
func newShares(c *SharedresourceV1alpha1Client) *shares {
	return &shares{
		client: c.RESTClient(),
	}
}

// Get takes name of the share, and returns the corresponding share object, and an error if there is any.
func (c *shares) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Share, err error) {
	result = &v1alpha1.Share{}
	err = c.client.Get().
		Resource("shares").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Shares that match those selectors.
func (c *shares) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ShareList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ShareList{}
	err = c.client.Get().
		Resource("shares").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested shares.
func (c *shares) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("shares").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a share and creates it.  Returns the server's representation of the share, and an error, if there is any.
func (c *shares) Create(ctx context.Context, share *v1alpha1.Share, opts v1.CreateOptions) (result *v1alpha1.Share, err error) {
	result = &v1alpha1.Share{}
	err = c.client.Post().
		Resource("shares").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(share).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a share and updates it. Returns the server's representation of the share, and an error, if there is any.
func (c *shares) Update(ctx context.Context, share *v1alpha1.Share, opts v1.UpdateOptions) (result *v1alpha1.Share, err error) {
	result = &v1alpha1.Share{}
	err = c.client.Put().
		Resource("shares").
		Name(share.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(share).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *shares) UpdateStatus(ctx context.Context, share *v1alpha1.Share, opts v1.UpdateOptions) (result *v1alpha1.Share, err error) {
	result = &v1alpha1.Share{}
	err = c.client.Put().
		Resource("shares").
		Name(share.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(share).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the share and deletes it. Returns an error if one occurs.
func (c *shares) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("shares").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *shares) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("shares").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched share.
func (c *shares) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Share, err error) {
	result = &v1alpha1.Share{}
	err = c.client.Patch(pt).
		Resource("shares").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/ca-gip/kotary/pkg/apis/ca-gip/v1"
	scheme "github.com/ca-gip/kotary/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ResourceQuotaClaimsGetter has a method to return a ResourceQuotaClaimInterface.
// A group's client should implement this interface.
type ResourceQuotaClaimsGetter interface {
	ResourceQuotaClaims(namespace string) ResourceQuotaClaimInterface
}

// ResourceQuotaClaimInterface has methods to work with ResourceQuotaClaim resources.
type ResourceQuotaClaimInterface interface {
	Create(*v1.ResourceQuotaClaim) (*v1.ResourceQuotaClaim, error)
	Update(*v1.ResourceQuotaClaim) (*v1.ResourceQuotaClaim, error)
	UpdateStatus(*v1.ResourceQuotaClaim) (*v1.ResourceQuotaClaim, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ResourceQuotaClaim, error)
	List(opts metav1.ListOptions) (*v1.ResourceQuotaClaimList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ResourceQuotaClaim, err error)
	ResourceQuotaClaimExpansion
}

// resourceQuotaClaims implements ResourceQuotaClaimInterface
type resourceQuotaClaims struct {
	client rest.Interface
	ns     string
}

// newResourceQuotaClaims returns a ResourceQuotaClaims
func newResourceQuotaClaims(c *CagipV1Client, namespace string) *resourceQuotaClaims {
	return &resourceQuotaClaims{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the resourceQuotaClaim, and returns the corresponding resourceQuotaClaim object, and an error if there is any.
func (c *resourceQuotaClaims) Get(name string, options metav1.GetOptions) (result *v1.ResourceQuotaClaim, err error) {
	result = &v1.ResourceQuotaClaim{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourcequotaclaims").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ResourceQuotaClaims that match those selectors.
func (c *resourceQuotaClaims) List(opts metav1.ListOptions) (result *v1.ResourceQuotaClaimList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ResourceQuotaClaimList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourcequotaclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested resourceQuotaClaims.
func (c *resourceQuotaClaims) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("resourcequotaclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a resourceQuotaClaim and creates it.  Returns the server's representation of the resourceQuotaClaim, and an error, if there is any.
func (c *resourceQuotaClaims) Create(resourceQuotaClaim *v1.ResourceQuotaClaim) (result *v1.ResourceQuotaClaim, err error) {
	result = &v1.ResourceQuotaClaim{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("resourcequotaclaims").
		Body(resourceQuotaClaim).
		Do().
		Into(result)
	return
}

// Update takes the representation of a resourceQuotaClaim and updates it. Returns the server's representation of the resourceQuotaClaim, and an error, if there is any.
func (c *resourceQuotaClaims) Update(resourceQuotaClaim *v1.ResourceQuotaClaim) (result *v1.ResourceQuotaClaim, err error) {
	result = &v1.ResourceQuotaClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resourcequotaclaims").
		Name(resourceQuotaClaim.Name).
		Body(resourceQuotaClaim).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *resourceQuotaClaims) UpdateStatus(resourceQuotaClaim *v1.ResourceQuotaClaim) (result *v1.ResourceQuotaClaim, err error) {
	result = &v1.ResourceQuotaClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resourcequotaclaims").
		Name(resourceQuotaClaim.Name).
		SubResource("status").
		Body(resourceQuotaClaim).
		Do().
		Into(result)
	return
}

// Delete takes name of the resourceQuotaClaim and deletes it. Returns an error if one occurs.
func (c *resourceQuotaClaims) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourcequotaclaims").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *resourceQuotaClaims) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourcequotaclaims").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched resourceQuotaClaim.
func (c *resourceQuotaClaims) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ResourceQuotaClaim, err error) {
	result = &v1.ResourceQuotaClaim{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("resourcequotaclaims").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

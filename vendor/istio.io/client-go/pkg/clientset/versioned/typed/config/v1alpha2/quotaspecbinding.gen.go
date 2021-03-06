// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	"time"

	v1alpha2 "istio.io/client-go/pkg/apis/config/v1alpha2"
	scheme "istio.io/client-go/pkg/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// QuotaSpecBindingsGetter has a method to return a QuotaSpecBindingInterface.
// A group's client should implement this interface.
type QuotaSpecBindingsGetter interface {
	QuotaSpecBindings(namespace string) QuotaSpecBindingInterface
}

// QuotaSpecBindingInterface has methods to work with QuotaSpecBinding resources.
type QuotaSpecBindingInterface interface {
	Create(*v1alpha2.QuotaSpecBinding) (*v1alpha2.QuotaSpecBinding, error)
	Update(*v1alpha2.QuotaSpecBinding) (*v1alpha2.QuotaSpecBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.QuotaSpecBinding, error)
	List(opts v1.ListOptions) (*v1alpha2.QuotaSpecBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.QuotaSpecBinding, err error)
	QuotaSpecBindingExpansion
}

// quotaSpecBindings implements QuotaSpecBindingInterface
type quotaSpecBindings struct {
	client rest.Interface
	ns     string
}

// newQuotaSpecBindings returns a QuotaSpecBindings
func newQuotaSpecBindings(c *ConfigV1alpha2Client, namespace string) *quotaSpecBindings {
	return &quotaSpecBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the quotaSpecBinding, and returns the corresponding quotaSpecBinding object, and an error if there is any.
func (c *quotaSpecBindings) Get(name string, options v1.GetOptions) (result *v1alpha2.QuotaSpecBinding, err error) {
	result = &v1alpha2.QuotaSpecBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("quotaspecbindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of QuotaSpecBindings that match those selectors.
func (c *quotaSpecBindings) List(opts v1.ListOptions) (result *v1alpha2.QuotaSpecBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.QuotaSpecBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("quotaspecbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested quotaSpecBindings.
func (c *quotaSpecBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("quotaspecbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a quotaSpecBinding and creates it.  Returns the server's representation of the quotaSpecBinding, and an error, if there is any.
func (c *quotaSpecBindings) Create(quotaSpecBinding *v1alpha2.QuotaSpecBinding) (result *v1alpha2.QuotaSpecBinding, err error) {
	result = &v1alpha2.QuotaSpecBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("quotaspecbindings").
		Body(quotaSpecBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a quotaSpecBinding and updates it. Returns the server's representation of the quotaSpecBinding, and an error, if there is any.
func (c *quotaSpecBindings) Update(quotaSpecBinding *v1alpha2.QuotaSpecBinding) (result *v1alpha2.QuotaSpecBinding, err error) {
	result = &v1alpha2.QuotaSpecBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("quotaspecbindings").
		Name(quotaSpecBinding.Name).
		Body(quotaSpecBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the quotaSpecBinding and deletes it. Returns an error if one occurs.
func (c *quotaSpecBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("quotaspecbindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *quotaSpecBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("quotaspecbindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched quotaSpecBinding.
func (c *quotaSpecBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.QuotaSpecBinding, err error) {
	result = &v1alpha2.QuotaSpecBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("quotaspecbindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

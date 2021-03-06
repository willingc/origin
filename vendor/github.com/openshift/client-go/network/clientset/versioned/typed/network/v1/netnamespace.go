package v1

import (
	v1 "github.com/openshift/api/network/v1"
	scheme "github.com/openshift/client-go/network/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NetNamespacesGetter has a method to return a NetNamespaceInterface.
// A group's client should implement this interface.
type NetNamespacesGetter interface {
	NetNamespaces() NetNamespaceInterface
}

// NetNamespaceInterface has methods to work with NetNamespace resources.
type NetNamespaceInterface interface {
	Create(*v1.NetNamespace) (*v1.NetNamespace, error)
	Update(*v1.NetNamespace) (*v1.NetNamespace, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.NetNamespace, error)
	List(opts meta_v1.ListOptions) (*v1.NetNamespaceList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.NetNamespace, err error)
	NetNamespaceExpansion
}

// netNamespaces implements NetNamespaceInterface
type netNamespaces struct {
	client rest.Interface
}

// newNetNamespaces returns a NetNamespaces
func newNetNamespaces(c *NetworkV1Client) *netNamespaces {
	return &netNamespaces{
		client: c.RESTClient(),
	}
}

// Get takes name of the netNamespace, and returns the corresponding netNamespace object, and an error if there is any.
func (c *netNamespaces) Get(name string, options meta_v1.GetOptions) (result *v1.NetNamespace, err error) {
	result = &v1.NetNamespace{}
	err = c.client.Get().
		Resource("netnamespaces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetNamespaces that match those selectors.
func (c *netNamespaces) List(opts meta_v1.ListOptions) (result *v1.NetNamespaceList, err error) {
	result = &v1.NetNamespaceList{}
	err = c.client.Get().
		Resource("netnamespaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested netNamespaces.
func (c *netNamespaces) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("netnamespaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a netNamespace and creates it.  Returns the server's representation of the netNamespace, and an error, if there is any.
func (c *netNamespaces) Create(netNamespace *v1.NetNamespace) (result *v1.NetNamespace, err error) {
	result = &v1.NetNamespace{}
	err = c.client.Post().
		Resource("netnamespaces").
		Body(netNamespace).
		Do().
		Into(result)
	return
}

// Update takes the representation of a netNamespace and updates it. Returns the server's representation of the netNamespace, and an error, if there is any.
func (c *netNamespaces) Update(netNamespace *v1.NetNamespace) (result *v1.NetNamespace, err error) {
	result = &v1.NetNamespace{}
	err = c.client.Put().
		Resource("netnamespaces").
		Name(netNamespace.Name).
		Body(netNamespace).
		Do().
		Into(result)
	return
}

// Delete takes name of the netNamespace and deletes it. Returns an error if one occurs.
func (c *netNamespaces) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("netnamespaces").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *netNamespaces) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Resource("netnamespaces").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched netNamespace.
func (c *netNamespaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.NetNamespace, err error) {
	result = &v1.NetNamespace{}
	err = c.client.Patch(pt).
		Resource("netnamespaces").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

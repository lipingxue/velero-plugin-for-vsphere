/*
Copyright the Velero contributors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	veleropluginv1 "github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/apis/veleroplugin/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDownloads implements DownloadInterface
type FakeDownloads struct {
	Fake *FakeVeleropluginV1
	ns   string
}

var downloadsResource = schema.GroupVersionResource{Group: "veleroplugin.io", Version: "v1", Resource: "downloads"}

var downloadsKind = schema.GroupVersionKind{Group: "veleroplugin.io", Version: "v1", Kind: "Download"}

// Get takes name of the download, and returns the corresponding download object, and an error if there is any.
func (c *FakeDownloads) Get(ctx context.Context, name string, options v1.GetOptions) (result *veleropluginv1.Download, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(downloadsResource, c.ns, name), &veleropluginv1.Download{})

	if obj == nil {
		return nil, err
	}
	return obj.(*veleropluginv1.Download), err
}

// List takes label and field selectors, and returns the list of Downloads that match those selectors.
func (c *FakeDownloads) List(ctx context.Context, opts v1.ListOptions) (result *veleropluginv1.DownloadList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(downloadsResource, downloadsKind, c.ns, opts), &veleropluginv1.DownloadList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &veleropluginv1.DownloadList{ListMeta: obj.(*veleropluginv1.DownloadList).ListMeta}
	for _, item := range obj.(*veleropluginv1.DownloadList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested downloads.
func (c *FakeDownloads) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(downloadsResource, c.ns, opts))

}

// Create takes the representation of a download and creates it.  Returns the server's representation of the download, and an error, if there is any.
func (c *FakeDownloads) Create(ctx context.Context, download *veleropluginv1.Download, opts v1.CreateOptions) (result *veleropluginv1.Download, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(downloadsResource, c.ns, download), &veleropluginv1.Download{})

	if obj == nil {
		return nil, err
	}
	return obj.(*veleropluginv1.Download), err
}

// Update takes the representation of a download and updates it. Returns the server's representation of the download, and an error, if there is any.
func (c *FakeDownloads) Update(ctx context.Context, download *veleropluginv1.Download, opts v1.UpdateOptions) (result *veleropluginv1.Download, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(downloadsResource, c.ns, download), &veleropluginv1.Download{})

	if obj == nil {
		return nil, err
	}
	return obj.(*veleropluginv1.Download), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDownloads) UpdateStatus(ctx context.Context, download *veleropluginv1.Download, opts v1.UpdateOptions) (*veleropluginv1.Download, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(downloadsResource, "status", c.ns, download), &veleropluginv1.Download{})

	if obj == nil {
		return nil, err
	}
	return obj.(*veleropluginv1.Download), err
}

// Delete takes name of the download and deletes it. Returns an error if one occurs.
func (c *FakeDownloads) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(downloadsResource, c.ns, name), &veleropluginv1.Download{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDownloads) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(downloadsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &veleropluginv1.DownloadList{})
	return err
}

// Patch applies the patch and returns the patched download.
func (c *FakeDownloads) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *veleropluginv1.Download, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(downloadsResource, c.ns, name, pt, data, subresources...), &veleropluginv1.Download{})

	if obj == nil {
		return nil, err
	}
	return obj.(*veleropluginv1.Download), err
}

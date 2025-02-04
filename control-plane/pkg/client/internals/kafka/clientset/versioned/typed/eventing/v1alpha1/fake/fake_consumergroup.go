/*
 * Copyright 2021 The Knative Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	autoscalingv1 "k8s.io/api/autoscaling/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "knative.dev/eventing-kafka-broker/control-plane/pkg/apis/internals/kafka/eventing/v1alpha1"
)

// FakeConsumerGroups implements ConsumerGroupInterface
type FakeConsumerGroups struct {
	Fake *FakeInternalV1alpha1
	ns   string
}

var consumergroupsResource = schema.GroupVersionResource{Group: "internal.kafka.eventing.knative.dev", Version: "v1alpha1", Resource: "consumergroups"}

var consumergroupsKind = schema.GroupVersionKind{Group: "internal.kafka.eventing.knative.dev", Version: "v1alpha1", Kind: "ConsumerGroup"}

// Get takes name of the consumerGroup, and returns the corresponding consumerGroup object, and an error if there is any.
func (c *FakeConsumerGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ConsumerGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(consumergroupsResource, c.ns, name), &v1alpha1.ConsumerGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConsumerGroup), err
}

// List takes label and field selectors, and returns the list of ConsumerGroups that match those selectors.
func (c *FakeConsumerGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ConsumerGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(consumergroupsResource, consumergroupsKind, c.ns, opts), &v1alpha1.ConsumerGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ConsumerGroupList{ListMeta: obj.(*v1alpha1.ConsumerGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.ConsumerGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested consumerGroups.
func (c *FakeConsumerGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(consumergroupsResource, c.ns, opts))

}

// Create takes the representation of a consumerGroup and creates it.  Returns the server's representation of the consumerGroup, and an error, if there is any.
func (c *FakeConsumerGroups) Create(ctx context.Context, consumerGroup *v1alpha1.ConsumerGroup, opts v1.CreateOptions) (result *v1alpha1.ConsumerGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(consumergroupsResource, c.ns, consumerGroup), &v1alpha1.ConsumerGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConsumerGroup), err
}

// Update takes the representation of a consumerGroup and updates it. Returns the server's representation of the consumerGroup, and an error, if there is any.
func (c *FakeConsumerGroups) Update(ctx context.Context, consumerGroup *v1alpha1.ConsumerGroup, opts v1.UpdateOptions) (result *v1alpha1.ConsumerGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(consumergroupsResource, c.ns, consumerGroup), &v1alpha1.ConsumerGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConsumerGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeConsumerGroups) UpdateStatus(ctx context.Context, consumerGroup *v1alpha1.ConsumerGroup, opts v1.UpdateOptions) (*v1alpha1.ConsumerGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(consumergroupsResource, "status", c.ns, consumerGroup), &v1alpha1.ConsumerGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConsumerGroup), err
}

// Delete takes name of the consumerGroup and deletes it. Returns an error if one occurs.
func (c *FakeConsumerGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(consumergroupsResource, c.ns, name), &v1alpha1.ConsumerGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConsumerGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(consumergroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ConsumerGroupList{})
	return err
}

// Patch applies the patch and returns the patched consumerGroup.
func (c *FakeConsumerGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ConsumerGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(consumergroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ConsumerGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConsumerGroup), err
}

// GetScale takes name of the consumerGroup, and returns the corresponding scale object, and an error if there is any.
func (c *FakeConsumerGroups) GetScale(ctx context.Context, consumerGroupName string, options v1.GetOptions) (result *autoscalingv1.Scale, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetSubresourceAction(consumergroupsResource, c.ns, "scale", consumerGroupName), &autoscalingv1.Scale{})

	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.Scale), err
}

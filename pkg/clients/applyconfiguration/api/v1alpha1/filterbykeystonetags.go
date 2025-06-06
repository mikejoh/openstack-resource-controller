/*
Copyright 2024 The ORC Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	apiv1alpha1 "github.com/k-orc/openstack-resource-controller/v2/api/v1alpha1"
)

// FilterByKeystoneTagsApplyConfiguration represents a declarative configuration of the FilterByKeystoneTags type for use
// with apply.
type FilterByKeystoneTagsApplyConfiguration struct {
	Tags       []apiv1alpha1.KeystoneTag `json:"tags,omitempty"`
	TagsAny    []apiv1alpha1.KeystoneTag `json:"tagsAny,omitempty"`
	NotTags    []apiv1alpha1.KeystoneTag `json:"notTags,omitempty"`
	NotTagsAny []apiv1alpha1.KeystoneTag `json:"notTagsAny,omitempty"`
}

// FilterByKeystoneTagsApplyConfiguration constructs a declarative configuration of the FilterByKeystoneTags type for use with
// apply.
func FilterByKeystoneTags() *FilterByKeystoneTagsApplyConfiguration {
	return &FilterByKeystoneTagsApplyConfiguration{}
}

// WithTags adds the given value to the Tags field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tags field.
func (b *FilterByKeystoneTagsApplyConfiguration) WithTags(values ...apiv1alpha1.KeystoneTag) *FilterByKeystoneTagsApplyConfiguration {
	for i := range values {
		b.Tags = append(b.Tags, values[i])
	}
	return b
}

// WithTagsAny adds the given value to the TagsAny field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the TagsAny field.
func (b *FilterByKeystoneTagsApplyConfiguration) WithTagsAny(values ...apiv1alpha1.KeystoneTag) *FilterByKeystoneTagsApplyConfiguration {
	for i := range values {
		b.TagsAny = append(b.TagsAny, values[i])
	}
	return b
}

// WithNotTags adds the given value to the NotTags field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NotTags field.
func (b *FilterByKeystoneTagsApplyConfiguration) WithNotTags(values ...apiv1alpha1.KeystoneTag) *FilterByKeystoneTagsApplyConfiguration {
	for i := range values {
		b.NotTags = append(b.NotTags, values[i])
	}
	return b
}

// WithNotTagsAny adds the given value to the NotTagsAny field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NotTagsAny field.
func (b *FilterByKeystoneTagsApplyConfiguration) WithNotTagsAny(values ...apiv1alpha1.KeystoneTag) *FilterByKeystoneTagsApplyConfiguration {
	for i := range values {
		b.NotTagsAny = append(b.NotTagsAny, values[i])
	}
	return b
}

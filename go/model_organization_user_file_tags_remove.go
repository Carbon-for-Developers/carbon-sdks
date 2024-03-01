/*
Carbon

Connect external data to LLMs, no matter the source.

API version: 1.0.0
*/

// Code generated by Konfig (https://konfigthis.com); DO NOT EDIT.

package carbon

import (
	"encoding/json"
)

// OrganizationUserFileTagsRemove struct for OrganizationUserFileTagsRemove
type OrganizationUserFileTagsRemove struct {
	Tags []string `json:"tags"`
	OrganizationUserFileId int32 `json:"organization_user_file_id"`
}

// NewOrganizationUserFileTagsRemove instantiates a new OrganizationUserFileTagsRemove object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationUserFileTagsRemove(tags []string, organizationUserFileId int32) *OrganizationUserFileTagsRemove {
	this := OrganizationUserFileTagsRemove{}
	this.Tags = tags
	this.OrganizationUserFileId = organizationUserFileId
	return &this
}

// NewOrganizationUserFileTagsRemoveWithDefaults instantiates a new OrganizationUserFileTagsRemove object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationUserFileTagsRemoveWithDefaults() *OrganizationUserFileTagsRemove {
	this := OrganizationUserFileTagsRemove{}
	return &this
}

// GetTags returns the Tags field value
func (o *OrganizationUserFileTagsRemove) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *OrganizationUserFileTagsRemove) GetTagsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *OrganizationUserFileTagsRemove) SetTags(v []string) {
	o.Tags = v
}

// GetOrganizationUserFileId returns the OrganizationUserFileId field value
func (o *OrganizationUserFileTagsRemove) GetOrganizationUserFileId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrganizationUserFileId
}

// GetOrganizationUserFileIdOk returns a tuple with the OrganizationUserFileId field value
// and a boolean to check if the value has been set.
func (o *OrganizationUserFileTagsRemove) GetOrganizationUserFileIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OrganizationUserFileId, true
}

// SetOrganizationUserFileId sets field value
func (o *OrganizationUserFileTagsRemove) SetOrganizationUserFileId(v int32) {
	o.OrganizationUserFileId = v
}

func (o OrganizationUserFileTagsRemove) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["organization_user_file_id"] = o.OrganizationUserFileId
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationUserFileTagsRemove struct {
	value *OrganizationUserFileTagsRemove
	isSet bool
}

func (v NullableOrganizationUserFileTagsRemove) Get() *OrganizationUserFileTagsRemove {
	return v.value
}

func (v *NullableOrganizationUserFileTagsRemove) Set(val *OrganizationUserFileTagsRemove) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationUserFileTagsRemove) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationUserFileTagsRemove) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationUserFileTagsRemove(val *OrganizationUserFileTagsRemove) *NullableOrganizationUserFileTagsRemove {
	return &NullableOrganizationUserFileTagsRemove{value: val, isSet: true}
}

func (v NullableOrganizationUserFileTagsRemove) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationUserFileTagsRemove) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



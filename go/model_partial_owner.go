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

// PartialOwner struct for PartialOwner
type PartialOwner struct {
	Id NullableString `json:"id"`
}

// NewPartialOwner instantiates a new PartialOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartialOwner(id NullableString) *PartialOwner {
	this := PartialOwner{}
	this.Id = id
	return &this
}

// NewPartialOwnerWithDefaults instantiates a new PartialOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartialOwnerWithDefaults() *PartialOwner {
	this := PartialOwner{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PartialOwner) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartialOwner) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *PartialOwner) SetId(v string) {
	o.Id.Set(&v)
}

func (o PartialOwner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePartialOwner struct {
	value *PartialOwner
	isSet bool
}

func (v NullablePartialOwner) Get() *PartialOwner {
	return v.value
}

func (v *NullablePartialOwner) Set(val *PartialOwner) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialOwner) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialOwner(val *PartialOwner) *NullablePartialOwner {
	return &NullablePartialOwner{value: val, isSet: true}
}

func (v NullablePartialOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



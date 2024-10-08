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

// HeadOrBase struct for HeadOrBase
type HeadOrBase struct {
	Label string `json:"label"`
	Ref string `json:"ref"`
	Sha string `json:"sha"`
}

// NewHeadOrBase instantiates a new HeadOrBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeadOrBase(label string, ref string, sha string) *HeadOrBase {
	this := HeadOrBase{}
	this.Label = label
	this.Ref = ref
	this.Sha = sha
	return &this
}

// NewHeadOrBaseWithDefaults instantiates a new HeadOrBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeadOrBaseWithDefaults() *HeadOrBase {
	this := HeadOrBase{}
	return &this
}

// GetLabel returns the Label field value
func (o *HeadOrBase) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *HeadOrBase) GetLabelOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *HeadOrBase) SetLabel(v string) {
	o.Label = v
}

// GetRef returns the Ref field value
func (o *HeadOrBase) GetRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ref
}

// GetRefOk returns a tuple with the Ref field value
// and a boolean to check if the value has been set.
func (o *HeadOrBase) GetRefOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Ref, true
}

// SetRef sets field value
func (o *HeadOrBase) SetRef(v string) {
	o.Ref = v
}

// GetSha returns the Sha field value
func (o *HeadOrBase) GetSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha
}

// GetShaOk returns a tuple with the Sha field value
// and a boolean to check if the value has been set.
func (o *HeadOrBase) GetShaOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Sha, true
}

// SetSha sets field value
func (o *HeadOrBase) SetSha(v string) {
	o.Sha = v
}

func (o HeadOrBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["ref"] = o.Ref
	}
	if true {
		toSerialize["sha"] = o.Sha
	}
	return json.Marshal(toSerialize)
}

type NullableHeadOrBase struct {
	value *HeadOrBase
	isSet bool
}

func (v NullableHeadOrBase) Get() *HeadOrBase {
	return v.value
}

func (v *NullableHeadOrBase) Set(val *HeadOrBase) {
	v.value = val
	v.isSet = true
}

func (v NullableHeadOrBase) IsSet() bool {
	return v.isSet
}

func (v *NullableHeadOrBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeadOrBase(val *HeadOrBase) *NullableHeadOrBase {
	return &NullableHeadOrBase{value: val, isSet: true}
}

func (v NullableHeadOrBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeadOrBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



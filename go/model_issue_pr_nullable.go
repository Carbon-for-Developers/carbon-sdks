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

// IssuePRNullable struct for IssuePRNullable
type IssuePRNullable struct {
	Url string `json:"url"`
	MergedAt NullableString `json:"merged_at"`
}

// NewIssuePRNullable instantiates a new IssuePRNullable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuePRNullable(url string, mergedAt NullableString) *IssuePRNullable {
	this := IssuePRNullable{}
	this.Url = url
	this.MergedAt = mergedAt
	return &this
}

// NewIssuePRNullableWithDefaults instantiates a new IssuePRNullable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuePRNullableWithDefaults() *IssuePRNullable {
	this := IssuePRNullable{}
	return &this
}

// GetUrl returns the Url field value
func (o *IssuePRNullable) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *IssuePRNullable) GetUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *IssuePRNullable) SetUrl(v string) {
	o.Url = v
}

// GetMergedAt returns the MergedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IssuePRNullable) GetMergedAt() string {
	if o == nil || o.MergedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.MergedAt.Get()
}

// GetMergedAtOk returns a tuple with the MergedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuePRNullable) GetMergedAtOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.MergedAt.Get(), o.MergedAt.IsSet()
}

// SetMergedAt sets field value
func (o *IssuePRNullable) SetMergedAt(v string) {
	o.MergedAt.Set(&v)
}

func (o IssuePRNullable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["merged_at"] = o.MergedAt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIssuePRNullable struct {
	value *IssuePRNullable
	isSet bool
}

func (v NullableIssuePRNullable) Get() *IssuePRNullable {
	return v.value
}

func (v *NullableIssuePRNullable) Set(val *IssuePRNullable) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuePRNullable) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuePRNullable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuePRNullable(val *IssuePRNullable) *NullableIssuePRNullable {
	return &NullableIssuePRNullable{value: val, isSet: true}
}

func (v NullableIssuePRNullable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuePRNullable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



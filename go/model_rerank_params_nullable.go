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

// RerankParamsNullable Parameters for reranking the chunks using a specified model. This field accepts an object with details of the reranking model to be used; either 'Cohere' or 'Jina'. If provided, the specified reranking model will reorder the retrieved chunks based on their relevance to the query.
type RerankParamsNullable struct {
	Model string `json:"model"`
}

// NewRerankParamsNullable instantiates a new RerankParamsNullable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRerankParamsNullable(model string) *RerankParamsNullable {
	this := RerankParamsNullable{}
	this.Model = model
	return &this
}

// NewRerankParamsNullableWithDefaults instantiates a new RerankParamsNullable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRerankParamsNullableWithDefaults() *RerankParamsNullable {
	this := RerankParamsNullable{}
	return &this
}

// GetModel returns the Model field value
func (o *RerankParamsNullable) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *RerankParamsNullable) GetModelOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *RerankParamsNullable) SetModel(v string) {
	o.Model = v
}

func (o RerankParamsNullable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["model"] = o.Model
	}
	return json.Marshal(toSerialize)
}

type NullableRerankParamsNullable struct {
	value *RerankParamsNullable
	isSet bool
}

func (v NullableRerankParamsNullable) Get() *RerankParamsNullable {
	return v.value
}

func (v *NullableRerankParamsNullable) Set(val *RerankParamsNullable) {
	v.value = val
	v.isSet = true
}

func (v NullableRerankParamsNullable) IsSet() bool {
	return v.isSet
}

func (v *NullableRerankParamsNullable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRerankParamsNullable(val *RerankParamsNullable) *NullableRerankParamsNullable {
	return &NullableRerankParamsNullable{value: val, isSet: true}
}

func (v NullableRerankParamsNullable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRerankParamsNullable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



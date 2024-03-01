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

// UserRequestContent struct for UserRequestContent
type UserRequestContent struct {
	CustomerId string `json:"customer_id"`
}

// NewUserRequestContent instantiates a new UserRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRequestContent(customerId string) *UserRequestContent {
	this := UserRequestContent{}
	this.CustomerId = customerId
	return &this
}

// NewUserRequestContentWithDefaults instantiates a new UserRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRequestContentWithDefaults() *UserRequestContent {
	this := UserRequestContent{}
	return &this
}

// GetCustomerId returns the CustomerId field value
func (o *UserRequestContent) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *UserRequestContent) GetCustomerIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *UserRequestContent) SetCustomerId(v string) {
	o.CustomerId = v
}

func (o UserRequestContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["customer_id"] = o.CustomerId
	}
	return json.Marshal(toSerialize)
}

type NullableUserRequestContent struct {
	value *UserRequestContent
	isSet bool
}

func (v NullableUserRequestContent) Get() *UserRequestContent {
	return v.value
}

func (v *NullableUserRequestContent) Set(val *UserRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRequestContent(val *UserRequestContent) *NullableUserRequestContent {
	return &NullableUserRequestContent{value: val, isSet: true}
}

func (v NullableUserRequestContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



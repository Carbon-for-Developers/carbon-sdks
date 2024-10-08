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

// UserNullable struct for UserNullable
type UserNullable struct {
	Id int32 `json:"id"`
	Login string `json:"login"`
}

// NewUserNullable instantiates a new UserNullable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserNullable(id int32, login string) *UserNullable {
	this := UserNullable{}
	this.Id = id
	this.Login = login
	return &this
}

// NewUserNullableWithDefaults instantiates a new UserNullable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserNullableWithDefaults() *UserNullable {
	this := UserNullable{}
	return &this
}

// GetId returns the Id field value
func (o *UserNullable) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserNullable) GetIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserNullable) SetId(v int32) {
	o.Id = v
}

// GetLogin returns the Login field value
func (o *UserNullable) GetLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Login
}

// GetLoginOk returns a tuple with the Login field value
// and a boolean to check if the value has been set.
func (o *UserNullable) GetLoginOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Login, true
}

// SetLogin sets field value
func (o *UserNullable) SetLogin(v string) {
	o.Login = v
}

func (o UserNullable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["login"] = o.Login
	}
	return json.Marshal(toSerialize)
}

type NullableUserNullable struct {
	value *UserNullable
	isSet bool
}

func (v NullableUserNullable) Get() *UserNullable {
	return v.value
}

func (v *NullableUserNullable) Set(val *UserNullable) {
	v.value = val
	v.isSet = true
}

func (v NullableUserNullable) IsSet() bool {
	return v.isSet
}

func (v *NullableUserNullable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserNullable(val *UserNullable) *NullableUserNullable {
	return &NullableUserNullable{value: val, isSet: true}
}

func (v NullableUserNullable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserNullable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



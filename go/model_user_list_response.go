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

// UserListResponse struct for UserListResponse
type UserListResponse struct {
	Users []ListUserResponse `json:"users"`
	Count NullableInt32 `json:"count"`
}

// NewUserListResponse instantiates a new UserListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserListResponse(users []ListUserResponse, count NullableInt32) *UserListResponse {
	this := UserListResponse{}
	this.Users = users
	this.Count = count
	return &this
}

// NewUserListResponseWithDefaults instantiates a new UserListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserListResponseWithDefaults() *UserListResponse {
	this := UserListResponse{}
	return &this
}

// GetUsers returns the Users field value
func (o *UserListResponse) GetUsers() []ListUserResponse {
	if o == nil {
		var ret []ListUserResponse
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *UserListResponse) GetUsersOk() ([]ListUserResponse, bool) {
	if o == nil {
    return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *UserListResponse) SetUsers(v []ListUserResponse) {
	o.Users = v
}

// GetCount returns the Count field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *UserListResponse) GetCount() int32 {
	if o == nil || o.Count.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Count.Get()
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserListResponse) GetCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Count.Get(), o.Count.IsSet()
}

// SetCount sets field value
func (o *UserListResponse) SetCount(v int32) {
	o.Count.Set(&v)
}

func (o UserListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["users"] = o.Users
	}
	if true {
		toSerialize["count"] = o.Count.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUserListResponse struct {
	value *UserListResponse
	isSet bool
}

func (v NullableUserListResponse) Get() *UserListResponse {
	return v.value
}

func (v *NullableUserListResponse) Set(val *UserListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserListResponse(val *UserListResponse) *NullableUserListResponse {
	return &NullableUserListResponse{value: val, isSet: true}
}

func (v NullableUserListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



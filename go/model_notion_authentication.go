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

// NotionAuthentication struct for NotionAuthentication
type NotionAuthentication struct {
	Source interface{} `json:"source"`
	AccessToken string `json:"access_token"`
	WorkspaceId string `json:"workspace_id"`
}

// NewNotionAuthentication instantiates a new NotionAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotionAuthentication(source interface{}, accessToken string, workspaceId string) *NotionAuthentication {
	this := NotionAuthentication{}
	this.Source = source
	this.AccessToken = accessToken
	this.WorkspaceId = workspaceId
	return &this
}

// NewNotionAuthenticationWithDefaults instantiates a new NotionAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotionAuthenticationWithDefaults() *NotionAuthentication {
	this := NotionAuthentication{}
	return &this
}

// GetSource returns the Source field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *NotionAuthentication) GetSource() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotionAuthentication) GetSourceOk() (*interface{}, bool) {
	if o == nil || isNil(o.Source) {
    return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *NotionAuthentication) SetSource(v interface{}) {
	o.Source = v
}

// GetAccessToken returns the AccessToken field value
func (o *NotionAuthentication) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *NotionAuthentication) GetAccessTokenOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *NotionAuthentication) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *NotionAuthentication) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *NotionAuthentication) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *NotionAuthentication) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

func (o NotionAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if true {
		toSerialize["workspace_id"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableNotionAuthentication struct {
	value *NotionAuthentication
	isSet bool
}

func (v NullableNotionAuthentication) Get() *NotionAuthentication {
	return v.value
}

func (v *NullableNotionAuthentication) Set(val *NotionAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableNotionAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableNotionAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotionAuthentication(val *NotionAuthentication) *NullableNotionAuthentication {
	return &NullableNotionAuthentication{value: val, isSet: true}
}

func (v NullableNotionAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotionAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



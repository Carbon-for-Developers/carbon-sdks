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

// OuthURLResponse struct for OuthURLResponse
type OuthURLResponse struct {
	OauthUrl string `json:"oauth_url"`
}

// NewOuthURLResponse instantiates a new OuthURLResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOuthURLResponse(oauthUrl string) *OuthURLResponse {
	this := OuthURLResponse{}
	this.OauthUrl = oauthUrl
	return &this
}

// NewOuthURLResponseWithDefaults instantiates a new OuthURLResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOuthURLResponseWithDefaults() *OuthURLResponse {
	this := OuthURLResponse{}
	return &this
}

// GetOauthUrl returns the OauthUrl field value
func (o *OuthURLResponse) GetOauthUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OauthUrl
}

// GetOauthUrlOk returns a tuple with the OauthUrl field value
// and a boolean to check if the value has been set.
func (o *OuthURLResponse) GetOauthUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OauthUrl, true
}

// SetOauthUrl sets field value
func (o *OuthURLResponse) SetOauthUrl(v string) {
	o.OauthUrl = v
}

func (o OuthURLResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["oauth_url"] = o.OauthUrl
	}
	return json.Marshal(toSerialize)
}

type NullableOuthURLResponse struct {
	value *OuthURLResponse
	isSet bool
}

func (v NullableOuthURLResponse) Get() *OuthURLResponse {
	return v.value
}

func (v *NullableOuthURLResponse) Set(val *OuthURLResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOuthURLResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOuthURLResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOuthURLResponse(val *OuthURLResponse) *NullableOuthURLResponse {
	return &NullableOuthURLResponse{value: val, isSet: true}
}

func (v NullableOuthURLResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOuthURLResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



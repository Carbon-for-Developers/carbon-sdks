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

// GuruAuthentication struct for GuruAuthentication
type GuruAuthentication struct {
	Source string `json:"source"`
	AccessToken string `json:"access_token"`
	Username string `json:"username"`
}

// NewGuruAuthentication instantiates a new GuruAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuruAuthentication(source string, accessToken string, username string) *GuruAuthentication {
	this := GuruAuthentication{}
	this.Source = source
	this.AccessToken = accessToken
	this.Username = username
	return &this
}

// NewGuruAuthenticationWithDefaults instantiates a new GuruAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuruAuthenticationWithDefaults() *GuruAuthentication {
	this := GuruAuthentication{}
	return &this
}

// GetSource returns the Source field value
func (o *GuruAuthentication) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *GuruAuthentication) GetSourceOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *GuruAuthentication) SetSource(v string) {
	o.Source = v
}

// GetAccessToken returns the AccessToken field value
func (o *GuruAuthentication) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *GuruAuthentication) GetAccessTokenOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *GuruAuthentication) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetUsername returns the Username field value
func (o *GuruAuthentication) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *GuruAuthentication) GetUsernameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *GuruAuthentication) SetUsername(v string) {
	o.Username = v
}

func (o GuruAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if true {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableGuruAuthentication struct {
	value *GuruAuthentication
	isSet bool
}

func (v NullableGuruAuthentication) Get() *GuruAuthentication {
	return v.value
}

func (v *NullableGuruAuthentication) Set(val *GuruAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableGuruAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableGuruAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuruAuthentication(val *GuruAuthentication) *NullableGuruAuthentication {
	return &NullableGuruAuthentication{value: val, isSet: true}
}

func (v NullableGuruAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuruAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



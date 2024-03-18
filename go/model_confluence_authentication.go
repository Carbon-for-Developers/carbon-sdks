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

// ConfluenceAuthentication struct for ConfluenceAuthentication
type ConfluenceAuthentication struct {
	Source interface{} `json:"source"`
	AccessToken string `json:"access_token"`
	RefreshToken NullableString `json:"refresh_token,omitempty"`
	Subdomain string `json:"subdomain"`
}

// NewConfluenceAuthentication instantiates a new ConfluenceAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfluenceAuthentication(source interface{}, accessToken string, subdomain string) *ConfluenceAuthentication {
	this := ConfluenceAuthentication{}
	this.Source = source
	this.AccessToken = accessToken
	this.Subdomain = subdomain
	return &this
}

// NewConfluenceAuthenticationWithDefaults instantiates a new ConfluenceAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfluenceAuthenticationWithDefaults() *ConfluenceAuthentication {
	this := ConfluenceAuthentication{}
	return &this
}

// GetSource returns the Source field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ConfluenceAuthentication) GetSource() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfluenceAuthentication) GetSourceOk() (*interface{}, bool) {
	if o == nil || isNil(o.Source) {
    return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *ConfluenceAuthentication) SetSource(v interface{}) {
	o.Source = v
}

// GetAccessToken returns the AccessToken field value
func (o *ConfluenceAuthentication) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *ConfluenceAuthentication) GetAccessTokenOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *ConfluenceAuthentication) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfluenceAuthentication) GetRefreshToken() string {
	if o == nil || isNil(o.RefreshToken.Get()) {
		var ret string
		return ret
	}
	return *o.RefreshToken.Get()
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfluenceAuthentication) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RefreshToken.Get(), o.RefreshToken.IsSet()
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *ConfluenceAuthentication) HasRefreshToken() bool {
	if o != nil && o.RefreshToken.IsSet() {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given NullableString and assigns it to the RefreshToken field.
func (o *ConfluenceAuthentication) SetRefreshToken(v string) {
	o.RefreshToken.Set(&v)
}
// SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil
func (o *ConfluenceAuthentication) SetRefreshTokenNil() {
	o.RefreshToken.Set(nil)
}

// UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
func (o *ConfluenceAuthentication) UnsetRefreshToken() {
	o.RefreshToken.Unset()
}

// GetSubdomain returns the Subdomain field value
func (o *ConfluenceAuthentication) GetSubdomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value
// and a boolean to check if the value has been set.
func (o *ConfluenceAuthentication) GetSubdomainOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Subdomain, true
}

// SetSubdomain sets field value
func (o *ConfluenceAuthentication) SetSubdomain(v string) {
	o.Subdomain = v
}

func (o ConfluenceAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.RefreshToken.IsSet() {
		toSerialize["refresh_token"] = o.RefreshToken.Get()
	}
	if true {
		toSerialize["subdomain"] = o.Subdomain
	}
	return json.Marshal(toSerialize)
}

type NullableConfluenceAuthentication struct {
	value *ConfluenceAuthentication
	isSet bool
}

func (v NullableConfluenceAuthentication) Get() *ConfluenceAuthentication {
	return v.value
}

func (v *NullableConfluenceAuthentication) Set(val *ConfluenceAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableConfluenceAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableConfluenceAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfluenceAuthentication(val *ConfluenceAuthentication) *NullableConfluenceAuthentication {
	return &NullableConfluenceAuthentication{value: val, isSet: true}
}

func (v NullableConfluenceAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfluenceAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


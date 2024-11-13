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

// OneDriveAuthentication struct for OneDriveAuthentication
type OneDriveAuthentication struct {
	Source string `json:"source"`
	AccessToken string `json:"access_token"`
	RefreshToken NullableString `json:"refresh_token,omitempty"`
	TenantName NullableString `json:"tenant_name,omitempty"`
}

// NewOneDriveAuthentication instantiates a new OneDriveAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOneDriveAuthentication(source string, accessToken string) *OneDriveAuthentication {
	this := OneDriveAuthentication{}
	this.Source = source
	this.AccessToken = accessToken
	return &this
}

// NewOneDriveAuthenticationWithDefaults instantiates a new OneDriveAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOneDriveAuthenticationWithDefaults() *OneDriveAuthentication {
	this := OneDriveAuthentication{}
	return &this
}

// GetSource returns the Source field value
func (o *OneDriveAuthentication) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *OneDriveAuthentication) GetSourceOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *OneDriveAuthentication) SetSource(v string) {
	o.Source = v
}

// GetAccessToken returns the AccessToken field value
func (o *OneDriveAuthentication) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *OneDriveAuthentication) GetAccessTokenOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *OneDriveAuthentication) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OneDriveAuthentication) GetRefreshToken() string {
	if o == nil || isNil(o.RefreshToken.Get()) {
		var ret string
		return ret
	}
	return *o.RefreshToken.Get()
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OneDriveAuthentication) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RefreshToken.Get(), o.RefreshToken.IsSet()
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *OneDriveAuthentication) HasRefreshToken() bool {
	if o != nil && o.RefreshToken.IsSet() {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given NullableString and assigns it to the RefreshToken field.
func (o *OneDriveAuthentication) SetRefreshToken(v string) {
	o.RefreshToken.Set(&v)
}
// SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil
func (o *OneDriveAuthentication) SetRefreshTokenNil() {
	o.RefreshToken.Set(nil)
}

// UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
func (o *OneDriveAuthentication) UnsetRefreshToken() {
	o.RefreshToken.Unset()
}

// GetTenantName returns the TenantName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OneDriveAuthentication) GetTenantName() string {
	if o == nil || isNil(o.TenantName.Get()) {
		var ret string
		return ret
	}
	return *o.TenantName.Get()
}

// GetTenantNameOk returns a tuple with the TenantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OneDriveAuthentication) GetTenantNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.TenantName.Get(), o.TenantName.IsSet()
}

// HasTenantName returns a boolean if a field has been set.
func (o *OneDriveAuthentication) HasTenantName() bool {
	if o != nil && o.TenantName.IsSet() {
		return true
	}

	return false
}

// SetTenantName gets a reference to the given NullableString and assigns it to the TenantName field.
func (o *OneDriveAuthentication) SetTenantName(v string) {
	o.TenantName.Set(&v)
}
// SetTenantNameNil sets the value for TenantName to be an explicit nil
func (o *OneDriveAuthentication) SetTenantNameNil() {
	o.TenantName.Set(nil)
}

// UnsetTenantName ensures that no value is present for TenantName, not even an explicit nil
func (o *OneDriveAuthentication) UnsetTenantName() {
	o.TenantName.Unset()
}

func (o OneDriveAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.RefreshToken.IsSet() {
		toSerialize["refresh_token"] = o.RefreshToken.Get()
	}
	if o.TenantName.IsSet() {
		toSerialize["tenant_name"] = o.TenantName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOneDriveAuthentication struct {
	value *OneDriveAuthentication
	isSet bool
}

func (v NullableOneDriveAuthentication) Get() *OneDriveAuthentication {
	return v.value
}

func (v *NullableOneDriveAuthentication) Set(val *OneDriveAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableOneDriveAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableOneDriveAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOneDriveAuthentication(val *OneDriveAuthentication) *NullableOneDriveAuthentication {
	return &NullableOneDriveAuthentication{value: val, isSet: true}
}

func (v NullableOneDriveAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOneDriveAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



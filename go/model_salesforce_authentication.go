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

// SalesforceAuthentication struct for SalesforceAuthentication
type SalesforceAuthentication struct {
	Source interface{} `json:"source"`
	AccessToken string `json:"access_token"`
	RefreshToken NullableString `json:"refresh_token,omitempty"`
	Domain string `json:"domain"`
}

// NewSalesforceAuthentication instantiates a new SalesforceAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesforceAuthentication(source interface{}, accessToken string, domain string) *SalesforceAuthentication {
	this := SalesforceAuthentication{}
	this.Source = source
	this.AccessToken = accessToken
	this.Domain = domain
	return &this
}

// NewSalesforceAuthenticationWithDefaults instantiates a new SalesforceAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesforceAuthenticationWithDefaults() *SalesforceAuthentication {
	this := SalesforceAuthentication{}
	return &this
}

// GetSource returns the Source field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SalesforceAuthentication) GetSource() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SalesforceAuthentication) GetSourceOk() (*interface{}, bool) {
	if o == nil || isNil(o.Source) {
    return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *SalesforceAuthentication) SetSource(v interface{}) {
	o.Source = v
}

// GetAccessToken returns the AccessToken field value
func (o *SalesforceAuthentication) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *SalesforceAuthentication) GetAccessTokenOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *SalesforceAuthentication) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SalesforceAuthentication) GetRefreshToken() string {
	if o == nil || isNil(o.RefreshToken.Get()) {
		var ret string
		return ret
	}
	return *o.RefreshToken.Get()
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SalesforceAuthentication) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RefreshToken.Get(), o.RefreshToken.IsSet()
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *SalesforceAuthentication) HasRefreshToken() bool {
	if o != nil && o.RefreshToken.IsSet() {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given NullableString and assigns it to the RefreshToken field.
func (o *SalesforceAuthentication) SetRefreshToken(v string) {
	o.RefreshToken.Set(&v)
}
// SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil
func (o *SalesforceAuthentication) SetRefreshTokenNil() {
	o.RefreshToken.Set(nil)
}

// UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
func (o *SalesforceAuthentication) UnsetRefreshToken() {
	o.RefreshToken.Unset()
}

// GetDomain returns the Domain field value
func (o *SalesforceAuthentication) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *SalesforceAuthentication) GetDomainOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *SalesforceAuthentication) SetDomain(v string) {
	o.Domain = v
}

func (o SalesforceAuthentication) MarshalJSON() ([]byte, error) {
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
		toSerialize["domain"] = o.Domain
	}
	return json.Marshal(toSerialize)
}

type NullableSalesforceAuthentication struct {
	value *SalesforceAuthentication
	isSet bool
}

func (v NullableSalesforceAuthentication) Get() *SalesforceAuthentication {
	return v.value
}

func (v *NullableSalesforceAuthentication) Set(val *SalesforceAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesforceAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesforceAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesforceAuthentication(val *SalesforceAuthentication) *NullableSalesforceAuthentication {
	return &NullableSalesforceAuthentication{value: val, isSet: true}
}

func (v NullableSalesforceAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesforceAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



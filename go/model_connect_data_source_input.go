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

// ConnectDataSourceInput struct for ConnectDataSourceInput
type ConnectDataSourceInput struct {
	Authentication OANSCZGFB `json:"authentication"`
	SyncOptions *SyncOptions `json:"sync_options,omitempty"`
}

// NewConnectDataSourceInput instantiates a new ConnectDataSourceInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectDataSourceInput(authentication OANSCZGFB) *ConnectDataSourceInput {
	this := ConnectDataSourceInput{}
	this.Authentication = authentication
	return &this
}

// NewConnectDataSourceInputWithDefaults instantiates a new ConnectDataSourceInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectDataSourceInputWithDefaults() *ConnectDataSourceInput {
	this := ConnectDataSourceInput{}
	return &this
}

// GetAuthentication returns the Authentication field value
func (o *ConnectDataSourceInput) GetAuthentication() OANSCZGFB {
	if o == nil {
		var ret OANSCZGFB
		return ret
	}

	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value
// and a boolean to check if the value has been set.
func (o *ConnectDataSourceInput) GetAuthenticationOk() (*OANSCZGFB, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Authentication, true
}

// SetAuthentication sets field value
func (o *ConnectDataSourceInput) SetAuthentication(v OANSCZGFB) {
	o.Authentication = v
}

// GetSyncOptions returns the SyncOptions field value if set, zero value otherwise.
func (o *ConnectDataSourceInput) GetSyncOptions() SyncOptions {
	if o == nil || isNil(o.SyncOptions) {
		var ret SyncOptions
		return ret
	}
	return *o.SyncOptions
}

// GetSyncOptionsOk returns a tuple with the SyncOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectDataSourceInput) GetSyncOptionsOk() (*SyncOptions, bool) {
	if o == nil || isNil(o.SyncOptions) {
    return nil, false
	}
	return o.SyncOptions, true
}

// HasSyncOptions returns a boolean if a field has been set.
func (o *ConnectDataSourceInput) HasSyncOptions() bool {
	if o != nil && !isNil(o.SyncOptions) {
		return true
	}

	return false
}

// SetSyncOptions gets a reference to the given SyncOptions and assigns it to the SyncOptions field.
func (o *ConnectDataSourceInput) SetSyncOptions(v SyncOptions) {
	o.SyncOptions = &v
}

func (o ConnectDataSourceInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["authentication"] = o.Authentication
	}
	if !isNil(o.SyncOptions) {
		toSerialize["sync_options"] = o.SyncOptions
	}
	return json.Marshal(toSerialize)
}

type NullableConnectDataSourceInput struct {
	value *ConnectDataSourceInput
	isSet bool
}

func (v NullableConnectDataSourceInput) Get() *ConnectDataSourceInput {
	return v.value
}

func (v *NullableConnectDataSourceInput) Set(val *ConnectDataSourceInput) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectDataSourceInput) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectDataSourceInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectDataSourceInput(val *ConnectDataSourceInput) *NullableConnectDataSourceInput {
	return &NullableConnectDataSourceInput{value: val, isSet: true}
}

func (v NullableConnectDataSourceInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectDataSourceInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// UpdateOrganizationInput struct for UpdateOrganizationInput
type UpdateOrganizationInput struct {
	GlobalUserConfig NullableUserConfigurationNullable `json:"global_user_config,omitempty"`
	// Used to set organization level defaults for configuration related to data sources.
	DataSourceConfigs map[string]DataSourceConfiguration `json:"data_source_configs,omitempty"`
}

// NewUpdateOrganizationInput instantiates a new UpdateOrganizationInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationInput() *UpdateOrganizationInput {
	this := UpdateOrganizationInput{}
	return &this
}

// NewUpdateOrganizationInputWithDefaults instantiates a new UpdateOrganizationInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationInputWithDefaults() *UpdateOrganizationInput {
	this := UpdateOrganizationInput{}
	return &this
}

// GetGlobalUserConfig returns the GlobalUserConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOrganizationInput) GetGlobalUserConfig() UserConfigurationNullable {
	if o == nil || isNil(o.GlobalUserConfig.Get()) {
		var ret UserConfigurationNullable
		return ret
	}
	return *o.GlobalUserConfig.Get()
}

// GetGlobalUserConfigOk returns a tuple with the GlobalUserConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOrganizationInput) GetGlobalUserConfigOk() (*UserConfigurationNullable, bool) {
	if o == nil {
    return nil, false
	}
	return o.GlobalUserConfig.Get(), o.GlobalUserConfig.IsSet()
}

// HasGlobalUserConfig returns a boolean if a field has been set.
func (o *UpdateOrganizationInput) HasGlobalUserConfig() bool {
	if o != nil && o.GlobalUserConfig.IsSet() {
		return true
	}

	return false
}

// SetGlobalUserConfig gets a reference to the given NullableUserConfigurationNullable and assigns it to the GlobalUserConfig field.
func (o *UpdateOrganizationInput) SetGlobalUserConfig(v UserConfigurationNullable) {
	o.GlobalUserConfig.Set(&v)
}
// SetGlobalUserConfigNil sets the value for GlobalUserConfig to be an explicit nil
func (o *UpdateOrganizationInput) SetGlobalUserConfigNil() {
	o.GlobalUserConfig.Set(nil)
}

// UnsetGlobalUserConfig ensures that no value is present for GlobalUserConfig, not even an explicit nil
func (o *UpdateOrganizationInput) UnsetGlobalUserConfig() {
	o.GlobalUserConfig.Unset()
}

// GetDataSourceConfigs returns the DataSourceConfigs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOrganizationInput) GetDataSourceConfigs() map[string]DataSourceConfiguration {
	if o == nil {
		var ret map[string]DataSourceConfiguration
		return ret
	}
	return o.DataSourceConfigs
}

// GetDataSourceConfigsOk returns a tuple with the DataSourceConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOrganizationInput) GetDataSourceConfigsOk() (*map[string]DataSourceConfiguration, bool) {
	if o == nil || isNil(o.DataSourceConfigs) {
    return nil, false
	}
	return &o.DataSourceConfigs, true
}

// HasDataSourceConfigs returns a boolean if a field has been set.
func (o *UpdateOrganizationInput) HasDataSourceConfigs() bool {
	if o != nil && isNil(o.DataSourceConfigs) {
		return true
	}

	return false
}

// SetDataSourceConfigs gets a reference to the given map[string]DataSourceConfiguration and assigns it to the DataSourceConfigs field.
func (o *UpdateOrganizationInput) SetDataSourceConfigs(v map[string]DataSourceConfiguration) {
	o.DataSourceConfigs = v
}

func (o UpdateOrganizationInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GlobalUserConfig.IsSet() {
		toSerialize["global_user_config"] = o.GlobalUserConfig.Get()
	}
	if o.DataSourceConfigs != nil {
		toSerialize["data_source_configs"] = o.DataSourceConfigs
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateOrganizationInput struct {
	value *UpdateOrganizationInput
	isSet bool
}

func (v NullableUpdateOrganizationInput) Get() *UpdateOrganizationInput {
	return v.value
}

func (v *NullableUpdateOrganizationInput) Set(val *UpdateOrganizationInput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationInput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationInput(val *UpdateOrganizationInput) *NullableUpdateOrganizationInput {
	return &NullableUpdateOrganizationInput{value: val, isSet: true}
}

func (v NullableUpdateOrganizationInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



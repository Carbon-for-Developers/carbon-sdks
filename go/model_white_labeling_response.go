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

// WhiteLabelingResponse struct for WhiteLabelingResponse
type WhiteLabelingResponse struct {
	RemoveBranding bool `json:"remove_branding"`
	Integrations map[string]interface{} `json:"integrations"`
	CustomLimits map[string]interface{} `json:"custom_limits"`
	ConnectorSettings map[string]interface{} `json:"connector_settings,omitempty"`
}

// NewWhiteLabelingResponse instantiates a new WhiteLabelingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhiteLabelingResponse(removeBranding bool, integrations map[string]interface{}, customLimits map[string]interface{}) *WhiteLabelingResponse {
	this := WhiteLabelingResponse{}
	this.RemoveBranding = removeBranding
	this.Integrations = integrations
	this.CustomLimits = customLimits
	return &this
}

// NewWhiteLabelingResponseWithDefaults instantiates a new WhiteLabelingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhiteLabelingResponseWithDefaults() *WhiteLabelingResponse {
	this := WhiteLabelingResponse{}
	return &this
}

// GetRemoveBranding returns the RemoveBranding field value
func (o *WhiteLabelingResponse) GetRemoveBranding() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RemoveBranding
}

// GetRemoveBrandingOk returns a tuple with the RemoveBranding field value
// and a boolean to check if the value has been set.
func (o *WhiteLabelingResponse) GetRemoveBrandingOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RemoveBranding, true
}

// SetRemoveBranding sets field value
func (o *WhiteLabelingResponse) SetRemoveBranding(v bool) {
	o.RemoveBranding = v
}

// GetIntegrations returns the Integrations field value
func (o *WhiteLabelingResponse) GetIntegrations() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Integrations
}

// GetIntegrationsOk returns a tuple with the Integrations field value
// and a boolean to check if the value has been set.
func (o *WhiteLabelingResponse) GetIntegrationsOk() (map[string]interface{}, bool) {
	if o == nil {
    return map[string]interface{}{}, false
	}
	return o.Integrations, true
}

// SetIntegrations sets field value
func (o *WhiteLabelingResponse) SetIntegrations(v map[string]interface{}) {
	o.Integrations = v
}

// GetCustomLimits returns the CustomLimits field value
func (o *WhiteLabelingResponse) GetCustomLimits() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.CustomLimits
}

// GetCustomLimitsOk returns a tuple with the CustomLimits field value
// and a boolean to check if the value has been set.
func (o *WhiteLabelingResponse) GetCustomLimitsOk() (map[string]interface{}, bool) {
	if o == nil {
    return map[string]interface{}{}, false
	}
	return o.CustomLimits, true
}

// SetCustomLimits sets field value
func (o *WhiteLabelingResponse) SetCustomLimits(v map[string]interface{}) {
	o.CustomLimits = v
}

// GetConnectorSettings returns the ConnectorSettings field value if set, zero value otherwise.
func (o *WhiteLabelingResponse) GetConnectorSettings() map[string]interface{} {
	if o == nil || isNil(o.ConnectorSettings) {
		var ret map[string]interface{}
		return ret
	}
	return o.ConnectorSettings
}

// GetConnectorSettingsOk returns a tuple with the ConnectorSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhiteLabelingResponse) GetConnectorSettingsOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.ConnectorSettings) {
    return map[string]interface{}{}, false
	}
	return o.ConnectorSettings, true
}

// HasConnectorSettings returns a boolean if a field has been set.
func (o *WhiteLabelingResponse) HasConnectorSettings() bool {
	if o != nil && !isNil(o.ConnectorSettings) {
		return true
	}

	return false
}

// SetConnectorSettings gets a reference to the given map[string]interface{} and assigns it to the ConnectorSettings field.
func (o *WhiteLabelingResponse) SetConnectorSettings(v map[string]interface{}) {
	o.ConnectorSettings = v
}

func (o WhiteLabelingResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["remove_branding"] = o.RemoveBranding
	}
	if true {
		toSerialize["integrations"] = o.Integrations
	}
	if true {
		toSerialize["custom_limits"] = o.CustomLimits
	}
	if !isNil(o.ConnectorSettings) {
		toSerialize["connector_settings"] = o.ConnectorSettings
	}
	return json.Marshal(toSerialize)
}

type NullableWhiteLabelingResponse struct {
	value *WhiteLabelingResponse
	isSet bool
}

func (v NullableWhiteLabelingResponse) Get() *WhiteLabelingResponse {
	return v.value
}

func (v *NullableWhiteLabelingResponse) Set(val *WhiteLabelingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWhiteLabelingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWhiteLabelingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhiteLabelingResponse(val *WhiteLabelingResponse) *NullableWhiteLabelingResponse {
	return &NullableWhiteLabelingResponse{value: val, isSet: true}
}

func (v NullableWhiteLabelingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhiteLabelingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



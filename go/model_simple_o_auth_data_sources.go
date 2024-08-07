/*
Carbon

Connect external data to LLMs, no matter the source.

API version: 1.0.0
*/

// Code generated by Konfig (https://konfigthis.com); DO NOT EDIT.

package carbon

import (
	"encoding/json"
	"fmt"
)

// SimpleOAuthDataSources the model 'SimpleOAuthDataSources'
type SimpleOAuthDataSources string

// List of SimpleOAuthDataSources
const (
	SIMPLEOAUTHDATASOURCES_GOOGLE_DRIVE SimpleOAuthDataSources = "GOOGLE_DRIVE"
	SIMPLEOAUTHDATASOURCES_INTERCOM SimpleOAuthDataSources = "INTERCOM"
	SIMPLEOAUTHDATASOURCES_DROPBOX SimpleOAuthDataSources = "DROPBOX"
	SIMPLEOAUTHDATASOURCES_ONEDRIVE SimpleOAuthDataSources = "ONEDRIVE"
	SIMPLEOAUTHDATASOURCES_BOX SimpleOAuthDataSources = "BOX"
	SIMPLEOAUTHDATASOURCES_GMAIL SimpleOAuthDataSources = "GMAIL"
	SIMPLEOAUTHDATASOURCES_OUTLOOK SimpleOAuthDataSources = "OUTLOOK"
	SIMPLEOAUTHDATASOURCES_SLACK SimpleOAuthDataSources = "SLACK"
	SIMPLEOAUTHDATASOURCES_GOOGLE_CLOUD_STORAGE SimpleOAuthDataSources = "GOOGLE_CLOUD_STORAGE"
)

// All allowed values of SimpleOAuthDataSources enum
var AllowedSimpleOAuthDataSourcesEnumValues = []SimpleOAuthDataSources{
	"GOOGLE_DRIVE",
	"INTERCOM",
	"DROPBOX",
	"ONEDRIVE",
	"BOX",
	"GMAIL",
	"OUTLOOK",
	"SLACK",
	"GOOGLE_CLOUD_STORAGE",
}

func (v *SimpleOAuthDataSources) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SimpleOAuthDataSources(value)
	for _, existing := range AllowedSimpleOAuthDataSourcesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SimpleOAuthDataSources", value)
}

// NewSimpleOAuthDataSourcesFromValue returns a pointer to a valid SimpleOAuthDataSources
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSimpleOAuthDataSourcesFromValue(v string) (*SimpleOAuthDataSources, error) {
	ev := SimpleOAuthDataSources(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SimpleOAuthDataSources: valid values are %v", v, AllowedSimpleOAuthDataSourcesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SimpleOAuthDataSources) IsValid() bool {
	for _, existing := range AllowedSimpleOAuthDataSourcesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SimpleOAuthDataSources value
func (v SimpleOAuthDataSources) Ptr() *SimpleOAuthDataSources {
	return &v
}

type NullableSimpleOAuthDataSources struct {
	value *SimpleOAuthDataSources
	isSet bool
}

func (v NullableSimpleOAuthDataSources) Get() *SimpleOAuthDataSources {
	return v.value
}

func (v *NullableSimpleOAuthDataSources) Set(val *SimpleOAuthDataSources) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleOAuthDataSources) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleOAuthDataSources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleOAuthDataSources(val *SimpleOAuthDataSources) *NullableSimpleOAuthDataSources {
	return &NullableSimpleOAuthDataSources{value: val, isSet: true}
}

func (v NullableSimpleOAuthDataSources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleOAuthDataSources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


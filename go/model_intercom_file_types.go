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

// IntercomFileTypes the model 'IntercomFileTypes'
type IntercomFileTypes string

// List of IntercomFileTypes
const (
	INTERCOMFILETYPES_ARTICLE IntercomFileTypes = "ARTICLE"
	INTERCOMFILETYPES_COLLECTION IntercomFileTypes = "COLLECTION"
	INTERCOMFILETYPES_HELP_CENTER IntercomFileTypes = "HELP_CENTER"
	INTERCOMFILETYPES_TICKET IntercomFileTypes = "TICKET"
	INTERCOMFILETYPES_TICKET_ATTACHMENT IntercomFileTypes = "TICKET_ATTACHMENT"
	INTERCOMFILETYPES_CONTACT IntercomFileTypes = "CONTACT"
	INTERCOMFILETYPES_ADMIN IntercomFileTypes = "ADMIN"
	INTERCOMFILETYPES_TEAM IntercomFileTypes = "TEAM"
	INTERCOMFILETYPES_CONVERSATION IntercomFileTypes = "CONVERSATION"
)

// All allowed values of IntercomFileTypes enum
var AllowedIntercomFileTypesEnumValues = []IntercomFileTypes{
	"ARTICLE",
	"COLLECTION",
	"HELP_CENTER",
	"TICKET",
	"TICKET_ATTACHMENT",
	"CONTACT",
	"ADMIN",
	"TEAM",
	"CONVERSATION",
}

func (v *IntercomFileTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IntercomFileTypes(value)
	for _, existing := range AllowedIntercomFileTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IntercomFileTypes", value)
}

// NewIntercomFileTypesFromValue returns a pointer to a valid IntercomFileTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIntercomFileTypesFromValue(v string) (*IntercomFileTypes, error) {
	ev := IntercomFileTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IntercomFileTypes: valid values are %v", v, AllowedIntercomFileTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IntercomFileTypes) IsValid() bool {
	for _, existing := range AllowedIntercomFileTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IntercomFileTypes value
func (v IntercomFileTypes) Ptr() *IntercomFileTypes {
	return &v
}

type NullableIntercomFileTypes struct {
	value *IntercomFileTypes
	isSet bool
}

func (v NullableIntercomFileTypes) Get() *IntercomFileTypes {
	return v.value
}

func (v *NullableIntercomFileTypes) Set(val *IntercomFileTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableIntercomFileTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableIntercomFileTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntercomFileTypes(val *IntercomFileTypes) *NullableIntercomFileTypes {
	return &NullableIntercomFileTypes{value: val, isSet: true}
}

func (v NullableIntercomFileTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntercomFileTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

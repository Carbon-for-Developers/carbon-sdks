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

// HSNFileTypes3 the model 'HSNFileTypes3'
type HSNFileTypes3 string

// List of HSNFileTypes3
const (
	HSNFILETYPES3_TICKET HSNFileTypes3 = "TICKET"
	HSNFILETYPES3_ARTICLE HSNFileTypes3 = "ARTICLE"
	HSNFILETYPES3_CONVERSATION HSNFileTypes3 = "CONVERSATION"
)

// All allowed values of HSNFileTypes3 enum
var AllowedHSNFileTypes3EnumValues = []HSNFileTypes3{
	"TICKET",
	"ARTICLE",
	"CONVERSATION",
}

func (v *HSNFileTypes3) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HSNFileTypes3(value)
	for _, existing := range AllowedHSNFileTypes3EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HSNFileTypes3", value)
}

// NewHSNFileTypes3FromValue returns a pointer to a valid HSNFileTypes3
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHSNFileTypes3FromValue(v string) (*HSNFileTypes3, error) {
	ev := HSNFileTypes3(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HSNFileTypes3: valid values are %v", v, AllowedHSNFileTypes3EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HSNFileTypes3) IsValid() bool {
	for _, existing := range AllowedHSNFileTypes3EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HSNFileTypes3 value
func (v HSNFileTypes3) Ptr() *HSNFileTypes3 {
	return &v
}

type NullableHSNFileTypes3 struct {
	value *HSNFileTypes3
	isSet bool
}

func (v NullableHSNFileTypes3) Get() *HSNFileTypes3 {
	return v.value
}

func (v *NullableHSNFileTypes3) Set(val *HSNFileTypes3) {
	v.value = val
	v.isSet = true
}

func (v NullableHSNFileTypes3) IsSet() bool {
	return v.isSet
}

func (v *NullableHSNFileTypes3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHSNFileTypes3(val *HSNFileTypes3) *NullableHSNFileTypes3 {
	return &NullableHSNFileTypes3{value: val, isSet: true}
}

func (v NullableHSNFileTypes3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHSNFileTypes3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


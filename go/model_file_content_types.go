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

// FileContentTypes the model 'FileContentTypes'
type FileContentTypes string

// List of FileContentTypes
const (
	FILECONTENTTYPES_TEXT FileContentTypes = "TEXT"
	FILECONTENTTYPES_IMAGE FileContentTypes = "IMAGE"
	FILECONTENTTYPES_AUDIO FileContentTypes = "AUDIO"
	FILECONTENTTYPES_VIDEO FileContentTypes = "VIDEO"
)

// All allowed values of FileContentTypes enum
var AllowedFileContentTypesEnumValues = []FileContentTypes{
	"TEXT",
	"IMAGE",
	"AUDIO",
	"VIDEO",
}

func (v *FileContentTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FileContentTypes(value)
	for _, existing := range AllowedFileContentTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FileContentTypes", value)
}

// NewFileContentTypesFromValue returns a pointer to a valid FileContentTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileContentTypesFromValue(v string) (*FileContentTypes, error) {
	ev := FileContentTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileContentTypes: valid values are %v", v, AllowedFileContentTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileContentTypes) IsValid() bool {
	for _, existing := range AllowedFileContentTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FileContentTypes value
func (v FileContentTypes) Ptr() *FileContentTypes {
	return &v
}

type NullableFileContentTypes struct {
	value *FileContentTypes
	isSet bool
}

func (v NullableFileContentTypes) Get() *FileContentTypes {
	return v.value
}

func (v *NullableFileContentTypes) Set(val *FileContentTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableFileContentTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableFileContentTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileContentTypes(val *FileContentTypes) *NullableFileContentTypes {
	return &NullableFileContentTypes{value: val, isSet: true}
}

func (v NullableFileContentTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileContentTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// SourceProperty struct for SourceProperty
type SourceProperty struct {
	DataSourceType *DataSourceType
	DataSourceTypeArray *[]DataSourceType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SourceProperty) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into DataSourceType
	err = json.Unmarshal(data, &dst.DataSourceType);
	if err == nil {
		jsonDataSourceType, _ := json.Marshal(dst.DataSourceType)
		if string(jsonDataSourceType) == "{}" { // empty struct
			dst.DataSourceType = nil
		} else {
			return nil // data stored in dst.DataSourceType, return on the first match
		}
	} else {
		dst.DataSourceType = nil
	}

	// try to unmarshal JSON data into []DataSourceType
	err = json.Unmarshal(data, &dst.DataSourceTypeArray);
	if err == nil {
		jsonDataSourceTypeArray, _ := json.Marshal(dst.DataSourceTypeArray)
		if string(jsonDataSourceTypeArray) == "{}" { // empty struct
			dst.DataSourceTypeArray = nil
		} else {
			return nil // data stored in dst.DataSourceTypeArray, return on the first match
		}
	} else {
		dst.DataSourceTypeArray = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SourceProperty)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SourceProperty) MarshalJSON() ([]byte, error) {
	if src.DataSourceType != nil {
		return json.Marshal(&src.DataSourceType)
	}

	if src.DataSourceTypeArray != nil {
		return json.Marshal(&src.DataSourceTypeArray)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSourceProperty struct {
	value *SourceProperty
	isSet bool
}

func (v NullableSourceProperty) Get() *SourceProperty {
	return v.value
}

func (v *NullableSourceProperty) Set(val *SourceProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceProperty(val *SourceProperty) *NullableSourceProperty {
	return &NullableSourceProperty{value: val, isSet: true}
}

func (v NullableSourceProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



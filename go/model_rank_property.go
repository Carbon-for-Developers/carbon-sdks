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

// RankProperty struct for RankProperty
type RankProperty struct {
	Float32 *float32
	Int32 *int32
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RankProperty) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into float32
	err = json.Unmarshal(data, &dst.Float32);
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			return nil // data stored in dst.Float32, return on the first match
		}
	} else {
		dst.Float32 = nil
	}

	// try to unmarshal JSON data into int32
	err = json.Unmarshal(data, &dst.Int32);
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			return nil // data stored in dst.Int32, return on the first match
		}
	} else {
		dst.Int32 = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RankProperty)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RankProperty) MarshalJSON() ([]byte, error) {
	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRankProperty struct {
	value *RankProperty
	isSet bool
}

func (v NullableRankProperty) Get() *RankProperty {
	return v.value
}

func (v *NullableRankProperty) Set(val *RankProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableRankProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableRankProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRankProperty(val *RankProperty) *NullableRankProperty {
	return &NullableRankProperty{value: val, isSet: true}
}

func (v NullableRankProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRankProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



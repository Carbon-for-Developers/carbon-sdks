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

// OpportunitiesOrderByNullable the model 'OpportunitiesOrderByNullable'
type OpportunitiesOrderByNullable string

// List of OpportunitiesOrderByNullable
const (
	OPPORTUNITIESORDERBYNULLABLE_CREATED_AT OpportunitiesOrderByNullable = "created_at"
	OPPORTUNITIESORDERBYNULLABLE_UPDATED_AT OpportunitiesOrderByNullable = "updated_at"
	OPPORTUNITIESORDERBYNULLABLE_AMOUNT OpportunitiesOrderByNullable = "amount"
	OPPORTUNITIESORDERBYNULLABLE_NAME OpportunitiesOrderByNullable = "name"
	OPPORTUNITIESORDERBYNULLABLE_LAST_ACTIVITY_AT OpportunitiesOrderByNullable = "last_activity_at"
	OPPORTUNITIESORDERBYNULLABLE_CLOSE_DATE OpportunitiesOrderByNullable = "close_date"
)

// All allowed values of OpportunitiesOrderByNullable enum
var AllowedOpportunitiesOrderByNullableEnumValues = []OpportunitiesOrderByNullable{
	"created_at",
	"updated_at",
	"amount",
	"name",
	"last_activity_at",
	"close_date",
}

func (v *OpportunitiesOrderByNullable) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OpportunitiesOrderByNullable(value)
	for _, existing := range AllowedOpportunitiesOrderByNullableEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OpportunitiesOrderByNullable", value)
}

// NewOpportunitiesOrderByNullableFromValue returns a pointer to a valid OpportunitiesOrderByNullable
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOpportunitiesOrderByNullableFromValue(v string) (*OpportunitiesOrderByNullable, error) {
	ev := OpportunitiesOrderByNullable(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OpportunitiesOrderByNullable: valid values are %v", v, AllowedOpportunitiesOrderByNullableEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OpportunitiesOrderByNullable) IsValid() bool {
	for _, existing := range AllowedOpportunitiesOrderByNullableEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OpportunitiesOrderByNullable value
func (v OpportunitiesOrderByNullable) Ptr() *OpportunitiesOrderByNullable {
	return &v
}

type NullableOpportunitiesOrderByNullable struct {
	value *OpportunitiesOrderByNullable
	isSet bool
}

func (v NullableOpportunitiesOrderByNullable) Get() *OpportunitiesOrderByNullable {
	return v.value
}

func (v *NullableOpportunitiesOrderByNullable) Set(val *OpportunitiesOrderByNullable) {
	v.value = val
	v.isSet = true
}

func (v NullableOpportunitiesOrderByNullable) IsSet() bool {
	return v.isSet
}

func (v *NullableOpportunitiesOrderByNullable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpportunitiesOrderByNullable(val *OpportunitiesOrderByNullable) *NullableOpportunitiesOrderByNullable {
	return &NullableOpportunitiesOrderByNullable{value: val, isSet: true}
}

func (v NullableOpportunitiesOrderByNullable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpportunitiesOrderByNullable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


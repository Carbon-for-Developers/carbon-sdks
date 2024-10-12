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

// OpportunitiesOrderBy the model 'OpportunitiesOrderBy'
type OpportunitiesOrderBy string

// List of OpportunitiesOrderBy
const (
	OPPORTUNITIESORDERBY_CREATED_AT OpportunitiesOrderBy = "created_at"
	OPPORTUNITIESORDERBY_UPDATED_AT OpportunitiesOrderBy = "updated_at"
	OPPORTUNITIESORDERBY_AMOUNT OpportunitiesOrderBy = "amount"
	OPPORTUNITIESORDERBY_NAME OpportunitiesOrderBy = "name"
	OPPORTUNITIESORDERBY_LAST_ACTIVITY_AT OpportunitiesOrderBy = "last_activity_at"
	OPPORTUNITIESORDERBY_CLOSE_DATE OpportunitiesOrderBy = "close_date"
)

// All allowed values of OpportunitiesOrderBy enum
var AllowedOpportunitiesOrderByEnumValues = []OpportunitiesOrderBy{
	"created_at",
	"updated_at",
	"amount",
	"name",
	"last_activity_at",
	"close_date",
}

func (v *OpportunitiesOrderBy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OpportunitiesOrderBy(value)
	for _, existing := range AllowedOpportunitiesOrderByEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OpportunitiesOrderBy", value)
}

// NewOpportunitiesOrderByFromValue returns a pointer to a valid OpportunitiesOrderBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOpportunitiesOrderByFromValue(v string) (*OpportunitiesOrderBy, error) {
	ev := OpportunitiesOrderBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OpportunitiesOrderBy: valid values are %v", v, AllowedOpportunitiesOrderByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OpportunitiesOrderBy) IsValid() bool {
	for _, existing := range AllowedOpportunitiesOrderByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OpportunitiesOrderBy value
func (v OpportunitiesOrderBy) Ptr() *OpportunitiesOrderBy {
	return &v
}

type NullableOpportunitiesOrderBy struct {
	value *OpportunitiesOrderBy
	isSet bool
}

func (v NullableOpportunitiesOrderBy) Get() *OpportunitiesOrderBy {
	return v.value
}

func (v *NullableOpportunitiesOrderBy) Set(val *OpportunitiesOrderBy) {
	v.value = val
	v.isSet = true
}

func (v NullableOpportunitiesOrderBy) IsSet() bool {
	return v.isSet
}

func (v *NullableOpportunitiesOrderBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpportunitiesOrderBy(val *OpportunitiesOrderBy) *NullableOpportunitiesOrderBy {
	return &NullableOpportunitiesOrderBy{value: val, isSet: true}
}

func (v NullableOpportunitiesOrderBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpportunitiesOrderBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

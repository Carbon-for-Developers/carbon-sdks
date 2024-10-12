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

// PROrderBy the model 'PROrderBy'
type PROrderBy string

// List of PROrderBy
const (
	PRORDERBY_CREATED PROrderBy = "created"
	PRORDERBY_UPDATED PROrderBy = "updated"
	PRORDERBY_POPULARITY PROrderBy = "popularity"
)

// All allowed values of PROrderBy enum
var AllowedPROrderByEnumValues = []PROrderBy{
	"created",
	"updated",
	"popularity",
}

func (v *PROrderBy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PROrderBy(value)
	for _, existing := range AllowedPROrderByEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PROrderBy", value)
}

// NewPROrderByFromValue returns a pointer to a valid PROrderBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPROrderByFromValue(v string) (*PROrderBy, error) {
	ev := PROrderBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PROrderBy: valid values are %v", v, AllowedPROrderByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PROrderBy) IsValid() bool {
	for _, existing := range AllowedPROrderByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PROrderBy value
func (v PROrderBy) Ptr() *PROrderBy {
	return &v
}

type NullablePROrderBy struct {
	value *PROrderBy
	isSet bool
}

func (v NullablePROrderBy) Get() *PROrderBy {
	return v.value
}

func (v *NullablePROrderBy) Set(val *PROrderBy) {
	v.value = val
	v.isSet = true
}

func (v NullablePROrderBy) IsSet() bool {
	return v.isSet
}

func (v *NullablePROrderBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePROrderBy(val *PROrderBy) *NullablePROrderBy {
	return &NullablePROrderBy{value: val, isSet: true}
}

func (v NullablePROrderBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePROrderBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

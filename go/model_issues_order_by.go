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

// IssuesOrderBy the model 'IssuesOrderBy'
type IssuesOrderBy string

// List of IssuesOrderBy
const (
	ISSUESORDERBY_CREATED IssuesOrderBy = "created"
	ISSUESORDERBY_UPDATED IssuesOrderBy = "updated"
	ISSUESORDERBY_COMMENTS IssuesOrderBy = "comments"
)

// All allowed values of IssuesOrderBy enum
var AllowedIssuesOrderByEnumValues = []IssuesOrderBy{
	"created",
	"updated",
	"comments",
}

func (v *IssuesOrderBy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IssuesOrderBy(value)
	for _, existing := range AllowedIssuesOrderByEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IssuesOrderBy", value)
}

// NewIssuesOrderByFromValue returns a pointer to a valid IssuesOrderBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIssuesOrderByFromValue(v string) (*IssuesOrderBy, error) {
	ev := IssuesOrderBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IssuesOrderBy: valid values are %v", v, AllowedIssuesOrderByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IssuesOrderBy) IsValid() bool {
	for _, existing := range AllowedIssuesOrderByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssuesOrderBy value
func (v IssuesOrderBy) Ptr() *IssuesOrderBy {
	return &v
}

type NullableIssuesOrderBy struct {
	value *IssuesOrderBy
	isSet bool
}

func (v NullableIssuesOrderBy) Get() *IssuesOrderBy {
	return v.value
}

func (v *NullableIssuesOrderBy) Set(val *IssuesOrderBy) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuesOrderBy) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuesOrderBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuesOrderBy(val *IssuesOrderBy) *NullableIssuesOrderBy {
	return &NullableIssuesOrderBy{value: val, isSet: true}
}

func (v NullableIssuesOrderBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuesOrderBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

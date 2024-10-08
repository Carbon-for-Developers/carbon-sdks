/*
Carbon

Connect external data to LLMs, no matter the source.

API version: 1.0.0
*/

// Code generated by Konfig (https://konfigthis.com); DO NOT EDIT.

package carbon

import (
	"encoding/json"
)

// PullRequestFilters struct for PullRequestFilters
type PullRequestFilters struct {
	State *PRStateInput `json:"state,omitempty"`
	// Filter pulls by base branch name
	Base NullableString `json:"base,omitempty"`
	// Filter pulls by head user or head organization and branch name in the format of user:ref-name or organization:ref-name
	Head NullableString `json:"head,omitempty"`
}

// NewPullRequestFilters instantiates a new PullRequestFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPullRequestFilters() *PullRequestFilters {
	this := PullRequestFilters{}
	return &this
}

// NewPullRequestFiltersWithDefaults instantiates a new PullRequestFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPullRequestFiltersWithDefaults() *PullRequestFilters {
	this := PullRequestFilters{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *PullRequestFilters) GetState() PRStateInput {
	if o == nil || isNil(o.State) {
		var ret PRStateInput
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequestFilters) GetStateOk() (*PRStateInput, bool) {
	if o == nil || isNil(o.State) {
    return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PullRequestFilters) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given PRStateInput and assigns it to the State field.
func (o *PullRequestFilters) SetState(v PRStateInput) {
	o.State = &v
}

// GetBase returns the Base field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PullRequestFilters) GetBase() string {
	if o == nil || isNil(o.Base.Get()) {
		var ret string
		return ret
	}
	return *o.Base.Get()
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PullRequestFilters) GetBaseOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Base.Get(), o.Base.IsSet()
}

// HasBase returns a boolean if a field has been set.
func (o *PullRequestFilters) HasBase() bool {
	if o != nil && o.Base.IsSet() {
		return true
	}

	return false
}

// SetBase gets a reference to the given NullableString and assigns it to the Base field.
func (o *PullRequestFilters) SetBase(v string) {
	o.Base.Set(&v)
}
// SetBaseNil sets the value for Base to be an explicit nil
func (o *PullRequestFilters) SetBaseNil() {
	o.Base.Set(nil)
}

// UnsetBase ensures that no value is present for Base, not even an explicit nil
func (o *PullRequestFilters) UnsetBase() {
	o.Base.Unset()
}

// GetHead returns the Head field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PullRequestFilters) GetHead() string {
	if o == nil || isNil(o.Head.Get()) {
		var ret string
		return ret
	}
	return *o.Head.Get()
}

// GetHeadOk returns a tuple with the Head field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PullRequestFilters) GetHeadOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Head.Get(), o.Head.IsSet()
}

// HasHead returns a boolean if a field has been set.
func (o *PullRequestFilters) HasHead() bool {
	if o != nil && o.Head.IsSet() {
		return true
	}

	return false
}

// SetHead gets a reference to the given NullableString and assigns it to the Head field.
func (o *PullRequestFilters) SetHead(v string) {
	o.Head.Set(&v)
}
// SetHeadNil sets the value for Head to be an explicit nil
func (o *PullRequestFilters) SetHeadNil() {
	o.Head.Set(nil)
}

// UnsetHead ensures that no value is present for Head, not even an explicit nil
func (o *PullRequestFilters) UnsetHead() {
	o.Head.Unset()
}

func (o PullRequestFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if o.Base.IsSet() {
		toSerialize["base"] = o.Base.Get()
	}
	if o.Head.IsSet() {
		toSerialize["head"] = o.Head.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePullRequestFilters struct {
	value *PullRequestFilters
	isSet bool
}

func (v NullablePullRequestFilters) Get() *PullRequestFilters {
	return v.value
}

func (v *NullablePullRequestFilters) Set(val *PullRequestFilters) {
	v.value = val
	v.isSet = true
}

func (v NullablePullRequestFilters) IsSet() bool {
	return v.isSet
}

func (v *NullablePullRequestFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePullRequestFilters(val *PullRequestFilters) *NullablePullRequestFilters {
	return &NullablePullRequestFilters{value: val, isSet: true}
}

func (v NullablePullRequestFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePullRequestFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



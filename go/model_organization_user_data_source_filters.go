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

// OrganizationUserDataSourceFilters struct for OrganizationUserDataSourceFilters
type OrganizationUserDataSourceFilters struct {
	Source NullableDataSourceTypeNullable `json:"source,omitempty"`
	Ids []int32 `json:"ids,omitempty"`
	RevokedAccess NullableBool `json:"revoked_access,omitempty"`
}

// NewOrganizationUserDataSourceFilters instantiates a new OrganizationUserDataSourceFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationUserDataSourceFilters() *OrganizationUserDataSourceFilters {
	this := OrganizationUserDataSourceFilters{}
	return &this
}

// NewOrganizationUserDataSourceFiltersWithDefaults instantiates a new OrganizationUserDataSourceFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationUserDataSourceFiltersWithDefaults() *OrganizationUserDataSourceFilters {
	this := OrganizationUserDataSourceFilters{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationUserDataSourceFilters) GetSource() DataSourceTypeNullable {
	if o == nil || isNil(o.Source.Get()) {
		var ret DataSourceTypeNullable
		return ret
	}
	return *o.Source.Get()
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationUserDataSourceFilters) GetSourceOk() (*DataSourceTypeNullable, bool) {
	if o == nil {
    return nil, false
	}
	return o.Source.Get(), o.Source.IsSet()
}

// HasSource returns a boolean if a field has been set.
func (o *OrganizationUserDataSourceFilters) HasSource() bool {
	if o != nil && o.Source.IsSet() {
		return true
	}

	return false
}

// SetSource gets a reference to the given NullableDataSourceTypeNullable and assigns it to the Source field.
func (o *OrganizationUserDataSourceFilters) SetSource(v DataSourceTypeNullable) {
	o.Source.Set(&v)
}
// SetSourceNil sets the value for Source to be an explicit nil
func (o *OrganizationUserDataSourceFilters) SetSourceNil() {
	o.Source.Set(nil)
}

// UnsetSource ensures that no value is present for Source, not even an explicit nil
func (o *OrganizationUserDataSourceFilters) UnsetSource() {
	o.Source.Unset()
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationUserDataSourceFilters) GetIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationUserDataSourceFilters) GetIdsOk() ([]int32, bool) {
	if o == nil || isNil(o.Ids) {
    return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *OrganizationUserDataSourceFilters) HasIds() bool {
	if o != nil && isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []int32 and assigns it to the Ids field.
func (o *OrganizationUserDataSourceFilters) SetIds(v []int32) {
	o.Ids = v
}

// GetRevokedAccess returns the RevokedAccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationUserDataSourceFilters) GetRevokedAccess() bool {
	if o == nil || isNil(o.RevokedAccess.Get()) {
		var ret bool
		return ret
	}
	return *o.RevokedAccess.Get()
}

// GetRevokedAccessOk returns a tuple with the RevokedAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationUserDataSourceFilters) GetRevokedAccessOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.RevokedAccess.Get(), o.RevokedAccess.IsSet()
}

// HasRevokedAccess returns a boolean if a field has been set.
func (o *OrganizationUserDataSourceFilters) HasRevokedAccess() bool {
	if o != nil && o.RevokedAccess.IsSet() {
		return true
	}

	return false
}

// SetRevokedAccess gets a reference to the given NullableBool and assigns it to the RevokedAccess field.
func (o *OrganizationUserDataSourceFilters) SetRevokedAccess(v bool) {
	o.RevokedAccess.Set(&v)
}
// SetRevokedAccessNil sets the value for RevokedAccess to be an explicit nil
func (o *OrganizationUserDataSourceFilters) SetRevokedAccessNil() {
	o.RevokedAccess.Set(nil)
}

// UnsetRevokedAccess ensures that no value is present for RevokedAccess, not even an explicit nil
func (o *OrganizationUserDataSourceFilters) UnsetRevokedAccess() {
	o.RevokedAccess.Unset()
}

func (o OrganizationUserDataSourceFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Source.IsSet() {
		toSerialize["source"] = o.Source.Get()
	}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.RevokedAccess.IsSet() {
		toSerialize["revoked_access"] = o.RevokedAccess.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationUserDataSourceFilters struct {
	value *OrganizationUserDataSourceFilters
	isSet bool
}

func (v NullableOrganizationUserDataSourceFilters) Get() *OrganizationUserDataSourceFilters {
	return v.value
}

func (v *NullableOrganizationUserDataSourceFilters) Set(val *OrganizationUserDataSourceFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationUserDataSourceFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationUserDataSourceFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationUserDataSourceFilters(val *OrganizationUserDataSourceFilters) *NullableOrganizationUserDataSourceFilters {
	return &NullableOrganizationUserDataSourceFilters{value: val, isSet: true}
}

func (v NullableOrganizationUserDataSourceFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationUserDataSourceFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


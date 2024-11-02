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

// ListItemsFilters struct for ListItemsFilters
type ListItemsFilters struct {
	ExternalIds []string `json:"external_ids,omitempty"`
	Ids []int32 `json:"ids,omitempty"`
	Name NullableString `json:"name,omitempty"`
	RootFilesOnly NullableBool `json:"root_files_only,omitempty"`
	FileFormats []StorageFileFormats `json:"file_formats,omitempty"`
	ItemTypes []SIHFNZCGABTypes `json:"item_types,omitempty"`
}

// NewListItemsFilters instantiates a new ListItemsFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListItemsFilters() *ListItemsFilters {
	this := ListItemsFilters{}
	return &this
}

// NewListItemsFiltersWithDefaults instantiates a new ListItemsFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListItemsFiltersWithDefaults() *ListItemsFilters {
	this := ListItemsFilters{}
	return &this
}

// GetExternalIds returns the ExternalIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListItemsFilters) GetExternalIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExternalIds
}

// GetExternalIdsOk returns a tuple with the ExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListItemsFilters) GetExternalIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ExternalIds) {
    return nil, false
	}
	return o.ExternalIds, true
}

// HasExternalIds returns a boolean if a field has been set.
func (o *ListItemsFilters) HasExternalIds() bool {
	if o != nil && isNil(o.ExternalIds) {
		return true
	}

	return false
}

// SetExternalIds gets a reference to the given []string and assigns it to the ExternalIds field.
func (o *ListItemsFilters) SetExternalIds(v []string) {
	o.ExternalIds = v
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListItemsFilters) GetIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListItemsFilters) GetIdsOk() ([]int32, bool) {
	if o == nil || isNil(o.Ids) {
    return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *ListItemsFilters) HasIds() bool {
	if o != nil && isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []int32 and assigns it to the Ids field.
func (o *ListItemsFilters) SetIds(v []int32) {
	o.Ids = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListItemsFilters) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListItemsFilters) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ListItemsFilters) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ListItemsFilters) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ListItemsFilters) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ListItemsFilters) UnsetName() {
	o.Name.Unset()
}

// GetRootFilesOnly returns the RootFilesOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListItemsFilters) GetRootFilesOnly() bool {
	if o == nil || isNil(o.RootFilesOnly.Get()) {
		var ret bool
		return ret
	}
	return *o.RootFilesOnly.Get()
}

// GetRootFilesOnlyOk returns a tuple with the RootFilesOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListItemsFilters) GetRootFilesOnlyOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.RootFilesOnly.Get(), o.RootFilesOnly.IsSet()
}

// HasRootFilesOnly returns a boolean if a field has been set.
func (o *ListItemsFilters) HasRootFilesOnly() bool {
	if o != nil && o.RootFilesOnly.IsSet() {
		return true
	}

	return false
}

// SetRootFilesOnly gets a reference to the given NullableBool and assigns it to the RootFilesOnly field.
func (o *ListItemsFilters) SetRootFilesOnly(v bool) {
	o.RootFilesOnly.Set(&v)
}
// SetRootFilesOnlyNil sets the value for RootFilesOnly to be an explicit nil
func (o *ListItemsFilters) SetRootFilesOnlyNil() {
	o.RootFilesOnly.Set(nil)
}

// UnsetRootFilesOnly ensures that no value is present for RootFilesOnly, not even an explicit nil
func (o *ListItemsFilters) UnsetRootFilesOnly() {
	o.RootFilesOnly.Unset()
}

// GetFileFormats returns the FileFormats field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListItemsFilters) GetFileFormats() []StorageFileFormats {
	if o == nil {
		var ret []StorageFileFormats
		return ret
	}
	return o.FileFormats
}

// GetFileFormatsOk returns a tuple with the FileFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListItemsFilters) GetFileFormatsOk() ([]StorageFileFormats, bool) {
	if o == nil || isNil(o.FileFormats) {
    return nil, false
	}
	return o.FileFormats, true
}

// HasFileFormats returns a boolean if a field has been set.
func (o *ListItemsFilters) HasFileFormats() bool {
	if o != nil && isNil(o.FileFormats) {
		return true
	}

	return false
}

// SetFileFormats gets a reference to the given []StorageFileFormats and assigns it to the FileFormats field.
func (o *ListItemsFilters) SetFileFormats(v []StorageFileFormats) {
	o.FileFormats = v
}

// GetItemTypes returns the ItemTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListItemsFilters) GetItemTypes() []SIHFNZCGABTypes {
	if o == nil {
		var ret []SIHFNZCGABTypes
		return ret
	}
	return o.ItemTypes
}

// GetItemTypesOk returns a tuple with the ItemTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListItemsFilters) GetItemTypesOk() ([]SIHFNZCGABTypes, bool) {
	if o == nil || isNil(o.ItemTypes) {
    return nil, false
	}
	return o.ItemTypes, true
}

// HasItemTypes returns a boolean if a field has been set.
func (o *ListItemsFilters) HasItemTypes() bool {
	if o != nil && isNil(o.ItemTypes) {
		return true
	}

	return false
}

// SetItemTypes gets a reference to the given []SIHFNZCGABTypes and assigns it to the ItemTypes field.
func (o *ListItemsFilters) SetItemTypes(v []SIHFNZCGABTypes) {
	o.ItemTypes = v
}

func (o ListItemsFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalIds != nil {
		toSerialize["external_ids"] = o.ExternalIds
	}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.RootFilesOnly.IsSet() {
		toSerialize["root_files_only"] = o.RootFilesOnly.Get()
	}
	if o.FileFormats != nil {
		toSerialize["file_formats"] = o.FileFormats
	}
	if o.ItemTypes != nil {
		toSerialize["item_types"] = o.ItemTypes
	}
	return json.Marshal(toSerialize)
}

type NullableListItemsFilters struct {
	value *ListItemsFilters
	isSet bool
}

func (v NullableListItemsFilters) Get() *ListItemsFilters {
	return v.value
}

func (v *NullableListItemsFilters) Set(val *ListItemsFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableListItemsFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableListItemsFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListItemsFilters(val *ListItemsFilters) *NullableListItemsFilters {
	return &NullableListItemsFilters{value: val, isSet: true}
}

func (v NullableListItemsFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListItemsFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



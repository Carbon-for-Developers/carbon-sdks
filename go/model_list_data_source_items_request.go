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

// ListDataSourceItemsRequest struct for ListDataSourceItemsRequest
type ListDataSourceItemsRequest struct {
	DataSourceId int32 `json:"data_source_id"`
	ParentId NullableString `json:"parent_id,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// NewListDataSourceItemsRequest instantiates a new ListDataSourceItemsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDataSourceItemsRequest(dataSourceId int32) *ListDataSourceItemsRequest {
	this := ListDataSourceItemsRequest{}
	this.DataSourceId = dataSourceId
	return &this
}

// NewListDataSourceItemsRequestWithDefaults instantiates a new ListDataSourceItemsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDataSourceItemsRequestWithDefaults() *ListDataSourceItemsRequest {
	this := ListDataSourceItemsRequest{}
	return &this
}

// GetDataSourceId returns the DataSourceId field value
func (o *ListDataSourceItemsRequest) GetDataSourceId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DataSourceId
}

// GetDataSourceIdOk returns a tuple with the DataSourceId field value
// and a boolean to check if the value has been set.
func (o *ListDataSourceItemsRequest) GetDataSourceIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DataSourceId, true
}

// SetDataSourceId sets field value
func (o *ListDataSourceItemsRequest) SetDataSourceId(v int32) {
	o.DataSourceId = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListDataSourceItemsRequest) GetParentId() string {
	if o == nil || isNil(o.ParentId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListDataSourceItemsRequest) GetParentIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *ListDataSourceItemsRequest) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableString and assigns it to the ParentId field.
func (o *ListDataSourceItemsRequest) SetParentId(v string) {
	o.ParentId.Set(&v)
}
// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *ListDataSourceItemsRequest) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *ListDataSourceItemsRequest) UnsetParentId() {
	o.ParentId.Unset()
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListDataSourceItemsRequest) GetPagination() Pagination {
	if o == nil || isNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDataSourceItemsRequest) GetPaginationOk() (*Pagination, bool) {
	if o == nil || isNil(o.Pagination) {
    return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListDataSourceItemsRequest) HasPagination() bool {
	if o != nil && !isNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListDataSourceItemsRequest) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ListDataSourceItemsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data_source_id"] = o.DataSourceId
	}
	if o.ParentId.IsSet() {
		toSerialize["parent_id"] = o.ParentId.Get()
	}
	if !isNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableListDataSourceItemsRequest struct {
	value *ListDataSourceItemsRequest
	isSet bool
}

func (v NullableListDataSourceItemsRequest) Get() *ListDataSourceItemsRequest {
	return v.value
}

func (v *NullableListDataSourceItemsRequest) Set(val *ListDataSourceItemsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListDataSourceItemsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListDataSourceItemsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDataSourceItemsRequest(val *ListDataSourceItemsRequest) *NullableListDataSourceItemsRequest {
	return &NullableListDataSourceItemsRequest{value: val, isSet: true}
}

func (v NullableListDataSourceItemsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDataSourceItemsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// IssuesInput struct for IssuesInput
type IssuesInput struct {
	DataSourceId int32 `json:"data_source_id"`
	IncludeRemoteData *bool `json:"include_remote_data,omitempty"`
	// Full name of the repository, denoted as {owner}/{repo}
	Repository string `json:"repository"`
	Page *int32 `json:"page,omitempty"`
	PageSize *int32 `json:"page_size,omitempty"`
	NextCursor NullableString `json:"next_cursor,omitempty"`
	Filters *IssuesFilter `json:"filters,omitempty"`
	OrderBy *IssuesOrderBy `json:"order_by,omitempty"`
	OrderDir NullableOrderDirV2Nullable `json:"order_dir,omitempty"`
}

// NewIssuesInput instantiates a new IssuesInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuesInput(dataSourceId int32, repository string) *IssuesInput {
	this := IssuesInput{}
	this.DataSourceId = dataSourceId
	var includeRemoteData bool = false
	this.IncludeRemoteData = &includeRemoteData
	this.Repository = repository
	var page int32 = 1
	this.Page = &page
	var pageSize int32 = 30
	this.PageSize = &pageSize
	var orderDir OrderDirV2Nullable = ORDERDIRV2NULLABLE_ASC
	this.OrderDir = *NewNullableOrderDirV2Nullable(&orderDir)
	return &this
}

// NewIssuesInputWithDefaults instantiates a new IssuesInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuesInputWithDefaults() *IssuesInput {
	this := IssuesInput{}
	var includeRemoteData bool = false
	this.IncludeRemoteData = &includeRemoteData
	var page int32 = 1
	this.Page = &page
	var pageSize int32 = 30
	this.PageSize = &pageSize
	var orderDir OrderDirV2Nullable = ORDERDIRV2NULLABLE_ASC
	this.OrderDir = *NewNullableOrderDirV2Nullable(&orderDir)
	return &this
}

// GetDataSourceId returns the DataSourceId field value
func (o *IssuesInput) GetDataSourceId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DataSourceId
}

// GetDataSourceIdOk returns a tuple with the DataSourceId field value
// and a boolean to check if the value has been set.
func (o *IssuesInput) GetDataSourceIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DataSourceId, true
}

// SetDataSourceId sets field value
func (o *IssuesInput) SetDataSourceId(v int32) {
	o.DataSourceId = v
}

// GetIncludeRemoteData returns the IncludeRemoteData field value if set, zero value otherwise.
func (o *IssuesInput) GetIncludeRemoteData() bool {
	if o == nil || isNil(o.IncludeRemoteData) {
		var ret bool
		return ret
	}
	return *o.IncludeRemoteData
}

// GetIncludeRemoteDataOk returns a tuple with the IncludeRemoteData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuesInput) GetIncludeRemoteDataOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeRemoteData) {
    return nil, false
	}
	return o.IncludeRemoteData, true
}

// HasIncludeRemoteData returns a boolean if a field has been set.
func (o *IssuesInput) HasIncludeRemoteData() bool {
	if o != nil && !isNil(o.IncludeRemoteData) {
		return true
	}

	return false
}

// SetIncludeRemoteData gets a reference to the given bool and assigns it to the IncludeRemoteData field.
func (o *IssuesInput) SetIncludeRemoteData(v bool) {
	o.IncludeRemoteData = &v
}

// GetRepository returns the Repository field value
func (o *IssuesInput) GetRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *IssuesInput) GetRepositoryOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *IssuesInput) SetRepository(v string) {
	o.Repository = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *IssuesInput) GetPage() int32 {
	if o == nil || isNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuesInput) GetPageOk() (*int32, bool) {
	if o == nil || isNil(o.Page) {
    return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *IssuesInput) HasPage() bool {
	if o != nil && !isNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *IssuesInput) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *IssuesInput) GetPageSize() int32 {
	if o == nil || isNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuesInput) GetPageSizeOk() (*int32, bool) {
	if o == nil || isNil(o.PageSize) {
    return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *IssuesInput) HasPageSize() bool {
	if o != nil && !isNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *IssuesInput) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuesInput) GetNextCursor() string {
	if o == nil || isNil(o.NextCursor.Get()) {
		var ret string
		return ret
	}
	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuesInput) GetNextCursorOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// HasNextCursor returns a boolean if a field has been set.
func (o *IssuesInput) HasNextCursor() bool {
	if o != nil && o.NextCursor.IsSet() {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given NullableString and assigns it to the NextCursor field.
func (o *IssuesInput) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}
// SetNextCursorNil sets the value for NextCursor to be an explicit nil
func (o *IssuesInput) SetNextCursorNil() {
	o.NextCursor.Set(nil)
}

// UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
func (o *IssuesInput) UnsetNextCursor() {
	o.NextCursor.Unset()
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *IssuesInput) GetFilters() IssuesFilter {
	if o == nil || isNil(o.Filters) {
		var ret IssuesFilter
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuesInput) GetFiltersOk() (*IssuesFilter, bool) {
	if o == nil || isNil(o.Filters) {
    return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *IssuesInput) HasFilters() bool {
	if o != nil && !isNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given IssuesFilter and assigns it to the Filters field.
func (o *IssuesInput) SetFilters(v IssuesFilter) {
	o.Filters = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *IssuesInput) GetOrderBy() IssuesOrderBy {
	if o == nil || isNil(o.OrderBy) {
		var ret IssuesOrderBy
		return ret
	}
	return *o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuesInput) GetOrderByOk() (*IssuesOrderBy, bool) {
	if o == nil || isNil(o.OrderBy) {
    return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *IssuesInput) HasOrderBy() bool {
	if o != nil && !isNil(o.OrderBy) {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given IssuesOrderBy and assigns it to the OrderBy field.
func (o *IssuesInput) SetOrderBy(v IssuesOrderBy) {
	o.OrderBy = &v
}

// GetOrderDir returns the OrderDir field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuesInput) GetOrderDir() OrderDirV2Nullable {
	if o == nil || isNil(o.OrderDir.Get()) {
		var ret OrderDirV2Nullable
		return ret
	}
	return *o.OrderDir.Get()
}

// GetOrderDirOk returns a tuple with the OrderDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuesInput) GetOrderDirOk() (*OrderDirV2Nullable, bool) {
	if o == nil {
    return nil, false
	}
	return o.OrderDir.Get(), o.OrderDir.IsSet()
}

// HasOrderDir returns a boolean if a field has been set.
func (o *IssuesInput) HasOrderDir() bool {
	if o != nil && o.OrderDir.IsSet() {
		return true
	}

	return false
}

// SetOrderDir gets a reference to the given NullableOrderDirV2Nullable and assigns it to the OrderDir field.
func (o *IssuesInput) SetOrderDir(v OrderDirV2Nullable) {
	o.OrderDir.Set(&v)
}
// SetOrderDirNil sets the value for OrderDir to be an explicit nil
func (o *IssuesInput) SetOrderDirNil() {
	o.OrderDir.Set(nil)
}

// UnsetOrderDir ensures that no value is present for OrderDir, not even an explicit nil
func (o *IssuesInput) UnsetOrderDir() {
	o.OrderDir.Unset()
}

func (o IssuesInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data_source_id"] = o.DataSourceId
	}
	if !isNil(o.IncludeRemoteData) {
		toSerialize["include_remote_data"] = o.IncludeRemoteData
	}
	if true {
		toSerialize["repository"] = o.Repository
	}
	if !isNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !isNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if o.NextCursor.IsSet() {
		toSerialize["next_cursor"] = o.NextCursor.Get()
	}
	if !isNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !isNil(o.OrderBy) {
		toSerialize["order_by"] = o.OrderBy
	}
	if o.OrderDir.IsSet() {
		toSerialize["order_dir"] = o.OrderDir.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIssuesInput struct {
	value *IssuesInput
	isSet bool
}

func (v NullableIssuesInput) Get() *IssuesInput {
	return v.value
}

func (v *NullableIssuesInput) Set(val *IssuesInput) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuesInput) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuesInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuesInput(val *IssuesInput) *NullableIssuesInput {
	return &NullableIssuesInput{value: val, isSet: true}
}

func (v NullableIssuesInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuesInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



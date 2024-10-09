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

// CommitsInput struct for CommitsInput
type CommitsInput struct {
	DataSourceId int32 `json:"data_source_id"`
	IncludeRemoteData *bool `json:"include_remote_data,omitempty"`
	// Full name of the repository, denoted as {owner}/{repo}
	Repository string `json:"repository"`
	Page *int32 `json:"page,omitempty"`
	PageSize *int32 `json:"page_size,omitempty"`
	NextCursor NullableString `json:"next_cursor,omitempty"`
	PullNumber int32 `json:"pull_number"`
}

// NewCommitsInput instantiates a new CommitsInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitsInput(dataSourceId int32, repository string, pullNumber int32) *CommitsInput {
	this := CommitsInput{}
	this.DataSourceId = dataSourceId
	var includeRemoteData bool = false
	this.IncludeRemoteData = &includeRemoteData
	this.Repository = repository
	var page int32 = 1
	this.Page = &page
	var pageSize int32 = 30
	this.PageSize = &pageSize
	this.PullNumber = pullNumber
	return &this
}

// NewCommitsInputWithDefaults instantiates a new CommitsInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitsInputWithDefaults() *CommitsInput {
	this := CommitsInput{}
	var includeRemoteData bool = false
	this.IncludeRemoteData = &includeRemoteData
	var page int32 = 1
	this.Page = &page
	var pageSize int32 = 30
	this.PageSize = &pageSize
	return &this
}

// GetDataSourceId returns the DataSourceId field value
func (o *CommitsInput) GetDataSourceId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DataSourceId
}

// GetDataSourceIdOk returns a tuple with the DataSourceId field value
// and a boolean to check if the value has been set.
func (o *CommitsInput) GetDataSourceIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DataSourceId, true
}

// SetDataSourceId sets field value
func (o *CommitsInput) SetDataSourceId(v int32) {
	o.DataSourceId = v
}

// GetIncludeRemoteData returns the IncludeRemoteData field value if set, zero value otherwise.
func (o *CommitsInput) GetIncludeRemoteData() bool {
	if o == nil || isNil(o.IncludeRemoteData) {
		var ret bool
		return ret
	}
	return *o.IncludeRemoteData
}

// GetIncludeRemoteDataOk returns a tuple with the IncludeRemoteData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitsInput) GetIncludeRemoteDataOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeRemoteData) {
    return nil, false
	}
	return o.IncludeRemoteData, true
}

// HasIncludeRemoteData returns a boolean if a field has been set.
func (o *CommitsInput) HasIncludeRemoteData() bool {
	if o != nil && !isNil(o.IncludeRemoteData) {
		return true
	}

	return false
}

// SetIncludeRemoteData gets a reference to the given bool and assigns it to the IncludeRemoteData field.
func (o *CommitsInput) SetIncludeRemoteData(v bool) {
	o.IncludeRemoteData = &v
}

// GetRepository returns the Repository field value
func (o *CommitsInput) GetRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *CommitsInput) GetRepositoryOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *CommitsInput) SetRepository(v string) {
	o.Repository = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *CommitsInput) GetPage() int32 {
	if o == nil || isNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitsInput) GetPageOk() (*int32, bool) {
	if o == nil || isNil(o.Page) {
    return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *CommitsInput) HasPage() bool {
	if o != nil && !isNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *CommitsInput) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *CommitsInput) GetPageSize() int32 {
	if o == nil || isNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitsInput) GetPageSizeOk() (*int32, bool) {
	if o == nil || isNil(o.PageSize) {
    return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *CommitsInput) HasPageSize() bool {
	if o != nil && !isNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *CommitsInput) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommitsInput) GetNextCursor() string {
	if o == nil || isNil(o.NextCursor.Get()) {
		var ret string
		return ret
	}
	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommitsInput) GetNextCursorOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// HasNextCursor returns a boolean if a field has been set.
func (o *CommitsInput) HasNextCursor() bool {
	if o != nil && o.NextCursor.IsSet() {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given NullableString and assigns it to the NextCursor field.
func (o *CommitsInput) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}
// SetNextCursorNil sets the value for NextCursor to be an explicit nil
func (o *CommitsInput) SetNextCursorNil() {
	o.NextCursor.Set(nil)
}

// UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
func (o *CommitsInput) UnsetNextCursor() {
	o.NextCursor.Unset()
}

// GetPullNumber returns the PullNumber field value
func (o *CommitsInput) GetPullNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PullNumber
}

// GetPullNumberOk returns a tuple with the PullNumber field value
// and a boolean to check if the value has been set.
func (o *CommitsInput) GetPullNumberOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PullNumber, true
}

// SetPullNumber sets field value
func (o *CommitsInput) SetPullNumber(v int32) {
	o.PullNumber = v
}

func (o CommitsInput) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["pull_number"] = o.PullNumber
	}
	return json.Marshal(toSerialize)
}

type NullableCommitsInput struct {
	value *CommitsInput
	isSet bool
}

func (v NullableCommitsInput) Get() *CommitsInput {
	return v.value
}

func (v *NullableCommitsInput) Set(val *CommitsInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitsInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitsInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitsInput(val *CommitsInput) *NullableCommitsInput {
	return &NullableCommitsInput{value: val, isSet: true}
}

func (v NullableCommitsInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitsInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



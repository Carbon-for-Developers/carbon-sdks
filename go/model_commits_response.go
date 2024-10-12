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

// CommitsResponse struct for CommitsResponse
type CommitsResponse struct {
	Data []PRCommit `json:"data"`
	NextCursor NullableString `json:"next_cursor"`
}

// NewCommitsResponse instantiates a new CommitsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitsResponse(data []PRCommit, nextCursor NullableString) *CommitsResponse {
	this := CommitsResponse{}
	this.Data = data
	this.NextCursor = nextCursor
	return &this
}

// NewCommitsResponseWithDefaults instantiates a new CommitsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitsResponseWithDefaults() *CommitsResponse {
	this := CommitsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CommitsResponse) GetData() []PRCommit {
	if o == nil {
		var ret []PRCommit
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CommitsResponse) GetDataOk() ([]PRCommit, bool) {
	if o == nil {
    return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *CommitsResponse) SetData(v []PRCommit) {
	o.Data = v
}

// GetNextCursor returns the NextCursor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommitsResponse) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommitsResponse) GetNextCursorOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// SetNextCursor sets field value
func (o *CommitsResponse) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}

func (o CommitsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["next_cursor"] = o.NextCursor.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCommitsResponse struct {
	value *CommitsResponse
	isSet bool
}

func (v NullableCommitsResponse) Get() *CommitsResponse {
	return v.value
}

func (v *NullableCommitsResponse) Set(val *CommitsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitsResponse(val *CommitsResponse) *NullableCommitsResponse {
	return &NullableCommitsResponse{value: val, isSet: true}
}

func (v NullableCommitsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


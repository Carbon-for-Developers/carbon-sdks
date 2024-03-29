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

// ChunkPropertiesNullable struct for ChunkPropertiesNullable
type ChunkPropertiesNullable struct {
	SetPageAsBoundary bool `json:"set_page_as_boundary"`
	PrependFilenameToChunks bool `json:"prepend_filename_to_chunks"`
	MaxItemsPerChunk NullableInt32 `json:"max_items_per_chunk"`
}

// NewChunkPropertiesNullable instantiates a new ChunkPropertiesNullable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChunkPropertiesNullable(setPageAsBoundary bool, prependFilenameToChunks bool, maxItemsPerChunk NullableInt32) *ChunkPropertiesNullable {
	this := ChunkPropertiesNullable{}
	this.SetPageAsBoundary = setPageAsBoundary
	this.PrependFilenameToChunks = prependFilenameToChunks
	this.MaxItemsPerChunk = maxItemsPerChunk
	return &this
}

// NewChunkPropertiesNullableWithDefaults instantiates a new ChunkPropertiesNullable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChunkPropertiesNullableWithDefaults() *ChunkPropertiesNullable {
	this := ChunkPropertiesNullable{}
	var setPageAsBoundary bool = false
	this.SetPageAsBoundary = setPageAsBoundary
	var prependFilenameToChunks bool = false
	this.PrependFilenameToChunks = prependFilenameToChunks
	return &this
}

// GetSetPageAsBoundary returns the SetPageAsBoundary field value
func (o *ChunkPropertiesNullable) GetSetPageAsBoundary() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SetPageAsBoundary
}

// GetSetPageAsBoundaryOk returns a tuple with the SetPageAsBoundary field value
// and a boolean to check if the value has been set.
func (o *ChunkPropertiesNullable) GetSetPageAsBoundaryOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SetPageAsBoundary, true
}

// SetSetPageAsBoundary sets field value
func (o *ChunkPropertiesNullable) SetSetPageAsBoundary(v bool) {
	o.SetPageAsBoundary = v
}

// GetPrependFilenameToChunks returns the PrependFilenameToChunks field value
func (o *ChunkPropertiesNullable) GetPrependFilenameToChunks() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PrependFilenameToChunks
}

// GetPrependFilenameToChunksOk returns a tuple with the PrependFilenameToChunks field value
// and a boolean to check if the value has been set.
func (o *ChunkPropertiesNullable) GetPrependFilenameToChunksOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PrependFilenameToChunks, true
}

// SetPrependFilenameToChunks sets field value
func (o *ChunkPropertiesNullable) SetPrependFilenameToChunks(v bool) {
	o.PrependFilenameToChunks = v
}

// GetMaxItemsPerChunk returns the MaxItemsPerChunk field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ChunkPropertiesNullable) GetMaxItemsPerChunk() int32 {
	if o == nil || o.MaxItemsPerChunk.Get() == nil {
		var ret int32
		return ret
	}

	return *o.MaxItemsPerChunk.Get()
}

// GetMaxItemsPerChunkOk returns a tuple with the MaxItemsPerChunk field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChunkPropertiesNullable) GetMaxItemsPerChunkOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxItemsPerChunk.Get(), o.MaxItemsPerChunk.IsSet()
}

// SetMaxItemsPerChunk sets field value
func (o *ChunkPropertiesNullable) SetMaxItemsPerChunk(v int32) {
	o.MaxItemsPerChunk.Set(&v)
}

func (o ChunkPropertiesNullable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["set_page_as_boundary"] = o.SetPageAsBoundary
	}
	if true {
		toSerialize["prepend_filename_to_chunks"] = o.PrependFilenameToChunks
	}
	if true {
		toSerialize["max_items_per_chunk"] = o.MaxItemsPerChunk.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableChunkPropertiesNullable struct {
	value *ChunkPropertiesNullable
	isSet bool
}

func (v NullableChunkPropertiesNullable) Get() *ChunkPropertiesNullable {
	return v.value
}

func (v *NullableChunkPropertiesNullable) Set(val *ChunkPropertiesNullable) {
	v.value = val
	v.isSet = true
}

func (v NullableChunkPropertiesNullable) IsSet() bool {
	return v.isSet
}

func (v *NullableChunkPropertiesNullable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChunkPropertiesNullable(val *ChunkPropertiesNullable) *NullableChunkPropertiesNullable {
	return &NullableChunkPropertiesNullable{value: val, isSet: true}
}

func (v NullableChunkPropertiesNullable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChunkPropertiesNullable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



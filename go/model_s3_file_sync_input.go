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

// S3FileSyncInput struct for S3FileSyncInput
type S3FileSyncInput struct {
	Tags map[string]interface{} `json:"tags,omitempty"`
	Ids []S3GetFileInput `json:"ids"`
	ChunkSize NullableInt32 `json:"chunk_size,omitempty"`
	ChunkOverlap NullableInt32 `json:"chunk_overlap,omitempty"`
	SkipEmbeddingGeneration NullableBool `json:"skip_embedding_generation,omitempty"`
	EmbeddingModel *EmbeddingGenerators `json:"embedding_model,omitempty"`
	GenerateSparseVectors NullableBool `json:"generate_sparse_vectors,omitempty"`
	PrependFilenameToChunks NullableBool `json:"prepend_filename_to_chunks,omitempty"`
	// Number of objects per chunk. For csv, tsv, xlsx, and json files only.
	MaxItemsPerChunk NullableInt32 `json:"max_items_per_chunk,omitempty"`
	SetPageAsBoundary *bool `json:"set_page_as_boundary,omitempty"`
	DataSourceId NullableInt32 `json:"data_source_id,omitempty"`
}

// NewS3FileSyncInput instantiates a new S3FileSyncInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3FileSyncInput(ids []S3GetFileInput) *S3FileSyncInput {
	this := S3FileSyncInput{}
	this.Ids = ids
	var chunkSize int32 = 1500
	this.ChunkSize = *NewNullableInt32(&chunkSize)
	var chunkOverlap int32 = 20
	this.ChunkOverlap = *NewNullableInt32(&chunkOverlap)
	var skipEmbeddingGeneration bool = false
	this.SkipEmbeddingGeneration = *NewNullableBool(&skipEmbeddingGeneration)
	var generateSparseVectors bool = false
	this.GenerateSparseVectors = *NewNullableBool(&generateSparseVectors)
	var prependFilenameToChunks bool = false
	this.PrependFilenameToChunks = *NewNullableBool(&prependFilenameToChunks)
	var setPageAsBoundary bool = false
	this.SetPageAsBoundary = &setPageAsBoundary
	return &this
}

// NewS3FileSyncInputWithDefaults instantiates a new S3FileSyncInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3FileSyncInputWithDefaults() *S3FileSyncInput {
	this := S3FileSyncInput{}
	var chunkSize int32 = 1500
	this.ChunkSize = *NewNullableInt32(&chunkSize)
	var chunkOverlap int32 = 20
	this.ChunkOverlap = *NewNullableInt32(&chunkOverlap)
	var skipEmbeddingGeneration bool = false
	this.SkipEmbeddingGeneration = *NewNullableBool(&skipEmbeddingGeneration)
	var generateSparseVectors bool = false
	this.GenerateSparseVectors = *NewNullableBool(&generateSparseVectors)
	var prependFilenameToChunks bool = false
	this.PrependFilenameToChunks = *NewNullableBool(&prependFilenameToChunks)
	var setPageAsBoundary bool = false
	this.SetPageAsBoundary = &setPageAsBoundary
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *S3FileSyncInput) GetTags() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3FileSyncInput) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Tags) {
    return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *S3FileSyncInput) HasTags() bool {
	if o != nil && isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *S3FileSyncInput) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetIds returns the Ids field value
func (o *S3FileSyncInput) GetIds() []S3GetFileInput {
	if o == nil {
		var ret []S3GetFileInput
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *S3FileSyncInput) GetIdsOk() ([]S3GetFileInput, bool) {
	if o == nil {
    return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *S3FileSyncInput) SetIds(v []S3GetFileInput) {
	o.Ids = v
}

// GetChunkSize returns the ChunkSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *S3FileSyncInput) GetChunkSize() int32 {
	if o == nil || isNil(o.ChunkSize.Get()) {
		var ret int32
		return ret
	}
	return *o.ChunkSize.Get()
}

// GetChunkSizeOk returns a tuple with the ChunkSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3FileSyncInput) GetChunkSizeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.ChunkSize.Get(), o.ChunkSize.IsSet()
}

// HasChunkSize returns a boolean if a field has been set.
func (o *S3FileSyncInput) HasChunkSize() bool {
	if o != nil && o.ChunkSize.IsSet() {
		return true
	}

	return false
}

// SetChunkSize gets a reference to the given NullableInt32 and assigns it to the ChunkSize field.
func (o *S3FileSyncInput) SetChunkSize(v int32) {
	o.ChunkSize.Set(&v)
}
// SetChunkSizeNil sets the value for ChunkSize to be an explicit nil
func (o *S3FileSyncInput) SetChunkSizeNil() {
	o.ChunkSize.Set(nil)
}

// UnsetChunkSize ensures that no value is present for ChunkSize, not even an explicit nil
func (o *S3FileSyncInput) UnsetChunkSize() {
	o.ChunkSize.Unset()
}

// GetChunkOverlap returns the ChunkOverlap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *S3FileSyncInput) GetChunkOverlap() int32 {
	if o == nil || isNil(o.ChunkOverlap.Get()) {
		var ret int32
		return ret
	}
	return *o.ChunkOverlap.Get()
}

// GetChunkOverlapOk returns a tuple with the ChunkOverlap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3FileSyncInput) GetChunkOverlapOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.ChunkOverlap.Get(), o.ChunkOverlap.IsSet()
}

// HasChunkOverlap returns a boolean if a field has been set.
func (o *S3FileSyncInput) HasChunkOverlap() bool {
	if o != nil && o.ChunkOverlap.IsSet() {
		return true
	}

	return false
}

// SetChunkOverlap gets a reference to the given NullableInt32 and assigns it to the ChunkOverlap field.
func (o *S3FileSyncInput) SetChunkOverlap(v int32) {
	o.ChunkOverlap.Set(&v)
}
// SetChunkOverlapNil sets the value for ChunkOverlap to be an explicit nil
func (o *S3FileSyncInput) SetChunkOverlapNil() {
	o.ChunkOverlap.Set(nil)
}

// UnsetChunkOverlap ensures that no value is present for ChunkOverlap, not even an explicit nil
func (o *S3FileSyncInput) UnsetChunkOverlap() {
	o.ChunkOverlap.Unset()
}

// GetSkipEmbeddingGeneration returns the SkipEmbeddingGeneration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *S3FileSyncInput) GetSkipEmbeddingGeneration() bool {
	if o == nil || isNil(o.SkipEmbeddingGeneration.Get()) {
		var ret bool
		return ret
	}
	return *o.SkipEmbeddingGeneration.Get()
}

// GetSkipEmbeddingGenerationOk returns a tuple with the SkipEmbeddingGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3FileSyncInput) GetSkipEmbeddingGenerationOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.SkipEmbeddingGeneration.Get(), o.SkipEmbeddingGeneration.IsSet()
}

// HasSkipEmbeddingGeneration returns a boolean if a field has been set.
func (o *S3FileSyncInput) HasSkipEmbeddingGeneration() bool {
	if o != nil && o.SkipEmbeddingGeneration.IsSet() {
		return true
	}

	return false
}

// SetSkipEmbeddingGeneration gets a reference to the given NullableBool and assigns it to the SkipEmbeddingGeneration field.
func (o *S3FileSyncInput) SetSkipEmbeddingGeneration(v bool) {
	o.SkipEmbeddingGeneration.Set(&v)
}
// SetSkipEmbeddingGenerationNil sets the value for SkipEmbeddingGeneration to be an explicit nil
func (o *S3FileSyncInput) SetSkipEmbeddingGenerationNil() {
	o.SkipEmbeddingGeneration.Set(nil)
}

// UnsetSkipEmbeddingGeneration ensures that no value is present for SkipEmbeddingGeneration, not even an explicit nil
func (o *S3FileSyncInput) UnsetSkipEmbeddingGeneration() {
	o.SkipEmbeddingGeneration.Unset()
}

// GetEmbeddingModel returns the EmbeddingModel field value if set, zero value otherwise.
func (o *S3FileSyncInput) GetEmbeddingModel() EmbeddingGenerators {
	if o == nil || isNil(o.EmbeddingModel) {
		var ret EmbeddingGenerators
		return ret
	}
	return *o.EmbeddingModel
}

// GetEmbeddingModelOk returns a tuple with the EmbeddingModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3FileSyncInput) GetEmbeddingModelOk() (*EmbeddingGenerators, bool) {
	if o == nil || isNil(o.EmbeddingModel) {
    return nil, false
	}
	return o.EmbeddingModel, true
}

// HasEmbeddingModel returns a boolean if a field has been set.
func (o *S3FileSyncInput) HasEmbeddingModel() bool {
	if o != nil && !isNil(o.EmbeddingModel) {
		return true
	}

	return false
}

// SetEmbeddingModel gets a reference to the given EmbeddingGenerators and assigns it to the EmbeddingModel field.
func (o *S3FileSyncInput) SetEmbeddingModel(v EmbeddingGenerators) {
	o.EmbeddingModel = &v
}

// GetGenerateSparseVectors returns the GenerateSparseVectors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *S3FileSyncInput) GetGenerateSparseVectors() bool {
	if o == nil || isNil(o.GenerateSparseVectors.Get()) {
		var ret bool
		return ret
	}
	return *o.GenerateSparseVectors.Get()
}

// GetGenerateSparseVectorsOk returns a tuple with the GenerateSparseVectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3FileSyncInput) GetGenerateSparseVectorsOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.GenerateSparseVectors.Get(), o.GenerateSparseVectors.IsSet()
}

// HasGenerateSparseVectors returns a boolean if a field has been set.
func (o *S3FileSyncInput) HasGenerateSparseVectors() bool {
	if o != nil && o.GenerateSparseVectors.IsSet() {
		return true
	}

	return false
}

// SetGenerateSparseVectors gets a reference to the given NullableBool and assigns it to the GenerateSparseVectors field.
func (o *S3FileSyncInput) SetGenerateSparseVectors(v bool) {
	o.GenerateSparseVectors.Set(&v)
}
// SetGenerateSparseVectorsNil sets the value for GenerateSparseVectors to be an explicit nil
func (o *S3FileSyncInput) SetGenerateSparseVectorsNil() {
	o.GenerateSparseVectors.Set(nil)
}

// UnsetGenerateSparseVectors ensures that no value is present for GenerateSparseVectors, not even an explicit nil
func (o *S3FileSyncInput) UnsetGenerateSparseVectors() {
	o.GenerateSparseVectors.Unset()
}

// GetPrependFilenameToChunks returns the PrependFilenameToChunks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *S3FileSyncInput) GetPrependFilenameToChunks() bool {
	if o == nil || isNil(o.PrependFilenameToChunks.Get()) {
		var ret bool
		return ret
	}
	return *o.PrependFilenameToChunks.Get()
}

// GetPrependFilenameToChunksOk returns a tuple with the PrependFilenameToChunks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3FileSyncInput) GetPrependFilenameToChunksOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.PrependFilenameToChunks.Get(), o.PrependFilenameToChunks.IsSet()
}

// HasPrependFilenameToChunks returns a boolean if a field has been set.
func (o *S3FileSyncInput) HasPrependFilenameToChunks() bool {
	if o != nil && o.PrependFilenameToChunks.IsSet() {
		return true
	}

	return false
}

// SetPrependFilenameToChunks gets a reference to the given NullableBool and assigns it to the PrependFilenameToChunks field.
func (o *S3FileSyncInput) SetPrependFilenameToChunks(v bool) {
	o.PrependFilenameToChunks.Set(&v)
}
// SetPrependFilenameToChunksNil sets the value for PrependFilenameToChunks to be an explicit nil
func (o *S3FileSyncInput) SetPrependFilenameToChunksNil() {
	o.PrependFilenameToChunks.Set(nil)
}

// UnsetPrependFilenameToChunks ensures that no value is present for PrependFilenameToChunks, not even an explicit nil
func (o *S3FileSyncInput) UnsetPrependFilenameToChunks() {
	o.PrependFilenameToChunks.Unset()
}

// GetMaxItemsPerChunk returns the MaxItemsPerChunk field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *S3FileSyncInput) GetMaxItemsPerChunk() int32 {
	if o == nil || isNil(o.MaxItemsPerChunk.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxItemsPerChunk.Get()
}

// GetMaxItemsPerChunkOk returns a tuple with the MaxItemsPerChunk field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3FileSyncInput) GetMaxItemsPerChunkOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxItemsPerChunk.Get(), o.MaxItemsPerChunk.IsSet()
}

// HasMaxItemsPerChunk returns a boolean if a field has been set.
func (o *S3FileSyncInput) HasMaxItemsPerChunk() bool {
	if o != nil && o.MaxItemsPerChunk.IsSet() {
		return true
	}

	return false
}

// SetMaxItemsPerChunk gets a reference to the given NullableInt32 and assigns it to the MaxItemsPerChunk field.
func (o *S3FileSyncInput) SetMaxItemsPerChunk(v int32) {
	o.MaxItemsPerChunk.Set(&v)
}
// SetMaxItemsPerChunkNil sets the value for MaxItemsPerChunk to be an explicit nil
func (o *S3FileSyncInput) SetMaxItemsPerChunkNil() {
	o.MaxItemsPerChunk.Set(nil)
}

// UnsetMaxItemsPerChunk ensures that no value is present for MaxItemsPerChunk, not even an explicit nil
func (o *S3FileSyncInput) UnsetMaxItemsPerChunk() {
	o.MaxItemsPerChunk.Unset()
}

// GetSetPageAsBoundary returns the SetPageAsBoundary field value if set, zero value otherwise.
func (o *S3FileSyncInput) GetSetPageAsBoundary() bool {
	if o == nil || isNil(o.SetPageAsBoundary) {
		var ret bool
		return ret
	}
	return *o.SetPageAsBoundary
}

// GetSetPageAsBoundaryOk returns a tuple with the SetPageAsBoundary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3FileSyncInput) GetSetPageAsBoundaryOk() (*bool, bool) {
	if o == nil || isNil(o.SetPageAsBoundary) {
    return nil, false
	}
	return o.SetPageAsBoundary, true
}

// HasSetPageAsBoundary returns a boolean if a field has been set.
func (o *S3FileSyncInput) HasSetPageAsBoundary() bool {
	if o != nil && !isNil(o.SetPageAsBoundary) {
		return true
	}

	return false
}

// SetSetPageAsBoundary gets a reference to the given bool and assigns it to the SetPageAsBoundary field.
func (o *S3FileSyncInput) SetSetPageAsBoundary(v bool) {
	o.SetPageAsBoundary = &v
}

// GetDataSourceId returns the DataSourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *S3FileSyncInput) GetDataSourceId() int32 {
	if o == nil || isNil(o.DataSourceId.Get()) {
		var ret int32
		return ret
	}
	return *o.DataSourceId.Get()
}

// GetDataSourceIdOk returns a tuple with the DataSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3FileSyncInput) GetDataSourceIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.DataSourceId.Get(), o.DataSourceId.IsSet()
}

// HasDataSourceId returns a boolean if a field has been set.
func (o *S3FileSyncInput) HasDataSourceId() bool {
	if o != nil && o.DataSourceId.IsSet() {
		return true
	}

	return false
}

// SetDataSourceId gets a reference to the given NullableInt32 and assigns it to the DataSourceId field.
func (o *S3FileSyncInput) SetDataSourceId(v int32) {
	o.DataSourceId.Set(&v)
}
// SetDataSourceIdNil sets the value for DataSourceId to be an explicit nil
func (o *S3FileSyncInput) SetDataSourceIdNil() {
	o.DataSourceId.Set(nil)
}

// UnsetDataSourceId ensures that no value is present for DataSourceId, not even an explicit nil
func (o *S3FileSyncInput) UnsetDataSourceId() {
	o.DataSourceId.Unset()
}

func (o S3FileSyncInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["ids"] = o.Ids
	}
	if o.ChunkSize.IsSet() {
		toSerialize["chunk_size"] = o.ChunkSize.Get()
	}
	if o.ChunkOverlap.IsSet() {
		toSerialize["chunk_overlap"] = o.ChunkOverlap.Get()
	}
	if o.SkipEmbeddingGeneration.IsSet() {
		toSerialize["skip_embedding_generation"] = o.SkipEmbeddingGeneration.Get()
	}
	if !isNil(o.EmbeddingModel) {
		toSerialize["embedding_model"] = o.EmbeddingModel
	}
	if o.GenerateSparseVectors.IsSet() {
		toSerialize["generate_sparse_vectors"] = o.GenerateSparseVectors.Get()
	}
	if o.PrependFilenameToChunks.IsSet() {
		toSerialize["prepend_filename_to_chunks"] = o.PrependFilenameToChunks.Get()
	}
	if o.MaxItemsPerChunk.IsSet() {
		toSerialize["max_items_per_chunk"] = o.MaxItemsPerChunk.Get()
	}
	if !isNil(o.SetPageAsBoundary) {
		toSerialize["set_page_as_boundary"] = o.SetPageAsBoundary
	}
	if o.DataSourceId.IsSet() {
		toSerialize["data_source_id"] = o.DataSourceId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableS3FileSyncInput struct {
	value *S3FileSyncInput
	isSet bool
}

func (v NullableS3FileSyncInput) Get() *S3FileSyncInput {
	return v.value
}

func (v *NullableS3FileSyncInput) Set(val *S3FileSyncInput) {
	v.value = val
	v.isSet = true
}

func (v NullableS3FileSyncInput) IsSet() bool {
	return v.isSet
}

func (v *NullableS3FileSyncInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3FileSyncInput(val *S3FileSyncInput) *NullableS3FileSyncInput {
	return &NullableS3FileSyncInput{value: val, isSet: true}
}

func (v NullableS3FileSyncInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3FileSyncInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// FreshDeskConnectRequest struct for FreshDeskConnectRequest
type FreshDeskConnectRequest struct {
	Tags map[string]interface{} `json:"tags,omitempty"`
	Domain string `json:"domain"`
	ApiKey string `json:"api_key"`
	ChunkSize NullableInt32 `json:"chunk_size,omitempty"`
	ChunkOverlap NullableInt32 `json:"chunk_overlap,omitempty"`
	SkipEmbeddingGeneration NullableBool `json:"skip_embedding_generation,omitempty"`
	EmbeddingModel NullableEmbeddingGeneratorsNullable `json:"embedding_model,omitempty"`
	GenerateSparseVectors NullableBool `json:"generate_sparse_vectors,omitempty"`
	PrependFilenameToChunks NullableBool `json:"prepend_filename_to_chunks,omitempty"`
	SyncFilesOnConnection NullableBool `json:"sync_files_on_connection,omitempty"`
	RequestId NullableString `json:"request_id,omitempty"`
	// Enabling this flag will fetch all available content from the source to be listed via list items endpoint
	SyncSourceItems *bool `json:"sync_source_items,omitempty"`
}

// NewFreshDeskConnectRequest instantiates a new FreshDeskConnectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreshDeskConnectRequest(domain string, apiKey string) *FreshDeskConnectRequest {
	this := FreshDeskConnectRequest{}
	this.Domain = domain
	this.ApiKey = apiKey
	var chunkSize int32 = 1500
	this.ChunkSize = *NewNullableInt32(&chunkSize)
	var chunkOverlap int32 = 20
	this.ChunkOverlap = *NewNullableInt32(&chunkOverlap)
	var skipEmbeddingGeneration bool = false
	this.SkipEmbeddingGeneration = *NewNullableBool(&skipEmbeddingGeneration)
	var embeddingModel EmbeddingGeneratorsNullable = EMBEDDINGGENERATORSNULLABLE_OPENAI
	this.EmbeddingModel = *NewNullableEmbeddingGeneratorsNullable(&embeddingModel)
	var generateSparseVectors bool = false
	this.GenerateSparseVectors = *NewNullableBool(&generateSparseVectors)
	var prependFilenameToChunks bool = false
	this.PrependFilenameToChunks = *NewNullableBool(&prependFilenameToChunks)
	var syncFilesOnConnection bool = true
	this.SyncFilesOnConnection = *NewNullableBool(&syncFilesOnConnection)
	var syncSourceItems bool = true
	this.SyncSourceItems = &syncSourceItems
	return &this
}

// NewFreshDeskConnectRequestWithDefaults instantiates a new FreshDeskConnectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreshDeskConnectRequestWithDefaults() *FreshDeskConnectRequest {
	this := FreshDeskConnectRequest{}
	var chunkSize int32 = 1500
	this.ChunkSize = *NewNullableInt32(&chunkSize)
	var chunkOverlap int32 = 20
	this.ChunkOverlap = *NewNullableInt32(&chunkOverlap)
	var skipEmbeddingGeneration bool = false
	this.SkipEmbeddingGeneration = *NewNullableBool(&skipEmbeddingGeneration)
	var embeddingModel EmbeddingGeneratorsNullable = EMBEDDINGGENERATORSNULLABLE_OPENAI
	this.EmbeddingModel = *NewNullableEmbeddingGeneratorsNullable(&embeddingModel)
	var generateSparseVectors bool = false
	this.GenerateSparseVectors = *NewNullableBool(&generateSparseVectors)
	var prependFilenameToChunks bool = false
	this.PrependFilenameToChunks = *NewNullableBool(&prependFilenameToChunks)
	var syncFilesOnConnection bool = true
	this.SyncFilesOnConnection = *NewNullableBool(&syncFilesOnConnection)
	var syncSourceItems bool = true
	this.SyncSourceItems = &syncSourceItems
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FreshDeskConnectRequest) GetTags() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FreshDeskConnectRequest) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Tags) {
    return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FreshDeskConnectRequest) HasTags() bool {
	if o != nil && isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *FreshDeskConnectRequest) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetDomain returns the Domain field value
func (o *FreshDeskConnectRequest) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *FreshDeskConnectRequest) GetDomainOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *FreshDeskConnectRequest) SetDomain(v string) {
	o.Domain = v
}

// GetApiKey returns the ApiKey field value
func (o *FreshDeskConnectRequest) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *FreshDeskConnectRequest) GetApiKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *FreshDeskConnectRequest) SetApiKey(v string) {
	o.ApiKey = v
}

// GetChunkSize returns the ChunkSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FreshDeskConnectRequest) GetChunkSize() int32 {
	if o == nil || isNil(o.ChunkSize.Get()) {
		var ret int32
		return ret
	}
	return *o.ChunkSize.Get()
}

// GetChunkSizeOk returns a tuple with the ChunkSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FreshDeskConnectRequest) GetChunkSizeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.ChunkSize.Get(), o.ChunkSize.IsSet()
}

// HasChunkSize returns a boolean if a field has been set.
func (o *FreshDeskConnectRequest) HasChunkSize() bool {
	if o != nil && o.ChunkSize.IsSet() {
		return true
	}

	return false
}

// SetChunkSize gets a reference to the given NullableInt32 and assigns it to the ChunkSize field.
func (o *FreshDeskConnectRequest) SetChunkSize(v int32) {
	o.ChunkSize.Set(&v)
}
// SetChunkSizeNil sets the value for ChunkSize to be an explicit nil
func (o *FreshDeskConnectRequest) SetChunkSizeNil() {
	o.ChunkSize.Set(nil)
}

// UnsetChunkSize ensures that no value is present for ChunkSize, not even an explicit nil
func (o *FreshDeskConnectRequest) UnsetChunkSize() {
	o.ChunkSize.Unset()
}

// GetChunkOverlap returns the ChunkOverlap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FreshDeskConnectRequest) GetChunkOverlap() int32 {
	if o == nil || isNil(o.ChunkOverlap.Get()) {
		var ret int32
		return ret
	}
	return *o.ChunkOverlap.Get()
}

// GetChunkOverlapOk returns a tuple with the ChunkOverlap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FreshDeskConnectRequest) GetChunkOverlapOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.ChunkOverlap.Get(), o.ChunkOverlap.IsSet()
}

// HasChunkOverlap returns a boolean if a field has been set.
func (o *FreshDeskConnectRequest) HasChunkOverlap() bool {
	if o != nil && o.ChunkOverlap.IsSet() {
		return true
	}

	return false
}

// SetChunkOverlap gets a reference to the given NullableInt32 and assigns it to the ChunkOverlap field.
func (o *FreshDeskConnectRequest) SetChunkOverlap(v int32) {
	o.ChunkOverlap.Set(&v)
}
// SetChunkOverlapNil sets the value for ChunkOverlap to be an explicit nil
func (o *FreshDeskConnectRequest) SetChunkOverlapNil() {
	o.ChunkOverlap.Set(nil)
}

// UnsetChunkOverlap ensures that no value is present for ChunkOverlap, not even an explicit nil
func (o *FreshDeskConnectRequest) UnsetChunkOverlap() {
	o.ChunkOverlap.Unset()
}

// GetSkipEmbeddingGeneration returns the SkipEmbeddingGeneration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FreshDeskConnectRequest) GetSkipEmbeddingGeneration() bool {
	if o == nil || isNil(o.SkipEmbeddingGeneration.Get()) {
		var ret bool
		return ret
	}
	return *o.SkipEmbeddingGeneration.Get()
}

// GetSkipEmbeddingGenerationOk returns a tuple with the SkipEmbeddingGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FreshDeskConnectRequest) GetSkipEmbeddingGenerationOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.SkipEmbeddingGeneration.Get(), o.SkipEmbeddingGeneration.IsSet()
}

// HasSkipEmbeddingGeneration returns a boolean if a field has been set.
func (o *FreshDeskConnectRequest) HasSkipEmbeddingGeneration() bool {
	if o != nil && o.SkipEmbeddingGeneration.IsSet() {
		return true
	}

	return false
}

// SetSkipEmbeddingGeneration gets a reference to the given NullableBool and assigns it to the SkipEmbeddingGeneration field.
func (o *FreshDeskConnectRequest) SetSkipEmbeddingGeneration(v bool) {
	o.SkipEmbeddingGeneration.Set(&v)
}
// SetSkipEmbeddingGenerationNil sets the value for SkipEmbeddingGeneration to be an explicit nil
func (o *FreshDeskConnectRequest) SetSkipEmbeddingGenerationNil() {
	o.SkipEmbeddingGeneration.Set(nil)
}

// UnsetSkipEmbeddingGeneration ensures that no value is present for SkipEmbeddingGeneration, not even an explicit nil
func (o *FreshDeskConnectRequest) UnsetSkipEmbeddingGeneration() {
	o.SkipEmbeddingGeneration.Unset()
}

// GetEmbeddingModel returns the EmbeddingModel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FreshDeskConnectRequest) GetEmbeddingModel() EmbeddingGeneratorsNullable {
	if o == nil || isNil(o.EmbeddingModel.Get()) {
		var ret EmbeddingGeneratorsNullable
		return ret
	}
	return *o.EmbeddingModel.Get()
}

// GetEmbeddingModelOk returns a tuple with the EmbeddingModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FreshDeskConnectRequest) GetEmbeddingModelOk() (*EmbeddingGeneratorsNullable, bool) {
	if o == nil {
    return nil, false
	}
	return o.EmbeddingModel.Get(), o.EmbeddingModel.IsSet()
}

// HasEmbeddingModel returns a boolean if a field has been set.
func (o *FreshDeskConnectRequest) HasEmbeddingModel() bool {
	if o != nil && o.EmbeddingModel.IsSet() {
		return true
	}

	return false
}

// SetEmbeddingModel gets a reference to the given NullableEmbeddingGeneratorsNullable and assigns it to the EmbeddingModel field.
func (o *FreshDeskConnectRequest) SetEmbeddingModel(v EmbeddingGeneratorsNullable) {
	o.EmbeddingModel.Set(&v)
}
// SetEmbeddingModelNil sets the value for EmbeddingModel to be an explicit nil
func (o *FreshDeskConnectRequest) SetEmbeddingModelNil() {
	o.EmbeddingModel.Set(nil)
}

// UnsetEmbeddingModel ensures that no value is present for EmbeddingModel, not even an explicit nil
func (o *FreshDeskConnectRequest) UnsetEmbeddingModel() {
	o.EmbeddingModel.Unset()
}

// GetGenerateSparseVectors returns the GenerateSparseVectors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FreshDeskConnectRequest) GetGenerateSparseVectors() bool {
	if o == nil || isNil(o.GenerateSparseVectors.Get()) {
		var ret bool
		return ret
	}
	return *o.GenerateSparseVectors.Get()
}

// GetGenerateSparseVectorsOk returns a tuple with the GenerateSparseVectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FreshDeskConnectRequest) GetGenerateSparseVectorsOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.GenerateSparseVectors.Get(), o.GenerateSparseVectors.IsSet()
}

// HasGenerateSparseVectors returns a boolean if a field has been set.
func (o *FreshDeskConnectRequest) HasGenerateSparseVectors() bool {
	if o != nil && o.GenerateSparseVectors.IsSet() {
		return true
	}

	return false
}

// SetGenerateSparseVectors gets a reference to the given NullableBool and assigns it to the GenerateSparseVectors field.
func (o *FreshDeskConnectRequest) SetGenerateSparseVectors(v bool) {
	o.GenerateSparseVectors.Set(&v)
}
// SetGenerateSparseVectorsNil sets the value for GenerateSparseVectors to be an explicit nil
func (o *FreshDeskConnectRequest) SetGenerateSparseVectorsNil() {
	o.GenerateSparseVectors.Set(nil)
}

// UnsetGenerateSparseVectors ensures that no value is present for GenerateSparseVectors, not even an explicit nil
func (o *FreshDeskConnectRequest) UnsetGenerateSparseVectors() {
	o.GenerateSparseVectors.Unset()
}

// GetPrependFilenameToChunks returns the PrependFilenameToChunks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FreshDeskConnectRequest) GetPrependFilenameToChunks() bool {
	if o == nil || isNil(o.PrependFilenameToChunks.Get()) {
		var ret bool
		return ret
	}
	return *o.PrependFilenameToChunks.Get()
}

// GetPrependFilenameToChunksOk returns a tuple with the PrependFilenameToChunks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FreshDeskConnectRequest) GetPrependFilenameToChunksOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.PrependFilenameToChunks.Get(), o.PrependFilenameToChunks.IsSet()
}

// HasPrependFilenameToChunks returns a boolean if a field has been set.
func (o *FreshDeskConnectRequest) HasPrependFilenameToChunks() bool {
	if o != nil && o.PrependFilenameToChunks.IsSet() {
		return true
	}

	return false
}

// SetPrependFilenameToChunks gets a reference to the given NullableBool and assigns it to the PrependFilenameToChunks field.
func (o *FreshDeskConnectRequest) SetPrependFilenameToChunks(v bool) {
	o.PrependFilenameToChunks.Set(&v)
}
// SetPrependFilenameToChunksNil sets the value for PrependFilenameToChunks to be an explicit nil
func (o *FreshDeskConnectRequest) SetPrependFilenameToChunksNil() {
	o.PrependFilenameToChunks.Set(nil)
}

// UnsetPrependFilenameToChunks ensures that no value is present for PrependFilenameToChunks, not even an explicit nil
func (o *FreshDeskConnectRequest) UnsetPrependFilenameToChunks() {
	o.PrependFilenameToChunks.Unset()
}

// GetSyncFilesOnConnection returns the SyncFilesOnConnection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FreshDeskConnectRequest) GetSyncFilesOnConnection() bool {
	if o == nil || isNil(o.SyncFilesOnConnection.Get()) {
		var ret bool
		return ret
	}
	return *o.SyncFilesOnConnection.Get()
}

// GetSyncFilesOnConnectionOk returns a tuple with the SyncFilesOnConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FreshDeskConnectRequest) GetSyncFilesOnConnectionOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.SyncFilesOnConnection.Get(), o.SyncFilesOnConnection.IsSet()
}

// HasSyncFilesOnConnection returns a boolean if a field has been set.
func (o *FreshDeskConnectRequest) HasSyncFilesOnConnection() bool {
	if o != nil && o.SyncFilesOnConnection.IsSet() {
		return true
	}

	return false
}

// SetSyncFilesOnConnection gets a reference to the given NullableBool and assigns it to the SyncFilesOnConnection field.
func (o *FreshDeskConnectRequest) SetSyncFilesOnConnection(v bool) {
	o.SyncFilesOnConnection.Set(&v)
}
// SetSyncFilesOnConnectionNil sets the value for SyncFilesOnConnection to be an explicit nil
func (o *FreshDeskConnectRequest) SetSyncFilesOnConnectionNil() {
	o.SyncFilesOnConnection.Set(nil)
}

// UnsetSyncFilesOnConnection ensures that no value is present for SyncFilesOnConnection, not even an explicit nil
func (o *FreshDeskConnectRequest) UnsetSyncFilesOnConnection() {
	o.SyncFilesOnConnection.Unset()
}

// GetRequestId returns the RequestId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FreshDeskConnectRequest) GetRequestId() string {
	if o == nil || isNil(o.RequestId.Get()) {
		var ret string
		return ret
	}
	return *o.RequestId.Get()
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FreshDeskConnectRequest) GetRequestIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RequestId.Get(), o.RequestId.IsSet()
}

// HasRequestId returns a boolean if a field has been set.
func (o *FreshDeskConnectRequest) HasRequestId() bool {
	if o != nil && o.RequestId.IsSet() {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given NullableString and assigns it to the RequestId field.
func (o *FreshDeskConnectRequest) SetRequestId(v string) {
	o.RequestId.Set(&v)
}
// SetRequestIdNil sets the value for RequestId to be an explicit nil
func (o *FreshDeskConnectRequest) SetRequestIdNil() {
	o.RequestId.Set(nil)
}

// UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
func (o *FreshDeskConnectRequest) UnsetRequestId() {
	o.RequestId.Unset()
}

// GetSyncSourceItems returns the SyncSourceItems field value if set, zero value otherwise.
func (o *FreshDeskConnectRequest) GetSyncSourceItems() bool {
	if o == nil || isNil(o.SyncSourceItems) {
		var ret bool
		return ret
	}
	return *o.SyncSourceItems
}

// GetSyncSourceItemsOk returns a tuple with the SyncSourceItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreshDeskConnectRequest) GetSyncSourceItemsOk() (*bool, bool) {
	if o == nil || isNil(o.SyncSourceItems) {
    return nil, false
	}
	return o.SyncSourceItems, true
}

// HasSyncSourceItems returns a boolean if a field has been set.
func (o *FreshDeskConnectRequest) HasSyncSourceItems() bool {
	if o != nil && !isNil(o.SyncSourceItems) {
		return true
	}

	return false
}

// SetSyncSourceItems gets a reference to the given bool and assigns it to the SyncSourceItems field.
func (o *FreshDeskConnectRequest) SetSyncSourceItems(v bool) {
	o.SyncSourceItems = &v
}

func (o FreshDeskConnectRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["domain"] = o.Domain
	}
	if true {
		toSerialize["api_key"] = o.ApiKey
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
	if o.EmbeddingModel.IsSet() {
		toSerialize["embedding_model"] = o.EmbeddingModel.Get()
	}
	if o.GenerateSparseVectors.IsSet() {
		toSerialize["generate_sparse_vectors"] = o.GenerateSparseVectors.Get()
	}
	if o.PrependFilenameToChunks.IsSet() {
		toSerialize["prepend_filename_to_chunks"] = o.PrependFilenameToChunks.Get()
	}
	if o.SyncFilesOnConnection.IsSet() {
		toSerialize["sync_files_on_connection"] = o.SyncFilesOnConnection.Get()
	}
	if o.RequestId.IsSet() {
		toSerialize["request_id"] = o.RequestId.Get()
	}
	if !isNil(o.SyncSourceItems) {
		toSerialize["sync_source_items"] = o.SyncSourceItems
	}
	return json.Marshal(toSerialize)
}

type NullableFreshDeskConnectRequest struct {
	value *FreshDeskConnectRequest
	isSet bool
}

func (v NullableFreshDeskConnectRequest) Get() *FreshDeskConnectRequest {
	return v.value
}

func (v *NullableFreshDeskConnectRequest) Set(val *FreshDeskConnectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFreshDeskConnectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFreshDeskConnectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreshDeskConnectRequest(val *FreshDeskConnectRequest) *NullableFreshDeskConnectRequest {
	return &NullableFreshDeskConnectRequest{value: val, isSet: true}
}

func (v NullableFreshDeskConnectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreshDeskConnectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



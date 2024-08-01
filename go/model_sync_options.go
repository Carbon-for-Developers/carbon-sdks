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

// SyncOptions struct for SyncOptions
type SyncOptions struct {
	Tags interface{} `json:"tags,omitempty"`
	ChunkSize NullableInt32 `json:"chunk_size,omitempty"`
	ChunkOverlap NullableInt32 `json:"chunk_overlap,omitempty"`
	SkipEmbeddingGeneration NullableBool `json:"skip_embedding_generation,omitempty"`
	EmbeddingModel NullableEmbeddingGeneratorsNullable `json:"embedding_model,omitempty"`
	GenerateSparseVectors NullableBool `json:"generate_sparse_vectors,omitempty"`
	PrependFilenameToChunks NullableBool `json:"prepend_filename_to_chunks,omitempty"`
	// Number of objects per chunk. For csv, tsv, xlsx, and json files only.
	MaxItemsPerChunk NullableInt32 `json:"max_items_per_chunk,omitempty"`
	// Used to specify whether Carbon should attempt to sync all your files automatically when authorization         is complete. This is only supported for a subset of connectors and will be ignored for the rest. Supported         connectors: Intercom, Zendesk, Gitbook, Confluence, Salesforce, Freshdesk
	SyncFilesOnConnection NullableBool `json:"sync_files_on_connection,omitempty"`
	SetPageAsBoundary *bool `json:"set_page_as_boundary,omitempty"`
	RequestId *string `json:"request_id,omitempty"`
	EnableFilePicker *bool `json:"enable_file_picker,omitempty"`
	// Enabling this flag will fetch all available content from the source to be listed via list items endpoint
	SyncSourceItems *bool `json:"sync_source_items,omitempty"`
	// Only sync files if they have not already been synced or if the embedding properties have changed.         This flag is currently supported by ONEDRIVE, GOOGLE_DRIVE, BOX, DROPBOX, INTERCOM, GMAIL, OUTLOOK. It will be ignored for other data sources.
	IncrementalSync *bool `json:"incremental_sync,omitempty"`
	FileSyncConfig NullableFileSyncConfigNullable `json:"file_sync_config,omitempty"`
}

// NewSyncOptions instantiates a new SyncOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncOptions() *SyncOptions {
	this := SyncOptions{}
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
	var setPageAsBoundary bool = false
	this.SetPageAsBoundary = &setPageAsBoundary
	var requestId string = "bdd2d0b8-c211-49bd-b70a-4889ae5fab99"
	this.RequestId = &requestId
	var enableFilePicker bool = true
	this.EnableFilePicker = &enableFilePicker
	var syncSourceItems bool = true
	this.SyncSourceItems = &syncSourceItems
	var incrementalSync bool = false
	this.IncrementalSync = &incrementalSync
	return &this
}

// NewSyncOptionsWithDefaults instantiates a new SyncOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncOptionsWithDefaults() *SyncOptions {
	this := SyncOptions{}
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
	var setPageAsBoundary bool = false
	this.SetPageAsBoundary = &setPageAsBoundary
	var requestId string = "bdd2d0b8-c211-49bd-b70a-4889ae5fab99"
	this.RequestId = &requestId
	var enableFilePicker bool = true
	this.EnableFilePicker = &enableFilePicker
	var syncSourceItems bool = true
	this.SyncSourceItems = &syncSourceItems
	var incrementalSync bool = false
	this.IncrementalSync = &incrementalSync
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyncOptions) GetTags() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyncOptions) GetTagsOk() (*interface{}, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SyncOptions) HasTags() bool {
	if o != nil && isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given interface{} and assigns it to the Tags field.
func (o *SyncOptions) SetTags(v interface{}) {
	o.Tags = v
}

// GetChunkSize returns the ChunkSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyncOptions) GetChunkSize() int32 {
	if o == nil || isNil(o.ChunkSize.Get()) {
		var ret int32
		return ret
	}
	return *o.ChunkSize.Get()
}

// GetChunkSizeOk returns a tuple with the ChunkSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyncOptions) GetChunkSizeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.ChunkSize.Get(), o.ChunkSize.IsSet()
}

// HasChunkSize returns a boolean if a field has been set.
func (o *SyncOptions) HasChunkSize() bool {
	if o != nil && o.ChunkSize.IsSet() {
		return true
	}

	return false
}

// SetChunkSize gets a reference to the given NullableInt32 and assigns it to the ChunkSize field.
func (o *SyncOptions) SetChunkSize(v int32) {
	o.ChunkSize.Set(&v)
}
// SetChunkSizeNil sets the value for ChunkSize to be an explicit nil
func (o *SyncOptions) SetChunkSizeNil() {
	o.ChunkSize.Set(nil)
}

// UnsetChunkSize ensures that no value is present for ChunkSize, not even an explicit nil
func (o *SyncOptions) UnsetChunkSize() {
	o.ChunkSize.Unset()
}

// GetChunkOverlap returns the ChunkOverlap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyncOptions) GetChunkOverlap() int32 {
	if o == nil || isNil(o.ChunkOverlap.Get()) {
		var ret int32
		return ret
	}
	return *o.ChunkOverlap.Get()
}

// GetChunkOverlapOk returns a tuple with the ChunkOverlap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyncOptions) GetChunkOverlapOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.ChunkOverlap.Get(), o.ChunkOverlap.IsSet()
}

// HasChunkOverlap returns a boolean if a field has been set.
func (o *SyncOptions) HasChunkOverlap() bool {
	if o != nil && o.ChunkOverlap.IsSet() {
		return true
	}

	return false
}

// SetChunkOverlap gets a reference to the given NullableInt32 and assigns it to the ChunkOverlap field.
func (o *SyncOptions) SetChunkOverlap(v int32) {
	o.ChunkOverlap.Set(&v)
}
// SetChunkOverlapNil sets the value for ChunkOverlap to be an explicit nil
func (o *SyncOptions) SetChunkOverlapNil() {
	o.ChunkOverlap.Set(nil)
}

// UnsetChunkOverlap ensures that no value is present for ChunkOverlap, not even an explicit nil
func (o *SyncOptions) UnsetChunkOverlap() {
	o.ChunkOverlap.Unset()
}

// GetSkipEmbeddingGeneration returns the SkipEmbeddingGeneration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyncOptions) GetSkipEmbeddingGeneration() bool {
	if o == nil || isNil(o.SkipEmbeddingGeneration.Get()) {
		var ret bool
		return ret
	}
	return *o.SkipEmbeddingGeneration.Get()
}

// GetSkipEmbeddingGenerationOk returns a tuple with the SkipEmbeddingGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyncOptions) GetSkipEmbeddingGenerationOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.SkipEmbeddingGeneration.Get(), o.SkipEmbeddingGeneration.IsSet()
}

// HasSkipEmbeddingGeneration returns a boolean if a field has been set.
func (o *SyncOptions) HasSkipEmbeddingGeneration() bool {
	if o != nil && o.SkipEmbeddingGeneration.IsSet() {
		return true
	}

	return false
}

// SetSkipEmbeddingGeneration gets a reference to the given NullableBool and assigns it to the SkipEmbeddingGeneration field.
func (o *SyncOptions) SetSkipEmbeddingGeneration(v bool) {
	o.SkipEmbeddingGeneration.Set(&v)
}
// SetSkipEmbeddingGenerationNil sets the value for SkipEmbeddingGeneration to be an explicit nil
func (o *SyncOptions) SetSkipEmbeddingGenerationNil() {
	o.SkipEmbeddingGeneration.Set(nil)
}

// UnsetSkipEmbeddingGeneration ensures that no value is present for SkipEmbeddingGeneration, not even an explicit nil
func (o *SyncOptions) UnsetSkipEmbeddingGeneration() {
	o.SkipEmbeddingGeneration.Unset()
}

// GetEmbeddingModel returns the EmbeddingModel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyncOptions) GetEmbeddingModel() EmbeddingGeneratorsNullable {
	if o == nil || isNil(o.EmbeddingModel.Get()) {
		var ret EmbeddingGeneratorsNullable
		return ret
	}
	return *o.EmbeddingModel.Get()
}

// GetEmbeddingModelOk returns a tuple with the EmbeddingModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyncOptions) GetEmbeddingModelOk() (*EmbeddingGeneratorsNullable, bool) {
	if o == nil {
    return nil, false
	}
	return o.EmbeddingModel.Get(), o.EmbeddingModel.IsSet()
}

// HasEmbeddingModel returns a boolean if a field has been set.
func (o *SyncOptions) HasEmbeddingModel() bool {
	if o != nil && o.EmbeddingModel.IsSet() {
		return true
	}

	return false
}

// SetEmbeddingModel gets a reference to the given NullableEmbeddingGeneratorsNullable and assigns it to the EmbeddingModel field.
func (o *SyncOptions) SetEmbeddingModel(v EmbeddingGeneratorsNullable) {
	o.EmbeddingModel.Set(&v)
}
// SetEmbeddingModelNil sets the value for EmbeddingModel to be an explicit nil
func (o *SyncOptions) SetEmbeddingModelNil() {
	o.EmbeddingModel.Set(nil)
}

// UnsetEmbeddingModel ensures that no value is present for EmbeddingModel, not even an explicit nil
func (o *SyncOptions) UnsetEmbeddingModel() {
	o.EmbeddingModel.Unset()
}

// GetGenerateSparseVectors returns the GenerateSparseVectors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyncOptions) GetGenerateSparseVectors() bool {
	if o == nil || isNil(o.GenerateSparseVectors.Get()) {
		var ret bool
		return ret
	}
	return *o.GenerateSparseVectors.Get()
}

// GetGenerateSparseVectorsOk returns a tuple with the GenerateSparseVectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyncOptions) GetGenerateSparseVectorsOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.GenerateSparseVectors.Get(), o.GenerateSparseVectors.IsSet()
}

// HasGenerateSparseVectors returns a boolean if a field has been set.
func (o *SyncOptions) HasGenerateSparseVectors() bool {
	if o != nil && o.GenerateSparseVectors.IsSet() {
		return true
	}

	return false
}

// SetGenerateSparseVectors gets a reference to the given NullableBool and assigns it to the GenerateSparseVectors field.
func (o *SyncOptions) SetGenerateSparseVectors(v bool) {
	o.GenerateSparseVectors.Set(&v)
}
// SetGenerateSparseVectorsNil sets the value for GenerateSparseVectors to be an explicit nil
func (o *SyncOptions) SetGenerateSparseVectorsNil() {
	o.GenerateSparseVectors.Set(nil)
}

// UnsetGenerateSparseVectors ensures that no value is present for GenerateSparseVectors, not even an explicit nil
func (o *SyncOptions) UnsetGenerateSparseVectors() {
	o.GenerateSparseVectors.Unset()
}

// GetPrependFilenameToChunks returns the PrependFilenameToChunks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyncOptions) GetPrependFilenameToChunks() bool {
	if o == nil || isNil(o.PrependFilenameToChunks.Get()) {
		var ret bool
		return ret
	}
	return *o.PrependFilenameToChunks.Get()
}

// GetPrependFilenameToChunksOk returns a tuple with the PrependFilenameToChunks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyncOptions) GetPrependFilenameToChunksOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.PrependFilenameToChunks.Get(), o.PrependFilenameToChunks.IsSet()
}

// HasPrependFilenameToChunks returns a boolean if a field has been set.
func (o *SyncOptions) HasPrependFilenameToChunks() bool {
	if o != nil && o.PrependFilenameToChunks.IsSet() {
		return true
	}

	return false
}

// SetPrependFilenameToChunks gets a reference to the given NullableBool and assigns it to the PrependFilenameToChunks field.
func (o *SyncOptions) SetPrependFilenameToChunks(v bool) {
	o.PrependFilenameToChunks.Set(&v)
}
// SetPrependFilenameToChunksNil sets the value for PrependFilenameToChunks to be an explicit nil
func (o *SyncOptions) SetPrependFilenameToChunksNil() {
	o.PrependFilenameToChunks.Set(nil)
}

// UnsetPrependFilenameToChunks ensures that no value is present for PrependFilenameToChunks, not even an explicit nil
func (o *SyncOptions) UnsetPrependFilenameToChunks() {
	o.PrependFilenameToChunks.Unset()
}

// GetMaxItemsPerChunk returns the MaxItemsPerChunk field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyncOptions) GetMaxItemsPerChunk() int32 {
	if o == nil || isNil(o.MaxItemsPerChunk.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxItemsPerChunk.Get()
}

// GetMaxItemsPerChunkOk returns a tuple with the MaxItemsPerChunk field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyncOptions) GetMaxItemsPerChunkOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxItemsPerChunk.Get(), o.MaxItemsPerChunk.IsSet()
}

// HasMaxItemsPerChunk returns a boolean if a field has been set.
func (o *SyncOptions) HasMaxItemsPerChunk() bool {
	if o != nil && o.MaxItemsPerChunk.IsSet() {
		return true
	}

	return false
}

// SetMaxItemsPerChunk gets a reference to the given NullableInt32 and assigns it to the MaxItemsPerChunk field.
func (o *SyncOptions) SetMaxItemsPerChunk(v int32) {
	o.MaxItemsPerChunk.Set(&v)
}
// SetMaxItemsPerChunkNil sets the value for MaxItemsPerChunk to be an explicit nil
func (o *SyncOptions) SetMaxItemsPerChunkNil() {
	o.MaxItemsPerChunk.Set(nil)
}

// UnsetMaxItemsPerChunk ensures that no value is present for MaxItemsPerChunk, not even an explicit nil
func (o *SyncOptions) UnsetMaxItemsPerChunk() {
	o.MaxItemsPerChunk.Unset()
}

// GetSyncFilesOnConnection returns the SyncFilesOnConnection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyncOptions) GetSyncFilesOnConnection() bool {
	if o == nil || isNil(o.SyncFilesOnConnection.Get()) {
		var ret bool
		return ret
	}
	return *o.SyncFilesOnConnection.Get()
}

// GetSyncFilesOnConnectionOk returns a tuple with the SyncFilesOnConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyncOptions) GetSyncFilesOnConnectionOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.SyncFilesOnConnection.Get(), o.SyncFilesOnConnection.IsSet()
}

// HasSyncFilesOnConnection returns a boolean if a field has been set.
func (o *SyncOptions) HasSyncFilesOnConnection() bool {
	if o != nil && o.SyncFilesOnConnection.IsSet() {
		return true
	}

	return false
}

// SetSyncFilesOnConnection gets a reference to the given NullableBool and assigns it to the SyncFilesOnConnection field.
func (o *SyncOptions) SetSyncFilesOnConnection(v bool) {
	o.SyncFilesOnConnection.Set(&v)
}
// SetSyncFilesOnConnectionNil sets the value for SyncFilesOnConnection to be an explicit nil
func (o *SyncOptions) SetSyncFilesOnConnectionNil() {
	o.SyncFilesOnConnection.Set(nil)
}

// UnsetSyncFilesOnConnection ensures that no value is present for SyncFilesOnConnection, not even an explicit nil
func (o *SyncOptions) UnsetSyncFilesOnConnection() {
	o.SyncFilesOnConnection.Unset()
}

// GetSetPageAsBoundary returns the SetPageAsBoundary field value if set, zero value otherwise.
func (o *SyncOptions) GetSetPageAsBoundary() bool {
	if o == nil || isNil(o.SetPageAsBoundary) {
		var ret bool
		return ret
	}
	return *o.SetPageAsBoundary
}

// GetSetPageAsBoundaryOk returns a tuple with the SetPageAsBoundary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncOptions) GetSetPageAsBoundaryOk() (*bool, bool) {
	if o == nil || isNil(o.SetPageAsBoundary) {
    return nil, false
	}
	return o.SetPageAsBoundary, true
}

// HasSetPageAsBoundary returns a boolean if a field has been set.
func (o *SyncOptions) HasSetPageAsBoundary() bool {
	if o != nil && !isNil(o.SetPageAsBoundary) {
		return true
	}

	return false
}

// SetSetPageAsBoundary gets a reference to the given bool and assigns it to the SetPageAsBoundary field.
func (o *SyncOptions) SetSetPageAsBoundary(v bool) {
	o.SetPageAsBoundary = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *SyncOptions) GetRequestId() string {
	if o == nil || isNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncOptions) GetRequestIdOk() (*string, bool) {
	if o == nil || isNil(o.RequestId) {
    return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *SyncOptions) HasRequestId() bool {
	if o != nil && !isNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *SyncOptions) SetRequestId(v string) {
	o.RequestId = &v
}

// GetEnableFilePicker returns the EnableFilePicker field value if set, zero value otherwise.
func (o *SyncOptions) GetEnableFilePicker() bool {
	if o == nil || isNil(o.EnableFilePicker) {
		var ret bool
		return ret
	}
	return *o.EnableFilePicker
}

// GetEnableFilePickerOk returns a tuple with the EnableFilePicker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncOptions) GetEnableFilePickerOk() (*bool, bool) {
	if o == nil || isNil(o.EnableFilePicker) {
    return nil, false
	}
	return o.EnableFilePicker, true
}

// HasEnableFilePicker returns a boolean if a field has been set.
func (o *SyncOptions) HasEnableFilePicker() bool {
	if o != nil && !isNil(o.EnableFilePicker) {
		return true
	}

	return false
}

// SetEnableFilePicker gets a reference to the given bool and assigns it to the EnableFilePicker field.
func (o *SyncOptions) SetEnableFilePicker(v bool) {
	o.EnableFilePicker = &v
}

// GetSyncSourceItems returns the SyncSourceItems field value if set, zero value otherwise.
func (o *SyncOptions) GetSyncSourceItems() bool {
	if o == nil || isNil(o.SyncSourceItems) {
		var ret bool
		return ret
	}
	return *o.SyncSourceItems
}

// GetSyncSourceItemsOk returns a tuple with the SyncSourceItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncOptions) GetSyncSourceItemsOk() (*bool, bool) {
	if o == nil || isNil(o.SyncSourceItems) {
    return nil, false
	}
	return o.SyncSourceItems, true
}

// HasSyncSourceItems returns a boolean if a field has been set.
func (o *SyncOptions) HasSyncSourceItems() bool {
	if o != nil && !isNil(o.SyncSourceItems) {
		return true
	}

	return false
}

// SetSyncSourceItems gets a reference to the given bool and assigns it to the SyncSourceItems field.
func (o *SyncOptions) SetSyncSourceItems(v bool) {
	o.SyncSourceItems = &v
}

// GetIncrementalSync returns the IncrementalSync field value if set, zero value otherwise.
func (o *SyncOptions) GetIncrementalSync() bool {
	if o == nil || isNil(o.IncrementalSync) {
		var ret bool
		return ret
	}
	return *o.IncrementalSync
}

// GetIncrementalSyncOk returns a tuple with the IncrementalSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncOptions) GetIncrementalSyncOk() (*bool, bool) {
	if o == nil || isNil(o.IncrementalSync) {
    return nil, false
	}
	return o.IncrementalSync, true
}

// HasIncrementalSync returns a boolean if a field has been set.
func (o *SyncOptions) HasIncrementalSync() bool {
	if o != nil && !isNil(o.IncrementalSync) {
		return true
	}

	return false
}

// SetIncrementalSync gets a reference to the given bool and assigns it to the IncrementalSync field.
func (o *SyncOptions) SetIncrementalSync(v bool) {
	o.IncrementalSync = &v
}

// GetFileSyncConfig returns the FileSyncConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyncOptions) GetFileSyncConfig() FileSyncConfigNullable {
	if o == nil || isNil(o.FileSyncConfig.Get()) {
		var ret FileSyncConfigNullable
		return ret
	}
	return *o.FileSyncConfig.Get()
}

// GetFileSyncConfigOk returns a tuple with the FileSyncConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyncOptions) GetFileSyncConfigOk() (*FileSyncConfigNullable, bool) {
	if o == nil {
    return nil, false
	}
	return o.FileSyncConfig.Get(), o.FileSyncConfig.IsSet()
}

// HasFileSyncConfig returns a boolean if a field has been set.
func (o *SyncOptions) HasFileSyncConfig() bool {
	if o != nil && o.FileSyncConfig.IsSet() {
		return true
	}

	return false
}

// SetFileSyncConfig gets a reference to the given NullableFileSyncConfigNullable and assigns it to the FileSyncConfig field.
func (o *SyncOptions) SetFileSyncConfig(v FileSyncConfigNullable) {
	o.FileSyncConfig.Set(&v)
}
// SetFileSyncConfigNil sets the value for FileSyncConfig to be an explicit nil
func (o *SyncOptions) SetFileSyncConfigNil() {
	o.FileSyncConfig.Set(nil)
}

// UnsetFileSyncConfig ensures that no value is present for FileSyncConfig, not even an explicit nil
func (o *SyncOptions) UnsetFileSyncConfig() {
	o.FileSyncConfig.Unset()
}

func (o SyncOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
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
	if o.MaxItemsPerChunk.IsSet() {
		toSerialize["max_items_per_chunk"] = o.MaxItemsPerChunk.Get()
	}
	if o.SyncFilesOnConnection.IsSet() {
		toSerialize["sync_files_on_connection"] = o.SyncFilesOnConnection.Get()
	}
	if !isNil(o.SetPageAsBoundary) {
		toSerialize["set_page_as_boundary"] = o.SetPageAsBoundary
	}
	if !isNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	if !isNil(o.EnableFilePicker) {
		toSerialize["enable_file_picker"] = o.EnableFilePicker
	}
	if !isNil(o.SyncSourceItems) {
		toSerialize["sync_source_items"] = o.SyncSourceItems
	}
	if !isNil(o.IncrementalSync) {
		toSerialize["incremental_sync"] = o.IncrementalSync
	}
	if o.FileSyncConfig.IsSet() {
		toSerialize["file_sync_config"] = o.FileSyncConfig.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSyncOptions struct {
	value *SyncOptions
	isSet bool
}

func (v NullableSyncOptions) Get() *SyncOptions {
	return v.value
}

func (v *NullableSyncOptions) Set(val *SyncOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncOptions(val *SyncOptions) *NullableSyncOptions {
	return &NullableSyncOptions{value: val, isSet: true}
}

func (v NullableSyncOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// GitbookConnectRequest struct for GitbookConnectRequest
type GitbookConnectRequest struct {
	Tags map[string]interface{} `json:"tags,omitempty"`
	Organization string `json:"organization"`
	AccessToken string `json:"access_token"`
	ChunkSize NullableInt32 `json:"chunk_size,omitempty"`
	ChunkOverlap NullableInt32 `json:"chunk_overlap,omitempty"`
	SkipEmbeddingGeneration NullableBool `json:"skip_embedding_generation,omitempty"`
	EmbeddingModel *EmbeddingGenerators `json:"embedding_model,omitempty"`
	GenerateSparseVectors NullableBool `json:"generate_sparse_vectors,omitempty"`
	PrependFilenameToChunks NullableBool `json:"prepend_filename_to_chunks,omitempty"`
	SyncFilesOnConnection NullableBool `json:"sync_files_on_connection,omitempty"`
	RequestId NullableString `json:"request_id,omitempty"`
	// Enabling this flag will fetch all available content from the source to be listed via list items endpoint
	SyncSourceItems *bool `json:"sync_source_items,omitempty"`
	FileSyncConfig NullableFileSyncConfigNullable `json:"file_sync_config,omitempty"`
	// Tags to be associated with the data source. If the data source already has tags set, then an upsert will be performed.
	DataSourceTags map[string]interface{} `json:"data_source_tags,omitempty"`
}

// NewGitbookConnectRequest instantiates a new GitbookConnectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitbookConnectRequest(organization string, accessToken string) *GitbookConnectRequest {
	this := GitbookConnectRequest{}
	this.Organization = organization
	this.AccessToken = accessToken
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
	var syncFilesOnConnection bool = true
	this.SyncFilesOnConnection = *NewNullableBool(&syncFilesOnConnection)
	var syncSourceItems bool = true
	this.SyncSourceItems = &syncSourceItems
	return &this
}

// NewGitbookConnectRequestWithDefaults instantiates a new GitbookConnectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitbookConnectRequestWithDefaults() *GitbookConnectRequest {
	this := GitbookConnectRequest{}
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
	var syncFilesOnConnection bool = true
	this.SyncFilesOnConnection = *NewNullableBool(&syncFilesOnConnection)
	var syncSourceItems bool = true
	this.SyncSourceItems = &syncSourceItems
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitbookConnectRequest) GetTags() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitbookConnectRequest) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Tags) {
    return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GitbookConnectRequest) HasTags() bool {
	if o != nil && isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *GitbookConnectRequest) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetOrganization returns the Organization field value
func (o *GitbookConnectRequest) GetOrganization() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *GitbookConnectRequest) GetOrganizationOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *GitbookConnectRequest) SetOrganization(v string) {
	o.Organization = v
}

// GetAccessToken returns the AccessToken field value
func (o *GitbookConnectRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *GitbookConnectRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *GitbookConnectRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetChunkSize returns the ChunkSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitbookConnectRequest) GetChunkSize() int32 {
	if o == nil || isNil(o.ChunkSize.Get()) {
		var ret int32
		return ret
	}
	return *o.ChunkSize.Get()
}

// GetChunkSizeOk returns a tuple with the ChunkSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitbookConnectRequest) GetChunkSizeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.ChunkSize.Get(), o.ChunkSize.IsSet()
}

// HasChunkSize returns a boolean if a field has been set.
func (o *GitbookConnectRequest) HasChunkSize() bool {
	if o != nil && o.ChunkSize.IsSet() {
		return true
	}

	return false
}

// SetChunkSize gets a reference to the given NullableInt32 and assigns it to the ChunkSize field.
func (o *GitbookConnectRequest) SetChunkSize(v int32) {
	o.ChunkSize.Set(&v)
}
// SetChunkSizeNil sets the value for ChunkSize to be an explicit nil
func (o *GitbookConnectRequest) SetChunkSizeNil() {
	o.ChunkSize.Set(nil)
}

// UnsetChunkSize ensures that no value is present for ChunkSize, not even an explicit nil
func (o *GitbookConnectRequest) UnsetChunkSize() {
	o.ChunkSize.Unset()
}

// GetChunkOverlap returns the ChunkOverlap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitbookConnectRequest) GetChunkOverlap() int32 {
	if o == nil || isNil(o.ChunkOverlap.Get()) {
		var ret int32
		return ret
	}
	return *o.ChunkOverlap.Get()
}

// GetChunkOverlapOk returns a tuple with the ChunkOverlap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitbookConnectRequest) GetChunkOverlapOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.ChunkOverlap.Get(), o.ChunkOverlap.IsSet()
}

// HasChunkOverlap returns a boolean if a field has been set.
func (o *GitbookConnectRequest) HasChunkOverlap() bool {
	if o != nil && o.ChunkOverlap.IsSet() {
		return true
	}

	return false
}

// SetChunkOverlap gets a reference to the given NullableInt32 and assigns it to the ChunkOverlap field.
func (o *GitbookConnectRequest) SetChunkOverlap(v int32) {
	o.ChunkOverlap.Set(&v)
}
// SetChunkOverlapNil sets the value for ChunkOverlap to be an explicit nil
func (o *GitbookConnectRequest) SetChunkOverlapNil() {
	o.ChunkOverlap.Set(nil)
}

// UnsetChunkOverlap ensures that no value is present for ChunkOverlap, not even an explicit nil
func (o *GitbookConnectRequest) UnsetChunkOverlap() {
	o.ChunkOverlap.Unset()
}

// GetSkipEmbeddingGeneration returns the SkipEmbeddingGeneration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitbookConnectRequest) GetSkipEmbeddingGeneration() bool {
	if o == nil || isNil(o.SkipEmbeddingGeneration.Get()) {
		var ret bool
		return ret
	}
	return *o.SkipEmbeddingGeneration.Get()
}

// GetSkipEmbeddingGenerationOk returns a tuple with the SkipEmbeddingGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitbookConnectRequest) GetSkipEmbeddingGenerationOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.SkipEmbeddingGeneration.Get(), o.SkipEmbeddingGeneration.IsSet()
}

// HasSkipEmbeddingGeneration returns a boolean if a field has been set.
func (o *GitbookConnectRequest) HasSkipEmbeddingGeneration() bool {
	if o != nil && o.SkipEmbeddingGeneration.IsSet() {
		return true
	}

	return false
}

// SetSkipEmbeddingGeneration gets a reference to the given NullableBool and assigns it to the SkipEmbeddingGeneration field.
func (o *GitbookConnectRequest) SetSkipEmbeddingGeneration(v bool) {
	o.SkipEmbeddingGeneration.Set(&v)
}
// SetSkipEmbeddingGenerationNil sets the value for SkipEmbeddingGeneration to be an explicit nil
func (o *GitbookConnectRequest) SetSkipEmbeddingGenerationNil() {
	o.SkipEmbeddingGeneration.Set(nil)
}

// UnsetSkipEmbeddingGeneration ensures that no value is present for SkipEmbeddingGeneration, not even an explicit nil
func (o *GitbookConnectRequest) UnsetSkipEmbeddingGeneration() {
	o.SkipEmbeddingGeneration.Unset()
}

// GetEmbeddingModel returns the EmbeddingModel field value if set, zero value otherwise.
func (o *GitbookConnectRequest) GetEmbeddingModel() EmbeddingGenerators {
	if o == nil || isNil(o.EmbeddingModel) {
		var ret EmbeddingGenerators
		return ret
	}
	return *o.EmbeddingModel
}

// GetEmbeddingModelOk returns a tuple with the EmbeddingModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitbookConnectRequest) GetEmbeddingModelOk() (*EmbeddingGenerators, bool) {
	if o == nil || isNil(o.EmbeddingModel) {
    return nil, false
	}
	return o.EmbeddingModel, true
}

// HasEmbeddingModel returns a boolean if a field has been set.
func (o *GitbookConnectRequest) HasEmbeddingModel() bool {
	if o != nil && !isNil(o.EmbeddingModel) {
		return true
	}

	return false
}

// SetEmbeddingModel gets a reference to the given EmbeddingGenerators and assigns it to the EmbeddingModel field.
func (o *GitbookConnectRequest) SetEmbeddingModel(v EmbeddingGenerators) {
	o.EmbeddingModel = &v
}

// GetGenerateSparseVectors returns the GenerateSparseVectors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitbookConnectRequest) GetGenerateSparseVectors() bool {
	if o == nil || isNil(o.GenerateSparseVectors.Get()) {
		var ret bool
		return ret
	}
	return *o.GenerateSparseVectors.Get()
}

// GetGenerateSparseVectorsOk returns a tuple with the GenerateSparseVectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitbookConnectRequest) GetGenerateSparseVectorsOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.GenerateSparseVectors.Get(), o.GenerateSparseVectors.IsSet()
}

// HasGenerateSparseVectors returns a boolean if a field has been set.
func (o *GitbookConnectRequest) HasGenerateSparseVectors() bool {
	if o != nil && o.GenerateSparseVectors.IsSet() {
		return true
	}

	return false
}

// SetGenerateSparseVectors gets a reference to the given NullableBool and assigns it to the GenerateSparseVectors field.
func (o *GitbookConnectRequest) SetGenerateSparseVectors(v bool) {
	o.GenerateSparseVectors.Set(&v)
}
// SetGenerateSparseVectorsNil sets the value for GenerateSparseVectors to be an explicit nil
func (o *GitbookConnectRequest) SetGenerateSparseVectorsNil() {
	o.GenerateSparseVectors.Set(nil)
}

// UnsetGenerateSparseVectors ensures that no value is present for GenerateSparseVectors, not even an explicit nil
func (o *GitbookConnectRequest) UnsetGenerateSparseVectors() {
	o.GenerateSparseVectors.Unset()
}

// GetPrependFilenameToChunks returns the PrependFilenameToChunks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitbookConnectRequest) GetPrependFilenameToChunks() bool {
	if o == nil || isNil(o.PrependFilenameToChunks.Get()) {
		var ret bool
		return ret
	}
	return *o.PrependFilenameToChunks.Get()
}

// GetPrependFilenameToChunksOk returns a tuple with the PrependFilenameToChunks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitbookConnectRequest) GetPrependFilenameToChunksOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.PrependFilenameToChunks.Get(), o.PrependFilenameToChunks.IsSet()
}

// HasPrependFilenameToChunks returns a boolean if a field has been set.
func (o *GitbookConnectRequest) HasPrependFilenameToChunks() bool {
	if o != nil && o.PrependFilenameToChunks.IsSet() {
		return true
	}

	return false
}

// SetPrependFilenameToChunks gets a reference to the given NullableBool and assigns it to the PrependFilenameToChunks field.
func (o *GitbookConnectRequest) SetPrependFilenameToChunks(v bool) {
	o.PrependFilenameToChunks.Set(&v)
}
// SetPrependFilenameToChunksNil sets the value for PrependFilenameToChunks to be an explicit nil
func (o *GitbookConnectRequest) SetPrependFilenameToChunksNil() {
	o.PrependFilenameToChunks.Set(nil)
}

// UnsetPrependFilenameToChunks ensures that no value is present for PrependFilenameToChunks, not even an explicit nil
func (o *GitbookConnectRequest) UnsetPrependFilenameToChunks() {
	o.PrependFilenameToChunks.Unset()
}

// GetSyncFilesOnConnection returns the SyncFilesOnConnection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitbookConnectRequest) GetSyncFilesOnConnection() bool {
	if o == nil || isNil(o.SyncFilesOnConnection.Get()) {
		var ret bool
		return ret
	}
	return *o.SyncFilesOnConnection.Get()
}

// GetSyncFilesOnConnectionOk returns a tuple with the SyncFilesOnConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitbookConnectRequest) GetSyncFilesOnConnectionOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.SyncFilesOnConnection.Get(), o.SyncFilesOnConnection.IsSet()
}

// HasSyncFilesOnConnection returns a boolean if a field has been set.
func (o *GitbookConnectRequest) HasSyncFilesOnConnection() bool {
	if o != nil && o.SyncFilesOnConnection.IsSet() {
		return true
	}

	return false
}

// SetSyncFilesOnConnection gets a reference to the given NullableBool and assigns it to the SyncFilesOnConnection field.
func (o *GitbookConnectRequest) SetSyncFilesOnConnection(v bool) {
	o.SyncFilesOnConnection.Set(&v)
}
// SetSyncFilesOnConnectionNil sets the value for SyncFilesOnConnection to be an explicit nil
func (o *GitbookConnectRequest) SetSyncFilesOnConnectionNil() {
	o.SyncFilesOnConnection.Set(nil)
}

// UnsetSyncFilesOnConnection ensures that no value is present for SyncFilesOnConnection, not even an explicit nil
func (o *GitbookConnectRequest) UnsetSyncFilesOnConnection() {
	o.SyncFilesOnConnection.Unset()
}

// GetRequestId returns the RequestId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitbookConnectRequest) GetRequestId() string {
	if o == nil || isNil(o.RequestId.Get()) {
		var ret string
		return ret
	}
	return *o.RequestId.Get()
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitbookConnectRequest) GetRequestIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RequestId.Get(), o.RequestId.IsSet()
}

// HasRequestId returns a boolean if a field has been set.
func (o *GitbookConnectRequest) HasRequestId() bool {
	if o != nil && o.RequestId.IsSet() {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given NullableString and assigns it to the RequestId field.
func (o *GitbookConnectRequest) SetRequestId(v string) {
	o.RequestId.Set(&v)
}
// SetRequestIdNil sets the value for RequestId to be an explicit nil
func (o *GitbookConnectRequest) SetRequestIdNil() {
	o.RequestId.Set(nil)
}

// UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
func (o *GitbookConnectRequest) UnsetRequestId() {
	o.RequestId.Unset()
}

// GetSyncSourceItems returns the SyncSourceItems field value if set, zero value otherwise.
func (o *GitbookConnectRequest) GetSyncSourceItems() bool {
	if o == nil || isNil(o.SyncSourceItems) {
		var ret bool
		return ret
	}
	return *o.SyncSourceItems
}

// GetSyncSourceItemsOk returns a tuple with the SyncSourceItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitbookConnectRequest) GetSyncSourceItemsOk() (*bool, bool) {
	if o == nil || isNil(o.SyncSourceItems) {
    return nil, false
	}
	return o.SyncSourceItems, true
}

// HasSyncSourceItems returns a boolean if a field has been set.
func (o *GitbookConnectRequest) HasSyncSourceItems() bool {
	if o != nil && !isNil(o.SyncSourceItems) {
		return true
	}

	return false
}

// SetSyncSourceItems gets a reference to the given bool and assigns it to the SyncSourceItems field.
func (o *GitbookConnectRequest) SetSyncSourceItems(v bool) {
	o.SyncSourceItems = &v
}

// GetFileSyncConfig returns the FileSyncConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitbookConnectRequest) GetFileSyncConfig() FileSyncConfigNullable {
	if o == nil || isNil(o.FileSyncConfig.Get()) {
		var ret FileSyncConfigNullable
		return ret
	}
	return *o.FileSyncConfig.Get()
}

// GetFileSyncConfigOk returns a tuple with the FileSyncConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitbookConnectRequest) GetFileSyncConfigOk() (*FileSyncConfigNullable, bool) {
	if o == nil {
    return nil, false
	}
	return o.FileSyncConfig.Get(), o.FileSyncConfig.IsSet()
}

// HasFileSyncConfig returns a boolean if a field has been set.
func (o *GitbookConnectRequest) HasFileSyncConfig() bool {
	if o != nil && o.FileSyncConfig.IsSet() {
		return true
	}

	return false
}

// SetFileSyncConfig gets a reference to the given NullableFileSyncConfigNullable and assigns it to the FileSyncConfig field.
func (o *GitbookConnectRequest) SetFileSyncConfig(v FileSyncConfigNullable) {
	o.FileSyncConfig.Set(&v)
}
// SetFileSyncConfigNil sets the value for FileSyncConfig to be an explicit nil
func (o *GitbookConnectRequest) SetFileSyncConfigNil() {
	o.FileSyncConfig.Set(nil)
}

// UnsetFileSyncConfig ensures that no value is present for FileSyncConfig, not even an explicit nil
func (o *GitbookConnectRequest) UnsetFileSyncConfig() {
	o.FileSyncConfig.Unset()
}

// GetDataSourceTags returns the DataSourceTags field value if set, zero value otherwise.
func (o *GitbookConnectRequest) GetDataSourceTags() map[string]interface{} {
	if o == nil || isNil(o.DataSourceTags) {
		var ret map[string]interface{}
		return ret
	}
	return o.DataSourceTags
}

// GetDataSourceTagsOk returns a tuple with the DataSourceTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitbookConnectRequest) GetDataSourceTagsOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.DataSourceTags) {
    return map[string]interface{}{}, false
	}
	return o.DataSourceTags, true
}

// HasDataSourceTags returns a boolean if a field has been set.
func (o *GitbookConnectRequest) HasDataSourceTags() bool {
	if o != nil && !isNil(o.DataSourceTags) {
		return true
	}

	return false
}

// SetDataSourceTags gets a reference to the given map[string]interface{} and assigns it to the DataSourceTags field.
func (o *GitbookConnectRequest) SetDataSourceTags(v map[string]interface{}) {
	o.DataSourceTags = v
}

func (o GitbookConnectRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["organization"] = o.Organization
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
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
	if o.SyncFilesOnConnection.IsSet() {
		toSerialize["sync_files_on_connection"] = o.SyncFilesOnConnection.Get()
	}
	if o.RequestId.IsSet() {
		toSerialize["request_id"] = o.RequestId.Get()
	}
	if !isNil(o.SyncSourceItems) {
		toSerialize["sync_source_items"] = o.SyncSourceItems
	}
	if o.FileSyncConfig.IsSet() {
		toSerialize["file_sync_config"] = o.FileSyncConfig.Get()
	}
	if !isNil(o.DataSourceTags) {
		toSerialize["data_source_tags"] = o.DataSourceTags
	}
	return json.Marshal(toSerialize)
}

type NullableGitbookConnectRequest struct {
	value *GitbookConnectRequest
	isSet bool
}

func (v NullableGitbookConnectRequest) Get() *GitbookConnectRequest {
	return v.value
}

func (v *NullableGitbookConnectRequest) Set(val *GitbookConnectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGitbookConnectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGitbookConnectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitbookConnectRequest(val *GitbookConnectRequest) *NullableGitbookConnectRequest {
	return &NullableGitbookConnectRequest{value: val, isSet: true}
}

func (v NullableGitbookConnectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitbookConnectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



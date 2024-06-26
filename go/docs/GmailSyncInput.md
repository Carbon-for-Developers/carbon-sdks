# GmailSyncInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **map[string]interface{}** |  | [optional] [default to {}]
**Filters** | **map[string]interface{}** |  | 
**ChunkSize** | Pointer to **NullableInt32** |  | [optional] [default to 1500]
**ChunkOverlap** | Pointer to **NullableInt32** |  | [optional] [default to 20]
**SkipEmbeddingGeneration** | Pointer to **NullableBool** |  | [optional] [default to false]
**EmbeddingModel** | Pointer to [**EmbeddingGenerators**](EmbeddingGenerators.md) |  | [optional] 
**GenerateSparseVectors** | Pointer to **NullableBool** |  | [optional] [default to false]
**PrependFilenameToChunks** | Pointer to **NullableBool** |  | [optional] [default to false]
**DataSourceId** | Pointer to **NullableInt32** |  | [optional] 
**RequestId** | Pointer to **NullableString** |  | [optional] 
**SyncAttachments** | Pointer to **NullableBool** |  | [optional] [default to false]
**FileSyncConfig** | Pointer to [**NullableFileSyncConfigNullable**](FileSyncConfigNullable.md) |  | [optional] 
**IncrementalSync** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewGmailSyncInput

`func NewGmailSyncInput(filters map[string]interface{}, ) *GmailSyncInput`

NewGmailSyncInput instantiates a new GmailSyncInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGmailSyncInputWithDefaults

`func NewGmailSyncInputWithDefaults() *GmailSyncInput`

NewGmailSyncInputWithDefaults instantiates a new GmailSyncInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *GmailSyncInput) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GmailSyncInput) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GmailSyncInput) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *GmailSyncInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *GmailSyncInput) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *GmailSyncInput) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetFilters

`func (o *GmailSyncInput) GetFilters() map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GmailSyncInput) GetFiltersOk() (*map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GmailSyncInput) SetFilters(v map[string]interface{})`

SetFilters sets Filters field to given value.


### GetChunkSize

`func (o *GmailSyncInput) GetChunkSize() int32`

GetChunkSize returns the ChunkSize field if non-nil, zero value otherwise.

### GetChunkSizeOk

`func (o *GmailSyncInput) GetChunkSizeOk() (*int32, bool)`

GetChunkSizeOk returns a tuple with the ChunkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkSize

`func (o *GmailSyncInput) SetChunkSize(v int32)`

SetChunkSize sets ChunkSize field to given value.

### HasChunkSize

`func (o *GmailSyncInput) HasChunkSize() bool`

HasChunkSize returns a boolean if a field has been set.

### SetChunkSizeNil

`func (o *GmailSyncInput) SetChunkSizeNil(b bool)`

 SetChunkSizeNil sets the value for ChunkSize to be an explicit nil

### UnsetChunkSize
`func (o *GmailSyncInput) UnsetChunkSize()`

UnsetChunkSize ensures that no value is present for ChunkSize, not even an explicit nil
### GetChunkOverlap

`func (o *GmailSyncInput) GetChunkOverlap() int32`

GetChunkOverlap returns the ChunkOverlap field if non-nil, zero value otherwise.

### GetChunkOverlapOk

`func (o *GmailSyncInput) GetChunkOverlapOk() (*int32, bool)`

GetChunkOverlapOk returns a tuple with the ChunkOverlap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkOverlap

`func (o *GmailSyncInput) SetChunkOverlap(v int32)`

SetChunkOverlap sets ChunkOverlap field to given value.

### HasChunkOverlap

`func (o *GmailSyncInput) HasChunkOverlap() bool`

HasChunkOverlap returns a boolean if a field has been set.

### SetChunkOverlapNil

`func (o *GmailSyncInput) SetChunkOverlapNil(b bool)`

 SetChunkOverlapNil sets the value for ChunkOverlap to be an explicit nil

### UnsetChunkOverlap
`func (o *GmailSyncInput) UnsetChunkOverlap()`

UnsetChunkOverlap ensures that no value is present for ChunkOverlap, not even an explicit nil
### GetSkipEmbeddingGeneration

`func (o *GmailSyncInput) GetSkipEmbeddingGeneration() bool`

GetSkipEmbeddingGeneration returns the SkipEmbeddingGeneration field if non-nil, zero value otherwise.

### GetSkipEmbeddingGenerationOk

`func (o *GmailSyncInput) GetSkipEmbeddingGenerationOk() (*bool, bool)`

GetSkipEmbeddingGenerationOk returns a tuple with the SkipEmbeddingGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipEmbeddingGeneration

`func (o *GmailSyncInput) SetSkipEmbeddingGeneration(v bool)`

SetSkipEmbeddingGeneration sets SkipEmbeddingGeneration field to given value.

### HasSkipEmbeddingGeneration

`func (o *GmailSyncInput) HasSkipEmbeddingGeneration() bool`

HasSkipEmbeddingGeneration returns a boolean if a field has been set.

### SetSkipEmbeddingGenerationNil

`func (o *GmailSyncInput) SetSkipEmbeddingGenerationNil(b bool)`

 SetSkipEmbeddingGenerationNil sets the value for SkipEmbeddingGeneration to be an explicit nil

### UnsetSkipEmbeddingGeneration
`func (o *GmailSyncInput) UnsetSkipEmbeddingGeneration()`

UnsetSkipEmbeddingGeneration ensures that no value is present for SkipEmbeddingGeneration, not even an explicit nil
### GetEmbeddingModel

`func (o *GmailSyncInput) GetEmbeddingModel() EmbeddingGenerators`

GetEmbeddingModel returns the EmbeddingModel field if non-nil, zero value otherwise.

### GetEmbeddingModelOk

`func (o *GmailSyncInput) GetEmbeddingModelOk() (*EmbeddingGenerators, bool)`

GetEmbeddingModelOk returns a tuple with the EmbeddingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingModel

`func (o *GmailSyncInput) SetEmbeddingModel(v EmbeddingGenerators)`

SetEmbeddingModel sets EmbeddingModel field to given value.

### HasEmbeddingModel

`func (o *GmailSyncInput) HasEmbeddingModel() bool`

HasEmbeddingModel returns a boolean if a field has been set.

### GetGenerateSparseVectors

`func (o *GmailSyncInput) GetGenerateSparseVectors() bool`

GetGenerateSparseVectors returns the GenerateSparseVectors field if non-nil, zero value otherwise.

### GetGenerateSparseVectorsOk

`func (o *GmailSyncInput) GetGenerateSparseVectorsOk() (*bool, bool)`

GetGenerateSparseVectorsOk returns a tuple with the GenerateSparseVectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateSparseVectors

`func (o *GmailSyncInput) SetGenerateSparseVectors(v bool)`

SetGenerateSparseVectors sets GenerateSparseVectors field to given value.

### HasGenerateSparseVectors

`func (o *GmailSyncInput) HasGenerateSparseVectors() bool`

HasGenerateSparseVectors returns a boolean if a field has been set.

### SetGenerateSparseVectorsNil

`func (o *GmailSyncInput) SetGenerateSparseVectorsNil(b bool)`

 SetGenerateSparseVectorsNil sets the value for GenerateSparseVectors to be an explicit nil

### UnsetGenerateSparseVectors
`func (o *GmailSyncInput) UnsetGenerateSparseVectors()`

UnsetGenerateSparseVectors ensures that no value is present for GenerateSparseVectors, not even an explicit nil
### GetPrependFilenameToChunks

`func (o *GmailSyncInput) GetPrependFilenameToChunks() bool`

GetPrependFilenameToChunks returns the PrependFilenameToChunks field if non-nil, zero value otherwise.

### GetPrependFilenameToChunksOk

`func (o *GmailSyncInput) GetPrependFilenameToChunksOk() (*bool, bool)`

GetPrependFilenameToChunksOk returns a tuple with the PrependFilenameToChunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrependFilenameToChunks

`func (o *GmailSyncInput) SetPrependFilenameToChunks(v bool)`

SetPrependFilenameToChunks sets PrependFilenameToChunks field to given value.

### HasPrependFilenameToChunks

`func (o *GmailSyncInput) HasPrependFilenameToChunks() bool`

HasPrependFilenameToChunks returns a boolean if a field has been set.

### SetPrependFilenameToChunksNil

`func (o *GmailSyncInput) SetPrependFilenameToChunksNil(b bool)`

 SetPrependFilenameToChunksNil sets the value for PrependFilenameToChunks to be an explicit nil

### UnsetPrependFilenameToChunks
`func (o *GmailSyncInput) UnsetPrependFilenameToChunks()`

UnsetPrependFilenameToChunks ensures that no value is present for PrependFilenameToChunks, not even an explicit nil
### GetDataSourceId

`func (o *GmailSyncInput) GetDataSourceId() int32`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *GmailSyncInput) GetDataSourceIdOk() (*int32, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *GmailSyncInput) SetDataSourceId(v int32)`

SetDataSourceId sets DataSourceId field to given value.

### HasDataSourceId

`func (o *GmailSyncInput) HasDataSourceId() bool`

HasDataSourceId returns a boolean if a field has been set.

### SetDataSourceIdNil

`func (o *GmailSyncInput) SetDataSourceIdNil(b bool)`

 SetDataSourceIdNil sets the value for DataSourceId to be an explicit nil

### UnsetDataSourceId
`func (o *GmailSyncInput) UnsetDataSourceId()`

UnsetDataSourceId ensures that no value is present for DataSourceId, not even an explicit nil
### GetRequestId

`func (o *GmailSyncInput) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GmailSyncInput) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GmailSyncInput) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *GmailSyncInput) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### SetRequestIdNil

`func (o *GmailSyncInput) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *GmailSyncInput) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetSyncAttachments

`func (o *GmailSyncInput) GetSyncAttachments() bool`

GetSyncAttachments returns the SyncAttachments field if non-nil, zero value otherwise.

### GetSyncAttachmentsOk

`func (o *GmailSyncInput) GetSyncAttachmentsOk() (*bool, bool)`

GetSyncAttachmentsOk returns a tuple with the SyncAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncAttachments

`func (o *GmailSyncInput) SetSyncAttachments(v bool)`

SetSyncAttachments sets SyncAttachments field to given value.

### HasSyncAttachments

`func (o *GmailSyncInput) HasSyncAttachments() bool`

HasSyncAttachments returns a boolean if a field has been set.

### SetSyncAttachmentsNil

`func (o *GmailSyncInput) SetSyncAttachmentsNil(b bool)`

 SetSyncAttachmentsNil sets the value for SyncAttachments to be an explicit nil

### UnsetSyncAttachments
`func (o *GmailSyncInput) UnsetSyncAttachments()`

UnsetSyncAttachments ensures that no value is present for SyncAttachments, not even an explicit nil
### GetFileSyncConfig

`func (o *GmailSyncInput) GetFileSyncConfig() FileSyncConfigNullable`

GetFileSyncConfig returns the FileSyncConfig field if non-nil, zero value otherwise.

### GetFileSyncConfigOk

`func (o *GmailSyncInput) GetFileSyncConfigOk() (*FileSyncConfigNullable, bool)`

GetFileSyncConfigOk returns a tuple with the FileSyncConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSyncConfig

`func (o *GmailSyncInput) SetFileSyncConfig(v FileSyncConfigNullable)`

SetFileSyncConfig sets FileSyncConfig field to given value.

### HasFileSyncConfig

`func (o *GmailSyncInput) HasFileSyncConfig() bool`

HasFileSyncConfig returns a boolean if a field has been set.

### SetFileSyncConfigNil

`func (o *GmailSyncInput) SetFileSyncConfigNil(b bool)`

 SetFileSyncConfigNil sets the value for FileSyncConfig to be an explicit nil

### UnsetFileSyncConfig
`func (o *GmailSyncInput) UnsetFileSyncConfig()`

UnsetFileSyncConfig ensures that no value is present for FileSyncConfig, not even an explicit nil
### GetIncrementalSync

`func (o *GmailSyncInput) GetIncrementalSync() bool`

GetIncrementalSync returns the IncrementalSync field if non-nil, zero value otherwise.

### GetIncrementalSyncOk

`func (o *GmailSyncInput) GetIncrementalSyncOk() (*bool, bool)`

GetIncrementalSyncOk returns a tuple with the IncrementalSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalSync

`func (o *GmailSyncInput) SetIncrementalSync(v bool)`

SetIncrementalSync sets IncrementalSync field to given value.

### HasIncrementalSync

`func (o *GmailSyncInput) HasIncrementalSync() bool`

HasIncrementalSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



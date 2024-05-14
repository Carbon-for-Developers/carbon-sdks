# SyncOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **interface{}** |  | [optional] 
**ChunkSize** | Pointer to **NullableInt32** |  | [optional] [default to 1500]
**ChunkOverlap** | Pointer to **NullableInt32** |  | [optional] [default to 20]
**SkipEmbeddingGeneration** | Pointer to **NullableBool** |  | [optional] [default to false]
**EmbeddingModel** | Pointer to [**NullableEmbeddingGeneratorsNullable**](EmbeddingGeneratorsNullable.md) |  | [optional] [default to EMBEDDINGGENERATORSNULLABLE_OPENAI]
**GenerateSparseVectors** | Pointer to **NullableBool** |  | [optional] [default to false]
**PrependFilenameToChunks** | Pointer to **NullableBool** |  | [optional] [default to false]
**MaxItemsPerChunk** | Pointer to **NullableInt32** | Number of objects per chunk. For csv, tsv, xlsx, and json files only. | [optional] 
**SyncFilesOnConnection** | Pointer to **NullableBool** | Used to specify whether Carbon should attempt to sync all your files automatically when authorization         is complete. This is only supported for a subset of connectors and will be ignored for the rest. Supported         connectors: Intercom, Zendesk, Gitbook, Confluence, Salesforce, Freshdesk | [optional] [default to true]
**SetPageAsBoundary** | Pointer to **bool** |  | [optional] [default to false]
**RequestId** | Pointer to **string** |  | [optional] [default to "30d4247e-9dcf-41e3-98a7-61ef4a78adad"]
**EnableFilePicker** | Pointer to **bool** |  | [optional] [default to true]
**SyncSourceItems** | Pointer to **bool** | Enabling this flag will fetch all available content from the source to be listed via list items endpoint | [optional] [default to true]
**IncrementalSync** | Pointer to **bool** | Only sync files if they have not already been synced or if the embedding properties have changed.         This flag is currently supported by ONEDRIVE, GOOGLE_DRIVE, BOX, DROPBOX. It will be ignored for other data sources. | [optional] [default to false]
**FileSyncConfig** | Pointer to [**NullableHelpdeskFileSyncConfigNullable**](HelpdeskFileSyncConfigNullable.md) |  | [optional] 

## Methods

### NewSyncOptions

`func NewSyncOptions() *SyncOptions`

NewSyncOptions instantiates a new SyncOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncOptionsWithDefaults

`func NewSyncOptionsWithDefaults() *SyncOptions`

NewSyncOptionsWithDefaults instantiates a new SyncOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *SyncOptions) GetTags() interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SyncOptions) GetTagsOk() (*interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SyncOptions) SetTags(v interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *SyncOptions) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *SyncOptions) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *SyncOptions) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetChunkSize

`func (o *SyncOptions) GetChunkSize() int32`

GetChunkSize returns the ChunkSize field if non-nil, zero value otherwise.

### GetChunkSizeOk

`func (o *SyncOptions) GetChunkSizeOk() (*int32, bool)`

GetChunkSizeOk returns a tuple with the ChunkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkSize

`func (o *SyncOptions) SetChunkSize(v int32)`

SetChunkSize sets ChunkSize field to given value.

### HasChunkSize

`func (o *SyncOptions) HasChunkSize() bool`

HasChunkSize returns a boolean if a field has been set.

### SetChunkSizeNil

`func (o *SyncOptions) SetChunkSizeNil(b bool)`

 SetChunkSizeNil sets the value for ChunkSize to be an explicit nil

### UnsetChunkSize
`func (o *SyncOptions) UnsetChunkSize()`

UnsetChunkSize ensures that no value is present for ChunkSize, not even an explicit nil
### GetChunkOverlap

`func (o *SyncOptions) GetChunkOverlap() int32`

GetChunkOverlap returns the ChunkOverlap field if non-nil, zero value otherwise.

### GetChunkOverlapOk

`func (o *SyncOptions) GetChunkOverlapOk() (*int32, bool)`

GetChunkOverlapOk returns a tuple with the ChunkOverlap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkOverlap

`func (o *SyncOptions) SetChunkOverlap(v int32)`

SetChunkOverlap sets ChunkOverlap field to given value.

### HasChunkOverlap

`func (o *SyncOptions) HasChunkOverlap() bool`

HasChunkOverlap returns a boolean if a field has been set.

### SetChunkOverlapNil

`func (o *SyncOptions) SetChunkOverlapNil(b bool)`

 SetChunkOverlapNil sets the value for ChunkOverlap to be an explicit nil

### UnsetChunkOverlap
`func (o *SyncOptions) UnsetChunkOverlap()`

UnsetChunkOverlap ensures that no value is present for ChunkOverlap, not even an explicit nil
### GetSkipEmbeddingGeneration

`func (o *SyncOptions) GetSkipEmbeddingGeneration() bool`

GetSkipEmbeddingGeneration returns the SkipEmbeddingGeneration field if non-nil, zero value otherwise.

### GetSkipEmbeddingGenerationOk

`func (o *SyncOptions) GetSkipEmbeddingGenerationOk() (*bool, bool)`

GetSkipEmbeddingGenerationOk returns a tuple with the SkipEmbeddingGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipEmbeddingGeneration

`func (o *SyncOptions) SetSkipEmbeddingGeneration(v bool)`

SetSkipEmbeddingGeneration sets SkipEmbeddingGeneration field to given value.

### HasSkipEmbeddingGeneration

`func (o *SyncOptions) HasSkipEmbeddingGeneration() bool`

HasSkipEmbeddingGeneration returns a boolean if a field has been set.

### SetSkipEmbeddingGenerationNil

`func (o *SyncOptions) SetSkipEmbeddingGenerationNil(b bool)`

 SetSkipEmbeddingGenerationNil sets the value for SkipEmbeddingGeneration to be an explicit nil

### UnsetSkipEmbeddingGeneration
`func (o *SyncOptions) UnsetSkipEmbeddingGeneration()`

UnsetSkipEmbeddingGeneration ensures that no value is present for SkipEmbeddingGeneration, not even an explicit nil
### GetEmbeddingModel

`func (o *SyncOptions) GetEmbeddingModel() EmbeddingGeneratorsNullable`

GetEmbeddingModel returns the EmbeddingModel field if non-nil, zero value otherwise.

### GetEmbeddingModelOk

`func (o *SyncOptions) GetEmbeddingModelOk() (*EmbeddingGeneratorsNullable, bool)`

GetEmbeddingModelOk returns a tuple with the EmbeddingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingModel

`func (o *SyncOptions) SetEmbeddingModel(v EmbeddingGeneratorsNullable)`

SetEmbeddingModel sets EmbeddingModel field to given value.

### HasEmbeddingModel

`func (o *SyncOptions) HasEmbeddingModel() bool`

HasEmbeddingModel returns a boolean if a field has been set.

### SetEmbeddingModelNil

`func (o *SyncOptions) SetEmbeddingModelNil(b bool)`

 SetEmbeddingModelNil sets the value for EmbeddingModel to be an explicit nil

### UnsetEmbeddingModel
`func (o *SyncOptions) UnsetEmbeddingModel()`

UnsetEmbeddingModel ensures that no value is present for EmbeddingModel, not even an explicit nil
### GetGenerateSparseVectors

`func (o *SyncOptions) GetGenerateSparseVectors() bool`

GetGenerateSparseVectors returns the GenerateSparseVectors field if non-nil, zero value otherwise.

### GetGenerateSparseVectorsOk

`func (o *SyncOptions) GetGenerateSparseVectorsOk() (*bool, bool)`

GetGenerateSparseVectorsOk returns a tuple with the GenerateSparseVectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateSparseVectors

`func (o *SyncOptions) SetGenerateSparseVectors(v bool)`

SetGenerateSparseVectors sets GenerateSparseVectors field to given value.

### HasGenerateSparseVectors

`func (o *SyncOptions) HasGenerateSparseVectors() bool`

HasGenerateSparseVectors returns a boolean if a field has been set.

### SetGenerateSparseVectorsNil

`func (o *SyncOptions) SetGenerateSparseVectorsNil(b bool)`

 SetGenerateSparseVectorsNil sets the value for GenerateSparseVectors to be an explicit nil

### UnsetGenerateSparseVectors
`func (o *SyncOptions) UnsetGenerateSparseVectors()`

UnsetGenerateSparseVectors ensures that no value is present for GenerateSparseVectors, not even an explicit nil
### GetPrependFilenameToChunks

`func (o *SyncOptions) GetPrependFilenameToChunks() bool`

GetPrependFilenameToChunks returns the PrependFilenameToChunks field if non-nil, zero value otherwise.

### GetPrependFilenameToChunksOk

`func (o *SyncOptions) GetPrependFilenameToChunksOk() (*bool, bool)`

GetPrependFilenameToChunksOk returns a tuple with the PrependFilenameToChunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrependFilenameToChunks

`func (o *SyncOptions) SetPrependFilenameToChunks(v bool)`

SetPrependFilenameToChunks sets PrependFilenameToChunks field to given value.

### HasPrependFilenameToChunks

`func (o *SyncOptions) HasPrependFilenameToChunks() bool`

HasPrependFilenameToChunks returns a boolean if a field has been set.

### SetPrependFilenameToChunksNil

`func (o *SyncOptions) SetPrependFilenameToChunksNil(b bool)`

 SetPrependFilenameToChunksNil sets the value for PrependFilenameToChunks to be an explicit nil

### UnsetPrependFilenameToChunks
`func (o *SyncOptions) UnsetPrependFilenameToChunks()`

UnsetPrependFilenameToChunks ensures that no value is present for PrependFilenameToChunks, not even an explicit nil
### GetMaxItemsPerChunk

`func (o *SyncOptions) GetMaxItemsPerChunk() int32`

GetMaxItemsPerChunk returns the MaxItemsPerChunk field if non-nil, zero value otherwise.

### GetMaxItemsPerChunkOk

`func (o *SyncOptions) GetMaxItemsPerChunkOk() (*int32, bool)`

GetMaxItemsPerChunkOk returns a tuple with the MaxItemsPerChunk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxItemsPerChunk

`func (o *SyncOptions) SetMaxItemsPerChunk(v int32)`

SetMaxItemsPerChunk sets MaxItemsPerChunk field to given value.

### HasMaxItemsPerChunk

`func (o *SyncOptions) HasMaxItemsPerChunk() bool`

HasMaxItemsPerChunk returns a boolean if a field has been set.

### SetMaxItemsPerChunkNil

`func (o *SyncOptions) SetMaxItemsPerChunkNil(b bool)`

 SetMaxItemsPerChunkNil sets the value for MaxItemsPerChunk to be an explicit nil

### UnsetMaxItemsPerChunk
`func (o *SyncOptions) UnsetMaxItemsPerChunk()`

UnsetMaxItemsPerChunk ensures that no value is present for MaxItemsPerChunk, not even an explicit nil
### GetSyncFilesOnConnection

`func (o *SyncOptions) GetSyncFilesOnConnection() bool`

GetSyncFilesOnConnection returns the SyncFilesOnConnection field if non-nil, zero value otherwise.

### GetSyncFilesOnConnectionOk

`func (o *SyncOptions) GetSyncFilesOnConnectionOk() (*bool, bool)`

GetSyncFilesOnConnectionOk returns a tuple with the SyncFilesOnConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncFilesOnConnection

`func (o *SyncOptions) SetSyncFilesOnConnection(v bool)`

SetSyncFilesOnConnection sets SyncFilesOnConnection field to given value.

### HasSyncFilesOnConnection

`func (o *SyncOptions) HasSyncFilesOnConnection() bool`

HasSyncFilesOnConnection returns a boolean if a field has been set.

### SetSyncFilesOnConnectionNil

`func (o *SyncOptions) SetSyncFilesOnConnectionNil(b bool)`

 SetSyncFilesOnConnectionNil sets the value for SyncFilesOnConnection to be an explicit nil

### UnsetSyncFilesOnConnection
`func (o *SyncOptions) UnsetSyncFilesOnConnection()`

UnsetSyncFilesOnConnection ensures that no value is present for SyncFilesOnConnection, not even an explicit nil
### GetSetPageAsBoundary

`func (o *SyncOptions) GetSetPageAsBoundary() bool`

GetSetPageAsBoundary returns the SetPageAsBoundary field if non-nil, zero value otherwise.

### GetSetPageAsBoundaryOk

`func (o *SyncOptions) GetSetPageAsBoundaryOk() (*bool, bool)`

GetSetPageAsBoundaryOk returns a tuple with the SetPageAsBoundary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetPageAsBoundary

`func (o *SyncOptions) SetSetPageAsBoundary(v bool)`

SetSetPageAsBoundary sets SetPageAsBoundary field to given value.

### HasSetPageAsBoundary

`func (o *SyncOptions) HasSetPageAsBoundary() bool`

HasSetPageAsBoundary returns a boolean if a field has been set.

### GetRequestId

`func (o *SyncOptions) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SyncOptions) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SyncOptions) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *SyncOptions) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetEnableFilePicker

`func (o *SyncOptions) GetEnableFilePicker() bool`

GetEnableFilePicker returns the EnableFilePicker field if non-nil, zero value otherwise.

### GetEnableFilePickerOk

`func (o *SyncOptions) GetEnableFilePickerOk() (*bool, bool)`

GetEnableFilePickerOk returns a tuple with the EnableFilePicker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFilePicker

`func (o *SyncOptions) SetEnableFilePicker(v bool)`

SetEnableFilePicker sets EnableFilePicker field to given value.

### HasEnableFilePicker

`func (o *SyncOptions) HasEnableFilePicker() bool`

HasEnableFilePicker returns a boolean if a field has been set.

### GetSyncSourceItems

`func (o *SyncOptions) GetSyncSourceItems() bool`

GetSyncSourceItems returns the SyncSourceItems field if non-nil, zero value otherwise.

### GetSyncSourceItemsOk

`func (o *SyncOptions) GetSyncSourceItemsOk() (*bool, bool)`

GetSyncSourceItemsOk returns a tuple with the SyncSourceItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncSourceItems

`func (o *SyncOptions) SetSyncSourceItems(v bool)`

SetSyncSourceItems sets SyncSourceItems field to given value.

### HasSyncSourceItems

`func (o *SyncOptions) HasSyncSourceItems() bool`

HasSyncSourceItems returns a boolean if a field has been set.

### GetIncrementalSync

`func (o *SyncOptions) GetIncrementalSync() bool`

GetIncrementalSync returns the IncrementalSync field if non-nil, zero value otherwise.

### GetIncrementalSyncOk

`func (o *SyncOptions) GetIncrementalSyncOk() (*bool, bool)`

GetIncrementalSyncOk returns a tuple with the IncrementalSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalSync

`func (o *SyncOptions) SetIncrementalSync(v bool)`

SetIncrementalSync sets IncrementalSync field to given value.

### HasIncrementalSync

`func (o *SyncOptions) HasIncrementalSync() bool`

HasIncrementalSync returns a boolean if a field has been set.

### GetFileSyncConfig

`func (o *SyncOptions) GetFileSyncConfig() HelpdeskFileSyncConfigNullable`

GetFileSyncConfig returns the FileSyncConfig field if non-nil, zero value otherwise.

### GetFileSyncConfigOk

`func (o *SyncOptions) GetFileSyncConfigOk() (*HelpdeskFileSyncConfigNullable, bool)`

GetFileSyncConfigOk returns a tuple with the FileSyncConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSyncConfig

`func (o *SyncOptions) SetFileSyncConfig(v HelpdeskFileSyncConfigNullable)`

SetFileSyncConfig sets FileSyncConfig field to given value.

### HasFileSyncConfig

`func (o *SyncOptions) HasFileSyncConfig() bool`

HasFileSyncConfig returns a boolean if a field has been set.

### SetFileSyncConfigNil

`func (o *SyncOptions) SetFileSyncConfigNil(b bool)`

 SetFileSyncConfigNil sets the value for FileSyncConfig to be an explicit nil

### UnsetFileSyncConfig
`func (o *SyncOptions) UnsetFileSyncConfig()`

UnsetFileSyncConfig ensures that no value is present for FileSyncConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



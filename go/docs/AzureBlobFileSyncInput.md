# AzureBlobFileSyncInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **map[string]interface{}** |  | [optional] [default to {}]
**Ids** | [**[]AzureBlobGetFileInput**](AzureBlobGetFileInput.md) |  | 
**ChunkSize** | Pointer to **NullableInt32** |  | [optional] [default to 1500]
**ChunkOverlap** | Pointer to **NullableInt32** |  | [optional] [default to 20]
**SkipEmbeddingGeneration** | Pointer to **NullableBool** |  | [optional] [default to false]
**EmbeddingModel** | Pointer to [**EmbeddingGenerators**](EmbeddingGenerators.md) |  | [optional] 
**GenerateSparseVectors** | Pointer to **NullableBool** |  | [optional] [default to false]
**PrependFilenameToChunks** | Pointer to **NullableBool** |  | [optional] [default to false]
**MaxItemsPerChunk** | Pointer to **NullableInt32** | Number of objects per chunk. For csv, tsv, xlsx, and json files only. | [optional] 
**SetPageAsBoundary** | Pointer to **bool** |  | [optional] [default to false]
**DataSourceId** | Pointer to **NullableInt32** |  | [optional] 
**RequestId** | Pointer to **NullableString** |  | [optional] 
**UseOcr** | Pointer to **NullableBool** |  | [optional] [default to false]
**ParsePdfTablesWithOcr** | Pointer to **NullableBool** |  | [optional] [default to false]
**FileSyncConfig** | Pointer to [**NullableFileSyncConfigNullable**](FileSyncConfigNullable.md) |  | [optional] 

## Methods

### NewAzureBlobFileSyncInput

`func NewAzureBlobFileSyncInput(ids []AzureBlobGetFileInput, ) *AzureBlobFileSyncInput`

NewAzureBlobFileSyncInput instantiates a new AzureBlobFileSyncInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureBlobFileSyncInputWithDefaults

`func NewAzureBlobFileSyncInputWithDefaults() *AzureBlobFileSyncInput`

NewAzureBlobFileSyncInputWithDefaults instantiates a new AzureBlobFileSyncInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *AzureBlobFileSyncInput) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AzureBlobFileSyncInput) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AzureBlobFileSyncInput) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *AzureBlobFileSyncInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AzureBlobFileSyncInput) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AzureBlobFileSyncInput) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetIds

`func (o *AzureBlobFileSyncInput) GetIds() []AzureBlobGetFileInput`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *AzureBlobFileSyncInput) GetIdsOk() (*[]AzureBlobGetFileInput, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *AzureBlobFileSyncInput) SetIds(v []AzureBlobGetFileInput)`

SetIds sets Ids field to given value.


### GetChunkSize

`func (o *AzureBlobFileSyncInput) GetChunkSize() int32`

GetChunkSize returns the ChunkSize field if non-nil, zero value otherwise.

### GetChunkSizeOk

`func (o *AzureBlobFileSyncInput) GetChunkSizeOk() (*int32, bool)`

GetChunkSizeOk returns a tuple with the ChunkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkSize

`func (o *AzureBlobFileSyncInput) SetChunkSize(v int32)`

SetChunkSize sets ChunkSize field to given value.

### HasChunkSize

`func (o *AzureBlobFileSyncInput) HasChunkSize() bool`

HasChunkSize returns a boolean if a field has been set.

### SetChunkSizeNil

`func (o *AzureBlobFileSyncInput) SetChunkSizeNil(b bool)`

 SetChunkSizeNil sets the value for ChunkSize to be an explicit nil

### UnsetChunkSize
`func (o *AzureBlobFileSyncInput) UnsetChunkSize()`

UnsetChunkSize ensures that no value is present for ChunkSize, not even an explicit nil
### GetChunkOverlap

`func (o *AzureBlobFileSyncInput) GetChunkOverlap() int32`

GetChunkOverlap returns the ChunkOverlap field if non-nil, zero value otherwise.

### GetChunkOverlapOk

`func (o *AzureBlobFileSyncInput) GetChunkOverlapOk() (*int32, bool)`

GetChunkOverlapOk returns a tuple with the ChunkOverlap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkOverlap

`func (o *AzureBlobFileSyncInput) SetChunkOverlap(v int32)`

SetChunkOverlap sets ChunkOverlap field to given value.

### HasChunkOverlap

`func (o *AzureBlobFileSyncInput) HasChunkOverlap() bool`

HasChunkOverlap returns a boolean if a field has been set.

### SetChunkOverlapNil

`func (o *AzureBlobFileSyncInput) SetChunkOverlapNil(b bool)`

 SetChunkOverlapNil sets the value for ChunkOverlap to be an explicit nil

### UnsetChunkOverlap
`func (o *AzureBlobFileSyncInput) UnsetChunkOverlap()`

UnsetChunkOverlap ensures that no value is present for ChunkOverlap, not even an explicit nil
### GetSkipEmbeddingGeneration

`func (o *AzureBlobFileSyncInput) GetSkipEmbeddingGeneration() bool`

GetSkipEmbeddingGeneration returns the SkipEmbeddingGeneration field if non-nil, zero value otherwise.

### GetSkipEmbeddingGenerationOk

`func (o *AzureBlobFileSyncInput) GetSkipEmbeddingGenerationOk() (*bool, bool)`

GetSkipEmbeddingGenerationOk returns a tuple with the SkipEmbeddingGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipEmbeddingGeneration

`func (o *AzureBlobFileSyncInput) SetSkipEmbeddingGeneration(v bool)`

SetSkipEmbeddingGeneration sets SkipEmbeddingGeneration field to given value.

### HasSkipEmbeddingGeneration

`func (o *AzureBlobFileSyncInput) HasSkipEmbeddingGeneration() bool`

HasSkipEmbeddingGeneration returns a boolean if a field has been set.

### SetSkipEmbeddingGenerationNil

`func (o *AzureBlobFileSyncInput) SetSkipEmbeddingGenerationNil(b bool)`

 SetSkipEmbeddingGenerationNil sets the value for SkipEmbeddingGeneration to be an explicit nil

### UnsetSkipEmbeddingGeneration
`func (o *AzureBlobFileSyncInput) UnsetSkipEmbeddingGeneration()`

UnsetSkipEmbeddingGeneration ensures that no value is present for SkipEmbeddingGeneration, not even an explicit nil
### GetEmbeddingModel

`func (o *AzureBlobFileSyncInput) GetEmbeddingModel() EmbeddingGenerators`

GetEmbeddingModel returns the EmbeddingModel field if non-nil, zero value otherwise.

### GetEmbeddingModelOk

`func (o *AzureBlobFileSyncInput) GetEmbeddingModelOk() (*EmbeddingGenerators, bool)`

GetEmbeddingModelOk returns a tuple with the EmbeddingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingModel

`func (o *AzureBlobFileSyncInput) SetEmbeddingModel(v EmbeddingGenerators)`

SetEmbeddingModel sets EmbeddingModel field to given value.

### HasEmbeddingModel

`func (o *AzureBlobFileSyncInput) HasEmbeddingModel() bool`

HasEmbeddingModel returns a boolean if a field has been set.

### GetGenerateSparseVectors

`func (o *AzureBlobFileSyncInput) GetGenerateSparseVectors() bool`

GetGenerateSparseVectors returns the GenerateSparseVectors field if non-nil, zero value otherwise.

### GetGenerateSparseVectorsOk

`func (o *AzureBlobFileSyncInput) GetGenerateSparseVectorsOk() (*bool, bool)`

GetGenerateSparseVectorsOk returns a tuple with the GenerateSparseVectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateSparseVectors

`func (o *AzureBlobFileSyncInput) SetGenerateSparseVectors(v bool)`

SetGenerateSparseVectors sets GenerateSparseVectors field to given value.

### HasGenerateSparseVectors

`func (o *AzureBlobFileSyncInput) HasGenerateSparseVectors() bool`

HasGenerateSparseVectors returns a boolean if a field has been set.

### SetGenerateSparseVectorsNil

`func (o *AzureBlobFileSyncInput) SetGenerateSparseVectorsNil(b bool)`

 SetGenerateSparseVectorsNil sets the value for GenerateSparseVectors to be an explicit nil

### UnsetGenerateSparseVectors
`func (o *AzureBlobFileSyncInput) UnsetGenerateSparseVectors()`

UnsetGenerateSparseVectors ensures that no value is present for GenerateSparseVectors, not even an explicit nil
### GetPrependFilenameToChunks

`func (o *AzureBlobFileSyncInput) GetPrependFilenameToChunks() bool`

GetPrependFilenameToChunks returns the PrependFilenameToChunks field if non-nil, zero value otherwise.

### GetPrependFilenameToChunksOk

`func (o *AzureBlobFileSyncInput) GetPrependFilenameToChunksOk() (*bool, bool)`

GetPrependFilenameToChunksOk returns a tuple with the PrependFilenameToChunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrependFilenameToChunks

`func (o *AzureBlobFileSyncInput) SetPrependFilenameToChunks(v bool)`

SetPrependFilenameToChunks sets PrependFilenameToChunks field to given value.

### HasPrependFilenameToChunks

`func (o *AzureBlobFileSyncInput) HasPrependFilenameToChunks() bool`

HasPrependFilenameToChunks returns a boolean if a field has been set.

### SetPrependFilenameToChunksNil

`func (o *AzureBlobFileSyncInput) SetPrependFilenameToChunksNil(b bool)`

 SetPrependFilenameToChunksNil sets the value for PrependFilenameToChunks to be an explicit nil

### UnsetPrependFilenameToChunks
`func (o *AzureBlobFileSyncInput) UnsetPrependFilenameToChunks()`

UnsetPrependFilenameToChunks ensures that no value is present for PrependFilenameToChunks, not even an explicit nil
### GetMaxItemsPerChunk

`func (o *AzureBlobFileSyncInput) GetMaxItemsPerChunk() int32`

GetMaxItemsPerChunk returns the MaxItemsPerChunk field if non-nil, zero value otherwise.

### GetMaxItemsPerChunkOk

`func (o *AzureBlobFileSyncInput) GetMaxItemsPerChunkOk() (*int32, bool)`

GetMaxItemsPerChunkOk returns a tuple with the MaxItemsPerChunk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxItemsPerChunk

`func (o *AzureBlobFileSyncInput) SetMaxItemsPerChunk(v int32)`

SetMaxItemsPerChunk sets MaxItemsPerChunk field to given value.

### HasMaxItemsPerChunk

`func (o *AzureBlobFileSyncInput) HasMaxItemsPerChunk() bool`

HasMaxItemsPerChunk returns a boolean if a field has been set.

### SetMaxItemsPerChunkNil

`func (o *AzureBlobFileSyncInput) SetMaxItemsPerChunkNil(b bool)`

 SetMaxItemsPerChunkNil sets the value for MaxItemsPerChunk to be an explicit nil

### UnsetMaxItemsPerChunk
`func (o *AzureBlobFileSyncInput) UnsetMaxItemsPerChunk()`

UnsetMaxItemsPerChunk ensures that no value is present for MaxItemsPerChunk, not even an explicit nil
### GetSetPageAsBoundary

`func (o *AzureBlobFileSyncInput) GetSetPageAsBoundary() bool`

GetSetPageAsBoundary returns the SetPageAsBoundary field if non-nil, zero value otherwise.

### GetSetPageAsBoundaryOk

`func (o *AzureBlobFileSyncInput) GetSetPageAsBoundaryOk() (*bool, bool)`

GetSetPageAsBoundaryOk returns a tuple with the SetPageAsBoundary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetPageAsBoundary

`func (o *AzureBlobFileSyncInput) SetSetPageAsBoundary(v bool)`

SetSetPageAsBoundary sets SetPageAsBoundary field to given value.

### HasSetPageAsBoundary

`func (o *AzureBlobFileSyncInput) HasSetPageAsBoundary() bool`

HasSetPageAsBoundary returns a boolean if a field has been set.

### GetDataSourceId

`func (o *AzureBlobFileSyncInput) GetDataSourceId() int32`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *AzureBlobFileSyncInput) GetDataSourceIdOk() (*int32, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *AzureBlobFileSyncInput) SetDataSourceId(v int32)`

SetDataSourceId sets DataSourceId field to given value.

### HasDataSourceId

`func (o *AzureBlobFileSyncInput) HasDataSourceId() bool`

HasDataSourceId returns a boolean if a field has been set.

### SetDataSourceIdNil

`func (o *AzureBlobFileSyncInput) SetDataSourceIdNil(b bool)`

 SetDataSourceIdNil sets the value for DataSourceId to be an explicit nil

### UnsetDataSourceId
`func (o *AzureBlobFileSyncInput) UnsetDataSourceId()`

UnsetDataSourceId ensures that no value is present for DataSourceId, not even an explicit nil
### GetRequestId

`func (o *AzureBlobFileSyncInput) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AzureBlobFileSyncInput) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AzureBlobFileSyncInput) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *AzureBlobFileSyncInput) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### SetRequestIdNil

`func (o *AzureBlobFileSyncInput) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *AzureBlobFileSyncInput) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetUseOcr

`func (o *AzureBlobFileSyncInput) GetUseOcr() bool`

GetUseOcr returns the UseOcr field if non-nil, zero value otherwise.

### GetUseOcrOk

`func (o *AzureBlobFileSyncInput) GetUseOcrOk() (*bool, bool)`

GetUseOcrOk returns a tuple with the UseOcr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOcr

`func (o *AzureBlobFileSyncInput) SetUseOcr(v bool)`

SetUseOcr sets UseOcr field to given value.

### HasUseOcr

`func (o *AzureBlobFileSyncInput) HasUseOcr() bool`

HasUseOcr returns a boolean if a field has been set.

### SetUseOcrNil

`func (o *AzureBlobFileSyncInput) SetUseOcrNil(b bool)`

 SetUseOcrNil sets the value for UseOcr to be an explicit nil

### UnsetUseOcr
`func (o *AzureBlobFileSyncInput) UnsetUseOcr()`

UnsetUseOcr ensures that no value is present for UseOcr, not even an explicit nil
### GetParsePdfTablesWithOcr

`func (o *AzureBlobFileSyncInput) GetParsePdfTablesWithOcr() bool`

GetParsePdfTablesWithOcr returns the ParsePdfTablesWithOcr field if non-nil, zero value otherwise.

### GetParsePdfTablesWithOcrOk

`func (o *AzureBlobFileSyncInput) GetParsePdfTablesWithOcrOk() (*bool, bool)`

GetParsePdfTablesWithOcrOk returns a tuple with the ParsePdfTablesWithOcr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsePdfTablesWithOcr

`func (o *AzureBlobFileSyncInput) SetParsePdfTablesWithOcr(v bool)`

SetParsePdfTablesWithOcr sets ParsePdfTablesWithOcr field to given value.

### HasParsePdfTablesWithOcr

`func (o *AzureBlobFileSyncInput) HasParsePdfTablesWithOcr() bool`

HasParsePdfTablesWithOcr returns a boolean if a field has been set.

### SetParsePdfTablesWithOcrNil

`func (o *AzureBlobFileSyncInput) SetParsePdfTablesWithOcrNil(b bool)`

 SetParsePdfTablesWithOcrNil sets the value for ParsePdfTablesWithOcr to be an explicit nil

### UnsetParsePdfTablesWithOcr
`func (o *AzureBlobFileSyncInput) UnsetParsePdfTablesWithOcr()`

UnsetParsePdfTablesWithOcr ensures that no value is present for ParsePdfTablesWithOcr, not even an explicit nil
### GetFileSyncConfig

`func (o *AzureBlobFileSyncInput) GetFileSyncConfig() FileSyncConfigNullable`

GetFileSyncConfig returns the FileSyncConfig field if non-nil, zero value otherwise.

### GetFileSyncConfigOk

`func (o *AzureBlobFileSyncInput) GetFileSyncConfigOk() (*FileSyncConfigNullable, bool)`

GetFileSyncConfigOk returns a tuple with the FileSyncConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSyncConfig

`func (o *AzureBlobFileSyncInput) SetFileSyncConfig(v FileSyncConfigNullable)`

SetFileSyncConfig sets FileSyncConfig field to given value.

### HasFileSyncConfig

`func (o *AzureBlobFileSyncInput) HasFileSyncConfig() bool`

HasFileSyncConfig returns a boolean if a field has been set.

### SetFileSyncConfigNil

`func (o *AzureBlobFileSyncInput) SetFileSyncConfigNil(b bool)`

 SetFileSyncConfigNil sets the value for FileSyncConfig to be an explicit nil

### UnsetFileSyncConfig
`func (o *AzureBlobFileSyncInput) UnsetFileSyncConfig()`

UnsetFileSyncConfig ensures that no value is present for FileSyncConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



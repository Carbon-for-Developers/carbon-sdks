# S3FileSyncInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **map[string]interface{}** |  | [optional] [default to {}]
**Ids** | [**[]S3GetFileInput**](S3GetFileInput.md) |  | 
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

## Methods

### NewS3FileSyncInput

`func NewS3FileSyncInput(ids []S3GetFileInput, ) *S3FileSyncInput`

NewS3FileSyncInput instantiates a new S3FileSyncInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3FileSyncInputWithDefaults

`func NewS3FileSyncInputWithDefaults() *S3FileSyncInput`

NewS3FileSyncInputWithDefaults instantiates a new S3FileSyncInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *S3FileSyncInput) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *S3FileSyncInput) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *S3FileSyncInput) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *S3FileSyncInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *S3FileSyncInput) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *S3FileSyncInput) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetIds

`func (o *S3FileSyncInput) GetIds() []S3GetFileInput`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *S3FileSyncInput) GetIdsOk() (*[]S3GetFileInput, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *S3FileSyncInput) SetIds(v []S3GetFileInput)`

SetIds sets Ids field to given value.


### GetChunkSize

`func (o *S3FileSyncInput) GetChunkSize() int32`

GetChunkSize returns the ChunkSize field if non-nil, zero value otherwise.

### GetChunkSizeOk

`func (o *S3FileSyncInput) GetChunkSizeOk() (*int32, bool)`

GetChunkSizeOk returns a tuple with the ChunkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkSize

`func (o *S3FileSyncInput) SetChunkSize(v int32)`

SetChunkSize sets ChunkSize field to given value.

### HasChunkSize

`func (o *S3FileSyncInput) HasChunkSize() bool`

HasChunkSize returns a boolean if a field has been set.

### SetChunkSizeNil

`func (o *S3FileSyncInput) SetChunkSizeNil(b bool)`

 SetChunkSizeNil sets the value for ChunkSize to be an explicit nil

### UnsetChunkSize
`func (o *S3FileSyncInput) UnsetChunkSize()`

UnsetChunkSize ensures that no value is present for ChunkSize, not even an explicit nil
### GetChunkOverlap

`func (o *S3FileSyncInput) GetChunkOverlap() int32`

GetChunkOverlap returns the ChunkOverlap field if non-nil, zero value otherwise.

### GetChunkOverlapOk

`func (o *S3FileSyncInput) GetChunkOverlapOk() (*int32, bool)`

GetChunkOverlapOk returns a tuple with the ChunkOverlap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkOverlap

`func (o *S3FileSyncInput) SetChunkOverlap(v int32)`

SetChunkOverlap sets ChunkOverlap field to given value.

### HasChunkOverlap

`func (o *S3FileSyncInput) HasChunkOverlap() bool`

HasChunkOverlap returns a boolean if a field has been set.

### SetChunkOverlapNil

`func (o *S3FileSyncInput) SetChunkOverlapNil(b bool)`

 SetChunkOverlapNil sets the value for ChunkOverlap to be an explicit nil

### UnsetChunkOverlap
`func (o *S3FileSyncInput) UnsetChunkOverlap()`

UnsetChunkOverlap ensures that no value is present for ChunkOverlap, not even an explicit nil
### GetSkipEmbeddingGeneration

`func (o *S3FileSyncInput) GetSkipEmbeddingGeneration() bool`

GetSkipEmbeddingGeneration returns the SkipEmbeddingGeneration field if non-nil, zero value otherwise.

### GetSkipEmbeddingGenerationOk

`func (o *S3FileSyncInput) GetSkipEmbeddingGenerationOk() (*bool, bool)`

GetSkipEmbeddingGenerationOk returns a tuple with the SkipEmbeddingGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipEmbeddingGeneration

`func (o *S3FileSyncInput) SetSkipEmbeddingGeneration(v bool)`

SetSkipEmbeddingGeneration sets SkipEmbeddingGeneration field to given value.

### HasSkipEmbeddingGeneration

`func (o *S3FileSyncInput) HasSkipEmbeddingGeneration() bool`

HasSkipEmbeddingGeneration returns a boolean if a field has been set.

### SetSkipEmbeddingGenerationNil

`func (o *S3FileSyncInput) SetSkipEmbeddingGenerationNil(b bool)`

 SetSkipEmbeddingGenerationNil sets the value for SkipEmbeddingGeneration to be an explicit nil

### UnsetSkipEmbeddingGeneration
`func (o *S3FileSyncInput) UnsetSkipEmbeddingGeneration()`

UnsetSkipEmbeddingGeneration ensures that no value is present for SkipEmbeddingGeneration, not even an explicit nil
### GetEmbeddingModel

`func (o *S3FileSyncInput) GetEmbeddingModel() EmbeddingGenerators`

GetEmbeddingModel returns the EmbeddingModel field if non-nil, zero value otherwise.

### GetEmbeddingModelOk

`func (o *S3FileSyncInput) GetEmbeddingModelOk() (*EmbeddingGenerators, bool)`

GetEmbeddingModelOk returns a tuple with the EmbeddingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingModel

`func (o *S3FileSyncInput) SetEmbeddingModel(v EmbeddingGenerators)`

SetEmbeddingModel sets EmbeddingModel field to given value.

### HasEmbeddingModel

`func (o *S3FileSyncInput) HasEmbeddingModel() bool`

HasEmbeddingModel returns a boolean if a field has been set.

### GetGenerateSparseVectors

`func (o *S3FileSyncInput) GetGenerateSparseVectors() bool`

GetGenerateSparseVectors returns the GenerateSparseVectors field if non-nil, zero value otherwise.

### GetGenerateSparseVectorsOk

`func (o *S3FileSyncInput) GetGenerateSparseVectorsOk() (*bool, bool)`

GetGenerateSparseVectorsOk returns a tuple with the GenerateSparseVectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateSparseVectors

`func (o *S3FileSyncInput) SetGenerateSparseVectors(v bool)`

SetGenerateSparseVectors sets GenerateSparseVectors field to given value.

### HasGenerateSparseVectors

`func (o *S3FileSyncInput) HasGenerateSparseVectors() bool`

HasGenerateSparseVectors returns a boolean if a field has been set.

### SetGenerateSparseVectorsNil

`func (o *S3FileSyncInput) SetGenerateSparseVectorsNil(b bool)`

 SetGenerateSparseVectorsNil sets the value for GenerateSparseVectors to be an explicit nil

### UnsetGenerateSparseVectors
`func (o *S3FileSyncInput) UnsetGenerateSparseVectors()`

UnsetGenerateSparseVectors ensures that no value is present for GenerateSparseVectors, not even an explicit nil
### GetPrependFilenameToChunks

`func (o *S3FileSyncInput) GetPrependFilenameToChunks() bool`

GetPrependFilenameToChunks returns the PrependFilenameToChunks field if non-nil, zero value otherwise.

### GetPrependFilenameToChunksOk

`func (o *S3FileSyncInput) GetPrependFilenameToChunksOk() (*bool, bool)`

GetPrependFilenameToChunksOk returns a tuple with the PrependFilenameToChunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrependFilenameToChunks

`func (o *S3FileSyncInput) SetPrependFilenameToChunks(v bool)`

SetPrependFilenameToChunks sets PrependFilenameToChunks field to given value.

### HasPrependFilenameToChunks

`func (o *S3FileSyncInput) HasPrependFilenameToChunks() bool`

HasPrependFilenameToChunks returns a boolean if a field has been set.

### SetPrependFilenameToChunksNil

`func (o *S3FileSyncInput) SetPrependFilenameToChunksNil(b bool)`

 SetPrependFilenameToChunksNil sets the value for PrependFilenameToChunks to be an explicit nil

### UnsetPrependFilenameToChunks
`func (o *S3FileSyncInput) UnsetPrependFilenameToChunks()`

UnsetPrependFilenameToChunks ensures that no value is present for PrependFilenameToChunks, not even an explicit nil
### GetMaxItemsPerChunk

`func (o *S3FileSyncInput) GetMaxItemsPerChunk() int32`

GetMaxItemsPerChunk returns the MaxItemsPerChunk field if non-nil, zero value otherwise.

### GetMaxItemsPerChunkOk

`func (o *S3FileSyncInput) GetMaxItemsPerChunkOk() (*int32, bool)`

GetMaxItemsPerChunkOk returns a tuple with the MaxItemsPerChunk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxItemsPerChunk

`func (o *S3FileSyncInput) SetMaxItemsPerChunk(v int32)`

SetMaxItemsPerChunk sets MaxItemsPerChunk field to given value.

### HasMaxItemsPerChunk

`func (o *S3FileSyncInput) HasMaxItemsPerChunk() bool`

HasMaxItemsPerChunk returns a boolean if a field has been set.

### SetMaxItemsPerChunkNil

`func (o *S3FileSyncInput) SetMaxItemsPerChunkNil(b bool)`

 SetMaxItemsPerChunkNil sets the value for MaxItemsPerChunk to be an explicit nil

### UnsetMaxItemsPerChunk
`func (o *S3FileSyncInput) UnsetMaxItemsPerChunk()`

UnsetMaxItemsPerChunk ensures that no value is present for MaxItemsPerChunk, not even an explicit nil
### GetSetPageAsBoundary

`func (o *S3FileSyncInput) GetSetPageAsBoundary() bool`

GetSetPageAsBoundary returns the SetPageAsBoundary field if non-nil, zero value otherwise.

### GetSetPageAsBoundaryOk

`func (o *S3FileSyncInput) GetSetPageAsBoundaryOk() (*bool, bool)`

GetSetPageAsBoundaryOk returns a tuple with the SetPageAsBoundary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetPageAsBoundary

`func (o *S3FileSyncInput) SetSetPageAsBoundary(v bool)`

SetSetPageAsBoundary sets SetPageAsBoundary field to given value.

### HasSetPageAsBoundary

`func (o *S3FileSyncInput) HasSetPageAsBoundary() bool`

HasSetPageAsBoundary returns a boolean if a field has been set.

### GetDataSourceId

`func (o *S3FileSyncInput) GetDataSourceId() int32`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *S3FileSyncInput) GetDataSourceIdOk() (*int32, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *S3FileSyncInput) SetDataSourceId(v int32)`

SetDataSourceId sets DataSourceId field to given value.

### HasDataSourceId

`func (o *S3FileSyncInput) HasDataSourceId() bool`

HasDataSourceId returns a boolean if a field has been set.

### SetDataSourceIdNil

`func (o *S3FileSyncInput) SetDataSourceIdNil(b bool)`

 SetDataSourceIdNil sets the value for DataSourceId to be an explicit nil

### UnsetDataSourceId
`func (o *S3FileSyncInput) UnsetDataSourceId()`

UnsetDataSourceId ensures that no value is present for DataSourceId, not even an explicit nil
### GetRequestId

`func (o *S3FileSyncInput) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *S3FileSyncInput) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *S3FileSyncInput) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *S3FileSyncInput) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### SetRequestIdNil

`func (o *S3FileSyncInput) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *S3FileSyncInput) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetUseOcr

`func (o *S3FileSyncInput) GetUseOcr() bool`

GetUseOcr returns the UseOcr field if non-nil, zero value otherwise.

### GetUseOcrOk

`func (o *S3FileSyncInput) GetUseOcrOk() (*bool, bool)`

GetUseOcrOk returns a tuple with the UseOcr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOcr

`func (o *S3FileSyncInput) SetUseOcr(v bool)`

SetUseOcr sets UseOcr field to given value.

### HasUseOcr

`func (o *S3FileSyncInput) HasUseOcr() bool`

HasUseOcr returns a boolean if a field has been set.

### SetUseOcrNil

`func (o *S3FileSyncInput) SetUseOcrNil(b bool)`

 SetUseOcrNil sets the value for UseOcr to be an explicit nil

### UnsetUseOcr
`func (o *S3FileSyncInput) UnsetUseOcr()`

UnsetUseOcr ensures that no value is present for UseOcr, not even an explicit nil
### GetParsePdfTablesWithOcr

`func (o *S3FileSyncInput) GetParsePdfTablesWithOcr() bool`

GetParsePdfTablesWithOcr returns the ParsePdfTablesWithOcr field if non-nil, zero value otherwise.

### GetParsePdfTablesWithOcrOk

`func (o *S3FileSyncInput) GetParsePdfTablesWithOcrOk() (*bool, bool)`

GetParsePdfTablesWithOcrOk returns a tuple with the ParsePdfTablesWithOcr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsePdfTablesWithOcr

`func (o *S3FileSyncInput) SetParsePdfTablesWithOcr(v bool)`

SetParsePdfTablesWithOcr sets ParsePdfTablesWithOcr field to given value.

### HasParsePdfTablesWithOcr

`func (o *S3FileSyncInput) HasParsePdfTablesWithOcr() bool`

HasParsePdfTablesWithOcr returns a boolean if a field has been set.

### SetParsePdfTablesWithOcrNil

`func (o *S3FileSyncInput) SetParsePdfTablesWithOcrNil(b bool)`

 SetParsePdfTablesWithOcrNil sets the value for ParsePdfTablesWithOcr to be an explicit nil

### UnsetParsePdfTablesWithOcr
`func (o *S3FileSyncInput) UnsetParsePdfTablesWithOcr()`

UnsetParsePdfTablesWithOcr ensures that no value is present for ParsePdfTablesWithOcr, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



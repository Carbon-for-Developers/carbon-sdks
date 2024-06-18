# SlackSyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **map[string]interface{}** |  | [optional] [default to {}]
**Filters** | [**SlackFilters**](SlackFilters.md) |  | 
**ChunkSize** | Pointer to **NullableInt32** |  | [optional] [default to 1500]
**ChunkOverlap** | Pointer to **NullableInt32** |  | [optional] [default to 20]
**SkipEmbeddingGeneration** | Pointer to **NullableBool** |  | [optional] [default to false]
**EmbeddingModel** | Pointer to [**EmbeddingGenerators**](EmbeddingGenerators.md) |  | [optional] 
**GenerateSparseVectors** | Pointer to **NullableBool** |  | [optional] [default to false]
**PrependFilenameToChunks** | Pointer to **NullableBool** |  | [optional] [default to false]
**DataSourceId** | Pointer to **NullableInt32** |  | [optional] 
**RequestId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSlackSyncRequest

`func NewSlackSyncRequest(filters SlackFilters, ) *SlackSyncRequest`

NewSlackSyncRequest instantiates a new SlackSyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlackSyncRequestWithDefaults

`func NewSlackSyncRequestWithDefaults() *SlackSyncRequest`

NewSlackSyncRequestWithDefaults instantiates a new SlackSyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *SlackSyncRequest) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SlackSyncRequest) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SlackSyncRequest) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *SlackSyncRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *SlackSyncRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *SlackSyncRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetFilters

`func (o *SlackSyncRequest) GetFilters() SlackFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *SlackSyncRequest) GetFiltersOk() (*SlackFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *SlackSyncRequest) SetFilters(v SlackFilters)`

SetFilters sets Filters field to given value.


### GetChunkSize

`func (o *SlackSyncRequest) GetChunkSize() int32`

GetChunkSize returns the ChunkSize field if non-nil, zero value otherwise.

### GetChunkSizeOk

`func (o *SlackSyncRequest) GetChunkSizeOk() (*int32, bool)`

GetChunkSizeOk returns a tuple with the ChunkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkSize

`func (o *SlackSyncRequest) SetChunkSize(v int32)`

SetChunkSize sets ChunkSize field to given value.

### HasChunkSize

`func (o *SlackSyncRequest) HasChunkSize() bool`

HasChunkSize returns a boolean if a field has been set.

### SetChunkSizeNil

`func (o *SlackSyncRequest) SetChunkSizeNil(b bool)`

 SetChunkSizeNil sets the value for ChunkSize to be an explicit nil

### UnsetChunkSize
`func (o *SlackSyncRequest) UnsetChunkSize()`

UnsetChunkSize ensures that no value is present for ChunkSize, not even an explicit nil
### GetChunkOverlap

`func (o *SlackSyncRequest) GetChunkOverlap() int32`

GetChunkOverlap returns the ChunkOverlap field if non-nil, zero value otherwise.

### GetChunkOverlapOk

`func (o *SlackSyncRequest) GetChunkOverlapOk() (*int32, bool)`

GetChunkOverlapOk returns a tuple with the ChunkOverlap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkOverlap

`func (o *SlackSyncRequest) SetChunkOverlap(v int32)`

SetChunkOverlap sets ChunkOverlap field to given value.

### HasChunkOverlap

`func (o *SlackSyncRequest) HasChunkOverlap() bool`

HasChunkOverlap returns a boolean if a field has been set.

### SetChunkOverlapNil

`func (o *SlackSyncRequest) SetChunkOverlapNil(b bool)`

 SetChunkOverlapNil sets the value for ChunkOverlap to be an explicit nil

### UnsetChunkOverlap
`func (o *SlackSyncRequest) UnsetChunkOverlap()`

UnsetChunkOverlap ensures that no value is present for ChunkOverlap, not even an explicit nil
### GetSkipEmbeddingGeneration

`func (o *SlackSyncRequest) GetSkipEmbeddingGeneration() bool`

GetSkipEmbeddingGeneration returns the SkipEmbeddingGeneration field if non-nil, zero value otherwise.

### GetSkipEmbeddingGenerationOk

`func (o *SlackSyncRequest) GetSkipEmbeddingGenerationOk() (*bool, bool)`

GetSkipEmbeddingGenerationOk returns a tuple with the SkipEmbeddingGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipEmbeddingGeneration

`func (o *SlackSyncRequest) SetSkipEmbeddingGeneration(v bool)`

SetSkipEmbeddingGeneration sets SkipEmbeddingGeneration field to given value.

### HasSkipEmbeddingGeneration

`func (o *SlackSyncRequest) HasSkipEmbeddingGeneration() bool`

HasSkipEmbeddingGeneration returns a boolean if a field has been set.

### SetSkipEmbeddingGenerationNil

`func (o *SlackSyncRequest) SetSkipEmbeddingGenerationNil(b bool)`

 SetSkipEmbeddingGenerationNil sets the value for SkipEmbeddingGeneration to be an explicit nil

### UnsetSkipEmbeddingGeneration
`func (o *SlackSyncRequest) UnsetSkipEmbeddingGeneration()`

UnsetSkipEmbeddingGeneration ensures that no value is present for SkipEmbeddingGeneration, not even an explicit nil
### GetEmbeddingModel

`func (o *SlackSyncRequest) GetEmbeddingModel() EmbeddingGenerators`

GetEmbeddingModel returns the EmbeddingModel field if non-nil, zero value otherwise.

### GetEmbeddingModelOk

`func (o *SlackSyncRequest) GetEmbeddingModelOk() (*EmbeddingGenerators, bool)`

GetEmbeddingModelOk returns a tuple with the EmbeddingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingModel

`func (o *SlackSyncRequest) SetEmbeddingModel(v EmbeddingGenerators)`

SetEmbeddingModel sets EmbeddingModel field to given value.

### HasEmbeddingModel

`func (o *SlackSyncRequest) HasEmbeddingModel() bool`

HasEmbeddingModel returns a boolean if a field has been set.

### GetGenerateSparseVectors

`func (o *SlackSyncRequest) GetGenerateSparseVectors() bool`

GetGenerateSparseVectors returns the GenerateSparseVectors field if non-nil, zero value otherwise.

### GetGenerateSparseVectorsOk

`func (o *SlackSyncRequest) GetGenerateSparseVectorsOk() (*bool, bool)`

GetGenerateSparseVectorsOk returns a tuple with the GenerateSparseVectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateSparseVectors

`func (o *SlackSyncRequest) SetGenerateSparseVectors(v bool)`

SetGenerateSparseVectors sets GenerateSparseVectors field to given value.

### HasGenerateSparseVectors

`func (o *SlackSyncRequest) HasGenerateSparseVectors() bool`

HasGenerateSparseVectors returns a boolean if a field has been set.

### SetGenerateSparseVectorsNil

`func (o *SlackSyncRequest) SetGenerateSparseVectorsNil(b bool)`

 SetGenerateSparseVectorsNil sets the value for GenerateSparseVectors to be an explicit nil

### UnsetGenerateSparseVectors
`func (o *SlackSyncRequest) UnsetGenerateSparseVectors()`

UnsetGenerateSparseVectors ensures that no value is present for GenerateSparseVectors, not even an explicit nil
### GetPrependFilenameToChunks

`func (o *SlackSyncRequest) GetPrependFilenameToChunks() bool`

GetPrependFilenameToChunks returns the PrependFilenameToChunks field if non-nil, zero value otherwise.

### GetPrependFilenameToChunksOk

`func (o *SlackSyncRequest) GetPrependFilenameToChunksOk() (*bool, bool)`

GetPrependFilenameToChunksOk returns a tuple with the PrependFilenameToChunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrependFilenameToChunks

`func (o *SlackSyncRequest) SetPrependFilenameToChunks(v bool)`

SetPrependFilenameToChunks sets PrependFilenameToChunks field to given value.

### HasPrependFilenameToChunks

`func (o *SlackSyncRequest) HasPrependFilenameToChunks() bool`

HasPrependFilenameToChunks returns a boolean if a field has been set.

### SetPrependFilenameToChunksNil

`func (o *SlackSyncRequest) SetPrependFilenameToChunksNil(b bool)`

 SetPrependFilenameToChunksNil sets the value for PrependFilenameToChunks to be an explicit nil

### UnsetPrependFilenameToChunks
`func (o *SlackSyncRequest) UnsetPrependFilenameToChunks()`

UnsetPrependFilenameToChunks ensures that no value is present for PrependFilenameToChunks, not even an explicit nil
### GetDataSourceId

`func (o *SlackSyncRequest) GetDataSourceId() int32`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *SlackSyncRequest) GetDataSourceIdOk() (*int32, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *SlackSyncRequest) SetDataSourceId(v int32)`

SetDataSourceId sets DataSourceId field to given value.

### HasDataSourceId

`func (o *SlackSyncRequest) HasDataSourceId() bool`

HasDataSourceId returns a boolean if a field has been set.

### SetDataSourceIdNil

`func (o *SlackSyncRequest) SetDataSourceIdNil(b bool)`

 SetDataSourceIdNil sets the value for DataSourceId to be an explicit nil

### UnsetDataSourceId
`func (o *SlackSyncRequest) UnsetDataSourceId()`

UnsetDataSourceId ensures that no value is present for DataSourceId, not even an explicit nil
### GetRequestId

`func (o *SlackSyncRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SlackSyncRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SlackSyncRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *SlackSyncRequest) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### SetRequestIdNil

`func (o *SlackSyncRequest) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *SlackSyncRequest) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



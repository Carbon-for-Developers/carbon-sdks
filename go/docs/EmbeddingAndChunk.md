# EmbeddingAndChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserFileId** | **int32** |  | 
**ChunkIndex** | **NullableInt32** |  | 
**SourceContent** | **string** |  | 
**Embedding** | **[]float32** |  | 
**ContentMetadata** | **map[string]interface{}** |  | 

## Methods

### NewEmbeddingAndChunk

`func NewEmbeddingAndChunk(userFileId int32, chunkIndex NullableInt32, sourceContent string, embedding []float32, contentMetadata map[string]interface{}, ) *EmbeddingAndChunk`

NewEmbeddingAndChunk instantiates a new EmbeddingAndChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddingAndChunkWithDefaults

`func NewEmbeddingAndChunkWithDefaults() *EmbeddingAndChunk`

NewEmbeddingAndChunkWithDefaults instantiates a new EmbeddingAndChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserFileId

`func (o *EmbeddingAndChunk) GetUserFileId() int32`

GetUserFileId returns the UserFileId field if non-nil, zero value otherwise.

### GetUserFileIdOk

`func (o *EmbeddingAndChunk) GetUserFileIdOk() (*int32, bool)`

GetUserFileIdOk returns a tuple with the UserFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFileId

`func (o *EmbeddingAndChunk) SetUserFileId(v int32)`

SetUserFileId sets UserFileId field to given value.


### GetChunkIndex

`func (o *EmbeddingAndChunk) GetChunkIndex() int32`

GetChunkIndex returns the ChunkIndex field if non-nil, zero value otherwise.

### GetChunkIndexOk

`func (o *EmbeddingAndChunk) GetChunkIndexOk() (*int32, bool)`

GetChunkIndexOk returns a tuple with the ChunkIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkIndex

`func (o *EmbeddingAndChunk) SetChunkIndex(v int32)`

SetChunkIndex sets ChunkIndex field to given value.


### SetChunkIndexNil

`func (o *EmbeddingAndChunk) SetChunkIndexNil(b bool)`

 SetChunkIndexNil sets the value for ChunkIndex to be an explicit nil

### UnsetChunkIndex
`func (o *EmbeddingAndChunk) UnsetChunkIndex()`

UnsetChunkIndex ensures that no value is present for ChunkIndex, not even an explicit nil
### GetSourceContent

`func (o *EmbeddingAndChunk) GetSourceContent() string`

GetSourceContent returns the SourceContent field if non-nil, zero value otherwise.

### GetSourceContentOk

`func (o *EmbeddingAndChunk) GetSourceContentOk() (*string, bool)`

GetSourceContentOk returns a tuple with the SourceContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceContent

`func (o *EmbeddingAndChunk) SetSourceContent(v string)`

SetSourceContent sets SourceContent field to given value.


### GetEmbedding

`func (o *EmbeddingAndChunk) GetEmbedding() []float32`

GetEmbedding returns the Embedding field if non-nil, zero value otherwise.

### GetEmbeddingOk

`func (o *EmbeddingAndChunk) GetEmbeddingOk() (*[]float32, bool)`

GetEmbeddingOk returns a tuple with the Embedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedding

`func (o *EmbeddingAndChunk) SetEmbedding(v []float32)`

SetEmbedding sets Embedding field to given value.


### SetEmbeddingNil

`func (o *EmbeddingAndChunk) SetEmbeddingNil(b bool)`

 SetEmbeddingNil sets the value for Embedding to be an explicit nil

### UnsetEmbedding
`func (o *EmbeddingAndChunk) UnsetEmbedding()`

UnsetEmbedding ensures that no value is present for Embedding, not even an explicit nil
### GetContentMetadata

`func (o *EmbeddingAndChunk) GetContentMetadata() map[string]interface{}`

GetContentMetadata returns the ContentMetadata field if non-nil, zero value otherwise.

### GetContentMetadataOk

`func (o *EmbeddingAndChunk) GetContentMetadataOk() (*map[string]interface{}, bool)`

GetContentMetadataOk returns a tuple with the ContentMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentMetadata

`func (o *EmbeddingAndChunk) SetContentMetadata(v map[string]interface{})`

SetContentMetadata sets ContentMetadata field to given value.


### SetContentMetadataNil

`func (o *EmbeddingAndChunk) SetContentMetadataNil(b bool)`

 SetContentMetadataNil sets the value for ContentMetadata to be an explicit nil

### UnsetContentMetadata
`func (o *EmbeddingAndChunk) UnsetContentMetadata()`

UnsetContentMetadata ensures that no value is present for ContentMetadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



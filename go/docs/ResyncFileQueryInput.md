# ResyncFileQueryInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **int32** |  | 
**ChunkSize** | Pointer to **NullableInt32** |  | [optional] 
**ChunkOverlap** | Pointer to **NullableInt32** |  | [optional] 
**ForceEmbeddingGeneration** | Pointer to **bool** |  | [optional] [default to false]
**SkipFileProcessing** | Pointer to **NullableBool** |  | [optional] [default to false]

## Methods

### NewResyncFileQueryInput

`func NewResyncFileQueryInput(fileId int32, ) *ResyncFileQueryInput`

NewResyncFileQueryInput instantiates a new ResyncFileQueryInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResyncFileQueryInputWithDefaults

`func NewResyncFileQueryInputWithDefaults() *ResyncFileQueryInput`

NewResyncFileQueryInputWithDefaults instantiates a new ResyncFileQueryInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *ResyncFileQueryInput) GetFileId() int32`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *ResyncFileQueryInput) GetFileIdOk() (*int32, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *ResyncFileQueryInput) SetFileId(v int32)`

SetFileId sets FileId field to given value.


### GetChunkSize

`func (o *ResyncFileQueryInput) GetChunkSize() int32`

GetChunkSize returns the ChunkSize field if non-nil, zero value otherwise.

### GetChunkSizeOk

`func (o *ResyncFileQueryInput) GetChunkSizeOk() (*int32, bool)`

GetChunkSizeOk returns a tuple with the ChunkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkSize

`func (o *ResyncFileQueryInput) SetChunkSize(v int32)`

SetChunkSize sets ChunkSize field to given value.

### HasChunkSize

`func (o *ResyncFileQueryInput) HasChunkSize() bool`

HasChunkSize returns a boolean if a field has been set.

### SetChunkSizeNil

`func (o *ResyncFileQueryInput) SetChunkSizeNil(b bool)`

 SetChunkSizeNil sets the value for ChunkSize to be an explicit nil

### UnsetChunkSize
`func (o *ResyncFileQueryInput) UnsetChunkSize()`

UnsetChunkSize ensures that no value is present for ChunkSize, not even an explicit nil
### GetChunkOverlap

`func (o *ResyncFileQueryInput) GetChunkOverlap() int32`

GetChunkOverlap returns the ChunkOverlap field if non-nil, zero value otherwise.

### GetChunkOverlapOk

`func (o *ResyncFileQueryInput) GetChunkOverlapOk() (*int32, bool)`

GetChunkOverlapOk returns a tuple with the ChunkOverlap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkOverlap

`func (o *ResyncFileQueryInput) SetChunkOverlap(v int32)`

SetChunkOverlap sets ChunkOverlap field to given value.

### HasChunkOverlap

`func (o *ResyncFileQueryInput) HasChunkOverlap() bool`

HasChunkOverlap returns a boolean if a field has been set.

### SetChunkOverlapNil

`func (o *ResyncFileQueryInput) SetChunkOverlapNil(b bool)`

 SetChunkOverlapNil sets the value for ChunkOverlap to be an explicit nil

### UnsetChunkOverlap
`func (o *ResyncFileQueryInput) UnsetChunkOverlap()`

UnsetChunkOverlap ensures that no value is present for ChunkOverlap, not even an explicit nil
### GetForceEmbeddingGeneration

`func (o *ResyncFileQueryInput) GetForceEmbeddingGeneration() bool`

GetForceEmbeddingGeneration returns the ForceEmbeddingGeneration field if non-nil, zero value otherwise.

### GetForceEmbeddingGenerationOk

`func (o *ResyncFileQueryInput) GetForceEmbeddingGenerationOk() (*bool, bool)`

GetForceEmbeddingGenerationOk returns a tuple with the ForceEmbeddingGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceEmbeddingGeneration

`func (o *ResyncFileQueryInput) SetForceEmbeddingGeneration(v bool)`

SetForceEmbeddingGeneration sets ForceEmbeddingGeneration field to given value.

### HasForceEmbeddingGeneration

`func (o *ResyncFileQueryInput) HasForceEmbeddingGeneration() bool`

HasForceEmbeddingGeneration returns a boolean if a field has been set.

### GetSkipFileProcessing

`func (o *ResyncFileQueryInput) GetSkipFileProcessing() bool`

GetSkipFileProcessing returns the SkipFileProcessing field if non-nil, zero value otherwise.

### GetSkipFileProcessingOk

`func (o *ResyncFileQueryInput) GetSkipFileProcessingOk() (*bool, bool)`

GetSkipFileProcessingOk returns a tuple with the SkipFileProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipFileProcessing

`func (o *ResyncFileQueryInput) SetSkipFileProcessing(v bool)`

SetSkipFileProcessing sets SkipFileProcessing field to given value.

### HasSkipFileProcessing

`func (o *ResyncFileQueryInput) HasSkipFileProcessing() bool`

HasSkipFileProcessing returns a boolean if a field has been set.

### SetSkipFileProcessingNil

`func (o *ResyncFileQueryInput) SetSkipFileProcessingNil(b bool)`

 SetSkipFileProcessingNil sets the value for SkipFileProcessing to be an explicit nil

### UnsetSkipFileProcessing
`func (o *ResyncFileQueryInput) UnsetSkipFileProcessing()`

UnsetSkipFileProcessing ensures that no value is present for SkipFileProcessing, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



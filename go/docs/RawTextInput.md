# RawTextInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | **string** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**ChunkSize** | Pointer to **NullableInt32** |  | [optional] 
**ChunkOverlap** | Pointer to **NullableInt32** |  | [optional] 
**SkipEmbeddingGeneration** | Pointer to **bool** |  | [optional] [default to false]
**OverwriteFileId** | Pointer to **NullableInt32** |  | [optional] 
**EmbeddingModel** | Pointer to [**NullableEmbeddingGeneratorsNullable**](EmbeddingGeneratorsNullable.md) |  | [optional] [default to EMBEDDINGGENERATORSNULLABLE_OPENAI]
**GenerateSparseVectors** | Pointer to **NullableBool** |  | [optional] [default to false]
**ColdStorageParams** | Pointer to [**ColdStorageProps**](ColdStorageProps.md) |  | [optional] 
**GenerateChunksOnly** | Pointer to **bool** | If this flag is enabled, the file will be chunked and stored with Carbon,         but no embeddings will be generated. This overrides the skip_embedding_generation flag. | [optional] [default to false]
**StoreFileOnly** | Pointer to **bool** | If this flag is enabled, the file will be stored with Carbon, but no processing will be done. | [optional] [default to false]

## Methods

### NewRawTextInput

`func NewRawTextInput(contents string, ) *RawTextInput`

NewRawTextInput instantiates a new RawTextInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRawTextInputWithDefaults

`func NewRawTextInputWithDefaults() *RawTextInput`

NewRawTextInputWithDefaults instantiates a new RawTextInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *RawTextInput) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *RawTextInput) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *RawTextInput) SetContents(v string)`

SetContents sets Contents field to given value.


### GetName

`func (o *RawTextInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RawTextInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RawTextInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RawTextInput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RawTextInput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RawTextInput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetChunkSize

`func (o *RawTextInput) GetChunkSize() int32`

GetChunkSize returns the ChunkSize field if non-nil, zero value otherwise.

### GetChunkSizeOk

`func (o *RawTextInput) GetChunkSizeOk() (*int32, bool)`

GetChunkSizeOk returns a tuple with the ChunkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkSize

`func (o *RawTextInput) SetChunkSize(v int32)`

SetChunkSize sets ChunkSize field to given value.

### HasChunkSize

`func (o *RawTextInput) HasChunkSize() bool`

HasChunkSize returns a boolean if a field has been set.

### SetChunkSizeNil

`func (o *RawTextInput) SetChunkSizeNil(b bool)`

 SetChunkSizeNil sets the value for ChunkSize to be an explicit nil

### UnsetChunkSize
`func (o *RawTextInput) UnsetChunkSize()`

UnsetChunkSize ensures that no value is present for ChunkSize, not even an explicit nil
### GetChunkOverlap

`func (o *RawTextInput) GetChunkOverlap() int32`

GetChunkOverlap returns the ChunkOverlap field if non-nil, zero value otherwise.

### GetChunkOverlapOk

`func (o *RawTextInput) GetChunkOverlapOk() (*int32, bool)`

GetChunkOverlapOk returns a tuple with the ChunkOverlap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkOverlap

`func (o *RawTextInput) SetChunkOverlap(v int32)`

SetChunkOverlap sets ChunkOverlap field to given value.

### HasChunkOverlap

`func (o *RawTextInput) HasChunkOverlap() bool`

HasChunkOverlap returns a boolean if a field has been set.

### SetChunkOverlapNil

`func (o *RawTextInput) SetChunkOverlapNil(b bool)`

 SetChunkOverlapNil sets the value for ChunkOverlap to be an explicit nil

### UnsetChunkOverlap
`func (o *RawTextInput) UnsetChunkOverlap()`

UnsetChunkOverlap ensures that no value is present for ChunkOverlap, not even an explicit nil
### GetSkipEmbeddingGeneration

`func (o *RawTextInput) GetSkipEmbeddingGeneration() bool`

GetSkipEmbeddingGeneration returns the SkipEmbeddingGeneration field if non-nil, zero value otherwise.

### GetSkipEmbeddingGenerationOk

`func (o *RawTextInput) GetSkipEmbeddingGenerationOk() (*bool, bool)`

GetSkipEmbeddingGenerationOk returns a tuple with the SkipEmbeddingGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipEmbeddingGeneration

`func (o *RawTextInput) SetSkipEmbeddingGeneration(v bool)`

SetSkipEmbeddingGeneration sets SkipEmbeddingGeneration field to given value.

### HasSkipEmbeddingGeneration

`func (o *RawTextInput) HasSkipEmbeddingGeneration() bool`

HasSkipEmbeddingGeneration returns a boolean if a field has been set.

### GetOverwriteFileId

`func (o *RawTextInput) GetOverwriteFileId() int32`

GetOverwriteFileId returns the OverwriteFileId field if non-nil, zero value otherwise.

### GetOverwriteFileIdOk

`func (o *RawTextInput) GetOverwriteFileIdOk() (*int32, bool)`

GetOverwriteFileIdOk returns a tuple with the OverwriteFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteFileId

`func (o *RawTextInput) SetOverwriteFileId(v int32)`

SetOverwriteFileId sets OverwriteFileId field to given value.

### HasOverwriteFileId

`func (o *RawTextInput) HasOverwriteFileId() bool`

HasOverwriteFileId returns a boolean if a field has been set.

### SetOverwriteFileIdNil

`func (o *RawTextInput) SetOverwriteFileIdNil(b bool)`

 SetOverwriteFileIdNil sets the value for OverwriteFileId to be an explicit nil

### UnsetOverwriteFileId
`func (o *RawTextInput) UnsetOverwriteFileId()`

UnsetOverwriteFileId ensures that no value is present for OverwriteFileId, not even an explicit nil
### GetEmbeddingModel

`func (o *RawTextInput) GetEmbeddingModel() EmbeddingGeneratorsNullable`

GetEmbeddingModel returns the EmbeddingModel field if non-nil, zero value otherwise.

### GetEmbeddingModelOk

`func (o *RawTextInput) GetEmbeddingModelOk() (*EmbeddingGeneratorsNullable, bool)`

GetEmbeddingModelOk returns a tuple with the EmbeddingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingModel

`func (o *RawTextInput) SetEmbeddingModel(v EmbeddingGeneratorsNullable)`

SetEmbeddingModel sets EmbeddingModel field to given value.

### HasEmbeddingModel

`func (o *RawTextInput) HasEmbeddingModel() bool`

HasEmbeddingModel returns a boolean if a field has been set.

### SetEmbeddingModelNil

`func (o *RawTextInput) SetEmbeddingModelNil(b bool)`

 SetEmbeddingModelNil sets the value for EmbeddingModel to be an explicit nil

### UnsetEmbeddingModel
`func (o *RawTextInput) UnsetEmbeddingModel()`

UnsetEmbeddingModel ensures that no value is present for EmbeddingModel, not even an explicit nil
### GetGenerateSparseVectors

`func (o *RawTextInput) GetGenerateSparseVectors() bool`

GetGenerateSparseVectors returns the GenerateSparseVectors field if non-nil, zero value otherwise.

### GetGenerateSparseVectorsOk

`func (o *RawTextInput) GetGenerateSparseVectorsOk() (*bool, bool)`

GetGenerateSparseVectorsOk returns a tuple with the GenerateSparseVectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateSparseVectors

`func (o *RawTextInput) SetGenerateSparseVectors(v bool)`

SetGenerateSparseVectors sets GenerateSparseVectors field to given value.

### HasGenerateSparseVectors

`func (o *RawTextInput) HasGenerateSparseVectors() bool`

HasGenerateSparseVectors returns a boolean if a field has been set.

### SetGenerateSparseVectorsNil

`func (o *RawTextInput) SetGenerateSparseVectorsNil(b bool)`

 SetGenerateSparseVectorsNil sets the value for GenerateSparseVectors to be an explicit nil

### UnsetGenerateSparseVectors
`func (o *RawTextInput) UnsetGenerateSparseVectors()`

UnsetGenerateSparseVectors ensures that no value is present for GenerateSparseVectors, not even an explicit nil
### GetColdStorageParams

`func (o *RawTextInput) GetColdStorageParams() ColdStorageProps`

GetColdStorageParams returns the ColdStorageParams field if non-nil, zero value otherwise.

### GetColdStorageParamsOk

`func (o *RawTextInput) GetColdStorageParamsOk() (*ColdStorageProps, bool)`

GetColdStorageParamsOk returns a tuple with the ColdStorageParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColdStorageParams

`func (o *RawTextInput) SetColdStorageParams(v ColdStorageProps)`

SetColdStorageParams sets ColdStorageParams field to given value.

### HasColdStorageParams

`func (o *RawTextInput) HasColdStorageParams() bool`

HasColdStorageParams returns a boolean if a field has been set.

### GetGenerateChunksOnly

`func (o *RawTextInput) GetGenerateChunksOnly() bool`

GetGenerateChunksOnly returns the GenerateChunksOnly field if non-nil, zero value otherwise.

### GetGenerateChunksOnlyOk

`func (o *RawTextInput) GetGenerateChunksOnlyOk() (*bool, bool)`

GetGenerateChunksOnlyOk returns a tuple with the GenerateChunksOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateChunksOnly

`func (o *RawTextInput) SetGenerateChunksOnly(v bool)`

SetGenerateChunksOnly sets GenerateChunksOnly field to given value.

### HasGenerateChunksOnly

`func (o *RawTextInput) HasGenerateChunksOnly() bool`

HasGenerateChunksOnly returns a boolean if a field has been set.

### GetStoreFileOnly

`func (o *RawTextInput) GetStoreFileOnly() bool`

GetStoreFileOnly returns the StoreFileOnly field if non-nil, zero value otherwise.

### GetStoreFileOnlyOk

`func (o *RawTextInput) GetStoreFileOnlyOk() (*bool, bool)`

GetStoreFileOnlyOk returns a tuple with the StoreFileOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreFileOnly

`func (o *RawTextInput) SetStoreFileOnly(v bool)`

SetStoreFileOnly sets StoreFileOnly field to given value.

### HasStoreFileOnly

`func (o *RawTextInput) HasStoreFileOnly() bool`

HasStoreFileOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



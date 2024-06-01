# UploadFileFromUrlInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**FileName** | Pointer to **NullableString** |  | [optional] 
**ChunkSize** | Pointer to **NullableInt32** |  | [optional] 
**ChunkOverlap** | Pointer to **NullableInt32** |  | [optional] 
**SkipEmbeddingGeneration** | Pointer to **bool** |  | [optional] [default to false]
**SetPageAsBoundary** | Pointer to **bool** |  | [optional] [default to false]
**EmbeddingModel** | Pointer to [**EmbeddingGenerators**](EmbeddingGenerators.md) |  | [optional] 
**GenerateSparseVectors** | Pointer to **bool** |  | [optional] [default to false]
**UseTextract** | Pointer to **bool** |  | [optional] [default to false]
**PrependFilenameToChunks** | Pointer to **bool** |  | [optional] [default to false]
**MaxItemsPerChunk** | Pointer to **NullableInt32** | Number of objects per chunk. For csv, tsv, xlsx, and json files only. | [optional] 
**ParsePdfTablesWithOcr** | Pointer to **bool** |  | [optional] [default to false]
**DetectAudioLanguage** | Pointer to **bool** |  | [optional] [default to false]
**MediaType** | Pointer to [**NullableFileContentTypesNullable**](FileContentTypesNullable.md) |  | [optional] 

## Methods

### NewUploadFileFromUrlInput

`func NewUploadFileFromUrlInput(url string, ) *UploadFileFromUrlInput`

NewUploadFileFromUrlInput instantiates a new UploadFileFromUrlInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadFileFromUrlInputWithDefaults

`func NewUploadFileFromUrlInputWithDefaults() *UploadFileFromUrlInput`

NewUploadFileFromUrlInputWithDefaults instantiates a new UploadFileFromUrlInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *UploadFileFromUrlInput) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UploadFileFromUrlInput) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UploadFileFromUrlInput) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetFileName

`func (o *UploadFileFromUrlInput) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *UploadFileFromUrlInput) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *UploadFileFromUrlInput) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *UploadFileFromUrlInput) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *UploadFileFromUrlInput) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *UploadFileFromUrlInput) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetChunkSize

`func (o *UploadFileFromUrlInput) GetChunkSize() int32`

GetChunkSize returns the ChunkSize field if non-nil, zero value otherwise.

### GetChunkSizeOk

`func (o *UploadFileFromUrlInput) GetChunkSizeOk() (*int32, bool)`

GetChunkSizeOk returns a tuple with the ChunkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkSize

`func (o *UploadFileFromUrlInput) SetChunkSize(v int32)`

SetChunkSize sets ChunkSize field to given value.

### HasChunkSize

`func (o *UploadFileFromUrlInput) HasChunkSize() bool`

HasChunkSize returns a boolean if a field has been set.

### SetChunkSizeNil

`func (o *UploadFileFromUrlInput) SetChunkSizeNil(b bool)`

 SetChunkSizeNil sets the value for ChunkSize to be an explicit nil

### UnsetChunkSize
`func (o *UploadFileFromUrlInput) UnsetChunkSize()`

UnsetChunkSize ensures that no value is present for ChunkSize, not even an explicit nil
### GetChunkOverlap

`func (o *UploadFileFromUrlInput) GetChunkOverlap() int32`

GetChunkOverlap returns the ChunkOverlap field if non-nil, zero value otherwise.

### GetChunkOverlapOk

`func (o *UploadFileFromUrlInput) GetChunkOverlapOk() (*int32, bool)`

GetChunkOverlapOk returns a tuple with the ChunkOverlap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkOverlap

`func (o *UploadFileFromUrlInput) SetChunkOverlap(v int32)`

SetChunkOverlap sets ChunkOverlap field to given value.

### HasChunkOverlap

`func (o *UploadFileFromUrlInput) HasChunkOverlap() bool`

HasChunkOverlap returns a boolean if a field has been set.

### SetChunkOverlapNil

`func (o *UploadFileFromUrlInput) SetChunkOverlapNil(b bool)`

 SetChunkOverlapNil sets the value for ChunkOverlap to be an explicit nil

### UnsetChunkOverlap
`func (o *UploadFileFromUrlInput) UnsetChunkOverlap()`

UnsetChunkOverlap ensures that no value is present for ChunkOverlap, not even an explicit nil
### GetSkipEmbeddingGeneration

`func (o *UploadFileFromUrlInput) GetSkipEmbeddingGeneration() bool`

GetSkipEmbeddingGeneration returns the SkipEmbeddingGeneration field if non-nil, zero value otherwise.

### GetSkipEmbeddingGenerationOk

`func (o *UploadFileFromUrlInput) GetSkipEmbeddingGenerationOk() (*bool, bool)`

GetSkipEmbeddingGenerationOk returns a tuple with the SkipEmbeddingGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipEmbeddingGeneration

`func (o *UploadFileFromUrlInput) SetSkipEmbeddingGeneration(v bool)`

SetSkipEmbeddingGeneration sets SkipEmbeddingGeneration field to given value.

### HasSkipEmbeddingGeneration

`func (o *UploadFileFromUrlInput) HasSkipEmbeddingGeneration() bool`

HasSkipEmbeddingGeneration returns a boolean if a field has been set.

### GetSetPageAsBoundary

`func (o *UploadFileFromUrlInput) GetSetPageAsBoundary() bool`

GetSetPageAsBoundary returns the SetPageAsBoundary field if non-nil, zero value otherwise.

### GetSetPageAsBoundaryOk

`func (o *UploadFileFromUrlInput) GetSetPageAsBoundaryOk() (*bool, bool)`

GetSetPageAsBoundaryOk returns a tuple with the SetPageAsBoundary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetPageAsBoundary

`func (o *UploadFileFromUrlInput) SetSetPageAsBoundary(v bool)`

SetSetPageAsBoundary sets SetPageAsBoundary field to given value.

### HasSetPageAsBoundary

`func (o *UploadFileFromUrlInput) HasSetPageAsBoundary() bool`

HasSetPageAsBoundary returns a boolean if a field has been set.

### GetEmbeddingModel

`func (o *UploadFileFromUrlInput) GetEmbeddingModel() EmbeddingGenerators`

GetEmbeddingModel returns the EmbeddingModel field if non-nil, zero value otherwise.

### GetEmbeddingModelOk

`func (o *UploadFileFromUrlInput) GetEmbeddingModelOk() (*EmbeddingGenerators, bool)`

GetEmbeddingModelOk returns a tuple with the EmbeddingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingModel

`func (o *UploadFileFromUrlInput) SetEmbeddingModel(v EmbeddingGenerators)`

SetEmbeddingModel sets EmbeddingModel field to given value.

### HasEmbeddingModel

`func (o *UploadFileFromUrlInput) HasEmbeddingModel() bool`

HasEmbeddingModel returns a boolean if a field has been set.

### GetGenerateSparseVectors

`func (o *UploadFileFromUrlInput) GetGenerateSparseVectors() bool`

GetGenerateSparseVectors returns the GenerateSparseVectors field if non-nil, zero value otherwise.

### GetGenerateSparseVectorsOk

`func (o *UploadFileFromUrlInput) GetGenerateSparseVectorsOk() (*bool, bool)`

GetGenerateSparseVectorsOk returns a tuple with the GenerateSparseVectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateSparseVectors

`func (o *UploadFileFromUrlInput) SetGenerateSparseVectors(v bool)`

SetGenerateSparseVectors sets GenerateSparseVectors field to given value.

### HasGenerateSparseVectors

`func (o *UploadFileFromUrlInput) HasGenerateSparseVectors() bool`

HasGenerateSparseVectors returns a boolean if a field has been set.

### GetUseTextract

`func (o *UploadFileFromUrlInput) GetUseTextract() bool`

GetUseTextract returns the UseTextract field if non-nil, zero value otherwise.

### GetUseTextractOk

`func (o *UploadFileFromUrlInput) GetUseTextractOk() (*bool, bool)`

GetUseTextractOk returns a tuple with the UseTextract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTextract

`func (o *UploadFileFromUrlInput) SetUseTextract(v bool)`

SetUseTextract sets UseTextract field to given value.

### HasUseTextract

`func (o *UploadFileFromUrlInput) HasUseTextract() bool`

HasUseTextract returns a boolean if a field has been set.

### GetPrependFilenameToChunks

`func (o *UploadFileFromUrlInput) GetPrependFilenameToChunks() bool`

GetPrependFilenameToChunks returns the PrependFilenameToChunks field if non-nil, zero value otherwise.

### GetPrependFilenameToChunksOk

`func (o *UploadFileFromUrlInput) GetPrependFilenameToChunksOk() (*bool, bool)`

GetPrependFilenameToChunksOk returns a tuple with the PrependFilenameToChunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrependFilenameToChunks

`func (o *UploadFileFromUrlInput) SetPrependFilenameToChunks(v bool)`

SetPrependFilenameToChunks sets PrependFilenameToChunks field to given value.

### HasPrependFilenameToChunks

`func (o *UploadFileFromUrlInput) HasPrependFilenameToChunks() bool`

HasPrependFilenameToChunks returns a boolean if a field has been set.

### GetMaxItemsPerChunk

`func (o *UploadFileFromUrlInput) GetMaxItemsPerChunk() int32`

GetMaxItemsPerChunk returns the MaxItemsPerChunk field if non-nil, zero value otherwise.

### GetMaxItemsPerChunkOk

`func (o *UploadFileFromUrlInput) GetMaxItemsPerChunkOk() (*int32, bool)`

GetMaxItemsPerChunkOk returns a tuple with the MaxItemsPerChunk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxItemsPerChunk

`func (o *UploadFileFromUrlInput) SetMaxItemsPerChunk(v int32)`

SetMaxItemsPerChunk sets MaxItemsPerChunk field to given value.

### HasMaxItemsPerChunk

`func (o *UploadFileFromUrlInput) HasMaxItemsPerChunk() bool`

HasMaxItemsPerChunk returns a boolean if a field has been set.

### SetMaxItemsPerChunkNil

`func (o *UploadFileFromUrlInput) SetMaxItemsPerChunkNil(b bool)`

 SetMaxItemsPerChunkNil sets the value for MaxItemsPerChunk to be an explicit nil

### UnsetMaxItemsPerChunk
`func (o *UploadFileFromUrlInput) UnsetMaxItemsPerChunk()`

UnsetMaxItemsPerChunk ensures that no value is present for MaxItemsPerChunk, not even an explicit nil
### GetParsePdfTablesWithOcr

`func (o *UploadFileFromUrlInput) GetParsePdfTablesWithOcr() bool`

GetParsePdfTablesWithOcr returns the ParsePdfTablesWithOcr field if non-nil, zero value otherwise.

### GetParsePdfTablesWithOcrOk

`func (o *UploadFileFromUrlInput) GetParsePdfTablesWithOcrOk() (*bool, bool)`

GetParsePdfTablesWithOcrOk returns a tuple with the ParsePdfTablesWithOcr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsePdfTablesWithOcr

`func (o *UploadFileFromUrlInput) SetParsePdfTablesWithOcr(v bool)`

SetParsePdfTablesWithOcr sets ParsePdfTablesWithOcr field to given value.

### HasParsePdfTablesWithOcr

`func (o *UploadFileFromUrlInput) HasParsePdfTablesWithOcr() bool`

HasParsePdfTablesWithOcr returns a boolean if a field has been set.

### GetDetectAudioLanguage

`func (o *UploadFileFromUrlInput) GetDetectAudioLanguage() bool`

GetDetectAudioLanguage returns the DetectAudioLanguage field if non-nil, zero value otherwise.

### GetDetectAudioLanguageOk

`func (o *UploadFileFromUrlInput) GetDetectAudioLanguageOk() (*bool, bool)`

GetDetectAudioLanguageOk returns a tuple with the DetectAudioLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectAudioLanguage

`func (o *UploadFileFromUrlInput) SetDetectAudioLanguage(v bool)`

SetDetectAudioLanguage sets DetectAudioLanguage field to given value.

### HasDetectAudioLanguage

`func (o *UploadFileFromUrlInput) HasDetectAudioLanguage() bool`

HasDetectAudioLanguage returns a boolean if a field has been set.

### GetMediaType

`func (o *UploadFileFromUrlInput) GetMediaType() FileContentTypesNullable`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *UploadFileFromUrlInput) GetMediaTypeOk() (*FileContentTypesNullable, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *UploadFileFromUrlInput) SetMediaType(v FileContentTypesNullable)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *UploadFileFromUrlInput) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### SetMediaTypeNil

`func (o *UploadFileFromUrlInput) SetMediaTypeNil(b bool)`

 SetMediaTypeNil sets the value for MediaType to be an explicit nil

### UnsetMediaType
`func (o *UploadFileFromUrlInput) UnsetMediaType()`

UnsetMediaType ensures that no value is present for MediaType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



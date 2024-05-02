# ChunksAndEmbeddingsUploadInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmbeddingModel** | [**EmbeddingGenerators**](EmbeddingGenerators.md) |  | 
**ChunksAndEmbeddings** | [**[]SingleChunksAndEmbeddingsUploadInput**](SingleChunksAndEmbeddingsUploadInput.md) |  | 
**OverwriteExisting** | Pointer to **bool** |  | [optional] [default to false]
**ChunksOnly** | Pointer to **bool** |  | [optional] [default to false]
**CustomCredentials** | Pointer to **map[string]map[string]interface{}** |  | [optional] [default to {}]

## Methods

### NewChunksAndEmbeddingsUploadInput

`func NewChunksAndEmbeddingsUploadInput(embeddingModel EmbeddingGenerators, chunksAndEmbeddings []SingleChunksAndEmbeddingsUploadInput, ) *ChunksAndEmbeddingsUploadInput`

NewChunksAndEmbeddingsUploadInput instantiates a new ChunksAndEmbeddingsUploadInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChunksAndEmbeddingsUploadInputWithDefaults

`func NewChunksAndEmbeddingsUploadInputWithDefaults() *ChunksAndEmbeddingsUploadInput`

NewChunksAndEmbeddingsUploadInputWithDefaults instantiates a new ChunksAndEmbeddingsUploadInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbeddingModel

`func (o *ChunksAndEmbeddingsUploadInput) GetEmbeddingModel() EmbeddingGenerators`

GetEmbeddingModel returns the EmbeddingModel field if non-nil, zero value otherwise.

### GetEmbeddingModelOk

`func (o *ChunksAndEmbeddingsUploadInput) GetEmbeddingModelOk() (*EmbeddingGenerators, bool)`

GetEmbeddingModelOk returns a tuple with the EmbeddingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingModel

`func (o *ChunksAndEmbeddingsUploadInput) SetEmbeddingModel(v EmbeddingGenerators)`

SetEmbeddingModel sets EmbeddingModel field to given value.


### GetChunksAndEmbeddings

`func (o *ChunksAndEmbeddingsUploadInput) GetChunksAndEmbeddings() []SingleChunksAndEmbeddingsUploadInput`

GetChunksAndEmbeddings returns the ChunksAndEmbeddings field if non-nil, zero value otherwise.

### GetChunksAndEmbeddingsOk

`func (o *ChunksAndEmbeddingsUploadInput) GetChunksAndEmbeddingsOk() (*[]SingleChunksAndEmbeddingsUploadInput, bool)`

GetChunksAndEmbeddingsOk returns a tuple with the ChunksAndEmbeddings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunksAndEmbeddings

`func (o *ChunksAndEmbeddingsUploadInput) SetChunksAndEmbeddings(v []SingleChunksAndEmbeddingsUploadInput)`

SetChunksAndEmbeddings sets ChunksAndEmbeddings field to given value.


### GetOverwriteExisting

`func (o *ChunksAndEmbeddingsUploadInput) GetOverwriteExisting() bool`

GetOverwriteExisting returns the OverwriteExisting field if non-nil, zero value otherwise.

### GetOverwriteExistingOk

`func (o *ChunksAndEmbeddingsUploadInput) GetOverwriteExistingOk() (*bool, bool)`

GetOverwriteExistingOk returns a tuple with the OverwriteExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteExisting

`func (o *ChunksAndEmbeddingsUploadInput) SetOverwriteExisting(v bool)`

SetOverwriteExisting sets OverwriteExisting field to given value.

### HasOverwriteExisting

`func (o *ChunksAndEmbeddingsUploadInput) HasOverwriteExisting() bool`

HasOverwriteExisting returns a boolean if a field has been set.

### GetChunksOnly

`func (o *ChunksAndEmbeddingsUploadInput) GetChunksOnly() bool`

GetChunksOnly returns the ChunksOnly field if non-nil, zero value otherwise.

### GetChunksOnlyOk

`func (o *ChunksAndEmbeddingsUploadInput) GetChunksOnlyOk() (*bool, bool)`

GetChunksOnlyOk returns a tuple with the ChunksOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunksOnly

`func (o *ChunksAndEmbeddingsUploadInput) SetChunksOnly(v bool)`

SetChunksOnly sets ChunksOnly field to given value.

### HasChunksOnly

`func (o *ChunksAndEmbeddingsUploadInput) HasChunksOnly() bool`

HasChunksOnly returns a boolean if a field has been set.

### GetCustomCredentials

`func (o *ChunksAndEmbeddingsUploadInput) GetCustomCredentials() map[string]map[string]interface{}`

GetCustomCredentials returns the CustomCredentials field if non-nil, zero value otherwise.

### GetCustomCredentialsOk

`func (o *ChunksAndEmbeddingsUploadInput) GetCustomCredentialsOk() (*map[string]map[string]interface{}, bool)`

GetCustomCredentialsOk returns a tuple with the CustomCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCredentials

`func (o *ChunksAndEmbeddingsUploadInput) SetCustomCredentials(v map[string]map[string]interface{})`

SetCustomCredentials sets CustomCredentials field to given value.

### HasCustomCredentials

`func (o *ChunksAndEmbeddingsUploadInput) HasCustomCredentials() bool`

HasCustomCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



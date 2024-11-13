# FileStatisticsNullable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileFormat** | Pointer to [**NullableFileFormatsNullable**](FileFormatsNullable.md) |  | [optional] 
**FileSize** | Pointer to **NullableInt32** |  | [optional] 
**NumCharacters** | Pointer to **NullableInt32** |  | [optional] 
**NumTokens** | Pointer to **NullableInt32** |  | [optional] 
**NumEmbeddings** | Pointer to **NullableInt32** |  | [optional] 
**MimeType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFileStatisticsNullable

`func NewFileStatisticsNullable() *FileStatisticsNullable`

NewFileStatisticsNullable instantiates a new FileStatisticsNullable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileStatisticsNullableWithDefaults

`func NewFileStatisticsNullableWithDefaults() *FileStatisticsNullable`

NewFileStatisticsNullableWithDefaults instantiates a new FileStatisticsNullable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileFormat

`func (o *FileStatisticsNullable) GetFileFormat() FileFormatsNullable`

GetFileFormat returns the FileFormat field if non-nil, zero value otherwise.

### GetFileFormatOk

`func (o *FileStatisticsNullable) GetFileFormatOk() (*FileFormatsNullable, bool)`

GetFileFormatOk returns a tuple with the FileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFormat

`func (o *FileStatisticsNullable) SetFileFormat(v FileFormatsNullable)`

SetFileFormat sets FileFormat field to given value.

### HasFileFormat

`func (o *FileStatisticsNullable) HasFileFormat() bool`

HasFileFormat returns a boolean if a field has been set.

### SetFileFormatNil

`func (o *FileStatisticsNullable) SetFileFormatNil(b bool)`

 SetFileFormatNil sets the value for FileFormat to be an explicit nil

### UnsetFileFormat
`func (o *FileStatisticsNullable) UnsetFileFormat()`

UnsetFileFormat ensures that no value is present for FileFormat, not even an explicit nil
### GetFileSize

`func (o *FileStatisticsNullable) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *FileStatisticsNullable) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *FileStatisticsNullable) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *FileStatisticsNullable) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### SetFileSizeNil

`func (o *FileStatisticsNullable) SetFileSizeNil(b bool)`

 SetFileSizeNil sets the value for FileSize to be an explicit nil

### UnsetFileSize
`func (o *FileStatisticsNullable) UnsetFileSize()`

UnsetFileSize ensures that no value is present for FileSize, not even an explicit nil
### GetNumCharacters

`func (o *FileStatisticsNullable) GetNumCharacters() int32`

GetNumCharacters returns the NumCharacters field if non-nil, zero value otherwise.

### GetNumCharactersOk

`func (o *FileStatisticsNullable) GetNumCharactersOk() (*int32, bool)`

GetNumCharactersOk returns a tuple with the NumCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCharacters

`func (o *FileStatisticsNullable) SetNumCharacters(v int32)`

SetNumCharacters sets NumCharacters field to given value.

### HasNumCharacters

`func (o *FileStatisticsNullable) HasNumCharacters() bool`

HasNumCharacters returns a boolean if a field has been set.

### SetNumCharactersNil

`func (o *FileStatisticsNullable) SetNumCharactersNil(b bool)`

 SetNumCharactersNil sets the value for NumCharacters to be an explicit nil

### UnsetNumCharacters
`func (o *FileStatisticsNullable) UnsetNumCharacters()`

UnsetNumCharacters ensures that no value is present for NumCharacters, not even an explicit nil
### GetNumTokens

`func (o *FileStatisticsNullable) GetNumTokens() int32`

GetNumTokens returns the NumTokens field if non-nil, zero value otherwise.

### GetNumTokensOk

`func (o *FileStatisticsNullable) GetNumTokensOk() (*int32, bool)`

GetNumTokensOk returns a tuple with the NumTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTokens

`func (o *FileStatisticsNullable) SetNumTokens(v int32)`

SetNumTokens sets NumTokens field to given value.

### HasNumTokens

`func (o *FileStatisticsNullable) HasNumTokens() bool`

HasNumTokens returns a boolean if a field has been set.

### SetNumTokensNil

`func (o *FileStatisticsNullable) SetNumTokensNil(b bool)`

 SetNumTokensNil sets the value for NumTokens to be an explicit nil

### UnsetNumTokens
`func (o *FileStatisticsNullable) UnsetNumTokens()`

UnsetNumTokens ensures that no value is present for NumTokens, not even an explicit nil
### GetNumEmbeddings

`func (o *FileStatisticsNullable) GetNumEmbeddings() int32`

GetNumEmbeddings returns the NumEmbeddings field if non-nil, zero value otherwise.

### GetNumEmbeddingsOk

`func (o *FileStatisticsNullable) GetNumEmbeddingsOk() (*int32, bool)`

GetNumEmbeddingsOk returns a tuple with the NumEmbeddings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumEmbeddings

`func (o *FileStatisticsNullable) SetNumEmbeddings(v int32)`

SetNumEmbeddings sets NumEmbeddings field to given value.

### HasNumEmbeddings

`func (o *FileStatisticsNullable) HasNumEmbeddings() bool`

HasNumEmbeddings returns a boolean if a field has been set.

### SetNumEmbeddingsNil

`func (o *FileStatisticsNullable) SetNumEmbeddingsNil(b bool)`

 SetNumEmbeddingsNil sets the value for NumEmbeddings to be an explicit nil

### UnsetNumEmbeddings
`func (o *FileStatisticsNullable) UnsetNumEmbeddings()`

UnsetNumEmbeddings ensures that no value is present for NumEmbeddings, not even an explicit nil
### GetMimeType

`func (o *FileStatisticsNullable) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *FileStatisticsNullable) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *FileStatisticsNullable) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *FileStatisticsNullable) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### SetMimeTypeNil

`func (o *FileStatisticsNullable) SetMimeTypeNil(b bool)`

 SetMimeTypeNil sets the value for MimeType to be an explicit nil

### UnsetMimeType
`func (o *FileStatisticsNullable) UnsetMimeType()`

UnsetMimeType ensures that no value is present for MimeType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



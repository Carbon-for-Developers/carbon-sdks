# UserConfigurationNullable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSyncEnabledSources** | Pointer to [**NullableAutoSyncEnabledSourcesProperty**](AutoSyncEnabledSourcesProperty.md) |  | [optional] 
**MaxFiles** | Pointer to **NullableInt32** | Custom file upload limit for the user over *all* user&#39;s files across all uploads.          If set, then the user will not be allowed to upload more files than this limit. If not set, or if set to -1,         then the user will have no limit. | [optional] 
**MaxFilesPerUpload** | Pointer to **NullableInt32** | Custom file upload limit for the user across a single upload.         If set, then the user will not be allowed to upload more files than this limit in a single upload. If not set,         or if set to -1, then the user will have no limit. | [optional] 
**MaxCharacters** | Pointer to **NullableInt32** | Custom character upload limit for the user over *all* user&#39;s files across all uploads.          If set, then the user will not be allowed to upload more characters than this limit. If not set, or if set to -1,         then the user will have no limit. | [optional] 
**MaxCharactersPerFile** | Pointer to **NullableInt32** | A single file upload from the user can not exceed this character limit.         If set, then the file will not be synced if it exceeds this limit. If not set, or if set to -1, then the          user will have no limit. | [optional] 
**MaxCharactersPerUpload** | Pointer to **NullableInt32** | Custom character upload limit for the user across a single upload.         If set, then the user won&#39;t be able to sync more than this many characters in one upload.          If not set, or if set to -1, then the user will have no limit. | [optional] 
**AutoSyncInterval** | Pointer to **NullableInt32** | The interval in hours at which the user&#39;s data sources should be synced. If not set or set to -1,          the user will be synced at the organization level interval or default interval if that is also not set.          Must be one of [3, 6, 12, 24] | [optional] 

## Methods

### NewUserConfigurationNullable

`func NewUserConfigurationNullable() *UserConfigurationNullable`

NewUserConfigurationNullable instantiates a new UserConfigurationNullable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserConfigurationNullableWithDefaults

`func NewUserConfigurationNullableWithDefaults() *UserConfigurationNullable`

NewUserConfigurationNullableWithDefaults instantiates a new UserConfigurationNullable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoSyncEnabledSources

`func (o *UserConfigurationNullable) GetAutoSyncEnabledSources() AutoSyncEnabledSourcesProperty`

GetAutoSyncEnabledSources returns the AutoSyncEnabledSources field if non-nil, zero value otherwise.

### GetAutoSyncEnabledSourcesOk

`func (o *UserConfigurationNullable) GetAutoSyncEnabledSourcesOk() (*AutoSyncEnabledSourcesProperty, bool)`

GetAutoSyncEnabledSourcesOk returns a tuple with the AutoSyncEnabledSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncEnabledSources

`func (o *UserConfigurationNullable) SetAutoSyncEnabledSources(v AutoSyncEnabledSourcesProperty)`

SetAutoSyncEnabledSources sets AutoSyncEnabledSources field to given value.

### HasAutoSyncEnabledSources

`func (o *UserConfigurationNullable) HasAutoSyncEnabledSources() bool`

HasAutoSyncEnabledSources returns a boolean if a field has been set.

### SetAutoSyncEnabledSourcesNil

`func (o *UserConfigurationNullable) SetAutoSyncEnabledSourcesNil(b bool)`

 SetAutoSyncEnabledSourcesNil sets the value for AutoSyncEnabledSources to be an explicit nil

### UnsetAutoSyncEnabledSources
`func (o *UserConfigurationNullable) UnsetAutoSyncEnabledSources()`

UnsetAutoSyncEnabledSources ensures that no value is present for AutoSyncEnabledSources, not even an explicit nil
### GetMaxFiles

`func (o *UserConfigurationNullable) GetMaxFiles() int32`

GetMaxFiles returns the MaxFiles field if non-nil, zero value otherwise.

### GetMaxFilesOk

`func (o *UserConfigurationNullable) GetMaxFilesOk() (*int32, bool)`

GetMaxFilesOk returns a tuple with the MaxFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFiles

`func (o *UserConfigurationNullable) SetMaxFiles(v int32)`

SetMaxFiles sets MaxFiles field to given value.

### HasMaxFiles

`func (o *UserConfigurationNullable) HasMaxFiles() bool`

HasMaxFiles returns a boolean if a field has been set.

### SetMaxFilesNil

`func (o *UserConfigurationNullable) SetMaxFilesNil(b bool)`

 SetMaxFilesNil sets the value for MaxFiles to be an explicit nil

### UnsetMaxFiles
`func (o *UserConfigurationNullable) UnsetMaxFiles()`

UnsetMaxFiles ensures that no value is present for MaxFiles, not even an explicit nil
### GetMaxFilesPerUpload

`func (o *UserConfigurationNullable) GetMaxFilesPerUpload() int32`

GetMaxFilesPerUpload returns the MaxFilesPerUpload field if non-nil, zero value otherwise.

### GetMaxFilesPerUploadOk

`func (o *UserConfigurationNullable) GetMaxFilesPerUploadOk() (*int32, bool)`

GetMaxFilesPerUploadOk returns a tuple with the MaxFilesPerUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFilesPerUpload

`func (o *UserConfigurationNullable) SetMaxFilesPerUpload(v int32)`

SetMaxFilesPerUpload sets MaxFilesPerUpload field to given value.

### HasMaxFilesPerUpload

`func (o *UserConfigurationNullable) HasMaxFilesPerUpload() bool`

HasMaxFilesPerUpload returns a boolean if a field has been set.

### SetMaxFilesPerUploadNil

`func (o *UserConfigurationNullable) SetMaxFilesPerUploadNil(b bool)`

 SetMaxFilesPerUploadNil sets the value for MaxFilesPerUpload to be an explicit nil

### UnsetMaxFilesPerUpload
`func (o *UserConfigurationNullable) UnsetMaxFilesPerUpload()`

UnsetMaxFilesPerUpload ensures that no value is present for MaxFilesPerUpload, not even an explicit nil
### GetMaxCharacters

`func (o *UserConfigurationNullable) GetMaxCharacters() int32`

GetMaxCharacters returns the MaxCharacters field if non-nil, zero value otherwise.

### GetMaxCharactersOk

`func (o *UserConfigurationNullable) GetMaxCharactersOk() (*int32, bool)`

GetMaxCharactersOk returns a tuple with the MaxCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacters

`func (o *UserConfigurationNullable) SetMaxCharacters(v int32)`

SetMaxCharacters sets MaxCharacters field to given value.

### HasMaxCharacters

`func (o *UserConfigurationNullable) HasMaxCharacters() bool`

HasMaxCharacters returns a boolean if a field has been set.

### SetMaxCharactersNil

`func (o *UserConfigurationNullable) SetMaxCharactersNil(b bool)`

 SetMaxCharactersNil sets the value for MaxCharacters to be an explicit nil

### UnsetMaxCharacters
`func (o *UserConfigurationNullable) UnsetMaxCharacters()`

UnsetMaxCharacters ensures that no value is present for MaxCharacters, not even an explicit nil
### GetMaxCharactersPerFile

`func (o *UserConfigurationNullable) GetMaxCharactersPerFile() int32`

GetMaxCharactersPerFile returns the MaxCharactersPerFile field if non-nil, zero value otherwise.

### GetMaxCharactersPerFileOk

`func (o *UserConfigurationNullable) GetMaxCharactersPerFileOk() (*int32, bool)`

GetMaxCharactersPerFileOk returns a tuple with the MaxCharactersPerFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharactersPerFile

`func (o *UserConfigurationNullable) SetMaxCharactersPerFile(v int32)`

SetMaxCharactersPerFile sets MaxCharactersPerFile field to given value.

### HasMaxCharactersPerFile

`func (o *UserConfigurationNullable) HasMaxCharactersPerFile() bool`

HasMaxCharactersPerFile returns a boolean if a field has been set.

### SetMaxCharactersPerFileNil

`func (o *UserConfigurationNullable) SetMaxCharactersPerFileNil(b bool)`

 SetMaxCharactersPerFileNil sets the value for MaxCharactersPerFile to be an explicit nil

### UnsetMaxCharactersPerFile
`func (o *UserConfigurationNullable) UnsetMaxCharactersPerFile()`

UnsetMaxCharactersPerFile ensures that no value is present for MaxCharactersPerFile, not even an explicit nil
### GetMaxCharactersPerUpload

`func (o *UserConfigurationNullable) GetMaxCharactersPerUpload() int32`

GetMaxCharactersPerUpload returns the MaxCharactersPerUpload field if non-nil, zero value otherwise.

### GetMaxCharactersPerUploadOk

`func (o *UserConfigurationNullable) GetMaxCharactersPerUploadOk() (*int32, bool)`

GetMaxCharactersPerUploadOk returns a tuple with the MaxCharactersPerUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharactersPerUpload

`func (o *UserConfigurationNullable) SetMaxCharactersPerUpload(v int32)`

SetMaxCharactersPerUpload sets MaxCharactersPerUpload field to given value.

### HasMaxCharactersPerUpload

`func (o *UserConfigurationNullable) HasMaxCharactersPerUpload() bool`

HasMaxCharactersPerUpload returns a boolean if a field has been set.

### SetMaxCharactersPerUploadNil

`func (o *UserConfigurationNullable) SetMaxCharactersPerUploadNil(b bool)`

 SetMaxCharactersPerUploadNil sets the value for MaxCharactersPerUpload to be an explicit nil

### UnsetMaxCharactersPerUpload
`func (o *UserConfigurationNullable) UnsetMaxCharactersPerUpload()`

UnsetMaxCharactersPerUpload ensures that no value is present for MaxCharactersPerUpload, not even an explicit nil
### GetAutoSyncInterval

`func (o *UserConfigurationNullable) GetAutoSyncInterval() int32`

GetAutoSyncInterval returns the AutoSyncInterval field if non-nil, zero value otherwise.

### GetAutoSyncIntervalOk

`func (o *UserConfigurationNullable) GetAutoSyncIntervalOk() (*int32, bool)`

GetAutoSyncIntervalOk returns a tuple with the AutoSyncInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncInterval

`func (o *UserConfigurationNullable) SetAutoSyncInterval(v int32)`

SetAutoSyncInterval sets AutoSyncInterval field to given value.

### HasAutoSyncInterval

`func (o *UserConfigurationNullable) HasAutoSyncInterval() bool`

HasAutoSyncInterval returns a boolean if a field has been set.

### SetAutoSyncIntervalNil

`func (o *UserConfigurationNullable) SetAutoSyncIntervalNil(b bool)`

 SetAutoSyncIntervalNil sets the value for AutoSyncInterval to be an explicit nil

### UnsetAutoSyncInterval
`func (o *UserConfigurationNullable) UnsetAutoSyncInterval()`

UnsetAutoSyncInterval ensures that no value is present for AutoSyncInterval, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UpdateUsersInput

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
**CustomerIds** | **[]string** | List of organization supplied user IDs | 

## Methods

### NewUpdateUsersInput

`func NewUpdateUsersInput(customerIds []string, ) *UpdateUsersInput`

NewUpdateUsersInput instantiates a new UpdateUsersInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUsersInputWithDefaults

`func NewUpdateUsersInputWithDefaults() *UpdateUsersInput`

NewUpdateUsersInputWithDefaults instantiates a new UpdateUsersInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoSyncEnabledSources

`func (o *UpdateUsersInput) GetAutoSyncEnabledSources() AutoSyncEnabledSourcesProperty`

GetAutoSyncEnabledSources returns the AutoSyncEnabledSources field if non-nil, zero value otherwise.

### GetAutoSyncEnabledSourcesOk

`func (o *UpdateUsersInput) GetAutoSyncEnabledSourcesOk() (*AutoSyncEnabledSourcesProperty, bool)`

GetAutoSyncEnabledSourcesOk returns a tuple with the AutoSyncEnabledSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncEnabledSources

`func (o *UpdateUsersInput) SetAutoSyncEnabledSources(v AutoSyncEnabledSourcesProperty)`

SetAutoSyncEnabledSources sets AutoSyncEnabledSources field to given value.

### HasAutoSyncEnabledSources

`func (o *UpdateUsersInput) HasAutoSyncEnabledSources() bool`

HasAutoSyncEnabledSources returns a boolean if a field has been set.

### SetAutoSyncEnabledSourcesNil

`func (o *UpdateUsersInput) SetAutoSyncEnabledSourcesNil(b bool)`

 SetAutoSyncEnabledSourcesNil sets the value for AutoSyncEnabledSources to be an explicit nil

### UnsetAutoSyncEnabledSources
`func (o *UpdateUsersInput) UnsetAutoSyncEnabledSources()`

UnsetAutoSyncEnabledSources ensures that no value is present for AutoSyncEnabledSources, not even an explicit nil
### GetMaxFiles

`func (o *UpdateUsersInput) GetMaxFiles() int32`

GetMaxFiles returns the MaxFiles field if non-nil, zero value otherwise.

### GetMaxFilesOk

`func (o *UpdateUsersInput) GetMaxFilesOk() (*int32, bool)`

GetMaxFilesOk returns a tuple with the MaxFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFiles

`func (o *UpdateUsersInput) SetMaxFiles(v int32)`

SetMaxFiles sets MaxFiles field to given value.

### HasMaxFiles

`func (o *UpdateUsersInput) HasMaxFiles() bool`

HasMaxFiles returns a boolean if a field has been set.

### SetMaxFilesNil

`func (o *UpdateUsersInput) SetMaxFilesNil(b bool)`

 SetMaxFilesNil sets the value for MaxFiles to be an explicit nil

### UnsetMaxFiles
`func (o *UpdateUsersInput) UnsetMaxFiles()`

UnsetMaxFiles ensures that no value is present for MaxFiles, not even an explicit nil
### GetMaxFilesPerUpload

`func (o *UpdateUsersInput) GetMaxFilesPerUpload() int32`

GetMaxFilesPerUpload returns the MaxFilesPerUpload field if non-nil, zero value otherwise.

### GetMaxFilesPerUploadOk

`func (o *UpdateUsersInput) GetMaxFilesPerUploadOk() (*int32, bool)`

GetMaxFilesPerUploadOk returns a tuple with the MaxFilesPerUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFilesPerUpload

`func (o *UpdateUsersInput) SetMaxFilesPerUpload(v int32)`

SetMaxFilesPerUpload sets MaxFilesPerUpload field to given value.

### HasMaxFilesPerUpload

`func (o *UpdateUsersInput) HasMaxFilesPerUpload() bool`

HasMaxFilesPerUpload returns a boolean if a field has been set.

### SetMaxFilesPerUploadNil

`func (o *UpdateUsersInput) SetMaxFilesPerUploadNil(b bool)`

 SetMaxFilesPerUploadNil sets the value for MaxFilesPerUpload to be an explicit nil

### UnsetMaxFilesPerUpload
`func (o *UpdateUsersInput) UnsetMaxFilesPerUpload()`

UnsetMaxFilesPerUpload ensures that no value is present for MaxFilesPerUpload, not even an explicit nil
### GetMaxCharacters

`func (o *UpdateUsersInput) GetMaxCharacters() int32`

GetMaxCharacters returns the MaxCharacters field if non-nil, zero value otherwise.

### GetMaxCharactersOk

`func (o *UpdateUsersInput) GetMaxCharactersOk() (*int32, bool)`

GetMaxCharactersOk returns a tuple with the MaxCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacters

`func (o *UpdateUsersInput) SetMaxCharacters(v int32)`

SetMaxCharacters sets MaxCharacters field to given value.

### HasMaxCharacters

`func (o *UpdateUsersInput) HasMaxCharacters() bool`

HasMaxCharacters returns a boolean if a field has been set.

### SetMaxCharactersNil

`func (o *UpdateUsersInput) SetMaxCharactersNil(b bool)`

 SetMaxCharactersNil sets the value for MaxCharacters to be an explicit nil

### UnsetMaxCharacters
`func (o *UpdateUsersInput) UnsetMaxCharacters()`

UnsetMaxCharacters ensures that no value is present for MaxCharacters, not even an explicit nil
### GetMaxCharactersPerFile

`func (o *UpdateUsersInput) GetMaxCharactersPerFile() int32`

GetMaxCharactersPerFile returns the MaxCharactersPerFile field if non-nil, zero value otherwise.

### GetMaxCharactersPerFileOk

`func (o *UpdateUsersInput) GetMaxCharactersPerFileOk() (*int32, bool)`

GetMaxCharactersPerFileOk returns a tuple with the MaxCharactersPerFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharactersPerFile

`func (o *UpdateUsersInput) SetMaxCharactersPerFile(v int32)`

SetMaxCharactersPerFile sets MaxCharactersPerFile field to given value.

### HasMaxCharactersPerFile

`func (o *UpdateUsersInput) HasMaxCharactersPerFile() bool`

HasMaxCharactersPerFile returns a boolean if a field has been set.

### SetMaxCharactersPerFileNil

`func (o *UpdateUsersInput) SetMaxCharactersPerFileNil(b bool)`

 SetMaxCharactersPerFileNil sets the value for MaxCharactersPerFile to be an explicit nil

### UnsetMaxCharactersPerFile
`func (o *UpdateUsersInput) UnsetMaxCharactersPerFile()`

UnsetMaxCharactersPerFile ensures that no value is present for MaxCharactersPerFile, not even an explicit nil
### GetMaxCharactersPerUpload

`func (o *UpdateUsersInput) GetMaxCharactersPerUpload() int32`

GetMaxCharactersPerUpload returns the MaxCharactersPerUpload field if non-nil, zero value otherwise.

### GetMaxCharactersPerUploadOk

`func (o *UpdateUsersInput) GetMaxCharactersPerUploadOk() (*int32, bool)`

GetMaxCharactersPerUploadOk returns a tuple with the MaxCharactersPerUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharactersPerUpload

`func (o *UpdateUsersInput) SetMaxCharactersPerUpload(v int32)`

SetMaxCharactersPerUpload sets MaxCharactersPerUpload field to given value.

### HasMaxCharactersPerUpload

`func (o *UpdateUsersInput) HasMaxCharactersPerUpload() bool`

HasMaxCharactersPerUpload returns a boolean if a field has been set.

### SetMaxCharactersPerUploadNil

`func (o *UpdateUsersInput) SetMaxCharactersPerUploadNil(b bool)`

 SetMaxCharactersPerUploadNil sets the value for MaxCharactersPerUpload to be an explicit nil

### UnsetMaxCharactersPerUpload
`func (o *UpdateUsersInput) UnsetMaxCharactersPerUpload()`

UnsetMaxCharactersPerUpload ensures that no value is present for MaxCharactersPerUpload, not even an explicit nil
### GetAutoSyncInterval

`func (o *UpdateUsersInput) GetAutoSyncInterval() int32`

GetAutoSyncInterval returns the AutoSyncInterval field if non-nil, zero value otherwise.

### GetAutoSyncIntervalOk

`func (o *UpdateUsersInput) GetAutoSyncIntervalOk() (*int32, bool)`

GetAutoSyncIntervalOk returns a tuple with the AutoSyncInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncInterval

`func (o *UpdateUsersInput) SetAutoSyncInterval(v int32)`

SetAutoSyncInterval sets AutoSyncInterval field to given value.

### HasAutoSyncInterval

`func (o *UpdateUsersInput) HasAutoSyncInterval() bool`

HasAutoSyncInterval returns a boolean if a field has been set.

### SetAutoSyncIntervalNil

`func (o *UpdateUsersInput) SetAutoSyncIntervalNil(b bool)`

 SetAutoSyncIntervalNil sets the value for AutoSyncInterval to be an explicit nil

### UnsetAutoSyncInterval
`func (o *UpdateUsersInput) UnsetAutoSyncInterval()`

UnsetAutoSyncInterval ensures that no value is present for AutoSyncInterval, not even an explicit nil
### GetCustomerIds

`func (o *UpdateUsersInput) GetCustomerIds() []string`

GetCustomerIds returns the CustomerIds field if non-nil, zero value otherwise.

### GetCustomerIdsOk

`func (o *UpdateUsersInput) GetCustomerIdsOk() (*[]string, bool)`

GetCustomerIdsOk returns a tuple with the CustomerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIds

`func (o *UpdateUsersInput) SetCustomerIds(v []string)`

SetCustomerIds sets CustomerIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UserConfigurationNullable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSyncEnabledSources** | Pointer to [**NullableAutoSyncEnabledSourcesProperty**](AutoSyncEnabledSourcesProperty.md) |  | [optional] 
**MaxFiles** | Pointer to **NullableInt32** | Custom file upload limit for the user over *all* user&#39;s files across all uploads.          If set, then the user will not be allowed to upload more files than this limit. If not set, or if set to -1,         then the user will have no limit. | [optional] 
**MaxFilesPerUpload** | Pointer to **NullableInt32** | Custom file upload limit for the user across a single upload.         If set, then the user will not be allowed to upload more files than this limit in a single upload. If not set,         or if set to -1, then the user will have no limit. | [optional] 

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



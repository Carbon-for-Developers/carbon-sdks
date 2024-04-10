# UpdateUsersInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerIds** | **[]string** | List of organization supplied user IDs | 
**AutoSyncEnabledSources** | Pointer to [**NullableAutoSyncEnabledSourcesProperty**](AutoSyncEnabledSourcesProperty.md) |  | [optional] 
**FileUploadLimit** | Pointer to **NullableInt32** | Custom file upload limit for the user. If set, then the user will not be allowed to          upload more files than this limit | [optional] 

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
### GetFileUploadLimit

`func (o *UpdateUsersInput) GetFileUploadLimit() int32`

GetFileUploadLimit returns the FileUploadLimit field if non-nil, zero value otherwise.

### GetFileUploadLimitOk

`func (o *UpdateUsersInput) GetFileUploadLimitOk() (*int32, bool)`

GetFileUploadLimitOk returns a tuple with the FileUploadLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUploadLimit

`func (o *UpdateUsersInput) SetFileUploadLimit(v int32)`

SetFileUploadLimit sets FileUploadLimit field to given value.

### HasFileUploadLimit

`func (o *UpdateUsersInput) HasFileUploadLimit() bool`

HasFileUploadLimit returns a boolean if a field has been set.

### SetFileUploadLimitNil

`func (o *UpdateUsersInput) SetFileUploadLimitNil(b bool)`

 SetFileUploadLimitNil sets the value for FileUploadLimit to be an explicit nil

### UnsetFileUploadLimit
`func (o *UpdateUsersInput) UnsetFileUploadLimit()`

UnsetFileUploadLimit ensures that no value is present for FileUploadLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



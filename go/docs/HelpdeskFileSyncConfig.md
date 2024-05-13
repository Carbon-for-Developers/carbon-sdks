# HelpdeskFileSyncConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileTypes** | Pointer to [**[]HelpdeskFileTypes**](HelpdeskFileTypes.md) |  | [optional] [default to ["ARTICLE"]]
**SyncAttachments** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewHelpdeskFileSyncConfig

`func NewHelpdeskFileSyncConfig() *HelpdeskFileSyncConfig`

NewHelpdeskFileSyncConfig instantiates a new HelpdeskFileSyncConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelpdeskFileSyncConfigWithDefaults

`func NewHelpdeskFileSyncConfigWithDefaults() *HelpdeskFileSyncConfig`

NewHelpdeskFileSyncConfigWithDefaults instantiates a new HelpdeskFileSyncConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileTypes

`func (o *HelpdeskFileSyncConfig) GetFileTypes() []HelpdeskFileTypes`

GetFileTypes returns the FileTypes field if non-nil, zero value otherwise.

### GetFileTypesOk

`func (o *HelpdeskFileSyncConfig) GetFileTypesOk() (*[]HelpdeskFileTypes, bool)`

GetFileTypesOk returns a tuple with the FileTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTypes

`func (o *HelpdeskFileSyncConfig) SetFileTypes(v []HelpdeskFileTypes)`

SetFileTypes sets FileTypes field to given value.

### HasFileTypes

`func (o *HelpdeskFileSyncConfig) HasFileTypes() bool`

HasFileTypes returns a boolean if a field has been set.

### GetSyncAttachments

`func (o *HelpdeskFileSyncConfig) GetSyncAttachments() bool`

GetSyncAttachments returns the SyncAttachments field if non-nil, zero value otherwise.

### GetSyncAttachmentsOk

`func (o *HelpdeskFileSyncConfig) GetSyncAttachmentsOk() (*bool, bool)`

GetSyncAttachmentsOk returns a tuple with the SyncAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncAttachments

`func (o *HelpdeskFileSyncConfig) SetSyncAttachments(v bool)`

SetSyncAttachments sets SyncAttachments field to given value.

### HasSyncAttachments

`func (o *HelpdeskFileSyncConfig) HasSyncAttachments() bool`

HasSyncAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



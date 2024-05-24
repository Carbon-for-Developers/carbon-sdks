# FileSyncConfigNullable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSyncedSourceTypes** | Pointer to [**[]HelpdeskFileTypes**](HelpdeskFileTypes.md) | File types to automatically sync when the data source connects. Only a subset of file types can be          controlled. If not supported, then they will always be synced | [optional] [default to ["ARTICLE"]]
**SyncAttachments** | Pointer to **bool** | Automatically sync attachments from files where supported. Currently applies to Helpdesk Tickets | [optional] [default to false]
**DetectAudioLanguage** | Pointer to **bool** | Detect audio language before transcription for audio files | [optional] [default to false]

## Methods

### NewFileSyncConfigNullable

`func NewFileSyncConfigNullable() *FileSyncConfigNullable`

NewFileSyncConfigNullable instantiates a new FileSyncConfigNullable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSyncConfigNullableWithDefaults

`func NewFileSyncConfigNullableWithDefaults() *FileSyncConfigNullable`

NewFileSyncConfigNullableWithDefaults instantiates a new FileSyncConfigNullable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoSyncedSourceTypes

`func (o *FileSyncConfigNullable) GetAutoSyncedSourceTypes() []HelpdeskFileTypes`

GetAutoSyncedSourceTypes returns the AutoSyncedSourceTypes field if non-nil, zero value otherwise.

### GetAutoSyncedSourceTypesOk

`func (o *FileSyncConfigNullable) GetAutoSyncedSourceTypesOk() (*[]HelpdeskFileTypes, bool)`

GetAutoSyncedSourceTypesOk returns a tuple with the AutoSyncedSourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncedSourceTypes

`func (o *FileSyncConfigNullable) SetAutoSyncedSourceTypes(v []HelpdeskFileTypes)`

SetAutoSyncedSourceTypes sets AutoSyncedSourceTypes field to given value.

### HasAutoSyncedSourceTypes

`func (o *FileSyncConfigNullable) HasAutoSyncedSourceTypes() bool`

HasAutoSyncedSourceTypes returns a boolean if a field has been set.

### GetSyncAttachments

`func (o *FileSyncConfigNullable) GetSyncAttachments() bool`

GetSyncAttachments returns the SyncAttachments field if non-nil, zero value otherwise.

### GetSyncAttachmentsOk

`func (o *FileSyncConfigNullable) GetSyncAttachmentsOk() (*bool, bool)`

GetSyncAttachmentsOk returns a tuple with the SyncAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncAttachments

`func (o *FileSyncConfigNullable) SetSyncAttachments(v bool)`

SetSyncAttachments sets SyncAttachments field to given value.

### HasSyncAttachments

`func (o *FileSyncConfigNullable) HasSyncAttachments() bool`

HasSyncAttachments returns a boolean if a field has been set.

### GetDetectAudioLanguage

`func (o *FileSyncConfigNullable) GetDetectAudioLanguage() bool`

GetDetectAudioLanguage returns the DetectAudioLanguage field if non-nil, zero value otherwise.

### GetDetectAudioLanguageOk

`func (o *FileSyncConfigNullable) GetDetectAudioLanguageOk() (*bool, bool)`

GetDetectAudioLanguageOk returns a tuple with the DetectAudioLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectAudioLanguage

`func (o *FileSyncConfigNullable) SetDetectAudioLanguage(v bool)`

SetDetectAudioLanguage sets DetectAudioLanguage field to given value.

### HasDetectAudioLanguage

`func (o *FileSyncConfigNullable) HasDetectAudioLanguage() bool`

HasDetectAudioLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# FileSyncConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSyncedSourceTypes** | Pointer to [**[]HelpdeskFileTypes**](HelpdeskFileTypes.md) | File types to automatically sync when the data source connects. Only a subset of file types can be          controlled. If not supported, then they will always be synced | [optional] [default to ["ARTICLE"]]
**SyncAttachments** | Pointer to **bool** | Automatically sync attachments from files where supported. Currently applies to Helpdesk Tickets | [optional] [default to false]
**DetectAudioLanguage** | Pointer to **bool** | Detect audio language before transcription for audio files | [optional] [default to false]
**TranscriptionService** | Pointer to [**NullableTranscriptionServiceNullable**](TranscriptionServiceNullable.md) |  | [optional] 
**SplitRows** | Pointer to **bool** | Whether to split tabular rows into chunks. Currently only valid for CSV, TSV, and XLSX files. | [optional] [default to false]

## Methods

### NewFileSyncConfig

`func NewFileSyncConfig() *FileSyncConfig`

NewFileSyncConfig instantiates a new FileSyncConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSyncConfigWithDefaults

`func NewFileSyncConfigWithDefaults() *FileSyncConfig`

NewFileSyncConfigWithDefaults instantiates a new FileSyncConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoSyncedSourceTypes

`func (o *FileSyncConfig) GetAutoSyncedSourceTypes() []HelpdeskFileTypes`

GetAutoSyncedSourceTypes returns the AutoSyncedSourceTypes field if non-nil, zero value otherwise.

### GetAutoSyncedSourceTypesOk

`func (o *FileSyncConfig) GetAutoSyncedSourceTypesOk() (*[]HelpdeskFileTypes, bool)`

GetAutoSyncedSourceTypesOk returns a tuple with the AutoSyncedSourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncedSourceTypes

`func (o *FileSyncConfig) SetAutoSyncedSourceTypes(v []HelpdeskFileTypes)`

SetAutoSyncedSourceTypes sets AutoSyncedSourceTypes field to given value.

### HasAutoSyncedSourceTypes

`func (o *FileSyncConfig) HasAutoSyncedSourceTypes() bool`

HasAutoSyncedSourceTypes returns a boolean if a field has been set.

### GetSyncAttachments

`func (o *FileSyncConfig) GetSyncAttachments() bool`

GetSyncAttachments returns the SyncAttachments field if non-nil, zero value otherwise.

### GetSyncAttachmentsOk

`func (o *FileSyncConfig) GetSyncAttachmentsOk() (*bool, bool)`

GetSyncAttachmentsOk returns a tuple with the SyncAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncAttachments

`func (o *FileSyncConfig) SetSyncAttachments(v bool)`

SetSyncAttachments sets SyncAttachments field to given value.

### HasSyncAttachments

`func (o *FileSyncConfig) HasSyncAttachments() bool`

HasSyncAttachments returns a boolean if a field has been set.

### GetDetectAudioLanguage

`func (o *FileSyncConfig) GetDetectAudioLanguage() bool`

GetDetectAudioLanguage returns the DetectAudioLanguage field if non-nil, zero value otherwise.

### GetDetectAudioLanguageOk

`func (o *FileSyncConfig) GetDetectAudioLanguageOk() (*bool, bool)`

GetDetectAudioLanguageOk returns a tuple with the DetectAudioLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectAudioLanguage

`func (o *FileSyncConfig) SetDetectAudioLanguage(v bool)`

SetDetectAudioLanguage sets DetectAudioLanguage field to given value.

### HasDetectAudioLanguage

`func (o *FileSyncConfig) HasDetectAudioLanguage() bool`

HasDetectAudioLanguage returns a boolean if a field has been set.

### GetTranscriptionService

`func (o *FileSyncConfig) GetTranscriptionService() TranscriptionServiceNullable`

GetTranscriptionService returns the TranscriptionService field if non-nil, zero value otherwise.

### GetTranscriptionServiceOk

`func (o *FileSyncConfig) GetTranscriptionServiceOk() (*TranscriptionServiceNullable, bool)`

GetTranscriptionServiceOk returns a tuple with the TranscriptionService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionService

`func (o *FileSyncConfig) SetTranscriptionService(v TranscriptionServiceNullable)`

SetTranscriptionService sets TranscriptionService field to given value.

### HasTranscriptionService

`func (o *FileSyncConfig) HasTranscriptionService() bool`

HasTranscriptionService returns a boolean if a field has been set.

### SetTranscriptionServiceNil

`func (o *FileSyncConfig) SetTranscriptionServiceNil(b bool)`

 SetTranscriptionServiceNil sets the value for TranscriptionService to be an explicit nil

### UnsetTranscriptionService
`func (o *FileSyncConfig) UnsetTranscriptionService()`

UnsetTranscriptionService ensures that no value is present for TranscriptionService, not even an explicit nil
### GetSplitRows

`func (o *FileSyncConfig) GetSplitRows() bool`

GetSplitRows returns the SplitRows field if non-nil, zero value otherwise.

### GetSplitRowsOk

`func (o *FileSyncConfig) GetSplitRowsOk() (*bool, bool)`

GetSplitRowsOk returns a tuple with the SplitRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitRows

`func (o *FileSyncConfig) SetSplitRows(v bool)`

SetSplitRows sets SplitRows field to given value.

### HasSplitRows

`func (o *FileSyncConfig) HasSplitRows() bool`

HasSplitRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



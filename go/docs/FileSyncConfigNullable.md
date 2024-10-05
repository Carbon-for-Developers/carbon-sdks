# FileSyncConfigNullable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSyncedSourceTypes** | Pointer to [**[]HSNFileTypes3**](HSNFileTypes3.md) | File types to automatically sync when the data source connects. Only a subset of file types can be          controlled. If not supported, then they will always be synced | [optional] [default to ["ARTICLE"]]
**SyncAttachments** | Pointer to **bool** | Automatically sync attachments from files where supported. Currently applies to Helpdesk Tickets | [optional] [default to false]
**DetectAudioLanguage** | Pointer to **bool** | Detect audio language before transcription for audio files | [optional] [default to false]
**TranscriptionService** | Pointer to [**NullableTranscriptionServiceNullable**](TranscriptionServiceNullable.md) |  | [optional] 
**IncludeSpeakerLabels** | Pointer to **bool** | Detect multiple speakers and label segments of speech by speaker for audio files. | [optional] [default to false]
**SplitRows** | Pointer to **bool** | Whether to split tabular rows into chunks. Currently only valid for CSV, TSV, and XLSX files. | [optional] [default to false]
**GenerateChunksOnly** | Pointer to **bool** | If this flag is enabled, the file will be chunked and stored with Carbon,           but no embeddings will be generated. This overrides the skip_embedding_generation flag. | [optional] [default to false]
**StoreFileOnly** | Pointer to **bool** | If this flag is enabled, the file will be stored with Carbon, but no chunks or embeddings will be generated.            This overrides the skip_embedding_generation and generate_chunks_only flags. | [optional] [default to false]
**SkipFileProcessing** | Pointer to **bool** | Setting this flag will create a new file record with Carbon but skip any and all processing.          This means that we do not download the remote file content or generate any chunks or embeddings. We will store         some metadata like name, external id, and external URL depending on the source you are syncing from. Note that this          flag overrides both skip_embedding_generation and generate_chunks_only flags. The file will be moved to          READY_TO_SYNC status. | [optional] [default to false]

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

`func (o *FileSyncConfigNullable) GetAutoSyncedSourceTypes() []HSNFileTypes3`

GetAutoSyncedSourceTypes returns the AutoSyncedSourceTypes field if non-nil, zero value otherwise.

### GetAutoSyncedSourceTypesOk

`func (o *FileSyncConfigNullable) GetAutoSyncedSourceTypesOk() (*[]HSNFileTypes3, bool)`

GetAutoSyncedSourceTypesOk returns a tuple with the AutoSyncedSourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncedSourceTypes

`func (o *FileSyncConfigNullable) SetAutoSyncedSourceTypes(v []HSNFileTypes3)`

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

### GetTranscriptionService

`func (o *FileSyncConfigNullable) GetTranscriptionService() TranscriptionServiceNullable`

GetTranscriptionService returns the TranscriptionService field if non-nil, zero value otherwise.

### GetTranscriptionServiceOk

`func (o *FileSyncConfigNullable) GetTranscriptionServiceOk() (*TranscriptionServiceNullable, bool)`

GetTranscriptionServiceOk returns a tuple with the TranscriptionService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionService

`func (o *FileSyncConfigNullable) SetTranscriptionService(v TranscriptionServiceNullable)`

SetTranscriptionService sets TranscriptionService field to given value.

### HasTranscriptionService

`func (o *FileSyncConfigNullable) HasTranscriptionService() bool`

HasTranscriptionService returns a boolean if a field has been set.

### SetTranscriptionServiceNil

`func (o *FileSyncConfigNullable) SetTranscriptionServiceNil(b bool)`

 SetTranscriptionServiceNil sets the value for TranscriptionService to be an explicit nil

### UnsetTranscriptionService
`func (o *FileSyncConfigNullable) UnsetTranscriptionService()`

UnsetTranscriptionService ensures that no value is present for TranscriptionService, not even an explicit nil
### GetIncludeSpeakerLabels

`func (o *FileSyncConfigNullable) GetIncludeSpeakerLabels() bool`

GetIncludeSpeakerLabels returns the IncludeSpeakerLabels field if non-nil, zero value otherwise.

### GetIncludeSpeakerLabelsOk

`func (o *FileSyncConfigNullable) GetIncludeSpeakerLabelsOk() (*bool, bool)`

GetIncludeSpeakerLabelsOk returns a tuple with the IncludeSpeakerLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSpeakerLabels

`func (o *FileSyncConfigNullable) SetIncludeSpeakerLabels(v bool)`

SetIncludeSpeakerLabels sets IncludeSpeakerLabels field to given value.

### HasIncludeSpeakerLabels

`func (o *FileSyncConfigNullable) HasIncludeSpeakerLabels() bool`

HasIncludeSpeakerLabels returns a boolean if a field has been set.

### GetSplitRows

`func (o *FileSyncConfigNullable) GetSplitRows() bool`

GetSplitRows returns the SplitRows field if non-nil, zero value otherwise.

### GetSplitRowsOk

`func (o *FileSyncConfigNullable) GetSplitRowsOk() (*bool, bool)`

GetSplitRowsOk returns a tuple with the SplitRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitRows

`func (o *FileSyncConfigNullable) SetSplitRows(v bool)`

SetSplitRows sets SplitRows field to given value.

### HasSplitRows

`func (o *FileSyncConfigNullable) HasSplitRows() bool`

HasSplitRows returns a boolean if a field has been set.

### GetGenerateChunksOnly

`func (o *FileSyncConfigNullable) GetGenerateChunksOnly() bool`

GetGenerateChunksOnly returns the GenerateChunksOnly field if non-nil, zero value otherwise.

### GetGenerateChunksOnlyOk

`func (o *FileSyncConfigNullable) GetGenerateChunksOnlyOk() (*bool, bool)`

GetGenerateChunksOnlyOk returns a tuple with the GenerateChunksOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateChunksOnly

`func (o *FileSyncConfigNullable) SetGenerateChunksOnly(v bool)`

SetGenerateChunksOnly sets GenerateChunksOnly field to given value.

### HasGenerateChunksOnly

`func (o *FileSyncConfigNullable) HasGenerateChunksOnly() bool`

HasGenerateChunksOnly returns a boolean if a field has been set.

### GetStoreFileOnly

`func (o *FileSyncConfigNullable) GetStoreFileOnly() bool`

GetStoreFileOnly returns the StoreFileOnly field if non-nil, zero value otherwise.

### GetStoreFileOnlyOk

`func (o *FileSyncConfigNullable) GetStoreFileOnlyOk() (*bool, bool)`

GetStoreFileOnlyOk returns a tuple with the StoreFileOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreFileOnly

`func (o *FileSyncConfigNullable) SetStoreFileOnly(v bool)`

SetStoreFileOnly sets StoreFileOnly field to given value.

### HasStoreFileOnly

`func (o *FileSyncConfigNullable) HasStoreFileOnly() bool`

HasStoreFileOnly returns a boolean if a field has been set.

### GetSkipFileProcessing

`func (o *FileSyncConfigNullable) GetSkipFileProcessing() bool`

GetSkipFileProcessing returns the SkipFileProcessing field if non-nil, zero value otherwise.

### GetSkipFileProcessingOk

`func (o *FileSyncConfigNullable) GetSkipFileProcessingOk() (*bool, bool)`

GetSkipFileProcessingOk returns a tuple with the SkipFileProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipFileProcessing

`func (o *FileSyncConfigNullable) SetSkipFileProcessing(v bool)`

SetSkipFileProcessing sets SkipFileProcessing field to given value.

### HasSkipFileProcessing

`func (o *FileSyncConfigNullable) HasSkipFileProcessing() bool`

HasSkipFileProcessing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



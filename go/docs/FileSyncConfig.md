# FileSyncConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSyncedSourceTypes** | Pointer to [**[]HSNFileTypes**](HSNFileTypes.md) | File types to automatically sync when the data source connects. Only a subset of file types can be          controlled. If not supported, then they will always be synced | [optional] [default to ["ARTICLE"]]
**SyncAttachments** | Pointer to **bool** | Automatically sync attachments from files where supported. Currently applies to Helpdesk Tickets | [optional] [default to false]
**DetectAudioLanguage** | Pointer to **bool** | Detect audio language before transcription for audio files | [optional] [default to false]
**TranscriptionService** | Pointer to [**NullableTranscriptionServiceNullable**](TranscriptionServiceNullable.md) |  | [optional] 
**IncludeSpeakerLabels** | Pointer to **bool** | Detect multiple speakers and label segments of speech by speaker for audio files. | [optional] [default to false]
**SplitRows** | Pointer to **bool** | Whether to split tabular rows into chunks. Currently only valid for CSV, TSV, and XLSX files. | [optional] [default to false]
**GenerateChunksOnly** | Pointer to **bool** | If this flag is enabled, the file will be chunked and stored with Carbon,           but no embeddings will be generated. This overrides the skip_embedding_generation flag. | [optional] [default to false]
**StoreFileOnly** | Pointer to **bool** | If this flag is enabled, the file will be stored with Carbon, but no chunks or embeddings will be generated.            This overrides the skip_embedding_generation and generate_chunks_only flags. | [optional] [default to false]
**SkipFileProcessing** | Pointer to **bool** | Setting this flag will create a new file record with Carbon but skip any and all processing.          This means that we do not download the remote file content or generate any chunks or embeddings. We will store         some metadata like name, external id, and external URL depending on the source you are syncing from. Note that this          flag overrides both skip_embedding_generation and generate_chunks_only flags. The file will be moved to          READY_TO_SYNC status. | [optional] [default to false]
**ParsedTextFormat** | Pointer to [**NullableParsedTextFormatsNullable**](ParsedTextFormatsNullable.md) |  | [optional] 

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

`func (o *FileSyncConfig) GetAutoSyncedSourceTypes() []HSNFileTypes`

GetAutoSyncedSourceTypes returns the AutoSyncedSourceTypes field if non-nil, zero value otherwise.

### GetAutoSyncedSourceTypesOk

`func (o *FileSyncConfig) GetAutoSyncedSourceTypesOk() (*[]HSNFileTypes, bool)`

GetAutoSyncedSourceTypesOk returns a tuple with the AutoSyncedSourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncedSourceTypes

`func (o *FileSyncConfig) SetAutoSyncedSourceTypes(v []HSNFileTypes)`

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
### GetIncludeSpeakerLabels

`func (o *FileSyncConfig) GetIncludeSpeakerLabels() bool`

GetIncludeSpeakerLabels returns the IncludeSpeakerLabels field if non-nil, zero value otherwise.

### GetIncludeSpeakerLabelsOk

`func (o *FileSyncConfig) GetIncludeSpeakerLabelsOk() (*bool, bool)`

GetIncludeSpeakerLabelsOk returns a tuple with the IncludeSpeakerLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSpeakerLabels

`func (o *FileSyncConfig) SetIncludeSpeakerLabels(v bool)`

SetIncludeSpeakerLabels sets IncludeSpeakerLabels field to given value.

### HasIncludeSpeakerLabels

`func (o *FileSyncConfig) HasIncludeSpeakerLabels() bool`

HasIncludeSpeakerLabels returns a boolean if a field has been set.

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

### GetGenerateChunksOnly

`func (o *FileSyncConfig) GetGenerateChunksOnly() bool`

GetGenerateChunksOnly returns the GenerateChunksOnly field if non-nil, zero value otherwise.

### GetGenerateChunksOnlyOk

`func (o *FileSyncConfig) GetGenerateChunksOnlyOk() (*bool, bool)`

GetGenerateChunksOnlyOk returns a tuple with the GenerateChunksOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateChunksOnly

`func (o *FileSyncConfig) SetGenerateChunksOnly(v bool)`

SetGenerateChunksOnly sets GenerateChunksOnly field to given value.

### HasGenerateChunksOnly

`func (o *FileSyncConfig) HasGenerateChunksOnly() bool`

HasGenerateChunksOnly returns a boolean if a field has been set.

### GetStoreFileOnly

`func (o *FileSyncConfig) GetStoreFileOnly() bool`

GetStoreFileOnly returns the StoreFileOnly field if non-nil, zero value otherwise.

### GetStoreFileOnlyOk

`func (o *FileSyncConfig) GetStoreFileOnlyOk() (*bool, bool)`

GetStoreFileOnlyOk returns a tuple with the StoreFileOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreFileOnly

`func (o *FileSyncConfig) SetStoreFileOnly(v bool)`

SetStoreFileOnly sets StoreFileOnly field to given value.

### HasStoreFileOnly

`func (o *FileSyncConfig) HasStoreFileOnly() bool`

HasStoreFileOnly returns a boolean if a field has been set.

### GetSkipFileProcessing

`func (o *FileSyncConfig) GetSkipFileProcessing() bool`

GetSkipFileProcessing returns the SkipFileProcessing field if non-nil, zero value otherwise.

### GetSkipFileProcessingOk

`func (o *FileSyncConfig) GetSkipFileProcessingOk() (*bool, bool)`

GetSkipFileProcessingOk returns a tuple with the SkipFileProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipFileProcessing

`func (o *FileSyncConfig) SetSkipFileProcessing(v bool)`

SetSkipFileProcessing sets SkipFileProcessing field to given value.

### HasSkipFileProcessing

`func (o *FileSyncConfig) HasSkipFileProcessing() bool`

HasSkipFileProcessing returns a boolean if a field has been set.

### GetParsedTextFormat

`func (o *FileSyncConfig) GetParsedTextFormat() ParsedTextFormatsNullable`

GetParsedTextFormat returns the ParsedTextFormat field if non-nil, zero value otherwise.

### GetParsedTextFormatOk

`func (o *FileSyncConfig) GetParsedTextFormatOk() (*ParsedTextFormatsNullable, bool)`

GetParsedTextFormatOk returns a tuple with the ParsedTextFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsedTextFormat

`func (o *FileSyncConfig) SetParsedTextFormat(v ParsedTextFormatsNullable)`

SetParsedTextFormat sets ParsedTextFormat field to given value.

### HasParsedTextFormat

`func (o *FileSyncConfig) HasParsedTextFormat() bool`

HasParsedTextFormat returns a boolean if a field has been set.

### SetParsedTextFormatNil

`func (o *FileSyncConfig) SetParsedTextFormatNil(b bool)`

 SetParsedTextFormatNil sets the value for ParsedTextFormat to be an explicit nil

### UnsetParsedTextFormat
`func (o *FileSyncConfig) UnsetParsedTextFormat()`

UnsetParsedTextFormat ensures that no value is present for ParsedTextFormat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



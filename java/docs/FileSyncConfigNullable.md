

# FileSyncConfigNullable

Used to configure file syncing for certain connectors when sync_files_on_connection is set to true

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**autoSyncedSourceTypes** | **List&lt;HSNFileTypes3&gt;** | File types to automatically sync when the data source connects. Only a subset of file types can be          controlled. If not supported, then they will always be synced |  [optional] |
|**syncAttachments** | **Boolean** | Automatically sync attachments from files where supported. Currently applies to Helpdesk Tickets |  [optional] |
|**detectAudioLanguage** | **Boolean** | Detect audio language before transcription for audio files |  [optional] |
|**transcriptionService** | **TranscriptionServiceNullable** |  |  [optional] |
|**includeSpeakerLabels** | **Boolean** | Detect multiple speakers and label segments of speech by speaker for audio files. |  [optional] |
|**splitRows** | **Boolean** | Whether to split tabular rows into chunks. Currently only valid for CSV, TSV, and XLSX files. |  [optional] |
|**generateChunksOnly** | **Boolean** | If this flag is enabled, the file will be chunked and stored with Carbon,           but no embeddings will be generated. This overrides the skip_embedding_generation flag. |  [optional] |
|**skipFileProcessing** | **Boolean** | Setting this flag will create a new file record with Carbon but skip any and all processing.          This means that we do not download the remote file content or generate any chunks or embeddings. We will store         some metadata like name, external id, and external URL depending on the source you are syncing from. Note that this          flag overrides both skip_embedding_generation and generate_chunks_only flags. The file will be moved to          READY_TO_SYNC status. |  [optional] |




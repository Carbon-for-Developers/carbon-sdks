

# UploadFileFromUrlInput


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**url** | **String** |  |  |
|**fileName** | **String** |  |  [optional] |
|**chunkSize** | **Integer** |  |  [optional] |
|**chunkOverlap** | **Integer** |  |  [optional] |
|**skipEmbeddingGeneration** | **Boolean** |  |  [optional] |
|**setPageAsBoundary** | **Boolean** |  |  [optional] |
|**embeddingModel** | **EmbeddingGenerators** |  |  [optional] |
|**generateSparseVectors** | **Boolean** |  |  [optional] |
|**useTextract** | **Boolean** |  |  [optional] |
|**prependFilenameToChunks** | **Boolean** |  |  [optional] |
|**maxItemsPerChunk** | **Integer** | Number of objects per chunk. For csv, tsv, xlsx, and json files only. |  [optional] |
|**parsePdfTablesWithOcr** | **Boolean** |  |  [optional] |
|**detectAudioLanguage** | **Boolean** |  |  [optional] |
|**transcriptionService** | **TranscriptionServiceNullable** |  |  [optional] |
|**includeSpeakerLabels** | **Boolean** |  |  [optional] |
|**mediaType** | **FileContentTypesNullable** |  |  [optional] |
|**splitRows** | **Boolean** |  |  [optional] |
|**coldStorageParams** | [**ColdStorageProps**](ColdStorageProps.md) |  |  [optional] |
|**generateChunksOnly** | **Boolean** | If this flag is enabled, the file will be chunked and stored with Carbon,         but no embeddings will be generated. This overrides the skip_embedding_generation flag. |  [optional] |




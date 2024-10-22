

# FreshDeskConnectRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**tags** | **Object** |  |  [optional] |
|**domain** | **String** |  |  |
|**apiKey** | **String** |  |  |
|**chunkSize** | **Integer** |  |  [optional] |
|**chunkOverlap** | **Integer** |  |  [optional] |
|**skipEmbeddingGeneration** | **Boolean** |  |  [optional] |
|**embeddingModel** | **EmbeddingGeneratorsNullable** |  |  [optional] |
|**generateSparseVectors** | **Boolean** |  |  [optional] |
|**prependFilenameToChunks** | **Boolean** |  |  [optional] |
|**syncFilesOnConnection** | **Boolean** |  |  [optional] |
|**requestId** | **String** |  |  [optional] |
|**syncSourceItems** | **Boolean** | Enabling this flag will fetch all available content from the source to be listed via list items endpoint |  [optional] |
|**fileSyncConfig** | [**FileSyncConfigNullable**](FileSyncConfigNullable.md) |  |  [optional] |
|**dataSourceTags** | **Object** | Tags to be associated with the data source. If the data source already has tags set, then an upsert will be performed. |  [optional] |




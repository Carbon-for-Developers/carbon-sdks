

# Document360ConnectRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**tags** | **Object** |  |  [optional] |
|**accountEmail** | **String** | This email will be used to identify your carbon data source. It should have access to the          Document360 account you wish to connect. |  |
|**accessToken** | **String** |  |  |
|**chunkSize** | **Integer** |  |  [optional] |
|**chunkOverlap** | **Integer** |  |  [optional] |
|**skipEmbeddingGeneration** | **Boolean** |  |  [optional] |
|**embeddingModel** | **EmbeddingGenerators** |  |  [optional] |
|**generateSparseVectors** | **Boolean** |  |  [optional] |
|**prependFilenameToChunks** | **Boolean** |  |  [optional] |
|**syncFilesOnConnection** | **Boolean** |  |  [optional] |
|**requestId** | **String** |  |  [optional] |
|**syncSourceItems** | **Boolean** | Enabling this flag will fetch all available content from the source to be listed via list items endpoint |  [optional] |
|**fileSyncConfig** | [**FileSyncConfigNullable**](FileSyncConfigNullable.md) |  |  [optional] |
|**dataSourceTags** | **Object** | Tags to be associated with the data source. If the data source already has tags set, then an upsert will be performed. |  [optional] |



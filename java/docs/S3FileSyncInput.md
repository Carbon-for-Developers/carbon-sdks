

# S3FileSyncInput


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**tags** | **Object** |  |  [optional] |
|**ids** | [**List&lt;S3GetFileInput&gt;**](S3GetFileInput.md) | Each input should be one of the following: A bucket name, a bucket name and a prefix, or a          bucket name and an object key. A prefix is the common path for all objects you want to sync.          Paths should end with a forward slash. |  |
|**chunkSize** | **Integer** |  |  [optional] |
|**chunkOverlap** | **Integer** |  |  [optional] |
|**skipEmbeddingGeneration** | **Boolean** |  |  [optional] |
|**embeddingModel** | **EmbeddingGenerators** |  |  [optional] |
|**generateSparseVectors** | **Boolean** |  |  [optional] |
|**prependFilenameToChunks** | **Boolean** |  |  [optional] |
|**maxItemsPerChunk** | **Integer** | Number of objects per chunk. For csv, tsv, xlsx, and json files only. |  [optional] |
|**setPageAsBoundary** | **Boolean** |  |  [optional] |
|**dataSourceId** | **Integer** |  |  [optional] |
|**requestId** | **String** |  |  [optional] |
|**useOcr** | **Boolean** |  |  [optional] |
|**parsePdfTablesWithOcr** | **Boolean** |  |  [optional] |
|**fileSyncConfig** | [**FileSyncConfigNullable**](FileSyncConfigNullable.md) |  |  [optional] |




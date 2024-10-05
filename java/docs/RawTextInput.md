

# RawTextInput


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**contents** | **String** |  |  |
|**name** | **String** |  |  [optional] |
|**chunkSize** | **Integer** |  |  [optional] |
|**chunkOverlap** | **Integer** |  |  [optional] |
|**skipEmbeddingGeneration** | **Boolean** |  |  [optional] |
|**overwriteFileId** | **Integer** |  |  [optional] |
|**embeddingModel** | **EmbeddingGeneratorsNullable** |  |  [optional] |
|**generateSparseVectors** | **Boolean** |  |  [optional] |
|**coldStorageParams** | [**ColdStorageProps**](ColdStorageProps.md) |  |  [optional] |
|**generateChunksOnly** | **Boolean** | If this flag is enabled, the file will be chunked and stored with Carbon,         but no embeddings will be generated. This overrides the skip_embedding_generation flag. |  [optional] |
|**storeFileOnly** | **Boolean** | If this flag is enabled, the file will be stored with Carbon, but no processing will be done. |  [optional] |




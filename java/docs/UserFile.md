

# UserFile


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**tags** | **Object** |  |  [optional] |
|**id** | **Integer** |  |  |
|**source** | **DataSourceType** |  |  |
|**organizationId** | **Integer** |  |  |
|**organizationUserId** | **Integer** |  |  |
|**organizationSuppliedUserId** | **String** |  |  |
|**organizationUserDataSourceId** | **Integer** |  |  [optional] |
|**externalFileId** | **String** |  |  |
|**externalUrl** | **String** |  |  [optional] |
|**syncStatus** | **ExternalFileSyncStatuses** |  |  |
|**syncErrorMessage** | **String** |  |  [optional] |
|**lastSync** | **OffsetDateTime** |  |  [optional] |
|**fileStatistics** | [**FileStatisticsNullable**](FileStatisticsNullable.md) |  |  [optional] |
|**fileMetadata** | **Object** |  |  [optional] |
|**embeddingProperties** | [**Map&lt;String, EmbeddingProperties&gt;**](EmbeddingProperties.md) |  |  [optional] |
|**chunkSize** | **Integer** |  |  [optional] |
|**chunkOverlap** | **Integer** |  |  [optional] |
|**chunkProperties** | [**ChunkPropertiesNullable**](ChunkPropertiesNullable.md) |  |  [optional] |
|**ocrProperties** | **Object** |  |  [optional] |
|**ocrJobStartedAt** | **OffsetDateTime** |  |  [optional] |
|**name** | **String** |  |  [optional] |
|**parentId** | **Integer** |  |  [optional] |
|**enableAutoSync** | **Boolean** |  |  [optional] |
|**presignedUrl** | **String** |  |  [optional] |
|**parsedTextUrl** | **String** |  |  [optional] |
|**additionalPresignedUrls** | **Object** |  |  [optional] |
|**skipEmbeddingGeneration** | **Boolean** |  |  |
|**sourceCreatedAt** | **OffsetDateTime** |  |  [optional] |
|**generateSparseVectors** | **Boolean** |  |  [optional] |
|**requestId** | **String** |  |  [optional] |
|**uploadId** | **String** |  |  [optional] |
|**syncProperties** | **Object** |  |  [optional] |
|**messagesMetadata** | **Object** |  |  [optional] |
|**fileContentsDeleted** | **Boolean** |  |  [optional] |
|**supportsColdStorage** | **Boolean** |  |  |
|**hotStorageTimeToLive** | **Integer** |  |  [optional] |
|**embeddingStorageStatus** | **EmbeddingStorageStatus** |  |  |
|**createdAt** | **OffsetDateTime** |  |  |
|**updatedAt** | **OffsetDateTime** |  |  |




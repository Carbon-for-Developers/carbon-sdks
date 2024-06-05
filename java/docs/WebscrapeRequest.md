

# WebscrapeRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**tags** | **Map&lt;String, Object&gt;** |  |  [optional] |
|**url** | **String** |  |  |
|**recursionDepth** | **Integer** |  |  [optional] |
|**maxPagesToScrape** | **Integer** |  |  [optional] |
|**chunkSize** | **Integer** |  |  [optional] |
|**chunkOverlap** | **Integer** |  |  [optional] |
|**skipEmbeddingGeneration** | **Boolean** |  |  [optional] |
|**enableAutoSync** | **Boolean** |  |  [optional] |
|**generateSparseVectors** | **Boolean** |  |  [optional] |
|**prependFilenameToChunks** | **Boolean** |  |  [optional] |
|**htmlTagsToSkip** | **List&lt;String&gt;** |  |  [optional] |
|**cssClassesToSkip** | **List&lt;String&gt;** |  |  [optional] |
|**cssSelectorsToSkip** | **List&lt;String&gt;** |  |  [optional] |
|**embeddingModel** | **EmbeddingGenerators** |  |  [optional] |
|**urlPathsToInclude** | **List&lt;String&gt;** | URL subpaths or directories that you want to include. For example if you want to only include         URLs that start with /questions in stackoverflow.com, you will add /questions/ in this input |  [optional] |






# SitemapScrapeRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**tags** | **Map&lt;String, Object&gt;** |  |  [optional] |
|**url** | **String** |  |  |
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
|**urlPathsToExclude** | **List&lt;String&gt;** | URL subpaths or directories that you want to exclude. For example if you want to exclude         URLs that start with /questions in stackoverflow.com, you will add /questions/ in this input |  [optional] |
|**urlsToScrape** | **List&lt;String&gt;** | You can submit a subset of URLs from the sitemap that should be scraped. To get the list of URLs,           you can check out /process_sitemap endpoint. If left empty, all URLs from the sitemap will be scraped. |  [optional] |




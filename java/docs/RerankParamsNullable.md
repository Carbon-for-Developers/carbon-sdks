

# RerankParamsNullable

Parameters for reranking the chunks using a specified model. This field accepts an object with details of the reranking model to be used; either 'Cohere' or 'Jina'. If provided, the specified reranking model will reorder the retrieved chunks based on their relevance to the query.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**model** | **String** | The name of the model to use for reranking. Options: [&#39;COHERE_RERANK_MULTILINGUAL_V3&#39;, &#39;JINA_RERANKER_MULTILINGUAL_V2&#39;, &#39;PONGO_RERANKER&#39;]. |  |




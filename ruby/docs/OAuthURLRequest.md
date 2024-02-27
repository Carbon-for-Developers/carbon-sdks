# Carbon::OAuthURLRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **tags** | **Object** |  | [optional] |
| **scope** | **String** |  | [optional] |
| **service** | [**DataSourceType**](DataSourceType.md) |  |  |
| **chunk_size** | **Integer** |  | [optional][default to 1500] |
| **chunk_overlap** | **Integer** |  | [optional][default to 20] |
| **skip_embedding_generation** | **Boolean** |  | [optional][default to false] |
| **embedding_model** | [**EmbeddingGeneratorsNullable**](EmbeddingGeneratorsNullable.md) |  | [optional][default to &#39;OPENAI&#39;] |
| **zendesk_subdomain** | **String** |  | [optional] |
| **microsoft_tenant** | **String** |  | [optional] |
| **sharepoint_site_name** | **String** |  | [optional] |
| **confluence_subdomain** | **String** |  | [optional] |
| **generate_sparse_vectors** | **Boolean** |  | [optional][default to false] |
| **prepend_filename_to_chunks** | **Boolean** |  | [optional][default to false] |
| **max_items_per_chunk** | **Integer** |  | [optional] |

## Example

```ruby
require 'carbon'

instance = Carbon::OAuthURLRequest.new(
  tags: null,
  scope: null,
  service: null,
  chunk_size: 1500,
  chunk_overlap: 20,
  skip_embedding_generation: false,
  embedding_model: OPENAI,
  zendesk_subdomain: null,
  microsoft_tenant: null,
  sharepoint_site_name: null,
  confluence_subdomain: null,
  generate_sparse_vectors: false,
  prepend_filename_to_chunks: false,
  max_items_per_chunk: null
)
```


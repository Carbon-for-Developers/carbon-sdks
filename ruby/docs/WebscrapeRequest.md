# Carbon::WebscrapeRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **tags** | [**Hash&lt;String, Tags1&gt;**](Tags1.md) |  | [optional] |
| **url** | **String** |  |  |
| **recursion_depth** | **Integer** |  | [optional][default to 3] |
| **max_pages_to_scrape** | **Integer** |  | [optional][default to 100] |
| **chunk_size** | **Integer** |  | [optional][default to 1500] |
| **chunk_overlap** | **Integer** |  | [optional][default to 20] |
| **skip_embedding_generation** | **Boolean** |  | [optional][default to false] |
| **enable_auto_sync** | **Boolean** |  | [optional][default to false] |
| **generate_sparse_vectors** | **Boolean** |  | [optional][default to false] |
| **prepend_filename_to_chunks** | **Boolean** |  | [optional][default to false] |
| **html_tags_to_skip** | **Array&lt;String&gt;** |  | [optional] |
| **css_classes_to_skip** | **Array&lt;String&gt;** |  | [optional] |
| **css_selectors_to_skip** | **Array&lt;String&gt;** |  | [optional] |

## Example

```ruby
require 'carbon'

instance = Carbon::WebscrapeRequest.new(
  tags: {},
  url: null,
  recursion_depth: 3,
  max_pages_to_scrape: 100,
  chunk_size: 1500,
  chunk_overlap: 20,
  skip_embedding_generation: false,
  enable_auto_sync: false,
  generate_sparse_vectors: false,
  prepend_filename_to_chunks: false,
  html_tags_to_skip: [],
  css_classes_to_skip: [],
  css_selectors_to_skip: []
)
```


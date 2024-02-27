# Carbon::UploadFileFromUrlInput

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **url** | **String** |  |  |
| **file_name** | **String** |  | [optional] |
| **chunk_size** | **Integer** |  | [optional] |
| **chunk_overlap** | **Integer** |  | [optional] |
| **skip_embedding_generation** | **Boolean** |  | [optional][default to false] |
| **set_page_as_boundary** | **Boolean** |  | [optional][default to false] |
| **embedding_model** | [**EmbeddingGenerators**](EmbeddingGenerators.md) |  | [optional][default to &#39;OPENAI&#39;] |
| **generate_sparse_vectors** | **Boolean** |  | [optional][default to false] |
| **use_textract** | **Boolean** |  | [optional][default to false] |
| **prepend_filename_to_chunks** | **Boolean** |  | [optional][default to false] |
| **max_items_per_chunk** | **Integer** |  | [optional] |

## Example

```ruby
require 'carbon'

instance = Carbon::UploadFileFromUrlInput.new(
  url: null,
  file_name: null,
  chunk_size: null,
  chunk_overlap: null,
  skip_embedding_generation: false,
  set_page_as_boundary: false,
  embedding_model: null,
  generate_sparse_vectors: false,
  use_textract: false,
  prepend_filename_to_chunks: false,
  max_items_per_chunk: null
)
```


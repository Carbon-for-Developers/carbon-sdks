# Carbon::SingleChunksAndEmbeddingsUploadInput

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **file_id** | **Integer** |  |  |
| **chunk_size** | **Integer** |  | [optional] |
| **chunk_overlap** | **Integer** |  | [optional] |
| **chunks_and_embeddings** | [**Array&lt;ChunksAndEmbeddings&gt;**](ChunksAndEmbeddings.md) |  |  |

## Example

```ruby
require 'carbon'

instance = Carbon::SingleChunksAndEmbeddingsUploadInput.new(
  file_id: null,
  chunk_size: null,
  chunk_overlap: null,
  chunks_and_embeddings: null
)
```


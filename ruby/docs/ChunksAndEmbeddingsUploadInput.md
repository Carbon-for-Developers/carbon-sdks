# Carbon::ChunksAndEmbeddingsUploadInput

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **embedding_model** | [**EmbeddingGenerators**](EmbeddingGenerators.md) |  |  |
| **chunks_and_embeddings** | [**Array&lt;SingleChunksAndEmbeddingsUploadInput&gt;**](SingleChunksAndEmbeddingsUploadInput.md) |  |  |
| **overwrite_existing** | **Boolean** |  | [optional][default to false] |

## Example

```ruby
require 'carbon'

instance = Carbon::ChunksAndEmbeddingsUploadInput.new(
  embedding_model: null,
  chunks_and_embeddings: null,
  overwrite_existing: false
)
```


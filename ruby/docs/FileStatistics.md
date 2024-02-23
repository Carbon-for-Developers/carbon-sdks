# Carbon::FileStatistics

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **file_format** | [**FileFormatsNullable**](FileFormatsNullable.md) |  |  |
| **file_size** | **Integer** |  |  |
| **num_characters** | **Integer** |  |  |
| **num_tokens** | **Integer** |  |  |
| **num_embeddings** | **Integer** |  |  |

## Example

```ruby
require 'carbon'

instance = Carbon::FileStatistics.new(
  file_format: null,
  file_size: null,
  num_characters: null,
  num_tokens: null,
  num_embeddings: null
)
```


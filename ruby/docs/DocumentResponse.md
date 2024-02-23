# Carbon::DocumentResponse

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **tags** | [**Hash&lt;String, Tags&gt;**](Tags.md) |  |  |
| **content** | **String** |  |  |
| **file_id** | **Integer** |  |  |
| **source** | **String** |  |  |
| **source_url** | **String** |  |  |
| **source_type** | [**DataSourceTypeNullable**](DataSourceTypeNullable.md) |  |  |
| **presigned_url** | **String** |  |  |
| **vector** | **Array&lt;Float&gt;** |  |  |
| **score** | **Float** |  |  |
| **rank** | [**RankProperty**](RankProperty.md) |  |  |
| **content_metadata** | **Object** |  |  |

## Example

```ruby
require 'carbon'

instance = Carbon::DocumentResponse.new(
  tags: null,
  content: null,
  file_id: null,
  source: null,
  source_url: null,
  source_type: null,
  presigned_url: null,
  vector: null,
  score: null,
  rank: null,
  content_metadata: null
)
```


# Carbon::UserResponse

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |
| **organization_id** | **Integer** |  |  |
| **organization_supplied_user_id** | **String** |  |  |
| **created_at** | **Time** |  |  |
| **updated_at** | **Time** |  |  |
| **deleted_at** | **Time** |  |  |
| **num_files_synced** | **Integer** |  |  |
| **num_characters_synced** | **Integer** |  |  |
| **num_tokens_synced** | **Integer** |  |  |
| **unique_file_tags** | **Array&lt;Object&gt;** |  |  |
| **enabled_features** | **Object** |  |  |

## Example

```ruby
require 'carbon'

instance = Carbon::UserResponse.new(
  id: null,
  organization_id: null,
  organization_supplied_user_id: null,
  created_at: null,
  updated_at: null,
  deleted_at: null,
  num_files_synced: null,
  num_characters_synced: null,
  num_tokens_synced: null,
  unique_file_tags: null,
  enabled_features: null
)
```


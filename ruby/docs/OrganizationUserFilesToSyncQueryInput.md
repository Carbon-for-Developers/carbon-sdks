# Carbon::OrganizationUserFilesToSyncQueryInput

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **pagination** | [**Pagination**](Pagination.md) |  | [optional] |
| **order_by** | [**OrganizationUserFilesToSyncOrderByTypes**](OrganizationUserFilesToSyncOrderByTypes.md) |  | [optional] |
| **order_dir** | [**OrderDir**](OrderDir.md) |  | [optional] |
| **filters** | [**OrganizationUserFilesToSyncFilters**](OrganizationUserFilesToSyncFilters.md) |  | [optional] |
| **include_raw_file** | **Boolean** |  | [optional] |
| **include_parsed_text_file** | **Boolean** |  | [optional] |
| **include_additional_files** | **Boolean** |  | [optional] |

## Example

```ruby
require 'carbon'

instance = Carbon::OrganizationUserFilesToSyncQueryInput.new(
  pagination: null,
  order_by: null,
  order_dir: null,
  filters: null,
  include_raw_file: null,
  include_parsed_text_file: null,
  include_additional_files: null
)
```


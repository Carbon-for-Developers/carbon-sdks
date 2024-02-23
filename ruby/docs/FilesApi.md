# Carbon::FilesApi

All URIs are relative to *https://api.carbon.ai*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**create_user_file_tags**](FilesApi.md#create_user_file_tags) | **POST** /create_user_file_tags | Create File Tags |
| [**delete**](FilesApi.md#delete) | **DELETE** /deletefile/{file_id} | Delete File Endpoint |
| [**delete_file_tags**](FilesApi.md#delete_file_tags) | **POST** /delete_user_file_tags | Delete File Tags |
| [**delete_many**](FilesApi.md#delete_many) | **POST** /delete_files | Delete Files Endpoint |
| [**get_parsed_file**](FilesApi.md#get_parsed_file) | **GET** /parsed_file/{file_id} | Parsed File |
| [**get_raw_file**](FilesApi.md#get_raw_file) | **GET** /raw_file/{file_id} | Raw File |
| [**query_user_files**](FilesApi.md#query_user_files) | **POST** /user_files_v2 | User Files V2 |
| [**query_user_files_deprecated**](FilesApi.md#query_user_files_deprecated) | **POST** /user_files | User Files |
| [**resync**](FilesApi.md#resync) | **POST** /resync_file | Resync File |
| [**upload**](FilesApi.md#upload) | **POST** /uploadfile | Create Upload File |
| [**upload_from_url**](FilesApi.md#upload_from_url) | **POST** /upload_file_from_url | Create Upload File From Url |
| [**upload_text**](FilesApi.md#upload_text) | **POST** /upload_text | Create Raw Text |

## create_user_file_tags

Create File Tags

A tag is a key-value pair that can be added to a file. This pair can then be used for searches (e.g. embedding searches) in order to narrow down the scope of the search. A file can have any number of tags. The following are reserved keys that cannot be used: - db_embedding_id - organization_id - user_id - organization_user_file_id  Carbon currently supports two data types for tag values - `string` and `list<string>`. Keys can only be `string`. If values other than `string` and `list<string>` are used, they're automatically converted to strings (e.g. 4 will become \"4\").

### Example

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'
Carbon.api_key = 'YOUR API KEY'
Carbon.customer_id = 'YOUR API KEY'

tags = {
        "key": "string_example",
    }
organization_user_file_id = 1

begin
  # Create File Tags
  result = Carbon::Files.create_user_file_tags(
                                                  tags: tags,
                                                  organization_user_file_id: organization_user_file_id,
                                                )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Files.create_user_file_tags: #{e}"
end
```

#### Using the create_user_file_tags_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
tags = {
        "key": "string_example",
    }
organization_user_file_id = 1

begin
  # Create File Tags
  data, status_code, headers, response = Carbon::Files.create_user_file_tags_with_http_info(
                                                                                               tags: tags,
                                                                                               organization_user_file_id: organization_user_file_id,
                                                                                             )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => UserFile
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Files.create_user_file_tags: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **organization_user_file_tag_create** | [**OrganizationUserFileTagCreate**](OrganizationUserFileTagCreate.md) |  |  |

### Return type

[**UserFile**](UserFile.md)

## delete

Delete File Endpoint

### Example

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'
Carbon.api_key = 'YOUR API KEY'
Carbon.customer_id = 'YOUR API KEY'

file_id = 1

begin
  # Delete File Endpoint
  result = Carbon::Files.delete(
                                   file_id: file_id,
                                 )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Files.delete: #{e}"
end
```

#### Using the delete_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
file_id = 1

begin
  # Delete File Endpoint
  data, status_code, headers, response = Carbon::Files.delete_with_http_info(
                                                                                file_id: file_id,
                                                                              )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => GenericSuccessResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Files.delete: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **file_id** | **Integer** |  |  |

### Return type

[**GenericSuccessResponse**](GenericSuccessResponse.md)

## delete_file_tags

Delete File Tags

### Example

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'
Carbon.api_key = 'YOUR API KEY'
Carbon.customer_id = 'YOUR API KEY'

tags = [
        "string_example"
    ]
organization_user_file_id = 1

begin
  # Delete File Tags
  result = Carbon::Files.delete_file_tags(
                                             tags: tags,
                                             organization_user_file_id: organization_user_file_id,
                                           )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Files.delete_file_tags: #{e}"
end
```

#### Using the delete_file_tags_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
tags = [
        "string_example"
    ]
organization_user_file_id = 1

begin
  # Delete File Tags
  data, status_code, headers, response = Carbon::Files.delete_file_tags_with_http_info(
                                                                                          tags: tags,
                                                                                          organization_user_file_id: organization_user_file_id,
                                                                                        )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => UserFile
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Files.delete_file_tags: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **organization_user_file_tags_remove** | [**OrganizationUserFileTagsRemove**](OrganizationUserFileTagsRemove.md) |  |  |

### Return type

[**UserFile**](UserFile.md)

## delete_many

Delete Files Endpoint

### Example

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'
Carbon.api_key = 'YOUR API KEY'
Carbon.customer_id = 'YOUR API KEY'

file_ids = [
        1
    ]
sync_statuses = [
        "string_example"
    ]
delete_non_synced_only = False

begin
  # Delete Files Endpoint
  result = Carbon::Files.delete_many(
                                        file_ids: file_ids,
                                        sync_statuses: sync_statuses,
                                        delete_non_synced_only: delete_non_synced_only,
                                      )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Files.delete_many: #{e}"
end
```

#### Using the delete_many_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
file_ids = [
        1
    ]
sync_statuses = [
        "string_example"
    ]
delete_non_synced_only = False

begin
  # Delete Files Endpoint
  data, status_code, headers, response = Carbon::Files.delete_many_with_http_info(
                                                                                     file_ids: file_ids,
                                                                                     sync_statuses: sync_statuses,
                                                                                     delete_non_synced_only: delete_non_synced_only,
                                                                                   )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => GenericSuccessResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Files.delete_many: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **delete_files_query_input** | [**DeleteFilesQueryInput**](DeleteFilesQueryInput.md) |  |  |

### Return type

[**GenericSuccessResponse**](GenericSuccessResponse.md)

## get_parsed_file

Parsed File

This route is deprecated. Use `/user_files_v2` instead.

### Example

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'
Carbon.api_key = 'YOUR API KEY'
Carbon.customer_id = 'YOUR API KEY'

file_id = 1

begin
  # Parsed File
  result = Carbon::Files.get_parsed_file(
                                            file_id: file_id,
                                          )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Files.get_parsed_file: #{e}"
end
```

#### Using the get_parsed_file_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
file_id = 1

begin
  # Parsed File
  data, status_code, headers, response = Carbon::Files.get_parsed_file_with_http_info(
                                                                                         file_id: file_id,
                                                                                       )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => PresignedURLResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Files.get_parsed_file: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **file_id** | **Integer** |  |  |

### Return type

[**PresignedURLResponse**](PresignedURLResponse.md)

## get_raw_file

Raw File

This route is deprecated. Use `/user_files_v2` instead.

### Example

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'
Carbon.api_key = 'YOUR API KEY'
Carbon.customer_id = 'YOUR API KEY'

file_id = 1

begin
  # Raw File
  result = Carbon::Files.get_raw_file(
                                         file_id: file_id,
                                       )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Files.get_raw_file: #{e}"
end
```

#### Using the get_raw_file_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
file_id = 1

begin
  # Raw File
  data, status_code, headers, response = Carbon::Files.get_raw_file_with_http_info(
                                                                                      file_id: file_id,
                                                                                    )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => PresignedURLResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Files.get_raw_file: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **file_id** | **Integer** |  |  |

### Return type

[**PresignedURLResponse**](PresignedURLResponse.md)

## query_user_files

User Files V2

For pre-filtering documents, using `tags_v2` is preferred to using `tags` (which is now deprecated). If both `tags_v2` and `tags` are specified, `tags` is ignored. `tags_v2` enables building complex filters through the use of \"AND\", \"OR\", and negation logic. Take the below input as an example: ```json {     \"OR\": [         {             \"key\": \"subject\",             \"value\": \"holy-bible\",             \"negate\": false         },         {             \"key\": \"person-of-interest\",             \"value\": \"jesus christ\",             \"negate\": false         },         {             \"key\": \"genre\",             \"value\": \"religion\",             \"negate\": true         }         {             \"AND\": [                 {                     \"key\": \"subject\",                     \"value\": \"tao-te-ching\",                     \"negate\": false                 },                 {                     \"key\": \"author\",                     \"value\": \"lao-tzu\",                     \"negate\": false                 }             ]         }     ] } ``` In this case, files will be filtered such that: 1. \"subject\" = \"holy-bible\" OR 2. \"person-of-interest\" = \"jesus christ\" OR 3. \"genre\" != \"religion\" OR 4. \"subject\" = \"tao-te-ching\" AND \"author\" = \"lao-tzu\"  Note that the top level of the query must be either an \"OR\" or \"AND\" array. Currently, nesting is limited to 3. For tag blocks (those with \"key\", \"value\", and \"negate\" keys), the following typing rules apply: 1. \"key\" isn't optional and must be a `string` 2. \"value\" isn't optional and can be `any` or list[`any`] 3. \"negate\" is optional and must be `true` or `false`. If present and `true`, then the filter block is negated in the resulting query. It is `false` by default.

### Example

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'
Carbon.api_key = 'YOUR API KEY'
Carbon.customer_id = 'YOUR API KEY'

pagination = {
        "limit" => 10,
        "offset" => 0,
    }
order_by = "created_at"
order_dir = "desc"
filters = {
    }
include_raw_file = True
include_parsed_text_file = True
include_additional_files = True

begin
  # User Files V2
  result = Carbon::Files.query_user_files(
                                             pagination: pagination,
                                             order_by: order_by,
                                             order_dir: order_dir,
                                             filters: filters,
                                             include_raw_file: include_raw_file,
                                             include_parsed_text_file: include_parsed_text_file,
                                             include_additional_files: include_additional_files,
                                           )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Files.query_user_files: #{e}"
end
```

#### Using the query_user_files_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
pagination = {
        "limit" => 10,
        "offset" => 0,
    }
order_by = "created_at"
order_dir = "desc"
filters = {
    }
include_raw_file = True
include_parsed_text_file = True
include_additional_files = True

begin
  # User Files V2
  data, status_code, headers, response = Carbon::Files.query_user_files_with_http_info(
                                                                                          pagination: pagination,
                                                                                          order_by: order_by,
                                                                                          order_dir: order_dir,
                                                                                          filters: filters,
                                                                                          include_raw_file: include_raw_file,
                                                                                          include_parsed_text_file: include_parsed_text_file,
                                                                                          include_additional_files: include_additional_files,
                                                                                        )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => UserFilesV2
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Files.query_user_files: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **organization_user_files_to_sync_query_input** | [**OrganizationUserFilesToSyncQueryInput**](OrganizationUserFilesToSyncQueryInput.md) |  |  |

### Return type

[**UserFilesV2**](UserFilesV2.md)

## query_user_files_deprecated

User Files

This route is deprecated. Use `/user_files_v2` instead.

### Example

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'
Carbon.api_key = 'YOUR API KEY'
Carbon.customer_id = 'YOUR API KEY'

pagination = {
        "limit" => 10,
        "offset" => 0,
    }
order_by = "created_at"
order_dir = "desc"
filters = {
    }
include_raw_file = True
include_parsed_text_file = True
include_additional_files = True

begin
  # User Files
  result = Carbon::Files.query_user_files_deprecated(
                                                        pagination: pagination,
                                                        order_by: order_by,
                                                        order_dir: order_dir,
                                                        filters: filters,
                                                        include_raw_file: include_raw_file,
                                                        include_parsed_text_file: include_parsed_text_file,
                                                        include_additional_files: include_additional_files,
                                                      )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Files.query_user_files_deprecated: #{e}"
end
```

#### Using the query_user_files_deprecated_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
pagination = {
        "limit" => 10,
        "offset" => 0,
    }
order_by = "created_at"
order_dir = "desc"
filters = {
    }
include_raw_file = True
include_parsed_text_file = True
include_additional_files = True

begin
  # User Files
  data, status_code, headers, response = Carbon::Files.query_user_files_deprecated_with_http_info(
                                                                                                     pagination: pagination,
                                                                                                     order_by: order_by,
                                                                                                     order_dir: order_dir,
                                                                                                     filters: filters,
                                                                                                     include_raw_file: include_raw_file,
                                                                                                     include_parsed_text_file: include_parsed_text_file,
                                                                                                     include_additional_files: include_additional_files,
                                                                                                   )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => Array<UserFile>
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Files.query_user_files_deprecated: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **organization_user_files_to_sync_query_input** | [**OrganizationUserFilesToSyncQueryInput**](OrganizationUserFilesToSyncQueryInput.md) |  |  |

### Return type

[**Array&lt;UserFile&gt;**](UserFile.md)

## resync

Resync File

### Example

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'
Carbon.api_key = 'YOUR API KEY'
Carbon.customer_id = 'YOUR API KEY'

file_id = 1
chunk_size = 1
chunk_overlap = 1

begin
  # Resync File
  result = Carbon::Files.resync(
                                   file_id: file_id,
                                   chunk_size: chunk_size,
                                   chunk_overlap: chunk_overlap,
                                 )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Files.resync: #{e}"
end
```

#### Using the resync_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
file_id = 1
chunk_size = 1
chunk_overlap = 1

begin
  # Resync File
  data, status_code, headers, response = Carbon::Files.resync_with_http_info(
                                                                                file_id: file_id,
                                                                                chunk_size: chunk_size,
                                                                                chunk_overlap: chunk_overlap,
                                                                              )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => UserFile
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Files.resync: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **resync_file_query_input** | [**ResyncFileQueryInput**](ResyncFileQueryInput.md) |  |  |

### Return type

[**UserFile**](UserFile.md)

## upload

Create Upload File

This endpoint is used to directly upload local files to Carbon. The `POST` request should be a multipart form request. Note that the `set_page_as_boundary` query parameter is applicable only to PDFs for now. When this value is set, PDF chunks are at most one page long. Additional information can be retrieved for each chunk, however, namely the coordinates of the bounding box around the chunk (this can be used for things like text highlighting). Following is a description of all possible query parameters: - `chunk_size`: the chunk size (in tokens) applied when splitting the document - `chunk_overlap`: the chunk overlap (in tokens) applied when splitting the document - `skip_embedding_generation`: whether or not to skip the generation of chunks and embeddings - `set_page_as_boundary`: described above - `embedding_model`: the model used to generate embeddings for the document chunks - `use_ocr`: whether or not to use OCR as a preprocessing step prior to generating chunks (only valid for PDFs currently) - `generate_sparse_vectors`: whether or not to generate sparse vectors for the file. Required for hybrid search. - `prepend_filename_to_chunks`: whether or not to prepend the filename to the chunk text   Carbon supports multiple models for use in generating embeddings for files. For images, we support Vertex AI's multimodal model; for text, we support OpenAI's `text-embedding-ada-002` and Cohere's embed-multilingual-v3.0. The model can be specified via the `embedding_model` parameter (in the POST body for `/embeddings`, and a query  parameter in `/uploadfile`). If no model is supplied, the `text-embedding-ada-002` is used by default. When performing embedding queries, embeddings from files that used the specified model will be considered in the query. For example, if files A and B have embeddings generated with `OPENAI`, and files C and D have embeddings generated with `COHERE_MULTILINGUAL_V3`, then by default, queries will only consider files A and B. If `COHERE_MULTILINGUAL_V3` is specified as the `embedding_model` in `/embeddings`, then only files C and D will be considered. Make sure that the set of all files you want considered for a query have embeddings generated via the same model. For now, **do not** set `VERTEX_MULTIMODAL` as an `embedding_model`. This model is used automatically by Carbon when it detects an image file.

### Example

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'
Carbon.api_key = 'YOUR API KEY'
Carbon.customer_id = 'YOUR API KEY'

file = open('/path/to/file', 'rb')
chunk_size = 1
chunk_overlap = 1
skip_embedding_generation = False
set_page_as_boundary = False
embedding_model = "OPENAI"
use_ocr = False
generate_sparse_vectors = False
prepend_filename_to_chunks = False
max_items_per_chunk = 1

begin
  # Create Upload File
  result = Carbon::Files.upload(
                                   file: file,
                                   chunk_size: chunk_size,
                                   chunk_overlap: chunk_overlap,
                                   skip_embedding_generation: skip_embedding_generation,
                                   set_page_as_boundary: set_page_as_boundary,
                                   embedding_model: embedding_model,
                                   use_ocr: use_ocr,
                                   generate_sparse_vectors: generate_sparse_vectors,
                                   prepend_filename_to_chunks: prepend_filename_to_chunks,
                                   max_items_per_chunk: max_items_per_chunk,
                                 )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Files.upload: #{e}"
end
```

#### Using the upload_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
file = open('/path/to/file', 'rb')
chunk_size = 1
chunk_overlap = 1
skip_embedding_generation = False
set_page_as_boundary = False
embedding_model = "OPENAI"
use_ocr = False
generate_sparse_vectors = False
prepend_filename_to_chunks = False
max_items_per_chunk = 1

begin
  # Create Upload File
  data, status_code, headers, response = Carbon::Files.upload_with_http_info(
                                                                                file: file,
                                                                                chunk_size: chunk_size,
                                                                                chunk_overlap: chunk_overlap,
                                                                                skip_embedding_generation: skip_embedding_generation,
                                                                                set_page_as_boundary: set_page_as_boundary,
                                                                                embedding_model: embedding_model,
                                                                                use_ocr: use_ocr,
                                                                                generate_sparse_vectors: generate_sparse_vectors,
                                                                                prepend_filename_to_chunks: prepend_filename_to_chunks,
                                                                                max_items_per_chunk: max_items_per_chunk,
                                                                              )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => UserFile
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Files.upload: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **file** | **File** |  |  |
| **body_create_upload_file_uploadfile_post** | [**BodyCreateUploadFileUploadfilePost**](BodyCreateUploadFileUploadfilePost.md) |  |  |
| **chunk_size** | **Integer** | Chunk size in tiktoken tokens to be used when processing file. | [optional] |
| **chunk_overlap** | **Integer** | Chunk overlap in tiktoken tokens to be used when processing file. | [optional] |
| **skip_embedding_generation** | **Boolean** | Flag to control whether or not embeddings should be generated and stored             when processing file. | [optional][default to false] |
| **set_page_as_boundary** | **Boolean** | Flag to control whether or not to set the a page&#39;s worth of content as the maximum             amount of content that can appear in a chunk. Only valid for PDFs. See description route description for             more information. | [optional][default to false] |
| **embedding_model** | [**TextEmbeddingGenerators**](.md) | Embedding model that will be used to embed file chunks. | [optional] |
| **use_ocr** | **Boolean** | Whether or not to use OCR when processing files. Only valid for PDFs. Useful for documents with             tables, images, and/or scanned text. | [optional][default to false] |
| **generate_sparse_vectors** | **Boolean** | Whether or not to generate sparse vectors for the file. This is *required* for the file to be a             candidate for hybrid search. | [optional][default to false] |
| **prepend_filename_to_chunks** | **Boolean** | Whether or not to prepend the file&#39;s name to chunks. | [optional][default to false] |
| **max_items_per_chunk** | **Integer** | Number of objects per chunk. For json files only. | [optional] |

### Return type

[**UserFile**](UserFile.md)

## upload_from_url

Create Upload File From Url

### Example

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'
Carbon.api_key = 'YOUR API KEY'
Carbon.customer_id = 'YOUR API KEY'

url = "string_example"
file_name = "string_example"
chunk_size = 1
chunk_overlap = 1
skip_embedding_generation = False
set_page_as_boundary = False
embedding_model = "OPENAI"
generate_sparse_vectors = False
use_textract = False
prepend_filename_to_chunks = False
max_items_per_chunk = 1

begin
  # Create Upload File From Url
  result = Carbon::Files.upload_from_url(
                                            url: url,
                                            file_name: file_name,
                                            chunk_size: chunk_size,
                                            chunk_overlap: chunk_overlap,
                                            skip_embedding_generation: skip_embedding_generation,
                                            set_page_as_boundary: set_page_as_boundary,
                                            embedding_model: embedding_model,
                                            generate_sparse_vectors: generate_sparse_vectors,
                                            use_textract: use_textract,
                                            prepend_filename_to_chunks: prepend_filename_to_chunks,
                                            max_items_per_chunk: max_items_per_chunk,
                                          )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Files.upload_from_url: #{e}"
end
```

#### Using the upload_from_url_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
url = "string_example"
file_name = "string_example"
chunk_size = 1
chunk_overlap = 1
skip_embedding_generation = False
set_page_as_boundary = False
embedding_model = "OPENAI"
generate_sparse_vectors = False
use_textract = False
prepend_filename_to_chunks = False
max_items_per_chunk = 1

begin
  # Create Upload File From Url
  data, status_code, headers, response = Carbon::Files.upload_from_url_with_http_info(
                                                                                         url: url,
                                                                                         file_name: file_name,
                                                                                         chunk_size: chunk_size,
                                                                                         chunk_overlap: chunk_overlap,
                                                                                         skip_embedding_generation: skip_embedding_generation,
                                                                                         set_page_as_boundary: set_page_as_boundary,
                                                                                         embedding_model: embedding_model,
                                                                                         generate_sparse_vectors: generate_sparse_vectors,
                                                                                         use_textract: use_textract,
                                                                                         prepend_filename_to_chunks: prepend_filename_to_chunks,
                                                                                         max_items_per_chunk: max_items_per_chunk,
                                                                                       )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => UserFile
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Files.upload_from_url: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **upload_file_from_url_input** | [**UploadFileFromUrlInput**](UploadFileFromUrlInput.md) |  |  |

### Return type

[**UserFile**](UserFile.md)

## upload_text

Create Raw Text

Carbon supports multiple models for use in generating embeddings for files. For images, we support Vertex AI's multimodal model; for text, we support OpenAI's `text-embedding-ada-002` and Cohere's embed-multilingual-v3.0. The model can be specified via the `embedding_model` parameter (in the POST body for `/embeddings`, and a query  parameter in `/uploadfile`). If no model is supplied, the `text-embedding-ada-002` is used by default. When performing embedding queries, embeddings from files that used the specified model will be considered in the query. For example, if files A and B have embeddings generated with `OPENAI`, and files C and D have embeddings generated with `COHERE_MULTILINGUAL_V3`, then by default, queries will only consider files A and B. If `COHERE_MULTILINGUAL_V3` is specified as the `embedding_model` in `/embeddings`, then only files C and D will be considered. Make sure that the set of all files you want considered for a query have embeddings generated via the same model. For now, **do not** set `VERTEX_MULTIMODAL` as an `embedding_model`. This model is used automatically by Carbon when it detects an image file.

### Example

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'
Carbon.api_key = 'YOUR API KEY'
Carbon.customer_id = 'YOUR API KEY'

contents = "string_example"
name = "string_example"
chunk_size = 1
chunk_overlap = 1
skip_embedding_generation = False
overwrite_file_id = 1
embedding_model = "OPENAI"
generate_sparse_vectors = False

begin
  # Create Raw Text
  result = Carbon::Files.upload_text(
                                        contents: contents,
                                        name: name,
                                        chunk_size: chunk_size,
                                        chunk_overlap: chunk_overlap,
                                        skip_embedding_generation: skip_embedding_generation,
                                        overwrite_file_id: overwrite_file_id,
                                        embedding_model: embedding_model,
                                        generate_sparse_vectors: generate_sparse_vectors,
                                      )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Files.upload_text: #{e}"
end
```

#### Using the upload_text_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
contents = "string_example"
name = "string_example"
chunk_size = 1
chunk_overlap = 1
skip_embedding_generation = False
overwrite_file_id = 1
embedding_model = "OPENAI"
generate_sparse_vectors = False

begin
  # Create Raw Text
  data, status_code, headers, response = Carbon::Files.upload_text_with_http_info(
                                                                                     contents: contents,
                                                                                     name: name,
                                                                                     chunk_size: chunk_size,
                                                                                     chunk_overlap: chunk_overlap,
                                                                                     skip_embedding_generation: skip_embedding_generation,
                                                                                     overwrite_file_id: overwrite_file_id,
                                                                                     embedding_model: embedding_model,
                                                                                     generate_sparse_vectors: generate_sparse_vectors,
                                                                                   )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => UserFile
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Files.upload_text: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **raw_text_input** | [**RawTextInput**](RawTextInput.md) |  |  |

### Return type

[**UserFile**](UserFile.md)


# Carbon::EmbeddingsApi

All URIs are relative to *https://api.carbon.ai*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**get_documents**](EmbeddingsApi.md#get_documents) | **POST** /embeddings | Embeddings |
| [**get_embeddings_and_chunks**](EmbeddingsApi.md#get_embeddings_and_chunks) | **POST** /text_chunks | Retrieve Embeddings And Content |
| [**upload_chunks_and_embeddings**](EmbeddingsApi.md#upload_chunks_and_embeddings) | **POST** /upload_chunks_and_embeddings | Upload Chunks And Embeddings |

## get_documents

Embeddings

For pre-filtering documents, using `tags_v2` is preferred to using `tags` (which is now deprecated). If both `tags_v2` and `tags` are specified, `tags` is ignored. `tags_v2` enables building complex filters through the use of \"AND\", \"OR\", and negation logic. Take the below input as an example: ```json {     \"OR\": [         {             \"key\": \"subject\",             \"value\": \"holy-bible\",             \"negate\": false         },         {             \"key\": \"person-of-interest\",             \"value\": \"jesus christ\",             \"negate\": false         },         {             \"key\": \"genre\",             \"value\": \"religion\",             \"negate\": true         }         {             \"AND\": [                 {                     \"key\": \"subject\",                     \"value\": \"tao-te-ching\",                     \"negate\": false                 },                 {                     \"key\": \"author\",                     \"value\": \"lao-tzu\",                     \"negate\": false                 }             ]         }     ] } ``` In this case, files will be filtered such that: 1. \"subject\" = \"holy-bible\" OR 2. \"person-of-interest\" = \"jesus christ\" OR 3. \"genre\" != \"religion\" OR 4. \"subject\" = \"tao-te-ching\" AND \"author\" = \"lao-tzu\"  Note that the top level of the query must be either an \"OR\" or \"AND\" array. Currently, nesting is limited to 3. For tag blocks (those with \"key\", \"value\", and \"negate\" keys), the following typing rules apply: 1. \"key\" isn't optional and must be a `string` 2. \"value\" isn't optional and can be `any` or list[`any`] 3. \"negate\" is optional and must be `true` or `false`. If present and `true`, then the filter block is negated in the resulting query. It is `false` by default.   When querying embeddings, you can optionally specify the `media_type` parameter in your request. By default (if not set), it is equal to \"TEXT\". This means that the query will be performed over files that have been parsed as text (for now, this covers all files except image files). If it is equal to \"IMAGE\", the query will be performed over image files (for now, `.jpg` and `.png` files). You can think of this field as an additional filter on top of any filters set in `file_ids` and   When `hybrid_search` is set to true, a combination of keyword search and semantic search are used to rank and select candidate embeddings during information retrieval. By default, these search methods are weighted equally during the ranking process. To adjust the weight (or \"importance\") of each search method, you can use the `hybrid_search_tuning_parameters` property. The description for the different tuning parameters are: - `weight_a`: weight to assign to semantic search - `weight_b`: weight to assign to keyword search  You must ensure that `sum(weight_a, weight_b,..., weight_n)` for all *n* weights is equal to 1. The equality has an error tolerance of 0.001 to account for possible floating point issues.  In order to use hybrid search for a customer across a set of documents, two flags need to be enabled: 1. Use the `/modify_user_configuration` endpoint to to enable `sparse_vectors` for the customer. The payload body for this request is below: ``` {   \"configuration_key_name\": \"sparse_vectors\",   \"value\": {     \"enabled\": true   } } ``` 2. Make sure hybrid search is enabled for the documents across which you want to perform the search. For the `/uploadfile` endpoint, this can be done by setting the following query parameter: `generate_sparse_vectors=true`   Carbon supports multiple models for use in generating embeddings for files. For images, we support Vertex AI's multimodal model; for text, we support OpenAI's `text-embedding-ada-002` and Cohere's embed-multilingual-v3.0. The model can be specified via the `embedding_model` parameter (in the POST body for `/embeddings`, and a query  parameter in `/uploadfile`). If no model is supplied, the `text-embedding-ada-002` is used by default. When performing embedding queries, embeddings from files that used the specified model will be considered in the query. For example, if files A and B have embeddings generated with `OPENAI`, and files C and D have embeddings generated with `COHERE_MULTILINGUAL_V3`, then by default, queries will only consider files A and B. If `COHERE_MULTILINGUAL_V3` is specified as the `embedding_model` in `/embeddings`, then only files C and D will be considered. Make sure that the set of all files you want considered for a query have embeddings generated via the same model. For now, **do not** set `VERTEX_MULTIMODAL` as an `embedding_model`. This model is used automatically by Carbon when it detects an image file.

### Example

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'
Carbon.api_key = 'YOUR API KEY'
Carbon.customer_id = 'YOUR API KEY'

query = "a"
k = 1
tags = {
        "key": "string_example",
    }
query_vector = [
        3.14
    ]
file_ids = [
        1
    ]
parent_file_ids = [
        1
    ]
tags_v2 = {
    }
include_tags = True
include_vectors = True
include_raw_file = True
hybrid_search = True
hybrid_search_tuning_parameters = {
        "weight_a" => 0.5,
        "weight_b" => 0.5,
    }
media_type = "TEXT"
embedding_model = "OPENAI"

begin
  # Embeddings
  result = Carbon::Embeddings.get_documents(
                                               query: query,
                                               k: k,
                                               tags: tags,
                                               query_vector: query_vector,
                                               file_ids: file_ids,
                                               parent_file_ids: parent_file_ids,
                                               tags_v2: tags_v2,
                                               include_tags: include_tags,
                                               include_vectors: include_vectors,
                                               include_raw_file: include_raw_file,
                                               hybrid_search: hybrid_search,
                                               hybrid_search_tuning_parameters: hybrid_search_tuning_parameters,
                                               media_type: media_type,
                                               embedding_model: embedding_model,
                                             )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Embeddings.get_documents: #{e}"
end
```

#### Using the get_documents_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
query = "a"
k = 1
tags = {
        "key": "string_example",
    }
query_vector = [
        3.14
    ]
file_ids = [
        1
    ]
parent_file_ids = [
        1
    ]
tags_v2 = {
    }
include_tags = True
include_vectors = True
include_raw_file = True
hybrid_search = True
hybrid_search_tuning_parameters = {
        "weight_a" => 0.5,
        "weight_b" => 0.5,
    }
media_type = "TEXT"
embedding_model = "OPENAI"

begin
  # Embeddings
  data, status_code, headers, response = Carbon::Embeddings.get_documents_with_http_info(
                                                                                            query: query,
                                                                                            k: k,
                                                                                            tags: tags,
                                                                                            query_vector: query_vector,
                                                                                            file_ids: file_ids,
                                                                                            parent_file_ids: parent_file_ids,
                                                                                            tags_v2: tags_v2,
                                                                                            include_tags: include_tags,
                                                                                            include_vectors: include_vectors,
                                                                                            include_raw_file: include_raw_file,
                                                                                            hybrid_search: hybrid_search,
                                                                                            hybrid_search_tuning_parameters: hybrid_search_tuning_parameters,
                                                                                            media_type: media_type,
                                                                                            embedding_model: embedding_model,
                                                                                          )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => DocumentResponseList
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Embeddings.get_documents: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **get_embedding_documents_body** | [**GetEmbeddingDocumentsBody**](GetEmbeddingDocumentsBody.md) |  |  |

### Return type

[**DocumentResponseList**](DocumentResponseList.md)

## get_embeddings_and_chunks

Retrieve Embeddings And Content

### Example

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'
Carbon.api_key = 'YOUR API KEY'
Carbon.customer_id = 'YOUR API KEY'

filters = {
        "user_file_id" => 1,
        "embedding_model" => "OPENAI",
    }
pagination = {
        "limit" => 10,
        "offset" => 0,
    }
order_by = "created_at"
order_dir = "desc"
include_vectors = False

begin
  # Retrieve Embeddings And Content
  result = Carbon::Embeddings.get_embeddings_and_chunks(
                                                           filters: filters,
                                                           pagination: pagination,
                                                           order_by: order_by,
                                                           order_dir: order_dir,
                                                           include_vectors: include_vectors,
                                                         )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Embeddings.get_embeddings_and_chunks: #{e}"
end
```

#### Using the get_embeddings_and_chunks_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
filters = {
        "user_file_id" => 1,
        "embedding_model" => "OPENAI",
    }
pagination = {
        "limit" => 10,
        "offset" => 0,
    }
order_by = "created_at"
order_dir = "desc"
include_vectors = False

begin
  # Retrieve Embeddings And Content
  data, status_code, headers, response = Carbon::Embeddings.get_embeddings_and_chunks_with_http_info(
                                                                                                        filters: filters,
                                                                                                        pagination: pagination,
                                                                                                        order_by: order_by,
                                                                                                        order_dir: order_dir,
                                                                                                        include_vectors: include_vectors,
                                                                                                      )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => EmbeddingsAndChunksResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Embeddings.get_embeddings_and_chunks: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **embeddings_and_chunks_query_input** | [**EmbeddingsAndChunksQueryInput**](EmbeddingsAndChunksQueryInput.md) |  |  |

### Return type

[**EmbeddingsAndChunksResponse**](EmbeddingsAndChunksResponse.md)

## upload_chunks_and_embeddings

Upload Chunks And Embeddings

### Example

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'
Carbon.api_key = 'YOUR API KEY'
Carbon.customer_id = 'YOUR API KEY'

embedding_model = "OPENAI"
chunks_and_embeddings = [
        {
            "file_id" => 1,
            "chunks_and_embeddings" => [
                {
                    "chunk_number" => 1,
                    "chunk" => "chunk_example",
                    "embedding" => [
                        3.14
                    ],
                }
            ],
        }
    ]
overwrite_existing = False

begin
  # Upload Chunks And Embeddings
  result = Carbon::Embeddings.upload_chunks_and_embeddings(
                                                              embedding_model: embedding_model,
                                                              chunks_and_embeddings: chunks_and_embeddings,
                                                              overwrite_existing: overwrite_existing,
                                                            )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Embeddings.upload_chunks_and_embeddings: #{e}"
end
```

#### Using the upload_chunks_and_embeddings_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
embedding_model = "OPENAI"
chunks_and_embeddings = [
        {
            "file_id" => 1,
            "chunks_and_embeddings" => [
                {
                    "chunk_number" => 1,
                    "chunk" => "chunk_example",
                    "embedding" => [
                        3.14
                    ],
                }
            ],
        }
    ]
overwrite_existing = False

begin
  # Upload Chunks And Embeddings
  data, status_code, headers, response = Carbon::Embeddings.upload_chunks_and_embeddings_with_http_info(
                                                                                                           embedding_model: embedding_model,
                                                                                                           chunks_and_embeddings: chunks_and_embeddings,
                                                                                                           overwrite_existing: overwrite_existing,
                                                                                                         )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => GenericSuccessResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Embeddings.upload_chunks_and_embeddings: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **chunks_and_embeddings_upload_input** | [**ChunksAndEmbeddingsUploadInput**](ChunksAndEmbeddingsUploadInput.md) |  |  |

### Return type

[**GenericSuccessResponse**](GenericSuccessResponse.md)


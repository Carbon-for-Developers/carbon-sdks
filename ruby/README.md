<div align="center">

[![Visit Carbon](./header.png)](https://carbon.ai)

# [Carbon](https://carbon.ai)<a id="carbon"></a>

Connect external data to LLMs, no matter the source.

[![npm](https://img.shields.io/badge/gem-v0.2.24-blue)](https://rubygems.org/gems/carbon_ruby_sdk/versions/0.2.24)

</div>

## Table of Contents<a id="table-of-contents"></a>

<!-- toc -->

- [Installation](#installation)
- [Getting Started](#getting-started)
- [Raw HTTP Response](#raw-http-response)
- [Reference](#reference)
  * [`carbon.auth.get_access_token`](#carbonauthget_access_token)
  * [`carbon.auth.get_white_labeling`](#carbonauthget_white_labeling)
  * [`carbon.data_sources.query_user_data_sources`](#carbondata_sourcesquery_user_data_sources)
  * [`carbon.data_sources.revoke_access_token`](#carbondata_sourcesrevoke_access_token)
  * [`carbon.embeddings.get_documents`](#carbonembeddingsget_documents)
  * [`carbon.embeddings.get_embeddings_and_chunks`](#carbonembeddingsget_embeddings_and_chunks)
  * [`carbon.embeddings.list`](#carbonembeddingslist)
  * [`carbon.embeddings.upload_chunks_and_embeddings`](#carbonembeddingsupload_chunks_and_embeddings)
  * [`carbon.files.create_user_file_tags`](#carbonfilescreate_user_file_tags)
  * [`carbon.files.delete`](#carbonfilesdelete)
  * [`carbon.files.delete_file_tags`](#carbonfilesdelete_file_tags)
  * [`carbon.files.delete_many`](#carbonfilesdelete_many)
  * [`carbon.files.delete_v2`](#carbonfilesdelete_v2)
  * [`carbon.files.get_parsed_file`](#carbonfilesget_parsed_file)
  * [`carbon.files.get_raw_file`](#carbonfilesget_raw_file)
  * [`carbon.files.modify_cold_storage_parameters`](#carbonfilesmodify_cold_storage_parameters)
  * [`carbon.files.move_to_hot_storage`](#carbonfilesmove_to_hot_storage)
  * [`carbon.files.query_user_files`](#carbonfilesquery_user_files)
  * [`carbon.files.query_user_files_deprecated`](#carbonfilesquery_user_files_deprecated)
  * [`carbon.files.resync`](#carbonfilesresync)
  * [`carbon.files.upload`](#carbonfilesupload)
  * [`carbon.files.upload_from_url`](#carbonfilesupload_from_url)
  * [`carbon.files.upload_text`](#carbonfilesupload_text)
  * [`carbon.integrations.cancel`](#carbonintegrationscancel)
  * [`carbon.integrations.connect_data_source`](#carbonintegrationsconnect_data_source)
  * [`carbon.integrations.connect_freshdesk`](#carbonintegrationsconnect_freshdesk)
  * [`carbon.integrations.connect_gitbook`](#carbonintegrationsconnect_gitbook)
  * [`carbon.integrations.create_aws_iam_user`](#carbonintegrationscreate_aws_iam_user)
  * [`carbon.integrations.get_oauth_url`](#carbonintegrationsget_oauth_url)
  * [`carbon.integrations.list_confluence_pages`](#carbonintegrationslist_confluence_pages)
  * [`carbon.integrations.list_conversations`](#carbonintegrationslist_conversations)
  * [`carbon.integrations.list_data_source_items`](#carbonintegrationslist_data_source_items)
  * [`carbon.integrations.list_folders`](#carbonintegrationslist_folders)
  * [`carbon.integrations.list_gitbook_spaces`](#carbonintegrationslist_gitbook_spaces)
  * [`carbon.integrations.list_labels`](#carbonintegrationslist_labels)
  * [`carbon.integrations.list_outlook_categories`](#carbonintegrationslist_outlook_categories)
  * [`carbon.integrations.list_repos`](#carbonintegrationslist_repos)
  * [`carbon.integrations.sync_confluence`](#carbonintegrationssync_confluence)
  * [`carbon.integrations.sync_data_source_items`](#carbonintegrationssync_data_source_items)
  * [`carbon.integrations.sync_files`](#carbonintegrationssync_files)
  * [`carbon.integrations.sync_git_hub`](#carbonintegrationssync_git_hub)
  * [`carbon.integrations.sync_gitbook`](#carbonintegrationssync_gitbook)
  * [`carbon.integrations.sync_gmail`](#carbonintegrationssync_gmail)
  * [`carbon.integrations.sync_outlook`](#carbonintegrationssync_outlook)
  * [`carbon.integrations.sync_repos`](#carbonintegrationssync_repos)
  * [`carbon.integrations.sync_rss_feed`](#carbonintegrationssync_rss_feed)
  * [`carbon.integrations.sync_s3_files`](#carbonintegrationssync_s3_files)
  * [`carbon.integrations.sync_slack`](#carbonintegrationssync_slack)
  * [`carbon.organizations.get`](#carbonorganizationsget)
  * [`carbon.organizations.update`](#carbonorganizationsupdate)
  * [`carbon.organizations.update_stats`](#carbonorganizationsupdate_stats)
  * [`carbon.users.delete`](#carbonusersdelete)
  * [`carbon.users.get`](#carbonusersget)
  * [`carbon.users.list`](#carbonuserslist)
  * [`carbon.users.toggle_user_features`](#carbonuserstoggle_user_features)
  * [`carbon.users.update_users`](#carbonusersupdate_users)
  * [`carbon.utilities.fetch_urls`](#carbonutilitiesfetch_urls)
  * [`carbon.utilities.fetch_webpage`](#carbonutilitiesfetch_webpage)
  * [`carbon.utilities.fetch_youtube_transcripts`](#carbonutilitiesfetch_youtube_transcripts)
  * [`carbon.utilities.process_sitemap`](#carbonutilitiesprocess_sitemap)
  * [`carbon.utilities.scrape_sitemap`](#carbonutilitiesscrape_sitemap)
  * [`carbon.utilities.scrape_web`](#carbonutilitiesscrape_web)
  * [`carbon.utilities.search_urls`](#carbonutilitiessearch_urls)
  * [`carbon.utilities.user_webpages`](#carbonutilitiesuser_webpages)
  * [`carbon.webhooks.add_url`](#carbonwebhooksadd_url)
  * [`carbon.webhooks.delete_url`](#carbonwebhooksdelete_url)
  * [`carbon.webhooks.urls`](#carbonwebhooksurls)

<!-- tocstop -->

## Installation<a id="installation"></a>

Add to Gemfile:

```ruby
gem 'carbon_ruby_sdk', '~> 0.2.24'
```

## Getting Started<a id="getting-started"></a>

```ruby
require 'carbon_ruby_sdk'

# 1) Get an access token for a customer
configuration = Carbon::Configuration.new
configuration.api_key = "YOUR_API_KEY"
configuration.customer_id = "YOUR_CUSTOMER_ID"
carbon = Carbon::Client.new(configuration)

token = carbon.auth.get_access_token

# 2) Use the access token to authenticate moving forward
configuration = Carbon::Configuration.new
configuration.access_token = token.access_token
carbon = Carbon::Client.new(configuration)

# use SDK as usual
white_labeling = carbon.auth.get_white_labeling
```

## Raw HTTP Response<a id="raw-http-response"></a>

To access the raw HTTP response, suffix any method with `_with_http_info`.

```ruby
result = carbon.auth.get_access_token_with_http_info
p result.data # [TokenResponse] Deserialized data
p.result.status_code # [Integer] HTTP status code
p.result.headers # [Hash] HTTP headers
p.result.response # [Faraday::Response] Raw HTTP response
```

## Reference<a id="reference"></a>


### `carbon.auth.get_access_token`<a id="carbonauthget_access_token"></a>

Get Access Token

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.auth.get_access_token
p result
```

#### 🔄 Return<a id="🔄-return"></a>

[TokenResponse](./lib/carbon_ruby_sdk/models/token_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/auth/v1/access_token` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.auth.get_white_labeling`<a id="carbonauthget_white_labeling"></a>

Returns whether or not the organization is white labeled and which integrations are white labeled

:param current_user: the current user
:param db: the database session
:return: a WhiteLabelingResponse

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.auth.get_white_labeling
p result
```

#### 🔄 Return<a id="🔄-return"></a>

[WhiteLabelingResponse](./lib/carbon_ruby_sdk/models/white_labeling_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/auth/v1/white_labeling` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.data_sources.query_user_data_sources`<a id="carbondata_sourcesquery_user_data_sources"></a>

User Data Sources

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.data_sources.query_user_data_sources(
  pagination: {
        "limit" => 10,
        "offset" => 0,
    },
  order_by: "created_at",
  order_dir: "desc",
  filters: {
        "source" => "GOOGLE_CLOUD_STORAGE",
    },
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### pagination: [`Pagination`](./lib/carbon_ruby_sdk/models/pagination.rb)<a id="pagination-paginationlibcarbon_ruby_sdkmodelspaginationrb"></a>
##### order_by: [`OrganizationUserDataSourceOrderByColumns`](./lib/carbon_ruby_sdk/models/organization_user_data_source_order_by_columns.rb)<a id="order_by-organizationuserdatasourceorderbycolumnslibcarbon_ruby_sdkmodelsorganization_user_data_source_order_by_columnsrb"></a>
##### order_dir: [`OrderDir`](./lib/carbon_ruby_sdk/models/order_dir.rb)<a id="order_dir-orderdirlibcarbon_ruby_sdkmodelsorder_dirrb"></a>
##### filters: [`OrganizationUserDataSourceFilters`](./lib/carbon_ruby_sdk/models/organization_user_data_source_filters.rb)<a id="filters-organizationuserdatasourcefilterslibcarbon_ruby_sdkmodelsorganization_user_data_source_filtersrb"></a>
#### 🔄 Return<a id="🔄-return"></a>

[OrganizationUserDataSourceResponse](./lib/carbon_ruby_sdk/models/organization_user_data_source_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/user_data_sources` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.data_sources.revoke_access_token`<a id="carbondata_sourcesrevoke_access_token"></a>

Revoke Access Token

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.data_sources.revoke_access_token(
  data_source_id: 1,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### data_source_id: `Integer`<a id="data_source_id-integer"></a>
#### 🔄 Return<a id="🔄-return"></a>

[GenericSuccessResponse](./lib/carbon_ruby_sdk/models/generic_success_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/revoke_access_token` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.embeddings.get_documents`<a id="carbonembeddingsget_documents"></a>

For pre-filtering documents, using `tags_v2` is preferred to using `tags` (which is now deprecated). If both `tags_v2`
and `tags` are specified, `tags` is ignored. `tags_v2` enables
building complex filters through the use of "AND", "OR", and negation logic. Take the below input as an example:
```json
{
    "OR": [
        {
            "key": "subject",
            "value": "holy-bible",
            "negate": false
        },
        {
            "key": "person-of-interest",
            "value": "jesus christ",
            "negate": false
        },
        {
            "key": "genre",
            "value": "religion",
            "negate": true
        }
        {
            "AND": [
                {
                    "key": "subject",
                    "value": "tao-te-ching",
                    "negate": false
                },
                {
                    "key": "author",
                    "value": "lao-tzu",
                    "negate": false
                }
            ]
        }
    ]
}
```
In this case, files will be filtered such that:
1. "subject" = "holy-bible" OR
2. "person-of-interest" = "jesus christ" OR
3. "genre" != "religion" OR
4. "subject" = "tao-te-ching" AND "author" = "lao-tzu"

Note that the top level of the query must be either an "OR" or "AND" array. Currently, nesting is limited to 3.
For tag blocks (those with "key", "value", and "negate" keys), the following typing rules apply:
1. "key" isn't optional and must be a `string`
2. "value" isn't optional and can be `any` or list[`any`]
3. "negate" is optional and must be `true` or `false`. If present and `true`, then the filter block is negated in
the resulting query. It is `false` by default.


When querying embeddings, you can optionally specify the `media_type` parameter in your request. By default (if
not set), it is equal to "TEXT". This means that the query will be performed over files that have
been parsed as text (for now, this covers all files except image files). If it is equal to "IMAGE",
the query will be performed over image files (for now, `.jpg` and `.png` files). You can think of this
field as an additional filter on top of any filters set in `file_ids` and


When `hybrid_search` is set to true, a combination of keyword search and semantic search are used to rank
and select candidate embeddings during information retrieval. By default, these search methods are weighted
equally during the ranking process. To adjust the weight (or "importance") of each search method, you can use
the `hybrid_search_tuning_parameters` property. The description for the different tuning parameters are:
- `weight_a`: weight to assign to semantic search
- `weight_b`: weight to assign to keyword search

You must ensure that `sum(weight_a, weight_b,..., weight_n)` for all *n* weights is equal to 1. The equality
has an error tolerance of 0.001 to account for possible floating point issues.

In order to use hybrid search for a customer across a set of documents, two flags need to be enabled:
1. Use the `/modify_user_configuration` endpoint to to enable `sparse_vectors` for the customer. The payload
body for this request is below:
```
{
  "configuration_key_name": "sparse_vectors",
  "value": {
    "enabled": true
  }
}
```
2. Make sure hybrid search is enabled for the documents across which you want to perform the search. For the
`/uploadfile` endpoint, this can be done by setting the following query parameter: `generate_sparse_vectors=true`


Carbon supports multiple models for use in generating embeddings for files. For images, we support Vertex AI's
multimodal model; for text, we support OpenAI's `text-embedding-ada-002` and Cohere's embed-multilingual-v3.0.
The model can be specified via the `embedding_model` parameter (in the POST body for `/embeddings`, and a query 
parameter in `/uploadfile`). If no model is supplied, the `text-embedding-ada-002` is used by default. When performing
embedding queries, embeddings from files that used the specified model will be considered in the query.
For example, if files A and B have embeddings generated with `OPENAI`, and files C and D have embeddings generated with
`COHERE_MULTILINGUAL_V3`, then by default, queries will only consider files A and B. If `COHERE_MULTILINGUAL_V3` is
specified as the `embedding_model` in `/embeddings`, then only files C and D will be considered. Make sure that
the set of all files you want considered for a query have embeddings generated via the same model. For now, **do not**
set `VERTEX_MULTIMODAL` as an `embedding_model`. This model is used automatically by Carbon when it detects an image file.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.embeddings.get_documents(
  query: "a",
  k: 1,
  tags: {
        "key": "string_example",
    },
  query_vector: [
        3.14
    ],
  file_ids: [
        1
    ],
  parent_file_ids: [
        1
    ],
  include_all_children: false,
  tags_v2: {
    },
  include_tags: true,
  include_vectors: true,
  include_raw_file: true,
  hybrid_search: true,
  hybrid_search_tuning_parameters: {
        "weight_a" => 0.5,
        "weight_b" => 0.5,
    },
  media_type: "TEXT",
  embedding_model: "OPENAI",
  include_file_level_metadata: false,
  high_accuracy: false,
  rerank: {
        "model" => "model_example",
    },
  file_types_at_source: [
        "string_example"
    ],
  exclude_cold_storage_files: false,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### query: `String`<a id="query-string"></a>
Query for which to get related chunks and embeddings.

##### k: `Integer`<a id="k-integer"></a>
Number of related chunks to return.

##### tags: Hash<string, [`Tags1`](./lib/carbon_ruby_sdk/models/tags1.rb)><a id="tags-hash"></a>
A set of tags to limit the search to. Deprecated and may be removed in the
future.

##### query_vector: Array<`Float`><a id="query_vector-array"></a>
Optional query vector for which to get related chunks and embeddings. It must
have been generated by the same model used to generate the embeddings across
which the search is being conducted. Cannot provide both `query` and
`query_vector`.

##### file_ids: Array<`Integer`><a id="file_ids-array"></a>
Optional list of file IDs to limit the search to

##### parent_file_ids: Array<`Integer`><a id="parent_file_ids-array"></a>
Optional list of parent file IDs to limit the search to. A parent file describes
a file to which another file belongs (e.g. a folder)

##### include_all_children: `Boolean`<a id="include_all_children-boolean"></a>
Flag to control whether or not to include all children of filtered files in the
embedding search.

##### tags_v2: `Object`<a id="tags_v2-object"></a>
A set of tags to limit the search to. Use this instead of `tags`, which is
deprecated.

##### include_tags: `Boolean`<a id="include_tags-boolean"></a>
Flag to control whether or not to include tags for each chunk in the response.

##### include_vectors: `Boolean`<a id="include_vectors-boolean"></a>
Flag to control whether or not to include embedding vectors in the response.

##### include_raw_file: `Boolean`<a id="include_raw_file-boolean"></a>
Flag to control whether or not to include a signed URL to the raw file
containing each chunk in the response.

##### hybrid_search: `Boolean`<a id="hybrid_search-boolean"></a>
Flag to control whether or not to perform hybrid search.

##### hybrid_search_tuning_parameters: [`HybridSearchTuningParamsNullable`](./lib/carbon_ruby_sdk/models/hybrid_search_tuning_params_nullable.rb)<a id="hybrid_search_tuning_parameters-hybridsearchtuningparamsnullablelibcarbon_ruby_sdkmodelshybrid_search_tuning_params_nullablerb"></a>
##### media_type: [`FileContentTypesNullable`](./lib/carbon_ruby_sdk/models/file_content_types_nullable.rb)<a id="media_type-filecontenttypesnullablelibcarbon_ruby_sdkmodelsfile_content_types_nullablerb"></a>
##### embedding_model: [`EmbeddingGeneratorsNullable`](./lib/carbon_ruby_sdk/models/embedding_generators_nullable.rb)<a id="embedding_model-embeddinggeneratorsnullablelibcarbon_ruby_sdkmodelsembedding_generators_nullablerb"></a>
##### include_file_level_metadata: `Boolean`<a id="include_file_level_metadata-boolean"></a>
Flag to control whether or not to include file-level metadata in the response.
This metadata will be included in the `content_metadata` field of each document
along with chunk/embedding level metadata.

##### high_accuracy: `Boolean`<a id="high_accuracy-boolean"></a>
Flag to control whether or not to perform a high accuracy embedding search. By
default, this is set to false. If true, the search may return more accurate
results, but may take longer to complete.

##### rerank: [`RerankParamsNullable`](./lib/carbon_ruby_sdk/models/rerank_params_nullable.rb)<a id="rerank-rerankparamsnullablelibcarbon_ruby_sdkmodelsrerank_params_nullablerb"></a>
##### file_types_at_source: Array<[`HelpdeskFileTypes`](./lib/carbon_ruby_sdk/models/helpdesk_file_types.rb)><a id="file_types_at_source-array"></a>
Filter files based on their type at the source (for example help center tickets
and articles)

##### exclude_cold_storage_files: `Boolean`<a id="exclude_cold_storage_files-boolean"></a>
Flag to control whether or not to exclude files that are not in hot storage. If
set to False, then an error will be returned if any filtered files are in cold
storage.

#### 🔄 Return<a id="🔄-return"></a>

[DocumentResponseList](./lib/carbon_ruby_sdk/models/document_response_list.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/embeddings` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.embeddings.get_embeddings_and_chunks`<a id="carbonembeddingsget_embeddings_and_chunks"></a>

Retrieve Embeddings And Content

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.embeddings.get_embeddings_and_chunks(
  filters: {
        "user_file_id" => 1,
        "embedding_model" => "OPENAI",
    },
  pagination: {
        "limit" => 10,
        "offset" => 0,
    },
  order_by: "created_at",
  order_dir: "desc",
  include_vectors: false,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### filters: [`EmbeddingsAndChunksFilters`](./lib/carbon_ruby_sdk/models/embeddings_and_chunks_filters.rb)<a id="filters-embeddingsandchunksfilterslibcarbon_ruby_sdkmodelsembeddings_and_chunks_filtersrb"></a>
##### pagination: [`Pagination`](./lib/carbon_ruby_sdk/models/pagination.rb)<a id="pagination-paginationlibcarbon_ruby_sdkmodelspaginationrb"></a>
##### order_by: [`EmbeddingsAndChunksOrderByColumns`](./lib/carbon_ruby_sdk/models/embeddings_and_chunks_order_by_columns.rb)<a id="order_by-embeddingsandchunksorderbycolumnslibcarbon_ruby_sdkmodelsembeddings_and_chunks_order_by_columnsrb"></a>
##### order_dir: [`OrderDir`](./lib/carbon_ruby_sdk/models/order_dir.rb)<a id="order_dir-orderdirlibcarbon_ruby_sdkmodelsorder_dirrb"></a>
##### include_vectors: `Boolean`<a id="include_vectors-boolean"></a>
#### 🔄 Return<a id="🔄-return"></a>

[EmbeddingsAndChunksResponse](./lib/carbon_ruby_sdk/models/embeddings_and_chunks_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/text_chunks` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.embeddings.list`<a id="carbonembeddingslist"></a>

Retrieve Embeddings And Content V2

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.embeddings.list(
  filters: {
        "include_all_children" => false,
        "non_synced_only" => false,
    },
  pagination: {
        "limit" => 10,
        "offset" => 0,
    },
  order_by: "created_at",
  order_dir: "desc",
  include_vectors: false,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### filters: [`OrganizationUserFilesToSyncFilters`](./lib/carbon_ruby_sdk/models/organization_user_files_to_sync_filters.rb)<a id="filters-organizationuserfilestosyncfilterslibcarbon_ruby_sdkmodelsorganization_user_files_to_sync_filtersrb"></a>
##### pagination: [`Pagination`](./lib/carbon_ruby_sdk/models/pagination.rb)<a id="pagination-paginationlibcarbon_ruby_sdkmodelspaginationrb"></a>
##### order_by: [`OrganizationUserFilesToSyncOrderByTypes`](./lib/carbon_ruby_sdk/models/organization_user_files_to_sync_order_by_types.rb)<a id="order_by-organizationuserfilestosyncorderbytypeslibcarbon_ruby_sdkmodelsorganization_user_files_to_sync_order_by_typesrb"></a>
##### order_dir: [`OrderDir`](./lib/carbon_ruby_sdk/models/order_dir.rb)<a id="order_dir-orderdirlibcarbon_ruby_sdkmodelsorder_dirrb"></a>
##### include_vectors: `Boolean`<a id="include_vectors-boolean"></a>
#### 🔄 Return<a id="🔄-return"></a>

[EmbeddingsAndChunksResponse](./lib/carbon_ruby_sdk/models/embeddings_and_chunks_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/list_chunks_and_embeddings` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.embeddings.upload_chunks_and_embeddings`<a id="carbonembeddingsupload_chunks_and_embeddings"></a>

Upload Chunks And Embeddings

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.embeddings.upload_chunks_and_embeddings(
  embedding_model: "OPENAI",
  chunks_and_embeddings: [
        {
            "file_id" => 1,
            "chunks_and_embeddings" => [
                {
                    "chunk_number" => 1,
                    "chunk" => "chunk_example",
                }
            ],
        }
    ],
  overwrite_existing: false,
  chunks_only: false,
  custom_credentials: {
        "key": {},
    },
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### embedding_model: [`EmbeddingGenerators`](./lib/carbon_ruby_sdk/models/embedding_generators.rb)<a id="embedding_model-embeddinggeneratorslibcarbon_ruby_sdkmodelsembedding_generatorsrb"></a>
##### chunks_and_embeddings: Array<[`SingleChunksAndEmbeddingsUploadInput`](./lib/carbon_ruby_sdk/models/single_chunks_and_embeddings_upload_input.rb)><a id="chunks_and_embeddings-array"></a>
##### overwrite_existing: `Boolean`<a id="overwrite_existing-boolean"></a>
##### chunks_only: `Boolean`<a id="chunks_only-boolean"></a>
##### custom_credentials: `Hash<String, Object>`<a id="custom_credentials-hash"></a>
#### 🔄 Return<a id="🔄-return"></a>

[GenericSuccessResponse](./lib/carbon_ruby_sdk/models/generic_success_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/upload_chunks_and_embeddings` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.create_user_file_tags`<a id="carbonfilescreate_user_file_tags"></a>

A tag is a key-value pair that can be added to a file. This pair can then be used
for searches (e.g. embedding searches) in order to narrow down the scope of the search.
A file can have any number of tags. The following are reserved keys that cannot be used:
- db_embedding_id
- organization_id
- user_id
- organization_user_file_id

Carbon currently supports two data types for tag values - `string` and `list<string>`.
Keys can only be `string`. If values other than `string` and `list<string>` are used,
they're automatically converted to strings (e.g. 4 will become "4").

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.files.create_user_file_tags(
  tags: {
        "key": "string_example",
    },
  organization_user_file_id: 1,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### tags: Hash<string, [`Tags1`](./lib/carbon_ruby_sdk/models/tags1.rb)><a id="tags-hash"></a>
##### organization_user_file_id: `Integer`<a id="organization_user_file_id-integer"></a>
#### 🔄 Return<a id="🔄-return"></a>

[UserFile](./lib/carbon_ruby_sdk/models/user_file.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/create_user_file_tags` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.delete`<a id="carbonfilesdelete"></a>
![Deprecated](https://img.shields.io/badge/deprecated-yellow)

Delete File Endpoint

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.files.delete(
  file_id: 1,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### file_id: `Integer`<a id="file_id-integer"></a>
#### 🔄 Return<a id="🔄-return"></a>

[GenericSuccessResponse](./lib/carbon_ruby_sdk/models/generic_success_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/deletefile/{file_id}` `DELETE`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.delete_file_tags`<a id="carbonfilesdelete_file_tags"></a>

Delete File Tags

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.files.delete_file_tags(
  tags: [
        "string_example"
    ],
  organization_user_file_id: 1,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### tags: Array<`String`><a id="tags-array"></a>
##### organization_user_file_id: `Integer`<a id="organization_user_file_id-integer"></a>
#### 🔄 Return<a id="🔄-return"></a>

[UserFile](./lib/carbon_ruby_sdk/models/user_file.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/delete_user_file_tags` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.delete_many`<a id="carbonfilesdelete_many"></a>
![Deprecated](https://img.shields.io/badge/deprecated-yellow)

Delete Files Endpoint

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.files.delete_many(
  file_ids: [
        1
    ],
  sync_statuses: [
        "string_example"
    ],
  delete_non_synced_only: false,
  send_webhook: false,
  delete_child_files: false,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### file_ids: Array<`Integer`><a id="file_ids-array"></a>
##### sync_statuses: Array<[`ExternalFileSyncStatuses`](./lib/carbon_ruby_sdk/models/external_file_sync_statuses.rb)><a id="sync_statuses-array"></a>
##### delete_non_synced_only: `Boolean`<a id="delete_non_synced_only-boolean"></a>
##### send_webhook: `Boolean`<a id="send_webhook-boolean"></a>
##### delete_child_files: `Boolean`<a id="delete_child_files-boolean"></a>
#### 🔄 Return<a id="🔄-return"></a>

[GenericSuccessResponse](./lib/carbon_ruby_sdk/models/generic_success_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/delete_files` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.delete_v2`<a id="carbonfilesdelete_v2"></a>

Delete Files V2 Endpoint

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.files.delete_v2(
  filters: {
        "include_all_children" => false,
        "non_synced_only" => false,
    },
  send_webhook: false,
  preserve_file_record: false,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### filters: [`OrganizationUserFilesToSyncFilters`](./lib/carbon_ruby_sdk/models/organization_user_files_to_sync_filters.rb)<a id="filters-organizationuserfilestosyncfilterslibcarbon_ruby_sdkmodelsorganization_user_files_to_sync_filtersrb"></a>
##### send_webhook: `Boolean`<a id="send_webhook-boolean"></a>
##### preserve_file_record: `Boolean`<a id="preserve_file_record-boolean"></a>
Whether or not to delete all data related to the file from the database, BUT to
preserve the file metadata, allowing for resyncs. By default
`preserve_file_record` is false, which means that all data related to the file
*as well as* its metadata will be deleted. Note that even if
`preserve_file_record` is true, raw files uploaded via the `uploadfile` endpoint
still cannot be resynced.

#### 🔄 Return<a id="🔄-return"></a>

[GenericSuccessResponse](./lib/carbon_ruby_sdk/models/generic_success_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/delete_files_v2` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.get_parsed_file`<a id="carbonfilesget_parsed_file"></a>
![Deprecated](https://img.shields.io/badge/deprecated-yellow)

This route is deprecated. Use `/user_files_v2` instead.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.files.get_parsed_file(
  file_id: 1,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### file_id: `Integer`<a id="file_id-integer"></a>
#### 🔄 Return<a id="🔄-return"></a>

[PresignedURLResponse](./lib/carbon_ruby_sdk/models/presigned_url_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/parsed_file/{file_id}` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.get_raw_file`<a id="carbonfilesget_raw_file"></a>
![Deprecated](https://img.shields.io/badge/deprecated-yellow)

This route is deprecated. Use `/user_files_v2` instead.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.files.get_raw_file(
  file_id: 1,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### file_id: `Integer`<a id="file_id-integer"></a>
#### 🔄 Return<a id="🔄-return"></a>

[PresignedURLResponse](./lib/carbon_ruby_sdk/models/presigned_url_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/raw_file/{file_id}` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.modify_cold_storage_parameters`<a id="carbonfilesmodify_cold_storage_parameters"></a>

Modify Cold Storage Parameters

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.files.modify_cold_storage_parameters(
  filters: {
        "include_all_children" => false,
        "non_synced_only" => false,
    },
  enable_cold_storage: true,
  hot_storage_time_to_live: 1,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### filters: [`OrganizationUserFilesToSyncFilters`](./lib/carbon_ruby_sdk/models/organization_user_files_to_sync_filters.rb)<a id="filters-organizationuserfilestosyncfilterslibcarbon_ruby_sdkmodelsorganization_user_files_to_sync_filtersrb"></a>
##### enable_cold_storage: `Boolean`<a id="enable_cold_storage-boolean"></a>
##### hot_storage_time_to_live: `Integer`<a id="hot_storage_time_to_live-integer"></a>
#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/modify_cold_storage_parameters` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.move_to_hot_storage`<a id="carbonfilesmove_to_hot_storage"></a>

Move To Hot Storage

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.files.move_to_hot_storage(
  filters: {
        "include_all_children" => false,
        "non_synced_only" => false,
    },
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### filters: [`OrganizationUserFilesToSyncFilters`](./lib/carbon_ruby_sdk/models/organization_user_files_to_sync_filters.rb)<a id="filters-organizationuserfilestosyncfilterslibcarbon_ruby_sdkmodelsorganization_user_files_to_sync_filtersrb"></a>
#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/move_to_hot_storage` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.query_user_files`<a id="carbonfilesquery_user_files"></a>

For pre-filtering documents, using `tags_v2` is preferred to using `tags` (which is now deprecated). If both `tags_v2`
and `tags` are specified, `tags` is ignored. `tags_v2` enables
building complex filters through the use of "AND", "OR", and negation logic. Take the below input as an example:
```json
{
    "OR": [
        {
            "key": "subject",
            "value": "holy-bible",
            "negate": false
        },
        {
            "key": "person-of-interest",
            "value": "jesus christ",
            "negate": false
        },
        {
            "key": "genre",
            "value": "religion",
            "negate": true
        }
        {
            "AND": [
                {
                    "key": "subject",
                    "value": "tao-te-ching",
                    "negate": false
                },
                {
                    "key": "author",
                    "value": "lao-tzu",
                    "negate": false
                }
            ]
        }
    ]
}
```
In this case, files will be filtered such that:
1. "subject" = "holy-bible" OR
2. "person-of-interest" = "jesus christ" OR
3. "genre" != "religion" OR
4. "subject" = "tao-te-ching" AND "author" = "lao-tzu"

Note that the top level of the query must be either an "OR" or "AND" array. Currently, nesting is limited to 3.
For tag blocks (those with "key", "value", and "negate" keys), the following typing rules apply:
1. "key" isn't optional and must be a `string`
2. "value" isn't optional and can be `any` or list[`any`]
3. "negate" is optional and must be `true` or `false`. If present and `true`, then the filter block is negated in
the resulting query. It is `false` by default.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.files.query_user_files(
  pagination: {
        "limit" => 10,
        "offset" => 0,
    },
  order_by: "created_at",
  order_dir: "desc",
  filters: {
        "include_all_children" => false,
        "non_synced_only" => false,
    },
  include_raw_file: true,
  include_parsed_text_file: true,
  include_additional_files: true,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### pagination: [`Pagination`](./lib/carbon_ruby_sdk/models/pagination.rb)<a id="pagination-paginationlibcarbon_ruby_sdkmodelspaginationrb"></a>
##### order_by: [`OrganizationUserFilesToSyncOrderByTypes`](./lib/carbon_ruby_sdk/models/organization_user_files_to_sync_order_by_types.rb)<a id="order_by-organizationuserfilestosyncorderbytypeslibcarbon_ruby_sdkmodelsorganization_user_files_to_sync_order_by_typesrb"></a>
##### order_dir: [`OrderDir`](./lib/carbon_ruby_sdk/models/order_dir.rb)<a id="order_dir-orderdirlibcarbon_ruby_sdkmodelsorder_dirrb"></a>
##### filters: [`OrganizationUserFilesToSyncFilters`](./lib/carbon_ruby_sdk/models/organization_user_files_to_sync_filters.rb)<a id="filters-organizationuserfilestosyncfilterslibcarbon_ruby_sdkmodelsorganization_user_files_to_sync_filtersrb"></a>
##### include_raw_file: `Boolean`<a id="include_raw_file-boolean"></a>
##### include_parsed_text_file: `Boolean`<a id="include_parsed_text_file-boolean"></a>
##### include_additional_files: `Boolean`<a id="include_additional_files-boolean"></a>
#### 🔄 Return<a id="🔄-return"></a>

[UserFilesV2](./lib/carbon_ruby_sdk/models/user_files_v2.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/user_files_v2` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.query_user_files_deprecated`<a id="carbonfilesquery_user_files_deprecated"></a>
![Deprecated](https://img.shields.io/badge/deprecated-yellow)

This route is deprecated. Use `/user_files_v2` instead.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.files.query_user_files_deprecated(
  pagination: {
        "limit" => 10,
        "offset" => 0,
    },
  order_by: "created_at",
  order_dir: "desc",
  filters: {
        "include_all_children" => false,
        "non_synced_only" => false,
    },
  include_raw_file: true,
  include_parsed_text_file: true,
  include_additional_files: true,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### pagination: [`Pagination`](./lib/carbon_ruby_sdk/models/pagination.rb)<a id="pagination-paginationlibcarbon_ruby_sdkmodelspaginationrb"></a>
##### order_by: [`OrganizationUserFilesToSyncOrderByTypes`](./lib/carbon_ruby_sdk/models/organization_user_files_to_sync_order_by_types.rb)<a id="order_by-organizationuserfilestosyncorderbytypeslibcarbon_ruby_sdkmodelsorganization_user_files_to_sync_order_by_typesrb"></a>
##### order_dir: [`OrderDir`](./lib/carbon_ruby_sdk/models/order_dir.rb)<a id="order_dir-orderdirlibcarbon_ruby_sdkmodelsorder_dirrb"></a>
##### filters: [`OrganizationUserFilesToSyncFilters`](./lib/carbon_ruby_sdk/models/organization_user_files_to_sync_filters.rb)<a id="filters-organizationuserfilestosyncfilterslibcarbon_ruby_sdkmodelsorganization_user_files_to_sync_filtersrb"></a>
##### include_raw_file: `Boolean`<a id="include_raw_file-boolean"></a>
##### include_parsed_text_file: `Boolean`<a id="include_parsed_text_file-boolean"></a>
##### include_additional_files: `Boolean`<a id="include_additional_files-boolean"></a>
#### 🔄 Return<a id="🔄-return"></a>

[UserFile](./lib/carbon_ruby_sdk/models/user_file.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/user_files` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.resync`<a id="carbonfilesresync"></a>

Resync File

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.files.resync(
  file_id: 1,
  chunk_size: 1,
  chunk_overlap: 1,
  force_embedding_generation: false,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### file_id: `Integer`<a id="file_id-integer"></a>
##### chunk_size: `Integer`<a id="chunk_size-integer"></a>
##### chunk_overlap: `Integer`<a id="chunk_overlap-integer"></a>
##### force_embedding_generation: `Boolean`<a id="force_embedding_generation-boolean"></a>
#### 🔄 Return<a id="🔄-return"></a>

[UserFile](./lib/carbon_ruby_sdk/models/user_file.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/resync_file` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.upload`<a id="carbonfilesupload"></a>

This endpoint is used to directly upload local files to Carbon. The `POST` request should be a multipart form request.
Note that the `set_page_as_boundary` query parameter is applicable only to PDFs for now. When this value is set,
PDF chunks are at most one page long. Additional information can be retrieved for each chunk, however, namely the coordinates
of the bounding box around the chunk (this can be used for things like text highlighting). Following is a description
of all possible query parameters:
- `chunk_size`: the chunk size (in tokens) applied when splitting the document
- `chunk_overlap`: the chunk overlap (in tokens) applied when splitting the document
- `skip_embedding_generation`: whether or not to skip the generation of chunks and embeddings
- `set_page_as_boundary`: described above
- `embedding_model`: the model used to generate embeddings for the document chunks
- `use_ocr`: whether or not to use OCR as a preprocessing step prior to generating chunks. Valid for PDFs, JPEGs, and PNGs
- `generate_sparse_vectors`: whether or not to generate sparse vectors for the file. Required for hybrid search.
- `prepend_filename_to_chunks`: whether or not to prepend the filename to the chunk text


Carbon supports multiple models for use in generating embeddings for files. For images, we support Vertex AI's
multimodal model; for text, we support OpenAI's `text-embedding-ada-002` and Cohere's embed-multilingual-v3.0.
The model can be specified via the `embedding_model` parameter (in the POST body for `/embeddings`, and a query 
parameter in `/uploadfile`). If no model is supplied, the `text-embedding-ada-002` is used by default. When performing
embedding queries, embeddings from files that used the specified model will be considered in the query.
For example, if files A and B have embeddings generated with `OPENAI`, and files C and D have embeddings generated with
`COHERE_MULTILINGUAL_V3`, then by default, queries will only consider files A and B. If `COHERE_MULTILINGUAL_V3` is
specified as the `embedding_model` in `/embeddings`, then only files C and D will be considered. Make sure that
the set of all files you want considered for a query have embeddings generated via the same model. For now, **do not**
set `VERTEX_MULTIMODAL` as an `embedding_model`. This model is used automatically by Carbon when it detects an image file.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.files.upload(
  file: File.open("path/to/file", "rb"),
  chunk_size: 1,
  chunk_overlap: 1,
  skip_embedding_generation: false,
  set_page_as_boundary: false,
  embedding_model: "string_example",
  use_ocr: false,
  generate_sparse_vectors: false,
  prepend_filename_to_chunks: false,
  max_items_per_chunk: 1,
  parse_pdf_tables_with_ocr: false,
  detect_audio_language: false,
  transcription_service: "assemblyai",
  include_speaker_labels: false,
  media_type: "TEXT",
  split_rows: false,
  enable_cold_storage: false,
  hot_storage_time_to_live: 1,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### file: `File`<a id="file-file"></a>
##### chunk_size: `Integer`<a id="chunk_size-integer"></a>
Chunk size in tiktoken tokens to be used when processing file.

##### chunk_overlap: `Integer`<a id="chunk_overlap-integer"></a>
Chunk overlap in tiktoken tokens to be used when processing file.

##### skip_embedding_generation: `Boolean`<a id="skip_embedding_generation-boolean"></a>
Flag to control whether or not embeddings should be generated and stored when
processing file.

##### set_page_as_boundary: `Boolean`<a id="set_page_as_boundary-boolean"></a>
Flag to control whether or not to set the a page's worth of content as the
maximum amount of content that can appear in a chunk. Only valid for PDFs. See
description route description for more information.

##### embedding_model: [`EmbeddingModel`](./lib/carbon_ruby_sdk/models/embedding_model.rb)<a id="embedding_model-embeddingmodellibcarbon_ruby_sdkmodelsembedding_modelrb"></a>
Embedding model that will be used to embed file chunks.

##### use_ocr: `Boolean`<a id="use_ocr-boolean"></a>
Whether or not to use OCR when processing files. Valid for PDFs, JPEGs, and
PNGs. Useful for documents with tables, images, and/or scanned text.

##### generate_sparse_vectors: `Boolean`<a id="generate_sparse_vectors-boolean"></a>
Whether or not to generate sparse vectors for the file. This is *required* for
the file to be a candidate for hybrid search.

##### prepend_filename_to_chunks: `Boolean`<a id="prepend_filename_to_chunks-boolean"></a>
Whether or not to prepend the file's name to chunks.

##### max_items_per_chunk: `Integer`<a id="max_items_per_chunk-integer"></a>
Number of objects per chunk. For csv, tsv, xlsx, and json files only.

##### parse_pdf_tables_with_ocr: `Boolean`<a id="parse_pdf_tables_with_ocr-boolean"></a>
Whether to use rich table parsing when `use_ocr` is enabled.

##### detect_audio_language: `Boolean`<a id="detect_audio_language-boolean"></a>
Whether to automatically detect the language of the uploaded audio file.

##### transcription_service: [`TranscriptionServiceNullable`](./lib/carbon_ruby_sdk/models/transcription_service_nullable.rb)<a id="transcription_service-transcriptionservicenullablelibcarbon_ruby_sdkmodelstranscription_service_nullablerb"></a>
The transcription service to use for audio files. If no service is specified,
'deepgram' will be used.

##### include_speaker_labels: `Boolean`<a id="include_speaker_labels-boolean"></a>
Detect multiple speakers and label segments of speech by speaker for audio
files.

##### media_type: [`FileContentTypesNullable`](./lib/carbon_ruby_sdk/models/file_content_types_nullable.rb)<a id="media_type-filecontenttypesnullablelibcarbon_ruby_sdkmodelsfile_content_types_nullablerb"></a>
The media type of the file. If not provided, it will be inferred from the file
extension.

##### split_rows: `Boolean`<a id="split_rows-boolean"></a>
Whether to split tabular rows into chunks. Currently only valid for CSV, TSV,
and XLSX files.

##### enable_cold_storage: `Boolean`<a id="enable_cold_storage-boolean"></a>
Enable cold storage for the file. If set to true, the file will be moved to cold
storage after a certain period of inactivity. Default is false.

##### hot_storage_time_to_live: `Integer`<a id="hot_storage_time_to_live-integer"></a>
Time in seconds after which the file will be moved to cold storage.

#### 🔄 Return<a id="🔄-return"></a>

[UserFile](./lib/carbon_ruby_sdk/models/user_file.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/uploadfile` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.upload_from_url`<a id="carbonfilesupload_from_url"></a>

Create Upload File From Url

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.files.upload_from_url(
  url: "string_example",
  file_name: "string_example",
  chunk_size: 1,
  chunk_overlap: 1,
  skip_embedding_generation: false,
  set_page_as_boundary: false,
  embedding_model: "OPENAI",
  generate_sparse_vectors: false,
  use_textract: false,
  prepend_filename_to_chunks: false,
  max_items_per_chunk: 1,
  parse_pdf_tables_with_ocr: false,
  detect_audio_language: false,
  transcription_service: "assemblyai",
  include_speaker_labels: false,
  media_type: "TEXT",
  split_rows: false,
  cold_storage_params: {
        "enable_cold_storage" => false,
    },
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### url: `String`<a id="url-string"></a>
##### file_name: `String`<a id="file_name-string"></a>
##### chunk_size: `Integer`<a id="chunk_size-integer"></a>
##### chunk_overlap: `Integer`<a id="chunk_overlap-integer"></a>
##### skip_embedding_generation: `Boolean`<a id="skip_embedding_generation-boolean"></a>
##### set_page_as_boundary: `Boolean`<a id="set_page_as_boundary-boolean"></a>
##### embedding_model: [`EmbeddingGenerators`](./lib/carbon_ruby_sdk/models/embedding_generators.rb)<a id="embedding_model-embeddinggeneratorslibcarbon_ruby_sdkmodelsembedding_generatorsrb"></a>
##### generate_sparse_vectors: `Boolean`<a id="generate_sparse_vectors-boolean"></a>
##### use_textract: `Boolean`<a id="use_textract-boolean"></a>
##### prepend_filename_to_chunks: `Boolean`<a id="prepend_filename_to_chunks-boolean"></a>
##### max_items_per_chunk: `Integer`<a id="max_items_per_chunk-integer"></a>
Number of objects per chunk. For csv, tsv, xlsx, and json files only.

##### parse_pdf_tables_with_ocr: `Boolean`<a id="parse_pdf_tables_with_ocr-boolean"></a>
##### detect_audio_language: `Boolean`<a id="detect_audio_language-boolean"></a>
##### transcription_service: [`TranscriptionServiceNullable`](./lib/carbon_ruby_sdk/models/transcription_service_nullable.rb)<a id="transcription_service-transcriptionservicenullablelibcarbon_ruby_sdkmodelstranscription_service_nullablerb"></a>
##### include_speaker_labels: `Boolean`<a id="include_speaker_labels-boolean"></a>
##### media_type: [`FileContentTypesNullable`](./lib/carbon_ruby_sdk/models/file_content_types_nullable.rb)<a id="media_type-filecontenttypesnullablelibcarbon_ruby_sdkmodelsfile_content_types_nullablerb"></a>
##### split_rows: `Boolean`<a id="split_rows-boolean"></a>
##### cold_storage_params: [`ColdStorageProps`](./lib/carbon_ruby_sdk/models/cold_storage_props.rb)<a id="cold_storage_params-coldstoragepropslibcarbon_ruby_sdkmodelscold_storage_propsrb"></a>
#### 🔄 Return<a id="🔄-return"></a>

[UserFile](./lib/carbon_ruby_sdk/models/user_file.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/upload_file_from_url` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.upload_text`<a id="carbonfilesupload_text"></a>

Carbon supports multiple models for use in generating embeddings for files. For images, we support Vertex AI's
multimodal model; for text, we support OpenAI's `text-embedding-ada-002` and Cohere's embed-multilingual-v3.0.
The model can be specified via the `embedding_model` parameter (in the POST body for `/embeddings`, and a query 
parameter in `/uploadfile`). If no model is supplied, the `text-embedding-ada-002` is used by default. When performing
embedding queries, embeddings from files that used the specified model will be considered in the query.
For example, if files A and B have embeddings generated with `OPENAI`, and files C and D have embeddings generated with
`COHERE_MULTILINGUAL_V3`, then by default, queries will only consider files A and B. If `COHERE_MULTILINGUAL_V3` is
specified as the `embedding_model` in `/embeddings`, then only files C and D will be considered. Make sure that
the set of all files you want considered for a query have embeddings generated via the same model. For now, **do not**
set `VERTEX_MULTIMODAL` as an `embedding_model`. This model is used automatically by Carbon when it detects an image file.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.files.upload_text(
  contents: "aaaaa",
  name: "string_example",
  chunk_size: 1,
  chunk_overlap: 1,
  skip_embedding_generation: false,
  overwrite_file_id: 1,
  embedding_model: "OPENAI",
  generate_sparse_vectors: false,
  cold_storage_params: {
        "enable_cold_storage" => false,
    },
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### contents: `String`<a id="contents-string"></a>
##### name: `String`<a id="name-string"></a>
##### chunk_size: `Integer`<a id="chunk_size-integer"></a>
##### chunk_overlap: `Integer`<a id="chunk_overlap-integer"></a>
##### skip_embedding_generation: `Boolean`<a id="skip_embedding_generation-boolean"></a>
##### overwrite_file_id: `Integer`<a id="overwrite_file_id-integer"></a>
##### embedding_model: [`EmbeddingGeneratorsNullable`](./lib/carbon_ruby_sdk/models/embedding_generators_nullable.rb)<a id="embedding_model-embeddinggeneratorsnullablelibcarbon_ruby_sdkmodelsembedding_generators_nullablerb"></a>
##### generate_sparse_vectors: `Boolean`<a id="generate_sparse_vectors-boolean"></a>
##### cold_storage_params: [`ColdStorageProps`](./lib/carbon_ruby_sdk/models/cold_storage_props.rb)<a id="cold_storage_params-coldstoragepropslibcarbon_ruby_sdkmodelscold_storage_propsrb"></a>
#### 🔄 Return<a id="🔄-return"></a>

[UserFile](./lib/carbon_ruby_sdk/models/user_file.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/upload_text` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.cancel`<a id="carbonintegrationscancel"></a>

Cancel Data Source Items Sync

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.cancel(
  data_source_id: 1,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### data_source_id: `Integer`<a id="data_source_id-integer"></a>
#### 🔄 Return<a id="🔄-return"></a>

[OrganizationUserDataSourceAPI](./lib/carbon_ruby_sdk/models/organization_user_data_source_api.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/items/sync/cancel` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.connect_data_source`<a id="carbonintegrationsconnect_data_source"></a>

Connect Data Source

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.connect_data_source(
  authentication: {
        "source" => "GOOGLE_DRIVE",
        "access_token" => "access_token_example",
    },
  sync_options: {
        "chunk_size" => 1500,
        "chunk_overlap" => 20,
        "skip_embedding_generation" => false,
        "embedding_model" => "OPENAI",
        "generate_sparse_vectors" => false,
        "prepend_filename_to_chunks" => false,
        "sync_files_on_connection" => true,
        "set_page_as_boundary" => false,
        "enable_file_picker" => true,
        "sync_source_items" => true,
        "incremental_sync" => false,
    },
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### authentication: [`AuthenticationProperty`](./lib/carbon_ruby_sdk/models/authentication_property.rb)<a id="authentication-authenticationpropertylibcarbon_ruby_sdkmodelsauthentication_propertyrb"></a>
##### sync_options: [`SyncOptions`](./lib/carbon_ruby_sdk/models/sync_options.rb)<a id="sync_options-syncoptionslibcarbon_ruby_sdkmodelssync_optionsrb"></a>
#### 🔄 Return<a id="🔄-return"></a>

[ConnectDataSourceResponse](./lib/carbon_ruby_sdk/models/connect_data_source_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/connect` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.connect_freshdesk`<a id="carbonintegrationsconnect_freshdesk"></a>

Refer this article to obtain an API key https://support.freshdesk.com/en/support/solutions/articles/215517.
Make sure that your API key has the permission to read solutions from your account and you are on a <b>paid</b> plan.
Once you have an API key, you can make a request to this endpoint along with your freshdesk domain. This will 
trigger an automatic sync of the articles in your "solutions" tab. Additional parameters below can be used to associate 
data with the synced articles or modify the sync behavior.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.connect_freshdesk(
  domain: "string_example",
  api_key: "string_example",
  tags: {},
  chunk_size: 1500,
  chunk_overlap: 20,
  skip_embedding_generation: false,
  embedding_model: "OPENAI",
  generate_sparse_vectors: false,
  prepend_filename_to_chunks: false,
  sync_files_on_connection: true,
  request_id: "string_example",
  sync_source_items: true,
  file_sync_config: {
        "auto_synced_source_types" => ["ARTICLE"],
        "sync_attachments" => false,
        "detect_audio_language" => false,
        "transcription_service" => "assemblyai",
        "include_speaker_labels" => false,
        "split_rows" => false,
    },
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### domain: `String`<a id="domain-string"></a>
##### api_key: `String`<a id="api_key-string"></a>
##### tags: `Object`<a id="tags-object"></a>
##### chunk_size: `Integer`<a id="chunk_size-integer"></a>
##### chunk_overlap: `Integer`<a id="chunk_overlap-integer"></a>
##### skip_embedding_generation: `Boolean`<a id="skip_embedding_generation-boolean"></a>
##### embedding_model: [`EmbeddingGeneratorsNullable`](./lib/carbon_ruby_sdk/models/embedding_generators_nullable.rb)<a id="embedding_model-embeddinggeneratorsnullablelibcarbon_ruby_sdkmodelsembedding_generators_nullablerb"></a>
##### generate_sparse_vectors: `Boolean`<a id="generate_sparse_vectors-boolean"></a>
##### prepend_filename_to_chunks: `Boolean`<a id="prepend_filename_to_chunks-boolean"></a>
##### sync_files_on_connection: `Boolean`<a id="sync_files_on_connection-boolean"></a>
##### request_id: `String`<a id="request_id-string"></a>
##### sync_source_items: `Boolean`<a id="sync_source_items-boolean"></a>
Enabling this flag will fetch all available content from the source to be listed
via list items endpoint

##### file_sync_config: [`FileSyncConfigNullable`](./lib/carbon_ruby_sdk/models/file_sync_config_nullable.rb)<a id="file_sync_config-filesyncconfignullablelibcarbon_ruby_sdkmodelsfile_sync_config_nullablerb"></a>
#### 🔄 Return<a id="🔄-return"></a>

[GenericSuccessResponse](./lib/carbon_ruby_sdk/models/generic_success_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/freshdesk` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.connect_gitbook`<a id="carbonintegrationsconnect_gitbook"></a>

You will need an access token to connect your Gitbook account. Note that the permissions will be defined by the user 
generating access token so make sure you have the permission to access spaces you will be syncing. 
Refer this article for more details https://developer.gitbook.com/gitbook-api/authentication. Additionally, you
need to specify the name of organization you will be syncing data from.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.connect_gitbook(
  organization: "string_example",
  access_token: "string_example",
  tags: {},
  chunk_size: 1500,
  chunk_overlap: 20,
  skip_embedding_generation: false,
  embedding_model: "OPENAI",
  generate_sparse_vectors: false,
  prepend_filename_to_chunks: false,
  sync_files_on_connection: true,
  request_id: "string_example",
  sync_source_items: true,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### organization: `String`<a id="organization-string"></a>
##### access_token: `String`<a id="access_token-string"></a>
##### tags: `Object`<a id="tags-object"></a>
##### chunk_size: `Integer`<a id="chunk_size-integer"></a>
##### chunk_overlap: `Integer`<a id="chunk_overlap-integer"></a>
##### skip_embedding_generation: `Boolean`<a id="skip_embedding_generation-boolean"></a>
##### embedding_model: [`EmbeddingGenerators`](./lib/carbon_ruby_sdk/models/embedding_generators.rb)<a id="embedding_model-embeddinggeneratorslibcarbon_ruby_sdkmodelsembedding_generatorsrb"></a>
##### generate_sparse_vectors: `Boolean`<a id="generate_sparse_vectors-boolean"></a>
##### prepend_filename_to_chunks: `Boolean`<a id="prepend_filename_to_chunks-boolean"></a>
##### sync_files_on_connection: `Boolean`<a id="sync_files_on_connection-boolean"></a>
##### request_id: `String`<a id="request_id-string"></a>
##### sync_source_items: `Boolean`<a id="sync_source_items-boolean"></a>
Enabling this flag will fetch all available content from the source to be listed
via list items endpoint

#### 🔄 Return<a id="🔄-return"></a>

[GenericSuccessResponse](./lib/carbon_ruby_sdk/models/generic_success_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/gitbook` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.create_aws_iam_user`<a id="carbonintegrationscreate_aws_iam_user"></a>

This endpoint can be used to connect S3 as well as Digital Ocean Spaces (S3 compatible)  
For S3, create a new IAM user with permissions to:
<ol>
<li>List all buckets.</li>
<li>Read from the specific buckets and objects to sync with Carbon. Ensure any future buckets or objects carry 
the same permissions.</li>
</ol>
Once created, generate an access key for this user and share the credentials with us. We recommend testing this key beforehand.  
For Digital Ocean Spaces, generate the above credentials in your Applications and API page here https://cloud.digitalocean.com/account/api/spaces.
Endpoint URL is required to connect Digital Ocean Spaces.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.create_aws_iam_user(
  access_key: "string_example",
  access_key_secret: "string_example",
  sync_source_items: true,
  endpoint_url: "string_example",
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### access_key: `String`<a id="access_key-string"></a>
##### access_key_secret: `String`<a id="access_key_secret-string"></a>
##### sync_source_items: `Boolean`<a id="sync_source_items-boolean"></a>
Enabling this flag will fetch all available content from the source to be listed
via list items endpoint

##### endpoint_url: `String`<a id="endpoint_url-string"></a>
You can specify a Digital Ocean endpoint URL to connect a Digital Ocean Space
through this endpoint. The URL should be of format
<region>.digitaloceanspaces.com. It's not required for S3 buckets.

#### 🔄 Return<a id="🔄-return"></a>

[OrganizationUserDataSourceAPI](./lib/carbon_ruby_sdk/models/organization_user_data_source_api.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/s3` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.get_oauth_url`<a id="carbonintegrationsget_oauth_url"></a>

This endpoint can be used to generate the following URLs
- An OAuth URL for OAuth based connectors
- A file syncing URL which skips the OAuth flow if the user already has a valid access token and takes them to the
success state.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.get_oauth_url(
  service: "GOOGLE_CLOUD_STORAGE",
  tags: None,
  scope: "string_example",
  chunk_size: 1500,
  chunk_overlap: 20,
  skip_embedding_generation: false,
  embedding_model: "OPENAI",
  zendesk_subdomain: "string_example",
  microsoft_tenant: "string_example",
  sharepoint_site_name: "string_example",
  confluence_subdomain: "string_example",
  generate_sparse_vectors: false,
  prepend_filename_to_chunks: false,
  max_items_per_chunk: 1,
  salesforce_domain: "string_example",
  sync_files_on_connection: true,
  set_page_as_boundary: false,
  data_source_id: 1,
  connecting_new_account: false,
  request_id: "string_example",
  use_ocr: false,
  parse_pdf_tables_with_ocr: false,
  enable_file_picker: true,
  sync_source_items: true,
  incremental_sync: false,
  file_sync_config: {
        "auto_synced_source_types" => ["ARTICLE"],
        "sync_attachments" => false,
        "detect_audio_language" => false,
        "transcription_service" => "assemblyai",
        "include_speaker_labels" => false,
        "split_rows" => false,
    },
  automatically_open_file_picker: true,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### service: [`DataSourceType`](./lib/carbon_ruby_sdk/models/data_source_type.rb)<a id="service-datasourcetypelibcarbon_ruby_sdkmodelsdata_source_typerb"></a>
##### tags: `Object`<a id="tags-object"></a>
##### scope: `String`<a id="scope-string"></a>
##### chunk_size: `Integer`<a id="chunk_size-integer"></a>
##### chunk_overlap: `Integer`<a id="chunk_overlap-integer"></a>
##### skip_embedding_generation: `Boolean`<a id="skip_embedding_generation-boolean"></a>
##### embedding_model: [`EmbeddingGeneratorsNullable`](./lib/carbon_ruby_sdk/models/embedding_generators_nullable.rb)<a id="embedding_model-embeddinggeneratorsnullablelibcarbon_ruby_sdkmodelsembedding_generators_nullablerb"></a>
##### zendesk_subdomain: `String`<a id="zendesk_subdomain-string"></a>
##### microsoft_tenant: `String`<a id="microsoft_tenant-string"></a>
##### sharepoint_site_name: `String`<a id="sharepoint_site_name-string"></a>
##### confluence_subdomain: `String`<a id="confluence_subdomain-string"></a>
##### generate_sparse_vectors: `Boolean`<a id="generate_sparse_vectors-boolean"></a>
##### prepend_filename_to_chunks: `Boolean`<a id="prepend_filename_to_chunks-boolean"></a>
##### max_items_per_chunk: `Integer`<a id="max_items_per_chunk-integer"></a>
Number of objects per chunk. For csv, tsv, xlsx, and json files only.

##### salesforce_domain: `String`<a id="salesforce_domain-string"></a>
##### sync_files_on_connection: `Boolean`<a id="sync_files_on_connection-boolean"></a>
Used to specify whether Carbon should attempt to sync all your files
automatically when authorization is complete. This is only supported for a
subset of connectors and will be ignored for the rest. Supported connectors:
Intercom, Zendesk, Gitbook, Confluence, Salesforce, Freshdesk

##### set_page_as_boundary: `Boolean`<a id="set_page_as_boundary-boolean"></a>
##### data_source_id: `Integer`<a id="data_source_id-integer"></a>
Used to specify a data source to sync from if you have multiple connected. It
can be skipped if you only have one data source of that type connected or are
connecting a new account.

##### connecting_new_account: `Boolean`<a id="connecting_new_account-boolean"></a>
Used to connect a new data source. If not specified, we will attempt to create a
sync URL for an existing data source based on type and ID.

##### request_id: `String`<a id="request_id-string"></a>
This request id will be added to all files that get synced using the generated
OAuth URL

##### use_ocr: `Boolean`<a id="use_ocr-boolean"></a>
Enable OCR for files that support it. Supported formats: pdf, jpg, png

##### parse_pdf_tables_with_ocr: `Boolean`<a id="parse_pdf_tables_with_ocr-boolean"></a>
##### enable_file_picker: `Boolean`<a id="enable_file_picker-boolean"></a>
Enable integration's file picker for sources that support it. Supported sources:
BOX, DROPBOX, GOOGLE_DRIVE, ONEDRIVE, SHAREPOINT

##### sync_source_items: `Boolean`<a id="sync_source_items-boolean"></a>
Enabling this flag will fetch all available content from the source to be listed
via list items endpoint

##### incremental_sync: `Boolean`<a id="incremental_sync-boolean"></a>
Only sync files if they have not already been synced or if the embedding
properties have changed. This flag is currently supported by ONEDRIVE,
GOOGLE_DRIVE, BOX, DROPBOX, INTERCOM, GMAIL, OUTLOOK, ZENDESK, CONFLUENCE,
NOTION, SHAREPOINT. It will be ignored for other data sources.

##### file_sync_config: [`FileSyncConfigNullable`](./lib/carbon_ruby_sdk/models/file_sync_config_nullable.rb)<a id="file_sync_config-filesyncconfignullablelibcarbon_ruby_sdkmodelsfile_sync_config_nullablerb"></a>
##### automatically_open_file_picker: `Boolean`<a id="automatically_open_file_picker-boolean"></a>
Automatically open source file picker after the OAuth flow is complete. This
flag is currently supported by BOX, DROPBOX, GOOGLE_DRIVE, ONEDRIVE, SHAREPOINT.
It will be ignored for other data sources.

#### 🔄 Return<a id="🔄-return"></a>

[OuthURLResponse](./lib/carbon_ruby_sdk/models/outh_url_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/oauth_url` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.list_confluence_pages`<a id="carbonintegrationslist_confluence_pages"></a>
![Deprecated](https://img.shields.io/badge/deprecated-yellow)

This endpoint has been deprecated. Use /integrations/items/list instead.

To begin listing a user's Confluence pages, at least a `data_source_id` of a connected
Confluence account must be specified. This base request returns a list of root pages for
every space the user has access to in a Confluence instance. To traverse further down
the user's page directory, additional requests to this endpoint can be made with the same
`data_source_id` and with `parent_id` set to the id of page from a previous request. For
convenience, the `has_children` property in each directory item in the response list will
flag which pages will return non-empty lists of pages when set as the `parent_id`.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.list_confluence_pages(
  data_source_id: 1,
  parent_id: "string_example",
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### data_source_id: `Integer`<a id="data_source_id-integer"></a>
##### parent_id: `String`<a id="parent_id-string"></a>
#### 🔄 Return<a id="🔄-return"></a>

[ListResponse](./lib/carbon_ruby_sdk/models/list_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/confluence/list` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.list_conversations`<a id="carbonintegrationslist_conversations"></a>

List all of your public and private channels, DMs, and Group DMs. The ID from response 
can be used as a filter to sync messages to Carbon   
types: Comma separated list of types. Available types are im (DMs), mpim (group DMs), public_channel, and private_channel.
Defaults to public_channel.    
cursor: Used for pagination. If next_cursor is returned in response, you need to pass it as the cursor in the next request    
data_source_id: Data source needs to be specified if you have linked multiple slack accounts  
exclude_archived: Should archived conversations be excluded, defaults to true

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.list_conversations(
  types: "public_channel",
  cursor: "string_example",
  data_source_id: 1,
  exclude_archived: true,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### types: `String`<a id="types-string"></a>
##### cursor: `String`<a id="cursor-string"></a>
##### data_source_id: `Integer`<a id="data_source_id-integer"></a>
##### exclude_archived: `Boolean`<a id="exclude_archived-boolean"></a>
#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/slack/conversations` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.list_data_source_items`<a id="carbonintegrationslist_data_source_items"></a>

List Data Source Items

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.list_data_source_items(
  data_source_id: 1,
  parent_id: "string_example",
  filters: {
    },
  pagination: {
        "limit" => 10,
        "offset" => 0,
    },
  order_by: "name",
  order_dir: "asc",
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### data_source_id: `Integer`<a id="data_source_id-integer"></a>
##### parent_id: `String`<a id="parent_id-string"></a>
##### filters: [`ListItemsFiltersNullable`](./lib/carbon_ruby_sdk/models/list_items_filters_nullable.rb)<a id="filters-listitemsfiltersnullablelibcarbon_ruby_sdkmodelslist_items_filters_nullablerb"></a>
##### pagination: [`Pagination`](./lib/carbon_ruby_sdk/models/pagination.rb)<a id="pagination-paginationlibcarbon_ruby_sdkmodelspaginationrb"></a>
##### order_by: [`ExternalSourceItemsOrderBy`](./lib/carbon_ruby_sdk/models/external_source_items_order_by.rb)<a id="order_by-externalsourceitemsorderbylibcarbon_ruby_sdkmodelsexternal_source_items_order_byrb"></a>
##### order_dir: [`OrderDirV2`](./lib/carbon_ruby_sdk/models/order_dir_v2.rb)<a id="order_dir-orderdirv2libcarbon_ruby_sdkmodelsorder_dir_v2rb"></a>
#### 🔄 Return<a id="🔄-return"></a>

[ListDataSourceItemsResponse](./lib/carbon_ruby_sdk/models/list_data_source_items_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/items/list` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.list_folders`<a id="carbonintegrationslist_folders"></a>

After connecting your Outlook account, you can use this endpoint to list all of your folders on outlook. This includes 
both system folders like "inbox" and user created folders.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.list_folders(
  data_source_id: 1,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### data_source_id: `Integer`<a id="data_source_id-integer"></a>
#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/outlook/user_folders` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.list_gitbook_spaces`<a id="carbonintegrationslist_gitbook_spaces"></a>

After connecting your Gitbook account, you can use this endpoint to list all of your spaces under current organization.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.list_gitbook_spaces(
  data_source_id: 1,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### data_source_id: `Integer`<a id="data_source_id-integer"></a>
#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/gitbook/spaces` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.list_labels`<a id="carbonintegrationslist_labels"></a>

After connecting your Gmail account, you can use this endpoint to list all of your labels. User created labels
will have the type "user" and Gmail's default labels will have the type "system"

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.list_labels(
  data_source_id: 1,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### data_source_id: `Integer`<a id="data_source_id-integer"></a>
#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/gmail/user_labels` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.list_outlook_categories`<a id="carbonintegrationslist_outlook_categories"></a>

After connecting your Outlook account, you can use this endpoint to list all of your categories on outlook. We currently
support listing up to 250 categories.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.list_outlook_categories(
  data_source_id: 1,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### data_source_id: `Integer`<a id="data_source_id-integer"></a>
#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/outlook/user_categories` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.list_repos`<a id="carbonintegrationslist_repos"></a>

Once you have connected your GitHub account, you can use this endpoint to list the 
    repositories your account has access to. You can use a data source ID or username to fetch from a specific account.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.list_repos(
  per_page: 30,
  page: 1,
  data_source_id: 1,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### per_page: `Integer`<a id="per_page-integer"></a>
##### page: `Integer`<a id="page-integer"></a>
##### data_source_id: `Integer`<a id="data_source_id-integer"></a>
#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/github/repos` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.sync_confluence`<a id="carbonintegrationssync_confluence"></a>
![Deprecated](https://img.shields.io/badge/deprecated-yellow)

This endpoint has been deprecated. Use /integrations/files/sync instead.

After listing pages in a user's Confluence account, the set of selected page `ids` and the
connected account's `data_source_id` can be passed into this endpoint to sync them into
Carbon. Additional parameters listed below can be used to associate data to the selected
pages or alter the behavior of the sync.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.sync_confluence(
  data_source_id: 1,
  ids: [
        "string_example"
    ],
  tags: {},
  chunk_size: 1500,
  chunk_overlap: 20,
  skip_embedding_generation: false,
  embedding_model: "OPENAI",
  generate_sparse_vectors: false,
  prepend_filename_to_chunks: false,
  max_items_per_chunk: 1,
  set_page_as_boundary: false,
  request_id: "string_example",
  use_ocr: false,
  parse_pdf_tables_with_ocr: false,
  incremental_sync: false,
  file_sync_config: {
        "auto_synced_source_types" => ["ARTICLE"],
        "sync_attachments" => false,
        "detect_audio_language" => false,
        "transcription_service" => "assemblyai",
        "include_speaker_labels" => false,
        "split_rows" => false,
    },
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### data_source_id: `Integer`<a id="data_source_id-integer"></a>
##### ids: [`IdsProperty`](./lib/carbon_ruby_sdk/models/ids_property.rb)<a id="ids-idspropertylibcarbon_ruby_sdkmodelsids_propertyrb"></a>
##### tags: `Object`<a id="tags-object"></a>
##### chunk_size: `Integer`<a id="chunk_size-integer"></a>
##### chunk_overlap: `Integer`<a id="chunk_overlap-integer"></a>
##### skip_embedding_generation: `Boolean`<a id="skip_embedding_generation-boolean"></a>
##### embedding_model: [`EmbeddingGeneratorsNullable`](./lib/carbon_ruby_sdk/models/embedding_generators_nullable.rb)<a id="embedding_model-embeddinggeneratorsnullablelibcarbon_ruby_sdkmodelsembedding_generators_nullablerb"></a>
##### generate_sparse_vectors: `Boolean`<a id="generate_sparse_vectors-boolean"></a>
##### prepend_filename_to_chunks: `Boolean`<a id="prepend_filename_to_chunks-boolean"></a>
##### max_items_per_chunk: `Integer`<a id="max_items_per_chunk-integer"></a>
Number of objects per chunk. For csv, tsv, xlsx, and json files only.

##### set_page_as_boundary: `Boolean`<a id="set_page_as_boundary-boolean"></a>
##### request_id: `String`<a id="request_id-string"></a>
##### use_ocr: `Boolean`<a id="use_ocr-boolean"></a>
##### parse_pdf_tables_with_ocr: `Boolean`<a id="parse_pdf_tables_with_ocr-boolean"></a>
##### incremental_sync: `Boolean`<a id="incremental_sync-boolean"></a>
Only sync files if they have not already been synced or if the embedding
properties have changed. This flag is currently supported by ONEDRIVE,
GOOGLE_DRIVE, BOX, DROPBOX, INTERCOM, GMAIL, OUTLOOK, ZENDESK, CONFLUENCE,
NOTION, SHAREPOINT. It will be ignored for other data sources.

##### file_sync_config: [`FileSyncConfigNullable`](./lib/carbon_ruby_sdk/models/file_sync_config_nullable.rb)<a id="file_sync_config-filesyncconfignullablelibcarbon_ruby_sdkmodelsfile_sync_config_nullablerb"></a>
#### 🔄 Return<a id="🔄-return"></a>

[GenericSuccessResponse](./lib/carbon_ruby_sdk/models/generic_success_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/confluence/sync` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.sync_data_source_items`<a id="carbonintegrationssync_data_source_items"></a>

Sync Data Source Items

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.sync_data_source_items(
  data_source_id: 1,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### data_source_id: `Integer`<a id="data_source_id-integer"></a>
#### 🔄 Return<a id="🔄-return"></a>

[OrganizationUserDataSourceAPI](./lib/carbon_ruby_sdk/models/organization_user_data_source_api.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/items/sync` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.sync_files`<a id="carbonintegrationssync_files"></a>

After listing files and folders via /integrations/items/sync and integrations/items/list, use the selected items' external ids 
as the ids in this endpoint to sync them into Carbon. Sharepoint items take an additional parameter root_id, which identifies
the drive the file or folder is in and is stored in root_external_id. That additional paramter is optional and excluding it will
tell the sync to assume the item is stored in the default Documents drive.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.sync_files(
  data_source_id: 1,
  ids: [
        "string_example"
    ],
  tags: {},
  chunk_size: 1500,
  chunk_overlap: 20,
  skip_embedding_generation: false,
  embedding_model: "OPENAI",
  generate_sparse_vectors: false,
  prepend_filename_to_chunks: false,
  max_items_per_chunk: 1,
  set_page_as_boundary: false,
  request_id: "string_example",
  use_ocr: false,
  parse_pdf_tables_with_ocr: false,
  incremental_sync: false,
  file_sync_config: {
        "auto_synced_source_types" => ["ARTICLE"],
        "sync_attachments" => false,
        "detect_audio_language" => false,
        "transcription_service" => "assemblyai",
        "include_speaker_labels" => false,
        "split_rows" => false,
    },
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### data_source_id: `Integer`<a id="data_source_id-integer"></a>
##### ids: [`IdsProperty`](./lib/carbon_ruby_sdk/models/ids_property.rb)<a id="ids-idspropertylibcarbon_ruby_sdkmodelsids_propertyrb"></a>
##### tags: `Object`<a id="tags-object"></a>
##### chunk_size: `Integer`<a id="chunk_size-integer"></a>
##### chunk_overlap: `Integer`<a id="chunk_overlap-integer"></a>
##### skip_embedding_generation: `Boolean`<a id="skip_embedding_generation-boolean"></a>
##### embedding_model: [`EmbeddingGeneratorsNullable`](./lib/carbon_ruby_sdk/models/embedding_generators_nullable.rb)<a id="embedding_model-embeddinggeneratorsnullablelibcarbon_ruby_sdkmodelsembedding_generators_nullablerb"></a>
##### generate_sparse_vectors: `Boolean`<a id="generate_sparse_vectors-boolean"></a>
##### prepend_filename_to_chunks: `Boolean`<a id="prepend_filename_to_chunks-boolean"></a>
##### max_items_per_chunk: `Integer`<a id="max_items_per_chunk-integer"></a>
Number of objects per chunk. For csv, tsv, xlsx, and json files only.

##### set_page_as_boundary: `Boolean`<a id="set_page_as_boundary-boolean"></a>
##### request_id: `String`<a id="request_id-string"></a>
##### use_ocr: `Boolean`<a id="use_ocr-boolean"></a>
##### parse_pdf_tables_with_ocr: `Boolean`<a id="parse_pdf_tables_with_ocr-boolean"></a>
##### incremental_sync: `Boolean`<a id="incremental_sync-boolean"></a>
Only sync files if they have not already been synced or if the embedding
properties have changed. This flag is currently supported by ONEDRIVE,
GOOGLE_DRIVE, BOX, DROPBOX, INTERCOM, GMAIL, OUTLOOK, ZENDESK, CONFLUENCE,
NOTION, SHAREPOINT. It will be ignored for other data sources.

##### file_sync_config: [`FileSyncConfigNullable`](./lib/carbon_ruby_sdk/models/file_sync_config_nullable.rb)<a id="file_sync_config-filesyncconfignullablelibcarbon_ruby_sdkmodelsfile_sync_config_nullablerb"></a>
#### 🔄 Return<a id="🔄-return"></a>

[GenericSuccessResponse](./lib/carbon_ruby_sdk/models/generic_success_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/files/sync` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.sync_git_hub`<a id="carbonintegrationssync_git_hub"></a>

Refer this article to obtain an access token https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens.
Make sure that your access token has the permission to read content from your desired repos. Note that if your access token
expires you will need to manually update it through this endpoint.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.sync_git_hub(
  username: "string_example",
  access_token: "string_example",
  sync_source_items: false,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### username: `String`<a id="username-string"></a>
##### access_token: `String`<a id="access_token-string"></a>
##### sync_source_items: `Boolean`<a id="sync_source_items-boolean"></a>
Enabling this flag will fetch all available content from the source to be listed
via list items endpoint

#### 🔄 Return<a id="🔄-return"></a>

[GenericSuccessResponse](./lib/carbon_ruby_sdk/models/generic_success_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/github` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.sync_gitbook`<a id="carbonintegrationssync_gitbook"></a>

You can sync upto 20 Gitbook spaces at a time using this endpoint. Additional parameters below can be used to associate 
data with the synced pages or modify the sync behavior.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.sync_gitbook(
  space_ids: [
        "string_example"
    ],
  data_source_id: 1,
  tags: {},
  chunk_size: 1500,
  chunk_overlap: 20,
  skip_embedding_generation: false,
  embedding_model: "OPENAI",
  generate_sparse_vectors: false,
  prepend_filename_to_chunks: false,
  request_id: "string_example",
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### space_ids: Array<`String`><a id="space_ids-array"></a>
##### data_source_id: `Integer`<a id="data_source_id-integer"></a>
##### tags: `Object`<a id="tags-object"></a>
##### chunk_size: `Integer`<a id="chunk_size-integer"></a>
##### chunk_overlap: `Integer`<a id="chunk_overlap-integer"></a>
##### skip_embedding_generation: `Boolean`<a id="skip_embedding_generation-boolean"></a>
##### embedding_model: [`EmbeddingGenerators`](./lib/carbon_ruby_sdk/models/embedding_generators.rb)<a id="embedding_model-embeddinggeneratorslibcarbon_ruby_sdkmodelsembedding_generatorsrb"></a>
##### generate_sparse_vectors: `Boolean`<a id="generate_sparse_vectors-boolean"></a>
##### prepend_filename_to_chunks: `Boolean`<a id="prepend_filename_to_chunks-boolean"></a>
##### request_id: `String`<a id="request_id-string"></a>
#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/gitbook/sync` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.sync_gmail`<a id="carbonintegrationssync_gmail"></a>

Once you have successfully connected your gmail account, you can choose which emails to sync with us
using the filters parameter. Filters is a JSON object with key value pairs. It also supports AND and OR operations.
For now, we support a limited set of keys listed below.

<b>label</b>: Inbuilt Gmail labels, for example "Important" or a custom label you created.  
<b>after</b> or <b>before</b>: A date in YYYY/mm/dd format (example 2023/12/31). Gets emails after/before a certain date.
You can also use them in combination to get emails from a certain period.  
<b>is</b>: Can have the following values - starred, important, snoozed, and unread  
<b>from</b>: Email address of the sender  
<b>to</b>: Email address of the recipient  

Using keys or values outside of the specified values can lead to unexpected behaviour.

An example of a basic query with filters can be
```json
{
    "filters": {
            "key": "label",
            "value": "Test"
        }
}
```
Which will list all emails that have the label "Test".

You can use AND and OR operation in the following way:
```json
{
    "filters": {
        "AND": [
            {
                "key": "after",
                "value": "2024/01/07"
            },
            {
                "OR": [
                    {
                        "key": "label",
                        "value": "Personal"
                    },
                    {
                        "key": "is",
                        "value": "starred"
                    }
                ]
            }
        ]
    }
}
```
This will return emails after 7th of Jan that are either starred or have the label "Personal". 
Note that this is the highest level of nesting we support, i.e. you can't add more AND/OR filters within the OR filter
in the above example.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.sync_gmail(
  filters: {},
  tags: {},
  chunk_size: 1500,
  chunk_overlap: 20,
  skip_embedding_generation: false,
  embedding_model: "OPENAI",
  generate_sparse_vectors: false,
  prepend_filename_to_chunks: false,
  data_source_id: 1,
  request_id: "string_example",
  sync_attachments: false,
  file_sync_config: {
        "auto_synced_source_types" => ["ARTICLE"],
        "sync_attachments" => false,
        "detect_audio_language" => false,
        "transcription_service" => "assemblyai",
        "include_speaker_labels" => false,
        "split_rows" => false,
    },
  incremental_sync: false,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### filters: `Object`<a id="filters-object"></a>
##### tags: `Object`<a id="tags-object"></a>
##### chunk_size: `Integer`<a id="chunk_size-integer"></a>
##### chunk_overlap: `Integer`<a id="chunk_overlap-integer"></a>
##### skip_embedding_generation: `Boolean`<a id="skip_embedding_generation-boolean"></a>
##### embedding_model: [`EmbeddingGenerators`](./lib/carbon_ruby_sdk/models/embedding_generators.rb)<a id="embedding_model-embeddinggeneratorslibcarbon_ruby_sdkmodelsembedding_generatorsrb"></a>
##### generate_sparse_vectors: `Boolean`<a id="generate_sparse_vectors-boolean"></a>
##### prepend_filename_to_chunks: `Boolean`<a id="prepend_filename_to_chunks-boolean"></a>
##### data_source_id: `Integer`<a id="data_source_id-integer"></a>
##### request_id: `String`<a id="request_id-string"></a>
##### sync_attachments: `Boolean`<a id="sync_attachments-boolean"></a>
##### file_sync_config: [`FileSyncConfigNullable`](./lib/carbon_ruby_sdk/models/file_sync_config_nullable.rb)<a id="file_sync_config-filesyncconfignullablelibcarbon_ruby_sdkmodelsfile_sync_config_nullablerb"></a>
##### incremental_sync: `Boolean`<a id="incremental_sync-boolean"></a>
#### 🔄 Return<a id="🔄-return"></a>

[GenericSuccessResponse](./lib/carbon_ruby_sdk/models/generic_success_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/gmail/sync` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.sync_outlook`<a id="carbonintegrationssync_outlook"></a>

Once you have successfully connected your Outlook account, you can choose which emails to sync with us
using the filters and folder parameter. "folder" should be the folder you want to sync from Outlook. By default
we get messages from your inbox folder.  
Filters is a JSON object with key value pairs. It also supports AND and OR operations.
For now, we support a limited set of keys listed below.

<b>category</b>: Custom categories that you created in Outlook.  
<b>after</b> or <b>before</b>: A date in YYYY/mm/dd format (example 2023/12/31). Gets emails after/before a certain date. You can also use them in combination to get emails from a certain period.    
<b>is</b>: Can have the following values: flagged  
<b>from</b>: Email address of the sender   

An example of a basic query with filters can be
```json
{
    "filters": {
            "key": "category",
            "value": "Test"
        }
}
```
Which will list all emails that have the category "Test".  

Specifying a custom folder in the same query
```json
{
    "folder": "Folder Name",
    "filters": {
            "key": "category",
            "value": "Test"
        }
}
```

You can use AND and OR operation in the following way:
```json
{
    "filters": {
        "AND": [
            {
                "key": "after",
                "value": "2024/01/07"
            },
            {
                "OR": [
                    {
                        "key": "category",
                        "value": "Personal"
                    },
                    {
                        "key": "category",
                        "value": "Test"
                    },
                ]
            }
        ]
    }
}
```
This will return emails after 7th of Jan that have either Personal or Test as category. 
Note that this is the highest level of nesting we support, i.e. you can't add more AND/OR filters within the OR filter
in the above example.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.sync_outlook(
  filters: {},
  tags: {},
  folder: "Inbox",
  chunk_size: 1500,
  chunk_overlap: 20,
  skip_embedding_generation: false,
  embedding_model: "OPENAI",
  generate_sparse_vectors: false,
  prepend_filename_to_chunks: false,
  data_source_id: 1,
  request_id: "string_example",
  sync_attachments: false,
  file_sync_config: {
        "auto_synced_source_types" => ["ARTICLE"],
        "sync_attachments" => false,
        "detect_audio_language" => false,
        "transcription_service" => "assemblyai",
        "include_speaker_labels" => false,
        "split_rows" => false,
    },
  incremental_sync: false,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### filters: `Object`<a id="filters-object"></a>
##### tags: `Object`<a id="tags-object"></a>
##### folder: `String`<a id="folder-string"></a>
##### chunk_size: `Integer`<a id="chunk_size-integer"></a>
##### chunk_overlap: `Integer`<a id="chunk_overlap-integer"></a>
##### skip_embedding_generation: `Boolean`<a id="skip_embedding_generation-boolean"></a>
##### embedding_model: [`EmbeddingGenerators`](./lib/carbon_ruby_sdk/models/embedding_generators.rb)<a id="embedding_model-embeddinggeneratorslibcarbon_ruby_sdkmodelsembedding_generatorsrb"></a>
##### generate_sparse_vectors: `Boolean`<a id="generate_sparse_vectors-boolean"></a>
##### prepend_filename_to_chunks: `Boolean`<a id="prepend_filename_to_chunks-boolean"></a>
##### data_source_id: `Integer`<a id="data_source_id-integer"></a>
##### request_id: `String`<a id="request_id-string"></a>
##### sync_attachments: `Boolean`<a id="sync_attachments-boolean"></a>
##### file_sync_config: [`FileSyncConfigNullable`](./lib/carbon_ruby_sdk/models/file_sync_config_nullable.rb)<a id="file_sync_config-filesyncconfignullablelibcarbon_ruby_sdkmodelsfile_sync_config_nullablerb"></a>
##### incremental_sync: `Boolean`<a id="incremental_sync-boolean"></a>
#### 🔄 Return<a id="🔄-return"></a>

[GenericSuccessResponse](./lib/carbon_ruby_sdk/models/generic_success_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/outlook/sync` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.sync_repos`<a id="carbonintegrationssync_repos"></a>

You can retreive repos your token has access to using /integrations/github/repos and sync their content. 
You can also pass full name of any public repository (username/repo-name). This will store the repo content with 
carbon which can be accessed through /integrations/items/list endpoint. Maximum of 25 repositories are accepted per request.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.sync_repos(
  repos: [
        "string_example"
    ],
  data_source_id: 1,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### repos: Array<`String`><a id="repos-array"></a>
##### data_source_id: `Integer`<a id="data_source_id-integer"></a>
#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/github/sync_repos` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.sync_rss_feed`<a id="carbonintegrationssync_rss_feed"></a>

Rss Feed

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.sync_rss_feed(
  url: "string_example",
  tags: {},
  chunk_size: 1500,
  chunk_overlap: 20,
  skip_embedding_generation: false,
  embedding_model: "OPENAI",
  generate_sparse_vectors: false,
  prepend_filename_to_chunks: false,
  request_id: "string_example",
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### url: `String`<a id="url-string"></a>
##### tags: `Object`<a id="tags-object"></a>
##### chunk_size: `Integer`<a id="chunk_size-integer"></a>
##### chunk_overlap: `Integer`<a id="chunk_overlap-integer"></a>
##### skip_embedding_generation: `Boolean`<a id="skip_embedding_generation-boolean"></a>
##### embedding_model: [`EmbeddingGenerators`](./lib/carbon_ruby_sdk/models/embedding_generators.rb)<a id="embedding_model-embeddinggeneratorslibcarbon_ruby_sdkmodelsembedding_generatorsrb"></a>
##### generate_sparse_vectors: `Boolean`<a id="generate_sparse_vectors-boolean"></a>
##### prepend_filename_to_chunks: `Boolean`<a id="prepend_filename_to_chunks-boolean"></a>
##### request_id: `String`<a id="request_id-string"></a>
#### 🔄 Return<a id="🔄-return"></a>

[GenericSuccessResponse](./lib/carbon_ruby_sdk/models/generic_success_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/rss_feed` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.sync_s3_files`<a id="carbonintegrationssync_s3_files"></a>

After optionally loading the items via /integrations/items/sync and integrations/items/list, use the bucket name 
and object key as the ID in this endpoint to sync them into Carbon. Additional parameters below can associate 
data with the selected items or modify the sync behavior

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.sync_s3_files(
  ids: [
        {
        }
    ],
  tags: {},
  chunk_size: 1500,
  chunk_overlap: 20,
  skip_embedding_generation: false,
  embedding_model: "OPENAI",
  generate_sparse_vectors: false,
  prepend_filename_to_chunks: false,
  max_items_per_chunk: 1,
  set_page_as_boundary: false,
  data_source_id: 1,
  request_id: "string_example",
  use_ocr: false,
  parse_pdf_tables_with_ocr: false,
  file_sync_config: {
        "auto_synced_source_types" => ["ARTICLE"],
        "sync_attachments" => false,
        "detect_audio_language" => false,
        "transcription_service" => "assemblyai",
        "include_speaker_labels" => false,
        "split_rows" => false,
    },
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### ids: Array<[`S3GetFileInput`](./lib/carbon_ruby_sdk/models/s3_get_file_input.rb)><a id="ids-array"></a>
##### tags: `Object`<a id="tags-object"></a>
##### chunk_size: `Integer`<a id="chunk_size-integer"></a>
##### chunk_overlap: `Integer`<a id="chunk_overlap-integer"></a>
##### skip_embedding_generation: `Boolean`<a id="skip_embedding_generation-boolean"></a>
##### embedding_model: [`EmbeddingGenerators`](./lib/carbon_ruby_sdk/models/embedding_generators.rb)<a id="embedding_model-embeddinggeneratorslibcarbon_ruby_sdkmodelsembedding_generatorsrb"></a>
##### generate_sparse_vectors: `Boolean`<a id="generate_sparse_vectors-boolean"></a>
##### prepend_filename_to_chunks: `Boolean`<a id="prepend_filename_to_chunks-boolean"></a>
##### max_items_per_chunk: `Integer`<a id="max_items_per_chunk-integer"></a>
Number of objects per chunk. For csv, tsv, xlsx, and json files only.

##### set_page_as_boundary: `Boolean`<a id="set_page_as_boundary-boolean"></a>
##### data_source_id: `Integer`<a id="data_source_id-integer"></a>
##### request_id: `String`<a id="request_id-string"></a>
##### use_ocr: `Boolean`<a id="use_ocr-boolean"></a>
##### parse_pdf_tables_with_ocr: `Boolean`<a id="parse_pdf_tables_with_ocr-boolean"></a>
##### file_sync_config: [`FileSyncConfigNullable`](./lib/carbon_ruby_sdk/models/file_sync_config_nullable.rb)<a id="file_sync_config-filesyncconfignullablelibcarbon_ruby_sdkmodelsfile_sync_config_nullablerb"></a>
#### 🔄 Return<a id="🔄-return"></a>

[GenericSuccessResponse](./lib/carbon_ruby_sdk/models/generic_success_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/s3/files` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.sync_slack`<a id="carbonintegrationssync_slack"></a>

You can list all conversations using the endpoint /integrations/slack/conversations. The ID of 
conversation will be used as an input for this endpoint with timestamps as optional filters.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.integrations.sync_slack(
  filters: {
        "conversation_id" => "conversation_id_example",
    },
  tags: {},
  chunk_size: 1500,
  chunk_overlap: 20,
  skip_embedding_generation: false,
  embedding_model: "OPENAI",
  generate_sparse_vectors: false,
  prepend_filename_to_chunks: false,
  data_source_id: 1,
  request_id: "string_example",
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### filters: [`SlackFilters`](./lib/carbon_ruby_sdk/models/slack_filters.rb)<a id="filters-slackfilterslibcarbon_ruby_sdkmodelsslack_filtersrb"></a>
##### tags: `Object`<a id="tags-object"></a>
##### chunk_size: `Integer`<a id="chunk_size-integer"></a>
##### chunk_overlap: `Integer`<a id="chunk_overlap-integer"></a>
##### skip_embedding_generation: `Boolean`<a id="skip_embedding_generation-boolean"></a>
##### embedding_model: [`EmbeddingGenerators`](./lib/carbon_ruby_sdk/models/embedding_generators.rb)<a id="embedding_model-embeddinggeneratorslibcarbon_ruby_sdkmodelsembedding_generatorsrb"></a>
##### generate_sparse_vectors: `Boolean`<a id="generate_sparse_vectors-boolean"></a>
##### prepend_filename_to_chunks: `Boolean`<a id="prepend_filename_to_chunks-boolean"></a>
##### data_source_id: `Integer`<a id="data_source_id-integer"></a>
##### request_id: `String`<a id="request_id-string"></a>
#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/integrations/slack/sync` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.organizations.get`<a id="carbonorganizationsget"></a>

Get Organization

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.organizations.get
p result
```

#### 🔄 Return<a id="🔄-return"></a>

[OrganizationResponse](./lib/carbon_ruby_sdk/models/organization_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/organization` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.organizations.update`<a id="carbonorganizationsupdate"></a>

Update Organization

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.organizations.update(
  global_user_config: {
    },
  data_source_configs: {
        "key": {
            "allowed_file_formats" => [],
        },
    },
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### global_user_config: [`UserConfigurationNullable`](./lib/carbon_ruby_sdk/models/user_configuration_nullable.rb)<a id="global_user_config-userconfigurationnullablelibcarbon_ruby_sdkmodelsuser_configuration_nullablerb"></a>
##### data_source_configs: Hash<string, [`DataSourceConfiguration`](./lib/carbon_ruby_sdk/models/data_source_configuration.rb)><a id="data_source_configs-hash"></a>
Used to set organization level defaults for configuration related to data
sources.

#### 🔄 Return<a id="🔄-return"></a>

[GenericSuccessResponse](./lib/carbon_ruby_sdk/models/generic_success_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/organization/update` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.organizations.update_stats`<a id="carbonorganizationsupdate_stats"></a>

Use this endpoint to reaggregate the statistics for an organization, for example aggregate_file_size. The reaggregation
process is asyncronous so a webhook will be sent with the event type being FILE_STATISTICS_AGGREGATED to notify when the
process is complee. After this aggregation is complete, the updated statistics can be retrieved using the /organization
endpoint. The response of /organization willalso contain a timestamp of the last time the statistics were reaggregated.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.organizations.update_stats
p result
```

#### 🔄 Return<a id="🔄-return"></a>

[GenericSuccessResponse](./lib/carbon_ruby_sdk/models/generic_success_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/organization/statistics` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.users.delete`<a id="carbonusersdelete"></a>

Delete Users

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.users.delete(
  customer_ids: [
        "string_example"
    ],
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### customer_ids: Array<`String`><a id="customer_ids-array"></a>
#### 🔄 Return<a id="🔄-return"></a>

[GenericSuccessResponse](./lib/carbon_ruby_sdk/models/generic_success_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/delete_users` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.users.get`<a id="carbonusersget"></a>

User Endpoint

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.users.get(
  customer_id: "string_example",
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### customer_id: `String`<a id="customer_id-string"></a>
#### 🔄 Return<a id="🔄-return"></a>

[UserResponse](./lib/carbon_ruby_sdk/models/user_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/user` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.users.list`<a id="carbonuserslist"></a>

List users within an organization

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.users.list(
  pagination: {
        "limit" => 10,
        "offset" => 0,
    },
  filters: {
    },
  order_by: "created_at",
  order_dir: "asc",
  include_count: false,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### pagination: [`Pagination`](./lib/carbon_ruby_sdk/models/pagination.rb)<a id="pagination-paginationlibcarbon_ruby_sdkmodelspaginationrb"></a>
##### filters: [`ListUsersFilters`](./lib/carbon_ruby_sdk/models/list_users_filters.rb)<a id="filters-listusersfilterslibcarbon_ruby_sdkmodelslist_users_filtersrb"></a>
##### order_by: [`ListUsersOrderByTypes`](./lib/carbon_ruby_sdk/models/list_users_order_by_types.rb)<a id="order_by-listusersorderbytypeslibcarbon_ruby_sdkmodelslist_users_order_by_typesrb"></a>
##### order_dir: [`OrderDirV2`](./lib/carbon_ruby_sdk/models/order_dir_v2.rb)<a id="order_dir-orderdirv2libcarbon_ruby_sdkmodelsorder_dir_v2rb"></a>
##### include_count: `Boolean`<a id="include_count-boolean"></a>
#### 🔄 Return<a id="🔄-return"></a>

[UserListResponse](./lib/carbon_ruby_sdk/models/user_list_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/list_users` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.users.toggle_user_features`<a id="carbonuserstoggle_user_features"></a>
![Deprecated](https://img.shields.io/badge/deprecated-yellow)

Toggle User Features

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.users.toggle_user_features(
  configuration_key_name: "string_example",
  value: {},
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### configuration_key_name: `String`<a id="configuration_key_name-string"></a>
##### value: `Object`<a id="value-object"></a>
#### 🔄 Return<a id="🔄-return"></a>

[GenericSuccessResponse](./lib/carbon_ruby_sdk/models/generic_success_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/modify_user_configuration` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.users.update_users`<a id="carbonusersupdate_users"></a>

Update Users

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.users.update_users(
  customer_ids: [
        "string_example"
    ],
  auto_sync_enabled_sources: [
        "string_example"
    ],
  max_files: -1,
  max_files_per_upload: -1,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### customer_ids: Array<`String`><a id="customer_ids-array"></a>
List of organization supplied user IDs

##### auto_sync_enabled_sources: [`AutoSyncEnabledSourcesProperty`](./lib/carbon_ruby_sdk/models/auto_sync_enabled_sources_property.rb)<a id="auto_sync_enabled_sources-autosyncenabledsourcespropertylibcarbon_ruby_sdkmodelsauto_sync_enabled_sources_propertyrb"></a>
##### max_files: `Integer`<a id="max_files-integer"></a>
Custom file upload limit for the user over *all* user's files across all
uploads. If set, then the user will not be allowed to upload more files than
this limit. If not set, or if set to -1, then the user will have no limit.

##### max_files_per_upload: `Integer`<a id="max_files_per_upload-integer"></a>
Custom file upload limit for the user across a single upload. If set, then the
user will not be allowed to upload more files than this limit in a single
upload. If not set, or if set to -1, then the user will have no limit.

#### 🔄 Return<a id="🔄-return"></a>

[GenericSuccessResponse](./lib/carbon_ruby_sdk/models/generic_success_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/update_users` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.utilities.fetch_urls`<a id="carbonutilitiesfetch_urls"></a>
![Deprecated](https://img.shields.io/badge/deprecated-yellow)

Extracts all URLs from a webpage. 

Args:
    url (str): URL of the webpage

Returns:
    FetchURLsResponse: A response object with a list of URLs extracted from the webpage and the webpage content.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.utilities.fetch_urls(
  url: "url_example",
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### url: `String`<a id="url-string"></a>
#### 🔄 Return<a id="🔄-return"></a>

[FetchURLsResponse](./lib/carbon_ruby_sdk/models/fetch_urls_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/fetch_urls` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.utilities.fetch_webpage`<a id="carbonutilitiesfetch_webpage"></a>

Fetch Urls V2

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.utilities.fetch_webpage(
  url: "string_example",
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### url: `String`<a id="url-string"></a>
#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/fetch_webpage` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.utilities.fetch_youtube_transcripts`<a id="carbonutilitiesfetch_youtube_transcripts"></a>

Fetches english transcripts from YouTube videos.

Args:
    id (str): The ID of the YouTube video. 
    raw (bool): Whether to return the raw transcript or not. Defaults to False.

Returns:
    dict: A dictionary with the transcript of the YouTube video.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.utilities.fetch_youtube_transcripts(
  id: "id_example",
  raw: false,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### id: `String`<a id="id-string"></a>
##### raw: `Boolean`<a id="raw-boolean"></a>
#### 🔄 Return<a id="🔄-return"></a>

[YoutubeTranscriptResponse](./lib/carbon_ruby_sdk/models/youtube_transcript_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/fetch_youtube_transcript` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.utilities.process_sitemap`<a id="carbonutilitiesprocess_sitemap"></a>

Retrieves all URLs from a sitemap, which can subsequently be utilized with our `web_scrape` endpoint.

<!--Args:
    url (str): URL of the sitemap

Returns:
    dict: A dictionary with a list of URLs extracted from the sitemap.-->

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.utilities.process_sitemap(
  url: "url_example",
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### url: `String`<a id="url-string"></a>
#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/process_sitemap` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.utilities.scrape_sitemap`<a id="carbonutilitiesscrape_sitemap"></a>

Extracts all URLs from a sitemap and performs a web scrape on each of them.

Args:
    sitemap_url (str): URL of the sitemap

Returns:
    dict: A response object with the status of the scraping job message.-->

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.utilities.scrape_sitemap(
  url: "string_example",
  tags: {
        "key": "string_example",
    },
  max_pages_to_scrape: 1,
  chunk_size: 1500,
  chunk_overlap: 20,
  skip_embedding_generation: false,
  enable_auto_sync: false,
  generate_sparse_vectors: false,
  prepend_filename_to_chunks: false,
  html_tags_to_skip: [],
  css_classes_to_skip: [],
  css_selectors_to_skip: [],
  embedding_model: "OPENAI",
  url_paths_to_include: [],
  url_paths_to_exclude: [],
  urls_to_scrape: [],
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### url: `String`<a id="url-string"></a>
##### tags: Hash<string, [`Tags1`](./lib/carbon_ruby_sdk/models/tags1.rb)><a id="tags-hash"></a>
##### max_pages_to_scrape: `Integer`<a id="max_pages_to_scrape-integer"></a>
##### chunk_size: `Integer`<a id="chunk_size-integer"></a>
##### chunk_overlap: `Integer`<a id="chunk_overlap-integer"></a>
##### skip_embedding_generation: `Boolean`<a id="skip_embedding_generation-boolean"></a>
##### enable_auto_sync: `Boolean`<a id="enable_auto_sync-boolean"></a>
##### generate_sparse_vectors: `Boolean`<a id="generate_sparse_vectors-boolean"></a>
##### prepend_filename_to_chunks: `Boolean`<a id="prepend_filename_to_chunks-boolean"></a>
##### html_tags_to_skip: Array<`String`><a id="html_tags_to_skip-array"></a>
##### css_classes_to_skip: Array<`String`><a id="css_classes_to_skip-array"></a>
##### css_selectors_to_skip: Array<`String`><a id="css_selectors_to_skip-array"></a>
##### embedding_model: [`EmbeddingGenerators`](./lib/carbon_ruby_sdk/models/embedding_generators.rb)<a id="embedding_model-embeddinggeneratorslibcarbon_ruby_sdkmodelsembedding_generatorsrb"></a>
##### url_paths_to_include: Array<`String`><a id="url_paths_to_include-array"></a>
URL subpaths or directories that you want to include. For example if you want to
only include URLs that start with /questions in stackoverflow.com, you will add
/questions/ in this input

##### url_paths_to_exclude: Array<`String`><a id="url_paths_to_exclude-array"></a>
URL subpaths or directories that you want to exclude. For example if you want to
exclude URLs that start with /questions in stackoverflow.com, you will add
/questions/ in this input

##### urls_to_scrape: Array<`String`><a id="urls_to_scrape-array"></a>
You can submit a subset of URLs from the sitemap that should be scraped. To get
the list of URLs, you can check out /process_sitemap endpoint. If left empty,
all URLs from the sitemap will be scraped.

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/scrape_sitemap` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.utilities.scrape_web`<a id="carbonutilitiesscrape_web"></a>

Conduct a web scrape on a given webpage URL. Our web scraper is fully compatible with JavaScript and supports recursion depth, enabling you to efficiently extract all content from the target website.

<!--Args:
    scraping_requests (List[WebscrapeRequest]): A list of WebscrapeRequest objects.

    
Returns:
    dict: A response object with the status of the scraping job message.-->

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.utilities.scrape_web(
  body: [
        {
            "url" => "url_example",
            "recursion_depth" => 3,
            "max_pages_to_scrape" => 100,
            "chunk_size" => 1500,
            "chunk_overlap" => 20,
            "skip_embedding_generation" => false,
            "enable_auto_sync" => false,
            "generate_sparse_vectors" => false,
            "prepend_filename_to_chunks" => false,
            "html_tags_to_skip" => [],
            "css_classes_to_skip" => [],
            "css_selectors_to_skip" => [],
            "embedding_model" => "OPENAI",
            "url_paths_to_include" => [],
        }
    ],
)
p result
```

#### ⚙️ body<a id="⚙️-body"></a>

Array<[`WebscrapeRequest`](./lib/carbon_ruby_sdk/models/webscrape_request.rb)>
#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/web_scrape` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.utilities.search_urls`<a id="carbonutilitiessearch_urls"></a>

Perform a web search and obtain a list of relevant URLs.

As an illustration, when you perform a search for “content related to MRNA,” you will receive a list of links such as the following:

    - https://tomrenz.substack.com/p/mrna-and-why-it-matters

    - https://www.statnews.com/2020/11/10/the-story-of-mrna-how-a-once-dismissed-idea-became-a-leading-technology-in-the-covid-vaccine-race/

    - https://www.statnews.com/2022/11/16/covid-19-vaccines-were-a-success-but-mrna-still-has-a-delivery-problem/
    
    - https://joomi.substack.com/p/were-still-being-misled-about-how

Subsequently, you can submit these links to the web_scrape endpoint in order to retrieve the content of the respective web pages.

Args:
    query (str): Query to search for

Returns:
    FetchURLsResponse: A response object with a list of URLs for a given search query.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.utilities.search_urls(
  query: "query_example",
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### query: `String`<a id="query-string"></a>
#### 🔄 Return<a id="🔄-return"></a>

[FetchURLsResponse](./lib/carbon_ruby_sdk/models/fetch_urls_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/search_urls` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.utilities.user_webpages`<a id="carbonutilitiesuser_webpages"></a>

User Web Pages

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.utilities.user_webpages(
  filters: {
    },
  pagination: {
        "limit" => 10,
        "offset" => 0,
    },
  order_by: "created_at",
  order_dir: "asc",
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### filters: [`UserWebPagesFilters`](./lib/carbon_ruby_sdk/models/user_web_pages_filters.rb)<a id="filters-userwebpagesfilterslibcarbon_ruby_sdkmodelsuser_web_pages_filtersrb"></a>
##### pagination: [`Pagination`](./lib/carbon_ruby_sdk/models/pagination.rb)<a id="pagination-paginationlibcarbon_ruby_sdkmodelspaginationrb"></a>
##### order_by: [`UserWebPageOrderByTypes`](./lib/carbon_ruby_sdk/models/user_web_page_order_by_types.rb)<a id="order_by-userwebpageorderbytypeslibcarbon_ruby_sdkmodelsuser_web_page_order_by_typesrb"></a>
##### order_dir: [`OrderDirV2`](./lib/carbon_ruby_sdk/models/order_dir_v2.rb)<a id="order_dir-orderdirv2libcarbon_ruby_sdkmodelsorder_dir_v2rb"></a>
#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/user_webpages` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.webhooks.add_url`<a id="carbonwebhooksadd_url"></a>

Add Webhook Url

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.webhooks.add_url(
  url: "string_example",
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### url: `String`<a id="url-string"></a>
#### 🔄 Return<a id="🔄-return"></a>

[Webhook](./lib/carbon_ruby_sdk/models/webhook.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/add_webhook` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.webhooks.delete_url`<a id="carbonwebhooksdelete_url"></a>

Delete Webhook Url

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.webhooks.delete_url(
  webhook_id: 1,
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### webhook_id: `Integer`<a id="webhook_id-integer"></a>
#### 🔄 Return<a id="🔄-return"></a>

[GenericSuccessResponse](./lib/carbon_ruby_sdk/models/generic_success_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/delete_webhook/{webhook_id}` `DELETE`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.webhooks.urls`<a id="carbonwebhooksurls"></a>

Webhook Urls

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
result = carbon.webhooks.urls(
  pagination: {
        "limit" => 10,
        "offset" => 0,
    },
  order_by: "created_at",
  order_dir: "desc",
  filters: {
        "ids" => [],
    },
)
p result
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### pagination: [`Pagination`](./lib/carbon_ruby_sdk/models/pagination.rb)<a id="pagination-paginationlibcarbon_ruby_sdkmodelspaginationrb"></a>
##### order_by: [`WebhookOrderByColumns`](./lib/carbon_ruby_sdk/models/webhook_order_by_columns.rb)<a id="order_by-webhookorderbycolumnslibcarbon_ruby_sdkmodelswebhook_order_by_columnsrb"></a>
##### order_dir: [`OrderDir`](./lib/carbon_ruby_sdk/models/order_dir.rb)<a id="order_dir-orderdirlibcarbon_ruby_sdkmodelsorder_dirrb"></a>
##### filters: [`WebhookFilters`](./lib/carbon_ruby_sdk/models/webhook_filters.rb)<a id="filters-webhookfilterslibcarbon_ruby_sdkmodelswebhook_filtersrb"></a>
#### 🔄 Return<a id="🔄-return"></a>

[WebhookQueryResponse](./lib/carbon_ruby_sdk/models/webhook_query_response.rb)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/webhooks` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


## Author<a id="author"></a>
This TypeScript package is automatically generated by [Konfig](https://konfigthis.com)

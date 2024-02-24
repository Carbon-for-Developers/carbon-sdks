<div align="center">

[![Visit Carbon](./header.png)](https://carbon.ai)

# [Carbon](https://carbon.ai)<a id="carbon"></a>

Connect external data to LLMs, no matter the source.

[![npm](https://img.shields.io/badge/gem-v0.1.0-blue)](https://rubygems.org/gems/carbon/versions/0.1.0)

</div>

## Table of Contents<a id="table-of-contents"></a>

<!-- toc -->

- [Installation](#installation)
- [Getting Started](#getting-started)
- [Reference](#reference)
  * [`carbon.auth.get_access_token`](#carbonauthget_access_token)
  * [`carbon.auth.get_white_labeling`](#carbonauthget_white_labeling)
  * [`carbon.data_sources.query_user_data_sources`](#carbondata_sourcesquery_user_data_sources)
  * [`carbon.data_sources.revoke_access_token`](#carbondata_sourcesrevoke_access_token)
  * [`carbon.embeddings.get_documents`](#carbonembeddingsget_documents)
  * [`carbon.embeddings.get_embeddings_and_chunks`](#carbonembeddingsget_embeddings_and_chunks)
  * [`carbon.embeddings.upload_chunks_and_embeddings`](#carbonembeddingsupload_chunks_and_embeddings)
  * [`carbon.files.create_user_file_tags`](#carbonfilescreate_user_file_tags)
  * [`carbon.files.delete`](#carbonfilesdelete)
  * [`carbon.files.delete_file_tags`](#carbonfilesdelete_file_tags)
  * [`carbon.files.delete_many`](#carbonfilesdelete_many)
  * [`carbon.files.get_parsed_file`](#carbonfilesget_parsed_file)
  * [`carbon.files.get_raw_file`](#carbonfilesget_raw_file)
  * [`carbon.files.query_user_files`](#carbonfilesquery_user_files)
  * [`carbon.files.query_user_files_deprecated`](#carbonfilesquery_user_files_deprecated)
  * [`carbon.files.resync`](#carbonfilesresync)
  * [`carbon.files.upload`](#carbonfilesupload)
  * [`carbon.files.upload_from_url`](#carbonfilesupload_from_url)
  * [`carbon.files.upload_text`](#carbonfilesupload_text)
  * [`carbon.health.check`](#carbonhealthcheck)
  * [`carbon.integrations.connect_freshdesk`](#carbonintegrationsconnect_freshdesk)
  * [`carbon.integrations.create_aws_iam_user`](#carbonintegrationscreate_aws_iam_user)
  * [`carbon.integrations.get_oauth_url`](#carbonintegrationsget_oauth_url)
  * [`carbon.integrations.list_confluence_pages`](#carbonintegrationslist_confluence_pages)
  * [`carbon.integrations.list_data_source_items`](#carbonintegrationslist_data_source_items)
  * [`carbon.integrations.list_folders`](#carbonintegrationslist_folders)
  * [`carbon.integrations.list_labels`](#carbonintegrationslist_labels)
  * [`carbon.integrations.sync_confluence`](#carbonintegrationssync_confluence)
  * [`carbon.integrations.sync_data_source_items`](#carbonintegrationssync_data_source_items)
  * [`carbon.integrations.sync_files`](#carbonintegrationssync_files)
  * [`carbon.integrations.sync_gmail`](#carbonintegrationssync_gmail)
  * [`carbon.integrations.sync_outlook`](#carbonintegrationssync_outlook)
  * [`carbon.integrations.sync_rss_feed`](#carbonintegrationssync_rss_feed)
  * [`carbon.integrations.sync_s3_files`](#carbonintegrationssync_s3_files)
  * [`carbon.organizations.get`](#carbonorganizationsget)
  * [`carbon.users.get`](#carbonusersget)
  * [`carbon.users.toggle_user_features`](#carbonuserstoggle_user_features)
  * [`carbon.utilities.fetch_urls`](#carbonutilitiesfetch_urls)
  * [`carbon.utilities.fetch_youtube_transcripts`](#carbonutilitiesfetch_youtube_transcripts)
  * [`carbon.utilities.process_sitemap`](#carbonutilitiesprocess_sitemap)
  * [`carbon.utilities.scrape_sitemap`](#carbonutilitiesscrape_sitemap)
  * [`carbon.utilities.scrape_web`](#carbonutilitiesscrape_web)
  * [`carbon.utilities.search_urls`](#carbonutilitiessearch_urls)
  * [`carbon.webhooks.add_url`](#carbonwebhooksadd_url)
  * [`carbon.webhooks.delete_url`](#carbonwebhooksdelete_url)
  * [`carbon.webhooks.urls`](#carbonwebhooksurls)

<!-- tocstop -->

## Installation<a id="installation"></a>

Add to Gemfile:

```ruby
gem 'carbon', '~> 0.1.0'
```

## Getting Started<a id="getting-started"></a>

```ruby
require 'carbon'
configuration = Carbon::Configuration.new
configuration.access_token = 'YOUR API KEY'
configuration.api_key = 'YOUR API KEY'
configuration.customer_id = 'YOUR API KEY'
carbon = Carbon::Client.new(configuration)

result = carbon.auth.get_access_token
p result
```

## Reference<a id="reference"></a>


### `carbon.auth.get_access_token`<a id="carbonauthget_access_token"></a>

Get Access Token

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby

result = carbon.auth.get_access_token
p result
```

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

---


### `carbon.data_sources.query_user_data_sources`<a id="carbondata_sourcesquery_user_data_sources"></a>

User Data Sources

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
pagination = {
        "limit" => 10,
        "offset" => 0,
    }
order_by = "created_at"
order_dir = "desc"
filters = {
        "source" => "GOOGLE_DRIVE",
    }

result = carbon.data_sources.query_user_data_sources(
  pagination: pagination,
  order_by: order_by,
  order_dir: order_dir,
  filters: filters,
)
p result
```

---


### `carbon.data_sources.revoke_access_token`<a id="carbondata_sourcesrevoke_access_token"></a>

Revoke Access Token

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
data_source_id = 1

result = carbon.data_sources.revoke_access_token(
  data_source_id: data_source_id,
)
p result
```

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

result = carbon.embeddings.get_documents(
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
```

---


### `carbon.embeddings.get_embeddings_and_chunks`<a id="carbonembeddingsget_embeddings_and_chunks"></a>

Retrieve Embeddings And Content

#### 🛠️ Usage<a id="🛠️-usage"></a>

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

result = carbon.embeddings.get_embeddings_and_chunks(
  filters: filters,
  pagination: pagination,
  order_by: order_by,
  order_dir: order_dir,
  include_vectors: include_vectors,
)
p result
```

---


### `carbon.embeddings.upload_chunks_and_embeddings`<a id="carbonembeddingsupload_chunks_and_embeddings"></a>

Upload Chunks And Embeddings

#### 🛠️ Usage<a id="🛠️-usage"></a>

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

result = carbon.embeddings.upload_chunks_and_embeddings(
  embedding_model: embedding_model,
  chunks_and_embeddings: chunks_and_embeddings,
  overwrite_existing: overwrite_existing,
)
p result
```

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
tags = {
        "key": "string_example",
    }
organization_user_file_id = 1

result = carbon.files.create_user_file_tags(
  tags: tags,
  organization_user_file_id: organization_user_file_id,
)
p result
```

---


### `carbon.files.delete`<a id="carbonfilesdelete"></a>

Delete File Endpoint

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
file_id = 1

result = carbon.files.delete(
  file_id: file_id,
)
p result
```

---


### `carbon.files.delete_file_tags`<a id="carbonfilesdelete_file_tags"></a>

Delete File Tags

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
tags = [
        "string_example"
    ]
organization_user_file_id = 1

result = carbon.files.delete_file_tags(
  tags: tags,
  organization_user_file_id: organization_user_file_id,
)
p result
```

---


### `carbon.files.delete_many`<a id="carbonfilesdelete_many"></a>

Delete Files Endpoint

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
file_ids = [
        1
    ]
sync_statuses = [
        "string_example"
    ]
delete_non_synced_only = False

result = carbon.files.delete_many(
  file_ids: file_ids,
  sync_statuses: sync_statuses,
  delete_non_synced_only: delete_non_synced_only,
)
p result
```

---


### `carbon.files.get_parsed_file`<a id="carbonfilesget_parsed_file"></a>

This route is deprecated. Use `/user_files_v2` instead.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
file_id = 1

result = carbon.files.get_parsed_file(
  file_id: file_id,
)
p result
```

---


### `carbon.files.get_raw_file`<a id="carbonfilesget_raw_file"></a>

This route is deprecated. Use `/user_files_v2` instead.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
file_id = 1

result = carbon.files.get_raw_file(
  file_id: file_id,
)
p result
```

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

result = carbon.files.query_user_files(
  pagination: pagination,
  order_by: order_by,
  order_dir: order_dir,
  filters: filters,
  include_raw_file: include_raw_file,
  include_parsed_text_file: include_parsed_text_file,
  include_additional_files: include_additional_files,
)
p result
```

---


### `carbon.files.query_user_files_deprecated`<a id="carbonfilesquery_user_files_deprecated"></a>

This route is deprecated. Use `/user_files_v2` instead.

#### 🛠️ Usage<a id="🛠️-usage"></a>

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

result = carbon.files.query_user_files_deprecated(
  pagination: pagination,
  order_by: order_by,
  order_dir: order_dir,
  filters: filters,
  include_raw_file: include_raw_file,
  include_parsed_text_file: include_parsed_text_file,
  include_additional_files: include_additional_files,
)
p result
```

---


### `carbon.files.resync`<a id="carbonfilesresync"></a>

Resync File

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
file_id = 1
chunk_size = 1
chunk_overlap = 1

result = carbon.files.resync(
  file_id: file_id,
  chunk_size: chunk_size,
  chunk_overlap: chunk_overlap,
)
p result
```

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
- `use_ocr`: whether or not to use OCR as a preprocessing step prior to generating chunks (only valid for PDFs currently)
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

result = carbon.files.upload(
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
```

---


### `carbon.files.upload_from_url`<a id="carbonfilesupload_from_url"></a>

Create Upload File From Url

#### 🛠️ Usage<a id="🛠️-usage"></a>

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

result = carbon.files.upload_from_url(
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
```

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
contents = "string_example"
name = "string_example"
chunk_size = 1
chunk_overlap = 1
skip_embedding_generation = False
overwrite_file_id = 1
embedding_model = "OPENAI"
generate_sparse_vectors = False

result = carbon.files.upload_text(
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
```

---


### `carbon.health.check`<a id="carbonhealthcheck"></a>

Health

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby

result = carbon.health.check
p result
```

---


### `carbon.integrations.connect_freshdesk`<a id="carbonintegrationsconnect_freshdesk"></a>

Refer this article to obtain an API key https://support.freshdesk.com/en/support/solutions/articles/215517.
Make sure that your API key has the permission to read solutions from your account and you are on a <b>paid</b> plan.
Once you have an API key, you can make a request to this endpoint along with your freshdesk domain. This will 
trigger an automatic sync of the articles in your "solutions" tab. Additional parameters below can be used to associate 
data with the synced articles or modify the sync behavior.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
domain = "string_example"
api_key = "string_example"
tags = {}
chunk_size = 1500
chunk_overlap = 20
skip_embedding_generation = False
embedding_model = "OPENAI"
generate_sparse_vectors = False
prepend_filename_to_chunks = False

result = carbon.integrations.connect_freshdesk(
  domain: domain,
  api_key: api_key,
  tags: tags,
  chunk_size: chunk_size,
  chunk_overlap: chunk_overlap,
  skip_embedding_generation: skip_embedding_generation,
  embedding_model: embedding_model,
  generate_sparse_vectors: generate_sparse_vectors,
  prepend_filename_to_chunks: prepend_filename_to_chunks,
)
p result
```

---


### `carbon.integrations.create_aws_iam_user`<a id="carbonintegrationscreate_aws_iam_user"></a>

Create a new IAM user with permissions to:
<ol>
<li>List all buckets.</li>
<li>Read from the specific buckets and objects to sync with Carbon. Ensure any future buckets or objects carry 
the same permissions.</li>
</ol>
Once created, generate an access key for this user and share the credentials with us. We recommend testing this key beforehand.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
access_key = "string_example"
access_key_secret = "string_example"

result = carbon.integrations.create_aws_iam_user(
  access_key: access_key,
  access_key_secret: access_key_secret,
)
p result
```

---


### `carbon.integrations.get_oauth_url`<a id="carbonintegrationsget_oauth_url"></a>

Get Oauth Url

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
service = "GOOGLE_DRIVE"
tags = None
scope = "string_example"
chunk_size = 1500
chunk_overlap = 20
skip_embedding_generation = False
embedding_model = "OPENAI"
zendesk_subdomain = "string_example"
microsoft_tenant = "string_example"
sharepoint_site_name = "string_example"
confluence_subdomain = "string_example"
generate_sparse_vectors = False
prepend_filename_to_chunks = False
max_items_per_chunk = 1

result = carbon.integrations.get_oauth_url(
  service: service,
  tags: tags,
  scope: scope,
  chunk_size: chunk_size,
  chunk_overlap: chunk_overlap,
  skip_embedding_generation: skip_embedding_generation,
  embedding_model: embedding_model,
  zendesk_subdomain: zendesk_subdomain,
  microsoft_tenant: microsoft_tenant,
  sharepoint_site_name: sharepoint_site_name,
  confluence_subdomain: confluence_subdomain,
  generate_sparse_vectors: generate_sparse_vectors,
  prepend_filename_to_chunks: prepend_filename_to_chunks,
  max_items_per_chunk: max_items_per_chunk,
)
p result
```

---


### `carbon.integrations.list_confluence_pages`<a id="carbonintegrationslist_confluence_pages"></a>

To begin listing a user's Confluence pages, at least a `data_source_id` of a connected
Confluence account must be specified. This base request returns a list of root pages for
every space the user has access to in a Confluence instance. To traverse further down
the user's page directory, additional requests to this endpoint can be made with the same
`data_source_id` and with `parent_id` set to the id of page from a previous request. For
convenience, the `has_children` property in each directory item in the response list will
flag which pages will return non-empty lists of pages when set as the `parent_id`.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
data_source_id = 1
parent_id = "string_example"

result = carbon.integrations.list_confluence_pages(
  data_source_id: data_source_id,
  parent_id: parent_id,
)
p result
```

---


### `carbon.integrations.list_data_source_items`<a id="carbonintegrationslist_data_source_items"></a>

List Data Source Items

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
data_source_id = 1
parent_id = "string_example"
pagination = {
        "limit" => 10,
        "offset" => 0,
    }

result = carbon.integrations.list_data_source_items(
  data_source_id: data_source_id,
  parent_id: parent_id,
  pagination: pagination,
)
p result
```

---


### `carbon.integrations.list_folders`<a id="carbonintegrationslist_folders"></a>

After connecting your Outlook account, you can use this endpoint to list all of your folders on outlook. This includes 
both system folders like "inbox" and user created folders.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby

result = carbon.integrations.list_folders
p result
```

---


### `carbon.integrations.list_labels`<a id="carbonintegrationslist_labels"></a>

After connecting your Gmail account, you can use this endpoint to list all of your labels. User created labels
will have the type "user" and Gmail's default labels will have the type "system"

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby

result = carbon.integrations.list_labels
p result
```

---


### `carbon.integrations.sync_confluence`<a id="carbonintegrationssync_confluence"></a>

After listing pages in a user's Confluence account, the set of selected page `ids` and the
connected account's `data_source_id` can be passed into this endpoint to sync them into
Carbon. Additional parameters listed below can be used to associate data to the selected
pages or alter the behavior of the sync.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
data_source_id = 1
ids = [
        "string_example"
    ]
tags = {}
chunk_size = 1500
chunk_overlap = 20
skip_embedding_generation = False
embedding_model = "OPENAI"
generate_sparse_vectors = False
prepend_filename_to_chunks = False
max_items_per_chunk = 1

result = carbon.integrations.sync_confluence(
  data_source_id: data_source_id,
  ids: ids,
  tags: tags,
  chunk_size: chunk_size,
  chunk_overlap: chunk_overlap,
  skip_embedding_generation: skip_embedding_generation,
  embedding_model: embedding_model,
  generate_sparse_vectors: generate_sparse_vectors,
  prepend_filename_to_chunks: prepend_filename_to_chunks,
  max_items_per_chunk: max_items_per_chunk,
)
p result
```

---


### `carbon.integrations.sync_data_source_items`<a id="carbonintegrationssync_data_source_items"></a>

Sync Data Source Items

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
data_source_id = 1

result = carbon.integrations.sync_data_source_items(
  data_source_id: data_source_id,
)
p result
```

---


### `carbon.integrations.sync_files`<a id="carbonintegrationssync_files"></a>

Sync Files

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
data_source_id = 1
ids = [
        "string_example"
    ]
tags = {}
chunk_size = 1500
chunk_overlap = 20
skip_embedding_generation = False
embedding_model = "OPENAI"
generate_sparse_vectors = False
prepend_filename_to_chunks = False
max_items_per_chunk = 1

result = carbon.integrations.sync_files(
  data_source_id: data_source_id,
  ids: ids,
  tags: tags,
  chunk_size: chunk_size,
  chunk_overlap: chunk_overlap,
  skip_embedding_generation: skip_embedding_generation,
  embedding_model: embedding_model,
  generate_sparse_vectors: generate_sparse_vectors,
  prepend_filename_to_chunks: prepend_filename_to_chunks,
  max_items_per_chunk: max_items_per_chunk,
)
p result
```

---


### `carbon.integrations.sync_gmail`<a id="carbonintegrationssync_gmail"></a>

Once you have successfully connected your gmail account, you can choose which emails to sync with us
using the filters parameter. Filters is a JSON object with key value pairs. It also supports AND and OR operations.
For now, we support a limited set of keys listed below.

<b>label</b>: Inbuilt Gmail labels, for example "Important" or a custom label you created.  
<b>after</b> or <b>before</b>: A date in YYYY/mm/dd format (example 2023/12/31). Gets emails after/before a certain date.
You can also use them in combination to get emails from a certain period.  
<b>is</b>: Can have the following values - starred, important, snoozed, and unread  

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
filters = {}
tags = {}
chunk_size = 1500
chunk_overlap = 20
skip_embedding_generation = False
embedding_model = "OPENAI"
generate_sparse_vectors = False
prepend_filename_to_chunks = False

result = carbon.integrations.sync_gmail(
  filters: filters,
  tags: tags,
  chunk_size: chunk_size,
  chunk_overlap: chunk_overlap,
  skip_embedding_generation: skip_embedding_generation,
  embedding_model: embedding_model,
  generate_sparse_vectors: generate_sparse_vectors,
  prepend_filename_to_chunks: prepend_filename_to_chunks,
)
p result
```

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
filters = {}
tags = {}
folder = "Inbox"
chunk_size = 1500
chunk_overlap = 20
skip_embedding_generation = False
embedding_model = "OPENAI"
generate_sparse_vectors = False
prepend_filename_to_chunks = False

result = carbon.integrations.sync_outlook(
  filters: filters,
  tags: tags,
  folder: folder,
  chunk_size: chunk_size,
  chunk_overlap: chunk_overlap,
  skip_embedding_generation: skip_embedding_generation,
  embedding_model: embedding_model,
  generate_sparse_vectors: generate_sparse_vectors,
  prepend_filename_to_chunks: prepend_filename_to_chunks,
)
p result
```

---


### `carbon.integrations.sync_rss_feed`<a id="carbonintegrationssync_rss_feed"></a>

Rss Feed

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
url = "string_example"
tags = {}
chunk_size = 1500
chunk_overlap = 20
skip_embedding_generation = False
embedding_model = "OPENAI"
generate_sparse_vectors = False
prepend_filename_to_chunks = False

result = carbon.integrations.sync_rss_feed(
  url: url,
  tags: tags,
  chunk_size: chunk_size,
  chunk_overlap: chunk_overlap,
  skip_embedding_generation: skip_embedding_generation,
  embedding_model: embedding_model,
  generate_sparse_vectors: generate_sparse_vectors,
  prepend_filename_to_chunks: prepend_filename_to_chunks,
)
p result
```

---


### `carbon.integrations.sync_s3_files`<a id="carbonintegrationssync_s3_files"></a>

After optionally loading the items via /integrations/items/sync and integrations/items/list, use the bucket name 
and object key as the ID in this endpoint to sync them into Carbon. Additional parameters below can associate 
data with the selected items or modify the sync behavior

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
ids = [
        {
        }
    ]
tags = {}
chunk_size = 1500
chunk_overlap = 20
skip_embedding_generation = False
embedding_model = "OPENAI"
generate_sparse_vectors = False
prepend_filename_to_chunks = False
max_items_per_chunk = 1

result = carbon.integrations.sync_s3_files(
  ids: ids,
  tags: tags,
  chunk_size: chunk_size,
  chunk_overlap: chunk_overlap,
  skip_embedding_generation: skip_embedding_generation,
  embedding_model: embedding_model,
  generate_sparse_vectors: generate_sparse_vectors,
  prepend_filename_to_chunks: prepend_filename_to_chunks,
  max_items_per_chunk: max_items_per_chunk,
)
p result
```

---


### `carbon.organizations.get`<a id="carbonorganizationsget"></a>

Get Organization

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby

result = carbon.organizations.get
p result
```

---


### `carbon.users.get`<a id="carbonusersget"></a>

User Endpoint

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
customer_id = "string_example"

result = carbon.users.get(
  customer_id: customer_id,
)
p result
```

---


### `carbon.users.toggle_user_features`<a id="carbonuserstoggle_user_features"></a>

Toggle User Features

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
configuration_key_name = "string_example"
value = {}

result = carbon.users.toggle_user_features(
  configuration_key_name: configuration_key_name,
  value: value,
)
p result
```

---


### `carbon.utilities.fetch_urls`<a id="carbonutilitiesfetch_urls"></a>

Extracts all URLs from a webpage. 

Args:
    url (str): URL of the webpage

Returns:
    FetchURLsResponse: A response object with a list of URLs extracted from the webpage and the webpage content.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
url = "url_example"

result = carbon.utilities.fetch_urls(
  url: url,
)
p result
```

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
id = "id_example"
raw = False

result = carbon.utilities.fetch_youtube_transcripts(
  id: id,
  raw: raw,
)
p result
```

---


### `carbon.utilities.process_sitemap`<a id="carbonutilitiesprocess_sitemap"></a>

Retrieves all URLs from a sitemap, which can subsequently be utilized with our `web_scrape` endpoint.

<!--Args:
    url (str): URL of the sitemap

Returns:
    dict: A dictionary with a list of URLs extracted from the sitemap.-->

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
url = "url_example"

result = carbon.utilities.process_sitemap(
  url: url,
)
p result
```

---


### `carbon.utilities.scrape_sitemap`<a id="carbonutilitiesscrape_sitemap"></a>

Extracts all URLs from a sitemap and performs a web scrape on each of them.

Args:
    sitemap_url (str): URL of the sitemap

Returns:
    dict: A response object with the status of the scraping job message.-->

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
url = "string_example"
tags = {
        "key": "string_example",
    }
max_pages_to_scrape = 100
chunk_size = 1500
chunk_overlap = 20
skip_embedding_generation = False
enable_auto_sync = False
generate_sparse_vectors = False
prepend_filename_to_chunks = False
html_tags_to_skip = []
css_classes_to_skip = []
css_selectors_to_skip = []

result = carbon.utilities.scrape_sitemap(
  url: url,
  tags: tags,
  max_pages_to_scrape: max_pages_to_scrape,
  chunk_size: chunk_size,
  chunk_overlap: chunk_overlap,
  skip_embedding_generation: skip_embedding_generation,
  enable_auto_sync: enable_auto_sync,
  generate_sparse_vectors: generate_sparse_vectors,
  prepend_filename_to_chunks: prepend_filename_to_chunks,
  html_tags_to_skip: html_tags_to_skip,
  css_classes_to_skip: css_classes_to_skip,
  css_selectors_to_skip: css_selectors_to_skip,
)
p result
```

---


### `carbon.utilities.scrape_web`<a id="carbonutilitiesscrape_web"></a>

Conduct a web scrape on a given webpage URL. Our web scraper is fully compatible with JavaScript and supports recursion depth, enabling you to efficiently extract all content from the target website.

<!--Args:
    scraping_requests (List[WebscrapeRequest]): A list of WebscrapeRequest objects.

    
Returns:
    dict: A response object with the status of the scraping job message.-->

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
body = [
        {
            "url" => "url_example",
            "recursion_depth" => 3,
            "max_pages_to_scrape" => 100,
            "chunk_size" => 1500,
            "chunk_overlap" => 20,
            "skip_embedding_generation" => False,
            "enable_auto_sync" => False,
            "generate_sparse_vectors" => False,
            "prepend_filename_to_chunks" => False,
            "html_tags_to_skip" => [],
            "css_classes_to_skip" => [],
            "css_selectors_to_skip" => [],
        }
    ]

result = carbon.utilities.scrape_web(
  body: body,
)
p result
```

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
query = "query_example"

result = carbon.utilities.search_urls(
  query: query,
)
p result
```

---


### `carbon.webhooks.add_url`<a id="carbonwebhooksadd_url"></a>

Add Webhook Url

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
url = "string_example"

result = carbon.webhooks.add_url(
  url: url,
)
p result
```

---


### `carbon.webhooks.delete_url`<a id="carbonwebhooksdelete_url"></a>

Delete Webhook Url

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
webhook_id = 1

result = carbon.webhooks.delete_url(
  webhook_id: webhook_id,
)
p result
```

---


### `carbon.webhooks.urls`<a id="carbonwebhooksurls"></a>

Webhook Urls

#### 🛠️ Usage<a id="🛠️-usage"></a>

```ruby
pagination = {
        "limit" => 10,
        "offset" => 0,
    }
order_by = "created_at"
order_dir = "desc"
filters = {
        "ids" => [],
    }

result = carbon.webhooks.urls(
  pagination: pagination,
  order_by: order_by,
  order_dir: order_dir,
  filters: filters,
)
p result
```

---


## Author<a id="author"></a>
This TypeScript package is automatically generated by [Konfig](https://konfigthis.com)

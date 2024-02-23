# Carbon::IntegrationsApi

All URIs are relative to *https://api.carbon.ai*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**connect_freshdesk**](IntegrationsApi.md#connect_freshdesk) | **POST** /integrations/freshdesk | Freshdesk Connect |
| [**create_aws_iam_user**](IntegrationsApi.md#create_aws_iam_user) | **POST** /integrations/s3 | S3 Auth |
| [**get_oauth_url**](IntegrationsApi.md#get_oauth_url) | **POST** /integrations/oauth_url | Get Oauth Url |
| [**list_confluence_pages**](IntegrationsApi.md#list_confluence_pages) | **POST** /integrations/confluence/list | Confluence List |
| [**list_data_source_items**](IntegrationsApi.md#list_data_source_items) | **POST** /integrations/items/list | List Data Source Items |
| [**list_folders**](IntegrationsApi.md#list_folders) | **GET** /integrations/outlook/user_folders | Outlook Folders |
| [**list_labels**](IntegrationsApi.md#list_labels) | **GET** /integrations/gmail/user_labels | Gmail Labels |
| [**sync_confluence**](IntegrationsApi.md#sync_confluence) | **POST** /integrations/confluence/sync | Confluence Sync |
| [**sync_data_source_items**](IntegrationsApi.md#sync_data_source_items) | **POST** /integrations/items/sync | Sync Data Source Items |
| [**sync_files**](IntegrationsApi.md#sync_files) | **POST** /integrations/files/sync | Sync Files |
| [**sync_gmail**](IntegrationsApi.md#sync_gmail) | **POST** /integrations/gmail/sync | Gmail Sync |
| [**sync_outlook**](IntegrationsApi.md#sync_outlook) | **POST** /integrations/outlook/sync | Outlook Sync |
| [**sync_rss_feed**](IntegrationsApi.md#sync_rss_feed) | **POST** /integrations/rss_feed | Rss Feed |
| [**sync_s3_files**](IntegrationsApi.md#sync_s3_files) | **POST** /integrations/s3/files | S3 Files |

## connect_freshdesk

Freshdesk Connect

Refer this article to obtain an API key https://support.freshdesk.com/en/support/solutions/articles/215517. Make sure that your API key has the permission to read solutions from your account and you are on a <b>paid</b> plan. Once you have an API key, you can make a request to this endpoint along with your freshdesk domain. This will  trigger an automatic sync of the articles in your \"solutions\" tab. Additional parameters below can be used to associate  data with the synced articles or modify the sync behavior.

### Example

```ruby
require 'carbon'
configuration = Carbon::Configuration.new
configuration.access_token = 'YOUR API KEY'
configuration.api_key = 'YOUR API KEY'
configuration.customer_id = 'YOUR API KEY'
carbon = Carbon::Client.new(configuration)

domain = "string_example"
api_key = "string_example"
tags = {}
chunk_size = 1500
chunk_overlap = 20
skip_embedding_generation = False
embedding_model = "OPENAI"
generate_sparse_vectors = False
prepend_filename_to_chunks = False

begin
  # Freshdesk Connect
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
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.connect_freshdesk: #{e}"
end
```

#### Using the connect_freshdesk_with_http_info variant

This returns an Array which contains the response data, status code and headers.

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

begin
  # Freshdesk Connect
  data, status_code, headers, response = Carbon::Integrations.connect_freshdesk_with_http_info(
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
  p status_code # => 2xx
  p headers # => { ... }
  p data # => GenericSuccessResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.connect_freshdesk: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **fresh_desk_connect_request** | [**FreshDeskConnectRequest**](FreshDeskConnectRequest.md) |  |  |

### Return type

[**GenericSuccessResponse**](GenericSuccessResponse.md)

## create_aws_iam_user

S3 Auth

Create a new IAM user with permissions to: <ol> <li>List all buckets.</li> <li>Read from the specific buckets and objects to sync with Carbon. Ensure any future buckets or objects carry  the same permissions.</li> </ol> Once created, generate an access key for this user and share the credentials with us. We recommend testing this key beforehand.

### Example

```ruby
require 'carbon'
configuration = Carbon::Configuration.new
configuration.access_token = 'YOUR API KEY'
configuration.api_key = 'YOUR API KEY'
configuration.customer_id = 'YOUR API KEY'
carbon = Carbon::Client.new(configuration)

access_key = "string_example"
access_key_secret = "string_example"

begin
  # S3 Auth
  result = carbon.integrations.create_aws_iam_user(
                                                      access_key: access_key,
                                                      access_key_secret: access_key_secret,
                                                    )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.create_aws_iam_user: #{e}"
end
```

#### Using the create_aws_iam_user_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
access_key = "string_example"
access_key_secret = "string_example"

begin
  # S3 Auth
  data, status_code, headers, response = Carbon::Integrations.create_aws_iam_user_with_http_info(
                                                                                                    access_key: access_key,
                                                                                                    access_key_secret: access_key_secret,
                                                                                                  )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => OrganizationUserDataSourceAPI
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.create_aws_iam_user: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **s3_auth_request** | [**S3AuthRequest**](S3AuthRequest.md) |  |  |

### Return type

[**OrganizationUserDataSourceAPI**](OrganizationUserDataSourceAPI.md)

## get_oauth_url

Get Oauth Url

### Example

```ruby
require 'carbon'
configuration = Carbon::Configuration.new
configuration.access_token = 'YOUR API KEY'
configuration.api_key = 'YOUR API KEY'
configuration.customer_id = 'YOUR API KEY'
carbon = Carbon::Client.new(configuration)

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

begin
  # Get Oauth Url
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
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.get_oauth_url: #{e}"
end
```

#### Using the get_oauth_url_with_http_info variant

This returns an Array which contains the response data, status code and headers.

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

begin
  # Get Oauth Url
  data, status_code, headers, response = Carbon::Integrations.get_oauth_url_with_http_info(
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
  p status_code # => 2xx
  p headers # => { ... }
  p data # => Object
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.get_oauth_url: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **o_auth_url_request** | [**OAuthURLRequest**](OAuthURLRequest.md) |  |  |

### Return type

**Object**

## list_confluence_pages

Confluence List

To begin listing a user's Confluence pages, at least a `data_source_id` of a connected Confluence account must be specified. This base request returns a list of root pages for every space the user has access to in a Confluence instance. To traverse further down the user's page directory, additional requests to this endpoint can be made with the same `data_source_id` and with `parent_id` set to the id of page from a previous request. For convenience, the `has_children` property in each directory item in the response list will flag which pages will return non-empty lists of pages when set as the `parent_id`.

### Example

```ruby
require 'carbon'
configuration = Carbon::Configuration.new
configuration.access_token = 'YOUR API KEY'
configuration.api_key = 'YOUR API KEY'
configuration.customer_id = 'YOUR API KEY'
carbon = Carbon::Client.new(configuration)

data_source_id = 1
parent_id = "string_example"

begin
  # Confluence List
  result = carbon.integrations.list_confluence_pages(
                                                        data_source_id: data_source_id,
                                                        parent_id: parent_id,
                                                      )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.list_confluence_pages: #{e}"
end
```

#### Using the list_confluence_pages_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
data_source_id = 1
parent_id = "string_example"

begin
  # Confluence List
  data, status_code, headers, response = Carbon::Integrations.list_confluence_pages_with_http_info(
                                                                                                      data_source_id: data_source_id,
                                                                                                      parent_id: parent_id,
                                                                                                    )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => ListResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.list_confluence_pages: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **list_request** | [**ListRequest**](ListRequest.md) |  |  |

### Return type

[**ListResponse**](ListResponse.md)

## list_data_source_items

List Data Source Items

### Example

```ruby
require 'carbon'
configuration = Carbon::Configuration.new
configuration.access_token = 'YOUR API KEY'
configuration.api_key = 'YOUR API KEY'
configuration.customer_id = 'YOUR API KEY'
carbon = Carbon::Client.new(configuration)

data_source_id = 1
parent_id = "string_example"
pagination = {
        "limit" => 10,
        "offset" => 0,
    }

begin
  # List Data Source Items
  result = carbon.integrations.list_data_source_items(
                                                         data_source_id: data_source_id,
                                                         parent_id: parent_id,
                                                         pagination: pagination,
                                                       )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.list_data_source_items: #{e}"
end
```

#### Using the list_data_source_items_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
data_source_id = 1
parent_id = "string_example"
pagination = {
        "limit" => 10,
        "offset" => 0,
    }

begin
  # List Data Source Items
  data, status_code, headers, response = Carbon::Integrations.list_data_source_items_with_http_info(
                                                                                                       data_source_id: data_source_id,
                                                                                                       parent_id: parent_id,
                                                                                                       pagination: pagination,
                                                                                                     )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => ListDataSourceItemsResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.list_data_source_items: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **list_data_source_items_request** | [**ListDataSourceItemsRequest**](ListDataSourceItemsRequest.md) |  |  |

### Return type

[**ListDataSourceItemsResponse**](ListDataSourceItemsResponse.md)

## list_folders

Outlook Folders

After connecting your Outlook account, you can use this endpoint to list all of your folders on outlook. This includes  both system folders like \"inbox\" and user created folders.

### Example

```ruby
require 'carbon'
configuration = Carbon::Configuration.new
configuration.access_token = 'YOUR API KEY'
configuration.api_key = 'YOUR API KEY'
configuration.customer_id = 'YOUR API KEY'
carbon = Carbon::Client.new(configuration)


begin
  # Outlook Folders
  result = carbon.integrations.list_folders
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.list_folders: #{e}"
end
```

#### Using the list_folders_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby

begin
  # Outlook Folders
  data, status_code, headers, response = Carbon::Integrations.list_folders_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => Object
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.list_folders: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

**Object**

## list_labels

Gmail Labels

After connecting your Gmail account, you can use this endpoint to list all of your labels. User created labels will have the type \"user\" and Gmail's default labels will have the type \"system\"

### Example

```ruby
require 'carbon'
configuration = Carbon::Configuration.new
configuration.access_token = 'YOUR API KEY'
configuration.api_key = 'YOUR API KEY'
configuration.customer_id = 'YOUR API KEY'
carbon = Carbon::Client.new(configuration)


begin
  # Gmail Labels
  result = carbon.integrations.list_labels
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.list_labels: #{e}"
end
```

#### Using the list_labels_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby

begin
  # Gmail Labels
  data, status_code, headers, response = Carbon::Integrations.list_labels_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => Object
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.list_labels: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

**Object**

## sync_confluence

Confluence Sync

After listing pages in a user's Confluence account, the set of selected page `ids` and the connected account's `data_source_id` can be passed into this endpoint to sync them into Carbon. Additional parameters listed below can be used to associate data to the selected pages or alter the behavior of the sync.

### Example

```ruby
require 'carbon'
configuration = Carbon::Configuration.new
configuration.access_token = 'YOUR API KEY'
configuration.api_key = 'YOUR API KEY'
configuration.customer_id = 'YOUR API KEY'
carbon = Carbon::Client.new(configuration)

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

begin
  # Confluence Sync
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
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.sync_confluence: #{e}"
end
```

#### Using the sync_confluence_with_http_info variant

This returns an Array which contains the response data, status code and headers.

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

begin
  # Confluence Sync
  data, status_code, headers, response = Carbon::Integrations.sync_confluence_with_http_info(
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
  p status_code # => 2xx
  p headers # => { ... }
  p data # => GenericSuccessResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.sync_confluence: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **sync_files_request** | [**SyncFilesRequest**](SyncFilesRequest.md) |  |  |

### Return type

[**GenericSuccessResponse**](GenericSuccessResponse.md)

## sync_data_source_items

Sync Data Source Items

### Example

```ruby
require 'carbon'
configuration = Carbon::Configuration.new
configuration.access_token = 'YOUR API KEY'
configuration.api_key = 'YOUR API KEY'
configuration.customer_id = 'YOUR API KEY'
carbon = Carbon::Client.new(configuration)

data_source_id = 1

begin
  # Sync Data Source Items
  result = carbon.integrations.sync_data_source_items(
                                                         data_source_id: data_source_id,
                                                       )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.sync_data_source_items: #{e}"
end
```

#### Using the sync_data_source_items_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
data_source_id = 1

begin
  # Sync Data Source Items
  data, status_code, headers, response = Carbon::Integrations.sync_data_source_items_with_http_info(
                                                                                                       data_source_id: data_source_id,
                                                                                                     )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => OrganizationUserDataSourceAPI
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.sync_data_source_items: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **sync_directory_request** | [**SyncDirectoryRequest**](SyncDirectoryRequest.md) |  |  |

### Return type

[**OrganizationUserDataSourceAPI**](OrganizationUserDataSourceAPI.md)

## sync_files

Sync Files

### Example

```ruby
require 'carbon'
configuration = Carbon::Configuration.new
configuration.access_token = 'YOUR API KEY'
configuration.api_key = 'YOUR API KEY'
configuration.customer_id = 'YOUR API KEY'
carbon = Carbon::Client.new(configuration)

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

begin
  # Sync Files
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
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.sync_files: #{e}"
end
```

#### Using the sync_files_with_http_info variant

This returns an Array which contains the response data, status code and headers.

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

begin
  # Sync Files
  data, status_code, headers, response = Carbon::Integrations.sync_files_with_http_info(
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
  p status_code # => 2xx
  p headers # => { ... }
  p data # => GenericSuccessResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.sync_files: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **sync_files_request** | [**SyncFilesRequest**](SyncFilesRequest.md) |  |  |

### Return type

[**GenericSuccessResponse**](GenericSuccessResponse.md)

## sync_gmail

Gmail Sync

Once you have successfully connected your gmail account, you can choose which emails to sync with us using the filters parameter. Filters is a JSON object with key value pairs. It also supports AND and OR operations. For now, we support a limited set of keys listed below.  <b>label</b>: Inbuilt Gmail labels, for example \"Important\" or a custom label you created.   <b>after</b> or <b>before</b>: A date in YYYY/mm/dd format (example 2023/12/31). Gets emails after/before a certain date. You can also use them in combination to get emails from a certain period.   <b>is</b>: Can have the following values - starred, important, snoozed, and unread    Using keys or values outside of the specified values can lead to unexpected behaviour.  An example of a basic query with filters can be ```json {     \"filters\": {             \"key\": \"label\",             \"value\": \"Test\"         } } ``` Which will list all emails that have the label \"Test\".  You can use AND and OR operation in the following way: ```json {     \"filters\": {         \"AND\": [             {                 \"key\": \"after\",                 \"value\": \"2024/01/07\"             },             {                 \"OR\": [                     {                         \"key\": \"label\",                         \"value\": \"Personal\"                     },                     {                         \"key\": \"is\",                         \"value\": \"starred\"                     }                 ]             }         ]     } } ``` This will return emails after 7th of Jan that are either starred or have the label \"Personal\".  Note that this is the highest level of nesting we support, i.e. you can't add more AND/OR filters within the OR filter in the above example.

### Example

```ruby
require 'carbon'
configuration = Carbon::Configuration.new
configuration.access_token = 'YOUR API KEY'
configuration.api_key = 'YOUR API KEY'
configuration.customer_id = 'YOUR API KEY'
carbon = Carbon::Client.new(configuration)

filters = {}
tags = {}
chunk_size = 1500
chunk_overlap = 20
skip_embedding_generation = False
embedding_model = "OPENAI"
generate_sparse_vectors = False
prepend_filename_to_chunks = False

begin
  # Gmail Sync
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
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.sync_gmail: #{e}"
end
```

#### Using the sync_gmail_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
filters = {}
tags = {}
chunk_size = 1500
chunk_overlap = 20
skip_embedding_generation = False
embedding_model = "OPENAI"
generate_sparse_vectors = False
prepend_filename_to_chunks = False

begin
  # Gmail Sync
  data, status_code, headers, response = Carbon::Integrations.sync_gmail_with_http_info(
                                                                                           filters: filters,
                                                                                           tags: tags,
                                                                                           chunk_size: chunk_size,
                                                                                           chunk_overlap: chunk_overlap,
                                                                                           skip_embedding_generation: skip_embedding_generation,
                                                                                           embedding_model: embedding_model,
                                                                                           generate_sparse_vectors: generate_sparse_vectors,
                                                                                           prepend_filename_to_chunks: prepend_filename_to_chunks,
                                                                                         )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => GenericSuccessResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.sync_gmail: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **gmail_sync_input** | [**GmailSyncInput**](GmailSyncInput.md) |  |  |

### Return type

[**GenericSuccessResponse**](GenericSuccessResponse.md)

## sync_outlook

Outlook Sync

Once you have successfully connected your Outlook account, you can choose which emails to sync with us using the filters and folder parameter. \"folder\" should be the folder you want to sync from Outlook. By default we get messages from your inbox folder.   Filters is a JSON object with key value pairs. It also supports AND and OR operations. For now, we support a limited set of keys listed below.  <b>category</b>: Custom categories that you created in Outlook.   <b>after</b> or <b>before</b>: A date in YYYY/mm/dd format (example 2023/12/31). Gets emails after/before a certain date. You can also use them in combination to get emails from a certain period.     <b>is</b>: Can have the following values: flagged     An example of a basic query with filters can be ```json {     \"filters\": {             \"key\": \"category\",             \"value\": \"Test\"         } } ``` Which will list all emails that have the category \"Test\".    Specifying a custom folder in the same query ```json {     \"folder\": \"Folder Name\",     \"filters\": {             \"key\": \"category\",             \"value\": \"Test\"         } } ```  You can use AND and OR operation in the following way: ```json {     \"filters\": {         \"AND\": [             {                 \"key\": \"after\",                 \"value\": \"2024/01/07\"             },             {                 \"OR\": [                     {                         \"key\": \"category\",                         \"value\": \"Personal\"                     },                     {                         \"key\": \"category\",                         \"value\": \"Test\"                     },                 ]             }         ]     } } ``` This will return emails after 7th of Jan that have either Personal or Test as category.  Note that this is the highest level of nesting we support, i.e. you can't add more AND/OR filters within the OR filter in the above example.

### Example

```ruby
require 'carbon'
configuration = Carbon::Configuration.new
configuration.access_token = 'YOUR API KEY'
configuration.api_key = 'YOUR API KEY'
configuration.customer_id = 'YOUR API KEY'
carbon = Carbon::Client.new(configuration)

filters = {}
tags = {}
folder = "Inbox"
chunk_size = 1500
chunk_overlap = 20
skip_embedding_generation = False
embedding_model = "OPENAI"
generate_sparse_vectors = False
prepend_filename_to_chunks = False

begin
  # Outlook Sync
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
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.sync_outlook: #{e}"
end
```

#### Using the sync_outlook_with_http_info variant

This returns an Array which contains the response data, status code and headers.

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

begin
  # Outlook Sync
  data, status_code, headers, response = Carbon::Integrations.sync_outlook_with_http_info(
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
  p status_code # => 2xx
  p headers # => { ... }
  p data # => GenericSuccessResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.sync_outlook: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **outlook_sync_input** | [**OutlookSyncInput**](OutlookSyncInput.md) |  |  |

### Return type

[**GenericSuccessResponse**](GenericSuccessResponse.md)

## sync_rss_feed

Rss Feed

### Example

```ruby
require 'carbon'
configuration = Carbon::Configuration.new
configuration.access_token = 'YOUR API KEY'
configuration.api_key = 'YOUR API KEY'
configuration.customer_id = 'YOUR API KEY'
carbon = Carbon::Client.new(configuration)

url = "string_example"
tags = {}
chunk_size = 1500
chunk_overlap = 20
skip_embedding_generation = False
embedding_model = "OPENAI"
generate_sparse_vectors = False
prepend_filename_to_chunks = False

begin
  # Rss Feed
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
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.sync_rss_feed: #{e}"
end
```

#### Using the sync_rss_feed_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
url = "string_example"
tags = {}
chunk_size = 1500
chunk_overlap = 20
skip_embedding_generation = False
embedding_model = "OPENAI"
generate_sparse_vectors = False
prepend_filename_to_chunks = False

begin
  # Rss Feed
  data, status_code, headers, response = Carbon::Integrations.sync_rss_feed_with_http_info(
                                                                                              url: url,
                                                                                              tags: tags,
                                                                                              chunk_size: chunk_size,
                                                                                              chunk_overlap: chunk_overlap,
                                                                                              skip_embedding_generation: skip_embedding_generation,
                                                                                              embedding_model: embedding_model,
                                                                                              generate_sparse_vectors: generate_sparse_vectors,
                                                                                              prepend_filename_to_chunks: prepend_filename_to_chunks,
                                                                                            )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => GenericSuccessResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.sync_rss_feed: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **rss_feed_input** | [**RSSFeedInput**](RSSFeedInput.md) |  |  |

### Return type

[**GenericSuccessResponse**](GenericSuccessResponse.md)

## sync_s3_files

S3 Files

After optionally loading the items via /integrations/items/sync and integrations/items/list, use the bucket name  and object key as the ID in this endpoint to sync them into Carbon. Additional parameters below can associate  data with the selected items or modify the sync behavior

### Example

```ruby
require 'carbon'
configuration = Carbon::Configuration.new
configuration.access_token = 'YOUR API KEY'
configuration.api_key = 'YOUR API KEY'
configuration.customer_id = 'YOUR API KEY'
carbon = Carbon::Client.new(configuration)

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

begin
  # S3 Files
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
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.sync_s3_files: #{e}"
end
```

#### Using the sync_s3_files_with_http_info variant

This returns an Array which contains the response data, status code and headers.

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

begin
  # S3 Files
  data, status_code, headers, response = Carbon::Integrations.sync_s3_files_with_http_info(
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
  p status_code # => 2xx
  p headers # => { ... }
  p data # => GenericSuccessResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Integrations.sync_s3_files: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **s3_file_sync_input** | [**S3FileSyncInput**](S3FileSyncInput.md) |  |  |

### Return type

[**GenericSuccessResponse**](GenericSuccessResponse.md)


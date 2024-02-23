# carbon

Connect external data to LLMs, no matter the source.


## Installation

Add to Gemfile:

```ruby
gem 'carbon', '~> 1.0.0'
```

## Getting Started

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'
Carbon.api_key = 'YOUR API KEY'
Carbon.customer_id = 'YOUR API KEY'


begin
  # Get Access Token
  result = Carbon::Auth.get_access_token
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Auth.get_access_token: #{e}"
end
```

## Documentation for API Endpoints

All URIs are relative to *https://api.carbon.ai*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*Carbon::AuthApi* | [**get_access_token**](docs/AuthApi.md#get_access_token) | **GET** /auth/v1/access_token | Get Access Token
*Carbon::AuthApi* | [**get_white_labeling**](docs/AuthApi.md#get_white_labeling) | **GET** /auth/v1/white_labeling | Get White Labeling
*Carbon::DataSourcesApi* | [**query_user_data_sources**](docs/DataSourcesApi.md#query_user_data_sources) | **POST** /user_data_sources | User Data Sources
*Carbon::DataSourcesApi* | [**revoke_access_token**](docs/DataSourcesApi.md#revoke_access_token) | **POST** /revoke_access_token | Revoke Access Token
*Carbon::EmbeddingsApi* | [**get_documents**](docs/EmbeddingsApi.md#get_documents) | **POST** /embeddings | Embeddings
*Carbon::EmbeddingsApi* | [**get_embeddings_and_chunks**](docs/EmbeddingsApi.md#get_embeddings_and_chunks) | **POST** /text_chunks | Retrieve Embeddings And Content
*Carbon::EmbeddingsApi* | [**upload_chunks_and_embeddings**](docs/EmbeddingsApi.md#upload_chunks_and_embeddings) | **POST** /upload_chunks_and_embeddings | Upload Chunks And Embeddings
*Carbon::FilesApi* | [**create_user_file_tags**](docs/FilesApi.md#create_user_file_tags) | **POST** /create_user_file_tags | Create File Tags
*Carbon::FilesApi* | [**delete**](docs/FilesApi.md#delete) | **DELETE** /deletefile/{file_id} | Delete File Endpoint
*Carbon::FilesApi* | [**delete_file_tags**](docs/FilesApi.md#delete_file_tags) | **POST** /delete_user_file_tags | Delete File Tags
*Carbon::FilesApi* | [**delete_many**](docs/FilesApi.md#delete_many) | **POST** /delete_files | Delete Files Endpoint
*Carbon::FilesApi* | [**get_parsed_file**](docs/FilesApi.md#get_parsed_file) | **GET** /parsed_file/{file_id} | Parsed File
*Carbon::FilesApi* | [**get_raw_file**](docs/FilesApi.md#get_raw_file) | **GET** /raw_file/{file_id} | Raw File
*Carbon::FilesApi* | [**query_user_files**](docs/FilesApi.md#query_user_files) | **POST** /user_files_v2 | User Files V2
*Carbon::FilesApi* | [**query_user_files_deprecated**](docs/FilesApi.md#query_user_files_deprecated) | **POST** /user_files | User Files
*Carbon::FilesApi* | [**resync**](docs/FilesApi.md#resync) | **POST** /resync_file | Resync File
*Carbon::FilesApi* | [**upload**](docs/FilesApi.md#upload) | **POST** /uploadfile | Create Upload File
*Carbon::FilesApi* | [**upload_from_url**](docs/FilesApi.md#upload_from_url) | **POST** /upload_file_from_url | Create Upload File From Url
*Carbon::FilesApi* | [**upload_text**](docs/FilesApi.md#upload_text) | **POST** /upload_text | Create Raw Text
*Carbon::HealthApi* | [**check**](docs/HealthApi.md#check) | **GET** /health | Health
*Carbon::IntegrationsApi* | [**connect_freshdesk**](docs/IntegrationsApi.md#connect_freshdesk) | **POST** /integrations/freshdesk | Freshdesk Connect
*Carbon::IntegrationsApi* | [**create_aws_iam_user**](docs/IntegrationsApi.md#create_aws_iam_user) | **POST** /integrations/s3 | S3 Auth
*Carbon::IntegrationsApi* | [**get_oauth_url**](docs/IntegrationsApi.md#get_oauth_url) | **POST** /integrations/oauth_url | Get Oauth Url
*Carbon::IntegrationsApi* | [**list_confluence_pages**](docs/IntegrationsApi.md#list_confluence_pages) | **POST** /integrations/confluence/list | Confluence List
*Carbon::IntegrationsApi* | [**list_data_source_items**](docs/IntegrationsApi.md#list_data_source_items) | **POST** /integrations/items/list | List Data Source Items
*Carbon::IntegrationsApi* | [**list_folders**](docs/IntegrationsApi.md#list_folders) | **GET** /integrations/outlook/user_folders | Outlook Folders
*Carbon::IntegrationsApi* | [**list_labels**](docs/IntegrationsApi.md#list_labels) | **GET** /integrations/gmail/user_labels | Gmail Labels
*Carbon::IntegrationsApi* | [**sync_confluence**](docs/IntegrationsApi.md#sync_confluence) | **POST** /integrations/confluence/sync | Confluence Sync
*Carbon::IntegrationsApi* | [**sync_data_source_items**](docs/IntegrationsApi.md#sync_data_source_items) | **POST** /integrations/items/sync | Sync Data Source Items
*Carbon::IntegrationsApi* | [**sync_files**](docs/IntegrationsApi.md#sync_files) | **POST** /integrations/files/sync | Sync Files
*Carbon::IntegrationsApi* | [**sync_gmail**](docs/IntegrationsApi.md#sync_gmail) | **POST** /integrations/gmail/sync | Gmail Sync
*Carbon::IntegrationsApi* | [**sync_outlook**](docs/IntegrationsApi.md#sync_outlook) | **POST** /integrations/outlook/sync | Outlook Sync
*Carbon::IntegrationsApi* | [**sync_rss_feed**](docs/IntegrationsApi.md#sync_rss_feed) | **POST** /integrations/rss_feed | Rss Feed
*Carbon::IntegrationsApi* | [**sync_s3_files**](docs/IntegrationsApi.md#sync_s3_files) | **POST** /integrations/s3/files | S3 Files
*Carbon::OrganizationsApi* | [**get**](docs/OrganizationsApi.md#get) | **GET** /organization | Get Organization
*Carbon::UsersApi* | [**get**](docs/UsersApi.md#get) | **POST** /user | User Endpoint
*Carbon::UsersApi* | [**toggle_user_features**](docs/UsersApi.md#toggle_user_features) | **POST** /modify_user_configuration | Toggle User Features
*Carbon::UtilitiesApi* | [**fetch_urls**](docs/UtilitiesApi.md#fetch_urls) | **GET** /fetch_urls | Fetch Urls
*Carbon::UtilitiesApi* | [**fetch_youtube_transcripts**](docs/UtilitiesApi.md#fetch_youtube_transcripts) | **GET** /fetch_youtube_transcript | Fetch Youtube Transcripts
*Carbon::UtilitiesApi* | [**process_sitemap**](docs/UtilitiesApi.md#process_sitemap) | **GET** /process_sitemap | Sitemap
*Carbon::UtilitiesApi* | [**scrape_sitemap**](docs/UtilitiesApi.md#scrape_sitemap) | **POST** /scrape_sitemap | Scrape Sitemap
*Carbon::UtilitiesApi* | [**scrape_web**](docs/UtilitiesApi.md#scrape_web) | **POST** /web_scrape | Web Scrape
*Carbon::UtilitiesApi* | [**search_urls**](docs/UtilitiesApi.md#search_urls) | **GET** /search_urls | Search Urls
*Carbon::WebhooksApi* | [**add_url**](docs/WebhooksApi.md#add_url) | **POST** /add_webhook | Add Webhook Url
*Carbon::WebhooksApi* | [**delete_url**](docs/WebhooksApi.md#delete_url) | **DELETE** /delete_webhook/{webhook_id} | Delete Webhook Url
*Carbon::WebhooksApi* | [**urls**](docs/WebhooksApi.md#urls) | **POST** /webhooks | Webhook Urls


## Documentation for Models

 - [Carbon::AddWebhookProps](docs/AddWebhookProps.md)
 - [Carbon::BodyCreateUploadFileUploadfilePost](docs/BodyCreateUploadFileUploadfilePost.md)
 - [Carbon::ChunkProperties](docs/ChunkProperties.md)
 - [Carbon::ChunkPropertiesNullable](docs/ChunkPropertiesNullable.md)
 - [Carbon::ChunksAndEmbeddings](docs/ChunksAndEmbeddings.md)
 - [Carbon::ChunksAndEmbeddingsUploadInput](docs/ChunksAndEmbeddingsUploadInput.md)
 - [Carbon::DataSourceLastSyncActions](docs/DataSourceLastSyncActions.md)
 - [Carbon::DataSourceSyncStatuses](docs/DataSourceSyncStatuses.md)
 - [Carbon::DataSourceType](docs/DataSourceType.md)
 - [Carbon::DataSourceTypeNullable](docs/DataSourceTypeNullable.md)
 - [Carbon::DeleteFilesQueryInput](docs/DeleteFilesQueryInput.md)
 - [Carbon::DirectoryItem](docs/DirectoryItem.md)
 - [Carbon::DocumentResponse](docs/DocumentResponse.md)
 - [Carbon::DocumentResponseList](docs/DocumentResponseList.md)
 - [Carbon::EmbeddingAndChunk](docs/EmbeddingAndChunk.md)
 - [Carbon::EmbeddingGenerators](docs/EmbeddingGenerators.md)
 - [Carbon::EmbeddingGeneratorsNullable](docs/EmbeddingGeneratorsNullable.md)
 - [Carbon::EmbeddingProperties](docs/EmbeddingProperties.md)
 - [Carbon::EmbeddingsAndChunksFilters](docs/EmbeddingsAndChunksFilters.md)
 - [Carbon::EmbeddingsAndChunksOrderByColumns](docs/EmbeddingsAndChunksOrderByColumns.md)
 - [Carbon::EmbeddingsAndChunksQueryInput](docs/EmbeddingsAndChunksQueryInput.md)
 - [Carbon::EmbeddingsAndChunksResponse](docs/EmbeddingsAndChunksResponse.md)
 - [Carbon::ExternalFileSyncStatuses](docs/ExternalFileSyncStatuses.md)
 - [Carbon::ExternalSourceItem](docs/ExternalSourceItem.md)
 - [Carbon::FetchURLsResponse](docs/FetchURLsResponse.md)
 - [Carbon::FileContentTypes](docs/FileContentTypes.md)
 - [Carbon::FileContentTypesNullable](docs/FileContentTypesNullable.md)
 - [Carbon::FileFormats](docs/FileFormats.md)
 - [Carbon::FileFormatsNullable](docs/FileFormatsNullable.md)
 - [Carbon::FileStatistics](docs/FileStatistics.md)
 - [Carbon::FileStatisticsNullable](docs/FileStatisticsNullable.md)
 - [Carbon::FreshDeskConnectRequest](docs/FreshDeskConnectRequest.md)
 - [Carbon::GenericSuccessResponse](docs/GenericSuccessResponse.md)
 - [Carbon::GetEmbeddingDocumentsBody](docs/GetEmbeddingDocumentsBody.md)
 - [Carbon::GmailSyncInput](docs/GmailSyncInput.md)
 - [Carbon::HTTPValidationError](docs/HTTPValidationError.md)
 - [Carbon::HybridSearchTuningParams](docs/HybridSearchTuningParams.md)
 - [Carbon::HybridSearchTuningParamsNullable](docs/HybridSearchTuningParamsNullable.md)
 - [Carbon::ListDataSourceItemsRequest](docs/ListDataSourceItemsRequest.md)
 - [Carbon::ListDataSourceItemsResponse](docs/ListDataSourceItemsResponse.md)
 - [Carbon::ListRequest](docs/ListRequest.md)
 - [Carbon::ListResponse](docs/ListResponse.md)
 - [Carbon::LocationPropertyInner](docs/LocationPropertyInner.md)
 - [Carbon::ModifyUserConfigurationInput](docs/ModifyUserConfigurationInput.md)
 - [Carbon::OAuthURLRequest](docs/OAuthURLRequest.md)
 - [Carbon::OrderDir](docs/OrderDir.md)
 - [Carbon::OrganizationResponse](docs/OrganizationResponse.md)
 - [Carbon::OrganizationUserDataSourceAPI](docs/OrganizationUserDataSourceAPI.md)
 - [Carbon::OrganizationUserDataSourceFilters](docs/OrganizationUserDataSourceFilters.md)
 - [Carbon::OrganizationUserDataSourceOrderByColumns](docs/OrganizationUserDataSourceOrderByColumns.md)
 - [Carbon::OrganizationUserDataSourceQueryInput](docs/OrganizationUserDataSourceQueryInput.md)
 - [Carbon::OrganizationUserDataSourceResponse](docs/OrganizationUserDataSourceResponse.md)
 - [Carbon::OrganizationUserFileTagCreate](docs/OrganizationUserFileTagCreate.md)
 - [Carbon::OrganizationUserFileTagsRemove](docs/OrganizationUserFileTagsRemove.md)
 - [Carbon::OrganizationUserFilesToSyncFilters](docs/OrganizationUserFilesToSyncFilters.md)
 - [Carbon::OrganizationUserFilesToSyncOrderByTypes](docs/OrganizationUserFilesToSyncOrderByTypes.md)
 - [Carbon::OrganizationUserFilesToSyncQueryInput](docs/OrganizationUserFilesToSyncQueryInput.md)
 - [Carbon::OutlookSyncInput](docs/OutlookSyncInput.md)
 - [Carbon::Pagination](docs/Pagination.md)
 - [Carbon::PresignedURLResponse](docs/PresignedURLResponse.md)
 - [Carbon::RSSFeedInput](docs/RSSFeedInput.md)
 - [Carbon::RankProperty](docs/RankProperty.md)
 - [Carbon::RawTextInput](docs/RawTextInput.md)
 - [Carbon::RawTranscriptPropertyInnerValue](docs/RawTranscriptPropertyInnerValue.md)
 - [Carbon::ResyncFileQueryInput](docs/ResyncFileQueryInput.md)
 - [Carbon::RevokeAccessTokenInput](docs/RevokeAccessTokenInput.md)
 - [Carbon::S3AuthRequest](docs/S3AuthRequest.md)
 - [Carbon::S3FileSyncInput](docs/S3FileSyncInput.md)
 - [Carbon::S3GetFileInput](docs/S3GetFileInput.md)
 - [Carbon::SingleChunksAndEmbeddingsUploadInput](docs/SingleChunksAndEmbeddingsUploadInput.md)
 - [Carbon::SitemapScrapeRequest](docs/SitemapScrapeRequest.md)
 - [Carbon::SourceProperty](docs/SourceProperty.md)
 - [Carbon::SyncDirectoryRequest](docs/SyncDirectoryRequest.md)
 - [Carbon::SyncFilesRequest](docs/SyncFilesRequest.md)
 - [Carbon::Tags](docs/Tags.md)
 - [Carbon::Tags1](docs/Tags1.md)
 - [Carbon::TextEmbeddingGenerators](docs/TextEmbeddingGenerators.md)
 - [Carbon::TokenResponse](docs/TokenResponse.md)
 - [Carbon::UploadFileFromUrlInput](docs/UploadFileFromUrlInput.md)
 - [Carbon::UserFile](docs/UserFile.md)
 - [Carbon::UserFilesV2](docs/UserFilesV2.md)
 - [Carbon::UserRequestContent](docs/UserRequestContent.md)
 - [Carbon::UserResponse](docs/UserResponse.md)
 - [Carbon::ValidationError](docs/ValidationError.md)
 - [Carbon::Webhook](docs/Webhook.md)
 - [Carbon::WebhookFilters](docs/WebhookFilters.md)
 - [Carbon::WebhookNoKey](docs/WebhookNoKey.md)
 - [Carbon::WebhookOrderByColumns](docs/WebhookOrderByColumns.md)
 - [Carbon::WebhookQueryInput](docs/WebhookQueryInput.md)
 - [Carbon::WebhookQueryResponse](docs/WebhookQueryResponse.md)
 - [Carbon::WebscrapeRequest](docs/WebscrapeRequest.md)
 - [Carbon::WhiteLabelingResponse](docs/WhiteLabelingResponse.md)
 - [Carbon::YoutubeTranscriptResponse](docs/YoutubeTranscriptResponse.md)


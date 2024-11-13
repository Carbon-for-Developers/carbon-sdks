<div align="center">

[![Visit Carbon](./header.png)](https://carbon.ai)

# [Carbon](https://carbon.ai)

Connect external data to LLMs, no matter the source.

[![Go Reference](https://pkg.go.dev/badge/github.com/Carbon-for-Developers/carbon-sdks/go)](https://pkg.go.dev/github.com/Carbon-for-Developers/carbon-sdks/go@v0.2.49)

</div>

## Installation

Add to your project:

```shell
go get github.com/Carbon-for-Developers/carbon-sdks/go
```

## Getting Started

```golang
package main

import (
	"fmt"

	carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
	// 1) Get an access token
	configuration := carbon.NewConfiguration()
	configuration.SetApiKey("API_KEY")
	configuration.SetCustomerId("CUSTOMER_ID")
	client := carbon.NewAPIClient(configuration)
	accessTokenRequest := client.AuthApi.GetAccessToken()
	accessTokenResponse, _, _ := accessTokenRequest.Execute()
	accessToken := accessTokenResponse.GetAccessToken()

	// 2) Use the access token to make requests
	configuration = carbon.NewConfiguration()
	configuration.SetAccessToken(accessToken)
	client = carbon.NewAPIClient(configuration)

	whiteLabelingRequest := client.AuthApi.GetWhiteLabeling()
	whiteLabelingResponse, _, _ := whiteLabelingRequest.Execute()
	integrations := whiteLabelingResponse.GetIntegrations()
	fmt.Println(integrations)
}

```

## Documentation for API Endpoints

All URIs are relative to *https://api.carbon.ai*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthApi* | [**GetAccessToken**](docs/AuthApi.md#getaccesstoken) | **Get** /auth/v1/access_token | Get Access Token
*AuthApi* | [**GetWhiteLabeling**](docs/AuthApi.md#getwhitelabeling) | **Get** /auth/v1/white_labeling | Get White Labeling
*CRMApi* | [**GetAccount**](docs/CRMApi.md#getaccount) | **Get** /integrations/data/crm/accounts/{id} | Get Account
*CRMApi* | [**GetAccounts**](docs/CRMApi.md#getaccounts) | **Post** /integrations/data/crm/accounts | Get Accounts
*CRMApi* | [**GetContact**](docs/CRMApi.md#getcontact) | **Get** /integrations/data/crm/contacts/{id} | Get Contact
*CRMApi* | [**GetContacts**](docs/CRMApi.md#getcontacts) | **Post** /integrations/data/crm/contacts | Get Contacts
*CRMApi* | [**GetLead**](docs/CRMApi.md#getlead) | **Get** /integrations/data/crm/leads/{id} | Get Lead
*CRMApi* | [**GetLeads**](docs/CRMApi.md#getleads) | **Post** /integrations/data/crm/leads | Get Leads
*CRMApi* | [**GetOpportunities**](docs/CRMApi.md#getopportunities) | **Post** /integrations/data/crm/opportunities | Get Opportunities
*CRMApi* | [**GetOpportunity**](docs/CRMApi.md#getopportunity) | **Get** /integrations/data/crm/opportunities/{id} | Get Opportunity
*DataSourcesApi* | [**AddTags**](docs/DataSourcesApi.md#addtags) | **Post** /data_sources/tags/add | Add Data Source Tags
*DataSourcesApi* | [**Query**](docs/DataSourcesApi.md#query) | **Post** /data_sources | Data Sources
*DataSourcesApi* | [**QueryUserDataSources**](docs/DataSourcesApi.md#queryuserdatasources) | **Post** /user_data_sources | User Data Sources
*DataSourcesApi* | [**RemoveTags**](docs/DataSourcesApi.md#removetags) | **Post** /data_sources/tags/remove | Remove Data Source Tags
*DataSourcesApi* | [**RevokeAccessToken**](docs/DataSourcesApi.md#revokeaccesstoken) | **Post** /revoke_access_token | Revoke Access Token
*EmbeddingsApi* | [**GetDocuments**](docs/EmbeddingsApi.md#getdocuments) | **Post** /embeddings | Embeddings
*EmbeddingsApi* | [**GetEmbeddingsAndChunks**](docs/EmbeddingsApi.md#getembeddingsandchunks) | **Post** /text_chunks | Retrieve Embeddings And Content
*EmbeddingsApi* | [**List**](docs/EmbeddingsApi.md#list) | **Post** /list_chunks_and_embeddings | Retrieve Embeddings And Content V2
*EmbeddingsApi* | [**UploadChunksAndEmbeddings**](docs/EmbeddingsApi.md#uploadchunksandembeddings) | **Post** /upload_chunks_and_embeddings | Upload Chunks And Embeddings
*FilesApi* | [**CreateUserFileTags**](docs/FilesApi.md#createuserfiletags) | **Post** /create_user_file_tags | Create File Tags
*FilesApi* | [**Delete**](docs/FilesApi.md#delete) | **Delete** /deletefile/{file_id} | Delete File Endpoint
*FilesApi* | [**DeleteFileTags**](docs/FilesApi.md#deletefiletags) | **Post** /delete_user_file_tags | Delete File Tags
*FilesApi* | [**DeleteMany**](docs/FilesApi.md#deletemany) | **Post** /delete_files | Delete Files Endpoint
*FilesApi* | [**DeleteV2**](docs/FilesApi.md#deletev2) | **Post** /delete_files_v2 | Delete Files V2 Endpoint
*FilesApi* | [**GetParsedFile**](docs/FilesApi.md#getparsedfile) | **Get** /parsed_file/{file_id} | Parsed File
*FilesApi* | [**GetRawFile**](docs/FilesApi.md#getrawfile) | **Get** /raw_file/{file_id} | Raw File
*FilesApi* | [**ModifyColdStorageParameters**](docs/FilesApi.md#modifycoldstorageparameters) | **Post** /modify_cold_storage_parameters | Modify Cold Storage Parameters
*FilesApi* | [**MoveToHotStorage**](docs/FilesApi.md#movetohotstorage) | **Post** /move_to_hot_storage | Move To Hot Storage
*FilesApi* | [**QueryUserFiles**](docs/FilesApi.md#queryuserfiles) | **Post** /user_files_v2 | User Files V2
*FilesApi* | [**QueryUserFilesDeprecated**](docs/FilesApi.md#queryuserfilesdeprecated) | **Post** /user_files | User Files
*FilesApi* | [**Resync**](docs/FilesApi.md#resync) | **Post** /resync_file | Resync File
*FilesApi* | [**Upload**](docs/FilesApi.md#upload) | **Post** /uploadfile | Create Upload File
*FilesApi* | [**UploadFromUrl**](docs/FilesApi.md#uploadfromurl) | **Post** /upload_file_from_url | Create Upload File From Url
*FilesApi* | [**UploadText**](docs/FilesApi.md#uploadtext) | **Post** /upload_text | Create Raw Text
*GithubApi* | [**GetIssue**](docs/GithubApi.md#getissue) | **Get** /integrations/data/github/issues/{issue_number} | Issue
*GithubApi* | [**GetIssues**](docs/GithubApi.md#getissues) | **Post** /integrations/data/github/issues | Issues
*GithubApi* | [**GetPr**](docs/GithubApi.md#getpr) | **Get** /integrations/data/github/pull_requests/{pull_number} | Get Pr
*GithubApi* | [**GetPrComments**](docs/GithubApi.md#getprcomments) | **Post** /integrations/data/github/pull_requests/comments | Pr Comments
*GithubApi* | [**GetPrCommits**](docs/GithubApi.md#getprcommits) | **Post** /integrations/data/github/pull_requests/commits | Pr Commits
*GithubApi* | [**GetPrFiles**](docs/GithubApi.md#getprfiles) | **Post** /integrations/data/github/pull_requests/files | Pr Files
*GithubApi* | [**GetPullRequests**](docs/GithubApi.md#getpullrequests) | **Post** /integrations/data/github/pull_requests | Get Prs
*IntegrationsApi* | [**Cancel**](docs/IntegrationsApi.md#cancel) | **Post** /integrations/items/sync/cancel | Cancel Data Source Items Sync
*IntegrationsApi* | [**ConnectDataSource**](docs/IntegrationsApi.md#connectdatasource) | **Post** /integrations/connect | Connect Data Source
*IntegrationsApi* | [**ConnectFreshdesk**](docs/IntegrationsApi.md#connectfreshdesk) | **Post** /integrations/freshdesk | Freshdesk Connect
*IntegrationsApi* | [**ConnectGitbook**](docs/IntegrationsApi.md#connectgitbook) | **Post** /integrations/gitbook | Gitbook Connect
*IntegrationsApi* | [**ConnectGuru**](docs/IntegrationsApi.md#connectguru) | **Post** /integrations/guru | Guru Connect
*IntegrationsApi* | [**CreateAwsIamUser**](docs/IntegrationsApi.md#createawsiamuser) | **Post** /integrations/s3 | S3 Auth
*IntegrationsApi* | [**GetOauthUrl**](docs/IntegrationsApi.md#getoauthurl) | **Post** /integrations/oauth_url | Get Oauth Url
*IntegrationsApi* | [**ListConfluencePages**](docs/IntegrationsApi.md#listconfluencepages) | **Post** /integrations/confluence/list | Confluence List
*IntegrationsApi* | [**ListConversations**](docs/IntegrationsApi.md#listconversations) | **Get** /integrations/slack/conversations | Slack List Conversations
*IntegrationsApi* | [**ListDataSourceItems**](docs/IntegrationsApi.md#listdatasourceitems) | **Post** /integrations/items/list | List Data Source Items
*IntegrationsApi* | [**ListFolders**](docs/IntegrationsApi.md#listfolders) | **Get** /integrations/outlook/user_folders | Outlook Folders
*IntegrationsApi* | [**ListGitbookSpaces**](docs/IntegrationsApi.md#listgitbookspaces) | **Get** /integrations/gitbook/spaces | Gitbook Spaces
*IntegrationsApi* | [**ListLabels**](docs/IntegrationsApi.md#listlabels) | **Get** /integrations/gmail/user_labels | Gmail Labels
*IntegrationsApi* | [**ListOutlookCategories**](docs/IntegrationsApi.md#listoutlookcategories) | **Get** /integrations/outlook/user_categories | Outlook Categories
*IntegrationsApi* | [**ListRepos**](docs/IntegrationsApi.md#listrepos) | **Get** /integrations/github/repos | Github List Repos
*IntegrationsApi* | [**ListSharepointSites**](docs/IntegrationsApi.md#listsharepointsites) | **Get** /integrations/sharepoint/sites/list | List Sharepoint Sites
*IntegrationsApi* | [**SyncAzureBlobFiles**](docs/IntegrationsApi.md#syncazureblobfiles) | **Post** /integrations/azure_blob_storage/files | Azure Blob Files
*IntegrationsApi* | [**SyncAzureBlobStorage**](docs/IntegrationsApi.md#syncazureblobstorage) | **Post** /integrations/azure_blob_storage | Azure Blob Storage Auth
*IntegrationsApi* | [**SyncConfluence**](docs/IntegrationsApi.md#syncconfluence) | **Post** /integrations/confluence/sync | Confluence Sync
*IntegrationsApi* | [**SyncDataSourceItems**](docs/IntegrationsApi.md#syncdatasourceitems) | **Post** /integrations/items/sync | Sync Data Source Items
*IntegrationsApi* | [**SyncFiles**](docs/IntegrationsApi.md#syncfiles) | **Post** /integrations/files/sync | Sync Files
*IntegrationsApi* | [**SyncGitHub**](docs/IntegrationsApi.md#syncgithub) | **Post** /integrations/github | Github Connect
*IntegrationsApi* | [**SyncGitbook**](docs/IntegrationsApi.md#syncgitbook) | **Post** /integrations/gitbook/sync | Gitbook Sync
*IntegrationsApi* | [**SyncGmail**](docs/IntegrationsApi.md#syncgmail) | **Post** /integrations/gmail/sync | Gmail Sync
*IntegrationsApi* | [**SyncOutlook**](docs/IntegrationsApi.md#syncoutlook) | **Post** /integrations/outlook/sync | Outlook Sync
*IntegrationsApi* | [**SyncRepos**](docs/IntegrationsApi.md#syncrepos) | **Post** /integrations/github/sync_repos | Github Sync Repos
*IntegrationsApi* | [**SyncRssFeed**](docs/IntegrationsApi.md#syncrssfeed) | **Post** /integrations/rss_feed | Rss Feed
*IntegrationsApi* | [**SyncS3Files**](docs/IntegrationsApi.md#syncs3files) | **Post** /integrations/s3/files | S3 Files
*IntegrationsApi* | [**SyncSlack**](docs/IntegrationsApi.md#syncslack) | **Post** /integrations/slack/sync | Slack Sync
*OrganizationsApi* | [**Get**](docs/OrganizationsApi.md#get) | **Get** /organization | Get Organization
*OrganizationsApi* | [**Update**](docs/OrganizationsApi.md#update) | **Post** /organization/update | Update Organization
*OrganizationsApi* | [**UpdateStats**](docs/OrganizationsApi.md#updatestats) | **Post** /organization/statistics | Update Organization Statistics
*UsersApi* | [**Delete**](docs/UsersApi.md#delete) | **Post** /delete_users | Delete Users
*UsersApi* | [**Get**](docs/UsersApi.md#get) | **Post** /user | User Endpoint
*UsersApi* | [**List**](docs/UsersApi.md#list) | **Post** /list_users | List Users Endpoint
*UsersApi* | [**ToggleUserFeatures**](docs/UsersApi.md#toggleuserfeatures) | **Post** /modify_user_configuration | Toggle User Features
*UsersApi* | [**UpdateUsers**](docs/UsersApi.md#updateusers) | **Post** /update_users | Update Users
*UsersApi* | [**WhoAmI**](docs/UsersApi.md#whoami) | **Get** /whoami | Me Endpoint
*UtilitiesApi* | [**FetchUrls**](docs/UtilitiesApi.md#fetchurls) | **Get** /fetch_urls | Fetch Urls
*UtilitiesApi* | [**FetchWebpage**](docs/UtilitiesApi.md#fetchwebpage) | **Post** /fetch_webpage | Fetch Urls V2
*UtilitiesApi* | [**FetchYoutubeTranscripts**](docs/UtilitiesApi.md#fetchyoutubetranscripts) | **Get** /fetch_youtube_transcript | Fetch Youtube Transcripts
*UtilitiesApi* | [**ProcessSitemap**](docs/UtilitiesApi.md#processsitemap) | **Get** /process_sitemap | Sitemap
*UtilitiesApi* | [**ScrapeSitemap**](docs/UtilitiesApi.md#scrapesitemap) | **Post** /scrape_sitemap | Scrape Sitemap
*UtilitiesApi* | [**ScrapeWeb**](docs/UtilitiesApi.md#scrapeweb) | **Post** /web_scrape | Web Scrape
*UtilitiesApi* | [**SearchUrls**](docs/UtilitiesApi.md#searchurls) | **Get** /search_urls | Search Urls
*UtilitiesApi* | [**UserWebpages**](docs/UtilitiesApi.md#userwebpages) | **Post** /user_webpages | User Web Pages
*WebhooksApi* | [**AddUrl**](docs/WebhooksApi.md#addurl) | **Post** /add_webhook | Add Webhook Url
*WebhooksApi* | [**DeleteUrl**](docs/WebhooksApi.md#deleteurl) | **Delete** /delete_webhook/{webhook_id} | Delete Webhook Url
*WebhooksApi* | [**Urls**](docs/WebhooksApi.md#urls) | **Post** /webhooks | Webhook Urls
*WhiteLabelApi* | [**Create**](docs/WhiteLabelApi.md#create) | **Post** /white_label/create | Create White Labels
*WhiteLabelApi* | [**Delete**](docs/WhiteLabelApi.md#delete) | **Post** /white_label/delete | Delete White Labels
*WhiteLabelApi* | [**List**](docs/WhiteLabelApi.md#list) | **Post** /white_label/list | List White Labels
*WhiteLabelApi* | [**Update**](docs/WhiteLabelApi.md#update) | **Post** /white_label/update | Update White Label


## Documentation For Models

 - [Account](docs/Account.md)
 - [AccountFilters](docs/AccountFilters.md)
 - [AccountResponse](docs/AccountResponse.md)
 - [AccountsOrderBy](docs/AccountsOrderBy.md)
 - [AccountsOrderByNullable](docs/AccountsOrderByNullable.md)
 - [AccountsRequest](docs/AccountsRequest.md)
 - [AddDataSourceTagsInput](docs/AddDataSourceTagsInput.md)
 - [AddWebhookProps](docs/AddWebhookProps.md)
 - [Address](docs/Address.md)
 - [AutoSyncEnabledSourcesProperty](docs/AutoSyncEnabledSourcesProperty.md)
 - [AzureBlobAuthRequest](docs/AzureBlobAuthRequest.md)
 - [AzureBlobFileSyncInput](docs/AzureBlobFileSyncInput.md)
 - [AzureBlobGetFileInput](docs/AzureBlobGetFileInput.md)
 - [AzureBlobStorageAuthentication](docs/AzureBlobStorageAuthentication.md)
 - [AzureBlobStorageTypes](docs/AzureBlobStorageTypes.md)
 - [BaseIncludes](docs/BaseIncludes.md)
 - [BodyCreateUploadFileUploadfilePost](docs/BodyCreateUploadFileUploadfilePost.md)
 - [ChunkProperties](docs/ChunkProperties.md)
 - [ChunkPropertiesNullable](docs/ChunkPropertiesNullable.md)
 - [ChunksAndEmbeddings](docs/ChunksAndEmbeddings.md)
 - [ChunksAndEmbeddingsUploadInput](docs/ChunksAndEmbeddingsUploadInput.md)
 - [ColdStorageProps](docs/ColdStorageProps.md)
 - [CommentsInput](docs/CommentsInput.md)
 - [CommentsOrderBy](docs/CommentsOrderBy.md)
 - [CommentsResponse](docs/CommentsResponse.md)
 - [Commit](docs/Commit.md)
 - [CommitUser](docs/CommitUser.md)
 - [CommitUserNullable](docs/CommitUserNullable.md)
 - [CommitsInput](docs/CommitsInput.md)
 - [CommitsResponse](docs/CommitsResponse.md)
 - [ConfigurationKeys](docs/ConfigurationKeys.md)
 - [ConfluenceAuthentication](docs/ConfluenceAuthentication.md)
 - [ConfluenceFileTypes](docs/ConfluenceFileTypes.md)
 - [ConnectDataSourceInput](docs/ConnectDataSourceInput.md)
 - [ConnectDataSourceResponse](docs/ConnectDataSourceResponse.md)
 - [Contact](docs/Contact.md)
 - [ContactFilters](docs/ContactFilters.md)
 - [ContactsOrderBy](docs/ContactsOrderBy.md)
 - [ContactsOrderByNullable](docs/ContactsOrderByNullable.md)
 - [ContactsRequest](docs/ContactsRequest.md)
 - [ContactsResponse](docs/ContactsResponse.md)
 - [Credentials](docs/Credentials.md)
 - [CustomCredentialsType](docs/CustomCredentialsType.md)
 - [DataSourceConfiguration](docs/DataSourceConfiguration.md)
 - [DataSourceExtendedInput](docs/DataSourceExtendedInput.md)
 - [DataSourceLastSyncActions](docs/DataSourceLastSyncActions.md)
 - [DataSourceSyncStatuses](docs/DataSourceSyncStatuses.md)
 - [DataSourceType](docs/DataSourceType.md)
 - [DataSourceTypeNullable](docs/DataSourceTypeNullable.md)
 - [DeleteFilesQueryInput](docs/DeleteFilesQueryInput.md)
 - [DeleteFilesV2QueryInput](docs/DeleteFilesV2QueryInput.md)
 - [DeleteUsersInput](docs/DeleteUsersInput.md)
 - [DeleteWhiteLabelRequest](docs/DeleteWhiteLabelRequest.md)
 - [DirectoryItem](docs/DirectoryItem.md)
 - [DocumentResponse](docs/DocumentResponse.md)
 - [DocumentResponseList](docs/DocumentResponseList.md)
 - [Email](docs/Email.md)
 - [EmbeddingAndChunk](docs/EmbeddingAndChunk.md)
 - [EmbeddingGenerators](docs/EmbeddingGenerators.md)
 - [EmbeddingGeneratorsNullable](docs/EmbeddingGeneratorsNullable.md)
 - [EmbeddingProperties](docs/EmbeddingProperties.md)
 - [EmbeddingStorageStatus](docs/EmbeddingStorageStatus.md)
 - [EmbeddingsAndChunksFilters](docs/EmbeddingsAndChunksFilters.md)
 - [EmbeddingsAndChunksOrderByColumns](docs/EmbeddingsAndChunksOrderByColumns.md)
 - [EmbeddingsAndChunksQueryInput](docs/EmbeddingsAndChunksQueryInput.md)
 - [EmbeddingsAndChunksQueryInputV2](docs/EmbeddingsAndChunksQueryInputV2.md)
 - [EmbeddingsAndChunksResponse](docs/EmbeddingsAndChunksResponse.md)
 - [Event](docs/Event.md)
 - [ExternalFileSyncStatuses](docs/ExternalFileSyncStatuses.md)
 - [ExternalSourceItem](docs/ExternalSourceItem.md)
 - [ExternalSourceItemsOrderBy](docs/ExternalSourceItemsOrderBy.md)
 - [FetchURLsRequest](docs/FetchURLsRequest.md)
 - [FetchURLsResponse](docs/FetchURLsResponse.md)
 - [FileContentTypes](docs/FileContentTypes.md)
 - [FileContentTypesNullable](docs/FileContentTypesNullable.md)
 - [FileFormats](docs/FileFormats.md)
 - [FileFormatsNullable](docs/FileFormatsNullable.md)
 - [FileStatistics](docs/FileStatistics.md)
 - [FileStatisticsNullable](docs/FileStatisticsNullable.md)
 - [FileStatus](docs/FileStatus.md)
 - [FileSyncConfig](docs/FileSyncConfig.md)
 - [FileSyncConfigNullable](docs/FileSyncConfigNullable.md)
 - [FilesInput](docs/FilesInput.md)
 - [FilesResponse](docs/FilesResponse.md)
 - [FreshDeskConnectRequest](docs/FreshDeskConnectRequest.md)
 - [FreskdeskAuthentication](docs/FreskdeskAuthentication.md)
 - [GDOSWhiteLabelInput](docs/GDOSWhiteLabelInput.md)
 - [GDOSWhiteLabelInput1](docs/GDOSWhiteLabelInput1.md)
 - [GenericSuccessResponse](docs/GenericSuccessResponse.md)
 - [GetEmbeddingDocumentsBody](docs/GetEmbeddingDocumentsBody.md)
 - [GitbookAuthetication](docs/GitbookAuthetication.md)
 - [GitbookConnectRequest](docs/GitbookConnectRequest.md)
 - [GitbookFileTypes](docs/GitbookFileTypes.md)
 - [GitbookSyncRequest](docs/GitbookSyncRequest.md)
 - [GithubAuthentication](docs/GithubAuthentication.md)
 - [GithubConnectRequest](docs/GithubConnectRequest.md)
 - [GithubFetchReposRequest](docs/GithubFetchReposRequest.md)
 - [GithubFileTypes](docs/GithubFileTypes.md)
 - [GmailSyncInput](docs/GmailSyncInput.md)
 - [GongAuthentication](docs/GongAuthentication.md)
 - [GongFileTypes](docs/GongFileTypes.md)
 - [GoogleDriveCredentials](docs/GoogleDriveCredentials.md)
 - [GoogleDriveCredentialsProperty](docs/GoogleDriveCredentialsProperty.md)
 - [GoogleDriveWhiteLabelInput](docs/GoogleDriveWhiteLabelInput.md)
 - [GuruAuthentication](docs/GuruAuthentication.md)
 - [GuruConnectRequest](docs/GuruConnectRequest.md)
 - [GuruFileTypes](docs/GuruFileTypes.md)
 - [HSNFileTypes](docs/HSNFileTypes.md)
 - [HSNFileTypes1](docs/HSNFileTypes1.md)
 - [HSNFileTypes2](docs/HSNFileTypes2.md)
 - [HSNFileTypes3](docs/HSNFileTypes3.md)
 - [HTTPValidationError](docs/HTTPValidationError.md)
 - [HeadOrBase](docs/HeadOrBase.md)
 - [HelpdeskFileTypes](docs/HelpdeskFileTypes.md)
 - [HybridSearchTuningParams](docs/HybridSearchTuningParams.md)
 - [HybridSearchTuningParamsNullable](docs/HybridSearchTuningParamsNullable.md)
 - [IntercomFileTypes](docs/IntercomFileTypes.md)
 - [Issue](docs/Issue.md)
 - [IssuePR](docs/IssuePR.md)
 - [IssuePRNullable](docs/IssuePRNullable.md)
 - [IssuesFilter](docs/IssuesFilter.md)
 - [IssuesInput](docs/IssuesInput.md)
 - [IssuesOrderBy](docs/IssuesOrderBy.md)
 - [IssuesResponse](docs/IssuesResponse.md)
 - [Label](docs/Label.md)
 - [Lead](docs/Lead.md)
 - [LeadFilters](docs/LeadFilters.md)
 - [LeadsOrderBy](docs/LeadsOrderBy.md)
 - [LeadsOrderByNullable](docs/LeadsOrderByNullable.md)
 - [LeadsRequest](docs/LeadsRequest.md)
 - [LeadsResponse](docs/LeadsResponse.md)
 - [ListDataSourceItemsRequest](docs/ListDataSourceItemsRequest.md)
 - [ListDataSourceItemsResponse](docs/ListDataSourceItemsResponse.md)
 - [ListItemsFilters](docs/ListItemsFilters.md)
 - [ListItemsFiltersNullable](docs/ListItemsFiltersNullable.md)
 - [ListRequest](docs/ListRequest.md)
 - [ListResponse](docs/ListResponse.md)
 - [ListUserResponse](docs/ListUserResponse.md)
 - [ListUsersFilters](docs/ListUsersFilters.md)
 - [ListUsersOrderByTypes](docs/ListUsersOrderByTypes.md)
 - [ListUsersRequest](docs/ListUsersRequest.md)
 - [ListWhiteLabelRequest](docs/ListWhiteLabelRequest.md)
 - [LocationPropertyInner](docs/LocationPropertyInner.md)
 - [ModifyColdStorageParametersQueryInput](docs/ModifyColdStorageParametersQueryInput.md)
 - [ModifyUserConfigurationInput](docs/ModifyUserConfigurationInput.md)
 - [MoveToHotStorageQueryInput](docs/MoveToHotStorageQueryInput.md)
 - [MultiModalEmbeddingGenerators](docs/MultiModalEmbeddingGenerators.md)
 - [NotionAuthentication](docs/NotionAuthentication.md)
 - [OANDSCZGFB](docs/OANDSCZGFB.md)
 - [OAuthAuthentication](docs/OAuthAuthentication.md)
 - [OAuthURLRequest](docs/OAuthURLRequest.md)
 - [OauthBasedConnectors](docs/OauthBasedConnectors.md)
 - [OneDriveAuthentication](docs/OneDriveAuthentication.md)
 - [OneDriveSharepointCredentials](docs/OneDriveSharepointCredentials.md)
 - [OneDriveSharepointWhiteLabelInput](docs/OneDriveSharepointWhiteLabelInput.md)
 - [OpportunitiesOrderBy](docs/OpportunitiesOrderBy.md)
 - [OpportunitiesOrderByNullable](docs/OpportunitiesOrderByNullable.md)
 - [OpportunitiesRequest](docs/OpportunitiesRequest.md)
 - [OpportunitiesResponse](docs/OpportunitiesResponse.md)
 - [Opportunity](docs/Opportunity.md)
 - [OpportunityFilters](docs/OpportunityFilters.md)
 - [OpportunityStatus](docs/OpportunityStatus.md)
 - [OpportunityStatusNullable](docs/OpportunityStatusNullable.md)
 - [OrderDir](docs/OrderDir.md)
 - [OrderDirV2](docs/OrderDirV2.md)
 - [OrderDirV2Nullable](docs/OrderDirV2Nullable.md)
 - [OrganizationResponse](docs/OrganizationResponse.md)
 - [OrganizationUserDataSourceAPI](docs/OrganizationUserDataSourceAPI.md)
 - [OrganizationUserDataSourceFilters](docs/OrganizationUserDataSourceFilters.md)
 - [OrganizationUserDataSourceOrderByColumns](docs/OrganizationUserDataSourceOrderByColumns.md)
 - [OrganizationUserDataSourceQueryInput](docs/OrganizationUserDataSourceQueryInput.md)
 - [OrganizationUserDataSourceResponse](docs/OrganizationUserDataSourceResponse.md)
 - [OrganizationUserFileTagCreate](docs/OrganizationUserFileTagCreate.md)
 - [OrganizationUserFileTagsRemove](docs/OrganizationUserFileTagsRemove.md)
 - [OrganizationUserFilesToSyncFilters](docs/OrganizationUserFilesToSyncFilters.md)
 - [OrganizationUserFilesToSyncOrderByTypes](docs/OrganizationUserFilesToSyncOrderByTypes.md)
 - [OrganizationUserFilesToSyncQueryInput](docs/OrganizationUserFilesToSyncQueryInput.md)
 - [OuthURLResponse](docs/OuthURLResponse.md)
 - [OutlookSyncInput](docs/OutlookSyncInput.md)
 - [PRComment](docs/PRComment.md)
 - [PRCommit](docs/PRCommit.md)
 - [PRFile](docs/PRFile.md)
 - [PROrderBy](docs/PROrderBy.md)
 - [PRState](docs/PRState.md)
 - [PRStateInput](docs/PRStateInput.md)
 - [Pagination](docs/Pagination.md)
 - [PartialAccount](docs/PartialAccount.md)
 - [PartialAccountNullable](docs/PartialAccountNullable.md)
 - [PartialContact](docs/PartialContact.md)
 - [PartialContactNullable](docs/PartialContactNullable.md)
 - [PartialOwner](docs/PartialOwner.md)
 - [PartialOwnerNullable](docs/PartialOwnerNullable.md)
 - [PhoneNumber](docs/PhoneNumber.md)
 - [PresignedURLResponse](docs/PresignedURLResponse.md)
 - [PullRequest](docs/PullRequest.md)
 - [PullRequestExtended](docs/PullRequestExtended.md)
 - [PullRequestFilters](docs/PullRequestFilters.md)
 - [PullRequestResponse](docs/PullRequestResponse.md)
 - [PullRequestsInput](docs/PullRequestsInput.md)
 - [RSSFeedInput](docs/RSSFeedInput.md)
 - [RankProperty](docs/RankProperty.md)
 - [RawTextInput](docs/RawTextInput.md)
 - [RawTranscriptPropertyInnerValue](docs/RawTranscriptPropertyInnerValue.md)
 - [RemoveDataSourceTagsInput](docs/RemoveDataSourceTagsInput.md)
 - [RerankParams](docs/RerankParams.md)
 - [RerankParamsNullable](docs/RerankParamsNullable.md)
 - [ResyncFileQueryInput](docs/ResyncFileQueryInput.md)
 - [RevokeAccessTokenInput](docs/RevokeAccessTokenInput.md)
 - [S3AuthRequest](docs/S3AuthRequest.md)
 - [S3Authentication](docs/S3Authentication.md)
 - [S3FileSyncInput](docs/S3FileSyncInput.md)
 - [S3GetFileInput](docs/S3GetFileInput.md)
 - [SIHFNZCGABTypes](docs/SIHFNZCGABTypes.md)
 - [SIHFNZCGABTypes1](docs/SIHFNZCGABTypes1.md)
 - [SalesforceAuthentication](docs/SalesforceAuthentication.md)
 - [SentWebhookPayload](docs/SentWebhookPayload.md)
 - [SentWebhookPayloadObject](docs/SentWebhookPayloadObject.md)
 - [SentWebhookPayloadObjectAdditionalInformation](docs/SentWebhookPayloadObjectAdditionalInformation.md)
 - [SentWebhookPayloadObjectObjectId](docs/SentWebhookPayloadObjectObjectId.md)
 - [SentWebhookRequestBody](docs/SentWebhookRequestBody.md)
 - [ServiceNowAuthentication](docs/ServiceNowAuthentication.md)
 - [ServiceNowCredentials](docs/ServiceNowCredentials.md)
 - [ServiceNowCredentialsNullable](docs/ServiceNowCredentialsNullable.md)
 - [ServiceNowFileTypes](docs/ServiceNowFileTypes.md)
 - [SharedSourceItemTypes](docs/SharedSourceItemTypes.md)
 - [SharepointAuthentication](docs/SharepointAuthentication.md)
 - [SimpleOAuthDataSources](docs/SimpleOAuthDataSources.md)
 - [SingleChunksAndEmbeddingsUploadInput](docs/SingleChunksAndEmbeddingsUploadInput.md)
 - [SitemapScrapeRequest](docs/SitemapScrapeRequest.md)
 - [SlackFilters](docs/SlackFilters.md)
 - [SlackSyncRequest](docs/SlackSyncRequest.md)
 - [SourceProperty](docs/SourceProperty.md)
 - [StorageFileFormats](docs/StorageFileFormats.md)
 - [SyncDirectoryRequest](docs/SyncDirectoryRequest.md)
 - [SyncFilesIds](docs/SyncFilesIds.md)
 - [SyncFilesIds1](docs/SyncFilesIds1.md)
 - [SyncFilesRequest](docs/SyncFilesRequest.md)
 - [SyncFilesRequestIdsInner](docs/SyncFilesRequestIdsInner.md)
 - [SyncOptions](docs/SyncOptions.md)
 - [TMEmbeddingGenerators](docs/TMEmbeddingGenerators.md)
 - [Tags](docs/Tags.md)
 - [Tags1](docs/Tags1.md)
 - [Task](docs/Task.md)
 - [Team](docs/Team.md)
 - [TextEmbeddingGenerators](docs/TextEmbeddingGenerators.md)
 - [TokenResponse](docs/TokenResponse.md)
 - [TranscriptionService](docs/TranscriptionService.md)
 - [TranscriptionServiceNullable](docs/TranscriptionServiceNullable.md)
 - [Tree](docs/Tree.md)
 - [UpdateOrganizationInput](docs/UpdateOrganizationInput.md)
 - [UpdateUsersInput](docs/UpdateUsersInput.md)
 - [UploadFileFromUrlInput](docs/UploadFileFromUrlInput.md)
 - [User](docs/User.md)
 - [UserConfiguration](docs/UserConfiguration.md)
 - [UserConfigurationNullable](docs/UserConfigurationNullable.md)
 - [UserFile](docs/UserFile.md)
 - [UserFilesV2](docs/UserFilesV2.md)
 - [UserListResponse](docs/UserListResponse.md)
 - [UserNullable](docs/UserNullable.md)
 - [UserRequestContent](docs/UserRequestContent.md)
 - [UserResponse](docs/UserResponse.md)
 - [UserWebPageOrderByTypes](docs/UserWebPageOrderByTypes.md)
 - [UserWebPagesFilters](docs/UserWebPagesFilters.md)
 - [UserWebPagesRequest](docs/UserWebPagesRequest.md)
 - [ValidationError](docs/ValidationError.md)
 - [Webhook](docs/Webhook.md)
 - [WebhookFilters](docs/WebhookFilters.md)
 - [WebhookNoKey](docs/WebhookNoKey.md)
 - [WebhookOrderByColumns](docs/WebhookOrderByColumns.md)
 - [WebhookQueryInput](docs/WebhookQueryInput.md)
 - [WebhookQueryResponse](docs/WebhookQueryResponse.md)
 - [WebhookStatus](docs/WebhookStatus.md)
 - [WebscrapeRequest](docs/WebscrapeRequest.md)
 - [WhiteLabelFilters](docs/WhiteLabelFilters.md)
 - [WhiteLabelInput](docs/WhiteLabelInput.md)
 - [WhiteLabelOrderByColumns](docs/WhiteLabelOrderByColumns.md)
 - [WhiteLabelingResponse](docs/WhiteLabelingResponse.md)
 - [YoutubeTranscriptResponse](docs/YoutubeTranscriptResponse.md)
 - [ZendeskAuthentication](docs/ZendeskAuthentication.md)
 - [ZendeskFileTypes](docs/ZendeskFileTypes.md)
 - [ZoteroAuthentication](docs/ZoteroAuthentication.md)

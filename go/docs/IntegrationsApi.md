# APIClient.IntegrationsApi

All URIs are relative to *https://api.carbon.ai*

Method | Path | Description
------------- | ------------- | -------------
[**Cancel**](IntegrationsApi.md#Cancel) | **Post** /integrations/items/sync/cancel | Cancel Data Source Items Sync
[**ConnectDataSource**](IntegrationsApi.md#ConnectDataSource) | **Post** /integrations/connect | Connect Data Source
[**ConnectFreshdesk**](IntegrationsApi.md#ConnectFreshdesk) | **Post** /integrations/freshdesk | Freshdesk Connect
[**ConnectGitbook**](IntegrationsApi.md#ConnectGitbook) | **Post** /integrations/gitbook | Gitbook Connect
[**ConnectGuru**](IntegrationsApi.md#ConnectGuru) | **Post** /integrations/guru | Guru Connect
[**CreateAwsIamUser**](IntegrationsApi.md#CreateAwsIamUser) | **Post** /integrations/s3 | S3 Auth
[**GetOauthUrl**](IntegrationsApi.md#GetOauthUrl) | **Post** /integrations/oauth_url | Get Oauth Url
[**ListConfluencePages**](IntegrationsApi.md#ListConfluencePages) | **Post** /integrations/confluence/list | Confluence List
[**ListConversations**](IntegrationsApi.md#ListConversations) | **Get** /integrations/slack/conversations | Slack List Conversations
[**ListDataSourceItems**](IntegrationsApi.md#ListDataSourceItems) | **Post** /integrations/items/list | List Data Source Items
[**ListFolders**](IntegrationsApi.md#ListFolders) | **Get** /integrations/outlook/user_folders | Outlook Folders
[**ListGitbookSpaces**](IntegrationsApi.md#ListGitbookSpaces) | **Get** /integrations/gitbook/spaces | Gitbook Spaces
[**ListLabels**](IntegrationsApi.md#ListLabels) | **Get** /integrations/gmail/user_labels | Gmail Labels
[**ListOutlookCategories**](IntegrationsApi.md#ListOutlookCategories) | **Get** /integrations/outlook/user_categories | Outlook Categories
[**ListRepos**](IntegrationsApi.md#ListRepos) | **Get** /integrations/github/repos | Github List Repos
[**SyncAzureBlobFiles**](IntegrationsApi.md#SyncAzureBlobFiles) | **Post** /integrations/azure_blob_storage/files | Azure Blob Files
[**SyncAzureBlobStorage**](IntegrationsApi.md#SyncAzureBlobStorage) | **Post** /integrations/azure_blob_storage | Azure Blob Storage Auth
[**SyncConfluence**](IntegrationsApi.md#SyncConfluence) | **Post** /integrations/confluence/sync | Confluence Sync
[**SyncDataSourceItems**](IntegrationsApi.md#SyncDataSourceItems) | **Post** /integrations/items/sync | Sync Data Source Items
[**SyncFiles**](IntegrationsApi.md#SyncFiles) | **Post** /integrations/files/sync | Sync Files
[**SyncGitHub**](IntegrationsApi.md#SyncGitHub) | **Post** /integrations/github | Github Connect
[**SyncGitbook**](IntegrationsApi.md#SyncGitbook) | **Post** /integrations/gitbook/sync | Gitbook Sync
[**SyncGmail**](IntegrationsApi.md#SyncGmail) | **Post** /integrations/gmail/sync | Gmail Sync
[**SyncOutlook**](IntegrationsApi.md#SyncOutlook) | **Post** /integrations/outlook/sync | Outlook Sync
[**SyncRepos**](IntegrationsApi.md#SyncRepos) | **Post** /integrations/github/sync_repos | Github Sync Repos
[**SyncRssFeed**](IntegrationsApi.md#SyncRssFeed) | **Post** /integrations/rss_feed | Rss Feed
[**SyncS3Files**](IntegrationsApi.md#SyncS3Files) | **Post** /integrations/s3/files | S3 Files
[**SyncSlack**](IntegrationsApi.md#SyncSlack) | **Post** /integrations/slack/sync | Slack Sync



## Cancel

Cancel Data Source Items Sync

### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    
    syncDirectoryRequest := *carbon.NewSyncDirectoryRequest(
        null,
    )
    
    request := client.IntegrationsApi.Cancel(
        syncDirectoryRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.Cancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `Cancel`: OrganizationUserDataSourceAPI
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.Cancel`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.Cancel.Id`: %v\n", resp.Id)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.Cancel.DataSourceExternalId`: %v\n", resp.DataSourceExternalId)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.Cancel.DataSourceType`: %v\n", resp.DataSourceType)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.Cancel.Token`: %v\n", resp.Token)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.Cancel.SyncStatus`: %v\n", resp.SyncStatus)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.Cancel.SourceItemsSyncedAt`: %v\n", resp.SourceItemsSyncedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.Cancel.OrganizationUserId`: %v\n", resp.OrganizationUserId)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.Cancel.OrganizationId`: %v\n", resp.OrganizationId)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.Cancel.OrganizationSuppliedUserId`: %v\n", resp.OrganizationSuppliedUserId)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.Cancel.RevokedAccess`: %v\n", resp.RevokedAccess)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.Cancel.LastSyncedAt`: %v\n", resp.LastSyncedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.Cancel.LastSyncAction`: %v\n", resp.LastSyncAction)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.Cancel.EnableAutoSync`: %v\n", resp.EnableAutoSync)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.Cancel.CreatedAt`: %v\n", resp.CreatedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.Cancel.UpdatedAt`: %v\n", resp.UpdatedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.Cancel.FilesSyncedAt`: %v\n", resp.FilesSyncedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.Cancel.DataSourceMetadata`: %v\n", resp.DataSourceMetadata)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectDataSource

Connect Data Source

### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    authentication := *carbon.NewOANDSCZGFB()
    syncOptions := *carbon.NewSyncOptions()
    
    connectDataSourceInput := *carbon.NewConnectDataSourceInput(
        authentication,
    )
    connectDataSourceInput.SetSyncOptions(syncOptions)
    
    request := client.IntegrationsApi.ConnectDataSource(
        connectDataSourceInput,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.ConnectDataSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `ConnectDataSource`: ConnectDataSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.ConnectDataSource`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `ConnectDataSourceResponse.ConnectDataSource.DataSource`: %v\n", resp.DataSource)
    fmt.Fprintf(os.Stdout, "Response from `ConnectDataSourceResponse.ConnectDataSource.SyncUrl`: %v\n", resp.SyncUrl)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectFreshdesk

Freshdesk Connect



### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    fileSyncConfig := *carbon.NewFileSyncConfigNullable()
    
    freshDeskConnectRequest := *carbon.NewFreshDeskConnectRequest(
        "null",
        "null",
    )
    freshDeskConnectRequest.SetTags({})
    freshDeskConnectRequest.SetChunkSize(1500)
    freshDeskConnectRequest.SetChunkOverlap(20)
    freshDeskConnectRequest.SetSkipEmbeddingGeneration(false)
    freshDeskConnectRequest.SetEmbeddingModel(null)
    freshDeskConnectRequest.SetGenerateSparseVectors(false)
    freshDeskConnectRequest.SetPrependFilenameToChunks(false)
    freshDeskConnectRequest.SetSyncFilesOnConnection(true)
    freshDeskConnectRequest.SetRequestId("null")
    freshDeskConnectRequest.SetSyncSourceItems(true)
    freshDeskConnectRequest.SetFileSyncConfig(fileSyncConfig)
    
    request := client.IntegrationsApi.ConnectFreshdesk(
        freshDeskConnectRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.ConnectFreshdesk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `ConnectFreshdesk`: GenericSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.ConnectFreshdesk`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `GenericSuccessResponse.ConnectFreshdesk.Success`: %v\n", resp.Success)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectGitbook

Gitbook Connect



### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    fileSyncConfig := *carbon.NewFileSyncConfigNullable()
    
    gitbookConnectRequest := *carbon.NewGitbookConnectRequest(
        "null",
        "null",
    )
    gitbookConnectRequest.SetTags({})
    gitbookConnectRequest.SetChunkSize(1500)
    gitbookConnectRequest.SetChunkOverlap(20)
    gitbookConnectRequest.SetSkipEmbeddingGeneration(false)
    gitbookConnectRequest.SetEmbeddingModel(null)
    gitbookConnectRequest.SetGenerateSparseVectors(false)
    gitbookConnectRequest.SetPrependFilenameToChunks(false)
    gitbookConnectRequest.SetSyncFilesOnConnection(true)
    gitbookConnectRequest.SetRequestId("null")
    gitbookConnectRequest.SetSyncSourceItems(true)
    gitbookConnectRequest.SetFileSyncConfig(fileSyncConfig)
    
    request := client.IntegrationsApi.ConnectGitbook(
        gitbookConnectRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.ConnectGitbook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `ConnectGitbook`: GenericSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.ConnectGitbook`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `GenericSuccessResponse.ConnectGitbook.Success`: %v\n", resp.Success)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectGuru

Guru Connect



### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    fileSyncConfig := *carbon.NewFileSyncConfigNullable()
    
    guruConnectRequest := *carbon.NewGuruConnectRequest(
        "null",
        "null",
    )
    guruConnectRequest.SetTags({})
    guruConnectRequest.SetChunkSize(1500)
    guruConnectRequest.SetChunkOverlap(20)
    guruConnectRequest.SetSkipEmbeddingGeneration(false)
    guruConnectRequest.SetEmbeddingModel(null)
    guruConnectRequest.SetGenerateSparseVectors(false)
    guruConnectRequest.SetPrependFilenameToChunks(false)
    guruConnectRequest.SetSyncFilesOnConnection(true)
    guruConnectRequest.SetRequestId("null")
    guruConnectRequest.SetSyncSourceItems(true)
    guruConnectRequest.SetFileSyncConfig(fileSyncConfig)
    
    request := client.IntegrationsApi.ConnectGuru(
        guruConnectRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.ConnectGuru``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `ConnectGuru`: GenericSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.ConnectGuru`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `GenericSuccessResponse.ConnectGuru.Success`: %v\n", resp.Success)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAwsIamUser

S3 Auth



### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    
    s3AuthRequest := *carbon.NewS3AuthRequest(
        "null",
        "null",
    )
    s3AuthRequest.SetSyncSourceItems(true)
    s3AuthRequest.SetEndpointUrl("null")
    
    request := client.IntegrationsApi.CreateAwsIamUser(
        s3AuthRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.CreateAwsIamUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `CreateAwsIamUser`: OrganizationUserDataSourceAPI
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.CreateAwsIamUser`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.CreateAwsIamUser.Id`: %v\n", resp.Id)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.CreateAwsIamUser.DataSourceExternalId`: %v\n", resp.DataSourceExternalId)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.CreateAwsIamUser.DataSourceType`: %v\n", resp.DataSourceType)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.CreateAwsIamUser.Token`: %v\n", resp.Token)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.CreateAwsIamUser.SyncStatus`: %v\n", resp.SyncStatus)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.CreateAwsIamUser.SourceItemsSyncedAt`: %v\n", resp.SourceItemsSyncedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.CreateAwsIamUser.OrganizationUserId`: %v\n", resp.OrganizationUserId)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.CreateAwsIamUser.OrganizationId`: %v\n", resp.OrganizationId)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.CreateAwsIamUser.OrganizationSuppliedUserId`: %v\n", resp.OrganizationSuppliedUserId)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.CreateAwsIamUser.RevokedAccess`: %v\n", resp.RevokedAccess)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.CreateAwsIamUser.LastSyncedAt`: %v\n", resp.LastSyncedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.CreateAwsIamUser.LastSyncAction`: %v\n", resp.LastSyncAction)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.CreateAwsIamUser.EnableAutoSync`: %v\n", resp.EnableAutoSync)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.CreateAwsIamUser.CreatedAt`: %v\n", resp.CreatedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.CreateAwsIamUser.UpdatedAt`: %v\n", resp.UpdatedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.CreateAwsIamUser.FilesSyncedAt`: %v\n", resp.FilesSyncedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.CreateAwsIamUser.DataSourceMetadata`: %v\n", resp.DataSourceMetadata)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOauthUrl

Get Oauth Url



### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    fileSyncConfig := *carbon.NewFileSyncConfigNullable()
    servicenowCredentials := *carbon.NewServiceNowCredentialsNullable()
    
    oAuthURLRequest := *carbon.NewOAuthURLRequest(
        null,
    )
    oAuthURLRequest.SetTags(null)
    oAuthURLRequest.SetScope("null")
    oAuthURLRequest.SetChunkSize(1500)
    oAuthURLRequest.SetChunkOverlap(20)
    oAuthURLRequest.SetSkipEmbeddingGeneration(false)
    oAuthURLRequest.SetEmbeddingModel(null)
    oAuthURLRequest.SetZendeskSubdomain("null")
    oAuthURLRequest.SetMicrosoftTenant("null")
    oAuthURLRequest.SetSharepointSiteName("null")
    oAuthURLRequest.SetConfluenceSubdomain("null")
    oAuthURLRequest.SetGenerateSparseVectors(false)
    oAuthURLRequest.SetPrependFilenameToChunks(false)
    oAuthURLRequest.SetMaxItemsPerChunk(null)
    oAuthURLRequest.SetSalesforceDomain("null")
    oAuthURLRequest.SetSyncFilesOnConnection(true)
    oAuthURLRequest.SetSetPageAsBoundary(false)
    oAuthURLRequest.SetDataSourceId(null)
    oAuthURLRequest.SetConnectingNewAccount(false)
    oAuthURLRequest.SetRequestId("null")
    oAuthURLRequest.SetUseOcr(false)
    oAuthURLRequest.SetParsePdfTablesWithOcr(false)
    oAuthURLRequest.SetEnableFilePicker(true)
    oAuthURLRequest.SetSyncSourceItems(true)
    oAuthURLRequest.SetIncrementalSync(false)
    oAuthURLRequest.SetFileSyncConfig(fileSyncConfig)
    oAuthURLRequest.SetAutomaticallyOpenFilePicker(null)
    oAuthURLRequest.SetGongAccountEmail("null")
    oAuthURLRequest.SetServicenowCredentials(servicenowCredentials)
    
    request := client.IntegrationsApi.GetOauthUrl(
        oAuthURLRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.GetOauthUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `GetOauthUrl`: OuthURLResponse
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.GetOauthUrl`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `OuthURLResponse.GetOauthUrl.OauthUrl`: %v\n", resp.OauthUrl)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConfluencePages
![Deprecated](https://img.shields.io/badge/deprecated-yellow)

Confluence List



### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    
    listRequest := *carbon.NewListRequest(
        null,
    )
    listRequest.SetParentId("null")
    
    request := client.IntegrationsApi.ListConfluencePages(
        listRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.ListConfluencePages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `ListConfluencePages`: ListResponse
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.ListConfluencePages`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `ListResponse.ListConfluencePages.Data`: %v\n", resp.Data)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConversations

Slack List Conversations



### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    request := client.IntegrationsApi.ListConversations(
    )
    request.Types(""public_channel"")
    request.Cursor("cursor_example")
    request.DataSourceId(56)
    request.ExcludeArchived(true)
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.ListConversations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `ListConversations`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.ListConversations`: %v\n", resp)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDataSourceItems

List Data Source Items

### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    filters := *carbon.NewListItemsFiltersNullable()
    pagination := *carbon.NewPagination()
    
    listDataSourceItemsRequest := *carbon.NewListDataSourceItemsRequest(
        null,
    )
    listDataSourceItemsRequest.SetParentId("null")
    listDataSourceItemsRequest.SetFilters(filters)
    listDataSourceItemsRequest.SetPagination(pagination)
    listDataSourceItemsRequest.SetOrderBy(null)
    listDataSourceItemsRequest.SetOrderDir(null)
    
    request := client.IntegrationsApi.ListDataSourceItems(
        listDataSourceItemsRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.ListDataSourceItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `ListDataSourceItems`: ListDataSourceItemsResponse
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.ListDataSourceItems`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `ListDataSourceItemsResponse.ListDataSourceItems.Items`: %v\n", resp.Items)
    fmt.Fprintf(os.Stdout, "Response from `ListDataSourceItemsResponse.ListDataSourceItems.Count`: %v\n", resp.Count)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFolders

Outlook Folders



### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    request := client.IntegrationsApi.ListFolders(
    )
    request.DataSourceId(56)
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.ListFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `ListFolders`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.ListFolders`: %v\n", resp)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGitbookSpaces

Gitbook Spaces



### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    request := client.IntegrationsApi.ListGitbookSpaces(
        56,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.ListGitbookSpaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `ListGitbookSpaces`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.ListGitbookSpaces`: %v\n", resp)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLabels

Gmail Labels



### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    request := client.IntegrationsApi.ListLabels(
    )
    request.DataSourceId(56)
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.ListLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `ListLabels`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.ListLabels`: %v\n", resp)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOutlookCategories

Outlook Categories



### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    request := client.IntegrationsApi.ListOutlookCategories(
    )
    request.DataSourceId(56)
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.ListOutlookCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `ListOutlookCategories`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.ListOutlookCategories`: %v\n", resp)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRepos

Github List Repos



### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    request := client.IntegrationsApi.ListRepos(
    )
    request.PerPage(30)
    request.Page(1)
    request.DataSourceId(56)
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.ListRepos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `ListRepos`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.ListRepos`: %v\n", resp)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncAzureBlobFiles

Azure Blob Files



### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    fileSyncConfig := *carbon.NewFileSyncConfigNullable()
    
    azureBlobFileSyncInput := *carbon.NewAzureBlobFileSyncInput(
        null,
    )
    azureBlobFileSyncInput.SetTags({})
    azureBlobFileSyncInput.SetChunkSize(1500)
    azureBlobFileSyncInput.SetChunkOverlap(20)
    azureBlobFileSyncInput.SetSkipEmbeddingGeneration(false)
    azureBlobFileSyncInput.SetEmbeddingModel(null)
    azureBlobFileSyncInput.SetGenerateSparseVectors(false)
    azureBlobFileSyncInput.SetPrependFilenameToChunks(false)
    azureBlobFileSyncInput.SetMaxItemsPerChunk(null)
    azureBlobFileSyncInput.SetSetPageAsBoundary(false)
    azureBlobFileSyncInput.SetDataSourceId(null)
    azureBlobFileSyncInput.SetRequestId("null")
    azureBlobFileSyncInput.SetUseOcr(false)
    azureBlobFileSyncInput.SetParsePdfTablesWithOcr(false)
    azureBlobFileSyncInput.SetFileSyncConfig(fileSyncConfig)
    
    request := client.IntegrationsApi.SyncAzureBlobFiles(
        azureBlobFileSyncInput,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.SyncAzureBlobFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `SyncAzureBlobFiles`: GenericSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.SyncAzureBlobFiles`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `GenericSuccessResponse.SyncAzureBlobFiles.Success`: %v\n", resp.Success)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncAzureBlobStorage

Azure Blob Storage Auth



### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    
    azureBlobAuthRequest := *carbon.NewAzureBlobAuthRequest(
        "null",
        "null",
    )
    azureBlobAuthRequest.SetSyncSourceItems(true)
    
    request := client.IntegrationsApi.SyncAzureBlobStorage(
        azureBlobAuthRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.SyncAzureBlobStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `SyncAzureBlobStorage`: OrganizationUserDataSourceAPI
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.SyncAzureBlobStorage`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncAzureBlobStorage.Id`: %v\n", resp.Id)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncAzureBlobStorage.DataSourceExternalId`: %v\n", resp.DataSourceExternalId)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncAzureBlobStorage.DataSourceType`: %v\n", resp.DataSourceType)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncAzureBlobStorage.Token`: %v\n", resp.Token)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncAzureBlobStorage.SyncStatus`: %v\n", resp.SyncStatus)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncAzureBlobStorage.SourceItemsSyncedAt`: %v\n", resp.SourceItemsSyncedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncAzureBlobStorage.OrganizationUserId`: %v\n", resp.OrganizationUserId)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncAzureBlobStorage.OrganizationId`: %v\n", resp.OrganizationId)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncAzureBlobStorage.OrganizationSuppliedUserId`: %v\n", resp.OrganizationSuppliedUserId)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncAzureBlobStorage.RevokedAccess`: %v\n", resp.RevokedAccess)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncAzureBlobStorage.LastSyncedAt`: %v\n", resp.LastSyncedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncAzureBlobStorage.LastSyncAction`: %v\n", resp.LastSyncAction)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncAzureBlobStorage.EnableAutoSync`: %v\n", resp.EnableAutoSync)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncAzureBlobStorage.CreatedAt`: %v\n", resp.CreatedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncAzureBlobStorage.UpdatedAt`: %v\n", resp.UpdatedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncAzureBlobStorage.FilesSyncedAt`: %v\n", resp.FilesSyncedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncAzureBlobStorage.DataSourceMetadata`: %v\n", resp.DataSourceMetadata)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncConfluence
![Deprecated](https://img.shields.io/badge/deprecated-yellow)

Confluence Sync



### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    fileSyncConfig := *carbon.NewFileSyncConfigNullable()
    
    syncFilesRequest := *carbon.NewSyncFilesRequest(
        null,
        null,
    )
    syncFilesRequest.SetTags({})
    syncFilesRequest.SetChunkSize(1500)
    syncFilesRequest.SetChunkOverlap(20)
    syncFilesRequest.SetSkipEmbeddingGeneration(false)
    syncFilesRequest.SetEmbeddingModel(null)
    syncFilesRequest.SetGenerateSparseVectors(false)
    syncFilesRequest.SetPrependFilenameToChunks(false)
    syncFilesRequest.SetMaxItemsPerChunk(null)
    syncFilesRequest.SetSetPageAsBoundary(false)
    syncFilesRequest.SetRequestId("null")
    syncFilesRequest.SetUseOcr(false)
    syncFilesRequest.SetParsePdfTablesWithOcr(false)
    syncFilesRequest.SetIncrementalSync(false)
    syncFilesRequest.SetFileSyncConfig(fileSyncConfig)
    
    request := client.IntegrationsApi.SyncConfluence(
        syncFilesRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.SyncConfluence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `SyncConfluence`: GenericSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.SyncConfluence`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `GenericSuccessResponse.SyncConfluence.Success`: %v\n", resp.Success)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncDataSourceItems

Sync Data Source Items

### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    
    syncDirectoryRequest := *carbon.NewSyncDirectoryRequest(
        null,
    )
    
    request := client.IntegrationsApi.SyncDataSourceItems(
        syncDirectoryRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.SyncDataSourceItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `SyncDataSourceItems`: OrganizationUserDataSourceAPI
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.SyncDataSourceItems`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncDataSourceItems.Id`: %v\n", resp.Id)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncDataSourceItems.DataSourceExternalId`: %v\n", resp.DataSourceExternalId)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncDataSourceItems.DataSourceType`: %v\n", resp.DataSourceType)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncDataSourceItems.Token`: %v\n", resp.Token)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncDataSourceItems.SyncStatus`: %v\n", resp.SyncStatus)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncDataSourceItems.SourceItemsSyncedAt`: %v\n", resp.SourceItemsSyncedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncDataSourceItems.OrganizationUserId`: %v\n", resp.OrganizationUserId)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncDataSourceItems.OrganizationId`: %v\n", resp.OrganizationId)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncDataSourceItems.OrganizationSuppliedUserId`: %v\n", resp.OrganizationSuppliedUserId)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncDataSourceItems.RevokedAccess`: %v\n", resp.RevokedAccess)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncDataSourceItems.LastSyncedAt`: %v\n", resp.LastSyncedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncDataSourceItems.LastSyncAction`: %v\n", resp.LastSyncAction)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncDataSourceItems.EnableAutoSync`: %v\n", resp.EnableAutoSync)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncDataSourceItems.CreatedAt`: %v\n", resp.CreatedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncDataSourceItems.UpdatedAt`: %v\n", resp.UpdatedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncDataSourceItems.FilesSyncedAt`: %v\n", resp.FilesSyncedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.SyncDataSourceItems.DataSourceMetadata`: %v\n", resp.DataSourceMetadata)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncFiles

Sync Files



### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    fileSyncConfig := *carbon.NewFileSyncConfigNullable()
    
    syncFilesRequest := *carbon.NewSyncFilesRequest(
        null,
        null,
    )
    syncFilesRequest.SetTags({})
    syncFilesRequest.SetChunkSize(1500)
    syncFilesRequest.SetChunkOverlap(20)
    syncFilesRequest.SetSkipEmbeddingGeneration(false)
    syncFilesRequest.SetEmbeddingModel(null)
    syncFilesRequest.SetGenerateSparseVectors(false)
    syncFilesRequest.SetPrependFilenameToChunks(false)
    syncFilesRequest.SetMaxItemsPerChunk(null)
    syncFilesRequest.SetSetPageAsBoundary(false)
    syncFilesRequest.SetRequestId("null")
    syncFilesRequest.SetUseOcr(false)
    syncFilesRequest.SetParsePdfTablesWithOcr(false)
    syncFilesRequest.SetIncrementalSync(false)
    syncFilesRequest.SetFileSyncConfig(fileSyncConfig)
    
    request := client.IntegrationsApi.SyncFiles(
        syncFilesRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.SyncFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `SyncFiles`: GenericSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.SyncFiles`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `GenericSuccessResponse.SyncFiles.Success`: %v\n", resp.Success)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncGitHub

Github Connect



### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    
    githubConnectRequest := *carbon.NewGithubConnectRequest(
        "null",
        "null",
    )
    githubConnectRequest.SetSyncSourceItems(false)
    
    request := client.IntegrationsApi.SyncGitHub(
        githubConnectRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.SyncGitHub``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `SyncGitHub`: GenericSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.SyncGitHub`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `GenericSuccessResponse.SyncGitHub.Success`: %v\n", resp.Success)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncGitbook

Gitbook Sync



### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    fileSyncConfig := *carbon.NewFileSyncConfigNullable()
    
    gitbookSyncRequest := *carbon.NewGitbookSyncRequest(
        null,
        null,
    )
    gitbookSyncRequest.SetTags({})
    gitbookSyncRequest.SetChunkSize(1500)
    gitbookSyncRequest.SetChunkOverlap(20)
    gitbookSyncRequest.SetSkipEmbeddingGeneration(false)
    gitbookSyncRequest.SetEmbeddingModel(null)
    gitbookSyncRequest.SetGenerateSparseVectors(false)
    gitbookSyncRequest.SetPrependFilenameToChunks(false)
    gitbookSyncRequest.SetRequestId("null")
    gitbookSyncRequest.SetFileSyncConfig(fileSyncConfig)
    
    request := client.IntegrationsApi.SyncGitbook(
        gitbookSyncRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.SyncGitbook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `SyncGitbook`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.SyncGitbook`: %v\n", resp)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncGmail

Gmail Sync



### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    fileSyncConfig := *carbon.NewFileSyncConfigNullable()
    
    gmailSyncInput := *carbon.NewGmailSyncInput(
        null,
    )
    gmailSyncInput.SetTags({})
    gmailSyncInput.SetChunkSize(1500)
    gmailSyncInput.SetChunkOverlap(20)
    gmailSyncInput.SetSkipEmbeddingGeneration(false)
    gmailSyncInput.SetEmbeddingModel(null)
    gmailSyncInput.SetGenerateSparseVectors(false)
    gmailSyncInput.SetPrependFilenameToChunks(false)
    gmailSyncInput.SetDataSourceId(null)
    gmailSyncInput.SetRequestId("null")
    gmailSyncInput.SetSyncAttachments(false)
    gmailSyncInput.SetFileSyncConfig(fileSyncConfig)
    gmailSyncInput.SetIncrementalSync(false)
    
    request := client.IntegrationsApi.SyncGmail(
        gmailSyncInput,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.SyncGmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `SyncGmail`: GenericSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.SyncGmail`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `GenericSuccessResponse.SyncGmail.Success`: %v\n", resp.Success)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncOutlook

Outlook Sync



### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    fileSyncConfig := *carbon.NewFileSyncConfigNullable()
    
    outlookSyncInput := *carbon.NewOutlookSyncInput(
        null,
    )
    outlookSyncInput.SetTags({})
    outlookSyncInput.SetFolder("Inbox")
    outlookSyncInput.SetChunkSize(1500)
    outlookSyncInput.SetChunkOverlap(20)
    outlookSyncInput.SetSkipEmbeddingGeneration(false)
    outlookSyncInput.SetEmbeddingModel(null)
    outlookSyncInput.SetGenerateSparseVectors(false)
    outlookSyncInput.SetPrependFilenameToChunks(false)
    outlookSyncInput.SetDataSourceId(null)
    outlookSyncInput.SetRequestId("null")
    outlookSyncInput.SetSyncAttachments(false)
    outlookSyncInput.SetFileSyncConfig(fileSyncConfig)
    outlookSyncInput.SetIncrementalSync(false)
    
    request := client.IntegrationsApi.SyncOutlook(
        outlookSyncInput,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.SyncOutlook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `SyncOutlook`: GenericSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.SyncOutlook`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `GenericSuccessResponse.SyncOutlook.Success`: %v\n", resp.Success)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncRepos

Github Sync Repos



### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    
    githubFetchReposRequest := *carbon.NewGithubFetchReposRequest(
        null,
    )
    githubFetchReposRequest.SetDataSourceId(null)
    
    request := client.IntegrationsApi.SyncRepos(
        githubFetchReposRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.SyncRepos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `SyncRepos`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.SyncRepos`: %v\n", resp)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncRssFeed

Rss Feed

### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    
    rSSFeedInput := *carbon.NewRSSFeedInput(
        "null",
    )
    rSSFeedInput.SetTags({})
    rSSFeedInput.SetChunkSize(1500)
    rSSFeedInput.SetChunkOverlap(20)
    rSSFeedInput.SetSkipEmbeddingGeneration(false)
    rSSFeedInput.SetEmbeddingModel(null)
    rSSFeedInput.SetGenerateSparseVectors(false)
    rSSFeedInput.SetPrependFilenameToChunks(false)
    rSSFeedInput.SetRequestId("null")
    
    request := client.IntegrationsApi.SyncRssFeed(
        rSSFeedInput,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.SyncRssFeed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `SyncRssFeed`: GenericSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.SyncRssFeed`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `GenericSuccessResponse.SyncRssFeed.Success`: %v\n", resp.Success)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncS3Files

S3 Files



### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    fileSyncConfig := *carbon.NewFileSyncConfigNullable()
    
    s3FileSyncInput := *carbon.NewS3FileSyncInput(
        null,
    )
    s3FileSyncInput.SetTags({})
    s3FileSyncInput.SetChunkSize(1500)
    s3FileSyncInput.SetChunkOverlap(20)
    s3FileSyncInput.SetSkipEmbeddingGeneration(false)
    s3FileSyncInput.SetEmbeddingModel(null)
    s3FileSyncInput.SetGenerateSparseVectors(false)
    s3FileSyncInput.SetPrependFilenameToChunks(false)
    s3FileSyncInput.SetMaxItemsPerChunk(null)
    s3FileSyncInput.SetSetPageAsBoundary(false)
    s3FileSyncInput.SetDataSourceId(null)
    s3FileSyncInput.SetRequestId("null")
    s3FileSyncInput.SetUseOcr(false)
    s3FileSyncInput.SetParsePdfTablesWithOcr(false)
    s3FileSyncInput.SetFileSyncConfig(fileSyncConfig)
    
    request := client.IntegrationsApi.SyncS3Files(
        s3FileSyncInput,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.SyncS3Files``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `SyncS3Files`: GenericSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.SyncS3Files`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `GenericSuccessResponse.SyncS3Files.Success`: %v\n", resp.Success)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncSlack

Slack Sync



### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    filters := *carbon.NewSlackFilters()
    
    slackSyncRequest := *carbon.NewSlackSyncRequest(
        filters,
    )
    slackSyncRequest.SetTags({})
    slackSyncRequest.SetChunkSize(1500)
    slackSyncRequest.SetChunkOverlap(20)
    slackSyncRequest.SetSkipEmbeddingGeneration(false)
    slackSyncRequest.SetEmbeddingModel(null)
    slackSyncRequest.SetGenerateSparseVectors(false)
    slackSyncRequest.SetPrependFilenameToChunks(false)
    slackSyncRequest.SetDataSourceId(null)
    slackSyncRequest.SetRequestId("null")
    
    request := client.IntegrationsApi.SyncSlack(
        slackSyncRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.SyncSlack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `SyncSlack`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.SyncSlack`: %v\n", resp)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


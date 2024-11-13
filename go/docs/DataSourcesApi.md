# APIClient.DataSourcesApi

All URIs are relative to *https://api.carbon.ai*

Method | Path | Description
------------- | ------------- | -------------
[**AddTags**](DataSourcesApi.md#AddTags) | **Post** /data_sources/tags/add | Add Data Source Tags
[**Query**](DataSourcesApi.md#Query) | **Post** /data_sources | Data Sources
[**QueryUserDataSources**](DataSourcesApi.md#QueryUserDataSources) | **Post** /user_data_sources | User Data Sources
[**RemoveTags**](DataSourcesApi.md#RemoveTags) | **Post** /data_sources/tags/remove | Remove Data Source Tags
[**RevokeAccessToken**](DataSourcesApi.md#RevokeAccessToken) | **Post** /revoke_access_token | Revoke Access Token



## AddTags

Add Data Source Tags

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

    
    addDataSourceTagsInput := *carbon.NewAddDataSourceTagsInput(
        null,
        null,
    )
    
    request := client.DataSourcesApi.AddTags(
        addDataSourceTagsInput,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesApi.AddTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `AddTags`: OrganizationUserDataSourceAPI
    fmt.Fprintf(os.Stdout, "Response from `DataSourcesApi.AddTags`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.AddTags.Tags`: %v\n", resp.Tags)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.AddTags.Id`: %v\n", resp.Id)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.AddTags.DataSourceExternalId`: %v\n", resp.DataSourceExternalId)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.AddTags.DataSourceType`: %v\n", resp.DataSourceType)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.AddTags.Token`: %v\n", resp.Token)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.AddTags.SyncStatus`: %v\n", resp.SyncStatus)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.AddTags.SourceItemsSyncedAt`: %v\n", resp.SourceItemsSyncedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.AddTags.OrganizationUserId`: %v\n", resp.OrganizationUserId)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.AddTags.OrganizationId`: %v\n", resp.OrganizationId)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.AddTags.OrganizationSuppliedUserId`: %v\n", resp.OrganizationSuppliedUserId)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.AddTags.RevokedAccess`: %v\n", resp.RevokedAccess)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.AddTags.LastSyncedAt`: %v\n", resp.LastSyncedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.AddTags.LastSyncAction`: %v\n", resp.LastSyncAction)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.AddTags.EnableAutoSync`: %v\n", resp.EnableAutoSync)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.AddTags.CreatedAt`: %v\n", resp.CreatedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.AddTags.UpdatedAt`: %v\n", resp.UpdatedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.AddTags.FilesSyncedAt`: %v\n", resp.FilesSyncedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.AddTags.DataSourceMetadata`: %v\n", resp.DataSourceMetadata)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Query

Data Sources

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
    configuration.SetApiKey("AUTHORIZATION")
    client := carbon.NewAPIClient(configuration)

    pagination := *carbon.NewPagination()
    filters := *carbon.NewOrganizationUserDataSourceFilters()
    
    organizationUserDataSourceQueryInput := *carbon.NewOrganizationUserDataSourceQueryInput()
    organizationUserDataSourceQueryInput.SetPagination(pagination)
    organizationUserDataSourceQueryInput.SetOrderBy(null)
    organizationUserDataSourceQueryInput.SetOrderDir(null)
    organizationUserDataSourceQueryInput.SetFilters(filters)
    
    request := client.DataSourcesApi.Query(
        organizationUserDataSourceQueryInput,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesApi.Query``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `Query`: OrganizationUserDataSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `DataSourcesApi.Query`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceResponse.Query.Results`: %v\n", resp.Results)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceResponse.Query.Count`: %v\n", resp.Count)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryUserDataSources

User Data Sources

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

    pagination := *carbon.NewPagination()
    filters := *carbon.NewOrganizationUserDataSourceFilters()
    
    organizationUserDataSourceQueryInput := *carbon.NewOrganizationUserDataSourceQueryInput()
    organizationUserDataSourceQueryInput.SetPagination(pagination)
    organizationUserDataSourceQueryInput.SetOrderBy(null)
    organizationUserDataSourceQueryInput.SetOrderDir(null)
    organizationUserDataSourceQueryInput.SetFilters(filters)
    
    request := client.DataSourcesApi.QueryUserDataSources(
        organizationUserDataSourceQueryInput,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesApi.QueryUserDataSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `QueryUserDataSources`: OrganizationUserDataSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `DataSourcesApi.QueryUserDataSources`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceResponse.QueryUserDataSources.Results`: %v\n", resp.Results)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceResponse.QueryUserDataSources.Count`: %v\n", resp.Count)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTags

Remove Data Source Tags

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

    
    removeDataSourceTagsInput := *carbon.NewRemoveDataSourceTagsInput(
        null,
    )
    removeDataSourceTagsInput.SetTagsToRemove([])
    removeDataSourceTagsInput.SetRemoveAllTags(false)
    
    request := client.DataSourcesApi.RemoveTags(
        removeDataSourceTagsInput,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesApi.RemoveTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `RemoveTags`: OrganizationUserDataSourceAPI
    fmt.Fprintf(os.Stdout, "Response from `DataSourcesApi.RemoveTags`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.RemoveTags.Tags`: %v\n", resp.Tags)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.RemoveTags.Id`: %v\n", resp.Id)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.RemoveTags.DataSourceExternalId`: %v\n", resp.DataSourceExternalId)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.RemoveTags.DataSourceType`: %v\n", resp.DataSourceType)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.RemoveTags.Token`: %v\n", resp.Token)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.RemoveTags.SyncStatus`: %v\n", resp.SyncStatus)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.RemoveTags.SourceItemsSyncedAt`: %v\n", resp.SourceItemsSyncedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.RemoveTags.OrganizationUserId`: %v\n", resp.OrganizationUserId)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.RemoveTags.OrganizationId`: %v\n", resp.OrganizationId)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.RemoveTags.OrganizationSuppliedUserId`: %v\n", resp.OrganizationSuppliedUserId)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.RemoveTags.RevokedAccess`: %v\n", resp.RevokedAccess)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.RemoveTags.LastSyncedAt`: %v\n", resp.LastSyncedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.RemoveTags.LastSyncAction`: %v\n", resp.LastSyncAction)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.RemoveTags.EnableAutoSync`: %v\n", resp.EnableAutoSync)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.RemoveTags.CreatedAt`: %v\n", resp.CreatedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.RemoveTags.UpdatedAt`: %v\n", resp.UpdatedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.RemoveTags.FilesSyncedAt`: %v\n", resp.FilesSyncedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationUserDataSourceAPI.RemoveTags.DataSourceMetadata`: %v\n", resp.DataSourceMetadata)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeAccessToken

Revoke Access Token

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

    
    revokeAccessTokenInput := *carbon.NewRevokeAccessTokenInput(
        null,
    )
    
    request := client.DataSourcesApi.RevokeAccessToken(
        revokeAccessTokenInput,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesApi.RevokeAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `RevokeAccessToken`: GenericSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `DataSourcesApi.RevokeAccessToken`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `GenericSuccessResponse.RevokeAccessToken.Success`: %v\n", resp.Success)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


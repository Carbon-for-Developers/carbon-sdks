# APIClient.UsersApi

All URIs are relative to *https://api.carbon.ai*

Method | Path | Description
------------- | ------------- | -------------
[**Delete**](UsersApi.md#Delete) | **Post** /delete_users | Delete Users
[**Get**](UsersApi.md#Get) | **Post** /user | User Endpoint
[**List**](UsersApi.md#List) | **Post** /list_users | List Users Endpoint
[**ToggleUserFeatures**](UsersApi.md#ToggleUserFeatures) | **Post** /modify_user_configuration | Toggle User Features
[**UpdateUsers**](UsersApi.md#UpdateUsers) | **Post** /update_users | Update Users



## Delete

Delete Users

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

    
    deleteUsersInput := *carbon.NewDeleteUsersInput(
        null,
    )
    
    request := client.UsersApi.Delete(
        deleteUsersInput,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `Delete`: GenericSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.Delete`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `GenericSuccessResponse.Delete.Success`: %v\n", resp.Success)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

User Endpoint

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

    
    userRequestContent := *carbon.NewUserRequestContent(
        "null",
    )
    
    request := client.UsersApi.Get(
        userRequestContent,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `Get`: UserResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.Get`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `UserResponse.Get.Id`: %v\n", resp.Id)
    fmt.Fprintf(os.Stdout, "Response from `UserResponse.Get.OrganizationId`: %v\n", resp.OrganizationId)
    fmt.Fprintf(os.Stdout, "Response from `UserResponse.Get.OrganizationSuppliedUserId`: %v\n", resp.OrganizationSuppliedUserId)
    fmt.Fprintf(os.Stdout, "Response from `UserResponse.Get.CreatedAt`: %v\n", resp.CreatedAt)
    fmt.Fprintf(os.Stdout, "Response from `UserResponse.Get.UpdatedAt`: %v\n", resp.UpdatedAt)
    fmt.Fprintf(os.Stdout, "Response from `UserResponse.Get.DeletedAt`: %v\n", resp.DeletedAt)
    fmt.Fprintf(os.Stdout, "Response from `UserResponse.Get.NumFilesSynced`: %v\n", resp.NumFilesSynced)
    fmt.Fprintf(os.Stdout, "Response from `UserResponse.Get.NumCharactersSynced`: %v\n", resp.NumCharactersSynced)
    fmt.Fprintf(os.Stdout, "Response from `UserResponse.Get.NumTokensSynced`: %v\n", resp.NumTokensSynced)
    fmt.Fprintf(os.Stdout, "Response from `UserResponse.Get.AggregateFileSize`: %v\n", resp.AggregateFileSize)
    fmt.Fprintf(os.Stdout, "Response from `UserResponse.Get.AggregateNumCharacters`: %v\n", resp.AggregateNumCharacters)
    fmt.Fprintf(os.Stdout, "Response from `UserResponse.Get.AggregateNumTokens`: %v\n", resp.AggregateNumTokens)
    fmt.Fprintf(os.Stdout, "Response from `UserResponse.Get.AggregateNumEmbeddings`: %v\n", resp.AggregateNumEmbeddings)
    fmt.Fprintf(os.Stdout, "Response from `UserResponse.Get.AggregateNumFilesBySource`: %v\n", resp.AggregateNumFilesBySource)
    fmt.Fprintf(os.Stdout, "Response from `UserResponse.Get.AggregateNumFilesByFileFormat`: %v\n", resp.AggregateNumFilesByFileFormat)
    fmt.Fprintf(os.Stdout, "Response from `UserResponse.Get.UniqueFileTags`: %v\n", resp.UniqueFileTags)
    fmt.Fprintf(os.Stdout, "Response from `UserResponse.Get.EnabledFeatures`: %v\n", resp.EnabledFeatures)
    fmt.Fprintf(os.Stdout, "Response from `UserResponse.Get.CustomLimits`: %v\n", resp.CustomLimits)
    fmt.Fprintf(os.Stdout, "Response from `UserResponse.Get.AutoSyncEnabledSources`: %v\n", resp.AutoSyncEnabledSources)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

List Users Endpoint



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
    filters := *carbon.NewListUsersFilters()
    
    listUsersRequest := *carbon.NewListUsersRequest()
    listUsersRequest.SetPagination(pagination)
    listUsersRequest.SetFilters(filters)
    listUsersRequest.SetOrderBy(null)
    listUsersRequest.SetOrderDir(null)
    listUsersRequest.SetIncludeCount(false)
    
    request := client.UsersApi.List(
        listUsersRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `List`: UserListResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.List`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `UserListResponse.List.Users`: %v\n", resp.Users)
    fmt.Fprintf(os.Stdout, "Response from `UserListResponse.List.Count`: %v\n", resp.Count)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleUserFeatures
![Deprecated](https://img.shields.io/badge/deprecated-yellow)

Toggle User Features

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

    
    modifyUserConfigurationInput := *carbon.NewModifyUserConfigurationInput(
        "null",
        null,
    )
    
    request := client.UsersApi.ToggleUserFeatures(
        modifyUserConfigurationInput,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ToggleUserFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `ToggleUserFeatures`: GenericSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ToggleUserFeatures`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `GenericSuccessResponse.ToggleUserFeatures.Success`: %v\n", resp.Success)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUsers

Update Users

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

    autoSyncEnabledSources := *carbon.NewAutoSyncEnabledSourcesProperty()
    
    updateUsersInput := *carbon.NewUpdateUsersInput(
        null,
    )
    updateUsersInput.SetAutoSyncEnabledSources(autoSyncEnabledSources)
    updateUsersInput.SetMaxFiles(null)
    updateUsersInput.SetMaxFilesPerUpload(null)
    
    request := client.UsersApi.UpdateUsers(
        updateUsersInput,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UpdateUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `UpdateUsers`: GenericSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UpdateUsers`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `GenericSuccessResponse.UpdateUsers.Success`: %v\n", resp.Success)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# APIClient.OrganizationsApi

All URIs are relative to *https://api.carbon.ai*

Method | Path | Description
------------- | ------------- | -------------
[**Get**](OrganizationsApi.md#Get) | **Get** /organization | Get Organization
[**Update**](OrganizationsApi.md#Update) | **Post** /organization/update | Update Organization
[**UpdateStats**](OrganizationsApi.md#UpdateStats) | **Post** /organization/statistics | Update Organization Statistics



## Get

Get Organization

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

    request := client.OrganizationsApi.Get(
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `Get`: OrganizationResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.Get`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationResponse.Get.Id`: %v\n", resp.Id)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationResponse.Get.Name`: %v\n", resp.Name)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationResponse.Get.Nickname`: %v\n", *resp.Nickname)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationResponse.Get.RemoveBranding`: %v\n", resp.RemoveBranding)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationResponse.Get.CustomBranding`: %v\n", *resp.CustomBranding)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationResponse.Get.CustomLimits`: %v\n", *resp.CustomLimits)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationResponse.Get.AggregateFileSize`: %v\n", resp.AggregateFileSize)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationResponse.Get.AggregateNumCharacters`: %v\n", resp.AggregateNumCharacters)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationResponse.Get.AggregateNumTokens`: %v\n", resp.AggregateNumTokens)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationResponse.Get.AggregateNumEmbeddings`: %v\n", resp.AggregateNumEmbeddings)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationResponse.Get.AggregateNumFilesBySource`: %v\n", resp.AggregateNumFilesBySource)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationResponse.Get.AggregateNumFilesByFileFormat`: %v\n", resp.AggregateNumFilesByFileFormat)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationResponse.Get.FileStatisticsAggregatedAt`: %v\n", resp.FileStatisticsAggregatedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationResponse.Get.PeriodEndsAt`: %v\n", *resp.PeriodEndsAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationResponse.Get.CancelAtPeriodEnd`: %v\n", *resp.CancelAtPeriodEnd)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationResponse.Get.ConnectorSettings`: %v\n", *resp.ConnectorSettings)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationResponse.Get.GlobalUserConfig`: %v\n", *resp.GlobalUserConfig)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationResponse.Get.FileSyncUsage`: %v\n", *resp.FileSyncUsage)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationResponse.Get.LoggingSettings`: %v\n", *resp.LoggingSettings)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationResponse.Get.CreatedAt`: %v\n", resp.CreatedAt)
    fmt.Fprintf(os.Stdout, "Response from `OrganizationResponse.Get.UpdatedAt`: %v\n", resp.UpdatedAt)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

Update Organization

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

    globalUserConfig := *carbon.NewUserConfigurationNullable()
    
    updateOrganizationInput := *carbon.NewUpdateOrganizationInput()
    updateOrganizationInput.SetGlobalUserConfig(globalUserConfig)
    updateOrganizationInput.SetDataSourceConfigs(null)
    
    request := client.OrganizationsApi.Update(
        updateOrganizationInput,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `Update`: GenericSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.Update`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `GenericSuccessResponse.Update.Success`: %v\n", resp.Success)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStats

Update Organization Statistics



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

    request := client.OrganizationsApi.UpdateStats(
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `UpdateStats`: GenericSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateStats`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `GenericSuccessResponse.UpdateStats.Success`: %v\n", resp.Success)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


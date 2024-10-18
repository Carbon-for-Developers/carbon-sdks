# APIClient.WhiteLabelApi

All URIs are relative to *https://api.carbon.ai*

Method | Path | Description
------------- | ------------- | -------------
[**Create**](WhiteLabelApi.md#Create) | **Post** /white_label/create | Create White Labels
[**Delete**](WhiteLabelApi.md#Delete) | **Post** /white_label/delete | Delete White Labels
[**List**](WhiteLabelApi.md#List) | **Post** /white_label/list | List White Labels
[**Update**](WhiteLabelApi.md#Update) | **Post** /white_label/update | Update White Label



## Create

Create White Labels

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

    
    gDOSWhiteLabelInput := *carbon.New[]GDOSWhiteLabelInput()
    
    request := client.WhiteLabelApi.Create(
        ,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhiteLabelApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `Create`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WhiteLabelApi.Create`: %v\n", resp)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

Delete White Labels

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

    
    deleteWhiteLabelRequest := *carbon.NewDeleteWhiteLabelRequest(
        null,
    )
    
    request := client.WhiteLabelApi.Delete(
        deleteWhiteLabelRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhiteLabelApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `Delete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WhiteLabelApi.Delete`: %v\n", resp)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

List White Labels

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
    filters := *carbon.NewWhiteLabelFilters()
    
    listWhiteLabelRequest := *carbon.NewListWhiteLabelRequest()
    listWhiteLabelRequest.SetPagination(pagination)
    listWhiteLabelRequest.SetOrderBy(null)
    listWhiteLabelRequest.SetOrderDir(null)
    listWhiteLabelRequest.SetFilters(filters)
    
    request := client.WhiteLabelApi.List(
        listWhiteLabelRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhiteLabelApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `List`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WhiteLabelApi.List`: %v\n", resp)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

Update White Label

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

    
    body := *carbon.NewGDOSWhiteLabelInput1()
    
    request := client.WhiteLabelApi.Update(
        ,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhiteLabelApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `Update`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WhiteLabelApi.Update`: %v\n", resp)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


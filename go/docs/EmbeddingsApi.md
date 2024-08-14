# APIClient.EmbeddingsApi

All URIs are relative to *https://api.carbon.ai*

Method | Path | Description
------------- | ------------- | -------------
[**GetDocuments**](EmbeddingsApi.md#GetDocuments) | **Post** /embeddings | Embeddings
[**GetEmbeddingsAndChunks**](EmbeddingsApi.md#GetEmbeddingsAndChunks) | **Post** /text_chunks | Retrieve Embeddings And Content
[**List**](EmbeddingsApi.md#List) | **Post** /list_chunks_and_embeddings | Retrieve Embeddings And Content V2
[**UploadChunksAndEmbeddings**](EmbeddingsApi.md#UploadChunksAndEmbeddings) | **Post** /upload_chunks_and_embeddings | Upload Chunks And Embeddings



## GetDocuments

Embeddings



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

    hybridSearchTuningParameters := *carbon.NewHybridSearchTuningParamsNullable()
    rerank := *carbon.NewRerankParamsNullable()
    
    getEmbeddingDocumentsBody := *carbon.NewGetEmbeddingDocumentsBody(
        "null",
        null,
    )
    getEmbeddingDocumentsBody.SetTags(null)
    getEmbeddingDocumentsBody.SetQueryVector(null)
    getEmbeddingDocumentsBody.SetFileIds(null)
    getEmbeddingDocumentsBody.SetParentFileIds(null)
    getEmbeddingDocumentsBody.SetIncludeAllChildren(false)
    getEmbeddingDocumentsBody.SetTagsV2({"OR":[{"key":"subject","negate":false,"value":"holy-bible"},{"key":"person-of-interest","negate":false,"value":"jesus christ"},{"key":"genre","negate":true,"value":"fiction"},{"AND":[{"key":"subject","negate":true,"value":"tao-te-ching"},{"key":"author","negate":false,"value":"lao-tzu"}]}]})
    getEmbeddingDocumentsBody.SetIncludeTags(null)
    getEmbeddingDocumentsBody.SetIncludeVectors(null)
    getEmbeddingDocumentsBody.SetIncludeRawFile(null)
    getEmbeddingDocumentsBody.SetHybridSearch(null)
    getEmbeddingDocumentsBody.SetHybridSearchTuningParameters(hybridSearchTuningParameters)
    getEmbeddingDocumentsBody.SetMediaType(null)
    getEmbeddingDocumentsBody.SetEmbeddingModel(null)
    getEmbeddingDocumentsBody.SetIncludeFileLevelMetadata(false)
    getEmbeddingDocumentsBody.SetHighAccuracy(false)
    getEmbeddingDocumentsBody.SetRerank(rerank)
    getEmbeddingDocumentsBody.SetFileTypesAtSource(null)
    getEmbeddingDocumentsBody.SetExcludeColdStorageFiles(false)
    
    request := client.EmbeddingsApi.GetDocuments(
        getEmbeddingDocumentsBody,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmbeddingsApi.GetDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `GetDocuments`: DocumentResponseList
    fmt.Fprintf(os.Stdout, "Response from `EmbeddingsApi.GetDocuments`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `DocumentResponseList.GetDocuments.Documents`: %v\n", resp.Documents)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmbeddingsAndChunks

Retrieve Embeddings And Content

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
    filters := *carbon.NewEmbeddingsAndChunksFilters()
    
    embeddingsAndChunksQueryInput := *carbon.NewEmbeddingsAndChunksQueryInput(
        filters,
    )
    embeddingsAndChunksQueryInput.SetPagination(pagination)
    embeddingsAndChunksQueryInput.SetOrderBy(null)
    embeddingsAndChunksQueryInput.SetOrderDir(null)
    embeddingsAndChunksQueryInput.SetIncludeVectors(false)
    
    request := client.EmbeddingsApi.GetEmbeddingsAndChunks(
        embeddingsAndChunksQueryInput,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmbeddingsApi.GetEmbeddingsAndChunks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `GetEmbeddingsAndChunks`: EmbeddingsAndChunksResponse
    fmt.Fprintf(os.Stdout, "Response from `EmbeddingsApi.GetEmbeddingsAndChunks`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `EmbeddingsAndChunksResponse.GetEmbeddingsAndChunks.Results`: %v\n", resp.Results)
    fmt.Fprintf(os.Stdout, "Response from `EmbeddingsAndChunksResponse.GetEmbeddingsAndChunks.Count`: %v\n", resp.Count)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

Retrieve Embeddings And Content V2

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
    filters := *carbon.NewOrganizationUserFilesToSyncFilters()
    
    embeddingsAndChunksQueryInputV2 := *carbon.NewEmbeddingsAndChunksQueryInputV2(
        filters,
    )
    embeddingsAndChunksQueryInputV2.SetPagination(pagination)
    embeddingsAndChunksQueryInputV2.SetOrderBy(null)
    embeddingsAndChunksQueryInputV2.SetOrderDir(null)
    embeddingsAndChunksQueryInputV2.SetIncludeVectors(false)
    
    request := client.EmbeddingsApi.List(
        embeddingsAndChunksQueryInputV2,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmbeddingsApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `List`: EmbeddingsAndChunksResponse
    fmt.Fprintf(os.Stdout, "Response from `EmbeddingsApi.List`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `EmbeddingsAndChunksResponse.List.Results`: %v\n", resp.Results)
    fmt.Fprintf(os.Stdout, "Response from `EmbeddingsAndChunksResponse.List.Count`: %v\n", resp.Count)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadChunksAndEmbeddings

Upload Chunks And Embeddings

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

    
    chunksAndEmbeddingsUploadInput := *carbon.NewChunksAndEmbeddingsUploadInput(
        null,
        null,
    )
    chunksAndEmbeddingsUploadInput.SetOverwriteExisting(false)
    chunksAndEmbeddingsUploadInput.SetChunksOnly(false)
    chunksAndEmbeddingsUploadInput.SetCustomCredentials({})
    
    request := client.EmbeddingsApi.UploadChunksAndEmbeddings(
        chunksAndEmbeddingsUploadInput,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmbeddingsApi.UploadChunksAndEmbeddings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `UploadChunksAndEmbeddings`: GenericSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `EmbeddingsApi.UploadChunksAndEmbeddings`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `GenericSuccessResponse.UploadChunksAndEmbeddings.Success`: %v\n", resp.Success)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# SitemapScrapeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to [**map[string]Tags1**](Tags1.md) |  | [optional] [default to {}]
**Url** | **string** |  | 
**MaxPagesToScrape** | Pointer to **NullableInt32** |  | [optional] 
**ChunkSize** | Pointer to **NullableInt32** |  | [optional] [default to 1500]
**ChunkOverlap** | Pointer to **NullableInt32** |  | [optional] [default to 20]
**SkipEmbeddingGeneration** | Pointer to **NullableBool** |  | [optional] [default to false]
**EnableAutoSync** | Pointer to **NullableBool** |  | [optional] [default to false]
**GenerateSparseVectors** | Pointer to **NullableBool** |  | [optional] [default to false]
**PrependFilenameToChunks** | Pointer to **NullableBool** |  | [optional] [default to false]
**HtmlTagsToSkip** | Pointer to **[]string** |  | [optional] [default to []]
**CssClassesToSkip** | Pointer to **[]string** |  | [optional] [default to []]
**CssSelectorsToSkip** | Pointer to **[]string** |  | [optional] [default to []]
**EmbeddingModel** | Pointer to [**EmbeddingGenerators**](EmbeddingGenerators.md) |  | [optional] 
**UrlPathsToInclude** | Pointer to **[]string** | URL subpaths or directories that you want to include. For example if you want to only include         URLs that start with /questions in stackoverflow.com, you will add /questions/ in this input | [optional] [default to []]
**UrlPathsToExclude** | Pointer to **[]string** | URL subpaths or directories that you want to exclude. For example if you want to exclude         URLs that start with /questions in stackoverflow.com, you will add /questions/ in this input | [optional] [default to []]
**UrlsToScrape** | Pointer to **[]string** | You can submit a subset of URLs from the sitemap that should be scraped. To get the list of URLs,           you can check out /process_sitemap endpoint. If left empty, all URLs from the sitemap will be scraped. | [optional] [default to []]
**DownloadCssAndMedia** | Pointer to **NullableBool** | Whether the scraper should download css and media from the page (images, fonts, etc). Scrapes          might take longer to finish with this flag enabled, but the success rate is improved. | [optional] [default to false]
**GenerateChunksOnly** | Pointer to **bool** | If this flag is enabled, the file will be chunked and stored with Carbon,           but no embeddings will be generated. This overrides the skip_embedding_generation flag. | [optional] [default to false]
**StoreFileOnly** | Pointer to **bool** | If this flag is enabled, the file will be stored with Carbon, but no processing will be done. | [optional] [default to false]
**UsePremiumProxies** | Pointer to **bool** | If the default proxies are blocked and not returning results, this flag can be enabled to use              alternate proxies (residential and office). Scrapes might take longer to finish with this flag enabled.          | [optional] [default to false]

## Methods

### NewSitemapScrapeRequest

`func NewSitemapScrapeRequest(url string, ) *SitemapScrapeRequest`

NewSitemapScrapeRequest instantiates a new SitemapScrapeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSitemapScrapeRequestWithDefaults

`func NewSitemapScrapeRequestWithDefaults() *SitemapScrapeRequest`

NewSitemapScrapeRequestWithDefaults instantiates a new SitemapScrapeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *SitemapScrapeRequest) GetTags() map[string]Tags1`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SitemapScrapeRequest) GetTagsOk() (*map[string]Tags1, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SitemapScrapeRequest) SetTags(v map[string]Tags1)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SitemapScrapeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *SitemapScrapeRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *SitemapScrapeRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetUrl

`func (o *SitemapScrapeRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SitemapScrapeRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SitemapScrapeRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMaxPagesToScrape

`func (o *SitemapScrapeRequest) GetMaxPagesToScrape() int32`

GetMaxPagesToScrape returns the MaxPagesToScrape field if non-nil, zero value otherwise.

### GetMaxPagesToScrapeOk

`func (o *SitemapScrapeRequest) GetMaxPagesToScrapeOk() (*int32, bool)`

GetMaxPagesToScrapeOk returns a tuple with the MaxPagesToScrape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPagesToScrape

`func (o *SitemapScrapeRequest) SetMaxPagesToScrape(v int32)`

SetMaxPagesToScrape sets MaxPagesToScrape field to given value.

### HasMaxPagesToScrape

`func (o *SitemapScrapeRequest) HasMaxPagesToScrape() bool`

HasMaxPagesToScrape returns a boolean if a field has been set.

### SetMaxPagesToScrapeNil

`func (o *SitemapScrapeRequest) SetMaxPagesToScrapeNil(b bool)`

 SetMaxPagesToScrapeNil sets the value for MaxPagesToScrape to be an explicit nil

### UnsetMaxPagesToScrape
`func (o *SitemapScrapeRequest) UnsetMaxPagesToScrape()`

UnsetMaxPagesToScrape ensures that no value is present for MaxPagesToScrape, not even an explicit nil
### GetChunkSize

`func (o *SitemapScrapeRequest) GetChunkSize() int32`

GetChunkSize returns the ChunkSize field if non-nil, zero value otherwise.

### GetChunkSizeOk

`func (o *SitemapScrapeRequest) GetChunkSizeOk() (*int32, bool)`

GetChunkSizeOk returns a tuple with the ChunkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkSize

`func (o *SitemapScrapeRequest) SetChunkSize(v int32)`

SetChunkSize sets ChunkSize field to given value.

### HasChunkSize

`func (o *SitemapScrapeRequest) HasChunkSize() bool`

HasChunkSize returns a boolean if a field has been set.

### SetChunkSizeNil

`func (o *SitemapScrapeRequest) SetChunkSizeNil(b bool)`

 SetChunkSizeNil sets the value for ChunkSize to be an explicit nil

### UnsetChunkSize
`func (o *SitemapScrapeRequest) UnsetChunkSize()`

UnsetChunkSize ensures that no value is present for ChunkSize, not even an explicit nil
### GetChunkOverlap

`func (o *SitemapScrapeRequest) GetChunkOverlap() int32`

GetChunkOverlap returns the ChunkOverlap field if non-nil, zero value otherwise.

### GetChunkOverlapOk

`func (o *SitemapScrapeRequest) GetChunkOverlapOk() (*int32, bool)`

GetChunkOverlapOk returns a tuple with the ChunkOverlap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkOverlap

`func (o *SitemapScrapeRequest) SetChunkOverlap(v int32)`

SetChunkOverlap sets ChunkOverlap field to given value.

### HasChunkOverlap

`func (o *SitemapScrapeRequest) HasChunkOverlap() bool`

HasChunkOverlap returns a boolean if a field has been set.

### SetChunkOverlapNil

`func (o *SitemapScrapeRequest) SetChunkOverlapNil(b bool)`

 SetChunkOverlapNil sets the value for ChunkOverlap to be an explicit nil

### UnsetChunkOverlap
`func (o *SitemapScrapeRequest) UnsetChunkOverlap()`

UnsetChunkOverlap ensures that no value is present for ChunkOverlap, not even an explicit nil
### GetSkipEmbeddingGeneration

`func (o *SitemapScrapeRequest) GetSkipEmbeddingGeneration() bool`

GetSkipEmbeddingGeneration returns the SkipEmbeddingGeneration field if non-nil, zero value otherwise.

### GetSkipEmbeddingGenerationOk

`func (o *SitemapScrapeRequest) GetSkipEmbeddingGenerationOk() (*bool, bool)`

GetSkipEmbeddingGenerationOk returns a tuple with the SkipEmbeddingGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipEmbeddingGeneration

`func (o *SitemapScrapeRequest) SetSkipEmbeddingGeneration(v bool)`

SetSkipEmbeddingGeneration sets SkipEmbeddingGeneration field to given value.

### HasSkipEmbeddingGeneration

`func (o *SitemapScrapeRequest) HasSkipEmbeddingGeneration() bool`

HasSkipEmbeddingGeneration returns a boolean if a field has been set.

### SetSkipEmbeddingGenerationNil

`func (o *SitemapScrapeRequest) SetSkipEmbeddingGenerationNil(b bool)`

 SetSkipEmbeddingGenerationNil sets the value for SkipEmbeddingGeneration to be an explicit nil

### UnsetSkipEmbeddingGeneration
`func (o *SitemapScrapeRequest) UnsetSkipEmbeddingGeneration()`

UnsetSkipEmbeddingGeneration ensures that no value is present for SkipEmbeddingGeneration, not even an explicit nil
### GetEnableAutoSync

`func (o *SitemapScrapeRequest) GetEnableAutoSync() bool`

GetEnableAutoSync returns the EnableAutoSync field if non-nil, zero value otherwise.

### GetEnableAutoSyncOk

`func (o *SitemapScrapeRequest) GetEnableAutoSyncOk() (*bool, bool)`

GetEnableAutoSyncOk returns a tuple with the EnableAutoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoSync

`func (o *SitemapScrapeRequest) SetEnableAutoSync(v bool)`

SetEnableAutoSync sets EnableAutoSync field to given value.

### HasEnableAutoSync

`func (o *SitemapScrapeRequest) HasEnableAutoSync() bool`

HasEnableAutoSync returns a boolean if a field has been set.

### SetEnableAutoSyncNil

`func (o *SitemapScrapeRequest) SetEnableAutoSyncNil(b bool)`

 SetEnableAutoSyncNil sets the value for EnableAutoSync to be an explicit nil

### UnsetEnableAutoSync
`func (o *SitemapScrapeRequest) UnsetEnableAutoSync()`

UnsetEnableAutoSync ensures that no value is present for EnableAutoSync, not even an explicit nil
### GetGenerateSparseVectors

`func (o *SitemapScrapeRequest) GetGenerateSparseVectors() bool`

GetGenerateSparseVectors returns the GenerateSparseVectors field if non-nil, zero value otherwise.

### GetGenerateSparseVectorsOk

`func (o *SitemapScrapeRequest) GetGenerateSparseVectorsOk() (*bool, bool)`

GetGenerateSparseVectorsOk returns a tuple with the GenerateSparseVectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateSparseVectors

`func (o *SitemapScrapeRequest) SetGenerateSparseVectors(v bool)`

SetGenerateSparseVectors sets GenerateSparseVectors field to given value.

### HasGenerateSparseVectors

`func (o *SitemapScrapeRequest) HasGenerateSparseVectors() bool`

HasGenerateSparseVectors returns a boolean if a field has been set.

### SetGenerateSparseVectorsNil

`func (o *SitemapScrapeRequest) SetGenerateSparseVectorsNil(b bool)`

 SetGenerateSparseVectorsNil sets the value for GenerateSparseVectors to be an explicit nil

### UnsetGenerateSparseVectors
`func (o *SitemapScrapeRequest) UnsetGenerateSparseVectors()`

UnsetGenerateSparseVectors ensures that no value is present for GenerateSparseVectors, not even an explicit nil
### GetPrependFilenameToChunks

`func (o *SitemapScrapeRequest) GetPrependFilenameToChunks() bool`

GetPrependFilenameToChunks returns the PrependFilenameToChunks field if non-nil, zero value otherwise.

### GetPrependFilenameToChunksOk

`func (o *SitemapScrapeRequest) GetPrependFilenameToChunksOk() (*bool, bool)`

GetPrependFilenameToChunksOk returns a tuple with the PrependFilenameToChunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrependFilenameToChunks

`func (o *SitemapScrapeRequest) SetPrependFilenameToChunks(v bool)`

SetPrependFilenameToChunks sets PrependFilenameToChunks field to given value.

### HasPrependFilenameToChunks

`func (o *SitemapScrapeRequest) HasPrependFilenameToChunks() bool`

HasPrependFilenameToChunks returns a boolean if a field has been set.

### SetPrependFilenameToChunksNil

`func (o *SitemapScrapeRequest) SetPrependFilenameToChunksNil(b bool)`

 SetPrependFilenameToChunksNil sets the value for PrependFilenameToChunks to be an explicit nil

### UnsetPrependFilenameToChunks
`func (o *SitemapScrapeRequest) UnsetPrependFilenameToChunks()`

UnsetPrependFilenameToChunks ensures that no value is present for PrependFilenameToChunks, not even an explicit nil
### GetHtmlTagsToSkip

`func (o *SitemapScrapeRequest) GetHtmlTagsToSkip() []string`

GetHtmlTagsToSkip returns the HtmlTagsToSkip field if non-nil, zero value otherwise.

### GetHtmlTagsToSkipOk

`func (o *SitemapScrapeRequest) GetHtmlTagsToSkipOk() (*[]string, bool)`

GetHtmlTagsToSkipOk returns a tuple with the HtmlTagsToSkip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlTagsToSkip

`func (o *SitemapScrapeRequest) SetHtmlTagsToSkip(v []string)`

SetHtmlTagsToSkip sets HtmlTagsToSkip field to given value.

### HasHtmlTagsToSkip

`func (o *SitemapScrapeRequest) HasHtmlTagsToSkip() bool`

HasHtmlTagsToSkip returns a boolean if a field has been set.

### SetHtmlTagsToSkipNil

`func (o *SitemapScrapeRequest) SetHtmlTagsToSkipNil(b bool)`

 SetHtmlTagsToSkipNil sets the value for HtmlTagsToSkip to be an explicit nil

### UnsetHtmlTagsToSkip
`func (o *SitemapScrapeRequest) UnsetHtmlTagsToSkip()`

UnsetHtmlTagsToSkip ensures that no value is present for HtmlTagsToSkip, not even an explicit nil
### GetCssClassesToSkip

`func (o *SitemapScrapeRequest) GetCssClassesToSkip() []string`

GetCssClassesToSkip returns the CssClassesToSkip field if non-nil, zero value otherwise.

### GetCssClassesToSkipOk

`func (o *SitemapScrapeRequest) GetCssClassesToSkipOk() (*[]string, bool)`

GetCssClassesToSkipOk returns a tuple with the CssClassesToSkip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssClassesToSkip

`func (o *SitemapScrapeRequest) SetCssClassesToSkip(v []string)`

SetCssClassesToSkip sets CssClassesToSkip field to given value.

### HasCssClassesToSkip

`func (o *SitemapScrapeRequest) HasCssClassesToSkip() bool`

HasCssClassesToSkip returns a boolean if a field has been set.

### SetCssClassesToSkipNil

`func (o *SitemapScrapeRequest) SetCssClassesToSkipNil(b bool)`

 SetCssClassesToSkipNil sets the value for CssClassesToSkip to be an explicit nil

### UnsetCssClassesToSkip
`func (o *SitemapScrapeRequest) UnsetCssClassesToSkip()`

UnsetCssClassesToSkip ensures that no value is present for CssClassesToSkip, not even an explicit nil
### GetCssSelectorsToSkip

`func (o *SitemapScrapeRequest) GetCssSelectorsToSkip() []string`

GetCssSelectorsToSkip returns the CssSelectorsToSkip field if non-nil, zero value otherwise.

### GetCssSelectorsToSkipOk

`func (o *SitemapScrapeRequest) GetCssSelectorsToSkipOk() (*[]string, bool)`

GetCssSelectorsToSkipOk returns a tuple with the CssSelectorsToSkip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssSelectorsToSkip

`func (o *SitemapScrapeRequest) SetCssSelectorsToSkip(v []string)`

SetCssSelectorsToSkip sets CssSelectorsToSkip field to given value.

### HasCssSelectorsToSkip

`func (o *SitemapScrapeRequest) HasCssSelectorsToSkip() bool`

HasCssSelectorsToSkip returns a boolean if a field has been set.

### SetCssSelectorsToSkipNil

`func (o *SitemapScrapeRequest) SetCssSelectorsToSkipNil(b bool)`

 SetCssSelectorsToSkipNil sets the value for CssSelectorsToSkip to be an explicit nil

### UnsetCssSelectorsToSkip
`func (o *SitemapScrapeRequest) UnsetCssSelectorsToSkip()`

UnsetCssSelectorsToSkip ensures that no value is present for CssSelectorsToSkip, not even an explicit nil
### GetEmbeddingModel

`func (o *SitemapScrapeRequest) GetEmbeddingModel() EmbeddingGenerators`

GetEmbeddingModel returns the EmbeddingModel field if non-nil, zero value otherwise.

### GetEmbeddingModelOk

`func (o *SitemapScrapeRequest) GetEmbeddingModelOk() (*EmbeddingGenerators, bool)`

GetEmbeddingModelOk returns a tuple with the EmbeddingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingModel

`func (o *SitemapScrapeRequest) SetEmbeddingModel(v EmbeddingGenerators)`

SetEmbeddingModel sets EmbeddingModel field to given value.

### HasEmbeddingModel

`func (o *SitemapScrapeRequest) HasEmbeddingModel() bool`

HasEmbeddingModel returns a boolean if a field has been set.

### GetUrlPathsToInclude

`func (o *SitemapScrapeRequest) GetUrlPathsToInclude() []string`

GetUrlPathsToInclude returns the UrlPathsToInclude field if non-nil, zero value otherwise.

### GetUrlPathsToIncludeOk

`func (o *SitemapScrapeRequest) GetUrlPathsToIncludeOk() (*[]string, bool)`

GetUrlPathsToIncludeOk returns a tuple with the UrlPathsToInclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPathsToInclude

`func (o *SitemapScrapeRequest) SetUrlPathsToInclude(v []string)`

SetUrlPathsToInclude sets UrlPathsToInclude field to given value.

### HasUrlPathsToInclude

`func (o *SitemapScrapeRequest) HasUrlPathsToInclude() bool`

HasUrlPathsToInclude returns a boolean if a field has been set.

### SetUrlPathsToIncludeNil

`func (o *SitemapScrapeRequest) SetUrlPathsToIncludeNil(b bool)`

 SetUrlPathsToIncludeNil sets the value for UrlPathsToInclude to be an explicit nil

### UnsetUrlPathsToInclude
`func (o *SitemapScrapeRequest) UnsetUrlPathsToInclude()`

UnsetUrlPathsToInclude ensures that no value is present for UrlPathsToInclude, not even an explicit nil
### GetUrlPathsToExclude

`func (o *SitemapScrapeRequest) GetUrlPathsToExclude() []string`

GetUrlPathsToExclude returns the UrlPathsToExclude field if non-nil, zero value otherwise.

### GetUrlPathsToExcludeOk

`func (o *SitemapScrapeRequest) GetUrlPathsToExcludeOk() (*[]string, bool)`

GetUrlPathsToExcludeOk returns a tuple with the UrlPathsToExclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPathsToExclude

`func (o *SitemapScrapeRequest) SetUrlPathsToExclude(v []string)`

SetUrlPathsToExclude sets UrlPathsToExclude field to given value.

### HasUrlPathsToExclude

`func (o *SitemapScrapeRequest) HasUrlPathsToExclude() bool`

HasUrlPathsToExclude returns a boolean if a field has been set.

### SetUrlPathsToExcludeNil

`func (o *SitemapScrapeRequest) SetUrlPathsToExcludeNil(b bool)`

 SetUrlPathsToExcludeNil sets the value for UrlPathsToExclude to be an explicit nil

### UnsetUrlPathsToExclude
`func (o *SitemapScrapeRequest) UnsetUrlPathsToExclude()`

UnsetUrlPathsToExclude ensures that no value is present for UrlPathsToExclude, not even an explicit nil
### GetUrlsToScrape

`func (o *SitemapScrapeRequest) GetUrlsToScrape() []string`

GetUrlsToScrape returns the UrlsToScrape field if non-nil, zero value otherwise.

### GetUrlsToScrapeOk

`func (o *SitemapScrapeRequest) GetUrlsToScrapeOk() (*[]string, bool)`

GetUrlsToScrapeOk returns a tuple with the UrlsToScrape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlsToScrape

`func (o *SitemapScrapeRequest) SetUrlsToScrape(v []string)`

SetUrlsToScrape sets UrlsToScrape field to given value.

### HasUrlsToScrape

`func (o *SitemapScrapeRequest) HasUrlsToScrape() bool`

HasUrlsToScrape returns a boolean if a field has been set.

### SetUrlsToScrapeNil

`func (o *SitemapScrapeRequest) SetUrlsToScrapeNil(b bool)`

 SetUrlsToScrapeNil sets the value for UrlsToScrape to be an explicit nil

### UnsetUrlsToScrape
`func (o *SitemapScrapeRequest) UnsetUrlsToScrape()`

UnsetUrlsToScrape ensures that no value is present for UrlsToScrape, not even an explicit nil
### GetDownloadCssAndMedia

`func (o *SitemapScrapeRequest) GetDownloadCssAndMedia() bool`

GetDownloadCssAndMedia returns the DownloadCssAndMedia field if non-nil, zero value otherwise.

### GetDownloadCssAndMediaOk

`func (o *SitemapScrapeRequest) GetDownloadCssAndMediaOk() (*bool, bool)`

GetDownloadCssAndMediaOk returns a tuple with the DownloadCssAndMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadCssAndMedia

`func (o *SitemapScrapeRequest) SetDownloadCssAndMedia(v bool)`

SetDownloadCssAndMedia sets DownloadCssAndMedia field to given value.

### HasDownloadCssAndMedia

`func (o *SitemapScrapeRequest) HasDownloadCssAndMedia() bool`

HasDownloadCssAndMedia returns a boolean if a field has been set.

### SetDownloadCssAndMediaNil

`func (o *SitemapScrapeRequest) SetDownloadCssAndMediaNil(b bool)`

 SetDownloadCssAndMediaNil sets the value for DownloadCssAndMedia to be an explicit nil

### UnsetDownloadCssAndMedia
`func (o *SitemapScrapeRequest) UnsetDownloadCssAndMedia()`

UnsetDownloadCssAndMedia ensures that no value is present for DownloadCssAndMedia, not even an explicit nil
### GetGenerateChunksOnly

`func (o *SitemapScrapeRequest) GetGenerateChunksOnly() bool`

GetGenerateChunksOnly returns the GenerateChunksOnly field if non-nil, zero value otherwise.

### GetGenerateChunksOnlyOk

`func (o *SitemapScrapeRequest) GetGenerateChunksOnlyOk() (*bool, bool)`

GetGenerateChunksOnlyOk returns a tuple with the GenerateChunksOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateChunksOnly

`func (o *SitemapScrapeRequest) SetGenerateChunksOnly(v bool)`

SetGenerateChunksOnly sets GenerateChunksOnly field to given value.

### HasGenerateChunksOnly

`func (o *SitemapScrapeRequest) HasGenerateChunksOnly() bool`

HasGenerateChunksOnly returns a boolean if a field has been set.

### GetStoreFileOnly

`func (o *SitemapScrapeRequest) GetStoreFileOnly() bool`

GetStoreFileOnly returns the StoreFileOnly field if non-nil, zero value otherwise.

### GetStoreFileOnlyOk

`func (o *SitemapScrapeRequest) GetStoreFileOnlyOk() (*bool, bool)`

GetStoreFileOnlyOk returns a tuple with the StoreFileOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreFileOnly

`func (o *SitemapScrapeRequest) SetStoreFileOnly(v bool)`

SetStoreFileOnly sets StoreFileOnly field to given value.

### HasStoreFileOnly

`func (o *SitemapScrapeRequest) HasStoreFileOnly() bool`

HasStoreFileOnly returns a boolean if a field has been set.

### GetUsePremiumProxies

`func (o *SitemapScrapeRequest) GetUsePremiumProxies() bool`

GetUsePremiumProxies returns the UsePremiumProxies field if non-nil, zero value otherwise.

### GetUsePremiumProxiesOk

`func (o *SitemapScrapeRequest) GetUsePremiumProxiesOk() (*bool, bool)`

GetUsePremiumProxiesOk returns a tuple with the UsePremiumProxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePremiumProxies

`func (o *SitemapScrapeRequest) SetUsePremiumProxies(v bool)`

SetUsePremiumProxies sets UsePremiumProxies field to given value.

### HasUsePremiumProxies

`func (o *SitemapScrapeRequest) HasUsePremiumProxies() bool`

HasUsePremiumProxies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



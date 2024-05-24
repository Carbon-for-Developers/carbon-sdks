# FetchURLsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Urls** | **[]string** |  | 
**HtmlContent** | **string** |  | 
**ErrorMessage** | **NullableString** |  | 

## Methods

### NewFetchURLsResponse

`func NewFetchURLsResponse(urls []string, htmlContent string, errorMessage NullableString, ) *FetchURLsResponse`

NewFetchURLsResponse instantiates a new FetchURLsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchURLsResponseWithDefaults

`func NewFetchURLsResponseWithDefaults() *FetchURLsResponse`

NewFetchURLsResponseWithDefaults instantiates a new FetchURLsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrls

`func (o *FetchURLsResponse) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *FetchURLsResponse) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *FetchURLsResponse) SetUrls(v []string)`

SetUrls sets Urls field to given value.


### GetHtmlContent

`func (o *FetchURLsResponse) GetHtmlContent() string`

GetHtmlContent returns the HtmlContent field if non-nil, zero value otherwise.

### GetHtmlContentOk

`func (o *FetchURLsResponse) GetHtmlContentOk() (*string, bool)`

GetHtmlContentOk returns a tuple with the HtmlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlContent

`func (o *FetchURLsResponse) SetHtmlContent(v string)`

SetHtmlContent sets HtmlContent field to given value.


### GetErrorMessage

`func (o *FetchURLsResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *FetchURLsResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *FetchURLsResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *FetchURLsResponse) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *FetchURLsResponse) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



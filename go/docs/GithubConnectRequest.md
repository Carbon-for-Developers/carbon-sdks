# GithubConnectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**AccessToken** | **string** |  | 
**SyncSourceItems** | Pointer to **bool** | Enabling this flag will fetch all available content from the source to be listed via list items endpoint | [optional] [default to false]

## Methods

### NewGithubConnectRequest

`func NewGithubConnectRequest(username string, accessToken string, ) *GithubConnectRequest`

NewGithubConnectRequest instantiates a new GithubConnectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubConnectRequestWithDefaults

`func NewGithubConnectRequestWithDefaults() *GithubConnectRequest`

NewGithubConnectRequestWithDefaults instantiates a new GithubConnectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *GithubConnectRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GithubConnectRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GithubConnectRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetAccessToken

`func (o *GithubConnectRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *GithubConnectRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *GithubConnectRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetSyncSourceItems

`func (o *GithubConnectRequest) GetSyncSourceItems() bool`

GetSyncSourceItems returns the SyncSourceItems field if non-nil, zero value otherwise.

### GetSyncSourceItemsOk

`func (o *GithubConnectRequest) GetSyncSourceItemsOk() (*bool, bool)`

GetSyncSourceItemsOk returns a tuple with the SyncSourceItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncSourceItems

`func (o *GithubConnectRequest) SetSyncSourceItems(v bool)`

SetSyncSourceItems sets SyncSourceItems field to given value.

### HasSyncSourceItems

`func (o *GithubConnectRequest) HasSyncSourceItems() bool`

HasSyncSourceItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



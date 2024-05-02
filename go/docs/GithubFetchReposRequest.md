# GithubFetchReposRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repos** | **[]string** |  | 
**DataSourceId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewGithubFetchReposRequest

`func NewGithubFetchReposRequest(repos []string, ) *GithubFetchReposRequest`

NewGithubFetchReposRequest instantiates a new GithubFetchReposRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubFetchReposRequestWithDefaults

`func NewGithubFetchReposRequestWithDefaults() *GithubFetchReposRequest`

NewGithubFetchReposRequestWithDefaults instantiates a new GithubFetchReposRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepos

`func (o *GithubFetchReposRequest) GetRepos() []string`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *GithubFetchReposRequest) GetReposOk() (*[]string, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *GithubFetchReposRequest) SetRepos(v []string)`

SetRepos sets Repos field to given value.


### GetDataSourceId

`func (o *GithubFetchReposRequest) GetDataSourceId() int32`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *GithubFetchReposRequest) GetDataSourceIdOk() (*int32, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *GithubFetchReposRequest) SetDataSourceId(v int32)`

SetDataSourceId sets DataSourceId field to given value.

### HasDataSourceId

`func (o *GithubFetchReposRequest) HasDataSourceId() bool`

HasDataSourceId returns a boolean if a field has been set.

### SetDataSourceIdNil

`func (o *GithubFetchReposRequest) SetDataSourceIdNil(b bool)`

 SetDataSourceIdNil sets the value for DataSourceId to be an explicit nil

### UnsetDataSourceId
`func (o *GithubFetchReposRequest) UnsetDataSourceId()`

UnsetDataSourceId ensures that no value is present for DataSourceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



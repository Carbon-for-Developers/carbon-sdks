# PullRequestFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**PRStateInput**](PRStateInput.md) |  | [optional] 
**Base** | Pointer to **NullableString** | Filter pulls by base branch name | [optional] 
**Head** | Pointer to **NullableString** | Filter pulls by head user or head organization and branch name in the format of user:ref-name or organization:ref-name | [optional] 

## Methods

### NewPullRequestFilters

`func NewPullRequestFilters() *PullRequestFilters`

NewPullRequestFilters instantiates a new PullRequestFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestFiltersWithDefaults

`func NewPullRequestFiltersWithDefaults() *PullRequestFilters`

NewPullRequestFiltersWithDefaults instantiates a new PullRequestFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *PullRequestFilters) GetState() PRStateInput`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PullRequestFilters) GetStateOk() (*PRStateInput, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PullRequestFilters) SetState(v PRStateInput)`

SetState sets State field to given value.

### HasState

`func (o *PullRequestFilters) HasState() bool`

HasState returns a boolean if a field has been set.

### GetBase

`func (o *PullRequestFilters) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *PullRequestFilters) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *PullRequestFilters) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *PullRequestFilters) HasBase() bool`

HasBase returns a boolean if a field has been set.

### SetBaseNil

`func (o *PullRequestFilters) SetBaseNil(b bool)`

 SetBaseNil sets the value for Base to be an explicit nil

### UnsetBase
`func (o *PullRequestFilters) UnsetBase()`

UnsetBase ensures that no value is present for Base, not even an explicit nil
### GetHead

`func (o *PullRequestFilters) GetHead() string`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *PullRequestFilters) GetHeadOk() (*string, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *PullRequestFilters) SetHead(v string)`

SetHead sets Head field to given value.

### HasHead

`func (o *PullRequestFilters) HasHead() bool`

HasHead returns a boolean if a field has been set.

### SetHeadNil

`func (o *PullRequestFilters) SetHeadNil(b bool)`

 SetHeadNil sets the value for Head to be an explicit nil

### UnsetHead
`func (o *PullRequestFilters) UnsetHead()`

UnsetHead ensures that no value is present for Head, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



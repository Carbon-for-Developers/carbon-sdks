# PullRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Id** | **int32** |  | 
**Url** | **string** |  | 
**Number** | **int32** |  | 
**State** | [**PRState**](PRState.md) |  | 
**User** | [**User**](User.md) |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**ClosedAt** | **NullableString** |  | 
**MergedAt** | **NullableString** |  | 
**RequestedReviewers** | [**[]User**](User.md) |  | 
**RequestedTeams** | [**[]Team**](Team.md) |  | 
**Labels** | [**[]Label**](Label.md) |  | 
**Draft** | **bool** |  | 
**Head** | [**HeadOrBase**](HeadOrBase.md) |  | 
**Base** | [**HeadOrBase**](HeadOrBase.md) |  | 
**RemoteData** | **map[string]interface{}** |  | 

## Methods

### NewPullRequest

`func NewPullRequest(title string, id int32, url string, number int32, state PRState, user User, createdAt string, updatedAt string, closedAt NullableString, mergedAt NullableString, requestedReviewers []User, requestedTeams []Team, labels []Label, draft bool, head HeadOrBase, base HeadOrBase, remoteData map[string]interface{}, ) *PullRequest`

NewPullRequest instantiates a new PullRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestWithDefaults

`func NewPullRequestWithDefaults() *PullRequest`

NewPullRequestWithDefaults instantiates a new PullRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *PullRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PullRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PullRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetId

`func (o *PullRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PullRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PullRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *PullRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PullRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PullRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNumber

`func (o *PullRequest) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PullRequest) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PullRequest) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetState

`func (o *PullRequest) GetState() PRState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PullRequest) GetStateOk() (*PRState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PullRequest) SetState(v PRState)`

SetState sets State field to given value.


### GetUser

`func (o *PullRequest) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PullRequest) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PullRequest) SetUser(v User)`

SetUser sets User field to given value.


### GetCreatedAt

`func (o *PullRequest) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PullRequest) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PullRequest) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PullRequest) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PullRequest) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PullRequest) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetClosedAt

`func (o *PullRequest) GetClosedAt() string`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *PullRequest) GetClosedAtOk() (*string, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *PullRequest) SetClosedAt(v string)`

SetClosedAt sets ClosedAt field to given value.


### SetClosedAtNil

`func (o *PullRequest) SetClosedAtNil(b bool)`

 SetClosedAtNil sets the value for ClosedAt to be an explicit nil

### UnsetClosedAt
`func (o *PullRequest) UnsetClosedAt()`

UnsetClosedAt ensures that no value is present for ClosedAt, not even an explicit nil
### GetMergedAt

`func (o *PullRequest) GetMergedAt() string`

GetMergedAt returns the MergedAt field if non-nil, zero value otherwise.

### GetMergedAtOk

`func (o *PullRequest) GetMergedAtOk() (*string, bool)`

GetMergedAtOk returns a tuple with the MergedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedAt

`func (o *PullRequest) SetMergedAt(v string)`

SetMergedAt sets MergedAt field to given value.


### SetMergedAtNil

`func (o *PullRequest) SetMergedAtNil(b bool)`

 SetMergedAtNil sets the value for MergedAt to be an explicit nil

### UnsetMergedAt
`func (o *PullRequest) UnsetMergedAt()`

UnsetMergedAt ensures that no value is present for MergedAt, not even an explicit nil
### GetRequestedReviewers

`func (o *PullRequest) GetRequestedReviewers() []User`

GetRequestedReviewers returns the RequestedReviewers field if non-nil, zero value otherwise.

### GetRequestedReviewersOk

`func (o *PullRequest) GetRequestedReviewersOk() (*[]User, bool)`

GetRequestedReviewersOk returns a tuple with the RequestedReviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedReviewers

`func (o *PullRequest) SetRequestedReviewers(v []User)`

SetRequestedReviewers sets RequestedReviewers field to given value.


### GetRequestedTeams

`func (o *PullRequest) GetRequestedTeams() []Team`

GetRequestedTeams returns the RequestedTeams field if non-nil, zero value otherwise.

### GetRequestedTeamsOk

`func (o *PullRequest) GetRequestedTeamsOk() (*[]Team, bool)`

GetRequestedTeamsOk returns a tuple with the RequestedTeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTeams

`func (o *PullRequest) SetRequestedTeams(v []Team)`

SetRequestedTeams sets RequestedTeams field to given value.


### GetLabels

`func (o *PullRequest) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PullRequest) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PullRequest) SetLabels(v []Label)`

SetLabels sets Labels field to given value.


### GetDraft

`func (o *PullRequest) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *PullRequest) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *PullRequest) SetDraft(v bool)`

SetDraft sets Draft field to given value.


### GetHead

`func (o *PullRequest) GetHead() HeadOrBase`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *PullRequest) GetHeadOk() (*HeadOrBase, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *PullRequest) SetHead(v HeadOrBase)`

SetHead sets Head field to given value.


### GetBase

`func (o *PullRequest) GetBase() HeadOrBase`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *PullRequest) GetBaseOk() (*HeadOrBase, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *PullRequest) SetBase(v HeadOrBase)`

SetBase sets Base field to given value.


### GetRemoteData

`func (o *PullRequest) GetRemoteData() map[string]interface{}`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *PullRequest) GetRemoteDataOk() (*map[string]interface{}, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *PullRequest) SetRemoteData(v map[string]interface{})`

SetRemoteData sets RemoteData field to given value.


### SetRemoteDataNil

`func (o *PullRequest) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *PullRequest) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



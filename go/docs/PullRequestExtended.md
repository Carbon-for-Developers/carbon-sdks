# PullRequestExtended

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
**Merged** | **bool** |  | 
**NumComments** | **int32** |  | 
**NumReviewComments** | **int32** |  | 
**NumCommits** | **int32** |  | 
**NumAdditions** | **int32** |  | 
**NumDeletions** | **int32** |  | 
**NumChangedFiles** | **int32** |  | 

## Methods

### NewPullRequestExtended

`func NewPullRequestExtended(title string, id int32, url string, number int32, state PRState, user User, createdAt string, updatedAt string, closedAt NullableString, mergedAt NullableString, requestedReviewers []User, requestedTeams []Team, labels []Label, draft bool, head HeadOrBase, base HeadOrBase, remoteData map[string]interface{}, merged bool, numComments int32, numReviewComments int32, numCommits int32, numAdditions int32, numDeletions int32, numChangedFiles int32, ) *PullRequestExtended`

NewPullRequestExtended instantiates a new PullRequestExtended object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestExtendedWithDefaults

`func NewPullRequestExtendedWithDefaults() *PullRequestExtended`

NewPullRequestExtendedWithDefaults instantiates a new PullRequestExtended object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *PullRequestExtended) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PullRequestExtended) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PullRequestExtended) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetId

`func (o *PullRequestExtended) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PullRequestExtended) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PullRequestExtended) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *PullRequestExtended) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PullRequestExtended) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PullRequestExtended) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNumber

`func (o *PullRequestExtended) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PullRequestExtended) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PullRequestExtended) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetState

`func (o *PullRequestExtended) GetState() PRState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PullRequestExtended) GetStateOk() (*PRState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PullRequestExtended) SetState(v PRState)`

SetState sets State field to given value.


### GetUser

`func (o *PullRequestExtended) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PullRequestExtended) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PullRequestExtended) SetUser(v User)`

SetUser sets User field to given value.


### GetCreatedAt

`func (o *PullRequestExtended) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PullRequestExtended) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PullRequestExtended) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PullRequestExtended) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PullRequestExtended) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PullRequestExtended) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetClosedAt

`func (o *PullRequestExtended) GetClosedAt() string`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *PullRequestExtended) GetClosedAtOk() (*string, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *PullRequestExtended) SetClosedAt(v string)`

SetClosedAt sets ClosedAt field to given value.


### SetClosedAtNil

`func (o *PullRequestExtended) SetClosedAtNil(b bool)`

 SetClosedAtNil sets the value for ClosedAt to be an explicit nil

### UnsetClosedAt
`func (o *PullRequestExtended) UnsetClosedAt()`

UnsetClosedAt ensures that no value is present for ClosedAt, not even an explicit nil
### GetMergedAt

`func (o *PullRequestExtended) GetMergedAt() string`

GetMergedAt returns the MergedAt field if non-nil, zero value otherwise.

### GetMergedAtOk

`func (o *PullRequestExtended) GetMergedAtOk() (*string, bool)`

GetMergedAtOk returns a tuple with the MergedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedAt

`func (o *PullRequestExtended) SetMergedAt(v string)`

SetMergedAt sets MergedAt field to given value.


### SetMergedAtNil

`func (o *PullRequestExtended) SetMergedAtNil(b bool)`

 SetMergedAtNil sets the value for MergedAt to be an explicit nil

### UnsetMergedAt
`func (o *PullRequestExtended) UnsetMergedAt()`

UnsetMergedAt ensures that no value is present for MergedAt, not even an explicit nil
### GetRequestedReviewers

`func (o *PullRequestExtended) GetRequestedReviewers() []User`

GetRequestedReviewers returns the RequestedReviewers field if non-nil, zero value otherwise.

### GetRequestedReviewersOk

`func (o *PullRequestExtended) GetRequestedReviewersOk() (*[]User, bool)`

GetRequestedReviewersOk returns a tuple with the RequestedReviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedReviewers

`func (o *PullRequestExtended) SetRequestedReviewers(v []User)`

SetRequestedReviewers sets RequestedReviewers field to given value.


### GetRequestedTeams

`func (o *PullRequestExtended) GetRequestedTeams() []Team`

GetRequestedTeams returns the RequestedTeams field if non-nil, zero value otherwise.

### GetRequestedTeamsOk

`func (o *PullRequestExtended) GetRequestedTeamsOk() (*[]Team, bool)`

GetRequestedTeamsOk returns a tuple with the RequestedTeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTeams

`func (o *PullRequestExtended) SetRequestedTeams(v []Team)`

SetRequestedTeams sets RequestedTeams field to given value.


### GetLabels

`func (o *PullRequestExtended) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PullRequestExtended) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PullRequestExtended) SetLabels(v []Label)`

SetLabels sets Labels field to given value.


### GetDraft

`func (o *PullRequestExtended) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *PullRequestExtended) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *PullRequestExtended) SetDraft(v bool)`

SetDraft sets Draft field to given value.


### GetHead

`func (o *PullRequestExtended) GetHead() HeadOrBase`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *PullRequestExtended) GetHeadOk() (*HeadOrBase, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *PullRequestExtended) SetHead(v HeadOrBase)`

SetHead sets Head field to given value.


### GetBase

`func (o *PullRequestExtended) GetBase() HeadOrBase`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *PullRequestExtended) GetBaseOk() (*HeadOrBase, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *PullRequestExtended) SetBase(v HeadOrBase)`

SetBase sets Base field to given value.


### GetRemoteData

`func (o *PullRequestExtended) GetRemoteData() map[string]interface{}`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *PullRequestExtended) GetRemoteDataOk() (*map[string]interface{}, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *PullRequestExtended) SetRemoteData(v map[string]interface{})`

SetRemoteData sets RemoteData field to given value.


### SetRemoteDataNil

`func (o *PullRequestExtended) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *PullRequestExtended) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil
### GetMerged

`func (o *PullRequestExtended) GetMerged() bool`

GetMerged returns the Merged field if non-nil, zero value otherwise.

### GetMergedOk

`func (o *PullRequestExtended) GetMergedOk() (*bool, bool)`

GetMergedOk returns a tuple with the Merged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerged

`func (o *PullRequestExtended) SetMerged(v bool)`

SetMerged sets Merged field to given value.


### GetNumComments

`func (o *PullRequestExtended) GetNumComments() int32`

GetNumComments returns the NumComments field if non-nil, zero value otherwise.

### GetNumCommentsOk

`func (o *PullRequestExtended) GetNumCommentsOk() (*int32, bool)`

GetNumCommentsOk returns a tuple with the NumComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumComments

`func (o *PullRequestExtended) SetNumComments(v int32)`

SetNumComments sets NumComments field to given value.


### GetNumReviewComments

`func (o *PullRequestExtended) GetNumReviewComments() int32`

GetNumReviewComments returns the NumReviewComments field if non-nil, zero value otherwise.

### GetNumReviewCommentsOk

`func (o *PullRequestExtended) GetNumReviewCommentsOk() (*int32, bool)`

GetNumReviewCommentsOk returns a tuple with the NumReviewComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumReviewComments

`func (o *PullRequestExtended) SetNumReviewComments(v int32)`

SetNumReviewComments sets NumReviewComments field to given value.


### GetNumCommits

`func (o *PullRequestExtended) GetNumCommits() int32`

GetNumCommits returns the NumCommits field if non-nil, zero value otherwise.

### GetNumCommitsOk

`func (o *PullRequestExtended) GetNumCommitsOk() (*int32, bool)`

GetNumCommitsOk returns a tuple with the NumCommits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCommits

`func (o *PullRequestExtended) SetNumCommits(v int32)`

SetNumCommits sets NumCommits field to given value.


### GetNumAdditions

`func (o *PullRequestExtended) GetNumAdditions() int32`

GetNumAdditions returns the NumAdditions field if non-nil, zero value otherwise.

### GetNumAdditionsOk

`func (o *PullRequestExtended) GetNumAdditionsOk() (*int32, bool)`

GetNumAdditionsOk returns a tuple with the NumAdditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAdditions

`func (o *PullRequestExtended) SetNumAdditions(v int32)`

SetNumAdditions sets NumAdditions field to given value.


### GetNumDeletions

`func (o *PullRequestExtended) GetNumDeletions() int32`

GetNumDeletions returns the NumDeletions field if non-nil, zero value otherwise.

### GetNumDeletionsOk

`func (o *PullRequestExtended) GetNumDeletionsOk() (*int32, bool)`

GetNumDeletionsOk returns a tuple with the NumDeletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDeletions

`func (o *PullRequestExtended) SetNumDeletions(v int32)`

SetNumDeletions sets NumDeletions field to given value.


### GetNumChangedFiles

`func (o *PullRequestExtended) GetNumChangedFiles() int32`

GetNumChangedFiles returns the NumChangedFiles field if non-nil, zero value otherwise.

### GetNumChangedFilesOk

`func (o *PullRequestExtended) GetNumChangedFilesOk() (*int32, bool)`

GetNumChangedFilesOk returns a tuple with the NumChangedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumChangedFiles

`func (o *PullRequestExtended) SetNumChangedFiles(v int32)`

SetNumChangedFiles sets NumChangedFiles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



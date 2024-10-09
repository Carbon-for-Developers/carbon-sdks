# Issue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Id** | **int32** |  | 
**Url** | **string** |  | 
**HtmlUrl** | **string** |  | 
**Number** | **int32** |  | 
**User** | [**User**](User.md) |  | 
**Labels** | [**[]Label**](Label.md) |  | 
**State** | [**PRState**](PRState.md) |  | 
**Locked** | **bool** |  | 
**NumComments** | **int32** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**ClosedAt** | **NullableString** |  | 
**Draft** | **bool** |  | 
**PullRequest** | [**NullableIssuePRNullable**](IssuePRNullable.md) |  | 
**Body** | **string** |  | 
**ClosedBy** | [**NullableUserNullable**](UserNullable.md) |  | 

## Methods

### NewIssue

`func NewIssue(title string, id int32, url string, htmlUrl string, number int32, user User, labels []Label, state PRState, locked bool, numComments int32, createdAt string, updatedAt string, closedAt NullableString, draft bool, pullRequest NullableIssuePRNullable, body string, closedBy NullableUserNullable, ) *Issue`

NewIssue instantiates a new Issue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueWithDefaults

`func NewIssueWithDefaults() *Issue`

NewIssueWithDefaults instantiates a new Issue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Issue) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Issue) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Issue) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetId

`func (o *Issue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Issue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Issue) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *Issue) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Issue) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Issue) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *Issue) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *Issue) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *Issue) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetNumber

`func (o *Issue) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Issue) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Issue) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetUser

`func (o *Issue) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Issue) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Issue) SetUser(v User)`

SetUser sets User field to given value.


### GetLabels

`func (o *Issue) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Issue) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Issue) SetLabels(v []Label)`

SetLabels sets Labels field to given value.


### GetState

`func (o *Issue) GetState() PRState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Issue) GetStateOk() (*PRState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Issue) SetState(v PRState)`

SetState sets State field to given value.


### GetLocked

`func (o *Issue) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Issue) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *Issue) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetNumComments

`func (o *Issue) GetNumComments() int32`

GetNumComments returns the NumComments field if non-nil, zero value otherwise.

### GetNumCommentsOk

`func (o *Issue) GetNumCommentsOk() (*int32, bool)`

GetNumCommentsOk returns a tuple with the NumComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumComments

`func (o *Issue) SetNumComments(v int32)`

SetNumComments sets NumComments field to given value.


### GetCreatedAt

`func (o *Issue) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Issue) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Issue) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Issue) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Issue) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Issue) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetClosedAt

`func (o *Issue) GetClosedAt() string`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *Issue) GetClosedAtOk() (*string, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *Issue) SetClosedAt(v string)`

SetClosedAt sets ClosedAt field to given value.


### SetClosedAtNil

`func (o *Issue) SetClosedAtNil(b bool)`

 SetClosedAtNil sets the value for ClosedAt to be an explicit nil

### UnsetClosedAt
`func (o *Issue) UnsetClosedAt()`

UnsetClosedAt ensures that no value is present for ClosedAt, not even an explicit nil
### GetDraft

`func (o *Issue) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *Issue) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *Issue) SetDraft(v bool)`

SetDraft sets Draft field to given value.


### GetPullRequest

`func (o *Issue) GetPullRequest() IssuePRNullable`

GetPullRequest returns the PullRequest field if non-nil, zero value otherwise.

### GetPullRequestOk

`func (o *Issue) GetPullRequestOk() (*IssuePRNullable, bool)`

GetPullRequestOk returns a tuple with the PullRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequest

`func (o *Issue) SetPullRequest(v IssuePRNullable)`

SetPullRequest sets PullRequest field to given value.


### SetPullRequestNil

`func (o *Issue) SetPullRequestNil(b bool)`

 SetPullRequestNil sets the value for PullRequest to be an explicit nil

### UnsetPullRequest
`func (o *Issue) UnsetPullRequest()`

UnsetPullRequest ensures that no value is present for PullRequest, not even an explicit nil
### GetBody

`func (o *Issue) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Issue) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Issue) SetBody(v string)`

SetBody sets Body field to given value.


### GetClosedBy

`func (o *Issue) GetClosedBy() UserNullable`

GetClosedBy returns the ClosedBy field if non-nil, zero value otherwise.

### GetClosedByOk

`func (o *Issue) GetClosedByOk() (*UserNullable, bool)`

GetClosedByOk returns a tuple with the ClosedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedBy

`func (o *Issue) SetClosedBy(v UserNullable)`

SetClosedBy sets ClosedBy field to given value.


### SetClosedByNil

`func (o *Issue) SetClosedByNil(b bool)`

 SetClosedByNil sets the value for ClosedBy to be an explicit nil

### UnsetClosedBy
`func (o *Issue) UnsetClosedBy()`

UnsetClosedBy ensures that no value is present for ClosedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



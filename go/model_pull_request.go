/*
Carbon

Connect external data to LLMs, no matter the source.

API version: 1.0.0
*/

// Code generated by Konfig (https://konfigthis.com); DO NOT EDIT.

package carbon

import (
	"encoding/json"
)

// PullRequest struct for PullRequest
type PullRequest struct {
	Title string `json:"title"`
	Id int32 `json:"id"`
	Url string `json:"url"`
	Number int32 `json:"number"`
	State PRState `json:"state"`
	User User `json:"user"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	ClosedAt NullableString `json:"closed_at"`
	MergedAt NullableString `json:"merged_at"`
	RequestedReviewers []User `json:"requested_reviewers"`
	RequestedTeams []Team `json:"requested_teams"`
	Labels []Label `json:"labels"`
	Draft bool `json:"draft"`
	Head HeadOrBase `json:"head"`
	Base HeadOrBase `json:"base"`
	RemoteData map[string]interface{} `json:"remote_data"`
}

// NewPullRequest instantiates a new PullRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPullRequest(title string, id int32, url string, number int32, state PRState, user User, createdAt string, updatedAt string, closedAt NullableString, mergedAt NullableString, requestedReviewers []User, requestedTeams []Team, labels []Label, draft bool, head HeadOrBase, base HeadOrBase, remoteData map[string]interface{}) *PullRequest {
	this := PullRequest{}
	this.Title = title
	this.Id = id
	this.Url = url
	this.Number = number
	this.State = state
	this.User = user
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.ClosedAt = closedAt
	this.MergedAt = mergedAt
	this.RequestedReviewers = requestedReviewers
	this.RequestedTeams = requestedTeams
	this.Labels = labels
	this.Draft = draft
	this.Head = head
	this.Base = base
	this.RemoteData = remoteData
	return &this
}

// NewPullRequestWithDefaults instantiates a new PullRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPullRequestWithDefaults() *PullRequest {
	this := PullRequest{}
	return &this
}

// GetTitle returns the Title field value
func (o *PullRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *PullRequest) SetTitle(v string) {
	o.Title = v
}

// GetId returns the Id field value
func (o *PullRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PullRequest) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *PullRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *PullRequest) SetUrl(v string) {
	o.Url = v
}

// GetNumber returns the Number field value
func (o *PullRequest) GetNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetNumberOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *PullRequest) SetNumber(v int32) {
	o.Number = v
}

// GetState returns the State field value
func (o *PullRequest) GetState() PRState {
	if o == nil {
		var ret PRState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetStateOk() (*PRState, bool) {
	if o == nil {
    return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *PullRequest) SetState(v PRState) {
	o.State = v
}

// GetUser returns the User field value
func (o *PullRequest) GetUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetUserOk() (*User, bool) {
	if o == nil {
    return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *PullRequest) SetUser(v User) {
	o.User = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PullRequest) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetCreatedAtOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PullRequest) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PullRequest) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PullRequest) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetClosedAt returns the ClosedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PullRequest) GetClosedAt() string {
	if o == nil || o.ClosedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.ClosedAt.Get()
}

// GetClosedAtOk returns a tuple with the ClosedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PullRequest) GetClosedAtOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ClosedAt.Get(), o.ClosedAt.IsSet()
}

// SetClosedAt sets field value
func (o *PullRequest) SetClosedAt(v string) {
	o.ClosedAt.Set(&v)
}

// GetMergedAt returns the MergedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PullRequest) GetMergedAt() string {
	if o == nil || o.MergedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.MergedAt.Get()
}

// GetMergedAtOk returns a tuple with the MergedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PullRequest) GetMergedAtOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.MergedAt.Get(), o.MergedAt.IsSet()
}

// SetMergedAt sets field value
func (o *PullRequest) SetMergedAt(v string) {
	o.MergedAt.Set(&v)
}

// GetRequestedReviewers returns the RequestedReviewers field value
func (o *PullRequest) GetRequestedReviewers() []User {
	if o == nil {
		var ret []User
		return ret
	}

	return o.RequestedReviewers
}

// GetRequestedReviewersOk returns a tuple with the RequestedReviewers field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetRequestedReviewersOk() ([]User, bool) {
	if o == nil {
    return nil, false
	}
	return o.RequestedReviewers, true
}

// SetRequestedReviewers sets field value
func (o *PullRequest) SetRequestedReviewers(v []User) {
	o.RequestedReviewers = v
}

// GetRequestedTeams returns the RequestedTeams field value
func (o *PullRequest) GetRequestedTeams() []Team {
	if o == nil {
		var ret []Team
		return ret
	}

	return o.RequestedTeams
}

// GetRequestedTeamsOk returns a tuple with the RequestedTeams field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetRequestedTeamsOk() ([]Team, bool) {
	if o == nil {
    return nil, false
	}
	return o.RequestedTeams, true
}

// SetRequestedTeams sets field value
func (o *PullRequest) SetRequestedTeams(v []Team) {
	o.RequestedTeams = v
}

// GetLabels returns the Labels field value
func (o *PullRequest) GetLabels() []Label {
	if o == nil {
		var ret []Label
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetLabelsOk() ([]Label, bool) {
	if o == nil {
    return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *PullRequest) SetLabels(v []Label) {
	o.Labels = v
}

// GetDraft returns the Draft field value
func (o *PullRequest) GetDraft() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Draft
}

// GetDraftOk returns a tuple with the Draft field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetDraftOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Draft, true
}

// SetDraft sets field value
func (o *PullRequest) SetDraft(v bool) {
	o.Draft = v
}

// GetHead returns the Head field value
func (o *PullRequest) GetHead() HeadOrBase {
	if o == nil {
		var ret HeadOrBase
		return ret
	}

	return o.Head
}

// GetHeadOk returns a tuple with the Head field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetHeadOk() (*HeadOrBase, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Head, true
}

// SetHead sets field value
func (o *PullRequest) SetHead(v HeadOrBase) {
	o.Head = v
}

// GetBase returns the Base field value
func (o *PullRequest) GetBase() HeadOrBase {
	if o == nil {
		var ret HeadOrBase
		return ret
	}

	return o.Base
}

// GetBaseOk returns a tuple with the Base field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetBaseOk() (*HeadOrBase, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Base, true
}

// SetBase sets field value
func (o *PullRequest) SetBase(v HeadOrBase) {
	o.Base = v
}

// GetRemoteData returns the RemoteData field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *PullRequest) GetRemoteData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.RemoteData
}

// GetRemoteDataOk returns a tuple with the RemoteData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PullRequest) GetRemoteDataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.RemoteData) {
    return map[string]interface{}{}, false
	}
	return o.RemoteData, true
}

// SetRemoteData sets field value
func (o *PullRequest) SetRemoteData(v map[string]interface{}) {
	o.RemoteData = v
}

func (o PullRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["number"] = o.Number
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["closed_at"] = o.ClosedAt.Get()
	}
	if true {
		toSerialize["merged_at"] = o.MergedAt.Get()
	}
	if true {
		toSerialize["requested_reviewers"] = o.RequestedReviewers
	}
	if true {
		toSerialize["requested_teams"] = o.RequestedTeams
	}
	if true {
		toSerialize["labels"] = o.Labels
	}
	if true {
		toSerialize["draft"] = o.Draft
	}
	if true {
		toSerialize["head"] = o.Head
	}
	if true {
		toSerialize["base"] = o.Base
	}
	if o.RemoteData != nil {
		toSerialize["remote_data"] = o.RemoteData
	}
	return json.Marshal(toSerialize)
}

type NullablePullRequest struct {
	value *PullRequest
	isSet bool
}

func (v NullablePullRequest) Get() *PullRequest {
	return v.value
}

func (v *NullablePullRequest) Set(val *PullRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePullRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePullRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePullRequest(val *PullRequest) *NullablePullRequest {
	return &NullablePullRequest{value: val, isSet: true}
}

func (v NullablePullRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePullRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



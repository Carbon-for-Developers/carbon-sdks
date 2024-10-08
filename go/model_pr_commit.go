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

// PRCommit struct for PRCommit
type PRCommit struct {
	Commit Commit `json:"commit"`
	Url string `json:"url"`
	Author NullableUserNullable `json:"author"`
	Committer NullableUserNullable `json:"committer"`
	RemoteData map[string]interface{} `json:"remote_data"`
}

// NewPRCommit instantiates a new PRCommit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPRCommit(commit Commit, url string, author NullableUserNullable, committer NullableUserNullable, remoteData map[string]interface{}) *PRCommit {
	this := PRCommit{}
	this.Commit = commit
	this.Url = url
	this.Author = author
	this.Committer = committer
	this.RemoteData = remoteData
	return &this
}

// NewPRCommitWithDefaults instantiates a new PRCommit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPRCommitWithDefaults() *PRCommit {
	this := PRCommit{}
	return &this
}

// GetCommit returns the Commit field value
func (o *PRCommit) GetCommit() Commit {
	if o == nil {
		var ret Commit
		return ret
	}

	return o.Commit
}

// GetCommitOk returns a tuple with the Commit field value
// and a boolean to check if the value has been set.
func (o *PRCommit) GetCommitOk() (*Commit, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Commit, true
}

// SetCommit sets field value
func (o *PRCommit) SetCommit(v Commit) {
	o.Commit = v
}

// GetUrl returns the Url field value
func (o *PRCommit) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *PRCommit) GetUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *PRCommit) SetUrl(v string) {
	o.Url = v
}

// GetAuthor returns the Author field value
// If the value is explicit nil, the zero value for UserNullable will be returned
func (o *PRCommit) GetAuthor() UserNullable {
	if o == nil || o.Author.Get() == nil {
		var ret UserNullable
		return ret
	}

	return *o.Author.Get()
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PRCommit) GetAuthorOk() (*UserNullable, bool) {
	if o == nil {
    return nil, false
	}
	return o.Author.Get(), o.Author.IsSet()
}

// SetAuthor sets field value
func (o *PRCommit) SetAuthor(v UserNullable) {
	o.Author.Set(&v)
}

// GetCommitter returns the Committer field value
// If the value is explicit nil, the zero value for UserNullable will be returned
func (o *PRCommit) GetCommitter() UserNullable {
	if o == nil || o.Committer.Get() == nil {
		var ret UserNullable
		return ret
	}

	return *o.Committer.Get()
}

// GetCommitterOk returns a tuple with the Committer field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PRCommit) GetCommitterOk() (*UserNullable, bool) {
	if o == nil {
    return nil, false
	}
	return o.Committer.Get(), o.Committer.IsSet()
}

// SetCommitter sets field value
func (o *PRCommit) SetCommitter(v UserNullable) {
	o.Committer.Set(&v)
}

// GetRemoteData returns the RemoteData field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *PRCommit) GetRemoteData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.RemoteData
}

// GetRemoteDataOk returns a tuple with the RemoteData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PRCommit) GetRemoteDataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.RemoteData) {
    return map[string]interface{}{}, false
	}
	return o.RemoteData, true
}

// SetRemoteData sets field value
func (o *PRCommit) SetRemoteData(v map[string]interface{}) {
	o.RemoteData = v
}

func (o PRCommit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["commit"] = o.Commit
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["author"] = o.Author.Get()
	}
	if true {
		toSerialize["committer"] = o.Committer.Get()
	}
	if o.RemoteData != nil {
		toSerialize["remote_data"] = o.RemoteData
	}
	return json.Marshal(toSerialize)
}

type NullablePRCommit struct {
	value *PRCommit
	isSet bool
}

func (v NullablePRCommit) Get() *PRCommit {
	return v.value
}

func (v *NullablePRCommit) Set(val *PRCommit) {
	v.value = val
	v.isSet = true
}

func (v NullablePRCommit) IsSet() bool {
	return v.isSet
}

func (v *NullablePRCommit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePRCommit(val *PRCommit) *NullablePRCommit {
	return &NullablePRCommit{value: val, isSet: true}
}

func (v NullablePRCommit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePRCommit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// Tree struct for Tree
type Tree struct {
	Sha string `json:"sha"`
	Url string `json:"url"`
}

// NewTree instantiates a new Tree object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTree(sha string, url string) *Tree {
	this := Tree{}
	this.Sha = sha
	this.Url = url
	return &this
}

// NewTreeWithDefaults instantiates a new Tree object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTreeWithDefaults() *Tree {
	this := Tree{}
	return &this
}

// GetSha returns the Sha field value
func (o *Tree) GetSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha
}

// GetShaOk returns a tuple with the Sha field value
// and a boolean to check if the value has been set.
func (o *Tree) GetShaOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Sha, true
}

// SetSha sets field value
func (o *Tree) SetSha(v string) {
	o.Sha = v
}

// GetUrl returns the Url field value
func (o *Tree) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Tree) GetUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Tree) SetUrl(v string) {
	o.Url = v
}

func (o Tree) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sha"] = o.Sha
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableTree struct {
	value *Tree
	isSet bool
}

func (v NullableTree) Get() *Tree {
	return v.value
}

func (v *NullableTree) Set(val *Tree) {
	v.value = val
	v.isSet = true
}

func (v NullableTree) IsSet() bool {
	return v.isSet
}

func (v *NullableTree) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTree(val *Tree) *NullableTree {
	return &NullableTree{value: val, isSet: true}
}

func (v NullableTree) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTree) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// PresignedURLResponse struct for PresignedURLResponse
type PresignedURLResponse struct {
	PresignedUrl string `json:"presigned_url"`
}

// NewPresignedURLResponse instantiates a new PresignedURLResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPresignedURLResponse(presignedUrl string) *PresignedURLResponse {
	this := PresignedURLResponse{}
	this.PresignedUrl = presignedUrl
	return &this
}

// NewPresignedURLResponseWithDefaults instantiates a new PresignedURLResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPresignedURLResponseWithDefaults() *PresignedURLResponse {
	this := PresignedURLResponse{}
	return &this
}

// GetPresignedUrl returns the PresignedUrl field value
func (o *PresignedURLResponse) GetPresignedUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PresignedUrl
}

// GetPresignedUrlOk returns a tuple with the PresignedUrl field value
// and a boolean to check if the value has been set.
func (o *PresignedURLResponse) GetPresignedUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PresignedUrl, true
}

// SetPresignedUrl sets field value
func (o *PresignedURLResponse) SetPresignedUrl(v string) {
	o.PresignedUrl = v
}

func (o PresignedURLResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["presigned_url"] = o.PresignedUrl
	}
	return json.Marshal(toSerialize)
}

type NullablePresignedURLResponse struct {
	value *PresignedURLResponse
	isSet bool
}

func (v NullablePresignedURLResponse) Get() *PresignedURLResponse {
	return v.value
}

func (v *NullablePresignedURLResponse) Set(val *PresignedURLResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePresignedURLResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePresignedURLResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePresignedURLResponse(val *PresignedURLResponse) *NullablePresignedURLResponse {
	return &NullablePresignedURLResponse{value: val, isSet: true}
}

func (v NullablePresignedURLResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePresignedURLResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



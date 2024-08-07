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

// SentWebhookPayloadObject struct for SentWebhookPayloadObject
type SentWebhookPayloadObject struct {
	ObjectType *string `json:"object_type,omitempty"`
	ObjectId NullableSentWebhookPayloadObjectObjectId `json:"object_id,omitempty"`
	AdditionalInformation NullableSentWebhookPayloadObjectAdditionalInformation `json:"additional_information,omitempty"`
}

// NewSentWebhookPayloadObject instantiates a new SentWebhookPayloadObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSentWebhookPayloadObject() *SentWebhookPayloadObject {
	this := SentWebhookPayloadObject{}
	return &this
}

// NewSentWebhookPayloadObjectWithDefaults instantiates a new SentWebhookPayloadObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSentWebhookPayloadObjectWithDefaults() *SentWebhookPayloadObject {
	this := SentWebhookPayloadObject{}
	return &this
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *SentWebhookPayloadObject) GetObjectType() string {
	if o == nil || isNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SentWebhookPayloadObject) GetObjectTypeOk() (*string, bool) {
	if o == nil || isNil(o.ObjectType) {
    return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *SentWebhookPayloadObject) HasObjectType() bool {
	if o != nil && !isNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *SentWebhookPayloadObject) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SentWebhookPayloadObject) GetObjectId() SentWebhookPayloadObjectObjectId {
	if o == nil || isNil(o.ObjectId.Get()) {
		var ret SentWebhookPayloadObjectObjectId
		return ret
	}
	return *o.ObjectId.Get()
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SentWebhookPayloadObject) GetObjectIdOk() (*SentWebhookPayloadObjectObjectId, bool) {
	if o == nil {
    return nil, false
	}
	return o.ObjectId.Get(), o.ObjectId.IsSet()
}

// HasObjectId returns a boolean if a field has been set.
func (o *SentWebhookPayloadObject) HasObjectId() bool {
	if o != nil && o.ObjectId.IsSet() {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given NullableSentWebhookPayloadObjectObjectId and assigns it to the ObjectId field.
func (o *SentWebhookPayloadObject) SetObjectId(v SentWebhookPayloadObjectObjectId) {
	o.ObjectId.Set(&v)
}
// SetObjectIdNil sets the value for ObjectId to be an explicit nil
func (o *SentWebhookPayloadObject) SetObjectIdNil() {
	o.ObjectId.Set(nil)
}

// UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil
func (o *SentWebhookPayloadObject) UnsetObjectId() {
	o.ObjectId.Unset()
}

// GetAdditionalInformation returns the AdditionalInformation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SentWebhookPayloadObject) GetAdditionalInformation() SentWebhookPayloadObjectAdditionalInformation {
	if o == nil || isNil(o.AdditionalInformation.Get()) {
		var ret SentWebhookPayloadObjectAdditionalInformation
		return ret
	}
	return *o.AdditionalInformation.Get()
}

// GetAdditionalInformationOk returns a tuple with the AdditionalInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SentWebhookPayloadObject) GetAdditionalInformationOk() (*SentWebhookPayloadObjectAdditionalInformation, bool) {
	if o == nil {
    return nil, false
	}
	return o.AdditionalInformation.Get(), o.AdditionalInformation.IsSet()
}

// HasAdditionalInformation returns a boolean if a field has been set.
func (o *SentWebhookPayloadObject) HasAdditionalInformation() bool {
	if o != nil && o.AdditionalInformation.IsSet() {
		return true
	}

	return false
}

// SetAdditionalInformation gets a reference to the given NullableSentWebhookPayloadObjectAdditionalInformation and assigns it to the AdditionalInformation field.
func (o *SentWebhookPayloadObject) SetAdditionalInformation(v SentWebhookPayloadObjectAdditionalInformation) {
	o.AdditionalInformation.Set(&v)
}
// SetAdditionalInformationNil sets the value for AdditionalInformation to be an explicit nil
func (o *SentWebhookPayloadObject) SetAdditionalInformationNil() {
	o.AdditionalInformation.Set(nil)
}

// UnsetAdditionalInformation ensures that no value is present for AdditionalInformation, not even an explicit nil
func (o *SentWebhookPayloadObject) UnsetAdditionalInformation() {
	o.AdditionalInformation.Unset()
}

func (o SentWebhookPayloadObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ObjectType) {
		toSerialize["object_type"] = o.ObjectType
	}
	if o.ObjectId.IsSet() {
		toSerialize["object_id"] = o.ObjectId.Get()
	}
	if o.AdditionalInformation.IsSet() {
		toSerialize["additional_information"] = o.AdditionalInformation.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSentWebhookPayloadObject struct {
	value *SentWebhookPayloadObject
	isSet bool
}

func (v NullableSentWebhookPayloadObject) Get() *SentWebhookPayloadObject {
	return v.value
}

func (v *NullableSentWebhookPayloadObject) Set(val *SentWebhookPayloadObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSentWebhookPayloadObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSentWebhookPayloadObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSentWebhookPayloadObject(val *SentWebhookPayloadObject) *NullableSentWebhookPayloadObject {
	return &NullableSentWebhookPayloadObject{value: val, isSet: true}
}

func (v NullableSentWebhookPayloadObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSentWebhookPayloadObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



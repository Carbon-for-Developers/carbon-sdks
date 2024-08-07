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

// SentWebhookPayload struct for SentWebhookPayload
type SentWebhookPayload struct {
	WebhookType *string `json:"webhook_type,omitempty"`
	CustomerId *string `json:"customer_id,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
	Object *SentWebhookPayloadObject `json:"object,omitempty"`
}

// NewSentWebhookPayload instantiates a new SentWebhookPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSentWebhookPayload() *SentWebhookPayload {
	this := SentWebhookPayload{}
	return &this
}

// NewSentWebhookPayloadWithDefaults instantiates a new SentWebhookPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSentWebhookPayloadWithDefaults() *SentWebhookPayload {
	this := SentWebhookPayload{}
	return &this
}

// GetWebhookType returns the WebhookType field value if set, zero value otherwise.
func (o *SentWebhookPayload) GetWebhookType() string {
	if o == nil || isNil(o.WebhookType) {
		var ret string
		return ret
	}
	return *o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SentWebhookPayload) GetWebhookTypeOk() (*string, bool) {
	if o == nil || isNil(o.WebhookType) {
    return nil, false
	}
	return o.WebhookType, true
}

// HasWebhookType returns a boolean if a field has been set.
func (o *SentWebhookPayload) HasWebhookType() bool {
	if o != nil && !isNil(o.WebhookType) {
		return true
	}

	return false
}

// SetWebhookType gets a reference to the given string and assigns it to the WebhookType field.
func (o *SentWebhookPayload) SetWebhookType(v string) {
	o.WebhookType = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *SentWebhookPayload) GetCustomerId() string {
	if o == nil || isNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SentWebhookPayload) GetCustomerIdOk() (*string, bool) {
	if o == nil || isNil(o.CustomerId) {
    return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *SentWebhookPayload) HasCustomerId() bool {
	if o != nil && !isNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *SentWebhookPayload) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *SentWebhookPayload) GetTimestamp() string {
	if o == nil || isNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SentWebhookPayload) GetTimestampOk() (*string, bool) {
	if o == nil || isNil(o.Timestamp) {
    return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *SentWebhookPayload) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *SentWebhookPayload) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *SentWebhookPayload) GetObject() SentWebhookPayloadObject {
	if o == nil || isNil(o.Object) {
		var ret SentWebhookPayloadObject
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SentWebhookPayload) GetObjectOk() (*SentWebhookPayloadObject, bool) {
	if o == nil || isNil(o.Object) {
    return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *SentWebhookPayload) HasObject() bool {
	if o != nil && !isNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given SentWebhookPayloadObject and assigns it to the Object field.
func (o *SentWebhookPayload) SetObject(v SentWebhookPayloadObject) {
	o.Object = &v
}

func (o SentWebhookPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WebhookType) {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if !isNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	if !isNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !isNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	return json.Marshal(toSerialize)
}

type NullableSentWebhookPayload struct {
	value *SentWebhookPayload
	isSet bool
}

func (v NullableSentWebhookPayload) Get() *SentWebhookPayload {
	return v.value
}

func (v *NullableSentWebhookPayload) Set(val *SentWebhookPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSentWebhookPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSentWebhookPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSentWebhookPayload(val *SentWebhookPayload) *NullableSentWebhookPayload {
	return &NullableSentWebhookPayload{value: val, isSet: true}
}

func (v NullableSentWebhookPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSentWebhookPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



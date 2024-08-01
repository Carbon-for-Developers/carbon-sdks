/*
Carbon

Connect external data to LLMs, no matter the source.

API version: 1.0.0
*/

// Code generated by Konfig (https://konfigthis.com); DO NOT EDIT.

package carbon

import (
	"encoding/json"
	"fmt"
)

// SentWebhookPayloadObjectAdditionalInformation - struct for SentWebhookPayloadObjectAdditionalInformation
type SentWebhookPayloadObjectAdditionalInformation struct {
	MapmapOfStringinterface *map[string]interface{}
	String *string
}

// map[string]interface{}AsSentWebhookPayloadObjectAdditionalInformation is a convenience function that returns map[string]interface{} wrapped in SentWebhookPayloadObjectAdditionalInformation
func MapmapOfStringinterfaceAsSentWebhookPayloadObjectAdditionalInformation(v *map[string]interface{}) SentWebhookPayloadObjectAdditionalInformation {
	return SentWebhookPayloadObjectAdditionalInformation{
		MapmapOfStringinterface: v,
	}
}

// stringAsSentWebhookPayloadObjectAdditionalInformation is a convenience function that returns string wrapped in SentWebhookPayloadObjectAdditionalInformation
func StringAsSentWebhookPayloadObjectAdditionalInformation(v *string) SentWebhookPayloadObjectAdditionalInformation {
	return SentWebhookPayloadObjectAdditionalInformation{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SentWebhookPayloadObjectAdditionalInformation) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into MapmapOfStringinterface
	err = newStrictDecoder(data).Decode(&dst.MapmapOfStringinterface)
	if err == nil {
		jsonMapmapOfStringinterface, _ := json.Marshal(dst.MapmapOfStringinterface)
		if string(jsonMapmapOfStringinterface) == "{}" { // empty struct
			dst.MapmapOfStringinterface = nil
		} else {
			match++
		}
	} else {
		dst.MapmapOfStringinterface = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MapmapOfStringinterface = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SentWebhookPayloadObjectAdditionalInformation)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SentWebhookPayloadObjectAdditionalInformation)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SentWebhookPayloadObjectAdditionalInformation) MarshalJSON() ([]byte, error) {
	if src.MapmapOfStringinterface != nil {
		return json.Marshal(&src.MapmapOfStringinterface)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SentWebhookPayloadObjectAdditionalInformation) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MapmapOfStringinterface != nil {
		return obj.MapmapOfStringinterface
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableSentWebhookPayloadObjectAdditionalInformation struct {
	value *SentWebhookPayloadObjectAdditionalInformation
	isSet bool
}

func (v NullableSentWebhookPayloadObjectAdditionalInformation) Get() *SentWebhookPayloadObjectAdditionalInformation {
	return v.value
}

func (v *NullableSentWebhookPayloadObjectAdditionalInformation) Set(val *SentWebhookPayloadObjectAdditionalInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableSentWebhookPayloadObjectAdditionalInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableSentWebhookPayloadObjectAdditionalInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSentWebhookPayloadObjectAdditionalInformation(val *SentWebhookPayloadObjectAdditionalInformation) *NullableSentWebhookPayloadObjectAdditionalInformation {
	return &NullableSentWebhookPayloadObjectAdditionalInformation{value: val, isSet: true}
}

func (v NullableSentWebhookPayloadObjectAdditionalInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSentWebhookPayloadObjectAdditionalInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



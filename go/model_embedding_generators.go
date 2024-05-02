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

// EmbeddingGenerators the model 'EmbeddingGenerators'
type EmbeddingGenerators string

// List of EmbeddingGenerators
const (
	EMBEDDINGGENERATORS_OPENAI EmbeddingGenerators = "OPENAI"
	EMBEDDINGGENERATORS_AZURE_OPENAI EmbeddingGenerators = "AZURE_OPENAI"
	EMBEDDINGGENERATORS_AZURE_ADA_LARGE_256 EmbeddingGenerators = "AZURE_ADA_LARGE_256"
	EMBEDDINGGENERATORS_AZURE_ADA_LARGE_1024 EmbeddingGenerators = "AZURE_ADA_LARGE_1024"
	EMBEDDINGGENERATORS_AZURE_ADA_LARGE_3072 EmbeddingGenerators = "AZURE_ADA_LARGE_3072"
	EMBEDDINGGENERATORS_AZURE_ADA_SMALL_512 EmbeddingGenerators = "AZURE_ADA_SMALL_512"
	EMBEDDINGGENERATORS_AZURE_ADA_SMALL_1536 EmbeddingGenerators = "AZURE_ADA_SMALL_1536"
	EMBEDDINGGENERATORS_COHERE_MULTILINGUAL_V3 EmbeddingGenerators = "COHERE_MULTILINGUAL_V3"
	EMBEDDINGGENERATORS_VERTEX_MULTIMODAL EmbeddingGenerators = "VERTEX_MULTIMODAL"
	EMBEDDINGGENERATORS_OPENAI_ADA_LARGE_256 EmbeddingGenerators = "OPENAI_ADA_LARGE_256"
	EMBEDDINGGENERATORS_OPENAI_ADA_LARGE_1024 EmbeddingGenerators = "OPENAI_ADA_LARGE_1024"
	EMBEDDINGGENERATORS_OPENAI_ADA_LARGE_3072 EmbeddingGenerators = "OPENAI_ADA_LARGE_3072"
	EMBEDDINGGENERATORS_OPENAI_ADA_SMALL_512 EmbeddingGenerators = "OPENAI_ADA_SMALL_512"
	EMBEDDINGGENERATORS_OPENAI_ADA_SMALL_1536 EmbeddingGenerators = "OPENAI_ADA_SMALL_1536"
	EMBEDDINGGENERATORS_SOLAR_1_MINI EmbeddingGenerators = "SOLAR_1_MINI"
)

// All allowed values of EmbeddingGenerators enum
var AllowedEmbeddingGeneratorsEnumValues = []EmbeddingGenerators{
	"OPENAI",
	"AZURE_OPENAI",
	"AZURE_ADA_LARGE_256",
	"AZURE_ADA_LARGE_1024",
	"AZURE_ADA_LARGE_3072",
	"AZURE_ADA_SMALL_512",
	"AZURE_ADA_SMALL_1536",
	"COHERE_MULTILINGUAL_V3",
	"VERTEX_MULTIMODAL",
	"OPENAI_ADA_LARGE_256",
	"OPENAI_ADA_LARGE_1024",
	"OPENAI_ADA_LARGE_3072",
	"OPENAI_ADA_SMALL_512",
	"OPENAI_ADA_SMALL_1536",
	"SOLAR_1_MINI",
}

func (v *EmbeddingGenerators) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EmbeddingGenerators(value)
	for _, existing := range AllowedEmbeddingGeneratorsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EmbeddingGenerators", value)
}

// NewEmbeddingGeneratorsFromValue returns a pointer to a valid EmbeddingGenerators
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEmbeddingGeneratorsFromValue(v string) (*EmbeddingGenerators, error) {
	ev := EmbeddingGenerators(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EmbeddingGenerators: valid values are %v", v, AllowedEmbeddingGeneratorsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EmbeddingGenerators) IsValid() bool {
	for _, existing := range AllowedEmbeddingGeneratorsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EmbeddingGenerators value
func (v EmbeddingGenerators) Ptr() *EmbeddingGenerators {
	return &v
}

type NullableEmbeddingGenerators struct {
	value *EmbeddingGenerators
	isSet bool
}

func (v NullableEmbeddingGenerators) Get() *EmbeddingGenerators {
	return v.value
}

func (v *NullableEmbeddingGenerators) Set(val *EmbeddingGenerators) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbeddingGenerators) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbeddingGenerators) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbeddingGenerators(val *EmbeddingGenerators) *NullableEmbeddingGenerators {
	return &NullableEmbeddingGenerators{value: val, isSet: true}
}

func (v NullableEmbeddingGenerators) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbeddingGenerators) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


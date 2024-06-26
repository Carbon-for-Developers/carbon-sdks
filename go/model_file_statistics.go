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

// FileStatistics struct for FileStatistics
type FileStatistics struct {
	FileFormat NullableFileFormatsNullable `json:"file_format"`
	FileSize NullableInt32 `json:"file_size"`
	NumCharacters NullableInt32 `json:"num_characters"`
	NumTokens NullableInt32 `json:"num_tokens"`
	NumEmbeddings NullableInt32 `json:"num_embeddings"`
	MimeType NullableString `json:"mime_type"`
}

// NewFileStatistics instantiates a new FileStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileStatistics(fileFormat NullableFileFormatsNullable, fileSize NullableInt32, numCharacters NullableInt32, numTokens NullableInt32, numEmbeddings NullableInt32, mimeType NullableString) *FileStatistics {
	this := FileStatistics{}
	this.FileFormat = fileFormat
	this.FileSize = fileSize
	this.NumCharacters = numCharacters
	this.NumTokens = numTokens
	this.NumEmbeddings = numEmbeddings
	this.MimeType = mimeType
	return &this
}

// NewFileStatisticsWithDefaults instantiates a new FileStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileStatisticsWithDefaults() *FileStatistics {
	this := FileStatistics{}
	return &this
}

// GetFileFormat returns the FileFormat field value
// If the value is explicit nil, the zero value for FileFormatsNullable will be returned
func (o *FileStatistics) GetFileFormat() FileFormatsNullable {
	if o == nil || o.FileFormat.Get() == nil {
		var ret FileFormatsNullable
		return ret
	}

	return *o.FileFormat.Get()
}

// GetFileFormatOk returns a tuple with the FileFormat field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileStatistics) GetFileFormatOk() (*FileFormatsNullable, bool) {
	if o == nil {
    return nil, false
	}
	return o.FileFormat.Get(), o.FileFormat.IsSet()
}

// SetFileFormat sets field value
func (o *FileStatistics) SetFileFormat(v FileFormatsNullable) {
	o.FileFormat.Set(&v)
}

// GetFileSize returns the FileSize field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *FileStatistics) GetFileSize() int32 {
	if o == nil || o.FileSize.Get() == nil {
		var ret int32
		return ret
	}

	return *o.FileSize.Get()
}

// GetFileSizeOk returns a tuple with the FileSize field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileStatistics) GetFileSizeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.FileSize.Get(), o.FileSize.IsSet()
}

// SetFileSize sets field value
func (o *FileStatistics) SetFileSize(v int32) {
	o.FileSize.Set(&v)
}

// GetNumCharacters returns the NumCharacters field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *FileStatistics) GetNumCharacters() int32 {
	if o == nil || o.NumCharacters.Get() == nil {
		var ret int32
		return ret
	}

	return *o.NumCharacters.Get()
}

// GetNumCharactersOk returns a tuple with the NumCharacters field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileStatistics) GetNumCharactersOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.NumCharacters.Get(), o.NumCharacters.IsSet()
}

// SetNumCharacters sets field value
func (o *FileStatistics) SetNumCharacters(v int32) {
	o.NumCharacters.Set(&v)
}

// GetNumTokens returns the NumTokens field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *FileStatistics) GetNumTokens() int32 {
	if o == nil || o.NumTokens.Get() == nil {
		var ret int32
		return ret
	}

	return *o.NumTokens.Get()
}

// GetNumTokensOk returns a tuple with the NumTokens field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileStatistics) GetNumTokensOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.NumTokens.Get(), o.NumTokens.IsSet()
}

// SetNumTokens sets field value
func (o *FileStatistics) SetNumTokens(v int32) {
	o.NumTokens.Set(&v)
}

// GetNumEmbeddings returns the NumEmbeddings field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *FileStatistics) GetNumEmbeddings() int32 {
	if o == nil || o.NumEmbeddings.Get() == nil {
		var ret int32
		return ret
	}

	return *o.NumEmbeddings.Get()
}

// GetNumEmbeddingsOk returns a tuple with the NumEmbeddings field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileStatistics) GetNumEmbeddingsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.NumEmbeddings.Get(), o.NumEmbeddings.IsSet()
}

// SetNumEmbeddings sets field value
func (o *FileStatistics) SetNumEmbeddings(v int32) {
	o.NumEmbeddings.Set(&v)
}

// GetMimeType returns the MimeType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FileStatistics) GetMimeType() string {
	if o == nil || o.MimeType.Get() == nil {
		var ret string
		return ret
	}

	return *o.MimeType.Get()
}

// GetMimeTypeOk returns a tuple with the MimeType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileStatistics) GetMimeTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.MimeType.Get(), o.MimeType.IsSet()
}

// SetMimeType sets field value
func (o *FileStatistics) SetMimeType(v string) {
	o.MimeType.Set(&v)
}

func (o FileStatistics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["file_format"] = o.FileFormat.Get()
	}
	if true {
		toSerialize["file_size"] = o.FileSize.Get()
	}
	if true {
		toSerialize["num_characters"] = o.NumCharacters.Get()
	}
	if true {
		toSerialize["num_tokens"] = o.NumTokens.Get()
	}
	if true {
		toSerialize["num_embeddings"] = o.NumEmbeddings.Get()
	}
	if true {
		toSerialize["mime_type"] = o.MimeType.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableFileStatistics struct {
	value *FileStatistics
	isSet bool
}

func (v NullableFileStatistics) Get() *FileStatistics {
	return v.value
}

func (v *NullableFileStatistics) Set(val *FileStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableFileStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableFileStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileStatistics(val *FileStatistics) *NullableFileStatistics {
	return &NullableFileStatistics{value: val, isSet: true}
}

func (v NullableFileStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



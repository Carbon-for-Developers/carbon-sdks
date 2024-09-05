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

// FileSyncConfig struct for FileSyncConfig
type FileSyncConfig struct {
	// File types to automatically sync when the data source connects. Only a subset of file types can be          controlled. If not supported, then they will always be synced
	AutoSyncedSourceTypes []HSNFileTypes `json:"auto_synced_source_types,omitempty"`
	// Automatically sync attachments from files where supported. Currently applies to Helpdesk Tickets
	SyncAttachments *bool `json:"sync_attachments,omitempty"`
	// Detect audio language before transcription for audio files
	DetectAudioLanguage *bool `json:"detect_audio_language,omitempty"`
	TranscriptionService NullableTranscriptionServiceNullable `json:"transcription_service,omitempty"`
	// Detect multiple speakers and label segments of speech by speaker for audio files.
	IncludeSpeakerLabels *bool `json:"include_speaker_labels,omitempty"`
	// Whether to split tabular rows into chunks. Currently only valid for CSV, TSV, and XLSX files.
	SplitRows *bool `json:"split_rows,omitempty"`
	// If this flag is enabled, the file will be chunked and stored with Carbon,           but no embeddings will be generated. This overrides the skip_embedding_generation flag.
	GenerateChunksOnly *bool `json:"generate_chunks_only,omitempty"`
}

// NewFileSyncConfig instantiates a new FileSyncConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileSyncConfig() *FileSyncConfig {
	this := FileSyncConfig{}
	var syncAttachments bool = false
	this.SyncAttachments = &syncAttachments
	var detectAudioLanguage bool = false
	this.DetectAudioLanguage = &detectAudioLanguage
	var includeSpeakerLabels bool = false
	this.IncludeSpeakerLabels = &includeSpeakerLabels
	var splitRows bool = false
	this.SplitRows = &splitRows
	var generateChunksOnly bool = false
	this.GenerateChunksOnly = &generateChunksOnly
	return &this
}

// NewFileSyncConfigWithDefaults instantiates a new FileSyncConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileSyncConfigWithDefaults() *FileSyncConfig {
	this := FileSyncConfig{}
	var syncAttachments bool = false
	this.SyncAttachments = &syncAttachments
	var detectAudioLanguage bool = false
	this.DetectAudioLanguage = &detectAudioLanguage
	var includeSpeakerLabels bool = false
	this.IncludeSpeakerLabels = &includeSpeakerLabels
	var splitRows bool = false
	this.SplitRows = &splitRows
	var generateChunksOnly bool = false
	this.GenerateChunksOnly = &generateChunksOnly
	return &this
}

// GetAutoSyncedSourceTypes returns the AutoSyncedSourceTypes field value if set, zero value otherwise.
func (o *FileSyncConfig) GetAutoSyncedSourceTypes() []HSNFileTypes {
	if o == nil || isNil(o.AutoSyncedSourceTypes) {
		var ret []HSNFileTypes
		return ret
	}
	return o.AutoSyncedSourceTypes
}

// GetAutoSyncedSourceTypesOk returns a tuple with the AutoSyncedSourceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSyncConfig) GetAutoSyncedSourceTypesOk() ([]HSNFileTypes, bool) {
	if o == nil || isNil(o.AutoSyncedSourceTypes) {
    return nil, false
	}
	return o.AutoSyncedSourceTypes, true
}

// HasAutoSyncedSourceTypes returns a boolean if a field has been set.
func (o *FileSyncConfig) HasAutoSyncedSourceTypes() bool {
	if o != nil && !isNil(o.AutoSyncedSourceTypes) {
		return true
	}

	return false
}

// SetAutoSyncedSourceTypes gets a reference to the given []HSNFileTypes and assigns it to the AutoSyncedSourceTypes field.
func (o *FileSyncConfig) SetAutoSyncedSourceTypes(v []HSNFileTypes) {
	o.AutoSyncedSourceTypes = v
}

// GetSyncAttachments returns the SyncAttachments field value if set, zero value otherwise.
func (o *FileSyncConfig) GetSyncAttachments() bool {
	if o == nil || isNil(o.SyncAttachments) {
		var ret bool
		return ret
	}
	return *o.SyncAttachments
}

// GetSyncAttachmentsOk returns a tuple with the SyncAttachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSyncConfig) GetSyncAttachmentsOk() (*bool, bool) {
	if o == nil || isNil(o.SyncAttachments) {
    return nil, false
	}
	return o.SyncAttachments, true
}

// HasSyncAttachments returns a boolean if a field has been set.
func (o *FileSyncConfig) HasSyncAttachments() bool {
	if o != nil && !isNil(o.SyncAttachments) {
		return true
	}

	return false
}

// SetSyncAttachments gets a reference to the given bool and assigns it to the SyncAttachments field.
func (o *FileSyncConfig) SetSyncAttachments(v bool) {
	o.SyncAttachments = &v
}

// GetDetectAudioLanguage returns the DetectAudioLanguage field value if set, zero value otherwise.
func (o *FileSyncConfig) GetDetectAudioLanguage() bool {
	if o == nil || isNil(o.DetectAudioLanguage) {
		var ret bool
		return ret
	}
	return *o.DetectAudioLanguage
}

// GetDetectAudioLanguageOk returns a tuple with the DetectAudioLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSyncConfig) GetDetectAudioLanguageOk() (*bool, bool) {
	if o == nil || isNil(o.DetectAudioLanguage) {
    return nil, false
	}
	return o.DetectAudioLanguage, true
}

// HasDetectAudioLanguage returns a boolean if a field has been set.
func (o *FileSyncConfig) HasDetectAudioLanguage() bool {
	if o != nil && !isNil(o.DetectAudioLanguage) {
		return true
	}

	return false
}

// SetDetectAudioLanguage gets a reference to the given bool and assigns it to the DetectAudioLanguage field.
func (o *FileSyncConfig) SetDetectAudioLanguage(v bool) {
	o.DetectAudioLanguage = &v
}

// GetTranscriptionService returns the TranscriptionService field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileSyncConfig) GetTranscriptionService() TranscriptionServiceNullable {
	if o == nil || isNil(o.TranscriptionService.Get()) {
		var ret TranscriptionServiceNullable
		return ret
	}
	return *o.TranscriptionService.Get()
}

// GetTranscriptionServiceOk returns a tuple with the TranscriptionService field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileSyncConfig) GetTranscriptionServiceOk() (*TranscriptionServiceNullable, bool) {
	if o == nil {
    return nil, false
	}
	return o.TranscriptionService.Get(), o.TranscriptionService.IsSet()
}

// HasTranscriptionService returns a boolean if a field has been set.
func (o *FileSyncConfig) HasTranscriptionService() bool {
	if o != nil && o.TranscriptionService.IsSet() {
		return true
	}

	return false
}

// SetTranscriptionService gets a reference to the given NullableTranscriptionServiceNullable and assigns it to the TranscriptionService field.
func (o *FileSyncConfig) SetTranscriptionService(v TranscriptionServiceNullable) {
	o.TranscriptionService.Set(&v)
}
// SetTranscriptionServiceNil sets the value for TranscriptionService to be an explicit nil
func (o *FileSyncConfig) SetTranscriptionServiceNil() {
	o.TranscriptionService.Set(nil)
}

// UnsetTranscriptionService ensures that no value is present for TranscriptionService, not even an explicit nil
func (o *FileSyncConfig) UnsetTranscriptionService() {
	o.TranscriptionService.Unset()
}

// GetIncludeSpeakerLabels returns the IncludeSpeakerLabels field value if set, zero value otherwise.
func (o *FileSyncConfig) GetIncludeSpeakerLabels() bool {
	if o == nil || isNil(o.IncludeSpeakerLabels) {
		var ret bool
		return ret
	}
	return *o.IncludeSpeakerLabels
}

// GetIncludeSpeakerLabelsOk returns a tuple with the IncludeSpeakerLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSyncConfig) GetIncludeSpeakerLabelsOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeSpeakerLabels) {
    return nil, false
	}
	return o.IncludeSpeakerLabels, true
}

// HasIncludeSpeakerLabels returns a boolean if a field has been set.
func (o *FileSyncConfig) HasIncludeSpeakerLabels() bool {
	if o != nil && !isNil(o.IncludeSpeakerLabels) {
		return true
	}

	return false
}

// SetIncludeSpeakerLabels gets a reference to the given bool and assigns it to the IncludeSpeakerLabels field.
func (o *FileSyncConfig) SetIncludeSpeakerLabels(v bool) {
	o.IncludeSpeakerLabels = &v
}

// GetSplitRows returns the SplitRows field value if set, zero value otherwise.
func (o *FileSyncConfig) GetSplitRows() bool {
	if o == nil || isNil(o.SplitRows) {
		var ret bool
		return ret
	}
	return *o.SplitRows
}

// GetSplitRowsOk returns a tuple with the SplitRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSyncConfig) GetSplitRowsOk() (*bool, bool) {
	if o == nil || isNil(o.SplitRows) {
    return nil, false
	}
	return o.SplitRows, true
}

// HasSplitRows returns a boolean if a field has been set.
func (o *FileSyncConfig) HasSplitRows() bool {
	if o != nil && !isNil(o.SplitRows) {
		return true
	}

	return false
}

// SetSplitRows gets a reference to the given bool and assigns it to the SplitRows field.
func (o *FileSyncConfig) SetSplitRows(v bool) {
	o.SplitRows = &v
}

// GetGenerateChunksOnly returns the GenerateChunksOnly field value if set, zero value otherwise.
func (o *FileSyncConfig) GetGenerateChunksOnly() bool {
	if o == nil || isNil(o.GenerateChunksOnly) {
		var ret bool
		return ret
	}
	return *o.GenerateChunksOnly
}

// GetGenerateChunksOnlyOk returns a tuple with the GenerateChunksOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSyncConfig) GetGenerateChunksOnlyOk() (*bool, bool) {
	if o == nil || isNil(o.GenerateChunksOnly) {
    return nil, false
	}
	return o.GenerateChunksOnly, true
}

// HasGenerateChunksOnly returns a boolean if a field has been set.
func (o *FileSyncConfig) HasGenerateChunksOnly() bool {
	if o != nil && !isNil(o.GenerateChunksOnly) {
		return true
	}

	return false
}

// SetGenerateChunksOnly gets a reference to the given bool and assigns it to the GenerateChunksOnly field.
func (o *FileSyncConfig) SetGenerateChunksOnly(v bool) {
	o.GenerateChunksOnly = &v
}

func (o FileSyncConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AutoSyncedSourceTypes) {
		toSerialize["auto_synced_source_types"] = o.AutoSyncedSourceTypes
	}
	if !isNil(o.SyncAttachments) {
		toSerialize["sync_attachments"] = o.SyncAttachments
	}
	if !isNil(o.DetectAudioLanguage) {
		toSerialize["detect_audio_language"] = o.DetectAudioLanguage
	}
	if o.TranscriptionService.IsSet() {
		toSerialize["transcription_service"] = o.TranscriptionService.Get()
	}
	if !isNil(o.IncludeSpeakerLabels) {
		toSerialize["include_speaker_labels"] = o.IncludeSpeakerLabels
	}
	if !isNil(o.SplitRows) {
		toSerialize["split_rows"] = o.SplitRows
	}
	if !isNil(o.GenerateChunksOnly) {
		toSerialize["generate_chunks_only"] = o.GenerateChunksOnly
	}
	return json.Marshal(toSerialize)
}

type NullableFileSyncConfig struct {
	value *FileSyncConfig
	isSet bool
}

func (v NullableFileSyncConfig) Get() *FileSyncConfig {
	return v.value
}

func (v *NullableFileSyncConfig) Set(val *FileSyncConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableFileSyncConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableFileSyncConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileSyncConfig(val *FileSyncConfig) *NullableFileSyncConfig {
	return &NullableFileSyncConfig{value: val, isSet: true}
}

func (v NullableFileSyncConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileSyncConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



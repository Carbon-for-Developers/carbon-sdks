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

// DataSourceConfiguration struct for DataSourceConfiguration
type DataSourceConfiguration struct {
	// List of allowed file formats for the organization. An empty list means all file formats are allowed.
	AllowedFileFormats []string `json:"allowed_file_formats,omitempty"`
	// Used to set the format Google Workspace documents (Docs, Sheets, Slides) are stored as for raw files uploads in S3. The default format is TXT
	GoogleWorkspaceDocsSaveAs NullableString `json:"google_workspace_docs_save_as,omitempty"`
}

// NewDataSourceConfiguration instantiates a new DataSourceConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSourceConfiguration() *DataSourceConfiguration {
	this := DataSourceConfiguration{}
	return &this
}

// NewDataSourceConfigurationWithDefaults instantiates a new DataSourceConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSourceConfigurationWithDefaults() *DataSourceConfiguration {
	this := DataSourceConfiguration{}
	return &this
}

// GetAllowedFileFormats returns the AllowedFileFormats field value if set, zero value otherwise.
func (o *DataSourceConfiguration) GetAllowedFileFormats() []string {
	if o == nil || isNil(o.AllowedFileFormats) {
		var ret []string
		return ret
	}
	return o.AllowedFileFormats
}

// GetAllowedFileFormatsOk returns a tuple with the AllowedFileFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceConfiguration) GetAllowedFileFormatsOk() ([]string, bool) {
	if o == nil || isNil(o.AllowedFileFormats) {
    return nil, false
	}
	return o.AllowedFileFormats, true
}

// HasAllowedFileFormats returns a boolean if a field has been set.
func (o *DataSourceConfiguration) HasAllowedFileFormats() bool {
	if o != nil && !isNil(o.AllowedFileFormats) {
		return true
	}

	return false
}

// SetAllowedFileFormats gets a reference to the given []string and assigns it to the AllowedFileFormats field.
func (o *DataSourceConfiguration) SetAllowedFileFormats(v []string) {
	o.AllowedFileFormats = v
}

// GetGoogleWorkspaceDocsSaveAs returns the GoogleWorkspaceDocsSaveAs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataSourceConfiguration) GetGoogleWorkspaceDocsSaveAs() string {
	if o == nil || isNil(o.GoogleWorkspaceDocsSaveAs.Get()) {
		var ret string
		return ret
	}
	return *o.GoogleWorkspaceDocsSaveAs.Get()
}

// GetGoogleWorkspaceDocsSaveAsOk returns a tuple with the GoogleWorkspaceDocsSaveAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataSourceConfiguration) GetGoogleWorkspaceDocsSaveAsOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.GoogleWorkspaceDocsSaveAs.Get(), o.GoogleWorkspaceDocsSaveAs.IsSet()
}

// HasGoogleWorkspaceDocsSaveAs returns a boolean if a field has been set.
func (o *DataSourceConfiguration) HasGoogleWorkspaceDocsSaveAs() bool {
	if o != nil && o.GoogleWorkspaceDocsSaveAs.IsSet() {
		return true
	}

	return false
}

// SetGoogleWorkspaceDocsSaveAs gets a reference to the given NullableString and assigns it to the GoogleWorkspaceDocsSaveAs field.
func (o *DataSourceConfiguration) SetGoogleWorkspaceDocsSaveAs(v string) {
	o.GoogleWorkspaceDocsSaveAs.Set(&v)
}
// SetGoogleWorkspaceDocsSaveAsNil sets the value for GoogleWorkspaceDocsSaveAs to be an explicit nil
func (o *DataSourceConfiguration) SetGoogleWorkspaceDocsSaveAsNil() {
	o.GoogleWorkspaceDocsSaveAs.Set(nil)
}

// UnsetGoogleWorkspaceDocsSaveAs ensures that no value is present for GoogleWorkspaceDocsSaveAs, not even an explicit nil
func (o *DataSourceConfiguration) UnsetGoogleWorkspaceDocsSaveAs() {
	o.GoogleWorkspaceDocsSaveAs.Unset()
}

func (o DataSourceConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AllowedFileFormats) {
		toSerialize["allowed_file_formats"] = o.AllowedFileFormats
	}
	if o.GoogleWorkspaceDocsSaveAs.IsSet() {
		toSerialize["google_workspace_docs_save_as"] = o.GoogleWorkspaceDocsSaveAs.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDataSourceConfiguration struct {
	value *DataSourceConfiguration
	isSet bool
}

func (v NullableDataSourceConfiguration) Get() *DataSourceConfiguration {
	return v.value
}

func (v *NullableDataSourceConfiguration) Set(val *DataSourceConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourceConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourceConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourceConfiguration(val *DataSourceConfiguration) *NullableDataSourceConfiguration {
	return &NullableDataSourceConfiguration{value: val, isSet: true}
}

func (v NullableDataSourceConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourceConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



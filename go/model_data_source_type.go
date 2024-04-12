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

// DataSourceType the model 'DataSourceType'
type DataSourceType string

// List of DataSourceType
const (
	DATASOURCETYPE_GOOGLE_DRIVE DataSourceType = "GOOGLE_DRIVE"
	DATASOURCETYPE_NOTION DataSourceType = "NOTION"
	DATASOURCETYPE_NOTION_DATABASE DataSourceType = "NOTION_DATABASE"
	DATASOURCETYPE_INTERCOM DataSourceType = "INTERCOM"
	DATASOURCETYPE_DROPBOX DataSourceType = "DROPBOX"
	DATASOURCETYPE_ONEDRIVE DataSourceType = "ONEDRIVE"
	DATASOURCETYPE_SHAREPOINT DataSourceType = "SHAREPOINT"
	DATASOURCETYPE_CONFLUENCE DataSourceType = "CONFLUENCE"
	DATASOURCETYPE_BOX DataSourceType = "BOX"
	DATASOURCETYPE_ZENDESK DataSourceType = "ZENDESK"
	DATASOURCETYPE_ZOTERO DataSourceType = "ZOTERO"
	DATASOURCETYPE_S3 DataSourceType = "S3"
	DATASOURCETYPE_GMAIL DataSourceType = "GMAIL"
	DATASOURCETYPE_OUTLOOK DataSourceType = "OUTLOOK"
	DATASOURCETYPE_TEXT DataSourceType = "TEXT"
	DATASOURCETYPE_CSV DataSourceType = "CSV"
	DATASOURCETYPE_TSV DataSourceType = "TSV"
	DATASOURCETYPE_PDF DataSourceType = "PDF"
	DATASOURCETYPE_DOCX DataSourceType = "DOCX"
	DATASOURCETYPE_PPTX DataSourceType = "PPTX"
	DATASOURCETYPE_XLSX DataSourceType = "XLSX"
	DATASOURCETYPE_MD DataSourceType = "MD"
	DATASOURCETYPE_RTF DataSourceType = "RTF"
	DATASOURCETYPE_JSON DataSourceType = "JSON"
	DATASOURCETYPE_RAW_TEXT DataSourceType = "RAW_TEXT"
	DATASOURCETYPE_WEB_SCRAPE DataSourceType = "WEB_SCRAPE"
	DATASOURCETYPE_RSS_FEED DataSourceType = "RSS_FEED"
	DATASOURCETYPE_FRESHDESK DataSourceType = "FRESHDESK"
	DATASOURCETYPE_GITBOOK DataSourceType = "GITBOOK"
	DATASOURCETYPE_SALESFORCE DataSourceType = "SALESFORCE"
	DATASOURCETYPE_JPG DataSourceType = "JPG"
	DATASOURCETYPE_PNG DataSourceType = "PNG"
	DATASOURCETYPE_JPEG DataSourceType = "JPEG"
	DATASOURCETYPE_MP3 DataSourceType = "MP3"
	DATASOURCETYPE_MP4 DataSourceType = "MP4"
	DATASOURCETYPE_MP2 DataSourceType = "MP2"
	DATASOURCETYPE_AAC DataSourceType = "AAC"
	DATASOURCETYPE_WAV DataSourceType = "WAV"
	DATASOURCETYPE_FLAC DataSourceType = "FLAC"
	DATASOURCETYPE_PCM DataSourceType = "PCM"
	DATASOURCETYPE_M4_A DataSourceType = "M4A"
	DATASOURCETYPE_OGG DataSourceType = "OGG"
	DATASOURCETYPE_OPUS DataSourceType = "OPUS"
	DATASOURCETYPE_WEBM DataSourceType = "WEBM"
)

// All allowed values of DataSourceType enum
var AllowedDataSourceTypeEnumValues = []DataSourceType{
	"GOOGLE_DRIVE",
	"NOTION",
	"NOTION_DATABASE",
	"INTERCOM",
	"DROPBOX",
	"ONEDRIVE",
	"SHAREPOINT",
	"CONFLUENCE",
	"BOX",
	"ZENDESK",
	"ZOTERO",
	"S3",
	"GMAIL",
	"OUTLOOK",
	"TEXT",
	"CSV",
	"TSV",
	"PDF",
	"DOCX",
	"PPTX",
	"XLSX",
	"MD",
	"RTF",
	"JSON",
	"RAW_TEXT",
	"WEB_SCRAPE",
	"RSS_FEED",
	"FRESHDESK",
	"GITBOOK",
	"SALESFORCE",
	"JPG",
	"PNG",
	"JPEG",
	"MP3",
	"MP4",
	"MP2",
	"AAC",
	"WAV",
	"FLAC",
	"PCM",
	"M4A",
	"OGG",
	"OPUS",
	"WEBM",
}

func (v *DataSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataSourceType(value)
	for _, existing := range AllowedDataSourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataSourceType", value)
}

// NewDataSourceTypeFromValue returns a pointer to a valid DataSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataSourceTypeFromValue(v string) (*DataSourceType, error) {
	ev := DataSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataSourceType: valid values are %v", v, AllowedDataSourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataSourceType) IsValid() bool {
	for _, existing := range AllowedDataSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataSourceType value
func (v DataSourceType) Ptr() *DataSourceType {
	return &v
}

type NullableDataSourceType struct {
	value *DataSourceType
	isSet bool
}

func (v NullableDataSourceType) Get() *DataSourceType {
	return v.value
}

func (v *NullableDataSourceType) Set(val *DataSourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourceType(val *DataSourceType) *NullableDataSourceType {
	return &NullableDataSourceType{value: val, isSet: true}
}

func (v NullableDataSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


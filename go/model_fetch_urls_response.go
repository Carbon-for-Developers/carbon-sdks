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

// FetchURLsResponse struct for FetchURLsResponse
type FetchURLsResponse struct {
	Urls []string `json:"urls"`
	HtmlContent string `json:"html_content"`
	ErrorMessage NullableString `json:"error_message"`
}

// NewFetchURLsResponse instantiates a new FetchURLsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchURLsResponse(urls []string, htmlContent string, errorMessage NullableString) *FetchURLsResponse {
	this := FetchURLsResponse{}
	this.Urls = urls
	this.HtmlContent = htmlContent
	this.ErrorMessage = errorMessage
	return &this
}

// NewFetchURLsResponseWithDefaults instantiates a new FetchURLsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchURLsResponseWithDefaults() *FetchURLsResponse {
	this := FetchURLsResponse{}
	return &this
}

// GetUrls returns the Urls field value
func (o *FetchURLsResponse) GetUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *FetchURLsResponse) GetUrlsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Urls, true
}

// SetUrls sets field value
func (o *FetchURLsResponse) SetUrls(v []string) {
	o.Urls = v
}

// GetHtmlContent returns the HtmlContent field value
func (o *FetchURLsResponse) GetHtmlContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlContent
}

// GetHtmlContentOk returns a tuple with the HtmlContent field value
// and a boolean to check if the value has been set.
func (o *FetchURLsResponse) GetHtmlContentOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.HtmlContent, true
}

// SetHtmlContent sets field value
func (o *FetchURLsResponse) SetHtmlContent(v string) {
	o.HtmlContent = v
}

// GetErrorMessage returns the ErrorMessage field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FetchURLsResponse) GetErrorMessage() string {
	if o == nil || o.ErrorMessage.Get() == nil {
		var ret string
		return ret
	}

	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FetchURLsResponse) GetErrorMessageOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// SetErrorMessage sets field value
func (o *FetchURLsResponse) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}

func (o FetchURLsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["urls"] = o.Urls
	}
	if true {
		toSerialize["html_content"] = o.HtmlContent
	}
	if true {
		toSerialize["error_message"] = o.ErrorMessage.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableFetchURLsResponse struct {
	value *FetchURLsResponse
	isSet bool
}

func (v NullableFetchURLsResponse) Get() *FetchURLsResponse {
	return v.value
}

func (v *NullableFetchURLsResponse) Set(val *FetchURLsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchURLsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchURLsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchURLsResponse(val *FetchURLsResponse) *NullableFetchURLsResponse {
	return &NullableFetchURLsResponse{value: val, isSet: true}
}

func (v NullableFetchURLsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchURLsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



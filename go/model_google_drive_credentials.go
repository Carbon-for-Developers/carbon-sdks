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

// GoogleDriveCredentials struct for GoogleDriveCredentials
type GoogleDriveCredentials struct {
	ClientId string `json:"client_id"`
	RedirectUri string `json:"redirect_uri"`
	ClientSecret NullableString `json:"client_secret,omitempty"`
	ApiKey NullableString `json:"api_key,omitempty"`
	FilePickerClientId NullableString `json:"file_picker_client_id,omitempty"`
	FilePickerRedirectUri NullableString `json:"file_picker_redirect_uri,omitempty"`
}

// NewGoogleDriveCredentials instantiates a new GoogleDriveCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleDriveCredentials(clientId string, redirectUri string) *GoogleDriveCredentials {
	this := GoogleDriveCredentials{}
	this.ClientId = clientId
	this.RedirectUri = redirectUri
	return &this
}

// NewGoogleDriveCredentialsWithDefaults instantiates a new GoogleDriveCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleDriveCredentialsWithDefaults() *GoogleDriveCredentials {
	this := GoogleDriveCredentials{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *GoogleDriveCredentials) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *GoogleDriveCredentials) GetClientIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *GoogleDriveCredentials) SetClientId(v string) {
	o.ClientId = v
}

// GetRedirectUri returns the RedirectUri field value
func (o *GoogleDriveCredentials) GetRedirectUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value
// and a boolean to check if the value has been set.
func (o *GoogleDriveCredentials) GetRedirectUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RedirectUri, true
}

// SetRedirectUri sets field value
func (o *GoogleDriveCredentials) SetRedirectUri(v string) {
	o.RedirectUri = v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleDriveCredentials) GetClientSecret() string {
	if o == nil || isNil(o.ClientSecret.Get()) {
		var ret string
		return ret
	}
	return *o.ClientSecret.Get()
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleDriveCredentials) GetClientSecretOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ClientSecret.Get(), o.ClientSecret.IsSet()
}

// HasClientSecret returns a boolean if a field has been set.
func (o *GoogleDriveCredentials) HasClientSecret() bool {
	if o != nil && o.ClientSecret.IsSet() {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given NullableString and assigns it to the ClientSecret field.
func (o *GoogleDriveCredentials) SetClientSecret(v string) {
	o.ClientSecret.Set(&v)
}
// SetClientSecretNil sets the value for ClientSecret to be an explicit nil
func (o *GoogleDriveCredentials) SetClientSecretNil() {
	o.ClientSecret.Set(nil)
}

// UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
func (o *GoogleDriveCredentials) UnsetClientSecret() {
	o.ClientSecret.Unset()
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleDriveCredentials) GetApiKey() string {
	if o == nil || isNil(o.ApiKey.Get()) {
		var ret string
		return ret
	}
	return *o.ApiKey.Get()
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleDriveCredentials) GetApiKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ApiKey.Get(), o.ApiKey.IsSet()
}

// HasApiKey returns a boolean if a field has been set.
func (o *GoogleDriveCredentials) HasApiKey() bool {
	if o != nil && o.ApiKey.IsSet() {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given NullableString and assigns it to the ApiKey field.
func (o *GoogleDriveCredentials) SetApiKey(v string) {
	o.ApiKey.Set(&v)
}
// SetApiKeyNil sets the value for ApiKey to be an explicit nil
func (o *GoogleDriveCredentials) SetApiKeyNil() {
	o.ApiKey.Set(nil)
}

// UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
func (o *GoogleDriveCredentials) UnsetApiKey() {
	o.ApiKey.Unset()
}

// GetFilePickerClientId returns the FilePickerClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleDriveCredentials) GetFilePickerClientId() string {
	if o == nil || isNil(o.FilePickerClientId.Get()) {
		var ret string
		return ret
	}
	return *o.FilePickerClientId.Get()
}

// GetFilePickerClientIdOk returns a tuple with the FilePickerClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleDriveCredentials) GetFilePickerClientIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.FilePickerClientId.Get(), o.FilePickerClientId.IsSet()
}

// HasFilePickerClientId returns a boolean if a field has been set.
func (o *GoogleDriveCredentials) HasFilePickerClientId() bool {
	if o != nil && o.FilePickerClientId.IsSet() {
		return true
	}

	return false
}

// SetFilePickerClientId gets a reference to the given NullableString and assigns it to the FilePickerClientId field.
func (o *GoogleDriveCredentials) SetFilePickerClientId(v string) {
	o.FilePickerClientId.Set(&v)
}
// SetFilePickerClientIdNil sets the value for FilePickerClientId to be an explicit nil
func (o *GoogleDriveCredentials) SetFilePickerClientIdNil() {
	o.FilePickerClientId.Set(nil)
}

// UnsetFilePickerClientId ensures that no value is present for FilePickerClientId, not even an explicit nil
func (o *GoogleDriveCredentials) UnsetFilePickerClientId() {
	o.FilePickerClientId.Unset()
}

// GetFilePickerRedirectUri returns the FilePickerRedirectUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleDriveCredentials) GetFilePickerRedirectUri() string {
	if o == nil || isNil(o.FilePickerRedirectUri.Get()) {
		var ret string
		return ret
	}
	return *o.FilePickerRedirectUri.Get()
}

// GetFilePickerRedirectUriOk returns a tuple with the FilePickerRedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleDriveCredentials) GetFilePickerRedirectUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.FilePickerRedirectUri.Get(), o.FilePickerRedirectUri.IsSet()
}

// HasFilePickerRedirectUri returns a boolean if a field has been set.
func (o *GoogleDriveCredentials) HasFilePickerRedirectUri() bool {
	if o != nil && o.FilePickerRedirectUri.IsSet() {
		return true
	}

	return false
}

// SetFilePickerRedirectUri gets a reference to the given NullableString and assigns it to the FilePickerRedirectUri field.
func (o *GoogleDriveCredentials) SetFilePickerRedirectUri(v string) {
	o.FilePickerRedirectUri.Set(&v)
}
// SetFilePickerRedirectUriNil sets the value for FilePickerRedirectUri to be an explicit nil
func (o *GoogleDriveCredentials) SetFilePickerRedirectUriNil() {
	o.FilePickerRedirectUri.Set(nil)
}

// UnsetFilePickerRedirectUri ensures that no value is present for FilePickerRedirectUri, not even an explicit nil
func (o *GoogleDriveCredentials) UnsetFilePickerRedirectUri() {
	o.FilePickerRedirectUri.Unset()
}

func (o GoogleDriveCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["redirect_uri"] = o.RedirectUri
	}
	if o.ClientSecret.IsSet() {
		toSerialize["client_secret"] = o.ClientSecret.Get()
	}
	if o.ApiKey.IsSet() {
		toSerialize["api_key"] = o.ApiKey.Get()
	}
	if o.FilePickerClientId.IsSet() {
		toSerialize["file_picker_client_id"] = o.FilePickerClientId.Get()
	}
	if o.FilePickerRedirectUri.IsSet() {
		toSerialize["file_picker_redirect_uri"] = o.FilePickerRedirectUri.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableGoogleDriveCredentials struct {
	value *GoogleDriveCredentials
	isSet bool
}

func (v NullableGoogleDriveCredentials) Get() *GoogleDriveCredentials {
	return v.value
}

func (v *NullableGoogleDriveCredentials) Set(val *GoogleDriveCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleDriveCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleDriveCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleDriveCredentials(val *GoogleDriveCredentials) *NullableGoogleDriveCredentials {
	return &NullableGoogleDriveCredentials{value: val, isSet: true}
}

func (v NullableGoogleDriveCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleDriveCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



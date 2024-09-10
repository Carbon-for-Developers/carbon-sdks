# GongAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **interface{}** |  | 
**AccessToken** | **string** |  | 
**RefreshToken** | Pointer to **NullableString** |  | [optional] 
**GongAccountEmail** | **string** |  | 

## Methods

### NewGongAuthentication

`func NewGongAuthentication(source interface{}, accessToken string, gongAccountEmail string, ) *GongAuthentication`

NewGongAuthentication instantiates a new GongAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGongAuthenticationWithDefaults

`func NewGongAuthenticationWithDefaults() *GongAuthentication`

NewGongAuthenticationWithDefaults instantiates a new GongAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *GongAuthentication) GetSource() interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GongAuthentication) GetSourceOk() (*interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GongAuthentication) SetSource(v interface{})`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *GongAuthentication) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *GongAuthentication) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetAccessToken

`func (o *GongAuthentication) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *GongAuthentication) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *GongAuthentication) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetRefreshToken

`func (o *GongAuthentication) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *GongAuthentication) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *GongAuthentication) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *GongAuthentication) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### SetRefreshTokenNil

`func (o *GongAuthentication) SetRefreshTokenNil(b bool)`

 SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil

### UnsetRefreshToken
`func (o *GongAuthentication) UnsetRefreshToken()`

UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
### GetGongAccountEmail

`func (o *GongAuthentication) GetGongAccountEmail() string`

GetGongAccountEmail returns the GongAccountEmail field if non-nil, zero value otherwise.

### GetGongAccountEmailOk

`func (o *GongAuthentication) GetGongAccountEmailOk() (*string, bool)`

GetGongAccountEmailOk returns a tuple with the GongAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGongAccountEmail

`func (o *GongAuthentication) SetGongAccountEmail(v string)`

SetGongAccountEmail sets GongAccountEmail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Email

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**EmailType** | **NullableString** |  | 

## Methods

### NewEmail

`func NewEmail(email string, emailType NullableString, ) *Email`

NewEmail instantiates a new Email object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailWithDefaults

`func NewEmailWithDefaults() *Email`

NewEmailWithDefaults instantiates a new Email object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *Email) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Email) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Email) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEmailType

`func (o *Email) GetEmailType() string`

GetEmailType returns the EmailType field if non-nil, zero value otherwise.

### GetEmailTypeOk

`func (o *Email) GetEmailTypeOk() (*string, bool)`

GetEmailTypeOk returns a tuple with the EmailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailType

`func (o *Email) SetEmailType(v string)`

SetEmailType sets EmailType field to given value.


### SetEmailTypeNil

`func (o *Email) SetEmailTypeNil(b bool)`

 SetEmailTypeNil sets the value for EmailType to be an explicit nil

### UnsetEmailType
`func (o *Email) UnsetEmailType()`

UnsetEmailType ensures that no value is present for EmailType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



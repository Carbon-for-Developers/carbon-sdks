# FreskdeskAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **interface{}** |  | 
**Domain** | **string** |  | 
**ApiKey** | **string** |  | 

## Methods

### NewFreskdeskAuthentication

`func NewFreskdeskAuthentication(source interface{}, domain string, apiKey string, ) *FreskdeskAuthentication`

NewFreskdeskAuthentication instantiates a new FreskdeskAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreskdeskAuthenticationWithDefaults

`func NewFreskdeskAuthenticationWithDefaults() *FreskdeskAuthentication`

NewFreskdeskAuthenticationWithDefaults instantiates a new FreskdeskAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *FreskdeskAuthentication) GetSource() interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FreskdeskAuthentication) GetSourceOk() (*interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FreskdeskAuthentication) SetSource(v interface{})`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *FreskdeskAuthentication) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *FreskdeskAuthentication) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetDomain

`func (o *FreskdeskAuthentication) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *FreskdeskAuthentication) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *FreskdeskAuthentication) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetApiKey

`func (o *FreskdeskAuthentication) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *FreskdeskAuthentication) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *FreskdeskAuthentication) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



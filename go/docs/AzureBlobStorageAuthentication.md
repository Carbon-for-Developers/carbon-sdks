# AzureBlobStorageAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **interface{}** |  | 
**AccountName** | **string** |  | 
**AccountKey** | **string** |  | 

## Methods

### NewAzureBlobStorageAuthentication

`func NewAzureBlobStorageAuthentication(source interface{}, accountName string, accountKey string, ) *AzureBlobStorageAuthentication`

NewAzureBlobStorageAuthentication instantiates a new AzureBlobStorageAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureBlobStorageAuthenticationWithDefaults

`func NewAzureBlobStorageAuthenticationWithDefaults() *AzureBlobStorageAuthentication`

NewAzureBlobStorageAuthenticationWithDefaults instantiates a new AzureBlobStorageAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *AzureBlobStorageAuthentication) GetSource() interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AzureBlobStorageAuthentication) GetSourceOk() (*interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AzureBlobStorageAuthentication) SetSource(v interface{})`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *AzureBlobStorageAuthentication) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *AzureBlobStorageAuthentication) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetAccountName

`func (o *AzureBlobStorageAuthentication) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *AzureBlobStorageAuthentication) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *AzureBlobStorageAuthentication) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetAccountKey

`func (o *AzureBlobStorageAuthentication) GetAccountKey() string`

GetAccountKey returns the AccountKey field if non-nil, zero value otherwise.

### GetAccountKeyOk

`func (o *AzureBlobStorageAuthentication) GetAccountKeyOk() (*string, bool)`

GetAccountKeyOk returns a tuple with the AccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKey

`func (o *AzureBlobStorageAuthentication) SetAccountKey(v string)`

SetAccountKey sets AccountKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



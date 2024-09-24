# AzureBlobAuthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | **string** |  | 
**AccountKey** | **string** |  | 
**SyncSourceItems** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewAzureBlobAuthRequest

`func NewAzureBlobAuthRequest(accountName string, accountKey string, ) *AzureBlobAuthRequest`

NewAzureBlobAuthRequest instantiates a new AzureBlobAuthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureBlobAuthRequestWithDefaults

`func NewAzureBlobAuthRequestWithDefaults() *AzureBlobAuthRequest`

NewAzureBlobAuthRequestWithDefaults instantiates a new AzureBlobAuthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *AzureBlobAuthRequest) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *AzureBlobAuthRequest) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *AzureBlobAuthRequest) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetAccountKey

`func (o *AzureBlobAuthRequest) GetAccountKey() string`

GetAccountKey returns the AccountKey field if non-nil, zero value otherwise.

### GetAccountKeyOk

`func (o *AzureBlobAuthRequest) GetAccountKeyOk() (*string, bool)`

GetAccountKeyOk returns a tuple with the AccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKey

`func (o *AzureBlobAuthRequest) SetAccountKey(v string)`

SetAccountKey sets AccountKey field to given value.


### GetSyncSourceItems

`func (o *AzureBlobAuthRequest) GetSyncSourceItems() bool`

GetSyncSourceItems returns the SyncSourceItems field if non-nil, zero value otherwise.

### GetSyncSourceItemsOk

`func (o *AzureBlobAuthRequest) GetSyncSourceItemsOk() (*bool, bool)`

GetSyncSourceItemsOk returns a tuple with the SyncSourceItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncSourceItems

`func (o *AzureBlobAuthRequest) SetSyncSourceItems(v bool)`

SetSyncSourceItems sets SyncSourceItems field to given value.

### HasSyncSourceItems

`func (o *AzureBlobAuthRequest) HasSyncSourceItems() bool`

HasSyncSourceItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



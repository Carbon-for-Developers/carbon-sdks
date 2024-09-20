# ConnectDataSourceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authentication** | [**OANSCZGFB**](OANSCZGFB.md) |  | 
**SyncOptions** | Pointer to [**SyncOptions**](SyncOptions.md) |  | [optional] 

## Methods

### NewConnectDataSourceInput

`func NewConnectDataSourceInput(authentication OANSCZGFB, ) *ConnectDataSourceInput`

NewConnectDataSourceInput instantiates a new ConnectDataSourceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectDataSourceInputWithDefaults

`func NewConnectDataSourceInputWithDefaults() *ConnectDataSourceInput`

NewConnectDataSourceInputWithDefaults instantiates a new ConnectDataSourceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthentication

`func (o *ConnectDataSourceInput) GetAuthentication() OANSCZGFB`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *ConnectDataSourceInput) GetAuthenticationOk() (*OANSCZGFB, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *ConnectDataSourceInput) SetAuthentication(v OANSCZGFB)`

SetAuthentication sets Authentication field to given value.


### GetSyncOptions

`func (o *ConnectDataSourceInput) GetSyncOptions() SyncOptions`

GetSyncOptions returns the SyncOptions field if non-nil, zero value otherwise.

### GetSyncOptionsOk

`func (o *ConnectDataSourceInput) GetSyncOptionsOk() (*SyncOptions, bool)`

GetSyncOptionsOk returns a tuple with the SyncOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncOptions

`func (o *ConnectDataSourceInput) SetSyncOptions(v SyncOptions)`

SetSyncOptions sets SyncOptions field to given value.

### HasSyncOptions

`func (o *ConnectDataSourceInput) HasSyncOptions() bool`

HasSyncOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



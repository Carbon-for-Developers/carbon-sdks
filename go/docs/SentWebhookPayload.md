# SentWebhookPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookType** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**Object** | Pointer to [**SentWebhookPayloadObject**](SentWebhookPayloadObject.md) |  | [optional] 

## Methods

### NewSentWebhookPayload

`func NewSentWebhookPayload() *SentWebhookPayload`

NewSentWebhookPayload instantiates a new SentWebhookPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSentWebhookPayloadWithDefaults

`func NewSentWebhookPayloadWithDefaults() *SentWebhookPayload`

NewSentWebhookPayloadWithDefaults instantiates a new SentWebhookPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookType

`func (o *SentWebhookPayload) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *SentWebhookPayload) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *SentWebhookPayload) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.

### HasWebhookType

`func (o *SentWebhookPayload) HasWebhookType() bool`

HasWebhookType returns a boolean if a field has been set.

### GetCustomerId

`func (o *SentWebhookPayload) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SentWebhookPayload) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SentWebhookPayload) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *SentWebhookPayload) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetTimestamp

`func (o *SentWebhookPayload) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SentWebhookPayload) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SentWebhookPayload) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SentWebhookPayload) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetObject

`func (o *SentWebhookPayload) GetObject() SentWebhookPayloadObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SentWebhookPayload) GetObjectOk() (*SentWebhookPayloadObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SentWebhookPayload) SetObject(v SentWebhookPayloadObject)`

SetObject sets Object field to given value.

### HasObject

`func (o *SentWebhookPayload) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SentWebhookPayloadObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to [**SentWebhookPayloadObjectObjectId**](SentWebhookPayloadObjectObjectId.md) |  | [optional] 
**AdditionalInformation** | Pointer to [**SentWebhookPayloadObjectAdditionalInformation**](SentWebhookPayloadObjectAdditionalInformation.md) |  | [optional] 

## Methods

### NewSentWebhookPayloadObject

`func NewSentWebhookPayloadObject() *SentWebhookPayloadObject`

NewSentWebhookPayloadObject instantiates a new SentWebhookPayloadObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSentWebhookPayloadObjectWithDefaults

`func NewSentWebhookPayloadObjectWithDefaults() *SentWebhookPayloadObject`

NewSentWebhookPayloadObjectWithDefaults instantiates a new SentWebhookPayloadObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *SentWebhookPayloadObject) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SentWebhookPayloadObject) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SentWebhookPayloadObject) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *SentWebhookPayloadObject) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectId

`func (o *SentWebhookPayloadObject) GetObjectId() SentWebhookPayloadObjectObjectId`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *SentWebhookPayloadObject) GetObjectIdOk() (*SentWebhookPayloadObjectObjectId, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *SentWebhookPayloadObject) SetObjectId(v SentWebhookPayloadObjectObjectId)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *SentWebhookPayloadObject) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *SentWebhookPayloadObject) GetAdditionalInformation() SentWebhookPayloadObjectAdditionalInformation`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *SentWebhookPayloadObject) GetAdditionalInformationOk() (*SentWebhookPayloadObjectAdditionalInformation, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *SentWebhookPayloadObject) SetAdditionalInformation(v SentWebhookPayloadObjectAdditionalInformation)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *SentWebhookPayloadObject) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



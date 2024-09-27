# LeadsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**NextCursor** | **NullableString** |  | 
**Data** | [**[]Lead**](Lead.md) |  | 

## Methods

### NewLeadsResponse

`func NewLeadsResponse(count int32, nextCursor NullableString, data []Lead, ) *LeadsResponse`

NewLeadsResponse instantiates a new LeadsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeadsResponseWithDefaults

`func NewLeadsResponseWithDefaults() *LeadsResponse`

NewLeadsResponseWithDefaults instantiates a new LeadsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *LeadsResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *LeadsResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *LeadsResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNextCursor

`func (o *LeadsResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *LeadsResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *LeadsResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *LeadsResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *LeadsResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetData

`func (o *LeadsResponse) GetData() []Lead`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LeadsResponse) GetDataOk() (*[]Lead, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LeadsResponse) SetData(v []Lead)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



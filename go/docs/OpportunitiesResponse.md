# OpportunitiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**NextCursor** | **NullableString** |  | 
**Data** | [**[]Opportunity**](Opportunity.md) |  | 

## Methods

### NewOpportunitiesResponse

`func NewOpportunitiesResponse(count int32, nextCursor NullableString, data []Opportunity, ) *OpportunitiesResponse`

NewOpportunitiesResponse instantiates a new OpportunitiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpportunitiesResponseWithDefaults

`func NewOpportunitiesResponseWithDefaults() *OpportunitiesResponse`

NewOpportunitiesResponseWithDefaults instantiates a new OpportunitiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *OpportunitiesResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *OpportunitiesResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *OpportunitiesResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNextCursor

`func (o *OpportunitiesResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *OpportunitiesResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *OpportunitiesResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *OpportunitiesResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *OpportunitiesResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetData

`func (o *OpportunitiesResponse) GetData() []Opportunity`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OpportunitiesResponse) GetDataOk() (*[]Opportunity, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OpportunitiesResponse) SetData(v []Opportunity)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



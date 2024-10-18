# ListWhiteLabelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**OrderBy** | Pointer to [**WhiteLabelOrderByColumns**](WhiteLabelOrderByColumns.md) |  | [optional] 
**OrderDir** | Pointer to [**OrderDir**](OrderDir.md) |  | [optional] 
**Filters** | Pointer to [**WhiteLabelFilters**](WhiteLabelFilters.md) |  | [optional] 

## Methods

### NewListWhiteLabelRequest

`func NewListWhiteLabelRequest() *ListWhiteLabelRequest`

NewListWhiteLabelRequest instantiates a new ListWhiteLabelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWhiteLabelRequestWithDefaults

`func NewListWhiteLabelRequestWithDefaults() *ListWhiteLabelRequest`

NewListWhiteLabelRequestWithDefaults instantiates a new ListWhiteLabelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *ListWhiteLabelRequest) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListWhiteLabelRequest) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListWhiteLabelRequest) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListWhiteLabelRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetOrderBy

`func (o *ListWhiteLabelRequest) GetOrderBy() WhiteLabelOrderByColumns`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *ListWhiteLabelRequest) GetOrderByOk() (*WhiteLabelOrderByColumns, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *ListWhiteLabelRequest) SetOrderBy(v WhiteLabelOrderByColumns)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *ListWhiteLabelRequest) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetOrderDir

`func (o *ListWhiteLabelRequest) GetOrderDir() OrderDir`

GetOrderDir returns the OrderDir field if non-nil, zero value otherwise.

### GetOrderDirOk

`func (o *ListWhiteLabelRequest) GetOrderDirOk() (*OrderDir, bool)`

GetOrderDirOk returns a tuple with the OrderDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDir

`func (o *ListWhiteLabelRequest) SetOrderDir(v OrderDir)`

SetOrderDir sets OrderDir field to given value.

### HasOrderDir

`func (o *ListWhiteLabelRequest) HasOrderDir() bool`

HasOrderDir returns a boolean if a field has been set.

### GetFilters

`func (o *ListWhiteLabelRequest) GetFilters() WhiteLabelFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ListWhiteLabelRequest) GetFiltersOk() (*WhiteLabelFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ListWhiteLabelRequest) SetFilters(v WhiteLabelFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ListWhiteLabelRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# EmbeddingsAndChunksQueryInputV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**OrderBy** | Pointer to [**OrganizationUserFilesToSyncOrderByTypes**](OrganizationUserFilesToSyncOrderByTypes.md) |  | [optional] 
**OrderDir** | Pointer to [**OrderDir**](OrderDir.md) |  | [optional] 
**Filters** | [**OrganizationUserFilesToSyncFilters**](OrganizationUserFilesToSyncFilters.md) |  | 
**IncludeVectors** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewEmbeddingsAndChunksQueryInputV2

`func NewEmbeddingsAndChunksQueryInputV2(filters OrganizationUserFilesToSyncFilters, ) *EmbeddingsAndChunksQueryInputV2`

NewEmbeddingsAndChunksQueryInputV2 instantiates a new EmbeddingsAndChunksQueryInputV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddingsAndChunksQueryInputV2WithDefaults

`func NewEmbeddingsAndChunksQueryInputV2WithDefaults() *EmbeddingsAndChunksQueryInputV2`

NewEmbeddingsAndChunksQueryInputV2WithDefaults instantiates a new EmbeddingsAndChunksQueryInputV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *EmbeddingsAndChunksQueryInputV2) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *EmbeddingsAndChunksQueryInputV2) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *EmbeddingsAndChunksQueryInputV2) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *EmbeddingsAndChunksQueryInputV2) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetOrderBy

`func (o *EmbeddingsAndChunksQueryInputV2) GetOrderBy() OrganizationUserFilesToSyncOrderByTypes`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *EmbeddingsAndChunksQueryInputV2) GetOrderByOk() (*OrganizationUserFilesToSyncOrderByTypes, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *EmbeddingsAndChunksQueryInputV2) SetOrderBy(v OrganizationUserFilesToSyncOrderByTypes)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *EmbeddingsAndChunksQueryInputV2) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetOrderDir

`func (o *EmbeddingsAndChunksQueryInputV2) GetOrderDir() OrderDir`

GetOrderDir returns the OrderDir field if non-nil, zero value otherwise.

### GetOrderDirOk

`func (o *EmbeddingsAndChunksQueryInputV2) GetOrderDirOk() (*OrderDir, bool)`

GetOrderDirOk returns a tuple with the OrderDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDir

`func (o *EmbeddingsAndChunksQueryInputV2) SetOrderDir(v OrderDir)`

SetOrderDir sets OrderDir field to given value.

### HasOrderDir

`func (o *EmbeddingsAndChunksQueryInputV2) HasOrderDir() bool`

HasOrderDir returns a boolean if a field has been set.

### GetFilters

`func (o *EmbeddingsAndChunksQueryInputV2) GetFilters() OrganizationUserFilesToSyncFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *EmbeddingsAndChunksQueryInputV2) GetFiltersOk() (*OrganizationUserFilesToSyncFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *EmbeddingsAndChunksQueryInputV2) SetFilters(v OrganizationUserFilesToSyncFilters)`

SetFilters sets Filters field to given value.


### GetIncludeVectors

`func (o *EmbeddingsAndChunksQueryInputV2) GetIncludeVectors() bool`

GetIncludeVectors returns the IncludeVectors field if non-nil, zero value otherwise.

### GetIncludeVectorsOk

`func (o *EmbeddingsAndChunksQueryInputV2) GetIncludeVectorsOk() (*bool, bool)`

GetIncludeVectorsOk returns a tuple with the IncludeVectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeVectors

`func (o *EmbeddingsAndChunksQueryInputV2) SetIncludeVectors(v bool)`

SetIncludeVectors sets IncludeVectors field to given value.

### HasIncludeVectors

`func (o *EmbeddingsAndChunksQueryInputV2) HasIncludeVectors() bool`

HasIncludeVectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



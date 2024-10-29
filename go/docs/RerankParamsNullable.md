# RerankParamsNullable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | The name of the model to use for reranking. Options: [&#39;COHERE_RERANK_MULTILINGUAL_V3&#39;, &#39;JINA_RERANKER_MULTILINGUAL_V2&#39;, &#39;PONGO_RERANKER&#39;]. | 

## Methods

### NewRerankParamsNullable

`func NewRerankParamsNullable(model string, ) *RerankParamsNullable`

NewRerankParamsNullable instantiates a new RerankParamsNullable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRerankParamsNullableWithDefaults

`func NewRerankParamsNullableWithDefaults() *RerankParamsNullable`

NewRerankParamsNullableWithDefaults instantiates a new RerankParamsNullable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *RerankParamsNullable) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *RerankParamsNullable) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *RerankParamsNullable) SetModel(v string)`

SetModel sets Model field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



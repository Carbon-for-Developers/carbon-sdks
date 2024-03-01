/*
Carbon

Connect external data to LLMs, no matter the source.

API version: 1.0.0
*/

// Code generated by Konfig (https://konfigthis.com); DO NOT EDIT.

package carbon

import (
	"encoding/json"
)

// WebscrapeRequest struct for WebscrapeRequest
type WebscrapeRequest struct {
	Tags map[string]Tags1 `json:"tags,omitempty"`
	Url string `json:"url"`
	RecursionDepth NullableInt32 `json:"recursion_depth,omitempty"`
	MaxPagesToScrape NullableInt32 `json:"max_pages_to_scrape,omitempty"`
	ChunkSize NullableInt32 `json:"chunk_size,omitempty"`
	ChunkOverlap NullableInt32 `json:"chunk_overlap,omitempty"`
	SkipEmbeddingGeneration NullableBool `json:"skip_embedding_generation,omitempty"`
	EnableAutoSync NullableBool `json:"enable_auto_sync,omitempty"`
	GenerateSparseVectors NullableBool `json:"generate_sparse_vectors,omitempty"`
	PrependFilenameToChunks NullableBool `json:"prepend_filename_to_chunks,omitempty"`
	HtmlTagsToSkip []string `json:"html_tags_to_skip,omitempty"`
	CssClassesToSkip []string `json:"css_classes_to_skip,omitempty"`
	CssSelectorsToSkip []string `json:"css_selectors_to_skip,omitempty"`
	EmbeddingModel *EmbeddingGenerators `json:"embedding_model,omitempty"`
}

// NewWebscrapeRequest instantiates a new WebscrapeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebscrapeRequest(url string) *WebscrapeRequest {
	this := WebscrapeRequest{}
	this.Url = url
	var recursionDepth int32 = 3
	this.RecursionDepth = *NewNullableInt32(&recursionDepth)
	var maxPagesToScrape int32 = 100
	this.MaxPagesToScrape = *NewNullableInt32(&maxPagesToScrape)
	var chunkSize int32 = 1500
	this.ChunkSize = *NewNullableInt32(&chunkSize)
	var chunkOverlap int32 = 20
	this.ChunkOverlap = *NewNullableInt32(&chunkOverlap)
	var skipEmbeddingGeneration bool = false
	this.SkipEmbeddingGeneration = *NewNullableBool(&skipEmbeddingGeneration)
	var enableAutoSync bool = false
	this.EnableAutoSync = *NewNullableBool(&enableAutoSync)
	var generateSparseVectors bool = false
	this.GenerateSparseVectors = *NewNullableBool(&generateSparseVectors)
	var prependFilenameToChunks bool = false
	this.PrependFilenameToChunks = *NewNullableBool(&prependFilenameToChunks)
	return &this
}

// NewWebscrapeRequestWithDefaults instantiates a new WebscrapeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebscrapeRequestWithDefaults() *WebscrapeRequest {
	this := WebscrapeRequest{}
	var recursionDepth int32 = 3
	this.RecursionDepth = *NewNullableInt32(&recursionDepth)
	var maxPagesToScrape int32 = 100
	this.MaxPagesToScrape = *NewNullableInt32(&maxPagesToScrape)
	var chunkSize int32 = 1500
	this.ChunkSize = *NewNullableInt32(&chunkSize)
	var chunkOverlap int32 = 20
	this.ChunkOverlap = *NewNullableInt32(&chunkOverlap)
	var skipEmbeddingGeneration bool = false
	this.SkipEmbeddingGeneration = *NewNullableBool(&skipEmbeddingGeneration)
	var enableAutoSync bool = false
	this.EnableAutoSync = *NewNullableBool(&enableAutoSync)
	var generateSparseVectors bool = false
	this.GenerateSparseVectors = *NewNullableBool(&generateSparseVectors)
	var prependFilenameToChunks bool = false
	this.PrependFilenameToChunks = *NewNullableBool(&prependFilenameToChunks)
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebscrapeRequest) GetTags() map[string]Tags1 {
	if o == nil {
		var ret map[string]Tags1
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebscrapeRequest) GetTagsOk() (*map[string]Tags1, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WebscrapeRequest) HasTags() bool {
	if o != nil && isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]Tags1 and assigns it to the Tags field.
func (o *WebscrapeRequest) SetTags(v map[string]Tags1) {
	o.Tags = v
}

// GetUrl returns the Url field value
func (o *WebscrapeRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WebscrapeRequest) GetUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WebscrapeRequest) SetUrl(v string) {
	o.Url = v
}

// GetRecursionDepth returns the RecursionDepth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebscrapeRequest) GetRecursionDepth() int32 {
	if o == nil || isNil(o.RecursionDepth.Get()) {
		var ret int32
		return ret
	}
	return *o.RecursionDepth.Get()
}

// GetRecursionDepthOk returns a tuple with the RecursionDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebscrapeRequest) GetRecursionDepthOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.RecursionDepth.Get(), o.RecursionDepth.IsSet()
}

// HasRecursionDepth returns a boolean if a field has been set.
func (o *WebscrapeRequest) HasRecursionDepth() bool {
	if o != nil && o.RecursionDepth.IsSet() {
		return true
	}

	return false
}

// SetRecursionDepth gets a reference to the given NullableInt32 and assigns it to the RecursionDepth field.
func (o *WebscrapeRequest) SetRecursionDepth(v int32) {
	o.RecursionDepth.Set(&v)
}
// SetRecursionDepthNil sets the value for RecursionDepth to be an explicit nil
func (o *WebscrapeRequest) SetRecursionDepthNil() {
	o.RecursionDepth.Set(nil)
}

// UnsetRecursionDepth ensures that no value is present for RecursionDepth, not even an explicit nil
func (o *WebscrapeRequest) UnsetRecursionDepth() {
	o.RecursionDepth.Unset()
}

// GetMaxPagesToScrape returns the MaxPagesToScrape field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebscrapeRequest) GetMaxPagesToScrape() int32 {
	if o == nil || isNil(o.MaxPagesToScrape.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxPagesToScrape.Get()
}

// GetMaxPagesToScrapeOk returns a tuple with the MaxPagesToScrape field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebscrapeRequest) GetMaxPagesToScrapeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxPagesToScrape.Get(), o.MaxPagesToScrape.IsSet()
}

// HasMaxPagesToScrape returns a boolean if a field has been set.
func (o *WebscrapeRequest) HasMaxPagesToScrape() bool {
	if o != nil && o.MaxPagesToScrape.IsSet() {
		return true
	}

	return false
}

// SetMaxPagesToScrape gets a reference to the given NullableInt32 and assigns it to the MaxPagesToScrape field.
func (o *WebscrapeRequest) SetMaxPagesToScrape(v int32) {
	o.MaxPagesToScrape.Set(&v)
}
// SetMaxPagesToScrapeNil sets the value for MaxPagesToScrape to be an explicit nil
func (o *WebscrapeRequest) SetMaxPagesToScrapeNil() {
	o.MaxPagesToScrape.Set(nil)
}

// UnsetMaxPagesToScrape ensures that no value is present for MaxPagesToScrape, not even an explicit nil
func (o *WebscrapeRequest) UnsetMaxPagesToScrape() {
	o.MaxPagesToScrape.Unset()
}

// GetChunkSize returns the ChunkSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebscrapeRequest) GetChunkSize() int32 {
	if o == nil || isNil(o.ChunkSize.Get()) {
		var ret int32
		return ret
	}
	return *o.ChunkSize.Get()
}

// GetChunkSizeOk returns a tuple with the ChunkSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebscrapeRequest) GetChunkSizeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.ChunkSize.Get(), o.ChunkSize.IsSet()
}

// HasChunkSize returns a boolean if a field has been set.
func (o *WebscrapeRequest) HasChunkSize() bool {
	if o != nil && o.ChunkSize.IsSet() {
		return true
	}

	return false
}

// SetChunkSize gets a reference to the given NullableInt32 and assigns it to the ChunkSize field.
func (o *WebscrapeRequest) SetChunkSize(v int32) {
	o.ChunkSize.Set(&v)
}
// SetChunkSizeNil sets the value for ChunkSize to be an explicit nil
func (o *WebscrapeRequest) SetChunkSizeNil() {
	o.ChunkSize.Set(nil)
}

// UnsetChunkSize ensures that no value is present for ChunkSize, not even an explicit nil
func (o *WebscrapeRequest) UnsetChunkSize() {
	o.ChunkSize.Unset()
}

// GetChunkOverlap returns the ChunkOverlap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebscrapeRequest) GetChunkOverlap() int32 {
	if o == nil || isNil(o.ChunkOverlap.Get()) {
		var ret int32
		return ret
	}
	return *o.ChunkOverlap.Get()
}

// GetChunkOverlapOk returns a tuple with the ChunkOverlap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebscrapeRequest) GetChunkOverlapOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.ChunkOverlap.Get(), o.ChunkOverlap.IsSet()
}

// HasChunkOverlap returns a boolean if a field has been set.
func (o *WebscrapeRequest) HasChunkOverlap() bool {
	if o != nil && o.ChunkOverlap.IsSet() {
		return true
	}

	return false
}

// SetChunkOverlap gets a reference to the given NullableInt32 and assigns it to the ChunkOverlap field.
func (o *WebscrapeRequest) SetChunkOverlap(v int32) {
	o.ChunkOverlap.Set(&v)
}
// SetChunkOverlapNil sets the value for ChunkOverlap to be an explicit nil
func (o *WebscrapeRequest) SetChunkOverlapNil() {
	o.ChunkOverlap.Set(nil)
}

// UnsetChunkOverlap ensures that no value is present for ChunkOverlap, not even an explicit nil
func (o *WebscrapeRequest) UnsetChunkOverlap() {
	o.ChunkOverlap.Unset()
}

// GetSkipEmbeddingGeneration returns the SkipEmbeddingGeneration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebscrapeRequest) GetSkipEmbeddingGeneration() bool {
	if o == nil || isNil(o.SkipEmbeddingGeneration.Get()) {
		var ret bool
		return ret
	}
	return *o.SkipEmbeddingGeneration.Get()
}

// GetSkipEmbeddingGenerationOk returns a tuple with the SkipEmbeddingGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebscrapeRequest) GetSkipEmbeddingGenerationOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.SkipEmbeddingGeneration.Get(), o.SkipEmbeddingGeneration.IsSet()
}

// HasSkipEmbeddingGeneration returns a boolean if a field has been set.
func (o *WebscrapeRequest) HasSkipEmbeddingGeneration() bool {
	if o != nil && o.SkipEmbeddingGeneration.IsSet() {
		return true
	}

	return false
}

// SetSkipEmbeddingGeneration gets a reference to the given NullableBool and assigns it to the SkipEmbeddingGeneration field.
func (o *WebscrapeRequest) SetSkipEmbeddingGeneration(v bool) {
	o.SkipEmbeddingGeneration.Set(&v)
}
// SetSkipEmbeddingGenerationNil sets the value for SkipEmbeddingGeneration to be an explicit nil
func (o *WebscrapeRequest) SetSkipEmbeddingGenerationNil() {
	o.SkipEmbeddingGeneration.Set(nil)
}

// UnsetSkipEmbeddingGeneration ensures that no value is present for SkipEmbeddingGeneration, not even an explicit nil
func (o *WebscrapeRequest) UnsetSkipEmbeddingGeneration() {
	o.SkipEmbeddingGeneration.Unset()
}

// GetEnableAutoSync returns the EnableAutoSync field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebscrapeRequest) GetEnableAutoSync() bool {
	if o == nil || isNil(o.EnableAutoSync.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableAutoSync.Get()
}

// GetEnableAutoSyncOk returns a tuple with the EnableAutoSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebscrapeRequest) GetEnableAutoSyncOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.EnableAutoSync.Get(), o.EnableAutoSync.IsSet()
}

// HasEnableAutoSync returns a boolean if a field has been set.
func (o *WebscrapeRequest) HasEnableAutoSync() bool {
	if o != nil && o.EnableAutoSync.IsSet() {
		return true
	}

	return false
}

// SetEnableAutoSync gets a reference to the given NullableBool and assigns it to the EnableAutoSync field.
func (o *WebscrapeRequest) SetEnableAutoSync(v bool) {
	o.EnableAutoSync.Set(&v)
}
// SetEnableAutoSyncNil sets the value for EnableAutoSync to be an explicit nil
func (o *WebscrapeRequest) SetEnableAutoSyncNil() {
	o.EnableAutoSync.Set(nil)
}

// UnsetEnableAutoSync ensures that no value is present for EnableAutoSync, not even an explicit nil
func (o *WebscrapeRequest) UnsetEnableAutoSync() {
	o.EnableAutoSync.Unset()
}

// GetGenerateSparseVectors returns the GenerateSparseVectors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebscrapeRequest) GetGenerateSparseVectors() bool {
	if o == nil || isNil(o.GenerateSparseVectors.Get()) {
		var ret bool
		return ret
	}
	return *o.GenerateSparseVectors.Get()
}

// GetGenerateSparseVectorsOk returns a tuple with the GenerateSparseVectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebscrapeRequest) GetGenerateSparseVectorsOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.GenerateSparseVectors.Get(), o.GenerateSparseVectors.IsSet()
}

// HasGenerateSparseVectors returns a boolean if a field has been set.
func (o *WebscrapeRequest) HasGenerateSparseVectors() bool {
	if o != nil && o.GenerateSparseVectors.IsSet() {
		return true
	}

	return false
}

// SetGenerateSparseVectors gets a reference to the given NullableBool and assigns it to the GenerateSparseVectors field.
func (o *WebscrapeRequest) SetGenerateSparseVectors(v bool) {
	o.GenerateSparseVectors.Set(&v)
}
// SetGenerateSparseVectorsNil sets the value for GenerateSparseVectors to be an explicit nil
func (o *WebscrapeRequest) SetGenerateSparseVectorsNil() {
	o.GenerateSparseVectors.Set(nil)
}

// UnsetGenerateSparseVectors ensures that no value is present for GenerateSparseVectors, not even an explicit nil
func (o *WebscrapeRequest) UnsetGenerateSparseVectors() {
	o.GenerateSparseVectors.Unset()
}

// GetPrependFilenameToChunks returns the PrependFilenameToChunks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebscrapeRequest) GetPrependFilenameToChunks() bool {
	if o == nil || isNil(o.PrependFilenameToChunks.Get()) {
		var ret bool
		return ret
	}
	return *o.PrependFilenameToChunks.Get()
}

// GetPrependFilenameToChunksOk returns a tuple with the PrependFilenameToChunks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebscrapeRequest) GetPrependFilenameToChunksOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.PrependFilenameToChunks.Get(), o.PrependFilenameToChunks.IsSet()
}

// HasPrependFilenameToChunks returns a boolean if a field has been set.
func (o *WebscrapeRequest) HasPrependFilenameToChunks() bool {
	if o != nil && o.PrependFilenameToChunks.IsSet() {
		return true
	}

	return false
}

// SetPrependFilenameToChunks gets a reference to the given NullableBool and assigns it to the PrependFilenameToChunks field.
func (o *WebscrapeRequest) SetPrependFilenameToChunks(v bool) {
	o.PrependFilenameToChunks.Set(&v)
}
// SetPrependFilenameToChunksNil sets the value for PrependFilenameToChunks to be an explicit nil
func (o *WebscrapeRequest) SetPrependFilenameToChunksNil() {
	o.PrependFilenameToChunks.Set(nil)
}

// UnsetPrependFilenameToChunks ensures that no value is present for PrependFilenameToChunks, not even an explicit nil
func (o *WebscrapeRequest) UnsetPrependFilenameToChunks() {
	o.PrependFilenameToChunks.Unset()
}

// GetHtmlTagsToSkip returns the HtmlTagsToSkip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebscrapeRequest) GetHtmlTagsToSkip() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.HtmlTagsToSkip
}

// GetHtmlTagsToSkipOk returns a tuple with the HtmlTagsToSkip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebscrapeRequest) GetHtmlTagsToSkipOk() ([]string, bool) {
	if o == nil || isNil(o.HtmlTagsToSkip) {
    return nil, false
	}
	return o.HtmlTagsToSkip, true
}

// HasHtmlTagsToSkip returns a boolean if a field has been set.
func (o *WebscrapeRequest) HasHtmlTagsToSkip() bool {
	if o != nil && isNil(o.HtmlTagsToSkip) {
		return true
	}

	return false
}

// SetHtmlTagsToSkip gets a reference to the given []string and assigns it to the HtmlTagsToSkip field.
func (o *WebscrapeRequest) SetHtmlTagsToSkip(v []string) {
	o.HtmlTagsToSkip = v
}

// GetCssClassesToSkip returns the CssClassesToSkip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebscrapeRequest) GetCssClassesToSkip() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CssClassesToSkip
}

// GetCssClassesToSkipOk returns a tuple with the CssClassesToSkip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebscrapeRequest) GetCssClassesToSkipOk() ([]string, bool) {
	if o == nil || isNil(o.CssClassesToSkip) {
    return nil, false
	}
	return o.CssClassesToSkip, true
}

// HasCssClassesToSkip returns a boolean if a field has been set.
func (o *WebscrapeRequest) HasCssClassesToSkip() bool {
	if o != nil && isNil(o.CssClassesToSkip) {
		return true
	}

	return false
}

// SetCssClassesToSkip gets a reference to the given []string and assigns it to the CssClassesToSkip field.
func (o *WebscrapeRequest) SetCssClassesToSkip(v []string) {
	o.CssClassesToSkip = v
}

// GetCssSelectorsToSkip returns the CssSelectorsToSkip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebscrapeRequest) GetCssSelectorsToSkip() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CssSelectorsToSkip
}

// GetCssSelectorsToSkipOk returns a tuple with the CssSelectorsToSkip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebscrapeRequest) GetCssSelectorsToSkipOk() ([]string, bool) {
	if o == nil || isNil(o.CssSelectorsToSkip) {
    return nil, false
	}
	return o.CssSelectorsToSkip, true
}

// HasCssSelectorsToSkip returns a boolean if a field has been set.
func (o *WebscrapeRequest) HasCssSelectorsToSkip() bool {
	if o != nil && isNil(o.CssSelectorsToSkip) {
		return true
	}

	return false
}

// SetCssSelectorsToSkip gets a reference to the given []string and assigns it to the CssSelectorsToSkip field.
func (o *WebscrapeRequest) SetCssSelectorsToSkip(v []string) {
	o.CssSelectorsToSkip = v
}

// GetEmbeddingModel returns the EmbeddingModel field value if set, zero value otherwise.
func (o *WebscrapeRequest) GetEmbeddingModel() EmbeddingGenerators {
	if o == nil || isNil(o.EmbeddingModel) {
		var ret EmbeddingGenerators
		return ret
	}
	return *o.EmbeddingModel
}

// GetEmbeddingModelOk returns a tuple with the EmbeddingModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebscrapeRequest) GetEmbeddingModelOk() (*EmbeddingGenerators, bool) {
	if o == nil || isNil(o.EmbeddingModel) {
    return nil, false
	}
	return o.EmbeddingModel, true
}

// HasEmbeddingModel returns a boolean if a field has been set.
func (o *WebscrapeRequest) HasEmbeddingModel() bool {
	if o != nil && !isNil(o.EmbeddingModel) {
		return true
	}

	return false
}

// SetEmbeddingModel gets a reference to the given EmbeddingGenerators and assigns it to the EmbeddingModel field.
func (o *WebscrapeRequest) SetEmbeddingModel(v EmbeddingGenerators) {
	o.EmbeddingModel = &v
}

func (o WebscrapeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.RecursionDepth.IsSet() {
		toSerialize["recursion_depth"] = o.RecursionDepth.Get()
	}
	if o.MaxPagesToScrape.IsSet() {
		toSerialize["max_pages_to_scrape"] = o.MaxPagesToScrape.Get()
	}
	if o.ChunkSize.IsSet() {
		toSerialize["chunk_size"] = o.ChunkSize.Get()
	}
	if o.ChunkOverlap.IsSet() {
		toSerialize["chunk_overlap"] = o.ChunkOverlap.Get()
	}
	if o.SkipEmbeddingGeneration.IsSet() {
		toSerialize["skip_embedding_generation"] = o.SkipEmbeddingGeneration.Get()
	}
	if o.EnableAutoSync.IsSet() {
		toSerialize["enable_auto_sync"] = o.EnableAutoSync.Get()
	}
	if o.GenerateSparseVectors.IsSet() {
		toSerialize["generate_sparse_vectors"] = o.GenerateSparseVectors.Get()
	}
	if o.PrependFilenameToChunks.IsSet() {
		toSerialize["prepend_filename_to_chunks"] = o.PrependFilenameToChunks.Get()
	}
	if o.HtmlTagsToSkip != nil {
		toSerialize["html_tags_to_skip"] = o.HtmlTagsToSkip
	}
	if o.CssClassesToSkip != nil {
		toSerialize["css_classes_to_skip"] = o.CssClassesToSkip
	}
	if o.CssSelectorsToSkip != nil {
		toSerialize["css_selectors_to_skip"] = o.CssSelectorsToSkip
	}
	if !isNil(o.EmbeddingModel) {
		toSerialize["embedding_model"] = o.EmbeddingModel
	}
	return json.Marshal(toSerialize)
}

type NullableWebscrapeRequest struct {
	value *WebscrapeRequest
	isSet bool
}

func (v NullableWebscrapeRequest) Get() *WebscrapeRequest {
	return v.value
}

func (v *NullableWebscrapeRequest) Set(val *WebscrapeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWebscrapeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWebscrapeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebscrapeRequest(val *WebscrapeRequest) *NullableWebscrapeRequest {
	return &NullableWebscrapeRequest{value: val, isSet: true}
}

func (v NullableWebscrapeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebscrapeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



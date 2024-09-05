/*
 * Carbon
 * Connect external data to LLMs, no matter the source.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by Konfig (https://konfigthis.com).
 * Do not edit the class manually.
 */


package com.konfigthis.carbonai.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import com.konfigthis.carbonai.client.model.ColdStorageProps;
import com.konfigthis.carbonai.client.model.EmbeddingGenerators;
import com.konfigthis.carbonai.client.model.FileContentTypesNullable;
import com.konfigthis.carbonai.client.model.TranscriptionServiceNullable;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import org.openapitools.jackson.nullable.JsonNullable;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonArray;
import com.google.gson.JsonDeserializationContext;
import com.google.gson.JsonDeserializer;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonParseException;
import com.google.gson.TypeAdapterFactory;
import com.google.gson.reflect.TypeToken;
import org.apache.commons.lang3.StringUtils;

import java.lang.reflect.Type;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Map.Entry;
import java.util.Set;

import com.konfigthis.carbonai.client.JSON;

/**
 * UploadFileFromUrlInput
 */@javax.annotation.Generated(value = "Generated by https://konfigthis.com")
public class UploadFileFromUrlInput {
  public static final String SERIALIZED_NAME_URL = "url";
  @SerializedName(SERIALIZED_NAME_URL)
  private String url;

  public static final String SERIALIZED_NAME_FILE_NAME = "file_name";
  @SerializedName(SERIALIZED_NAME_FILE_NAME)
  private String fileName;

  public static final String SERIALIZED_NAME_CHUNK_SIZE = "chunk_size";
  @SerializedName(SERIALIZED_NAME_CHUNK_SIZE)
  private Integer chunkSize;

  public static final String SERIALIZED_NAME_CHUNK_OVERLAP = "chunk_overlap";
  @SerializedName(SERIALIZED_NAME_CHUNK_OVERLAP)
  private Integer chunkOverlap;

  public static final String SERIALIZED_NAME_SKIP_EMBEDDING_GENERATION = "skip_embedding_generation";
  @SerializedName(SERIALIZED_NAME_SKIP_EMBEDDING_GENERATION)
  private Boolean skipEmbeddingGeneration = false;

  public static final String SERIALIZED_NAME_SET_PAGE_AS_BOUNDARY = "set_page_as_boundary";
  @SerializedName(SERIALIZED_NAME_SET_PAGE_AS_BOUNDARY)
  private Boolean setPageAsBoundary = false;

  public static final String SERIALIZED_NAME_EMBEDDING_MODEL = "embedding_model";
  @SerializedName(SERIALIZED_NAME_EMBEDDING_MODEL)
  private EmbeddingGenerators embeddingModel;

  public static final String SERIALIZED_NAME_GENERATE_SPARSE_VECTORS = "generate_sparse_vectors";
  @SerializedName(SERIALIZED_NAME_GENERATE_SPARSE_VECTORS)
  private Boolean generateSparseVectors = false;

  public static final String SERIALIZED_NAME_USE_TEXTRACT = "use_textract";
  @SerializedName(SERIALIZED_NAME_USE_TEXTRACT)
  private Boolean useTextract = false;

  public static final String SERIALIZED_NAME_PREPEND_FILENAME_TO_CHUNKS = "prepend_filename_to_chunks";
  @SerializedName(SERIALIZED_NAME_PREPEND_FILENAME_TO_CHUNKS)
  private Boolean prependFilenameToChunks = false;

  public static final String SERIALIZED_NAME_MAX_ITEMS_PER_CHUNK = "max_items_per_chunk";
  @SerializedName(SERIALIZED_NAME_MAX_ITEMS_PER_CHUNK)
  private Integer maxItemsPerChunk;

  public static final String SERIALIZED_NAME_PARSE_PDF_TABLES_WITH_OCR = "parse_pdf_tables_with_ocr";
  @SerializedName(SERIALIZED_NAME_PARSE_PDF_TABLES_WITH_OCR)
  private Boolean parsePdfTablesWithOcr = false;

  public static final String SERIALIZED_NAME_DETECT_AUDIO_LANGUAGE = "detect_audio_language";
  @SerializedName(SERIALIZED_NAME_DETECT_AUDIO_LANGUAGE)
  private Boolean detectAudioLanguage = false;

  public static final String SERIALIZED_NAME_TRANSCRIPTION_SERVICE = "transcription_service";
  @SerializedName(SERIALIZED_NAME_TRANSCRIPTION_SERVICE)
  private TranscriptionServiceNullable transcriptionService;

  public static final String SERIALIZED_NAME_INCLUDE_SPEAKER_LABELS = "include_speaker_labels";
  @SerializedName(SERIALIZED_NAME_INCLUDE_SPEAKER_LABELS)
  private Boolean includeSpeakerLabels = false;

  public static final String SERIALIZED_NAME_MEDIA_TYPE = "media_type";
  @SerializedName(SERIALIZED_NAME_MEDIA_TYPE)
  private FileContentTypesNullable mediaType;

  public static final String SERIALIZED_NAME_SPLIT_ROWS = "split_rows";
  @SerializedName(SERIALIZED_NAME_SPLIT_ROWS)
  private Boolean splitRows = false;

  public static final String SERIALIZED_NAME_COLD_STORAGE_PARAMS = "cold_storage_params";
  @SerializedName(SERIALIZED_NAME_COLD_STORAGE_PARAMS)
  private ColdStorageProps coldStorageParams;

  public static final String SERIALIZED_NAME_GENERATE_CHUNKS_ONLY = "generate_chunks_only";
  @SerializedName(SERIALIZED_NAME_GENERATE_CHUNKS_ONLY)
  private Boolean generateChunksOnly = false;

  public UploadFileFromUrlInput() {
  }

  public UploadFileFromUrlInput url(String url) {
    
    
    
    
    this.url = url;
    return this;
  }

   /**
   * Get url
   * @return url
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")

  public String getUrl() {
    return url;
  }


  public void setUrl(String url) {
    
    
    
    this.url = url;
  }


  public UploadFileFromUrlInput fileName(String fileName) {
    
    
    
    
    this.fileName = fileName;
    return this;
  }

   /**
   * Get fileName
   * @return fileName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getFileName() {
    return fileName;
  }


  public void setFileName(String fileName) {
    
    
    
    this.fileName = fileName;
  }


  public UploadFileFromUrlInput chunkSize(Integer chunkSize) {
    
    
    
    
    this.chunkSize = chunkSize;
    return this;
  }

   /**
   * Get chunkSize
   * @return chunkSize
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getChunkSize() {
    return chunkSize;
  }


  public void setChunkSize(Integer chunkSize) {
    
    
    
    this.chunkSize = chunkSize;
  }


  public UploadFileFromUrlInput chunkOverlap(Integer chunkOverlap) {
    
    
    
    
    this.chunkOverlap = chunkOverlap;
    return this;
  }

   /**
   * Get chunkOverlap
   * @return chunkOverlap
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getChunkOverlap() {
    return chunkOverlap;
  }


  public void setChunkOverlap(Integer chunkOverlap) {
    
    
    
    this.chunkOverlap = chunkOverlap;
  }


  public UploadFileFromUrlInput skipEmbeddingGeneration(Boolean skipEmbeddingGeneration) {
    
    
    
    
    this.skipEmbeddingGeneration = skipEmbeddingGeneration;
    return this;
  }

   /**
   * Get skipEmbeddingGeneration
   * @return skipEmbeddingGeneration
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "false", value = "")

  public Boolean getSkipEmbeddingGeneration() {
    return skipEmbeddingGeneration;
  }


  public void setSkipEmbeddingGeneration(Boolean skipEmbeddingGeneration) {
    
    
    
    this.skipEmbeddingGeneration = skipEmbeddingGeneration;
  }


  public UploadFileFromUrlInput setPageAsBoundary(Boolean setPageAsBoundary) {
    
    
    
    
    this.setPageAsBoundary = setPageAsBoundary;
    return this;
  }

   /**
   * Get setPageAsBoundary
   * @return setPageAsBoundary
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "false", value = "")

  public Boolean getSetPageAsBoundary() {
    return setPageAsBoundary;
  }


  public void setSetPageAsBoundary(Boolean setPageAsBoundary) {
    
    
    
    this.setPageAsBoundary = setPageAsBoundary;
  }


  public UploadFileFromUrlInput embeddingModel(EmbeddingGenerators embeddingModel) {
    
    
    
    
    this.embeddingModel = embeddingModel;
    return this;
  }

   /**
   * Get embeddingModel
   * @return embeddingModel
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public EmbeddingGenerators getEmbeddingModel() {
    return embeddingModel;
  }


  public void setEmbeddingModel(EmbeddingGenerators embeddingModel) {
    
    
    
    this.embeddingModel = embeddingModel;
  }


  public UploadFileFromUrlInput generateSparseVectors(Boolean generateSparseVectors) {
    
    
    
    
    this.generateSparseVectors = generateSparseVectors;
    return this;
  }

   /**
   * Get generateSparseVectors
   * @return generateSparseVectors
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "false", value = "")

  public Boolean getGenerateSparseVectors() {
    return generateSparseVectors;
  }


  public void setGenerateSparseVectors(Boolean generateSparseVectors) {
    
    
    
    this.generateSparseVectors = generateSparseVectors;
  }


  public UploadFileFromUrlInput useTextract(Boolean useTextract) {
    
    
    
    
    this.useTextract = useTextract;
    return this;
  }

   /**
   * Get useTextract
   * @return useTextract
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "false", value = "")

  public Boolean getUseTextract() {
    return useTextract;
  }


  public void setUseTextract(Boolean useTextract) {
    
    
    
    this.useTextract = useTextract;
  }


  public UploadFileFromUrlInput prependFilenameToChunks(Boolean prependFilenameToChunks) {
    
    
    
    
    this.prependFilenameToChunks = prependFilenameToChunks;
    return this;
  }

   /**
   * Get prependFilenameToChunks
   * @return prependFilenameToChunks
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "false", value = "")

  public Boolean getPrependFilenameToChunks() {
    return prependFilenameToChunks;
  }


  public void setPrependFilenameToChunks(Boolean prependFilenameToChunks) {
    
    
    
    this.prependFilenameToChunks = prependFilenameToChunks;
  }


  public UploadFileFromUrlInput maxItemsPerChunk(Integer maxItemsPerChunk) {
    
    
    
    
    this.maxItemsPerChunk = maxItemsPerChunk;
    return this;
  }

   /**
   * Number of objects per chunk. For csv, tsv, xlsx, and json files only.
   * @return maxItemsPerChunk
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Number of objects per chunk. For csv, tsv, xlsx, and json files only.")

  public Integer getMaxItemsPerChunk() {
    return maxItemsPerChunk;
  }


  public void setMaxItemsPerChunk(Integer maxItemsPerChunk) {
    
    
    
    this.maxItemsPerChunk = maxItemsPerChunk;
  }


  public UploadFileFromUrlInput parsePdfTablesWithOcr(Boolean parsePdfTablesWithOcr) {
    
    
    
    
    this.parsePdfTablesWithOcr = parsePdfTablesWithOcr;
    return this;
  }

   /**
   * Get parsePdfTablesWithOcr
   * @return parsePdfTablesWithOcr
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "false", value = "")

  public Boolean getParsePdfTablesWithOcr() {
    return parsePdfTablesWithOcr;
  }


  public void setParsePdfTablesWithOcr(Boolean parsePdfTablesWithOcr) {
    
    
    
    this.parsePdfTablesWithOcr = parsePdfTablesWithOcr;
  }


  public UploadFileFromUrlInput detectAudioLanguage(Boolean detectAudioLanguage) {
    
    
    
    
    this.detectAudioLanguage = detectAudioLanguage;
    return this;
  }

   /**
   * Get detectAudioLanguage
   * @return detectAudioLanguage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "false", value = "")

  public Boolean getDetectAudioLanguage() {
    return detectAudioLanguage;
  }


  public void setDetectAudioLanguage(Boolean detectAudioLanguage) {
    
    
    
    this.detectAudioLanguage = detectAudioLanguage;
  }


  public UploadFileFromUrlInput transcriptionService(TranscriptionServiceNullable transcriptionService) {
    
    
    
    
    this.transcriptionService = transcriptionService;
    return this;
  }

   /**
   * Get transcriptionService
   * @return transcriptionService
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public TranscriptionServiceNullable getTranscriptionService() {
    return transcriptionService;
  }


  public void setTranscriptionService(TranscriptionServiceNullable transcriptionService) {
    
    
    
    this.transcriptionService = transcriptionService;
  }


  public UploadFileFromUrlInput includeSpeakerLabels(Boolean includeSpeakerLabels) {
    
    
    
    
    this.includeSpeakerLabels = includeSpeakerLabels;
    return this;
  }

   /**
   * Get includeSpeakerLabels
   * @return includeSpeakerLabels
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "false", value = "")

  public Boolean getIncludeSpeakerLabels() {
    return includeSpeakerLabels;
  }


  public void setIncludeSpeakerLabels(Boolean includeSpeakerLabels) {
    
    
    
    this.includeSpeakerLabels = includeSpeakerLabels;
  }


  public UploadFileFromUrlInput mediaType(FileContentTypesNullable mediaType) {
    
    
    
    
    this.mediaType = mediaType;
    return this;
  }

   /**
   * Get mediaType
   * @return mediaType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public FileContentTypesNullable getMediaType() {
    return mediaType;
  }


  public void setMediaType(FileContentTypesNullable mediaType) {
    
    
    
    this.mediaType = mediaType;
  }


  public UploadFileFromUrlInput splitRows(Boolean splitRows) {
    
    
    
    
    this.splitRows = splitRows;
    return this;
  }

   /**
   * Get splitRows
   * @return splitRows
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "false", value = "")

  public Boolean getSplitRows() {
    return splitRows;
  }


  public void setSplitRows(Boolean splitRows) {
    
    
    
    this.splitRows = splitRows;
  }


  public UploadFileFromUrlInput coldStorageParams(ColdStorageProps coldStorageParams) {
    
    
    
    
    this.coldStorageParams = coldStorageParams;
    return this;
  }

   /**
   * Get coldStorageParams
   * @return coldStorageParams
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ColdStorageProps getColdStorageParams() {
    return coldStorageParams;
  }


  public void setColdStorageParams(ColdStorageProps coldStorageParams) {
    
    
    
    this.coldStorageParams = coldStorageParams;
  }


  public UploadFileFromUrlInput generateChunksOnly(Boolean generateChunksOnly) {
    
    
    
    
    this.generateChunksOnly = generateChunksOnly;
    return this;
  }

   /**
   * If this flag is enabled, the file will be chunked and stored with Carbon,         but no embeddings will be generated. This overrides the skip_embedding_generation flag.
   * @return generateChunksOnly
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "false", value = "If this flag is enabled, the file will be chunked and stored with Carbon,         but no embeddings will be generated. This overrides the skip_embedding_generation flag.")

  public Boolean getGenerateChunksOnly() {
    return generateChunksOnly;
  }


  public void setGenerateChunksOnly(Boolean generateChunksOnly) {
    
    
    
    this.generateChunksOnly = generateChunksOnly;
  }

  /**
   * A container for additional, undeclared properties.
   * This is a holder for any undeclared properties as specified with
   * the 'additionalProperties' keyword in the OAS document.
   */
  private Map<String, Object> additionalProperties;

  /**
   * Set the additional (undeclared) property with the specified name and value.
   * If the property does not already exist, create it otherwise replace it.
   *
   * @param key name of the property
   * @param value value of the property
   * @return the UploadFileFromUrlInput instance itself
   */
  public UploadFileFromUrlInput putAdditionalProperty(String key, Object value) {
    if (this.additionalProperties == null) {
        this.additionalProperties = new HashMap<String, Object>();
    }
    this.additionalProperties.put(key, value);
    return this;
  }

  /**
   * Return the additional (undeclared) property.
   *
   * @return a map of objects
   */
  public Map<String, Object> getAdditionalProperties() {
    return additionalProperties;
  }

  /**
   * Return the additional (undeclared) property with the specified name.
   *
   * @param key name of the property
   * @return an object
   */
  public Object getAdditionalProperty(String key) {
    if (this.additionalProperties == null) {
        return null;
    }
    return this.additionalProperties.get(key);
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    UploadFileFromUrlInput uploadFileFromUrlInput = (UploadFileFromUrlInput) o;
    return Objects.equals(this.url, uploadFileFromUrlInput.url) &&
        Objects.equals(this.fileName, uploadFileFromUrlInput.fileName) &&
        Objects.equals(this.chunkSize, uploadFileFromUrlInput.chunkSize) &&
        Objects.equals(this.chunkOverlap, uploadFileFromUrlInput.chunkOverlap) &&
        Objects.equals(this.skipEmbeddingGeneration, uploadFileFromUrlInput.skipEmbeddingGeneration) &&
        Objects.equals(this.setPageAsBoundary, uploadFileFromUrlInput.setPageAsBoundary) &&
        Objects.equals(this.embeddingModel, uploadFileFromUrlInput.embeddingModel) &&
        Objects.equals(this.generateSparseVectors, uploadFileFromUrlInput.generateSparseVectors) &&
        Objects.equals(this.useTextract, uploadFileFromUrlInput.useTextract) &&
        Objects.equals(this.prependFilenameToChunks, uploadFileFromUrlInput.prependFilenameToChunks) &&
        Objects.equals(this.maxItemsPerChunk, uploadFileFromUrlInput.maxItemsPerChunk) &&
        Objects.equals(this.parsePdfTablesWithOcr, uploadFileFromUrlInput.parsePdfTablesWithOcr) &&
        Objects.equals(this.detectAudioLanguage, uploadFileFromUrlInput.detectAudioLanguage) &&
        Objects.equals(this.transcriptionService, uploadFileFromUrlInput.transcriptionService) &&
        Objects.equals(this.includeSpeakerLabels, uploadFileFromUrlInput.includeSpeakerLabels) &&
        Objects.equals(this.mediaType, uploadFileFromUrlInput.mediaType) &&
        Objects.equals(this.splitRows, uploadFileFromUrlInput.splitRows) &&
        Objects.equals(this.coldStorageParams, uploadFileFromUrlInput.coldStorageParams) &&
        Objects.equals(this.generateChunksOnly, uploadFileFromUrlInput.generateChunksOnly)&&
        Objects.equals(this.additionalProperties, uploadFileFromUrlInput.additionalProperties);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(url, fileName, chunkSize, chunkOverlap, skipEmbeddingGeneration, setPageAsBoundary, embeddingModel, generateSparseVectors, useTextract, prependFilenameToChunks, maxItemsPerChunk, parsePdfTablesWithOcr, detectAudioLanguage, transcriptionService, includeSpeakerLabels, mediaType, splitRows, coldStorageParams, generateChunksOnly, additionalProperties);
  }

  private static <T> int hashCodeNullable(JsonNullable<T> a) {
    if (a == null) {
      return 1;
    }
    return a.isPresent() ? Arrays.deepHashCode(new Object[]{a.get()}) : 31;
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UploadFileFromUrlInput {\n");
    sb.append("    url: ").append(toIndentedString(url)).append("\n");
    sb.append("    fileName: ").append(toIndentedString(fileName)).append("\n");
    sb.append("    chunkSize: ").append(toIndentedString(chunkSize)).append("\n");
    sb.append("    chunkOverlap: ").append(toIndentedString(chunkOverlap)).append("\n");
    sb.append("    skipEmbeddingGeneration: ").append(toIndentedString(skipEmbeddingGeneration)).append("\n");
    sb.append("    setPageAsBoundary: ").append(toIndentedString(setPageAsBoundary)).append("\n");
    sb.append("    embeddingModel: ").append(toIndentedString(embeddingModel)).append("\n");
    sb.append("    generateSparseVectors: ").append(toIndentedString(generateSparseVectors)).append("\n");
    sb.append("    useTextract: ").append(toIndentedString(useTextract)).append("\n");
    sb.append("    prependFilenameToChunks: ").append(toIndentedString(prependFilenameToChunks)).append("\n");
    sb.append("    maxItemsPerChunk: ").append(toIndentedString(maxItemsPerChunk)).append("\n");
    sb.append("    parsePdfTablesWithOcr: ").append(toIndentedString(parsePdfTablesWithOcr)).append("\n");
    sb.append("    detectAudioLanguage: ").append(toIndentedString(detectAudioLanguage)).append("\n");
    sb.append("    transcriptionService: ").append(toIndentedString(transcriptionService)).append("\n");
    sb.append("    includeSpeakerLabels: ").append(toIndentedString(includeSpeakerLabels)).append("\n");
    sb.append("    mediaType: ").append(toIndentedString(mediaType)).append("\n");
    sb.append("    splitRows: ").append(toIndentedString(splitRows)).append("\n");
    sb.append("    coldStorageParams: ").append(toIndentedString(coldStorageParams)).append("\n");
    sb.append("    generateChunksOnly: ").append(toIndentedString(generateChunksOnly)).append("\n");
    sb.append("    additionalProperties: ").append(toIndentedString(additionalProperties)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }


  public static HashSet<String> openapiFields;
  public static HashSet<String> openapiRequiredFields;

  static {
    // a set of all properties/fields (JSON key names)
    openapiFields = new HashSet<String>();
    openapiFields.add("url");
    openapiFields.add("file_name");
    openapiFields.add("chunk_size");
    openapiFields.add("chunk_overlap");
    openapiFields.add("skip_embedding_generation");
    openapiFields.add("set_page_as_boundary");
    openapiFields.add("embedding_model");
    openapiFields.add("generate_sparse_vectors");
    openapiFields.add("use_textract");
    openapiFields.add("prepend_filename_to_chunks");
    openapiFields.add("max_items_per_chunk");
    openapiFields.add("parse_pdf_tables_with_ocr");
    openapiFields.add("detect_audio_language");
    openapiFields.add("transcription_service");
    openapiFields.add("include_speaker_labels");
    openapiFields.add("media_type");
    openapiFields.add("split_rows");
    openapiFields.add("cold_storage_params");
    openapiFields.add("generate_chunks_only");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("url");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to UploadFileFromUrlInput
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!UploadFileFromUrlInput.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in UploadFileFromUrlInput is not found in the empty JSON string", UploadFileFromUrlInput.openapiRequiredFields.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : UploadFileFromUrlInput.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if (!jsonObj.get("url").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `url` to be a primitive type in the JSON string but got `%s`", jsonObj.get("url").toString()));
      }
      if (!jsonObj.get("file_name").isJsonNull() && (jsonObj.get("file_name") != null && !jsonObj.get("file_name").isJsonNull()) && !jsonObj.get("file_name").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `file_name` to be a primitive type in the JSON string but got `%s`", jsonObj.get("file_name").toString()));
      }
      // validate the optional field `cold_storage_params`
      if (jsonObj.get("cold_storage_params") != null && !jsonObj.get("cold_storage_params").isJsonNull()) {
        ColdStorageProps.validateJsonObject(jsonObj.getAsJsonObject("cold_storage_params"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!UploadFileFromUrlInput.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'UploadFileFromUrlInput' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<UploadFileFromUrlInput> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(UploadFileFromUrlInput.class));

       return (TypeAdapter<T>) new TypeAdapter<UploadFileFromUrlInput>() {
           @Override
           public void write(JsonWriter out, UploadFileFromUrlInput value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             obj.remove("additionalProperties");
             // serialize additonal properties
             if (value.getAdditionalProperties() != null) {
               for (Map.Entry<String, Object> entry : value.getAdditionalProperties().entrySet()) {
                 if (entry.getValue() instanceof String)
                   obj.addProperty(entry.getKey(), (String) entry.getValue());
                 else if (entry.getValue() instanceof Number)
                   obj.addProperty(entry.getKey(), (Number) entry.getValue());
                 else if (entry.getValue() instanceof Boolean)
                   obj.addProperty(entry.getKey(), (Boolean) entry.getValue());
                 else if (entry.getValue() instanceof Character)
                   obj.addProperty(entry.getKey(), (Character) entry.getValue());
                 else if (entry.getValue() == null) {
                   obj.addProperty(entry.getKey(), (String) null);
                 } else {
                   obj.add(entry.getKey(), gson.toJsonTree(entry.getValue()).getAsJsonObject());
                 }
               }
             }
             elementAdapter.write(out, obj);
           }

           @Override
           public UploadFileFromUrlInput read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             // store additional fields in the deserialized instance
             UploadFileFromUrlInput instance = thisAdapter.fromJsonTree(jsonObj);
             for (Map.Entry<String, JsonElement> entry : jsonObj.entrySet()) {
               if (!openapiFields.contains(entry.getKey())) {
                 if (entry.getValue().isJsonPrimitive()) { // primitive type
                   if (entry.getValue().getAsJsonPrimitive().isString())
                     instance.putAdditionalProperty(entry.getKey(), entry.getValue().getAsString());
                   else if (entry.getValue().getAsJsonPrimitive().isNumber())
                     instance.putAdditionalProperty(entry.getKey(), entry.getValue().getAsNumber());
                   else if (entry.getValue().getAsJsonPrimitive().isBoolean())
                     instance.putAdditionalProperty(entry.getKey(), entry.getValue().getAsBoolean());
                   else
                     throw new IllegalArgumentException(String.format("The field `%s` has unknown primitive type. Value: %s", entry.getKey(), entry.getValue().toString()));
                 } else if (entry.getValue().isJsonArray()) {
                     instance.putAdditionalProperty(entry.getKey(), gson.fromJson(entry.getValue(), List.class));
                 } else { // JSON object
                     instance.putAdditionalProperty(entry.getKey(), gson.fromJson(entry.getValue(), HashMap.class));
                 }
               }
             }
             return instance;
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of UploadFileFromUrlInput given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of UploadFileFromUrlInput
  * @throws IOException if the JSON string is invalid with respect to UploadFileFromUrlInput
  */
  public static UploadFileFromUrlInput fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, UploadFileFromUrlInput.class);
  }

 /**
  * Convert an instance of UploadFileFromUrlInput to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}


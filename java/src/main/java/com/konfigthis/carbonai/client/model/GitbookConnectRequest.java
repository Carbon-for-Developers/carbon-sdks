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
import com.konfigthis.carbonai.client.model.EmbeddingGenerators;
import com.konfigthis.carbonai.client.model.FileSyncConfigNullable;
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
 * GitbookConnectRequest
 */@javax.annotation.Generated(value = "Generated by https://konfigthis.com")
public class GitbookConnectRequest {
  public static final String SERIALIZED_NAME_TAGS = "tags";
  @SerializedName(SERIALIZED_NAME_TAGS)
  private Object tags;

  public static final String SERIALIZED_NAME_ORGANIZATION = "organization";
  @SerializedName(SERIALIZED_NAME_ORGANIZATION)
  private String organization;

  public static final String SERIALIZED_NAME_ACCESS_TOKEN = "access_token";
  @SerializedName(SERIALIZED_NAME_ACCESS_TOKEN)
  private String accessToken;

  public static final String SERIALIZED_NAME_CHUNK_SIZE = "chunk_size";
  @SerializedName(SERIALIZED_NAME_CHUNK_SIZE)
  private Integer chunkSize = 1500;

  public static final String SERIALIZED_NAME_CHUNK_OVERLAP = "chunk_overlap";
  @SerializedName(SERIALIZED_NAME_CHUNK_OVERLAP)
  private Integer chunkOverlap = 20;

  public static final String SERIALIZED_NAME_SKIP_EMBEDDING_GENERATION = "skip_embedding_generation";
  @SerializedName(SERIALIZED_NAME_SKIP_EMBEDDING_GENERATION)
  private Boolean skipEmbeddingGeneration = false;

  public static final String SERIALIZED_NAME_EMBEDDING_MODEL = "embedding_model";
  @SerializedName(SERIALIZED_NAME_EMBEDDING_MODEL)
  private EmbeddingGenerators embeddingModel;

  public static final String SERIALIZED_NAME_GENERATE_SPARSE_VECTORS = "generate_sparse_vectors";
  @SerializedName(SERIALIZED_NAME_GENERATE_SPARSE_VECTORS)
  private Boolean generateSparseVectors = false;

  public static final String SERIALIZED_NAME_PREPEND_FILENAME_TO_CHUNKS = "prepend_filename_to_chunks";
  @SerializedName(SERIALIZED_NAME_PREPEND_FILENAME_TO_CHUNKS)
  private Boolean prependFilenameToChunks = false;

  public static final String SERIALIZED_NAME_SYNC_FILES_ON_CONNECTION = "sync_files_on_connection";
  @SerializedName(SERIALIZED_NAME_SYNC_FILES_ON_CONNECTION)
  private Boolean syncFilesOnConnection = true;

  public static final String SERIALIZED_NAME_REQUEST_ID = "request_id";
  @SerializedName(SERIALIZED_NAME_REQUEST_ID)
  private String requestId;

  public static final String SERIALIZED_NAME_SYNC_SOURCE_ITEMS = "sync_source_items";
  @SerializedName(SERIALIZED_NAME_SYNC_SOURCE_ITEMS)
  private Boolean syncSourceItems = true;

  public static final String SERIALIZED_NAME_FILE_SYNC_CONFIG = "file_sync_config";
  @SerializedName(SERIALIZED_NAME_FILE_SYNC_CONFIG)
  private FileSyncConfigNullable fileSyncConfig;

  public static final String SERIALIZED_NAME_DATA_SOURCE_TAGS = "data_source_tags";
  @SerializedName(SERIALIZED_NAME_DATA_SOURCE_TAGS)
  private Object dataSourceTags;

  public GitbookConnectRequest() {
  }

  public GitbookConnectRequest tags(Object tags) {
    
    
    
    
    this.tags = tags;
    return this;
  }

   /**
   * Get tags
   * @return tags
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "{}", value = "")

  public Object getTags() {
    return tags;
  }


  public void setTags(Object tags) {
    
    
    
    this.tags = tags;
  }


  public GitbookConnectRequest organization(String organization) {
    
    
    
    
    this.organization = organization;
    return this;
  }

   /**
   * Get organization
   * @return organization
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")

  public String getOrganization() {
    return organization;
  }


  public void setOrganization(String organization) {
    
    
    
    this.organization = organization;
  }


  public GitbookConnectRequest accessToken(String accessToken) {
    
    
    
    
    this.accessToken = accessToken;
    return this;
  }

   /**
   * Get accessToken
   * @return accessToken
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")

  public String getAccessToken() {
    return accessToken;
  }


  public void setAccessToken(String accessToken) {
    
    
    
    this.accessToken = accessToken;
  }


  public GitbookConnectRequest chunkSize(Integer chunkSize) {
    
    
    
    
    this.chunkSize = chunkSize;
    return this;
  }

   /**
   * Get chunkSize
   * @return chunkSize
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "1500", value = "")

  public Integer getChunkSize() {
    return chunkSize;
  }


  public void setChunkSize(Integer chunkSize) {
    
    
    
    this.chunkSize = chunkSize;
  }


  public GitbookConnectRequest chunkOverlap(Integer chunkOverlap) {
    
    
    
    
    this.chunkOverlap = chunkOverlap;
    return this;
  }

   /**
   * Get chunkOverlap
   * @return chunkOverlap
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "20", value = "")

  public Integer getChunkOverlap() {
    return chunkOverlap;
  }


  public void setChunkOverlap(Integer chunkOverlap) {
    
    
    
    this.chunkOverlap = chunkOverlap;
  }


  public GitbookConnectRequest skipEmbeddingGeneration(Boolean skipEmbeddingGeneration) {
    
    
    
    
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


  public GitbookConnectRequest embeddingModel(EmbeddingGenerators embeddingModel) {
    
    
    
    
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


  public GitbookConnectRequest generateSparseVectors(Boolean generateSparseVectors) {
    
    
    
    
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


  public GitbookConnectRequest prependFilenameToChunks(Boolean prependFilenameToChunks) {
    
    
    
    
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


  public GitbookConnectRequest syncFilesOnConnection(Boolean syncFilesOnConnection) {
    
    
    
    
    this.syncFilesOnConnection = syncFilesOnConnection;
    return this;
  }

   /**
   * Get syncFilesOnConnection
   * @return syncFilesOnConnection
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "true", value = "")

  public Boolean getSyncFilesOnConnection() {
    return syncFilesOnConnection;
  }


  public void setSyncFilesOnConnection(Boolean syncFilesOnConnection) {
    
    
    
    this.syncFilesOnConnection = syncFilesOnConnection;
  }


  public GitbookConnectRequest requestId(String requestId) {
    
    
    
    
    this.requestId = requestId;
    return this;
  }

   /**
   * Get requestId
   * @return requestId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRequestId() {
    return requestId;
  }


  public void setRequestId(String requestId) {
    
    
    
    this.requestId = requestId;
  }


  public GitbookConnectRequest syncSourceItems(Boolean syncSourceItems) {
    
    
    
    
    this.syncSourceItems = syncSourceItems;
    return this;
  }

   /**
   * Enabling this flag will fetch all available content from the source to be listed via list items endpoint
   * @return syncSourceItems
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "true", value = "Enabling this flag will fetch all available content from the source to be listed via list items endpoint")

  public Boolean getSyncSourceItems() {
    return syncSourceItems;
  }


  public void setSyncSourceItems(Boolean syncSourceItems) {
    
    
    
    this.syncSourceItems = syncSourceItems;
  }


  public GitbookConnectRequest fileSyncConfig(FileSyncConfigNullable fileSyncConfig) {
    
    
    
    
    this.fileSyncConfig = fileSyncConfig;
    return this;
  }

   /**
   * Get fileSyncConfig
   * @return fileSyncConfig
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public FileSyncConfigNullable getFileSyncConfig() {
    return fileSyncConfig;
  }


  public void setFileSyncConfig(FileSyncConfigNullable fileSyncConfig) {
    
    
    
    this.fileSyncConfig = fileSyncConfig;
  }


  public GitbookConnectRequest dataSourceTags(Object dataSourceTags) {
    
    
    
    
    this.dataSourceTags = dataSourceTags;
    return this;
  }

   /**
   * Tags to be associated with the data source. If the data source already has tags set, then an upsert will be performed.
   * @return dataSourceTags
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "{}", value = "Tags to be associated with the data source. If the data source already has tags set, then an upsert will be performed.")

  public Object getDataSourceTags() {
    return dataSourceTags;
  }


  public void setDataSourceTags(Object dataSourceTags) {
    
    
    
    this.dataSourceTags = dataSourceTags;
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
   * @return the GitbookConnectRequest instance itself
   */
  public GitbookConnectRequest putAdditionalProperty(String key, Object value) {
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
    GitbookConnectRequest gitbookConnectRequest = (GitbookConnectRequest) o;
    return Objects.equals(this.tags, gitbookConnectRequest.tags) &&
        Objects.equals(this.organization, gitbookConnectRequest.organization) &&
        Objects.equals(this.accessToken, gitbookConnectRequest.accessToken) &&
        Objects.equals(this.chunkSize, gitbookConnectRequest.chunkSize) &&
        Objects.equals(this.chunkOverlap, gitbookConnectRequest.chunkOverlap) &&
        Objects.equals(this.skipEmbeddingGeneration, gitbookConnectRequest.skipEmbeddingGeneration) &&
        Objects.equals(this.embeddingModel, gitbookConnectRequest.embeddingModel) &&
        Objects.equals(this.generateSparseVectors, gitbookConnectRequest.generateSparseVectors) &&
        Objects.equals(this.prependFilenameToChunks, gitbookConnectRequest.prependFilenameToChunks) &&
        Objects.equals(this.syncFilesOnConnection, gitbookConnectRequest.syncFilesOnConnection) &&
        Objects.equals(this.requestId, gitbookConnectRequest.requestId) &&
        Objects.equals(this.syncSourceItems, gitbookConnectRequest.syncSourceItems) &&
        Objects.equals(this.fileSyncConfig, gitbookConnectRequest.fileSyncConfig) &&
        Objects.equals(this.dataSourceTags, gitbookConnectRequest.dataSourceTags)&&
        Objects.equals(this.additionalProperties, gitbookConnectRequest.additionalProperties);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(tags, organization, accessToken, chunkSize, chunkOverlap, skipEmbeddingGeneration, embeddingModel, generateSparseVectors, prependFilenameToChunks, syncFilesOnConnection, requestId, syncSourceItems, fileSyncConfig, dataSourceTags, additionalProperties);
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
    sb.append("class GitbookConnectRequest {\n");
    sb.append("    tags: ").append(toIndentedString(tags)).append("\n");
    sb.append("    organization: ").append(toIndentedString(organization)).append("\n");
    sb.append("    accessToken: ").append(toIndentedString(accessToken)).append("\n");
    sb.append("    chunkSize: ").append(toIndentedString(chunkSize)).append("\n");
    sb.append("    chunkOverlap: ").append(toIndentedString(chunkOverlap)).append("\n");
    sb.append("    skipEmbeddingGeneration: ").append(toIndentedString(skipEmbeddingGeneration)).append("\n");
    sb.append("    embeddingModel: ").append(toIndentedString(embeddingModel)).append("\n");
    sb.append("    generateSparseVectors: ").append(toIndentedString(generateSparseVectors)).append("\n");
    sb.append("    prependFilenameToChunks: ").append(toIndentedString(prependFilenameToChunks)).append("\n");
    sb.append("    syncFilesOnConnection: ").append(toIndentedString(syncFilesOnConnection)).append("\n");
    sb.append("    requestId: ").append(toIndentedString(requestId)).append("\n");
    sb.append("    syncSourceItems: ").append(toIndentedString(syncSourceItems)).append("\n");
    sb.append("    fileSyncConfig: ").append(toIndentedString(fileSyncConfig)).append("\n");
    sb.append("    dataSourceTags: ").append(toIndentedString(dataSourceTags)).append("\n");
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
    openapiFields.add("tags");
    openapiFields.add("organization");
    openapiFields.add("access_token");
    openapiFields.add("chunk_size");
    openapiFields.add("chunk_overlap");
    openapiFields.add("skip_embedding_generation");
    openapiFields.add("embedding_model");
    openapiFields.add("generate_sparse_vectors");
    openapiFields.add("prepend_filename_to_chunks");
    openapiFields.add("sync_files_on_connection");
    openapiFields.add("request_id");
    openapiFields.add("sync_source_items");
    openapiFields.add("file_sync_config");
    openapiFields.add("data_source_tags");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("organization");
    openapiRequiredFields.add("access_token");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to GitbookConnectRequest
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!GitbookConnectRequest.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in GitbookConnectRequest is not found in the empty JSON string", GitbookConnectRequest.openapiRequiredFields.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : GitbookConnectRequest.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if (!jsonObj.get("organization").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `organization` to be a primitive type in the JSON string but got `%s`", jsonObj.get("organization").toString()));
      }
      if (!jsonObj.get("access_token").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `access_token` to be a primitive type in the JSON string but got `%s`", jsonObj.get("access_token").toString()));
      }
      if (!jsonObj.get("request_id").isJsonNull() && (jsonObj.get("request_id") != null && !jsonObj.get("request_id").isJsonNull()) && !jsonObj.get("request_id").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `request_id` to be a primitive type in the JSON string but got `%s`", jsonObj.get("request_id").toString()));
      }
      // validate the optional field `file_sync_config`
      if (jsonObj.get("file_sync_config") != null && !jsonObj.get("file_sync_config").isJsonNull()) {
        FileSyncConfigNullable.validateJsonObject(jsonObj.getAsJsonObject("file_sync_config"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!GitbookConnectRequest.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'GitbookConnectRequest' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<GitbookConnectRequest> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(GitbookConnectRequest.class));

       return (TypeAdapter<T>) new TypeAdapter<GitbookConnectRequest>() {
           @Override
           public void write(JsonWriter out, GitbookConnectRequest value) throws IOException {
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
           public GitbookConnectRequest read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             // store additional fields in the deserialized instance
             GitbookConnectRequest instance = thisAdapter.fromJsonTree(jsonObj);
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
  * Create an instance of GitbookConnectRequest given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of GitbookConnectRequest
  * @throws IOException if the JSON string is invalid with respect to GitbookConnectRequest
  */
  public static GitbookConnectRequest fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, GitbookConnectRequest.class);
  }

 /**
  * Convert an instance of GitbookConnectRequest to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}


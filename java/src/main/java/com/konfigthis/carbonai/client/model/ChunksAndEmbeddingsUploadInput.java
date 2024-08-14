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
import com.konfigthis.carbonai.client.model.SingleChunksAndEmbeddingsUploadInput;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

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
 * ChunksAndEmbeddingsUploadInput
 */@javax.annotation.Generated(value = "Generated by https://konfigthis.com")
public class ChunksAndEmbeddingsUploadInput {
  public static final String SERIALIZED_NAME_EMBEDDING_MODEL = "embedding_model";
  @SerializedName(SERIALIZED_NAME_EMBEDDING_MODEL)
  private EmbeddingGenerators embeddingModel;

  public static final String SERIALIZED_NAME_CHUNKS_AND_EMBEDDINGS = "chunks_and_embeddings";
  @SerializedName(SERIALIZED_NAME_CHUNKS_AND_EMBEDDINGS)
  private List<SingleChunksAndEmbeddingsUploadInput> chunksAndEmbeddings = new ArrayList<>();

  public static final String SERIALIZED_NAME_OVERWRITE_EXISTING = "overwrite_existing";
  @SerializedName(SERIALIZED_NAME_OVERWRITE_EXISTING)
  private Boolean overwriteExisting = false;

  public static final String SERIALIZED_NAME_CHUNKS_ONLY = "chunks_only";
  @SerializedName(SERIALIZED_NAME_CHUNKS_ONLY)
  private Boolean chunksOnly = false;

  public static final String SERIALIZED_NAME_CUSTOM_CREDENTIALS = "custom_credentials";
  @SerializedName(SERIALIZED_NAME_CUSTOM_CREDENTIALS)
  private Map<String, Object> customCredentials = null;

  public ChunksAndEmbeddingsUploadInput() {
  }

  public ChunksAndEmbeddingsUploadInput embeddingModel(EmbeddingGenerators embeddingModel) {
    
    
    
    
    this.embeddingModel = embeddingModel;
    return this;
  }

   /**
   * Get embeddingModel
   * @return embeddingModel
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")

  public EmbeddingGenerators getEmbeddingModel() {
    return embeddingModel;
  }


  public void setEmbeddingModel(EmbeddingGenerators embeddingModel) {
    
    
    
    this.embeddingModel = embeddingModel;
  }


  public ChunksAndEmbeddingsUploadInput chunksAndEmbeddings(List<SingleChunksAndEmbeddingsUploadInput> chunksAndEmbeddings) {
    
    
    
    
    this.chunksAndEmbeddings = chunksAndEmbeddings;
    return this;
  }

  public ChunksAndEmbeddingsUploadInput addChunksAndEmbeddingsItem(SingleChunksAndEmbeddingsUploadInput chunksAndEmbeddingsItem) {
    this.chunksAndEmbeddings.add(chunksAndEmbeddingsItem);
    return this;
  }

   /**
   * Get chunksAndEmbeddings
   * @return chunksAndEmbeddings
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")

  public List<SingleChunksAndEmbeddingsUploadInput> getChunksAndEmbeddings() {
    return chunksAndEmbeddings;
  }


  public void setChunksAndEmbeddings(List<SingleChunksAndEmbeddingsUploadInput> chunksAndEmbeddings) {
    
    
    
    this.chunksAndEmbeddings = chunksAndEmbeddings;
  }


  public ChunksAndEmbeddingsUploadInput overwriteExisting(Boolean overwriteExisting) {
    
    
    
    
    this.overwriteExisting = overwriteExisting;
    return this;
  }

   /**
   * Get overwriteExisting
   * @return overwriteExisting
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "false", value = "")

  public Boolean getOverwriteExisting() {
    return overwriteExisting;
  }


  public void setOverwriteExisting(Boolean overwriteExisting) {
    
    
    
    this.overwriteExisting = overwriteExisting;
  }


  public ChunksAndEmbeddingsUploadInput chunksOnly(Boolean chunksOnly) {
    
    
    
    
    this.chunksOnly = chunksOnly;
    return this;
  }

   /**
   * Get chunksOnly
   * @return chunksOnly
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "false", value = "")

  public Boolean getChunksOnly() {
    return chunksOnly;
  }


  public void setChunksOnly(Boolean chunksOnly) {
    
    
    
    this.chunksOnly = chunksOnly;
  }


  public ChunksAndEmbeddingsUploadInput customCredentials(Map<String, Object> customCredentials) {
    
    
    
    
    this.customCredentials = customCredentials;
    return this;
  }

  public ChunksAndEmbeddingsUploadInput putCustomCredentialsItem(String key, Object customCredentialsItem) {
    if (this.customCredentials == null) {
      this.customCredentials = new HashMap<>();
    }
    this.customCredentials.put(key, customCredentialsItem);
    return this;
  }

   /**
   * Get customCredentials
   * @return customCredentials
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "{}", value = "")

  public Map<String, Object> getCustomCredentials() {
    return customCredentials;
  }


  public void setCustomCredentials(Map<String, Object> customCredentials) {
    
    
    
    this.customCredentials = customCredentials;
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
   * @return the ChunksAndEmbeddingsUploadInput instance itself
   */
  public ChunksAndEmbeddingsUploadInput putAdditionalProperty(String key, Object value) {
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
    ChunksAndEmbeddingsUploadInput chunksAndEmbeddingsUploadInput = (ChunksAndEmbeddingsUploadInput) o;
    return Objects.equals(this.embeddingModel, chunksAndEmbeddingsUploadInput.embeddingModel) &&
        Objects.equals(this.chunksAndEmbeddings, chunksAndEmbeddingsUploadInput.chunksAndEmbeddings) &&
        Objects.equals(this.overwriteExisting, chunksAndEmbeddingsUploadInput.overwriteExisting) &&
        Objects.equals(this.chunksOnly, chunksAndEmbeddingsUploadInput.chunksOnly) &&
        Objects.equals(this.customCredentials, chunksAndEmbeddingsUploadInput.customCredentials)&&
        Objects.equals(this.additionalProperties, chunksAndEmbeddingsUploadInput.additionalProperties);
  }

  @Override
  public int hashCode() {
    return Objects.hash(embeddingModel, chunksAndEmbeddings, overwriteExisting, chunksOnly, customCredentials, additionalProperties);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ChunksAndEmbeddingsUploadInput {\n");
    sb.append("    embeddingModel: ").append(toIndentedString(embeddingModel)).append("\n");
    sb.append("    chunksAndEmbeddings: ").append(toIndentedString(chunksAndEmbeddings)).append("\n");
    sb.append("    overwriteExisting: ").append(toIndentedString(overwriteExisting)).append("\n");
    sb.append("    chunksOnly: ").append(toIndentedString(chunksOnly)).append("\n");
    sb.append("    customCredentials: ").append(toIndentedString(customCredentials)).append("\n");
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
    openapiFields.add("embedding_model");
    openapiFields.add("chunks_and_embeddings");
    openapiFields.add("overwrite_existing");
    openapiFields.add("chunks_only");
    openapiFields.add("custom_credentials");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("embedding_model");
    openapiRequiredFields.add("chunks_and_embeddings");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to ChunksAndEmbeddingsUploadInput
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!ChunksAndEmbeddingsUploadInput.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in ChunksAndEmbeddingsUploadInput is not found in the empty JSON string", ChunksAndEmbeddingsUploadInput.openapiRequiredFields.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : ChunksAndEmbeddingsUploadInput.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      // ensure the json data is an array
      if (!jsonObj.get("chunks_and_embeddings").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `chunks_and_embeddings` to be an array in the JSON string but got `%s`", jsonObj.get("chunks_and_embeddings").toString()));
      }

      JsonArray jsonArraychunksAndEmbeddings = jsonObj.getAsJsonArray("chunks_and_embeddings");
      // validate the required field `chunks_and_embeddings` (array)
      for (int i = 0; i < jsonArraychunksAndEmbeddings.size(); i++) {
        SingleChunksAndEmbeddingsUploadInput.validateJsonObject(jsonArraychunksAndEmbeddings.get(i).getAsJsonObject());
      };
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!ChunksAndEmbeddingsUploadInput.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'ChunksAndEmbeddingsUploadInput' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<ChunksAndEmbeddingsUploadInput> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(ChunksAndEmbeddingsUploadInput.class));

       return (TypeAdapter<T>) new TypeAdapter<ChunksAndEmbeddingsUploadInput>() {
           @Override
           public void write(JsonWriter out, ChunksAndEmbeddingsUploadInput value) throws IOException {
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
           public ChunksAndEmbeddingsUploadInput read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             // store additional fields in the deserialized instance
             ChunksAndEmbeddingsUploadInput instance = thisAdapter.fromJsonTree(jsonObj);
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
  * Create an instance of ChunksAndEmbeddingsUploadInput given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of ChunksAndEmbeddingsUploadInput
  * @throws IOException if the JSON string is invalid with respect to ChunksAndEmbeddingsUploadInput
  */
  public static ChunksAndEmbeddingsUploadInput fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, ChunksAndEmbeddingsUploadInput.class);
  }

 /**
  * Convert an instance of ChunksAndEmbeddingsUploadInput to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}


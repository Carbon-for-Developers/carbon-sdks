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
 * ChunkProperties
 */@javax.annotation.Generated(value = "Generated by https://konfigthis.com")
public class ChunkProperties {
  public static final String SERIALIZED_NAME_SET_PAGE_AS_BOUNDARY = "set_page_as_boundary";
  @SerializedName(SERIALIZED_NAME_SET_PAGE_AS_BOUNDARY)
  private Boolean setPageAsBoundary = false;

  public static final String SERIALIZED_NAME_PREPEND_FILENAME_TO_CHUNKS = "prepend_filename_to_chunks";
  @SerializedName(SERIALIZED_NAME_PREPEND_FILENAME_TO_CHUNKS)
  private Boolean prependFilenameToChunks = false;

  public static final String SERIALIZED_NAME_MAX_ITEMS_PER_CHUNK = "max_items_per_chunk";
  @SerializedName(SERIALIZED_NAME_MAX_ITEMS_PER_CHUNK)
  private Integer maxItemsPerChunk;

  public ChunkProperties() {
  }

  public ChunkProperties setPageAsBoundary(Boolean setPageAsBoundary) {
    
    
    
    
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


  public ChunkProperties prependFilenameToChunks(Boolean prependFilenameToChunks) {
    
    
    
    
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


  public ChunkProperties maxItemsPerChunk(Integer maxItemsPerChunk) {
    
    
    
    
    this.maxItemsPerChunk = maxItemsPerChunk;
    return this;
  }

   /**
   * Get maxItemsPerChunk
   * @return maxItemsPerChunk
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getMaxItemsPerChunk() {
    return maxItemsPerChunk;
  }


  public void setMaxItemsPerChunk(Integer maxItemsPerChunk) {
    
    
    
    this.maxItemsPerChunk = maxItemsPerChunk;
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
   * @return the ChunkProperties instance itself
   */
  public ChunkProperties putAdditionalProperty(String key, Object value) {
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
    ChunkProperties chunkProperties = (ChunkProperties) o;
    return Objects.equals(this.setPageAsBoundary, chunkProperties.setPageAsBoundary) &&
        Objects.equals(this.prependFilenameToChunks, chunkProperties.prependFilenameToChunks) &&
        Objects.equals(this.maxItemsPerChunk, chunkProperties.maxItemsPerChunk)&&
        Objects.equals(this.additionalProperties, chunkProperties.additionalProperties);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(setPageAsBoundary, prependFilenameToChunks, maxItemsPerChunk, additionalProperties);
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
    sb.append("class ChunkProperties {\n");
    sb.append("    setPageAsBoundary: ").append(toIndentedString(setPageAsBoundary)).append("\n");
    sb.append("    prependFilenameToChunks: ").append(toIndentedString(prependFilenameToChunks)).append("\n");
    sb.append("    maxItemsPerChunk: ").append(toIndentedString(maxItemsPerChunk)).append("\n");
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
    openapiFields.add("set_page_as_boundary");
    openapiFields.add("prepend_filename_to_chunks");
    openapiFields.add("max_items_per_chunk");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to ChunkProperties
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!ChunkProperties.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in ChunkProperties is not found in the empty JSON string", ChunkProperties.openapiRequiredFields.toString()));
        }
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!ChunkProperties.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'ChunkProperties' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<ChunkProperties> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(ChunkProperties.class));

       return (TypeAdapter<T>) new TypeAdapter<ChunkProperties>() {
           @Override
           public void write(JsonWriter out, ChunkProperties value) throws IOException {
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
           public ChunkProperties read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             // store additional fields in the deserialized instance
             ChunkProperties instance = thisAdapter.fromJsonTree(jsonObj);
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
  * Create an instance of ChunkProperties given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of ChunkProperties
  * @throws IOException if the JSON string is invalid with respect to ChunkProperties
  */
  public static ChunkProperties fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, ChunkProperties.class);
  }

 /**
  * Convert an instance of ChunkProperties to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}


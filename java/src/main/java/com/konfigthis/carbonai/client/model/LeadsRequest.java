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
import com.konfigthis.carbonai.client.model.BaseIncludes;
import com.konfigthis.carbonai.client.model.LeadFilters;
import com.konfigthis.carbonai.client.model.LeadsOrderByNullable;
import com.konfigthis.carbonai.client.model.OrderDirV2Nullable;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
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
 * LeadsRequest
 */@javax.annotation.Generated(value = "Generated by https://konfigthis.com")
public class LeadsRequest {
  public static final String SERIALIZED_NAME_DATA_SOURCE_ID = "data_source_id";
  @SerializedName(SERIALIZED_NAME_DATA_SOURCE_ID)
  private Integer dataSourceId;

  public static final String SERIALIZED_NAME_INCLUDE_REMOTE_DATA = "include_remote_data";
  @SerializedName(SERIALIZED_NAME_INCLUDE_REMOTE_DATA)
  private Boolean includeRemoteData = false;

  public static final String SERIALIZED_NAME_NEXT_CURSOR = "next_cursor";
  @SerializedName(SERIALIZED_NAME_NEXT_CURSOR)
  private String nextCursor;

  public static final String SERIALIZED_NAME_PAGE_SIZE = "page_size";
  @SerializedName(SERIALIZED_NAME_PAGE_SIZE)
  private Integer pageSize;

  public static final String SERIALIZED_NAME_ORDER_DIR = "order_dir";
  @SerializedName(SERIALIZED_NAME_ORDER_DIR)
  private OrderDirV2Nullable orderDir = OrderDirV2Nullable.ASC;

  public static final String SERIALIZED_NAME_INCLUDES = "includes";
  @SerializedName(SERIALIZED_NAME_INCLUDES)
  private List<BaseIncludes> includes = null;

  public static final String SERIALIZED_NAME_FILTERS = "filters";
  @SerializedName(SERIALIZED_NAME_FILTERS)
  private LeadFilters filters;

  public static final String SERIALIZED_NAME_ORDER_BY = "order_by";
  @SerializedName(SERIALIZED_NAME_ORDER_BY)
  private LeadsOrderByNullable orderBy;

  public LeadsRequest() {
  }

  public LeadsRequest dataSourceId(Integer dataSourceId) {
    
    
    
    
    this.dataSourceId = dataSourceId;
    return this;
  }

   /**
   * Get dataSourceId
   * @return dataSourceId
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")

  public Integer getDataSourceId() {
    return dataSourceId;
  }


  public void setDataSourceId(Integer dataSourceId) {
    
    
    
    this.dataSourceId = dataSourceId;
  }


  public LeadsRequest includeRemoteData(Boolean includeRemoteData) {
    
    
    
    
    this.includeRemoteData = includeRemoteData;
    return this;
  }

   /**
   * Get includeRemoteData
   * @return includeRemoteData
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "false", value = "")

  public Boolean getIncludeRemoteData() {
    return includeRemoteData;
  }


  public void setIncludeRemoteData(Boolean includeRemoteData) {
    
    
    
    this.includeRemoteData = includeRemoteData;
  }


  public LeadsRequest nextCursor(String nextCursor) {
    
    
    
    
    this.nextCursor = nextCursor;
    return this;
  }

   /**
   * Get nextCursor
   * @return nextCursor
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getNextCursor() {
    return nextCursor;
  }


  public void setNextCursor(String nextCursor) {
    
    
    
    this.nextCursor = nextCursor;
  }


  public LeadsRequest pageSize(Integer pageSize) {
    
    
    
    
    this.pageSize = pageSize;
    return this;
  }

   /**
   * Get pageSize
   * @return pageSize
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getPageSize() {
    return pageSize;
  }


  public void setPageSize(Integer pageSize) {
    
    
    
    this.pageSize = pageSize;
  }


  public LeadsRequest orderDir(OrderDirV2Nullable orderDir) {
    
    
    
    
    this.orderDir = orderDir;
    return this;
  }

   /**
   * Get orderDir
   * @return orderDir
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OrderDirV2Nullable getOrderDir() {
    return orderDir;
  }


  public void setOrderDir(OrderDirV2Nullable orderDir) {
    
    
    
    this.orderDir = orderDir;
  }


  public LeadsRequest includes(List<BaseIncludes> includes) {
    
    
    
    
    this.includes = includes;
    return this;
  }

  public LeadsRequest addIncludesItem(BaseIncludes includesItem) {
    if (this.includes == null) {
      this.includes = new ArrayList<>();
    }
    this.includes.add(includesItem);
    return this;
  }

   /**
   * Get includes
   * @return includes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "[]", value = "")

  public List<BaseIncludes> getIncludes() {
    return includes;
  }


  public void setIncludes(List<BaseIncludes> includes) {
    
    
    
    this.includes = includes;
  }


  public LeadsRequest filters(LeadFilters filters) {
    
    
    
    
    this.filters = filters;
    return this;
  }

   /**
   * Get filters
   * @return filters
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public LeadFilters getFilters() {
    return filters;
  }


  public void setFilters(LeadFilters filters) {
    
    
    
    this.filters = filters;
  }


  public LeadsRequest orderBy(LeadsOrderByNullable orderBy) {
    
    
    
    
    this.orderBy = orderBy;
    return this;
  }

   /**
   * Get orderBy
   * @return orderBy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public LeadsOrderByNullable getOrderBy() {
    return orderBy;
  }


  public void setOrderBy(LeadsOrderByNullable orderBy) {
    
    
    
    this.orderBy = orderBy;
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
   * @return the LeadsRequest instance itself
   */
  public LeadsRequest putAdditionalProperty(String key, Object value) {
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
    LeadsRequest leadsRequest = (LeadsRequest) o;
    return Objects.equals(this.dataSourceId, leadsRequest.dataSourceId) &&
        Objects.equals(this.includeRemoteData, leadsRequest.includeRemoteData) &&
        Objects.equals(this.nextCursor, leadsRequest.nextCursor) &&
        Objects.equals(this.pageSize, leadsRequest.pageSize) &&
        Objects.equals(this.orderDir, leadsRequest.orderDir) &&
        Objects.equals(this.includes, leadsRequest.includes) &&
        Objects.equals(this.filters, leadsRequest.filters) &&
        Objects.equals(this.orderBy, leadsRequest.orderBy)&&
        Objects.equals(this.additionalProperties, leadsRequest.additionalProperties);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(dataSourceId, includeRemoteData, nextCursor, pageSize, orderDir, includes, filters, orderBy, additionalProperties);
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
    sb.append("class LeadsRequest {\n");
    sb.append("    dataSourceId: ").append(toIndentedString(dataSourceId)).append("\n");
    sb.append("    includeRemoteData: ").append(toIndentedString(includeRemoteData)).append("\n");
    sb.append("    nextCursor: ").append(toIndentedString(nextCursor)).append("\n");
    sb.append("    pageSize: ").append(toIndentedString(pageSize)).append("\n");
    sb.append("    orderDir: ").append(toIndentedString(orderDir)).append("\n");
    sb.append("    includes: ").append(toIndentedString(includes)).append("\n");
    sb.append("    filters: ").append(toIndentedString(filters)).append("\n");
    sb.append("    orderBy: ").append(toIndentedString(orderBy)).append("\n");
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
    openapiFields.add("data_source_id");
    openapiFields.add("include_remote_data");
    openapiFields.add("next_cursor");
    openapiFields.add("page_size");
    openapiFields.add("order_dir");
    openapiFields.add("includes");
    openapiFields.add("filters");
    openapiFields.add("order_by");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("data_source_id");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to LeadsRequest
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!LeadsRequest.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in LeadsRequest is not found in the empty JSON string", LeadsRequest.openapiRequiredFields.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : LeadsRequest.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if (!jsonObj.get("next_cursor").isJsonNull() && (jsonObj.get("next_cursor") != null && !jsonObj.get("next_cursor").isJsonNull()) && !jsonObj.get("next_cursor").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `next_cursor` to be a primitive type in the JSON string but got `%s`", jsonObj.get("next_cursor").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("includes") != null && !jsonObj.get("includes").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `includes` to be an array in the JSON string but got `%s`", jsonObj.get("includes").toString()));
      }
      // validate the optional field `filters`
      if (jsonObj.get("filters") != null && !jsonObj.get("filters").isJsonNull()) {
        LeadFilters.validateJsonObject(jsonObj.getAsJsonObject("filters"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!LeadsRequest.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'LeadsRequest' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<LeadsRequest> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(LeadsRequest.class));

       return (TypeAdapter<T>) new TypeAdapter<LeadsRequest>() {
           @Override
           public void write(JsonWriter out, LeadsRequest value) throws IOException {
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
           public LeadsRequest read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             // store additional fields in the deserialized instance
             LeadsRequest instance = thisAdapter.fromJsonTree(jsonObj);
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
  * Create an instance of LeadsRequest given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of LeadsRequest
  * @throws IOException if the JSON string is invalid with respect to LeadsRequest
  */
  public static LeadsRequest fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, LeadsRequest.class);
  }

 /**
  * Convert an instance of LeadsRequest to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}


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
import com.konfigthis.carbonai.client.model.OrderDir;
import com.konfigthis.carbonai.client.model.OrganizationUserFilesToSyncFilters;
import com.konfigthis.carbonai.client.model.OrganizationUserFilesToSyncOrderByTypes;
import com.konfigthis.carbonai.client.model.Pagination;
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
 * OrganizationUserFilesToSyncQueryInput
 */@javax.annotation.Generated(value = "Generated by https://konfigthis.com")
public class OrganizationUserFilesToSyncQueryInput {
  public static final String SERIALIZED_NAME_PAGINATION = "pagination";
  @SerializedName(SERIALIZED_NAME_PAGINATION)
  private Pagination pagination;

  public static final String SERIALIZED_NAME_ORDER_BY = "order_by";
  @SerializedName(SERIALIZED_NAME_ORDER_BY)
  private OrganizationUserFilesToSyncOrderByTypes orderBy;

  public static final String SERIALIZED_NAME_ORDER_DIR = "order_dir";
  @SerializedName(SERIALIZED_NAME_ORDER_DIR)
  private OrderDir orderDir;

  public static final String SERIALIZED_NAME_FILTERS = "filters";
  @SerializedName(SERIALIZED_NAME_FILTERS)
  private OrganizationUserFilesToSyncFilters filters;

  public static final String SERIALIZED_NAME_INCLUDE_RAW_FILE = "include_raw_file";
  @SerializedName(SERIALIZED_NAME_INCLUDE_RAW_FILE)
  private Boolean includeRawFile;

  public static final String SERIALIZED_NAME_INCLUDE_PARSED_TEXT_FILE = "include_parsed_text_file";
  @SerializedName(SERIALIZED_NAME_INCLUDE_PARSED_TEXT_FILE)
  private Boolean includeParsedTextFile;

  public static final String SERIALIZED_NAME_INCLUDE_ADDITIONAL_FILES = "include_additional_files";
  @SerializedName(SERIALIZED_NAME_INCLUDE_ADDITIONAL_FILES)
  private Boolean includeAdditionalFiles;

  public static final String SERIALIZED_NAME_PRESIGNED_URL_EXPIRY_TIME_SECONDS = "presigned_url_expiry_time_seconds";
  @SerializedName(SERIALIZED_NAME_PRESIGNED_URL_EXPIRY_TIME_SECONDS)
  private Integer presignedUrlExpiryTimeSeconds = 3600;

  public OrganizationUserFilesToSyncQueryInput() {
  }

  public OrganizationUserFilesToSyncQueryInput pagination(Pagination pagination) {
    
    
    
    
    this.pagination = pagination;
    return this;
  }

   /**
   * Get pagination
   * @return pagination
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Pagination getPagination() {
    return pagination;
  }


  public void setPagination(Pagination pagination) {
    
    
    
    this.pagination = pagination;
  }


  public OrganizationUserFilesToSyncQueryInput orderBy(OrganizationUserFilesToSyncOrderByTypes orderBy) {
    
    
    
    
    this.orderBy = orderBy;
    return this;
  }

   /**
   * Get orderBy
   * @return orderBy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OrganizationUserFilesToSyncOrderByTypes getOrderBy() {
    return orderBy;
  }


  public void setOrderBy(OrganizationUserFilesToSyncOrderByTypes orderBy) {
    
    
    
    this.orderBy = orderBy;
  }


  public OrganizationUserFilesToSyncQueryInput orderDir(OrderDir orderDir) {
    
    
    
    
    this.orderDir = orderDir;
    return this;
  }

   /**
   * Get orderDir
   * @return orderDir
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OrderDir getOrderDir() {
    return orderDir;
  }


  public void setOrderDir(OrderDir orderDir) {
    
    
    
    this.orderDir = orderDir;
  }


  public OrganizationUserFilesToSyncQueryInput filters(OrganizationUserFilesToSyncFilters filters) {
    
    
    
    
    this.filters = filters;
    return this;
  }

   /**
   * Get filters
   * @return filters
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OrganizationUserFilesToSyncFilters getFilters() {
    return filters;
  }


  public void setFilters(OrganizationUserFilesToSyncFilters filters) {
    
    
    
    this.filters = filters;
  }


  public OrganizationUserFilesToSyncQueryInput includeRawFile(Boolean includeRawFile) {
    
    
    
    
    this.includeRawFile = includeRawFile;
    return this;
  }

   /**
   * If true, the query will return presigned URLs for the raw file. Only relevant for the /user_files_v2 endpoint.
   * @return includeRawFile
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "If true, the query will return presigned URLs for the raw file. Only relevant for the /user_files_v2 endpoint.")

  public Boolean getIncludeRawFile() {
    return includeRawFile;
  }


  public void setIncludeRawFile(Boolean includeRawFile) {
    
    
    
    this.includeRawFile = includeRawFile;
  }


  public OrganizationUserFilesToSyncQueryInput includeParsedTextFile(Boolean includeParsedTextFile) {
    
    
    
    
    this.includeParsedTextFile = includeParsedTextFile;
    return this;
  }

   /**
   * If true, the query will return presigned URLs for the parsed text file. Only relevant for the /user_files_v2 endpoint.
   * @return includeParsedTextFile
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "If true, the query will return presigned URLs for the parsed text file. Only relevant for the /user_files_v2 endpoint.")

  public Boolean getIncludeParsedTextFile() {
    return includeParsedTextFile;
  }


  public void setIncludeParsedTextFile(Boolean includeParsedTextFile) {
    
    
    
    this.includeParsedTextFile = includeParsedTextFile;
  }


  public OrganizationUserFilesToSyncQueryInput includeAdditionalFiles(Boolean includeAdditionalFiles) {
    
    
    
    
    this.includeAdditionalFiles = includeAdditionalFiles;
    return this;
  }

   /**
   * If true, the query will return presigned URLs for additional files. Only relevant for the /user_files_v2 endpoint.
   * @return includeAdditionalFiles
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "If true, the query will return presigned URLs for additional files. Only relevant for the /user_files_v2 endpoint.")

  public Boolean getIncludeAdditionalFiles() {
    return includeAdditionalFiles;
  }


  public void setIncludeAdditionalFiles(Boolean includeAdditionalFiles) {
    
    
    
    this.includeAdditionalFiles = includeAdditionalFiles;
  }


  public OrganizationUserFilesToSyncQueryInput presignedUrlExpiryTimeSeconds(Integer presignedUrlExpiryTimeSeconds) {
    
    
    
    
    this.presignedUrlExpiryTimeSeconds = presignedUrlExpiryTimeSeconds;
    return this;
  }

   /**
   * The expiry time for the presigned URLs. Only relevant for the /user_files_v2 endpoint.
   * @return presignedUrlExpiryTimeSeconds
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "3600", value = "The expiry time for the presigned URLs. Only relevant for the /user_files_v2 endpoint.")

  public Integer getPresignedUrlExpiryTimeSeconds() {
    return presignedUrlExpiryTimeSeconds;
  }


  public void setPresignedUrlExpiryTimeSeconds(Integer presignedUrlExpiryTimeSeconds) {
    
    
    
    this.presignedUrlExpiryTimeSeconds = presignedUrlExpiryTimeSeconds;
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
   * @return the OrganizationUserFilesToSyncQueryInput instance itself
   */
  public OrganizationUserFilesToSyncQueryInput putAdditionalProperty(String key, Object value) {
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
    OrganizationUserFilesToSyncQueryInput organizationUserFilesToSyncQueryInput = (OrganizationUserFilesToSyncQueryInput) o;
    return Objects.equals(this.pagination, organizationUserFilesToSyncQueryInput.pagination) &&
        Objects.equals(this.orderBy, organizationUserFilesToSyncQueryInput.orderBy) &&
        Objects.equals(this.orderDir, organizationUserFilesToSyncQueryInput.orderDir) &&
        Objects.equals(this.filters, organizationUserFilesToSyncQueryInput.filters) &&
        Objects.equals(this.includeRawFile, organizationUserFilesToSyncQueryInput.includeRawFile) &&
        Objects.equals(this.includeParsedTextFile, organizationUserFilesToSyncQueryInput.includeParsedTextFile) &&
        Objects.equals(this.includeAdditionalFiles, organizationUserFilesToSyncQueryInput.includeAdditionalFiles) &&
        Objects.equals(this.presignedUrlExpiryTimeSeconds, organizationUserFilesToSyncQueryInput.presignedUrlExpiryTimeSeconds)&&
        Objects.equals(this.additionalProperties, organizationUserFilesToSyncQueryInput.additionalProperties);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(pagination, orderBy, orderDir, filters, includeRawFile, includeParsedTextFile, includeAdditionalFiles, presignedUrlExpiryTimeSeconds, additionalProperties);
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
    sb.append("class OrganizationUserFilesToSyncQueryInput {\n");
    sb.append("    pagination: ").append(toIndentedString(pagination)).append("\n");
    sb.append("    orderBy: ").append(toIndentedString(orderBy)).append("\n");
    sb.append("    orderDir: ").append(toIndentedString(orderDir)).append("\n");
    sb.append("    filters: ").append(toIndentedString(filters)).append("\n");
    sb.append("    includeRawFile: ").append(toIndentedString(includeRawFile)).append("\n");
    sb.append("    includeParsedTextFile: ").append(toIndentedString(includeParsedTextFile)).append("\n");
    sb.append("    includeAdditionalFiles: ").append(toIndentedString(includeAdditionalFiles)).append("\n");
    sb.append("    presignedUrlExpiryTimeSeconds: ").append(toIndentedString(presignedUrlExpiryTimeSeconds)).append("\n");
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
    openapiFields.add("pagination");
    openapiFields.add("order_by");
    openapiFields.add("order_dir");
    openapiFields.add("filters");
    openapiFields.add("include_raw_file");
    openapiFields.add("include_parsed_text_file");
    openapiFields.add("include_additional_files");
    openapiFields.add("presigned_url_expiry_time_seconds");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to OrganizationUserFilesToSyncQueryInput
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!OrganizationUserFilesToSyncQueryInput.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in OrganizationUserFilesToSyncQueryInput is not found in the empty JSON string", OrganizationUserFilesToSyncQueryInput.openapiRequiredFields.toString()));
        }
      }
      // validate the optional field `pagination`
      if (jsonObj.get("pagination") != null && !jsonObj.get("pagination").isJsonNull()) {
        Pagination.validateJsonObject(jsonObj.getAsJsonObject("pagination"));
      }
      // validate the optional field `filters`
      if (jsonObj.get("filters") != null && !jsonObj.get("filters").isJsonNull()) {
        OrganizationUserFilesToSyncFilters.validateJsonObject(jsonObj.getAsJsonObject("filters"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!OrganizationUserFilesToSyncQueryInput.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'OrganizationUserFilesToSyncQueryInput' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<OrganizationUserFilesToSyncQueryInput> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(OrganizationUserFilesToSyncQueryInput.class));

       return (TypeAdapter<T>) new TypeAdapter<OrganizationUserFilesToSyncQueryInput>() {
           @Override
           public void write(JsonWriter out, OrganizationUserFilesToSyncQueryInput value) throws IOException {
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
           public OrganizationUserFilesToSyncQueryInput read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             // store additional fields in the deserialized instance
             OrganizationUserFilesToSyncQueryInput instance = thisAdapter.fromJsonTree(jsonObj);
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
  * Create an instance of OrganizationUserFilesToSyncQueryInput given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of OrganizationUserFilesToSyncQueryInput
  * @throws IOException if the JSON string is invalid with respect to OrganizationUserFilesToSyncQueryInput
  */
  public static OrganizationUserFilesToSyncQueryInput fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, OrganizationUserFilesToSyncQueryInput.class);
  }

 /**
  * Convert an instance of OrganizationUserFilesToSyncQueryInput to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}


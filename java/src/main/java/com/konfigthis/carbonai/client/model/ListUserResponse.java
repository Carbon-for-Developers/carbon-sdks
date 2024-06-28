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
import java.time.OffsetDateTime;
import java.util.ArrayList;
import java.util.List;

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
 * ListUserResponse
 */@javax.annotation.Generated(value = "Generated by https://konfigthis.com")
public class ListUserResponse {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Integer id;

  public static final String SERIALIZED_NAME_ORGANIZATION_ID = "organization_id";
  @SerializedName(SERIALIZED_NAME_ORGANIZATION_ID)
  private Integer organizationId;

  public static final String SERIALIZED_NAME_ORGANIZATION_SUPPLIED_USER_ID = "organization_supplied_user_id";
  @SerializedName(SERIALIZED_NAME_ORGANIZATION_SUPPLIED_USER_ID)
  private String organizationSuppliedUserId;

  public static final String SERIALIZED_NAME_CREATED_AT = "created_at";
  @SerializedName(SERIALIZED_NAME_CREATED_AT)
  private OffsetDateTime createdAt;

  public static final String SERIALIZED_NAME_UPDATED_AT = "updated_at";
  @SerializedName(SERIALIZED_NAME_UPDATED_AT)
  private OffsetDateTime updatedAt;

  public static final String SERIALIZED_NAME_DELETED_AT = "deleted_at";
  @SerializedName(SERIALIZED_NAME_DELETED_AT)
  private OffsetDateTime deletedAt;

  public static final String SERIALIZED_NAME_ENABLED_FEATURES = "enabled_features";
  @SerializedName(SERIALIZED_NAME_ENABLED_FEATURES)
  private Object enabledFeatures;

  public static final String SERIALIZED_NAME_CUSTOM_LIMITS = "custom_limits";
  @SerializedName(SERIALIZED_NAME_CUSTOM_LIMITS)
  private Object customLimits;

  public static final String SERIALIZED_NAME_AUTO_SYNC_ENABLED_SOURCES = "auto_sync_enabled_sources";
  @SerializedName(SERIALIZED_NAME_AUTO_SYNC_ENABLED_SOURCES)
  private List<Object> autoSyncEnabledSources = new ArrayList<>();

  public ListUserResponse() {
  }

  public ListUserResponse id(Integer id) {
    
    
    
    
    this.id = id;
    return this;
  }

   /**
   * Get id
   * @return id
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")

  public Integer getId() {
    return id;
  }


  public void setId(Integer id) {
    
    
    
    this.id = id;
  }


  public ListUserResponse organizationId(Integer organizationId) {
    
    
    
    
    this.organizationId = organizationId;
    return this;
  }

   /**
   * Get organizationId
   * @return organizationId
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")

  public Integer getOrganizationId() {
    return organizationId;
  }


  public void setOrganizationId(Integer organizationId) {
    
    
    
    this.organizationId = organizationId;
  }


  public ListUserResponse organizationSuppliedUserId(String organizationSuppliedUserId) {
    
    
    
    
    this.organizationSuppliedUserId = organizationSuppliedUserId;
    return this;
  }

   /**
   * Get organizationSuppliedUserId
   * @return organizationSuppliedUserId
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")

  public String getOrganizationSuppliedUserId() {
    return organizationSuppliedUserId;
  }


  public void setOrganizationSuppliedUserId(String organizationSuppliedUserId) {
    
    
    
    this.organizationSuppliedUserId = organizationSuppliedUserId;
  }


  public ListUserResponse createdAt(OffsetDateTime createdAt) {
    
    
    
    
    this.createdAt = createdAt;
    return this;
  }

   /**
   * Get createdAt
   * @return createdAt
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")

  public OffsetDateTime getCreatedAt() {
    return createdAt;
  }


  public void setCreatedAt(OffsetDateTime createdAt) {
    
    
    
    this.createdAt = createdAt;
  }


  public ListUserResponse updatedAt(OffsetDateTime updatedAt) {
    
    
    
    
    this.updatedAt = updatedAt;
    return this;
  }

   /**
   * Get updatedAt
   * @return updatedAt
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")

  public OffsetDateTime getUpdatedAt() {
    return updatedAt;
  }


  public void setUpdatedAt(OffsetDateTime updatedAt) {
    
    
    
    this.updatedAt = updatedAt;
  }


  public ListUserResponse deletedAt(OffsetDateTime deletedAt) {
    
    
    
    
    this.deletedAt = deletedAt;
    return this;
  }

   /**
   * Get deletedAt
   * @return deletedAt
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(required = true, value = "")

  public OffsetDateTime getDeletedAt() {
    return deletedAt;
  }


  public void setDeletedAt(OffsetDateTime deletedAt) {
    
    
    
    this.deletedAt = deletedAt;
  }


  public ListUserResponse enabledFeatures(Object enabledFeatures) {
    
    
    
    
    this.enabledFeatures = enabledFeatures;
    return this;
  }

   /**
   * Get enabledFeatures
   * @return enabledFeatures
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(required = true, value = "")

  public Object getEnabledFeatures() {
    return enabledFeatures;
  }


  public void setEnabledFeatures(Object enabledFeatures) {
    
    
    
    this.enabledFeatures = enabledFeatures;
  }


  public ListUserResponse customLimits(Object customLimits) {
    
    
    
    
    this.customLimits = customLimits;
    return this;
  }

   /**
   * Get customLimits
   * @return customLimits
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(example = "{}", required = true, value = "")

  public Object getCustomLimits() {
    return customLimits;
  }


  public void setCustomLimits(Object customLimits) {
    
    
    
    this.customLimits = customLimits;
  }


  public ListUserResponse autoSyncEnabledSources(List<Object> autoSyncEnabledSources) {
    
    
    
    
    this.autoSyncEnabledSources = autoSyncEnabledSources;
    return this;
  }

  public ListUserResponse addAutoSyncEnabledSourcesItem(Object autoSyncEnabledSourcesItem) {
    this.autoSyncEnabledSources.add(autoSyncEnabledSourcesItem);
    return this;
  }

   /**
   * Get autoSyncEnabledSources
   * @return autoSyncEnabledSources
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(example = "[]", required = true, value = "")

  public List<Object> getAutoSyncEnabledSources() {
    return autoSyncEnabledSources;
  }


  public void setAutoSyncEnabledSources(List<Object> autoSyncEnabledSources) {
    
    
    
    this.autoSyncEnabledSources = autoSyncEnabledSources;
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
   * @return the ListUserResponse instance itself
   */
  public ListUserResponse putAdditionalProperty(String key, Object value) {
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
    ListUserResponse listUserResponse = (ListUserResponse) o;
    return Objects.equals(this.id, listUserResponse.id) &&
        Objects.equals(this.organizationId, listUserResponse.organizationId) &&
        Objects.equals(this.organizationSuppliedUserId, listUserResponse.organizationSuppliedUserId) &&
        Objects.equals(this.createdAt, listUserResponse.createdAt) &&
        Objects.equals(this.updatedAt, listUserResponse.updatedAt) &&
        Objects.equals(this.deletedAt, listUserResponse.deletedAt) &&
        Objects.equals(this.enabledFeatures, listUserResponse.enabledFeatures) &&
        Objects.equals(this.customLimits, listUserResponse.customLimits) &&
        Objects.equals(this.autoSyncEnabledSources, listUserResponse.autoSyncEnabledSources)&&
        Objects.equals(this.additionalProperties, listUserResponse.additionalProperties);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, organizationId, organizationSuppliedUserId, createdAt, updatedAt, deletedAt, enabledFeatures, customLimits, autoSyncEnabledSources, additionalProperties);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ListUserResponse {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    organizationId: ").append(toIndentedString(organizationId)).append("\n");
    sb.append("    organizationSuppliedUserId: ").append(toIndentedString(organizationSuppliedUserId)).append("\n");
    sb.append("    createdAt: ").append(toIndentedString(createdAt)).append("\n");
    sb.append("    updatedAt: ").append(toIndentedString(updatedAt)).append("\n");
    sb.append("    deletedAt: ").append(toIndentedString(deletedAt)).append("\n");
    sb.append("    enabledFeatures: ").append(toIndentedString(enabledFeatures)).append("\n");
    sb.append("    customLimits: ").append(toIndentedString(customLimits)).append("\n");
    sb.append("    autoSyncEnabledSources: ").append(toIndentedString(autoSyncEnabledSources)).append("\n");
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
    openapiFields.add("id");
    openapiFields.add("organization_id");
    openapiFields.add("organization_supplied_user_id");
    openapiFields.add("created_at");
    openapiFields.add("updated_at");
    openapiFields.add("deleted_at");
    openapiFields.add("enabled_features");
    openapiFields.add("custom_limits");
    openapiFields.add("auto_sync_enabled_sources");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("id");
    openapiRequiredFields.add("organization_id");
    openapiRequiredFields.add("organization_supplied_user_id");
    openapiRequiredFields.add("created_at");
    openapiRequiredFields.add("updated_at");
    openapiRequiredFields.add("deleted_at");
    openapiRequiredFields.add("enabled_features");
    openapiRequiredFields.add("custom_limits");
    openapiRequiredFields.add("auto_sync_enabled_sources");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to ListUserResponse
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!ListUserResponse.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in ListUserResponse is not found in the empty JSON string", ListUserResponse.openapiRequiredFields.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : ListUserResponse.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if (!jsonObj.get("organization_supplied_user_id").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `organization_supplied_user_id` to be a primitive type in the JSON string but got `%s`", jsonObj.get("organization_supplied_user_id").toString()));
      }
      // ensure the required json array is present
      if (jsonObj.get("auto_sync_enabled_sources") == null) {
        throw new IllegalArgumentException("Expected the field `linkedContent` to be an array in the JSON string but got `null`");
      } else if (!jsonObj.get("auto_sync_enabled_sources").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `auto_sync_enabled_sources` to be an array in the JSON string but got `%s`", jsonObj.get("auto_sync_enabled_sources").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!ListUserResponse.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'ListUserResponse' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<ListUserResponse> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(ListUserResponse.class));

       return (TypeAdapter<T>) new TypeAdapter<ListUserResponse>() {
           @Override
           public void write(JsonWriter out, ListUserResponse value) throws IOException {
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
                 else {
                   obj.add(entry.getKey(), gson.toJsonTree(entry.getValue()).getAsJsonObject());
                 }
               }
             }
             elementAdapter.write(out, obj);
           }

           @Override
           public ListUserResponse read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             // store additional fields in the deserialized instance
             ListUserResponse instance = thisAdapter.fromJsonTree(jsonObj);
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
  * Create an instance of ListUserResponse given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of ListUserResponse
  * @throws IOException if the JSON string is invalid with respect to ListUserResponse
  */
  public static ListUserResponse fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, ListUserResponse.class);
  }

 /**
  * Convert an instance of ListUserResponse to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}


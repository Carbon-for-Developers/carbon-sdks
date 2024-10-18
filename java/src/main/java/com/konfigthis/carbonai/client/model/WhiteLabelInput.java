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
import com.konfigthis.carbonai.client.model.Credentials;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

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
 * WhiteLabelInput
 */@javax.annotation.Generated(value = "Generated by https://konfigthis.com")
public class WhiteLabelInput {
  /**
   * Gets or Sets dataSourceType
   */
  @JsonAdapter(DataSourceTypeEnum.Adapter.class)
 public enum DataSourceTypeEnum {
    INTERCOM("INTERCOM"),
    
    NOTION("NOTION"),
    
    SLACK("SLACK"),
    
    ZENDESK("ZENDESK"),
    
    OUTLOOK("OUTLOOK"),
    
    GMAIL("GMAIL"),
    
    SERVICENOW("SERVICENOW"),
    
    SALESFORCE("SALESFORCE"),
    
    ZOTERO("ZOTERO"),
    
    BOX("BOX"),
    
    CONFLUENCE("CONFLUENCE"),
    
    DROPBOX("DROPBOX"),
    
    GOOGLE_CLOUD_STORAGE("GOOGLE_CLOUD_STORAGE"),
    
    GONG("GONG");

    private String value;

    DataSourceTypeEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static DataSourceTypeEnum fromValue(String value) {
      for (DataSourceTypeEnum b : DataSourceTypeEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<DataSourceTypeEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final DataSourceTypeEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public DataSourceTypeEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return DataSourceTypeEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_DATA_SOURCE_TYPE = "data_source_type";
  @SerializedName(SERIALIZED_NAME_DATA_SOURCE_TYPE)
  private DataSourceTypeEnum dataSourceType;

  public static final String SERIALIZED_NAME_CREDENTIALS = "credentials";
  @SerializedName(SERIALIZED_NAME_CREDENTIALS)
  private Credentials credentials;

  public WhiteLabelInput() {
  }

  public WhiteLabelInput dataSourceType(DataSourceTypeEnum dataSourceType) {
    
    
    
    
    this.dataSourceType = dataSourceType;
    return this;
  }

   /**
   * Get dataSourceType
   * @return dataSourceType
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")

  public DataSourceTypeEnum getDataSourceType() {
    return dataSourceType;
  }


  public void setDataSourceType(DataSourceTypeEnum dataSourceType) {
    
    
    
    this.dataSourceType = dataSourceType;
  }


  public WhiteLabelInput credentials(Credentials credentials) {
    
    
    
    
    this.credentials = credentials;
    return this;
  }

   /**
   * Get credentials
   * @return credentials
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")

  public Credentials getCredentials() {
    return credentials;
  }


  public void setCredentials(Credentials credentials) {
    
    
    
    this.credentials = credentials;
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
   * @return the WhiteLabelInput instance itself
   */
  public WhiteLabelInput putAdditionalProperty(String key, Object value) {
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
    WhiteLabelInput whiteLabelInput = (WhiteLabelInput) o;
    return Objects.equals(this.dataSourceType, whiteLabelInput.dataSourceType) &&
        Objects.equals(this.credentials, whiteLabelInput.credentials)&&
        Objects.equals(this.additionalProperties, whiteLabelInput.additionalProperties);
  }

  @Override
  public int hashCode() {
    return Objects.hash(dataSourceType, credentials, additionalProperties);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class WhiteLabelInput {\n");
    sb.append("    dataSourceType: ").append(toIndentedString(dataSourceType)).append("\n");
    sb.append("    credentials: ").append(toIndentedString(credentials)).append("\n");
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
    openapiFields.add("data_source_type");
    openapiFields.add("credentials");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("data_source_type");
    openapiRequiredFields.add("credentials");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to WhiteLabelInput
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!WhiteLabelInput.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in WhiteLabelInput is not found in the empty JSON string", WhiteLabelInput.openapiRequiredFields.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : WhiteLabelInput.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if (!jsonObj.get("data_source_type").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `data_source_type` to be a primitive type in the JSON string but got `%s`", jsonObj.get("data_source_type").toString()));
      }
      // validate the required field `credentials`
      Credentials.validateJsonObject(jsonObj.getAsJsonObject("credentials"));
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!WhiteLabelInput.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'WhiteLabelInput' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<WhiteLabelInput> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(WhiteLabelInput.class));

       return (TypeAdapter<T>) new TypeAdapter<WhiteLabelInput>() {
           @Override
           public void write(JsonWriter out, WhiteLabelInput value) throws IOException {
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
           public WhiteLabelInput read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             // store additional fields in the deserialized instance
             WhiteLabelInput instance = thisAdapter.fromJsonTree(jsonObj);
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
  * Create an instance of WhiteLabelInput given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of WhiteLabelInput
  * @throws IOException if the JSON string is invalid with respect to WhiteLabelInput
  */
  public static WhiteLabelInput fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, WhiteLabelInput.class);
  }

 /**
  * Convert an instance of WhiteLabelInput to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}


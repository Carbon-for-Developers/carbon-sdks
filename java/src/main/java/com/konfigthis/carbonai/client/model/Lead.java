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
import com.konfigthis.carbonai.client.model.Address;
import com.konfigthis.carbonai.client.model.Email;
import com.konfigthis.carbonai.client.model.Event;
import com.konfigthis.carbonai.client.model.PartialAccountNullable;
import com.konfigthis.carbonai.client.model.PartialContactNullable;
import com.konfigthis.carbonai.client.model.PartialOwner;
import com.konfigthis.carbonai.client.model.PhoneNumber;
import com.konfigthis.carbonai.client.model.Task;
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
 * Lead
 */@javax.annotation.Generated(value = "Generated by https://konfigthis.com")
public class Lead {
  public static final String SERIALIZED_NAME_TITLE = "title";
  @SerializedName(SERIALIZED_NAME_TITLE)
  private String title;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private String id;

  public static final String SERIALIZED_NAME_OWNER = "owner";
  @SerializedName(SERIALIZED_NAME_OWNER)
  private PartialOwner owner;

  public static final String SERIALIZED_NAME_SOURCE = "source";
  @SerializedName(SERIALIZED_NAME_SOURCE)
  private String source;

  public static final String SERIALIZED_NAME_STATUS = "status";
  @SerializedName(SERIALIZED_NAME_STATUS)
  private String status;

  public static final String SERIALIZED_NAME_COMPANY = "company";
  @SerializedName(SERIALIZED_NAME_COMPANY)
  private String company;

  public static final String SERIALIZED_NAME_FIRST_NAME = "first_name";
  @SerializedName(SERIALIZED_NAME_FIRST_NAME)
  private String firstName;

  public static final String SERIALIZED_NAME_LAST_NAME = "last_name";
  @SerializedName(SERIALIZED_NAME_LAST_NAME)
  private String lastName;

  public static final String SERIALIZED_NAME_ADDRESSES = "addresses";
  @SerializedName(SERIALIZED_NAME_ADDRESSES)
  private List<Address> addresses = new ArrayList<>();

  public static final String SERIALIZED_NAME_PHONE_NUMBERS = "phone_numbers";
  @SerializedName(SERIALIZED_NAME_PHONE_NUMBERS)
  private List<PhoneNumber> phoneNumbers = new ArrayList<>();

  public static final String SERIALIZED_NAME_EMAILS = "emails";
  @SerializedName(SERIALIZED_NAME_EMAILS)
  private List<Email> emails = new ArrayList<>();

  public static final String SERIALIZED_NAME_CONVERTED_AT = "converted_at";
  @SerializedName(SERIALIZED_NAME_CONVERTED_AT)
  private String convertedAt;

  public static final String SERIALIZED_NAME_CONVERTED_ACCOUNT = "converted_account";
  @SerializedName(SERIALIZED_NAME_CONVERTED_ACCOUNT)
  private PartialAccountNullable convertedAccount;

  public static final String SERIALIZED_NAME_CONVERTED_CONTACT = "converted_contact";
  @SerializedName(SERIALIZED_NAME_CONVERTED_CONTACT)
  private PartialContactNullable convertedContact;

  public static final String SERIALIZED_NAME_LAST_ACTIVITY_AT = "last_activity_at";
  @SerializedName(SERIALIZED_NAME_LAST_ACTIVITY_AT)
  private String lastActivityAt;

  public static final String SERIALIZED_NAME_CREATED_AT = "created_at";
  @SerializedName(SERIALIZED_NAME_CREATED_AT)
  private String createdAt;

  public static final String SERIALIZED_NAME_UPDATED_AT = "updated_at";
  @SerializedName(SERIALIZED_NAME_UPDATED_AT)
  private String updatedAt;

  public static final String SERIALIZED_NAME_IS_DELETED = "is_deleted";
  @SerializedName(SERIALIZED_NAME_IS_DELETED)
  private Boolean isDeleted;

  public static final String SERIALIZED_NAME_TASKS = "tasks";
  @SerializedName(SERIALIZED_NAME_TASKS)
  private List<Task> tasks = null;

  public static final String SERIALIZED_NAME_EVENTS = "events";
  @SerializedName(SERIALIZED_NAME_EVENTS)
  private List<Event> events = null;

  public static final String SERIALIZED_NAME_REMOTE_DATA = "remote_data";
  @SerializedName(SERIALIZED_NAME_REMOTE_DATA)
  private Object remoteData;

  public Lead() {
  }

  public Lead title(String title) {
    
    
    
    
    this.title = title;
    return this;
  }

   /**
   * Get title
   * @return title
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(required = true, value = "")

  public String getTitle() {
    return title;
  }


  public void setTitle(String title) {
    
    
    
    this.title = title;
  }


  public Lead description(String description) {
    
    
    
    
    this.description = description;
    return this;
  }

   /**
   * Get description
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(required = true, value = "")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    
    
    
    this.description = description;
  }


  public Lead id(String id) {
    
    
    
    
    this.id = id;
    return this;
  }

   /**
   * Get id
   * @return id
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")

  public String getId() {
    return id;
  }


  public void setId(String id) {
    
    
    
    this.id = id;
  }


  public Lead owner(PartialOwner owner) {
    
    
    
    
    this.owner = owner;
    return this;
  }

   /**
   * Get owner
   * @return owner
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")

  public PartialOwner getOwner() {
    return owner;
  }


  public void setOwner(PartialOwner owner) {
    
    
    
    this.owner = owner;
  }


  public Lead source(String source) {
    
    
    
    
    this.source = source;
    return this;
  }

   /**
   * Get source
   * @return source
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(required = true, value = "")

  public String getSource() {
    return source;
  }


  public void setSource(String source) {
    
    
    
    this.source = source;
  }


  public Lead status(String status) {
    
    
    
    
    this.status = status;
    return this;
  }

   /**
   * Get status
   * @return status
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(required = true, value = "")

  public String getStatus() {
    return status;
  }


  public void setStatus(String status) {
    
    
    
    this.status = status;
  }


  public Lead company(String company) {
    
    
    
    
    this.company = company;
    return this;
  }

   /**
   * Get company
   * @return company
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(required = true, value = "")

  public String getCompany() {
    return company;
  }


  public void setCompany(String company) {
    
    
    
    this.company = company;
  }


  public Lead firstName(String firstName) {
    
    
    
    
    this.firstName = firstName;
    return this;
  }

   /**
   * Get firstName
   * @return firstName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(required = true, value = "")

  public String getFirstName() {
    return firstName;
  }


  public void setFirstName(String firstName) {
    
    
    
    this.firstName = firstName;
  }


  public Lead lastName(String lastName) {
    
    
    
    
    this.lastName = lastName;
    return this;
  }

   /**
   * Get lastName
   * @return lastName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(required = true, value = "")

  public String getLastName() {
    return lastName;
  }


  public void setLastName(String lastName) {
    
    
    
    this.lastName = lastName;
  }


  public Lead addresses(List<Address> addresses) {
    
    
    
    
    this.addresses = addresses;
    return this;
  }

  public Lead addAddressesItem(Address addressesItem) {
    this.addresses.add(addressesItem);
    return this;
  }

   /**
   * Get addresses
   * @return addresses
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")

  public List<Address> getAddresses() {
    return addresses;
  }


  public void setAddresses(List<Address> addresses) {
    
    
    
    this.addresses = addresses;
  }


  public Lead phoneNumbers(List<PhoneNumber> phoneNumbers) {
    
    
    
    
    this.phoneNumbers = phoneNumbers;
    return this;
  }

  public Lead addPhoneNumbersItem(PhoneNumber phoneNumbersItem) {
    this.phoneNumbers.add(phoneNumbersItem);
    return this;
  }

   /**
   * Get phoneNumbers
   * @return phoneNumbers
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")

  public List<PhoneNumber> getPhoneNumbers() {
    return phoneNumbers;
  }


  public void setPhoneNumbers(List<PhoneNumber> phoneNumbers) {
    
    
    
    this.phoneNumbers = phoneNumbers;
  }


  public Lead emails(List<Email> emails) {
    
    
    
    
    this.emails = emails;
    return this;
  }

  public Lead addEmailsItem(Email emailsItem) {
    this.emails.add(emailsItem);
    return this;
  }

   /**
   * Get emails
   * @return emails
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")

  public List<Email> getEmails() {
    return emails;
  }


  public void setEmails(List<Email> emails) {
    
    
    
    this.emails = emails;
  }


  public Lead convertedAt(String convertedAt) {
    
    
    
    
    this.convertedAt = convertedAt;
    return this;
  }

   /**
   * Get convertedAt
   * @return convertedAt
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(required = true, value = "")

  public String getConvertedAt() {
    return convertedAt;
  }


  public void setConvertedAt(String convertedAt) {
    
    
    
    this.convertedAt = convertedAt;
  }


  public Lead convertedAccount(PartialAccountNullable convertedAccount) {
    
    
    
    
    this.convertedAccount = convertedAccount;
    return this;
  }

   /**
   * Get convertedAccount
   * @return convertedAccount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(required = true, value = "")

  public PartialAccountNullable getConvertedAccount() {
    return convertedAccount;
  }


  public void setConvertedAccount(PartialAccountNullable convertedAccount) {
    
    
    
    this.convertedAccount = convertedAccount;
  }


  public Lead convertedContact(PartialContactNullable convertedContact) {
    
    
    
    
    this.convertedContact = convertedContact;
    return this;
  }

   /**
   * Get convertedContact
   * @return convertedContact
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(required = true, value = "")

  public PartialContactNullable getConvertedContact() {
    return convertedContact;
  }


  public void setConvertedContact(PartialContactNullable convertedContact) {
    
    
    
    this.convertedContact = convertedContact;
  }


  public Lead lastActivityAt(String lastActivityAt) {
    
    
    
    
    this.lastActivityAt = lastActivityAt;
    return this;
  }

   /**
   * Get lastActivityAt
   * @return lastActivityAt
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(required = true, value = "")

  public String getLastActivityAt() {
    return lastActivityAt;
  }


  public void setLastActivityAt(String lastActivityAt) {
    
    
    
    this.lastActivityAt = lastActivityAt;
  }


  public Lead createdAt(String createdAt) {
    
    
    
    
    this.createdAt = createdAt;
    return this;
  }

   /**
   * Get createdAt
   * @return createdAt
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")

  public String getCreatedAt() {
    return createdAt;
  }


  public void setCreatedAt(String createdAt) {
    
    
    
    this.createdAt = createdAt;
  }


  public Lead updatedAt(String updatedAt) {
    
    
    
    
    this.updatedAt = updatedAt;
    return this;
  }

   /**
   * Get updatedAt
   * @return updatedAt
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")

  public String getUpdatedAt() {
    return updatedAt;
  }


  public void setUpdatedAt(String updatedAt) {
    
    
    
    this.updatedAt = updatedAt;
  }


  public Lead isDeleted(Boolean isDeleted) {
    
    
    
    
    this.isDeleted = isDeleted;
    return this;
  }

   /**
   * Get isDeleted
   * @return isDeleted
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")

  public Boolean getIsDeleted() {
    return isDeleted;
  }


  public void setIsDeleted(Boolean isDeleted) {
    
    
    
    this.isDeleted = isDeleted;
  }


  public Lead tasks(List<Task> tasks) {
    
    
    
    
    this.tasks = tasks;
    return this;
  }

  public Lead addTasksItem(Task tasksItem) {
    if (this.tasks == null) {
      this.tasks = new ArrayList<>();
    }
    this.tasks.add(tasksItem);
    return this;
  }

   /**
   * Get tasks
   * @return tasks
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Task> getTasks() {
    return tasks;
  }


  public void setTasks(List<Task> tasks) {
    
    
    
    this.tasks = tasks;
  }


  public Lead events(List<Event> events) {
    
    
    
    
    this.events = events;
    return this;
  }

  public Lead addEventsItem(Event eventsItem) {
    if (this.events == null) {
      this.events = new ArrayList<>();
    }
    this.events.add(eventsItem);
    return this;
  }

   /**
   * Get events
   * @return events
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Event> getEvents() {
    return events;
  }


  public void setEvents(List<Event> events) {
    
    
    
    this.events = events;
  }


  public Lead remoteData(Object remoteData) {
    
    
    
    
    this.remoteData = remoteData;
    return this;
  }

   /**
   * Get remoteData
   * @return remoteData
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(required = true, value = "")

  public Object getRemoteData() {
    return remoteData;
  }


  public void setRemoteData(Object remoteData) {
    
    
    
    this.remoteData = remoteData;
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
   * @return the Lead instance itself
   */
  public Lead putAdditionalProperty(String key, Object value) {
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
    Lead lead = (Lead) o;
    return Objects.equals(this.title, lead.title) &&
        Objects.equals(this.description, lead.description) &&
        Objects.equals(this.id, lead.id) &&
        Objects.equals(this.owner, lead.owner) &&
        Objects.equals(this.source, lead.source) &&
        Objects.equals(this.status, lead.status) &&
        Objects.equals(this.company, lead.company) &&
        Objects.equals(this.firstName, lead.firstName) &&
        Objects.equals(this.lastName, lead.lastName) &&
        Objects.equals(this.addresses, lead.addresses) &&
        Objects.equals(this.phoneNumbers, lead.phoneNumbers) &&
        Objects.equals(this.emails, lead.emails) &&
        Objects.equals(this.convertedAt, lead.convertedAt) &&
        Objects.equals(this.convertedAccount, lead.convertedAccount) &&
        Objects.equals(this.convertedContact, lead.convertedContact) &&
        Objects.equals(this.lastActivityAt, lead.lastActivityAt) &&
        Objects.equals(this.createdAt, lead.createdAt) &&
        Objects.equals(this.updatedAt, lead.updatedAt) &&
        Objects.equals(this.isDeleted, lead.isDeleted) &&
        Objects.equals(this.tasks, lead.tasks) &&
        Objects.equals(this.events, lead.events) &&
        Objects.equals(this.remoteData, lead.remoteData)&&
        Objects.equals(this.additionalProperties, lead.additionalProperties);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(title, description, id, owner, source, status, company, firstName, lastName, addresses, phoneNumbers, emails, convertedAt, convertedAccount, convertedContact, lastActivityAt, createdAt, updatedAt, isDeleted, tasks, events, remoteData, additionalProperties);
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
    sb.append("class Lead {\n");
    sb.append("    title: ").append(toIndentedString(title)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    owner: ").append(toIndentedString(owner)).append("\n");
    sb.append("    source: ").append(toIndentedString(source)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    company: ").append(toIndentedString(company)).append("\n");
    sb.append("    firstName: ").append(toIndentedString(firstName)).append("\n");
    sb.append("    lastName: ").append(toIndentedString(lastName)).append("\n");
    sb.append("    addresses: ").append(toIndentedString(addresses)).append("\n");
    sb.append("    phoneNumbers: ").append(toIndentedString(phoneNumbers)).append("\n");
    sb.append("    emails: ").append(toIndentedString(emails)).append("\n");
    sb.append("    convertedAt: ").append(toIndentedString(convertedAt)).append("\n");
    sb.append("    convertedAccount: ").append(toIndentedString(convertedAccount)).append("\n");
    sb.append("    convertedContact: ").append(toIndentedString(convertedContact)).append("\n");
    sb.append("    lastActivityAt: ").append(toIndentedString(lastActivityAt)).append("\n");
    sb.append("    createdAt: ").append(toIndentedString(createdAt)).append("\n");
    sb.append("    updatedAt: ").append(toIndentedString(updatedAt)).append("\n");
    sb.append("    isDeleted: ").append(toIndentedString(isDeleted)).append("\n");
    sb.append("    tasks: ").append(toIndentedString(tasks)).append("\n");
    sb.append("    events: ").append(toIndentedString(events)).append("\n");
    sb.append("    remoteData: ").append(toIndentedString(remoteData)).append("\n");
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
    openapiFields.add("title");
    openapiFields.add("description");
    openapiFields.add("id");
    openapiFields.add("owner");
    openapiFields.add("source");
    openapiFields.add("status");
    openapiFields.add("company");
    openapiFields.add("first_name");
    openapiFields.add("last_name");
    openapiFields.add("addresses");
    openapiFields.add("phone_numbers");
    openapiFields.add("emails");
    openapiFields.add("converted_at");
    openapiFields.add("converted_account");
    openapiFields.add("converted_contact");
    openapiFields.add("last_activity_at");
    openapiFields.add("created_at");
    openapiFields.add("updated_at");
    openapiFields.add("is_deleted");
    openapiFields.add("tasks");
    openapiFields.add("events");
    openapiFields.add("remote_data");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("title");
    openapiRequiredFields.add("description");
    openapiRequiredFields.add("id");
    openapiRequiredFields.add("owner");
    openapiRequiredFields.add("source");
    openapiRequiredFields.add("status");
    openapiRequiredFields.add("company");
    openapiRequiredFields.add("first_name");
    openapiRequiredFields.add("last_name");
    openapiRequiredFields.add("addresses");
    openapiRequiredFields.add("phone_numbers");
    openapiRequiredFields.add("emails");
    openapiRequiredFields.add("converted_at");
    openapiRequiredFields.add("converted_account");
    openapiRequiredFields.add("converted_contact");
    openapiRequiredFields.add("last_activity_at");
    openapiRequiredFields.add("created_at");
    openapiRequiredFields.add("updated_at");
    openapiRequiredFields.add("is_deleted");
    openapiRequiredFields.add("remote_data");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to Lead
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!Lead.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in Lead is not found in the empty JSON string", Lead.openapiRequiredFields.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : Lead.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if (!jsonObj.get("title").isJsonNull() && !jsonObj.get("title").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `title` to be a primitive type in the JSON string but got `%s`", jsonObj.get("title").toString()));
      }
      if (!jsonObj.get("description").isJsonNull() && !jsonObj.get("description").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `description` to be a primitive type in the JSON string but got `%s`", jsonObj.get("description").toString()));
      }
      if (!jsonObj.get("id").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `id` to be a primitive type in the JSON string but got `%s`", jsonObj.get("id").toString()));
      }
      // validate the required field `owner`
      PartialOwner.validateJsonObject(jsonObj.getAsJsonObject("owner"));
      if (!jsonObj.get("source").isJsonNull() && !jsonObj.get("source").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `source` to be a primitive type in the JSON string but got `%s`", jsonObj.get("source").toString()));
      }
      if (!jsonObj.get("status").isJsonNull() && !jsonObj.get("status").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `status` to be a primitive type in the JSON string but got `%s`", jsonObj.get("status").toString()));
      }
      if (!jsonObj.get("company").isJsonNull() && !jsonObj.get("company").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `company` to be a primitive type in the JSON string but got `%s`", jsonObj.get("company").toString()));
      }
      if (!jsonObj.get("first_name").isJsonNull() && !jsonObj.get("first_name").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `first_name` to be a primitive type in the JSON string but got `%s`", jsonObj.get("first_name").toString()));
      }
      if (!jsonObj.get("last_name").isJsonNull() && !jsonObj.get("last_name").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `last_name` to be a primitive type in the JSON string but got `%s`", jsonObj.get("last_name").toString()));
      }
      // ensure the json data is an array
      if (!jsonObj.get("addresses").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `addresses` to be an array in the JSON string but got `%s`", jsonObj.get("addresses").toString()));
      }

      JsonArray jsonArrayaddresses = jsonObj.getAsJsonArray("addresses");
      // validate the required field `addresses` (array)
      for (int i = 0; i < jsonArrayaddresses.size(); i++) {
        Address.validateJsonObject(jsonArrayaddresses.get(i).getAsJsonObject());
      };
      // ensure the json data is an array
      if (!jsonObj.get("phone_numbers").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `phone_numbers` to be an array in the JSON string but got `%s`", jsonObj.get("phone_numbers").toString()));
      }

      JsonArray jsonArrayphoneNumbers = jsonObj.getAsJsonArray("phone_numbers");
      // validate the required field `phone_numbers` (array)
      for (int i = 0; i < jsonArrayphoneNumbers.size(); i++) {
        PhoneNumber.validateJsonObject(jsonArrayphoneNumbers.get(i).getAsJsonObject());
      };
      // ensure the json data is an array
      if (!jsonObj.get("emails").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `emails` to be an array in the JSON string but got `%s`", jsonObj.get("emails").toString()));
      }

      JsonArray jsonArrayemails = jsonObj.getAsJsonArray("emails");
      // validate the required field `emails` (array)
      for (int i = 0; i < jsonArrayemails.size(); i++) {
        Email.validateJsonObject(jsonArrayemails.get(i).getAsJsonObject());
      };
      if (!jsonObj.get("converted_at").isJsonNull() && !jsonObj.get("converted_at").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `converted_at` to be a primitive type in the JSON string but got `%s`", jsonObj.get("converted_at").toString()));
      }
      // validate the required field `converted_account`
      if (!jsonObj.get("converted_account").isJsonNull()) {
        PartialAccountNullable.validateJsonObject(jsonObj.getAsJsonObject("converted_account"));
      }
      // validate the required field `converted_contact`
      if (!jsonObj.get("converted_contact").isJsonNull()) {
        PartialContactNullable.validateJsonObject(jsonObj.getAsJsonObject("converted_contact"));
      }
      if (!jsonObj.get("last_activity_at").isJsonNull() && !jsonObj.get("last_activity_at").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `last_activity_at` to be a primitive type in the JSON string but got `%s`", jsonObj.get("last_activity_at").toString()));
      }
      if (!jsonObj.get("created_at").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `created_at` to be a primitive type in the JSON string but got `%s`", jsonObj.get("created_at").toString()));
      }
      if (!jsonObj.get("updated_at").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `updated_at` to be a primitive type in the JSON string but got `%s`", jsonObj.get("updated_at").toString()));
      }
      if (jsonObj.get("tasks") != null && !jsonObj.get("tasks").isJsonNull()) {
        JsonArray jsonArraytasks = jsonObj.getAsJsonArray("tasks");
        if (jsonArraytasks != null) {
          // ensure the json data is an array
          if (!jsonObj.get("tasks").isJsonArray()) {
            throw new IllegalArgumentException(String.format("Expected the field `tasks` to be an array in the JSON string but got `%s`", jsonObj.get("tasks").toString()));
          }

          // validate the optional field `tasks` (array)
          for (int i = 0; i < jsonArraytasks.size(); i++) {
            Task.validateJsonObject(jsonArraytasks.get(i).getAsJsonObject());
          };
        }
      }
      if (jsonObj.get("events") != null && !jsonObj.get("events").isJsonNull()) {
        JsonArray jsonArrayevents = jsonObj.getAsJsonArray("events");
        if (jsonArrayevents != null) {
          // ensure the json data is an array
          if (!jsonObj.get("events").isJsonArray()) {
            throw new IllegalArgumentException(String.format("Expected the field `events` to be an array in the JSON string but got `%s`", jsonObj.get("events").toString()));
          }

          // validate the optional field `events` (array)
          for (int i = 0; i < jsonArrayevents.size(); i++) {
            Event.validateJsonObject(jsonArrayevents.get(i).getAsJsonObject());
          };
        }
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!Lead.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'Lead' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<Lead> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(Lead.class));

       return (TypeAdapter<T>) new TypeAdapter<Lead>() {
           @Override
           public void write(JsonWriter out, Lead value) throws IOException {
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
           public Lead read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             // store additional fields in the deserialized instance
             Lead instance = thisAdapter.fromJsonTree(jsonObj);
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
  * Create an instance of Lead given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of Lead
  * @throws IOException if the JSON string is invalid with respect to Lead
  */
  public static Lead fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, Lead.class);
  }

 /**
  * Convert an instance of Lead to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}


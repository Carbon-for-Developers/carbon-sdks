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
import com.google.gson.annotations.SerializedName;

import java.io.IOException;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;

/**
 * Gets or Sets ZendeskFileTypes
 */
@JsonAdapter(ZendeskFileTypes.Adapter.class)public enum ZendeskFileTypes {
  
  TICKET("TICKET"),
  
  ARTICLE("ARTICLE"),
  
  TICKET_ATTACHMENT("TICKET_ATTACHMENT"),
  
  SECTION("SECTION"),
  
  CATEGORY("CATEGORY");

  private String value;

  ZendeskFileTypes(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static ZendeskFileTypes fromValue(String value) {
    for (ZendeskFileTypes b : ZendeskFileTypes.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<ZendeskFileTypes> {
    @Override
    public void write(final JsonWriter jsonWriter, final ZendeskFileTypes enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public ZendeskFileTypes read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return ZendeskFileTypes.fromValue(value);
    }
  }
}


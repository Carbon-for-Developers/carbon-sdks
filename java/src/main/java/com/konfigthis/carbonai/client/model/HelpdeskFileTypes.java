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
 * Gets or Sets HelpdeskFileTypes
 */
@JsonAdapter(HelpdeskFileTypes.Adapter.class)public enum HelpdeskFileTypes {
  
  TICKET("TICKET"),
  
  ARTICLE("ARTICLE"),
  
  CONVERSATION("CONVERSATION");

  private String value;

  HelpdeskFileTypes(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static HelpdeskFileTypes fromValue(String value) {
    for (HelpdeskFileTypes b : HelpdeskFileTypes.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<HelpdeskFileTypes> {
    @Override
    public void write(final JsonWriter jsonWriter, final HelpdeskFileTypes enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public HelpdeskFileTypes read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return HelpdeskFileTypes.fromValue(value);
    }
  }
}


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
 * Gets or Sets HSNFileTypes1
 */
@JsonAdapter(HSNFileTypes1.Adapter.class)public enum HSNFileTypes1 {
  
  TICKET("TICKET"),
  
  ARTICLE("ARTICLE");

  private String value;

  HSNFileTypes1(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static HSNFileTypes1 fromValue(String value) {
    for (HSNFileTypes1 b : HSNFileTypes1.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<HSNFileTypes1> {
    @Override
    public void write(final JsonWriter jsonWriter, final HSNFileTypes1 enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public HSNFileTypes1 read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return HSNFileTypes1.fromValue(value);
    }
  }
}


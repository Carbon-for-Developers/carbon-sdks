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
 * Gets or Sets PRState
 */
@JsonAdapter(PRState.Adapter.class)public enum PRState {
  
  CLOSED("closed"),
  
  OPEN("open");

  private String value;

  PRState(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static PRState fromValue(String value) {
    for (PRState b : PRState.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<PRState> {
    @Override
    public void write(final JsonWriter jsonWriter, final PRState enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public PRState read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return PRState.fromValue(value);
    }
  }
}


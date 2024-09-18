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
 * Gets or Sets ExternalFileSyncStatuses
 */
@JsonAdapter(ExternalFileSyncStatuses.Adapter.class)public enum ExternalFileSyncStatuses {
  
  DELAYED("DELAYED"),
  
  QUEUED_FOR_SYNC("QUEUED_FOR_SYNC"),
  
  SYNCING("SYNCING"),
  
  READY("READY"),
  
  SYNC_ERROR("SYNC_ERROR"),
  
  EVALUATING_RESYNC("EVALUATING_RESYNC"),
  
  RATE_LIMITED("RATE_LIMITED"),
  
  SYNC_ABORTED("SYNC_ABORTED"),
  
  QUEUED_FOR_OCR("QUEUED_FOR_OCR"),
  
  READY_TO_SYNC("READY_TO_SYNC");

  private String value;

  ExternalFileSyncStatuses(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static ExternalFileSyncStatuses fromValue(String value) {
    for (ExternalFileSyncStatuses b : ExternalFileSyncStatuses.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<ExternalFileSyncStatuses> {
    @Override
    public void write(final JsonWriter jsonWriter, final ExternalFileSyncStatuses enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public ExternalFileSyncStatuses read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return ExternalFileSyncStatuses.fromValue(value);
    }
  }
}


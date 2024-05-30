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
 * Gets or Sets TMEmbeddingGenerators
 */
@JsonAdapter(TMEmbeddingGenerators.Adapter.class)public enum TMEmbeddingGenerators {
  
  OPENAI("OPENAI"),
  
  AZURE_OPENAI("AZURE_OPENAI"),
  
  COHERE_MULTILINGUAL_V3("COHERE_MULTILINGUAL_V3"),
  
  OPENAI_ADA_LARGE_256("OPENAI_ADA_LARGE_256"),
  
  OPENAI_ADA_LARGE_1024("OPENAI_ADA_LARGE_1024"),
  
  OPENAI_ADA_LARGE_3072("OPENAI_ADA_LARGE_3072"),
  
  OPENAI_ADA_SMALL_512("OPENAI_ADA_SMALL_512"),
  
  OPENAI_ADA_SMALL_1536("OPENAI_ADA_SMALL_1536"),
  
  AZURE_ADA_LARGE_256("AZURE_ADA_LARGE_256"),
  
  AZURE_ADA_LARGE_1024("AZURE_ADA_LARGE_1024"),
  
  AZURE_ADA_LARGE_3072("AZURE_ADA_LARGE_3072"),
  
  AZURE_ADA_SMALL_512("AZURE_ADA_SMALL_512"),
  
  AZURE_ADA_SMALL_1536("AZURE_ADA_SMALL_1536"),
  
  SOLAR_1_MINI("SOLAR_1_MINI");

  private String value;

  TMEmbeddingGenerators(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static TMEmbeddingGenerators fromValue(String value) {
    for (TMEmbeddingGenerators b : TMEmbeddingGenerators.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<TMEmbeddingGenerators> {
    @Override
    public void write(final JsonWriter jsonWriter, final TMEmbeddingGenerators enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public TMEmbeddingGenerators read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return TMEmbeddingGenerators.fromValue(value);
    }
  }
}


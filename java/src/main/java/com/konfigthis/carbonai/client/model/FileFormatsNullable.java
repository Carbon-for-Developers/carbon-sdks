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
 * Gets or Sets FileFormatsNullable
 */
@JsonAdapter(FileFormatsNullable.Adapter.class)public enum FileFormatsNullable {
  
  TXT("TXT"),
  
  CSV("CSV"),
  
  TSV("TSV"),
  
  PDF("PDF"),
  
  DOCX("DOCX"),
  
  PPTX("PPTX"),
  
  XLSX("XLSX"),
  
  MD("MD"),
  
  RTF("RTF"),
  
  JSON("JSON"),
  
  HTML("HTML"),
  
  NOTION("NOTION"),
  
  GOOGLE_DOCS("GOOGLE_DOCS"),
  
  GOOGLE_SHEETS("GOOGLE_SHEETS"),
  
  GOOGLE_SLIDES("GOOGLE_SLIDES"),
  
  INTERCOM("INTERCOM"),
  
  CONFLUENCE("CONFLUENCE"),
  
  RSS_FEED("RSS_FEED"),
  
  GMAIL("GMAIL"),
  
  OUTLOOK("OUTLOOK"),
  
  ZENDESK("ZENDESK"),
  
  FRESHDESK("FRESHDESK"),
  
  WEB_SCRAPE("WEB_SCRAPE"),
  
  GITBOOK("GITBOOK"),
  
  SALESFORCE("SALESFORCE"),
  
  GITHUB("GITHUB"),
  
  SLACK("SLACK"),
  
  JPG("JPG"),
  
  PNG("PNG"),
  
  MP3("MP3"),
  
  MP2("MP2"),
  
  AAC("AAC"),
  
  WAV("WAV"),
  
  FLAC("FLAC"),
  
  PCM("PCM"),
  
  M4A("M4A"),
  
  OGG("OGG"),
  
  OPUS("OPUS"),
  
  MPEG("MPEG"),
  
  MPG("MPG"),
  
  MP4("MP4"),
  
  WMV("WMV"),
  
  AVI("AVI"),
  
  MOV("MOV"),
  
  MKV("MKV"),
  
  FLV("FLV"),
  
  WEBM("WEBM");

  private String value;

  FileFormatsNullable(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static FileFormatsNullable fromValue(String value) {
    for (FileFormatsNullable b : FileFormatsNullable.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    return null;
  }

  public static class Adapter extends TypeAdapter<FileFormatsNullable> {
    @Override
    public void write(final JsonWriter jsonWriter, final FileFormatsNullable enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public FileFormatsNullable read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return FileFormatsNullable.fromValue(value);
    }
  }
}


=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'date'
require 'time'

module Carbon
  class FileFormats
    TXT = "TXT".freeze
    CSV = "CSV".freeze
    TSV = "TSV".freeze
    PDF = "PDF".freeze
    DOCX = "DOCX".freeze
    PPTX = "PPTX".freeze
    XLSX = "XLSX".freeze
    MD = "MD".freeze
    RTF = "RTF".freeze
    JSON = "JSON".freeze
    HTML = "HTML".freeze
    NOTION = "NOTION".freeze
    GOOGLE_DOCS = "GOOGLE_DOCS".freeze
    GOOGLE_SHEETS = "GOOGLE_SHEETS".freeze
    GOOGLE_SLIDES = "GOOGLE_SLIDES".freeze
    INTERCOM = "INTERCOM".freeze
    CONFLUENCE = "CONFLUENCE".freeze
    RSS_FEED = "RSS_FEED".freeze
    GMAIL = "GMAIL".freeze
    OUTLOOK = "OUTLOOK".freeze
    ZENDESK = "ZENDESK".freeze
    FRESHDESK = "FRESHDESK".freeze
    WEB_SCRAPE = "WEB_SCRAPE".freeze
    GITBOOK = "GITBOOK".freeze
    SALESFORCE = "SALESFORCE".freeze
    GITHUB = "GITHUB".freeze
    SLACK = "SLACK".freeze
    JPG = "JPG".freeze
    PNG = "PNG".freeze
    MP3 = "MP3".freeze
    MP2 = "MP2".freeze
    AAC = "AAC".freeze
    WAV = "WAV".freeze
    FLAC = "FLAC".freeze
    PCM = "PCM".freeze
    M4_A = "M4A".freeze
    OGG = "OGG".freeze
    OPUS = "OPUS".freeze
    MPEG = "MPEG".freeze
    MPG = "MPG".freeze
    MP4 = "MP4".freeze
    WMV = "WMV".freeze
    AVI = "AVI".freeze
    MOV = "MOV".freeze
    MKV = "MKV".freeze
    FLV = "FLV".freeze
    WEBM = "WEBM".freeze
    EML = "EML".freeze

    def self.all_vars
      @all_vars ||= [TXT, CSV, TSV, PDF, DOCX, PPTX, XLSX, MD, RTF, JSON, HTML, NOTION, GOOGLE_DOCS, GOOGLE_SHEETS, GOOGLE_SLIDES, INTERCOM, CONFLUENCE, RSS_FEED, GMAIL, OUTLOOK, ZENDESK, FRESHDESK, WEB_SCRAPE, GITBOOK, SALESFORCE, GITHUB, SLACK, JPG, PNG, MP3, MP2, AAC, WAV, FLAC, PCM, M4_A, OGG, OPUS, MPEG, MPG, MP4, WMV, AVI, MOV, MKV, FLV, WEBM, EML].freeze
    end

    # Builds the enum from string
    # @param [String] The enum value in the form of the string
    # @return [String] The enum value
    def self.build_from_hash(value)
      new.build_from_hash(value)
    end

    # Builds the enum from string
    # @param [String] The enum value in the form of the string
    # @return [String] The enum value
    def build_from_hash(value)
      return value if FileFormats.all_vars.include?(value)
      raise "Invalid ENUM value #{value} for class #FileFormats"
    end
  end
end

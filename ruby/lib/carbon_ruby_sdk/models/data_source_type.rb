=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'date'
require 'time'

module Carbon
  class DataSourceType
    GOOGLE_DRIVE = "GOOGLE_DRIVE".freeze
    NOTION = "NOTION".freeze
    NOTION_DATABASE = "NOTION_DATABASE".freeze
    INTERCOM = "INTERCOM".freeze
    DROPBOX = "DROPBOX".freeze
    ONEDRIVE = "ONEDRIVE".freeze
    SHAREPOINT = "SHAREPOINT".freeze
    CONFLUENCE = "CONFLUENCE".freeze
    BOX = "BOX".freeze
    ZENDESK = "ZENDESK".freeze
    ZOTERO = "ZOTERO".freeze
    S3 = "S3".freeze
    GMAIL = "GMAIL".freeze
    OUTLOOK = "OUTLOOK".freeze
    TEXT = "TEXT".freeze
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
    RAW_TEXT = "RAW_TEXT".freeze
    WEB_SCRAPE = "WEB_SCRAPE".freeze
    RSS_FEED = "RSS_FEED".freeze
    FRESHDESK = "FRESHDESK".freeze
    GITBOOK = "GITBOOK".freeze
    SALESFORCE = "SALESFORCE".freeze
    GITHUB = "GITHUB".freeze
    SLACK = "SLACK".freeze
    JPG = "JPG".freeze
    PNG = "PNG".freeze
    JPEG = "JPEG".freeze
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

    def self.all_vars
      @all_vars ||= [GOOGLE_DRIVE, NOTION, NOTION_DATABASE, INTERCOM, DROPBOX, ONEDRIVE, SHAREPOINT, CONFLUENCE, BOX, ZENDESK, ZOTERO, S3, GMAIL, OUTLOOK, TEXT, CSV, TSV, PDF, DOCX, PPTX, XLSX, MD, RTF, JSON, HTML, RAW_TEXT, WEB_SCRAPE, RSS_FEED, FRESHDESK, GITBOOK, SALESFORCE, GITHUB, SLACK, JPG, PNG, JPEG, MP3, MP2, AAC, WAV, FLAC, PCM, M4_A, OGG, OPUS, MPEG, MPG, MP4, WMV, AVI, MOV, MKV, FLV, WEBM].freeze
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
      return value if DataSourceType.all_vars.include?(value)
      raise "Invalid ENUM value #{value} for class #DataSourceType"
    end
  end
end

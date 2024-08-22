=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'date'
require 'time'

module Carbon
  class ExternalDataSourceType
    BOX = "BOX".freeze
    CONFLUENCE = "CONFLUENCE".freeze
    DROPBOX = "DROPBOX".freeze
    GMAIL = "GMAIL".freeze
    GOOGLE_DRIVE = "GOOGLE_DRIVE".freeze
    GOOGLE_CLOUD_STORAGE = "GOOGLE_CLOUD_STORAGE".freeze
    INTERCOM = "INTERCOM".freeze
    NOTION = "NOTION".freeze
    ONEDRIVE = "ONEDRIVE".freeze
    OUTLOOK = "OUTLOOK".freeze
    SALESFORCE = "SALESFORCE".freeze
    SHAREPOINT = "SHAREPOINT".freeze
    SLACK = "SLACK".freeze
    ZENDESK = "ZENDESK".freeze
    ZOTERO = "ZOTERO".freeze

    def self.all_vars
      @all_vars ||= [BOX, CONFLUENCE, DROPBOX, GMAIL, GOOGLE_DRIVE, GOOGLE_CLOUD_STORAGE, INTERCOM, NOTION, ONEDRIVE, OUTLOOK, SALESFORCE, SHAREPOINT, SLACK, ZENDESK, ZOTERO].freeze
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
      return value if ExternalDataSourceType.all_vars.include?(value)
      raise "Invalid ENUM value #{value} for class #ExternalDataSourceType"
    end
  end
end

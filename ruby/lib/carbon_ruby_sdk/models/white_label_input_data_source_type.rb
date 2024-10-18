=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'date'
require 'time'

module Carbon
  class WhiteLabelInputDataSourceType
    INTERCOM = "INTERCOM".freeze
    NOTION = "NOTION".freeze
    SLACK = "SLACK".freeze
    ZENDESK = "ZENDESK".freeze
    OUTLOOK = "OUTLOOK".freeze
    GMAIL = "GMAIL".freeze
    SERVICENOW = "SERVICENOW".freeze
    SALESFORCE = "SALESFORCE".freeze
    ZOTERO = "ZOTERO".freeze
    BOX = "BOX".freeze
    CONFLUENCE = "CONFLUENCE".freeze
    DROPBOX = "DROPBOX".freeze
    GOOGLE_CLOUD_STORAGE = "GOOGLE_CLOUD_STORAGE".freeze
    GONG = "GONG".freeze

    def self.all_vars
      @all_vars ||= [INTERCOM, NOTION, SLACK, ZENDESK, OUTLOOK, GMAIL, SERVICENOW, SALESFORCE, ZOTERO, BOX, CONFLUENCE, DROPBOX, GOOGLE_CLOUD_STORAGE, GONG].freeze
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
      return value if WhiteLabelInputDataSourceType.all_vars.include?(value)
      raise "Invalid ENUM value #{value} for class #WhiteLabelInputDataSourceType"
    end
  end
end

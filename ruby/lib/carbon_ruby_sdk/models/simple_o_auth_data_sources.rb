=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'date'
require 'time'

module Carbon
  class SimpleOAuthDataSources
    GOOGLE_DRIVE = "GOOGLE_DRIVE".freeze
    INTERCOM = "INTERCOM".freeze
    DROPBOX = "DROPBOX".freeze
    ONEDRIVE = "ONEDRIVE".freeze
    BOX = "BOX".freeze
    GMAIL = "GMAIL".freeze
    OUTLOOK = "OUTLOOK".freeze
    SLACK = "SLACK".freeze
    GOOGLE_CLOUD_STORAGE = "GOOGLE_CLOUD_STORAGE".freeze

    def self.all_vars
      @all_vars ||= [GOOGLE_DRIVE, INTERCOM, DROPBOX, ONEDRIVE, BOX, GMAIL, OUTLOOK, SLACK, GOOGLE_CLOUD_STORAGE].freeze
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
      return value if SimpleOAuthDataSources.all_vars.include?(value)
      raise "Invalid ENUM value #{value} for class #SimpleOAuthDataSources"
    end
  end
end

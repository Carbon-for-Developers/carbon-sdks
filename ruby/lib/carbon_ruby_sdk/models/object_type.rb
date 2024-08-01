=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'date'
require 'time'

module Carbon
  class ObjectType
    CHUNK_LIST = "CHUNK_LIST".freeze
    DATA_SOURCE = "DATA_SOURCE".freeze
    FILE = "FILE".freeze
    FILE_LIST = "FILE_LIST".freeze
    NONE = "NONE".freeze
    ORGANIZATION_USER = "ORGANIZATION_USER".freeze
    UPLOAD_REQUEST_ID = "UPLOAD_REQUEST_ID".freeze
    WEBPAGE = "WEBPAGE".freeze

    def self.all_vars
      @all_vars ||= [CHUNK_LIST, DATA_SOURCE, FILE, FILE_LIST, NONE, ORGANIZATION_USER, UPLOAD_REQUEST_ID, WEBPAGE].freeze
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
      return value if ObjectType.all_vars.include?(value)
      raise "Invalid ENUM value #{value} for class #ObjectType"
    end
  end
end

=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'date'
require 'time'

module Carbon
  class ServiceNowFileTypes
    TABLE = "TABLE".freeze
    INCIDENT = "INCIDENT".freeze
    ATTACHMENT = "ATTACHMENT".freeze

    def self.all_vars
      @all_vars ||= [TABLE, INCIDENT, ATTACHMENT].freeze
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
      return value if ServiceNowFileTypes.all_vars.include?(value)
      raise "Invalid ENUM value #{value} for class #ServiceNowFileTypes"
    end
  end
end

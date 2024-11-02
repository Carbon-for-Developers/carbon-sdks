=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'date'
require 'time'

module Carbon
  class IntercomFileTypes
    ARTICLE = "ARTICLE".freeze
    COLLECTION = "COLLECTION".freeze
    HELP_CENTER = "HELP_CENTER".freeze
    TICKET = "TICKET".freeze
    TICKET_ATTACHMENT = "TICKET_ATTACHMENT".freeze
    CONTACT = "CONTACT".freeze
    ADMIN = "ADMIN".freeze
    TEAM = "TEAM".freeze
    CONVERSATION = "CONVERSATION".freeze

    def self.all_vars
      @all_vars ||= [ARTICLE, COLLECTION, HELP_CENTER, TICKET, TICKET_ATTACHMENT, CONTACT, ADMIN, TEAM, CONVERSATION].freeze
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
      return value if IntercomFileTypes.all_vars.include?(value)
      raise "Invalid ENUM value #{value} for class #IntercomFileTypes"
    end
  end
end

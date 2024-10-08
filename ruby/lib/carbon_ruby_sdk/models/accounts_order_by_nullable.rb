=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'date'
require 'time'

module Carbon
  class AccountsOrderByNullable
    CREATED_AT = "created_at".freeze
    UPDATED_AT = "updated_at".freeze
    NUMBER_OF_EMPLOYEES = "number_of_employees".freeze
    NAME = "name".freeze
    LAST_ACTIVITY_AT = "last_activity_at".freeze

    def self.all_vars
      @all_vars ||= [CREATED_AT, UPDATED_AT, NUMBER_OF_EMPLOYEES, NAME, LAST_ACTIVITY_AT].freeze
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
      return value if AccountsOrderByNullable.all_vars.include?(value)
      raise "Invalid ENUM value #{value} for class #AccountsOrderByNullable"
    end
  end
end

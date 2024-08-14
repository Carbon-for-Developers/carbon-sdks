=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'date'
require 'time'

module Carbon
  class EmbeddingStorageStatus
    HOT_STORAGE = "HOT_STORAGE".freeze
    HOT_TO_COLD = "HOT_TO_COLD".freeze
    COLD_STORAGE = "COLD_STORAGE".freeze
    COLD_TO_HOT = "COLD_TO_HOT".freeze

    def self.all_vars
      @all_vars ||= [HOT_STORAGE, HOT_TO_COLD, COLD_STORAGE, COLD_TO_HOT].freeze
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
      return value if EmbeddingStorageStatus.all_vars.include?(value)
      raise "Invalid ENUM value #{value} for class #EmbeddingStorageStatus"
    end
  end
end

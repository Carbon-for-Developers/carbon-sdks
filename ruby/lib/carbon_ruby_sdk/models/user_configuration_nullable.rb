=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'date'
require 'time'

module Carbon
  # Used to set organization level defaults for user settings. These settings will apply to all users under         the organization that don't have them defined.
  class UserConfigurationNullable
    attr_accessor :auto_sync_enabled_sources

    # Custom file upload limit for the user over *all* user's files across all uploads.          If set, then the user will not be allowed to upload more files than this limit. If not set, or if set to -1,         then the user will have no limit.
    attr_accessor :max_files

    # Custom file upload limit for the user across a single upload.         If set, then the user will not be allowed to upload more files than this limit in a single upload. If not set,         or if set to -1, then the user will have no limit.
    attr_accessor :max_files_per_upload

    # Custom character upload limit for the user over *all* user's files across all uploads.          If set, then the user will not be allowed to upload more characters than this limit. If not set, or if set to -1,         then the user will have no limit.
    attr_accessor :max_characters

    # A single file upload from the user can not exceed this character limit.         If set, then the file will not be synced if it exceeds this limit. If not set, or if set to -1, then the          user will have no limit.
    attr_accessor :max_characters_per_file

    # Custom character upload limit for the user across a single upload.         If set, then the user won't be able to sync more than this many characters in one upload.          If not set, or if set to -1, then the user will have no limit.
    attr_accessor :max_characters_per_upload

    # Attribute mapping from ruby-style variable name to JSON key.
    def self.attribute_map
      {
        :'auto_sync_enabled_sources' => :'auto_sync_enabled_sources',
        :'max_files' => :'max_files',
        :'max_files_per_upload' => :'max_files_per_upload',
        :'max_characters' => :'max_characters',
        :'max_characters_per_file' => :'max_characters_per_file',
        :'max_characters_per_upload' => :'max_characters_per_upload'
      }
    end

    # Returns all the JSON keys this model knows about
    def self.acceptable_attributes
      attribute_map.values
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'auto_sync_enabled_sources' => :'AutoSyncEnabledSourcesProperty',
        :'max_files' => :'Integer',
        :'max_files_per_upload' => :'Integer',
        :'max_characters' => :'Integer',
        :'max_characters_per_file' => :'Integer',
        :'max_characters_per_upload' => :'Integer'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
        :'auto_sync_enabled_sources',
        :'max_files',
        :'max_files_per_upload',
        :'max_characters',
        :'max_characters_per_file',
        :'max_characters_per_upload'
      ])
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `Carbon::UserConfigurationNullable` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `Carbon::UserConfigurationNullable`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'auto_sync_enabled_sources')
        self.auto_sync_enabled_sources = attributes[:'auto_sync_enabled_sources']
      end

      if attributes.key?(:'max_files')
        self.max_files = attributes[:'max_files']
      end

      if attributes.key?(:'max_files_per_upload')
        self.max_files_per_upload = attributes[:'max_files_per_upload']
      end

      if attributes.key?(:'max_characters')
        self.max_characters = attributes[:'max_characters']
      end

      if attributes.key?(:'max_characters_per_file')
        self.max_characters_per_file = attributes[:'max_characters_per_file']
      end

      if attributes.key?(:'max_characters_per_upload')
        self.max_characters_per_upload = attributes[:'max_characters_per_upload']
      end
    end

    # Show invalid properties with the reasons. Usually used together with valid?
    # @return Array for valid properties with the reasons
    def list_invalid_properties
      invalid_properties = Array.new
      if !@max_files.nil? && @max_files < -1
        invalid_properties.push('invalid value for "max_files", must be greater than or equal to -1.')
      end

      if !@max_files_per_upload.nil? && @max_files_per_upload < -1
        invalid_properties.push('invalid value for "max_files_per_upload", must be greater than or equal to -1.')
      end

      if !@max_characters.nil? && @max_characters < -1
        invalid_properties.push('invalid value for "max_characters", must be greater than or equal to -1.')
      end

      if !@max_characters_per_file.nil? && @max_characters_per_file < -1
        invalid_properties.push('invalid value for "max_characters_per_file", must be greater than or equal to -1.')
      end

      if !@max_characters_per_upload.nil? && @max_characters_per_upload < -1
        invalid_properties.push('invalid value for "max_characters_per_upload", must be greater than or equal to -1.')
      end

      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      return false if !@max_files.nil? && @max_files < -1
      return false if !@max_files_per_upload.nil? && @max_files_per_upload < -1
      return false if !@max_characters.nil? && @max_characters < -1
      return false if !@max_characters_per_file.nil? && @max_characters_per_file < -1
      return false if !@max_characters_per_upload.nil? && @max_characters_per_upload < -1
      true
    end

    # Custom attribute writer method with validation
    # @param [Object] max_files Value to be assigned
    def max_files=(max_files)
      if !max_files.nil? && max_files < -1
        fail ArgumentError, 'invalid value for "max_files", must be greater than or equal to -1.'
      end

      @max_files = max_files
    end

    # Custom attribute writer method with validation
    # @param [Object] max_files_per_upload Value to be assigned
    def max_files_per_upload=(max_files_per_upload)
      if !max_files_per_upload.nil? && max_files_per_upload < -1
        fail ArgumentError, 'invalid value for "max_files_per_upload", must be greater than or equal to -1.'
      end

      @max_files_per_upload = max_files_per_upload
    end

    # Custom attribute writer method with validation
    # @param [Object] max_characters Value to be assigned
    def max_characters=(max_characters)
      if !max_characters.nil? && max_characters < -1
        fail ArgumentError, 'invalid value for "max_characters", must be greater than or equal to -1.'
      end

      @max_characters = max_characters
    end

    # Custom attribute writer method with validation
    # @param [Object] max_characters_per_file Value to be assigned
    def max_characters_per_file=(max_characters_per_file)
      if !max_characters_per_file.nil? && max_characters_per_file < -1
        fail ArgumentError, 'invalid value for "max_characters_per_file", must be greater than or equal to -1.'
      end

      @max_characters_per_file = max_characters_per_file
    end

    # Custom attribute writer method with validation
    # @param [Object] max_characters_per_upload Value to be assigned
    def max_characters_per_upload=(max_characters_per_upload)
      if !max_characters_per_upload.nil? && max_characters_per_upload < -1
        fail ArgumentError, 'invalid value for "max_characters_per_upload", must be greater than or equal to -1.'
      end

      @max_characters_per_upload = max_characters_per_upload
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          auto_sync_enabled_sources == o.auto_sync_enabled_sources &&
          max_files == o.max_files &&
          max_files_per_upload == o.max_files_per_upload &&
          max_characters == o.max_characters &&
          max_characters_per_file == o.max_characters_per_file &&
          max_characters_per_upload == o.max_characters_per_upload
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [auto_sync_enabled_sources, max_files, max_files_per_upload, max_characters, max_characters_per_file, max_characters_per_upload].hash
    end

    # Builds the object from hash
    # @param [Hash] attributes Model attributes in the form of hash
    # @return [Object] Returns the model itself
    def self.build_from_hash(attributes)
      new.build_from_hash(attributes)
    end

    # Builds the object from hash
    # @param [Hash] attributes Model attributes in the form of hash
    # @return [Object] Returns the model itself
    def build_from_hash(attributes)
      return nil unless attributes.is_a?(Hash)
      attributes = attributes.transform_keys(&:to_sym)
      self.class.openapi_types.each_pair do |key, type|
        if attributes[self.class.attribute_map[key]].nil? && self.class.openapi_nullable.include?(key)
          self.send("#{key}=", nil)
        elsif type =~ /\AArray<(.*)>/i
          # check to ensure the input is an array given that the attribute
          # is documented as an array but the input is not
          if attributes[self.class.attribute_map[key]].is_a?(Array)
            self.send("#{key}=", attributes[self.class.attribute_map[key]].map { |v| _deserialize($1, v) })
          end
        elsif !attributes[self.class.attribute_map[key]].nil?
          self.send("#{key}=", _deserialize(type, attributes[self.class.attribute_map[key]]))
        end
      end

      self
    end

    # Deserializes the data based on type
    # @param string type Data type
    # @param string value Value to be deserialized
    # @return [Object] Deserialized data
    def _deserialize(type, value)
      case type.to_sym
      when :Time
        Time.parse(value)
      when :Date
        Date.parse(value)
      when :String
        value.to_s
      when :Integer
        value.to_i
      when :Float
        value.to_f
      when :Boolean
        if value.to_s =~ /\A(true|t|yes|y|1)\z/i
          true
        else
          false
        end
      when :Object
        # generic object (usually a Hash), return directly
        value
      when /\AArray<(?<inner_type>.+)>\z/
        inner_type = Regexp.last_match[:inner_type]
        value.map { |v| _deserialize(inner_type, v) }
      when /\AHash<(?<k_type>.+?), (?<v_type>.+)>\z/
        k_type = Regexp.last_match[:k_type]
        v_type = Regexp.last_match[:v_type]
        {}.tap do |hash|
          value.each do |k, v|
            hash[_deserialize(k_type, k)] = _deserialize(v_type, v)
          end
        end
      else # model
        # models (e.g. Pet) or oneOf
        klass = Carbon.const_get(type)
        klass.respond_to?(:openapi_one_of) ? klass.build(value) : klass.build_from_hash(value)
      end
    end

    # Returns the string representation of the object
    # @return [String] String presentation of the object
    def to_s
      to_hash.to_s
    end

    # to_body is an alias to to_hash (backward compatibility)
    # @return [Hash] Returns the object in the form of hash
    def to_body
      to_hash
    end

    # Returns the object in the form of hash
    # @return [Hash] Returns the object in the form of hash
    def to_hash
      hash = {}
      self.class.attribute_map.each_pair do |attr, param|
        value = self.send(attr)
        if value.nil?
          is_nullable = self.class.openapi_nullable.include?(attr)
          next if !is_nullable || (is_nullable && !instance_variable_defined?(:"@#{attr}"))
        end

        hash[param] = _to_hash(value)
      end
      hash
    end

    # Outputs non-array value in the form of hash
    # For object, use to_hash. Otherwise, just return the value
    # @param [Object] value Any valid value
    # @return [Hash] Returns the value in the form of hash
    def _to_hash(value)
      if value.is_a?(Array)
        value.compact.map { |v| _to_hash(v) }
      elsif value.is_a?(Hash)
        {}.tap do |hash|
          value.each { |k, v| hash[k] = _to_hash(v) }
        end
      elsif value.respond_to? :to_hash
        value.to_hash
      else
        value
      end
    end

  end

end

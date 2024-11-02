=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'date'
require 'time'

module Carbon
  class ListItemsFiltersNullable
    attr_accessor :external_ids

    attr_accessor :ids

    attr_accessor :name

    attr_accessor :root_files_only

    attr_accessor :file_formats

    attr_accessor :item_types

    # Attribute mapping from ruby-style variable name to JSON key.
    def self.attribute_map
      {
        :'external_ids' => :'external_ids',
        :'ids' => :'ids',
        :'name' => :'name',
        :'root_files_only' => :'root_files_only',
        :'file_formats' => :'file_formats',
        :'item_types' => :'item_types'
      }
    end

    # Returns all the JSON keys this model knows about
    def self.acceptable_attributes
      attribute_map.values
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'external_ids' => :'Array<String>',
        :'ids' => :'Array<Integer>',
        :'name' => :'String',
        :'root_files_only' => :'Boolean',
        :'file_formats' => :'Array<StorageFileFormats>',
        :'item_types' => :'Array<ItemTypesPropertyInner>'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
        :'external_ids',
        :'ids',
        :'name',
        :'root_files_only',
        :'file_formats',
        :'item_types'
      ])
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `Carbon::ListItemsFiltersNullable` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `Carbon::ListItemsFiltersNullable`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'external_ids')
        if (value = attributes[:'external_ids']).is_a?(Array)
          self.external_ids = value
        end
      end

      if attributes.key?(:'ids')
        if (value = attributes[:'ids']).is_a?(Array)
          self.ids = value
        end
      end

      if attributes.key?(:'name')
        self.name = attributes[:'name']
      end

      if attributes.key?(:'root_files_only')
        self.root_files_only = attributes[:'root_files_only']
      end

      if attributes.key?(:'file_formats')
        if (value = attributes[:'file_formats']).is_a?(Array)
          self.file_formats = value
        end
      end

      if attributes.key?(:'item_types')
        if (value = attributes[:'item_types']).is_a?(Array)
          self.item_types = value
        end
      end
    end

    # Show invalid properties with the reasons. Usually used together with valid?
    # @return Array for valid properties with the reasons
    def list_invalid_properties
      invalid_properties = Array.new
      if !@file_formats.nil? && @file_formats.length > 50
        invalid_properties.push('invalid value for "file_formats", number of items must be less than or equal to 50.')
      end

      if !@item_types.nil? && @item_types.length > 50
        invalid_properties.push('invalid value for "item_types", number of items must be less than or equal to 50.')
      end

      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      return false if !@file_formats.nil? && @file_formats.length > 50
      return false if !@item_types.nil? && @item_types.length > 50
      true
    end

    # Custom attribute writer method with validation
    # @param [Object] file_formats Value to be assigned
    def file_formats=(file_formats)
      if !file_formats.nil? && file_formats.length > 50
        fail ArgumentError, 'invalid value for "file_formats", number of items must be less than or equal to 50.'
      end

      @file_formats = file_formats
    end

    # Custom attribute writer method with validation
    # @param [Object] item_types Value to be assigned
    def item_types=(item_types)
      if !item_types.nil? && item_types.length > 50
        fail ArgumentError, 'invalid value for "item_types", number of items must be less than or equal to 50.'
      end

      @item_types = item_types
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          external_ids == o.external_ids &&
          ids == o.ids &&
          name == o.name &&
          root_files_only == o.root_files_only &&
          file_formats == o.file_formats &&
          item_types == o.item_types
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [external_ids, ids, name, root_files_only, file_formats, item_types].hash
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

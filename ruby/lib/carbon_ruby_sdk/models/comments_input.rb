=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'date'
require 'time'

module Carbon
  class CommentsInput
    attr_accessor :data_source_id

    attr_accessor :include_remote_data

    # Full name of the repository, denoted as {owner}/{repo}
    attr_accessor :repository

    attr_accessor :page

    attr_accessor :page_size

    attr_accessor :next_cursor

    attr_accessor :pull_number

    attr_accessor :order_by

    attr_accessor :order_dir

    # Attribute mapping from ruby-style variable name to JSON key.
    def self.attribute_map
      {
        :'data_source_id' => :'data_source_id',
        :'include_remote_data' => :'include_remote_data',
        :'repository' => :'repository',
        :'page' => :'page',
        :'page_size' => :'page_size',
        :'next_cursor' => :'next_cursor',
        :'pull_number' => :'pull_number',
        :'order_by' => :'order_by',
        :'order_dir' => :'order_dir'
      }
    end

    # Returns all the JSON keys this model knows about
    def self.acceptable_attributes
      attribute_map.values
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'data_source_id' => :'Integer',
        :'include_remote_data' => :'Boolean',
        :'repository' => :'String',
        :'page' => :'Integer',
        :'page_size' => :'Integer',
        :'next_cursor' => :'String',
        :'pull_number' => :'Integer',
        :'order_by' => :'CommentsOrderBy',
        :'order_dir' => :'OrderDirV2Nullable'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
        :'next_cursor',
        :'order_dir'
      ])
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `Carbon::CommentsInput` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `Carbon::CommentsInput`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'data_source_id')
        self.data_source_id = attributes[:'data_source_id']
      end

      if attributes.key?(:'include_remote_data')
        self.include_remote_data = attributes[:'include_remote_data']
      else
        self.include_remote_data = false
      end

      if attributes.key?(:'repository')
        self.repository = attributes[:'repository']
      end

      if attributes.key?(:'page')
        self.page = attributes[:'page']
      else
        self.page = 1
      end

      if attributes.key?(:'page_size')
        self.page_size = attributes[:'page_size']
      else
        self.page_size = 30
      end

      if attributes.key?(:'next_cursor')
        self.next_cursor = attributes[:'next_cursor']
      end

      if attributes.key?(:'pull_number')
        self.pull_number = attributes[:'pull_number']
      end

      if attributes.key?(:'order_by')
        self.order_by = attributes[:'order_by']
      else
        self.order_by = 'created'
      end

      if attributes.key?(:'order_dir')
        self.order_dir = attributes[:'order_dir']
      else
        self.order_dir = 'asc'
      end
    end

    # Show invalid properties with the reasons. Usually used together with valid?
    # @return Array for valid properties with the reasons
    def list_invalid_properties
      invalid_properties = Array.new
      if @data_source_id.nil?
        invalid_properties.push('invalid value for "data_source_id", data_source_id cannot be nil.')
      end

      if @repository.nil?
        invalid_properties.push('invalid value for "repository", repository cannot be nil.')
      end

      if !@page_size.nil? && @page_size > 100
        invalid_properties.push('invalid value for "page_size", must be smaller than or equal to 100.')
      end

      if @pull_number.nil?
        invalid_properties.push('invalid value for "pull_number", pull_number cannot be nil.')
      end

      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      return false if @data_source_id.nil?
      return false if @repository.nil?
      return false if !@page_size.nil? && @page_size > 100
      return false if @pull_number.nil?
      true
    end

    # Custom attribute writer method with validation
    # @param [Object] page_size Value to be assigned
    def page_size=(page_size)
      if !page_size.nil? && page_size > 100
        fail ArgumentError, 'invalid value for "page_size", must be smaller than or equal to 100.'
      end

      @page_size = page_size
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          data_source_id == o.data_source_id &&
          include_remote_data == o.include_remote_data &&
          repository == o.repository &&
          page == o.page &&
          page_size == o.page_size &&
          next_cursor == o.next_cursor &&
          pull_number == o.pull_number &&
          order_by == o.order_by &&
          order_dir == o.order_dir
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [data_source_id, include_remote_data, repository, page, page_size, next_cursor, pull_number, order_by, order_dir].hash
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

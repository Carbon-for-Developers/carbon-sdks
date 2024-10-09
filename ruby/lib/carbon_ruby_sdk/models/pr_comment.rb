=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'date'
require 'time'

module Carbon
  class PRComment
    attr_accessor :id

    attr_accessor :pull_request_review_id

    attr_accessor :url

    attr_accessor :diff_hunk

    attr_accessor :path

    attr_accessor :user

    attr_accessor :body

    attr_accessor :created_at

    attr_accessor :updated_at

    attr_accessor :start_line

    attr_accessor :line

    attr_accessor :remote_data

    # Attribute mapping from ruby-style variable name to JSON key.
    def self.attribute_map
      {
        :'id' => :'id',
        :'pull_request_review_id' => :'pull_request_review_id',
        :'url' => :'url',
        :'diff_hunk' => :'diff_hunk',
        :'path' => :'path',
        :'user' => :'user',
        :'body' => :'body',
        :'created_at' => :'created_at',
        :'updated_at' => :'updated_at',
        :'start_line' => :'start_line',
        :'line' => :'line',
        :'remote_data' => :'remote_data'
      }
    end

    # Returns all the JSON keys this model knows about
    def self.acceptable_attributes
      attribute_map.values
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'id' => :'Integer',
        :'pull_request_review_id' => :'Integer',
        :'url' => :'String',
        :'diff_hunk' => :'String',
        :'path' => :'String',
        :'user' => :'User',
        :'body' => :'String',
        :'created_at' => :'String',
        :'updated_at' => :'String',
        :'start_line' => :'Integer',
        :'line' => :'Integer',
        :'remote_data' => :'Object'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
        :'pull_request_review_id',
        :'start_line',
        :'line',
        :'remote_data'
      ])
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `Carbon::PRComment` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `Carbon::PRComment`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'id')
        self.id = attributes[:'id']
      end

      if attributes.key?(:'pull_request_review_id')
        self.pull_request_review_id = attributes[:'pull_request_review_id']
      end

      if attributes.key?(:'url')
        self.url = attributes[:'url']
      end

      if attributes.key?(:'diff_hunk')
        self.diff_hunk = attributes[:'diff_hunk']
      end

      if attributes.key?(:'path')
        self.path = attributes[:'path']
      end

      if attributes.key?(:'user')
        self.user = attributes[:'user']
      end

      if attributes.key?(:'body')
        self.body = attributes[:'body']
      end

      if attributes.key?(:'created_at')
        self.created_at = attributes[:'created_at']
      end

      if attributes.key?(:'updated_at')
        self.updated_at = attributes[:'updated_at']
      end

      if attributes.key?(:'start_line')
        self.start_line = attributes[:'start_line']
      end

      if attributes.key?(:'line')
        self.line = attributes[:'line']
      end

      if attributes.key?(:'remote_data')
        self.remote_data = attributes[:'remote_data']
      end
    end

    # Show invalid properties with the reasons. Usually used together with valid?
    # @return Array for valid properties with the reasons
    def list_invalid_properties
      invalid_properties = Array.new
      if @id.nil?
        invalid_properties.push('invalid value for "id", id cannot be nil.')
      end

      if @url.nil?
        invalid_properties.push('invalid value for "url", url cannot be nil.')
      end

      if @diff_hunk.nil?
        invalid_properties.push('invalid value for "diff_hunk", diff_hunk cannot be nil.')
      end

      if @path.nil?
        invalid_properties.push('invalid value for "path", path cannot be nil.')
      end

      if @user.nil?
        invalid_properties.push('invalid value for "user", user cannot be nil.')
      end

      if @body.nil?
        invalid_properties.push('invalid value for "body", body cannot be nil.')
      end

      if @created_at.nil?
        invalid_properties.push('invalid value for "created_at", created_at cannot be nil.')
      end

      if @updated_at.nil?
        invalid_properties.push('invalid value for "updated_at", updated_at cannot be nil.')
      end

      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      return false if @id.nil?
      return false if @url.nil?
      return false if @diff_hunk.nil?
      return false if @path.nil?
      return false if @user.nil?
      return false if @body.nil?
      return false if @created_at.nil?
      return false if @updated_at.nil?
      true
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          id == o.id &&
          pull_request_review_id == o.pull_request_review_id &&
          url == o.url &&
          diff_hunk == o.diff_hunk &&
          path == o.path &&
          user == o.user &&
          body == o.body &&
          created_at == o.created_at &&
          updated_at == o.updated_at &&
          start_line == o.start_line &&
          line == o.line &&
          remote_data == o.remote_data
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [id, pull_request_review_id, url, diff_hunk, path, user, body, created_at, updated_at, start_line, line, remote_data].hash
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

=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'date'
require 'time'

module Carbon
  class Issue
    attr_accessor :title

    attr_accessor :id

    attr_accessor :url

    attr_accessor :html_url

    attr_accessor :number

    attr_accessor :user

    attr_accessor :labels

    attr_accessor :state

    attr_accessor :locked

    attr_accessor :num_comments

    attr_accessor :created_at

    attr_accessor :updated_at

    attr_accessor :closed_at

    attr_accessor :draft

    attr_accessor :pull_request

    attr_accessor :body

    attr_accessor :closed_by

    # Attribute mapping from ruby-style variable name to JSON key.
    def self.attribute_map
      {
        :'title' => :'title',
        :'id' => :'id',
        :'url' => :'url',
        :'html_url' => :'html_url',
        :'number' => :'number',
        :'user' => :'user',
        :'labels' => :'labels',
        :'state' => :'state',
        :'locked' => :'locked',
        :'num_comments' => :'num_comments',
        :'created_at' => :'created_at',
        :'updated_at' => :'updated_at',
        :'closed_at' => :'closed_at',
        :'draft' => :'draft',
        :'pull_request' => :'pull_request',
        :'body' => :'body',
        :'closed_by' => :'closed_by'
      }
    end

    # Returns all the JSON keys this model knows about
    def self.acceptable_attributes
      attribute_map.values
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'title' => :'String',
        :'id' => :'Integer',
        :'url' => :'String',
        :'html_url' => :'String',
        :'number' => :'Integer',
        :'user' => :'User',
        :'labels' => :'Array<Label>',
        :'state' => :'PRState',
        :'locked' => :'Boolean',
        :'num_comments' => :'Integer',
        :'created_at' => :'String',
        :'updated_at' => :'String',
        :'closed_at' => :'String',
        :'draft' => :'Boolean',
        :'pull_request' => :'IssuePRNullable',
        :'body' => :'String',
        :'closed_by' => :'UserNullable'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
        :'closed_at',
        :'pull_request',
        :'closed_by'
      ])
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `Carbon::Issue` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `Carbon::Issue`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'title')
        self.title = attributes[:'title']
      end

      if attributes.key?(:'id')
        self.id = attributes[:'id']
      end

      if attributes.key?(:'url')
        self.url = attributes[:'url']
      end

      if attributes.key?(:'html_url')
        self.html_url = attributes[:'html_url']
      end

      if attributes.key?(:'number')
        self.number = attributes[:'number']
      end

      if attributes.key?(:'user')
        self.user = attributes[:'user']
      end

      if attributes.key?(:'labels')
        if (value = attributes[:'labels']).is_a?(Array)
          self.labels = value
        end
      end

      if attributes.key?(:'state')
        self.state = attributes[:'state']
      end

      if attributes.key?(:'locked')
        self.locked = attributes[:'locked']
      end

      if attributes.key?(:'num_comments')
        self.num_comments = attributes[:'num_comments']
      end

      if attributes.key?(:'created_at')
        self.created_at = attributes[:'created_at']
      end

      if attributes.key?(:'updated_at')
        self.updated_at = attributes[:'updated_at']
      end

      if attributes.key?(:'closed_at')
        self.closed_at = attributes[:'closed_at']
      end

      if attributes.key?(:'draft')
        self.draft = attributes[:'draft']
      end

      if attributes.key?(:'pull_request')
        self.pull_request = attributes[:'pull_request']
      end

      if attributes.key?(:'body')
        self.body = attributes[:'body']
      end

      if attributes.key?(:'closed_by')
        self.closed_by = attributes[:'closed_by']
      end
    end

    # Show invalid properties with the reasons. Usually used together with valid?
    # @return Array for valid properties with the reasons
    def list_invalid_properties
      invalid_properties = Array.new
      if @title.nil?
        invalid_properties.push('invalid value for "title", title cannot be nil.')
      end

      if @id.nil?
        invalid_properties.push('invalid value for "id", id cannot be nil.')
      end

      if @url.nil?
        invalid_properties.push('invalid value for "url", url cannot be nil.')
      end

      if @html_url.nil?
        invalid_properties.push('invalid value for "html_url", html_url cannot be nil.')
      end

      if @number.nil?
        invalid_properties.push('invalid value for "number", number cannot be nil.')
      end

      if @user.nil?
        invalid_properties.push('invalid value for "user", user cannot be nil.')
      end

      if @labels.nil?
        invalid_properties.push('invalid value for "labels", labels cannot be nil.')
      end

      if @state.nil?
        invalid_properties.push('invalid value for "state", state cannot be nil.')
      end

      if @locked.nil?
        invalid_properties.push('invalid value for "locked", locked cannot be nil.')
      end

      if @num_comments.nil?
        invalid_properties.push('invalid value for "num_comments", num_comments cannot be nil.')
      end

      if @created_at.nil?
        invalid_properties.push('invalid value for "created_at", created_at cannot be nil.')
      end

      if @updated_at.nil?
        invalid_properties.push('invalid value for "updated_at", updated_at cannot be nil.')
      end

      if @draft.nil?
        invalid_properties.push('invalid value for "draft", draft cannot be nil.')
      end

      if @body.nil?
        invalid_properties.push('invalid value for "body", body cannot be nil.')
      end

      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      return false if @title.nil?
      return false if @id.nil?
      return false if @url.nil?
      return false if @html_url.nil?
      return false if @number.nil?
      return false if @user.nil?
      return false if @labels.nil?
      return false if @state.nil?
      return false if @locked.nil?
      return false if @num_comments.nil?
      return false if @created_at.nil?
      return false if @updated_at.nil?
      return false if @draft.nil?
      return false if @body.nil?
      true
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          title == o.title &&
          id == o.id &&
          url == o.url &&
          html_url == o.html_url &&
          number == o.number &&
          user == o.user &&
          labels == o.labels &&
          state == o.state &&
          locked == o.locked &&
          num_comments == o.num_comments &&
          created_at == o.created_at &&
          updated_at == o.updated_at &&
          closed_at == o.closed_at &&
          draft == o.draft &&
          pull_request == o.pull_request &&
          body == o.body &&
          closed_by == o.closed_by
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [title, id, url, html_url, number, user, labels, state, locked, num_comments, created_at, updated_at, closed_at, draft, pull_request, body, closed_by].hash
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

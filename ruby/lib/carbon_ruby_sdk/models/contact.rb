=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'date'
require 'time'

module Carbon
  class Contact
    attr_accessor :title

    attr_accessor :description

    attr_accessor :id

    attr_accessor :owner

    attr_accessor :first_name

    attr_accessor :last_name

    attr_accessor :name

    attr_accessor :department

    attr_accessor :addresses

    attr_accessor :phone_numbers

    attr_accessor :emails

    attr_accessor :account

    attr_accessor :last_activity_at

    attr_accessor :created_at

    attr_accessor :updated_at

    attr_accessor :is_deleted

    attr_accessor :tasks

    attr_accessor :events

    attr_accessor :remote_data

    # Attribute mapping from ruby-style variable name to JSON key.
    def self.attribute_map
      {
        :'title' => :'title',
        :'description' => :'description',
        :'id' => :'id',
        :'owner' => :'owner',
        :'first_name' => :'first_name',
        :'last_name' => :'last_name',
        :'name' => :'name',
        :'department' => :'department',
        :'addresses' => :'addresses',
        :'phone_numbers' => :'phone_numbers',
        :'emails' => :'emails',
        :'account' => :'account',
        :'last_activity_at' => :'last_activity_at',
        :'created_at' => :'created_at',
        :'updated_at' => :'updated_at',
        :'is_deleted' => :'is_deleted',
        :'tasks' => :'tasks',
        :'events' => :'events',
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
        :'title' => :'String',
        :'description' => :'String',
        :'id' => :'String',
        :'owner' => :'PartialOwnerNullable',
        :'first_name' => :'String',
        :'last_name' => :'String',
        :'name' => :'String',
        :'department' => :'String',
        :'addresses' => :'Array<Address>',
        :'phone_numbers' => :'Array<PhoneNumber>',
        :'emails' => :'Array<Email>',
        :'account' => :'PartialAccountNullable',
        :'last_activity_at' => :'String',
        :'created_at' => :'String',
        :'updated_at' => :'String',
        :'is_deleted' => :'Boolean',
        :'tasks' => :'Array<Task>',
        :'events' => :'Array<Event>',
        :'remote_data' => :'Object'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
        :'title',
        :'description',
        :'owner',
        :'first_name',
        :'last_name',
        :'name',
        :'department',
        :'account',
        :'last_activity_at',
        :'tasks',
        :'events',
        :'remote_data'
      ])
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `Carbon::Contact` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `Carbon::Contact`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'title')
        self.title = attributes[:'title']
      end

      if attributes.key?(:'description')
        self.description = attributes[:'description']
      end

      if attributes.key?(:'id')
        self.id = attributes[:'id']
      end

      if attributes.key?(:'owner')
        self.owner = attributes[:'owner']
      end

      if attributes.key?(:'first_name')
        self.first_name = attributes[:'first_name']
      end

      if attributes.key?(:'last_name')
        self.last_name = attributes[:'last_name']
      end

      if attributes.key?(:'name')
        self.name = attributes[:'name']
      end

      if attributes.key?(:'department')
        self.department = attributes[:'department']
      end

      if attributes.key?(:'addresses')
        if (value = attributes[:'addresses']).is_a?(Array)
          self.addresses = value
        end
      end

      if attributes.key?(:'phone_numbers')
        if (value = attributes[:'phone_numbers']).is_a?(Array)
          self.phone_numbers = value
        end
      end

      if attributes.key?(:'emails')
        if (value = attributes[:'emails']).is_a?(Array)
          self.emails = value
        end
      end

      if attributes.key?(:'account')
        self.account = attributes[:'account']
      end

      if attributes.key?(:'last_activity_at')
        self.last_activity_at = attributes[:'last_activity_at']
      end

      if attributes.key?(:'created_at')
        self.created_at = attributes[:'created_at']
      end

      if attributes.key?(:'updated_at')
        self.updated_at = attributes[:'updated_at']
      end

      if attributes.key?(:'is_deleted')
        self.is_deleted = attributes[:'is_deleted']
      end

      if attributes.key?(:'tasks')
        if (value = attributes[:'tasks']).is_a?(Array)
          self.tasks = value
        end
      end

      if attributes.key?(:'events')
        if (value = attributes[:'events']).is_a?(Array)
          self.events = value
        end
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

      if @addresses.nil?
        invalid_properties.push('invalid value for "addresses", addresses cannot be nil.')
      end

      if @phone_numbers.nil?
        invalid_properties.push('invalid value for "phone_numbers", phone_numbers cannot be nil.')
      end

      if @emails.nil?
        invalid_properties.push('invalid value for "emails", emails cannot be nil.')
      end

      if @created_at.nil?
        invalid_properties.push('invalid value for "created_at", created_at cannot be nil.')
      end

      if @updated_at.nil?
        invalid_properties.push('invalid value for "updated_at", updated_at cannot be nil.')
      end

      if @is_deleted.nil?
        invalid_properties.push('invalid value for "is_deleted", is_deleted cannot be nil.')
      end

      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      return false if @id.nil?
      return false if @addresses.nil?
      return false if @phone_numbers.nil?
      return false if @emails.nil?
      return false if @created_at.nil?
      return false if @updated_at.nil?
      return false if @is_deleted.nil?
      true
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          title == o.title &&
          description == o.description &&
          id == o.id &&
          owner == o.owner &&
          first_name == o.first_name &&
          last_name == o.last_name &&
          name == o.name &&
          department == o.department &&
          addresses == o.addresses &&
          phone_numbers == o.phone_numbers &&
          emails == o.emails &&
          account == o.account &&
          last_activity_at == o.last_activity_at &&
          created_at == o.created_at &&
          updated_at == o.updated_at &&
          is_deleted == o.is_deleted &&
          tasks == o.tasks &&
          events == o.events &&
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
      [title, description, id, owner, first_name, last_name, name, department, addresses, phone_numbers, emails, account, last_activity_at, created_at, updated_at, is_deleted, tasks, events, remote_data].hash
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

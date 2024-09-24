=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'date'
require 'time'

module Carbon
  class AuthenticationProperty
    attr_accessor :source

    attr_accessor :access_token

    attr_accessor :refresh_token

    attr_accessor :workspace_id

    attr_accessor :tenant_name

    attr_accessor :site_name

    attr_accessor :subdomain

    attr_accessor :access_token_secret

    attr_accessor :username

    attr_accessor :zotero_id

    attr_accessor :organization_name

    attr_accessor :domain

    attr_accessor :api_key

    attr_accessor :access_key

    attr_accessor :access_key_secret

    # You can specify a Digital Ocean endpoint URL to connect a Digital Ocean Space through this endpoint.         The URL should be of format <region>.digitaloceanspaces.com. It's not required for S3 buckets.
    attr_accessor :endpoint_url

    attr_accessor :account_name

    attr_accessor :account_key

    attr_accessor :instance_subdomain

    attr_accessor :client_id

    attr_accessor :client_secret

    attr_accessor :redirect_uri

    attr_accessor :gong_account_email

    # Attribute mapping from ruby-style variable name to JSON key.
    def self.attribute_map
      {
        :'source' => :'source',
        :'access_token' => :'access_token',
        :'refresh_token' => :'refresh_token',
        :'workspace_id' => :'workspace_id',
        :'tenant_name' => :'tenant_name',
        :'site_name' => :'site_name',
        :'subdomain' => :'subdomain',
        :'access_token_secret' => :'access_token_secret',
        :'username' => :'username',
        :'zotero_id' => :'zotero_id',
        :'organization_name' => :'organization_name',
        :'domain' => :'domain',
        :'api_key' => :'api_key',
        :'access_key' => :'access_key',
        :'access_key_secret' => :'access_key_secret',
        :'endpoint_url' => :'endpoint_url',
        :'account_name' => :'account_name',
        :'account_key' => :'account_key',
        :'instance_subdomain' => :'instance_subdomain',
        :'client_id' => :'client_id',
        :'client_secret' => :'client_secret',
        :'redirect_uri' => :'redirect_uri',
        :'gong_account_email' => :'gong_account_email'
      }
    end

    # Returns all the JSON keys this model knows about
    def self.acceptable_attributes
      attribute_map.values
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'source' => :'Object',
        :'access_token' => :'String',
        :'refresh_token' => :'String',
        :'workspace_id' => :'String',
        :'tenant_name' => :'String',
        :'site_name' => :'String',
        :'subdomain' => :'String',
        :'access_token_secret' => :'String',
        :'username' => :'String',
        :'zotero_id' => :'String',
        :'organization_name' => :'String',
        :'domain' => :'String',
        :'api_key' => :'String',
        :'access_key' => :'String',
        :'access_key_secret' => :'String',
        :'endpoint_url' => :'String',
        :'account_name' => :'String',
        :'account_key' => :'String',
        :'instance_subdomain' => :'String',
        :'client_id' => :'String',
        :'client_secret' => :'String',
        :'redirect_uri' => :'String',
        :'gong_account_email' => :'String'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
        :'source',
        :'refresh_token',
        :'tenant_name',
        :'endpoint_url',
      ])
    end

    # List of class defined in anyOf (OpenAPI v3)
    def self.openapi_any_of
      [
      :'AzureBlobStorageAuthentication',
      :'ConfluenceAuthentication',
      :'FreskdeskAuthentication',
      :'GitbookAuthetication',
      :'GithubAuthentication',
      :'GongAuthentication',
      :'GuruAuthentication',
      :'NotionAuthentication',
      :'OAuthAuthentication',
      :'OneDriveAuthentication',
      :'S3Authentication',
      :'SalesforceAuthentication',
      :'ServiceNowAuthentication',
      :'SharepointAuthentication',
      :'ZendeskAuthentication',
      :'ZoteroAuthentication'
      ]
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `Carbon::AuthenticationProperty` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `Carbon::AuthenticationProperty`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'source')
        self.source = attributes[:'source']
      end

      if attributes.key?(:'access_token')
        self.access_token = attributes[:'access_token']
      end

      if attributes.key?(:'refresh_token')
        self.refresh_token = attributes[:'refresh_token']
      end

      if attributes.key?(:'workspace_id')
        self.workspace_id = attributes[:'workspace_id']
      end

      if attributes.key?(:'tenant_name')
        self.tenant_name = attributes[:'tenant_name']
      end

      if attributes.key?(:'site_name')
        self.site_name = attributes[:'site_name']
      end

      if attributes.key?(:'subdomain')
        self.subdomain = attributes[:'subdomain']
      end

      if attributes.key?(:'access_token_secret')
        self.access_token_secret = attributes[:'access_token_secret']
      end

      if attributes.key?(:'username')
        self.username = attributes[:'username']
      end

      if attributes.key?(:'zotero_id')
        self.zotero_id = attributes[:'zotero_id']
      end

      if attributes.key?(:'organization_name')
        self.organization_name = attributes[:'organization_name']
      end

      if attributes.key?(:'domain')
        self.domain = attributes[:'domain']
      end

      if attributes.key?(:'api_key')
        self.api_key = attributes[:'api_key']
      end

      if attributes.key?(:'access_key')
        self.access_key = attributes[:'access_key']
      end

      if attributes.key?(:'access_key_secret')
        self.access_key_secret = attributes[:'access_key_secret']
      end

      if attributes.key?(:'endpoint_url')
        self.endpoint_url = attributes[:'endpoint_url']
      end

      if attributes.key?(:'account_name')
        self.account_name = attributes[:'account_name']
      end

      if attributes.key?(:'account_key')
        self.account_key = attributes[:'account_key']
      end

      if attributes.key?(:'instance_subdomain')
        self.instance_subdomain = attributes[:'instance_subdomain']
      end

      if attributes.key?(:'client_id')
        self.client_id = attributes[:'client_id']
      end

      if attributes.key?(:'client_secret')
        self.client_secret = attributes[:'client_secret']
      end

      if attributes.key?(:'redirect_uri')
        self.redirect_uri = attributes[:'redirect_uri']
      end

      if attributes.key?(:'gong_account_email')
        self.gong_account_email = attributes[:'gong_account_email']
      end
    end

    # Show invalid properties with the reasons. Usually used together with valid?
    # @return Array for valid properties with the reasons
    def list_invalid_properties
      invalid_properties = Array.new
      if @access_token.nil?
        invalid_properties.push('invalid value for "access_token", access_token cannot be nil.')
      end

      if @workspace_id.nil?
        invalid_properties.push('invalid value for "workspace_id", workspace_id cannot be nil.')
      end

      if @site_name.nil?
        invalid_properties.push('invalid value for "site_name", site_name cannot be nil.')
      end

      if @subdomain.nil?
        invalid_properties.push('invalid value for "subdomain", subdomain cannot be nil.')
      end

      if @access_token_secret.nil?
        invalid_properties.push('invalid value for "access_token_secret", access_token_secret cannot be nil.')
      end

      if @username.nil?
        invalid_properties.push('invalid value for "username", username cannot be nil.')
      end

      if @zotero_id.nil?
        invalid_properties.push('invalid value for "zotero_id", zotero_id cannot be nil.')
      end

      if @organization_name.nil?
        invalid_properties.push('invalid value for "organization_name", organization_name cannot be nil.')
      end

      if @domain.nil?
        invalid_properties.push('invalid value for "domain", domain cannot be nil.')
      end

      if @api_key.nil?
        invalid_properties.push('invalid value for "api_key", api_key cannot be nil.')
      end

      if @access_key.nil?
        invalid_properties.push('invalid value for "access_key", access_key cannot be nil.')
      end

      if @access_key_secret.nil?
        invalid_properties.push('invalid value for "access_key_secret", access_key_secret cannot be nil.')
      end

      if @account_name.nil?
        invalid_properties.push('invalid value for "account_name", account_name cannot be nil.')
      end

      if @account_key.nil?
        invalid_properties.push('invalid value for "account_key", account_key cannot be nil.')
      end

      if @instance_subdomain.nil?
        invalid_properties.push('invalid value for "instance_subdomain", instance_subdomain cannot be nil.')
      end

      if @client_id.nil?
        invalid_properties.push('invalid value for "client_id", client_id cannot be nil.')
      end

      if @client_secret.nil?
        invalid_properties.push('invalid value for "client_secret", client_secret cannot be nil.')
      end

      if @redirect_uri.nil?
        invalid_properties.push('invalid value for "redirect_uri", redirect_uri cannot be nil.')
      end

      if @gong_account_email.nil?
        invalid_properties.push('invalid value for "gong_account_email", gong_account_email cannot be nil.')
      end

      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      return false if @access_token.nil?
      return false if @workspace_id.nil?
      return false if @site_name.nil?
      return false if @subdomain.nil?
      return false if @access_token_secret.nil?
      return false if @username.nil?
      return false if @zotero_id.nil?
      return false if @organization_name.nil?
      return false if @domain.nil?
      return false if @api_key.nil?
      return false if @access_key.nil?
      return false if @access_key_secret.nil?
      return false if @account_name.nil?
      return false if @account_key.nil?
      return false if @instance_subdomain.nil?
      return false if @client_id.nil?
      return false if @client_secret.nil?
      return false if @redirect_uri.nil?
      return false if @gong_account_email.nil?
      _any_of_found = false
      self.class.openapi_any_of.each do |_class|
        _any_of = Carbon.const_get(_class).build_from_hash(self.to_hash)
        if _any_of.valid?
          _any_of_found = true
        end
      end

      if !_any_of_found
        return false
      end

      true
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          source == o.source &&
          access_token == o.access_token &&
          refresh_token == o.refresh_token &&
          workspace_id == o.workspace_id &&
          tenant_name == o.tenant_name &&
          site_name == o.site_name &&
          subdomain == o.subdomain &&
          access_token_secret == o.access_token_secret &&
          username == o.username &&
          zotero_id == o.zotero_id &&
          organization_name == o.organization_name &&
          domain == o.domain &&
          api_key == o.api_key &&
          access_key == o.access_key &&
          access_key_secret == o.access_key_secret &&
          endpoint_url == o.endpoint_url &&
          account_name == o.account_name &&
          account_key == o.account_key &&
          instance_subdomain == o.instance_subdomain &&
          client_id == o.client_id &&
          client_secret == o.client_secret &&
          redirect_uri == o.redirect_uri &&
          gong_account_email == o.gong_account_email
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [source, access_token, refresh_token, workspace_id, tenant_name, site_name, subdomain, access_token_secret, username, zotero_id, organization_name, domain, api_key, access_key, access_key_secret, endpoint_url, account_name, account_key, instance_subdomain, client_id, client_secret, redirect_uri, gong_account_email].hash
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

=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'date'
require 'time'

module Carbon
  class PullRequestExtended
    attr_accessor :title

    attr_accessor :id

    attr_accessor :url

    attr_accessor :number

    attr_accessor :state

    attr_accessor :user

    attr_accessor :created_at

    attr_accessor :updated_at

    attr_accessor :closed_at

    attr_accessor :merged_at

    attr_accessor :requested_reviewers

    attr_accessor :requested_teams

    attr_accessor :labels

    attr_accessor :draft

    attr_accessor :head

    attr_accessor :base

    attr_accessor :remote_data

    attr_accessor :merged

    attr_accessor :num_comments

    attr_accessor :num_review_comments

    attr_accessor :num_commits

    attr_accessor :num_additions

    attr_accessor :num_deletions

    attr_accessor :num_changed_files

    # Attribute mapping from ruby-style variable name to JSON key.
    def self.attribute_map
      {
        :'title' => :'title',
        :'id' => :'id',
        :'url' => :'url',
        :'number' => :'number',
        :'state' => :'state',
        :'user' => :'user',
        :'created_at' => :'created_at',
        :'updated_at' => :'updated_at',
        :'closed_at' => :'closed_at',
        :'merged_at' => :'merged_at',
        :'requested_reviewers' => :'requested_reviewers',
        :'requested_teams' => :'requested_teams',
        :'labels' => :'labels',
        :'draft' => :'draft',
        :'head' => :'head',
        :'base' => :'base',
        :'remote_data' => :'remote_data',
        :'merged' => :'merged',
        :'num_comments' => :'num_comments',
        :'num_review_comments' => :'num_review_comments',
        :'num_commits' => :'num_commits',
        :'num_additions' => :'num_additions',
        :'num_deletions' => :'num_deletions',
        :'num_changed_files' => :'num_changed_files'
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
        :'number' => :'Integer',
        :'state' => :'PRState',
        :'user' => :'User',
        :'created_at' => :'String',
        :'updated_at' => :'String',
        :'closed_at' => :'String',
        :'merged_at' => :'String',
        :'requested_reviewers' => :'Array<User>',
        :'requested_teams' => :'Array<Team>',
        :'labels' => :'Array<Label>',
        :'draft' => :'Boolean',
        :'head' => :'HeadOrBase',
        :'base' => :'HeadOrBase',
        :'remote_data' => :'Object',
        :'merged' => :'Boolean',
        :'num_comments' => :'Integer',
        :'num_review_comments' => :'Integer',
        :'num_commits' => :'Integer',
        :'num_additions' => :'Integer',
        :'num_deletions' => :'Integer',
        :'num_changed_files' => :'Integer'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
        :'closed_at',
        :'merged_at',
        :'remote_data',
      ])
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `Carbon::PullRequestExtended` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `Carbon::PullRequestExtended`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
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

      if attributes.key?(:'number')
        self.number = attributes[:'number']
      end

      if attributes.key?(:'state')
        self.state = attributes[:'state']
      end

      if attributes.key?(:'user')
        self.user = attributes[:'user']
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

      if attributes.key?(:'merged_at')
        self.merged_at = attributes[:'merged_at']
      end

      if attributes.key?(:'requested_reviewers')
        if (value = attributes[:'requested_reviewers']).is_a?(Array)
          self.requested_reviewers = value
        end
      end

      if attributes.key?(:'requested_teams')
        if (value = attributes[:'requested_teams']).is_a?(Array)
          self.requested_teams = value
        end
      end

      if attributes.key?(:'labels')
        if (value = attributes[:'labels']).is_a?(Array)
          self.labels = value
        end
      end

      if attributes.key?(:'draft')
        self.draft = attributes[:'draft']
      end

      if attributes.key?(:'head')
        self.head = attributes[:'head']
      end

      if attributes.key?(:'base')
        self.base = attributes[:'base']
      end

      if attributes.key?(:'remote_data')
        self.remote_data = attributes[:'remote_data']
      end

      if attributes.key?(:'merged')
        self.merged = attributes[:'merged']
      end

      if attributes.key?(:'num_comments')
        self.num_comments = attributes[:'num_comments']
      end

      if attributes.key?(:'num_review_comments')
        self.num_review_comments = attributes[:'num_review_comments']
      end

      if attributes.key?(:'num_commits')
        self.num_commits = attributes[:'num_commits']
      end

      if attributes.key?(:'num_additions')
        self.num_additions = attributes[:'num_additions']
      end

      if attributes.key?(:'num_deletions')
        self.num_deletions = attributes[:'num_deletions']
      end

      if attributes.key?(:'num_changed_files')
        self.num_changed_files = attributes[:'num_changed_files']
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

      if @number.nil?
        invalid_properties.push('invalid value for "number", number cannot be nil.')
      end

      if @state.nil?
        invalid_properties.push('invalid value for "state", state cannot be nil.')
      end

      if @user.nil?
        invalid_properties.push('invalid value for "user", user cannot be nil.')
      end

      if @created_at.nil?
        invalid_properties.push('invalid value for "created_at", created_at cannot be nil.')
      end

      if @updated_at.nil?
        invalid_properties.push('invalid value for "updated_at", updated_at cannot be nil.')
      end

      if @requested_reviewers.nil?
        invalid_properties.push('invalid value for "requested_reviewers", requested_reviewers cannot be nil.')
      end

      if @requested_teams.nil?
        invalid_properties.push('invalid value for "requested_teams", requested_teams cannot be nil.')
      end

      if @labels.nil?
        invalid_properties.push('invalid value for "labels", labels cannot be nil.')
      end

      if @draft.nil?
        invalid_properties.push('invalid value for "draft", draft cannot be nil.')
      end

      if @head.nil?
        invalid_properties.push('invalid value for "head", head cannot be nil.')
      end

      if @base.nil?
        invalid_properties.push('invalid value for "base", base cannot be nil.')
      end

      if @merged.nil?
        invalid_properties.push('invalid value for "merged", merged cannot be nil.')
      end

      if @num_comments.nil?
        invalid_properties.push('invalid value for "num_comments", num_comments cannot be nil.')
      end

      if @num_review_comments.nil?
        invalid_properties.push('invalid value for "num_review_comments", num_review_comments cannot be nil.')
      end

      if @num_commits.nil?
        invalid_properties.push('invalid value for "num_commits", num_commits cannot be nil.')
      end

      if @num_additions.nil?
        invalid_properties.push('invalid value for "num_additions", num_additions cannot be nil.')
      end

      if @num_deletions.nil?
        invalid_properties.push('invalid value for "num_deletions", num_deletions cannot be nil.')
      end

      if @num_changed_files.nil?
        invalid_properties.push('invalid value for "num_changed_files", num_changed_files cannot be nil.')
      end

      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      return false if @title.nil?
      return false if @id.nil?
      return false if @url.nil?
      return false if @number.nil?
      return false if @state.nil?
      return false if @user.nil?
      return false if @created_at.nil?
      return false if @updated_at.nil?
      return false if @requested_reviewers.nil?
      return false if @requested_teams.nil?
      return false if @labels.nil?
      return false if @draft.nil?
      return false if @head.nil?
      return false if @base.nil?
      return false if @merged.nil?
      return false if @num_comments.nil?
      return false if @num_review_comments.nil?
      return false if @num_commits.nil?
      return false if @num_additions.nil?
      return false if @num_deletions.nil?
      return false if @num_changed_files.nil?
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
          number == o.number &&
          state == o.state &&
          user == o.user &&
          created_at == o.created_at &&
          updated_at == o.updated_at &&
          closed_at == o.closed_at &&
          merged_at == o.merged_at &&
          requested_reviewers == o.requested_reviewers &&
          requested_teams == o.requested_teams &&
          labels == o.labels &&
          draft == o.draft &&
          head == o.head &&
          base == o.base &&
          remote_data == o.remote_data &&
          merged == o.merged &&
          num_comments == o.num_comments &&
          num_review_comments == o.num_review_comments &&
          num_commits == o.num_commits &&
          num_additions == o.num_additions &&
          num_deletions == o.num_deletions &&
          num_changed_files == o.num_changed_files
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [title, id, url, number, state, user, created_at, updated_at, closed_at, merged_at, requested_reviewers, requested_teams, labels, draft, head, base, remote_data, merged, num_comments, num_review_comments, num_commits, num_additions, num_deletions, num_changed_files].hash
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

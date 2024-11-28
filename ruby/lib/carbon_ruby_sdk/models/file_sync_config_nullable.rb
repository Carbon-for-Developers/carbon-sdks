=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'date'
require 'time'

module Carbon
  # Used to configure file syncing for certain connectors when sync_files_on_connection is set to true
  class FileSyncConfigNullable
    # File types to automatically sync when the data source connects. Only a subset of file types can be          controlled. If not supported, then they will always be synced
    attr_accessor :auto_synced_source_types

    # Automatically sync attachments from files where supported. Currently applies to Helpdesk Tickets
    attr_accessor :sync_attachments

    # Detect audio language before transcription for audio files
    attr_accessor :detect_audio_language

    attr_accessor :transcription_service

    # Detect multiple speakers and label segments of speech by speaker for audio files.
    attr_accessor :include_speaker_labels

    # Whether to split tabular rows into chunks. Currently only valid for CSV, TSV, and XLSX files.
    attr_accessor :split_rows

    # If this flag is enabled, the file will be chunked and stored with Carbon,           but no embeddings will be generated. This overrides the skip_embedding_generation flag.
    attr_accessor :generate_chunks_only

    # If this flag is enabled, the file will be stored with Carbon, but no chunks or embeddings will be generated.            This overrides the skip_embedding_generation and generate_chunks_only flags.
    attr_accessor :store_file_only

    # Setting this flag will create a new file record with Carbon but skip any and all processing.          This means that we do not download the remote file content or generate any chunks or embeddings. We will store         some metadata like name, external id, and external URL depending on the source you are syncing from. Note that this          flag overrides both skip_embedding_generation and generate_chunks_only flags. The file will be moved to          READY_TO_SYNC status.
    attr_accessor :skip_file_processing

    # The format in which the text content of the file should be parsed. Notion files are the only files         that support this feature. The default parsed text format is plain text.
    attr_accessor :parsed_text_format

    # Attribute mapping from ruby-style variable name to JSON key.
    def self.attribute_map
      {
        :'auto_synced_source_types' => :'auto_synced_source_types',
        :'sync_attachments' => :'sync_attachments',
        :'detect_audio_language' => :'detect_audio_language',
        :'transcription_service' => :'transcription_service',
        :'include_speaker_labels' => :'include_speaker_labels',
        :'split_rows' => :'split_rows',
        :'generate_chunks_only' => :'generate_chunks_only',
        :'store_file_only' => :'store_file_only',
        :'skip_file_processing' => :'skip_file_processing',
        :'parsed_text_format' => :'parsed_text_format'
      }
    end

    # Returns all the JSON keys this model knows about
    def self.acceptable_attributes
      attribute_map.values
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'auto_synced_source_types' => :'Array<AutoSyncedSourceTypesPropertyInner>',
        :'sync_attachments' => :'Boolean',
        :'detect_audio_language' => :'Boolean',
        :'transcription_service' => :'TranscriptionServiceNullable',
        :'include_speaker_labels' => :'Boolean',
        :'split_rows' => :'Boolean',
        :'generate_chunks_only' => :'Boolean',
        :'store_file_only' => :'Boolean',
        :'skip_file_processing' => :'Boolean',
        :'parsed_text_format' => :'ParsedTextFormatsNullable'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
        :'transcription_service',
        :'parsed_text_format'
      ])
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `Carbon::FileSyncConfigNullable` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `Carbon::FileSyncConfigNullable`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'auto_synced_source_types')
        if (value = attributes[:'auto_synced_source_types']).is_a?(Array)
          self.auto_synced_source_types = value
        end
      end

      if attributes.key?(:'sync_attachments')
        self.sync_attachments = attributes[:'sync_attachments']
      else
        self.sync_attachments = false
      end

      if attributes.key?(:'detect_audio_language')
        self.detect_audio_language = attributes[:'detect_audio_language']
      else
        self.detect_audio_language = false
      end

      if attributes.key?(:'transcription_service')
        self.transcription_service = attributes[:'transcription_service']
      end

      if attributes.key?(:'include_speaker_labels')
        self.include_speaker_labels = attributes[:'include_speaker_labels']
      else
        self.include_speaker_labels = false
      end

      if attributes.key?(:'split_rows')
        self.split_rows = attributes[:'split_rows']
      else
        self.split_rows = false
      end

      if attributes.key?(:'generate_chunks_only')
        self.generate_chunks_only = attributes[:'generate_chunks_only']
      else
        self.generate_chunks_only = false
      end

      if attributes.key?(:'store_file_only')
        self.store_file_only = attributes[:'store_file_only']
      else
        self.store_file_only = false
      end

      if attributes.key?(:'skip_file_processing')
        self.skip_file_processing = attributes[:'skip_file_processing']
      else
        self.skip_file_processing = false
      end

      if attributes.key?(:'parsed_text_format')
        self.parsed_text_format = attributes[:'parsed_text_format']
      end
    end

    # Show invalid properties with the reasons. Usually used together with valid?
    # @return Array for valid properties with the reasons
    def list_invalid_properties
      invalid_properties = Array.new
      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      true
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          auto_synced_source_types == o.auto_synced_source_types &&
          sync_attachments == o.sync_attachments &&
          detect_audio_language == o.detect_audio_language &&
          transcription_service == o.transcription_service &&
          include_speaker_labels == o.include_speaker_labels &&
          split_rows == o.split_rows &&
          generate_chunks_only == o.generate_chunks_only &&
          store_file_only == o.store_file_only &&
          skip_file_processing == o.skip_file_processing &&
          parsed_text_format == o.parsed_text_format
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [auto_synced_source_types, sync_attachments, detect_audio_language, transcription_service, include_speaker_labels, split_rows, generate_chunks_only, store_file_only, skip_file_processing, parsed_text_format].hash
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

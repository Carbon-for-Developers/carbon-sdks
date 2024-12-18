=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'date'
require 'time'

module Carbon
  class SitemapScrapeRequest
    attr_accessor :tags

    attr_accessor :url

    attr_accessor :max_pages_to_scrape

    attr_accessor :chunk_size

    attr_accessor :chunk_overlap

    attr_accessor :skip_embedding_generation

    attr_accessor :enable_auto_sync

    attr_accessor :generate_sparse_vectors

    attr_accessor :prepend_filename_to_chunks

    attr_accessor :html_tags_to_skip

    attr_accessor :css_classes_to_skip

    attr_accessor :css_selectors_to_skip

    attr_accessor :embedding_model

    # URL subpaths or directories that you want to include. For example if you want to only include         URLs that start with /questions in stackoverflow.com, you will add /questions/ in this input
    attr_accessor :url_paths_to_include

    # URL subpaths or directories that you want to exclude. For example if you want to exclude         URLs that start with /questions in stackoverflow.com, you will add /questions/ in this input
    attr_accessor :url_paths_to_exclude

    # You can submit a subset of URLs from the sitemap that should be scraped. To get the list of URLs,           you can check out /process_sitemap endpoint. If left empty, all URLs from the sitemap will be scraped.
    attr_accessor :urls_to_scrape

    # Whether the scraper should download css and media from the page (images, fonts, etc). Scrapes          might take longer to finish with this flag enabled, but the success rate is improved.
    attr_accessor :download_css_and_media

    # If this flag is enabled, the file will be chunked and stored with Carbon,           but no embeddings will be generated. This overrides the skip_embedding_generation flag.
    attr_accessor :generate_chunks_only

    # If this flag is enabled, the file will be stored with Carbon, but no processing will be done.
    attr_accessor :store_file_only

    # If the default proxies are blocked and not returning results, this flag can be enabled to use              alternate proxies (residential and office). Scrapes might take longer to finish with this flag enabled.         
    attr_accessor :use_premium_proxies

    # Attribute mapping from ruby-style variable name to JSON key.
    def self.attribute_map
      {
        :'tags' => :'tags',
        :'url' => :'url',
        :'max_pages_to_scrape' => :'max_pages_to_scrape',
        :'chunk_size' => :'chunk_size',
        :'chunk_overlap' => :'chunk_overlap',
        :'skip_embedding_generation' => :'skip_embedding_generation',
        :'enable_auto_sync' => :'enable_auto_sync',
        :'generate_sparse_vectors' => :'generate_sparse_vectors',
        :'prepend_filename_to_chunks' => :'prepend_filename_to_chunks',
        :'html_tags_to_skip' => :'html_tags_to_skip',
        :'css_classes_to_skip' => :'css_classes_to_skip',
        :'css_selectors_to_skip' => :'css_selectors_to_skip',
        :'embedding_model' => :'embedding_model',
        :'url_paths_to_include' => :'url_paths_to_include',
        :'url_paths_to_exclude' => :'url_paths_to_exclude',
        :'urls_to_scrape' => :'urls_to_scrape',
        :'download_css_and_media' => :'download_css_and_media',
        :'generate_chunks_only' => :'generate_chunks_only',
        :'store_file_only' => :'store_file_only',
        :'use_premium_proxies' => :'use_premium_proxies'
      }
    end

    # Returns all the JSON keys this model knows about
    def self.acceptable_attributes
      attribute_map.values
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'tags' => :'Hash<String, Tags1>',
        :'url' => :'String',
        :'max_pages_to_scrape' => :'Integer',
        :'chunk_size' => :'Integer',
        :'chunk_overlap' => :'Integer',
        :'skip_embedding_generation' => :'Boolean',
        :'enable_auto_sync' => :'Boolean',
        :'generate_sparse_vectors' => :'Boolean',
        :'prepend_filename_to_chunks' => :'Boolean',
        :'html_tags_to_skip' => :'Array<String>',
        :'css_classes_to_skip' => :'Array<String>',
        :'css_selectors_to_skip' => :'Array<String>',
        :'embedding_model' => :'EmbeddingGenerators',
        :'url_paths_to_include' => :'Array<String>',
        :'url_paths_to_exclude' => :'Array<String>',
        :'urls_to_scrape' => :'Array<String>',
        :'download_css_and_media' => :'Boolean',
        :'generate_chunks_only' => :'Boolean',
        :'store_file_only' => :'Boolean',
        :'use_premium_proxies' => :'Boolean'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
        :'tags',
        :'max_pages_to_scrape',
        :'chunk_size',
        :'chunk_overlap',
        :'skip_embedding_generation',
        :'enable_auto_sync',
        :'generate_sparse_vectors',
        :'prepend_filename_to_chunks',
        :'html_tags_to_skip',
        :'css_classes_to_skip',
        :'css_selectors_to_skip',
        :'url_paths_to_include',
        :'url_paths_to_exclude',
        :'urls_to_scrape',
        :'download_css_and_media',
      ])
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `Carbon::SitemapScrapeRequest` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `Carbon::SitemapScrapeRequest`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'tags')
        if (value = attributes[:'tags']).is_a?(Hash)
          self.tags = value
        end
      end

      if attributes.key?(:'url')
        self.url = attributes[:'url']
      end

      if attributes.key?(:'max_pages_to_scrape')
        self.max_pages_to_scrape = attributes[:'max_pages_to_scrape']
      end

      if attributes.key?(:'chunk_size')
        self.chunk_size = attributes[:'chunk_size']
      else
        self.chunk_size = 1500
      end

      if attributes.key?(:'chunk_overlap')
        self.chunk_overlap = attributes[:'chunk_overlap']
      else
        self.chunk_overlap = 20
      end

      if attributes.key?(:'skip_embedding_generation')
        self.skip_embedding_generation = attributes[:'skip_embedding_generation']
      else
        self.skip_embedding_generation = false
      end

      if attributes.key?(:'enable_auto_sync')
        self.enable_auto_sync = attributes[:'enable_auto_sync']
      else
        self.enable_auto_sync = false
      end

      if attributes.key?(:'generate_sparse_vectors')
        self.generate_sparse_vectors = attributes[:'generate_sparse_vectors']
      else
        self.generate_sparse_vectors = false
      end

      if attributes.key?(:'prepend_filename_to_chunks')
        self.prepend_filename_to_chunks = attributes[:'prepend_filename_to_chunks']
      else
        self.prepend_filename_to_chunks = false
      end

      if attributes.key?(:'html_tags_to_skip')
        if (value = attributes[:'html_tags_to_skip']).is_a?(Array)
          self.html_tags_to_skip = value
        end
      end

      if attributes.key?(:'css_classes_to_skip')
        if (value = attributes[:'css_classes_to_skip']).is_a?(Array)
          self.css_classes_to_skip = value
        end
      end

      if attributes.key?(:'css_selectors_to_skip')
        if (value = attributes[:'css_selectors_to_skip']).is_a?(Array)
          self.css_selectors_to_skip = value
        end
      end

      if attributes.key?(:'embedding_model')
        self.embedding_model = attributes[:'embedding_model']
      end

      if attributes.key?(:'url_paths_to_include')
        if (value = attributes[:'url_paths_to_include']).is_a?(Array)
          self.url_paths_to_include = value
        end
      end

      if attributes.key?(:'url_paths_to_exclude')
        if (value = attributes[:'url_paths_to_exclude']).is_a?(Array)
          self.url_paths_to_exclude = value
        end
      end

      if attributes.key?(:'urls_to_scrape')
        if (value = attributes[:'urls_to_scrape']).is_a?(Array)
          self.urls_to_scrape = value
        end
      end

      if attributes.key?(:'download_css_and_media')
        self.download_css_and_media = attributes[:'download_css_and_media']
      else
        self.download_css_and_media = false
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

      if attributes.key?(:'use_premium_proxies')
        self.use_premium_proxies = attributes[:'use_premium_proxies']
      else
        self.use_premium_proxies = false
      end
    end

    # Show invalid properties with the reasons. Usually used together with valid?
    # @return Array for valid properties with the reasons
    def list_invalid_properties
      invalid_properties = Array.new
      if @url.nil?
        invalid_properties.push('invalid value for "url", url cannot be nil.')
      end

      if !@max_pages_to_scrape.nil? && @max_pages_to_scrape < 1
        invalid_properties.push('invalid value for "max_pages_to_scrape", must be greater than or equal to 1.')
      end

      if !@url_paths_to_include.nil? && @url_paths_to_include.length > 10
        invalid_properties.push('invalid value for "url_paths_to_include", number of items must be less than or equal to 10.')
      end

      if !@url_paths_to_exclude.nil? && @url_paths_to_exclude.length > 10
        invalid_properties.push('invalid value for "url_paths_to_exclude", number of items must be less than or equal to 10.')
      end

      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      return false if @url.nil?
      return false if !@max_pages_to_scrape.nil? && @max_pages_to_scrape < 1
      return false if !@url_paths_to_include.nil? && @url_paths_to_include.length > 10
      return false if !@url_paths_to_exclude.nil? && @url_paths_to_exclude.length > 10
      true
    end

    # Custom attribute writer method with validation
    # @param [Object] max_pages_to_scrape Value to be assigned
    def max_pages_to_scrape=(max_pages_to_scrape)
      if !max_pages_to_scrape.nil? && max_pages_to_scrape < 1
        fail ArgumentError, 'invalid value for "max_pages_to_scrape", must be greater than or equal to 1.'
      end

      @max_pages_to_scrape = max_pages_to_scrape
    end

    # Custom attribute writer method with validation
    # @param [Object] url_paths_to_include Value to be assigned
    def url_paths_to_include=(url_paths_to_include)
      if !url_paths_to_include.nil? && url_paths_to_include.length > 10
        fail ArgumentError, 'invalid value for "url_paths_to_include", number of items must be less than or equal to 10.'
      end

      @url_paths_to_include = url_paths_to_include
    end

    # Custom attribute writer method with validation
    # @param [Object] url_paths_to_exclude Value to be assigned
    def url_paths_to_exclude=(url_paths_to_exclude)
      if !url_paths_to_exclude.nil? && url_paths_to_exclude.length > 10
        fail ArgumentError, 'invalid value for "url_paths_to_exclude", number of items must be less than or equal to 10.'
      end

      @url_paths_to_exclude = url_paths_to_exclude
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          tags == o.tags &&
          url == o.url &&
          max_pages_to_scrape == o.max_pages_to_scrape &&
          chunk_size == o.chunk_size &&
          chunk_overlap == o.chunk_overlap &&
          skip_embedding_generation == o.skip_embedding_generation &&
          enable_auto_sync == o.enable_auto_sync &&
          generate_sparse_vectors == o.generate_sparse_vectors &&
          prepend_filename_to_chunks == o.prepend_filename_to_chunks &&
          html_tags_to_skip == o.html_tags_to_skip &&
          css_classes_to_skip == o.css_classes_to_skip &&
          css_selectors_to_skip == o.css_selectors_to_skip &&
          embedding_model == o.embedding_model &&
          url_paths_to_include == o.url_paths_to_include &&
          url_paths_to_exclude == o.url_paths_to_exclude &&
          urls_to_scrape == o.urls_to_scrape &&
          download_css_and_media == o.download_css_and_media &&
          generate_chunks_only == o.generate_chunks_only &&
          store_file_only == o.store_file_only &&
          use_premium_proxies == o.use_premium_proxies
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [tags, url, max_pages_to_scrape, chunk_size, chunk_overlap, skip_embedding_generation, enable_auto_sync, generate_sparse_vectors, prepend_filename_to_chunks, html_tags_to_skip, css_classes_to_skip, css_selectors_to_skip, embedding_model, url_paths_to_include, url_paths_to_exclude, urls_to_scrape, download_css_and_media, generate_chunks_only, store_file_only, use_premium_proxies].hash
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

=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


=end

require 'cgi'

module Carbon
  class OrganizationsApi
    attr_accessor :api_client

    def initialize(api_client = ApiClient.default)
      @api_client = api_client
    end

    # Get Organization
    #
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get(extra: {})
      data, _status_code, _headers = get_with_http_info_impl(extra)
      data
    end

    # Get Organization
    #
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_with_http_info(extra: {})
      get_with_http_info_impl(extra)
    end

    # Get Organization
    # @param [Hash] opts the optional parameters
    # @return [OrganizationResponse]
    private def get_impl(opts = {})
      data, _status_code, _headers = get_with_http_info(opts)
      data
    end

    # Get Organization
    # @param [Hash] opts the optional parameters
    # @return [Array<(OrganizationResponse, Integer, Hash)>] OrganizationResponse data, response status code and response headers
    private def get_with_http_info_impl(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: OrganizationsApi.get ...'
      end
      # resource path
      local_var_path = '/organization'

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'OrganizationResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['accessToken']

      new_options = opts.merge(
        :operation => :"OrganizationsApi.get",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: OrganizationsApi#get\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers, response
    end
  end

  # top-level client access to avoid having the user to insantiate their own API instances
  Organizations = OrganizationsApi::new
end

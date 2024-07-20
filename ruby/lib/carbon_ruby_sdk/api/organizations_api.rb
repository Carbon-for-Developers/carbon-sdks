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
      api_response = get_with_http_info_impl(extra)
      api_response.data
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
    # @return [APIResponse] data is OrganizationResponse, status code, headers and response
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
      auth_names = opts[:debug_auth_names] || ['apiKey']

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
      APIResponse::new(data, status_code, headers, response)
    end


    # Update Organization
    #
    # @param global_user_config [UserConfigurationNullable] 
    # @param data_source_configs [Hash<String, DataSourceConfiguration>] Used to set organization level defaults for configuration related to data sources.
    # @param body [UpdateOrganizationInput] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def update(global_user_config: SENTINEL, data_source_configs: SENTINEL, extra: {})
      _body = {}
      _body[:global_user_config] = global_user_config if global_user_config != SENTINEL
      _body[:data_source_configs] = data_source_configs if data_source_configs != SENTINEL
      update_organization_input = _body
      api_response = update_with_http_info_impl(update_organization_input, extra)
      api_response.data
    end

    # Update Organization
    #
    # @param global_user_config [UserConfigurationNullable] 
    # @param data_source_configs [Hash<String, DataSourceConfiguration>] Used to set organization level defaults for configuration related to data sources.
    # @param body [UpdateOrganizationInput] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def update_with_http_info(global_user_config: SENTINEL, data_source_configs: SENTINEL, extra: {})
      _body = {}
      _body[:global_user_config] = global_user_config if global_user_config != SENTINEL
      _body[:data_source_configs] = data_source_configs if data_source_configs != SENTINEL
      update_organization_input = _body
      update_with_http_info_impl(update_organization_input, extra)
    end

    # Update Organization
    # @param update_organization_input [UpdateOrganizationInput] 
    # @param [Hash] opts the optional parameters
    # @return [GenericSuccessResponse]
    private def update_impl(update_organization_input, opts = {})
      data, _status_code, _headers = update_with_http_info(update_organization_input, opts)
      data
    end

    # Update Organization
    # @param update_organization_input [UpdateOrganizationInput] 
    # @param [Hash] opts the optional parameters
    # @return [APIResponse] data is GenericSuccessResponse, status code, headers and response
    private def update_with_http_info_impl(update_organization_input, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: OrganizationsApi.update ...'
      end
      # verify the required parameter 'update_organization_input' is set
      if @api_client.config.client_side_validation && update_organization_input.nil?
        fail ArgumentError, "Missing the required parameter 'update_organization_input' when calling OrganizationsApi.update"
      end
      # resource path
      local_var_path = '/organization/update'

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])
      # HTTP header 'Content-Type'
      content_type = @api_client.select_header_content_type(['application/json'])
      if !content_type.nil?
        header_params['Content-Type'] = content_type
      end

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body] || @api_client.object_to_http_body(update_organization_input)

      # return_type
      return_type = opts[:debug_return_type] || 'GenericSuccessResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['apiKey']

      new_options = opts.merge(
        :operation => :"OrganizationsApi.update",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: OrganizationsApi#update\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end


    # Update Organization Statistics
    #
    # Use this endpoint to reaggregate the statistics for an organization, for example aggregate_file_size. The reaggregation
    # process is asyncronous so a webhook will be sent with the event type being FILE_STATISTICS_AGGREGATED to notify when the
    # process is complee. After this aggregation is complete, the updated statistics can be retrieved using the /organization
    # endpoint. The response of /organization willalso contain a timestamp of the last time the statistics were reaggregated.
    #
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def update_stats(extra: {})
      api_response = update_stats_with_http_info_impl(extra)
      api_response.data
    end

    # Update Organization Statistics
    #
    # Use this endpoint to reaggregate the statistics for an organization, for example aggregate_file_size. The reaggregation
    # process is asyncronous so a webhook will be sent with the event type being FILE_STATISTICS_AGGREGATED to notify when the
    # process is complee. After this aggregation is complete, the updated statistics can be retrieved using the /organization
    # endpoint. The response of /organization willalso contain a timestamp of the last time the statistics were reaggregated.
    #
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def update_stats_with_http_info(extra: {})
      update_stats_with_http_info_impl(extra)
    end

    # Update Organization Statistics
    # Use this endpoint to reaggregate the statistics for an organization, for example aggregate_file_size. The reaggregation process is asyncronous so a webhook will be sent with the event type being FILE_STATISTICS_AGGREGATED to notify when the process is complee. After this aggregation is complete, the updated statistics can be retrieved using the /organization endpoint. The response of /organization willalso contain a timestamp of the last time the statistics were reaggregated.
    # @param [Hash] opts the optional parameters
    # @return [GenericSuccessResponse]
    private def update_stats_impl(opts = {})
      data, _status_code, _headers = update_stats_with_http_info(opts)
      data
    end

    # Update Organization Statistics
    # Use this endpoint to reaggregate the statistics for an organization, for example aggregate_file_size. The reaggregation process is asyncronous so a webhook will be sent with the event type being FILE_STATISTICS_AGGREGATED to notify when the process is complee. After this aggregation is complete, the updated statistics can be retrieved using the /organization endpoint. The response of /organization willalso contain a timestamp of the last time the statistics were reaggregated.
    # @param [Hash] opts the optional parameters
    # @return [APIResponse] data is GenericSuccessResponse, status code, headers and response
    private def update_stats_with_http_info_impl(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: OrganizationsApi.update_stats ...'
      end
      # resource path
      local_var_path = '/organization/statistics'

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
      return_type = opts[:debug_return_type] || 'GenericSuccessResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['apiKey']

      new_options = opts.merge(
        :operation => :"OrganizationsApi.update_stats",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: OrganizationsApi#update_stats\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end
  end

  # top-level client access to avoid having the user to insantiate their own API instances
  Organizations = OrganizationsApi::new
end

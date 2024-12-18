=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'cgi'

module Carbon
  class DataSourcesApi
    attr_accessor :api_client

    def initialize(api_client = ApiClient.default)
      @api_client = api_client
    end

    # Add Data Source Tags
    #
    # @param tags [Object] 
    # @param data_source_id [Integer] 
    # @param body [AddDataSourceTagsInput] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def add_tags(tags:, data_source_id:, extra: {})
      _body = {}
      _body[:tags] = tags if tags != SENTINEL
      _body[:data_source_id] = data_source_id if data_source_id != SENTINEL
      add_data_source_tags_input = _body
      api_response = add_tags_with_http_info_impl(add_data_source_tags_input, extra)
      api_response.data
    end

    # Add Data Source Tags
    #
    # @param tags [Object] 
    # @param data_source_id [Integer] 
    # @param body [AddDataSourceTagsInput] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def add_tags_with_http_info(tags:, data_source_id:, extra: {})
      _body = {}
      _body[:tags] = tags if tags != SENTINEL
      _body[:data_source_id] = data_source_id if data_source_id != SENTINEL
      add_data_source_tags_input = _body
      add_tags_with_http_info_impl(add_data_source_tags_input, extra)
    end

    # Add Data Source Tags
    # @param add_data_source_tags_input [AddDataSourceTagsInput] 
    # @param [Hash] opts the optional parameters
    # @return [OrganizationUserDataSourceAPI]
    private def add_tags_impl(add_data_source_tags_input, opts = {})
      data, _status_code, _headers = add_tags_with_http_info(add_data_source_tags_input, opts)
      data
    end

    # Add Data Source Tags
    # @param add_data_source_tags_input [AddDataSourceTagsInput] 
    # @param [Hash] opts the optional parameters
    # @return [APIResponse] data is OrganizationUserDataSourceAPI, status code, headers and response
    private def add_tags_with_http_info_impl(add_data_source_tags_input, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: DataSourcesApi.add_tags ...'
      end
      # verify the required parameter 'add_data_source_tags_input' is set
      if @api_client.config.client_side_validation && add_data_source_tags_input.nil?
        fail ArgumentError, "Missing the required parameter 'add_data_source_tags_input' when calling DataSourcesApi.add_tags"
      end
      # resource path
      local_var_path = '/data_sources/tags/add'

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
      post_body = opts[:debug_body] || @api_client.object_to_http_body(add_data_source_tags_input)

      # return_type
      return_type = opts[:debug_return_type] || 'OrganizationUserDataSourceAPI'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['accessToken', 'apiKey', 'customerId']

      new_options = opts.merge(
        :operation => :"DataSourcesApi.add_tags",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: DataSourcesApi#add_tags\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end


    # Data Sources
    #
    # @param pagination [Pagination] 
    # @param order_by [OrganizationUserDataSourceOrderByColumns] 
    # @param order_dir [OrderDir] 
    # @param filters [OrganizationUserDataSourceFilters] 
    # @param body [OrganizationUserDataSourceQueryInput] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def query(pagination: SENTINEL, order_by: SENTINEL, order_dir: SENTINEL, filters: SENTINEL, extra: {})
      _body = {}
      _body[:pagination] = pagination if pagination != SENTINEL
      _body[:order_by] = order_by if order_by != SENTINEL
      _body[:order_dir] = order_dir if order_dir != SENTINEL
      _body[:filters] = filters if filters != SENTINEL
      organization_user_data_source_query_input = _body
      api_response = query_with_http_info_impl(organization_user_data_source_query_input, extra)
      api_response.data
    end

    # Data Sources
    #
    # @param pagination [Pagination] 
    # @param order_by [OrganizationUserDataSourceOrderByColumns] 
    # @param order_dir [OrderDir] 
    # @param filters [OrganizationUserDataSourceFilters] 
    # @param body [OrganizationUserDataSourceQueryInput] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def query_with_http_info(pagination: SENTINEL, order_by: SENTINEL, order_dir: SENTINEL, filters: SENTINEL, extra: {})
      _body = {}
      _body[:pagination] = pagination if pagination != SENTINEL
      _body[:order_by] = order_by if order_by != SENTINEL
      _body[:order_dir] = order_dir if order_dir != SENTINEL
      _body[:filters] = filters if filters != SENTINEL
      organization_user_data_source_query_input = _body
      query_with_http_info_impl(organization_user_data_source_query_input, extra)
    end

    # Data Sources
    # @param organization_user_data_source_query_input [OrganizationUserDataSourceQueryInput] 
    # @param [Hash] opts the optional parameters
    # @return [OrganizationUserDataSourceResponse]
    private def query_impl(organization_user_data_source_query_input, opts = {})
      data, _status_code, _headers = query_with_http_info(organization_user_data_source_query_input, opts)
      data
    end

    # Data Sources
    # @param organization_user_data_source_query_input [OrganizationUserDataSourceQueryInput] 
    # @param [Hash] opts the optional parameters
    # @return [APIResponse] data is OrganizationUserDataSourceResponse, status code, headers and response
    private def query_with_http_info_impl(organization_user_data_source_query_input, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: DataSourcesApi.query ...'
      end
      # verify the required parameter 'organization_user_data_source_query_input' is set
      if @api_client.config.client_side_validation && organization_user_data_source_query_input.nil?
        fail ArgumentError, "Missing the required parameter 'organization_user_data_source_query_input' when calling DataSourcesApi.query"
      end
      # resource path
      local_var_path = '/data_sources'

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
      post_body = opts[:debug_body] || @api_client.object_to_http_body(organization_user_data_source_query_input)

      # return_type
      return_type = opts[:debug_return_type] || 'OrganizationUserDataSourceResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['apiKey']

      new_options = opts.merge(
        :operation => :"DataSourcesApi.query",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: DataSourcesApi#query\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end


    # User Data Sources
    #
    # @param pagination [Pagination] 
    # @param order_by [OrganizationUserDataSourceOrderByColumns] 
    # @param order_dir [OrderDir] 
    # @param filters [OrganizationUserDataSourceFilters] 
    # @param body [OrganizationUserDataSourceQueryInput] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def query_user_data_sources(pagination: SENTINEL, order_by: SENTINEL, order_dir: SENTINEL, filters: SENTINEL, extra: {})
      _body = {}
      _body[:pagination] = pagination if pagination != SENTINEL
      _body[:order_by] = order_by if order_by != SENTINEL
      _body[:order_dir] = order_dir if order_dir != SENTINEL
      _body[:filters] = filters if filters != SENTINEL
      organization_user_data_source_query_input = _body
      api_response = query_user_data_sources_with_http_info_impl(organization_user_data_source_query_input, extra)
      api_response.data
    end

    # User Data Sources
    #
    # @param pagination [Pagination] 
    # @param order_by [OrganizationUserDataSourceOrderByColumns] 
    # @param order_dir [OrderDir] 
    # @param filters [OrganizationUserDataSourceFilters] 
    # @param body [OrganizationUserDataSourceQueryInput] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def query_user_data_sources_with_http_info(pagination: SENTINEL, order_by: SENTINEL, order_dir: SENTINEL, filters: SENTINEL, extra: {})
      _body = {}
      _body[:pagination] = pagination if pagination != SENTINEL
      _body[:order_by] = order_by if order_by != SENTINEL
      _body[:order_dir] = order_dir if order_dir != SENTINEL
      _body[:filters] = filters if filters != SENTINEL
      organization_user_data_source_query_input = _body
      query_user_data_sources_with_http_info_impl(organization_user_data_source_query_input, extra)
    end

    # User Data Sources
    # @param organization_user_data_source_query_input [OrganizationUserDataSourceQueryInput] 
    # @param [Hash] opts the optional parameters
    # @return [OrganizationUserDataSourceResponse]
    private def query_user_data_sources_impl(organization_user_data_source_query_input, opts = {})
      data, _status_code, _headers = query_user_data_sources_with_http_info(organization_user_data_source_query_input, opts)
      data
    end

    # User Data Sources
    # @param organization_user_data_source_query_input [OrganizationUserDataSourceQueryInput] 
    # @param [Hash] opts the optional parameters
    # @return [APIResponse] data is OrganizationUserDataSourceResponse, status code, headers and response
    private def query_user_data_sources_with_http_info_impl(organization_user_data_source_query_input, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: DataSourcesApi.query_user_data_sources ...'
      end
      # verify the required parameter 'organization_user_data_source_query_input' is set
      if @api_client.config.client_side_validation && organization_user_data_source_query_input.nil?
        fail ArgumentError, "Missing the required parameter 'organization_user_data_source_query_input' when calling DataSourcesApi.query_user_data_sources"
      end
      # resource path
      local_var_path = '/user_data_sources'

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
      post_body = opts[:debug_body] || @api_client.object_to_http_body(organization_user_data_source_query_input)

      # return_type
      return_type = opts[:debug_return_type] || 'OrganizationUserDataSourceResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['accessToken', 'apiKey', 'customerId']

      new_options = opts.merge(
        :operation => :"DataSourcesApi.query_user_data_sources",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: DataSourcesApi#query_user_data_sources\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end


    # Remove Data Source Tags
    #
    # @param data_source_id [Integer] 
    # @param tags_to_remove [Array<String>] 
    # @param remove_all_tags [Boolean] 
    # @param body [RemoveDataSourceTagsInput] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def remove_tags(data_source_id:, tags_to_remove: SENTINEL, remove_all_tags: false, extra: {})
      _body = {}
      _body[:data_source_id] = data_source_id if data_source_id != SENTINEL
      _body[:tags_to_remove] = tags_to_remove if tags_to_remove != SENTINEL
      _body[:remove_all_tags] = remove_all_tags if remove_all_tags != SENTINEL
      remove_data_source_tags_input = _body
      api_response = remove_tags_with_http_info_impl(remove_data_source_tags_input, extra)
      api_response.data
    end

    # Remove Data Source Tags
    #
    # @param data_source_id [Integer] 
    # @param tags_to_remove [Array<String>] 
    # @param remove_all_tags [Boolean] 
    # @param body [RemoveDataSourceTagsInput] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def remove_tags_with_http_info(data_source_id:, tags_to_remove: SENTINEL, remove_all_tags: false, extra: {})
      _body = {}
      _body[:data_source_id] = data_source_id if data_source_id != SENTINEL
      _body[:tags_to_remove] = tags_to_remove if tags_to_remove != SENTINEL
      _body[:remove_all_tags] = remove_all_tags if remove_all_tags != SENTINEL
      remove_data_source_tags_input = _body
      remove_tags_with_http_info_impl(remove_data_source_tags_input, extra)
    end

    # Remove Data Source Tags
    # @param remove_data_source_tags_input [RemoveDataSourceTagsInput] 
    # @param [Hash] opts the optional parameters
    # @return [OrganizationUserDataSourceAPI]
    private def remove_tags_impl(remove_data_source_tags_input, opts = {})
      data, _status_code, _headers = remove_tags_with_http_info(remove_data_source_tags_input, opts)
      data
    end

    # Remove Data Source Tags
    # @param remove_data_source_tags_input [RemoveDataSourceTagsInput] 
    # @param [Hash] opts the optional parameters
    # @return [APIResponse] data is OrganizationUserDataSourceAPI, status code, headers and response
    private def remove_tags_with_http_info_impl(remove_data_source_tags_input, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: DataSourcesApi.remove_tags ...'
      end
      # verify the required parameter 'remove_data_source_tags_input' is set
      if @api_client.config.client_side_validation && remove_data_source_tags_input.nil?
        fail ArgumentError, "Missing the required parameter 'remove_data_source_tags_input' when calling DataSourcesApi.remove_tags"
      end
      # resource path
      local_var_path = '/data_sources/tags/remove'

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
      post_body = opts[:debug_body] || @api_client.object_to_http_body(remove_data_source_tags_input)

      # return_type
      return_type = opts[:debug_return_type] || 'OrganizationUserDataSourceAPI'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['accessToken', 'apiKey', 'customerId']

      new_options = opts.merge(
        :operation => :"DataSourcesApi.remove_tags",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: DataSourcesApi#remove_tags\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end


    # Revoke Access Token
    #
    # @param data_source_id [Integer] 
    # @param body [RevokeAccessTokenInput] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def revoke_access_token(data_source_id:, extra: {})
      _body = {}
      _body[:data_source_id] = data_source_id if data_source_id != SENTINEL
      revoke_access_token_input = _body
      api_response = revoke_access_token_with_http_info_impl(revoke_access_token_input, extra)
      api_response.data
    end

    # Revoke Access Token
    #
    # @param data_source_id [Integer] 
    # @param body [RevokeAccessTokenInput] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def revoke_access_token_with_http_info(data_source_id:, extra: {})
      _body = {}
      _body[:data_source_id] = data_source_id if data_source_id != SENTINEL
      revoke_access_token_input = _body
      revoke_access_token_with_http_info_impl(revoke_access_token_input, extra)
    end

    # Revoke Access Token
    # @param revoke_access_token_input [RevokeAccessTokenInput] 
    # @param [Hash] opts the optional parameters
    # @return [GenericSuccessResponse]
    private def revoke_access_token_impl(revoke_access_token_input, opts = {})
      data, _status_code, _headers = revoke_access_token_with_http_info(revoke_access_token_input, opts)
      data
    end

    # Revoke Access Token
    # @param revoke_access_token_input [RevokeAccessTokenInput] 
    # @param [Hash] opts the optional parameters
    # @return [APIResponse] data is GenericSuccessResponse, status code, headers and response
    private def revoke_access_token_with_http_info_impl(revoke_access_token_input, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: DataSourcesApi.revoke_access_token ...'
      end
      # verify the required parameter 'revoke_access_token_input' is set
      if @api_client.config.client_side_validation && revoke_access_token_input.nil?
        fail ArgumentError, "Missing the required parameter 'revoke_access_token_input' when calling DataSourcesApi.revoke_access_token"
      end
      # resource path
      local_var_path = '/revoke_access_token'

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
      post_body = opts[:debug_body] || @api_client.object_to_http_body(revoke_access_token_input)

      # return_type
      return_type = opts[:debug_return_type] || 'GenericSuccessResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['accessToken', 'apiKey', 'customerId']

      new_options = opts.merge(
        :operation => :"DataSourcesApi.revoke_access_token",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: DataSourcesApi#revoke_access_token\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end
  end

  # top-level client access to avoid having the user to insantiate their own API instances
  DataSources = DataSourcesApi::new
end

=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'cgi'

module Carbon
  class WhiteLabelApi
    attr_accessor :api_client

    def initialize(api_client = ApiClient.default)
      @api_client = api_client
    end

    # Create White Labels
    #
    # @param body [Array<WhiteLabelCreateRequestInner>] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def create(body:, extra: {})
      white_label_create_request_inner = body
      api_response = create_with_http_info_impl(white_label_create_request_inner, extra)
      api_response.data
    end

    # Create White Labels
    #
    # @param body [Array<WhiteLabelCreateRequestInner>] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def create_with_http_info(body:, extra: {})
      white_label_create_request_inner = body
      create_with_http_info_impl(white_label_create_request_inner, extra)
    end

    # Create White Labels
    # @param white_label_create_request_inner [Array<WhiteLabelCreateRequestInner>] 
    # @param [Hash] opts the optional parameters
    # @return [Object]
    private def create_impl(white_label_create_request_inner, opts = {})
      data, _status_code, _headers = create_with_http_info(white_label_create_request_inner, opts)
      data
    end

    # Create White Labels
    # @param white_label_create_request_inner [Array<WhiteLabelCreateRequestInner>] 
    # @param [Hash] opts the optional parameters
    # @return [APIResponse] data is Object, status code, headers and response
    private def create_with_http_info_impl(white_label_create_request_inner, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: WhiteLabelApi.create ...'
      end
      # verify the required parameter 'white_label_create_request_inner' is set
      if @api_client.config.client_side_validation && white_label_create_request_inner.nil?
        fail ArgumentError, "Missing the required parameter 'white_label_create_request_inner' when calling WhiteLabelApi.create"
      end
      # resource path
      local_var_path = '/white_label/create'

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
      post_body = opts[:debug_body] || @api_client.object_to_http_body(white_label_create_request_inner)

      # return_type
      return_type = opts[:debug_return_type] || 'Object'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['apiKey']

      new_options = opts.merge(
        :operation => :"WhiteLabelApi.create",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: WhiteLabelApi#create\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end


    # Delete White Labels
    #
    # @param ids [Array<Integer>] 
    # @param body [DeleteWhiteLabelRequest] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def delete(ids:, extra: {})
      _body = {}
      _body[:ids] = ids if ids != SENTINEL
      delete_white_label_request = _body
      api_response = delete_with_http_info_impl(delete_white_label_request, extra)
      api_response.data
    end

    # Delete White Labels
    #
    # @param ids [Array<Integer>] 
    # @param body [DeleteWhiteLabelRequest] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def delete_with_http_info(ids:, extra: {})
      _body = {}
      _body[:ids] = ids if ids != SENTINEL
      delete_white_label_request = _body
      delete_with_http_info_impl(delete_white_label_request, extra)
    end

    # Delete White Labels
    # @param delete_white_label_request [DeleteWhiteLabelRequest] 
    # @param [Hash] opts the optional parameters
    # @return [Object]
    private def delete_impl(delete_white_label_request, opts = {})
      data, _status_code, _headers = delete_with_http_info(delete_white_label_request, opts)
      data
    end

    # Delete White Labels
    # @param delete_white_label_request [DeleteWhiteLabelRequest] 
    # @param [Hash] opts the optional parameters
    # @return [APIResponse] data is Object, status code, headers and response
    private def delete_with_http_info_impl(delete_white_label_request, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: WhiteLabelApi.delete ...'
      end
      # verify the required parameter 'delete_white_label_request' is set
      if @api_client.config.client_side_validation && delete_white_label_request.nil?
        fail ArgumentError, "Missing the required parameter 'delete_white_label_request' when calling WhiteLabelApi.delete"
      end
      # resource path
      local_var_path = '/white_label/delete'

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
      post_body = opts[:debug_body] || @api_client.object_to_http_body(delete_white_label_request)

      # return_type
      return_type = opts[:debug_return_type] || 'Object'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['apiKey']

      new_options = opts.merge(
        :operation => :"WhiteLabelApi.delete",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: WhiteLabelApi#delete\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end


    # List White Labels
    #
    # @param pagination [Pagination] 
    # @param order_by [WhiteLabelOrderByColumns] 
    # @param order_dir [OrderDir] 
    # @param filters [WhiteLabelFilters] 
    # @param body [ListWhiteLabelRequest] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def list(pagination: SENTINEL, order_by: 'created_at', order_dir: 'asc', filters: SENTINEL, extra: {})
      _body = {}
      _body[:pagination] = pagination if pagination != SENTINEL
      _body[:order_by] = order_by if order_by != SENTINEL
      _body[:order_dir] = order_dir if order_dir != SENTINEL
      _body[:filters] = filters if filters != SENTINEL
      list_white_label_request = _body
      api_response = list_with_http_info_impl(list_white_label_request, extra)
      api_response.data
    end

    # List White Labels
    #
    # @param pagination [Pagination] 
    # @param order_by [WhiteLabelOrderByColumns] 
    # @param order_dir [OrderDir] 
    # @param filters [WhiteLabelFilters] 
    # @param body [ListWhiteLabelRequest] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def list_with_http_info(pagination: SENTINEL, order_by: 'created_at', order_dir: 'asc', filters: SENTINEL, extra: {})
      _body = {}
      _body[:pagination] = pagination if pagination != SENTINEL
      _body[:order_by] = order_by if order_by != SENTINEL
      _body[:order_dir] = order_dir if order_dir != SENTINEL
      _body[:filters] = filters if filters != SENTINEL
      list_white_label_request = _body
      list_with_http_info_impl(list_white_label_request, extra)
    end

    # List White Labels
    # @param list_white_label_request [ListWhiteLabelRequest] 
    # @param [Hash] opts the optional parameters
    # @return [Object]
    private def list_impl(list_white_label_request, opts = {})
      data, _status_code, _headers = list_with_http_info(list_white_label_request, opts)
      data
    end

    # List White Labels
    # @param list_white_label_request [ListWhiteLabelRequest] 
    # @param [Hash] opts the optional parameters
    # @return [APIResponse] data is Object, status code, headers and response
    private def list_with_http_info_impl(list_white_label_request, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: WhiteLabelApi.list ...'
      end
      # verify the required parameter 'list_white_label_request' is set
      if @api_client.config.client_side_validation && list_white_label_request.nil?
        fail ArgumentError, "Missing the required parameter 'list_white_label_request' when calling WhiteLabelApi.list"
      end
      # resource path
      local_var_path = '/white_label/list'

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
      post_body = opts[:debug_body] || @api_client.object_to_http_body(list_white_label_request)

      # return_type
      return_type = opts[:debug_return_type] || 'Object'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['apiKey']

      new_options = opts.merge(
        :operation => :"WhiteLabelApi.list",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: WhiteLabelApi#list\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end


    # Update White Label
    #
    # @param data_source_type [WhiteLabelInputDataSourceType] 
    # @param credentials [Credentials] 
    # @param body [WhiteLabelUpdateRequest] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def update(data_source_type: SENTINEL, credentials: SENTINEL, body:, extra: {})
      white_label_update_request = body
      api_response = update_with_http_info_impl(white_label_update_request, extra)
      api_response.data
    end

    # Update White Label
    #
    # @param data_source_type [WhiteLabelInputDataSourceType] 
    # @param credentials [Credentials] 
    # @param body [WhiteLabelUpdateRequest] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def update_with_http_info(data_source_type: SENTINEL, credentials: SENTINEL, body:, extra: {})
      white_label_update_request = body
      update_with_http_info_impl(white_label_update_request, extra)
    end

    # Update White Label
    # @param white_label_update_request [WhiteLabelUpdateRequest] 
    # @param [Hash] opts the optional parameters
    # @return [Object]
    private def update_impl(white_label_update_request, opts = {})
      data, _status_code, _headers = update_with_http_info(white_label_update_request, opts)
      data
    end

    # Update White Label
    # @param white_label_update_request [WhiteLabelUpdateRequest] 
    # @param [Hash] opts the optional parameters
    # @return [APIResponse] data is Object, status code, headers and response
    private def update_with_http_info_impl(white_label_update_request, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: WhiteLabelApi.update ...'
      end
      # verify the required parameter 'white_label_update_request' is set
      if @api_client.config.client_side_validation && white_label_update_request.nil?
        fail ArgumentError, "Missing the required parameter 'white_label_update_request' when calling WhiteLabelApi.update"
      end
      # resource path
      local_var_path = '/white_label/update'

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
      post_body = opts[:debug_body] || @api_client.object_to_http_body(white_label_update_request)

      # return_type
      return_type = opts[:debug_return_type] || 'Object'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['apiKey']

      new_options = opts.merge(
        :operation => :"WhiteLabelApi.update",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: WhiteLabelApi#update\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end
  end

  # top-level client access to avoid having the user to insantiate their own API instances
  WhiteLabel = WhiteLabelApi::new
end

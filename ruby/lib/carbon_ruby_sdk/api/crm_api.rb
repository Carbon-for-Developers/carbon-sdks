=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'cgi'

module Carbon
  class CRMApi
    attr_accessor :api_client

    def initialize(api_client = ApiClient.default)
      @api_client = api_client
    end

    # Get Account
    #
    # @param id [String] 
    # @param data_source_id [Integer] 
    # @param include_remote_data [Boolean] 
    # @param includes [Array<BaseIncludes>] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_account(id:, data_source_id:, include_remote_data: false, includes: SENTINEL, extra: {})
      extra[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      extra[:includes] = includes if includes != SENTINEL
      api_response = get_account_with_http_info_impl(id, data_source_id, extra)
      api_response.data
    end

    # Get Account
    #
    # @param id [String] 
    # @param data_source_id [Integer] 
    # @param include_remote_data [Boolean] 
    # @param includes [Array<BaseIncludes>] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_account_with_http_info(id:, data_source_id:, include_remote_data: false, includes: SENTINEL, extra: {})
      extra[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      extra[:includes] = includes if includes != SENTINEL
      get_account_with_http_info_impl(id, data_source_id, extra)
    end

    # Get Account
    # @param id [String] 
    # @param data_source_id [Integer] 
    # @param [Hash] opts the optional parameters
    # @option opts [Boolean] :include_remote_data  (default to false)
    # @option opts [Array<BaseIncludes>] :includes 
    # @return [Account]
    private def get_account_impl(id, data_source_id, opts = {})
      data, _status_code, _headers = get_account_with_http_info(id, data_source_id, opts)
      data
    end

    # Get Account
    # @param id [String] 
    # @param data_source_id [Integer] 
    # @param [Hash] opts the optional parameters
    # @option opts [Boolean] :include_remote_data  (default to false)
    # @option opts [Array<BaseIncludes>] :includes 
    # @return [APIResponse] data is Account, status code, headers and response
    private def get_account_with_http_info_impl(id, data_source_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: CRMApi.get_account ...'
      end
      # verify the required parameter 'id' is set
      if @api_client.config.client_side_validation && id.nil?
        fail ArgumentError, "Missing the required parameter 'id' when calling CRMApi.get_account"
      end
      # verify the required parameter 'data_source_id' is set
      if @api_client.config.client_side_validation && data_source_id.nil?
        fail ArgumentError, "Missing the required parameter 'data_source_id' when calling CRMApi.get_account"
      end
      # resource path
      local_var_path = '/integrations/data/crm/accounts/{id}'.sub('{' + 'id' + '}', CGI.escape(id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}
      query_params[:'data_source_id'] = data_source_id
      query_params[:'include_remote_data'] = opts[:'include_remote_data'] if !opts[:'include_remote_data'].nil?
      query_params[:'includes'] = @api_client.build_collection_param(opts[:'includes'], :multi) if !opts[:'includes'].nil?

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'Account'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['accessToken', 'apiKey', 'customerId']

      new_options = opts.merge(
        :operation => :"CRMApi.get_account",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: CRMApi#get_account\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end


    # Get Accounts
    #
    # @param data_source_id [Integer] 
    # @param include_remote_data [Boolean] 
    # @param next_cursor [String] 
    # @param page_size [Integer] 
    # @param order_dir [OrderDirV2Nullable] 
    # @param includes [Array<BaseIncludes>] 
    # @param filters [AccountFilters] 
    # @param order_by [AccountsOrderByNullable] 
    # @param body [AccountsRequest] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_accounts(data_source_id:, include_remote_data: false, next_cursor: SENTINEL, page_size: SENTINEL, order_dir: 'asc', includes: SENTINEL, filters: SENTINEL, order_by: SENTINEL, extra: {})
      _body = {}
      _body[:data_source_id] = data_source_id if data_source_id != SENTINEL
      _body[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      _body[:next_cursor] = next_cursor if next_cursor != SENTINEL
      _body[:page_size] = page_size if page_size != SENTINEL
      _body[:order_dir] = order_dir if order_dir != SENTINEL
      _body[:includes] = includes if includes != SENTINEL
      _body[:filters] = filters if filters != SENTINEL
      _body[:order_by] = order_by if order_by != SENTINEL
      accounts_request = _body
      api_response = get_accounts_with_http_info_impl(accounts_request, extra)
      api_response.data
    end

    # Get Accounts
    #
    # @param data_source_id [Integer] 
    # @param include_remote_data [Boolean] 
    # @param next_cursor [String] 
    # @param page_size [Integer] 
    # @param order_dir [OrderDirV2Nullable] 
    # @param includes [Array<BaseIncludes>] 
    # @param filters [AccountFilters] 
    # @param order_by [AccountsOrderByNullable] 
    # @param body [AccountsRequest] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_accounts_with_http_info(data_source_id:, include_remote_data: false, next_cursor: SENTINEL, page_size: SENTINEL, order_dir: 'asc', includes: SENTINEL, filters: SENTINEL, order_by: SENTINEL, extra: {})
      _body = {}
      _body[:data_source_id] = data_source_id if data_source_id != SENTINEL
      _body[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      _body[:next_cursor] = next_cursor if next_cursor != SENTINEL
      _body[:page_size] = page_size if page_size != SENTINEL
      _body[:order_dir] = order_dir if order_dir != SENTINEL
      _body[:includes] = includes if includes != SENTINEL
      _body[:filters] = filters if filters != SENTINEL
      _body[:order_by] = order_by if order_by != SENTINEL
      accounts_request = _body
      get_accounts_with_http_info_impl(accounts_request, extra)
    end

    # Get Accounts
    # @param accounts_request [AccountsRequest] 
    # @param [Hash] opts the optional parameters
    # @return [AccountResponse]
    private def get_accounts_impl(accounts_request, opts = {})
      data, _status_code, _headers = get_accounts_with_http_info(accounts_request, opts)
      data
    end

    # Get Accounts
    # @param accounts_request [AccountsRequest] 
    # @param [Hash] opts the optional parameters
    # @return [APIResponse] data is AccountResponse, status code, headers and response
    private def get_accounts_with_http_info_impl(accounts_request, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: CRMApi.get_accounts ...'
      end
      # verify the required parameter 'accounts_request' is set
      if @api_client.config.client_side_validation && accounts_request.nil?
        fail ArgumentError, "Missing the required parameter 'accounts_request' when calling CRMApi.get_accounts"
      end
      # resource path
      local_var_path = '/integrations/data/crm/accounts'

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
      post_body = opts[:debug_body] || @api_client.object_to_http_body(accounts_request)

      # return_type
      return_type = opts[:debug_return_type] || 'AccountResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['accessToken', 'apiKey', 'customerId']

      new_options = opts.merge(
        :operation => :"CRMApi.get_accounts",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: CRMApi#get_accounts\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end


    # Get Contact
    #
    # @param id [String] 
    # @param data_source_id [Integer] 
    # @param include_remote_data [Boolean] 
    # @param includes [Array<BaseIncludes>] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_contact(id:, data_source_id:, include_remote_data: false, includes: SENTINEL, extra: {})
      extra[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      extra[:includes] = includes if includes != SENTINEL
      api_response = get_contact_with_http_info_impl(id, data_source_id, extra)
      api_response.data
    end

    # Get Contact
    #
    # @param id [String] 
    # @param data_source_id [Integer] 
    # @param include_remote_data [Boolean] 
    # @param includes [Array<BaseIncludes>] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_contact_with_http_info(id:, data_source_id:, include_remote_data: false, includes: SENTINEL, extra: {})
      extra[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      extra[:includes] = includes if includes != SENTINEL
      get_contact_with_http_info_impl(id, data_source_id, extra)
    end

    # Get Contact
    # @param id [String] 
    # @param data_source_id [Integer] 
    # @param [Hash] opts the optional parameters
    # @option opts [Boolean] :include_remote_data  (default to false)
    # @option opts [Array<BaseIncludes>] :includes 
    # @return [Contact]
    private def get_contact_impl(id, data_source_id, opts = {})
      data, _status_code, _headers = get_contact_with_http_info(id, data_source_id, opts)
      data
    end

    # Get Contact
    # @param id [String] 
    # @param data_source_id [Integer] 
    # @param [Hash] opts the optional parameters
    # @option opts [Boolean] :include_remote_data  (default to false)
    # @option opts [Array<BaseIncludes>] :includes 
    # @return [APIResponse] data is Contact, status code, headers and response
    private def get_contact_with_http_info_impl(id, data_source_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: CRMApi.get_contact ...'
      end
      # verify the required parameter 'id' is set
      if @api_client.config.client_side_validation && id.nil?
        fail ArgumentError, "Missing the required parameter 'id' when calling CRMApi.get_contact"
      end
      # verify the required parameter 'data_source_id' is set
      if @api_client.config.client_side_validation && data_source_id.nil?
        fail ArgumentError, "Missing the required parameter 'data_source_id' when calling CRMApi.get_contact"
      end
      # resource path
      local_var_path = '/integrations/data/crm/contacts/{id}'.sub('{' + 'id' + '}', CGI.escape(id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}
      query_params[:'data_source_id'] = data_source_id
      query_params[:'include_remote_data'] = opts[:'include_remote_data'] if !opts[:'include_remote_data'].nil?
      query_params[:'includes'] = @api_client.build_collection_param(opts[:'includes'], :multi) if !opts[:'includes'].nil?

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'Contact'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['accessToken', 'apiKey', 'customerId']

      new_options = opts.merge(
        :operation => :"CRMApi.get_contact",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: CRMApi#get_contact\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end


    # Get Contacts
    #
    # @param data_source_id [Integer] 
    # @param include_remote_data [Boolean] 
    # @param next_cursor [String] 
    # @param page_size [Integer] 
    # @param order_dir [OrderDirV2Nullable] 
    # @param includes [Array<BaseIncludes>] 
    # @param filters [ContactFilters] 
    # @param order_by [ContactsOrderByNullable] 
    # @param body [ContactsRequest] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_contacts(data_source_id:, include_remote_data: false, next_cursor: SENTINEL, page_size: SENTINEL, order_dir: 'asc', includes: SENTINEL, filters: SENTINEL, order_by: SENTINEL, extra: {})
      _body = {}
      _body[:data_source_id] = data_source_id if data_source_id != SENTINEL
      _body[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      _body[:next_cursor] = next_cursor if next_cursor != SENTINEL
      _body[:page_size] = page_size if page_size != SENTINEL
      _body[:order_dir] = order_dir if order_dir != SENTINEL
      _body[:includes] = includes if includes != SENTINEL
      _body[:filters] = filters if filters != SENTINEL
      _body[:order_by] = order_by if order_by != SENTINEL
      contacts_request = _body
      api_response = get_contacts_with_http_info_impl(contacts_request, extra)
      api_response.data
    end

    # Get Contacts
    #
    # @param data_source_id [Integer] 
    # @param include_remote_data [Boolean] 
    # @param next_cursor [String] 
    # @param page_size [Integer] 
    # @param order_dir [OrderDirV2Nullable] 
    # @param includes [Array<BaseIncludes>] 
    # @param filters [ContactFilters] 
    # @param order_by [ContactsOrderByNullable] 
    # @param body [ContactsRequest] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_contacts_with_http_info(data_source_id:, include_remote_data: false, next_cursor: SENTINEL, page_size: SENTINEL, order_dir: 'asc', includes: SENTINEL, filters: SENTINEL, order_by: SENTINEL, extra: {})
      _body = {}
      _body[:data_source_id] = data_source_id if data_source_id != SENTINEL
      _body[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      _body[:next_cursor] = next_cursor if next_cursor != SENTINEL
      _body[:page_size] = page_size if page_size != SENTINEL
      _body[:order_dir] = order_dir if order_dir != SENTINEL
      _body[:includes] = includes if includes != SENTINEL
      _body[:filters] = filters if filters != SENTINEL
      _body[:order_by] = order_by if order_by != SENTINEL
      contacts_request = _body
      get_contacts_with_http_info_impl(contacts_request, extra)
    end

    # Get Contacts
    # @param contacts_request [ContactsRequest] 
    # @param [Hash] opts the optional parameters
    # @return [ContactsResponse]
    private def get_contacts_impl(contacts_request, opts = {})
      data, _status_code, _headers = get_contacts_with_http_info(contacts_request, opts)
      data
    end

    # Get Contacts
    # @param contacts_request [ContactsRequest] 
    # @param [Hash] opts the optional parameters
    # @return [APIResponse] data is ContactsResponse, status code, headers and response
    private def get_contacts_with_http_info_impl(contacts_request, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: CRMApi.get_contacts ...'
      end
      # verify the required parameter 'contacts_request' is set
      if @api_client.config.client_side_validation && contacts_request.nil?
        fail ArgumentError, "Missing the required parameter 'contacts_request' when calling CRMApi.get_contacts"
      end
      # resource path
      local_var_path = '/integrations/data/crm/contacts'

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
      post_body = opts[:debug_body] || @api_client.object_to_http_body(contacts_request)

      # return_type
      return_type = opts[:debug_return_type] || 'ContactsResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['accessToken', 'apiKey', 'customerId']

      new_options = opts.merge(
        :operation => :"CRMApi.get_contacts",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: CRMApi#get_contacts\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end


    # Get Lead
    #
    # @param id [String] 
    # @param data_source_id [Integer] 
    # @param include_remote_data [Boolean] 
    # @param includes [Array<BaseIncludes>] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_lead(id:, data_source_id:, include_remote_data: false, includes: SENTINEL, extra: {})
      extra[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      extra[:includes] = includes if includes != SENTINEL
      api_response = get_lead_with_http_info_impl(id, data_source_id, extra)
      api_response.data
    end

    # Get Lead
    #
    # @param id [String] 
    # @param data_source_id [Integer] 
    # @param include_remote_data [Boolean] 
    # @param includes [Array<BaseIncludes>] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_lead_with_http_info(id:, data_source_id:, include_remote_data: false, includes: SENTINEL, extra: {})
      extra[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      extra[:includes] = includes if includes != SENTINEL
      get_lead_with_http_info_impl(id, data_source_id, extra)
    end

    # Get Lead
    # @param id [String] 
    # @param data_source_id [Integer] 
    # @param [Hash] opts the optional parameters
    # @option opts [Boolean] :include_remote_data  (default to false)
    # @option opts [Array<BaseIncludes>] :includes 
    # @return [Lead]
    private def get_lead_impl(id, data_source_id, opts = {})
      data, _status_code, _headers = get_lead_with_http_info(id, data_source_id, opts)
      data
    end

    # Get Lead
    # @param id [String] 
    # @param data_source_id [Integer] 
    # @param [Hash] opts the optional parameters
    # @option opts [Boolean] :include_remote_data  (default to false)
    # @option opts [Array<BaseIncludes>] :includes 
    # @return [APIResponse] data is Lead, status code, headers and response
    private def get_lead_with_http_info_impl(id, data_source_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: CRMApi.get_lead ...'
      end
      # verify the required parameter 'id' is set
      if @api_client.config.client_side_validation && id.nil?
        fail ArgumentError, "Missing the required parameter 'id' when calling CRMApi.get_lead"
      end
      # verify the required parameter 'data_source_id' is set
      if @api_client.config.client_side_validation && data_source_id.nil?
        fail ArgumentError, "Missing the required parameter 'data_source_id' when calling CRMApi.get_lead"
      end
      # resource path
      local_var_path = '/integrations/data/crm/leads/{id}'.sub('{' + 'id' + '}', CGI.escape(id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}
      query_params[:'data_source_id'] = data_source_id
      query_params[:'include_remote_data'] = opts[:'include_remote_data'] if !opts[:'include_remote_data'].nil?
      query_params[:'includes'] = @api_client.build_collection_param(opts[:'includes'], :multi) if !opts[:'includes'].nil?

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'Lead'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['accessToken', 'apiKey', 'customerId']

      new_options = opts.merge(
        :operation => :"CRMApi.get_lead",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: CRMApi#get_lead\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end


    # Get Leads
    #
    # @param data_source_id [Integer] 
    # @param include_remote_data [Boolean] 
    # @param next_cursor [String] 
    # @param page_size [Integer] 
    # @param order_dir [OrderDirV2Nullable] 
    # @param includes [Array<BaseIncludes>] 
    # @param filters [LeadFilters] 
    # @param order_by [LeadsOrderByNullable] 
    # @param body [LeadsRequest] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_leads(data_source_id:, include_remote_data: false, next_cursor: SENTINEL, page_size: SENTINEL, order_dir: 'asc', includes: SENTINEL, filters: SENTINEL, order_by: SENTINEL, extra: {})
      _body = {}
      _body[:data_source_id] = data_source_id if data_source_id != SENTINEL
      _body[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      _body[:next_cursor] = next_cursor if next_cursor != SENTINEL
      _body[:page_size] = page_size if page_size != SENTINEL
      _body[:order_dir] = order_dir if order_dir != SENTINEL
      _body[:includes] = includes if includes != SENTINEL
      _body[:filters] = filters if filters != SENTINEL
      _body[:order_by] = order_by if order_by != SENTINEL
      leads_request = _body
      api_response = get_leads_with_http_info_impl(leads_request, extra)
      api_response.data
    end

    # Get Leads
    #
    # @param data_source_id [Integer] 
    # @param include_remote_data [Boolean] 
    # @param next_cursor [String] 
    # @param page_size [Integer] 
    # @param order_dir [OrderDirV2Nullable] 
    # @param includes [Array<BaseIncludes>] 
    # @param filters [LeadFilters] 
    # @param order_by [LeadsOrderByNullable] 
    # @param body [LeadsRequest] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_leads_with_http_info(data_source_id:, include_remote_data: false, next_cursor: SENTINEL, page_size: SENTINEL, order_dir: 'asc', includes: SENTINEL, filters: SENTINEL, order_by: SENTINEL, extra: {})
      _body = {}
      _body[:data_source_id] = data_source_id if data_source_id != SENTINEL
      _body[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      _body[:next_cursor] = next_cursor if next_cursor != SENTINEL
      _body[:page_size] = page_size if page_size != SENTINEL
      _body[:order_dir] = order_dir if order_dir != SENTINEL
      _body[:includes] = includes if includes != SENTINEL
      _body[:filters] = filters if filters != SENTINEL
      _body[:order_by] = order_by if order_by != SENTINEL
      leads_request = _body
      get_leads_with_http_info_impl(leads_request, extra)
    end

    # Get Leads
    # @param leads_request [LeadsRequest] 
    # @param [Hash] opts the optional parameters
    # @return [LeadsResponse]
    private def get_leads_impl(leads_request, opts = {})
      data, _status_code, _headers = get_leads_with_http_info(leads_request, opts)
      data
    end

    # Get Leads
    # @param leads_request [LeadsRequest] 
    # @param [Hash] opts the optional parameters
    # @return [APIResponse] data is LeadsResponse, status code, headers and response
    private def get_leads_with_http_info_impl(leads_request, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: CRMApi.get_leads ...'
      end
      # verify the required parameter 'leads_request' is set
      if @api_client.config.client_side_validation && leads_request.nil?
        fail ArgumentError, "Missing the required parameter 'leads_request' when calling CRMApi.get_leads"
      end
      # resource path
      local_var_path = '/integrations/data/crm/leads'

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
      post_body = opts[:debug_body] || @api_client.object_to_http_body(leads_request)

      # return_type
      return_type = opts[:debug_return_type] || 'LeadsResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['accessToken', 'apiKey', 'customerId']

      new_options = opts.merge(
        :operation => :"CRMApi.get_leads",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: CRMApi#get_leads\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end


    # Get Opportunities
    #
    # @param data_source_id [Integer] 
    # @param include_remote_data [Boolean] 
    # @param next_cursor [String] 
    # @param page_size [Integer] 
    # @param order_dir [OrderDirV2Nullable] 
    # @param includes [Array<BaseIncludes>] 
    # @param filters [OpportunityFilters] 
    # @param order_by [OpportunitiesOrderByNullable] 
    # @param body [OpportunitiesRequest] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_opportunities(data_source_id:, include_remote_data: false, next_cursor: SENTINEL, page_size: SENTINEL, order_dir: 'asc', includes: SENTINEL, filters: SENTINEL, order_by: SENTINEL, extra: {})
      _body = {}
      _body[:data_source_id] = data_source_id if data_source_id != SENTINEL
      _body[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      _body[:next_cursor] = next_cursor if next_cursor != SENTINEL
      _body[:page_size] = page_size if page_size != SENTINEL
      _body[:order_dir] = order_dir if order_dir != SENTINEL
      _body[:includes] = includes if includes != SENTINEL
      _body[:filters] = filters if filters != SENTINEL
      _body[:order_by] = order_by if order_by != SENTINEL
      opportunities_request = _body
      api_response = get_opportunities_with_http_info_impl(opportunities_request, extra)
      api_response.data
    end

    # Get Opportunities
    #
    # @param data_source_id [Integer] 
    # @param include_remote_data [Boolean] 
    # @param next_cursor [String] 
    # @param page_size [Integer] 
    # @param order_dir [OrderDirV2Nullable] 
    # @param includes [Array<BaseIncludes>] 
    # @param filters [OpportunityFilters] 
    # @param order_by [OpportunitiesOrderByNullable] 
    # @param body [OpportunitiesRequest] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_opportunities_with_http_info(data_source_id:, include_remote_data: false, next_cursor: SENTINEL, page_size: SENTINEL, order_dir: 'asc', includes: SENTINEL, filters: SENTINEL, order_by: SENTINEL, extra: {})
      _body = {}
      _body[:data_source_id] = data_source_id if data_source_id != SENTINEL
      _body[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      _body[:next_cursor] = next_cursor if next_cursor != SENTINEL
      _body[:page_size] = page_size if page_size != SENTINEL
      _body[:order_dir] = order_dir if order_dir != SENTINEL
      _body[:includes] = includes if includes != SENTINEL
      _body[:filters] = filters if filters != SENTINEL
      _body[:order_by] = order_by if order_by != SENTINEL
      opportunities_request = _body
      get_opportunities_with_http_info_impl(opportunities_request, extra)
    end

    # Get Opportunities
    # @param opportunities_request [OpportunitiesRequest] 
    # @param [Hash] opts the optional parameters
    # @return [OpportunitiesResponse]
    private def get_opportunities_impl(opportunities_request, opts = {})
      data, _status_code, _headers = get_opportunities_with_http_info(opportunities_request, opts)
      data
    end

    # Get Opportunities
    # @param opportunities_request [OpportunitiesRequest] 
    # @param [Hash] opts the optional parameters
    # @return [APIResponse] data is OpportunitiesResponse, status code, headers and response
    private def get_opportunities_with_http_info_impl(opportunities_request, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: CRMApi.get_opportunities ...'
      end
      # verify the required parameter 'opportunities_request' is set
      if @api_client.config.client_side_validation && opportunities_request.nil?
        fail ArgumentError, "Missing the required parameter 'opportunities_request' when calling CRMApi.get_opportunities"
      end
      # resource path
      local_var_path = '/integrations/data/crm/opportunities'

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
      post_body = opts[:debug_body] || @api_client.object_to_http_body(opportunities_request)

      # return_type
      return_type = opts[:debug_return_type] || 'OpportunitiesResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['accessToken', 'apiKey', 'customerId']

      new_options = opts.merge(
        :operation => :"CRMApi.get_opportunities",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: CRMApi#get_opportunities\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end


    # Get Opportunity
    #
    # @param id [String] 
    # @param data_source_id [Integer] 
    # @param include_remote_data [Boolean] 
    # @param includes [Array<BaseIncludes>] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_opportunity(id:, data_source_id:, include_remote_data: false, includes: SENTINEL, extra: {})
      extra[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      extra[:includes] = includes if includes != SENTINEL
      api_response = get_opportunity_with_http_info_impl(id, data_source_id, extra)
      api_response.data
    end

    # Get Opportunity
    #
    # @param id [String] 
    # @param data_source_id [Integer] 
    # @param include_remote_data [Boolean] 
    # @param includes [Array<BaseIncludes>] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_opportunity_with_http_info(id:, data_source_id:, include_remote_data: false, includes: SENTINEL, extra: {})
      extra[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      extra[:includes] = includes if includes != SENTINEL
      get_opportunity_with_http_info_impl(id, data_source_id, extra)
    end

    # Get Opportunity
    # @param id [String] 
    # @param data_source_id [Integer] 
    # @param [Hash] opts the optional parameters
    # @option opts [Boolean] :include_remote_data  (default to false)
    # @option opts [Array<BaseIncludes>] :includes 
    # @return [Opportunity]
    private def get_opportunity_impl(id, data_source_id, opts = {})
      data, _status_code, _headers = get_opportunity_with_http_info(id, data_source_id, opts)
      data
    end

    # Get Opportunity
    # @param id [String] 
    # @param data_source_id [Integer] 
    # @param [Hash] opts the optional parameters
    # @option opts [Boolean] :include_remote_data  (default to false)
    # @option opts [Array<BaseIncludes>] :includes 
    # @return [APIResponse] data is Opportunity, status code, headers and response
    private def get_opportunity_with_http_info_impl(id, data_source_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: CRMApi.get_opportunity ...'
      end
      # verify the required parameter 'id' is set
      if @api_client.config.client_side_validation && id.nil?
        fail ArgumentError, "Missing the required parameter 'id' when calling CRMApi.get_opportunity"
      end
      # verify the required parameter 'data_source_id' is set
      if @api_client.config.client_side_validation && data_source_id.nil?
        fail ArgumentError, "Missing the required parameter 'data_source_id' when calling CRMApi.get_opportunity"
      end
      # resource path
      local_var_path = '/integrations/data/crm/opportunities/{id}'.sub('{' + 'id' + '}', CGI.escape(id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}
      query_params[:'data_source_id'] = data_source_id
      query_params[:'include_remote_data'] = opts[:'include_remote_data'] if !opts[:'include_remote_data'].nil?
      query_params[:'includes'] = @api_client.build_collection_param(opts[:'includes'], :multi) if !opts[:'includes'].nil?

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'Opportunity'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['accessToken', 'apiKey', 'customerId']

      new_options = opts.merge(
        :operation => :"CRMApi.get_opportunity",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: CRMApi#get_opportunity\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end
  end

  # top-level client access to avoid having the user to insantiate their own API instances
  CRM = CRMApi::new
end

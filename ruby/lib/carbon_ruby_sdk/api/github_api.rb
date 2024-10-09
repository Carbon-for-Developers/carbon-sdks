=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'cgi'

module Carbon
  class GithubApi
    attr_accessor :api_client

    def initialize(api_client = ApiClient.default)
      @api_client = api_client
    end

    # Issue
    #
    # @param issue_number [Integer] 
    # @param include_remote_data [Boolean] 
    # @param data_source_id [Integer] 
    # @param repository [String] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_issue(issue_number:, include_remote_data: false, data_source_id: SENTINEL, repository: SENTINEL, extra: {})
      extra[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      extra[:data_source_id] = data_source_id if data_source_id != SENTINEL
      extra[:repository] = repository if repository != SENTINEL
      api_response = get_issue_with_http_info_impl(issue_number, extra)
      api_response.data
    end

    # Issue
    #
    # @param issue_number [Integer] 
    # @param include_remote_data [Boolean] 
    # @param data_source_id [Integer] 
    # @param repository [String] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_issue_with_http_info(issue_number:, include_remote_data: false, data_source_id: SENTINEL, repository: SENTINEL, extra: {})
      extra[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      extra[:data_source_id] = data_source_id if data_source_id != SENTINEL
      extra[:repository] = repository if repository != SENTINEL
      get_issue_with_http_info_impl(issue_number, extra)
    end

    # Issue
    # @param issue_number [Integer] 
    # @param [Hash] opts the optional parameters
    # @option opts [Boolean] :include_remote_data  (default to false)
    # @option opts [Integer] :data_source_id 
    # @option opts [String] :repository 
    # @return [Issue]
    private def get_issue_impl(issue_number, opts = {})
      data, _status_code, _headers = get_issue_with_http_info(issue_number, opts)
      data
    end

    # Issue
    # @param issue_number [Integer] 
    # @param [Hash] opts the optional parameters
    # @option opts [Boolean] :include_remote_data  (default to false)
    # @option opts [Integer] :data_source_id 
    # @option opts [String] :repository 
    # @return [APIResponse] data is Issue, status code, headers and response
    private def get_issue_with_http_info_impl(issue_number, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: GithubApi.get_issue ...'
      end
      # verify the required parameter 'issue_number' is set
      if @api_client.config.client_side_validation && issue_number.nil?
        fail ArgumentError, "Missing the required parameter 'issue_number' when calling GithubApi.get_issue"
      end
      # resource path
      local_var_path = '/integrations/data/github/issues/{issue_number}'.sub('{' + 'issue_number' + '}', CGI.escape(issue_number.to_s))

      # query parameters
      query_params = opts[:query_params] || {}
      query_params[:'include_remote_data'] = opts[:'include_remote_data'] if !opts[:'include_remote_data'].nil?
      query_params[:'data_source_id'] = opts[:'data_source_id'] if !opts[:'data_source_id'].nil?
      query_params[:'repository'] = opts[:'repository'] if !opts[:'repository'].nil?

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'Issue'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['accessToken', 'apiKey', 'customerId']

      new_options = opts.merge(
        :operation => :"GithubApi.get_issue",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: GithubApi#get_issue\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end


    # Issues
    #
    # @param data_source_id [Integer] 
    # @param repository [String] Full name of the repository, denoted as {owner}/{repo}
    # @param include_remote_data [Boolean] 
    # @param page [Integer] 
    # @param page_size [Integer] 
    # @param next_cursor [String] 
    # @param filters [IssuesFilter] 
    # @param order_by [IssuesOrderBy] 
    # @param order_dir [OrderDirV2Nullable] 
    # @param body [IssuesInput] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_issues(data_source_id:, repository:, include_remote_data: false, page: 1, page_size: 30, next_cursor: SENTINEL, filters: SENTINEL, order_by: 'created', order_dir: 'asc', extra: {})
      _body = {}
      _body[:data_source_id] = data_source_id if data_source_id != SENTINEL
      _body[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      _body[:repository] = repository if repository != SENTINEL
      _body[:page] = page if page != SENTINEL
      _body[:page_size] = page_size if page_size != SENTINEL
      _body[:next_cursor] = next_cursor if next_cursor != SENTINEL
      _body[:filters] = filters if filters != SENTINEL
      _body[:order_by] = order_by if order_by != SENTINEL
      _body[:order_dir] = order_dir if order_dir != SENTINEL
      issues_input = _body
      api_response = get_issues_with_http_info_impl(issues_input, extra)
      api_response.data
    end

    # Issues
    #
    # @param data_source_id [Integer] 
    # @param repository [String] Full name of the repository, denoted as {owner}/{repo}
    # @param include_remote_data [Boolean] 
    # @param page [Integer] 
    # @param page_size [Integer] 
    # @param next_cursor [String] 
    # @param filters [IssuesFilter] 
    # @param order_by [IssuesOrderBy] 
    # @param order_dir [OrderDirV2Nullable] 
    # @param body [IssuesInput] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_issues_with_http_info(data_source_id:, repository:, include_remote_data: false, page: 1, page_size: 30, next_cursor: SENTINEL, filters: SENTINEL, order_by: 'created', order_dir: 'asc', extra: {})
      _body = {}
      _body[:data_source_id] = data_source_id if data_source_id != SENTINEL
      _body[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      _body[:repository] = repository if repository != SENTINEL
      _body[:page] = page if page != SENTINEL
      _body[:page_size] = page_size if page_size != SENTINEL
      _body[:next_cursor] = next_cursor if next_cursor != SENTINEL
      _body[:filters] = filters if filters != SENTINEL
      _body[:order_by] = order_by if order_by != SENTINEL
      _body[:order_dir] = order_dir if order_dir != SENTINEL
      issues_input = _body
      get_issues_with_http_info_impl(issues_input, extra)
    end

    # Issues
    # @param issues_input [IssuesInput] 
    # @param [Hash] opts the optional parameters
    # @return [IssuesResponse]
    private def get_issues_impl(issues_input, opts = {})
      data, _status_code, _headers = get_issues_with_http_info(issues_input, opts)
      data
    end

    # Issues
    # @param issues_input [IssuesInput] 
    # @param [Hash] opts the optional parameters
    # @return [APIResponse] data is IssuesResponse, status code, headers and response
    private def get_issues_with_http_info_impl(issues_input, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: GithubApi.get_issues ...'
      end
      # verify the required parameter 'issues_input' is set
      if @api_client.config.client_side_validation && issues_input.nil?
        fail ArgumentError, "Missing the required parameter 'issues_input' when calling GithubApi.get_issues"
      end
      # resource path
      local_var_path = '/integrations/data/github/issues'

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
      post_body = opts[:debug_body] || @api_client.object_to_http_body(issues_input)

      # return_type
      return_type = opts[:debug_return_type] || 'IssuesResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['accessToken', 'apiKey', 'customerId']

      new_options = opts.merge(
        :operation => :"GithubApi.get_issues",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: GithubApi#get_issues\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end


    # Get Pr
    #
    # @param pull_number [Integer] 
    # @param include_remote_data [Boolean] 
    # @param data_source_id [Integer] 
    # @param repository [String] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_pr(pull_number:, include_remote_data: false, data_source_id: SENTINEL, repository: SENTINEL, extra: {})
      extra[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      extra[:data_source_id] = data_source_id if data_source_id != SENTINEL
      extra[:repository] = repository if repository != SENTINEL
      api_response = get_pr_with_http_info_impl(pull_number, extra)
      api_response.data
    end

    # Get Pr
    #
    # @param pull_number [Integer] 
    # @param include_remote_data [Boolean] 
    # @param data_source_id [Integer] 
    # @param repository [String] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_pr_with_http_info(pull_number:, include_remote_data: false, data_source_id: SENTINEL, repository: SENTINEL, extra: {})
      extra[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      extra[:data_source_id] = data_source_id if data_source_id != SENTINEL
      extra[:repository] = repository if repository != SENTINEL
      get_pr_with_http_info_impl(pull_number, extra)
    end

    # Get Pr
    # @param pull_number [Integer] 
    # @param [Hash] opts the optional parameters
    # @option opts [Boolean] :include_remote_data  (default to false)
    # @option opts [Integer] :data_source_id 
    # @option opts [String] :repository 
    # @return [PullRequestExtended]
    private def get_pr_impl(pull_number, opts = {})
      data, _status_code, _headers = get_pr_with_http_info(pull_number, opts)
      data
    end

    # Get Pr
    # @param pull_number [Integer] 
    # @param [Hash] opts the optional parameters
    # @option opts [Boolean] :include_remote_data  (default to false)
    # @option opts [Integer] :data_source_id 
    # @option opts [String] :repository 
    # @return [APIResponse] data is PullRequestExtended, status code, headers and response
    private def get_pr_with_http_info_impl(pull_number, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: GithubApi.get_pr ...'
      end
      # verify the required parameter 'pull_number' is set
      if @api_client.config.client_side_validation && pull_number.nil?
        fail ArgumentError, "Missing the required parameter 'pull_number' when calling GithubApi.get_pr"
      end
      # resource path
      local_var_path = '/integrations/data/github/pull_requests/{pull_number}'.sub('{' + 'pull_number' + '}', CGI.escape(pull_number.to_s))

      # query parameters
      query_params = opts[:query_params] || {}
      query_params[:'include_remote_data'] = opts[:'include_remote_data'] if !opts[:'include_remote_data'].nil?
      query_params[:'data_source_id'] = opts[:'data_source_id'] if !opts[:'data_source_id'].nil?
      query_params[:'repository'] = opts[:'repository'] if !opts[:'repository'].nil?

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'PullRequestExtended'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['accessToken', 'apiKey', 'customerId']

      new_options = opts.merge(
        :operation => :"GithubApi.get_pr",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: GithubApi#get_pr\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end


    # Pr Comments
    #
    # @param data_source_id [Integer] 
    # @param repository [String] Full name of the repository, denoted as {owner}/{repo}
    # @param pull_number [Integer] 
    # @param include_remote_data [Boolean] 
    # @param page [Integer] 
    # @param page_size [Integer] 
    # @param next_cursor [String] 
    # @param order_by [CommentsOrderBy] 
    # @param order_dir [OrderDirV2Nullable] 
    # @param body [CommentsInput] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_pr_comments(data_source_id:, repository:, pull_number:, include_remote_data: false, page: 1, page_size: 30, next_cursor: SENTINEL, order_by: 'created', order_dir: 'asc', extra: {})
      _body = {}
      _body[:data_source_id] = data_source_id if data_source_id != SENTINEL
      _body[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      _body[:repository] = repository if repository != SENTINEL
      _body[:page] = page if page != SENTINEL
      _body[:page_size] = page_size if page_size != SENTINEL
      _body[:next_cursor] = next_cursor if next_cursor != SENTINEL
      _body[:pull_number] = pull_number if pull_number != SENTINEL
      _body[:order_by] = order_by if order_by != SENTINEL
      _body[:order_dir] = order_dir if order_dir != SENTINEL
      comments_input = _body
      api_response = get_pr_comments_with_http_info_impl(comments_input, extra)
      api_response.data
    end

    # Pr Comments
    #
    # @param data_source_id [Integer] 
    # @param repository [String] Full name of the repository, denoted as {owner}/{repo}
    # @param pull_number [Integer] 
    # @param include_remote_data [Boolean] 
    # @param page [Integer] 
    # @param page_size [Integer] 
    # @param next_cursor [String] 
    # @param order_by [CommentsOrderBy] 
    # @param order_dir [OrderDirV2Nullable] 
    # @param body [CommentsInput] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_pr_comments_with_http_info(data_source_id:, repository:, pull_number:, include_remote_data: false, page: 1, page_size: 30, next_cursor: SENTINEL, order_by: 'created', order_dir: 'asc', extra: {})
      _body = {}
      _body[:data_source_id] = data_source_id if data_source_id != SENTINEL
      _body[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      _body[:repository] = repository if repository != SENTINEL
      _body[:page] = page if page != SENTINEL
      _body[:page_size] = page_size if page_size != SENTINEL
      _body[:next_cursor] = next_cursor if next_cursor != SENTINEL
      _body[:pull_number] = pull_number if pull_number != SENTINEL
      _body[:order_by] = order_by if order_by != SENTINEL
      _body[:order_dir] = order_dir if order_dir != SENTINEL
      comments_input = _body
      get_pr_comments_with_http_info_impl(comments_input, extra)
    end

    # Pr Comments
    # @param comments_input [CommentsInput] 
    # @param [Hash] opts the optional parameters
    # @return [CommentsResponse]
    private def get_pr_comments_impl(comments_input, opts = {})
      data, _status_code, _headers = get_pr_comments_with_http_info(comments_input, opts)
      data
    end

    # Pr Comments
    # @param comments_input [CommentsInput] 
    # @param [Hash] opts the optional parameters
    # @return [APIResponse] data is CommentsResponse, status code, headers and response
    private def get_pr_comments_with_http_info_impl(comments_input, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: GithubApi.get_pr_comments ...'
      end
      # verify the required parameter 'comments_input' is set
      if @api_client.config.client_side_validation && comments_input.nil?
        fail ArgumentError, "Missing the required parameter 'comments_input' when calling GithubApi.get_pr_comments"
      end
      # resource path
      local_var_path = '/integrations/data/github/pull_requests/comments'

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
      post_body = opts[:debug_body] || @api_client.object_to_http_body(comments_input)

      # return_type
      return_type = opts[:debug_return_type] || 'CommentsResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['accessToken', 'apiKey', 'customerId']

      new_options = opts.merge(
        :operation => :"GithubApi.get_pr_comments",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: GithubApi#get_pr_comments\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end


    # Pr Commits
    #
    # @param data_source_id [Integer] 
    # @param repository [String] Full name of the repository, denoted as {owner}/{repo}
    # @param pull_number [Integer] 
    # @param include_remote_data [Boolean] 
    # @param page [Integer] 
    # @param page_size [Integer] 
    # @param next_cursor [String] 
    # @param body [CommitsInput] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_pr_commits(data_source_id:, repository:, pull_number:, include_remote_data: false, page: 1, page_size: 30, next_cursor: SENTINEL, extra: {})
      _body = {}
      _body[:data_source_id] = data_source_id if data_source_id != SENTINEL
      _body[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      _body[:repository] = repository if repository != SENTINEL
      _body[:page] = page if page != SENTINEL
      _body[:page_size] = page_size if page_size != SENTINEL
      _body[:next_cursor] = next_cursor if next_cursor != SENTINEL
      _body[:pull_number] = pull_number if pull_number != SENTINEL
      commits_input = _body
      api_response = get_pr_commits_with_http_info_impl(commits_input, extra)
      api_response.data
    end

    # Pr Commits
    #
    # @param data_source_id [Integer] 
    # @param repository [String] Full name of the repository, denoted as {owner}/{repo}
    # @param pull_number [Integer] 
    # @param include_remote_data [Boolean] 
    # @param page [Integer] 
    # @param page_size [Integer] 
    # @param next_cursor [String] 
    # @param body [CommitsInput] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_pr_commits_with_http_info(data_source_id:, repository:, pull_number:, include_remote_data: false, page: 1, page_size: 30, next_cursor: SENTINEL, extra: {})
      _body = {}
      _body[:data_source_id] = data_source_id if data_source_id != SENTINEL
      _body[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      _body[:repository] = repository if repository != SENTINEL
      _body[:page] = page if page != SENTINEL
      _body[:page_size] = page_size if page_size != SENTINEL
      _body[:next_cursor] = next_cursor if next_cursor != SENTINEL
      _body[:pull_number] = pull_number if pull_number != SENTINEL
      commits_input = _body
      get_pr_commits_with_http_info_impl(commits_input, extra)
    end

    # Pr Commits
    # @param commits_input [CommitsInput] 
    # @param [Hash] opts the optional parameters
    # @return [CommitsResponse]
    private def get_pr_commits_impl(commits_input, opts = {})
      data, _status_code, _headers = get_pr_commits_with_http_info(commits_input, opts)
      data
    end

    # Pr Commits
    # @param commits_input [CommitsInput] 
    # @param [Hash] opts the optional parameters
    # @return [APIResponse] data is CommitsResponse, status code, headers and response
    private def get_pr_commits_with_http_info_impl(commits_input, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: GithubApi.get_pr_commits ...'
      end
      # verify the required parameter 'commits_input' is set
      if @api_client.config.client_side_validation && commits_input.nil?
        fail ArgumentError, "Missing the required parameter 'commits_input' when calling GithubApi.get_pr_commits"
      end
      # resource path
      local_var_path = '/integrations/data/github/pull_requests/commits'

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
      post_body = opts[:debug_body] || @api_client.object_to_http_body(commits_input)

      # return_type
      return_type = opts[:debug_return_type] || 'CommitsResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['accessToken', 'apiKey', 'customerId']

      new_options = opts.merge(
        :operation => :"GithubApi.get_pr_commits",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: GithubApi#get_pr_commits\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end


    # Pr Files
    #
    # @param data_source_id [Integer] 
    # @param repository [String] Full name of the repository, denoted as {owner}/{repo}
    # @param pull_number [Integer] 
    # @param include_remote_data [Boolean] 
    # @param page [Integer] 
    # @param page_size [Integer] 
    # @param next_cursor [String] 
    # @param body [FilesInput] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_pr_files(data_source_id:, repository:, pull_number:, include_remote_data: false, page: 1, page_size: 30, next_cursor: SENTINEL, extra: {})
      _body = {}
      _body[:data_source_id] = data_source_id if data_source_id != SENTINEL
      _body[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      _body[:repository] = repository if repository != SENTINEL
      _body[:page] = page if page != SENTINEL
      _body[:page_size] = page_size if page_size != SENTINEL
      _body[:next_cursor] = next_cursor if next_cursor != SENTINEL
      _body[:pull_number] = pull_number if pull_number != SENTINEL
      files_input = _body
      api_response = get_pr_files_with_http_info_impl(files_input, extra)
      api_response.data
    end

    # Pr Files
    #
    # @param data_source_id [Integer] 
    # @param repository [String] Full name of the repository, denoted as {owner}/{repo}
    # @param pull_number [Integer] 
    # @param include_remote_data [Boolean] 
    # @param page [Integer] 
    # @param page_size [Integer] 
    # @param next_cursor [String] 
    # @param body [FilesInput] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_pr_files_with_http_info(data_source_id:, repository:, pull_number:, include_remote_data: false, page: 1, page_size: 30, next_cursor: SENTINEL, extra: {})
      _body = {}
      _body[:data_source_id] = data_source_id if data_source_id != SENTINEL
      _body[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      _body[:repository] = repository if repository != SENTINEL
      _body[:page] = page if page != SENTINEL
      _body[:page_size] = page_size if page_size != SENTINEL
      _body[:next_cursor] = next_cursor if next_cursor != SENTINEL
      _body[:pull_number] = pull_number if pull_number != SENTINEL
      files_input = _body
      get_pr_files_with_http_info_impl(files_input, extra)
    end

    # Pr Files
    # @param files_input [FilesInput] 
    # @param [Hash] opts the optional parameters
    # @return [FilesResponse]
    private def get_pr_files_impl(files_input, opts = {})
      data, _status_code, _headers = get_pr_files_with_http_info(files_input, opts)
      data
    end

    # Pr Files
    # @param files_input [FilesInput] 
    # @param [Hash] opts the optional parameters
    # @return [APIResponse] data is FilesResponse, status code, headers and response
    private def get_pr_files_with_http_info_impl(files_input, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: GithubApi.get_pr_files ...'
      end
      # verify the required parameter 'files_input' is set
      if @api_client.config.client_side_validation && files_input.nil?
        fail ArgumentError, "Missing the required parameter 'files_input' when calling GithubApi.get_pr_files"
      end
      # resource path
      local_var_path = '/integrations/data/github/pull_requests/files'

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
      post_body = opts[:debug_body] || @api_client.object_to_http_body(files_input)

      # return_type
      return_type = opts[:debug_return_type] || 'FilesResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['accessToken', 'apiKey', 'customerId']

      new_options = opts.merge(
        :operation => :"GithubApi.get_pr_files",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: GithubApi#get_pr_files\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end


    # Get Prs
    #
    # @param data_source_id [Integer] 
    # @param repository [String] Full name of the repository, denoted as {owner}/{repo}
    # @param include_remote_data [Boolean] 
    # @param page [Integer] 
    # @param page_size [Integer] 
    # @param next_cursor [String] 
    # @param filters [PullRequestFilters] 
    # @param order_by [PROrderBy] 
    # @param order_dir [OrderDirV2Nullable] 
    # @param body [PullRequestsInput] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_pull_requests(data_source_id:, repository:, include_remote_data: false, page: 1, page_size: 30, next_cursor: SENTINEL, filters: SENTINEL, order_by: 'created', order_dir: 'asc', extra: {})
      _body = {}
      _body[:data_source_id] = data_source_id if data_source_id != SENTINEL
      _body[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      _body[:repository] = repository if repository != SENTINEL
      _body[:page] = page if page != SENTINEL
      _body[:page_size] = page_size if page_size != SENTINEL
      _body[:next_cursor] = next_cursor if next_cursor != SENTINEL
      _body[:filters] = filters if filters != SENTINEL
      _body[:order_by] = order_by if order_by != SENTINEL
      _body[:order_dir] = order_dir if order_dir != SENTINEL
      pull_requests_input = _body
      api_response = get_pull_requests_with_http_info_impl(pull_requests_input, extra)
      api_response.data
    end

    # Get Prs
    #
    # @param data_source_id [Integer] 
    # @param repository [String] Full name of the repository, denoted as {owner}/{repo}
    # @param include_remote_data [Boolean] 
    # @param page [Integer] 
    # @param page_size [Integer] 
    # @param next_cursor [String] 
    # @param filters [PullRequestFilters] 
    # @param order_by [PROrderBy] 
    # @param order_dir [OrderDirV2Nullable] 
    # @param body [PullRequestsInput] 
    # @param [Hash] extra additional parameters to pass along through :header_params, :query_params, or parameter name
    def get_pull_requests_with_http_info(data_source_id:, repository:, include_remote_data: false, page: 1, page_size: 30, next_cursor: SENTINEL, filters: SENTINEL, order_by: 'created', order_dir: 'asc', extra: {})
      _body = {}
      _body[:data_source_id] = data_source_id if data_source_id != SENTINEL
      _body[:include_remote_data] = include_remote_data if include_remote_data != SENTINEL
      _body[:repository] = repository if repository != SENTINEL
      _body[:page] = page if page != SENTINEL
      _body[:page_size] = page_size if page_size != SENTINEL
      _body[:next_cursor] = next_cursor if next_cursor != SENTINEL
      _body[:filters] = filters if filters != SENTINEL
      _body[:order_by] = order_by if order_by != SENTINEL
      _body[:order_dir] = order_dir if order_dir != SENTINEL
      pull_requests_input = _body
      get_pull_requests_with_http_info_impl(pull_requests_input, extra)
    end

    # Get Prs
    # @param pull_requests_input [PullRequestsInput] 
    # @param [Hash] opts the optional parameters
    # @return [PullRequestResponse]
    private def get_pull_requests_impl(pull_requests_input, opts = {})
      data, _status_code, _headers = get_pull_requests_with_http_info(pull_requests_input, opts)
      data
    end

    # Get Prs
    # @param pull_requests_input [PullRequestsInput] 
    # @param [Hash] opts the optional parameters
    # @return [APIResponse] data is PullRequestResponse, status code, headers and response
    private def get_pull_requests_with_http_info_impl(pull_requests_input, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: GithubApi.get_pull_requests ...'
      end
      # verify the required parameter 'pull_requests_input' is set
      if @api_client.config.client_side_validation && pull_requests_input.nil?
        fail ArgumentError, "Missing the required parameter 'pull_requests_input' when calling GithubApi.get_pull_requests"
      end
      # resource path
      local_var_path = '/integrations/data/github/pull_requests'

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
      post_body = opts[:debug_body] || @api_client.object_to_http_body(pull_requests_input)

      # return_type
      return_type = opts[:debug_return_type] || 'PullRequestResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['accessToken', 'apiKey', 'customerId']

      new_options = opts.merge(
        :operation => :"GithubApi.get_pull_requests",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers, response = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: GithubApi#get_pull_requests\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      APIResponse::new(data, status_code, headers, response)
    end
  end

  # top-level client access to avoid having the user to insantiate their own API instances
  Github = GithubApi::new
end

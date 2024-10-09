=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'

# Unit tests for Carbon::GithubApi
describe 'GithubApi' do
  before do
    # run before each test
    @api_instance = Carbon::GithubApi.new
  end

  after do
    # run after each test
  end

  describe 'test an instance of GithubApi' do
    it 'should create an instance of GithubApi' do
      expect(@api_instance).to be_instance_of(Carbon::GithubApi)
    end
  end

  # unit tests for get_issue
  # Issue
  # @param issue_number 
  # @param [Hash] opts the optional parameters
  # @option opts [Boolean] :include_remote_data 
  # @option opts [Integer] :data_source_id 
  # @option opts [String] :repository 
  # @return [Issue]
  describe 'get_issue test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_issues
  # Issues
  # @param issues_input 
  # @param [Hash] opts the optional parameters
  # @return [IssuesResponse]
  describe 'get_issues test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_pr
  # Get Pr
  # @param pull_number 
  # @param [Hash] opts the optional parameters
  # @option opts [Boolean] :include_remote_data 
  # @option opts [Integer] :data_source_id 
  # @option opts [String] :repository 
  # @return [PullRequestExtended]
  describe 'get_pr test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_pr_comments
  # Pr Comments
  # @param comments_input 
  # @param [Hash] opts the optional parameters
  # @return [CommentsResponse]
  describe 'get_pr_comments test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_pr_commits
  # Pr Commits
  # @param commits_input 
  # @param [Hash] opts the optional parameters
  # @return [CommitsResponse]
  describe 'get_pr_commits test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_pr_files
  # Pr Files
  # @param files_input 
  # @param [Hash] opts the optional parameters
  # @return [FilesResponse]
  describe 'get_pr_files test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_pull_requests
  # Get Prs
  # @param pull_requests_input 
  # @param [Hash] opts the optional parameters
  # @return [PullRequestResponse]
  describe 'get_pull_requests test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

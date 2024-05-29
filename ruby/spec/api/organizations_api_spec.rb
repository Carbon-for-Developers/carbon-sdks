=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'

# Unit tests for Carbon::OrganizationsApi
describe 'OrganizationsApi' do
  before do
    # run before each test
    @api_instance = Carbon::OrganizationsApi.new
  end

  after do
    # run after each test
  end

  describe 'test an instance of OrganizationsApi' do
    it 'should create an instance of OrganizationsApi' do
      expect(@api_instance).to be_instance_of(Carbon::OrganizationsApi)
    end
  end

  # unit tests for get
  # Get Organization
  # @param [Hash] opts the optional parameters
  # @return [OrganizationResponse]
  describe 'get test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for update
  # Update Organization
  # @param update_organization_input 
  # @param [Hash] opts the optional parameters
  # @return [GenericSuccessResponse]
  describe 'update test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for update_stats
  # Update Organization Statistics
  # Use this endpoint to reaggregate the statistics for an organization, for example aggregate_file_size. The reaggregation process is asyncronous so a webhook will be sent with the event type being FILE_STATISTICS_AGGREGATED to notify when the process is complee. After this aggregation is complete, the updated statistics can be retrieved using the /organization endpoint. The response of /organization willalso contain a timestamp of the last time the statistics were reaggregated.
  # @param [Hash] opts the optional parameters
  # @return [GenericSuccessResponse]
  describe 'update_stats test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

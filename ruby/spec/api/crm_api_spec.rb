=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'

# Unit tests for Carbon::CRMApi
describe 'CRMApi' do
  before do
    # run before each test
    @api_instance = Carbon::CRMApi.new
  end

  after do
    # run after each test
  end

  describe 'test an instance of CRMApi' do
    it 'should create an instance of CRMApi' do
      expect(@api_instance).to be_instance_of(Carbon::CRMApi)
    end
  end

  # unit tests for get_account
  # Get Account
  # @param id 
  # @param data_source_id 
  # @param [Hash] opts the optional parameters
  # @option opts [Boolean] :include_remote_data 
  # @option opts [Array<BaseIncludes>] :includes 
  # @return [Account]
  describe 'get_account test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_accounts
  # Get Accounts
  # @param accounts_request 
  # @param [Hash] opts the optional parameters
  # @return [AccountResponse]
  describe 'get_accounts test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_contact
  # Get Contact
  # @param id 
  # @param data_source_id 
  # @param [Hash] opts the optional parameters
  # @option opts [Boolean] :include_remote_data 
  # @option opts [Array<BaseIncludes>] :includes 
  # @return [Contact]
  describe 'get_contact test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_contacts
  # Get Contacts
  # @param contacts_request 
  # @param [Hash] opts the optional parameters
  # @return [ContactsResponse]
  describe 'get_contacts test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_lead
  # Get Lead
  # @param id 
  # @param data_source_id 
  # @param [Hash] opts the optional parameters
  # @option opts [Boolean] :include_remote_data 
  # @option opts [Array<BaseIncludes>] :includes 
  # @return [Lead]
  describe 'get_lead test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_leads
  # Get Leads
  # @param leads_request 
  # @param [Hash] opts the optional parameters
  # @return [LeadsResponse]
  describe 'get_leads test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_opportunities
  # Get Opportunities
  # @param opportunities_request 
  # @param [Hash] opts the optional parameters
  # @return [OpportunitiesResponse]
  describe 'get_opportunities test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_opportunity
  # Get Opportunity
  # @param id 
  # @param data_source_id 
  # @param [Hash] opts the optional parameters
  # @option opts [Boolean] :include_remote_data 
  # @option opts [Array<BaseIncludes>] :includes 
  # @return [Opportunity]
  describe 'get_opportunity test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

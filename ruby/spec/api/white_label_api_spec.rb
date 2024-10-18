=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'

# Unit tests for Carbon::WhiteLabelApi
describe 'WhiteLabelApi' do
  before do
    # run before each test
    @api_instance = Carbon::WhiteLabelApi.new
  end

  after do
    # run after each test
  end

  describe 'test an instance of WhiteLabelApi' do
    it 'should create an instance of WhiteLabelApi' do
      expect(@api_instance).to be_instance_of(Carbon::WhiteLabelApi)
    end
  end

  # unit tests for create
  # Create White Labels
  # @param white_label_create_request_inner 
  # @param [Hash] opts the optional parameters
  # @return [Object]
  describe 'create test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for delete
  # Delete White Labels
  # @param delete_white_label_request 
  # @param [Hash] opts the optional parameters
  # @return [Object]
  describe 'delete test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for list
  # List White Labels
  # @param list_white_label_request 
  # @param [Hash] opts the optional parameters
  # @return [Object]
  describe 'list test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for update
  # Update White Label
  # @param white_label_update_request 
  # @param [Hash] opts the optional parameters
  # @return [Object]
  describe 'update test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

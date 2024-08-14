=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::ModifyColdStorageParametersQueryInput
describe Carbon::ModifyColdStorageParametersQueryInput do
  let(:instance) { Carbon::ModifyColdStorageParametersQueryInput.new }

  describe 'test an instance of ModifyColdStorageParametersQueryInput' do
    it 'should create an instance of ModifyColdStorageParametersQueryInput' do
      expect(instance).to be_instance_of(Carbon::ModifyColdStorageParametersQueryInput)
    end
  end
  describe 'test attribute "filters"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "enable_cold_storage"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "hot_storage_time_to_live"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

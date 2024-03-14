=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::ConnectDataSourceResponse
describe Carbon::ConnectDataSourceResponse do
  let(:instance) { Carbon::ConnectDataSourceResponse.new }

  describe 'test an instance of ConnectDataSourceResponse' do
    it 'should create an instance of ConnectDataSourceResponse' do
      expect(instance).to be_instance_of(Carbon::ConnectDataSourceResponse)
    end
  end
  describe 'test attribute "data_source"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "sync_url"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

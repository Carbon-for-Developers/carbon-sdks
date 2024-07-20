=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::DataSourceConfiguration
describe Carbon::DataSourceConfiguration do
  let(:instance) { Carbon::DataSourceConfiguration.new }

  describe 'test an instance of DataSourceConfiguration' do
    it 'should create an instance of DataSourceConfiguration' do
      expect(instance).to be_instance_of(Carbon::DataSourceConfiguration)
    end
  end
  describe 'test attribute "allowed_file_formats"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "google_workspace_docs_save_as"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::UserConfigurationNullable
describe Carbon::UserConfigurationNullable do
  let(:instance) { Carbon::UserConfigurationNullable.new }

  describe 'test an instance of UserConfigurationNullable' do
    it 'should create an instance of UserConfigurationNullable' do
      expect(instance).to be_instance_of(Carbon::UserConfigurationNullable)
    end
  end
  describe 'test attribute "auto_sync_enabled_sources"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "max_files"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "max_files_per_upload"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

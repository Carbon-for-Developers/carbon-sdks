=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::HelpdeskFileSyncConfig
describe Carbon::HelpdeskFileSyncConfig do
  let(:instance) { Carbon::HelpdeskFileSyncConfig.new }

  describe 'test an instance of HelpdeskFileSyncConfig' do
    it 'should create an instance of HelpdeskFileSyncConfig' do
      expect(instance).to be_instance_of(Carbon::HelpdeskFileSyncConfig)
    end
  end
  describe 'test attribute "file_types"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "sync_attachments"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

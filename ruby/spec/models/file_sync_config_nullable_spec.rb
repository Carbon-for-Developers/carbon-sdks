=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::FileSyncConfigNullable
describe Carbon::FileSyncConfigNullable do
  let(:instance) { Carbon::FileSyncConfigNullable.new }

  describe 'test an instance of FileSyncConfigNullable' do
    it 'should create an instance of FileSyncConfigNullable' do
      expect(instance).to be_instance_of(Carbon::FileSyncConfigNullable)
    end
  end
  describe 'test attribute "auto_synced_source_types"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "sync_attachments"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "detect_audio_language"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

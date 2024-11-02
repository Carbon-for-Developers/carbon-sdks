=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::StorageFileFormats
describe Carbon::StorageFileFormats do
  let(:instance) { Carbon::StorageFileFormats.new }

  describe 'test an instance of StorageFileFormats' do
    it 'should create an instance of StorageFileFormats' do
      expect(instance).to be_instance_of(Carbon::StorageFileFormats)
    end
  end
end

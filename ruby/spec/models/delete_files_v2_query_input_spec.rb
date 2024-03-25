=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::DeleteFilesV2QueryInput
describe Carbon::DeleteFilesV2QueryInput do
  let(:instance) { Carbon::DeleteFilesV2QueryInput.new }

  describe 'test an instance of DeleteFilesV2QueryInput' do
    it 'should create an instance of DeleteFilesV2QueryInput' do
      expect(instance).to be_instance_of(Carbon::DeleteFilesV2QueryInput)
    end
  end
  describe 'test attribute "filters"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "send_webhook"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

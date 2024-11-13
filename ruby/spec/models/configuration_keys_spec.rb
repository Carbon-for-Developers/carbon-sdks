=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::ConfigurationKeys
describe Carbon::ConfigurationKeys do
  let(:instance) { Carbon::ConfigurationKeys.new }

  describe 'test an instance of ConfigurationKeys' do
    it 'should create an instance of ConfigurationKeys' do
      expect(instance).to be_instance_of(Carbon::ConfigurationKeys)
    end
  end
end

=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::ZendeskFileTypes
describe Carbon::ZendeskFileTypes do
  let(:instance) { Carbon::ZendeskFileTypes.new }

  describe 'test an instance of ZendeskFileTypes' do
    it 'should create an instance of ZendeskFileTypes' do
      expect(instance).to be_instance_of(Carbon::ZendeskFileTypes)
    end
  end
end

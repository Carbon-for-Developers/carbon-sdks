=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::ServiceNowFileTypes
describe Carbon::ServiceNowFileTypes do
  let(:instance) { Carbon::ServiceNowFileTypes.new }

  describe 'test an instance of ServiceNowFileTypes' do
    it 'should create an instance of ServiceNowFileTypes' do
      expect(instance).to be_instance_of(Carbon::ServiceNowFileTypes)
    end
  end
end

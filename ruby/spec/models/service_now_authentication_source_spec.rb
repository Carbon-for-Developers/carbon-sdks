=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::ServiceNowAuthenticationSource
describe Carbon::ServiceNowAuthenticationSource do
  let(:instance) { Carbon::ServiceNowAuthenticationSource.new }

  describe 'test an instance of ServiceNowAuthenticationSource' do
    it 'should create an instance of ServiceNowAuthenticationSource' do
      expect(instance).to be_instance_of(Carbon::ServiceNowAuthenticationSource)
    end
  end
end

=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::ZendeskAuthenticationSource
describe Carbon::ZendeskAuthenticationSource do
  let(:instance) { Carbon::ZendeskAuthenticationSource.new }

  describe 'test an instance of ZendeskAuthenticationSource' do
    it 'should create an instance of ZendeskAuthenticationSource' do
      expect(instance).to be_instance_of(Carbon::ZendeskAuthenticationSource)
    end
  end
end

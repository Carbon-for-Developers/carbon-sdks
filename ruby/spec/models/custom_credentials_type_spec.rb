=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::CustomCredentialsType
describe Carbon::CustomCredentialsType do
  let(:instance) { Carbon::CustomCredentialsType.new }

  describe 'test an instance of CustomCredentialsType' do
    it 'should create an instance of CustomCredentialsType' do
      expect(instance).to be_instance_of(Carbon::CustomCredentialsType)
    end
  end
end

=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::WhiteLabelUpdateRequest
describe Carbon::WhiteLabelUpdateRequest do
  let(:instance) { Carbon::WhiteLabelUpdateRequest.new }

  describe 'test an instance of WhiteLabelUpdateRequest' do
    it 'should create an instance of WhiteLabelUpdateRequest' do
      expect(instance).to be_instance_of(Carbon::WhiteLabelUpdateRequest)
    end
  end
  describe 'test attribute "data_source_type"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "credentials"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

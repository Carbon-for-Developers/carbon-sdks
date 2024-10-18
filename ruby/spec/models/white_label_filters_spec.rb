=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::WhiteLabelFilters
describe Carbon::WhiteLabelFilters do
  let(:instance) { Carbon::WhiteLabelFilters.new }

  describe 'test an instance of WhiteLabelFilters' do
    it 'should create an instance of WhiteLabelFilters' do
      expect(instance).to be_instance_of(Carbon::WhiteLabelFilters)
    end
  end
  describe 'test attribute "ids"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "data_source_type"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

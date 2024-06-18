=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::SlackFilters
describe Carbon::SlackFilters do
  let(:instance) { Carbon::SlackFilters.new }

  describe 'test an instance of SlackFilters' do
    it 'should create an instance of SlackFilters' do
      expect(instance).to be_instance_of(Carbon::SlackFilters)
    end
  end
  describe 'test attribute "conversation_id"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "after"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

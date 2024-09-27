=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::AccountFilters
describe Carbon::AccountFilters do
  let(:instance) { Carbon::AccountFilters.new }

  describe 'test an instance of AccountFilters' do
    it 'should create an instance of AccountFilters' do
      expect(instance).to be_instance_of(Carbon::AccountFilters)
    end
  end
  describe 'test attribute "owner_id"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "name"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

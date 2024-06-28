=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::ListUsersFilters
describe Carbon::ListUsersFilters do
  let(:instance) { Carbon::ListUsersFilters.new }

  describe 'test an instance of ListUsersFilters' do
    it 'should create an instance of ListUsersFilters' do
      expect(instance).to be_instance_of(Carbon::ListUsersFilters)
    end
  end
  describe 'test attribute "customer_ids"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "carbon_user_ids"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

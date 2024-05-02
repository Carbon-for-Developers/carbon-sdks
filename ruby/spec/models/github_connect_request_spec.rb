=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::GithubConnectRequest
describe Carbon::GithubConnectRequest do
  let(:instance) { Carbon::GithubConnectRequest.new }

  describe 'test an instance of GithubConnectRequest' do
    it 'should create an instance of GithubConnectRequest' do
      expect(instance).to be_instance_of(Carbon::GithubConnectRequest)
    end
  end
  describe 'test attribute "username"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "access_token"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "sync_source_items"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::OuthURLResponse
describe Carbon::OuthURLResponse do
  let(:instance) { Carbon::OuthURLResponse.new }

  describe 'test an instance of OuthURLResponse' do
    it 'should create an instance of OuthURLResponse' do
      expect(instance).to be_instance_of(Carbon::OuthURLResponse)
    end
  end
  describe 'test attribute "oauth_url"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

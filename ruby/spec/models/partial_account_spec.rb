=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::PartialAccount
describe Carbon::PartialAccount do
  let(:instance) { Carbon::PartialAccount.new }

  describe 'test an instance of PartialAccount' do
    it 'should create an instance of PartialAccount' do
      expect(instance).to be_instance_of(Carbon::PartialAccount)
    end
  end
  describe 'test attribute "id"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

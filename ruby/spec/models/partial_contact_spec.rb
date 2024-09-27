=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::PartialContact
describe Carbon::PartialContact do
  let(:instance) { Carbon::PartialContact.new }

  describe 'test an instance of PartialContact' do
    it 'should create an instance of PartialContact' do
      expect(instance).to be_instance_of(Carbon::PartialContact)
    end
  end
  describe 'test attribute "id"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

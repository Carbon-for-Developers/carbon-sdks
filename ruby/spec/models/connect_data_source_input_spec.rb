=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::ConnectDataSourceInput
describe Carbon::ConnectDataSourceInput do
  let(:instance) { Carbon::ConnectDataSourceInput.new }

  describe 'test an instance of ConnectDataSourceInput' do
    it 'should create an instance of ConnectDataSourceInput' do
      expect(instance).to be_instance_of(Carbon::ConnectDataSourceInput)
    end
  end
  describe 'test attribute "authentication"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "sync_options"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

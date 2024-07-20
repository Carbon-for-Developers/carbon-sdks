=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::UpdateOrganizationInput
describe Carbon::UpdateOrganizationInput do
  let(:instance) { Carbon::UpdateOrganizationInput.new }

  describe 'test an instance of UpdateOrganizationInput' do
    it 'should create an instance of UpdateOrganizationInput' do
      expect(instance).to be_instance_of(Carbon::UpdateOrganizationInput)
    end
  end
  describe 'test attribute "global_user_config"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "data_source_configs"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

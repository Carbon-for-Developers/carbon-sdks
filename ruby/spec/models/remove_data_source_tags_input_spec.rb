=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::RemoveDataSourceTagsInput
describe Carbon::RemoveDataSourceTagsInput do
  let(:instance) { Carbon::RemoveDataSourceTagsInput.new }

  describe 'test an instance of RemoveDataSourceTagsInput' do
    it 'should create an instance of RemoveDataSourceTagsInput' do
      expect(instance).to be_instance_of(Carbon::RemoveDataSourceTagsInput)
    end
  end
  describe 'test attribute "data_source_id"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "tags_to_remove"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "remove_all_tags"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::DocumentResponseList
describe Carbon::DocumentResponseList do
  let(:instance) { Carbon::DocumentResponseList.new }

  describe 'test an instance of DocumentResponseList' do
    it 'should create an instance of DocumentResponseList' do
      expect(instance).to be_instance_of(Carbon::DocumentResponseList)
    end
  end
  describe 'test attribute "documents"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

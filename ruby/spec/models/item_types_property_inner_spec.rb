=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::ItemTypesPropertyInner
describe Carbon::ItemTypesPropertyInner do
  let(:instance) { Carbon::ItemTypesPropertyInner.new }

  describe 'test an instance of ItemTypesPropertyInner' do
    it 'should create an instance of ItemTypesPropertyInner' do
      expect(instance).to be_instance_of(Carbon::ItemTypesPropertyInner)
    end
  end
end

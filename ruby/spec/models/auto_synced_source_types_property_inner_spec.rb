=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::AutoSyncedSourceTypesPropertyInner
describe Carbon::AutoSyncedSourceTypesPropertyInner do
  let(:instance) { Carbon::AutoSyncedSourceTypesPropertyInner.new }

  describe 'test an instance of AutoSyncedSourceTypesPropertyInner' do
    it 'should create an instance of AutoSyncedSourceTypesPropertyInner' do
      expect(instance).to be_instance_of(Carbon::AutoSyncedSourceTypesPropertyInner)
    end
  end
end

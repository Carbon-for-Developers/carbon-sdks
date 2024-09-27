=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::OrderDirV2Nullable
describe Carbon::OrderDirV2Nullable do
  let(:instance) { Carbon::OrderDirV2Nullable.new }

  describe 'test an instance of OrderDirV2Nullable' do
    it 'should create an instance of OrderDirV2Nullable' do
      expect(instance).to be_instance_of(Carbon::OrderDirV2Nullable)
    end
  end
end

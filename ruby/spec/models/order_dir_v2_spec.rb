=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::OrderDirV2
describe Carbon::OrderDirV2 do
  let(:instance) { Carbon::OrderDirV2.new }

  describe 'test an instance of OrderDirV2' do
    it 'should create an instance of OrderDirV2' do
      expect(instance).to be_instance_of(Carbon::OrderDirV2)
    end
  end
end

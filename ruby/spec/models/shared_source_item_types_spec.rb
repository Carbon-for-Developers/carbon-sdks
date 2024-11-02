=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::SharedSourceItemTypes
describe Carbon::SharedSourceItemTypes do
  let(:instance) { Carbon::SharedSourceItemTypes.new }

  describe 'test an instance of SharedSourceItemTypes' do
    it 'should create an instance of SharedSourceItemTypes' do
      expect(instance).to be_instance_of(Carbon::SharedSourceItemTypes)
    end
  end
end

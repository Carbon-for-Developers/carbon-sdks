=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::WhiteLabelOrderByColumns
describe Carbon::WhiteLabelOrderByColumns do
  let(:instance) { Carbon::WhiteLabelOrderByColumns.new }

  describe 'test an instance of WhiteLabelOrderByColumns' do
    it 'should create an instance of WhiteLabelOrderByColumns' do
      expect(instance).to be_instance_of(Carbon::WhiteLabelOrderByColumns)
    end
  end
end

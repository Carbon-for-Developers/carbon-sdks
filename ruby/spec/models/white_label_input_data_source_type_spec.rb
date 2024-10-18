=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::WhiteLabelInputDataSourceType
describe Carbon::WhiteLabelInputDataSourceType do
  let(:instance) { Carbon::WhiteLabelInputDataSourceType.new }

  describe 'test an instance of WhiteLabelInputDataSourceType' do
    it 'should create an instance of WhiteLabelInputDataSourceType' do
      expect(instance).to be_instance_of(Carbon::WhiteLabelInputDataSourceType)
    end
  end
end

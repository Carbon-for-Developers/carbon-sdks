=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::GoogleDriveWhiteLabelInputDataSourceType
describe Carbon::GoogleDriveWhiteLabelInputDataSourceType do
  let(:instance) { Carbon::GoogleDriveWhiteLabelInputDataSourceType.new }

  describe 'test an instance of GoogleDriveWhiteLabelInputDataSourceType' do
    it 'should create an instance of GoogleDriveWhiteLabelInputDataSourceType' do
      expect(instance).to be_instance_of(Carbon::GoogleDriveWhiteLabelInputDataSourceType)
    end
  end
end

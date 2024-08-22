=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::ExternalDataSourceType
describe Carbon::ExternalDataSourceType do
  let(:instance) { Carbon::ExternalDataSourceType.new }

  describe 'test an instance of ExternalDataSourceType' do
    it 'should create an instance of ExternalDataSourceType' do
      expect(instance).to be_instance_of(Carbon::ExternalDataSourceType)
    end
  end
end

=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::DataSourceExtendedInput
describe Carbon::DataSourceExtendedInput do
  let(:instance) { Carbon::DataSourceExtendedInput.new }

  describe 'test an instance of DataSourceExtendedInput' do
    it 'should create an instance of DataSourceExtendedInput' do
      expect(instance).to be_instance_of(Carbon::DataSourceExtendedInput)
    end
  end
end

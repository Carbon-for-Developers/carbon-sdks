=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::AutoSyncEnabledSourcesProperty
describe Carbon::AutoSyncEnabledSourcesProperty do
  let(:instance) { Carbon::AutoSyncEnabledSourcesProperty.new }

  describe 'test an instance of AutoSyncEnabledSourcesProperty' do
    it 'should create an instance of AutoSyncEnabledSourcesProperty' do
      expect(instance).to be_instance_of(Carbon::AutoSyncEnabledSourcesProperty)
    end
  end
end

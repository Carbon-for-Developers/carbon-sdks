=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::ExternalSourceItemsOrderBy
describe Carbon::ExternalSourceItemsOrderBy do
  let(:instance) { Carbon::ExternalSourceItemsOrderBy.new }

  describe 'test an instance of ExternalSourceItemsOrderBy' do
    it 'should create an instance of ExternalSourceItemsOrderBy' do
      expect(instance).to be_instance_of(Carbon::ExternalSourceItemsOrderBy)
    end
  end
end

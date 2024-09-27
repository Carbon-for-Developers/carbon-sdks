=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::OpportunityStatus
describe Carbon::OpportunityStatus do
  let(:instance) { Carbon::OpportunityStatus.new }

  describe 'test an instance of OpportunityStatus' do
    it 'should create an instance of OpportunityStatus' do
      expect(instance).to be_instance_of(Carbon::OpportunityStatus)
    end
  end
end

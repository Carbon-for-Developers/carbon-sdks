=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::LeadsOrderBy
describe Carbon::LeadsOrderBy do
  let(:instance) { Carbon::LeadsOrderBy.new }

  describe 'test an instance of LeadsOrderBy' do
    it 'should create an instance of LeadsOrderBy' do
      expect(instance).to be_instance_of(Carbon::LeadsOrderBy)
    end
  end
end

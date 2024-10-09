=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::PROrderBy
describe Carbon::PROrderBy do
  let(:instance) { Carbon::PROrderBy.new }

  describe 'test an instance of PROrderBy' do
    it 'should create an instance of PROrderBy' do
      expect(instance).to be_instance_of(Carbon::PROrderBy)
    end
  end
end

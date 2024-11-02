=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::GitbookFileTypes
describe Carbon::GitbookFileTypes do
  let(:instance) { Carbon::GitbookFileTypes.new }

  describe 'test an instance of GitbookFileTypes' do
    it 'should create an instance of GitbookFileTypes' do
      expect(instance).to be_instance_of(Carbon::GitbookFileTypes)
    end
  end
end

=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::HelpdeskFileTypes
describe Carbon::HelpdeskFileTypes do
  let(:instance) { Carbon::HelpdeskFileTypes.new }

  describe 'test an instance of HelpdeskFileTypes' do
    it 'should create an instance of HelpdeskFileTypes' do
      expect(instance).to be_instance_of(Carbon::HelpdeskFileTypes)
    end
  end
end

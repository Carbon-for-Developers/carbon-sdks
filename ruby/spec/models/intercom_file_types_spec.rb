=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::IntercomFileTypes
describe Carbon::IntercomFileTypes do
  let(:instance) { Carbon::IntercomFileTypes.new }

  describe 'test an instance of IntercomFileTypes' do
    it 'should create an instance of IntercomFileTypes' do
      expect(instance).to be_instance_of(Carbon::IntercomFileTypes)
    end
  end
end

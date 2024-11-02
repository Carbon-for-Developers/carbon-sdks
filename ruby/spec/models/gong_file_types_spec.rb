=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::GongFileTypes
describe Carbon::GongFileTypes do
  let(:instance) { Carbon::GongFileTypes.new }

  describe 'test an instance of GongFileTypes' do
    it 'should create an instance of GongFileTypes' do
      expect(instance).to be_instance_of(Carbon::GongFileTypes)
    end
  end
end

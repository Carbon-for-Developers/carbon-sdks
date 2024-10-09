=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::PRStateInput
describe Carbon::PRStateInput do
  let(:instance) { Carbon::PRStateInput.new }

  describe 'test an instance of PRStateInput' do
    it 'should create an instance of PRStateInput' do
      expect(instance).to be_instance_of(Carbon::PRStateInput)
    end
  end
end

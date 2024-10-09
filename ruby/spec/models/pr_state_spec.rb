=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::PRState
describe Carbon::PRState do
  let(:instance) { Carbon::PRState.new }

  describe 'test an instance of PRState' do
    it 'should create an instance of PRState' do
      expect(instance).to be_instance_of(Carbon::PRState)
    end
  end
end

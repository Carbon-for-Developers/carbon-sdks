=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::Source
describe Carbon::Source do
  let(:instance) { Carbon::Source.new }

  describe 'test an instance of Source' do
    it 'should create an instance of Source' do
      expect(instance).to be_instance_of(Carbon::Source)
    end
  end
end

=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::ParsedTextFormats
describe Carbon::ParsedTextFormats do
  let(:instance) { Carbon::ParsedTextFormats.new }

  describe 'test an instance of ParsedTextFormats' do
    it 'should create an instance of ParsedTextFormats' do
      expect(instance).to be_instance_of(Carbon::ParsedTextFormats)
    end
  end
end

=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::ParsedTextFormatsNullable
describe Carbon::ParsedTextFormatsNullable do
  let(:instance) { Carbon::ParsedTextFormatsNullable.new }

  describe 'test an instance of ParsedTextFormatsNullable' do
    it 'should create an instance of ParsedTextFormatsNullable' do
      expect(instance).to be_instance_of(Carbon::ParsedTextFormatsNullable)
    end
  end
end

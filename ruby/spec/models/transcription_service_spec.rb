=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::TranscriptionService
describe Carbon::TranscriptionService do
  let(:instance) { Carbon::TranscriptionService.new }

  describe 'test an instance of TranscriptionService' do
    it 'should create an instance of TranscriptionService' do
      expect(instance).to be_instance_of(Carbon::TranscriptionService)
    end
  end
end

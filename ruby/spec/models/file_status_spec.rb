=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::FileStatus
describe Carbon::FileStatus do
  let(:instance) { Carbon::FileStatus.new }

  describe 'test an instance of FileStatus' do
    it 'should create an instance of FileStatus' do
      expect(instance).to be_instance_of(Carbon::FileStatus)
    end
  end
end

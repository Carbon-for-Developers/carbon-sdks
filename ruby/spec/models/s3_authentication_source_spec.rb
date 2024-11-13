=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::S3AuthenticationSource
describe Carbon::S3AuthenticationSource do
  let(:instance) { Carbon::S3AuthenticationSource.new }

  describe 'test an instance of S3AuthenticationSource' do
    it 'should create an instance of S3AuthenticationSource' do
      expect(instance).to be_instance_of(Carbon::S3AuthenticationSource)
    end
  end
end

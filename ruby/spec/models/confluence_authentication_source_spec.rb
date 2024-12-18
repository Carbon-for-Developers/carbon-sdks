=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::ConfluenceAuthenticationSource
describe Carbon::ConfluenceAuthenticationSource do
  let(:instance) { Carbon::ConfluenceAuthenticationSource.new }

  describe 'test an instance of ConfluenceAuthenticationSource' do
    it 'should create an instance of ConfluenceAuthenticationSource' do
      expect(instance).to be_instance_of(Carbon::ConfluenceAuthenticationSource)
    end
  end
end

=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::ZoteroAuthenticationSource
describe Carbon::ZoteroAuthenticationSource do
  let(:instance) { Carbon::ZoteroAuthenticationSource.new }

  describe 'test an instance of ZoteroAuthenticationSource' do
    it 'should create an instance of ZoteroAuthenticationSource' do
      expect(instance).to be_instance_of(Carbon::ZoteroAuthenticationSource)
    end
  end
end

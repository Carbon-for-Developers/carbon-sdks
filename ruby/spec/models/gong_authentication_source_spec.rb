=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::GongAuthenticationSource
describe Carbon::GongAuthenticationSource do
  let(:instance) { Carbon::GongAuthenticationSource.new }

  describe 'test an instance of GongAuthenticationSource' do
    it 'should create an instance of GongAuthenticationSource' do
      expect(instance).to be_instance_of(Carbon::GongAuthenticationSource)
    end
  end
end

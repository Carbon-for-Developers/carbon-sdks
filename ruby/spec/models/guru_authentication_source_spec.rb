=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::GuruAuthenticationSource
describe Carbon::GuruAuthenticationSource do
  let(:instance) { Carbon::GuruAuthenticationSource.new }

  describe 'test an instance of GuruAuthenticationSource' do
    it 'should create an instance of GuruAuthenticationSource' do
      expect(instance).to be_instance_of(Carbon::GuruAuthenticationSource)
    end
  end
end

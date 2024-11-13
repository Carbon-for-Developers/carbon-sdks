=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::SharepointAuthenticationSource
describe Carbon::SharepointAuthenticationSource do
  let(:instance) { Carbon::SharepointAuthenticationSource.new }

  describe 'test an instance of SharepointAuthenticationSource' do
    it 'should create an instance of SharepointAuthenticationSource' do
      expect(instance).to be_instance_of(Carbon::SharepointAuthenticationSource)
    end
  end
end

=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::BaseIncludes
describe Carbon::BaseIncludes do
  let(:instance) { Carbon::BaseIncludes.new }

  describe 'test an instance of BaseIncludes' do
    it 'should create an instance of BaseIncludes' do
      expect(instance).to be_instance_of(Carbon::BaseIncludes)
    end
  end
end

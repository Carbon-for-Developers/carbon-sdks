=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::UserWebPageOrderByTypes
describe Carbon::UserWebPageOrderByTypes do
  let(:instance) { Carbon::UserWebPageOrderByTypes.new }

  describe 'test an instance of UserWebPageOrderByTypes' do
    it 'should create an instance of UserWebPageOrderByTypes' do
      expect(instance).to be_instance_of(Carbon::UserWebPageOrderByTypes)
    end
  end
end

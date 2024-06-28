=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::ListUsersOrderByTypes
describe Carbon::ListUsersOrderByTypes do
  let(:instance) { Carbon::ListUsersOrderByTypes.new }

  describe 'test an instance of ListUsersOrderByTypes' do
    it 'should create an instance of ListUsersOrderByTypes' do
      expect(instance).to be_instance_of(Carbon::ListUsersOrderByTypes)
    end
  end
end

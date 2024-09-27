=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::ContactsOrderBy
describe Carbon::ContactsOrderBy do
  let(:instance) { Carbon::ContactsOrderBy.new }

  describe 'test an instance of ContactsOrderBy' do
    it 'should create an instance of ContactsOrderBy' do
      expect(instance).to be_instance_of(Carbon::ContactsOrderBy)
    end
  end
end

=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::ContactsOrderByNullable
describe Carbon::ContactsOrderByNullable do
  let(:instance) { Carbon::ContactsOrderByNullable.new }

  describe 'test an instance of ContactsOrderByNullable' do
    it 'should create an instance of ContactsOrderByNullable' do
      expect(instance).to be_instance_of(Carbon::ContactsOrderByNullable)
    end
  end
end

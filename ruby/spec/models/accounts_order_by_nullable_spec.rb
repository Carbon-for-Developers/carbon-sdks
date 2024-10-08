=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::AccountsOrderByNullable
describe Carbon::AccountsOrderByNullable do
  let(:instance) { Carbon::AccountsOrderByNullable.new }

  describe 'test an instance of AccountsOrderByNullable' do
    it 'should create an instance of AccountsOrderByNullable' do
      expect(instance).to be_instance_of(Carbon::AccountsOrderByNullable)
    end
  end
end

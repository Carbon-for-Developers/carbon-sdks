=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::AccountsOrderBy
describe Carbon::AccountsOrderBy do
  let(:instance) { Carbon::AccountsOrderBy.new }

  describe 'test an instance of AccountsOrderBy' do
    it 'should create an instance of AccountsOrderBy' do
      expect(instance).to be_instance_of(Carbon::AccountsOrderBy)
    end
  end
end

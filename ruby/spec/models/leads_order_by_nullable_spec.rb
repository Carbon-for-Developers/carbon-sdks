=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::LeadsOrderByNullable
describe Carbon::LeadsOrderByNullable do
  let(:instance) { Carbon::LeadsOrderByNullable.new }

  describe 'test an instance of LeadsOrderByNullable' do
    it 'should create an instance of LeadsOrderByNullable' do
      expect(instance).to be_instance_of(Carbon::LeadsOrderByNullable)
    end
  end
end

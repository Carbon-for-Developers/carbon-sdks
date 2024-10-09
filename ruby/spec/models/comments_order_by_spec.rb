=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::CommentsOrderBy
describe Carbon::CommentsOrderBy do
  let(:instance) { Carbon::CommentsOrderBy.new }

  describe 'test an instance of CommentsOrderBy' do
    it 'should create an instance of CommentsOrderBy' do
      expect(instance).to be_instance_of(Carbon::CommentsOrderBy)
    end
  end
end

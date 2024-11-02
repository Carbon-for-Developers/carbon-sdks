=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::GithubFileTypes
describe Carbon::GithubFileTypes do
  let(:instance) { Carbon::GithubFileTypes.new }

  describe 'test an instance of GithubFileTypes' do
    it 'should create an instance of GithubFileTypes' do
      expect(instance).to be_instance_of(Carbon::GithubFileTypes)
    end
  end
end

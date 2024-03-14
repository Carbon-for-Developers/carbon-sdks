=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::SimpleOAuthDataSources
describe Carbon::SimpleOAuthDataSources do
  let(:instance) { Carbon::SimpleOAuthDataSources.new }

  describe 'test an instance of SimpleOAuthDataSources' do
    it 'should create an instance of SimpleOAuthDataSources' do
      expect(instance).to be_instance_of(Carbon::SimpleOAuthDataSources)
    end
  end
end

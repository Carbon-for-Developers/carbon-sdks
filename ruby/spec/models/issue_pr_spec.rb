=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::IssuePR
describe Carbon::IssuePR do
  let(:instance) { Carbon::IssuePR.new }

  describe 'test an instance of IssuePR' do
    it 'should create an instance of IssuePR' do
      expect(instance).to be_instance_of(Carbon::IssuePR)
    end
  end
  describe 'test attribute "url"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "merged_at"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

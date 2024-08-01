=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::SentWebhookRequestBody
describe Carbon::SentWebhookRequestBody do
  let(:instance) { Carbon::SentWebhookRequestBody.new }

  describe 'test an instance of SentWebhookRequestBody' do
    it 'should create an instance of SentWebhookRequestBody' do
      expect(instance).to be_instance_of(Carbon::SentWebhookRequestBody)
    end
  end
  describe 'test attribute "payload"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::SentWebhookPayload
describe Carbon::SentWebhookPayload do
  let(:instance) { Carbon::SentWebhookPayload.new }

  describe 'test an instance of SentWebhookPayload' do
    it 'should create an instance of SentWebhookPayload' do
      expect(instance).to be_instance_of(Carbon::SentWebhookPayload)
    end
  end
  describe 'test attribute "webhook_type"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
      # validator = Petstore::EnumTest::EnumAttributeValidator.new('String', ["ADD", "ALL_UPLOADED_FILES_QUEUED", "CANCEL", "CHECKUP", "DATA_SOURCE_READY", "EMBEDDING_STORAGE_MODIFIED", "FILES_CREATED", "FILES_SKIPPED", "FILE_DELETED", "FILE_ERROR", "FILE_READY", "FILE_STATISTICS_AGGREGATED", "FILE_SYNC_LIMIT_REACHED", "MOVED_TO_COLD_STORAGE", "MOVED_TO_HOT_STORAGE", "ORGANIZATION_USER_DELETED", "RATE_LIMIT_ERROR", "REVOKE", "SPARSE_VECTOR_QUEUE_STATUS", "UPDATE", "WEBPAGE_ERROR", "WEBPAGE_READY", "WEBSCRAPE_URLS_READY"])
      # validator.allowable_values.each do |value|
      #   expect { instance.webhook_type = value }.not_to raise_error
      # end
    end
  end

  describe 'test attribute "customer_id"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "timestamp"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "object"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

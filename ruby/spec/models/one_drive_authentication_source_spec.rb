=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::OneDriveAuthenticationSource
describe Carbon::OneDriveAuthenticationSource do
  let(:instance) { Carbon::OneDriveAuthenticationSource.new }

  describe 'test an instance of OneDriveAuthenticationSource' do
    it 'should create an instance of OneDriveAuthenticationSource' do
      expect(instance).to be_instance_of(Carbon::OneDriveAuthenticationSource)
    end
  end
end

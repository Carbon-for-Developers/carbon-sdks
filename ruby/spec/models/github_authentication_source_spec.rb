=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for Carbon::GithubAuthenticationSource
describe Carbon::GithubAuthenticationSource do
  let(:instance) { Carbon::GithubAuthenticationSource.new }

  describe 'test an instance of GithubAuthenticationSource' do
    it 'should create an instance of GithubAuthenticationSource' do
      expect(instance).to be_instance_of(Carbon::GithubAuthenticationSource)
    end
  end
end

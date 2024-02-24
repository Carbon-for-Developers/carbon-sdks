require 'spec_helper'
require 'json'
require 'securerandom'
require 'debug'

describe 'GettingStarted' do
  before do
  end

  after do
    # run after each test
  end

  it 'data_sources.query_user_data_sources' do
    debugger
    configuration = Carbon::Configuration.new
    configuration.access_token = 'YOUR API KEY'
    configuration.api_key = 'YOUR API KEY'
    configuration.customer_id = 'YOUR API KEY'
    configuration.host = 'http://127.0.0.1:4010'
    carbon = Carbon::Client.new(configuration)
    pagination = {
        "limit" => 10,
        "offset" => 0,
    }
    order_by = "created_at"
    order_dir = "desc"
    filters = {
            "source" => "GOOGLE_DRIVE",
        }
    result = carbon.data_sources.query_user_data_sources(
      pagination: pagination,
      order_by: order_by,
      order_dir: order_dir,
      filters: filters,
    )
    p result
  end
end

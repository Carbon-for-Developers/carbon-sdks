# Carbon::UtilitiesApi

All URIs are relative to *https://api.carbon.ai*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**fetch_urls**](UtilitiesApi.md#fetch_urls) | **GET** /fetch_urls | Fetch Urls |
| [**fetch_youtube_transcripts**](UtilitiesApi.md#fetch_youtube_transcripts) | **GET** /fetch_youtube_transcript | Fetch Youtube Transcripts |
| [**process_sitemap**](UtilitiesApi.md#process_sitemap) | **GET** /process_sitemap | Sitemap |
| [**scrape_sitemap**](UtilitiesApi.md#scrape_sitemap) | **POST** /scrape_sitemap | Scrape Sitemap |
| [**scrape_web**](UtilitiesApi.md#scrape_web) | **POST** /web_scrape | Web Scrape |
| [**search_urls**](UtilitiesApi.md#search_urls) | **GET** /search_urls | Search Urls |

## fetch_urls

Fetch Urls

Extracts all URLs from a webpage.   Args:     url (str): URL of the webpage  Returns:     FetchURLsResponse: A response object with a list of URLs extracted from the webpage and the webpage content.

### Example

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'
Carbon.api_key = 'YOUR API KEY'
Carbon.customer_id = 'YOUR API KEY'

url = "url_example"

begin
  # Fetch Urls
  result = Carbon::Utilities.fetch_urls(
                                           url: url,
                                         )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Utilities.fetch_urls: #{e}"
end
```

#### Using the fetch_urls_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
url = "url_example"

begin
  # Fetch Urls
  data, status_code, headers, response = Carbon::Utilities.fetch_urls_with_http_info(
                                                                                        url: url,
                                                                                      )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => FetchURLsResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Utilities.fetch_urls: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **url** | **String** |  |  |

### Return type

[**FetchURLsResponse**](FetchURLsResponse.md)

## fetch_youtube_transcripts

Fetch Youtube Transcripts

Fetches english transcripts from YouTube videos.  Args:     id (str): The ID of the YouTube video.      raw (bool): Whether to return the raw transcript or not. Defaults to False.  Returns:     dict: A dictionary with the transcript of the YouTube video.

### Example

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'
Carbon.api_key = 'YOUR API KEY'
Carbon.customer_id = 'YOUR API KEY'

id = "id_example"
raw = False

begin
  # Fetch Youtube Transcripts
  result = Carbon::Utilities.fetch_youtube_transcripts(
                                                          id: id,
                                                          raw: raw,
                                                        )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Utilities.fetch_youtube_transcripts: #{e}"
end
```

#### Using the fetch_youtube_transcripts_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
id = "id_example"
raw = False

begin
  # Fetch Youtube Transcripts
  data, status_code, headers, response = Carbon::Utilities.fetch_youtube_transcripts_with_http_info(
                                                                                                       id: id,
                                                                                                       raw: raw,
                                                                                                     )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => YoutubeTranscriptResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Utilities.fetch_youtube_transcripts: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **raw** | **Boolean** |  | [optional][default to false] |

### Return type

[**YoutubeTranscriptResponse**](YoutubeTranscriptResponse.md)

## process_sitemap

Sitemap

Retrieves all URLs from a sitemap, which can subsequently be utilized with our `web_scrape` endpoint.  <!--Args:     url (str): URL of the sitemap  Returns:     dict: A dictionary with a list of URLs extracted from the sitemap.-->

### Example

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'
Carbon.api_key = 'YOUR API KEY'
Carbon.customer_id = 'YOUR API KEY'

url = "url_example"

begin
  # Sitemap
  result = Carbon::Utilities.process_sitemap(
                                                url: url,
                                              )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Utilities.process_sitemap: #{e}"
end
```

#### Using the process_sitemap_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
url = "url_example"

begin
  # Sitemap
  data, status_code, headers, response = Carbon::Utilities.process_sitemap_with_http_info(
                                                                                             url: url,
                                                                                           )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => Object
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Utilities.process_sitemap: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **url** | **String** |  |  |

### Return type

**Object**

## scrape_sitemap

Scrape Sitemap

Extracts all URLs from a sitemap and performs a web scrape on each of them.  Args:     sitemap_url (str): URL of the sitemap  Returns:     dict: A response object with the status of the scraping job message.-->

### Example

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'
Carbon.api_key = 'YOUR API KEY'
Carbon.customer_id = 'YOUR API KEY'

url = "string_example"
tags = {
        "key": "string_example",
    }
max_pages_to_scrape = 100
chunk_size = 1500
chunk_overlap = 20
skip_embedding_generation = False
enable_auto_sync = False
generate_sparse_vectors = False
prepend_filename_to_chunks = False
html_tags_to_skip = []
css_classes_to_skip = []
css_selectors_to_skip = []

begin
  # Scrape Sitemap
  result = Carbon::Utilities.scrape_sitemap(
                                               url: url,
                                               tags: tags,
                                               max_pages_to_scrape: max_pages_to_scrape,
                                               chunk_size: chunk_size,
                                               chunk_overlap: chunk_overlap,
                                               skip_embedding_generation: skip_embedding_generation,
                                               enable_auto_sync: enable_auto_sync,
                                               generate_sparse_vectors: generate_sparse_vectors,
                                               prepend_filename_to_chunks: prepend_filename_to_chunks,
                                               html_tags_to_skip: html_tags_to_skip,
                                               css_classes_to_skip: css_classes_to_skip,
                                               css_selectors_to_skip: css_selectors_to_skip,
                                             )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Utilities.scrape_sitemap: #{e}"
end
```

#### Using the scrape_sitemap_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
url = "string_example"
tags = {
        "key": "string_example",
    }
max_pages_to_scrape = 100
chunk_size = 1500
chunk_overlap = 20
skip_embedding_generation = False
enable_auto_sync = False
generate_sparse_vectors = False
prepend_filename_to_chunks = False
html_tags_to_skip = []
css_classes_to_skip = []
css_selectors_to_skip = []

begin
  # Scrape Sitemap
  data, status_code, headers, response = Carbon::Utilities.scrape_sitemap_with_http_info(
                                                                                            url: url,
                                                                                            tags: tags,
                                                                                            max_pages_to_scrape: max_pages_to_scrape,
                                                                                            chunk_size: chunk_size,
                                                                                            chunk_overlap: chunk_overlap,
                                                                                            skip_embedding_generation: skip_embedding_generation,
                                                                                            enable_auto_sync: enable_auto_sync,
                                                                                            generate_sparse_vectors: generate_sparse_vectors,
                                                                                            prepend_filename_to_chunks: prepend_filename_to_chunks,
                                                                                            html_tags_to_skip: html_tags_to_skip,
                                                                                            css_classes_to_skip: css_classes_to_skip,
                                                                                            css_selectors_to_skip: css_selectors_to_skip,
                                                                                          )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => Object
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Utilities.scrape_sitemap: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **sitemap_scrape_request** | [**SitemapScrapeRequest**](SitemapScrapeRequest.md) |  |  |

### Return type

**Object**

## scrape_web

Web Scrape

Conduct a web scrape on a given webpage URL. Our web scraper is fully compatible with JavaScript and supports recursion depth, enabling you to efficiently extract all content from the target website.  <!--Args:     scraping_requests (List[WebscrapeRequest]): A list of WebscrapeRequest objects.       Returns:     dict: A response object with the status of the scraping job message.-->

### Example

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'
Carbon.api_key = 'YOUR API KEY'
Carbon.customer_id = 'YOUR API KEY'

body = [
        {
            "url" => "url_example",
            "recursion_depth" => 3,
            "max_pages_to_scrape" => 100,
            "chunk_size" => 1500,
            "chunk_overlap" => 20,
            "skip_embedding_generation" => False,
            "enable_auto_sync" => False,
            "generate_sparse_vectors" => False,
            "prepend_filename_to_chunks" => False,
            "html_tags_to_skip" => [],
            "css_classes_to_skip" => [],
            "css_selectors_to_skip" => [],
        }
    ]

begin
  # Web Scrape
  result = Carbon::Utilities.scrape_web(
                                           body: body,
                                         )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Utilities.scrape_web: #{e}"
end
```

#### Using the scrape_web_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
body = [
        {
            "url" => "url_example",
            "recursion_depth" => 3,
            "max_pages_to_scrape" => 100,
            "chunk_size" => 1500,
            "chunk_overlap" => 20,
            "skip_embedding_generation" => False,
            "enable_auto_sync" => False,
            "generate_sparse_vectors" => False,
            "prepend_filename_to_chunks" => False,
            "html_tags_to_skip" => [],
            "css_classes_to_skip" => [],
            "css_selectors_to_skip" => [],
        }
    ]

begin
  # Web Scrape
  data, status_code, headers, response = Carbon::Utilities.scrape_web_with_http_info(
                                                                                        body: body,
                                                                                      )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => Object
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Utilities.scrape_web: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **webscrape_request** | [**Array&lt;WebscrapeRequest&gt;**](WebscrapeRequest.md) |  |  |

### Return type

**Object**

## search_urls

Search Urls

Perform a web search and obtain a list of relevant URLs.  As an illustration, when you perform a search for “content related to MRNA,” you will receive a list of links such as the following:      - https://tomrenz.substack.com/p/mrna-and-why-it-matters      - https://www.statnews.com/2020/11/10/the-story-of-mrna-how-a-once-dismissed-idea-became-a-leading-technology-in-the-covid-vaccine-race/      - https://www.statnews.com/2022/11/16/covid-19-vaccines-were-a-success-but-mrna-still-has-a-delivery-problem/          - https://joomi.substack.com/p/were-still-being-misled-about-how  Subsequently, you can submit these links to the web_scrape endpoint in order to retrieve the content of the respective web pages.  Args:     query (str): Query to search for  Returns:     FetchURLsResponse: A response object with a list of URLs for a given search query.

### Example

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'
Carbon.api_key = 'YOUR API KEY'
Carbon.customer_id = 'YOUR API KEY'

query = "query_example"

begin
  # Search Urls
  result = Carbon::Utilities.search_urls(
                                            query: query,
                                          )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Utilities.search_urls: #{e}"
end
```

#### Using the search_urls_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
query = "query_example"

begin
  # Search Urls
  data, status_code, headers, response = Carbon::Utilities.search_urls_with_http_info(
                                                                                         query: query,
                                                                                       )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => FetchURLsResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Utilities.search_urls: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **query** | **String** |  |  |

### Return type

[**FetchURLsResponse**](FetchURLsResponse.md)


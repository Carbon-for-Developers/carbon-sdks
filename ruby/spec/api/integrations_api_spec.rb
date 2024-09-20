=begin
#Carbon

#Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0
=end

require 'spec_helper'
require 'json'

# Unit tests for Carbon::IntegrationsApi
describe 'IntegrationsApi' do
  before do
    # run before each test
    @api_instance = Carbon::IntegrationsApi.new
  end

  after do
    # run after each test
  end

  describe 'test an instance of IntegrationsApi' do
    it 'should create an instance of IntegrationsApi' do
      expect(@api_instance).to be_instance_of(Carbon::IntegrationsApi)
    end
  end

  # unit tests for cancel
  # Cancel Data Source Items Sync
  # @param sync_directory_request 
  # @param [Hash] opts the optional parameters
  # @return [OrganizationUserDataSourceAPI]
  describe 'cancel test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for connect_data_source
  # Connect Data Source
  # @param connect_data_source_input 
  # @param [Hash] opts the optional parameters
  # @return [ConnectDataSourceResponse]
  describe 'connect_data_source test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for connect_freshdesk
  # Freshdesk Connect
  # Refer this article to obtain an API key https://support.freshdesk.com/en/support/solutions/articles/215517. Make sure that your API key has the permission to read solutions from your account and you are on a &lt;b&gt;paid&lt;/b&gt; plan. Once you have an API key, you can make a request to this endpoint along with your freshdesk domain. This will  trigger an automatic sync of the articles in your \&quot;solutions\&quot; tab. Additional parameters below can be used to associate  data with the synced articles or modify the sync behavior.
  # @param fresh_desk_connect_request 
  # @param [Hash] opts the optional parameters
  # @return [GenericSuccessResponse]
  describe 'connect_freshdesk test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for connect_gitbook
  # Gitbook Connect
  # You will need an access token to connect your Gitbook account. Note that the permissions will be defined by the user  generating access token so make sure you have the permission to access spaces you will be syncing.  Refer this article for more details https://developer.gitbook.com/gitbook-api/authentication. Additionally, you need to specify the name of organization you will be syncing data from.
  # @param gitbook_connect_request 
  # @param [Hash] opts the optional parameters
  # @return [GenericSuccessResponse]
  describe 'connect_gitbook test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for connect_guru
  # Guru Connect
  # You will need an access token to connect your Guru account. To obtain an access token, follow the steps highlighted here https://help.getguru.com/docs/gurus-api#obtaining-a-user-token. The username should be your Guru username.
  # @param guru_connect_request 
  # @param [Hash] opts the optional parameters
  # @return [GenericSuccessResponse]
  describe 'connect_guru test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for create_aws_iam_user
  # S3 Auth
  # This endpoint can be used to connect S3 as well as Digital Ocean Spaces (S3 compatible)   For S3, create a new IAM user with permissions to: &lt;ol&gt; &lt;li&gt;List all buckets.&lt;/li&gt; &lt;li&gt;Read from the specific buckets and objects to sync with Carbon. Ensure any future buckets or objects carry  the same permissions.&lt;/li&gt; &lt;/ol&gt; Once created, generate an access key for this user and share the credentials with us. We recommend testing this key beforehand.   For Digital Ocean Spaces, generate the above credentials in your Applications and API page here https://cloud.digitalocean.com/account/api/spaces. Endpoint URL is required to connect Digital Ocean Spaces. It should look like &lt;&lt;region&gt;&gt;.digitaloceanspaces.com
  # @param s3_auth_request 
  # @param [Hash] opts the optional parameters
  # @return [OrganizationUserDataSourceAPI]
  describe 'create_aws_iam_user test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_oauth_url
  # Get Oauth Url
  # This endpoint can be used to generate the following URLs - An OAuth URL for OAuth based connectors - A file syncing URL which skips the OAuth flow if the user already has a valid access token and takes them to the success state.
  # @param o_auth_url_request 
  # @param [Hash] opts the optional parameters
  # @return [OuthURLResponse]
  describe 'get_oauth_url test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for list_confluence_pages
  # Confluence List
  # This endpoint has been deprecated. Use /integrations/items/list instead.  To begin listing a user&#39;s Confluence pages, at least a &#x60;data_source_id&#x60; of a connected Confluence account must be specified. This base request returns a list of root pages for every space the user has access to in a Confluence instance. To traverse further down the user&#39;s page directory, additional requests to this endpoint can be made with the same &#x60;data_source_id&#x60; and with &#x60;parent_id&#x60; set to the id of page from a previous request. For convenience, the &#x60;has_children&#x60; property in each directory item in the response list will flag which pages will return non-empty lists of pages when set as the &#x60;parent_id&#x60;.
  # @param list_request 
  # @param [Hash] opts the optional parameters
  # @return [ListResponse]
  describe 'list_confluence_pages test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for list_conversations
  # Slack List Conversations
  # List all of your public and private channels, DMs, and Group DMs. The ID from response  can be used as a filter to sync messages to Carbon    types: Comma separated list of types. Available types are im (DMs), mpim (group DMs), public_channel, and private_channel. Defaults to public_channel.     cursor: Used for pagination. If next_cursor is returned in response, you need to pass it as the cursor in the next request     data_source_id: Data source needs to be specified if you have linked multiple slack accounts   exclude_archived: Should archived conversations be excluded, defaults to true
  # @param [Hash] opts the optional parameters
  # @option opts [String] :types 
  # @option opts [String] :cursor 
  # @option opts [Integer] :data_source_id 
  # @option opts [Boolean] :exclude_archived 
  # @return [Object]
  describe 'list_conversations test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for list_data_source_items
  # List Data Source Items
  # @param list_data_source_items_request 
  # @param [Hash] opts the optional parameters
  # @return [ListDataSourceItemsResponse]
  describe 'list_data_source_items test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for list_folders
  # Outlook Folders
  # After connecting your Outlook account, you can use this endpoint to list all of your folders on outlook. This includes  both system folders like \&quot;inbox\&quot; and user created folders.
  # @param [Hash] opts the optional parameters
  # @option opts [Integer] :data_source_id 
  # @return [Object]
  describe 'list_folders test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for list_gitbook_spaces
  # Gitbook Spaces
  # After connecting your Gitbook account, you can use this endpoint to list all of your spaces under current organization.
  # @param data_source_id 
  # @param [Hash] opts the optional parameters
  # @return [Object]
  describe 'list_gitbook_spaces test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for list_labels
  # Gmail Labels
  # After connecting your Gmail account, you can use this endpoint to list all of your labels. User created labels will have the type \&quot;user\&quot; and Gmail&#39;s default labels will have the type \&quot;system\&quot;
  # @param [Hash] opts the optional parameters
  # @option opts [Integer] :data_source_id 
  # @return [Object]
  describe 'list_labels test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for list_outlook_categories
  # Outlook Categories
  # After connecting your Outlook account, you can use this endpoint to list all of your categories on outlook. We currently support listing up to 250 categories.
  # @param [Hash] opts the optional parameters
  # @option opts [Integer] :data_source_id 
  # @return [Object]
  describe 'list_outlook_categories test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for list_repos
  # Github List Repos
  # Once you have connected your GitHub account, you can use this endpoint to list the      repositories your account has access to. You can use a data source ID or username to fetch from a specific account.
  # @param [Hash] opts the optional parameters
  # @option opts [Integer] :per_page 
  # @option opts [Integer] :page 
  # @option opts [Integer] :data_source_id 
  # @return [Object]
  describe 'list_repos test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for sync_azure_blob_files
  # Azure Blob Files
  # After optionally loading the items via /integrations/items/sync and integrations/items/list, use the container name  and file name as the ID in this endpoint to sync them into Carbon. Additional parameters below can associate  data with the selected items or modify the sync behavior
  # @param azure_blob_file_sync_input 
  # @param [Hash] opts the optional parameters
  # @return [GenericSuccessResponse]
  describe 'sync_azure_blob_files test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for sync_azure_blob_storage
  # Azure Blob Storage Auth
  # This endpoint can be used to connect Azure Blob Storage.  For Azure Blob Storage, follow these steps: &lt;ol&gt;   &lt;li&gt;Create a new Azure Storage account and grant the following permissions:     &lt;ul&gt;       &lt;li&gt;List containers.&lt;/li&gt;       &lt;li&gt;Read from specific containers and blobs to sync with Carbon. Ensure any future containers or blobs carry the same permissions.&lt;/li&gt;     &lt;/ul&gt;   &lt;/li&gt;   &lt;li&gt;Generate a shared access signature (SAS) token or an access key for the storage account.&lt;/li&gt; &lt;/ol&gt;  Once created, provide us with the following details to generate the connection URL: &lt;ol&gt;   &lt;li&gt;Storage Account KeyName.&lt;/li&gt;   &lt;li&gt;Storage Account Name.&lt;/li&gt; &lt;/ol&gt;
  # @param azure_blob_auth_request 
  # @param [Hash] opts the optional parameters
  # @return [OrganizationUserDataSourceAPI]
  describe 'sync_azure_blob_storage test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for sync_confluence
  # Confluence Sync
  # This endpoint has been deprecated. Use /integrations/files/sync instead.  After listing pages in a user&#39;s Confluence account, the set of selected page &#x60;ids&#x60; and the connected account&#39;s &#x60;data_source_id&#x60; can be passed into this endpoint to sync them into Carbon. Additional parameters listed below can be used to associate data to the selected pages or alter the behavior of the sync.
  # @param sync_files_request 
  # @param [Hash] opts the optional parameters
  # @return [GenericSuccessResponse]
  describe 'sync_confluence test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for sync_data_source_items
  # Sync Data Source Items
  # @param sync_directory_request 
  # @param [Hash] opts the optional parameters
  # @return [OrganizationUserDataSourceAPI]
  describe 'sync_data_source_items test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for sync_files
  # Sync Files
  # After listing files and folders via /integrations/items/sync and integrations/items/list, use the selected items&#39; external ids  as the ids in this endpoint to sync them into Carbon. Sharepoint items take an additional parameter root_id, which identifies the drive the file or folder is in and is stored in root_external_id. That additional paramter is optional and excluding it will tell the sync to assume the item is stored in the default Documents drive.
  # @param sync_files_request 
  # @param [Hash] opts the optional parameters
  # @return [GenericSuccessResponse]
  describe 'sync_files test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for sync_git_hub
  # Github Connect
  # Refer this article to obtain an access token https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens. Make sure that your access token has the permission to read content from your desired repos. Note that if your access token expires you will need to manually update it through this endpoint.
  # @param github_connect_request 
  # @param [Hash] opts the optional parameters
  # @return [GenericSuccessResponse]
  describe 'sync_git_hub test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for sync_gitbook
  # Gitbook Sync
  # You can sync upto 20 Gitbook spaces at a time using this endpoint. Additional parameters below can be used to associate  data with the synced pages or modify the sync behavior.
  # @param gitbook_sync_request 
  # @param [Hash] opts the optional parameters
  # @return [Object]
  describe 'sync_gitbook test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for sync_gmail
  # Gmail Sync
  # Once you have successfully connected your gmail account, you can choose which emails to sync with us using the filters parameter. Filters is a JSON object with key value pairs. It also supports AND and OR operations. For now, we support a limited set of keys listed below.  &lt;b&gt;label&lt;/b&gt;: Inbuilt Gmail labels, for example \&quot;Important\&quot; or a custom label you created.   &lt;b&gt;after&lt;/b&gt; or &lt;b&gt;before&lt;/b&gt;: A date in YYYY/mm/dd format (example 2023/12/31). Gets emails after/before a certain date. You can also use them in combination to get emails from a certain period.   &lt;b&gt;is&lt;/b&gt;: Can have the following values - starred, important, snoozed, and unread   &lt;b&gt;from&lt;/b&gt;: Email address of the sender   &lt;b&gt;to&lt;/b&gt;: Email address of the recipient   &lt;b&gt;in&lt;/b&gt;: Can have the following values - sent (sync emails sent by the user)   &lt;b&gt;has&lt;/b&gt;: Can have the following values - attachment (sync emails that have attachments)    Using keys or values outside of the specified values can lead to unexpected behaviour.  An example of a basic query with filters can be &#x60;&#x60;&#x60;json {     \&quot;filters\&quot;: {             \&quot;key\&quot;: \&quot;label\&quot;,             \&quot;value\&quot;: \&quot;Test\&quot;         } } &#x60;&#x60;&#x60; Which will list all emails that have the label \&quot;Test\&quot;.  You can use AND and OR operation in the following way: &#x60;&#x60;&#x60;json {     \&quot;filters\&quot;: {         \&quot;AND\&quot;: [             {                 \&quot;key\&quot;: \&quot;after\&quot;,                 \&quot;value\&quot;: \&quot;2024/01/07\&quot;             },             {                 \&quot;OR\&quot;: [                     {                         \&quot;key\&quot;: \&quot;label\&quot;,                         \&quot;value\&quot;: \&quot;Personal\&quot;                     },                     {                         \&quot;key\&quot;: \&quot;is\&quot;,                         \&quot;value\&quot;: \&quot;starred\&quot;                     }                 ]             }         ]     } } &#x60;&#x60;&#x60; This will return emails after 7th of Jan that are either starred or have the label \&quot;Personal\&quot;.  Note that this is the highest level of nesting we support, i.e. you can&#39;t add more AND/OR filters within the OR filter in the above example.
  # @param gmail_sync_input 
  # @param [Hash] opts the optional parameters
  # @return [GenericSuccessResponse]
  describe 'sync_gmail test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for sync_outlook
  # Outlook Sync
  # Once you have successfully connected your Outlook account, you can choose which emails to sync with us using the filters and folder parameter. \&quot;folder\&quot; should be the folder you want to sync from Outlook. By default we get messages from your inbox folder.   Filters is a JSON object with key value pairs. It also supports AND and OR operations. For now, we support a limited set of keys listed below.  &lt;b&gt;category&lt;/b&gt;: Custom categories that you created in Outlook.   &lt;b&gt;after&lt;/b&gt; or &lt;b&gt;before&lt;/b&gt;: A date in YYYY/mm/dd format (example 2023/12/31). Gets emails after/before a certain date. You can also use them in combination to get emails from a certain period.     &lt;b&gt;is&lt;/b&gt;: Can have the following values: flagged   &lt;b&gt;from&lt;/b&gt;: Email address of the sender     An example of a basic query with filters can be &#x60;&#x60;&#x60;json {     \&quot;filters\&quot;: {             \&quot;key\&quot;: \&quot;category\&quot;,             \&quot;value\&quot;: \&quot;Test\&quot;         } } &#x60;&#x60;&#x60; Which will list all emails that have the category \&quot;Test\&quot;.    Specifying a custom folder in the same query &#x60;&#x60;&#x60;json {     \&quot;folder\&quot;: \&quot;Folder Name\&quot;,     \&quot;filters\&quot;: {             \&quot;key\&quot;: \&quot;category\&quot;,             \&quot;value\&quot;: \&quot;Test\&quot;         } } &#x60;&#x60;&#x60;  You can use AND and OR operation in the following way: &#x60;&#x60;&#x60;json {     \&quot;filters\&quot;: {         \&quot;AND\&quot;: [             {                 \&quot;key\&quot;: \&quot;after\&quot;,                 \&quot;value\&quot;: \&quot;2024/01/07\&quot;             },             {                 \&quot;OR\&quot;: [                     {                         \&quot;key\&quot;: \&quot;category\&quot;,                         \&quot;value\&quot;: \&quot;Personal\&quot;                     },                     {                         \&quot;key\&quot;: \&quot;category\&quot;,                         \&quot;value\&quot;: \&quot;Test\&quot;                     },                 ]             }         ]     } } &#x60;&#x60;&#x60; This will return emails after 7th of Jan that have either Personal or Test as category.  Note that this is the highest level of nesting we support, i.e. you can&#39;t add more AND/OR filters within the OR filter in the above example.
  # @param outlook_sync_input 
  # @param [Hash] opts the optional parameters
  # @return [GenericSuccessResponse]
  describe 'sync_outlook test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for sync_repos
  # Github Sync Repos
  # You can retreive repos your token has access to using /integrations/github/repos and sync their content.  You can also pass full name of any public repository (username/repo-name). This will store the repo content with  carbon which can be accessed through /integrations/items/list endpoint. Maximum of 25 repositories are accepted per request.
  # @param github_fetch_repos_request 
  # @param [Hash] opts the optional parameters
  # @return [Object]
  describe 'sync_repos test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for sync_rss_feed
  # Rss Feed
  # @param rss_feed_input 
  # @param [Hash] opts the optional parameters
  # @return [GenericSuccessResponse]
  describe 'sync_rss_feed test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for sync_s3_files
  # S3 Files
  # After optionally loading the items via /integrations/items/sync and integrations/items/list, use the bucket name  and object key as the ID in this endpoint to sync them into Carbon. Additional parameters below can associate  data with the selected items or modify the sync behavior
  # @param s3_file_sync_input 
  # @param [Hash] opts the optional parameters
  # @return [GenericSuccessResponse]
  describe 'sync_s3_files test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for sync_slack
  # Slack Sync
  # You can list all conversations using the endpoint /integrations/slack/conversations. The ID of  conversation will be used as an input for this endpoint with timestamps as optional filters.
  # @param slack_sync_request 
  # @param [Hash] opts the optional parameters
  # @return [Object]
  describe 'sync_slack test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

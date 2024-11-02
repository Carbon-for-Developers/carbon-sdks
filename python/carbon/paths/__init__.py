# do not import all endpoints into this module because that uses a lot of memory and stack frames
# if you need the ability to import all endpoints from this module, import them with
# from carbon.apis.path_to_api import path_to_api

import enum


class PathValues(str, enum.Enum):
    INTEGRATIONS_OAUTH_URL = "/integrations/oauth_url"
    INTEGRATIONS_CONNECT = "/integrations/connect"
    INTEGRATIONS_ITEMS_SYNC = "/integrations/items/sync"
    INTEGRATIONS_ITEMS_SYNC_CANCEL = "/integrations/items/sync/cancel"
    INTEGRATIONS_ITEMS_LIST = "/integrations/items/list"
    INTEGRATIONS_FILES_SYNC = "/integrations/files/sync"
    INTEGRATIONS_CONFLUENCE_LIST = "/integrations/confluence/list"
    INTEGRATIONS_CONFLUENCE_SYNC = "/integrations/confluence/sync"
    INTEGRATIONS_AZURE_BLOB_STORAGE = "/integrations/azure_blob_storage"
    INTEGRATIONS_S3 = "/integrations/s3"
    INTEGRATIONS_S3_FILES = "/integrations/s3/files"
    INTEGRATIONS_AZURE_BLOB_STORAGE_FILES = "/integrations/azure_blob_storage/files"
    INTEGRATIONS_GMAIL_SYNC = "/integrations/gmail/sync"
    INTEGRATIONS_RSS_FEED = "/integrations/rss_feed"
    INTEGRATIONS_OUTLOOK_SYNC = "/integrations/outlook/sync"
    INTEGRATIONS_FRESHDESK = "/integrations/freshdesk"
    INTEGRATIONS_OUTLOOK_USER_FOLDERS = "/integrations/outlook/user_folders"
    INTEGRATIONS_OUTLOOK_USER_CATEGORIES = "/integrations/outlook/user_categories"
    INTEGRATIONS_GMAIL_USER_LABELS = "/integrations/gmail/user_labels"
    INTEGRATIONS_GITBOOK = "/integrations/gitbook"
    INTEGRATIONS_GITBOOK_SPACES = "/integrations/gitbook/spaces"
    INTEGRATIONS_GITBOOK_SYNC = "/integrations/gitbook/sync"
    INTEGRATIONS_GITHUB = "/integrations/github"
    INTEGRATIONS_GITHUB_REPOS = "/integrations/github/repos"
    INTEGRATIONS_GITHUB_SYNC_REPOS = "/integrations/github/sync_repos"
    INTEGRATIONS_SLACK_CONVERSATIONS = "/integrations/slack/conversations"
    INTEGRATIONS_SLACK_SYNC = "/integrations/slack/sync"
    INTEGRATIONS_GURU = "/integrations/guru"
    INTEGRATIONS_SHAREPOINT_SITES_LIST = "/integrations/sharepoint/sites/list"
    AUTH_V1_ACCESS_TOKEN = "/auth/v1/access_token"
    AUTH_V1_WHITE_LABELING = "/auth/v1/white_labeling"
    EMBEDDINGS = "/embeddings"
    TEXT_CHUNKS = "/text_chunks"
    LIST_CHUNKS_AND_EMBEDDINGS = "/list_chunks_and_embeddings"
    UPLOAD_CHUNKS_AND_EMBEDDINGS = "/upload_chunks_and_embeddings"
    ORGANIZATION = "/organization"
    ORGANIZATION_UPDATE = "/organization/update"
    ORGANIZATION_STATISTICS = "/organization/statistics"
    USER = "/user"
    WHOAMI = "/whoami"
    MODIFY_USER_CONFIGURATION = "/modify_user_configuration"
    DELETE_USERS = "/delete_users"
    UPDATE_USERS = "/update_users"
    LIST_USERS = "/list_users"
    UPLOADFILE = "/uploadfile"
    UPLOAD_FILE_FROM_URL = "/upload_file_from_url"
    UPLOAD_TEXT = "/upload_text"
    DELETE_FILES_V2 = "/delete_files_v2"
    USER_FILES_V2 = "/user_files_v2"
    CREATE_USER_FILE_TAGS = "/create_user_file_tags"
    DELETE_USER_FILE_TAGS = "/delete_user_file_tags"
    RESYNC_FILE = "/resync_file"
    MOVE_TO_HOT_STORAGE = "/move_to_hot_storage"
    MODIFY_COLD_STORAGE_PARAMETERS = "/modify_cold_storage_parameters"
    RAW_FILE_FILE_ID = "/raw_file/{file_id}"
    PARSED_FILE_FILE_ID = "/parsed_file/{file_id}"
    USER_FILES = "/user_files"
    DELETEFILE_FILE_ID = "/deletefile/{file_id}"
    DELETE_FILES = "/delete_files"
    WEBHOOKS = "/webhooks"
    ADD_WEBHOOK = "/add_webhook"
    DELETE_WEBHOOK_WEBHOOK_ID = "/delete_webhook/{webhook_id}"
    USER_DATA_SOURCES = "/user_data_sources"
    REVOKE_ACCESS_TOKEN = "/revoke_access_token"
    DATA_SOURCES_TAGS_ADD = "/data_sources/tags/add"
    DATA_SOURCES_TAGS_REMOVE = "/data_sources/tags/remove"
    WEB_SCRAPE = "/web_scrape"
    PROCESS_SITEMAP = "/process_sitemap"
    SCRAPE_SITEMAP = "/scrape_sitemap"
    FETCH_URLS = "/fetch_urls"
    FETCH_WEBPAGE = "/fetch_webpage"
    USER_WEBPAGES = "/user_webpages"
    SEARCH_URLS = "/search_urls"
    FETCH_YOUTUBE_TRANSCRIPT = "/fetch_youtube_transcript"
    INTEGRATIONS_DATA_CRM_ACCOUNTS = "/integrations/data/crm/accounts"
    INTEGRATIONS_DATA_CRM_ACCOUNTS_ID = "/integrations/data/crm/accounts/{id}"
    INTEGRATIONS_DATA_CRM_LEADS = "/integrations/data/crm/leads"
    INTEGRATIONS_DATA_CRM_LEADS_ID = "/integrations/data/crm/leads/{id}"
    INTEGRATIONS_DATA_CRM_CONTACTS = "/integrations/data/crm/contacts"
    INTEGRATIONS_DATA_CRM_CONTACTS_ID = "/integrations/data/crm/contacts/{id}"
    INTEGRATIONS_DATA_CRM_OPPORTUNITIES = "/integrations/data/crm/opportunities"
    INTEGRATIONS_DATA_CRM_OPPORTUNITIES_ID = "/integrations/data/crm/opportunities/{id}"
    INTEGRATIONS_DATA_GITHUB_PULL_REQUESTS = "/integrations/data/github/pull_requests"
    INTEGRATIONS_DATA_GITHUB_PULL_REQUESTS_PULL_NUMBER = "/integrations/data/github/pull_requests/{pull_number}"
    INTEGRATIONS_DATA_GITHUB_PULL_REQUESTS_COMMENTS = "/integrations/data/github/pull_requests/comments"
    INTEGRATIONS_DATA_GITHUB_PULL_REQUESTS_FILES = "/integrations/data/github/pull_requests/files"
    INTEGRATIONS_DATA_GITHUB_PULL_REQUESTS_COMMITS = "/integrations/data/github/pull_requests/commits"
    INTEGRATIONS_DATA_GITHUB_ISSUES = "/integrations/data/github/issues"
    INTEGRATIONS_DATA_GITHUB_ISSUES_ISSUE_NUMBER = "/integrations/data/github/issues/{issue_number}"
    WHITE_LABEL_CREATE = "/white_label/create"
    WHITE_LABEL_LIST = "/white_label/list"
    WHITE_LABEL_DELETE = "/white_label/delete"
    WHITE_LABEL_UPDATE = "/white_label/update"

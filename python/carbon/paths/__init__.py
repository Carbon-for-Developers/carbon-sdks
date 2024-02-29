# do not import all endpoints into this module because that uses a lot of memory and stack frames
# if you need the ability to import all endpoints from this module, import them with
# from carbon.apis.path_to_api import path_to_api

import enum


class PathValues(str, enum.Enum):
    INTEGRATIONS_OAUTH_URL = "/integrations/oauth_url"
    INTEGRATIONS_ITEMS_SYNC = "/integrations/items/sync"
    INTEGRATIONS_ITEMS_LIST = "/integrations/items/list"
    INTEGRATIONS_FILES_SYNC = "/integrations/files/sync"
    INTEGRATIONS_CONFLUENCE_LIST = "/integrations/confluence/list"
    INTEGRATIONS_CONFLUENCE_SYNC = "/integrations/confluence/sync"
    INTEGRATIONS_S3 = "/integrations/s3"
    INTEGRATIONS_S3_FILES = "/integrations/s3/files"
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
    AUTH_V1_ACCESS_TOKEN = "/auth/v1/access_token"
    AUTH_V1_WHITE_LABELING = "/auth/v1/white_labeling"
    EMBEDDINGS = "/embeddings"
    TEXT_CHUNKS = "/text_chunks"
    UPLOAD_CHUNKS_AND_EMBEDDINGS = "/upload_chunks_and_embeddings"
    ORGANIZATION = "/organization"
    USER = "/user"
    MODIFY_USER_CONFIGURATION = "/modify_user_configuration"
    UPLOADFILE = "/uploadfile"
    UPLOAD_FILE_FROM_URL = "/upload_file_from_url"
    UPLOAD_TEXT = "/upload_text"
    DELETEFILE_FILE_ID = "/deletefile/{file_id}"
    DELETE_FILES = "/delete_files"
    USER_FILES_V2 = "/user_files_v2"
    CREATE_USER_FILE_TAGS = "/create_user_file_tags"
    DELETE_USER_FILE_TAGS = "/delete_user_file_tags"
    RESYNC_FILE = "/resync_file"
    RAW_FILE_FILE_ID = "/raw_file/{file_id}"
    PARSED_FILE_FILE_ID = "/parsed_file/{file_id}"
    USER_FILES = "/user_files"
    WEBHOOKS = "/webhooks"
    ADD_WEBHOOK = "/add_webhook"
    DELETE_WEBHOOK_WEBHOOK_ID = "/delete_webhook/{webhook_id}"
    USER_DATA_SOURCES = "/user_data_sources"
    REVOKE_ACCESS_TOKEN = "/revoke_access_token"
    WEB_SCRAPE = "/web_scrape"
    PROCESS_SITEMAP = "/process_sitemap"
    SCRAPE_SITEMAP = "/scrape_sitemap"
    FETCH_URLS = "/fetch_urls"
    SEARCH_URLS = "/search_urls"
    FETCH_YOUTUBE_TRANSCRIPT = "/fetch_youtube_transcript"
    HEALTH = "/health"

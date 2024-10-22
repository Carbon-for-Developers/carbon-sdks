import typing_extensions

from carbon.paths import PathValues
from carbon.apis.paths.integrations_oauth_url import IntegrationsOauthUrl
from carbon.apis.paths.integrations_connect import IntegrationsConnect
from carbon.apis.paths.integrations_items_sync import IntegrationsItemsSync
from carbon.apis.paths.integrations_items_sync_cancel import IntegrationsItemsSyncCancel
from carbon.apis.paths.integrations_items_list import IntegrationsItemsList
from carbon.apis.paths.integrations_files_sync import IntegrationsFilesSync
from carbon.apis.paths.integrations_confluence_list import IntegrationsConfluenceList
from carbon.apis.paths.integrations_confluence_sync import IntegrationsConfluenceSync
from carbon.apis.paths.integrations_azure_blob_storage import IntegrationsAzureBlobStorage
from carbon.apis.paths.integrations_s3 import IntegrationsS3
from carbon.apis.paths.integrations_s3_files import IntegrationsS3Files
from carbon.apis.paths.integrations_azure_blob_storage_files import IntegrationsAzureBlobStorageFiles
from carbon.apis.paths.integrations_gmail_sync import IntegrationsGmailSync
from carbon.apis.paths.integrations_rss_feed import IntegrationsRssFeed
from carbon.apis.paths.integrations_outlook_sync import IntegrationsOutlookSync
from carbon.apis.paths.integrations_freshdesk import IntegrationsFreshdesk
from carbon.apis.paths.integrations_outlook_user_folders import IntegrationsOutlookUserFolders
from carbon.apis.paths.integrations_outlook_user_categories import IntegrationsOutlookUserCategories
from carbon.apis.paths.integrations_gmail_user_labels import IntegrationsGmailUserLabels
from carbon.apis.paths.integrations_gitbook import IntegrationsGitbook
from carbon.apis.paths.integrations_gitbook_spaces import IntegrationsGitbookSpaces
from carbon.apis.paths.integrations_gitbook_sync import IntegrationsGitbookSync
from carbon.apis.paths.integrations_github import IntegrationsGithub
from carbon.apis.paths.integrations_github_repos import IntegrationsGithubRepos
from carbon.apis.paths.integrations_github_sync_repos import IntegrationsGithubSyncRepos
from carbon.apis.paths.integrations_slack_conversations import IntegrationsSlackConversations
from carbon.apis.paths.integrations_slack_sync import IntegrationsSlackSync
from carbon.apis.paths.integrations_guru import IntegrationsGuru
from carbon.apis.paths.auth_v1_access_token import AuthV1AccessToken
from carbon.apis.paths.auth_v1_white_labeling import AuthV1WhiteLabeling
from carbon.apis.paths.embeddings import Embeddings
from carbon.apis.paths.text_chunks import TextChunks
from carbon.apis.paths.list_chunks_and_embeddings import ListChunksAndEmbeddings
from carbon.apis.paths.upload_chunks_and_embeddings import UploadChunksAndEmbeddings
from carbon.apis.paths.organization import Organization
from carbon.apis.paths.organization_update import OrganizationUpdate
from carbon.apis.paths.organization_statistics import OrganizationStatistics
from carbon.apis.paths.user import User
from carbon.apis.paths.whoami import Whoami
from carbon.apis.paths.modify_user_configuration import ModifyUserConfiguration
from carbon.apis.paths.delete_users import DeleteUsers
from carbon.apis.paths.update_users import UpdateUsers
from carbon.apis.paths.list_users import ListUsers
from carbon.apis.paths.uploadfile import Uploadfile
from carbon.apis.paths.upload_file_from_url import UploadFileFromUrl
from carbon.apis.paths.upload_text import UploadText
from carbon.apis.paths.delete_files_v2 import DeleteFilesV2
from carbon.apis.paths.user_files_v2 import UserFilesV2
from carbon.apis.paths.create_user_file_tags import CreateUserFileTags
from carbon.apis.paths.delete_user_file_tags import DeleteUserFileTags
from carbon.apis.paths.resync_file import ResyncFile
from carbon.apis.paths.move_to_hot_storage import MoveToHotStorage
from carbon.apis.paths.modify_cold_storage_parameters import ModifyColdStorageParameters
from carbon.apis.paths.raw_file_file_id import RawFileFileId
from carbon.apis.paths.parsed_file_file_id import ParsedFileFileId
from carbon.apis.paths.user_files import UserFiles
from carbon.apis.paths.deletefile_file_id import DeletefileFileId
from carbon.apis.paths.delete_files import DeleteFiles
from carbon.apis.paths.webhooks import Webhooks
from carbon.apis.paths.add_webhook import AddWebhook
from carbon.apis.paths.delete_webhook_webhook_id import DeleteWebhookWebhookId
from carbon.apis.paths.user_data_sources import UserDataSources
from carbon.apis.paths.revoke_access_token import RevokeAccessToken
from carbon.apis.paths.data_sources_tags_add import DataSourcesTagsAdd
from carbon.apis.paths.data_sources_tags_remove import DataSourcesTagsRemove
from carbon.apis.paths.web_scrape import WebScrape
from carbon.apis.paths.process_sitemap import ProcessSitemap
from carbon.apis.paths.scrape_sitemap import ScrapeSitemap
from carbon.apis.paths.fetch_urls import FetchUrls
from carbon.apis.paths.fetch_webpage import FetchWebpage
from carbon.apis.paths.user_webpages import UserWebpages
from carbon.apis.paths.search_urls import SearchUrls
from carbon.apis.paths.fetch_youtube_transcript import FetchYoutubeTranscript
from carbon.apis.paths.integrations_data_crm_accounts import IntegrationsDataCrmAccounts
from carbon.apis.paths.integrations_data_crm_accounts_id import IntegrationsDataCrmAccountsId
from carbon.apis.paths.integrations_data_crm_leads import IntegrationsDataCrmLeads
from carbon.apis.paths.integrations_data_crm_leads_id import IntegrationsDataCrmLeadsId
from carbon.apis.paths.integrations_data_crm_contacts import IntegrationsDataCrmContacts
from carbon.apis.paths.integrations_data_crm_contacts_id import IntegrationsDataCrmContactsId
from carbon.apis.paths.integrations_data_crm_opportunities import IntegrationsDataCrmOpportunities
from carbon.apis.paths.integrations_data_crm_opportunities_id import IntegrationsDataCrmOpportunitiesId
from carbon.apis.paths.integrations_data_github_pull_requests import IntegrationsDataGithubPullRequests
from carbon.apis.paths.integrations_data_github_pull_requests_pull_number import IntegrationsDataGithubPullRequestsPullNumber
from carbon.apis.paths.integrations_data_github_pull_requests_comments import IntegrationsDataGithubPullRequestsComments
from carbon.apis.paths.integrations_data_github_pull_requests_files import IntegrationsDataGithubPullRequestsFiles
from carbon.apis.paths.integrations_data_github_pull_requests_commits import IntegrationsDataGithubPullRequestsCommits
from carbon.apis.paths.integrations_data_github_issues import IntegrationsDataGithubIssues
from carbon.apis.paths.integrations_data_github_issues_issue_number import IntegrationsDataGithubIssuesIssueNumber
from carbon.apis.paths.white_label_create import WhiteLabelCreate
from carbon.apis.paths.white_label_list import WhiteLabelList
from carbon.apis.paths.white_label_delete import WhiteLabelDelete
from carbon.apis.paths.white_label_update import WhiteLabelUpdate

PathToApi = typing_extensions.TypedDict(
    'PathToApi',
    {
        PathValues.INTEGRATIONS_OAUTH_URL: IntegrationsOauthUrl,
        PathValues.INTEGRATIONS_CONNECT: IntegrationsConnect,
        PathValues.INTEGRATIONS_ITEMS_SYNC: IntegrationsItemsSync,
        PathValues.INTEGRATIONS_ITEMS_SYNC_CANCEL: IntegrationsItemsSyncCancel,
        PathValues.INTEGRATIONS_ITEMS_LIST: IntegrationsItemsList,
        PathValues.INTEGRATIONS_FILES_SYNC: IntegrationsFilesSync,
        PathValues.INTEGRATIONS_CONFLUENCE_LIST: IntegrationsConfluenceList,
        PathValues.INTEGRATIONS_CONFLUENCE_SYNC: IntegrationsConfluenceSync,
        PathValues.INTEGRATIONS_AZURE_BLOB_STORAGE: IntegrationsAzureBlobStorage,
        PathValues.INTEGRATIONS_S3: IntegrationsS3,
        PathValues.INTEGRATIONS_S3_FILES: IntegrationsS3Files,
        PathValues.INTEGRATIONS_AZURE_BLOB_STORAGE_FILES: IntegrationsAzureBlobStorageFiles,
        PathValues.INTEGRATIONS_GMAIL_SYNC: IntegrationsGmailSync,
        PathValues.INTEGRATIONS_RSS_FEED: IntegrationsRssFeed,
        PathValues.INTEGRATIONS_OUTLOOK_SYNC: IntegrationsOutlookSync,
        PathValues.INTEGRATIONS_FRESHDESK: IntegrationsFreshdesk,
        PathValues.INTEGRATIONS_OUTLOOK_USER_FOLDERS: IntegrationsOutlookUserFolders,
        PathValues.INTEGRATIONS_OUTLOOK_USER_CATEGORIES: IntegrationsOutlookUserCategories,
        PathValues.INTEGRATIONS_GMAIL_USER_LABELS: IntegrationsGmailUserLabels,
        PathValues.INTEGRATIONS_GITBOOK: IntegrationsGitbook,
        PathValues.INTEGRATIONS_GITBOOK_SPACES: IntegrationsGitbookSpaces,
        PathValues.INTEGRATIONS_GITBOOK_SYNC: IntegrationsGitbookSync,
        PathValues.INTEGRATIONS_GITHUB: IntegrationsGithub,
        PathValues.INTEGRATIONS_GITHUB_REPOS: IntegrationsGithubRepos,
        PathValues.INTEGRATIONS_GITHUB_SYNC_REPOS: IntegrationsGithubSyncRepos,
        PathValues.INTEGRATIONS_SLACK_CONVERSATIONS: IntegrationsSlackConversations,
        PathValues.INTEGRATIONS_SLACK_SYNC: IntegrationsSlackSync,
        PathValues.INTEGRATIONS_GURU: IntegrationsGuru,
        PathValues.AUTH_V1_ACCESS_TOKEN: AuthV1AccessToken,
        PathValues.AUTH_V1_WHITE_LABELING: AuthV1WhiteLabeling,
        PathValues.EMBEDDINGS: Embeddings,
        PathValues.TEXT_CHUNKS: TextChunks,
        PathValues.LIST_CHUNKS_AND_EMBEDDINGS: ListChunksAndEmbeddings,
        PathValues.UPLOAD_CHUNKS_AND_EMBEDDINGS: UploadChunksAndEmbeddings,
        PathValues.ORGANIZATION: Organization,
        PathValues.ORGANIZATION_UPDATE: OrganizationUpdate,
        PathValues.ORGANIZATION_STATISTICS: OrganizationStatistics,
        PathValues.USER: User,
        PathValues.WHOAMI: Whoami,
        PathValues.MODIFY_USER_CONFIGURATION: ModifyUserConfiguration,
        PathValues.DELETE_USERS: DeleteUsers,
        PathValues.UPDATE_USERS: UpdateUsers,
        PathValues.LIST_USERS: ListUsers,
        PathValues.UPLOADFILE: Uploadfile,
        PathValues.UPLOAD_FILE_FROM_URL: UploadFileFromUrl,
        PathValues.UPLOAD_TEXT: UploadText,
        PathValues.DELETE_FILES_V2: DeleteFilesV2,
        PathValues.USER_FILES_V2: UserFilesV2,
        PathValues.CREATE_USER_FILE_TAGS: CreateUserFileTags,
        PathValues.DELETE_USER_FILE_TAGS: DeleteUserFileTags,
        PathValues.RESYNC_FILE: ResyncFile,
        PathValues.MOVE_TO_HOT_STORAGE: MoveToHotStorage,
        PathValues.MODIFY_COLD_STORAGE_PARAMETERS: ModifyColdStorageParameters,
        PathValues.RAW_FILE_FILE_ID: RawFileFileId,
        PathValues.PARSED_FILE_FILE_ID: ParsedFileFileId,
        PathValues.USER_FILES: UserFiles,
        PathValues.DELETEFILE_FILE_ID: DeletefileFileId,
        PathValues.DELETE_FILES: DeleteFiles,
        PathValues.WEBHOOKS: Webhooks,
        PathValues.ADD_WEBHOOK: AddWebhook,
        PathValues.DELETE_WEBHOOK_WEBHOOK_ID: DeleteWebhookWebhookId,
        PathValues.USER_DATA_SOURCES: UserDataSources,
        PathValues.REVOKE_ACCESS_TOKEN: RevokeAccessToken,
        PathValues.DATA_SOURCES_TAGS_ADD: DataSourcesTagsAdd,
        PathValues.DATA_SOURCES_TAGS_REMOVE: DataSourcesTagsRemove,
        PathValues.WEB_SCRAPE: WebScrape,
        PathValues.PROCESS_SITEMAP: ProcessSitemap,
        PathValues.SCRAPE_SITEMAP: ScrapeSitemap,
        PathValues.FETCH_URLS: FetchUrls,
        PathValues.FETCH_WEBPAGE: FetchWebpage,
        PathValues.USER_WEBPAGES: UserWebpages,
        PathValues.SEARCH_URLS: SearchUrls,
        PathValues.FETCH_YOUTUBE_TRANSCRIPT: FetchYoutubeTranscript,
        PathValues.INTEGRATIONS_DATA_CRM_ACCOUNTS: IntegrationsDataCrmAccounts,
        PathValues.INTEGRATIONS_DATA_CRM_ACCOUNTS_ID: IntegrationsDataCrmAccountsId,
        PathValues.INTEGRATIONS_DATA_CRM_LEADS: IntegrationsDataCrmLeads,
        PathValues.INTEGRATIONS_DATA_CRM_LEADS_ID: IntegrationsDataCrmLeadsId,
        PathValues.INTEGRATIONS_DATA_CRM_CONTACTS: IntegrationsDataCrmContacts,
        PathValues.INTEGRATIONS_DATA_CRM_CONTACTS_ID: IntegrationsDataCrmContactsId,
        PathValues.INTEGRATIONS_DATA_CRM_OPPORTUNITIES: IntegrationsDataCrmOpportunities,
        PathValues.INTEGRATIONS_DATA_CRM_OPPORTUNITIES_ID: IntegrationsDataCrmOpportunitiesId,
        PathValues.INTEGRATIONS_DATA_GITHUB_PULL_REQUESTS: IntegrationsDataGithubPullRequests,
        PathValues.INTEGRATIONS_DATA_GITHUB_PULL_REQUESTS_PULL_NUMBER: IntegrationsDataGithubPullRequestsPullNumber,
        PathValues.INTEGRATIONS_DATA_GITHUB_PULL_REQUESTS_COMMENTS: IntegrationsDataGithubPullRequestsComments,
        PathValues.INTEGRATIONS_DATA_GITHUB_PULL_REQUESTS_FILES: IntegrationsDataGithubPullRequestsFiles,
        PathValues.INTEGRATIONS_DATA_GITHUB_PULL_REQUESTS_COMMITS: IntegrationsDataGithubPullRequestsCommits,
        PathValues.INTEGRATIONS_DATA_GITHUB_ISSUES: IntegrationsDataGithubIssues,
        PathValues.INTEGRATIONS_DATA_GITHUB_ISSUES_ISSUE_NUMBER: IntegrationsDataGithubIssuesIssueNumber,
        PathValues.WHITE_LABEL_CREATE: WhiteLabelCreate,
        PathValues.WHITE_LABEL_LIST: WhiteLabelList,
        PathValues.WHITE_LABEL_DELETE: WhiteLabelDelete,
        PathValues.WHITE_LABEL_UPDATE: WhiteLabelUpdate,
    }
)

path_to_api = PathToApi(
    {
        PathValues.INTEGRATIONS_OAUTH_URL: IntegrationsOauthUrl,
        PathValues.INTEGRATIONS_CONNECT: IntegrationsConnect,
        PathValues.INTEGRATIONS_ITEMS_SYNC: IntegrationsItemsSync,
        PathValues.INTEGRATIONS_ITEMS_SYNC_CANCEL: IntegrationsItemsSyncCancel,
        PathValues.INTEGRATIONS_ITEMS_LIST: IntegrationsItemsList,
        PathValues.INTEGRATIONS_FILES_SYNC: IntegrationsFilesSync,
        PathValues.INTEGRATIONS_CONFLUENCE_LIST: IntegrationsConfluenceList,
        PathValues.INTEGRATIONS_CONFLUENCE_SYNC: IntegrationsConfluenceSync,
        PathValues.INTEGRATIONS_AZURE_BLOB_STORAGE: IntegrationsAzureBlobStorage,
        PathValues.INTEGRATIONS_S3: IntegrationsS3,
        PathValues.INTEGRATIONS_S3_FILES: IntegrationsS3Files,
        PathValues.INTEGRATIONS_AZURE_BLOB_STORAGE_FILES: IntegrationsAzureBlobStorageFiles,
        PathValues.INTEGRATIONS_GMAIL_SYNC: IntegrationsGmailSync,
        PathValues.INTEGRATIONS_RSS_FEED: IntegrationsRssFeed,
        PathValues.INTEGRATIONS_OUTLOOK_SYNC: IntegrationsOutlookSync,
        PathValues.INTEGRATIONS_FRESHDESK: IntegrationsFreshdesk,
        PathValues.INTEGRATIONS_OUTLOOK_USER_FOLDERS: IntegrationsOutlookUserFolders,
        PathValues.INTEGRATIONS_OUTLOOK_USER_CATEGORIES: IntegrationsOutlookUserCategories,
        PathValues.INTEGRATIONS_GMAIL_USER_LABELS: IntegrationsGmailUserLabels,
        PathValues.INTEGRATIONS_GITBOOK: IntegrationsGitbook,
        PathValues.INTEGRATIONS_GITBOOK_SPACES: IntegrationsGitbookSpaces,
        PathValues.INTEGRATIONS_GITBOOK_SYNC: IntegrationsGitbookSync,
        PathValues.INTEGRATIONS_GITHUB: IntegrationsGithub,
        PathValues.INTEGRATIONS_GITHUB_REPOS: IntegrationsGithubRepos,
        PathValues.INTEGRATIONS_GITHUB_SYNC_REPOS: IntegrationsGithubSyncRepos,
        PathValues.INTEGRATIONS_SLACK_CONVERSATIONS: IntegrationsSlackConversations,
        PathValues.INTEGRATIONS_SLACK_SYNC: IntegrationsSlackSync,
        PathValues.INTEGRATIONS_GURU: IntegrationsGuru,
        PathValues.AUTH_V1_ACCESS_TOKEN: AuthV1AccessToken,
        PathValues.AUTH_V1_WHITE_LABELING: AuthV1WhiteLabeling,
        PathValues.EMBEDDINGS: Embeddings,
        PathValues.TEXT_CHUNKS: TextChunks,
        PathValues.LIST_CHUNKS_AND_EMBEDDINGS: ListChunksAndEmbeddings,
        PathValues.UPLOAD_CHUNKS_AND_EMBEDDINGS: UploadChunksAndEmbeddings,
        PathValues.ORGANIZATION: Organization,
        PathValues.ORGANIZATION_UPDATE: OrganizationUpdate,
        PathValues.ORGANIZATION_STATISTICS: OrganizationStatistics,
        PathValues.USER: User,
        PathValues.WHOAMI: Whoami,
        PathValues.MODIFY_USER_CONFIGURATION: ModifyUserConfiguration,
        PathValues.DELETE_USERS: DeleteUsers,
        PathValues.UPDATE_USERS: UpdateUsers,
        PathValues.LIST_USERS: ListUsers,
        PathValues.UPLOADFILE: Uploadfile,
        PathValues.UPLOAD_FILE_FROM_URL: UploadFileFromUrl,
        PathValues.UPLOAD_TEXT: UploadText,
        PathValues.DELETE_FILES_V2: DeleteFilesV2,
        PathValues.USER_FILES_V2: UserFilesV2,
        PathValues.CREATE_USER_FILE_TAGS: CreateUserFileTags,
        PathValues.DELETE_USER_FILE_TAGS: DeleteUserFileTags,
        PathValues.RESYNC_FILE: ResyncFile,
        PathValues.MOVE_TO_HOT_STORAGE: MoveToHotStorage,
        PathValues.MODIFY_COLD_STORAGE_PARAMETERS: ModifyColdStorageParameters,
        PathValues.RAW_FILE_FILE_ID: RawFileFileId,
        PathValues.PARSED_FILE_FILE_ID: ParsedFileFileId,
        PathValues.USER_FILES: UserFiles,
        PathValues.DELETEFILE_FILE_ID: DeletefileFileId,
        PathValues.DELETE_FILES: DeleteFiles,
        PathValues.WEBHOOKS: Webhooks,
        PathValues.ADD_WEBHOOK: AddWebhook,
        PathValues.DELETE_WEBHOOK_WEBHOOK_ID: DeleteWebhookWebhookId,
        PathValues.USER_DATA_SOURCES: UserDataSources,
        PathValues.REVOKE_ACCESS_TOKEN: RevokeAccessToken,
        PathValues.DATA_SOURCES_TAGS_ADD: DataSourcesTagsAdd,
        PathValues.DATA_SOURCES_TAGS_REMOVE: DataSourcesTagsRemove,
        PathValues.WEB_SCRAPE: WebScrape,
        PathValues.PROCESS_SITEMAP: ProcessSitemap,
        PathValues.SCRAPE_SITEMAP: ScrapeSitemap,
        PathValues.FETCH_URLS: FetchUrls,
        PathValues.FETCH_WEBPAGE: FetchWebpage,
        PathValues.USER_WEBPAGES: UserWebpages,
        PathValues.SEARCH_URLS: SearchUrls,
        PathValues.FETCH_YOUTUBE_TRANSCRIPT: FetchYoutubeTranscript,
        PathValues.INTEGRATIONS_DATA_CRM_ACCOUNTS: IntegrationsDataCrmAccounts,
        PathValues.INTEGRATIONS_DATA_CRM_ACCOUNTS_ID: IntegrationsDataCrmAccountsId,
        PathValues.INTEGRATIONS_DATA_CRM_LEADS: IntegrationsDataCrmLeads,
        PathValues.INTEGRATIONS_DATA_CRM_LEADS_ID: IntegrationsDataCrmLeadsId,
        PathValues.INTEGRATIONS_DATA_CRM_CONTACTS: IntegrationsDataCrmContacts,
        PathValues.INTEGRATIONS_DATA_CRM_CONTACTS_ID: IntegrationsDataCrmContactsId,
        PathValues.INTEGRATIONS_DATA_CRM_OPPORTUNITIES: IntegrationsDataCrmOpportunities,
        PathValues.INTEGRATIONS_DATA_CRM_OPPORTUNITIES_ID: IntegrationsDataCrmOpportunitiesId,
        PathValues.INTEGRATIONS_DATA_GITHUB_PULL_REQUESTS: IntegrationsDataGithubPullRequests,
        PathValues.INTEGRATIONS_DATA_GITHUB_PULL_REQUESTS_PULL_NUMBER: IntegrationsDataGithubPullRequestsPullNumber,
        PathValues.INTEGRATIONS_DATA_GITHUB_PULL_REQUESTS_COMMENTS: IntegrationsDataGithubPullRequestsComments,
        PathValues.INTEGRATIONS_DATA_GITHUB_PULL_REQUESTS_FILES: IntegrationsDataGithubPullRequestsFiles,
        PathValues.INTEGRATIONS_DATA_GITHUB_PULL_REQUESTS_COMMITS: IntegrationsDataGithubPullRequestsCommits,
        PathValues.INTEGRATIONS_DATA_GITHUB_ISSUES: IntegrationsDataGithubIssues,
        PathValues.INTEGRATIONS_DATA_GITHUB_ISSUES_ISSUE_NUMBER: IntegrationsDataGithubIssuesIssueNumber,
        PathValues.WHITE_LABEL_CREATE: WhiteLabelCreate,
        PathValues.WHITE_LABEL_LIST: WhiteLabelList,
        PathValues.WHITE_LABEL_DELETE: WhiteLabelDelete,
        PathValues.WHITE_LABEL_UPDATE: WhiteLabelUpdate,
    }
)

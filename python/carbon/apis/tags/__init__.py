# do not import all endpoints into this module because that uses a lot of memory and stack frames
# if you need the ability to import all endpoints from this module, import them with
# from carbon.apis.tag_to_api import tag_to_api

import enum


class TagValues(str, enum.Enum):
    INTEGRATIONS = "Integrations"
    FILES = "Files"
    UTILITIES = "Utilities"
    CRM = "CRM"
    GITHUB = "Github"
    USERS = "Users"
    EMBEDDINGS = "Embeddings"
    DATA_SOURCES = "Data Sources"
    WHITE_LABEL = "White Label"
    WEBHOOKS = "Webhooks"
    ORGANIZATIONS = "Organizations"
    AUTH = "Auth"
    GETTING_STARTED = "Getting Started"
    CARBON_CONNECT = "Carbon Connect"
    HEALTH = "Health"
    CONTACT_US = "Contact Us"

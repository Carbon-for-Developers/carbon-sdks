# coding: utf-8

"""
    Carbon

    Connect external data to LLMs, no matter the source.

    The version of the OpenAPI document: 1.0.0
    Generated by: https://konfigthis.com
"""

from carbon.paths.integrations_data_crm_accounts_id.get import GetAccountRaw
from carbon.paths.integrations_data_crm_accounts.post import GetAccountsRaw
from carbon.paths.integrations_data_crm_contacts_id.get import GetContactRaw
from carbon.paths.integrations_data_crm_contacts.post import GetContactsRaw
from carbon.paths.integrations_data_crm_leads_id.get import GetLeadRaw
from carbon.paths.integrations_data_crm_leads.post import GetLeadsRaw
from carbon.paths.integrations_data_crm_opportunities.post import GetOpportunitiesRaw
from carbon.paths.integrations_data_crm_opportunities_id.get import GetOpportunityRaw


class CRMApiRaw(
    GetAccountRaw,
    GetAccountsRaw,
    GetContactRaw,
    GetContactsRaw,
    GetLeadRaw,
    GetLeadsRaw,
    GetOpportunitiesRaw,
    GetOpportunityRaw,
):
    """NOTE:
    This class is auto generated by Konfig (https://konfigthis.com)
    """
    pass

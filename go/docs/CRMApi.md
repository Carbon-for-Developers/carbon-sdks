# APIClient.CRMApi

All URIs are relative to *https://api.carbon.ai*

Method | Path | Description
------------- | ------------- | -------------
[**GetAccount**](CRMApi.md#GetAccount) | **Get** /integrations/data/crm/accounts/{id} | Get Account
[**GetAccounts**](CRMApi.md#GetAccounts) | **Post** /integrations/data/crm/accounts | Get Accounts
[**GetContact**](CRMApi.md#GetContact) | **Get** /integrations/data/crm/contacts/{id} | Get Contact
[**GetContacts**](CRMApi.md#GetContacts) | **Post** /integrations/data/crm/contacts | Get Contacts
[**GetLead**](CRMApi.md#GetLead) | **Get** /integrations/data/crm/leads/{id} | Get Lead
[**GetLeads**](CRMApi.md#GetLeads) | **Post** /integrations/data/crm/leads | Get Leads
[**GetOpportunities**](CRMApi.md#GetOpportunities) | **Post** /integrations/data/crm/opportunities | Get Opportunities
[**GetOpportunity**](CRMApi.md#GetOpportunity) | **Get** /integrations/data/crm/opportunities/{id} | Get Opportunity



## GetAccount

Get Account

### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    request := client.CRMApi.GetAccount(
        "id_example",
        56,
    )
    request.IncludeRemoteData(false)
    request.Includes()
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CRMApi.GetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `GetAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `CRMApi.GetAccount`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `Account.GetAccount.Description`: %v\n", resp.Description)
    fmt.Fprintf(os.Stdout, "Response from `Account.GetAccount.Id`: %v\n", resp.Id)
    fmt.Fprintf(os.Stdout, "Response from `Account.GetAccount.Owner`: %v\n", resp.Owner)
    fmt.Fprintf(os.Stdout, "Response from `Account.GetAccount.Name`: %v\n", resp.Name)
    fmt.Fprintf(os.Stdout, "Response from `Account.GetAccount.Industry`: %v\n", resp.Industry)
    fmt.Fprintf(os.Stdout, "Response from `Account.GetAccount.Website`: %v\n", resp.Website)
    fmt.Fprintf(os.Stdout, "Response from `Account.GetAccount.NumberOfEmployees`: %v\n", resp.NumberOfEmployees)
    fmt.Fprintf(os.Stdout, "Response from `Account.GetAccount.Addresses`: %v\n", resp.Addresses)
    fmt.Fprintf(os.Stdout, "Response from `Account.GetAccount.PhoneNumbers`: %v\n", resp.PhoneNumbers)
    fmt.Fprintf(os.Stdout, "Response from `Account.GetAccount.LastActivityAt`: %v\n", resp.LastActivityAt)
    fmt.Fprintf(os.Stdout, "Response from `Account.GetAccount.CreatedAt`: %v\n", resp.CreatedAt)
    fmt.Fprintf(os.Stdout, "Response from `Account.GetAccount.UpdatedAt`: %v\n", resp.UpdatedAt)
    fmt.Fprintf(os.Stdout, "Response from `Account.GetAccount.IsDeleted`: %v\n", resp.IsDeleted)
    fmt.Fprintf(os.Stdout, "Response from `Account.GetAccount.Tasks`: %v\n", *resp.Tasks)
    fmt.Fprintf(os.Stdout, "Response from `Account.GetAccount.Events`: %v\n", *resp.Events)
    fmt.Fprintf(os.Stdout, "Response from `Account.GetAccount.RemoteData`: %v\n", resp.RemoteData)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccounts

Get Accounts

### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    filters := *carbon.NewAccountFilters()
    
    accountsRequest := *carbon.NewAccountsRequest(
        null,
    )
    accountsRequest.SetIncludeRemoteData(false)
    accountsRequest.SetNextCursor("null")
    accountsRequest.SetPageSize(null)
    accountsRequest.SetOrderDir(null)
    accountsRequest.SetIncludes([])
    accountsRequest.SetFilters(filters)
    accountsRequest.SetOrderBy(null)
    
    request := client.CRMApi.GetAccounts(
        accountsRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CRMApi.GetAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `GetAccounts`: AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `CRMApi.GetAccounts`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `AccountResponse.GetAccounts.Count`: %v\n", resp.Count)
    fmt.Fprintf(os.Stdout, "Response from `AccountResponse.GetAccounts.NextCursor`: %v\n", resp.NextCursor)
    fmt.Fprintf(os.Stdout, "Response from `AccountResponse.GetAccounts.Data`: %v\n", resp.Data)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContact

Get Contact

### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    request := client.CRMApi.GetContact(
        "id_example",
        56,
    )
    request.IncludeRemoteData(false)
    request.Includes()
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CRMApi.GetContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `GetContact`: Contact
    fmt.Fprintf(os.Stdout, "Response from `CRMApi.GetContact`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `Contact.GetContact.Title`: %v\n", resp.Title)
    fmt.Fprintf(os.Stdout, "Response from `Contact.GetContact.Description`: %v\n", resp.Description)
    fmt.Fprintf(os.Stdout, "Response from `Contact.GetContact.Id`: %v\n", resp.Id)
    fmt.Fprintf(os.Stdout, "Response from `Contact.GetContact.Owner`: %v\n", resp.Owner)
    fmt.Fprintf(os.Stdout, "Response from `Contact.GetContact.FirstName`: %v\n", resp.FirstName)
    fmt.Fprintf(os.Stdout, "Response from `Contact.GetContact.LastName`: %v\n", resp.LastName)
    fmt.Fprintf(os.Stdout, "Response from `Contact.GetContact.Name`: %v\n", resp.Name)
    fmt.Fprintf(os.Stdout, "Response from `Contact.GetContact.Department`: %v\n", resp.Department)
    fmt.Fprintf(os.Stdout, "Response from `Contact.GetContact.Addresses`: %v\n", resp.Addresses)
    fmt.Fprintf(os.Stdout, "Response from `Contact.GetContact.PhoneNumbers`: %v\n", resp.PhoneNumbers)
    fmt.Fprintf(os.Stdout, "Response from `Contact.GetContact.Emails`: %v\n", resp.Emails)
    fmt.Fprintf(os.Stdout, "Response from `Contact.GetContact.Account`: %v\n", resp.Account)
    fmt.Fprintf(os.Stdout, "Response from `Contact.GetContact.LastActivityAt`: %v\n", resp.LastActivityAt)
    fmt.Fprintf(os.Stdout, "Response from `Contact.GetContact.CreatedAt`: %v\n", resp.CreatedAt)
    fmt.Fprintf(os.Stdout, "Response from `Contact.GetContact.UpdatedAt`: %v\n", resp.UpdatedAt)
    fmt.Fprintf(os.Stdout, "Response from `Contact.GetContact.IsDeleted`: %v\n", resp.IsDeleted)
    fmt.Fprintf(os.Stdout, "Response from `Contact.GetContact.Tasks`: %v\n", *resp.Tasks)
    fmt.Fprintf(os.Stdout, "Response from `Contact.GetContact.Events`: %v\n", *resp.Events)
    fmt.Fprintf(os.Stdout, "Response from `Contact.GetContact.RemoteData`: %v\n", resp.RemoteData)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContacts

Get Contacts

### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    filters := *carbon.NewContactFilters()
    
    contactsRequest := *carbon.NewContactsRequest(
        null,
    )
    contactsRequest.SetIncludeRemoteData(false)
    contactsRequest.SetNextCursor("null")
    contactsRequest.SetPageSize(null)
    contactsRequest.SetOrderDir(null)
    contactsRequest.SetIncludes([])
    contactsRequest.SetFilters(filters)
    contactsRequest.SetOrderBy(null)
    
    request := client.CRMApi.GetContacts(
        contactsRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CRMApi.GetContacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `GetContacts`: ContactsResponse
    fmt.Fprintf(os.Stdout, "Response from `CRMApi.GetContacts`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `ContactsResponse.GetContacts.Count`: %v\n", resp.Count)
    fmt.Fprintf(os.Stdout, "Response from `ContactsResponse.GetContacts.NextCursor`: %v\n", resp.NextCursor)
    fmt.Fprintf(os.Stdout, "Response from `ContactsResponse.GetContacts.Data`: %v\n", resp.Data)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLead

Get Lead

### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    request := client.CRMApi.GetLead(
        "id_example",
        56,
    )
    request.IncludeRemoteData(false)
    request.Includes()
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CRMApi.GetLead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `GetLead`: Lead
    fmt.Fprintf(os.Stdout, "Response from `CRMApi.GetLead`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `Lead.GetLead.Title`: %v\n", resp.Title)
    fmt.Fprintf(os.Stdout, "Response from `Lead.GetLead.Description`: %v\n", resp.Description)
    fmt.Fprintf(os.Stdout, "Response from `Lead.GetLead.Id`: %v\n", resp.Id)
    fmt.Fprintf(os.Stdout, "Response from `Lead.GetLead.Owner`: %v\n", resp.Owner)
    fmt.Fprintf(os.Stdout, "Response from `Lead.GetLead.Source`: %v\n", resp.Source)
    fmt.Fprintf(os.Stdout, "Response from `Lead.GetLead.Status`: %v\n", resp.Status)
    fmt.Fprintf(os.Stdout, "Response from `Lead.GetLead.Company`: %v\n", resp.Company)
    fmt.Fprintf(os.Stdout, "Response from `Lead.GetLead.FirstName`: %v\n", resp.FirstName)
    fmt.Fprintf(os.Stdout, "Response from `Lead.GetLead.LastName`: %v\n", resp.LastName)
    fmt.Fprintf(os.Stdout, "Response from `Lead.GetLead.Addresses`: %v\n", resp.Addresses)
    fmt.Fprintf(os.Stdout, "Response from `Lead.GetLead.PhoneNumbers`: %v\n", resp.PhoneNumbers)
    fmt.Fprintf(os.Stdout, "Response from `Lead.GetLead.Emails`: %v\n", resp.Emails)
    fmt.Fprintf(os.Stdout, "Response from `Lead.GetLead.ConvertedAt`: %v\n", resp.ConvertedAt)
    fmt.Fprintf(os.Stdout, "Response from `Lead.GetLead.ConvertedAccount`: %v\n", resp.ConvertedAccount)
    fmt.Fprintf(os.Stdout, "Response from `Lead.GetLead.ConvertedContact`: %v\n", resp.ConvertedContact)
    fmt.Fprintf(os.Stdout, "Response from `Lead.GetLead.LastActivityAt`: %v\n", resp.LastActivityAt)
    fmt.Fprintf(os.Stdout, "Response from `Lead.GetLead.CreatedAt`: %v\n", resp.CreatedAt)
    fmt.Fprintf(os.Stdout, "Response from `Lead.GetLead.UpdatedAt`: %v\n", resp.UpdatedAt)
    fmt.Fprintf(os.Stdout, "Response from `Lead.GetLead.IsDeleted`: %v\n", resp.IsDeleted)
    fmt.Fprintf(os.Stdout, "Response from `Lead.GetLead.Tasks`: %v\n", *resp.Tasks)
    fmt.Fprintf(os.Stdout, "Response from `Lead.GetLead.Events`: %v\n", *resp.Events)
    fmt.Fprintf(os.Stdout, "Response from `Lead.GetLead.RemoteData`: %v\n", resp.RemoteData)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLeads

Get Leads

### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    filters := *carbon.NewLeadFilters()
    
    leadsRequest := *carbon.NewLeadsRequest(
        null,
    )
    leadsRequest.SetIncludeRemoteData(false)
    leadsRequest.SetNextCursor("null")
    leadsRequest.SetPageSize(null)
    leadsRequest.SetOrderDir(null)
    leadsRequest.SetIncludes([])
    leadsRequest.SetFilters(filters)
    leadsRequest.SetOrderBy(null)
    
    request := client.CRMApi.GetLeads(
        leadsRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CRMApi.GetLeads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `GetLeads`: LeadsResponse
    fmt.Fprintf(os.Stdout, "Response from `CRMApi.GetLeads`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `LeadsResponse.GetLeads.Count`: %v\n", resp.Count)
    fmt.Fprintf(os.Stdout, "Response from `LeadsResponse.GetLeads.NextCursor`: %v\n", resp.NextCursor)
    fmt.Fprintf(os.Stdout, "Response from `LeadsResponse.GetLeads.Data`: %v\n", resp.Data)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpportunities

Get Opportunities

### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    filters := *carbon.NewOpportunityFilters()
    
    opportunitiesRequest := *carbon.NewOpportunitiesRequest(
        null,
    )
    opportunitiesRequest.SetIncludeRemoteData(false)
    opportunitiesRequest.SetNextCursor("null")
    opportunitiesRequest.SetPageSize(null)
    opportunitiesRequest.SetOrderDir(null)
    opportunitiesRequest.SetIncludes([])
    opportunitiesRequest.SetFilters(filters)
    opportunitiesRequest.SetOrderBy(null)
    
    request := client.CRMApi.GetOpportunities(
        opportunitiesRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CRMApi.GetOpportunities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `GetOpportunities`: OpportunitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `CRMApi.GetOpportunities`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `OpportunitiesResponse.GetOpportunities.Count`: %v\n", resp.Count)
    fmt.Fprintf(os.Stdout, "Response from `OpportunitiesResponse.GetOpportunities.NextCursor`: %v\n", resp.NextCursor)
    fmt.Fprintf(os.Stdout, "Response from `OpportunitiesResponse.GetOpportunities.Data`: %v\n", resp.Data)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpportunity

Get Opportunity

### Example

```go
package main

import (
    "fmt"
    "os"
    carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func main() {
    configuration := carbon.NewConfiguration()
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)

    request := client.CRMApi.GetOpportunity(
        "id_example",
        56,
    )
    request.IncludeRemoteData(false)
    request.Includes()
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CRMApi.GetOpportunity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `GetOpportunity`: Opportunity
    fmt.Fprintf(os.Stdout, "Response from `CRMApi.GetOpportunity`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `Opportunity.GetOpportunity.Description`: %v\n", resp.Description)
    fmt.Fprintf(os.Stdout, "Response from `Opportunity.GetOpportunity.Id`: %v\n", resp.Id)
    fmt.Fprintf(os.Stdout, "Response from `Opportunity.GetOpportunity.Owner`: %v\n", resp.Owner)
    fmt.Fprintf(os.Stdout, "Response from `Opportunity.GetOpportunity.Name`: %v\n", resp.Name)
    fmt.Fprintf(os.Stdout, "Response from `Opportunity.GetOpportunity.Amount`: %v\n", resp.Amount)
    fmt.Fprintf(os.Stdout, "Response from `Opportunity.GetOpportunity.Account`: %v\n", resp.Account)
    fmt.Fprintf(os.Stdout, "Response from `Opportunity.GetOpportunity.Contact`: %v\n", resp.Contact)
    fmt.Fprintf(os.Stdout, "Response from `Opportunity.GetOpportunity.Stage`: %v\n", resp.Stage)
    fmt.Fprintf(os.Stdout, "Response from `Opportunity.GetOpportunity.Status`: %v\n", resp.Status)
    fmt.Fprintf(os.Stdout, "Response from `Opportunity.GetOpportunity.CloseDate`: %v\n", resp.CloseDate)
    fmt.Fprintf(os.Stdout, "Response from `Opportunity.GetOpportunity.LastActivityAt`: %v\n", resp.LastActivityAt)
    fmt.Fprintf(os.Stdout, "Response from `Opportunity.GetOpportunity.CreatedAt`: %v\n", resp.CreatedAt)
    fmt.Fprintf(os.Stdout, "Response from `Opportunity.GetOpportunity.UpdatedAt`: %v\n", resp.UpdatedAt)
    fmt.Fprintf(os.Stdout, "Response from `Opportunity.GetOpportunity.IsDeleted`: %v\n", resp.IsDeleted)
    fmt.Fprintf(os.Stdout, "Response from `Opportunity.GetOpportunity.Tasks`: %v\n", *resp.Tasks)
    fmt.Fprintf(os.Stdout, "Response from `Opportunity.GetOpportunity.Events`: %v\n", *resp.Events)
    fmt.Fprintf(os.Stdout, "Response from `Opportunity.GetOpportunity.RemoteData`: %v\n", resp.RemoteData)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


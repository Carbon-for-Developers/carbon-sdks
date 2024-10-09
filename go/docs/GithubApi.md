# APIClient.GithubApi

All URIs are relative to *https://api.carbon.ai*

Method | Path | Description
------------- | ------------- | -------------
[**GetIssue**](GithubApi.md#GetIssue) | **Get** /integrations/data/github/issues/{issue_number} | Issue
[**GetIssues**](GithubApi.md#GetIssues) | **Post** /integrations/data/github/issues | Issues
[**GetPr**](GithubApi.md#GetPr) | **Get** /integrations/data/github/pull_requests/{pull_number} | Get Pr
[**GetPrComments**](GithubApi.md#GetPrComments) | **Post** /integrations/data/github/pull_requests/comments | Pr Comments
[**GetPrCommits**](GithubApi.md#GetPrCommits) | **Post** /integrations/data/github/pull_requests/commits | Pr Commits
[**GetPrFiles**](GithubApi.md#GetPrFiles) | **Post** /integrations/data/github/pull_requests/files | Pr Files
[**GetPullRequests**](GithubApi.md#GetPullRequests) | **Post** /integrations/data/github/pull_requests | Get Prs



## GetIssue

Issue

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

    request := client.GithubApi.GetIssue(
        56,
    )
    request.IncludeRemoteData(false)
    request.DataSourceId(56)
    request.Repository("repository_example")
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GithubApi.GetIssue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `GetIssue`: Issue
    fmt.Fprintf(os.Stdout, "Response from `GithubApi.GetIssue`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `Issue.GetIssue.Title`: %v\n", resp.Title)
    fmt.Fprintf(os.Stdout, "Response from `Issue.GetIssue.Id`: %v\n", resp.Id)
    fmt.Fprintf(os.Stdout, "Response from `Issue.GetIssue.Url`: %v\n", resp.Url)
    fmt.Fprintf(os.Stdout, "Response from `Issue.GetIssue.HtmlUrl`: %v\n", resp.HtmlUrl)
    fmt.Fprintf(os.Stdout, "Response from `Issue.GetIssue.Number`: %v\n", resp.Number)
    fmt.Fprintf(os.Stdout, "Response from `Issue.GetIssue.User`: %v\n", resp.User)
    fmt.Fprintf(os.Stdout, "Response from `Issue.GetIssue.Labels`: %v\n", resp.Labels)
    fmt.Fprintf(os.Stdout, "Response from `Issue.GetIssue.State`: %v\n", resp.State)
    fmt.Fprintf(os.Stdout, "Response from `Issue.GetIssue.Locked`: %v\n", resp.Locked)
    fmt.Fprintf(os.Stdout, "Response from `Issue.GetIssue.NumComments`: %v\n", resp.NumComments)
    fmt.Fprintf(os.Stdout, "Response from `Issue.GetIssue.CreatedAt`: %v\n", resp.CreatedAt)
    fmt.Fprintf(os.Stdout, "Response from `Issue.GetIssue.UpdatedAt`: %v\n", resp.UpdatedAt)
    fmt.Fprintf(os.Stdout, "Response from `Issue.GetIssue.ClosedAt`: %v\n", resp.ClosedAt)
    fmt.Fprintf(os.Stdout, "Response from `Issue.GetIssue.Draft`: %v\n", resp.Draft)
    fmt.Fprintf(os.Stdout, "Response from `Issue.GetIssue.PullRequest`: %v\n", resp.PullRequest)
    fmt.Fprintf(os.Stdout, "Response from `Issue.GetIssue.Body`: %v\n", resp.Body)
    fmt.Fprintf(os.Stdout, "Response from `Issue.GetIssue.ClosedBy`: %v\n", resp.ClosedBy)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssues

Issues

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

    filters := *carbon.NewIssuesFilter()
    
    issuesInput := *carbon.NewIssuesInput(
        null,
        "null",
    )
    issuesInput.SetIncludeRemoteData(false)
    issuesInput.SetPage(1)
    issuesInput.SetPageSize(30)
    issuesInput.SetNextCursor("null")
    issuesInput.SetFilters(filters)
    issuesInput.SetOrderBy(null)
    issuesInput.SetOrderDir(null)
    
    request := client.GithubApi.GetIssues(
        issuesInput,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GithubApi.GetIssues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `GetIssues`: IssuesResponse
    fmt.Fprintf(os.Stdout, "Response from `GithubApi.GetIssues`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `IssuesResponse.GetIssues.Data`: %v\n", resp.Data)
    fmt.Fprintf(os.Stdout, "Response from `IssuesResponse.GetIssues.NextCursor`: %v\n", resp.NextCursor)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPr

Get Pr

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

    request := client.GithubApi.GetPr(
        56,
    )
    request.IncludeRemoteData(false)
    request.DataSourceId(56)
    request.Repository("repository_example")
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GithubApi.GetPr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `GetPr`: PullRequestExtended
    fmt.Fprintf(os.Stdout, "Response from `GithubApi.GetPr`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestExtended.GetPr.Title`: %v\n", resp.Title)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestExtended.GetPr.Id`: %v\n", resp.Id)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestExtended.GetPr.Url`: %v\n", resp.Url)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestExtended.GetPr.Number`: %v\n", resp.Number)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestExtended.GetPr.State`: %v\n", resp.State)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestExtended.GetPr.User`: %v\n", resp.User)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestExtended.GetPr.CreatedAt`: %v\n", resp.CreatedAt)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestExtended.GetPr.UpdatedAt`: %v\n", resp.UpdatedAt)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestExtended.GetPr.ClosedAt`: %v\n", resp.ClosedAt)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestExtended.GetPr.MergedAt`: %v\n", resp.MergedAt)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestExtended.GetPr.RequestedReviewers`: %v\n", resp.RequestedReviewers)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestExtended.GetPr.RequestedTeams`: %v\n", resp.RequestedTeams)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestExtended.GetPr.Labels`: %v\n", resp.Labels)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestExtended.GetPr.Draft`: %v\n", resp.Draft)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestExtended.GetPr.Head`: %v\n", resp.Head)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestExtended.GetPr.Base`: %v\n", resp.Base)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestExtended.GetPr.RemoteData`: %v\n", resp.RemoteData)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestExtended.GetPr.Merged`: %v\n", resp.Merged)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestExtended.GetPr.NumComments`: %v\n", resp.NumComments)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestExtended.GetPr.NumReviewComments`: %v\n", resp.NumReviewComments)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestExtended.GetPr.NumCommits`: %v\n", resp.NumCommits)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestExtended.GetPr.NumAdditions`: %v\n", resp.NumAdditions)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestExtended.GetPr.NumDeletions`: %v\n", resp.NumDeletions)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestExtended.GetPr.NumChangedFiles`: %v\n", resp.NumChangedFiles)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrComments

Pr Comments

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

    
    commentsInput := *carbon.NewCommentsInput(
        null,
        "null",
        null,
    )
    commentsInput.SetIncludeRemoteData(false)
    commentsInput.SetPage(1)
    commentsInput.SetPageSize(30)
    commentsInput.SetNextCursor("null")
    commentsInput.SetOrderBy(null)
    commentsInput.SetOrderDir(null)
    
    request := client.GithubApi.GetPrComments(
        commentsInput,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GithubApi.GetPrComments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `GetPrComments`: CommentsResponse
    fmt.Fprintf(os.Stdout, "Response from `GithubApi.GetPrComments`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `CommentsResponse.GetPrComments.Data`: %v\n", resp.Data)
    fmt.Fprintf(os.Stdout, "Response from `CommentsResponse.GetPrComments.NextCursor`: %v\n", resp.NextCursor)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrCommits

Pr Commits

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

    
    commitsInput := *carbon.NewCommitsInput(
        null,
        "null",
        null,
    )
    commitsInput.SetIncludeRemoteData(false)
    commitsInput.SetPage(1)
    commitsInput.SetPageSize(30)
    commitsInput.SetNextCursor("null")
    
    request := client.GithubApi.GetPrCommits(
        commitsInput,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GithubApi.GetPrCommits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `GetPrCommits`: CommitsResponse
    fmt.Fprintf(os.Stdout, "Response from `GithubApi.GetPrCommits`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `CommitsResponse.GetPrCommits.Data`: %v\n", resp.Data)
    fmt.Fprintf(os.Stdout, "Response from `CommitsResponse.GetPrCommits.NextCursor`: %v\n", resp.NextCursor)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrFiles

Pr Files

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

    
    filesInput := *carbon.NewFilesInput(
        null,
        "null",
        null,
    )
    filesInput.SetIncludeRemoteData(false)
    filesInput.SetPage(1)
    filesInput.SetPageSize(30)
    filesInput.SetNextCursor("null")
    
    request := client.GithubApi.GetPrFiles(
        filesInput,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GithubApi.GetPrFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `GetPrFiles`: FilesResponse
    fmt.Fprintf(os.Stdout, "Response from `GithubApi.GetPrFiles`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `FilesResponse.GetPrFiles.Data`: %v\n", resp.Data)
    fmt.Fprintf(os.Stdout, "Response from `FilesResponse.GetPrFiles.NextCursor`: %v\n", resp.NextCursor)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPullRequests

Get Prs

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

    filters := *carbon.NewPullRequestFilters()
    
    pullRequestsInput := *carbon.NewPullRequestsInput(
        null,
        "null",
    )
    pullRequestsInput.SetIncludeRemoteData(false)
    pullRequestsInput.SetPage(1)
    pullRequestsInput.SetPageSize(30)
    pullRequestsInput.SetNextCursor("null")
    pullRequestsInput.SetFilters(filters)
    pullRequestsInput.SetOrderBy(null)
    pullRequestsInput.SetOrderDir(null)
    
    request := client.GithubApi.GetPullRequests(
        pullRequestsInput,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GithubApi.GetPullRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `GetPullRequests`: PullRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `GithubApi.GetPullRequests`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestResponse.GetPullRequests.Data`: %v\n", resp.Data)
    fmt.Fprintf(os.Stdout, "Response from `PullRequestResponse.GetPullRequests.NextCursor`: %v\n", resp.NextCursor)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


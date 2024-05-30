/*
Carbon

Testing IntegrationsApiService

*/

// Code generated by Konfig (https://konfigthis.com);

package carbon

import (
    "testing"
    // "github.com/stretchr/testify/assert"
    // "github.com/stretchr/testify/require"
    // carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func Test_carbon_IntegrationsApiService(t *testing.T) {

    // configuration := carbon.NewConfiguration()
    // configuration.SetHost("http://127.0.0.1:4010")
    /* 
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)
    */

    t.Run("Test IntegrationsApiService ConnectDataSource", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        authentication := *carbon.NewAuthenticationProperty()
        syncOptions := *carbon.NewSyncOptions()
        
        connectDataSourceInput := *carbon.NewConnectDataSourceInput(
            authentication,
        )
        connectDataSourceInput.SetSyncOptions(syncOptions)
        
        request := client.IntegrationsApi.ConnectDataSource(
            connectDataSourceInput,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test IntegrationsApiService ConnectFreshdesk", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        fileSyncConfig := *carbon.NewFileSyncConfigNullable()
        
        freshDeskConnectRequest := *carbon.NewFreshDeskConnectRequest(
            "null",
            "null",
        )
        freshDeskConnectRequest.SetTags({})
        freshDeskConnectRequest.SetChunkSize(1500)
        freshDeskConnectRequest.SetChunkOverlap(20)
        freshDeskConnectRequest.SetSkipEmbeddingGeneration(false)
        freshDeskConnectRequest.SetEmbeddingModel(null)
        freshDeskConnectRequest.SetGenerateSparseVectors(false)
        freshDeskConnectRequest.SetPrependFilenameToChunks(false)
        freshDeskConnectRequest.SetSyncFilesOnConnection(true)
        freshDeskConnectRequest.SetRequestId("null")
        freshDeskConnectRequest.SetSyncSourceItems(true)
        freshDeskConnectRequest.SetFileSyncConfig(fileSyncConfig)
        
        request := client.IntegrationsApi.ConnectFreshdesk(
            freshDeskConnectRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test IntegrationsApiService ConnectGitbook", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        
        gitbookConnectRequest := *carbon.NewGitbookConnectRequest(
            "null",
            "null",
        )
        gitbookConnectRequest.SetTags({})
        gitbookConnectRequest.SetChunkSize(1500)
        gitbookConnectRequest.SetChunkOverlap(20)
        gitbookConnectRequest.SetSkipEmbeddingGeneration(false)
        gitbookConnectRequest.SetEmbeddingModel(null)
        gitbookConnectRequest.SetGenerateSparseVectors(false)
        gitbookConnectRequest.SetPrependFilenameToChunks(false)
        gitbookConnectRequest.SetSyncFilesOnConnection(true)
        gitbookConnectRequest.SetRequestId("null")
        gitbookConnectRequest.SetSyncSourceItems(true)
        
        request := client.IntegrationsApi.ConnectGitbook(
            gitbookConnectRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test IntegrationsApiService CreateAwsIamUser", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        
        s3AuthRequest := *carbon.NewS3AuthRequest(
            "null",
            "null",
        )
        s3AuthRequest.SetSyncSourceItems(true)
        
        request := client.IntegrationsApi.CreateAwsIamUser(
            s3AuthRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test IntegrationsApiService GetOauthUrl", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        fileSyncConfig := *carbon.NewFileSyncConfigNullable()
        
        oAuthURLRequest := *carbon.NewOAuthURLRequest(
            null,
        )
        oAuthURLRequest.SetTags(null)
        oAuthURLRequest.SetScope("null")
        oAuthURLRequest.SetChunkSize(1500)
        oAuthURLRequest.SetChunkOverlap(20)
        oAuthURLRequest.SetSkipEmbeddingGeneration(false)
        oAuthURLRequest.SetEmbeddingModel(null)
        oAuthURLRequest.SetZendeskSubdomain("null")
        oAuthURLRequest.SetMicrosoftTenant("null")
        oAuthURLRequest.SetSharepointSiteName("null")
        oAuthURLRequest.SetConfluenceSubdomain("null")
        oAuthURLRequest.SetGenerateSparseVectors(false)
        oAuthURLRequest.SetPrependFilenameToChunks(false)
        oAuthURLRequest.SetMaxItemsPerChunk(null)
        oAuthURLRequest.SetSalesforceDomain("null")
        oAuthURLRequest.SetSyncFilesOnConnection(true)
        oAuthURLRequest.SetSetPageAsBoundary(false)
        oAuthURLRequest.SetDataSourceId(null)
        oAuthURLRequest.SetConnectingNewAccount(false)
        oAuthURLRequest.SetRequestId("ce1b1ec8-be64-491c-9159-c40f85fa0073")
        oAuthURLRequest.SetUseOcr(false)
        oAuthURLRequest.SetParsePdfTablesWithOcr(false)
        oAuthURLRequest.SetEnableFilePicker(true)
        oAuthURLRequest.SetSyncSourceItems(true)
        oAuthURLRequest.SetIncrementalSync(false)
        oAuthURLRequest.SetFileSyncConfig(fileSyncConfig)
        
        request := client.IntegrationsApi.GetOauthUrl(
            oAuthURLRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test IntegrationsApiService ListConfluencePages", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        
        listRequest := *carbon.NewListRequest(
            null,
        )
        listRequest.SetParentId("null")
        
        request := client.IntegrationsApi.ListConfluencePages(
            listRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test IntegrationsApiService ListDataSourceItems", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        filters := *carbon.NewListItemsFiltersNullable()
        pagination := *carbon.NewPagination()
        
        listDataSourceItemsRequest := *carbon.NewListDataSourceItemsRequest(
            null,
        )
        listDataSourceItemsRequest.SetParentId("null")
        listDataSourceItemsRequest.SetFilters(filters)
        listDataSourceItemsRequest.SetPagination(pagination)
        listDataSourceItemsRequest.SetOrderBy(null)
        listDataSourceItemsRequest.SetOrderDir(null)
        
        request := client.IntegrationsApi.ListDataSourceItems(
            listDataSourceItemsRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test IntegrationsApiService ListFolders", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        request := client.IntegrationsApi.ListFolders(
        )
        request.DataSourceId(56)
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test IntegrationsApiService ListGitbookSpaces", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        request := client.IntegrationsApi.ListGitbookSpaces(
            56,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test IntegrationsApiService ListLabels", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        request := client.IntegrationsApi.ListLabels(
        )
        request.DataSourceId(56)
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test IntegrationsApiService ListOutlookCategories", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        request := client.IntegrationsApi.ListOutlookCategories(
        )
        request.DataSourceId(56)
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test IntegrationsApiService ListRepos", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        request := client.IntegrationsApi.ListRepos(
        )
        request.PerPage(30)
        request.Page(1)
        request.DataSourceId(56)
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test IntegrationsApiService SyncConfluence", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        ids := *carbon.NewIdsProperty()
        fileSyncConfig := *carbon.NewFileSyncConfigNullable()
        
        syncFilesRequest := *carbon.NewSyncFilesRequest(
            null,
            ids,
        )
        syncFilesRequest.SetTags({})
        syncFilesRequest.SetChunkSize(1500)
        syncFilesRequest.SetChunkOverlap(20)
        syncFilesRequest.SetSkipEmbeddingGeneration(false)
        syncFilesRequest.SetEmbeddingModel(null)
        syncFilesRequest.SetGenerateSparseVectors(false)
        syncFilesRequest.SetPrependFilenameToChunks(false)
        syncFilesRequest.SetMaxItemsPerChunk(null)
        syncFilesRequest.SetSetPageAsBoundary(false)
        syncFilesRequest.SetRequestId("9fe9190e-384f-4baa-a416-d51ed93d1be7")
        syncFilesRequest.SetUseOcr(false)
        syncFilesRequest.SetParsePdfTablesWithOcr(false)
        syncFilesRequest.SetIncrementalSync(false)
        syncFilesRequest.SetFileSyncConfig(fileSyncConfig)
        
        request := client.IntegrationsApi.SyncConfluence(
            syncFilesRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test IntegrationsApiService SyncDataSourceItems", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        
        syncDirectoryRequest := *carbon.NewSyncDirectoryRequest(
            null,
        )
        
        request := client.IntegrationsApi.SyncDataSourceItems(
            syncDirectoryRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test IntegrationsApiService SyncFiles", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        ids := *carbon.NewIdsProperty()
        fileSyncConfig := *carbon.NewFileSyncConfigNullable()
        
        syncFilesRequest := *carbon.NewSyncFilesRequest(
            null,
            ids,
        )
        syncFilesRequest.SetTags({})
        syncFilesRequest.SetChunkSize(1500)
        syncFilesRequest.SetChunkOverlap(20)
        syncFilesRequest.SetSkipEmbeddingGeneration(false)
        syncFilesRequest.SetEmbeddingModel(null)
        syncFilesRequest.SetGenerateSparseVectors(false)
        syncFilesRequest.SetPrependFilenameToChunks(false)
        syncFilesRequest.SetMaxItemsPerChunk(null)
        syncFilesRequest.SetSetPageAsBoundary(false)
        syncFilesRequest.SetRequestId("9fe9190e-384f-4baa-a416-d51ed93d1be7")
        syncFilesRequest.SetUseOcr(false)
        syncFilesRequest.SetParsePdfTablesWithOcr(false)
        syncFilesRequest.SetIncrementalSync(false)
        syncFilesRequest.SetFileSyncConfig(fileSyncConfig)
        
        request := client.IntegrationsApi.SyncFiles(
            syncFilesRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test IntegrationsApiService SyncGitHub", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        
        githubConnectRequest := *carbon.NewGithubConnectRequest(
            "null",
            "null",
        )
        githubConnectRequest.SetSyncSourceItems(false)
        
        request := client.IntegrationsApi.SyncGitHub(
            githubConnectRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test IntegrationsApiService SyncGitbook", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        
        gitbookSyncRequest := *carbon.NewGitbookSyncRequest(
            null,
            null,
        )
        gitbookSyncRequest.SetTags({})
        gitbookSyncRequest.SetChunkSize(1500)
        gitbookSyncRequest.SetChunkOverlap(20)
        gitbookSyncRequest.SetSkipEmbeddingGeneration(false)
        gitbookSyncRequest.SetEmbeddingModel(null)
        gitbookSyncRequest.SetGenerateSparseVectors(false)
        gitbookSyncRequest.SetPrependFilenameToChunks(false)
        gitbookSyncRequest.SetRequestId("null")
        
        request := client.IntegrationsApi.SyncGitbook(
            gitbookSyncRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test IntegrationsApiService SyncGmail", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        fileSyncConfig := *carbon.NewFileSyncConfigNullable()
        
        gmailSyncInput := *carbon.NewGmailSyncInput(
            null,
        )
        gmailSyncInput.SetTags({})
        gmailSyncInput.SetChunkSize(1500)
        gmailSyncInput.SetChunkOverlap(20)
        gmailSyncInput.SetSkipEmbeddingGeneration(false)
        gmailSyncInput.SetEmbeddingModel(null)
        gmailSyncInput.SetGenerateSparseVectors(false)
        gmailSyncInput.SetPrependFilenameToChunks(false)
        gmailSyncInput.SetDataSourceId(null)
        gmailSyncInput.SetRequestId("null")
        gmailSyncInput.SetSyncAttachments(false)
        gmailSyncInput.SetFileSyncConfig(fileSyncConfig)
        gmailSyncInput.SetIncrementalSync(false)
        
        request := client.IntegrationsApi.SyncGmail(
            gmailSyncInput,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test IntegrationsApiService SyncOutlook", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        fileSyncConfig := *carbon.NewFileSyncConfigNullable()
        
        outlookSyncInput := *carbon.NewOutlookSyncInput(
            null,
        )
        outlookSyncInput.SetTags({})
        outlookSyncInput.SetFolder("Inbox")
        outlookSyncInput.SetChunkSize(1500)
        outlookSyncInput.SetChunkOverlap(20)
        outlookSyncInput.SetSkipEmbeddingGeneration(false)
        outlookSyncInput.SetEmbeddingModel(null)
        outlookSyncInput.SetGenerateSparseVectors(false)
        outlookSyncInput.SetPrependFilenameToChunks(false)
        outlookSyncInput.SetDataSourceId(null)
        outlookSyncInput.SetRequestId("null")
        outlookSyncInput.SetSyncAttachments(false)
        outlookSyncInput.SetFileSyncConfig(fileSyncConfig)
        outlookSyncInput.SetIncrementalSync(false)
        
        request := client.IntegrationsApi.SyncOutlook(
            outlookSyncInput,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test IntegrationsApiService SyncRepos", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        
        githubFetchReposRequest := *carbon.NewGithubFetchReposRequest(
            null,
        )
        githubFetchReposRequest.SetDataSourceId(null)
        
        request := client.IntegrationsApi.SyncRepos(
            githubFetchReposRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test IntegrationsApiService SyncRssFeed", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        
        rSSFeedInput := *carbon.NewRSSFeedInput(
            "null",
        )
        rSSFeedInput.SetTags({})
        rSSFeedInput.SetChunkSize(1500)
        rSSFeedInput.SetChunkOverlap(20)
        rSSFeedInput.SetSkipEmbeddingGeneration(false)
        rSSFeedInput.SetEmbeddingModel(null)
        rSSFeedInput.SetGenerateSparseVectors(false)
        rSSFeedInput.SetPrependFilenameToChunks(false)
        rSSFeedInput.SetRequestId("null")
        
        request := client.IntegrationsApi.SyncRssFeed(
            rSSFeedInput,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test IntegrationsApiService SyncS3Files", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        fileSyncConfig := *carbon.NewFileSyncConfigNullable()
        
        s3FileSyncInput := *carbon.NewS3FileSyncInput(
            null,
        )
        s3FileSyncInput.SetTags({})
        s3FileSyncInput.SetChunkSize(1500)
        s3FileSyncInput.SetChunkOverlap(20)
        s3FileSyncInput.SetSkipEmbeddingGeneration(false)
        s3FileSyncInput.SetEmbeddingModel(null)
        s3FileSyncInput.SetGenerateSparseVectors(false)
        s3FileSyncInput.SetPrependFilenameToChunks(false)
        s3FileSyncInput.SetMaxItemsPerChunk(null)
        s3FileSyncInput.SetSetPageAsBoundary(false)
        s3FileSyncInput.SetDataSourceId(null)
        s3FileSyncInput.SetRequestId("null")
        s3FileSyncInput.SetUseOcr(false)
        s3FileSyncInput.SetParsePdfTablesWithOcr(false)
        s3FileSyncInput.SetFileSyncConfig(fileSyncConfig)
        
        request := client.IntegrationsApi.SyncS3Files(
            s3FileSyncInput,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

}
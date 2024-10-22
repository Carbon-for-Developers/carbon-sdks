/*
Carbon

Testing DataSourcesApiService

*/

// Code generated by Konfig (https://konfigthis.com);

package carbon

import (
    "testing"
    // "github.com/stretchr/testify/assert"
    // "github.com/stretchr/testify/require"
    // carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func Test_carbon_DataSourcesApiService(t *testing.T) {

    // configuration := carbon.NewConfiguration()
    // configuration.SetHost("http://127.0.0.1:4010")
    /* 
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)
    */

    t.Run("Test DataSourcesApiService AddTags", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        
        addDataSourceTagsInput := *carbon.NewAddDataSourceTagsInput(
            null,
            null,
        )
        
        request := client.DataSourcesApi.AddTags(
            addDataSourceTagsInput,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test DataSourcesApiService QueryUserDataSources", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        pagination := *carbon.NewPagination()
        filters := *carbon.NewOrganizationUserDataSourceFilters()
        
        organizationUserDataSourceQueryInput := *carbon.NewOrganizationUserDataSourceQueryInput()
        organizationUserDataSourceQueryInput.SetPagination(pagination)
        organizationUserDataSourceQueryInput.SetOrderBy(null)
        organizationUserDataSourceQueryInput.SetOrderDir(null)
        organizationUserDataSourceQueryInput.SetFilters(filters)
        
        request := client.DataSourcesApi.QueryUserDataSources(
            organizationUserDataSourceQueryInput,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test DataSourcesApiService RemoveTags", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        
        removeDataSourceTagsInput := *carbon.NewRemoveDataSourceTagsInput(
            null,
        )
        removeDataSourceTagsInput.SetTagsToRemove([])
        removeDataSourceTagsInput.SetRemoveAllTags(false)
        
        request := client.DataSourcesApi.RemoveTags(
            removeDataSourceTagsInput,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test DataSourcesApiService RevokeAccessToken", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        
        revokeAccessTokenInput := *carbon.NewRevokeAccessTokenInput(
            null,
        )
        
        request := client.DataSourcesApi.RevokeAccessToken(
            revokeAccessTokenInput,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

}
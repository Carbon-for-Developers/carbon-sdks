/*
Carbon

Testing OrganizationsApiService

*/

// Code generated by Konfig (https://konfigthis.com);

package carbon

import (
    "testing"
    // "github.com/stretchr/testify/assert"
    // "github.com/stretchr/testify/require"
    // carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func Test_carbon_OrganizationsApiService(t *testing.T) {

    // configuration := carbon.NewConfiguration()
    // configuration.SetHost("http://127.0.0.1:4010")
    /* 
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)
    */

    t.Run("Test OrganizationsApiService Get", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        request := client.OrganizationsApi.Get(
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test OrganizationsApiService Update", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        globalUserConfig := *carbon.NewUserConfigurationNullable()
        
        updateOrganizationInput := *carbon.NewUpdateOrganizationInput()
        updateOrganizationInput.SetGlobalUserConfig(globalUserConfig)
        updateOrganizationInput.SetDataSourceConfigs(null)
        
        request := client.OrganizationsApi.Update(
            updateOrganizationInput,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test OrganizationsApiService UpdateStats", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        request := client.OrganizationsApi.UpdateStats(
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

}
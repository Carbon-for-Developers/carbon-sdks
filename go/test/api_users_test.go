/*
Carbon

Testing UsersApiService

*/

// Code generated by Konfig (https://konfigthis.com);

package carbon

import (
    "testing"
    // "github.com/stretchr/testify/assert"
    // "github.com/stretchr/testify/require"
    // carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
)

func Test_carbon_UsersApiService(t *testing.T) {

    // configuration := carbon.NewConfiguration()
    // configuration.SetHost("http://127.0.0.1:4010")
    /* 
    configuration.SetAccessToken("AUTHORIZATION")
    configuration.SetApiKey("AUTHORIZATION")
    configuration.SetCustomerId("CUSTOMER_ID")
    client := carbon.NewAPIClient(configuration)
    */

    t.Run("Test UsersApiService Delete", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        
        deleteUsersInput := *carbon.NewDeleteUsersInput(
            null,
        )
        
        request := client.UsersApi.Delete(
            deleteUsersInput,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test UsersApiService Get", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        
        userRequestContent := *carbon.NewUserRequestContent(
            "null",
        )
        
        request := client.UsersApi.Get(
            userRequestContent,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test UsersApiService ToggleUserFeatures", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        
        modifyUserConfigurationInput := *carbon.NewModifyUserConfigurationInput(
            "null",
            null,
        )
        
        request := client.UsersApi.ToggleUserFeatures(
            modifyUserConfigurationInput,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test UsersApiService UpdateUsers", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        autoSyncEnabledSources := *carbon.NewAutoSyncEnabledSourcesProperty()
        
        updateUsersInput := *carbon.NewUpdateUsersInput(
            null,
        )
        updateUsersInput.SetAutoSyncEnabledSources(autoSyncEnabledSources)
        
        request := client.UsersApi.UpdateUsers(
            updateUsersInput,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

}
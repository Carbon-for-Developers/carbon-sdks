package carbon

import (
	"testing"

	carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_carbon_(t *testing.T) {

	configuration := carbon.NewConfiguration()
	configuration.SetHost("http://127.0.0.1:4010")
	configuration.SetAccessToken("AUTHORIZATION")
	configuration.SetApiKey("AUTHORIZATION")
	configuration.SetCustomerId("CUSTOMER_ID")
	client := carbon.NewAPIClient(configuration)

	t.Run("Test GettingStarted GetAccessToken", func(t *testing.T) {
		request := client.AuthApi.GetAccessToken()
		client.AuthApi.GetAccessToken().Execute()

		resp, httpRes, err := request.Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

}

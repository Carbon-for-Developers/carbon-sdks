package carbon

import (
	"testing"

	carbon "github.com/Carbon-for-Developers/carbon-sdks/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_carbon_(t *testing.T) {

	t.Run("Test GetAccessToken with apiKey and customerId", func(t *testing.T) {
		configuration := carbon.NewConfiguration()
		configuration.SetHost("http://127.0.0.1:4010")
		configuration.SetApiKey("API_KEY")
		configuration.SetCustomerId("CUSTOMER_ID")
		client := carbon.NewAPIClient(configuration)
		request := client.AuthApi.GetAccessToken()

		resp, httpRes, err := request.Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test GetAccessToken with apiKey and customerId", func(t *testing.T) {
		configuration := carbon.NewConfiguration()
		configuration.SetHost("http://127.0.0.1:4010")
		configuration.SetAccessToken("ACCESS_TOKEN")
		client := carbon.NewAPIClient(configuration)
		request := client.AuthApi.GetAccessToken()

		resp, httpRes, err := request.Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

}

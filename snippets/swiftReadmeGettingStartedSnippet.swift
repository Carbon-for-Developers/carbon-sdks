import CarbonAI

// 1) Get an access token for a customer
let carbon = CarbonAIClient(
	accessToken: nil,
    apiKey: "API_KEY",
    customerId: "CUSTOMER_ID"
)

let token = try await carbon.auth.getAccessToken()

// 2) Use the access token to authenticate moving forward
let carbonWithToken = CarbonAIClient(
	accessToken: token!.accessToken,
	apiKey: nil,
	customerId: nil
)

// use SDK as usual
white_labeling = try await carbonWithToken.auth.get_white_labeling()
// etc.

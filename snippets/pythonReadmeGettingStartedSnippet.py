from carbon import Carbon

# 1) Get an access token for a customer
carbon = Carbon(
    api_key="YOUR_API_KEY",
    customer_id="YOUR_CUSTOMER_ID",
)

token = carbon.auth.get_access_token()

# 2) Use the access token to authenticate moving forward
carbon = Carbon(access_token=token.access_token)

# use SDK as usual
white_labeling = carbon.auth.get_white_labeling()
# etc.

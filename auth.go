package vinderman

import (
	"gitlab.com/8h9x/vinderman/auth"
)

func (c *Client) GetExchangeCode(credentials auth.TokenResponse) (auth.ExchangeResponse, error) {
	return auth.GetExchangeCode(c.HttpClient, c.CredentialsMap[c.ClientID])
}

func (c Client) CreateDeviceAuth() (auth.DeviceAuthResponse, error) {
	return auth.CreateDeviceAuth(c.HttpClient, c.CredentialsMap[c.ClientID])
}

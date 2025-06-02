package vinderman

import (
	"encoding/base64"
	"fmt"
	"gitlab.com/8h9x/Vinderman/consts"
	"net/http"
	"net/url"
	"time"

	"gitlab.com/8h9x/Vinderman/eos"
	"gitlab.com/8h9x/Vinderman/request"
)

type AuthClient struct {
	ClientId     string
	ClientSecret string
}

func (ac *AuthClient) String() string {
	return fmt.Sprintf("AuthClient{ClientId: %s}", ac.ClientId)
}

func (ac *AuthClient) BasicToken() string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", ac.ClientId, ac.ClientSecret)))
}

type ClientCredentials eos.ClientCredentials

type DeviceAuth struct {
	AccountID string `json:"accountId"`
	Created   struct {
		DateTime  time.Time `json:"dateTime"`
		IpAddress string    `json:"ipAddress"`
		Location  string    `json:"location"`
	} `json:"created"`
	DeviceId  string `json:"deviceId"`
	Secret    string `json:"secret"`
	UserAgent string `json:"userAgent"`
}

type Exchange eos.Exchange

type UserCredentials struct {
	AccessToken      string    `json:"access_token"`
	AccountID        string    `json:"account_id"`
	ApplicationId    string    `json:"application_id"`
	ClientId         string    `json:"client_id"`
	ExpiresAt        time.Time `json:"expires_at"`
	ExpiresIn        int       `json:"expires_in"`
	RefreshExpiresAt time.Time `json:"refresh_expires_at"`
	RefreshExpiresIn int       `json:"refresh_expires_in"`
	RefreshToken     string    `json:"refresh_token"`
	Scope            []string  `json:"scope"`
	TokenType        string    `json:"token_type"`
}

func (c Client) CreateDeviceAuth(credentials UserCredentials) (deviceAuth DeviceAuth, err error) {
	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprint("Bearer ", credentials.AccessToken))

	resp, err := c.Request("POST", fmt.Sprintf("%s/account/api/public/account/%s/deviceAuth", consts.ACCOUNT_SERVICE, credentials.AccountID), headers, "")
	if err != nil {
		return
	}

	res, err := request.ResponseParser[DeviceAuth](resp)

	return res.Body, err
}

func (c Client) GetClientCredentials(ac AuthClient) (credentials ClientCredentials, err error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/x-www-form-urlencoded")
	headers.Set("Authorization", fmt.Sprint("Basic ", ac.BasicToken()))

	v := url.Values{}
	v.Set("grant_type", "client_credentials")
	body := v.Encode()

	resp, err := c.Request("POST", consts.ACCOUNT_AUTH+"/token", headers, body)
	if err != nil {
		return
	}

	res, err := request.ResponseParser[ClientCredentials](resp)

	return res.Body, err
}

func (c Client) GetExchangeCode(credentials UserCredentials) (exchange Exchange, err error) {
	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprint("Bearer ", credentials.AccessToken))

	resp, err := c.Request("GET", consts.ACCOUNT_AUTH+"/exchange", headers, "")
	if err != nil {
		return
	}

	res, err := request.ResponseParser[Exchange](resp)

	return res.Body, err
}

func (c Client) RefreshTokenLogin(ac AuthClient, refreshToken string) (credentials UserCredentials, err error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/x-www-form-urlencoded")
	headers.Set("Authorization", fmt.Sprint("Basic ", ac.BasicToken()))

	v := url.Values{}
	v.Set("grant_type", "refresh_token")
	v.Set("refresh_token", refreshToken)
	body := v.Encode()

	resp, err := c.Request("POST", consts.ACCOUNT_AUTH+"/token", headers, body)
	if err != nil {
		return
	}

	res, err := request.ResponseParser[UserCredentials](resp)

	return res.Body, err
}

func (c Client) ExchangeCodeLogin(ac AuthClient, code string) (credentials UserCredentials, err error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/x-www-form-urlencoded")
	headers.Set("Authorization", fmt.Sprint("Basic ", ac.BasicToken()))

	v := url.Values{}
	v.Set("grant_type", "exchange_code")
	v.Set("exchange_code", code)
	v.Set("scope", "offline_access")
	body := v.Encode()

	resp, err := c.Request("POST", consts.ACCOUNT_AUTH+"/token", headers, body)
	if err != nil {
		return
	}

	res, err := request.ResponseParser[UserCredentials](resp)

	return res.Body, err
}

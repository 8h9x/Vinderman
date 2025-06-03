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

type AuthPayloadAuthorizationCode struct {
	Code string `json:"code"`
}

type AuthPayloadContinuationToken struct {
	ContinuationToken string `json:"continuation_token"`
}

type AuthPayloadDeviceAuth struct {
	AccountID string `json:"account_id"`
	DeviceID  string `json:"device_id"`
	Secret    string `json:"secret"`
}

type AuthPayloadDeviceCode struct {
	DeviceCode string `json:"device_code"`
}

type AuthPayloadExchangeCode struct {
	ExchangeCode string `json:"exchange_code"`
}

type AuthPayloadExternalAuth struct {
	ExternalAuthToken string `json:"external_auth_token"`
}

type AuthPayloadOTP struct {
	OTP string `json:"otp"`
}

type AuthPayloadPassword struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthPayloadRefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

type AuthPayloadTokenToToken struct {
	AccessToken string `json:"access_token"`
}

type AuthPayload interface {
	AuthPayloadAuthorizationCode | AuthPayloadContinuationToken | AuthPayloadDeviceAuth |
	AuthPayloadDeviceCode | AuthPayloadExchangeCode | AuthPayloadExternalAuth | AuthPayloadOTP |
	AuthPayloadPassword | AuthPayloadRefreshToken | AuthPayloadTokenToToken
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

type DeviceAuthorization eos.DeviceAuthorization

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

func (c Client) GetDeviceCode(credentials ClientCredentials) (deviceAuth DeviceAuthorization, err error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/x-www-form-urlencoded")
	headers.Set("Authorization", fmt.Sprint("Bearer ", credentials.AccessToken))

	v := url.Values{}
	v.Set("prompt", "login")
	body := v.Encode()

	resp, err := c.Request("POST", consts.ACCOUNT_AUTH+"/deviceAuthorization", headers, body)
	if err != nil {
		return
	}

	res, err := request.ResponseParser[DeviceAuthorization](resp)

	return res.Body, err
}

func (c Client) WaitForDeviceCodeAccept(ac AuthClient, deviceCode string) (credentials UserCredentials, err error) {
	credentials, err = c.Login(ac, LoginBuilder(AuthPayloadDeviceCode{
		DeviceCode: deviceCode,
	}))

	if err != nil {
		if err.(*request.Error[EpicErrorResponse]).Raw.ErrorCode == consts.ErrorAuthorizationPending {
			time.Sleep(10 * time.Second)
			return c.WaitForDeviceCodeAccept(ac, deviceCode)
		}

		return
	}

	return
}

type BuiltAuthPayload struct {
	raw interface{}
}

func LoginBuilder[T AuthPayload](payload T) BuiltAuthPayload {
	return BuiltAuthPayload{raw: payload}
}

func (c Client) Login(ac AuthClient, payload BuiltAuthPayload) (credentials UserCredentials, err error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/x-www-form-urlencoded")
	headers.Set("Authorization", fmt.Sprint("Basic ", ac.BasicToken()))

	v := url.Values{}

	switch p := payload.raw.(type) {
	case AuthPayloadAuthorizationCode:
		v.Set("grant_type", "authorization_code")
		v.Set("code", p.Code)
	case AuthPayloadContinuationToken:
		v.Set("grant_type", "continuation_token")
		v.Set("continuation_token", p.ContinuationToken)
	case AuthPayloadDeviceAuth:
		v.Set("grant_type", "device_auth")
		v.Set("account_id", p.AccountID)
		v.Set("device_id", p.DeviceID)
		v.Set("secret", p.Secret)
	case AuthPayloadDeviceCode:
		v.Set("grant_type", "device_code")
		v.Set("device_code", p.DeviceCode)
	case AuthPayloadExchangeCode:
		v.Set("grant_type", "exchange_code")
		v.Set("exchange_code", p.ExchangeCode)
	case AuthPayloadExternalAuth:
		v.Set("grant_type", "external_auth")
		v.Set("external_auth_token", p.ExternalAuthToken)
	case AuthPayloadOTP:
		v.Set("grant_type", "otp")
		v.Set("otp", p.OTP)
	case AuthPayloadPassword:
		v.Set("grant_type", "password")
		v.Set("username", p.Username)
		v.Set("password", p.Password)
	case AuthPayloadRefreshToken:
		v.Set("grant_type", "refresh_token")
		v.Set("refresh_token", p.RefreshToken)
	case AuthPayloadTokenToToken:
		v.Set("grant_type", "token_to_token")
		v.Set("access_token", p.AccessToken)
	default:
		return UserCredentials{}, fmt.Errorf("unsupported payload type: %T", payload)
	}

	body := v.Encode()
	resp, err := c.Request("POST", consts.ACCOUNT_AUTH+"/token", headers, body)
	if err != nil {
		return
	}

	res, err := request.ResponseParser[UserCredentials](resp)
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

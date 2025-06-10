package vinderman

import (
	"encoding/base64"
	"fmt"
	"gitlab.com/8h9x/vinderman/consts"
	"gitlab.com/8h9x/vinderman/request"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type AuthPayloadAuthorizationCode struct {
	Code string `json:"code"`
}

type AuthPayloadClientCredentials struct{}

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
	AuthPayloadAuthorizationCode | AuthPayloadClientCredentials | AuthPayloadContinuationToken |
		AuthPayloadDeviceAuth | AuthPayloadDeviceCode | AuthPayloadExchangeCode | AuthPayloadExternalAuth |
		AuthPayloadOTP | AuthPayloadPassword | AuthPayloadRefreshToken | AuthPayloadTokenToToken
}

func Authenticate[T AuthPayload](clientId, clientSecret string, payload T, eg1 bool) (AuthTokenResponse, error) {
	v := url.Values{}
	if eg1 == true {
		v.Set("token_type", "eg1")
	}

	switch p := any(payload).(type) {
	case AuthPayloadAuthorizationCode:
		v.Set("grant_type", "authorization_code")
		v.Set("code", p.Code)
		break
	case AuthPayloadClientCredentials:
		v.Set("grant_type", "client_credentials")
		break
	case AuthPayloadContinuationToken:
		v.Set("grant_type", "continuation_token")
		v.Set("continuation_token", p.ContinuationToken)
		break
	case AuthPayloadDeviceAuth:
		v.Set("grant_type", "device_auth")
		v.Set("account_id", p.AccountID)
		v.Set("device_id", p.DeviceID)
		v.Set("secret", p.Secret)
		break
	case AuthPayloadDeviceCode:
		v.Set("grant_type", "device_code")
		v.Set("device_code", p.DeviceCode)
		break
	case AuthPayloadExchangeCode:
		v.Set("grant_type", "exchange_code")
		v.Set("exchange_code", p.ExchangeCode)
		break
	case AuthPayloadExternalAuth:
		v.Set("grant_type", "external_auth")
		v.Set("external_auth_token", p.ExternalAuthToken)
		break
	case AuthPayloadOTP:
		v.Set("grant_type", "otp")
		v.Set("otp", p.OTP)
		break
	case AuthPayloadPassword:
		v.Set("grant_type", "password")
		v.Set("username", p.Username)
		v.Set("password", p.Password)
		break
	case AuthPayloadRefreshToken:
		v.Set("grant_type", "refresh_token")
		v.Set("refresh_token", p.RefreshToken)
		break
	case AuthPayloadTokenToToken:
		v.Set("grant_type", "token_to_token")
		v.Set("access_token", p.AccessToken)
		break
	}

	httpClient := &http.Client{}
	req, err := http.NewRequest("POST", consts.AccountProxyService+"/account/api/oauth/token", strings.NewReader(v.Encode()))
	if err != nil {
		return AuthTokenResponse{}, err
	}

	basicToken := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", clientId, clientSecret)))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", fmt.Sprint("Basic ", basicToken))

	resp, err := httpClient.Do(req)
	if err != nil {
		return AuthTokenResponse{}, err
	}

	res, err := request.ResponseParser[AuthTokenResponse](resp)
	if err != nil {
		return AuthTokenResponse{}, err
	}

	return res.Body, err
}

type AuthTokenResponse struct {
	AccessToken      string    `json:"access_token"`
	ExpiresIn        int       `json:"expires_in"`
	ExpiresAt        time.Time `json:"expires_at"`
	TokenType        string    `json:"token_type"`
	RefreshToken     string    `json:"refresh_token,omitempty"`
	RefreshExpires   int       `json:"refresh_expires,omitempty"`
	RefreshExpiresAt time.Time `json:"refresh_expires_at,omitempty"`
	AccountId        string    `json:"account_id,omitempty"`
	ClientId         string    `json:"client_id"`
	InternalClient   bool      `json:"internal_client"`
	ClientService    string    `json:"client_service"`
	DisplayName      string    `json:"displayName,omitempty"`
	App              string    `json:"app,omitempty"`
	InAppId          string    `json:"in_app_id,omitempty"`
	DeviceId         string    `json:"device_id,omitempty"`
	ProductId        string    `json:"product_id"`
	ApplicationId    string    `json:"application_id"`
	Acr              string    `json:"acr,omitempty"`
	AuthTime         time.Time `json:"auth_time,omitempty"`
}

func VerifyToken(accessToken string, includePerms bool) (AuthVerifyResponse, error) {
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", consts.AccountProxyService+"/account/api/oauth/verify?", nil)
	if err != nil {
		return AuthVerifyResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprint("Bearer ", accessToken))

	if includePerms == true {
		req.URL.Query().Set("includePerms", "true")
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return AuthVerifyResponse{}, err
	}

	res, err := request.ResponseParser[AuthVerifyResponse](resp)
	if err != nil {
		return AuthVerifyResponse{}, err
	}

	return res.Body, err
}

type AuthVerifyResponse struct {
	Token          string    `json:"token"`
	SessionId      string    `json:"session_id"`
	TokenType      string    `json:"token_type"`
	ClientId       string    `json:"client_id"`
	InternalClient bool      `json:"internal_client"`
	ClientService  string    `json:"client_service"`
	AccountId      string    `json:"account_id"`
	ExpiresIn      int       `json:"expires_in"`
	ExpiresAt      time.Time `json:"expires_at"`
	AuthMethod     string    `json:"auth_method"`
	DisplayName    string    `json:"display_name"`
	App            string    `json:"app"`
	InAppId        string    `json:"in_app_id"`
	DeviceId       string    `json:"device_id"`
	Scope          []string  `json:"scope"`
	ProductId      string    `json:"product_id"`
	SandboxId      string    `json:"sandbox_id"`
	DeploymentId   string    `json:"deployment_id"`
	ApplicationId  string    `json:"application_id"`
	Acr            string    `json:"acr"`
	AuthTime       time.Time `json:"auth_time"`
	Perms          []struct {
		Resource string `json:"resource"`
		Action   string `json:"action"`
	} `json:"perms"`
}

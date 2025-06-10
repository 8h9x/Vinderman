package auth

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

type PayloadAuthorizationCode struct {
	Code string `json:"code"`
}

type PayloadClientCredentials struct{}

type PayloadContinuationToken struct {
	ContinuationToken string `json:"continuation_token"`
}

type PayloadDeviceAuth struct {
	AccountID string `json:"account_id"`
	DeviceID  string `json:"device_id"`
	Secret    string `json:"secret"`
}

type PayloadDeviceCode struct {
	DeviceCode string `json:"device_code"`
}

type PayloadExchangeCode struct {
	ExchangeCode string `json:"exchange_code"`
}

type PayloadExternalAuth struct {
	ExternalAuthToken string `json:"external_auth_token"`
}

type PayloadOTP struct {
	OTP string `json:"otp"`
}

type PayloadPassword struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PayloadRefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

type PayloadTokenToToken struct {
	AccessToken string `json:"access_token"`
}

type Payload interface {
	PayloadAuthorizationCode | PayloadClientCredentials | PayloadContinuationToken |
		PayloadDeviceAuth | PayloadDeviceCode | PayloadExchangeCode | PayloadExternalAuth |
		PayloadOTP | PayloadPassword | PayloadRefreshToken | PayloadTokenToToken
}

func Authenticate[T Payload](clientId, clientSecret string, payload T, eg1 bool) (TokenResponse, error) {
	v := url.Values{}
	if eg1 == true {
		v.Set("token_type", "eg1")
	}

	switch p := any(payload).(type) {
	case PayloadAuthorizationCode:
		v.Set("grant_type", "authorization_code")
		v.Set("code", p.Code)
		break
	case PayloadClientCredentials:
		v.Set("grant_type", "client_credentials")
		break
	case PayloadContinuationToken:
		v.Set("grant_type", "continuation_token")
		v.Set("continuation_token", p.ContinuationToken)
		break
	case PayloadDeviceAuth:
		v.Set("grant_type", "device_auth")
		v.Set("account_id", p.AccountID)
		v.Set("device_id", p.DeviceID)
		v.Set("secret", p.Secret)
		break
	case PayloadDeviceCode:
		v.Set("grant_type", "device_code")
		v.Set("device_code", p.DeviceCode)
		break
	case PayloadExchangeCode:
		v.Set("grant_type", "exchange_code")
		v.Set("exchange_code", p.ExchangeCode)
		break
	case PayloadExternalAuth:
		v.Set("grant_type", "external_auth")
		v.Set("external_auth_token", p.ExternalAuthToken)
		break
	case PayloadOTP:
		v.Set("grant_type", "otp")
		v.Set("otp", p.OTP)
		break
	case PayloadPassword:
		v.Set("grant_type", "password")
		v.Set("username", p.Username)
		v.Set("password", p.Password)
		break
	case PayloadRefreshToken:
		v.Set("grant_type", "refresh_token")
		v.Set("refresh_token", p.RefreshToken)
		break
	case PayloadTokenToToken:
		v.Set("grant_type", "token_to_token")
		v.Set("access_token", p.AccessToken)
		break
	}

	httpClient := &http.Client{}
	req, err := http.NewRequest("POST", consts.AccountProxyService+"/account/api/oauth/token", strings.NewReader(v.Encode()))
	if err != nil {
		return TokenResponse{}, err
	}

	basicToken := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", clientId, clientSecret)))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", fmt.Sprint("Basic ", basicToken))

	resp, err := httpClient.Do(req)
	if err != nil {
		return TokenResponse{}, err
	}

	res, err := request.ResponseParser[TokenResponse](resp)
	if err != nil {
		return TokenResponse{}, err
	}

	return res.Body, err
}

type TokenResponse struct {
	AccessToken      string    `json:"access_token"`
	ExpiresIn        int       `json:"expires_in"`
	ExpiresAt        time.Time `json:"expires_at"`
	TokenType        string    `json:"token_type"`
	RefreshToken     string    `json:"refresh_token"`
	RefreshExpires   int       `json:"refresh_expires"`
	RefreshExpiresAt time.Time `json:"refresh_expires_at"`
	AccountId        string    `json:"account_id"`
	ClientId         string    `json:"client_id"`
	InternalClient   bool      `json:"internal_client"`
	ClientService    string    `json:"client_service"`
	DisplayName      string    `json:"displayName"`
	App              string    `json:"app"`
	InAppId          string    `json:"in_app_id"`
	DeviceId         string    `json:"device_id"`
	ProductId        string    `json:"product_id"`
	ApplicationId    string    `json:"application_id"`
	Acr              string    `json:"acr"`
	AuthTime         time.Time `json:"auth_time"`
}

func VerifyToken(accessToken string, includePerms bool) (VerifyTokenResponse, error) {
	httpClient := &http.Client{}
	v := url.Values{}

	if includePerms {
		v.Set("includePerms", "true")
	}

	req, err := http.NewRequest("GET", consts.AccountProxyService+"/account/api/oauth/verify", strings.NewReader(v.Encode()))
	if err != nil {
		return VerifyTokenResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprint("Bearer ", accessToken))

	resp, err := httpClient.Do(req)
	if err != nil {
		return VerifyTokenResponse{}, err
	}

	res, err := request.ResponseParser[VerifyTokenResponse](resp)
	if err != nil {
		return VerifyTokenResponse{}, err
	}

	return res.Body, err
}

type VerifyTokenResponse struct {
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

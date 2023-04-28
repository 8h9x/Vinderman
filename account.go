package vinderman

import (
	"fmt"
	"net/http"
	"time"

	"github.com/0xDistrust/Vinderman/consts"
	"github.com/0xDistrust/Vinderman/request"
)

type AccountExternalLookup []struct {
	ID            string `json:"id"`
	DisplayName   string `json:"displayName"`
	ExternalAuths map[string]struct {
		AccountID           string        `json:"accountId"`
		Type                string        `json:"type"`
		ExternalAuthIDType  string        `json:"externalAuthIdType"`
		ExternalDisplayName string        `json:"externalDisplayName"`
		AuthIDs             []interface{} `json:"authIds"`
	} `json:"externalAuths"`
}

type AccountLookup struct {
	AgeGroup                   string    `json:"ageGroup"`
	CabinedMode                bool      `json:"cabinedMode"`
	CanUpdateDisplayName       bool      `json:"canUpdateDisplayName"`
	CanUpdateDisplayNameNext   time.Time `json:"canUpdateDisplayNameNext"`
	Country                    string    `json:"country"`
	DisplayName                string    `json:"displayName"`
	Email                      string    `json:"email"`
	EmailVerified              bool      `json:"emailVerified"`
	FailedLoginAttempts        int       `json:"failedLoginAttempts"`
	HasHashedEmail             bool      `json:"hasHashedEmail"`
	Headless                   bool      `json:"headless"`
	ID                         string    `json:"id"`
	LastDisplayNameChange      time.Time `json:"lastDisplayNameChange"`
	LastLogin                  time.Time `json:"lastLogin"`
	LastName                   string    `json:"lastName"`
	MinorExpected              bool      `json:"minorExpected"`
	MinorStatus                string    `json:"minorStatus"`
	MinorVerified              bool      `json:"minorVerified"`
	Name                       string    `json:"name"`
	NumberOfDisplayNameChanges int       `json:"numberOfDisplayNameChanges"`
	PreferredLanguage          string    `json:"preferredLanguage"`
	TFAEnabled                 bool      `json:"tfaEnabled"`
}

func (c Client) FetchMe(credentials UserCredentials) (AccountLookup, error) {
	return c.FetchAccountByID(credentials, credentials.AccountID)
}

func (c Client) FetchAccountByID(credentials UserCredentials, accountID string) (userLookup AccountLookup, err error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	resp, err := c.Request("GET", fmt.Sprintf("%s/account/api/public/account/%s", consts.ACCOUNT_SERVICE, accountID), headers, "")
	if err != nil {
		return
	}

	res, err := request.ResponseParser[AccountLookup](resp)

	return res.Body, err
}

func (c Client) FetchAccountByDisplayName(credentials UserCredentials, displayName string) (userLookup AccountLookup, err error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	resp, err := c.Request("GET", fmt.Sprintf("%s/account/api/public/account/displayName/%s", consts.ACCOUNT_SERVICE, displayName), headers, "")
	if err != nil {
		return
	}

	res, err := request.ResponseParser[AccountLookup](resp)

	return res.Body, err
}

func (c Client) FetchAccountByExternalDisplayName(credentials UserCredentials, displayName string, platform ExternalAuthType) (userExternalLookup AccountExternalLookup, err error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	resp, err := c.Request("GET", fmt.Sprintf("%s/account/api/public/account/lookup/externalAuth/%s/displayName/%s?caseInsensitive=true", consts.ACCOUNT_SERVICE, platform, displayName), headers, "")
	if err != nil {
		return
	}

	res, err := request.ResponseParser[AccountExternalLookup](resp)

	return res.Body, err
}

package vinderman

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"gitlab.com/8h9x/Vinderman/consts"
	"gitlab.com/8h9x/Vinderman/request"
)

type UserExternalLookup []struct {
	ID            string `json:"id"`
	DisplayName   string `json:"displayName"`
	ExternalAuths map[string]struct {
		AccountID           string        `json:"accountId"`
		Type                string        `json:"type"`
		ExternalAuthIDType  string        `json:"externalAuthIdType"`
		ExternalDisplayName string        `json:"externalDisplayName"`
		AuthIds             []interface{} `json:"authIds"`
	} `json:"externalAuths"`
}

type UserLookup struct {
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

type UserSearchData struct {
	AccountID string `json:"accountId"`
	Matches   []struct {
		Value    string `json:"value"`
		Platform string `json:"platform"`
	} `json:"matches"`
	MatchType    string `json:"matchType"`
	EpicMutuals  int    `json:"epicMutuals"`
	SortPosition int    `json:"sortPosition"`
}

func (c Client) FetchMe(credentials UserCredentials) (UserLookup, error) {
	return c.FetchUserByID(credentials, credentials.AccountID)
}

func (c Client) FetchUserByID(credentials UserCredentials, accountID string) (userLookup UserLookup, err error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	resp, err := c.Request("GET", fmt.Sprintf("%s/account/api/public/account/%s", consts.ACCOUNT_SERVICE, accountID), headers, "")
	if err != nil {
		return
	}

	res, err := request.ResponseParser[UserLookup](resp)

	return res.Body, err
}

func (c Client) FetchUserByDisplayName(credentials UserCredentials, displayName string) (userLookup UserLookup, err error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	resp, err := c.Request("GET", fmt.Sprintf("%s/account/api/public/account/displayName/%s", consts.ACCOUNT_SERVICE, displayName), headers, "")
	if err != nil {
		return
	}

	res, err := request.ResponseParser[UserLookup](resp)

	return res.Body, err
}

func (c Client) FetchUserByExternalDisplayName(credentials UserCredentials, displayName string, platform ExternalAuthType) (userExternalLookup UserExternalLookup, err error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	resp, err := c.Request("GET", fmt.Sprintf("%s/account/api/public/account/lookup/externalAuth/%s/displayName/%s?caseInsensitive=true", consts.ACCOUNT_SERVICE, platform, displayName), headers, "")
	if err != nil {
		return
	}

	res, err := request.ResponseParser[UserExternalLookup](resp)

	return res.Body, err
}

func (c Client) SearchUsers(credentials UserCredentials, displayName string, platform ExternalAuthType) (offers []UserSearchData, err error) {
	values := url.Values{}
	values.Add("prefix", displayName)
	values.Add("platform", string(platform))

	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	resp, err := c.Request("GET", fmt.Sprintf("%s/api/v1/search/%s?%s", consts.USERSEARCH_SERVICE, credentials.AccountID, values.Encode()), headers, "")
	if err != nil {
		return
	}

	res, err := request.ResponseParser[[]UserSearchData](resp)

	return res.Body, err
}

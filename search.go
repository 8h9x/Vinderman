package vinderman

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/0xDistrust/Vinderman/consts"
	"github.com/0xDistrust/Vinderman/request"
)

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

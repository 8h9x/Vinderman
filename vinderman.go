package vinderman

import (
	"fmt"
	"gitlab.com/8h9x/vinderman/auth"
	"net/http"
)

type Client struct {
	HttpClient     *http.Client
	Header         http.Header
	ClientID       string
	CredentialsMap map[string]auth.TokenResponse
}

func NewClient(httpClient *http.Client, initCredentials auth.TokenResponse) (*Client, error) {
	client := &Client{
		HttpClient:     httpClient,
		Header:         make(http.Header),
		ClientID:       initCredentials.ClientId,
		CredentialsMap: make(map[string]auth.TokenResponse),
	}

	client.CredentialsMap[initCredentials.ClientId] = initCredentials

	_, err := auth.VerifyToken(initCredentials.AccessToken, false)
	if err != nil {
		return client, err
	}

	mcpVersionData, err := client.GetMCPVersion()
	if err != nil {
		return client, err
	}

	client.Header.Set("User-Agent", fmt.Sprintf("Fortnite/++Fortnite+%s-CL-%s Windows/10.0.26100.1.256.64bit", mcpVersionData.Branch, mcpVersionData.Cln))

	return client, nil
}

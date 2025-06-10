package vinderman

import (
	"net/http"
)

type Client struct {
	HttpClient     *http.Client
	CredentialsMap map[string]AuthTokenResponse
}

func NewClient(httpClient *http.Client, initCredentials AuthTokenResponse) (*Client, error) {
	_, err := VerifyToken(initCredentials.AccessToken, true)
	if err != nil {
		return &Client{}, err
	}

	credentialsMap := make(map[string]AuthTokenResponse)
	credentialsMap[initCredentials.ClientId] = initCredentials

	return &Client{
		HttpClient:     httpClient,
		CredentialsMap: credentialsMap,
	}, nil
}

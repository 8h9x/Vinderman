package vinderman

import (
	"net/http"

	"gitlab.com/8h9x/Vinderman/eos"
	"gitlab.com/8h9x/Vinderman/external/fnapicom"
)

type Client struct {
	HttpClient *http.Client
	EOS        *eos.Client
	FNApiCom   *fnapicom.Client
}

type EpicErrorResponse eos.EpicErrorResponse

func New() *Client {
	httpClient := &http.Client{}

	return &Client{
		HttpClient: httpClient,
		EOS: &eos.Client{
			HttpClient: httpClient,
		},
		FNApiCom: &fnapicom.Client{},
	}
}

func (c Client) Request(method string, url string, headers http.Header, body string) (*http.Response, error) {
	return c.EOS.Request(method, url, headers, body)
}

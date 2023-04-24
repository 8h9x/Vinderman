package vinderman

import (
	"net/http"

	"github.com/0xDistrust/Vinderman/eos"
	"github.com/0xDistrust/Vinderman/external/fnapicom"
	"github.com/0xDistrust/Vinderman/external/nitestats"
)

type Client struct {
	HttpClient *http.Client
	EOS        *eos.Client
	FNApiCom   *fnapicom.Client
	Nitestats  *nitestats.Client
}

type EpicErrorResponse eos.EpicErrorResponse

func New() *Client {
	httpClient := &http.Client{}

	return &Client{
		HttpClient: httpClient,
		EOS: &eos.Client{
			HttpClient: httpClient,
		},
		FNApiCom:  &fnapicom.Client{},
		Nitestats: &nitestats.Client{},
	}
}

func (c Client) Request(method string, url string, headers http.Header, body string) (*http.Response, error) {
	return c.EOS.Request(method, url, headers, body)
}
package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/0xDistrust/Vinderman/consts"
	"github.com/0xDistrust/Vinderman/eos"
	"github.com/0xDistrust/Vinderman/external/fnapicom"
	"github.com/0xDistrust/Vinderman/external/nitestats"
)

type Client struct {
	HttpClient *http.Client
	FNApiCom   *fnapicom.Client
	Nitestats  *nitestats.Client
	EOS        *eos.Client
}

type EpicErrorResponse eos.EpicErrorResponse

func New() *Client {
	httpClient := &http.Client{}

	return &Client{
		FNApiCom:  &fnapicom.Client{},
		Nitestats: &nitestats.Client{},
		EOS: &eos.Client{
			HttpClient: httpClient,
		},
	}
}

func (c Client) Request(method string, url string, headers http.Header, body string) (*http.Response, error) {
	return c.EOS.Request(method, url, headers, body)
}

func main() {
	vinderman := New()

	log.Println(vinderman.Nitestats.FetchFLToken())
	log.Println(vinderman.FNApiCom.CosmeticSearch(url.Values{
		"name": {"Raven"},
	}))
	log.Println(vinderman.EOS.GetClientCredentials(consts.FORTNITE_PC_CLIENT_ID, consts.FORTNITE_PC_CLIENT_SECRET))
}

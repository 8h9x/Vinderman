package eos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/0xDistrust/Vinderman/request"
)

type Client struct {
	HttpClient *http.Client
}

type EpicErrorResponse struct {
	ErrorCode          string `json:"errorCode"`
	ErrorMessage       string `json:"errorMessage"`
	NumericErrorCode   int    `json:"numericErrorCode"`
	OriginatingService string `json:"originatingService"`
	Intent             string `json:"intent"`
}

func New() *Client {
	return &Client{
		HttpClient: &http.Client{},
	}
}

func (c Client) Request(method string, url string, headers http.Header, body string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	
	req.Header.Set("User-Agent", "FortniteGame/++Fortnite+Release-24.20-CL-25029190 Windows/10.0.22000.1.768.64bit")

	for key, value := range headers {
		req.Header.Set(key, value[0])
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		if resp.Body != nil {
			defer resp.Body.Close()

			var res EpicErrorResponse
			err = json.NewDecoder(resp.Body).Decode(&res)
			if err != nil {
				return nil, err
			}

			if res.ErrorMessage != "" {
				return nil, &request.Error[EpicErrorResponse]{
					StatusCode: resp.StatusCode,
					Message:    fmt.Sprintf("%s request to %s failed with error message: %s", method, url, res.ErrorMessage),
					Raw:        res,
				}
			}
		}

		return nil, fmt.Errorf("%s request to %s failed with status code %d", method, url, resp.StatusCode)
	}

	return resp, nil
}

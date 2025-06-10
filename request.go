package vinderman

import (
	"encoding/json"
	"fmt"
	"gitlab.com/8h9x/vinderman/request"
	"net/http"
	"strings"
)

type EpicErrorResponse struct {
	ErrorCode          string `json:"errorCode"`
	ErrorMessage       string `json:"errorMessage"`
	NumericErrorCode   int    `json:"numericErrorCode"`
	OriginatingService string `json:"originatingService"`
	Intent             string `json:"intent"`
}

func (c *Client) Request(method string, url string, header http.Header, body string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	for key, value := range c.Header {
		req.Header.Set(key, value[0])
	}

	for key, value := range header {
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

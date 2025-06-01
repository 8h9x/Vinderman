package vinderman

import (
	"encoding/json"
	"net/http"

	"github.com/0xDistrust/Vinderman/consts"
	"github.com/0xDistrust/Vinderman/request"
)

type CalderaRequest struct {
	AccountId    string `json:"account_id"`
	ExchangeCode string `json:"exchange_code"`
	TestMode     bool   `json:"test_mode"`
	EpicApp      string `json:"epic_app"`
	Nvidia       bool   `json:"nvidia"`
}

type CalderaToken struct {
	Provider string
	Jwt      string
}

func (c Client) FetchCaldera() (caldera CalderaToken, err error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")

	payload := CalderaRequest{
		Nvidia: true,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return
	}

	resp, err := c.Request("POST", consts.CALDERA_SERVICE+"/caldera/api/v1/launcher/racp", headers, string(body))
	if err != nil {
		return
	}

	res, err := request.ResponseParser[CalderaToken](resp)

	return res.Body, err
}

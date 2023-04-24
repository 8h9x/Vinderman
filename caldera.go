package vinderman

import (
	"encoding/json"
	"net/http"
	
	"github.com/0xDistrust/Vinderman/consts"
	"github.com/0xDistrust/Vinderman/request"
)

type Caldera struct {
	Provider string
	Jwt      string
}

func (c Client) FetchCaldera() (caldera Caldera, err error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	
	rawBody := map[string]bool{"nvidia": true}
	body, err := json.Marshal(rawBody)
	if err != nil {
		return
	}
	
	resp, err := c.Request("POST", consts.CALDERA_SERVICE+"/caldera/api/v1/launcher/racp", headers, string(body))
	if err != nil {
		return
	}

	res, err := request.ResponseParser[Caldera](resp)

	return res.Body, err
}
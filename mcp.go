package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/0xDistrust/Vinderman/consts"
)

func (c Client) ComposeProfileOperation(credentials UserCredentials, operation string, profileID string, payload string) (resp *http.Response, err error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	resp, err = c.Request("POST", fmt.Sprintf("%s/profile/%s/client/%s?profileId=%s&rvn=-1", consts.FORTNITE_GAME, credentials.AccountID, operation, profileID), headers, payload)
	return
}

func (c Client) ProfileOperation(credentials UserCredentials, operation string, profileId string, payload any) (any, error) {
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return c.ComposeProfileOperation(credentials, operation, profileId, string(bodyBytes))
}

//func ProfileOperationExample(credentials UserCredentials, operation string, profileID string, payload string) () {
//	client := New()
//	res, err := client.ComposeProfileOperation(credentials, operation, profileID, payload)
//	if err != nil {
//		log.Println(err)
//	}
//	
//	data, err := request.ResponseParser[UserCredentials](res)
//	if err != nil {
//		log.Println(err)
//	}
//}
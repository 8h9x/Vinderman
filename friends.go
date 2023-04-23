package main

import (
	"fmt"
	"net/http"

	"github.com/0xDistrust/Vinderman/consts"
)

func (c Client) AddFriend(credentials UserCredentials, friendID string) (err error) {
	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprint("Bearer ", credentials.AccessToken))

	resp, err := c.Request("POST", fmt.Sprintf("%s/friends/api/v1/%s/friends/%s", consts.FRIENDS_SERVICE, credentials.AccountID, friendID), headers, "")
	if err != nil {
		return
	}

	if resp.StatusCode >= 300 {
		return fmt.Errorf("failed to add friend: %s", resp.Status)
	}
	
	return
}
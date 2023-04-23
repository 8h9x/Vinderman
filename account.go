package main

import (
	"fmt"
	"net/http"
	
	"github.com/0xDistrust/Vinderman/consts"
	"github.com/0xDistrust/Vinderman/request"
)

type BRInventory struct {
	Stash struct {
		Globalcash int `json:"globalcash"`
	} `json:"stash"`
}

func (c Client) FetchBRInventory(credentials UserCredentials) (inventory BRInventory, err error) {
	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprint("Bearer ", credentials.AccessToken))

	resp, err := c.Request("GET", fmt.Sprintf("%s/br-inventory/account/%s", consts.FORTNITE_GAME, credentials.AccountID), headers, "")
	if err != nil {
		return
	}

	res, err := request.ResponseParser[BRInventory](resp)

	return res.Body, err
}
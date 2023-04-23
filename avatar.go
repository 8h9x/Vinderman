package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/0xDistrust/Vinderman/consts"
	"github.com/0xDistrust/Vinderman/request"
)

type Avatar struct {
	AccountID string `json:"accountId"`
	Namespace string `json:"namespace"`
	AvatarID  string `json:"avatarId"`
}

func (c Client) FetchAvatar(credentials UserCredentials) (avatar Avatar, err error) {
	avatars, err := c.FetchAvatarBulk(credentials, credentials.AccountID)
	if err != nil {
		return
	}

	return avatars[0], nil
}

func (c Client) FetchAvatarBulk(credentials UserCredentials, accountIDs ...string) (avatar []Avatar, err error) {
	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprint("Bearer ", credentials.AccessToken))

	resp, err := c.Request("GET", fmt.Sprintf("%s/v1/avatar/fortnite/ids?accountIds=%s", consts.AVATAR_SERVICE, strings.Join(accountIDs, ",")), headers, "")
	if err != nil {
		return
	}

	res, err := request.ResponseParser[[]Avatar](resp)

	return res.Body, err
}

func (c Client) FetchAvatarURL(credentials UserCredentials) (string, error) {
	avatar, err := c.FetchAvatar(credentials)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://fortnite-api.com/images/cosmetics/br/%s/icon.png", strings.Replace(avatar.AvatarID, "ATHENACHARACTER:", "", -1)), nil
}
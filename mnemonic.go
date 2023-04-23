package main

import (
	"fmt"
	"net/http"
	"time"
	
	"github.com/0xDistrust/Vinderman/consts"
	"github.com/0xDistrust/Vinderman/request"
)

type MnemonicInfo struct {
	AccountID       string    `json:"accountId"`
	Active          bool      `json:"active"`
	Created         time.Time `json:"createdAt"`
	CreatorName     string    `json:"creatorName"`
	DescriptionTags []string  `json:"descriptionTags"`
	Disabled        bool      `json:"disabled"`
	LinkType        string    `json:"linkType"`
	Metadata        struct {
		DynamicXp struct {
			CalibrationPhase string `json:"calibrationPhase"`
			// UniqueGameVersion string `json:"uniqueGameVersion"` // somtimes int, sometimes string
		} `json:"dynamicXp"`
		GeneratedIslandUrlsOld struct {
			URL  string `json:"url"`
			URLM string `json:"url_m"`
			URLS string `json:"url_s"`
		} `json:"generated_island_urls_old"`
		ImageURL     string `json:"image_url"`
		Introduction string `json:"introduction"`
		IslandType   string `json:"islandType"`
		Locale       string `json:"locale"`
		Matchmaking  struct {
			BAllowJoinInProgress   bool   `json:"bAllowJoinInProgress"`
			JoinInProgressTeam     int    `json:"joinInProgressTeam"`
			JoinInProgressType     string `json:"joinInProgressType"`
			MaximumNumberOfPlayers int    `json:"maximumNumberOfPlayers"`
			MinimumNumberOfPlayers int    `json:"minimumNumberOfPlayers"`
			MmsType                string `json:"mmsType"`
			NumberOfTeams          int    `json:"numberOfTeams"`
			OverridePlaylist       string `json:"override_Playlist"`
			PlayerCount            int    `json:"playerCount"`
			PlayersPerTeam         int    `json:"playersPerTeam"`
		} `json:"matchmaking"`
		SupportCode string `json:"supportCode"`
		Tagline     string `json:"tagline"`
		Title       string `json:"title"`
	} `json:"metadata"`
	Mnemonic         string    `json:"mnemonic"`
	ModerationStatus string    `json:"moderationStatus"`
	Namespace        string    `json:"namespace"`
	Published        time.Time `json:"published"`
	Version          int       `json:"version"`
}

func (c Client) FavoriteMnemonic(credentials UserCredentials, mnemonic string) (err error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	resp, err := c.Request("POST", fmt.Sprintf("%s/api/v1/links/favorites/%s/%s", consts.DISCOVERY_SERVICE, credentials.AccountID, mnemonic), headers, "{}")
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to favorite mnemonic: %s", resp.Status)
	}

	return nil
}

func (c Client) FetchMnemonicInfo(credentials UserCredentials, mnemonic string) (info MnemonicInfo, err error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	resp, err := c.Request("GET", fmt.Sprintf("%s/links/api/fn/mnemonic/%s", consts.LINKS_SERVICE, mnemonic), headers, "")
	if err != nil {
		return
	}

	res, err := request.ResponseParser[MnemonicInfo](resp)

	return res.Body, err
}

func (c Client) UnfavoriteMnemonic(credentials UserCredentials, mnemonic string) (err error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	resp, err := c.Request("DELETE", fmt.Sprintf("%s/api/v1/links/favorites/%s/%s", consts.DISCOVERY_SERVICE, credentials.AccountID, mnemonic), headers, "{}")
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to unfavorite mnemonic: %s", resp.Status)
	}

	return nil
}
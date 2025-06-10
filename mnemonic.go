package vinderman

import (
	"gitlab.com/8h9x/vinderman/consts"
	"gitlab.com/8h9x/vinderman/request"
	"net/http"
	"time"
)

type GetMnemonicResponse struct {
	Namespace   string `json:"namespace"`
	AccountId   string `json:"accountId"`
	CreatorName string `json:"creatorName"`
	Mnemonic    string `json:"mnemonic"`
	LinkType    string `json:"linkType"`
	Metadata    struct {
		OwnershipToken           string `json:"ownership_token"`
		LobbyBackgroundImageUrls struct {
			Url string `json:"url"`
		} `json:"lobby_background_image_urls"`
		BlogCategory    string            `json:"blog_category"`
		ImageUrl        string            `json:"image_url"`
		AltIntroduction map[string]string `json:"alt_introduction"`
		ImageUrls       struct {
			UrlS string `json:"url_s"`
			UrlM string `json:"url_m"`
			Url  string `json:"url"`
		} `json:"image_urls"`
		Locale      string `json:"locale"`
		Title       string `json:"title"`
		Matchmaking struct {
			JoinInProgressType     string `json:"joinInProgressType"`
			PlayersPerTeam         int    `json:"playersPerTeam"`
			MaximumNumberOfPlayers int    `json:"maximumNumberOfPlayers"`
			OverridePlaylist       string `json:"override_Playlist"`
			PlayerCount            int    `json:"playerCount"`
			MmsType                string `json:"mmsType"`
			MmsPrivacy             string `json:"mmsPrivacy"`
			NumberOfTeams          int    `json:"numberOfTeams"`
			BAllowJoinInProgress   bool   `json:"bAllowJoinInProgress"`
			MinimumNumberOfPlayers int    `json:"minimumNumberOfPlayers"`
			JoinInProgressTeam     int    `json:"joinInProgressTeam"`
		} `json:"matchmaking"`
		VideoVUID           string   `json:"video_vuid"`
		DisallowedPlatforms []string `json:"disallowedPlatforms"`
		AltImageUrls        map[string]struct {
			UrlS string `json:"url_s"`
			UrlM string `json:"url_m"`
			Url  string `json:"url"`
		} `json:"alt_image_urls"`
		AltTitle   map[string]string `json:"alt_title"`
		AltTagline map[string]string `json:"alt_tagline"`
		ProductTag string            `json:"product_tag"`
		Ratings    struct {
			RatingReceivedTime time.Time `json:"rating_received_time"`
			Boards             map[string]struct {
				Descriptors         []string `json:"descriptors"`
				RatingOverridden    bool     `json:"rating_overridden"`
				Rating              string   `json:"rating"`
				InitialRating       string   `json:"initial_rating"`
				InteractiveElements []string `json:"interactive_elements"`
			} `json:"boards"`
		} `json:"ratings"`
		Tagline   string `json:"tagline"`
		DynamicXp struct {
			UniqueGameVersion string `json:"uniqueGameVersion"`
			CalibrationPhase  string `json:"calibrationPhase"`
		} `json:"dynamicXp"`
		SquareImageUrls struct {
			Url string `json:"url"`
		} `json:"square_image_urls"`
		Introduction string `json:"introduction"`
	} `json:"metadata"`
	Version          int       `json:"version"`
	Active           bool      `json:"active"`
	Disabled         bool      `json:"disabled"`
	Created          time.Time `json:"created"`
	Published        time.Time `json:"published"`
	DescriptionTags  []string  `json:"descriptionTags"`
	ModerationStatus string    `json:"moderationStatus"`
	DiscoveryIntent  string    `json:"discoveryIntent"`
}

func (c *Client) GetMnemonic(mnemonic string) (GetMnemonicResponse, error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.CredentialsMap[c.ClientID].AccessToken)

	resp, err := c.Request("GET", consts.LinksService+"/links/api/fn/mnemonic/"+mnemonic, headers, "")
	if err != nil {
		return GetMnemonicResponse{}, err
	}

	res, err := request.ResponseParser[GetMnemonicResponse](resp)
	if err != nil {
		return GetMnemonicResponse{}, err
	}

	return res.Body, err
}

// GetRelatedMnemonics TODO: Actually implement the struct for this
func (c *Client) GetRelatedMnemonics(mnemonic string) (GetMnemonicResponse, error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.CredentialsMap[c.ClientID].AccessToken)

	resp, err := c.Request("GET", consts.LinksService+"/links/api/fn/mnemonic/"+mnemonic+"/related", headers, "")
	if err != nil {
		return GetMnemonicResponse{}, err
	}

	res, err := request.ResponseParser[GetMnemonicResponse](resp)
	if err != nil {
		return GetMnemonicResponse{}, err
	}

	return res.Body, err
}

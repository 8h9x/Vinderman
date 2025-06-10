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
		DynamicXP struct {
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

type GetRelatedMnemonicsResponse struct {
	ParentLinks []struct {
		Namespace   string `json:"namespace"`
		AccountId   string `json:"accountId"`
		CreatorName string `json:"creatorName"`
		Mnemonic    string `json:"mnemonic"`
		LinkType    string `json:"linkType"`
		Metadata    struct {
			BlogCategory string   `json:"blog_category"`
			ImageUrl     string   `json:"image_url"`
			GenreLabels  []string `json:"genre_labels"`
			ImageUrls    struct {
				UrlS string `json:"url_s"`
				UrlM string `json:"url_m"`
				Url  string `json:"url"`
			} `json:"image_urls"`
			Title        string            `json:"title"`
			Locale       string            `json:"locale"`
			VideoVUID    string            `json:"video_vuid"`
			SubLinkCodes []string          `json:"sub_link_codes"`
			AltTitle     map[string]string `json:"alt_title"`
			AltTagline   map[string]string `json:"alt_tagline"`
			Ratings      struct {
				RatingReceivedTime time.Time `json:"rating_received_time"`
				Boards             map[string]struct {
					Descriptors         []string `json:"descriptors"`
					RatingOverridden    bool     `json:"rating_overridden"`
					Rating              string   `json:"rating"`
					InitialRating       string   `json:"initial_rating"`
					InteractiveElements []string `json:"interactive_elements"`
				} `json:"boards"`
			} `json:"ratings"`
			ProductTag    string `json:"product_tag"`
			FallbackLinks struct {
				Graceful string `json:"graceful"`
			} `json:"fallback_links"`
			Tagline        string `json:"tagline"`
			ExtraImageUrls []struct {
				UrlS string `json:"url_s"`
				UrlM string `json:"url_m"`
				Url  string `json:"url"`
			} `json:"extra_image_urls,omitempty"`
			CategoryLabels    []string `json:"category_labels"`
			CorrespondingSets struct {
				Ranked   string `json:"ranked,omitempty"`
				Unranked string `json:"unranked,omitempty"`
			} `json:"corresponding_sets"`
			DefaultSubLinkCode string `json:"default_sub_link_code"`
			UnlockConditions   struct {
				AllOf struct {
					Conditions []struct {
						Id   string `json:"id"`
						Type string `json:"type"`
					} `json:"conditions"`
				} `json:"allOf"`
				PartyEligibility string `json:"partyEligibility"`
			} `json:"unlockConditions,omitempty"`
			MatchmakingV2 struct {
				RatingType string `json:"ratingType"`
				IsRanked   bool   `json:"isRanked"`
			} `json:"matchmakingV2,omitempty"`
			AltImageUrls map[string]struct {
				UrlS string `json:"url_s"`
				UrlM string `json:"url_m"`
				Url  string `json:"url"`
			} `json:"alt_image_urls"`
		} `json:"metadata"`
		Version           int           `json:"version"`
		Active            bool          `json:"active"`
		Disabled          bool          `json:"disabled"`
		Created           time.Time     `json:"created"`
		Published         time.Time     `json:"published"`
		DescriptionTags   []interface{} `json:"descriptionTags"`
		ModerationStatus  string        `json:"moderationStatus"`
		LastActivatedDate time.Time     `json:"lastActivatedDate"`
		DiscoveryIntent   string        `json:"discoveryIntent"`
	} `json:"parentLinks"`
	Links map[string]struct {
		Namespace   string `json:"namespace"`
		AccountId   string `json:"accountId"`
		CreatorName string `json:"creatorName"`
		Mnemonic    string `json:"mnemonic"`
		LinkType    string `json:"linkType"`
		Metadata    struct {
			ParentSet           string            `json:"parent_set"`
			FavoriteOverride    string            `json:"favorite_override"`
			PlayHistoryOverride string            `json:"play_history_override"`
			AltTitle            map[string]string `json:"alt_title"`
			ImageUrl            string            `json:"image_url"`
			ProductTag          string            `json:"product_tag"`
			ImageUrls           struct {
				UrlS  string `json:"url_s"`
				UrlXS string `json:"url_xs"`
				UrlM  string `json:"url_m"`
				Url   string `json:"url"`
			} `json:"image_urls"`
			DynamicXP struct {
				UniqueGameVersion int    `json:"uniqueGameVersion"`
				CalibrationPhase  string `json:"calibrationPhase"`
			} `json:"dynamicXp"`
			Matchmaking struct {
				OverridePlaylist string `json:"override_playlist"`
			} `json:"matchmaking"`
			VideoVuid string `json:"video_vuid"`
			Title     string `json:"title"`
		} `json:"metadata"`
		Version          int           `json:"version"`
		Active           bool          `json:"active"`
		Disabled         bool          `json:"disabled"`
		Created          time.Time     `json:"created"`
		Published        time.Time     `json:"published"`
		DescriptionTags  []interface{} `json:"descriptionTags"`
		ModerationStatus string        `json:"moderationStatus"`
	} `json:"links"`
}

func (c *Client) GetRelatedMnemonics(mnemonic string) (GetRelatedMnemonicsResponse, error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.CredentialsMap[c.ClientID].AccessToken)

	resp, err := c.Request("GET", consts.LinksService+"/links/api/fn/mnemonic/"+mnemonic+"/related", headers, "")
	if err != nil {
		return GetRelatedMnemonicsResponse{}, err
	}

	res, err := request.ResponseParser[GetRelatedMnemonicsResponse](resp)
	if err != nil {
		return GetRelatedMnemonicsResponse{}, err
	}

	return res.Body, err
}

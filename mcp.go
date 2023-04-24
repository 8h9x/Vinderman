package vinderman

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/0xDistrust/Vinderman/consts"
)

type ProfileStatsType interface {
	AthenaProfileStats | CampaignProfileStats | CollectionBookPeopleProfileStats | CollectionBookSchematicsProfileStats | CollectionsProfileStats | CommonPublicProfileStats |
		CommonCoreProfileStats | CreativeProfileStats | MetadataProfileStats | OutpostProfileStats | RecycleBinProfileStats | Theater0ProfileStats | Theater1ProfileStats | Theater2ProfileStats
}

type Profile[T ProfileStatsType] struct {
	ProfileRevision            int    `json:"profileRevision"`
	ProfileID                  string `json:"profileId"`
	ProfileChangesBaseRevision int    `json:"profileChangesBaseRevision"`
	ProfileChanges             []struct {
		ChangeType string `json:"changeType"`
		Profile    struct {
			Created         string                     `json:"created"`
			Updated         string                     `json:"updated"`
			RVN             int                        `json:"rvn"`
			WipeNumber      int                        `json:"wipeNumber"`
			AccountId       string                     `json:"accountId"`
			ProfileId       string                     `json:"profileId"`
			Version         string                     `json:"version"`
			Items           map[string]json.RawMessage `json:"items"`
			Stats           T                          `json:"stats"`
			CommandRevision int                        `json:"commandRevision"`
		} `json:"profile"`
	} `json:"profileChanges"`
	ProfileCommandRevision int       `json:"profileCommandRevision"`
	ServerTime             time.Time `json:"serverTime"`
	ResponseVersion        int       `json:"responseVersion"`
}

func (c Client) ComposeProfileOperation(credentials UserCredentials, operation string, profileID string, payload string) (resp *http.Response, err error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	resp, err = c.Request("POST", fmt.Sprintf("%s/profile/%s/client/%s?profileId=%s&rvn=-1", consts.FORTNITE_GAME, credentials.AccountID, operation, profileID), headers, payload)
	return
}

func (c Client) ProfileOperation(credentials UserCredentials, operation string, profileId string, payload any) (resp *http.Response, err error) {
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return
	}

	return c.ComposeProfileOperation(credentials, operation, profileId, string(bodyBytes))
}

func (c Client) QueryProfile(credentials UserCredentials, profileId string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "QueryProfile", profileId, "{}")
}

//func ProfileOperationExample() () {
//	client := New()
//
//	res, err := client.QueryProfile(UserCredentials{}, "athena")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	data, err := request.ResponseParser[Profile[AthenaProfileStats]](res)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var skins []AthenaCosmeticItem
//
//	for _, item := range data.Body.ProfileChanges[0].Profile.Items {
//		var cosmetic AthenaCosmeticItem
//		if err = json.Unmarshal(item, &cosmetic); err != nil {
//			// not a skin; (you should probably add an additional check to ensure that it isnt some other type of error occurring); TODO: abstract this to a helper function that properly error checks and returns an empty state of the type passed if the type of data doesnt match
//			continue
//		}
//
//		if strings.HasPrefix(cosmetic.TemplateID, "AthenaCharacter") {
//			skins = append(skins, cosmetic)
//		}
//	}
//
//	log.Println("Account Level:", data.Body.ProfileChanges[0].Profile.Stats.Attributes.AccountLevel)
//	log.Println("Skin Count:", len(skins))
//}
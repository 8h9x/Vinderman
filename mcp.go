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

type Profile[ST ProfileStatsType, NT CampaignNotifications | []interface{}] struct {
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
			AccountID       string                     `json:"accountId"`
			ProfileID       string                     `json:"profileId"`
			Version         string                     `json:"version"`
			Items           map[string]json.RawMessage `json:"items"`
			Stats           ST                         `json:"stats"`
			CommandRevision int                        `json:"commandRevision"`
		} `json:"profile"`
	} `json:"profileChanges"`
	ProfileCommandRevision int       `json:"profileCommandRevision"`
	ServerTime             time.Time `json:"serverTime"`
	ResponseVersion        int       `json:"responseVersion"`
	Notifications          NT        `json:"notifications"`
}

type AbandonExpeditionPayload struct {
	ExpeditionID string `json:"expeditionId"`
}

type ActivateConsumablePayload struct {
	TargetItemID    string `json:"targetItemId"`
	TargetAccountID string `json:"targetAccountId"`
}

type AddToCollectionPayload struct {
	Category    string              `json:"category"`
	Variant     string              `json:"variant"`
	ContextTags []string            `json:"contextTags"`
	Properties  interface{}         `json:"properties"`
	SeenState   EFortCollectedState `json:"seenState"`
	Count       int                 `json:"count"`
}

type ApplyVotePayload struct {
	OfferID string `json:"offerId"`
}

type AssignGadgetToLoadoutPayload struct {
	GadgetID  string `json:"gadgetId"`
	LoadoutID string `json:"loadoutId"`
	SlotIndex uint8  `json:"slotIndex"` // either 0 or 1
}

type AssignHeroToLoadoutPayload struct {
	HeroID    string `json:"heroId"`
	LoadoutID string `json:"loadoutId"`
	SlotIndex uint8  `json:"slotIndex"` // either 0 or 1
}

type AssignTeamPerkToLoadoutPayload struct {
	TeamPerkID string `json:"teamPerkId"`
	LoadoutID  string `json:"loadoutId"`
}

type AssignWorkerToSquadPayload struct {
	CharacterID string         `json:"characterId"`
	SquadID     SquadAttribute `json:"squadId"`
	SlotIndex   uint8          `json:"slotIndex"`
}

type AssignWorkerToSquadBatchPayload struct {
	CharacterIDs []string         `json:"characterIds"`
	SquadIDs     []SquadAttribute `json:"squadIds"`
	SlotIndices  []uint8          `json:"slotIndices"`
}

type AthenaPinQuestPayload struct {
	PinnedQuest string `json:"pinnedQuest"`
}

type AthenaRemoveQuestsPayload struct {
	RemovedQuests []string `json:"removedQuests"`
}

type BulkUpdateCollectionsPayload struct {
	Items []AddToCollectionPayload `json:"items"`
}

type CancelOrResumeSubscriptionPayload struct {
	AppStore             string `json:"appStore"`
	UniqueSubscriptionId string `json:"uniqueSubscriptionId"`
	WillAutoRenew        bool   `json:"willAutoRenew"`
}

type ChallengeBundleLevelUpPayload struct {
	BundleIdToLevel string `json:"bundleIdToLevel"`
}

type ClaimCollectedResourcesPayload struct {
	CollectorsToClaim []string `json:"collectorsToClaim"`
}

type ClaimCollectionBookPageRewardsPayload struct {
	PageTemplateID      string `json:"pageTemplateId"`
	SectionID           string `json:"sectionId"`
	SelectedRewardIndex int    `json:"selectedRewardIndex"`
}

type ClaimCollectionBookRewardsPayload struct {
	RequiredXP          int `json:"requiredXp"`
	SelectedRewardIndex int `json:"selectedRewardIndex"`
}

type ClaimImportFriendsRewardPayload struct {
	Network ESocialImportPanelPlatform `json:"network"`
}

type ClaimMFAEnabledPayload struct {
	ClaimForSTW bool `json:"bClaimForStw"`
}

type ClaimQuestRewardPayload struct {
	QuestID             string `json:"questId"`
	SelectedRewardIndex int    `json:"selectedRewardIndex"`
}

type ClaimSubscriptionRewardsPayload struct {
	AppStore             string `json:"appStore"`
	UniqueSubscriptionId string `json:"uniqueSubscriptionId"`
	ReceiptInfo          string `json:"receiptInfo"`
}

type ClearHeroLoadoutPayload struct {
	LoadoutID string `json:"loadoutId"`
}

type ClientQuestLoginPayload struct {
	StreamingAppKey string `json:"streamingAppKey"`
}

type CollectExpeditionPayload struct {
	ExpeditionTemplate string `json:"expeditionTemplate"`
	ExpeditionId       string `json:"expeditionId"`
}

type CompletePlayerSurveyPayload struct {
	SurveyID                 string `json:"surveyId"`
	UpdateAllSurveysMetadata bool   `json:"bUpdateAllSurveysMetadata"`
}

type ConvertItemPayload struct {
	TargetItemID    string `json:"targetItemId"`
	ConversionIndex uint8  `json:"conversionIndex"`
}

type ConvertSlottedItemPayload struct {
	TargetItemID    string `json:"targetItemId"`
	ConversionIndex uint8  `json:"conversionIndex"`
}

type CopyCosmeticLoadoutPayload struct {
	SourceIndex         int    `json:"sourceIndex"`
	TargetIndex         int    `json:"targetIndex"`
	OptNewNameForTarget string `json:"optNewNameForTarget"`
}

type CraftWorldItemPayload struct {
	TargetSchematicItemID string `json:"targetSchematicItemId"`
	NumTimesToCraft       int    `json:"numTimesToCraft"`
	TargetSchematicTier   string `json:"targetSchematicTier"`
}

type DeleteCosmeticLoadoutPayload struct {
	Index                int  `json:"index"`
	FallbackLoadoutIndex int  `json:"fallbackLoadoutIndex"`
	LeaveNullSlot        bool `json:"leaveNullSlot"`
}

type DestroyWorldItemsPayload struct {
	ItemIDs []string `json:"itemIds"`
}

type DisassembleWorldItemsPayload struct {
	TargetItemIdAndQuantityPairs []struct {
		ItemID   string `json:"itemId"`
		Quantity int    `json:"quantity"`
	} `json:"targetItemIdAndQuantityPairs"`
}

type ExchangeGameCurrencyForBattlePassOfferPayload struct {
	OfferItemIDList []string `json:"offerItemIdList"`
}

func (c Client) ComposeProfileOperation(credentials UserCredentials, operation string, profileID string, payload string) (resp *http.Response, err error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	resp, err = c.Request("POST", fmt.Sprintf("%s/profile/%s/client/%s?profileId=%s&rvn=-1", consts.FORTNITE_GAME, credentials.AccountID, operation, profileID), headers, payload)
	return
}

func (c Client) ProfileOperation(credentials UserCredentials, operation string, profileID string, payload any) (resp *http.Response, err error) {
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return
	}

	return c.ComposeProfileOperation(credentials, operation, profileID, string(bodyBytes))
}

func (c Client) AbandonExpedition(credentials UserCredentials, expeditionId string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "AbandonExpedition", "campaign", AbandonExpeditionPayload{ExpeditionID: expeditionId})
}

func (c Client) ActivateConsumable(credentials UserCredentials, payload ActivateConsumablePayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "ActivateConsumable", "campaign", payload)
}

func (c Client) AddToCollection(credentials UserCredentials, payload AddToCollectionPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "AddToCollection", "collections", payload)
}

func (c Client) ApplyVote(credentials UserCredentials, offerID string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "ApplyVote", "athena", ApplyVotePayload{OfferID: offerID})
}

func (c Client) AssignGadgetToLoadout(credentials UserCredentials, payload AssignGadgetToLoadoutPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "AssignGadgetToLoadout", "campaign", payload)
}

func (c Client) AssignHeroToLoadout(credentials UserCredentials, payload AssignHeroToLoadoutPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "AssignHeroToLoadout", "campaign", payload)
}

func (c Client) AssignTeamPerkToLoadout(credentials UserCredentials, payload AssignTeamPerkToLoadoutPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "AssignTeamPerkToLoadout", "campaign", payload)
}

func (c Client) AssignWorkerToSquad(credentials UserCredentials, payload AssignWorkerToSquadPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "AssignWorkerToSquad", "campaign", payload)
}

func (c Client) AssignWorkerToSquadBatch(credentials UserCredentials, payload AssignWorkerToSquadBatchPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "AssignWorkerToSquadBatch", "campaign", payload)
}

func (c Client) AthenaPinQuest(credentials UserCredentials, pinnedQuest string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "AthenaPinQuest", "athena", AthenaPinQuestPayload{PinnedQuest: pinnedQuest})
}

func (c Client) AthenaRemoveQuests(credentials UserCredentials, removedQuests []string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "AthenaRemoveQuests", "athena", AthenaRemoveQuestsPayload{RemovedQuests: removedQuests})
}

func (c Client) BulkUpdateCollections(credentials UserCredentials, payload BulkUpdateCollectionsPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "BulkUpdateCollections", "collections", payload)
}

func (c Client) CancelOrResumeSubscription(credentials UserCredentials, payload CancelOrResumeSubscriptionPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "CancelOrResumeSubscription", "common_core", payload)
}

func (c Client) ChallengeBundleLevelUp(credentials UserCredentials, bundleIdToLevel string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "ChallengeBundleLevelUp", "athena", ChallengeBundleLevelUpPayload{BundleIdToLevel: bundleIdToLevel})
}

func (c Client) ClaimCollectedResources(credentials UserCredentials, collectorsToClaim []string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "ClaimCollectedResources", "campaign", ClaimCollectedResourcesPayload{CollectorsToClaim: collectorsToClaim})
}

// ClaimCollectionBookPageRewards collection_book_people0, collection_book_schematics0
func (c Client) ClaimCollectionBookPageRewards(credentials UserCredentials, profileID string, payload ClaimCollectionBookPageRewardsPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "ClaimCollectionBookPageRewards", profileID, payload)
}

func (c Client) ClaimCollectionBookRewards(credentials UserCredentials, payload ClaimCollectionBookRewardsPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "ClaimCollectionBookPageRewards", "campaign", payload)
}

func (c Client) ClaimImportFriendsReward(credentials UserCredentials, network ESocialImportPanelPlatform) (*http.Response, error) {
	return c.ProfileOperation(credentials, "ClaimImportFriendsReward", "common_core", ClaimImportFriendsRewardPayload{Network: network})
}

func (c Client) ClaimLoginReward(credentials UserCredentials) (*http.Response, error) {
	return c.ComposeProfileOperation(credentials, "ClaimLoginReward", "campaign", "{}")
}

func (c Client) ClaimMFAEnabled(credentials UserCredentials, claimForSTW bool) (*http.Response, error) {
	return c.ProfileOperation(credentials, "ClaimMfaEnabled", "common_core", ClaimMFAEnabledPayload{ClaimForSTW: claimForSTW})
}

func (c Client) ClaimMissionAlertRewards(credentials UserCredentials) (*http.Response, error) {
	return c.ComposeProfileOperation(credentials, "ClaimMissionAlertRewards", "campaign", "{}")
}

// ClaimQuestReward athena, campaign
func (c Client) ClaimQuestReward(credentials UserCredentials, profileID string, payload ClaimQuestRewardPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "ClaimQuestReward", profileID, payload)
}

func (c Client) ClaimSubscriptionRewards(credentials UserCredentials, payload ClaimSubscriptionRewardsPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "ClaimSubscriptionRewards", "common_core", payload)
}

func (c Client) ClearHeroLoadout(credentials UserCredentials, loadoutID string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "ClearHeroLoadout", "campaign", ClearHeroLoadoutPayload{LoadoutID: loadoutID})
}

// ClientQuestLogin athena, campaign
func (c Client) ClientQuestLogin(credentials UserCredentials, profileID string, streamingAppKey string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "ClientQuestLogin", profileID, ClientQuestLoginPayload{StreamingAppKey: streamingAppKey})
}

func (c Client) CollectExpedition(credentials UserCredentials, payload CollectExpeditionPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "CollectExpedition", "campaign", payload)
}

func (c Client) CompletePlayerSurvey(credentials UserCredentials, payload CompletePlayerSurveyPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "CompletePlayerSurvey", "common_core", payload)
}

func (c Client) ConvertItem(credentials UserCredentials, payload ConvertItemPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "ConvertItem", "campaign", payload)
}

// ConvertSlottedItem collection_book_people0, collection_book_schematics0
func (c Client) ConvertSlottedItem(credentials UserCredentials, profileID string, payload ConvertSlottedItemPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "ConvertSlottedItem", profileID, payload)
}

// CopyCosmeticLoadout athena, campaign
func (c Client) CopyCosmeticLoadout(credentials UserCredentials, profileID string, payload CopyCosmeticLoadoutPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "CopyCosmeticLoadout", profileID, payload)
}

func (c Client) CraftWorldItem(credentials UserCredentials, payload CraftWorldItemPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "CraftWorldItem", "theater0", payload)
}

func (c Client) DeleteBattleLabIsland(credentials UserCredentials) (*http.Response, error) {
	return c.ComposeProfileOperation(credentials, "DeleteBattleLabIsland", "creative", "{}")
}

// DeleteCosmeticLoadout athena, campaign
func (c Client) DeleteCosmeticLoadout(credentials UserCredentials, profileID string, payload DeleteCosmeticLoadoutPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "DeleteCosmeticLoadout", profileID, payload)
}

// DestroyWorldItems outpost0, theater0, theater1, theater2
func (c Client) DestroyWorldItems(credentials UserCredentials, profileID string, itemIDs []string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "DestroyWorldItems", profileID, DestroyWorldItemsPayload{ItemIDs: itemIDs})
}

// DisassembleWorldItems theater0, theater1, theater2
func (c Client) DisassembleWorldItems(credentials UserCredentials, profileID string, payload DisassembleWorldItemsPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "DisassembleWorldItems", profileID, payload)
}

func (c Client) ExchangeGameCurrencyForBattlePassOffer(credentials UserCredentials, offerItemIDList []string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "ExchangeGameCurrencyForBattlePassOffer", "athena", ExchangeGameCurrencyForBattlePassOfferPayload{OfferItemIDList: offerItemIDList})
}

func (c Client) ExchangeGiftToken(credentials UserCredentials) (*http.Response, error) {
	return c.ComposeProfileOperation(credentials, "ExchangeGiftToken", "athena", "{}")
}

func (c Client) QueryProfile(credentials UserCredentials, profileID string) (*http.Response, error) {
	return c.ComposeProfileOperation(credentials, "QueryProfile", profileID, "{}")
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
//			// not a skin; (you should probably add an additional check to ensure that it isn't some other type of error occurring); TODO: abstract this to a helper function that properly error checks and returns an empty state of the type passed if the type of data doesnt match
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

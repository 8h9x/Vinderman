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

type FortRerollDailyQuestPayload struct {
	QuestID string `json:"questId"`
}

type GiftCatalogEntryPayload struct {
	OfferID            string   `json:"offerId"`
	Currency           string   `json:"currency"`
	CurrencySubType    string   `json:"currencySubType"`
	ExpectedTotalPrice int      `json:"expectedTotalPrice"`
	GameContext        string   `json:"gameContext"`
	ReceiverAccountIds []string `json:"receiverAccountIds"`
	GiftWrapTemplateID string   `json:"giftWrapTemplateId"`
	PersonalMessage    string   `json:"personalMessage"`
}

type InitializeTheaterPayload struct {
	TheaterGUID string `json:"theaterGuid"`
}

type IssueFriendCodePayload struct {
	CodeTokenType string `json:"codeTokenType"`
}

type MarkCollectedItemsSeenPayload struct {
	Variants []MarkCollectedItemsSeenPayloadVariants `json:"variants"`
}

type MarkCollectedItemsSeenPayloadVariants struct {
	Category string `json:"category"`
	Variant  string `json:"variant"`
}

type MarkItemSeenPayload struct {
	ItemIDs []string `json:"itemIds"`
}

type MarkNewQuestNotificationSentPayload struct {
	ItemIDs []string `json:"itemIds"`
}

type ModifyQuickbarPayload struct {
	PrimaryQuickbarChoices  []string `json:"primaryQuickbarChoices"`
	SecondaryQuickbarChoice string   `json:"secondaryQuickbarChoice"`
}

type OpenCardPackPayload struct {
	CardPackItemID string `json:"cardPackItemId"`
	SelectionIndex int    `json:"selectionIdx"`
}

type OpenCardPackBatchPayload struct {
	CardPackItemIDs []string `json:"cardPackItemIds"`
}

type PromoteItemPayload struct {
	TargetItemID string `json:"targetItemId"`
}

type PurchaseCatalogEntryPayload struct {
	OfferID            string `json:"offerId"`
	PurchaseQuantity   int    `json:"purchaseQuantity"`
	Currency           string `json:"currency"`
	CurrencySubType    string `json:"currencySubType"`
	ExpectedTotalPrice int    `json:"expectedTotalPrice"`
	GameContext        string `json:"gameContext"`
}

type PurchaseMultipleCatalogEntriesPayload struct {
	PurchaseInfoList []PurchaseCatalogEntryPayload `json:"purchaseInfoList"`
}

type PurchaseOrUpgradeHomebaseNodePayload struct {
	NodeID string `json:"nodeId"`
}

type PurchaseResearchStatUpgradePayload struct {
	StatID string `json:"statId"`
}

type RecycleItemPayload struct {
	TargetItemID string `json:"targetItemId"`
}

type RecycleItemBatchPayload struct {
	TargetItemIDs []string `json:"targetItemIds"`
}

type RedeemRealMoneyPurchasesPayload struct {
	AppStore              string                                `json:"appStore"`
	AuthTokens            []string                              `json:"authTokens"`
	ReceiptIDs            []string                              `json:"receiptIds"`
	RefreshType           RealMoneyPurchaseRefreshType          `json:"refreshType"`
	VerifierModeOverride  RealMoneyPurchaseVerifierModeOverride `json:"verifierModeOverride"`
	PurchaseCorrelationID string                                `json:"purchaseCorrelationId"`
}

type RefundItemPayload struct {
	TargetItemID string `json:"targetItemId"`
}

type RefundMtxPurchasePayload struct {
	PurchaseID  string `json:"purchaseId"`
	QuickReturn bool   `json:"quickReturn"`
	GameContext string `json:"gameContext"`
}

type RemoveGiftBoxPayload struct {
	GiftBoxItemIDs []string `json:"giftBoxItemIds"`
}

type RequestRestedStateIncreasePayload struct {
	TimeToCompensateFor    int `json:"timeToCompensateFor"`
	RestedXPGenAccumulated int `json:"restedXpGenAccumulated"`
}

type ResearchItemFromCollectionBookPayload struct {
	TemplateID string `json:"templateId"`
}

type RespecAlterationPayload struct {
	TargetItemID   string `json:"targetItemId"`
	AlterationSlot int    `json:"alterationSlot"`
	AlterationID   string `json:"alterationId"`
}

type SetActiveHeroLoadoutPayload struct {
	SelectedLoadout string `json:"selectedLoadout"`
}

type SetAffiliateNamePayload struct {
	AffiliateName string `json:"affiliateName"`
}

type SetCosmeticLockerBannerPayload struct {
	LockerItem              string `json:"lockerItem"`
	BannerIconTemplateName  string `json:"bannerIconTemplateName"`
	BannerColorTemplateName string `json:"bannerColorTemplateName"`
}

type SetCosmeticLockerNamePayload struct {
	LockerItem string `json:"lockerItem"`
	Name       string `json:"name"`
}

type SetCosmeticLockerSlotPayload struct {
	LockerItem                string                                      `json:"lockerItem"`
	Category                  string                                      `json:"category"`
	ItemToSlot                string                                      `json:"itemToSlot"`
	SlotIndex                 uint8                                       `json:"slotIndex"`
	VariantUpdates            []SetCosmeticLockerSlotPayloadVariantUpdate `json:"variantUpdates"`
	OptLockerUseCountOverride int                                         `json:"optLockerUseCountOverride"`
}

type SetCosmeticLockerSlotPayloadVariantUpdate struct {
	Channel string   `json:"channel"`
	Active  string   `json:"active"`
	Owned   []string `json:"owned"`
}

type SetCosmeticLockerSlotsPayload struct {
	LockerItem  string                                     `json:"lockerItem"`
	LoadoutData []SetCosmeticLockerSlotsPayloadLoadoutData `json:"loadoutData"`
}

type SetCosmeticLockerSlotsPayloadLoadoutData struct {
	Category        string `json:"category"`
	ItemToSlot      string `json:"itemToSlot"`
	IndexWithinSlot uint8  `json:"indexWithinSlot"`
}

type SetForcedIntroPlayedPayload struct {
	ForcedIntroName string `json:"forcedIntroName"`
}

type SetHardcoreModifierPayload struct {
	Updates []SetHardcoreModifierPayloadUpdates `json:"updates"`
}

type SetHardcoreModifierPayloadUpdates struct {
	ModifierID string `json:"modifierId"`
	Enabled    bool   `json:"bEnabled"`
}

type SetHeroCosmeticVariantsPayload struct {
	HeroItem          string                                   `json:"heroItem"`
	OutfitVariants    []SetHeroCosmeticVariantsPayloadVariants `json:"outfitVariants"`
	BackblingVariants []SetHeroCosmeticVariantsPayloadVariants `json:"backblingVariants"`
}

type SetHeroCosmeticVariantsPayloadVariants struct {
	Channel string        `json:"channel"`
	Active  string        `json:"active"`
	Owned   []interface{} `json:"owned"` // TODO: proper type
}

type SetHomebaseBannerPayload struct {
	HomebaseBannerIconID  string `json:"homebaseBannerIconId"`
	HomebaseBannerColorID string `json:"homebaseBannerColorId"`
}

type SetHomebaseNamePayload struct {
	HomebaseName string `json:"homebaseName"`
}

type SetItemArchivedStatusBatchPayload struct {
	ItemIDs  []string `json:"itemIds"`
	Archived bool     `json:"archived"`
}

type SetItemFavoriteStatusPayload struct {
	TargetItemID string `json:"targetItemId"`
	Favorite     bool   `json:"bFavorite"`
}

type SetItemFavoriteStatusBatchPayload struct {
	ItemIDs       []string `json:"itemIds"`
	ItemFavStatus []bool   `json:"itemFavStatus"`
}

type SetMtxPlatformPayload struct {
	NewPlatform string `json:"newPlatform"`
}

type SetPinnedQuestsPayload struct {
	PinnedQuestIDs []string `json:"pinnedQuestIds"`
}

type SetRandomCosmeticLoadoutFlagPayload struct {
	Random bool `json:"random"`
}

type SetReceiveGiftsEnabledPayload struct {
	ReceiveGifts bool `json:"bReceiveGifts"`
}

type SetRewardGraphConfigPayload struct {
	State         []string `json:"state"`
	RewardGraphID string   `json:"rewardGraphId"`
}

type StartExpeditionPayload struct {
	ExpeditionID string   `json:"expeditionId"`
	SquadID      string   `json:"squadId"`
	ItemIDs      []string `json:"itemIds"`
	SlotIndices  []int    `json:"slotIndices"`
}

type StorageTransferPayload struct {
	TransferOperations []StorageTransferPayloadTransferOperations `json:"transferOperations"`
}

type StorageTransferPayloadTransferOperations struct {
	ItemID        string `json:"itemId"`
	Quantity      int    `json:"quantity"`
	ToStorage     bool   `json:"toStorage"`
	NewItemIdHint string `json:"newItemIdHint"`
}

type ToggleQuestActiveStatePayload struct {
	QuestIDs []string `json:"questIds"`
}

type UnassignAllSquadsPayload struct {
	SquadIDs []SquadAttribute `json:"squadIds"`
}

type UnlockRewardNodePayload struct {
	NodeID        string `json:"nodeId"`
	RewardGraphID string `json:"rewardGraphId"`
	RewardCFG     string `json:"rewardCfg"`
}

type UnslotItemFromCollectionBookPayload struct {
	TemplateID string `json:"templateId"`
	ItemID     string `json:"itemId"`
	Specific   string `json:"specific"`
}

type UpdateQuestClientObjectivesPayload struct {
	Advance []UpdateQuestClientObjectivesPayloadAdvance `json:"advance"`
}

type UpdateQuestClientObjectivesPayloadAdvance struct {
	StatName        string `json:"statName"`
	Count           int    `json:"count"`
	TimestampOffset int    `json:"timestampOffset"`
}

type UpgradeAlterationPayload struct {
	TargetItemID   string `json:"targetItemId"`
	AlterationSlot uint8  `json:"alterationSlot"`
}

type UpgradeItemPayload struct {
	TargetItemID string `json:"targetItemId"`
}

type UpgradeItemBulkPayload struct {
	TargetItemID                string `json:"targetItemId"`
	DesiredLevel                int    `json:"desiredLevel"`
	DesiredTier                 string `json:"desiredTier"`
	ConversionRecipeIndexChoice int    `json:"conversionRecipeIndexChoice"`
}

type UpgradeItemRarityPayload struct {
	TargetItemID string `json:"targetItemId"`
}

type UpgradeSlottedItemPayload struct {
	TargetItemID string `json:"targetItemId"`
	DesiredLevel int    `json:"desiredLevel"`
}

type VerifyRealMoneyPurchasePayload struct {
	AppStore              string `json:"appStore"`
	AppStoreID            string `json:"appStoreId"`
	ReceiptID             string `json:"receiptId"`
	ReceiptInfo           string `json:"receiptInfo"`
	PurchaseCorrelationID string `json:"purchaseCorrelationId"`
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

// FortRerollDailyQuest athena, campaign
func (c Client) FortRerollDailyQuest(credentials UserCredentials, profileID string, questID string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "FortRerollDailyQuest", profileID, FortRerollDailyQuestPayload{QuestID: questID})
}

func (c Client) GiftCatalogEntry(credentials UserCredentials, payload GiftCatalogEntryPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "GiftCatalogEntry", "common_core", payload)
}

func (c Client) IssueFriendCode(credentials UserCredentials, codeTokenType string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "IssueFriendCode", "common_core", IssueFriendCodePayload{CodeTokenType: codeTokenType})
}

func (c Client) MarkCollectedItemsSeen(credentials UserCredentials, payload MarkCollectedItemsSeenPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "MarkCollectedItemsSeen", "collections", payload)
}

// MarkItemSeen athena, campaign, common_core
func (c Client) MarkItemSeen(credentials UserCredentials, profileID string, itemIDs []string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "MarkItemSeen", profileID, MarkItemSeenPayload{ItemIDs: itemIDs})
}

// MarkNewQuestNotificationSent athena, campaign
func (c Client) MarkNewQuestNotificationSent(credentials UserCredentials, profileID string, itemIDs []string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "MarkNewQuestNotificationSent", profileID, MarkNewQuestNotificationSentPayload{ItemIDs: itemIDs})
}

// ModifyQuickbar theater0, theater1, theater2
func (c Client) ModifyQuickbar(credentials UserCredentials, profileID string, payload ModifyQuickbarPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "ModifyQuickbar", profileID, payload)
}

func (c Client) OpenCardPack(credentials UserCredentials, payload OpenCardPackPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "OpenCardPack", "campaign", payload)
}

func (c Client) OpenCardPackBatch(credentials UserCredentials, cardPackitemIDs []string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "OpenCardPackBatch", "campaign", OpenCardPackBatchPayload{CardPackItemIDs: cardPackitemIDs})
}

func (c Client) PopulatePrerolledOffers(credentials UserCredentials) (*http.Response, error) {
	return c.ComposeProfileOperation(credentials, "PopulatePrerolledOffers", "campaign", "{}")
}

// PromoteItem campaign, collection_book_people0, collection_book_schematics0
func (c Client) PromoteItem(credentials UserCredentials, profileID string, targetItemID string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "PromoteItem", profileID, PromoteItemPayload{TargetItemID: targetItemID})
}

func (c Client) PurchaseCatalogEntry(credentials UserCredentials, payload PurchaseCatalogEntryPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "PurchaseCatalogEntry", "common_core", payload)
}

func (c Client) PurchaseMultipleCatalogEntries(credentials UserCredentials, payload PurchaseMultipleCatalogEntriesPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "PurchaseMultipleCatalogEntries", "common_core", payload)
}

func (c Client) PurchaseOrUpgradeHomebaseNode(credentials UserCredentials, nodeID string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "PurchaseOrUpgradeHomebaseNode", "campaign", PurchaseOrUpgradeHomebaseNodePayload{NodeID: nodeID})
}

func (c Client) PurchaseResearchStatUpgrade(credentials UserCredentials, statID string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "PurchaseResearchStatUpgrade", "campaign", PurchaseResearchStatUpgradePayload{StatID: statID})
}

func (c Client) QueryProfile(credentials UserCredentials, profileID string) (*http.Response, error) {
	return c.ComposeProfileOperation(credentials, "QueryProfile", profileID, "{}")
}

// TODO: QueryPublicProfile

func (c Client) RecycleItem(credentials UserCredentials, targetItemID string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "RecycleItem", "campaign", RecycleItemPayload{TargetItemID: targetItemID})
}

func (c Client) RecycleItemBatch(credentials UserCredentials, targetItemIDs []string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "RecycleItemBatch", "campaign", RecycleItemBatchPayload{TargetItemIDs: targetItemIDs})
}

func (c Client) RedeemRealMoneyPurchases(credentials UserCredentials, payload RedeemRealMoneyPurchasesPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "RedeemRealMoneyPurchases", "common_core", payload)
}

func (c Client) RedeemSTWAccoladeTokens(credentials UserCredentials) (*http.Response, error) {
	return c.ComposeProfileOperation(credentials, "RedeemSTWAccoladeTokens", "athena", "{}")
}

func (c Client) RefreshExpeditions(credentials UserCredentials) (*http.Response, error) {
	return c.ComposeProfileOperation(credentials, "RefreshExpeditions", "campaign", "{}")
}

func (c Client) RefundItem(credentials UserCredentials, targetItemID string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "RefundItem", "campaign", RefundItemPayload{TargetItemID: targetItemID})
}

func (c Client) RefundMtxPurchase(credentials UserCredentials, payload RefundMtxPurchasePayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "RefundMtxPurchase", "common_core", payload)
}

func (c Client) RemoveGiftBox(credentials UserCredentials, profileID string, giftBoxItemIDs []string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "RemoveGiftBox", profileID, RemoveGiftBoxPayload{GiftBoxItemIDs: giftBoxItemIDs})
}

func (c Client) RequestRestedStateIncrease(credentials UserCredentials, payload RequestRestedStateIncreasePayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "RequestRestedStateIncrease", "athena", payload)
}

// ResearchItemFromCollectionBook campaign, theater0, theater1, theater2
func (c Client) ResearchItemFromCollectionBook(credentials UserCredentials, profileID string, templateID string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "ResearchItemFromCollectionBook", profileID, ResearchItemFromCollectionBookPayload{TemplateID: templateID})
}

func (c Client) RespecAlteration(credentials UserCredentials, payload RespecAlterationPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "RespecAlteration", "campaign", payload)
}

func (c Client) RespecResearch(credentials UserCredentials) (*http.Response, error) {
	return c.ComposeProfileOperation(credentials, "RespecResearch", "campaign", "{}")
}

func (c Client) RespecUpgrades(credentials UserCredentials) (*http.Response, error) {
	return c.ComposeProfileOperation(credentials, "RespecUpgrades", "campaign", "{}")
}

func (c Client) SetActiveHeroLoadout(credentials UserCredentials, selectedLoadout string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "SetActiveHeroLoadout", "campaign", SetActiveHeroLoadoutPayload{SelectedLoadout: selectedLoadout})
}

func (c Client) SetAffiliateNameLoadout(credentials UserCredentials, affiliateName string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "SetActiveHeroLoadout", "common_core", SetAffiliateNamePayload{AffiliateName: affiliateName})
}

// Deprecated: SetCosmeticLockerBanner Modern versions of the fortnite mcp system use a separate locker api--use locker.UpdateActiveLockerLoadout()
func (c Client) SetCosmeticLockerBanner(credentials UserCredentials, profileID string, payload SetCosmeticLockerBannerPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "SetCosmeticLockerBanner", profileID, payload)
}

// Deprecated: SetCosmeticLockerName Modern versions of the fortnite mcp system use a separate locker api--use locker.UpdateActiveLockerLoadout()
func (c Client) SetCosmeticLockerName(credentials UserCredentials, profileID string, payload SetCosmeticLockerNamePayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "SetCosmeticLockerName", profileID, payload)
}

// Deprecated: SetCosmeticLockerSlot Modern versions of the fortnite mcp system use a separate locker api--use locker.UpdateActiveLockerLoadout()
func (c Client) SetCosmeticLockerSlot(credentials UserCredentials, profileID string, payload SetCosmeticLockerSlotPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "SetCosmeticLockerSlot", profileID, payload)
}

// Deprecated: SetCosmeticLockerSlots Modern versions of the fortnite mcp system use a separate locker api--use locker.UpdateActiveLockerLoadout()
func (c Client) SetCosmeticLockerSlots(credentials UserCredentials, profileID string, payload SetCosmeticLockerSlotsPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "SetCosmeticLockerSlots", profileID, payload)
}

func (c Client) SetForcedIntroPlayed(credentials UserCredentials, forcedIntroName string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "SetForcedIntroPlayed", "common_core", SetForcedIntroPlayedPayload{ForcedIntroName: forcedIntroName})
}

func (c Client) SetHardcoreModifier(credentials UserCredentials, payload SetHardcoreModifierPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "SetHardcoreModifier", "athena", payload)
}

func (c Client) SetHeroCosmeticVariants(credentials UserCredentials, payload SetHeroCosmeticVariantsPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "SetHeroCosmeticVariants", "campaign", payload)
}

func (c Client) SetHomebaseBanner(credentials UserCredentials, payload SetHomebaseBannerPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "SetHomebaseBanner", "common_public", payload)
}

// Deprecated: SetHomebaseName The home base naming system has been removed from the game.
func (c Client) SetHomebaseName(credentials UserCredentials, homebaseName string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "SetHomebaseName", "common_core", SetHomebaseNamePayload{HomebaseName: homebaseName})
}

func (c Client) SetIntroGamePlayed(credentials UserCredentials) (*http.Response, error) {
	return c.ComposeProfileOperation(credentials, "SetIntroGamePlayed", "common_core", "{}")
}

func (c Client) SetItemArchivedStatusBatch(credentials UserCredentials, payload SetItemArchivedStatusBatchPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "SetItemArchivedStatusBatch", "athena", payload)
}

// SetItemFavoriteStatus athena, campaign
func (c Client) SetItemFavoriteStatus(credentials UserCredentials, profileID string, payload SetItemFavoriteStatusPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "SetItemFavoriteStatus", profileID, payload)
}

// SetItemFavoriteStatusBatch athena, campaign
func (c Client) SetItemFavoriteStatusBatch(credentials UserCredentials, profileID string, payload SetItemFavoriteStatusBatchPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "SetItemFavoriteStatusBatch", profileID, payload)
}

func (c Client) SetMatchmakingBansViewed(credentials UserCredentials) (*http.Response, error) {
	return c.ComposeProfileOperation(credentials, "SetMatchmakingBansViewed", "common_core", "{}")
}

func (c Client) SetMtxPlatform(credentials UserCredentials, newPlatform string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "SetMtxPlatform", "common_core", SetMtxPlatformPayload{NewPlatform: newPlatform})
}

func (c Client) SetPinnedQuests(credentials UserCredentials, pinnedQuestIDs []string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "SetPinnedQuests", "common_core", SetPinnedQuestsPayload{PinnedQuestIDs: pinnedQuestIDs})
}

// SetRandomCosmeticLoadoutFlag athena, campaign
func (c Client) SetRandomCosmeticLoadoutFlag(credentials UserCredentials, profileID string, random bool) (*http.Response, error) {
	return c.ProfileOperation(credentials, "SetRandomCosmeticLoadoutFlag", profileID, SetRandomCosmeticLoadoutFlagPayload{Random: random})
}

func (c Client) SetReceiveGiftsEnabled(credentials UserCredentials, receiveGifts bool) (*http.Response, error) {
	return c.ProfileOperation(credentials, "SetReceiveGiftsEnabled", "common_core", SetReceiveGiftsEnabledPayload{ReceiveGifts: receiveGifts})
}

func (c Client) SetRewardGraphConfig(credentials UserCredentials, payload SetRewardGraphConfigPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "SetRewardGraphConfig", "athena", payload)
}

func (c Client) StartExpedition(credentials UserCredentials, payload StartExpeditionPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "StartExpedition", "campaign", payload)
}

func (c Client) StorageTransfer(credentials UserCredentials, payload StorageTransferPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "StorageTransfer", "theater0", payload)
}

func (c Client) ToggleQuestActiveState(credentials UserCredentials, questIDs []string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "ToggleQuestActiveState", "athena", ToggleQuestActiveStatePayload{QuestIDs: questIDs})
}

func (c Client) UnassignAllSquads(credentials UserCredentials, squadIDs []SquadAttribute) (*http.Response, error) {
	return c.ProfileOperation(credentials, "UnassignAllSquads", "campaign", UnassignAllSquadsPayload{SquadIDs: squadIDs})
}

func (c Client) UnlockRewardNode(credentials UserCredentials, payload UnlockRewardNodePayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "UnlockRewardNode", "athena", payload)
}

// UnslotItemFromCollectionBook campaign, theater0, theater1, theater2
func (c Client) UnslotItemFromCollectionBook(credentials UserCredentials, profileID string, payload UnslotItemFromCollectionBookPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "UnslotItemFromCollectionBook", profileID, payload)
}

// UpdateQuestClientObjectives athena, campaign
func (c Client) UpdateQuestClientObjectives(credentials UserCredentials, profileID string, payload UpdateQuestClientObjectivesPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "UpdateQuestClientObjectives", profileID, payload)
}

func (c Client) UpgradeAlteration(credentials UserCredentials, payload UpgradeAlterationPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "UpgradeAlteration", "campaign", payload)
}

func (c Client) UpgradeItem(credentials UserCredentials, targetItemID string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "UpgradeItem", "campaign", UpgradeItemPayload{TargetItemID: targetItemID})
}

func (c Client) UpgradeItemBulk(credentials UserCredentials, payload UpgradeItemBulkPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "UpgradeItemBulk", "campaign", payload)
}

func (c Client) UpgradeItemRarity(credentials UserCredentials, targetItemID string) (*http.Response, error) {
	return c.ProfileOperation(credentials, "UpgradeItemRarity", "campaign", UpgradeItemRarityPayload{TargetItemID: targetItemID})
}

// UpgradeSlottedItem collection_book_people0, collection_book_schematics0
func (c Client) UpgradeSlottedItem(credentials UserCredentials, profileID string, payload UpgradeSlottedItemPayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "UpgradeSlottedItem", profileID, payload)
}

func (c Client) VerifyRealMoneyPurchase(credentials UserCredentials, payload VerifyRealMoneyPurchasePayload) (*http.Response, error) {
	return c.ProfileOperation(credentials, "VerifyRealMoneyPurchase", "common_core", payload)
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

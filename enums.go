package vinderman

type EFortCollectedState string

const (
	EFortCollectedStateUnknown        EFortCollectedState = "Unknown"
	EFortCollectedStateNew            EFortCollectedState = "New"
	EFortCollectedStateKnown          EFortCollectedState = "Known"
	EFortCollectedStateNewlyCollected EFortCollectedState = "NewlyCollected"
	EFortCollectedStateCollected      EFortCollectedState = "Collected"
	EFortCollectedStateNewBest        EFortCollectedState = "NewBest"
	EFortCollectedStateNewRecord      EFortCollectedState = "NewRecord"
	EFortCollectedStateNewLocation    EFortCollectedState = "NewLocation"
	EFortCollectedStateNewlyCompleted EFortCollectedState = "NewlyCompleted"
	EFortCollectedStateComplete       EFortCollectedState = "Complete"
)

type SquadAttribute string

const (
	SquadAttributeMedicineEMTSquad            SquadAttribute = "SquadAttribute_Medicine_EMTSquad"
	SquadAttributeArmsFireTeamAlpha           SquadAttribute = "SquadAttribute_Arms_FireTeamAlpha"
	SquadAttributeScavengingGadgeteers        SquadAttribute = "SquadAttribute_Scavenging_Gadgeteers"
	SquadAttributeSynthesisCorpsOfEngineering SquadAttribute = "SquadAttribute_Synthesis_CorpsofEngineering"
	SquadAttributeMedicineTrainingTeam        SquadAttribute = "SquadAttribute_Medicine_TrainingTeam"
	SquadAttributeArmsCloseAssaultSquad       SquadAttribute = "SquadAttribute_Arms_CloseAssaultSquad"
	SquadAttributeScavengingScoutingParty     SquadAttribute = "SquadAttribute_Scavenging_ScoutingParty"
	SquadAttributeSynthesisTheThinkTank       SquadAttribute = "SquadAttribute_Synthesis_TheThinkTank"
)

type ESocialImportPanelPlatform string

const (
	ESocialImportPanelPlatformFacebook    ESocialImportPanelPlatform = "Facebook"
	ESocialImportPanelPlatformVK          ESocialImportPanelPlatform = "VK"
	ESocialImportPanelPlatformSteam       ESocialImportPanelPlatform = "Steam"
	ESocialImportPanelPlatformXbox        ESocialImportPanelPlatform = "Xbox"
	ESocialImportPanelPlatformPlaystation ESocialImportPanelPlatform = "Playstation"
	ESocialImportPanelPlatformSwitch      ESocialImportPanelPlatform = "Switch"
)

type RealMoneyPurchaseRefreshType string

const (
	RealMoneyPurchaseRefreshTypeDefault           RealMoneyPurchaseRefreshType = "Default"
	RealMoneyPurchaseRefreshTypeForceAll          RealMoneyPurchaseRefreshType = "ForceAll"
	RealMoneyPurchaseRefreshTypeForceCurrent      RealMoneyPurchaseRefreshType = "ForceCurrent"
	RealMoneyPurchaseRefreshTypeUpdateOfflineAuth RealMoneyPurchaseRefreshType = "UpdateOfflineAuth"
)

type RealMoneyPurchaseVerifierModeOverride string

const (
	RealMoneyPurchaseVerifierModeOverrideDefaultToConfig              RealMoneyPurchaseVerifierModeOverride = "DefaultToConfig"
	RealMoneyPurchaseVerifierModeOverrideOccurrenceOnly               RealMoneyPurchaseVerifierModeOverride = "OccurrenceOnly"
	RealMoneyPurchaseVerifierModeOverrideOccurrenceOnlyRemoveReceipts RealMoneyPurchaseVerifierModeOverride = "OccurrenceOnlyRemoveReceipts"
	RealMoneyPurchaseVerifierModeOverrideOccurrencePrimary            RealMoneyPurchaseVerifierModeOverride = "OccurrencePrimary"
	RealMoneyPurchaseVerifierModeOverrideReceiptOnly                  RealMoneyPurchaseVerifierModeOverride = "ReceiptOnly"
	RealMoneyPurchaseVerifierModeOverrideReceiptPrimary               RealMoneyPurchaseVerifierModeOverride = "ReceiptPrimary"
)

type ExternalAuthType string

const (
	ExternalAuthTypeSteam    ExternalAuthType = "steam"
	ExternalAuthTypeGithub   ExternalAuthType = "github"
	ExternalAuthTypeTwitch   ExternalAuthType = "twitch"
	ExternalAuthTypeXBL      ExternalAuthType = "xbl"
	ExternalAuthTypePSN      ExternalAuthType = "psn"
	ExternalAuthTypeNintendo ExternalAuthType = "nintendo"
)

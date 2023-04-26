package vinderman

type AthenaProfileStats struct {
	Attributes struct {
		UseRandomLoadout bool `json:"use_random_loadout"`
		PastSeasons      []struct {
			SeasonNumber    int64 `json:"seasonNumber"`
			NumWins         int64 `json:"numWins"`
			NumHighBracket  int64 `json:"numHighBracket"`
			NumLowBracket   int64 `json:"numLowBracket"`
			SeasonXP        int64 `json:"seasonXp"`
			SeasonLevel     int64 `json:"seasonLevel"`
			BookXP          int64 `json:"bookXp"`
			BookLevel       int64 `json:"bookLevel"`
			PurchasedVIP    bool  `json:"purchasedVIP"`
			NumRoyalRoyales int64 `json:"numRoyalRoyales"`
		} `json:"past_seasons"`
		SeasonMatchBoost          int64    `json:"season_match_boost"`
		Loadouts                  []string `json:"loadouts"`
		RestedXPOverflow          int64    `json:"rested_xp_overflow"`
		MFARewardClaimed          bool     `json:"mfa_reward_claimed"`
		LastXPInteraction         string   `json:"last_xp_interaction"`
		RestedXPGoldenPathGranted int64    `json:"rested_xp_golden_path_granted"`
		QuestManager              struct {
			DailyLoginInterval string `json:"dailyLoginInterval"`
			DailyQuestRerolls  int64  `json:"dailyQuestRerolls"`
		} `json:"quest_manager"`
		BookLevel         int64 `json:"book_level"`
		SeasonNum         int64 `json:"season_num"`
		SeasonUpdate      int64 `json:"season_update"`
		CreativeDynamicXP struct {
			Timespan          float64 `json:"timespan"`
			BucketXP          int64   `json:"bucketXp"`
			BankXP            int64   `json:"bankXp"`
			BankXpMult        float64 `json:"bankXpMult"`
			BoosterBucketXP   int64   `json:"boosterBucketXp"`
			BoosterXpMult     float64 `json:"boosterXpMult"`
			DailyExcessXPMult float64 `json:"dailyExcessXpMult"`
			CurrentDayXP      int64   `json:"currentDayXp"`
			CurrentDay        int64   `json:"currentDay"`
		} `json:"creative_dynamic_xp"`
		Season struct {
			NumWins        int64 `json:"numWins"`
			NumHighBracket int64 `json:"numHighBracket"`
			NumLowBracket  int64 `json:"numLowBracket"`
		} `json:"season"`
		Battlestars int64 `json:"battlestars"`
		VoteData    struct {
			ElectionId  string `json:"electionId"`
			VoteHistory map[string]struct {
				VoteCount   int64  `json:"voteCount"`
				FirstVoteAt string `json:"firstVoteAt"`
				LastVoteAt  string `json:"lastVoteAt"`
			} `json:"voteHistory"`
			VotesRemaining  int64  `json:"votesRemaining"`
			LastVoteGranted string `json:"lastVoteGranted"`
		} `json:"vote_data"`
		BattlestarsSeasonTotal        int64  `json:"battlestars_season_total"`
		LifetimeWins                  int64  `json:"lifetime_wins"`
		PartyAssistQuest              string `json:"party_assist_quest"`
		BookPurchased                 bool   `json:"book_purchased"`
		PurchasedBattlePassTierOffers []struct {
			ID    string `json:"id"`
			Count int64  `json:"count"`
		} `json:"purchased_battle_pass_tier_offers"`
		RestedXPExchange       float64 `json:"rested_xp_exchange"`
		Level                  int64   `json:"level"`
		RestedXPMult           float64 `json:"rested_xp_mult"`
		AccountLevel           int64   `json:"accountLevel"`
		PinnedQuest            string  `json:"pinned_quest"`
		LastAppliedLoadout     string  `json:"last_applied_loadout"`
		XP                     int64   `json:"xp"`
		SeasonFriendMatchBoost int64   `json:"season_friend_match_boost"`
		PurchasedBPOffers      []struct {
			OfferID           string `json:"offerId"`
			BIsFreePassReward bool   `json:"bIsFreePassReward"`
			PurchaseDate      string `json:"purchaseDate"`
			LootResult        []struct {
				ItemType    string `json:"itemType"`
				ItemGUID    string `json:"itemGuid"`
				ItemProfile string `json:"itemProfile"`
				Attributes  struct {
					Platform string `json:"platform"`
				} `json:"attributes"`
				Quantity int64 `json:"quantity"`
			} `json:"lootResult"`
			CurrencyType      string `json:"currencyType"`
			TotalCurrencyPaid int64  `json:"totalCurrencyPaid"`
		} `json:"purchased_bp_offers"`
		LastMatchEndDatetime            string `json:"last_match_end_datetime"`
		LastSTWAccoladeTransferDatetime string `json:"last_stw_accolade_transfer_datetime"`
		MtxPurchaseHistoryCopy          []struct {
			PurchaseID   string   `json:"purchaseId"`
			PurchaseDate string   `json:"purchaseDate"`
			TemplateIDs  []string `json:"templateIds"`
		} `json:"mtx_purchase_history_copy"`
	} `json:"attributes"`
}

type CampaignProfileStats struct {
	Attributes struct {
		NodeCosts struct {
			HomebaseNodeDefaultPage struct {
				TokenHomebasepoints int `json:"Token:homebasepoints"`
			} `json:"homebase_node_default_page"`
			ResearchNodeDefaultPage map[string]int `json:"research_node_default_page"`
		} `json:"node_costs"`
		MissionAlertRedemptionRecord struct {
			ClaimData []struct {
				MissionAlertID         string `json:"missionAlertId"`
				RedemptionDateUTC      string `json:"redemptionDateUtc"`
				EvictClaimDataAfterUTC string `json:"evictClaimDataAfterUtc"`
			} `json:"claimData"`
		} `json:"mission_alert_redemption_record"`
		ClientSettings struct {
			PinnedQuestInstances []interface{} `json:"pinnedQuestInstances"`
		} `json:"client_settings"`
		ResearchLevels struct {
			Fortitude  int `json:"fortitude"`
			Offense    int `json:"offense"`
			Resistance int `json:"resistance"`
			Technology int `json:"technology"`
		} `json:"research_levels"`
		Level               int           `json:"level"`
		SelectedHeroLoadout string        `json:"selected_hero_loadout"`
		Loadouts            []interface{} `json:"loadouts"`
		CollectionBook      struct {
			MaxBookXPLevelAchieved int `json:"maxBookXpLevelAchieved"`
		} `json:"collection_book"`
		LatentXPMarker   string `json:"latent_xp_marker"`
		MFARewardClaimed bool   `json:"mfa_reward_claimed"`
		QuestManager     struct {
			DailyLoginInterval string `json:"dailyLoginInterval"`
			DailyQuestRerolls  int    `json:"dailyQuestRerolls"`
			QuestPoolStats     struct {
				PoolStats []struct {
					PoolName         string   `json:"poolName"`
					NextRefresh      string   `json:"nextRefresh"`
					RerollsRemaining int      `json:"rerollsRemaining"`
					QuestHistory     []string `json:"questHistory"`
				} `json:"poolStats"`
				DailyLoginInterval string `json:"dailyLoginInterval"`
				PoolLockouts       struct {
					PoolLockouts []struct {
						LockoutName string `json:"lockoutName"`
					} `json:"poolLockouts"`
				} `json:"poolLockouts"`
			} `json:"questPoolStats"`
		} `json:"quest_manager"`
		LegacyResearchPointsSpent int `json:"legacy_research_points_spent"`
		GameplayStats             []struct {
			StatName  string `json:"statName"`
			StatValue int    `json:"statValue"`
		} `json:"gameplay_stats"`
		EventCurrency struct {
			TemplateID string  `json:"templateId"`
			CF         float64 `json:"cf"`
		} `json:"event_currency"`
		MatchesPlayed int           `json:"matches_played"`
		ModeLoadouts  []interface{} `json:"mode_loadouts"`
		DailyRewards  struct {
			NextDefaultReward   int    `json:"nextDefaultReward"`
			TotalDaysLoggedIn   int    `json:"totalDaysLoggedIn"`
			LastClaimDate       string `json:"lastClaimDate"`
			AdditionalSchedules struct {
				Founderspackdailyrewardtoken struct {
					RewardsClaimed int  `json:"rewardsClaimed"`
					ClaimedToday   bool `json:"claimedToday"`
				} `json:"founderspackdailyrewardtoken"`
			} `json:"additionalSchedules"`
		} `json:"daily_rewards"`
		LastAppliedLoadout string `json:"last_applied_loadout"`
		XP                 int    `json:"xp"`
		PacksGranted       int    `json:"packs_granted"`
	} `json:"attributes"`
}

type CollectionBookPeopleProfileStats struct {
	Attributes struct{} `json:"attributes"`
}

type CollectionBookSchematicsProfileStats struct {
	Attributes struct{} `json:"attributes"`
}

type CollectionsProfileStats struct {
	Attributes struct {
		CurrentSeason int `json:"current_season"`
	} `json:"attributes"`
}

type CommonCoreProfileStats struct {
	Attributes struct {
		SurveyData struct {
			AllSurveysMetadata struct {
				NumTimesCompleted int    `json:"numTimesCompleted"`
				LastTimeCompleted string `json:"lastTimeCompleted"`
			} `json:"allSurveysMetadata"`
			Metadata map[string]struct {
				NumTimesCompleted int    `json:"numTimesCompleted"`
				LastTimeCompleted string `json:"lastTimeCompleted"`
			} `json:"metadata"`
		} `json:"survey_data"`
		IntroGamePlayed    bool `json:"intro_game_played"`
		MtxPurchaseHistory struct {
			RefundsUsed               int    `json:"refundsUsed"`
			RefundCredits             int    `json:"refundCredits"`
			TokenRefreshReferenceTime string `json:"tokenRefreshReferenceTime"`
			Purchases                 []struct {
				PurchaseID         string        `json:"purchaseId"`
				OfferID            string        `json:"offerId"`
				PurchaseDate       string        `json:"purchaseDate"`
				FreeRefundEligible bool          `json:"freeRefundEligible"`
				Fulfillments       []interface{} `json:"fulfillments"`
				LootResult         []struct {
					ItemType    string `json:"itemType"`
					ItemGUID    string `json:"itemGuid"`
					ItemProfile string `json:"itemProfile"`
					Quantity    int    `json:"quantity"`
				} `json:"lootResult"`
				TotalMtxPaid int `json:"totalMtxPaid"`
				Metadata     struct {
				} `json:"metadata"`
				GameContext string `json:"gameContext"`
				RefundDate  string `json:"refundDate,omitempty"`
				UndoTimeout string `json:"undoTimeout,omitempty"`
			} `json:"purchases"`
		} `json:"mtx_purchase_history"`
		UndoCooldowns []struct {
			OfferID         string `json:"offerId"`
			CooldownExpires string `json:"cooldownExpires"`
		} `json:"undo_cooldowns"`
		MtxAffiliateSetTime string `json:"mtx_affiliate_set_time"`
		CurrentMtxPlatform  string `json:"current_mtx_platform"`
		MtxAffiliate        string `json:"mtx_affiliate"`
		WeeklyPurchases     struct {
			LastInterval string         `json:"lastInterval"`
			PurchaseList map[string]int `json:"purchaseList"`
		} `json:"weekly_purchases"`
		DailyPurchases struct {
			LastInterval string         `json:"lastInterval"`
			PurchaseList map[string]int `json:"purchaseList"`
		} `json:"daily_purchases"`
		InAppPurchases struct {
			Receipts          []string       `json:"receipts"`
			IgnoredReceipts   []interface{}  `json:"ignoredReceipts"`
			FulfillmentCounts map[string]int `json:"fulfillmentCounts"`
			RefreshTimers     map[string]struct {
				NextEntitlementRefresh string `json:"nextEntitlementRefresh"`
			} `json:"refreshTimers"`
		} `json:"in_app_purchases"`
		ForcedIntroPlayed  string `json:"forced_intro_played"`
		RmtPurchaseHistory struct {
			Purchases []struct {
				FulfillmentID string `json:"fulfillmentId"`
				PurchaseDate  string `json:"purchaseDate"`
				LootResult    []struct {
					ItemType    string `json:"itemType"`
					ItemGUID    string `json:"itemGuid"`
					ItemProfile string `json:"itemProfile"`
					Attributes  struct {
						Platform string `json:"platform"`
					} `json:"attributes"`
					Quantity int `json:"quantity"`
				} `json:"lootResult"`
			} `json:"purchases"`
		} `json:"rmt_purchase_history"`
		UndoTimeout      string `json:"undo_timeout"`
		MonthlyPurchases struct {
			LastInterval string         `json:"lastInterval"`
			PurchaseList map[string]int `json:"purchaseList"`
		} `json:"monthly_purchases"`
		AllowedToSendGifts    bool   `json:"allowed_to_send_gifts"`
		MFAEnabled            bool   `json:"mfa_enabled"`
		AllowedToReceiveGifts bool   `json:"allowed_to_receive_gifts"`
		MtxAffiliateID        string `json:"mtx_affiliate_id"`
		GiftHistory           struct {
			NumSent      int               `json:"num_sent"`
			SentTo       map[string]string `json:"sentTo"`
			NumReceived  int               `json:"num_received"`
			ReceivedFrom map[string]string `json:"receivedFrom"`
			Gifts        []struct {
				Date        string `json:"date"`
				OfferID     string `json:"offerId"`
				ToAccountID string `json:"toAccountId"`
			} `json:"gifts"`
		} `json:"gift_history"`
	} `json:"attributes"`
}

type CommonPublicProfileStats struct {
	Attributes struct {
		BannerColor  string `json:"banner_color"`
		HomebaseName string `json:"homebase_name"`
		BannerIcon   string `json:"banner_icon"`
	} `json:"attributes"`
}

type CreativeProfileStats struct {
	Attributes struct {
		LastUsedProject string `json:"last_used_project"`
		MaxIslandPlots  int    `json:"max_island_plots"`
		LastUsedPlot    string `json:"last_used_plot"`
	} `json:"attributes"`
}

type MetadataProfileStats struct {
	Attributes struct{} `json:"attributes"`
}

type OutpostProfileStats struct {
	Attributes struct{} `json:"attributes"`
}

type RecycleBinProfileStats struct {
	Attributes struct {
		NextReceiptSequence int `json:"next_receipt_sequence"`
	} `json:"attributes"`
}

type Theater0ProfileStats struct {
	Attributes struct {
		PlayerLoadout struct {
			BPlayerIsNew          bool `json:"bPlayerIsNew"`
			PrimaryQuickBarRecord struct {
				Slots []struct {
					Items []string `json:"items"`
				} `json:"slots"`
			} `json:"primaryQuickBarRecord"`
			SecondaryQuickBarRecord struct {
				Slots []struct {
					Items []string `json:"items"`
				} `json:"slots"`
			} `json:"secondaryQuickBarRecord"`
			ZonesCompleted int `json:"zonesCompleted"`
		} `json:"player_loadout"`
	} `json:"attributes"`
}

type Theater1ProfileStats struct {
	Attributes struct {
		PlayerLoadout struct {
			BPlayerIsNew             bool          `json:"bPlayerIsNew"`
			PinnedSchematicInstances []interface{} `json:"pinnedSchematicInstances"`
			PrimaryQuickBarRecord    struct {
				CurrentFocusedSlot   int `json:"currentFocusedSlot"`
				PreviousFocusedSlot  int `json:"previousFocusedSlot"`
				SecondaryFocusedSlot int `json:"secondaryFocusedSlot"`
				Slots                []struct {
					Items []string `json:"items"`
				} `json:"slots"`
				DataDefinition struct {
					QuickbarSlots []struct {
						AcceptedItemTypes []string `json:"acceptedItemTypes"`
						BStaticSlot       bool     `json:"bStaticSlot"`
						DefaultItem       string   `json:"defaultItem"`
					} `json:"quickbarSlots"`
				} `json:"dataDefinition"`
			} `json:"primaryQuickBarRecord"`
			SecondaryQuickBarRecord struct {
				CurrentFocusedSlot   int `json:"currentFocusedSlot"`
				PreviousFocusedSlot  int `json:"previousFocusedSlot"`
				SecondaryFocusedSlot int `json:"secondaryFocusedSlot"`
				Slots                []struct {
					Items []string `json:"items"`
				} `json:"slots"`
				DataDefinition struct {
					QuickbarSlots []struct {
						AcceptedItemTypes []string `json:"acceptedItemTypes"`
						BStaticSlot       bool     `json:"bStaticSlot"`
						DefaultItem       string   `json:"defaultItem"`
					} `json:"quickbarSlots"`
				} `json:"dataDefinition"`
			} `json:"secondaryQuickBarRecord"`
			ZonesCompleted int `json:"zonesCompleted"`
		} `json:"player_loadout"`
	} `json:"attributes"`
}

type Theater2ProfileStats struct {
	Attributes struct {
		PlayerLoadout struct {
			BPlayerIsNew          bool `json:"bPlayerIsNew"`
			PrimaryQuickBarRecord struct {
				Slots []struct {
					Items []string `json:"items,omitempty"`
				} `json:"slots"`
			} `json:"primaryQuickBarRecord"`
			ZonesCompleted int `json:"zonesCompleted"`
		} `json:"player_loadout"`
		LastEventInstanceKey string `json:"last_event_instance_key"`
	} `json:"attributes"`
}

type CampaignNotifications []struct {
	Type         string `json:"type"`
	Primary      bool   `json:"primary"`
	DaysLoggedIn int    `json:"daysLoggedIn"`
	Items        []struct {
		ItemType    string `json:"itemType"`
		ItemGuid    string `json:"itemGuid"`
		ItemProfile string `json:"itemProfile"`
		Quantity    int    `json:"quantity"`
	} `json:"items"`
}

type AthenaCosmeticItem struct {
	TemplateID string `json:"templateId"`
	Attributes struct {
		Favorite      bool `json:"favorite,omitempty"`
		Archived      bool `json:"archived,omitempty"`
		ItemSeen      bool `json:"item_seen"`
		Level         int  `json:"level"`
		MaxLevelBonus int  `json:"max_level_bonus"`
		RndSelCnt     int  `json:"rnd_sel_cnt"`
		Variants      []struct {
			Channel string   `json:"channel"`
			Active  string   `json:"active"`
			Owned   []string `json:"owned"`
		} `json:"variants,omitempty"`
		XP int `json:"xp"`
	} `json:"attributes"`
	Quantity int `json:"quantity"`
}

type CommonCoreMtxItem struct {
	TemplateID string `json:"templateId"`
	Quantity   int    `json:"quantity"`
	Attributes struct {
		Platform string `json:"platform"`
	} `json:"attributes"`
}

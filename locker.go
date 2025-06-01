package vinderman

import (
	"encoding/json"
	"fmt"
	"github.com/0xDistrust/Vinderman/consts"
	"github.com/0xDistrust/Vinderman/request"
	"net/http"
	"time"
)

type Loadout struct {
	LoadoutSlots []struct {
		SlotTemplate       string `json:"slotTemplate"`
		EquippedItemId     string `json:"equippedItemId"`
		ItemCustomizations []struct {
			ChannelTag     string `json:"channelTag"`
			VariantTag     string `json:"variantTag"`
			AdditionalData string `json:"additionalData"`
		} `json:"itemCustomizations"`
	} `json:"loadoutSlots"`
	ShuffleType string `json:"shuffleType"`
}

type ItemCustomization struct {
	ChannelTag     string `json:"channelTag"`
	VariantTag     string `json:"variantTag"`
	AdditionalData string `json:"additionalData,omitempty"`
}

type LoadoutSlot struct {
	SlotTemplate       string              `json:"slotTemplate"`
	EquippedItemId     string              `json:"equippedItemId"`
	ItemCustomizations []ItemCustomization `json:"itemCustomizations"`
}

type LoadoutPreset struct {
	DeploymentId string        `json:"deploymentId"`
	AccountId    string        `json:"accountId"`
	LoadoutType  string        `json:"loadoutType"`
	PresetId     string        `json:"presetId"`
	PresetIndex  int           `json:"presetIndex"`
	AthenaItemId string        `json:"athenaItemId"`
	CreationTime time.Time     `json:"creationTime"`
	UpdatedTime  time.Time     `json:"updatedTime"`
	LoadoutSlots []LoadoutSlot `json:"loadoutSlots"`
}

type ActiveLoadoutGroup struct {
	DeploymentId string               `json:"deploymentId"`
	AccountId    string               `json:"accountId"`
	AthenaItemId string               `json:"athenaItemId"`
	CreationTime time.Time            `json:"creationTime"`
	UpdatedTime  time.Time            `json:"updatedTime"`
	Loadouts     []map[string]Loadout `json:"loadouts"`
	ShuffleType  string               `json:"shuffleType"`
}

type LockerItems struct {
	ActiveLoadoutGroup  ActiveLoadoutGroup `json:"activeLoadoutGroup"`
	LoadoutGroupPresets []any              `json:"loadoutGroupPresets"`
	LoadoutPresets      []LoadoutPreset    `json:"loadoutPresets"`
}

type UpdateActiveLockerLoadoutRequest struct {
	Loadouts             []map[string]Loadout `json:"loadouts"`
	ShuffleType          string               `json:"shuffleType"`
	EquippedPresetItemId string               `json:"equippedPresetItemId"`
	AthenaItemId         string               `json:"athenaItemId"`
	CreationTime         time.Time            `json:"creationTime"`
}

func (c Client) QueryLockerItems(credentials UserCredentials, deploymentId string) (items LockerItems, err error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	resp, err := c.Request("GET", fmt.Sprintf("%s/api/locker/v4/%s/account/%s/items", consts.FORTNITE_SERVICE, deploymentId, credentials.AccountID), headers, "")
	if err != nil {
		return LockerItems{}, err
	}

	res, err := request.ResponseParser[LockerItems](resp)

	return res.Body, err
	// https://fngw-svc-gc-livefn.ol.epicgames.com/api/locker/v4/:deploymentId/account/:accountId/items
}

func (c Client) UpdateActiveLockerLoadout(credentials UserCredentials, deploymentId string, payload UpdateActiveLockerLoadoutRequest) (activeLoadoutGroup ActiveLoadoutGroup, err error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return ActiveLoadoutGroup{}, err
	}

	resp, err := c.Request("PUT", fmt.Sprintf("%s/api/locker/v4/%s/account/%s/active-loadout-group", consts.FORTNITE_SERVICE, deploymentId, credentials.AccountID), headers, string(bodyBytes))
	if err != nil {
		return ActiveLoadoutGroup{}, err
	}

	res, err := request.ResponseParser[ActiveLoadoutGroup](resp)

	return res.Body, err
	// https://fngw-svc-gc-livefn.ol.epicgames.com/api/locker/v4/:deploymentId/account/:accountId/active-loadout-group
}

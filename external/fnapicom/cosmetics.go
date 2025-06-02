package fnapicom

import (
	"net/url"
	"time"

	"gitlab.com/8h9x/Vinderman/consts"
	"gitlab.com/8h9x/Vinderman/request"
)

type Cosmetic struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        struct {
		Value        string `json:"value"`
		DisplayValue string `json:"displayValue"`
		BackendValue string `json:"backendValue"`
	} `json:"type"`
	Rarity struct {
		Value        string `json:"value"`
		DisplayValue string `json:"displayValue"`
		BackendValue string `json:"backendValue"`
	} `json:"rarity"`
	Series       interface{} `json:"series"`
	Set          interface{} `json:"set"`
	Introduction struct {
		Chapter      string `json:"chapter"`
		Season       string `json:"season"`
		Text         string `json:"text"`
		BackendValue int    `json:"backendValue"`
	} `json:"introduction"`
	Images struct {
		SmallIcon string      `json:"smallIcon"`
		Icon      string      `json:"icon"`
		Featured  interface{} `json:"featured"`
		Other     interface{} `json:"other"`
	} `json:"images"`
	Variants []struct {
		Channel string `json:"channel"`
		Type    string `json:"type"`
		Options []struct {
			Tag                string `json:"tag"`
			Name               string `json:"name"`
			Image              string `json:"image"`
			UnlockRequirements string `json:"unlockRequirements,omitempty"`
		} `json:"options"`
	} `json:"variants"`
	SearchTags       interface{} `json:"searchTags"`
	GameplayTags     []string    `json:"gameplayTags"`
	MetaTags         interface{} `json:"metaTags"`
	ShowcaseVideo    interface{} `json:"showcaseVideo"`
	DynamicPakID     interface{} `json:"dynamicPakId"`
	DisplayAssetPath interface{} `json:"displayAssetPath"`
	DefinitionPath   interface{} `json:"definitionPath"`
	Path             string      `json:"path"`
	Added            time.Time   `json:"added"`
	ShopHistory      interface{} `json:"shopHistory"`
}

type CosmeticSearchRes struct {
	Status int      `json:"status"`
	Data   Cosmetic `json:"data"`
}

func (c Client) CosmeticSearch(params url.Values) (Cosmetic, error) {
	res, err := request.Getf[CosmeticSearchRes]("%s/cosmetics/br/search?%s", consts.FNAPICOM_API, params.Encode())
	return res.Body.Data, err
}

package vinderman

import (
	"fmt"
	"net/http"
	"time"
	
	"github.com/0xDistrust/Vinderman/consts"
	"github.com/0xDistrust/Vinderman/request"
)

type Catalog struct {
	RefreshIntervalHrs int       `json:"refreshIntervalHrs"`
	DailyPurchaseHrs   int       `json:"dailyPurchaseHrs"`
	Expiration         time.Time `json:"expiration"`
	Storefronts        []struct {
		Name           string `json:"name"`
		CatalogEntries []struct {
			OfferID   string `json:"offerId"`
			DevName   string `json:"devName"`
			OfferType string `json:"offerType"`
			Prices    []struct {
				CurrencyType        string    `json:"currencyType"`
				CurrencySubType     string    `json:"currencySubType"`
				RegularPrice        int       `json:"regularPrice"`
				DynamicRegularPrice int       `json:"dynamicRegularPrice"`
				FinalPrice          int       `json:"finalPrice"`
				SaleExpiration      time.Time `json:"saleExpiration"`
				BasePrice           int       `json:"basePrice"`
			} `json:"prices"`
			Categories   []interface{} `json:"categories"`
			DailyLimit   int           `json:"dailyLimit"`
			WeeklyLimit  int           `json:"weeklyLimit"`
			MonthlyLimit int           `json:"monthlyLimit"`
			Refundable   bool          `json:"refundable"`
			AppStoreID   []string      `json:"appStoreId"`
			Requirements []interface{} `json:"requirements"`
			MetaInfo     []struct {
				Key   string `json:"key,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"metaInfo"`
			CatalogGroup         string        `json:"catalogGroup"`
			CatalogGroupPriority int           `json:"catalogGroupPriority"`
			SortPriority         int           `json:"sortPriority"`
			Title                string        `json:"title"`
			ShortDescription     string        `json:"shortDescription"`
			Description          string        `json:"description"`
			DisplayAssetPath     string        `json:"displayAssetPath"`
			ItemGrants           []interface{} `json:"itemGrants"`
		} `json:"catalogEntries"`
	} `json:"storefronts"`
}

func (c Client) FetchCatalog(credentials UserCredentials) (catalog Catalog, err error) {
	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprint("Bearer ", credentials.AccessToken))

	resp, err := c.Request("POST", consts.FORTNITE_SERVICE+"/storefront/v2/catalog", headers, "")
	if err != nil {
		return
	}

	res, err := request.ResponseParser[Catalog](resp)

	return res.Body, err
}
package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/0xDistrust/Vinderman/consts"
	"github.com/0xDistrust/Vinderman/request"
)

type Offers map[string]struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	LongDescription string `json:"longDescription"`
	KeyImages       []struct {
		Type         string    `json:"type"`
		URL          string    `json:"url"`
		MD5          string    `json:"md5"`
		Width        int       `json:"width"`
		Height       int       `json:"height"`
		Size         int       `json:"size"`
		UploadedDate time.Time `json:"uploadedDate"`
	} `json:"keyImages"`
	Categories []struct {
		Path string `json:"path"`
	} `json:"categories"`
	Namespace        string    `json:"namespace"`
	Status           string    `json:"status"`
	CreationDate     time.Time `json:"creationDate"`
	LastModifiedDate time.Time `json:"lastModifiedDate"`
	CustomAttributes struct {
	} `json:"customAttributes"`
	InternalName string `json:"internalName"`
	Recurrence   string `json:"recurrence"`
	Items        []struct {
		ID              string        `json:"id"`
		Title           string        `json:"title"`
		Description     string        `json:"description"`
		LongDescription string        `json:"longDescription"`
		KeyImages       []interface{} `json:"keyImages"`
		Categories      []struct {
			Path string `json:"path"`
		} `json:"categories"`
		Namespace           string        `json:"namespace"`
		Status              string        `json:"status"`
		CreationDate        time.Time     `json:"creationDate"`
		LastModifiedDate    time.Time     `json:"lastModifiedDate"`
		EntitlementName     string        `json:"entitlementName"`
		EntitlementType     string        `json:"entitlementType"`
		ItemType            string        `json:"itemType"`
		ReleaseInfo         []interface{} `json:"releaseInfo"`
		Developer           string        `json:"developer"`
		DeveloperID         string        `json:"developerId"`
		UseCount            int           `json:"useCount"`
		EulaIds             []interface{} `json:"eulaIds"`
		EndOfSupport        bool          `json:"endOfSupport"`
		NsMajorItems        []interface{} `json:"nsMajorItems"`
		NsDependsOnDlcItems []interface{} `json:"nsDependsOnDlcItems"`
		Unsearchable        bool          `json:"unsearchable"`
	} `json:"items"`
	CurrencyCode          string `json:"currencyCode"`
	CurrentPrice          int    `json:"currentPrice"`
	Price                 int    `json:"price"`
	BasePrice             int    `json:"basePrice"`
	BasePriceCurrencyCode string `json:"basePriceCurrencyCode"`
	RecurringPrice        int    `json:"recurringPrice"`
	FreeDays              int    `json:"freeDays"`
	MaxBillingCycles      int    `json:"maxBillingCycles"`
	Seller                struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"seller"`
	ViewableDate                       time.Time     `json:"viewableDate"`
	EffectiveDate                      time.Time     `json:"effectiveDate"`
	VatIncluded                        bool          `json:"vatIncluded"`
	IsCodeRedemptionOnly               bool          `json:"isCodeRedemptionOnly"`
	IsFeatured                         bool          `json:"isFeatured"`
	TaxSkuID                           string        `json:"taxSkuId"`
	MerchantGroup                      string        `json:"merchantGroup"`
	PriceTier                          string        `json:"priceTier"`
	URLSlug                            string        `json:"urlSlug"`
	RoleNamesToGrant                   []interface{} `json:"roleNamesToGrant"`
	Tags                               []interface{} `json:"tags"`
	PurchaseLimit                      int           `json:"purchaseLimit"`
	IgnoreOrder                        bool          `json:"ignoreOrder"`
	FulfillToGroup                     bool          `json:"fulfillToGroup"`
	FraudItemType                      string        `json:"fraudItemType"`
	ShareRevenue                       bool          `json:"shareRevenue"`
	OfferType                          string        `json:"offerType"`
	Unsearchable                       bool          `json:"unsearchable"`
	ReleaseOffer                       string        `json:"releaseOffer"`
	Title4Sort                         string        `json:"title4Sort"`
	SelfRefundable                     bool          `json:"selfRefundable"`
	RefundType                         string        `json:"refundType"`
	VisibilityType                     string        `json:"visibilityType"`
	CurrencyDecimals                   int           `json:"currencyDecimals"`
	AllowPurchaseForPartialOwned       bool          `json:"allowPurchaseForPartialOwned"`
	ShareRevenueWithUnderageAffiliates bool          `json:"shareRevenueWithUnderageAffiliates"`
	PlatformWhitelist                  []interface{} `json:"platformWhitelist"`
	PlatformBlacklist                  []interface{} `json:"platformBlacklist"`
	PartialItemPrerequisiteCheck       bool          `json:"partialItemPrerequisiteCheck"`
}

func (c Client) FetchOffers(credentials UserCredentials, offerIDs ...string) (offers Offers, err error) {
	values := url.Values{}
	for _, offerID := range offerIDs {
		values.Add("id", offerID)
	}

	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	resp, err := c.Request("GET", fmt.Sprintf("%s/catalog/api/shared/bulk/offers?%s", consts.CATALOG_SERVICE, values.Encode()), headers, "")
	if err != nil {
		return
	}

	res, err := request.ResponseParser[Offers](resp)

	return res.Body, err
}
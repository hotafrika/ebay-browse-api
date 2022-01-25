package browse

type GetItemsByItemGroupResponse struct {
	CommonDescriptions []CommonDescription `json:"commonDescriptions"`
	Items              []Item              `json:"items"`
	// Warnings        []ErrorDetailV3 `json:"warnings"` // TODO
}

type CommonDescription struct {
	Description string   `json:"description"`
	ItemIDs     []string `json:"itemIds"`
}

type Item struct {
	ItemID           string `json:"itemId"`
	Title            string `json:"title"`
	ShortDescription string `json:"shortDescription"`
	Price            Price  `json:"price"`
	CategoryPath     string `json:"categoryPath"`
	Condition        string `json:"condition"`
	ConditionID      string `json:"conditionId"`
	//itemLocation
	Image                   Image
	AdditionalImages        []Image                 `json:"additionalImages"`
	Brand                   string                  `json:"brand"`
	Seller                  Seller                  `json:"seller"`
	GTIN                    string                  `json:"gtin"`
	MPN                     string                  `json:"mpn"`
	EstimatedAvailabilities []EstimatedAvailability `json:"estimatedAvailabilities"`
	ShippingOptions         []ShippingOption        `json:"shippingOptions"`
	// shipToLocations
	// returnTerms
	// taxes
	LocalizedAspects         []LocalizedAspect `json:"localizedAspects"`
	TopRatedBuyingExperience bool              `json:"topRatedBuyingExperience"`
	BuyingOptions            []string          `json:"buyingOptions"`
	ItemWebURL               string            `json:"itemWebUrl"`
	//paymentMethods
	PrimaryItemGroup PrimaryItemGroup `json:"primaryItemGroup"`
	//enabledForGuestCheckout
	//eligibleForInlineCheckout
	LegacyItemID string `json:"legacyItemId"`
	//priorityListing
	//adultOnly
	CategoryID string `json:"categoryId"`
}

type PrimaryItemGroup struct {
	ItemGroupID               string  `json:"itemGroupId"`
	ItemGroupType             string  `json:"itemGroupType"`
	ItemGroupHref             string  `json:"itemGroupHref"`
	ItemGroupTitle            string  `json:"itemGroupTitle"`
	ItemGroupImage            Image   `json:"itemGroupImage"`
	ItemGroupAdditionalImages []Image `json:"itemGroupAdditionalImages"`
}

type LocalizedAspect struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type EstimatedAvailability struct {
	EstimatedAvailabilityStatus string   `json:"estimatedAvailabilityStatus"`
	EstimatedAvailableQuantity  int      `json:"estimatedAvailableQuantity"`
	EstimatedSoldQuantity       int      `json:"estimatedSoldQuantity"`
	DeliveryOptions             []string `json:"deliveryOptions"`
}

package browse

type ItemSummarySearchResponse struct {
	AutoCorrections AutoCorrection `json:"autoCorrections"`
	Href            string         `json:"href"`
	Next            string         `json:"next"`
	Prev            string         `json:"prev"`
	ItemSummaries   []ItemSummary  `json:"itemSummaries"`
	Limit           int            `json:"limit"`
	Offset          int            `json:"offset"`
	Total           int            `json:"total"`
	// Warnings        []ErrorDetailV3 `json:"warnings"` // TODO
	// Refinement Refinement `json:"refinement"` // TODO
}

type AutoCorrection struct {
	Q string `json:"q"`
}

type ItemSummary struct {
	ItemID                   string           `json:"itemId"`
	Title                    string           `json:"title"`
	ShortDescription         string           `json:"shortDescription"`
	Image                    Image            `json:"image"`
	BidCount                 int              `json:"bidCount"`
	WatchCount               int              `json:"watchCount"`
	ItemEndDate              string           `json:"itemEndDate"`
	ItemAffiliateWebURL      string           `json:"itemAffiliateWebUrl"`
	Price                    Price            `json:"price"`
	ItemHref                 string           `json:"itemHref"`
	Seller                   Seller           `json:"seller"`
	MarketingPrice           MarketingPrice   `json:"marketingPrice"`
	Condition                string           `json:"condition"`
	ConditionID              string           `json:"conditionId"`
	ThumbnailImages          []Image          `json:"thumbnailImages"`
	ShippingOptions          []ShippingOption `json:"shippingOptions"`
	BuyingOptions            []string         `json:"buyingOptions"`
	CurrentBidPrice          Price            `json:"currentBidPrice"`
	Epid                     string           `json:"epid"`
	ItemWebURL               string           `json:"itemWebUrl"`
	ItemLocation             ItemLocation     `json:"itemLocation"`
	Categories               []Category       `json:"categories"`
	AdditionalImages         []Image          `json:"additionalImages"`
	LegacyItemID             string           `json:"legacyItemId"`
	AdultOnly                bool             `json:"adultOnly"`
	AvailableCoupons         bool             `json:"availableCoupons"`
	PriorityListing          bool             `json:"priorityListing"`
	TopRatedBuyingExperience bool             `json:"topRatedBuyingExperience"`
	// compatibilityMatch
	// compatibilityProperties
	// distanceFromPickupLocation
	// energyEfficiencyClass
	// itemGroupHref
	// itemGroupType
	// pickupOptions
	// priceDisplayCondition
	// qualifiedPrograms
	// tyreLabelImageUrl
	// unitPrice
	// unitPricingMeasure
}

type Image struct {
	ImageURL string `json:"imageUrl"`
}

type Price struct {
	Value                 string `json:"value"`
	Currency              string `json:"currency"`
	ConvertedFromValue    string `json:"convertedFromValue"`
	ConvertedFromCurrency string `json:"convertedFromCurrency"`
}

type Seller struct {
	Username           string `json:"username"`
	FeedbackPercentage string `json:"feedbackPercentage"`
	FeedbackScore      int    `json:"feedbackScore"`
	SellerAccountType  string `json:"sellerAccountType"`
}

type MarketingPrice struct {
	OriginalPrice      Price  `json:"originalPrice"`
	DiscountPercentage string `json:"discountPercentage"`
	DiscountAmount     Price  `json:"discountAmount"`
	PriceTreatment     string `json:"priceTreatment"`
}

type ShippingOption struct {
	GuaranteedDelivery       bool   `json:"guaranteedDelivery"`
	MaxEstimatedDeliveryDate string `json:"maxEstimatedDeliveryDate"`
	MinEstimatedDeliveryDate string `json:"minEstimatedDeliveryDate"`
	ShippingCostType         string `json:"shippingCostType"`
	ShippingCost             Price  `json:"shippingCost"`
}

type ItemLocation struct {
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
}

type Category struct {
	CategoryID string `json:"categoryId"`
}

type Refinement struct {
}

type ErrorDetailV3 struct {
}

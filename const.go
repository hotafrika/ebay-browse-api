package browse

const URLProductionItemSummarySearch = "https://api.ebay.com/buy/browse/v1/item_summary/search"
const URLSandboxItemSummarySearch = "https://api.sandbox.ebay.com/buy/browse/v1/item_summary/search"

const URLProductionGetItemByLegacyId = "https://api.ebay.com/buy/browse/v1/item/get_item_by_legacy_id"
const URLSandboxGetItemByLegacyId = "https://api.sandbox.ebay.com/buy/browse/v1/item/get_item_by_legacy_id"

const URLProductionGetItemsByItemGroup = "https://api.ebay.com/buy/browse/v1/item/get_items_by_item_group"
const URLSandboxGetItemsByItemGroup = "https://api.sandbox.ebay.com/buy/browse/v1/item/get_items_by_item_group"

type EbayMarketplace string

// https://developer.ebay.com/api-docs/buy/static/ref-marketplace-supported.html
const (
	MarketplaceUSA EbayMarketplace = "EBAY_US"
	MarketplaceGB  EbayMarketplace = "EBAY_GB"
	// TODO add
)

type FilterType string

// https://developer.ebay.com/api-docs/buy/static/ref-buy-browse-filters.html
const (
	FilterDeliveryCountry    FilterType = "deliveryCountry"
	FilterDeliveryPostalCode FilterType = "deliveryPostalCode"

	// TODO add
)

type FieldGroupsOptions string

const (
	FieldGroupsPRODUCT FieldGroupsOptions = "PRODUCT"
)

package main

type Listing1 struct {
	ItemName                string `tsv:"item-name"`
	ItemDescription         string `tsv:"item-description"`
	ListingId               string `tsv:"listing-id"`
	SellerSku               string `tsv:"seller-sku"`
	Price                   string `tsv:"price"`
	Quantity                string `tsv:"quantity"`
	OpenDate                string `tsv:"open-date"`
	ImageUrl                string `tsv:"image-url"`
	ItemIsMarketplace       string `tsv:"item-is-marketplace"`
	ProductIdType           string `tsv:"product-id-type"`
	ZshopShippingFee        string `tsv:"zshop-shipping-fee"`
	ItemNote                string `tsv:"item-note"`
	ItemCondition           string `tsv:"item-condition"`
	ZshopCategory1          string `tsv:"zshop-category1"`
	ZshopBrowsePath         string `tsv:"zshop-browse-path"`
	ZshopStorefrontFeature  string `tsv:"zshop-storefront-feature"`
	Asin1                   string `tsv:"asin1"`
	Asin2                   string `tsv:"asin2"`
	Asin3                   string `tsv:"asin3"`
	WillShipInternationally string `tsv:"will-ship-internationally"`
	ExpeditedShipping       string `tsv:"expedited-shipping"`
	ZshopBoldface           string `tsv:"zshop-boldface"`
	ProductId               string `tsv:"product-id"`
	BidForFeaturedPlacement string `tsv:"bid-for-featured-placement"`
	AddDelete               string `tsv:"add-delete"`
	PendingQuantity         string `tsv:"pending-quantity"`
	FulfillmentChannel      string `tsv:"fulfillment-channel"`
	MerchantShippingGroup   string `tsv:"merchant-shipping-group"`
	Status                  string `tsv:"status"`
}

package models

import (
   // "gorm.io/gorm"

	"go-admin/common/models"
)

type Product struct {
    models.Model
    
    DataId string `json:"dataId" gorm:"type:bigint;comment:data_id"` 
    Puid string `json:"puid" gorm:"type:varchar(100);comment:null"` 
    ShopId string `json:"shopId" gorm:"type:bigint;comment:4462"` 
    MarketplaceId string `json:"marketplaceId" gorm:"type:varchar(100);comment:ATVPDKIKX0DER"` 
    FullCid string `json:"fullCid" gorm:"type:varchar(100);comment:FullCid"` 
    ParentId string `json:"parentId" gorm:"type:bigint;comment:4762270"` 
    DxmState string `json:"dxmState" gorm:"type:varchar(100);comment:online_amazon"` 
    DxmPublishState string `json:"dxmPublishState" gorm:"type:varchar(100);comment:null"` 
    ErrMsg string `json:"errMsg" gorm:"type:varchar(100);comment:null"` 
    Asin string `json:"asin" gorm:"type:varchar(100);comment:B08KGSXPZS"` 
    ListingId string `json:"listingId" gorm:"type:varchar(100);comment:0324YYDEASD"` 
    ParentAsin string `json:"parentAsin" gorm:"type:varchar(100);comment:null"` 
    ChildAsins string `json:"childAsins" gorm:"type:varchar(100);comment:null"` 
    IsVariation string `json:"isVariation" gorm:"type:bigint;comment:0"` 
    Sku string `json:"sku" gorm:"type:varchar(100);comment:K7-PQU4-OE8V"` 
    StandardProductType string `json:"standardProductType" gorm:"type:varchar(100);comment:null"` 
    StandardProductId string `json:"standardProductId" gorm:"type:varchar(100);comment:null"` 
    Title string `json:"title" gorm:"type:varchar(100);comment:product title"` 
    Brand string `json:"brand" gorm:"type:varchar(100);comment:null"` 
    Manufacturer string `json:"manufacturer" gorm:"type:varchar(100);comment:null"` 
    Description string `json:"description" gorm:"type:varchar(100);comment:null"` 
    BulletPoint string `json:"bulletPoint" gorm:"type:varchar(100);comment:null"` 
    ConditionType string `json:"conditionType" gorm:"type:varchar(100);comment:null"` 
    ConditionValue string `json:"conditionValue" gorm:"type:varchar(100);comment:null"` 
    StandardPrice string `json:"standardPrice" gorm:"type:double;comment:14.99"` 
    ListingPricing string `json:"listingPricing" gorm:"type:double;comment:14.99"` 
    SwitchFulfillmentTo string `json:"switchFulfillmentTo" gorm:"type:varchar(100);comment:AFN"` 
    FulfillmentChannel string `json:"fulfillmentChannel" gorm:"type:varchar(100);comment:null"` 
    PartNumber string `json:"partNumber" gorm:"type:varchar(100);comment:null"` 
    MainImage string `json:"mainImage" gorm:"type:varchar(100);comment:https://m.media-amazon.com/images/I/613gR7poDQL._SL75_.jpg"` 
    SaleStartDate string `json:"saleStartDate" gorm:"type:varchar(100);comment:null"` 
    SaleEndDate string `json:"saleEndDate" gorm:"type:varchar(100);comment:null"` 
    SaleSalePrice string `json:"saleSalePrice" gorm:"type:varchar(100);comment:null"` 
    Quantity string `json:"quantity" gorm:"type:bigint;comment:0"` 
    OpenDate string `json:"openDate" gorm:"type:varchar(100);comment:2021-03-23 19:09:08"` 
    ItemDimensions string `json:"itemDimensions" gorm:"type:varchar(100);comment:null"` 
    PackageDimensions string `json:"packageDimensions" gorm:"type:varchar(100);comment:null"` 
    StandardPriceCurrency string `json:"standardPriceCurrency" gorm:"type:varchar(100);comment:USD"` 
    Specifics string `json:"specifics" gorm:"type:varchar(100);comment:null"` 
    VariationChildStr string `json:"variationChildStr" gorm:"type:varchar(100);comment:null"` 
    OnlineStatus string `json:"onlineStatus" gorm:"type:varchar(100);comment:Active"` 
    SaleNum string `json:"saleNum" gorm:"type:bigint;comment:54"` 
    Fnsku string `json:"fnsku" gorm:"type:varchar(100);comment:X002UH0IRH"` 
    Warehousing string `json:"warehousing" gorm:"type:bigint;comment:542"` 
    Unsellable string `json:"unsellable" gorm:"type:bigint;comment:0"` 
    ReservedQty string `json:"reservedQty" gorm:"type:bigint;comment:17"` 
    PurchaseCost string `json:"purchaseCost" gorm:"type:bigint;comment:0"` 
    HeadTripCost string `json:"headTripCost" gorm:"type:bigint;comment:0"` 
    ShipCost string `json:"shipCost" gorm:"type:varchar(100);comment:null"` 
    PriceFeedStatus string `json:"priceFeedStatus" gorm:"type:varchar(100);comment:null"` 
    InventoryFeedStatus string `json:"inventoryFeedStatus" gorm:"type:varchar(100);comment:null"` 
    UpdateStandardPrice string `json:"updateStandardPrice" gorm:"type:varchar(100);comment:null"` 
    UpdateQuantity string `json:"updateQuantity" gorm:"type:varchar(100);comment:null"` 
    Rating string `json:"rating" gorm:"type:varchar(100);comment:null"` 
    RatingCount string `json:"ratingCount" gorm:"type:varchar(100);comment:null"` 
    Rank string `json:"rank" gorm:"type:varchar(100);comment:null"` 
    BsrRank string `json:"bsrRank" gorm:"type:varchar(100);comment:fuck"` 
    CommodityId string `json:"commodityId" gorm:"type:varchar(100);comment:null"` 
    DevId string `json:"devId" gorm:"type:varchar(100);comment:DevId"` 
    CreateId string `json:"createId" gorm:"type:varchar(100);comment:null"` 
    UpdateId string `json:"updateId" gorm:"type:varchar(100);comment:null"` 
    LpriceUpdateTime string `json:"lpriceUpdateTime" gorm:"type:varchar(100);comment:null"` 
    LpriceStatus string `json:"lpriceStatus" gorm:"type:varchar(100);comment:null"` 
    ChildList string `json:"childList" gorm:"type:varchar(100);comment:null"` 
    SalePrices string `json:"salePrices" gorm:"type:double;comment:809.46"` 
    AdCosts string `json:"adCosts" gorm:"type:double;comment:130.33"` 
    Currency string `json:"currency" gorm:"type:varchar(100);comment:USD"` 
    CostCurrency string `json:"costCurrency" gorm:"type:varchar(100);comment:null"` 
    CommoditySku string `json:"commoditySku" gorm:"type:varchar(100);comment:null"` 
    CommodityName string `json:"commodityName" gorm:"type:varchar(100);comment:null"` 
    ShopName string `json:"shopName" gorm:"type:varchar(100);comment:A3LR0SXN8QVIE0-na-US"` 
    SiteName string `json:"siteName" gorm:"type:varchar(100);comment:美国"` 
    Domain string `json:"domain" gorm:"type:varchar(100);comment:www.amazon.com"` 
    ChildNum string `json:"childNum" gorm:"type:varchar(100);comment:null"` 
    DevIds string `json:"devIds" gorm:"type:varchar(100);comment:null"` 
    DevName string `json:"devName" gorm:"type:varchar(100);comment:null"` 
    VariationStrMap string `json:"variationStrMap" gorm:"type:varchar(100);comment:null"` 
    ListingPrice string `json:"listingPrice" gorm:"type:double;comment:14.99"` 
    ShippingPrice string `json:"shippingPrice" gorm:"type:bigint;comment:0"` 
    BuyBoxWinner string `json:"buyBoxWinner" gorm:"type:varchar(100);comment:false"` 
    BuyBoxCurrency string `json:"buyBoxCurrency" gorm:"type:varchar(100);comment:USD"` 
    TotalFee string `json:"totalFee" gorm:"type:double;comment:7.15"` 
    ReferralFee string `json:"referralFee" gorm:"type:double;comment:2.25"` 
    VariableClosingFee string `json:"variableClosingFee" gorm:"type:bigint;comment:0"` 
    PerItemFee string `json:"perItemFee" gorm:"type:bigint;comment:0"` 
    FbaFees string `json:"fbaFees" gorm:"type:double;comment:4.9"` 
    FeeCurrency string `json:"feeCurrency" gorm:"type:varchar(100);comment:USD"` 
    VartationStr string `json:"vartationStr" gorm:"type:varchar(100);comment:VartationStr"` 
    BsrRankStr string `json:"bsrRankStr" gorm:"type:varchar(100);comment:null"` 
    models.ModelTime
    models.ControlBy
}

func (Product) TableName() string {
    return "product"
}

func (e *Product) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Product) GetId() interface{} {
	return e.Id
}
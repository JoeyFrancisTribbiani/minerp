package models

import (
   // "gorm.io/gorm"

	"go-admin/common/models"
)

type ErpProductSpecifics struct {
    models.Model
    
    Puid string `json:"puid" gorm:"type:varchar(100);comment:父uid"` 
    ShopId string `json:"shopId" gorm:"type:varchar(100);comment:店铺id"` 
    MarketplaceId string `json:"marketplaceId" gorm:"type:varchar(100);comment:市场id"` 
    FullCid string `json:"fullCid" gorm:"type:varchar(100);comment:FullCID"` 
    ParentId string `json:"parentId" gorm:"type:bigint;comment:父ID"` 
    DxmState string `json:"dxmState" gorm:"type:varchar(100);comment:产品状态"` 
    DxmPublishState string `json:"dxmPublishState" gorm:"type:varchar(100);comment:产品发布状态"` 
    ErrMsg string `json:"errMsg" gorm:"type:varchar(100);comment:错误信息"` 
    Asin string `json:"asin" gorm:"type:varchar(100);comment:ASIN"` 
    ListingId string `json:"listingId" gorm:"type:varchar(100);comment:Listing ID"` 
    ParentAsin string `json:"parentAsin" gorm:"type:varchar(100);comment:父ASIN"` 
    ChildAsins string `json:"childAsins" gorm:"type:varchar(100);comment:子ASIN"` 
    IsVariation string `json:"isVariation" gorm:"type:bigint;comment:是否变化"` 
    Sku string `json:"sku" gorm:"type:varchar(100);comment:SKU"` 
    StandardProductType string `json:"standardProductType" gorm:"type:varchar(100);comment:标准产品类型"` 
    StandardProductId string `json:"standardProductId" gorm:"type:varchar(100);comment:标准产品ID"` 
    Title string `json:"title" gorm:"type:varchar(100);comment:产品标题"` 
    Brand string `json:"brand" gorm:"type:varchar(100);comment:品牌"` 
    Manufacturer string `json:"manufacturer" gorm:"type:varchar(100);comment:制造商"` 
    Description string `json:"description" gorm:"type:varchar(100);comment:描述"` 
    BulletPoint string `json:"bulletPoint" gorm:"type:varchar(100);comment:重点"` 
    ConditionType string `json:"conditionType" gorm:"type:varchar(100);comment:条件类型"` 
    ConditionValue string `json:"conditionValue" gorm:"type:varchar(100);comment:条件值"` 
    StandardPrice string `json:"standardPrice" gorm:"type:double;comment:标准价格"` 
    ListingPricing string `json:"listingPricing" gorm:"type:double;comment:价格"` 
    SwitchFulfillmentTo string `json:"switchFulfillmentTo" gorm:"type:varchar(100);comment:配送转换"` 
    FulfillmentChannel string `json:"fulfillmentChannel" gorm:"type:varchar(100);comment:配送渠道"` 
    PartNumber string `json:"partNumber" gorm:"type:varchar(100);comment:产品型号"` 
    MainImage string `json:"mainImage" gorm:"type:varchar(100);comment:主图"` 
    SaleStartDate string `json:"saleStartDate" gorm:"type:varchar(100);comment:开始售出日期"` 
    SaleEndDate string `json:"saleEndDate" gorm:"type:varchar(100);comment:销售结束日期"` 
    SaleSalePrice string `json:"saleSalePrice" gorm:"type:varchar(100);comment:售价"` 
    Quantity string `json:"quantity" gorm:"type:bigint;comment:数量"` 
    OpenDate string `json:"openDate" gorm:"type:varchar(100);comment:开放时间"` 
    ItemDimensions string `json:"itemDimensions" gorm:"type:varchar(100);comment:产品体积"` 
    PackageDimensions string `json:"packageDimensions" gorm:"type:varchar(100);comment:包裹体积"` 
    StandardPriceCurrency string `json:"standardPriceCurrency" gorm:"type:varchar(100);comment:标准价格币种"` 
    Specifics string `json:"specifics" gorm:"type:varchar(100);comment:详情"` 
    VariationChildStr string `json:"variationChildStr" gorm:"type:varchar(100);comment:变体？"` 
    OnlineStatus string `json:"onlineStatus" gorm:"type:varchar(100);comment:在线状态"` 
    SaleNum string `json:"saleNum" gorm:"type:bigint;comment:销量"` 
    Fnsku string `json:"fnsku" gorm:"type:varchar(100);comment:FNSKU"` 
    Warehousing string `json:"warehousing" gorm:"type:bigint;comment:库存"` 
    Unsellable string `json:"unsellable" gorm:"type:bigint;comment:不可售数量"` 
    ReservedQty string `json:"reservedQty" gorm:"type:bigint;comment:预留数量"` 
    PurchaseCost string `json:"purchaseCost" gorm:"type:bigint;comment:购买成本"` 
    HeadTripCost string `json:"headTripCost" gorm:"type:bigint;comment:头程费用"` 
    ShipCost string `json:"shipCost" gorm:"type:varchar(100);comment:运费"` 
    PriceFeedStatus string `json:"priceFeedStatus" gorm:"type:varchar(100);comment:价格推送状态"` 
    InventoryFeedStatus string `json:"inventoryFeedStatus" gorm:"type:varchar(100);comment:库存推送状态"` 
    UpdateStandardPrice string `json:"updateStandardPrice" gorm:"type:varchar(100);comment:更新标准价格"` 
    UpdateQuantity string `json:"updateQuantity" gorm:"type:varchar(100);comment:更新数量"` 
    Rating string `json:"rating" gorm:"type:varchar(100);comment:Rating"` 
    RatingCount string `json:"ratingCount" gorm:"type:varchar(100);comment:Rating数"` 
    Rank string `json:"rank" gorm:"type:varchar(100);comment:排名"` 
    BsrRank string `json:"bsrRank" gorm:"type:varchar(100);comment:BestSeller排名"` 
    CommodityId string `json:"commodityId" gorm:"type:varchar(100);comment:商品id"` 
    DevId string `json:"devId" gorm:"type:varchar(100);comment:开发者id"` 
    CreateId string `json:"createId" gorm:"type:varchar(100);comment:创建id"` 
    UpdateId string `json:"updateId" gorm:"type:varchar(100);comment:更新id"` 
    LpriceUpdateTime string `json:"lpriceUpdateTime" gorm:"type:varchar(100);comment:l价格更新时间"` 
    LpriceStatus string `json:"lpriceStatus" gorm:"type:varchar(100);comment:l价格状态"` 
    ChildList string `json:"childList" gorm:"type:varchar(100);comment:子listing"` 
    SalePrices string `json:"salePrices" gorm:"type:double;comment:销售额"` 
    AdCosts string `json:"adCosts" gorm:"type:double;comment:广告费"` 
    Currency string `json:"currency" gorm:"type:varchar(100);comment:币种"` 
    CostCurrency string `json:"costCurrency" gorm:"type:varchar(100);comment:广告费币种"` 
    CommoditySku string `json:"commoditySku" gorm:"type:varchar(100);comment:商品SKU"` 
    CommodityName string `json:"commodityName" gorm:"type:varchar(100);comment:商品名称"` 
    ShopName string `json:"shopName" gorm:"type:varchar(100);comment:店铺名称"` 
    SiteName string `json:"siteName" gorm:"type:varchar(100);comment:站点名称"` 
    Domain string `json:"domain" gorm:"type:varchar(100);comment:区域域名"` 
    ChildNum string `json:"childNum" gorm:"type:varchar(100);comment:子体数"` 
    DevIds string `json:"devIds" gorm:"type:varchar(100);comment:业务员id"` 
    DevName string `json:"devName" gorm:"type:varchar(100);comment:业务员名称"` 
    VariationStrMap string `json:"variationStrMap" gorm:"type:varchar(100);comment:变化映射"` 
    ListingPrice string `json:"listingPrice" gorm:"type:double;comment:Buy Box价格"` 
    ShippingPrice string `json:"shippingPrice" gorm:"type:bigint;comment:运费"` 
    BuyBoxWinner string `json:"buyBoxWinner" gorm:"type:varchar(100);comment:Buy Box资格"` 
    BuyBoxCurrency string `json:"buyBoxCurrency" gorm:"type:varchar(100);comment:Buy Box币种"` 
    TotalFee string `json:"totalFee" gorm:"type:double;comment:总费用"` 
    ReferralFee string `json:"referralFee" gorm:"type:double;comment:销售佣金"` 
    VariableClosingFee string `json:"variableClosingFee" gorm:"type:bigint;comment:交易手续费"` 
    PerItemFee string `json:"perItemFee" gorm:"type:bigint;comment:计件费用"` 
    FbaFees string `json:"fbaFees" gorm:"type:double;comment:FBA费用"` 
    FeeCurrency string `json:"feeCurrency" gorm:"type:varchar(100);comment:费用币种"` 
    VartationStr string `json:"vartationStr" gorm:"type:varchar(100);comment:变化字符"` 
    BsrRankStr string `json:"bsrRankStr" gorm:"type:varchar(100);comment:BestSeller排名字符"` 
    models.ModelTime
    models.ControlBy
}

func (ErpProductSpecifics) TableName() string {
    return "erp_product_specifics"
}

func (e *ErpProductSpecifics) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *ErpProductSpecifics) GetId() interface{} {
	return e.Id
}
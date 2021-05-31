package models

import (
   // "gorm.io/gorm"

	"go-admin/common/models"
)

type TransactionDetail struct {
    models.Model
    
    PaymentTime string `json:"paymentTime" gorm:"type:varchar(100);comment:结算日期"` 
    EarliestShipDate string `json:"earliestShipDate" gorm:"type:varchar(100);comment:发货日期"` 
    MainImage string `json:"mainImage" gorm:"type:varchar(100);comment:图片"` 
    Asin string `json:"asin" gorm:"type:varchar(100);comment:ASIN"` 
    CommodityName string `json:"commodityName" gorm:"type:varchar(100);comment:品名"` 
    CommoditySku string `json:"commoditySku" gorm:"type:varchar(100);comment:商品SKU"` 
    PurchaseCost string `json:"purchaseCost" gorm:"type:varchar(100);comment:采购成本"` 
    HeadTripCost string `json:"headTripCost" gorm:"type:varchar(100);comment:头程费用"` 
    GrossProfit string `json:"grossProfit" gorm:"type:varchar(100);comment:毛利润"` 
    Currency string `json:"currency" gorm:"type:varchar(100);comment:币种"` 
    Domain string `json:"domain" gorm:"type:varchar(100);comment:domain"` 
    ShopId string `json:"shopId" gorm:"type:bigint;comment:店铺id"` 
    ShopName string `json:"shopName" gorm:"type:varchar(100);comment:店铺名字"` 
    MarketPlaceId string `json:"marketPlaceId" gorm:"type:varchar(100);comment:市场ID"` 
    SettlementId string `json:"settlementId" gorm:"type:varchar(100);comment:结算ID"` 
    Type string `json:"type" gorm:"type:varchar(100);comment:交易类型"` 
    Description string `json:"description" gorm:"type:varchar(100);comment:描述"` 
    OrderId string `json:"orderId" gorm:"type:varchar(100);comment:订单号"` 
    AccountType string `json:"accountType" gorm:"type:varchar(100);comment:类型"` 
    Sku string `json:"sku" gorm:"type:varchar(100);comment:MSKU"` 
    Marketplace string `json:"marketplace" gorm:"type:varchar(100);comment:销售市场"` 
    Fulfillment string `json:"fulfillment" gorm:"type:varchar(100);comment:发货方式"` 
    Quantity string `json:"quantity" gorm:"type:bigint;comment:数量"` 
    ProductSales string `json:"productSales" gorm:"type:varchar(100);comment:销售价格"` 
    ShippingCredits string `json:"shippingCredits" gorm:"type:varchar(100);comment:运费"` 
    GiftWrapCredits string `json:"giftWrapCredits" gorm:"type:varchar(100);comment:礼品包装费"` 
    PromotionalRebates string `json:"promotionalRebates" gorm:"type:varchar(100);comment:促销返点"` 
    ProductSalesTax string `json:"productSalesTax" gorm:"type:varchar(100);comment:销售税"` 
    MarketplaceWithheldTax string `json:"marketplaceWithheldTax" gorm:"type:varchar(100);comment:市场税"` 
    SellingFees string `json:"sellingFees" gorm:"type:varchar(100);comment:平台佣金"` 
    FbaFees string `json:"fbaFees" gorm:"type:varchar(100);comment:FBA Fees"` 
    OtherTransactionFees string `json:"otherTransactionFees" gorm:"type:varchar(100);comment:其他交易费"` 
    Other string `json:"other" gorm:"type:varchar(100);comment:其他费"` 
    Total string `json:"total" gorm:"type:varchar(100);comment:亚马逊结算小计"` 
    Integral string `json:"integral" gorm:"type:varchar(100);comment:积分"` 
    OrderOtherFee string `json:"orderOtherFee" gorm:"type:varchar(100);comment:订单其他佣金"` 
    EvaluationFee string `json:"evaluationFee" gorm:"type:varchar(100);comment:评估费用"` 
    FbmShipCost string `json:"fbmShipCost" gorm:"type:varchar(100);comment:FBM 运费"` 
    models.ModelTime
    models.ControlBy
}

func (TransactionDetail) TableName() string {
    return "erp_transaction_detail"
}

func (e *TransactionDetail) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *TransactionDetail) GetId() interface{} {
	return e.Id
}
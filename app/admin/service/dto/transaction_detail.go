package dto

import (
    "errors"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type TransactionDetailSearch struct {
	dto.Pagination     `search:"-"`
    ShopId string `form:"shopId"  search:"type:exact;column:shop_id;table:erp_transaction_detail" comment:"店铺id"`
    ShopName string `form:"shopName"  search:"type:exact;column:shop_name;table:erp_transaction_detail" comment:"店铺名字"`
    Type string `form:"type"  search:"type:exact;column:type;table:erp_transaction_detail" comment:"交易类型"`
    Sku string `form:"sku"  search:"type:contains;column:sku;table:erp_transaction_detail" comment:"MSKU"`
    Marketplace string `form:"marketplace"  search:"type:exact;column:marketplace;table:erp_transaction_detail" comment:"销售市场"`
}

func (m *TransactionDetailSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *TransactionDetailSearch) Generate() dto.Index {
	o := *m
	return &o
}

func (m *TransactionDetailSearch) Bind(ctx *gin.Context) error {
	log := api.GetRequestLogger(ctx)
    err := ctx.ShouldBind(m)
    if err != nil {
    	log.Errorf("ShouldBind error: %s", err.Error())
    }
    return err
}

type TransactionDetailControl struct {
    Id int `uri:"id" comment:"id"` // id
    PaymentTime string `json:"paymentTime" comment:"结算日期"`
    EarliestShipDate string `json:"earliestShipDate" comment:"发货日期"`
    MainImage string `json:"mainImage" comment:"图片"`
    Asin string `json:"asin" comment:"ASIN"`
    CommodityName string `json:"commodityName" comment:"品名"`
    CommoditySku string `json:"commoditySku" comment:"商品SKU"`
    PurchaseCost string `json:"purchaseCost" comment:"采购成本"`
    HeadTripCost string `json:"headTripCost" comment:"头程费用"`
    GrossProfit string `json:"grossProfit" comment:"毛利润"`
    Currency string `json:"currency" comment:"币种"`
    Domain string `json:"domain" comment:"domain"`
    ShopId string `json:"shopId" comment:"店铺id"`
    ShopName string `json:"shopName" comment:"店铺名字"`
    MarketPlaceId string `json:"marketPlaceId" comment:"市场ID"`
    SettlementId string `json:"settlementId" comment:"结算ID"`
    Type string `json:"type" comment:"交易类型"`
    Description string `json:"description" comment:"描述"`
    OrderId string `json:"orderId" comment:"订单号"`
    AccountType string `json:"accountType" comment:"类型"`
    Sku string `json:"sku" comment:"MSKU"`
    Marketplace string `json:"marketplace" comment:"销售市场"`
    Fulfillment string `json:"fulfillment" comment:"发货方式"`
    Quantity string `json:"quantity" comment:"数量"`
    ProductSales string `json:"productSales" comment:"销售价格"`
    ShippingCredits string `json:"shippingCredits" comment:"运费"`
    GiftWrapCredits string `json:"giftWrapCredits" comment:"礼品包装费"`
    PromotionalRebates string `json:"promotionalRebates" comment:"促销返点"`
    ProductSalesTax string `json:"productSalesTax" comment:"销售税"`
    MarketplaceWithheldTax string `json:"marketplaceWithheldTax" comment:"市场税"`
    SellingFees string `json:"sellingFees" comment:"平台佣金"`
    FbaFees string `json:"fbaFees" comment:"FBA Fees"`
    OtherTransactionFees string `json:"otherTransactionFees" comment:"其他交易费"`
    Other string `json:"other" comment:"其他费"`
    Total string `json:"total" comment:"亚马逊结算小计"`
    Integral string `json:"integral" comment:"积分"`
    OrderOtherFee string `json:"orderOtherFee" comment:"订单其他佣金"`
    EvaluationFee string `json:"evaluationFee" comment:"评估费用"`
    FbmShipCost string `json:"fbmShipCost" comment:"FBM 运费"`
}

func (s *TransactionDetailControl) Generate() dto.Control {
	o := *s
	return &o
}

func (s *TransactionDetailControl) Bind(ctx *gin.Context) error {
	log := api.GetRequestLogger(ctx)
    err := ctx.ShouldBindUri(s)
    if err != nil {
        log.Errorf("ShouldBindUri error: %s", err.Error())
        return errors.New("数据绑定出错")
    }
    err = ctx.ShouldBind(s)
    if err != nil {
        log.Errorf("ShouldBind error: %s", err.Error())
        err = errors.New("数据绑定出错")
    }
    return err
}

func (s *TransactionDetailControl) GenerateM() (common.ActiveRecord, error) {
	return &models.TransactionDetail{
        Model:     common.Model{ Id: s.Id },
        PaymentTime:  s.PaymentTime,
        EarliestShipDate:  s.EarliestShipDate,
        MainImage:  s.MainImage,
        Asin:  s.Asin,
        CommodityName:  s.CommodityName,
        CommoditySku:  s.CommoditySku,
        PurchaseCost:  s.PurchaseCost,
        HeadTripCost:  s.HeadTripCost,
        GrossProfit:  s.GrossProfit,
        Currency:  s.Currency,
        Domain:  s.Domain,
        ShopId:  s.ShopId,
        ShopName:  s.ShopName,
        MarketPlaceId:  s.MarketPlaceId,
        SettlementId:  s.SettlementId,
        Type:  s.Type,
        Description:  s.Description,
        OrderId:  s.OrderId,
        AccountType:  s.AccountType,
        Sku:  s.Sku,
        Marketplace:  s.Marketplace,
        Fulfillment:  s.Fulfillment,
        Quantity:  s.Quantity,
        ProductSales:  s.ProductSales,
        ShippingCredits:  s.ShippingCredits,
        GiftWrapCredits:  s.GiftWrapCredits,
        PromotionalRebates:  s.PromotionalRebates,
        ProductSalesTax:  s.ProductSalesTax,
        MarketplaceWithheldTax:  s.MarketplaceWithheldTax,
        SellingFees:  s.SellingFees,
        FbaFees:  s.FbaFees,
        OtherTransactionFees:  s.OtherTransactionFees,
        Other:  s.Other,
        Total:  s.Total,
        Integral:  s.Integral,
        OrderOtherFee:  s.OrderOtherFee,
        EvaluationFee:  s.EvaluationFee,
        FbmShipCost:  s.FbmShipCost,
	}, nil
}

func (s *TransactionDetailControl) GetId() interface{} {
	return s.Id
}

type TransactionDetailById struct {
	dto.ObjectById
}

func (s *TransactionDetailById) GetId() interface{} {
	if len(s.Ids) > 0 {
		s.Ids = append(s.Ids, s.Id)
		return s.Ids
	}
	return s.Id
}

func (s *TransactionDetailById) Bind(ctx *gin.Context) error {
	log := api.GetRequestLogger(ctx)
	err := ctx.ShouldBindUri(s)
	if err != nil {
		log.Errorf("ShouldBindUri error: %s", err.Error())
		return errors.New("数据绑定出错")
	}
	err = ctx.ShouldBind(s)
	if err != nil {
		log.Errorf("ShouldBind error: %s", err.Error())
		err = errors.New("数据绑定出错")
	}
	return err
}

func (s *TransactionDetailById) SetUpdateBy(id int) {

}

func (s *TransactionDetailById) Generate() dto.Control {
	o := *s
	return &o
}

func (s *TransactionDetailById) GenerateM() (common.ActiveRecord, error) {
	return &models.TransactionDetail{}, nil
}
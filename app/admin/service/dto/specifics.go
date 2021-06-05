package dto

import (
    "errors"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type ErpProductSpecificsSearch struct {
	dto.Pagination     `search:"-"`
    ShopId string `form:"shopId"  search:"type:exact;column:shop_id;table:erp_product_specifics" comment:"店铺id"`
    MarketplaceId string `form:"marketplaceId"  search:"type:exact;column:marketplace_id;table:erp_product_specifics" comment:"市场id"`
    DxmState string `form:"dxmState"  search:"type:exact;column:dxm_state;table:erp_product_specifics" comment:"产品状态"`
    Sku string `form:"sku"  search:"type:exact;column:sku;table:erp_product_specifics" comment:"SKU"`
    Title string `form:"title"  search:"type:exact;column:title;table:erp_product_specifics" comment:"产品标题"`
    SaleStartDate string `form:"saleStartDate"  search:"type:exact;column:sale_start_date;table:erp_product_specifics" comment:"开始售出日期"`
    SaleEndDate string `form:"saleEndDate"  search:"type:exact;column:sale_end_date;table:erp_product_specifics" comment:"销售结束日期"`
    Fnsku string `form:"fnsku"  search:"type:exact;column:fnsku;table:erp_product_specifics" comment:"FNSKU"`
    CommodityName string `form:"commodityName"  search:"type:exact;column:commodity_name;table:erp_product_specifics" comment:"商品名称"`
    SiteName string `form:"siteName"  search:"type:exact;column:site_name;table:erp_product_specifics" comment:"站点名称"`
}

func (m *ErpProductSpecificsSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *ErpProductSpecificsSearch) Generate() dto.Index {
	o := *m
	return &o
}

func (m *ErpProductSpecificsSearch) Bind(ctx *gin.Context) error {
	log := api.GetRequestLogger(ctx)
    err := ctx.ShouldBind(m)
    if err != nil {
    	log.Errorf("ShouldBind error: %s", err.Error())
    }
    return err
}

type ErpProductSpecificsControl struct {
    Id int `uri:"id" comment:"id"` // id
    Puid string `json:"puid" comment:"父uid"`
    ShopId string `json:"shopId" comment:"店铺id"`
    MarketplaceId string `json:"marketplaceId" comment:"市场id"`
    FullCid string `json:"fullCid" comment:"FullCID"`
    ParentId string `json:"parentId" comment:"父ID"`
    DxmState string `json:"dxmState" comment:"产品状态"`
    DxmPublishState string `json:"dxmPublishState" comment:"产品发布状态"`
    ErrMsg string `json:"errMsg" comment:"错误信息"`
    Asin string `json:"asin" comment:"ASIN"`
    ListingId string `json:"listingId" comment:"Listing ID"`
    ParentAsin string `json:"parentAsin" comment:"父ASIN"`
    ChildAsins string `json:"childAsins" comment:"子ASIN"`
    IsVariation string `json:"isVariation" comment:"是否变化"`
    Sku string `json:"sku" comment:"SKU"`
    StandardProductType string `json:"standardProductType" comment:"标准产品类型"`
    StandardProductId string `json:"standardProductId" comment:"标准产品ID"`
    Title string `json:"title" comment:"产品标题"`
    Brand string `json:"brand" comment:"品牌"`
    Manufacturer string `json:"manufacturer" comment:"制造商"`
    Description string `json:"description" comment:"描述"`
    BulletPoint string `json:"bulletPoint" comment:"重点"`
    ConditionType string `json:"conditionType" comment:"条件类型"`
    ConditionValue string `json:"conditionValue" comment:"条件值"`
    StandardPrice string `json:"standardPrice" comment:"标准价格"`
    ListingPricing string `json:"listingPricing" comment:"价格"`
    SwitchFulfillmentTo string `json:"switchFulfillmentTo" comment:"配送转换"`
    FulfillmentChannel string `json:"fulfillmentChannel" comment:"配送渠道"`
    PartNumber string `json:"partNumber" comment:"产品型号"`
    MainImage string `json:"mainImage" comment:"主图"`
    SaleStartDate string `json:"saleStartDate" comment:"开始售出日期"`
    SaleEndDate string `json:"saleEndDate" comment:"销售结束日期"`
    SaleSalePrice string `json:"saleSalePrice" comment:"售价"`
    Quantity string `json:"quantity" comment:"数量"`
    OpenDate string `json:"openDate" comment:"开放时间"`
    ItemDimensions string `json:"itemDimensions" comment:"产品体积"`
    PackageDimensions string `json:"packageDimensions" comment:"包裹体积"`
    StandardPriceCurrency string `json:"standardPriceCurrency" comment:"标准价格币种"`
    Specifics string `json:"specifics" comment:"详情"`
    VariationChildStr string `json:"variationChildStr" comment:"变体？"`
    OnlineStatus string `json:"onlineStatus" comment:"在线状态"`
    SaleNum string `json:"saleNum" comment:"销量"`
    Fnsku string `json:"fnsku" comment:"FNSKU"`
    Warehousing string `json:"warehousing" comment:"库存"`
    Unsellable string `json:"unsellable" comment:"不可售数量"`
    ReservedQty string `json:"reservedQty" comment:"预留数量"`
    PurchaseCost string `json:"purchaseCost" comment:"购买成本"`
    HeadTripCost string `json:"headTripCost" comment:"头程费用"`
    ShipCost string `json:"shipCost" comment:"运费"`
    PriceFeedStatus string `json:"priceFeedStatus" comment:"价格推送状态"`
    InventoryFeedStatus string `json:"inventoryFeedStatus" comment:"库存推送状态"`
    UpdateStandardPrice string `json:"updateStandardPrice" comment:"更新标准价格"`
    UpdateQuantity string `json:"updateQuantity" comment:"更新数量"`
    Rating string `json:"rating" comment:"Rating"`
    RatingCount string `json:"ratingCount" comment:"Rating数"`
    Rank string `json:"rank" comment:"排名"`
    BsrRank string `json:"bsrRank" comment:"BestSeller排名"`
    CommodityId string `json:"commodityId" comment:"商品id"`
    DevId string `json:"devId" comment:"开发者id"`
    CreateId string `json:"createId" comment:"创建id"`
    UpdateId string `json:"updateId" comment:"更新id"`
    LpriceUpdateTime string `json:"lpriceUpdateTime" comment:"l价格更新时间"`
    LpriceStatus string `json:"lpriceStatus" comment:"l价格状态"`
    ChildList string `json:"childList" comment:"子listing"`
    SalePrices string `json:"salePrices" comment:"销售额"`
    AdCosts string `json:"adCosts" comment:"广告费"`
    Currency string `json:"currency" comment:"币种"`
    CostCurrency string `json:"costCurrency" comment:"广告费币种"`
    CommoditySku string `json:"commoditySku" comment:"商品SKU"`
    CommodityName string `json:"commodityName" comment:"商品名称"`
    ShopName string `json:"shopName" comment:"店铺名称"`
    SiteName string `json:"siteName" comment:"站点名称"`
    Domain string `json:"domain" comment:"区域域名"`
    ChildNum string `json:"childNum" comment:"子体数"`
    DevIds string `json:"devIds" comment:"业务员id"`
    DevName string `json:"devName" comment:"业务员名称"`
    VariationStrMap string `json:"variationStrMap" comment:"变化映射"`
    ListingPrice string `json:"listingPrice" comment:"Buy Box价格"`
    ShippingPrice string `json:"shippingPrice" comment:"运费"`
    BuyBoxWinner string `json:"buyBoxWinner" comment:"Buy Box资格"`
    BuyBoxCurrency string `json:"buyBoxCurrency" comment:"Buy Box币种"`
    TotalFee string `json:"totalFee" comment:"总费用"`
    ReferralFee string `json:"referralFee" comment:"销售佣金"`
    VariableClosingFee string `json:"variableClosingFee" comment:"交易手续费"`
    PerItemFee string `json:"perItemFee" comment:"计件费用"`
    FbaFees string `json:"fbaFees" comment:"FBA费用"`
    FeeCurrency string `json:"feeCurrency" comment:"费用币种"`
    VartationStr string `json:"vartationStr" comment:"变化字符"`
    BsrRankStr string `json:"bsrRankStr" comment:"BestSeller排名字符"`
}

func (s *ErpProductSpecificsControl) Generate() dto.Control {
	o := *s
	return &o
}

func (s *ErpProductSpecificsControl) Bind(ctx *gin.Context) error {
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

func (s *ErpProductSpecificsControl) GenerateM() (common.ActiveRecord, error) {
	return &models.ErpProductSpecifics{
        Model:     common.Model{ Id: s.Id },
        Puid:  s.Puid,
        ShopId:  s.ShopId,
        MarketplaceId:  s.MarketplaceId,
        FullCid:  s.FullCid,
        ParentId:  s.ParentId,
        DxmState:  s.DxmState,
        DxmPublishState:  s.DxmPublishState,
        ErrMsg:  s.ErrMsg,
        Asin:  s.Asin,
        ListingId:  s.ListingId,
        ParentAsin:  s.ParentAsin,
        ChildAsins:  s.ChildAsins,
        IsVariation:  s.IsVariation,
        Sku:  s.Sku,
        StandardProductType:  s.StandardProductType,
        StandardProductId:  s.StandardProductId,
        Title:  s.Title,
        Brand:  s.Brand,
        Manufacturer:  s.Manufacturer,
        Description:  s.Description,
        BulletPoint:  s.BulletPoint,
        ConditionType:  s.ConditionType,
        ConditionValue:  s.ConditionValue,
        StandardPrice:  s.StandardPrice,
        ListingPricing:  s.ListingPricing,
        SwitchFulfillmentTo:  s.SwitchFulfillmentTo,
        FulfillmentChannel:  s.FulfillmentChannel,
        PartNumber:  s.PartNumber,
        MainImage:  s.MainImage,
        SaleStartDate:  s.SaleStartDate,
        SaleEndDate:  s.SaleEndDate,
        SaleSalePrice:  s.SaleSalePrice,
        Quantity:  s.Quantity,
        OpenDate:  s.OpenDate,
        ItemDimensions:  s.ItemDimensions,
        PackageDimensions:  s.PackageDimensions,
        StandardPriceCurrency:  s.StandardPriceCurrency,
        Specifics:  s.Specifics,
        VariationChildStr:  s.VariationChildStr,
        OnlineStatus:  s.OnlineStatus,
        SaleNum:  s.SaleNum,
        Fnsku:  s.Fnsku,
        Warehousing:  s.Warehousing,
        Unsellable:  s.Unsellable,
        ReservedQty:  s.ReservedQty,
        PurchaseCost:  s.PurchaseCost,
        HeadTripCost:  s.HeadTripCost,
        ShipCost:  s.ShipCost,
        PriceFeedStatus:  s.PriceFeedStatus,
        InventoryFeedStatus:  s.InventoryFeedStatus,
        UpdateStandardPrice:  s.UpdateStandardPrice,
        UpdateQuantity:  s.UpdateQuantity,
        Rating:  s.Rating,
        RatingCount:  s.RatingCount,
        Rank:  s.Rank,
        BsrRank:  s.BsrRank,
        CommodityId:  s.CommodityId,
        DevId:  s.DevId,
        CreateId:  s.CreateId,
        UpdateId:  s.UpdateId,
        LpriceUpdateTime:  s.LpriceUpdateTime,
        LpriceStatus:  s.LpriceStatus,
        ChildList:  s.ChildList,
        SalePrices:  s.SalePrices,
        AdCosts:  s.AdCosts,
        Currency:  s.Currency,
        CostCurrency:  s.CostCurrency,
        CommoditySku:  s.CommoditySku,
        CommodityName:  s.CommodityName,
        ShopName:  s.ShopName,
        SiteName:  s.SiteName,
        Domain:  s.Domain,
        ChildNum:  s.ChildNum,
        DevIds:  s.DevIds,
        DevName:  s.DevName,
        VariationStrMap:  s.VariationStrMap,
        ListingPrice:  s.ListingPrice,
        ShippingPrice:  s.ShippingPrice,
        BuyBoxWinner:  s.BuyBoxWinner,
        BuyBoxCurrency:  s.BuyBoxCurrency,
        TotalFee:  s.TotalFee,
        ReferralFee:  s.ReferralFee,
        VariableClosingFee:  s.VariableClosingFee,
        PerItemFee:  s.PerItemFee,
        FbaFees:  s.FbaFees,
        FeeCurrency:  s.FeeCurrency,
        VartationStr:  s.VartationStr,
        BsrRankStr:  s.BsrRankStr,
	}, nil
}

func (s *ErpProductSpecificsControl) GetId() interface{} {
	return s.Id
}

type ErpProductSpecificsById struct {
	dto.ObjectById
}

func (s *ErpProductSpecificsById) GetId() interface{} {
	if len(s.Ids) > 0 {
		s.Ids = append(s.Ids, s.Id)
		return s.Ids
	}
	return s.Id
}

func (s *ErpProductSpecificsById) Bind(ctx *gin.Context) error {
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

func (s *ErpProductSpecificsById) SetUpdateBy(id int) {

}

func (s *ErpProductSpecificsById) Generate() dto.Control {
	o := *s
	return &o
}

func (s *ErpProductSpecificsById) GenerateM() (common.ActiveRecord, error) {
	return &models.ErpProductSpecifics{}, nil
}
package dto

import (
    "errors"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type ProductSearch struct {
	dto.Pagination     `search:"-"`
    Title string `form:"title"  search:"type:exact;column:title;table:product" comment:"product title"`
    Brand string `form:"brand"  search:"type:exact;column:brand;table:product" comment:"null"`
    Manufacturer string `form:"manufacturer"  search:"type:exact;column:manufacturer;table:product" comment:"null"`
    Description string `form:"description"  search:"type:exact;column:description;table:product" comment:"null"`
    ListingPricing string `form:"listingPricing"  search:"type:exact;column:listing_pricing;table:product" comment:"14.99"`
    SwitchFulfillmentTo string `form:"switchFulfillmentTo"  search:"type:exact;column:switch_fulfillment_to;table:product" comment:"AFN"`
    PartNumber string `form:"partNumber"  search:"type:exact;column:part_number;table:product" comment:"null"`
    MainImage string `form:"mainImage"  search:"type:exact;column:main_image;table:product" comment:"https://m.media-amazon.com/images/I/613gR7poDQL._SL75_.jpg"`
    SaleStartDate string `form:"saleStartDate"  search:"type:exact;column:sale_start_date;table:product" comment:"null"`
    SaleEndDate string `form:"saleEndDate"  search:"type:exact;column:sale_end_date;table:product" comment:"null"`
}

func (m *ProductSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *ProductSearch) Generate() dto.Index {
	o := *m
	return &o
}

func (m *ProductSearch) Bind(ctx *gin.Context) error {
	log := api.GetRequestLogger(ctx)
    err := ctx.ShouldBind(m)
    if err != nil {
    	log.Errorf("ShouldBind error: %s", err.Error())
    }
    return err
}

type ProductControl struct {
    Id int `uri:"id" comment:"product_id"` // product_id
    DataId string `json:"dataId" comment:"data_id"`
    Puid string `json:"puid" comment:"null"`
    ShopId string `json:"shopId" comment:"4462"`
    MarketplaceId string `json:"marketplaceId" comment:"ATVPDKIKX0DER"`
    FullCid string `json:"fullCid" comment:""`
    ParentId string `json:"parentId" comment:"4762270"`
    DxmState string `json:"dxmState" comment:"online_amazon"`
    DxmPublishState string `json:"dxmPublishState" comment:"null"`
    ErrMsg string `json:"errMsg" comment:"null"`
    Asin string `json:"asin" comment:"B08KGSXPZS"`
    ListingId string `json:"listingId" comment:"0324YYDEASD"`
    ParentAsin string `json:"parentAsin" comment:"null"`
    ChildAsins string `json:"childAsins" comment:"null"`
    IsVariation string `json:"isVariation" comment:"0"`
    Sku string `json:"sku" comment:"K7-PQU4-OE8V"`
    StandardProductType string `json:"standardProductType" comment:"null"`
    StandardProductId string `json:"standardProductId" comment:"null"`
    Title string `json:"title" comment:"product title"`
    Brand string `json:"brand" comment:"null"`
    Manufacturer string `json:"manufacturer" comment:"null"`
    Description string `json:"description" comment:"null"`
    BulletPoint string `json:"bulletPoint" comment:"null"`
    ConditionType string `json:"conditionType" comment:"null"`
    ConditionValue string `json:"conditionValue" comment:"null"`
    StandardPrice string `json:"standardPrice" comment:"14.99"`
    ListingPricing string `json:"listingPricing" comment:"14.99"`
    SwitchFulfillmentTo string `json:"switchFulfillmentTo" comment:"AFN"`
    FulfillmentChannel string `json:"fulfillmentChannel" comment:"null"`
    PartNumber string `json:"partNumber" comment:"null"`
    MainImage string `json:"mainImage" comment:"https://m.media-amazon.com/images/I/613gR7poDQL._SL75_.jpg"`
    SaleStartDate string `json:"saleStartDate" comment:"null"`
    SaleEndDate string `json:"saleEndDate" comment:"null"`
    SaleSalePrice string `json:"saleSalePrice" comment:"null"`
    Quantity string `json:"quantity" comment:"0"`
    OpenDate string `json:"openDate" comment:"2021-03-23 19:09:08"`
    ItemDimensions string `json:"itemDimensions" comment:"null"`
    PackageDimensions string `json:"packageDimensions" comment:"null"`
    StandardPriceCurrency string `json:"standardPriceCurrency" comment:"USD"`
    Specifics string `json:"specifics" comment:"null"`
    VariationChildStr string `json:"variationChildStr" comment:"null"`
    OnlineStatus string `json:"onlineStatus" comment:"Active"`
    SaleNum string `json:"saleNum" comment:"54"`
    Fnsku string `json:"fnsku" comment:"X002UH0IRH"`
    Warehousing string `json:"warehousing" comment:"542"`
    Unsellable string `json:"unsellable" comment:"0"`
    ReservedQty string `json:"reservedQty" comment:"17"`
    PurchaseCost string `json:"purchaseCost" comment:"0"`
    HeadTripCost string `json:"headTripCost" comment:"0"`
    ShipCost string `json:"shipCost" comment:"null"`
    PriceFeedStatus string `json:"priceFeedStatus" comment:"null"`
    InventoryFeedStatus string `json:"inventoryFeedStatus" comment:"null"`
    UpdateStandardPrice string `json:"updateStandardPrice" comment:"null"`
    UpdateQuantity string `json:"updateQuantity" comment:"null"`
    Rating string `json:"rating" comment:"null"`
    RatingCount string `json:"ratingCount" comment:"null"`
    Rank string `json:"rank" comment:"null"`
    BsrRank string `json:"bsrRank" comment:"fuck"`
    CommodityId string `json:"commodityId" comment:"null"`
    DevId string `json:"devId" comment:""`
    CreateId string `json:"createId" comment:"null"`
    UpdateId string `json:"updateId" comment:"null"`
    LpriceUpdateTime string `json:"lpriceUpdateTime" comment:"null"`
    LpriceStatus string `json:"lpriceStatus" comment:"null"`
    ChildList string `json:"childList" comment:"null"`
    SalePrices string `json:"salePrices" comment:"809.46"`
    AdCosts string `json:"adCosts" comment:"130.33"`
    Currency string `json:"currency" comment:"USD"`
    CostCurrency string `json:"costCurrency" comment:"null"`
    CommoditySku string `json:"commoditySku" comment:"null"`
    CommodityName string `json:"commodityName" comment:"null"`
    ShopName string `json:"shopName" comment:"A3LR0SXN8QVIE0-na-US"`
    SiteName string `json:"siteName" comment:"美国"`
    Domain string `json:"domain" comment:"www.amazon.com"`
    ChildNum string `json:"childNum" comment:"null"`
    DevIds string `json:"devIds" comment:"null"`
    DevName string `json:"devName" comment:"null"`
    VariationStrMap string `json:"variationStrMap" comment:"null"`
    ListingPrice string `json:"listingPrice" comment:"14.99"`
    ShippingPrice string `json:"shippingPrice" comment:"0"`
    BuyBoxWinner string `json:"buyBoxWinner" comment:"false"`
    BuyBoxCurrency string `json:"buyBoxCurrency" comment:"USD"`
    TotalFee string `json:"totalFee" comment:"7.15"`
    ReferralFee string `json:"referralFee" comment:"2.25"`
    VariableClosingFee string `json:"variableClosingFee" comment:"0"`
    PerItemFee string `json:"perItemFee" comment:"0"`
    FbaFees string `json:"fbaFees" comment:"4.9"`
    FeeCurrency string `json:"feeCurrency" comment:"USD"`
    VartationStr string `json:"vartationStr" comment:""`
    BsrRankStr string `json:"bsrRankStr" comment:"null"`
}

func (s *ProductControl) Generate() dto.Control {
	o := *s
	return &o
}

func (s *ProductControl) Bind(ctx *gin.Context) error {
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

func (s *ProductControl) GenerateM() (common.ActiveRecord, error) {
	return &models.Product{
        Model:     common.Model{ Id: s.Id },
        DataId:  s.DataId,
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

func (s *ProductControl) GetId() interface{} {
	return s.Id
}

type ProductById struct {
	dto.ObjectById
}

func (s *ProductById) GetId() interface{} {
	if len(s.Ids) > 0 {
		s.Ids = append(s.Ids, s.Id)
		return s.Ids
	}
	return s.Id
}

func (s *ProductById) Bind(ctx *gin.Context) error {
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

func (s *ProductById) SetUpdateBy(id int) {

}

func (s *ProductById) Generate() dto.Control {
	o := *s
	return &o
}

func (s *ProductById) GenerateM() (common.ActiveRecord, error) {
	return &models.Product{}, nil
}
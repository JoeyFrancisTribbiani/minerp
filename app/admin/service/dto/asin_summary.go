package dto

import (
    "errors"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type ErpFinanceAsinSummarySearch struct {
	dto.Pagination     `search:"-"`
    CommoditySku string `form:"commoditySku"  search:"type:contains;column:commodity_sku;table:erp_finance_asin_summary" comment:"商品SKU"`
    CommodityName string `form:"commodityName"  search:"type:contains;column:commodity_name;table:erp_finance_asin_summary" comment:"商品名称"`
    ShopId string `form:"shopId"  search:"type:exact;column:shop_id;table:erp_finance_asin_summary" comment:"店铺id"`
    MarketplaceId string `form:"marketplaceId"  search:"type:exact;column:marketplace_id;table:erp_finance_asin_summary" comment:"市场id"`
    Currency string `form:"currency"  search:"type:exact;column:currency;table:erp_finance_asin_summary" comment:"币种"`
    Asin string `form:"asin"  search:"type:contains;column:asin;table:erp_finance_asin_summary" comment:"ASIN"`
    Sku string `form:"sku"  search:"type:exact;column:sku;table:erp_finance_asin_summary" comment:"SKU"`
    Month string `form:"month"  search:"type:exact;column:month;table:erp_finance_asin_summary" comment:"月份"`
}

func (m *ErpFinanceAsinSummarySearch) GetNeedSearch() interface{} {
	return *m
}

func (m *ErpFinanceAsinSummarySearch) Generate() dto.Index {
	o := *m
	return &o
}

func (m *ErpFinanceAsinSummarySearch) Bind(ctx *gin.Context) error {
	log := api.GetRequestLogger(ctx)
    err := ctx.ShouldBind(m)
    if err != nil {
    	log.Errorf("ShouldBind error: %s", err.Error())
    }
    return err
}

type ErpFinanceAsinSummaryControl struct {
    Id int `uri:"id" comment:"id"` // id
    MainImg string `json:"mainImg" comment:"主图"`
    CommoditySku string `json:"commoditySku" comment:"商品SKU"`
    CommodityName string `json:"commodityName" comment:"商品名称"`
    ShopId string `json:"shopId" comment:"店铺id"`
    ShopName string `json:"shopName" comment:"店铺名称"`
    MarketplaceId string `json:"marketplaceId" comment:"市场id"`
    MarketplaceName string `json:"marketplaceName" comment:"市场名称"`
    Currency string `json:"currency" comment:"币种"`
    Asin string `json:"asin" comment:"ASIN"`
    AsinList string `json:"asinList" comment:"ASIN列表"`
    CommodityList string `json:"commodityList" comment:"商品列表"`
    SumOrderFee string `json:"sumOrderFee" comment:"订单费用"`
    Domain string `json:"domain" comment:"区域域名"`
    ParentAsin string `json:"parentAsin" comment:"父ASIN"`
    Sku string `json:"sku" comment:"SKU"`
    Month string `json:"month" comment:"月份"`
    ProductSales string `json:"productSales" comment:"销售额"`
    FbaSales string `json:"fbaSales" comment:"FBA销售额"`
    FbmSales string `json:"fbmSales" comment:"FBM销售额"`
    CpcSales string `json:"cpcSales" comment:"CPC销售额"`
    FbaSalesNum string `json:"fbaSalesNum" comment:"亚马逊FBA销量"`
    FbmSalesNum string `json:"fbmSalesNum" comment:"亚马逊FBM销量"`
    MultiChannelNum string `json:"multiChannelNum" comment:"多渠道销量"`
    FbaReissueNum string `json:"fbaReissueNum" comment:"FBA补发量"`
    CpcSalesNum string `json:"cpcSalesNum" comment:"CPC销量"`
    CpcSdSalesNum string `json:"cpcSdSalesNum" comment:"CPC广告销量"`
    CpcSpSalesNum string `json:"cpcSpSalesNum" comment:"CPC推广销量"`
    FbaShippingCredits string `json:"fbaShippingCredits" comment:"FBA运输信用分"`
    FbmShippingCredits string `json:"fbmShippingCredits" comment:"FBM运输信用分"`
    PromotionalRebates string `json:"promotionalRebates" comment:"促销折扣"`
    AdjCompensation string `json:"adjCompensation" comment:"库存调整赔偿"`
    AdjCompensationPcFee string `json:"adjCompensationPcFee" comment:"FBA库存赔偿"`
    AdjCompensationNum string `json:"adjCompensationNum" comment:"库存调整数量"`
    Refund string `json:"refund" comment:"退款"`
    FbaRefund string `json:"fbaRefund" comment:"FBA退款"`
    FbmRefund string `json:"fbmRefund" comment:"FBM退款"`
    RefundNum string `json:"refundNum" comment:"退款量"`
    FbaRefundNum string `json:"fbaRefundNum" comment:"FBA退款量"`
    FbmRefundNum string `json:"fbmRefundNum" comment:"FBM退款量"`
    RefundRate string `json:"refundRate" comment:"退款率"`
    ReturnNum string `json:"returnNum" comment:"退货量"`
    ReturnSellableNum string `json:"returnSellableNum" comment:"退货可售商品数"`
    ReturnUnsellableNum string `json:"returnUnsellableNum" comment:"退货不可售商品数"`
    ReturnRate string `json:"returnRate" comment:"退货比例"`
    SaleOrderFee string `json:"saleOrderFee" comment:"销售订单"`
    SaleSellingFee string `json:"saleSellingFee" comment:"销售成本"`
    SaleFbaFee string `json:"saleFbaFee" comment:"FBA发货费"`
    SaleFbmFee string `json:"saleFbmFee" comment:"FBM发货费"`
    SaleOtherFee string `json:"saleOtherFee" comment:"其他费"`
    MultiChannelOrderFee string `json:"multiChannelOrderFee" comment:"多渠道订单"`
    McFbaShipFee string `json:"mcFbaShipFee" comment:"FBA发货费"`
    RefundOrderFees string `json:"refundOrderFees" comment:"退款订单"`
    RefundSellingFee string `json:"refundSellingFee" comment:"平台费"`
    RefundFbaFee string `json:"refundFbaFee" comment:"FBA发货费"`
    RefundOtherFee string `json:"refundOtherFee" comment:"其他费"`
    SumCpcCost string `json:"sumCpcCost" comment:"广告费"`
    CpcCostDiff string `json:"cpcCostDiff" comment:"CPC广告费差异"`
    CpcSpCost string `json:"cpcSpCost" comment:"CPC推广费"`
    CpcSdCost string `json:"cpcSdCost" comment:"CPC广告成本"`
    SumPromoteFee string `json:"sumPromoteFee" comment:"推广费用"`
    LdPromoteFee string `json:"ldPromoteFee" comment:"LD费"`
    CouponsPromoteFee string `json:"couponsPromoteFee" comment:"优惠劵"`
    ErpPromoteFee string `json:"erpPromoteFee" comment:"早期评论人计划"`
    SumFbaStorageFee string `json:"sumFbaStorageFee" comment:"仓储费用"`
    FbaStorageFee string `json:"fbaStorageFee" comment:"FBA月仓储费"`
    FbaLongStorageFee string `json:"fbaLongStorageFee" comment:"FBA长期仓储费"`
    FbaStorageFeeDiff string `json:"fbaStorageFeeDiff" comment:"FBA月仓储费差"`
    SumStorageOtherFee string `json:"sumStorageOtherFee" comment:"其他仓储费用"`
    FbaDisposalNum string `json:"fbaDisposalNum" comment:"FBA销毁量"`
    FbaDisposalFee string `json:"fbaDisposalFee" comment:"FBA销毁费"`
    FbaRemovalNum string `json:"fbaRemovalNum" comment:"FBA移除量"`
    FbaRemovalFee string `json:"fbaRemovalFee" comment:"FBA移除费"`
    FbaReturnFee string `json:"fbaReturnFee" comment:"退货入仓费"`
    FbaShipFee string `json:"fbaShipFee" comment:"亚马逊合作承运费"`
    FbaStorageOtherFee string `json:"fbaStorageOtherFee" comment:"其他费"`
    FbaInventoryPlacementFee string `json:"fbaInventoryPlacementFee" comment:"库存配置费"`
    SumFbaAdjustmentFee string `json:"sumFbaAdjustmentFee" comment:"FBA仓储调整费用合计"`
    FbaAdjustmentFee string `json:"fbaAdjustmentFee" comment:"库存调整费"`
    FbaAdjustmentDiffFee string `json:"fbaAdjustmentDiffFee" comment:"FBA库存调整费差异"`
    SumOtherFee string `json:"sumOtherFee" comment:"其他费合计"`
    SubscriptionFee string `json:"subscriptionFee" comment:"订阅费"`
    OtherFee string `json:"otherFee" comment:"其他费"`
    Tax string `json:"tax" comment:"税费"`
    SumPurchaseFee string `json:"sumPurchaseFee" comment:"总采购成本"`
    PurchasePcFee string `json:"purchasePcFee" comment:"采购成本"`
    PurchaseDpcFee string `json:"purchaseDpcFee" comment:"采购成本（销毁）"`
    PurchaseRtpcFee string `json:"purchaseRtpcFee" comment:"采购成本（退货）"`
    PurchaseAcpcFee string `json:"purchaseAcpcFee" comment:"采购成本（赔偿）"`
    PurchaseMcpcFee string `json:"purchaseMcpcFee" comment:"采购成本"`
    PurchaseRpcFee string `json:"purchaseRpcFee" comment:"采购成本"`
    SumHeadTripFee string `json:"sumHeadTripFee" comment:"总头程费用"`
    HeadTripPcFee string `json:"headTripPcFee" comment:"头程费用"`
    HeadTripDpcFee string `json:"headTripDpcFee" comment:"头程费用（销毁）"`
    HeadTripRtpcFee string `json:"headTripRtpcFee" comment:"头程费用（退货）"`
    HeadTripAcpcFee string `json:"headTripAcpcFee" comment:"头程费用（赔偿）"`
    HeadTripMcpcFee string `json:"headTripMcpcFee" comment:"头程费用"`
    HeadTripRpcFee string `json:"headTripRpcFee" comment:"头程费用"`
    SumBatchPurchaseFee string `json:"sumBatchPurchaseFee" comment:"批次采购成本合计"`
    BatchPurchasePcFee string `json:"batchPurchasePcFee" comment:"批次采购成本"`
    SumBatchPurchaseOtherFee string `json:"sumBatchPurchaseOtherFee" comment:"批次采购其他成本合计"`
    BatchPurchaseCrpcFee string `json:"batchPurchaseCrpcFee" comment:"批次采购成本（退货）"`
    BatchPurchaseApcFee string `json:"batchPurchaseApcFee" comment:"批次采购成本（赔偿）"`
    BatchPurchaseVrpcFee string `json:"batchPurchaseVrpcFee" comment:"批次采购成本（供应商退货）"`
    BatchPurchaseDfpcFee string `json:"batchPurchaseDfpcFee" comment:"批次采购成本（差异）"`
    SumBatchHeadTripFee string `json:"sumBatchHeadTripFee" comment:"批次头程费用合计"`
    BatchTripPcFee string `json:"batchTripPcFee" comment:"批次头程费用"`
    SumBatchHeadTripOtherFee string `json:"sumBatchHeadTripOtherFee" comment:"批次头程其他费用合计"`
    BatchTripCrpcFee string `json:"batchTripCrpcFee" comment:"批次头程费用（退货）"`
    BatchTripApcFee string `json:"batchTripApcFee" comment:"批次头程费用（赔偿）"`
    BatchTripVrpcFee string `json:"batchTripVrpcFee" comment:"批次头程费用（供应商退货）"`
    BatchTripDfpcFee string `json:"batchTripDfpcFee" comment:"批次头程费用（差异）"`
    SumCustomizeFee string `json:"sumCustomizeFee" comment:"自定义费用"`
    EvaluationFee string `json:"evaluationFee" comment:"测评费"`
    EvaluationCapital string `json:"evaluationCapital" comment:"测评本金"`
    EvaluationCommission string `json:"evaluationCommission" comment:"测评佣金"`
    ShopOtherFee string `json:"shopOtherFee" comment:"店铺其他费"`
    AsinOtherFee string `json:"asinOtherFee" comment:"ASIN其他费"`
    FixedFee string `json:"fixedFee" comment:"固定费用"`
    GrossProfit string `json:"grossProfit" comment:"毛利润"`
    GrossProfitRate string `json:"grossProfitRate" comment:"毛利率"`
    Roi string `json:"roi" comment:"ROI"`
    SaleRatioByAsin string `json:"saleRatioByAsin" comment:"ASIN销售率"`
    SaleRatioByShop string `json:"saleRatioByShop" comment:"店铺销售率"`
    AsinOtherFees string `json:"asinOtherFees" comment:"ASIN其他费"`
    ShopOtherFees string `json:"shopOtherFees" comment:"店铺其他费用"`
    SumSalesNum string `json:"sumSalesNum" comment:"销售数量"`
}

func (s *ErpFinanceAsinSummaryControl) Generate() dto.Control {
	o := *s
	return &o
}

func (s *ErpFinanceAsinSummaryControl) Bind(ctx *gin.Context) error {
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

func (s *ErpFinanceAsinSummaryControl) GenerateM() (common.ActiveRecord, error) {
	return &models.ErpFinanceAsinSummary{
        Model:     common.Model{ Id: s.Id },
        MainImg:  s.MainImg,
        CommoditySku:  s.CommoditySku,
        CommodityName:  s.CommodityName,
        ShopId:  s.ShopId,
        ShopName:  s.ShopName,
        MarketplaceId:  s.MarketplaceId,
        MarketplaceName:  s.MarketplaceName,
        Currency:  s.Currency,
        Asin:  s.Asin,
        AsinList:  s.AsinList,
        CommodityList:  s.CommodityList,
        SumOrderFee:  s.SumOrderFee,
        Domain:  s.Domain,
        ParentAsin:  s.ParentAsin,
        Sku:  s.Sku,
        Month:  s.Month,
        ProductSales:  s.ProductSales,
        FbaSales:  s.FbaSales,
        FbmSales:  s.FbmSales,
        CpcSales:  s.CpcSales,
        FbaSalesNum:  s.FbaSalesNum,
        FbmSalesNum:  s.FbmSalesNum,
        MultiChannelNum:  s.MultiChannelNum,
        FbaReissueNum:  s.FbaReissueNum,
        CpcSalesNum:  s.CpcSalesNum,
        CpcSdSalesNum:  s.CpcSdSalesNum,
        CpcSpSalesNum:  s.CpcSpSalesNum,
        FbaShippingCredits:  s.FbaShippingCredits,
        FbmShippingCredits:  s.FbmShippingCredits,
        PromotionalRebates:  s.PromotionalRebates,
        AdjCompensation:  s.AdjCompensation,
        AdjCompensationPcFee:  s.AdjCompensationPcFee,
        AdjCompensationNum:  s.AdjCompensationNum,
        Refund:  s.Refund,
        FbaRefund:  s.FbaRefund,
        FbmRefund:  s.FbmRefund,
        RefundNum:  s.RefundNum,
        FbaRefundNum:  s.FbaRefundNum,
        FbmRefundNum:  s.FbmRefundNum,
        RefundRate:  s.RefundRate,
        ReturnNum:  s.ReturnNum,
        ReturnSellableNum:  s.ReturnSellableNum,
        ReturnUnsellableNum:  s.ReturnUnsellableNum,
        ReturnRate:  s.ReturnRate,
        SaleOrderFee:  s.SaleOrderFee,
        SaleSellingFee:  s.SaleSellingFee,
        SaleFbaFee:  s.SaleFbaFee,
        SaleFbmFee:  s.SaleFbmFee,
        SaleOtherFee:  s.SaleOtherFee,
        MultiChannelOrderFee:  s.MultiChannelOrderFee,
        McFbaShipFee:  s.McFbaShipFee,
        RefundOrderFees:  s.RefundOrderFees,
        RefundSellingFee:  s.RefundSellingFee,
        RefundFbaFee:  s.RefundFbaFee,
        RefundOtherFee:  s.RefundOtherFee,
        SumCpcCost:  s.SumCpcCost,
        CpcCostDiff:  s.CpcCostDiff,
        CpcSpCost:  s.CpcSpCost,
        CpcSdCost:  s.CpcSdCost,
        SumPromoteFee:  s.SumPromoteFee,
        LdPromoteFee:  s.LdPromoteFee,
        CouponsPromoteFee:  s.CouponsPromoteFee,
        ErpPromoteFee:  s.ErpPromoteFee,
        SumFbaStorageFee:  s.SumFbaStorageFee,
        FbaStorageFee:  s.FbaStorageFee,
        FbaLongStorageFee:  s.FbaLongStorageFee,
        FbaStorageFeeDiff:  s.FbaStorageFeeDiff,
        SumStorageOtherFee:  s.SumStorageOtherFee,
        FbaDisposalNum:  s.FbaDisposalNum,
        FbaDisposalFee:  s.FbaDisposalFee,
        FbaRemovalNum:  s.FbaRemovalNum,
        FbaRemovalFee:  s.FbaRemovalFee,
        FbaReturnFee:  s.FbaReturnFee,
        FbaShipFee:  s.FbaShipFee,
        FbaStorageOtherFee:  s.FbaStorageOtherFee,
        FbaInventoryPlacementFee:  s.FbaInventoryPlacementFee,
        SumFbaAdjustmentFee:  s.SumFbaAdjustmentFee,
        FbaAdjustmentFee:  s.FbaAdjustmentFee,
        FbaAdjustmentDiffFee:  s.FbaAdjustmentDiffFee,
        SumOtherFee:  s.SumOtherFee,
        SubscriptionFee:  s.SubscriptionFee,
        OtherFee:  s.OtherFee,
        Tax:  s.Tax,
        SumPurchaseFee:  s.SumPurchaseFee,
        PurchasePcFee:  s.PurchasePcFee,
        PurchaseDpcFee:  s.PurchaseDpcFee,
        PurchaseRtpcFee:  s.PurchaseRtpcFee,
        PurchaseAcpcFee:  s.PurchaseAcpcFee,
        PurchaseMcpcFee:  s.PurchaseMcpcFee,
        PurchaseRpcFee:  s.PurchaseRpcFee,
        SumHeadTripFee:  s.SumHeadTripFee,
        HeadTripPcFee:  s.HeadTripPcFee,
        HeadTripDpcFee:  s.HeadTripDpcFee,
        HeadTripRtpcFee:  s.HeadTripRtpcFee,
        HeadTripAcpcFee:  s.HeadTripAcpcFee,
        HeadTripMcpcFee:  s.HeadTripMcpcFee,
        HeadTripRpcFee:  s.HeadTripRpcFee,
        SumBatchPurchaseFee:  s.SumBatchPurchaseFee,
        BatchPurchasePcFee:  s.BatchPurchasePcFee,
        SumBatchPurchaseOtherFee:  s.SumBatchPurchaseOtherFee,
        BatchPurchaseCrpcFee:  s.BatchPurchaseCrpcFee,
        BatchPurchaseApcFee:  s.BatchPurchaseApcFee,
        BatchPurchaseVrpcFee:  s.BatchPurchaseVrpcFee,
        BatchPurchaseDfpcFee:  s.BatchPurchaseDfpcFee,
        SumBatchHeadTripFee:  s.SumBatchHeadTripFee,
        BatchTripPcFee:  s.BatchTripPcFee,
        SumBatchHeadTripOtherFee:  s.SumBatchHeadTripOtherFee,
        BatchTripCrpcFee:  s.BatchTripCrpcFee,
        BatchTripApcFee:  s.BatchTripApcFee,
        BatchTripVrpcFee:  s.BatchTripVrpcFee,
        BatchTripDfpcFee:  s.BatchTripDfpcFee,
        SumCustomizeFee:  s.SumCustomizeFee,
        EvaluationFee:  s.EvaluationFee,
        EvaluationCapital:  s.EvaluationCapital,
        EvaluationCommission:  s.EvaluationCommission,
        ShopOtherFee:  s.ShopOtherFee,
        AsinOtherFee:  s.AsinOtherFee,
        FixedFee:  s.FixedFee,
        GrossProfit:  s.GrossProfit,
        GrossProfitRate:  s.GrossProfitRate,
        Roi:  s.Roi,
        SaleRatioByAsin:  s.SaleRatioByAsin,
        SaleRatioByShop:  s.SaleRatioByShop,
        AsinOtherFees:  s.AsinOtherFees,
        ShopOtherFees:  s.ShopOtherFees,
        SumSalesNum:  s.SumSalesNum,
	}, nil
}

func (s *ErpFinanceAsinSummaryControl) GetId() interface{} {
	return s.Id
}

type ErpFinanceAsinSummaryById struct {
	dto.ObjectById
}

func (s *ErpFinanceAsinSummaryById) GetId() interface{} {
	if len(s.Ids) > 0 {
		s.Ids = append(s.Ids, s.Id)
		return s.Ids
	}
	return s.Id
}

func (s *ErpFinanceAsinSummaryById) Bind(ctx *gin.Context) error {
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

func (s *ErpFinanceAsinSummaryById) SetUpdateBy(id int) {

}

func (s *ErpFinanceAsinSummaryById) Generate() dto.Control {
	o := *s
	return &o
}

func (s *ErpFinanceAsinSummaryById) GenerateM() (common.ActiveRecord, error) {
	return &models.ErpFinanceAsinSummary{}, nil
}
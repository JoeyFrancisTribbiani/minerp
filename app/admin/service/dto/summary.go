package dto

import (
    "errors"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type ErpFinanceSummarySearch struct {
	dto.Pagination     `search:"-"`
}

func (m *ErpFinanceSummarySearch) GetNeedSearch() interface{} {
	return *m
}

func (m *ErpFinanceSummarySearch) Generate() dto.Index {
	o := *m
	return &o
}

func (m *ErpFinanceSummarySearch) Bind(ctx *gin.Context) error {
	log := api.GetRequestLogger(ctx)
    err := ctx.ShouldBind(m)
    if err != nil {
    	log.Errorf("ShouldBind error: %s", err.Error())
    }
    return err
}

type ErpFinanceSummaryControl struct {
    Id int `uri:"id" comment:"id"` // id
    Month string `json:"month" comment:"null"`
    ShopId string `json:"shopId" comment:"店铺id"`
    ShopName string `json:"shopName" comment:"店铺名称"`
    MarketplaceId string `json:"marketplaceId" comment:"市场id"`
    Marketplace string `json:"marketplace" comment:"市场"`
    SumSalesNum string `json:"sumSalesNum" comment:"销售数量"`
    FbaSalesNum string `json:"fbaSalesNum" comment:"亚马逊FBA销量"`
    FbmSalesNum string `json:"fbmSalesNum" comment:"亚马逊FBM销量"`
    FbaReissueNum string `json:"fbaReissueNum" comment:"亚马逊补发销量"`
    MultiChannelNum string `json:"multiChannelNum" comment:"多渠道销量"`
    RefundNum string `json:"refundNum" comment:"退款量"`
    ReturnNum string `json:"returnNum" comment:"退货量"`
    CompensationNum string `json:"compensationNum" comment:"赔偿量"`
    DisposedNum string `json:"disposedNum" comment:"销毁量"`
    RemovalNum string `json:"removalNum" comment:"移除量"`
    SumOrderSales string `json:"sumOrderSales" comment:"订单收入"`
    ProductSales string `json:"productSales" comment:"销售额"`
    ShippingCredits string `json:"shippingCredits" comment:"买家运费"`
    PromotionalRebates string `json:"promotionalRebates" comment:"促销折扣"`
    Refund string `json:"refund" comment:"退款"`
    Compensation string `json:"compensation" comment:"赔偿收入"`
    AdjCompensation string `json:"adjCompensation" comment:"库存调整赔偿"`
    SumOrderFee string `json:"sumOrderFee" comment:"订单费用"`
    SaleOrderFee string `json:"saleOrderFee" comment:"销售订单"`
    SaleSellingFee string `json:"saleSellingFee" comment:"平台费"`
    SaleFbaFee string `json:"saleFbaFee" comment:"FBA发货费"`
    SaleFbmFee string `json:"saleFbmFee" comment:"FBM发货费"`
    SaleOtherFee string `json:"saleOtherFee" comment:"其他费"`
    MultiChannelOrderFee string `json:"multiChannelOrderFee" comment:"多渠道订单"`
    McFbaShipFee string `json:"mcFbaShipFee" comment:"FBA发货费"`
    RefundOrderFees string `json:"refundOrderFees" comment:"退款订单"`
    RefundSellingFee string `json:"refundSellingFee" comment:"平台费"`
    RefundFbaFee string `json:"refundFbaFee" comment:"FBA发货费"`
    RefundOtherFee string `json:"refundOtherFee" comment:"其他费"`
    SumFbaStorageFee string `json:"sumFbaStorageFee" comment:"仓储费用"`
    SumShopFbaStorageFee string `json:"sumShopFbaStorageFee" comment:"店铺仓储费用"`
    FbaStorageFee string `json:"fbaStorageFee" comment:"FBA月仓储费"`
    FbaStorageFeeDiff string `json:"fbaStorageFeeDiff" comment:"FBA月仓储费差"`
    FbaLongStorageFee string `json:"fbaLongStorageFee" comment:"0.00"`
    SumShopFbaStorageOtherFee string `json:"sumShopFbaStorageOtherFee" comment:"店铺其他仓储费用汇总"`
    FbaAdjustmentFee string `json:"fbaAdjustmentFee" comment:"库存调整费"`
    FbaDisposalFee string `json:"fbaDisposalFee" comment:"FBA销毁费"`
    FbaRemovalFee string `json:"fbaRemovalFee" comment:"FBA移除费"`
    FbaReturnFee string `json:"fbaReturnFee" comment:"退货入仓费"`
    FbaShipFee string `json:"fbaShipFee" comment:"亚马逊合作承运费"`
    FbaStorageOtherFee string `json:"fbaStorageOtherFee" comment:"其他费"`
    FbaInventoryPlacementFee string `json:"fbaInventoryPlacementFee" comment:"库存配置费"`
    SumCpcCost string `json:"sumCpcCost" comment:"广告费"`
    CpcCost string `json:"cpcCost" comment:"CPC广告费"`
    CpcCostDiff string `json:"cpcCostDiff" comment:"CPC广告费差异"`
    SumPromoteFee string `json:"sumPromoteFee" comment:"推广费用"`
    LdPromoteFee string `json:"ldPromoteFee" comment:"LD费"`
    CouponsPromoteFee string `json:"couponsPromoteFee" comment:"优惠劵"`
    ErpPromoteFee string `json:"erpPromoteFee" comment:"早期评论人计划"`
    SumTax string `json:"sumTax" comment:"税费"`
    Tax string `json:"tax" comment:"税费"`
    SumProductCostFee string `json:"sumProductCostFee" comment:"商品成本"`
    SumPurchaseFee string `json:"sumPurchaseFee" comment:"总采购成本"`
    SumHeadTripFee string `json:"sumHeadTripFee" comment:"总头程费用"`
    SalePcFee string `json:"salePcFee" comment:"售出商品"`
    PurchasePcFee string `json:"purchasePcFee" comment:"采购成本"`
    HeadTripPcFee string `json:"headTripPcFee" comment:"头程费用"`
    DisposedPcFee string `json:"disposedPcFee" comment:"销毁商品"`
    PurchaseDpcFee string `json:"purchaseDpcFee" comment:"采购成本（销毁）"`
    HeadTripDpcFee string `json:"headTripDpcFee" comment:"头程费用（销毁）"`
    ReturnPcFee string `json:"returnPcFee" comment:"退货成本"`
    PurchaseRtpcFee string `json:"purchaseRtpcFee" comment:"采购成本（退货）"`
    HeadTripRtpcFee string `json:"headTripRtpcFee" comment:"头程费用（退货）"`
    AdjCompensationPcFee string `json:"adjCompensationPcFee" comment:"FBA库存赔偿"`
    PurchaseAcpcFee string `json:"purchaseAcpcFee" comment:"采购成本（赔偿）"`
    HeadTripAcpcFee string `json:"headTripAcpcFee" comment:"头程费用（赔偿）"`
    MultiChannelPcFee string `json:"multiChannelPcFee" comment:"多渠道订单商品"`
    PurchaseMcpcFee string `json:"purchaseMcpcFee" comment:"采购成本"`
    HeadTripMcpcFee string `json:"headTripMcpcFee" comment:"头程费用"`
    ReissuePcFee string `json:"reissuePcFee" comment:"补发成本"`
    PurchaseRpcFee string `json:"purchaseRpcFee" comment:"采购成本"`
    HeadTripRpcFee string `json:"headTripRpcFee" comment:"头程费用"`
    SumBatchProductCostFee string `json:"sumBatchProductCostFee" comment:"批次商品成本合计"`
    SumBatchPurchaseFee string `json:"sumBatchPurchaseFee" comment:"批次采购成本合计"`
    SumBatchHeadTripFee string `json:"sumBatchHeadTripFee" comment:"批次头程费用合计"`
    SumBatchPurchaseOtherFee string `json:"sumBatchPurchaseOtherFee" comment:"批次采购其他成本合计"`
    SumBatchHeadTripOtherFee string `json:"sumBatchHeadTripOtherFee" comment:"批次头程其他费用合计"`
    BatchSalePcFee string `json:"batchSalePcFee" comment:"批次售出商品"`
    BatchPurchasePcFee string `json:"batchPurchasePcFee" comment:"批次采购成本"`
    BatchTripPcFee string `json:"batchTripPcFee" comment:"批次头程费用"`
    BatchCustomerReturnsPcFee string `json:"batchCustomerReturnsPcFee" comment:"批次退货成本"`
    BatchPurchaseCrpcFee string `json:"batchPurchaseCrpcFee" comment:"批次采购成本（退货）"`
    BatchTripCrpcFee string `json:"batchTripCrpcFee" comment:"批次头程费用（退货）"`
    BatchAdjustmentsPcFee string `json:"batchAdjustmentsPcFee" comment:"批次FBA库存赔偿"`
    BatchPurchaseApcFee string `json:"batchPurchaseApcFee" comment:"批次采购成本（赔偿）"`
    BatchTripApcFee string `json:"batchTripApcFee" comment:"批次头程费用（赔偿）"`
    BatchVendorReturnsPcFee string `json:"batchVendorReturnsPcFee" comment:"批次供应商退货"`
    BatchPurchaseVrpcFee string `json:"batchPurchaseVrpcFee" comment:"批次采购成本（供应商退货）"`
    BatchTripVrpcFee string `json:"batchTripVrpcFee" comment:"批次头程费用（供应商退货）"`
    BatchDifferencePcFee string `json:"batchDifferencePcFee" comment:"批次差异"`
    BatchPurchaseDfpcFee string `json:"batchPurchaseDfpcFee" comment:"批次采购成本（差异）"`
    BatchTripDfpcFee string `json:"batchTripDfpcFee" comment:"批次头程费用（差异）"`
    SumOtherFee string `json:"sumOtherFee" comment:"平台其他费用"`
    SubscriptionFee string `json:"subscriptionFee" comment:"订阅费"`
    OtherFee string `json:"otherFee" comment:"其他费"`
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
    OriCurrency string `json:"oriCurrency" comment:"ori_currency"`
    Currency string `json:"currency" comment:"币种"`
    ShopOtherFees string `json:"shopOtherFees" comment:"店铺其他费用"`
    BadDataItems string `json:"badDataItems" comment:"坏数据项"`
}

func (s *ErpFinanceSummaryControl) Generate() dto.Control {
	o := *s
	return &o
}

func (s *ErpFinanceSummaryControl) Bind(ctx *gin.Context) error {
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

func (s *ErpFinanceSummaryControl) GenerateM() (common.ActiveRecord, error) {
	return &models.ErpFinanceSummary{
        Model:     common.Model{ Id: s.Id },
        Month:  s.Month,
        ShopId:  s.ShopId,
        ShopName:  s.ShopName,
        MarketplaceId:  s.MarketplaceId,
        Marketplace:  s.Marketplace,
        SumSalesNum:  s.SumSalesNum,
        FbaSalesNum:  s.FbaSalesNum,
        FbmSalesNum:  s.FbmSalesNum,
        FbaReissueNum:  s.FbaReissueNum,
        MultiChannelNum:  s.MultiChannelNum,
        RefundNum:  s.RefundNum,
        ReturnNum:  s.ReturnNum,
        CompensationNum:  s.CompensationNum,
        DisposedNum:  s.DisposedNum,
        RemovalNum:  s.RemovalNum,
        SumOrderSales:  s.SumOrderSales,
        ProductSales:  s.ProductSales,
        ShippingCredits:  s.ShippingCredits,
        PromotionalRebates:  s.PromotionalRebates,
        Refund:  s.Refund,
        Compensation:  s.Compensation,
        AdjCompensation:  s.AdjCompensation,
        SumOrderFee:  s.SumOrderFee,
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
        SumFbaStorageFee:  s.SumFbaStorageFee,
        SumShopFbaStorageFee:  s.SumShopFbaStorageFee,
        FbaStorageFee:  s.FbaStorageFee,
        FbaStorageFeeDiff:  s.FbaStorageFeeDiff,
        FbaLongStorageFee:  s.FbaLongStorageFee,
        SumShopFbaStorageOtherFee:  s.SumShopFbaStorageOtherFee,
        FbaAdjustmentFee:  s.FbaAdjustmentFee,
        FbaDisposalFee:  s.FbaDisposalFee,
        FbaRemovalFee:  s.FbaRemovalFee,
        FbaReturnFee:  s.FbaReturnFee,
        FbaShipFee:  s.FbaShipFee,
        FbaStorageOtherFee:  s.FbaStorageOtherFee,
        FbaInventoryPlacementFee:  s.FbaInventoryPlacementFee,
        SumCpcCost:  s.SumCpcCost,
        CpcCost:  s.CpcCost,
        CpcCostDiff:  s.CpcCostDiff,
        SumPromoteFee:  s.SumPromoteFee,
        LdPromoteFee:  s.LdPromoteFee,
        CouponsPromoteFee:  s.CouponsPromoteFee,
        ErpPromoteFee:  s.ErpPromoteFee,
        SumTax:  s.SumTax,
        Tax:  s.Tax,
        SumProductCostFee:  s.SumProductCostFee,
        SumPurchaseFee:  s.SumPurchaseFee,
        SumHeadTripFee:  s.SumHeadTripFee,
        SalePcFee:  s.SalePcFee,
        PurchasePcFee:  s.PurchasePcFee,
        HeadTripPcFee:  s.HeadTripPcFee,
        DisposedPcFee:  s.DisposedPcFee,
        PurchaseDpcFee:  s.PurchaseDpcFee,
        HeadTripDpcFee:  s.HeadTripDpcFee,
        ReturnPcFee:  s.ReturnPcFee,
        PurchaseRtpcFee:  s.PurchaseRtpcFee,
        HeadTripRtpcFee:  s.HeadTripRtpcFee,
        AdjCompensationPcFee:  s.AdjCompensationPcFee,
        PurchaseAcpcFee:  s.PurchaseAcpcFee,
        HeadTripAcpcFee:  s.HeadTripAcpcFee,
        MultiChannelPcFee:  s.MultiChannelPcFee,
        PurchaseMcpcFee:  s.PurchaseMcpcFee,
        HeadTripMcpcFee:  s.HeadTripMcpcFee,
        ReissuePcFee:  s.ReissuePcFee,
        PurchaseRpcFee:  s.PurchaseRpcFee,
        HeadTripRpcFee:  s.HeadTripRpcFee,
        SumBatchProductCostFee:  s.SumBatchProductCostFee,
        SumBatchPurchaseFee:  s.SumBatchPurchaseFee,
        SumBatchHeadTripFee:  s.SumBatchHeadTripFee,
        SumBatchPurchaseOtherFee:  s.SumBatchPurchaseOtherFee,
        SumBatchHeadTripOtherFee:  s.SumBatchHeadTripOtherFee,
        BatchSalePcFee:  s.BatchSalePcFee,
        BatchPurchasePcFee:  s.BatchPurchasePcFee,
        BatchTripPcFee:  s.BatchTripPcFee,
        BatchCustomerReturnsPcFee:  s.BatchCustomerReturnsPcFee,
        BatchPurchaseCrpcFee:  s.BatchPurchaseCrpcFee,
        BatchTripCrpcFee:  s.BatchTripCrpcFee,
        BatchAdjustmentsPcFee:  s.BatchAdjustmentsPcFee,
        BatchPurchaseApcFee:  s.BatchPurchaseApcFee,
        BatchTripApcFee:  s.BatchTripApcFee,
        BatchVendorReturnsPcFee:  s.BatchVendorReturnsPcFee,
        BatchPurchaseVrpcFee:  s.BatchPurchaseVrpcFee,
        BatchTripVrpcFee:  s.BatchTripVrpcFee,
        BatchDifferencePcFee:  s.BatchDifferencePcFee,
        BatchPurchaseDfpcFee:  s.BatchPurchaseDfpcFee,
        BatchTripDfpcFee:  s.BatchTripDfpcFee,
        SumOtherFee:  s.SumOtherFee,
        SubscriptionFee:  s.SubscriptionFee,
        OtherFee:  s.OtherFee,
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
        OriCurrency:  s.OriCurrency,
        Currency:  s.Currency,
        ShopOtherFees:  s.ShopOtherFees,
        BadDataItems:  s.BadDataItems,
	}, nil
}

func (s *ErpFinanceSummaryControl) GetId() interface{} {
	return s.Id
}

type ErpFinanceSummaryById struct {
	dto.ObjectById
}

func (s *ErpFinanceSummaryById) GetId() interface{} {
	if len(s.Ids) > 0 {
		s.Ids = append(s.Ids, s.Id)
		return s.Ids
	}
	return s.Id
}

func (s *ErpFinanceSummaryById) Bind(ctx *gin.Context) error {
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

func (s *ErpFinanceSummaryById) SetUpdateBy(id int) {

}

func (s *ErpFinanceSummaryById) Generate() dto.Control {
	o := *s
	return &o
}

func (s *ErpFinanceSummaryById) GenerateM() (common.ActiveRecord, error) {
	return &models.ErpFinanceSummary{}, nil
}
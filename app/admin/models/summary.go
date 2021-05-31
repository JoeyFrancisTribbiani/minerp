package models

import (
   // "gorm.io/gorm"

	"go-admin/common/models"
)

type ErpFinanceSummary struct {
    models.Model
    
    Month string `json:"month" gorm:"type:varchar(100);comment:null"` 
    ShopId string `json:"shopId" gorm:"type:varchar(100);comment:店铺id"` 
    ShopName string `json:"shopName" gorm:"type:varchar(100);comment:店铺名称"` 
    MarketplaceId string `json:"marketplaceId" gorm:"type:varchar(100);comment:市场id"` 
    Marketplace string `json:"marketplace" gorm:"type:varchar(100);comment:市场"` 
    SumSalesNum string `json:"sumSalesNum" gorm:"type:bigint;comment:销售数量"` 
    FbaSalesNum string `json:"fbaSalesNum" gorm:"type:bigint;comment:亚马逊FBA销量"` 
    FbmSalesNum string `json:"fbmSalesNum" gorm:"type:bigint;comment:亚马逊FBM销量"` 
    FbaReissueNum string `json:"fbaReissueNum" gorm:"type:bigint;comment:亚马逊补发销量"` 
    MultiChannelNum string `json:"multiChannelNum" gorm:"type:bigint;comment:多渠道销量"` 
    RefundNum string `json:"refundNum" gorm:"type:bigint;comment:退款量"` 
    ReturnNum string `json:"returnNum" gorm:"type:bigint;comment:退货量"` 
    CompensationNum string `json:"compensationNum" gorm:"type:bigint;comment:赔偿量"` 
    DisposedNum string `json:"disposedNum" gorm:"type:bigint;comment:销毁量"` 
    RemovalNum string `json:"removalNum" gorm:"type:bigint;comment:移除量"` 
    SumOrderSales string `json:"sumOrderSales" gorm:"type:varchar(100);comment:订单收入"` 
    ProductSales string `json:"productSales" gorm:"type:varchar(100);comment:销售额"` 
    ShippingCredits string `json:"shippingCredits" gorm:"type:varchar(100);comment:买家运费"` 
    PromotionalRebates string `json:"promotionalRebates" gorm:"type:varchar(100);comment:促销折扣"` 
    Refund string `json:"refund" gorm:"type:varchar(100);comment:退款"` 
    Compensation string `json:"compensation" gorm:"type:varchar(100);comment:赔偿收入"` 
    AdjCompensation string `json:"adjCompensation" gorm:"type:varchar(100);comment:库存调整赔偿"` 
    SumOrderFee string `json:"sumOrderFee" gorm:"type:varchar(100);comment:订单费用"` 
    SaleOrderFee string `json:"saleOrderFee" gorm:"type:varchar(100);comment:销售订单"` 
    SaleSellingFee string `json:"saleSellingFee" gorm:"type:varchar(100);comment:平台费"` 
    SaleFbaFee string `json:"saleFbaFee" gorm:"type:varchar(100);comment:FBA发货费"` 
    SaleFbmFee string `json:"saleFbmFee" gorm:"type:varchar(100);comment:FBM发货费"` 
    SaleOtherFee string `json:"saleOtherFee" gorm:"type:varchar(100);comment:其他费"` 
    MultiChannelOrderFee string `json:"multiChannelOrderFee" gorm:"type:varchar(100);comment:多渠道订单"` 
    McFbaShipFee string `json:"mcFbaShipFee" gorm:"type:varchar(100);comment:FBA发货费"` 
    RefundOrderFees string `json:"refundOrderFees" gorm:"type:varchar(100);comment:退款订单"` 
    RefundSellingFee string `json:"refundSellingFee" gorm:"type:varchar(100);comment:平台费"` 
    RefundFbaFee string `json:"refundFbaFee" gorm:"type:varchar(100);comment:FBA发货费"` 
    RefundOtherFee string `json:"refundOtherFee" gorm:"type:varchar(99);comment:其他费"` 
    SumFbaStorageFee string `json:"sumFbaStorageFee" gorm:"type:varchar(100);comment:仓储费用"` 
    SumShopFbaStorageFee string `json:"sumShopFbaStorageFee" gorm:"type:varchar(100);comment:店铺仓储费用"` 
    FbaStorageFee string `json:"fbaStorageFee" gorm:"type:varchar(100);comment:FBA月仓储费"` 
    FbaStorageFeeDiff string `json:"fbaStorageFeeDiff" gorm:"type:varchar(100);comment:FBA月仓储费差"` 
    FbaLongStorageFee string `json:"fbaLongStorageFee" gorm:"type:varchar(100);comment:0.00"` 
    SumShopFbaStorageOtherFee string `json:"sumShopFbaStorageOtherFee" gorm:"type:varchar(100);comment:店铺其他仓储费用汇总"` 
    FbaAdjustmentFee string `json:"fbaAdjustmentFee" gorm:"type:varchar(100);comment:库存调整费"` 
    FbaDisposalFee string `json:"fbaDisposalFee" gorm:"type:varchar(100);comment:FBA销毁费"` 
    FbaRemovalFee string `json:"fbaRemovalFee" gorm:"type:varchar(100);comment:FBA移除费"` 
    FbaReturnFee string `json:"fbaReturnFee" gorm:"type:varchar(100);comment:退货入仓费"` 
    FbaShipFee string `json:"fbaShipFee" gorm:"type:varchar(100);comment:亚马逊合作承运费"` 
    FbaStorageOtherFee string `json:"fbaStorageOtherFee" gorm:"type:varchar(100);comment:其他费"` 
    FbaInventoryPlacementFee string `json:"fbaInventoryPlacementFee" gorm:"type:varchar(100);comment:库存配置费"` 
    SumCpcCost string `json:"sumCpcCost" gorm:"type:varchar(100);comment:广告费"` 
    CpcCost string `json:"cpcCost" gorm:"type:varchar(100);comment:CPC广告费"` 
    CpcCostDiff string `json:"cpcCostDiff" gorm:"type:varchar(100);comment:CPC广告费差异"` 
    SumPromoteFee string `json:"sumPromoteFee" gorm:"type:varchar(100);comment:推广费用"` 
    LdPromoteFee string `json:"ldPromoteFee" gorm:"type:varchar(100);comment:LD费"` 
    CouponsPromoteFee string `json:"couponsPromoteFee" gorm:"type:varchar(100);comment:优惠劵"` 
    ErpPromoteFee string `json:"erpPromoteFee" gorm:"type:varchar(100);comment:早期评论人计划"` 
    SumTax string `json:"sumTax" gorm:"type:varchar(100);comment:税费"` 
    Tax string `json:"tax" gorm:"type:varchar(100);comment:税费"` 
    SumProductCostFee string `json:"sumProductCostFee" gorm:"type:varchar(100);comment:商品成本"` 
    SumPurchaseFee string `json:"sumPurchaseFee" gorm:"type:varchar(100);comment:总采购成本"` 
    SumHeadTripFee string `json:"sumHeadTripFee" gorm:"type:varchar(100);comment:总头程费用"` 
    SalePcFee string `json:"salePcFee" gorm:"type:varchar(100);comment:售出商品"` 
    PurchasePcFee string `json:"purchasePcFee" gorm:"type:varchar(100);comment:采购成本"` 
    HeadTripPcFee string `json:"headTripPcFee" gorm:"type:varchar(100);comment:头程费用"` 
    DisposedPcFee string `json:"disposedPcFee" gorm:"type:varchar(100);comment:销毁商品"` 
    PurchaseDpcFee string `json:"purchaseDpcFee" gorm:"type:varchar(100);comment:采购成本（销毁）"` 
    HeadTripDpcFee string `json:"headTripDpcFee" gorm:"type:varchar(100);comment:头程费用（销毁）"` 
    ReturnPcFee string `json:"returnPcFee" gorm:"type:varchar(100);comment:退货成本"` 
    PurchaseRtpcFee string `json:"purchaseRtpcFee" gorm:"type:varchar(100);comment:采购成本（退货）"` 
    HeadTripRtpcFee string `json:"headTripRtpcFee" gorm:"type:varchar(100);comment:头程费用（退货）"` 
    AdjCompensationPcFee string `json:"adjCompensationPcFee" gorm:"type:varchar(100);comment:FBA库存赔偿"` 
    PurchaseAcpcFee string `json:"purchaseAcpcFee" gorm:"type:varchar(100);comment:采购成本（赔偿）"` 
    HeadTripAcpcFee string `json:"headTripAcpcFee" gorm:"type:varchar(100);comment:头程费用（赔偿）"` 
    MultiChannelPcFee string `json:"multiChannelPcFee" gorm:"type:varchar(100);comment:多渠道订单商品"` 
    PurchaseMcpcFee string `json:"purchaseMcpcFee" gorm:"type:varchar(100);comment:采购成本"` 
    HeadTripMcpcFee string `json:"headTripMcpcFee" gorm:"type:varchar(100);comment:头程费用"` 
    ReissuePcFee string `json:"reissuePcFee" gorm:"type:varchar(100);comment:补发成本"` 
    PurchaseRpcFee string `json:"purchaseRpcFee" gorm:"type:varchar(100);comment:采购成本"` 
    HeadTripRpcFee string `json:"headTripRpcFee" gorm:"type:varchar(100);comment:头程费用"` 
    SumBatchProductCostFee string `json:"sumBatchProductCostFee" gorm:"type:varchar(100);comment:批次商品成本合计"` 
    SumBatchPurchaseFee string `json:"sumBatchPurchaseFee" gorm:"type:varchar(100);comment:批次采购成本合计"` 
    SumBatchHeadTripFee string `json:"sumBatchHeadTripFee" gorm:"type:varchar(100);comment:批次头程费用合计"` 
    SumBatchPurchaseOtherFee string `json:"sumBatchPurchaseOtherFee" gorm:"type:varchar(100);comment:批次采购其他成本合计"` 
    SumBatchHeadTripOtherFee string `json:"sumBatchHeadTripOtherFee" gorm:"type:varchar(100);comment:批次头程其他费用合计"` 
    BatchSalePcFee string `json:"batchSalePcFee" gorm:"type:varchar(100);comment:批次售出商品"` 
    BatchPurchasePcFee string `json:"batchPurchasePcFee" gorm:"type:varchar(100);comment:批次采购成本"` 
    BatchTripPcFee string `json:"batchTripPcFee" gorm:"type:varchar(100);comment:批次头程费用"` 
    BatchCustomerReturnsPcFee string `json:"batchCustomerReturnsPcFee" gorm:"type:varchar(100);comment:批次退货成本"` 
    BatchPurchaseCrpcFee string `json:"batchPurchaseCrpcFee" gorm:"type:varchar(100);comment:批次采购成本（退货）"` 
    BatchTripCrpcFee string `json:"batchTripCrpcFee" gorm:"type:varchar(100);comment:批次头程费用（退货）"` 
    BatchAdjustmentsPcFee string `json:"batchAdjustmentsPcFee" gorm:"type:varchar(100);comment:批次FBA库存赔偿"` 
    BatchPurchaseApcFee string `json:"batchPurchaseApcFee" gorm:"type:varchar(100);comment:批次采购成本（赔偿）"` 
    BatchTripApcFee string `json:"batchTripApcFee" gorm:"type:varchar(100);comment:批次头程费用（赔偿）"` 
    BatchVendorReturnsPcFee string `json:"batchVendorReturnsPcFee" gorm:"type:varchar(100);comment:批次供应商退货"` 
    BatchPurchaseVrpcFee string `json:"batchPurchaseVrpcFee" gorm:"type:varchar(100);comment:批次采购成本（供应商退货）"` 
    BatchTripVrpcFee string `json:"batchTripVrpcFee" gorm:"type:varchar(100);comment:批次头程费用（供应商退货）"` 
    BatchDifferencePcFee string `json:"batchDifferencePcFee" gorm:"type:varchar(100);comment:批次差异"` 
    BatchPurchaseDfpcFee string `json:"batchPurchaseDfpcFee" gorm:"type:varchar(100);comment:批次采购成本（差异）"` 
    BatchTripDfpcFee string `json:"batchTripDfpcFee" gorm:"type:varchar(100);comment:批次头程费用（差异）"` 
    SumOtherFee string `json:"sumOtherFee" gorm:"type:varchar(100);comment:平台其他费用"` 
    SubscriptionFee string `json:"subscriptionFee" gorm:"type:varchar(100);comment:订阅费"` 
    OtherFee string `json:"otherFee" gorm:"type:varchar(100);comment:其他费"` 
    SumCustomizeFee string `json:"sumCustomizeFee" gorm:"type:varchar(100);comment:自定义费用"` 
    EvaluationFee string `json:"evaluationFee" gorm:"type:varchar(100);comment:测评费"` 
    EvaluationCapital string `json:"evaluationCapital" gorm:"type:varchar(100);comment:测评本金"` 
    EvaluationCommission string `json:"evaluationCommission" gorm:"type:varchar(100);comment:测评佣金"` 
    ShopOtherFee string `json:"shopOtherFee" gorm:"type:varchar(100);comment:店铺其他费"` 
    AsinOtherFee string `json:"asinOtherFee" gorm:"type:varchar(100);comment:ASIN其他费"` 
    FixedFee string `json:"fixedFee" gorm:"type:varchar(100);comment:固定费用"` 
    GrossProfit string `json:"grossProfit" gorm:"type:varchar(100);comment:毛利润"` 
    GrossProfitRate string `json:"grossProfitRate" gorm:"type:varchar(100);comment:毛利率"` 
    Roi string `json:"roi" gorm:"type:varchar(100);comment:ROI"` 
    OriCurrency string `json:"oriCurrency" gorm:"type:varchar(100);comment:ori_currency"` 
    Currency string `json:"currency" gorm:"type:varchar(100);comment:币种"` 
    ShopOtherFees string `json:"shopOtherFees" gorm:"type:varchar(100);comment:店铺其他费用"` 
    BadDataItems string `json:"badDataItems" gorm:"type:varchar(100);comment:坏数据项"` 
    models.ModelTime
    models.ControlBy
}

func (ErpFinanceSummary) TableName() string {
    return "erp_finance_summary"
}

func (e *ErpFinanceSummary) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *ErpFinanceSummary) GetId() interface{} {
	return e.Id
}
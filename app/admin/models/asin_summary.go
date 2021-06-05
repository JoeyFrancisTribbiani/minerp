package models

import (
   // "gorm.io/gorm"

	"go-admin/common/models"
)

type ErpFinanceAsinSummary struct {
    models.Model
    
    MainImg string `json:"mainImg" gorm:"type:varchar(100);comment:主图"` 
    CommoditySku string `json:"commoditySku" gorm:"type:varchar(100);comment:商品SKU"` 
    CommodityName string `json:"commodityName" gorm:"type:varchar(100);comment:商品名称"` 
    ShopId string `json:"shopId" gorm:"type:varchar(100);comment:店铺id"` 
    ShopName string `json:"shopName" gorm:"type:varchar(100);comment:店铺名称"` 
    MarketplaceId string `json:"marketplaceId" gorm:"type:varchar(100);comment:市场id"` 
    MarketplaceName string `json:"marketplaceName" gorm:"type:varchar(100);comment:市场名称"` 
    Currency string `json:"currency" gorm:"type:varchar(100);comment:币种"` 
    Asin string `json:"asin" gorm:"type:varchar(100);comment:ASIN"` 
    AsinList string `json:"asinList" gorm:"type:varchar(100);comment:ASIN列表"` 
    CommodityList string `json:"commodityList" gorm:"type:varchar(100);comment:商品列表"` 
    SumOrderFee string `json:"sumOrderFee" gorm:"type:varchar(100);comment:订单费用"` 
    Domain string `json:"domain" gorm:"type:varchar(100);comment:区域域名"` 
    ParentAsin string `json:"parentAsin" gorm:"type:varchar(100);comment:父ASIN"` 
    Sku string `json:"sku" gorm:"type:varchar(100);comment:SKU"` 
    Month string `json:"month" gorm:"type:varchar(100);comment:月份"` 
    ProductSales string `json:"productSales" gorm:"type:varchar(100);comment:销售额"` 
    FbaSales string `json:"fbaSales" gorm:"type:varchar(100);comment:FBA销售额"` 
    FbmSales string `json:"fbmSales" gorm:"type:varchar(100);comment:FBM销售额"` 
    CpcSales string `json:"cpcSales" gorm:"type:varchar(100);comment:CPC销售额"` 
    FbaSalesNum string `json:"fbaSalesNum" gorm:"type:bigint;comment:亚马逊FBA销量"` 
    FbmSalesNum string `json:"fbmSalesNum" gorm:"type:bigint;comment:亚马逊FBM销量"` 
    MultiChannelNum string `json:"multiChannelNum" gorm:"type:bigint;comment:多渠道销量"` 
    FbaReissueNum string `json:"fbaReissueNum" gorm:"type:bigint;comment:FBA补发量"` 
    CpcSalesNum string `json:"cpcSalesNum" gorm:"type:bigint;comment:CPC销量"` 
    CpcSdSalesNum string `json:"cpcSdSalesNum" gorm:"type:bigint;comment:CPC广告销量"` 
    CpcSpSalesNum string `json:"cpcSpSalesNum" gorm:"type:bigint;comment:CPC推广销量"` 
    FbaShippingCredits string `json:"fbaShippingCredits" gorm:"type:varchar(100);comment:FBA运输信用分"` 
    FbmShippingCredits string `json:"fbmShippingCredits" gorm:"type:varchar(100);comment:FBM运输信用分"` 
    PromotionalRebates string `json:"promotionalRebates" gorm:"type:varchar(100);comment:促销折扣"` 
    AdjCompensation string `json:"adjCompensation" gorm:"type:varchar(100);comment:库存调整赔偿"` 
    AdjCompensationPcFee string `json:"adjCompensationPcFee" gorm:"type:varchar(100);comment:FBA库存赔偿"` 
    AdjCompensationNum string `json:"adjCompensationNum" gorm:"type:bigint;comment:库存调整数量"` 
    Refund string `json:"refund" gorm:"type:varchar(100);comment:退款"` 
    FbaRefund string `json:"fbaRefund" gorm:"type:varchar(100);comment:FBA退款"` 
    FbmRefund string `json:"fbmRefund" gorm:"type:varchar(100);comment:FBM退款"` 
    RefundNum string `json:"refundNum" gorm:"type:bigint;comment:退款量"` 
    FbaRefundNum string `json:"fbaRefundNum" gorm:"type:bigint;comment:FBA退款量"` 
    FbmRefundNum string `json:"fbmRefundNum" gorm:"type:bigint;comment:FBM退款量"` 
    RefundRate string `json:"refundRate" gorm:"type:varchar(100);comment:退款率"` 
    ReturnNum string `json:"returnNum" gorm:"type:bigint;comment:退货量"` 
    ReturnSellableNum string `json:"returnSellableNum" gorm:"type:bigint;comment:退货可售商品数"` 
    ReturnUnsellableNum string `json:"returnUnsellableNum" gorm:"type:bigint;comment:退货不可售商品数"` 
    ReturnRate string `json:"returnRate" gorm:"type:varchar(100);comment:退货比例"` 
    SaleOrderFee string `json:"saleOrderFee" gorm:"type:varchar(100);comment:销售订单"` 
    SaleSellingFee string `json:"saleSellingFee" gorm:"type:varchar(100);comment:销售成本"` 
    SaleFbaFee string `json:"saleFbaFee" gorm:"type:varchar(100);comment:FBA发货费"` 
    SaleFbmFee string `json:"saleFbmFee" gorm:"type:varchar(100);comment:FBM发货费"` 
    SaleOtherFee string `json:"saleOtherFee" gorm:"type:varchar(100);comment:其他费"` 
    MultiChannelOrderFee string `json:"multiChannelOrderFee" gorm:"type:varchar(100);comment:多渠道订单"` 
    McFbaShipFee string `json:"mcFbaShipFee" gorm:"type:varchar(100);comment:FBA发货费"` 
    RefundOrderFees string `json:"refundOrderFees" gorm:"type:varchar(100);comment:退款订单"` 
    RefundSellingFee string `json:"refundSellingFee" gorm:"type:varchar(100);comment:平台费"` 
    RefundFbaFee string `json:"refundFbaFee" gorm:"type:varchar(100);comment:FBA发货费"` 
    RefundOtherFee string `json:"refundOtherFee" gorm:"type:varchar(99);comment:其他费"` 
    SumCpcCost string `json:"sumCpcCost" gorm:"type:varchar(100);comment:广告费"` 
    CpcCostDiff string `json:"cpcCostDiff" gorm:"type:varchar(100);comment:CPC广告费差异"` 
    CpcSpCost string `json:"cpcSpCost" gorm:"type:varchar(100);comment:CPC推广费"` 
    CpcSdCost string `json:"cpcSdCost" gorm:"type:varchar(100);comment:CPC广告成本"` 
    SumPromoteFee string `json:"sumPromoteFee" gorm:"type:varchar(100);comment:推广费用"` 
    LdPromoteFee string `json:"ldPromoteFee" gorm:"type:varchar(100);comment:LD费"` 
    CouponsPromoteFee string `json:"couponsPromoteFee" gorm:"type:varchar(100);comment:优惠劵"` 
    ErpPromoteFee string `json:"erpPromoteFee" gorm:"type:varchar(100);comment:早期评论人计划"` 
    SumFbaStorageFee string `json:"sumFbaStorageFee" gorm:"type:varchar(100);comment:仓储费用"` 
    FbaStorageFee string `json:"fbaStorageFee" gorm:"type:varchar(100);comment:FBA月仓储费"` 
    FbaLongStorageFee string `json:"fbaLongStorageFee" gorm:"type:varchar(100);comment:FBA长期仓储费"` 
    FbaStorageFeeDiff string `json:"fbaStorageFeeDiff" gorm:"type:varchar(100);comment:FBA月仓储费差"` 
    SumStorageOtherFee string `json:"sumStorageOtherFee" gorm:"type:varchar(100);comment:其他仓储费用"` 
    FbaDisposalNum string `json:"fbaDisposalNum" gorm:"type:bigint;comment:FBA销毁量"` 
    FbaDisposalFee string `json:"fbaDisposalFee" gorm:"type:varchar(100);comment:FBA销毁费"` 
    FbaRemovalNum string `json:"fbaRemovalNum" gorm:"type:bigint;comment:FBA移除量"` 
    FbaRemovalFee string `json:"fbaRemovalFee" gorm:"type:varchar(100);comment:FBA移除费"` 
    FbaReturnFee string `json:"fbaReturnFee" gorm:"type:varchar(100);comment:退货入仓费"` 
    FbaShipFee string `json:"fbaShipFee" gorm:"type:varchar(100);comment:亚马逊合作承运费"` 
    FbaStorageOtherFee string `json:"fbaStorageOtherFee" gorm:"type:varchar(100);comment:其他费"` 
    FbaInventoryPlacementFee string `json:"fbaInventoryPlacementFee" gorm:"type:varchar(100);comment:库存配置费"` 
    SumFbaAdjustmentFee string `json:"sumFbaAdjustmentFee" gorm:"type:varchar(100);comment:FBA仓储调整费用合计"` 
    FbaAdjustmentFee string `json:"fbaAdjustmentFee" gorm:"type:varchar(100);comment:库存调整费"` 
    FbaAdjustmentDiffFee string `json:"fbaAdjustmentDiffFee" gorm:"type:varchar(100);comment:FBA库存调整费差异"` 
    SumOtherFee string `json:"sumOtherFee" gorm:"type:varchar(100);comment:其他费合计"` 
    SubscriptionFee string `json:"subscriptionFee" gorm:"type:varchar(100);comment:订阅费"` 
    OtherFee string `json:"otherFee" gorm:"type:varchar(100);comment:其他费"` 
    Tax string `json:"tax" gorm:"type:varchar(100);comment:税费"` 
    SumPurchaseFee string `json:"sumPurchaseFee" gorm:"type:varchar(100);comment:总采购成本"` 
    PurchasePcFee string `json:"purchasePcFee" gorm:"type:varchar(100);comment:采购成本"` 
    PurchaseDpcFee string `json:"purchaseDpcFee" gorm:"type:varchar(100);comment:采购成本（销毁）"` 
    PurchaseRtpcFee string `json:"purchaseRtpcFee" gorm:"type:varchar(100);comment:采购成本（退货）"` 
    PurchaseAcpcFee string `json:"purchaseAcpcFee" gorm:"type:varchar(100);comment:采购成本（赔偿）"` 
    PurchaseMcpcFee string `json:"purchaseMcpcFee" gorm:"type:varchar(100);comment:采购成本"` 
    PurchaseRpcFee string `json:"purchaseRpcFee" gorm:"type:varchar(100);comment:采购成本"` 
    SumHeadTripFee string `json:"sumHeadTripFee" gorm:"type:varchar(100);comment:总头程费用"` 
    HeadTripPcFee string `json:"headTripPcFee" gorm:"type:varchar(100);comment:头程费用"` 
    HeadTripDpcFee string `json:"headTripDpcFee" gorm:"type:varchar(100);comment:头程费用（销毁）"` 
    HeadTripRtpcFee string `json:"headTripRtpcFee" gorm:"type:varchar(100);comment:头程费用（退货）"` 
    HeadTripAcpcFee string `json:"headTripAcpcFee" gorm:"type:varchar(100);comment:头程费用（赔偿）"` 
    HeadTripMcpcFee string `json:"headTripMcpcFee" gorm:"type:varchar(100);comment:头程费用"` 
    HeadTripRpcFee string `json:"headTripRpcFee" gorm:"type:varchar(100);comment:头程费用"` 
    SumBatchPurchaseFee string `json:"sumBatchPurchaseFee" gorm:"type:varchar(100);comment:批次采购成本合计"` 
    BatchPurchasePcFee string `json:"batchPurchasePcFee" gorm:"type:varchar(100);comment:批次采购成本"` 
    SumBatchPurchaseOtherFee string `json:"sumBatchPurchaseOtherFee" gorm:"type:varchar(100);comment:批次采购其他成本合计"` 
    BatchPurchaseCrpcFee string `json:"batchPurchaseCrpcFee" gorm:"type:varchar(100);comment:批次采购成本（退货）"` 
    BatchPurchaseApcFee string `json:"batchPurchaseApcFee" gorm:"type:varchar(100);comment:批次采购成本（赔偿）"` 
    BatchPurchaseVrpcFee string `json:"batchPurchaseVrpcFee" gorm:"type:varchar(100);comment:批次采购成本（供应商退货）"` 
    BatchPurchaseDfpcFee string `json:"batchPurchaseDfpcFee" gorm:"type:varchar(100);comment:批次采购成本（差异）"` 
    SumBatchHeadTripFee string `json:"sumBatchHeadTripFee" gorm:"type:varchar(100);comment:批次头程费用合计"` 
    BatchTripPcFee string `json:"batchTripPcFee" gorm:"type:varchar(100);comment:批次头程费用"` 
    SumBatchHeadTripOtherFee string `json:"sumBatchHeadTripOtherFee" gorm:"type:varchar(100);comment:批次头程其他费用合计"` 
    BatchTripCrpcFee string `json:"batchTripCrpcFee" gorm:"type:varchar(100);comment:批次头程费用（退货）"` 
    BatchTripApcFee string `json:"batchTripApcFee" gorm:"type:varchar(100);comment:批次头程费用（赔偿）"` 
    BatchTripVrpcFee string `json:"batchTripVrpcFee" gorm:"type:varchar(100);comment:批次头程费用（供应商退货）"` 
    BatchTripDfpcFee string `json:"batchTripDfpcFee" gorm:"type:varchar(100);comment:批次头程费用（差异）"` 
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
    SaleRatioByAsin string `json:"saleRatioByAsin" gorm:"type:varchar(100);comment:ASIN销售率"` 
    SaleRatioByShop string `json:"saleRatioByShop" gorm:"type:varchar(100);comment:店铺销售率"` 
    AsinOtherFees string `json:"asinOtherFees" gorm:"type:varchar(100);comment:ASIN其他费"` 
    ShopOtherFees string `json:"shopOtherFees" gorm:"type:varchar(100);comment:店铺其他费用"` 
    SumSalesNum string `json:"sumSalesNum" gorm:"type:bigint;comment:销售数量"` 
    models.ModelTime
    models.ControlBy
}

func (ErpFinanceAsinSummary) TableName() string {
    return "erp_finance_asin_summary"
}

func (e *ErpFinanceAsinSummary) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *ErpFinanceAsinSummary) GetId() interface{} {
	return e.Id
}
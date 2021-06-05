CREATE TABLE `erp_finance_performance_report`(
    id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `create_by` int(11) unsigned DEFAULT NULL,
    `update_by` int(11) unsigned DEFAULT NULL,
    user_name varchar(100) NULL COMMENT '用户名'
    asin_num bigint(11) NULL COMMENT 'ASIN数量',
    sales_price Double NULL COMMENT '销售额',
    ad_sales_price bigint(11) NULL COMMENT '广告销售额',
    asoas bigint(11) NULL COMMENT 'ASoAS',
    product_total_num bigint(11) NULL COMMENT '销量',
    refund_num bigint(11) NULL COMMENT '退款量',
    refund_rate varchar(100) NULL COMMENT '退款率',
    return_num bigint(11) NULL COMMENT '退货量',
    return_rate varchar(100) NULL COMMENT '退货率',
    refund_price bigint(11) NULL COMMENT '退款',
    promotional_rebates_price bigint(11) NULL COMMENT '促销折扣',
    promotional_rebates_price_rate bigint(11) NULL COMMENT '促销折扣占比',
    sale_fee Double NULL COMMENT '发货费',
    sale_selling_fee varchar(100) NULL COMMENT '销售成本',
    sale_other_fee varchar(100) NULL COMMENT '其他费',
    sale_selling_fee_rate Double NULL COMMENT '平台费占比',
    adj_compensation_price bigint(11) NULL COMMENT 'FBA库存赔偿',
    sum_cpc_cost varchar(100) NULL COMMENT '广告费',
    acos bigint(11) NULL COMMENT 'ACOS',
    sum_promote_fee varchar(100) NULL COMMENT '推广费用',
    sum_fba_storage_fee varchar(100) NULL COMMENT '仓储费用',
    sum_fba_storage_fee_rate Double NULL COMMENT 'FBA仓储费占比',
    sum_storage_other_fee varchar(100) NULL COMMENT '其他仓储费用',
    fba_adjustment_fee varchar(100) NULL COMMENT '库存调整费',
    sum_other_fee varchar(100) NULL COMMENT '其他费合计',
    tax varchar(100) NULL COMMENT '税费',
    sum_purchase_fee varchar(100) NULL COMMENT '采购成本',
    sum_head_trip_fee varchar(100) NULL COMMENT '头程费用',
    sum_purchase_fee_rate Double NULL COMMENT '采购成本占比',
    sum_head_trip_fee_rate Double NULL COMMENT '头程费用占比',
    evaluation_fee varchar(100) NULL COMMENT '测评费',
    asin_other_fee varchar(100) NULL COMMENT 'ASIN其他费',
    shop_other_fee varchar(100) NULL COMMENT '店铺其他费',
    fixed_fee varchar(100) NULL COMMENT '固定费用',
    gross_profit varchar(100) NULL COMMENT '毛利润',
    gross_profit_rate varchar(100) NULL COMMENT '毛利率',
    currency varchar(100) NULL COMMENT '币种',
    PRIMARY KEY (`id`) USING BTREE
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='ASIN月费';

-- Request URL: https://demo.sellfox.com/api/performanceReport/getSum.json?orderBy=&desc=&shopIds=&type=1&marketplaceIds=&startMonth=202104&endMonth=202104&currency=USD&uids=
-- orderBy: 
-- desc: 
-- shopIds: 
-- type: 1
-- marketplaceIds: 
-- startMonth: 202104
-- endMonth: 202104
-- currency: USD
-- uids: 

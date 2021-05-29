CREATE TABLE `data`(
    data_id bigint(100) NOT NULL AUTO_INCREMENT COMMENT 'data_id',
    user_name varchar(100) NULL COMMENT '演示用户-9914',
    root_id bigint(100) NULL COMMENT 'root_id',
    asin_num bigint(11) NULL COMMENT '25',
    sales_price Double NULL COMMENT '144.35',
    ad_sales_price bigint(11) NULL COMMENT '0',
    asoas bigint(11) NULL COMMENT '0',
    product_total_num bigint(11) NULL COMMENT '11',
    refund_num bigint(11) NULL COMMENT '0',
    refund_rate bigint(11) NULL COMMENT '0',
    return_num bigint(11) NULL COMMENT '1',
    return_rate Double NULL COMMENT '9.09',
    refund_price bigint(11) NULL COMMENT '0',
    promotional_rebates_price bigint(11) NULL COMMENT '0',
    promotional_rebates_price_rate bigint(11) NULL COMMENT '0',
    sale_fee Double NULL COMMENT '-44.42',
    sale_selling_fee Double NULL COMMENT '-13.36',
    sale_other_fee bigint(11) NULL COMMENT '0',
    sale_selling_fee_rate Double NULL COMMENT '9.26',
    adj_compensation_price bigint(11) NULL COMMENT '0',
    sum_cpc_cost Double NULL COMMENT '-82.96',
    acos bigint(11) NULL COMMENT '0',
    sum_promote_fee bigint(11) NULL COMMENT '0',
    sum_fba_storage_fee Double NULL COMMENT '-0.87',
    sum_fba_storage_fee_rate Double NULL COMMENT '0.6',
    sum_storage_other_fee bigint(11) NULL COMMENT '0',
    fba_adjustment_fee bigint(11) NULL COMMENT '0',
    sum_other_fee Double NULL COMMENT '10.13',
    tax Double NULL COMMENT '13.34',
    sum_purchase_fee Double NULL COMMENT '-37.49',
    sum_head_trip_fee Double NULL COMMENT '-42.99',
    sum_purchase_fee_rate Double NULL COMMENT '25.97',
    sum_head_trip_fee_rate Double NULL COMMENT '29.78',
    evaluation_fee bigint(11) NULL COMMENT '0',
    asin_other_fee bigint(11) NULL COMMENT '0',
    shop_other_fee bigint(11) NULL COMMENT '0',
    fixed_fee bigint(11) NULL COMMENT '0',
    gross_profit Double NULL COMMENT '-54.28',
    gross_profit_rate Double NULL COMMENT '-37.6',
    currency varchar(100) NULL COMMENT 'USD',
    PRIMARY KEY (`data_id`) USING BTREE,
    KEY `root_id` (`root_id`),
    CONSTRAINT `data_fk_root_id` FOREIGN KEY (`root_id`) REFERENCES `root` (`root_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='data';

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
CREATE TABLE `erp_finance_lot_cost`(
    id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `create_by` int(11) unsigned DEFAULT NULL,
    `update_by` int(11) unsigned DEFAULT NULL,
    rows_id bigint(100) NOT NULL AUTO_INCREMENT COMMENT 'rows_id',
    shop_id varchar(100) NULL COMMENT '店铺id',
    page_id bigint(100) NULL COMMENT 'page_id',
    region_shop_id bigint(11) NULL COMMENT '387',
    shop_name varchar(100) NULL COMMENT '店铺名称',
    marketplace_id varchar(100) NULL COMMENT '市场id',
    marketplace_name varchar(100) NULL COMMENT '市场名称',
    sku varchar(100) NULL COMMENT 'SKU',
    fnsku varchar(100) NULL COMMENT 'MSKU',
    attr bigint(11) NULL COMMENT '类型id',
    attr_name varchar(100) NULL COMMENT '类型名称',
    sellable bigint(11) NULL COMMENT '是否可售id',
    sellable_name varchar(100) NULL COMMENT '库存属性',
    quantity bigint(11) NULL COMMENT '数量',
    document_number varchar(100) NULL COMMENT '单据号',
    fba_shipment_id varchar(100) NULL COMMENT '货件批次',
    batch_purchase_pc_fee varchar(100) NULL COMMENT '批次采购成本',
    invoice varchar(100) NULL COMMENT '发货单',
    inventory_cost bigint(11) NULL COMMENT '库存成本',
    transport_cost Double NULL COMMENT '头程费用',
    currency varchar(100) NULL COMMENT '币种',
    received_date varchar(100) NULL COMMENT '时间',
    cost_type bigint(11) NULL COMMENT '成本类型',
    PRIMARY KEY (`id`) USING BTREE
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='批次成本';

CREATE TABLE `erp_finance_lot_cost_summary`(
    id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `create_by` int(11) unsigned DEFAULT NULL,
    `update_by` int(11) unsigned DEFAULT NULL,
    quantity bigint(11) NULL COMMENT '数量合计',
    inventory_cost Double NULL COMMENT '库存成本合计',
    transport_cost Double NULL COMMENT '头程费用合计',
    currency varchar(100) NULL COMMENT '币种',
    PRIMARY KEY (`id`) USING BTREE
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='批次成本合计';
-- Request URL: https://demo.sellfox.com/api/batchCost/pageList.json?pageNo=1&pageSize=20&orderBy=receivedDate&desc=false&regionId=142&shopIds=1860,387,1861,1859,1862&sids=&sellable=&startDate=2021-04-30&endDate=2021-05-29&attrs=&searchType=sku&searchContent=
-- pageNo: 1
-- pageSize: 20
-- orderBy: receivedDate
-- desc: false
-- regionId: 142
-- shopIds: 1860,387,1861,1859,1862
-- sids: 
-- sellable: 
-- startDate: 2021-04-30
-- endDate: 2021-05-29
-- attrs: 
-- searchType: sku
-- searchContent: 

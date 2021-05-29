CREATE TABLE `rows`(
    rows_id bigint(100) NOT NULL AUTO_INCREMENT COMMENT 'rows_id',
    shop_id varchar(100) NULL COMMENT 'null',
    page_id bigint(100) NULL COMMENT 'page_id',
    region_shop_id bigint(11) NULL COMMENT '387',
    shop_name varchar(100) NULL COMMENT 'null',
    marketplace_id varchar(100) NULL COMMENT 'null',
    marketplace varchar(100) NULL COMMENT 'null',
    marketplace_name varchar(100) NULL COMMENT 'null',
    sku varchar(100) NULL COMMENT 'SDSKU-5003474',
    fnsku varchar(100) NULL COMMENT 'SDSKU-6459857',
    attr bigint(11) NULL COMMENT '1',
    attr_name varchar(100) NULL COMMENT '月初库存',
    sellable bigint(11) NULL COMMENT '1',
    sellable_name varchar(100) NULL COMMENT '可售库存',
    quantity bigint(11) NULL COMMENT '10',
    document_number varchar(100) NULL COMMENT 'null',
    fba_shipment_id varchar(100) NULL COMMENT 'null',
    batch varchar(100) NULL COMMENT '初始化库存',
    invoice varchar(100) NULL COMMENT 'null',
    inventory_cost bigint(11) NULL COMMENT '491',
    transport_cost Double NULL COMMENT '253.8',
    currency varchar(100) NULL COMMENT 'CNY',
    received_date varchar(100) NULL COMMENT '2021-05-01',
    cost_type bigint(11) NULL COMMENT '1',
    PRIMARY KEY (`rows_id`) USING BTREE,
    KEY `page_id` (`page_id`),
    CONSTRAINT `rows_fk_page_id` FOREIGN KEY (`page_id`) REFERENCES `page` (`page_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='rows';
CREATE TABLE `statistics`(
    statistics_id bigint(100) NOT NULL AUTO_INCREMENT COMMENT 'statistics_id',
    quantity bigint(11) NULL COMMENT '122',
    data_id bigint(100) NULL COMMENT 'data_id',
    inventory_cost Double NULL COMMENT '5990.2',
    transport_cost Double NULL COMMENT '3096.36',
    currency varchar(100) NULL COMMENT 'CNY',
    PRIMARY KEY (`statistics_id`) USING BTREE,
    KEY `data_id` (`data_id`),
    CONSTRAINT `statistics_fk_data_id` FOREIGN KEY (`data_id`) REFERENCES `data` (`data_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='statistics';


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
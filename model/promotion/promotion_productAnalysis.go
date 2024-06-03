package promotion

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespProductAnalysis) GetData(r ReqPromotionAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(GetProductSqlByParam("list", r))
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()
	i.Records = []ProductAnalysisNode{}
	i.Sum = []ProductAnalysisNode{}

	fmt.Println("RespProductAnalysis sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespProductAnalysis sql\n", mainSql)
		return e
	}
	i.GetDataCount(r)
	if len(i.Records) == 0 {
		fmt.Println("RespProductAnalysis sql\n", mainSql)
		i.Records = []ProductAnalysisNode{}
		return nil
	}

	// i.GetDataSum(r)
	return nil

}

type DCount struct {
	Count int64 `json:"count"`
}

func (i *RespProductAnalysis) GetDataSum(r ReqPromotionAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_PROMOTION_PRODUCTLIST_Sum)
	sqlp.SetSql(GetProductSqlByParam("sum", r))
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("RespProductAnalysis sum sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Sum).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespProductAnalysis sum sql\n", mainSql)
		return e
	}

	return nil
}

func (i *RespProductAnalysis) GetDataCount(r ReqPromotionAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(GetProductSqlByParam("count", r))
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()
	var count DCount

	fmt.Println("RespProductAnalysis count sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&count).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespProductAnalysis count sql\n", mainSql)
		return e
	}
	i.Count = int64(count.Count)
	return nil
}

// 根据参数返回不同的 Sql
func GetProductSqlByParam(sType string, r ReqPromotionAllSearch) string {
	var tempSql = SQL_PROMOTION_PRODUCTLIST
	if sType == "count" {
		tempSql = SQL_PROMOTION_PRODUCTLIST_Count

	}
	if sType == "sum" {
		tempSql = SQL_PROMOTION_PRODUCTLIST_Sum
	}
	var tableType = "bidtype"
	if len(r.KeywordFilter) >= 1 {
		tableType = "keyword"

	}
	if len(r.AudienceFilter) >= 1 {
		tableType = "crowd"

	}
	if tableType == "bidtype" {
		tempSql = fmt.Sprintf(tempSql, " from wanxiang_product w ", " AND (w.bid_type in :bidtype  ) ")
	}
	if tableType == "keyword" {
		tempSql = fmt.Sprintf(tempSql, " from wanxiang_keywords w ", " AND (w.keyword_name in :keyword  ) ")
	}
	if tableType == "crowd" {
		tempSql = fmt.Sprintf(tempSql, " from wanxiang_audience w ", " AND (w.crowd_type in :crowd  ) ")
	}
	// table
	// from wanxiang_product w  或者 from wanxiang_keywords w 或者 from wanxiang_audience w

	// where
	// AND (w.bid_type in :bidtype  )
	// AND (w.keyword_name in :keyword   )
	// AND (w.crowd_type in :crowd   )

	return tempSql
}

const SQL_PROMOTION_PRODUCTLIST_Count = `
with bp1 as (
	-- 获取结束日期的商品货盘
		select 
			bpc.product_id,
			bpc.pallet,
			bpc.pre_pallet,
			bpc.pallet_change
		from biz_product_classes bpc
	WHERE 
			bpc.statistic_date = :enddate   
	       AND (bpc.pallet in :pallet  )
	
	),
	wp1 as (
	select 
	bp1.pallet,
	w.product_id,
	bp.product_name,
	bp.product_alias,
	bp.category_name as level_three_category,
	sum(w.spend) as cost,
	sum(w.gmv) as gmv,
	sum(w.gmv)/NULLIF(sum(w.spend),0.0) AS roi,
	sum(w.dir_sell_amount)/NULLIF(sum(w.spend),0.0) AS direct_roi,
	sum(w.idr_sell_amount)/NULLIF(sum(w.spend),0.0) AS indirect_roi,
	SUM(w.new_customers)/NULLIF(sum(w.buyer_count),0.0) as new_customer_percentage,
	1-SUM(w.new_customers)/NULLIF(sum(w.buyer_count),0.0) as existing_customer_percentage,
	sum(w.clicktraffic) as clicks,
	sum(w.clicktraffic)/sum(w.impressions) as click_through_rate,
	sum(w.gmv_count)/sum(w.clicktraffic) as conversion_rate,
	sum(w.gmv)/sum(w.buyer_count) as average_order_value,
	sum(w.dir_sell_amount) as direct_transaction_amount,
	sum(w.dir_sell_count) as direct_transaction_count,
	sum(w.idr_sell_amount) as indirect_transaction_amount,
	sum(w.idr_sell_count) as indirect_transaction_count,
	sum(w.coll_add_count)/sum(w.clicktraffic) as favorite_add_to_cart_rate
	-- 后面两段替换来源数据表
	` +
	`%s` + // from wanxiang_product w  或者 from wanxiang_keywords w 或者 from wanxiang_audience w
	`left join bp1 
	on w.product_id = bp1.product_id
	left join biz_product bp 
	on w.product_id = bp.product_id
	where w.datetimekey >= :startdate AND w.datetimekey <= :enddate 
	AND (bp1.pallet in :pallet  )
	AND (bp1.pallet in :subpallet  )
	` +
	`%s` + // AND (w.bid_type in :bidtype  )  或者 AND (w.keyword_name in :keyword  ) 或者 AND (w.crowd_type in :crowd   )
	`
	-- AND (w.product_id in :ids  )	
	AND (promotion_type in  :scene  ) 
	AND (bp.responsible in :resperson )
	-- 缺出价方式的筛选条件
	 
	group by bp1.pallet,
	w.product_id,
	bp.product_name,
	bp.product_alias,
	bp.category_name
	),
	wp2 as (
	select 
	sum(cost) as cost_all,
	sum(gmv) as gmv_all
	from wp1
	)
	
	select
	count( DISTINCT wp1.product_id) as count
	from wp1
	left join wp2 on 1=1


`

const SQL_PROMOTION_PRODUCTLIST_Sum = `
with bp1 as (
	-- 获取结束日期的商品货盘
		select 
			bpc.product_id,
			bpc.pallet,
			bpc.pre_pallet,
			bpc.pallet_change
		from biz_product_classes bpc
	WHERE 
			bpc.statistic_date = :enddate   
	       AND (bpc.pallet in :pallet  )
	
	),
	wp1 as (
	select 
	bp1.pallet,
	w.product_id,
	bp.product_name,
	bp.product_alias,
	bp.category_name as level_three_category,
	sum(w.spend) as cost,
	sum(w.gmv) as gmv,
	sum(w.gmv)/NULLIF(sum(w.spend),0.0) AS roi,
	sum(w.dir_sell_amount)/NULLIF(sum(w.spend),0.0) AS direct_roi,
	sum(w.idr_sell_amount)/NULLIF(sum(w.spend),0.0) AS indirect_roi,
	SUM(w.new_customers)/NULLIF(sum(w.buyer_count),0.0) as new_customer_percentage,
	1-SUM(w.new_customers)/NULLIF(sum(w.buyer_count),0.0) as existing_customer_percentage,
	sum(w.clicktraffic) as clicks,
	sum(w.impressions) as impressions,
	sum(w.clicktraffic)/sum(w.impressions) as click_through_rate,
	sum(w.gmv_count)/sum(w.clicktraffic) as conversion_rate,
	sum(w.gmv)/sum(w.buyer_count) as average_order_value,
	sum(w.dir_sell_amount) as direct_transaction_amount,
	sum(w.dir_sell_count) as direct_transaction_count,
	sum(w.idr_sell_amount) as indirect_transaction_amount,
	sum(w.idr_sell_count) as indirect_transaction_count,
	sum(w.coll_add_count)/sum(w.clicktraffic) as favorite_add_to_cart_rate
	-- 后面两段替换来源数据表
	` +
	`%s` + // from wanxiang_product w  或者 from wanxiang_keywords w 或者 from wanxiang_audience w
	`left join bp1 
	on w.product_id = bp1.product_id
	left join biz_product bp 
	on w.product_id = bp.product_id
	where w.datetimekey >= :startdate AND w.datetimekey <= :enddate 
	AND (bp1.pallet in :pallet  )
	AND (bp1.pallet in :subpallet  )
	` +
	`%s` + // AND (w.bid_type in :bidtype  )  或者 AND (w.keyword_name in :keyword  ) 或者 AND (w.crowd_type in :crowd   )
	`
	-- AND (w.product_id in :ids  )	
	AND (promotion_type in  :scene  ) 
	AND (bp.responsible in :resperson )
	-- 缺出价方式的筛选条件
	 
	group by bp1.pallet,
	w.product_id,
	bp.product_name,
	bp.product_alias,
	bp.category_name
	),
	wp2 as (
	select 
	sum(cost) as cost_all,
	sum(gmv) as gmv_all
	from wp1
	)

	select 
    sum(wp1.cost) as cost,
  	sum(wp1.cost) /  wp2.cost_all  AS cost_percentage,
  	NULL AS cost_trend,
  	sum(wp1.gmv) as gmv,
  	sum(wp1.gmv) /  wp2.gmv_all  AS gmv_percentage,
  	NULL AS gmv_trend,
  	avg(wp1.roi) as roi ,
  	NULL AS roi_trend,
 	  avg(wp1.direct_roi),
	avg(wp1.indirect_roi) as indirect_roi,
	sum(
		wp1.new_customer_percentage
	) as new_customer_percentage,
	sum(
		wp1.existing_customer_percentage
	) as existing_customer_percentage,
	sum(wp1.clicks) as clicks,
	sum(wp1.clicks)/sum(w.impressions) as click_through_rate,
	-- sum(w.clicktraffic)/sum(w.impressions) as click_through_rate,
	sum(wp1.gmv_counts)/sum(wp1.clicks) as conversion_rate,
	-- avg(wp1.conversion_rate) as conversion_rate,
	sum(wp1.average_order_value) as average_order_value,
	sum(
		wp1.direct_transaction_amount
	) as direct_transaction_amount,
	sum(
		wp1.direct_transaction_count
	) as direct_transaction_count,
	sum(
		wp1.indirect_transaction_amount
	) as indirect_transaction_amount,
	sum(
		wp1.indirect_transaction_count
	) as indirect_transaction_count,
	AVG( wp1.favorite_add_to_cart_rate) as favorite_add_to_cart_rate
	from wp1
	left join wp2 on 1=1

`

const SQL_PROMOTION_PRODUCTLIST = `
with bp1 as (
	-- 获取结束日期的商品货盘
		select 
			bpc.product_id,
			bpc.pallet,
			bpc.pre_pallet,
			bpc.pallet_change
		from biz_product_classes bpc
	WHERE 
			bpc.statistic_date = :enddate   
	       AND (bpc.pallet in :pallet  )
	
	),
	wp1 as (
	select 
	bp1.pallet,
	w.product_id,
	bp.product_name,
	bp.product_alias,
	bp.category_name as level_three_category,
	sum(w.spend) as cost,
	sum(w.gmv) as gmv,
	sum(w.gmv)/NULLIF(sum(w.spend),0.0) AS roi,
	sum(w.dir_sell_amount)/NULLIF(sum(w.spend),0.0) AS direct_roi,
	sum(w.idr_sell_amount)/NULLIF(sum(w.spend),0.0) AS indirect_roi,
	SUM(w.new_customers)/NULLIF(sum(w.buyer_count),0.0) as new_customer_percentage,
	1-SUM(w.new_customers)/NULLIF(sum(w.buyer_count),0.0) as existing_customer_percentage,
	sum(w.clicktraffic) as clicks,
	sum(w.clicktraffic)/sum(w.impressions) as click_through_rate,
	sum(w.gmv_count)/sum(w.clicktraffic) as conversion_rate,
	sum(w.gmv)/sum(w.buyer_count) as average_order_value,
	sum(w.dir_sell_amount) as direct_transaction_amount,
	sum(w.dir_sell_count) as direct_transaction_count,
	sum(w.idr_sell_amount) as indirect_transaction_amount,
	sum(w.idr_sell_count) as indirect_transaction_count,
	sum(w.coll_add_count)/sum(w.clicktraffic) as favorite_add_to_cart_rate
	-- 后面两段替换来源数据表
	` +
	` %s ` + // from wanxiang_product w  或者 from wanxiang_keywords w 或者 from wanxiang_audience w
	`left join bp1 
	on w.product_id = bp1.product_id
	left join biz_product bp 
	on w.product_id = bp.product_id
	where w.datetimekey >= :startdate AND w.datetimekey <= :enddate 
	AND (bp1.pallet in :pallet  )
	AND (bp1.pallet in :subpallet  )
	` +
	`%s` + // AND (w.bid_type in :bidtype  )  或者 AND (w.keyword_name in :keyword  ) 或者 AND (w.crowd_type in :crowd   )
	`
	-- AND (w.product_id in :ids  )
	AND (promotion_type in  :scene  ) 
	AND (bp.responsible in :resperson )
 
	group by bp1.pallet,
	w.product_id,
	bp.product_name,
	bp.product_alias,
	bp.category_name
	),
	wp2 as (
	select 
	sum(cost) as cost_all,
	sum(gmv) as gmv_all
	from wp1
	)
	select 
	wp1.pallet,
	wp1.product_id,
	wp1.product_name,
	wp1.product_alias,
	wp1.level_three_category,
	wp1.cost,
	wp1.cost/wp2.cost_all as cost_percentage,
	null as cost_trend,
	wp1.gmv,
	wp1.gmv/wp2.gmv_all as gmv_percentage,
	null as gmv_trend,
	wp1.roi,
	null as roi_trend,
	wp1.direct_roi,
	wp1.indirect_roi,
	wp1.new_customer_percentage,
	wp1.existing_customer_percentage,
	wp1.clicks,
	wp1.click_through_rate,
	wp1.conversion_rate,
	wp1.average_order_value,
	wp1.direct_transaction_amount,
	wp1.direct_transaction_count,
	wp1.indirect_transaction_amount,
	wp1.indirect_transaction_count,
	wp1.favorite_add_to_cart_rate
	from wp1
	left join wp2 on 1=1
	LIMIT :offset , :pageSize  

`

	// table
	// from wanxiang_product w  或者 from wanxiang_keywords w 或者 from wanxiang_audience w

	// where
	// AND (w.bid_type in :bidtype  )
	// AND (w.keyword_name in :keyword   )
	// AND (w.crowd_type in :crowd   )

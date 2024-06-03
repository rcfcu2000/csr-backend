package promotion

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespPlanAnalysis) GetData(r ReqPromotionAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(GetPlanSqlByParam("list", r))
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()
	i.Records = []PlanAnalysisNode{}
	i.Sum = []PlanAnalysisNode{}

	fmt.Println("RespPlanAnalysis sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespPlanAnalysis sql\n", mainSql)
		return e
	}
	i.GetDataCount(r)
	if len(i.Records) == 0 {
		fmt.Println("RespPlanAnalysis sql\n", mainSql)
		i.Records = []PlanAnalysisNode{}
		return nil
	}
	i.GetDataSum(r)
	return nil
}

func (i *RespPlanAnalysis) GetDataSum(r ReqPromotionAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(GetPlanSqlByParam("sum", r))
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("RespPlanAnalysis sum sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Sum).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespPlanAnalysis sum sql\n", mainSql)
		return e
	}

	return nil
}

func (i *RespPlanAnalysis) GetDataCount(r ReqPromotionAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(GetPlanSqlByParam("count", r))
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()
	var count DCount

	fmt.Println("RespPlanAnalysis count sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&count).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespPlanAnalysis count sql\n", mainSql)
		return e
	}
	i.Count = int64(count.Count)
	return nil
}

// 根据参数返回不同的 Sql
func GetPlanSqlByParam(sType string, r ReqPromotionAllSearch) string {
	var tempSql = SQL_PROMOTION_PLANLIST
	if sType == "count" {
		tempSql = SQL_PROMOTION_PLANLIST_COUNT

	}
	if sType == "sum" {
		tempSql = SQL_PROMOTION_PLANLIST_SUM
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

const SQL_PROMOTION_PLANLIST_COUNT = `

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
			AND (bpc.product_id in :ids  )
            AND (bpc.pallet in :pallet  )
	
	)

	select 
	count( DISTINCT  w.promotion_type,  w.plan_id,  w.bid_type, w.plan_name ) as count
	` +
	` %s ` + // from wanxiang_product w  或者 from wanxiang_keywords w 或者 from wanxiang_audience w
	`left join bp1 
	on w.product_id = bp1.product_id
	left join biz_product bp 
	on w.product_id = bp.product_id
	where w.datetimekey >= :startdate  AND w.datetimekey <= :enddate  
		AND (bp1.pallet in :pallet  )
		AND (bp1.pallet in :subpallet  )
		` +
	`%s` + // AND (w.bid_type in :bidtype  )  或者 AND (w.keyword_name in :keyword  ) 或者 AND (w.crowd_type in :crowd   )
	`
	 	AND (promotion_type in :scene  ) -- 场景分类筛选条件
	    AND (bp.responsible in :resperson )

`

const SQL_PROMOTION_PLANLIST_SUM = `
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
			AND (bpc.product_id in :ids  )
 
            AND (bpc.pallet in :pallet  )
	
	)
	select 
 
	sum(w.spend) as spend,
	null as spend_trend,
	sum(w.gmv) as gmv,
	null as gmv_trend,
	sum(w.gmv)/NULLIF(sum(w.spend),0.0) AS roi,
	null as roi_trend,
	sum(w.dir_sell_amount)/NULLIF(sum(w.spend),0.0) AS direct_roi,
	sum(w.idr_sell_amount)/NULLIF(sum(w.spend),0.0) AS indirect_roi,
	sum(w.clicktraffic) as clicks,
	sum(w.clicktraffic)/sum(w.impressions) as click_through_rate,
	sum(w.spend)/sum(w.clicktraffic) as cpc,
	sum(w.gmv_count)/sum(w.clicktraffic) as conversion_rate,
	sum(w.spend)/sum(w.shopcart_count) as add_to_cart_cost,
	sum(w.spend)/sum(gmv_count) as transaction_cost,
	sum(w.dir_sell_amount) as direct_transaction_amount,
	sum(w.dir_sell_count) as direct_transaction_count,
	sum(w.idr_sell_amount) as indirect_transaction_amount,
	sum(w.idr_sell_count) as indirect_transaction_count
	-- 后面两段替换来源数据表
	` +
	` %s ` + // from wanxiang_product w  或者 from wanxiang_keywords w 或者 from wanxiang_audience w
	`left join bp1 
	on w.product_id = bp1.product_id
	left join biz_product bp 
	on w.product_id = bp.product_id
	where w.datetimekey >= :startdate  AND w.datetimekey <= :enddate  
		AND (bp1.pallet in :pallet  )
		AND (bp1.pallet in :subpallet  )
		` +
	`%s` + // AND (w.bid_type in :bidtype  )  或者 AND (w.keyword_name in :keyword  ) 或者 AND (w.crowd_type in :crowd   )
	`
	 	AND (promotion_type in :scene  ) -- 场景分类筛选条件
	    AND (bp.responsible in :resperson )
	 
	--     AND (w.product_id in )-- 点击上方商品筛选下方计划明细（不确定用商品ID还是商品名称过滤）
	--     缺出价方式的筛选条件
	
	 
	 
	--	group by 
	-- w.promotion_type,
	-- w.plan_id,
	-- w.bid_type,
	-- w.plan_name
 
	
`

const SQL_PROMOTION_PLANLIST = `
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
			AND (bpc.product_id in :ids  )
 
            AND (bpc.pallet in :pallet  )
	
	)
	select 
	w.promotion_type ,
	w.plan_id,
	w.bid_type as bid_type,
	w.plan_name as campaign_name,
	sum(w.spend) as spend,
	null as spend_trend,
	sum(w.gmv) as gmv,
	null as gmv_trend,
	sum(w.gmv)/NULLIF(sum(w.spend),0.0) AS roi,
	null as roi_trend,
	sum(w.dir_sell_amount)/NULLIF(sum(w.spend),0.0) AS direct_roi,
	sum(w.idr_sell_amount)/NULLIF(sum(w.spend),0.0) AS indirect_roi,
	sum(w.clicktraffic) as clicks,
	sum(w.clicktraffic)/sum(w.impressions) as click_through_rate,
	sum(w.spend)/sum(w.clicktraffic) as cpc,
	sum(w.gmv_count)/sum(w.clicktraffic) as conversion_rate,
	sum(w.spend)/sum(w.shopcart_count) as add_to_cart_cost,
	sum(w.spend)/sum(gmv_count) as transaction_cost,
	sum(w.dir_sell_amount) as direct_transaction_amount,
	sum(w.dir_sell_count) as direct_transaction_count,
	sum(w.idr_sell_amount) as indirect_transaction_amount,
	sum(w.idr_sell_count) as indirect_transaction_count
	-- 后面两段替换来源数据表
	` +
	` %s ` + // from wanxiang_product w  或者 from wanxiang_keywords w 或者 from wanxiang_audience w
	`left join bp1 
	on w.product_id = bp1.product_id
	left join biz_product bp 
	on w.product_id = bp.product_id

	where w.datetimekey >= :startdate  AND w.datetimekey <= :enddate  
		AND (bp1.pallet in :pallet  )
		AND (bp1.pallet in :subpallet  )
		` +
	`%s` + // AND (w.bid_type in :bidtype  )  或者 AND (w.keyword_name in :keyword  ) 或者 AND (w.crowd_type in :crowd   )
	`
	 	AND (promotion_type in :scene  ) -- 场景分类筛选条件
	    AND (bp.responsible in :resperson )
	 
	--     AND (w.product_id in )-- 点击上方商品筛选下方计划明细（不确定用商品ID还是商品名称过滤）
	--     缺出价方式的筛选条件
	
	 
	 
	group by 
	w.promotion_type,
	w.plan_id,
	w.bid_type,
	w.plan_name

	LIMIT :offset , :pageSize  
	
`

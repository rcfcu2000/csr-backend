package promotion

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespPlanTrend) GetData(r ReqPromotionThendSearch) error {
	sqlp := &common.SQLProccesor{}

	l := []PlanTrendNode{}

	sqlp.SetSql(GetPlanThendSqlByParam("list", r))
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()
	i.Records = []PlanOneTrend{}
	i.Sum = PlanOneTrend{Records: []PlanTrendNode{}}

	e := global.GVA_DB.Raw(mainSql).Scan(&l).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespPlanTrend sql\n", mainSql)
		return e
	}
	if len(l) == 0 {
		fmt.Println("RespPlanTrend sql\n", mainSql)
	}
	i.Prodata(l, r)
	i.GetDataSum(r)
	return nil
}

func (i *RespPlanTrend) GetDataSum(r ReqPromotionThendSearch) error {
	sqlp := &common.SQLProccesor{}

	sqlp.SetSql(GetPlanThendSqlByParam("sum", r))
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	e := global.GVA_DB.Raw(mainSql).Scan(&i.Sum.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespPlanTrend sql\n", mainSql)
		return e
	}

	return nil
}

func (i *RespPlanTrend) Prodata(l []PlanTrendNode, r ReqPromotionThendSearch) {
	for m := 0; m < len(r.Ids); m++ {
		one := &PlanOneTrend{PlanId: r.Ids[m]}
		for n := 0; n < len(l); n++ {
			if one.PlanId == l[n].PlanId {
				one.Records = append(one.Records, l[n])

			}
		}
		i.Records = append(i.Records, *one)
	}

}

// 根据参数返回不同的 Sql
func GetPlanThendSqlByParam(sType string, r ReqPromotionThendSearch) string {
	var tempSql = SQL_PROMOTION_PLANTHEND_LIST

	if sType == "sum" {
		tempSql = SQL_PROMOTION_PLANTHEND_LIST_SUM
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

const SQL_PROMOTION_PLANTHEND_LIST_SUM = `
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
	
	)
	select 
 
	w.datetimekey as date,
	sum(w.spend) as spend_trend,
	sum(w.gmv) as gmv_trend,
	sum(w.gmv)/NULLIF(sum(w.spend),0.0) AS roi_trend

	` +
	` %s ` + // from wanxiang_product w  或者 from wanxiang_keywords w 或者 from wanxiang_audience w
	`left join bp1 
	on w.product_id = bp1.product_id
	left join biz_product bp 
	on w.product_id = bp.product_id
	where w.datetimekey >= :startdate  AND w.datetimekey <= :enddate  
		AND (w.plan_id in :ids  )
		AND (bp1.pallet in :pallet  )
		` +
	` %s ` + // AND (w.bid_type in :bidtype  )  或者 AND (w.keyword_name in :keyword  ) 或者 AND (w.crowd_type in :crowd   )
	`					
	 	AND (promotion_type in :scene  ) -- 场景分类筛选条件
	    AND (bp.responsible in :resperson )

	group by 
	w.datetimekey  order by  w.datetimekey
  
`

const SQL_PROMOTION_PLANTHEND_LIST = `
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
	
	)
	select 
	w.plan_id,
	w.datetimekey as date,
	sum(w.spend) as spend_trend,
	sum(w.gmv) as gmv_trend,
	sum(w.gmv)/NULLIF(sum(w.spend),0.0) AS roi_trend
	
	` +
	` %s ` + // from wanxiang_product w  或者 from wanxiang_keywords w 或者 from wanxiang_audience w
	`left join bp1 
	on w.product_id = bp1.product_id
	left join biz_product bp 
	on w.product_id = bp.product_id
	where w.datetimekey >= :startdate  AND w.datetimekey <= :enddate  
		AND (w.plan_id in :ids  )
		AND (bp1.pallet in :pallet  )
		` + ` %s ` + // AND (w.bid_type in :bidtype  )  或者 AND (w.keyword_name in :keyword  ) 或者 AND (w.crowd_type in :crowd   )
	`AND (promotion_type in :scene  ) -- 场景分类筛选条件
	    AND (bp.responsible in :resperson )

	group by 
	w.plan_id,w.datetimekey  order by w.plan_id, w.datetimekey
  
`

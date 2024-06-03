package product

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespIndexTrendList) GetData(r ReqProductAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_PRODUCT_TrendList)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("RespIndexTrendList sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespIndexTrendList sql\n", mainSql)
		return e
	}
	// i.Prodata()
	return nil

}

const SQL_PRODUCT_TrendList = `

select v.product_id as "product_id",
DATE_FORMAT(v.statistic_date, '%Y-%m-%d')  as "date",
sum(v.visitors_count) as "visitors_count", -- 商品访客数
sum(v.paid_amount) as "gmv",  -- gmv
sum(v.paid_buyers) / sum(v.visitors_count) as "payment_conversion_rate",  -- 订单支付转化率
sum(v.returning_buyers_paid) / sum(v.paid_buyers) as "returning_customer_ratio",  -- 老客户占比

sum(v.returning_buyers_paid) / sum(v.paid_buyers) as "returning_customer_ratio",  -- 老客户占比
 
CASE
WHEN  sum(bpt_gmv) > 0 THEN sum(bpt_search_gmv) / sum(v.paid_amount)
ELSE 0
END AS search_gmv_ratio, -- 搜索渠道贡献的GMV占比

CASE
WHEN sum(bpt_visitors_count) > 0 THEN sum(bpt_search_count) / sum(v.visitors_count)
ELSE 0
END AS search_visitor_ratio,  -- 搜索渠道访问商品的访客比例	

sum(v.successful_refund_amount) / sum(v.paid_amount) as "refund_rate",  -- 退款率
 
sum(free_click_rate) as "free_search_click_rate",  -- 免费搜索点击率
avg(associated_leaf_category) as "bundle_purchase",  -- 连带率
avg(repurchase_rate) as "repeat_purchase_rate"    -- 复购率

from v_product_day_all  v -- v_bpp_bp =>v_product_day_all
where (v.product_id =:productid )
and ( v.statistic_date >=:startdate and v.statistic_date <= :enddate )
group by statistic_date
order by  statistic_date 

`

package product

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespProductDay) GetData(r ReqProductAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_PRODUCT_DAY)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()
	i.Records = []ProductPerformanceNode{}
	fmt.Println("RespProductDay sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespProductDay sql\n", mainSql)
		return e
	}
	i.Prodata()
	i.GetDataCount(r)
	i.GetDataSum(r)
	return nil

}
func (i *RespProductDay) Prodata() {
	// sum := 0
	// for m := 0; m < len(i.Records); m++ {
	// 	sum += int(i.Records[m].GMV)
	// }
	// for m := 0; m < len(i.Records); m++ {
	// 	i.Records[m].GMVPercentage = i.Records[m].GMV / float64(sum)
	// }
}

func (i *RespProductDay) GetDataSum(r ReqProductAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_PRODUCT_DAY_SUM)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("RespProductDay sum sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Sum).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespProductDay sum sql\n", mainSql)
		return e
	}

	return nil

}

func (i *RespProductDay) GetDataCount(r ReqProductAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_PRODUCT_DAY_COUNT)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()
	var count DCount

	fmt.Println("RespProductDay count sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&count).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespProductDay count sql\n", mainSql)
		return e
	}
	i.Count = int64(count.Count)
	return nil
}

const SQL_PRODUCT_DAY_COUNT = `
select  count( v.statistic_date  ) as count
from v_bpp_bp v
where (v.product_id =:productid )
and ( v.statistic_date >=:startdate and v.statistic_date <= :enddate )
`

const SQL_PRODUCT_DAY_SUM = `
select v.product_id as "product_id",
"合计"  as "date",
sum(v.visitors_count) as "product_visitor_count", -- 商品访客数
sum(v.paid_amount) as "gmv",  -- gmv
sum(v.paid_buyers) / sum(v.visitors_count) as "payment_conversion_rate",  -- 订单支付转化率
CONVERT(avg(price_strength), SIGNED)   as "price_power_stars",  -- 价格力
avg(price_strength_exposure) as "price_power_extra_exposure",  -- 价格力额外曝光

CASE
WHEN  sum(bpt_gmv) > 0 THEN sum(bpt_search_gmv) / sum(v.paid_amount)
ELSE 0
END AS search_gmv_ratio, -- 搜索渠道贡献的GMV占比

CASE
WHEN sum(bpt_visitors_count) > 0 THEN sum(bpt_search_count) / sum(v.visitors_count)
ELSE 0
END AS search_visitor_ratio,  -- 搜索渠道访问商品的访客比例	

sum(v.returning_buyers_paid) / sum(v.paid_buyers) as "returning_customer_ratio",  -- 老客户占比
 

sum(v.successful_refund_amount) / sum(v.paid_amount) as "refund_rate",  -- 退款率
sum(free_click_rate) as "free_search_click_through_rate",  -- 免费搜索点击率
avg(associated_leaf_category) as "bundle_purchase",  -- 连带率
0 as "associated_purchase_subcategory_width",  -- 连带率2
avg(repurchase_rate) as "repeat_purchase_rate",   -- 复购率
sum(spend) as "promotion_cost",   -- 成本
sum(spend)/sum(v.paid_amount) as "promotion_roi"    -- roi

from v_product_day_all  v   -- v_bpp_bp
where (v.product_id =:productid )
and ( v.statistic_date >=:startdate and v.statistic_date <= :enddate )
`

const SQL_PRODUCT_DAY = `

select v.product_id as "product_id",
DATE_FORMAT(v.statistic_date, '%Y-%m-%d')  as "date",
sum(v.visitors_count) as "product_visitor_count", -- 商品访客数
sum(v.paid_amount) as "gmv",  -- gmv
sum(v.paid_buyers) / sum(v.visitors_count) as "payment_conversion_rate",  -- 订单支付转化率
CONVERT(avg(price_strength), SIGNED)   as "price_power_stars",  -- 价格力
avg(price_strength_exposure) as "price_power_extra_exposure",  -- 价格力额外曝光

CASE
WHEN  sum(bpt_gmv) > 0 THEN sum(bpt_search_gmv) / sum(v.paid_amount)
ELSE 0
END AS search_gmv_ratio, -- 搜索渠道贡献的GMV占比

CASE
WHEN sum(bpt_visitors_count) > 0 THEN sum(bpt_search_count) / sum(v.visitors_count)
ELSE 0
END AS search_visitor_ratio,  -- 搜索渠道访问商品的访客比例	

sum(v.returning_buyers_paid) / sum(v.paid_buyers) as "returning_customer_ratio",  -- 老客户占比
 
sum(v.successful_refund_amount) / sum(v.paid_amount) as "refund_rate",  -- 退款率
sum(free_click_rate) as "free_search_click_through_rate",  -- 免费搜索点击率
avg(associated_leaf_category) as "bundle_purchase",  -- 连带率
avg(associated_leaf_category) as "associated_purchase_subcategory_width",  -- 连带率2
avg(repurchase_rate) as "repeat_purchase_rate",   -- 复购率

sum(spend) as "promotion_cost",   -- 成本
sum(spend)/sum(v.paid_amount) as "promotion_roi"    -- roi


 
from v_product_day_all  v   -- v_bpp_bp
where (v.product_id =:productid )
and ( v.statistic_date >=:startdate and v.statistic_date <= :enddate )
group by statistic_date
order by  statistic_date 
LIMIT :offset , :pageSize  

`

// GMV	指标	商品每日数据	支付金额
// 支付转化率	指标	商品每日数据	支付转化率=支付买家数/访客数
// 搜索访客占比	指标	商品流量数据	取来源类型=手淘搜索的访客数
// 搜索GMV占比	指标	商品每日数据
// 商品流量数据	分子取商品流量数据，来源效果=每一次，二级来源=手淘搜索，三级来源=汇总的支付金额；分母取商品每日数据的支付金额
// 老客占比	指标	商品每日数据	老买家人数占比=支付老买家数/支付买家数
// 退款率	指标	商品每日数据	退款率=成功退款金额/支付金额
// 价格力星级	指标	商品每日数据（补充指标）
// 价格力额外曝光	指标	商品每日数据（补充指标）
// 免费搜索点击率	指标	达摩盘-货品洞察
// 连带购买也子类目宽度	指标	达摩盘-货品洞察
// 复购率	指标	达摩盘-货品洞察
// 推广花费	指标	宝贝主体报表	花费
// 推广ROI	指标	宝贝主体报表	推广ROI=总成交金额 / 花费

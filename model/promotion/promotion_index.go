package promotion

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

type PromotionIndex struct {
	RespPromotionIndex1
	RespPromotionIndex2
}

func (i *PromotionIndex) GetData(r ReqPromotionAllSearch) error {
	i.GetDataIndex1(r)
	i.GetDataIndex2(r) // 先2 后1 不能乱

	return nil
}

func (i *PromotionIndex) GetDataIndex1(r ReqPromotionAllSearch) error {

	tempIndex1 := RespPromotionIndex1{}
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_PROMOTION_INDEX1)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()
	fmt.Println("GetDataIndex1 sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&tempIndex1).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("GetDataIndex1 sql\n", mainSql)
		return e
	}
	fmt.Println("tempIndex1", tempIndex1)
	i.OverallGMV = tempIndex1.OverallGMV
	i.OverallAddToCartRate = tempIndex1.OverallAddToCartRate
	i.OverallConversionRate = tempIndex1.OverallConversionRate
	i.OverallROI = tempIndex1.OverallROI

	return nil
}

func (i *PromotionIndex) GetDataIndex2(r ReqPromotionAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_PROMOTION_INDEX2)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()
	i.Records = []PromotionIndex2Node{}

	tempIndex2 := RespPromotionIndex2{}
	fmt.Println("GetDataIndex2 sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&tempIndex2.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("GetDataIndex2 sql\n", mainSql)
		return e
	}
	i.Records = append(i.Records, tempIndex2.Records...)
	i.ProData(i.Records)
	return nil
}

func (i *PromotionIndex) ProData(Records []PromotionIndex2Node) {
	sumval := 0.0
	sumclick := 0
	sumvistor := 0
	sumcar := 0
	sumgmvcount := 0
	for m := 0; m < len(i.Records); m++ {
		sumval += i.Records[m].GMV
		i.PromotionCost += i.Records[m].Spend
		i.PromotionGMV += i.Records[m].GMV
		sumclick += int(i.Records[m].Clicks)
		sumvistor += int(i.Records[m].GuidedVisitors)
		sumcar += int(i.Records[m].ShopcartCount)
		sumgmvcount += int(i.Records[m].GmvCount)

	}
	if i.OverallGMV != 0 {
		i.PromotionGMVPercentage = i.PromotionGMV / i.OverallGMV
		i.CostPercentage = i.PromotionCost / i.OverallGMV
	}
	if sumvistor != 0 {
		i.PromotionTrafficPercentage = float64(sumclick) / float64(sumvistor)
	}
	if sumclick != 0 {
		i.PromotionAddToCartRate = float64(sumcar) / float64(sumclick)
		i.PromotionConversionRate = float64(sumgmvcount) / float64(sumclick)
	}

	for m := 0; m < len(i.Records); m++ {
		if sumval != 0 {
			if sumval != 0 {
				i.Records[m].ChannelPercentage = i.Records[m].GMV / sumval
			}
		}
	}
}

/*
参考
SELECT

	sum(visitors_count) AS "商品访客数",
	sum(v.paid_amount) AS "支付金额",
	-- gmv
	sum(v.add_to_cart_buyers) / sum(visitors_count) AS 全店加购率,
	-- 全店加购率
	sum(v.paid_buyers) / sum(visitors_count) AS 全店转化率,
	-- overall_conversion_rate
	sum(paid_amount) / sum(paid_amount) AS 全店ROI --  分母？？  全店ROI=商品每日数据的支付金额/ 宝贝主体报表的花费
	-- sum(adders_count) AS "商品加购人数",
	-- sum(payers) AS "支付买家数",
	-- sum(pay_payment) AS overall_gmv,

FROM

	xtt.v_bpp_bp_bpc v

WHERE

	pallet IN ("S", "A")

AND v.statistic_date >= "2024-01-01"
AND statistic_date <= "2024-01-11"
AND responsible IN ("李祥") -- and  v.promotion_name in ("场景推广","关键词推广","精准人群推广")
LIMIT 1000;
overall_gmv
overall_add_to_cart_rate
overall_conversion_rate
overall_roi
*/
const SQL_PROMOTION_INDEX1 = `

SELECT
	sum(visitors_count) AS "商品访客数",
	sum(v.paid_amount) AS overall_gmv,
	-- gmv
	sum(v.add_to_cart_buyers) / sum(visitors_count) AS overall_add_to_cart_rate,
	-- 全店加购率
	sum(v.paid_buyers) / sum(visitors_count) AS overall_conversion_rate,
	-- overall_conversion_rate
	sum(paid_amount) / sum(paid_amount) AS overall_roi --  分母？？  全店ROI=商品每日数据的支付金额/ 宝贝主体报表的花费
	-- sum(adders_count) AS "商品加购人数",
	-- sum(payers) AS "支付买家数",
	-- sum(pay_payment) AS overall_gmv,
FROM
	xtt.v_bpp_bp_bpc v
WHERE
	pallet IN :pallet  
AND v.statistic_date >= :startdate 
AND statistic_date <= :enddate 
AND responsible IN :resperson 
 
`

//  -- and  v.promotion_name in ("场景推广","关键词推广","精准人群推广")

/* 参考
select
w.promotion_type as "场景类型",
sum(spend) as "花费",
sum(spend)/sum(gmv_count) as "成交成本",
sum(gmv) as "总成交金额",
sum(gmv)/sum(gmv) as "GMV占比", -- 需要再计算
sum(gmv)/sum(spend) as "推广ROI",
sum(w.clicktraffic) as "点击量",
sum(w.clicktraffic)/sum(w.impressions) as "点击率",
sum(spend)/sum(w.clicktraffic) as "平均点击花费(CPC)",
sum(shopcart_count)/sum(w.clicktraffic) as "加购率",
sum(spend)/sum(shopcart_count) as "加购成本",
sum(wangwang_count) as "旺旺咨询量",
sum(w.impressions) as "展现量",
sum(gmv_count) as "总成交笔数",
sum(shopcart_count) as "总购物车数",
sum(guided_visits) as "引导访问量",
sum(guided_visitors) as "引导访问人数",
sum(buyer_count) as "成交人数"

from v_wxp_bp_bpc w
WHERE  pallet IN ( 'S','A','B','C','D' )
AND w.datetimekey >= '2024-01-01'
AND w.datetimekey <= '2024-01-10'
AND responsible IN (select DISTINCT responsible  from biz_product t)
-- and  v.promotion_name in ("场景推广","关键词推广","精准人群推广")

GROUP BY promotion_type

*/

const SQL_PROMOTION_INDEX2 = `
select
w.promotion_type as "scene_category",
sum(spend) as "spend",
sum(spend)/sum(gmv_count) as "transaction_cost",
sum(gmv) as "gmv",
sum(gmv)/sum(gmv) as "GMV占比", -- 需要再计算
sum(gmv)/sum(spend) as "promotion_roi",
sum(w.clicktraffic) as "clicks",
sum(w.clicktraffic)/sum(w.impressions) as "click_through_rate",
sum(spend)/sum(w.clicktraffic) as "cpc",
sum(shopcart_count)/sum(w.clicktraffic) as "add_to_cart_rate",
sum(spend)/sum(shopcart_count) as "add_to_cart_cost",
sum(wangwang_count) as "ali_wang_wang_inquiries",
sum(w.impressions) as "展现量",
sum(gmv_count) as "gmv_count",
sum(shopcart_count) as "shopcart_count",
sum(guided_visits) as "guided_visits",
sum(guided_visitors) as "guided_visitors",
sum(buyer_count) as "buyer_count"

from v_wxp_bp_bpc w
WHERE  pallet IN :pallet   
AND w.datetimekey >= :startdate  
AND w.datetimekey <= :enddate  
AND responsible IN :resperson   
and  w.promotion_type in :scene     

GROUP BY promotion_type

`

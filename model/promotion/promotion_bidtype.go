package promotion

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespBidTypeAnalysis) GetData(r ReqPromotionAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_PROMOTION_BIDTYPE)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()
	i.Records = []BidTypeAnalysisNode{}

	fmt.Println("RespBidTypeAnalysis sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespBidTypeAnalysis sql\n", mainSql)
		return e
	}
	if len(i.Records) == 0 {
		i.Records = []BidTypeAnalysisNode{}
		return nil
	}
	i.Prodata()
	return nil

}
func (i *RespBidTypeAnalysis) Prodata() {
	sum := 0
	for m := 0; m < len(i.Records); m++ {
		sum += int(i.Records[m].GMV)
	}
	for m := 0; m < len(i.Records); m++ {
		if sum != 0 {
			i.Records[m].GMVPercentage = i.Records[m].GMV / float64(sum)
		}
	}
}

const SQL_PROMOTION_BIDTYPE = `

select 
w.bid_type as "bid_type",
sum(spend) as "spend",
sum(gmv) as "gmv",
sum(clicktraffic) as clicks , -- 点击量 
sum(w.clicktraffic)/sum(w.impressions) as click_through_rate,  --  点击率 =  点击量 / 展现量
sum(w.spend)/sum(w.shopcart_count) as add_to_cart_cost, -- 加购成本 =  花费 / 总购物车数
sum(gmv)  /sum(spend) as "promotion_roi",  -- 推广ROI, -- 总成交金额 / 花费
sum(w.spend)/sum(gmv_count) as transaction_cost, -- 成交成本, -- 花费 / 总成交笔数
sum(w.gmv_count)/sum(w.clicktraffic) as conversion_rate -- 推广转化率 -- 总成交笔数 / 点击量
 
from v_wxp_bp_bpc w
WHERE  product_id in ( select product_id from biz_product_classes where statistic_date =:enddate and  pallet IN :subpallet   )
AND w.datetimekey >= :startdate  
AND w.datetimekey <= :enddate  
AND responsible IN :resperson  
and  w.promotion_type in :scene      
GROUP BY w.bid_type

`

// // 出价类型分析
// type BidTypeAnalysisNode struct {
// 	// 出价类型
// 	BidType string `json:"bid_type"`

// 	// 花费
// 	Spend float64 `json:"spend"`

// 	// GMV（商品交易总额）
// 	GMV float64 `json:"gmv"`

// 	// GMV占比（此次广告活动产生的GMV占总体GMV的比例）
// 	GMVPercentage float64 `json:"gmv_percentage"`

// 	// 点击量
// 	Clicks int64 `json:"clicks"`

// 	// 点击率（点击量占展现量的比例）
// 	ClickThroughRate float64 `json:"click_through_rate"`

// 	// 加购成本（每产生一次加购行为的平均成本）
// 	AddToCartCost float64 `json:"add_to_cart_cost"`

// 	// 推广投资回报率（ROI）
// 	PromotionROI float64 `json:"promotion_roi"`
// }

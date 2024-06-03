package target

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespTargetIndexData) GetData(r ReqTargetAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_TARGET_Index)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("SQL_TARGET_Index sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(i).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("SQL_TARGET_Index sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespTargetIndexData) Prodata() {

}

// GMV	指标	商品每日数据	支付金额
// 净利润	指标	万相台-宝贝主体报表
// 商品每日数据	净利润=（b.支付金额-b.成功退款金额）*c.预估毛利率-b.支付件数*c.发货费用-c.人工费用*天数-a.花费
//                    c是商品主数据表 b是bpp商品每日数据， 是主表，a是万相台宝贝主体报表
// 净利润率	指标	万相台-宝贝主体报表
// 商品每日数据	净利润率=净利润/（支付金额-成功退款金额）
// 推广花费	指标	万相台-宝贝主体报表	花费
// 推广占比	指标	万相台-宝贝主体报表
// 商品每日数据	推广占比=万相台-宝贝主体报表的总成交金额汇总/商品每日数据的支付金额汇总
// 累计GMV	指标	商品每日数据	按照结束日期所在月份，累计当月支付金额汇总，左侧为全店汇总等于右侧货盘明细的合计值
// GMV目标	指标	货盘目标（补录表）	取补录表中结束日期所在月份的GMV目标，左侧为全店汇总等于右侧货盘明细的合计值
// GMV达成率	指标		GMV达成率=累计GMV/GMV目标 GMV达成率小于达成率目标为红色，大于等于达成率目标为绿色
// 达成率目标	指标		达成率目标=结束日期的天数/结束日期所在月的总天数

const SQL_TARGET_Index = `
WITH t1 as(

	select 
	"one"  as id,
	sum(paid_amount) as gmv, 
	sum(jlr) as profit, sum(jlr)/sum(gmv_refund) as  profit_rate,   
	sum(spend) as spend, sum(wp_gmv)/ sum(paid_amount) as  promotion_percentage,
	sum(paid_amount) as month_gmv 
	from v_target_bpp2_wp_bspt  v
	where  v.statistic_date >=:startdate  and v.statistic_date <=:enddate 
	),
	 
	 t2 as( 
	select 
	  "one"  as id,
		 sum(gmv_target) as  target_gmv,
		0.0 as  target_gmv_rate,
		( 1-DATEDIFF(LAST_DAY(:enddate ),:enddate )/(DATEDIFF(:enddate ,:startdate )+1+DATEDIFF(LAST_DAY(:enddate ),:enddate )) ) as  target_day_rate
	from biz_shop_pallet_targets  bspt 
	where  bspt.statistic_date >=:startdate  and bspt.statistic_date <= LAST_DAY(:enddate ) 
	) 
	
	select t1.* , t2.target_gmv, t1.gmv /t2.target_gmv as target_gmv_rate, target_day_rate  
	from t1 
	LEFT JOIN t2 on t1.id = t2.id
`

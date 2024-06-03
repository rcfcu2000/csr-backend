package target

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespProductManagerTargetListData) GetData(r ReqTargetAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_TARGET_ProductManagerList)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("RespPalletTargetListData sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespPalletTargetListData sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespProductManagerTargetListData) Prodata() {

}

const SQL_TARGET_ProductManagerList = `
 

	SELECT
	tmain.responsible as manager,
	tmain.category_lv3 as category_lv3,
	sum(paid_amount) as gmv,	
	sum(gmv_target) as target_gmv,
	sum(paid_amount)/sum(gmv_target) as target_gmv_rate, 
	( 1-DATEDIFF(LAST_DAY(:enddate ),:enddate )/(DATEDIFF(:enddate ,:startdate )+1+DATEDIFF(LAST_DAY(:enddate ),:enddate )) ) as  target_day_rate,
	sum(jlr) as profit, 
	sum(profit_target) as  profit_rate, -- 经营利润目标
	sum(jlr)/sum(profit_target) as  profit_rate  -- 经营利润达成率
 
FROM
	v_target_bpp2_wx_bpt tmain
where
	tmain.statistic_date >= :startdate 
	AND tmain.statistic_date <=  :enddate  
group by  	tmain.responsible, tmain.category_lv3
 
`

// 责任人
// 三级类目
// GMV
// 时间进度
// GMV达成率
// GMV目标
// 经营利润
// 经营利润目标
// 经营利润达成率

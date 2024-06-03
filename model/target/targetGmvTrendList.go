package target

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespGmvTrendListData) GetData(r ReqTargetAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_TARGET_GmvTrendList)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("RespGmvTrendListData sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespGmvTrendListData sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespGmvTrendListData) Prodata() {

}

const SQL_TARGET_GmvTrendList = `
	select 
	-- statistic_date as date, 
	DATE_FORMAT(statistic_date, '%m-%d')  as date,
	sum(paid_amount) as gmv, sum(spend) as spend, sum(wp_gmv) as promotion_gmv  
	from v_bpp_wp v
	where v.statistic_date >= :startdate  and v.statistic_date <= :enddate  
	GROUP BY statistic_date
`

package shophome

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespShopHomeSumTrendData) GetData(r ReqShopHomeAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_SHOPHOME_SumTrend)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("SQL_SHOPHOME_SumTrend sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("SQL_SHOPHOME_SumTrend sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespShopHomeSumTrendData) Prodata() {

}

const SQL_SHOPHOME_SumTrend = `
select 
DATE_FORMAT(bpp.statistic_date, '%m-%d')  as date,
sum(visitors_count) as visitors,sum(paid_amount) as gmv
from biz_product_performance bpp
where bpp.statistic_date >= :startdate    
and bpp.statistic_date <= :enddate 
group by bpp.statistic_date
order by bpp.statistic_date

`

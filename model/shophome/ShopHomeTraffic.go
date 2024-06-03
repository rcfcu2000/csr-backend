package shophome

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespShopHomeTrafficData) GetData(r ReqShopHomeAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_SHOPHOME_Traffic)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("SQL_SHOPHOME_Traffic sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("SQL_SHOPHOME_Traffic sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespShopHomeTrafficData) Prodata() {

}

const SQL_SHOPHOME_Traffic = `
select 
tertiary_source as l3,
sum(visitors_count) as visitors,
sum(paid_amount) as gmv
from biz_shop_traffic bst
where tertiary_source <> '汇总'
and date >= :startdate  
and  date <= :enddate  
group by tertiary_source
`

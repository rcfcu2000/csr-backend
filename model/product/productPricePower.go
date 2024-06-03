package product

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespPricePower) GetData(r ReqProductAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_PRODUCT_PricePower)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("RespProductDay sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespProductDay sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespPricePower) Prodata() {

}

const SQL_PRODUCT_PricePower = `

select  DATE_FORMAT(v.statistic_date, '%Y-%m-%d')  as date,  
v.price_strength as "pp_level", -- 价格力星级
v.unit_price as "unit_price"  -- 单价
from v_bpdi_bp v
where (v.product_id =:productid )
and ( v.statistic_date >=:startdate and v.statistic_date <= :enddate )
order by  statistic_date 
 
`

/*  参考
select v.statistic_date as "date",
v.price_strength as "pp_level", -- 价格力星级
v.unit_price as "unit_price"  -- 单价
from v_bpdi_bp v
where (v.product_id ="" )
and ( v.statistic_date >="2024-01-01" and v.statistic_date <= "2024-01-11" )
order by  statistic_date


*/

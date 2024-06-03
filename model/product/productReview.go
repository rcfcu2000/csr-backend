package product

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespReview) GetData(r ReqProductAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_PRODUCT_Review)
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
func (i *RespReview) Prodata() {

}

const SQL_PRODUCT_Review = `

select 
DATE_FORMAT(v.statistic_date, '%Y-%m-%d')  as date,  
v.word as "keyword", -- 关键词
v.count as "count"  -- 计数
from v_bpr_bp v
where (v.product_id =:productid )
and ( v.statistic_date >=:startdate and v.statistic_date <= :enddate )
order by  statistic_date 
 
`

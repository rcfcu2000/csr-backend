package shop

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *Content) GetData(r ReqShopAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_SHOP_Content)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	i.Records = []ContentTrendNode{}

	fmt.Println("Content sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("Content sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *Content) Prodata() {

}

const SQL_SHOP_Content = `

select type, sum( amount) as amount 
from biz_shop_content t 
where  channel ="全部" and  statistic_date>=:startdate and statistic_date<=:enddate  
GROUP BY type 

`

/////////////////////////////////////////////////////////////

func (i *ContentTrendList) GetData(r ReqShopAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_SHOP_ContentTrend)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()
	i.Records = []ContentTrendNode{}

	fmt.Println("ContentTrendList sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("ContentTrendList sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *ContentTrendList) Prodata() {

}

const SQL_SHOP_ContentTrend = `
with 
  bsc as (
select statistic_date, type, sum( amount)  as amount
from biz_shop_content bsc
where  channel ="全部" and statistic_date>=:startdate   and statistic_date<=:enddate  
GROUP BY type, statistic_date
),

bpp as (
	select  statistic_date, sum( paid_amount ) as sum_amount  
	from biz_product_performance  
	where   statistic_date<=:enddate   and    statistic_date>=:startdate 
	group by statistic_date
	)

select 
bsc.statistic_date as date ,type, bsc.amount as amount, bsc.amount / bpp.sum_amount as proportion
from bsc
left join bpp on  bsc.statistic_date = bpp.statistic_date
  
 
`

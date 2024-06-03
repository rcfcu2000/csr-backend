package shop

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

// 缺  动销率

func (i *ShopIndex) GetData(r ReqShopAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_SHOP_Index)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	// fmt.Println("ShopIndex sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(i).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("ShopIndex sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *ShopIndex) Prodata() {

}

const SQL_SHOP_Index = `
select 
(
	select count(DISTINCT product_id) from biz_product_performance t 
where t.paid_amount >0 and  t.statistic_date<=:enddate  and  t.statistic_date>=:startdate 
)/ (
	select count(*) from biz_product
) as turnover_rate,

t.level, t.ranking 
from biz_shop_level t where (t.shop_name=:shop_name )
and t.statistic_date=:enddate ;
`

// 动销率
// select count( DISTINCT t.product_id) as count from biz_product_performance t
// where t.statistic_date>=:startdate and t.statistic_date>=:enddate and t.shop_name=:shop_name ;

// select count( DISTINCT t.product_id) as tcount  from biz_product t
// where t.shop_name=:shop_name ;

/////////////////////////////////////////////////////////////////////
//ShopIndexTrendList

func (i *ShopIndexTrendList) GetData(r ReqShopAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_SHOP_IndexTrend)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()
	i.Records = []ShopIndexTrendNode{}

	fmt.Println("ShopIndexTrendList sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("ShopIndexTrendList sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *ShopIndexTrendList) Prodata() {

}

const SQL_SHOP_IndexTrend = `

select  DATE_FORMAT(t.statistic_date, '%Y-%m-%d')  as date, bpp.turnover_rate, t.level, t.ranking from biz_shop_level t 

left join (
	select bpp.statistic_date, count(DISTINCT product_id)/(select count(*) from biz_product) as  turnover_rate
	from biz_product_performance bpp
	where bpp.paid_amount >0 and  bpp.statistic_date<=:enddate  and  bpp.statistic_date>=:startdate  
			group by statistic_date
) bpp on ( t.statistic_date = bpp.statistic_date )

 
where t.statistic_date>=:startdate and t.statistic_date<=:enddate and t.shop_name=:shop_name ;

`

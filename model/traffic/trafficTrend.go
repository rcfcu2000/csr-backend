package traffic

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespTrafficTrendData) GetData(r ReqTrafficAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_TRAFFIC_Trend)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("RespTrafficTrendData sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespTrafficTrendData sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespTrafficTrendData) Prodata() {

}

const SQL_TRAFFIC_Trend = `

select 
 
	DATE_FORMAT(date, '%m-%d')  as date,
	tertiary_source as tertiary_source, 
	sum(t.visitors_count) as visitors_count, 
	sum(t.paid_amount) as paid_amount 
	
from biz_shop_traffic t
where t.date<=:enddate  and  t.date>=:startdate 
-- and t.src=:belong 
-- AND (bp.responsible = :resperson OR :resperson IS NULL OR :resperson = '')
and t.tertiary_source <>"汇总"
and tertiary_source in :channel 
group by t.tertiary_source , date
 

UNION
 
select 
 
	DATE_FORMAT(date, '%m-%d')  as date,
	secondary_source as tertiary_source, 
	sum(t.visitors_count) as visitors_count, 
	sum(t.paid_amount) as paid_amount 
	
from biz_shop_traffic t
where t.date<=:enddate  and  t.date>=:startdate 
-- and t.src=:belong 
-- AND (bp.responsible = :resperson OR :resperson IS NULL OR :resperson = '')
and t.tertiary_source <>"汇总"
and secondary_source in :channel 
group by t.secondary_source , date
order by  'tertiary_source', date 
 

`

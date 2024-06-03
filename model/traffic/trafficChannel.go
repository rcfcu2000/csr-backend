package traffic

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

// select  DISTINCT t.source_type_2 from biz_product_traffic_stats t where t.source_type_1 ="平台流量" LIMIT 0, 100;

// select  DISTINCT t.source_type_3 from biz_product_traffic_stats t where t.source_type_1 ="广告流量" LIMIT 0, 100;

func (i *RespTrafficChannelData) GetData(r ReqTrafficAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_TRAFFIC_Channel)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("RespTrafficChannelData sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespTrafficChannelData sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespTrafficChannelData) Prodata() {

}

func (i *RespChannelsData) GetData(r ReqTrafficAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_TRAFFIC_Channels)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("SQL_TRAFFIC_Channels sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("SQL_TRAFFIC_Channels sql\n", mainSql)
		return e
	}

	return nil

}

const SQL_TRAFFIC_Channels = `

select  DISTINCT t.source_type_2 as "channel" from biz_product_traffic_stats t where t.source_type_1 ="平台流量" and  source_type_2<>"汇总" 
UNION
select  DISTINCT t.source_type_3 as "channel"  from biz_product_traffic_stats t where t.source_type_1 ="广告流量" ;

`

const SQL_TRAFFIC_Channel = `
 
select tertiary_source as tertiary_source, sum(t.visitors_count) as visitors_count, sum(t.paid_amount) as paid_amount 
from biz_shop_traffic t
where t.date<=:enddate  and  t.date>=:startdate 
-- and t.src=:belong 
-- AND (bp.responsible = :resperson OR :resperson IS NULL OR :resperson = '')
and t.tertiary_source <>"汇总"
and tertiary_source in :channel 
group by t.tertiary_source 
UNION

select secondary_source as tertiary_source, sum(t.visitors_count) as visitors_count, sum(t.paid_amount) as paid_amount 
from biz_shop_traffic t
where t.date<=:enddate  and  t.date>=:startdate 
-- and t.src=:belong 
-- AND (bp.responsible = :resperson OR :resperson IS NULL OR :resperson = '')
and t.tertiary_source <>"汇总"
and secondary_source in :channel 
group by t.secondary_source 
order by  'tertiary_source' 



`

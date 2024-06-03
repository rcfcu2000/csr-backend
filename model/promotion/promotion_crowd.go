package promotion

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespCrowdSpend) GetData(r ReqPromotionAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_PROMOTION_CROWD)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()
	i.Records = []CrowdSpendNode{}

	fmt.Println("RespCrowdSpend sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespCrowdSpend sql\n", mainSql)
		return e
	}
	if len(i.Records) == 0 {
		fmt.Println("RespCrowdSpend sql\n", mainSql)
		i.Records = []CrowdSpendNode{}
		return nil
	}
	i.ProData()
	return nil

}

func (i *RespCrowdSpend) ProData() {
	sum := 0.0
	for m := 0; m < len(i.Records); m++ {
		sum += i.Records[m].Spend
	}
	for m := 0; m < len(i.Records); m++ {
		if sum != 0 {
			i.Records[m].SpendPercentage = i.Records[m].Spend / sum
		}
	}
}

const SQL_PROMOTION_CROWD = `
	
	select 
	w.crowd_type as "crowd",
	sum(spend) as "spend",
	sum(gmv) as "gmv",
 	IFNULL( (sum(gmv) -sum(spend))/sum(spend),0.0) as "roi" 
	from v_wxa_bp_bpc w
	WHERE  product_id in ( select product_id from biz_product_classes where statistic_date =:enddate and  pallet IN :subpallet   )
	AND w.datetimekey >= :startdate  
	AND w.datetimekey <= :enddate  
	AND responsible IN  :resperson  
	and  w.promotion_name in :scene    
	GROUP BY w.crowd_type
	order by gmv desc LIMIT 20

	`

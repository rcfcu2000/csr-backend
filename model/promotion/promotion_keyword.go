package promotion

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespKeywordCost) GetData(r ReqPromotionAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_PROMOTION_KEYWORD)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()
	i.Records = []KeywordCostNode{}

	fmt.Println("RespKeywordCost sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespKeywordCost sql\n", mainSql)
		return e
	}
	if len(i.Records) == 0 {
		fmt.Println("RespKeywordCost sql\n", mainSql)
		i.Records = []KeywordCostNode{}
		return nil
	}
	i.ProData()
	return nil

}

func (i *RespKeywordCost) ProData() {
	sum := 0.0
	for m := 0; m < len(i.Records); m++ {
		sum += i.Records[m].Cost
	}
	for m := 0; m < len(i.Records); m++ {
		if sum != 0 {
			i.Records[m].CostPercentage = i.Records[m].Cost / sum
		}
	}
}

const SQL_PROMOTION_KEYWORD = `

select 
w.keyword_name as "keyword",
sum(spend) as "cost",
sum(gmv) as "gmv",
(sum(gmv) -sum(spend))/sum(spend) as "roi" 
from v_wxk_bp_bpc w
WHERE  product_id in ( select product_id from biz_product_classes where statistic_date =:enddate and  pallet IN :subpallet   )
AND w.datetimekey >= :startdate  
AND w.datetimekey <= :enddate  
AND responsible IN  :resperson  
and  w.promotion_name in :scene         
GROUP BY w.keyword_name
order by gmv desc LIMIT 20
 
`

package promotion

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespPalletCost) GetData(r ReqPromotionAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_PROMOTION_PALLETCOST)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()
	i.Records = []PalletCostNode{}

	fmt.Println("RespPalletCost sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespPalletCost sql\n", mainSql)
		return e
	}
	i.ProData()
	return nil

}

func (i *RespPalletCost) GetData4ProductList(r ReqPromotionAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_PROMOTION_PALLETCOST2)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()
	i.Records = []PalletCostNode{}

	fmt.Println("RespPalletCost sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespPalletCost sql\n", mainSql)
		return e
	}
	if len(i.Records) == 0 {
		fmt.Println("RespPalletCost sql\n", mainSql)
		i.Records = []PalletCostNode{}
		return nil
	}
	i.ProData()
	return nil

}

func (i *RespPalletCost) ProData() {
	sumcost := 0.0
	sumgmv := 0.0
	for m := 0; m < len(i.Records); m++ {
		sumcost += i.Records[m].Cost
		sumgmv += i.Records[m].GMV
	}
	for m := 0; m < len(i.Records); m++ {
		if sumcost != 0 {
			i.Records[m].CostRate = i.Records[m].Cost / sumcost
		}
		if sumgmv != 0 {
			i.Records[m].GMVRate = i.Records[m].GMV / sumgmv
		}
	}
}

/*
case
when denominator > 0 then numerator / denominator
else 0 -- 或者任何你想在除数为零时返回的值，比如NULL或者预设的一个常数
end as result

(sum(gmv) -sum(spend))/sum(spend) as "roi"
*/

const SQL_PROMOTION_PALLETCOST = `
select 
t.pallet as "pallet",
sum(spend) as "cost",
sum(gmv) as "gmv",
case 
	when sum(spend) >0 then sum(gmv)/sum(spend)
	else 0 
	end  as "roi"
 
from wanxiang_product w
left join 
( select product_id,pallet from biz_product_classes where statistic_date =:enddate and  pallet IN :subpallet   ) t 
on w.product_id = t.product_id
left join biz_product bp on w.product_id = bp.product_id
 
where w.datetimekey >= :startdate  
AND w.datetimekey <= :enddate  
AND bp.responsible IN :resperson  
and  w.promotion_type in :scene      
GROUP BY t.pallet
 
`

// 区别， 给产品清单用 ; :pallet
const SQL_PROMOTION_PALLETCOST2 = `
select 
w.pallet as "pallet",
sum(spend) as "cost",
sum(gmv) as "gmv",
COUNT(DISTINCT product_id) as "product_num",
case 
	when sum(spend) >0 then (sum(gmv) -sum(spend))/sum(spend)
	else 0 
	end  as "roi"
 
from v_wxp_bp_bpc w
WHERE  product_id in ( select product_id from biz_product_classes where statistic_date =:enddate and  pallet IN :pallet   )
AND w.datetimekey >= :startdate  
AND w.datetimekey <= :enddate  
AND responsible IN :resperson  
and  w.promotion_type in :scene      
GROUP BY w.pallet

`

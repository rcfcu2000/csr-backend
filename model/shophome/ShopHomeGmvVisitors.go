package shophome

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespShopHomeGmvVisitorsData) GetData(r ReqShopHomeAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_SHOPHOME_GmvVisitor)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("SQL_SHOPHOME_GmvVisitor sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("SQL_SHOPHOME_GmvVisitor sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespShopHomeGmvVisitorsData) Prodata() {

	for _, item := range i.Records {
		i.Sum.Gmv += item.Gmv
		i.Sum.TargetGmv += item.TargetGmv
		i.Sum.TargetDayRate = item.TargetDayRate
	}
	if i.Sum.TargetGmv != 0 {
		i.Sum.TargetGmvRate = i.Sum.Gmv / i.Sum.TargetGmv
	}

}

const SQL_SHOPHOME_GmvVisitor = `
WITH t_id_pallet_gmv AS (
	SELECT
		sum(bpp.paid_amount) AS gmv,
		max(
			CASE
			WHEN bpp.statistic_date = :enddate     THEN
				bpp.pallet
			ELSE
				'-'
			END
		) AS pallet
	FROM
		biz_pallet_product bpp
	WHERE
		bpp.statistic_date >= :startdate    
	AND bpp.statistic_date <= :enddate    
	GROUP BY
		pallet
),
 t_pallet_target_gmv AS (
	SELECT
		pallet,
		sum(gmv_target) AS gmv_target
	FROM
		biz_shop_pallet_targets bspt
	WHERE
		bspt.statistic_date >= :startdate   
	AND bspt.statistic_date <= LAST_DAY(:enddate    )
	GROUP BY
		pallet
) 
 

SELECT
	tmain.pallet,
	gmv,
	gmv_target as target_gmv,
  (gmv / IFNULL(gmv_target, 0)) AS target_gmv_rate,
  ( 1-DATEDIFF(LAST_DAY(:enddate    ),:enddate    )/(DATEDIFF(:enddate    ,:startdate     )+1+DATEDIFF(LAST_DAY(:enddate    ),:enddate    )) ) as  target_day_rate
 
FROM
	t_id_pallet_gmv tmain
LEFT JOIN t_pallet_target_gmv bspt ON bspt.pallet = tmain.pallet 

`

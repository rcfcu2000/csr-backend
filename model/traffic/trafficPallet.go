package traffic

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespPalletData) GetData(r ReqTrafficAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_TRAFFIC_Pallet)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("RespPalletData sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("RespPalletData sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespPalletData) Prodata() {

}

const SQL_TRAFFIC_Pallet = `
 
WITH t_product_pallet AS (
	SELECT
	 bpp.product_id,
	max(case when bpp.statistic_date = :enddate    then bpp.pallet else '-' end) as pallet 

	FROM
		biz_pallet_product bpp
		where bpp.statistic_date >= :startdate   and bpp.statistic_date<=:enddate 
		-- AND (bpp.responsible = :resperson  OR :resperson  IS NULL OR :resperson  = '')
	GROUP BY
	bpp.product_id
)
select pallet, count( product_id ) as count from t_product_pallet GROUP BY pallet 
ORDER BY FIELD(pallet, 'S','A','B','C','D','-');



`

package shophome

import (
	"fmt"
	"xtt/global"
	"xtt/model/common"

	"go.uber.org/zap"
)

func (i *RespShopHomeExperienceScoreData) GetData(r ReqShopHomeAllSearch) error {
	sqlp := &common.SQLProccesor{}
	sqlp.SetSql(SQL_SHOPHOME_ExperienceScore)
	(&r).SetToSQLProccesor(sqlp)
	mainSql := sqlp.GetResult()

	fmt.Println("SQL_SHOPHOME_ExperienceScore sql\n", mainSql)
	e := global.GVA_DB.Raw(mainSql).Scan(&i.Records).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", mainSql))
		fmt.Println("SQL_SHOPHOME_ExperienceScore sql\n", mainSql)
		return e
	}
	i.Prodata()
	return nil

}
func (i *RespShopHomeExperienceScoreData) Prodata() {

}

const SQL_SHOPHOME_ExperienceScore = `

select 
 DATE_FORMAT(statistic_date, '%m-%d')  as date,
 overall_experience_score,
product_experience_score,
logistics_experience_score,
service_experience_score
 from  biz_shop_experience_score t
where  t.statistic_date >= :startdate 
AND t.statistic_date <= :enddate 
`

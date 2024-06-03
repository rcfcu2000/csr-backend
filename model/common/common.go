package common

import "fmt"

func ToInStringInt(i []int64) string {
	if len(i) == 0 {
		return "(-1,1,0,100)"
	}
	str := ""
	for _, item := range i {
		str += fmt.Sprintf("%d,", item)
		if item == 999 {
			return "(-1,1,0,100)"
		}
	}
	str = str[0 : len(str)-1]
	str = fmt.Sprintf("( %s )", str)
	return str
}

// []string  =>  （"id1","id2"）  or  ( select id from xxx )
func ListToInStringStr(i []string, idtype string) string {
	if len(i) == 0 {
		if idtype == "" {
			return "('S','A','B','C','D','-')"
		}
		if idtype == "pallet" {
			return "('S','A','B','C','D','-')"
		}
		if idtype == "resperson" {
			return " (select DISTINCT responsible  from biz_product  )  "
		}

		if idtype == "scene" {
			return `( "场景推广", "关键词推广", "精准人群推广" )`
		}
		if idtype == "keyword" {
			return `( select DISTINCT keyword_name from wanxiang_keywords )`
		}
		if idtype == "crowd" {
			return `( select DISTINCT crowd_type from wanxiang_audience )`
		}

		if idtype == "channel" {
			return `( select  DISTINCT t.source_type_2 as "channel" from biz_product_traffic_stats t where t.source_type_1 ="平台流量" and  source_type_2<>"汇总" 
			UNION
			select  DISTINCT t.source_type_3 as "channel"  from biz_product_traffic_stats t where t.source_type_1 ="广告流量" 
			 )`
		}

		return "('')"
	}
	str := ""
	for _, item := range i {
		str += fmt.Sprintf("'%s',", item)

	}
	str = str[0 : len(str)-1]
	str = fmt.Sprintf("( %s )", str)
	return str
}

func PalletToInStringStr(i []string) string {
	if len(i) == 0 {
		return "('S','A','B','C','D','-')"
	}
	str := ""
	for _, item := range i {
		str += fmt.Sprintf("'%s',", item)

	}
	str = str[0 : len(str)-1]
	str = fmt.Sprintf("( %s )", str)
	return str
}

func PidToInStringStr(i []string) string {
	if len(i) == 0 {
		return " (select DISTINCT product_id  from biz_product  ) "
	}
	str := ""
	for _, item := range i {
		str += fmt.Sprintf("'%s',", item)

	}
	str = str[0 : len(str)-1]
	str = fmt.Sprintf("( %s )", str)
	return str
}

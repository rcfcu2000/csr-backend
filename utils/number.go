package utils

import (
	"fmt"
	"math"
	"strconv"
)

// Round
/**
 * @Description 用于对浮点数进行四舍五入，并保留指定的小数位数
 * @Date 2023-05-17 16:42:35
 */
func Round(value float64, decimalPlaces int) float64 {
	decimalPow := math.Pow10(decimalPlaces)
	rounded := math.Round(value*decimalPow) / decimalPow
	rounded, _ = strconv.ParseFloat(strconv.FormatFloat(rounded, 'f', decimalPlaces, 64), 64)
	return rounded
}

// TruncateFloat
/**
* @Description 指定小数位截断
* @Date 2023-06-28
*/
func TruncateFloat(num float64, decimalPlaces int) float64 {
	shift := math.Pow(10, float64(decimalPlaces))
	return math.Floor(num*shift) / shift
}

// Float2
/**
* @Description 浮点数忽略精度丢失的小数位
* @Date 2023-06-15
*/
func Float2(num float64) float64{
		fNum, _ := strconv.ParseFloat(fmt.Sprintf("%.8f", num), 64)
		return fNum
}

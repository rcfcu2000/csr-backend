package utils

import (
	"fmt"
	"time"

	"io/ioutil"

	"github.com/bwmarrin/snowflake"
)

func init() {
	__n, _ = snowflake.NewNode(time.Now().UnixMilli() % 1024)
}

// 返回指定日期范围内的所有日期组成的字符串数组
func GetDatesInRange(startDateStr string, endDateStr string) ([]string, error) {
	// 解析输入的日期字符串
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return nil, fmt.Errorf("解析开始日期时出错: %v", err)
	}
	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return nil, fmt.Errorf("解析结束日期时出错: %v", err)
	}

	// 确保开始日期不晚于结束日期
	if startDate.After(endDate) {
		return nil, fmt.Errorf("开始日期晚于结束日期")
	}

	// 创建一个空的日期字符串切片来存储结果
	var dates []string

	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		// 将时间格式化为字符串并添加到结果切片中
		dateStr := d.Format("2006-01-02")
		dates = append(dates, dateStr)
	}

	return dates, nil
}

var __n *snowflake.Node

func init() {
	__n, _ = snowflake.NewNode(time.Now().UnixMilli() % 1024)
}

func GenId() snowflake.ID {
	id := __n.Generate()
	return id
	// id.String()
	// return fmt.Sprintf("%d", id.Int64()), id
}

func CopyFile(sourceFile, destinationFile string) error {
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {
		fmt.Println("Error creating", destinationFile)
		fmt.Println(err)
		return err
	}
	return nil
}

// 范围所在月的天数  time.February  GetDaysInMonth(2024, time.February) GetDaysInMonth(2024, 2)
func GetDaysInMonth(year int, month time.Month) int {
	return time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC).Day() - 1
}

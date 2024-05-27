package main

import (
	"your_project/controllers"
	"your_project/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto Migrate the schema
	db.AutoMigrate(&models.BizQa{}, &models.BizQaType{}, &models.BizQuestionType{}, &models.BizQaQuestion{}, &models.BizMerchant{}, &models.BizMerchantLink{}, &models.BizMTag{}, &models.BizMerchantTag{}, &models.BizMerchantParameters{}, &models.BizMerchantKeypoints{})

	r := gin.Default()

	// Initialize controllers
	bizQaController := controllers.NewBizQaController(db)

	// Define routes
	r.POST("/biz_qa", bizQaController.CreateBizQa)
	r.GET("/biz_qa/:id", bizQaController.GetBizQa)
	r.PUT("/biz_qa/:id", bizQaController.UpdateBizQa)
	r.DELETE("/biz_qa/:id", bizQaController.DeleteBizQa)

	// Run the server
	r.Run(":8080")
}

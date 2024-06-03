package main

import (
	"csr-backend/controllers"
	"csr-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	docs "csr-backend/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var db *gorm.DB
var err error

func main() {
	dsn := "root:pwd123OK@tcp(47.109.94.69:3306)/csr?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db = db.Debug()

	if err != nil {
		panic("failed to connect database")
	}

	// Auto Migrate the schema
	db.AutoMigrate(&models.BizQa{}, &models.BizQaType{}, &models.BizMessage{})

	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/"

	// Initialize controllers
	bizQaController := controllers.NewBizQaController(db)
	bizQaTypeController := controllers.NewBizQaTypeController(db)
	bizMessageController := controllers.NewBizMessageController(db)

	// Register routes
	api := r.Group("/api/v1")
	registerRoutes(api, bizQaController, bizQaTypeController, bizMessageController)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Run the server
	r.Run(":8080")
}

func registerRoutes(r *gin.RouterGroup, bizQaController *controllers.BizQaController, bizQaTypeController *controllers.BizQaTypeController, bizMessageController *controllers.BizMessageController) {

	r.POST("/biz_qa_complex", bizQaController.CreateBizQaComplex)
	r.GET("/biz_qa/:id", bizQaController.GetBizQa)
	r.PUT("/biz_qa/:id", bizQaController.UpdateBizQa)
	r.DELETE("/biz_qa/:id", bizQaController.DeleteBizQa)

	// BizQaType routes
	r.POST("/biz_qa_type", bizQaTypeController.CreateBizQaType)
	r.GET("/biz_qa_type/:id", bizQaTypeController.GetBizQaType)
	r.GET("/biz_qa_types", bizQaTypeController.GetAllBizQaTypes)
	r.PUT("/biz_qa_type/:id", bizQaTypeController.UpdateBizQaType)
	r.DELETE("/biz_qa_type/:id", bizQaTypeController.DeleteBizQaType)

	// BizMessage routes
	r.POST("/biz_messages", bizMessageController.CreateBizMessages)
}

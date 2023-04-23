package main

import (
	"LearnGo/component/appctx"
	"LearnGo/component/uploadprovider"
	"LearnGo/middleware"
	"LearnGo/module/restaurant/transport/ginrestaurant"
	"LearnGo/module/upload/uploadtransport/ginupload"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dsn := os.Getenv("MYSQL_STR")
	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	secretKey := os.Getenv("SYSTEM_SECRET")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connect DB failed", err)
	}
	log.Println("Connect DB success", db)

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)
	appContext := appctx.NewAppContext(db, s3Provider, secretKey)
	db = db.Debug()

	r := gin.Default()

	r.Use(middleware.Recover(appContext))

	v1 := r.Group("v1")

	v1.POST("upload", ginupload.Upload(appContext))

	restaurants := v1.Group("restaurants")

	restaurants.POST("/", ginrestaurant.CreateRestaurant(appContext))

	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))

	restaurants.GET("/", ginrestaurant.ListRestaurant(appContext))

	restaurants.GET("/:id", ginrestaurant.GetRestaurant(appContext))

	restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appContext))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

package main

import (
	"LearnGo/component/appctx"
	"LearnGo/component/uploadprovider"
	"LearnGo/middleware"
	localPb "LearnGo/pubsub/localpub"
	"LearnGo/route/admin"
	"LearnGo/route/client"
	"LearnGo/route/user"
	"LearnGo/subscriber"
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
	db = db.Debug()

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)
	pubsub := localPb.NewPubSub()
	appContext := appctx.NewAppContext(db, s3Provider, secretKey, pubsub)

	// use this line as an alternative for Setup
	if err := subscriber.NewEngine(appContext).Start(); err != nil {
		log.Fatalln()
	}

	////setup subscriber
	//subscriber.Setup(appContext)

	r := gin.Default()

	r.Use(middleware.Recover(appContext))

	// route
	v1 := r.Group("v1")
	admin.AdminRoute(appContext, v1)
	client.ClientRoute(appContext, v1)
	user.UserRoute(appContext, v1)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

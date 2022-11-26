package main

import (
	"context"
	"log"
	"os"

	"bds/component/appctx"
	upldprv "bds/component/uploadprovider"
	"bds/middleware"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("MYSQL_CONN_STRING")
	secretKey := os.Getenv("SYSTEM_SECRET")
	cloudStorageProjectID := os.Getenv("CLOUD_STORAGE_PROJECT_ID")
	cloudStorageBucketName := os.Getenv("CLOUD_STORAGE_BUCKET_NAME")
	cloudStorageDomain := os.Getenv("CLOUD_STORAGE_DOMAIN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	cloudStorageClient, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	cloudStorageProvider := upldprv.NewCloudStorageProvider(
		cloudStorageClient,
		cloudStorageProjectID,
		cloudStorageBucketName,
		cloudStorageDomain,
	)

	appContext := appctx.NewAppContext(db, cloudStorageProvider, secretKey)

	r := gin.Default()
	r.Use(middleware.Recover(appContext))

	v1 := r.Group("/v1")

	setupAdminRoute(appContext, v1)

	if err = r.Run(); err != nil {
		return
	}
}

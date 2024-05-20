package main

import (
	"flowerpot/models"
	"flowerpot/routers"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitlab.com/1hopin/go-module/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	var uuid string = "start-app"
	utils.SetInitLog(os.Getenv("APP_NAME"), os.Getenv("ENVIRONMENT"), uuid)
	if os.Getenv("ENVIRONMENT") != "local" && err != nil {
		utils.LoggerError("App init", "Error loading .env file", nil, nil, err, false)
	}
}

func main() {
	db, err := connectDatabase()

	if err != nil {
		utils.LoggerError("ConnectDatabase", "Error connect database", nil, nil, err, false)
		log.Fatalf("Error connect database: %v", err)
	}

	//Start Setup Routers
	router := gin.New()
	router.Use(gin.Recovery())

	port := os.Getenv("SERVER_PORT")

	routers.SetupRouter(router, db)
	utils.LoggerInfo("Healthcheck", "Start "+os.Getenv("APP_NAME")+"("+os.Getenv("ENVIRONMENT")+") on:"+port, nil, nil, os.Getenv("ENVIRONMENT") != "local")
	log.Fatal(router.Run(":" + port))
}

func connectDatabase() (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt: true,
		})
	if err != nil {
		return nil, err
	}

	// AutoMigrate the schema
	db.AutoMigrate(
		&models.Member{},
	)

	return db, nil
}

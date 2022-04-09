package config

import (
	"fmt"
	"os"

	"github.com/allanurbayramgeldiyev209/learngin/helpers"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDbConn() *gorm.DB {
	errEnv := godotenv.Load(".env")
	helpers.CheckErr(errEnv)

	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ashgabat", db_host, db_username, db_password, db_name, db_port)
	db, errConn := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helpers.CheckErr(errConn)

	return db
}

func CloseDbConn(db *gorm.DB) {
	dbSQL, err := db.DB()
	helpers.CheckErr(err)

	errCloseDbConn := dbSQL.Close()
	helpers.CheckErr(errCloseDbConn)
}

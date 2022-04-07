package config

import (
	"fmt"
	"os"

	"github.com/allanurbayramgeldiyev209/learngin/helpers"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
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

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",db_username,db_password,db_host,db_port,db_name)
  db, errConn := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helpers.CheckErr(errConn)

	return db
} 

func CloseDbConn( db *gorm.DB) {
	dbSQL , err := db.DB()
	helpers.CheckErr(err)

	errCloseDbConn := dbSQL.Close()
	helpers.CheckErr(errCloseDbConn)
}
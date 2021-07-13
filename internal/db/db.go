package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var db *gorp.DbMap
var db *gorm.DB

//Init ...
func Init() {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	fmt.Printf("Connecting to %s\n", dbURI)
	var err error
	db, err = gorm.Open(mysql.Open(dbURI), &gorm.Config{})

	if err != nil {
		log.Fatalf("%+v", err)
	}

}

func GetDB() *gorm.DB {
	return db
}

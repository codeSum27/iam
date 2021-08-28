package api

import (
	"fmt"
	"github.com/codeSum27/iam/pkg/common"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
	"log"
)

var gGormDB *gorm.DB

func CloseDB() {
	if gGormDB != nil {
		gGormDB.Close()
	}
}

func Init() {
	db, err := gorm.Open("mysql", getDBConfig())
	if err != nil {
		log.Println("DB Connection error ", err)
		panic(err)
	}
	gGormDB = db

	log.Println("Start DB Migration ... ")
	log.Println("Start DB ... ")

	if err := db.AutoMigrate(&User{}).Error; err != nil {
		log.Println("DB Migration error ", err)
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return gGormDB
}

func getDBConfig() (dbConnString string) {
	dbHost := common.Cnf.Database.Mariadb.Host
	dbUser := common.Cnf.Database.Mariadb.User
	dbPass := common.Cnf.Database.Mariadb.Password
	dbName := common.Cnf.Database.Mariadb.Database

	dbConnString = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbName,
	)

	return
}
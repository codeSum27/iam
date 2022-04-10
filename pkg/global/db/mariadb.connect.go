package db

import (
	"fmt"
	"os"

	"github.com/codeSum27/iam/pkg/api"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gGormDB *gorm.DB

func DBInit() {
    fmt.Println("Start to initialize Database Connection")
    db, err := gorm.Open(mysql.Open(getDBDsn()), &gorm.Config{})
    if err != nil {
        panic(err.Error())
    }

    gGormDB = db

    err = db.AutoMigrate(&api.User{})
    if err != nil {
        panic(err.Error())
    }

    fmt.Println("Success to initialize Database Connection")
}

func GetDBClient() *gorm.DB {
	return gGormDB
}

func getDBDsn() (dsn string) {
    user := os.Getenv("DB_USER")
    pass := os.Getenv("DB_PASS")
    host := os.Getenv("DB_HOST")
    name := os.Getenv("DB_NAME")
    
    dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
    user, pass, host, name)

    return dsn
}
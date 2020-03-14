package common

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"mygin/modle"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	//dirverNmae := "mysql"
	////dirver := "jdbc"
	//host := "localhost"
	//addr := "3306"
	//database := "ginessential"
	//username := "root"
	//password := "adminwss"
	//charset := "utf8"
	//args := fmt.Sprintf("%s:%s&tcp(%s:%s)/%s?charset=%s&parseTime=ture",
	//	//dirver,
	//	username,
	//	password,
	//	host,
	//	addr,
	//	database,
	//	charset)

	db, err := gorm.Open("mysql", "root:adminwss@/ginessential?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database ,err: " + err.Error())
	}

	db.AutoMigrate(&modle.User{})

	return db
}

func GetDBEngine() *gorm.DB {
	return DB
}

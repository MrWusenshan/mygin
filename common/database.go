package common

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"mygin/modle"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	/*dirverNmae := viper.GetString("datasource.dirverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=ture",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(dirverNmae, args)

	//args:=fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=ture&loc=Local",username,password,database,charset)
	//db,err:=gorm.Open(dirverNmae,args) */
	db, err := gorm.Open("mysql", "root:adminwss@/ginessential?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database ,err: " + err.Error())
	}

	db.AutoMigrate(&modle.User{})
	DB = db
	return db
}

func GetDBEngine() *gorm.DB {
	return DB
}

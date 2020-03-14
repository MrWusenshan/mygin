package orm

import (
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(110);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

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

	db.AutoMigrate(&User{})

	return db
}

package configs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB() *gorm.DB {
	//open a db connection
	db, err := gorm.Open("mysql", "root:@/dragon?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}
	return db
}

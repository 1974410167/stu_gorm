package init_db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB


func InitDB(){
	dsn := "root:wsghy.5637@tcp(gz-cynosdbmysql-grp-0ku3m9v9.sql.tencentcdb.com:29978)/database1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatalln(err)
	}
	DB = db
}

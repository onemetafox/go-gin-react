package database

import (
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
type Database struct{
	*gorm.DB
}

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB () (*gorm.DB)  {
	var err error
	dsn := os.Getenv("Mysql_user") +":"+ os.Getenv("Mysql_password") +"@tcp"+ "(" + os.Getenv("Mysql_host") + ":3306)/go-gin-react?" + "parseTime=true&loc=Local"
   
   	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

   	if err != nil {
		panic(err)
   	}

   	return db
}
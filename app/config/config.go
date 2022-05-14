package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const JwtSecret = "12345"
const BaseURLRO = "https://api.rajaongkir.com/starter/"
const Key = "677971ea001e7e7d868fc03c52412452"

func InitDB() (DB *gorm.DB) {
	username := "root"
	password := ""
	host := "127.0.0.1"
	port := "3306"
	dbname := "latihan-db-2"
	var err error

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return
}

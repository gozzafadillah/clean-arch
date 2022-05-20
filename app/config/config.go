package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DBNAME string
	DBUSER string
	DBPASS string
	DBHOST string
	DBPORT string

	JWTSecret string
}

const BaseURLRO = "https://api.rajaongkir.com/starter/"
const Key = "677971ea001e7e7d868fc03c52412452"

var Conf Config

func InitDB() (DB *gorm.DB) {
	Conf = Config{
		// DBNAME:    os.Getenv("DBNAME"),
		// DBUSER:    os.Getenv("DBUSER"),
		// DBPASS:    os.Getenv("DBPASS"),
		// DBHOST:    os.Getenv("DBHOST"),
		// DBPORT:    os.Getenv("DBPORT"),
		DBUSER:    "root",
		DBPASS:    "",
		DBHOST:    "host.docker.internal",
		DBPORT:    "3306",
		DBNAME:    "latihan-db-2",
		JWTSecret: "12345",
	}

	var err error

	dsn := Conf.DBUSER + ":" + Conf.DBPASS + "@tcp(" + Conf.DBHOST + ":" + Conf.DBPORT + ")/" + Conf.DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return
}

package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT      = 3000
	SECRETKEY []byte
	CERTFILE  = ""
	KEYFILE   = ""
	DBDRIVER  = ""
	DBURL     = ""
)

func Load() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Println(err)
	}
	DBDRIVER = os.Getenv("DB_DRIVER")
	SECRETKEY = []byte(os.Getenv("API_SECRET"))
	KEYFILE = os.Getenv("KEYFILE")
	CERTFILE = os.Getenv("CERTFILE")
	DBDRIVER = os.Getenv("DB_DRIVER")
	DBURL = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	//db, err := gorm.Open("mysql", "user:password@dbname?charset=utf8&parseTime=True&loc=Local")

}

package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tawfeeq0/go_ssl_server/server/config"
)

func Connect() (*gorm.DB,error){
	db, err := gorm.Open(config.DBDRIVER,config.DBURL)
	if err != nil {
		return nil,err
	}
	return db,nil
}
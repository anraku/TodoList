package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	DB *gorm.DB
}

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:password@/app?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	return db, nil
}

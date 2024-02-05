package gorestapi

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	DB struct {
		orm *gorm.DB
	}
)

func NewDB(config *Config) (db *DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DbName,
	)
	orm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db = &DB{
		orm: orm,
	}
	return
}

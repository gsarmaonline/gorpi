package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	DB struct {
		Orm *gorm.DB
	}
	ResourceModel interface {
		String() string
		Ancestor() ResourceModel
	}
)

func NewDB(username, password, host, port, dbName string) (db *DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbName,
	)
	orm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db = &DB{
		Orm: orm,
	}
	return
}

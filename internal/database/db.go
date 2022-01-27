package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

type DB struct {
	ormMaster *gorm.DB
}

func New() (*DB, error) {
	ormMaster, err := gorm.Open(mysql.Open(), &gorm.Config{})
	if err != nil {
		return err
	}
	
}
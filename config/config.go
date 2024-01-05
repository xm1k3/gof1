package config

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectSqlite3(dbName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to connect database")
	}
	return db, nil

}

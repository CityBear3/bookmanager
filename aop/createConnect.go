package aop

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	user := "root"
	pass := "testPassword"
	protcol := "tcp(127.0.0.1:53306)"
	dbName := "testDB"

	dsn := user + ":" + pass + "@" + protcol + "/" + dbName
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

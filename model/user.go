package model

import (
	"app/aop"
)

type User struct {
	Id       uint   `gorm:"primary_key"`
	Uid      string `gorm:"column:uid"`
	Password string `gorm:"column:password"`
}

func (u *User) GetUserById() error {
	db, err := aop.Connect()
	if err != nil {
		return err
	}

	dbCloser, err := db.DB()
	if err != nil {
		return err
	}
	defer dbCloser.Close()

	db.Find(u)
	db.Logger.LogMode(4)

	return nil
}

func (u *User) RegisterUser() error {
	db, err := aop.Connect()
	if err != nil {
		return err
	}

	dbCloser, err := db.DB()
	if err != nil {
		return err
	}
	defer dbCloser.Close()

	db.Create(u)
	return nil
}

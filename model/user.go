package model

type User struct {
	Id       uint   `gorm:"primary_key"`
	Uid      string `gorm:"column:uid"`
	Password string `gorm:"column:password"`
}

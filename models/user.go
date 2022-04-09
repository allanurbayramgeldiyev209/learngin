package models

import (
	"github.com/allanurbayramgeldiyev209/learngin/config"
	"github.com/allanurbayramgeldiyev209/learngin/helpers"
)

type User struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"index;type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Token    string `gorm:"-" json:"token"`
}

func (u *User) Add() {

	db := config.SetupDbConn()

	err := db.Create(&u).Error
	helpers.CheckErr(err)

	config.CloseDbConn(db)

}

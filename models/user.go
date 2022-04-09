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
	Token    string `gorm:"text" json:"token"`
}

func (u User) Add() {

	db := config.SetupDbConn()
	// config.CloseDbConn(db)

	err := db.Create(&u).Error
	helpers.CheckErr(err)

}

func (u User) GetUser(email string) (error, User) {

	db := config.SetupDbConn()
	// config.CloseDbConn(db)

	err := db.First(&u, "email = ?", email).Error
	if err != nil {
		return err, User{}
	}
	return nil, u

}

func (u User) UpdateToken(email, token string) {

	db := config.SetupDbConn()

	err := db.Model(&u).Where("email = ?", email).Update("token", token).Error
	helpers.CheckErr(err)

}

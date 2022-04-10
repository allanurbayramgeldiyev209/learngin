package models

import (
	"github.com/allanurbayramgeldiyev209/learngin/config"
	"github.com/allanurbayramgeldiyev209/learngin/helpers"
)

type User struct {
	ID            uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string `gorm:"type:varchar(255)" json:"name"`
	Email         string `gorm:"index;type:varchar(255)" json:"email"`
	Password      string `gorm:"->;<-;not null" json:"-"`
	AccessToken   string `gorm:"text" json:"access_token"`
	ResfreshToken string `gorm:"text" json:"refresh_token"`
}

func (u User) Add() {

	db := config.SetupDbConn()

	err := db.Create(&u).Error
	helpers.CheckErr(err)

}

func (u User) GetUser(email string) (error, User) {

	db := config.SetupDbConn()

	err := db.First(&u, "email = ?", email).Error
	if err != nil {
		return err, User{}
	}
	return nil, u

}

func (u User) UpdateToken(email, access_token, refresh_token string) {

	db := config.SetupDbConn()

	err := db.Model(u).Where("email = ?", email).Updates(User{AccessToken: access_token, ResfreshToken: refresh_token}).Error
	helpers.CheckErr(err)

}

func (u User) UpdateTokenByID(id uint, access_token, refresh_token string) {

	db := config.SetupDbConn()

	err := db.Model(u).Where("id = ?", id).Updates(User{AccessToken: access_token, ResfreshToken: refresh_token}).Error
	helpers.CheckErr(err)

}

func (u User) GetUserByID(id uint) (error, User) {

	db := config.SetupDbConn()

	err := db.First(&u, "id = ?", id).Error
	if err != nil {
		return err, User{}
	}
	return nil, u

}

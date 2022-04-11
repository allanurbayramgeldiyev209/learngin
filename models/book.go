package models

import (
	"github.com/allanurbayramgeldiyev209/learngin/config"
	"github.com/allanurbayramgeldiyev209/learngin/helpers"
)

type Book struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
}

func (b Book) Add() {

	db := config.SetupDbConn()

	err := db.Create(&b).Error
	helpers.CheckErr(err)

}

func (b Book) GetAll() []Book {

	books := []Book{}

	db := config.SetupDbConn()

	err := db.Find(&books).Error
	helpers.CheckErr(err)

	return books

}

func (b Book) Get(id uint) Book {
	db := config.SetupDbConn()

	db.First(&b, "id = ?", id)
	return b

}

func (b Book) Update(id uint, title, description string) {

	db := config.SetupDbConn()

	err := db.Model(b).Where("id = ?", id).Updates(Book{Title: title, Description: description}).Error
	helpers.CheckErr(err)

}

func (b Book) Delete(id uint) {

	db := config.SetupDbConn()

	err := db.Delete(&b, id).Error
	helpers.CheckErr(err)

}

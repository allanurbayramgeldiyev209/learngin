package models

type Book struct {
	ID uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Title string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	UserID uint64 `gorm:"not null" json:"user_id"`
	User User `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
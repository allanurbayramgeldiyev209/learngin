package models

type User struct {
	ID uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(255)" json:"name"`
	Email string `gorm:"index;type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Token string `gorm:"-" json:"token"`
}
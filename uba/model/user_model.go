package model

type User struct {
	Id       int    `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(300);not null"`
	LastName string `gorm:"type:varchar(300);not null"`
	Email    string `gorm:"type:varchar(500);not null;unique"`
	Password string `gorm:"type:varchar(200);not null"`
	UserType bool   `gorm:"type:boolean;not null"`
	Dni      int    `gorm:"type int;not null"`
}

type Users []User

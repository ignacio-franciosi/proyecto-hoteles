package model

type Booking struct {
	Id        int    `gorm:"primaryKey"`
	StartDate string `gorm:"type:date;not null"`
	EndDate   string `gorm:"type:date;not null"`
	IdMongo   string `gorm:"type:varchar(250);not null"`
	IdUser    int    `gorm:"type:integer;not null"`
}

type Bookings []Booking

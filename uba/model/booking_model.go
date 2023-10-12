package model

import "time"

type Booking struct {
	Id         int       `gorm:"primaryKey"`
	StartDate  time.Time `gorm:"type:date;not null"`
	EndDate    time.Time `gorm:"type:date;not null"`
	TotalPrice float32   `gorm:"type:decimal;unsigned;not null"`
	IdHotel    int       `gorm:"type:integer;not null"`
	IdUser     int       `gorm:"type:integer;not null"`
}

type Bookings []Booking

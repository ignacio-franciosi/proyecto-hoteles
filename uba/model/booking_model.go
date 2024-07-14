package model

type Booking struct {
	IdBooking  int     `gorm:"primaryKey"`
	StartDate  string  `gorm:"type:varchar(16);not null"`
	EndDate    string  `gorm:"type:varchar(16);not null"`
	IdMongo    string  `gorm:"type:varchar(250);not null"`
	IdUser     int     `gorm:"type:integer;not null"`
	TotalPrice float64 `gorm:"type:decimal(10,2); not null"`
}

type Bookings []Booking

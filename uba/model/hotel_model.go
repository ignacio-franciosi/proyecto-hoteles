package model

type Hotel struct {
	Id        int     `gorm:"primaryKey"`
	IdMongo   string  `gorm:"type:varchar(250);not null;unique"`
	IdAmadeus string  `gorm:"type:varchar(250);not null;unique"`
	Price     float32 `gorm:"type:decimal;unsigned;not null"`
}

type Hotels []Hotel

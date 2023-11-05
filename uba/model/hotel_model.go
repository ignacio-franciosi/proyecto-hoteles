package model

type Hotel struct {
	Id        int    `gorm:"primaryKey"`
	IdMongo   string `gorm:"type:varchar(250);not null;unique"`
	IdAmadeus string `gorm:"type:varchar(250);not null;unique"`
}

type Hotels []Hotel

package models

type Hotel struct {
	Id int `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
	Floor int `gorm:"not null;" json:"floor"`
	District int `gorm:"not null;" json:"district"`
	Buildin int `gorm:"not null;" json:"buildin"`
	Image string `gorm:"size:255;not null" json:"image"`
}

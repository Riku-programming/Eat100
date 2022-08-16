package entity

import "gorm.io/gorm"

type ShopDetail struct {
	gorm.Model

	ShopName    string
	Reservable  bool
	Address     string
	Time        string
	Payment     string
	PhoneNumber string
	Cost        string
	URL         string
}

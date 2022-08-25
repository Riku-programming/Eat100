package entity

import "gorm.io/gorm"

type ShopDetail struct {
	gorm.Model

	Category    string
	ShopName    string
	Reservable  bool
	Address     string
	Time        string
	Payment     string
	PhoneNumber string
	Cost        int
	URL         string
}

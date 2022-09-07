package entity

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model

	Category    string
	ShopName    string
	Reservable  bool
	Address     string
	Time        string
	Payment     bool
	PhoneNumber string
	Cost        int
	URL         string
}
